package infra

import (
	"context"
	"database/sql"
	"errors"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

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

// ================================================
// QUERY BUILDER
// ================================================

const (
	WhereClausePlaceholder = "/*WHERE+PARAMS*/"
)

type QueryBuilder struct {
	dbx          *dbx.DB
	querySQL     string
	whereClauses []string
	params       dbx.Params
}

func (builder *QueryBuilder) Build() *QueryExecutor {

	processedSQL := builder.processSQL()
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Building query for SQL: %s", processedSQL)

	var query = builder.dbx.NewQuery(processedSQL)
	var queryExecutor = withParams(query, builder.params)
	return queryExecutor
}

func (builder *QueryBuilder) processSQL() string {
	var processedSQL = builder.querySQL
	if len(builder.whereClauses) > 0 {
		var whereStatement = " WHERE 1=1 " + strings.Join(builder.whereClauses, " ")
		processedSQL = strings.Replace(processedSQL, WhereClausePlaceholder, whereStatement, 1)
	}
	return processedSQL
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

// ================================================
// DB ADAPTER
// ================================================

type RDBMSAdapter struct {
	config         RDBMSConfiguration
	connectionPool *sql.DB
	dbx            *dbx.DB
}

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

func (adapter *RDBMSAdapter) BuildQuery(sql string) *QueryBuilder {
	return &QueryBuilder{
		dbx:          adapter.dbx,
		querySQL:     sql,
		params:       dbx.Params{},
		whereClauses: make([]string, 0),
	}
}

func (adapter *RDBMSAdapter) Insert(model interface{}) error {
	return adapter.dbx.Model(model).Insert()
}

func (adapter *RDBMSAdapter) UpdateListedFields(model interface{}, fields ...string) error {
	return adapter.dbx.Model(model).Update(fields...)
}

func BuildDatabaseAdapter(config *Configuration) *RDBMSAdapter {
	return &RDBMSAdapter{config: config.RdbmsConfig}
}
