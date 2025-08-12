package rdbms

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/benizzio/open-asset-allocator/infra"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
	"github.com/lib/pq"
)

// =================================================
// RDBMS ADAPTER - Implementation
// =================================================

type Adapter struct {
	config         infra.RDBMSConfiguration
	connectionPool *sql.DB
	dbx            *dbx.DB
}

// ------------------------------------------------
// SETUP FUNCTIONS
// ------------------------------------------------

func (adapter *Adapter) openPool() {

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Opening connection to %s at %s", adapter.config.DriverName, adapter.config.RdbmsURL)

	var connectionError error
	adapter.connectionPool, connectionError = sql.Open(adapter.config.DriverName, adapter.config.RdbmsURL)
	if connectionError != nil {
		glog.Fatal("Error opening database connection: ", connectionError)
		return
	}
}

func (adapter *Adapter) configPool() {
	adapter.connectionPool.SetMaxOpenConns(5)
	adapter.connectionPool.SetMaxIdleConns(5)
	adapter.connectionPool.SetConnMaxLifetime(20 * time.Minute)
	adapter.connectionPool.SetConnMaxIdleTime(5 * time.Minute)
}

func (adapter *Adapter) buildDBX() {
	adapter.dbx = dbx.NewFromDB(adapter.connectionPool, adapter.config.DriverName)
}

func (adapter *Adapter) Init() {
	glog.Infof("Opening connection pool")
	adapter.openPool()
	glog.Infof("Configuring connection pool")
	adapter.configPool()
	glog.Infof("Configuring ozzo-dbx enhancer")
	adapter.buildDBX()
}

func (adapter *Adapter) Stop() {
	err := adapter.connectionPool.Close()
	if err != nil {
		glog.Fatal("Error closing database connection: ", err)
		return
	}
}

func (adapter *Adapter) Ping() {

	glog.Info("Pinging database to test connection")

	pingContext, cancel := buildPingContext()
	defer cancel()

	err := adapter.connectionPool.PingContext(pingContext)
	if err != nil {
		glog.Fatal("Error pinging database connection: ", err)
		return
	}

	glog.Info("Ping successful!")
}

func buildPingContext() (context.Context, context.CancelFunc) {
	pingContext, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go func() {
		<-pingContext.Done()
		if errors.Is(pingContext.Err(), context.DeadlineExceeded) {
			glog.Fatal("Error pinging database connection: timeout")
		}
	}()
	return pingContext, cancel
}

// ------------------------------------------------
// USAGE FUNCTIONS
// ------------------------------------------------

func (adapter *Adapter) BuildQuery(sql string) *QueryBuilder {
	return &QueryBuilder{
		dbx:          adapter.dbx,
		querySQL:     sql,
		params:       dbx.Params{},
		whereClauses: make([]string, 0),
	}
}

func BuildQueryInTransaction[T any](
	transContext *SQLTransactionalContext,
	sql string,
) *SQLTransactionalQueryBuilder[T] {
	return &SQLTransactionalQueryBuilder[T]{
		transaction:  transContext.GetTransaction(),
		querySQL:     sql,
		params:       make([]any, 0),
		whereClauses: make([]string, 0),
	}
}

func (adapter *Adapter) buildTransactionalContext() (*SQLTransactionalContext, error) {
	var transactionContext, err = withTransaction(adapter.connectionPool)
	if err != nil {
		return nil, err
	}
	return transactionContext, nil
}

func (adapter *Adapter) RunInTransaction(
	transactionalFunction func(transContext *SQLTransactionalContext) error,
) error {

	transContext, err := adapter.buildTransactionalContext()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			glog.Errorf("Recovered from panic during transactional operation: %v", r)
			var transaction = transContext.GetTransaction()
			if rollbackErr := transaction.Rollback(); rollbackErr != nil {
				glog.Errorf("Transaction rollback failed: %v", rollbackErr)
			}
		}
	}()

	return adapter.runInTransaction(transContext, transactionalFunction)
}

func (adapter *Adapter) runInTransaction(
	transContext *SQLTransactionalContext,
	transactionalFunction func(transContext *SQLTransactionalContext) error,
) error {

	err := transactionalFunction(transContext)

	var transaction = transContext.GetTransaction()
	if err != nil {
		if rollbackErr := transaction.Rollback(); rollbackErr != nil {
			return infra.BuildAppError(
				"transaction rollback failed: "+rollbackErr.Error()+"; original error: "+err.Error(),
				adapter,
			)
		}
		return err
	}

	return transaction.Commit()
}

func (adapter *Adapter) Insert(model interface{}) error {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Inserting model %T", model)
	return adapter.dbx.Model(model).Insert()
}

func (adapter *Adapter) UpdateListedFields(model interface{}, fields ...string) error {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Updating model %T with fields %v", model, fields)
	return adapter.dbx.Model(model).Update(fields...)
}

func (adapter *Adapter) Read(model interface{}, id any) error {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Reading model %T with id %v", model, id)
	return adapter.dbx.Select().Model(id, model)
}

func (adapter *Adapter) ExecuteInTransaction(
	transContext *SQLTransactionalContext,
	sql string,
	params ...any,
) (sql.Result, error) {
	var transaction = transContext.GetTransaction()
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing statement in transaction %s \n with params %s", sql, params)
	return transaction.Exec(sql, processParamsForPostgreSQL(params...)...)
}

func (adapter *Adapter) InsertBulkInTransaction(
	transContext *SQLTransactionalContext,
	tableName string,
	columns []string,
	values [][]any,
) error {

	var transaction = transContext.GetTransaction()

	statement, err := createBulkInsertPreparedStatement(tableName, columns, transaction)
	if err != nil {
		return err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			glog.Errorf("Error closing prepared statement: %v", err)
		}
	}(statement)

	return executeBulkInsertPreparedStatement(statement, values)
}

func createBulkInsertPreparedStatement(tableName string, columns []string, transaction *sql.Tx) (*sql.Stmt, error) {
	var copyInSQL = pq.CopyIn(tableName, columns...)
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Preparing statement in transaction %s", copyInSQL)
	return transaction.Prepare(copyInSQL)
}

func executeBulkInsertPreparedStatement(copyStatement *sql.Stmt, values [][]any) error {

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing statement in transaction with values %s", values)

	var err error
	for _, value := range values {
		_, err = copyStatement.Exec(value...)
		if err != nil {
			return err
		}
	}

	_, err = copyStatement.Exec()
	return err
}

func BuildDatabaseAdapter(config *infra.Configuration) *Adapter {
	return &Adapter{config: config.RdbmsConfig}
}

// =================================================
// RDBMS ADAPTER - Interfaces
// =================================================

type RepositoryRDBMSAdapter interface {
	BuildQuery(sql string) *QueryBuilder
	Insert(model interface{}) error
	UpdateListedFields(model interface{}, fields ...string) error
	Read(model interface{}, id any) error
	ExecuteInTransaction(transContext *SQLTransactionalContext, sql string, params ...any) (sql.Result, error)
	InsertBulkInTransaction(
		transContext *SQLTransactionalContext,
		tableName string,
		columns []string,
		values [][]any,
	) error
}

type TransactionManager interface {
	RunInTransaction(transactionalFunction func(transContext *SQLTransactionalContext) error) error
}
