package util

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
)

const (
	DBDriverName           = "postgres"
	PostgresqlImage        = "postgres:17.4-bullseye"
	PostgresqlDatabaseName = "postgres"
	PostgresqlUsername
	PostgresqlDefaultScheme              = "postgres:"
	PostgresqlJDBCScheme                 = "jdbc:postgresql:"
	PostgresqlGoScheme                   = "postgresql:"
	PostgresqlPassword                   = "localadmin"
	PostgresqlConnectionStringParameters = "sslmode=disable"
	FlywayImage                          = "flyway/flyway:10"
	TestAPIURLprefix                     = "http://localhost:8081/api"
)

var PostgresqlConnectionString string

func ExecuteDBQuery(sql string) error {

	db, err := dbx.Open(DBDriverName, PostgresqlConnectionString+PostgresqlConnectionStringParameters)
	if err != nil {
		glog.Errorf("Error openingDB connection: %s", err)
		return err
	}

	defer func(db *dbx.DB) {
		err := db.Close()
		if err != nil {
			glog.Errorf("Error closing DB connection: %s", err)
		}
	}(db)

	var query = db.NewQuery(sql)
	_, err = query.Execute()
	if err != nil {
		glog.Errorf("Error executing query: %s", err)
		return err
	}

	return nil
}
