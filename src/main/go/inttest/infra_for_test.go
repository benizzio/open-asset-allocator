package inttest

import (
	"context"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/inttest/util"
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

func TestMain(m *testing.M) {

	if infra.ConfigLogger() {
		return
	}

	ctx := context.Background()

	var (
		postgresContainer *postgres.PostgresContainer
		err               error
	)
	postgresContainer, util.PostgresqlConnectionString, err = buildAndRunPostgresqlTestcontainer(ctx)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			glog.Errorf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		os.Exit(1)
	}

	err = runFlywayTestcontainer(ctx)
	if err != nil {
		os.Exit(1)
	}

	err = util.InitializeDBState()
	if err != nil {
		os.Exit(1)
	}

	app := buildAndStartApplication()
	defer func() {
		app.Stop()
	}()

	// run tests
	var exitVal = m.Run()

	os.Exit(exitVal)
}

func buildAndRunPostgresqlTestcontainer(ctx context.Context) (*postgres.PostgresContainer, string, error) {

	glog.Info("Starting PostgreSQL testcontainer...")
	postgresContainer, err := postgres.Run(
		ctx, util.PostgresqlImage,
		postgres.WithDatabase(util.PostgresqlDatabaseName),
		postgres.WithUsername(util.PostgresqlUsername),
		postgres.WithPassword(util.PostgresqlPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)

	if err != nil {
		glog.Errorf("failed to start container: %s", err)
		return nil, "", err
	}

	connectionString, err := postgresContainer.ConnectionString(ctx)
	if err != nil {
		glog.Errorf("failed to obtain connection string: %s", err)
		return nil, "", err
	}

	glog.Info("PostgreSQL testcontainer initialized with no errors as ", connectionString)

	state, err := postgresContainer.State(ctx)
	if err != nil {
		glog.Errorf("failed to get container state: %s", err)
		return nil, "", err
	}
	glog.Infof("PostgreSQL container state is '%s'", state.Status)

	return postgresContainer, connectionString, nil
}

func runFlywayTestcontainer(ctx context.Context) error {

	// Set up migrations
	var flywayMigrationsPath = filepath.Join("..", "..", "flyway", "sql")
	var flywayConfigPath = filepath.Join("..", "..", "flyway", "conf")

	var flywayConnectionString = strings.Replace(
		util.PostgresqlConnectionString,
		util.PostgresqlUsername+":"+util.PostgresqlPassword+"@localhost",
		"172.17.0.1",
		1,
	)
	flywayConnectionString = strings.Replace(flywayConnectionString, "postgres:", "jdbc:postgresql:", 1)

	glog.Info("Starting flyway testcontainer with connection ", flywayConnectionString)
	var flywayContainerRequest = testcontainers.ContainerRequest{
		Image: util.FlywayImage,
		Cmd: []string{
			"-url=" + flywayConnectionString,
			"-user=" + util.PostgresqlUsername,
			"-password=" + util.PostgresqlPassword,
			"-connectRetries=5",
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
		WaitingFor: wait.ForLog("Successfully applied").WithStartupTimeout(5 * time.Second),
	}

	flywayContainer, err := testcontainers.GenericContainer(
		ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: flywayContainerRequest,
			Started:          true,
		},
	)

	defer func() {
		if err := flywayContainer.Terminate(ctx); err != nil {
			glog.Errorf("failed to terminate flyway container: %s", err)
		}
	}()

	// Even if the container failed to start properly, try to get the logs
	// The container might have started but exited quickly with an error
	containerID := flywayContainer.GetContainerID()
	if containerID != "" {

		// Get logs even if the container exited
		logReader, err := flywayContainer.Logs(ctx)
		if err != nil {
			glog.Errorf("failed to retrieve logs: %s", err)
		} else {
			defer func(logReader io.ReadCloser) {
				err := logReader.Close()
				if err != nil {
					glog.Errorf("failed to close log reader: %s", err)
				}
			}(logReader)
			logContent, _ := io.ReadAll(logReader)
			glog.Infof("Flyway container logs:\n%s", string(logContent))
		}
	}

	if err != nil {
		glog.Errorf("failed to start flyway container: %s", err)
		return err
	}

	return nil
}

func buildAndStartApplication() root.App {

	// set up application test server
	var appConnectionString = util.PostgresqlConnectionString + util.PostgresqlConnectionStringParameters

	var ginServerConfig = infra.GinServerConfiguration{
		Port:    "8081",
		ApiOnly: true,
	}
	var dbConfig = infra.RDBMSConfiguration{
		DriverName: util.DBDriverName,
		RdbmsURL:   appConnectionString,
	}

	var testConfig = infra.Configuration{
		GinServerConfig: ginServerConfig,
		RdbmsConfig:     dbConfig,
	}

	var app = root.App{}
	app.StartOverridingConfigs(&testConfig)

	return app
}
