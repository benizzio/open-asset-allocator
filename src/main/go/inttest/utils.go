package inttest

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
)

const (
	dbDriverName           = "postgres"
	postgresqlImage        = "postgres:17.4-bullseye"
	postgresqlDatabaseName = "postgres"
	postgresqlUsername
	postgresqlPassword                   = "localadmin"
	postgresqlConnectionStringParameters = "sslmode=disable"
	flywayImage                          = "flyway/flyway:10"
	testAPIURLprefix                     = "http://localhost:8081/api"
)

var postgresqlConnectionString string

func executeDBQuery(sql string) error {

	db, err := dbx.Open(dbDriverName, postgresqlConnectionString+postgresqlConnectionStringParameters)
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
