package test

import (
	"context"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/root"
	"github.com/golang/glog"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) { //TODO clean code

	if infra.ConfigLogger() {
		return
	}

	ctx := context.Background()

	glog.Info("Starting PostgreSQL testcontainer...")
	postgresContainer, err := postgres.Run(
		ctx, "postgres:17.4-bullseye",
		//postgres.WithInitScripts(filepath.Join("..", "..", "postgres", "init-user-db.sh")),
		//postgres.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("localadmin"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		glog.Errorf("failed to start container: %s", err)
		os.Exit(1)
	}

	connectionString, err := postgresContainer.ConnectionString(ctx)
	if err != nil {
		glog.Errorf("failed to obtain connection string: %s", err)
		os.Exit(1)
	}
	glog.Info("PostgreSQL testcontainer initialized with no errors as ", connectionString)

	state, err := postgresContainer.State(ctx)
	if err != nil {
		glog.Errorf("failed to get container state: %s", err)
		os.Exit(1)
	}
	glog.Infof("PostgreSQL container state: %s", state.Status)

	//TODO setup DB migrations

	var ginServerConfig = infra.GinServerConfiguration{
		Port:    "8081",
		ApiOnly: true,
	}
	var dbConfig = infra.RDBMSConfiguration{
		DriverName: "postgres",
		RdbmsURL:   connectionString + "sslmode=disable",
	}

	var testConfig = infra.Configuration{
		GinServerConfig: ginServerConfig,
		RdbmsConfig:     dbConfig,
	}

	var app = root.App{}
	app.StartOverridingConfigs(&testConfig)
	time.Sleep(10 * time.Second)

	var exitVal = m.Run()

	app.Stop()

	if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
		glog.Errorf("failed to terminate container: %s", err)
	}

	os.Exit(exitVal)
}
