package inttest

import (
	"context"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/util"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
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

const deferRegistryKey = "deferRegistry"

func TestMain(m *testing.M) {

	ctx := context.Background()

	var deferRegistry = util.BuildDeferRegistry()
	ctx = context.WithValue(ctx, deferRegistryKey, deferRegistry)

	var exitVal = setupAndRunTests(ctx, m)

	deferRegistry.Execute()
	os.Exit(exitVal)
}

func setupAndRunTests(ctx context.Context, m *testing.M) int {

	var exitVal = setupTestInfra(ctx)
	if exitVal != 0 {
		return exitVal
	}

	var app = buildAndStartApplication()
	defer func() {
		app.Stop()
	}()

	// run tests
	return m.Run()
}

func setupTestInfra(ctx context.Context) int {

	if infra.ConfigLogger() {
		return 1
	}

	var err error

	inttestutil.PostgresqlConnectionString, err = buildAndRunPostgresqlTestcontainer(ctx)
	if err != nil {
		return 1
	}

	err = runFlywayTestcontainer(ctx)
	if err != nil {
		return 1
	}

	err = inttestutil.InitializeDBState()
	if err != nil {
		return 1
	}

	return 0
}

func buildAndRunPostgresqlTestcontainer(ctx context.Context) (string, error) {

	glog.Info("Starting PostgreSQL testcontainer...")
	postgresContainer, err := postgres.Run(
		ctx, inttestutil.PostgresqlImage,
		postgres.WithDatabase(inttestutil.PostgresqlDatabaseName),
		postgres.WithUsername(inttestutil.PostgresqlUsername),
		postgres.WithPassword(inttestutil.PostgresqlPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)

	var deferRegistry = ctx.Value(deferRegistryKey).(*util.DeferRegistry)
	var terminateContainerDefer = func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			glog.Errorf("failed to terminate container: %s", err)
		}
	}
	deferRegistry.RegisterDefer(terminateContainerDefer)

	if err != nil {
		glog.Errorf("failed to start container: %s", err)
		return "", err
	}

	connectionString, err := postgresContainer.ConnectionString(ctx)
	if err != nil {
		glog.Errorf("failed to obtain connection string: %s", err)
		return "", err
	}

	glog.Info("PostgreSQL testcontainer initialized with no errors as ", connectionString)

	state, err := postgresContainer.State(ctx)
	if err != nil {
		glog.Errorf("failed to get container state: %s", err)
		return "", err
	}
	glog.Infof("PostgreSQL container state is '%s'", state.Status)

	return connectionString, nil
}

func runFlywayTestcontainer(ctx context.Context) error {

	// Set up migrations
	var flywayMigrationsPath = filepath.Join("..", "..", "flyway", "sql")
	var flywayConfigPath = filepath.Join("..", "..", "flyway", "conf")

	var flywayConnectionString = strings.Replace(
		inttestutil.PostgresqlConnectionString,
		inttestutil.PostgresqlUsername+":"+inttestutil.PostgresqlPassword+"@localhost",
		"172.17.0.1",
		1,
	)
	flywayConnectionString = strings.Replace(
		flywayConnectionString,
		inttestutil.PostgresqlDefaultScheme,
		inttestutil.PostgresqlJDBCScheme,
		1,
	)

	glog.Info("Starting flyway testcontainer with connection ", flywayConnectionString)
	var flywayContainerRequest = testcontainers.ContainerRequest{
		Image: inttestutil.FlywayImage,
		Cmd: []string{
			"-url=" + flywayConnectionString,
			"-user=" + inttestutil.PostgresqlUsername,
			"-password=" + inttestutil.PostgresqlPassword,
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
	var appConnectionString = strings.Replace(
		inttestutil.PostgresqlConnectionString,
		inttestutil.PostgresqlDefaultScheme,
		inttestutil.PostgresqlGoScheme,
		1,
	) + inttestutil.PostgresqlConnectionStringParameters

	var ginServerConfig = infra.GinServerConfiguration{
		Port:    "8081",
		ApiOnly: true,
	}
	var dbConfig = infra.RDBMSConfiguration{
		DriverName: inttestutil.DBDriverName,
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
