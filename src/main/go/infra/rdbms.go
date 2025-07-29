package infra

import (
	"context"
	"database/sql"
	"errors"
	"github.com/benizzio/open-asset-allocator/langext"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
	"github.com/lib/pq"
	"strings"
	"time"
)

// TODO split in rdbms package

// ================================================
// TRANSACTIONAL CONTEXT
// ================================================

const sqlTransactionContextKey = "TRANSACTION"

type TransactionalContext struct {
	context.Context
}

func (transactionalContext *TransactionalContext) GetTransaction() *sql.Tx {
	return transactionalContext.Context.Value(sqlTransactionContextKey).(*sql.Tx)
}

func withTransaction(db *sql.DB) (*TransactionalContext, error) {

	var transaction, err = db.Begin()
	if err != nil {
		return nil, err
	}

	var parentContext = context.WithValue(context.Background(), sqlTransactionContextKey, transaction)
	return &TransactionalContext{parentContext}, nil
}

// ================================================
// QUERY EXECUTOR
// ================================================

type QueryExecutor struct {
	query *dbx.Query
}

func withParams(query *dbx.Query, params dbx.Params) *QueryExecutor {
	if len(params) >= 0 {
		query.Bind(params)
	}
	return &QueryExecutor{query: query}
}

func (executor *QueryExecutor) FindInto(target any) error {
	return executor.query.All(target)
}

func (executor *QueryExecutor) GetInto(target any) error {
	return executor.query.One(target)
}

func (executor *QueryExecutor) GetRows() (*dbx.Rows, error) {
	return executor.query.Rows()
}

type SQLTransactionalQueryExecutor[T any] struct {
	queryBuilder *SQLTransactionalQueryBuilder[T]
}

type RowScanner[T any] func(*sql.Rows) (T, error)

