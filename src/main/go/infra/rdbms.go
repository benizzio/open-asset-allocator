package infra

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"time"
)

type DatabaseAdapter struct {
	config         RDBMSConfiguration
	connectionPool *sql.DB
}

func (adapter *DatabaseAdapter) openPool() {

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Opening connection to %s at %s", adapter.config.driverName, adapter.config.rdbmsURL)

	var connectionError error
	adapter.connectionPool, connectionError = sql.Open(adapter.config.driverName, adapter.config.rdbmsURL)
	if connectionError != nil {
		glog.Fatal("Error opening database connection: ", connectionError)
		return
	}
}

func (adapter *DatabaseAdapter) configPool() {
	adapter.connectionPool.SetMaxOpenConns(5)
	adapter.connectionPool.SetMaxIdleConns(5)
	adapter.connectionPool.SetConnMaxLifetime(20 * time.Minute)
	adapter.connectionPool.SetConnMaxIdleTime(5 * time.Minute)
}

func (adapter *DatabaseAdapter) Init() {
	glog.Infof("Opening connection pool")
	adapter.openPool()
	glog.Infof("Configuring connection pool")
	adapter.configPool()
}

func (adapter *DatabaseAdapter) Stop() {
	err := adapter.connectionPool.Close()
	if err != nil {
		glog.Fatal("Error closing database connection: ", err)
		return
	}
}

func (adapter *DatabaseAdapter) Ping() {

	glog.Info("Pinging database to test connection")

	pingContext, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := adapter.connectionPool.PingContext(pingContext)
	if err != nil {
		glog.Fatal("Error pinging database connection: ", err)
		return
	}

	go func() {
		<-pingContext.Done()
		if errors.Is(pingContext.Err(), context.DeadlineExceeded) {
			glog.Fatal("Error pinging database connection: timeout")
		}
	}()

	glog.Info("Ping successful!")
}

func BuildDatabaseAdapter(config Configuration) *DatabaseAdapter {
	return &DatabaseAdapter{config: config.rdbmsConfig}
}
