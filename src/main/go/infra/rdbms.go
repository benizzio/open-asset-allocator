package infra

import (
	"context"
	"database/sql"
	"errors"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"time"
)

type QueryBuilder struct {
	query *dbx.Query
}

func (builder *QueryBuilder) AddParam(name string, value any) *QueryBuilder {
	builder.query.Bind(dbx.Params{name: value})
	return builder
}

func (builder *QueryBuilder) FindInto(target any) error {
	return builder.query.All(target)
}

func (builder *QueryBuilder) FetchInto(target any) error {
	return builder.query.One(target)
}

type RDBMSAdapter struct {
	config         RDBMSConfiguration
	connectionPool *sql.DB
	dbx            *dbx.DB
}

func (adapter *RDBMSAdapter) openPool() {

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Opening connection to %s at %s", adapter.config.driverName, adapter.config.rdbmsURL)

	var connectionError error
	adapter.connectionPool, connectionError = sql.Open(adapter.config.driverName, adapter.config.rdbmsURL)
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
	adapter.dbx = dbx.NewFromDB(adapter.connectionPool, adapter.config.driverName)
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

func (adapter *RDBMSAdapter) BuildQuery(sql string) *QueryBuilder {
	return &QueryBuilder{adapter.dbx.NewQuery(sql)}
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

func BuildDatabaseAdapter(config Configuration) *RDBMSAdapter {
	return &RDBMSAdapter{config: config.rdbmsConfig}
}