func (executor *SQLTransactionalQueryExecutor[T]) Find(rowScanner RowScanner[T]) ([]T, error) {

	var builder = executor.queryBuilder

	rows, err := builder.transaction.Query(builder.querySQL, builder.params...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			glog.Errorf("Error closing rows: %v", closeErr)
		}
	}()

	var result = make([]T, 0)
	var index = 0
	for rows.Next() {

		rowValue, scanErr := rowScanner(rows)
		if scanErr != nil {
			glog.Errorf("Error scanning row: %v", index)
			return nil, scanErr
		}

		result = append(result, rowValue)
		index++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// ================================================
// QUERY BUILDER
// ================================================

const (
	WhereClausePlaceholder = "/*WHERE+PARAMS*/"
)

func processSQL(querySQL string, whereClauses []string) string {

	var processedSQL = querySQL

	if len(whereClauses) > 0 {
		var whereStatement = " WHERE 1=1 " + strings.Join(whereClauses, " ")
		processedSQL = strings.Replace(processedSQL, WhereClausePlaceholder, whereStatement, 1)
	} else {
		processedSQL = strings.Replace(processedSQL, WhereClausePlaceholder, "", 1)
	}

	return processedSQL
}

type QueryBuilder struct {
	dbx          *dbx.DB
	querySQL     string
	whereClauses []string
	params       dbx.Params
}

func (builder *QueryBuilder) Build() *QueryExecutor {

	var processedSQL = processSQL(builder.querySQL, builder.whereClauses)
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Building query for SQL: %s", processedSQL)

	var query = builder.dbx.NewQuery(processedSQL)
	var queryExecutor = withParams(query, builder.params)
	return queryExecutor
}

func (builder *QueryBuilder) AddParam(name string, value any) *QueryBuilder {
	builder.params[name] = value
	return builder
}

func (builder *QueryBuilder) AddWhereClause(whereClause string) *QueryBuilder {
	builder.whereClauses = append(builder.whereClauses, whereClause)
	return builder
}

func (builder *QueryBuilder) AddWhereClauseAndParam(whereClause string, name string, value any) *QueryBuilder {
	return builder.AddWhereClause(whereClause).AddParam(name, value)
}

type SQLTransactionalQueryBuilder[T any] struct {
	transaction  *sql.Tx
	querySQL     string
	whereClauses []string
	params       []any
}

func (builder *SQLTransactionalQueryBuilder[T]) AddParams(value ...any) *SQLTransactionalQueryBuilder[T] {
	builder.params = append(builder.params, value)
	return builder
}

func (builder *SQLTransactionalQueryBuilder[T]) AddWhereClause(whereClause string) *SQLTransactionalQueryBuilder[T] {
	builder.whereClauses = append(builder.whereClauses, whereClause)
	return builder
}

// processParamsForPostgreSQL converts slice parameters to pq.Array for PostgreSQL compatibility.
//
// Parameters:
//   - params: Variable number of parameters that may include slices
//
// Returns:
//   - []any: Processed parameters with slices converted to pq.Array
//
// Authored by: GitHub Copilot
func processParamsForPostgreSQL(params ...any) []any {

	var processedParams = make([]any, len(params))

	for i, param := range params {
		if langext.IsSlice(param) {
			processedParams[i] = pq.Array(param)
		} else {
			processedParams[i] = param
		}
	}

	return processedParams
}

// AddWhereClauseAndParams adds a WHERE clause and its parameters to the query builder.
//
// This method automatically converts slice parameters to pq.Array for PostgreSQL
// compatibility, allowing seamless use of slices in SQL IN clauses and array operations.
//
// Parameters:
//   - whereClause: The WHERE clause SQL fragment to add
//   - params: Variable number of parameters, with slices automatically converted to pq.Array
//
// Returns:
//   - *SQLTransactionalQueryBuilder[T]: The same builder instance for method chaining
//
// Co-authored by: GitHub Copilot
func (builder *SQLTransactionalQueryBuilder[T]) AddWhereClauseAndParams(
	whereClause string,
	params ...any,
) *SQLTransactionalQueryBuilder[T] {

	builder.whereClauses = append(builder.whereClauses, whereClause)

	var processedParams = processParamsForPostgreSQL(params...)
	builder.params = append(builder.params, processedParams...)

	return builder
}

func (builder *SQLTransactionalQueryBuilder[T]) Build() *SQLTransactionalQueryExecutor[T] {

	var processedSQL = processSQL(builder.querySQL, builder.whereClauses)
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Building transactional query for SQL: %s", processedSQL)
	builder.querySQL = processedSQL

	return &SQLTransactionalQueryExecutor[T]{
		queryBuilder: builder,
	}
}

// ================================================
// DB ADAPTER
// ================================================

type RDBMSAdapter struct {
	config         RDBMSConfiguration
	connectionPool *sql.DB
	dbx            *dbx.DB
}

type RepositoryRDBMSAdapter interface {
	BuildQuery(sql string) *QueryBuilder
	Insert(model interface{}) error
	UpdateListedFields(model interface{}, fields ...string) error
	Read(model interface{}, id any) error
	ExecuteInTransaction(transContext *TransactionalContext, sql string, params ...any) (sql.Result, error)
	InsertBulkInTransaction(
		transContext *TransactionalContext,
		tableName string,
		columns []string,
		values [][]any,
	) error
}

type TransactionManager interface {
	RunInTransaction(transactionalFunction func(transContext *TransactionalContext) error) error
}

// ------------------------------------------------
// SETUP FUNCTIONS
// ------------------------------------------------

func (adapter *RDBMSAdapter) openPool() {

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Opening connection to %s at %s", adapter.config.DriverName, adapter.config.RdbmsURL)

	var connectionError error
	adapter.connectionPool, connectionError = sql.Open(adapter.config.DriverName, adapter.config.RdbmsURL)
	if connectionError != nil {
		glog.Fatal("Error opening database connection: ", connectionError)
		return
	}
}

func (adapter *RDBMSAdapter) configPool() {
	adapter.connectionPool.SetMaxOpenConns(5)
	adapter.connectionPool.SetMaxIdleConns(5)
	adapter.connectionPool.SetConnMaxLifetime(20 * time.Minute)
	adapter.connectionPool.SetConnMaxIdleTime(5 * time.Minute)
}

func (adapter *RDBMSAdapter) buildDBX() {
	adapter.dbx = dbx.NewFromDB(adapter.connectionPool, adapter.config.DriverName)
}

func (adapter *RDBMSAdapter) Init() {
	glog.Infof("Opening connection pool")
	adapter.openPool()
	glog.Infof("Configuring connection pool")
	adapter.configPool()
	glog.Infof("Configuring ozzo-dbx enhancer")
	adapter.buildDBX()
}

func (adapter *RDBMSAdapter) Stop() {
	err := adapter.connectionPool.Close()
	if err != nil {
		glog.Fatal("Error closing database connection: ", err)
		return
	}
}

func (adapter *RDBMSAdapter) Ping() {

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

func (adapter *RDBMSAdapter) BuildQuery(sql string) *QueryBuilder {
	return &QueryBuilder{
		dbx:          adapter.dbx,
		querySQL:     sql,
		params:       dbx.Params{},
		whereClauses: make([]string, 0),
	}
}

func BuildQueryWithinTransaction[T any](
	transContext *TransactionalContext,
	sql string,
) *SQLTransactionalQueryBuilder[T] {
	return &SQLTransactionalQueryBuilder[T]{
		transaction:  transContext.GetTransaction(),
		querySQL:     sql,
		params:       make([]any, 0),
		whereClauses: make([]string, 0),
	}
}

func (adapter *RDBMSAdapter) buildTransactionalContext() (*TransactionalContext, error) {
	var transactionContext, err = withTransaction(adapter.connectionPool)
	if err != nil {
		return nil, err
	}
	return transactionContext, nil
}

func (adapter *RDBMSAdapter) RunInTransaction(
	transactionalFunction func(transContext *TransactionalContext) error,
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

func (adapter *RDBMSAdapter) runInTransaction(
	transContext *TransactionalContext,
	transactionalFunction func(transContext *TransactionalContext) error,
) error {

	err := transactionalFunction(transContext)

	var transaction = transContext.GetTransaction()
	if err != nil {
		if rollbackErr := transaction.Rollback(); rollbackErr != nil {
			return BuildAppError(
				"transaction rollback failed: "+rollbackErr.Error()+"; original error: "+err.Error(),
				adapter,
			)
		}
		return err
	}

	return transaction.Commit()
}

func (adapter *RDBMSAdapter) Insert(model interface{}) error {
	return adapter.dbx.Model(model).Insert()
}

func (adapter *RDBMSAdapter) UpdateListedFields(model interface{}, fields ...string) error {
	return adapter.dbx.Model(model).Update(fields...)
}

func (adapter *RDBMSAdapter) Read(model interface{}, id any) error {
	return adapter.dbx.Select().Model(id, model)
}

func (adapter *RDBMSAdapter) ExecuteInTransaction(
	transContext *TransactionalContext,
	sql string,
	params ...any,
) (sql.Result, error) {
	var transaction = transContext.GetTransaction()
	return transaction.Exec(sql, processParamsForPostgreSQL(params...)...)
}

func (adapter *RDBMSAdapter) InsertBulkInTransaction(
	transContext *TransactionalContext,
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
	return transaction.Prepare(copyInSQL)
}

func executeBulkInsertPreparedStatement(copyStatement *sql.Stmt, values [][]any) error {

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

func BuildDatabaseAdapter(config *Configuration) *RDBMSAdapter {
	return &RDBMSAdapter{config: config.RdbmsConfig}
}
