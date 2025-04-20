package test

import (
	"context"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/root"
	"github.com/docker/docker/api/types/container"
	"github.com/golang/glog"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) { //TODO clean code

	if infra.ConfigLogger() {
		return
	}

	// Set up a test container for PostgreSQL
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

	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			glog.Errorf("failed to terminate container: %s", err)
		}
	}()

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

	// Set up migrations
	var flywayMigrationsPath = filepath.Join("..", "..", "flyway", "sql")
	var flywayConfigPath = filepath.Join("..", "..", "flyway", "conf")
	var flywayConnectionString = strings.ReplaceAll(connectionString, "postgres:localadmin@localhost", "172.17.0.1")
	flywayConnectionString = strings.ReplaceAll(flywayConnectionString, "postgres:", "jdbc:postgresql:")
	var flywayContainerRequest = testcontainers.ContainerRequest{
		Image: "flyway/flyway:10",
		Cmd: []string{
			"-url=" + flywayConnectionString,
			"-user=postgres",
			"-password=localadmin",
			"-connectRetries=10",
			"migrate",
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      flywayMigrationsPath,
				ContainerFilePath: "/flyway/sql",
				FileMode:          0755,
			},
			{
				HostFilePath:      flywayConfigPath,
				ContainerFilePath: "/flyway/conf",
				FileMode:          0755,
			},
		},
		HostConfigModifier: func(hc *container.HostConfig) {
			hc.NetworkMode = "host"
		},
		Env: map[string]string{
			"FLYWAY_DEBUG": "true",
		},
		WaitingFor: wait.ForLog("Successfully applied").WithStartupTimeout(10 * time.Second),
	}

	flywayContainer, err := testcontainers.GenericContainer(
		ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: flywayContainerRequest,
			Started:          true,
		},
	)

	// Even if the container failed to start properly, try to get the logs
	// The container might have started but exited quickly with an error
	containerID := flywayContainer.GetContainerID()
	if containerID != "" {
		// Get logs even if the container exited
		logReader, err := flywayContainer.Logs(ctx)
		if err != nil {
			glog.Errorf("failed to retrieve logs: %s", err)
		} else {
			defer logReader.Close()
			logContent, _ := io.ReadAll(logReader)
			glog.Infof("Flyway container logs:\n%s", string(logContent))
		}
	}

	if err != nil {
		glog.Errorf("failed to start flyway container: %s", err)
		os.Exit(1)
	}

	defer func() {
		if err := flywayContainer.Terminate(ctx); err != nil {
			glog.Errorf("failed to terminate flyway container: %s", err)
		}
	}()

	// set up application test server
	var appConnectionString = connectionString + "sslmode=disable"

	var ginServerConfig = infra.GinServerConfiguration{
		Port:    "8081",
		ApiOnly: true,
	}
	var dbConfig = infra.RDBMSConfiguration{
		DriverName: "postgres",
		RdbmsURL:   appConnectionString,
	}

	var testConfig = infra.Configuration{
		GinServerConfig: ginServerConfig,
		RdbmsConfig:     dbConfig,
	}

	var app = root.App{}
	app.StartOverridingConfigs(&testConfig)

	var exitVal = m.Run()

	// Stop and cleanup
	app.Stop()

	os.Exit(exitVal)
}
