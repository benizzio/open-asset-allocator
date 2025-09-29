package infra

import (
	"context"
	"io"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/root"
	"github.com/docker/docker/api/types/container"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	DBDriverName           = "postgres"
	PostgresqlImage        = "postgres:17.5-bullseye"
	PostgresqlDatabaseName = "postgres"
	PostgresqlUsername
	PostgresqlDefaultScheme              = "postgres:"
	PostgresqlJDBCScheme                 = "jdbc:postgresql:"
	PostgresqlGoScheme                   = "postgresql:"
	PostgresqlPassword                   = "localadmin"
	PostgresqlConnectionStringParameters = "sslmode=disable"
	FlywayImage                          = "flyway/flyway:10"
	TestAPIURLPrefix                     = "http://localhost:8081/api"
	DeferRegistryKey                     = "deferRegistry"
)

var (
	PostgresqlConnectionString string
	DatabaseConnection         *dbx.DB
	postgresLogFan             = newLogFanOutConsumer(5000)
)

func SetupTestInfra(ctx context.Context) int {

	if infra.ConfigLogger() {
		return 1
	}

	var err error

	PostgresqlConnectionString, err = buildAndRunPostgresqlTestcontainer(ctx)
	if err != nil {
		return 1
	}

	err = runFlywayTestcontainer(ctx)
	if err != nil {
		return 1
	}

	err = buildDatabaseConnection(ctx)
	if err != nil {
		return 1
	}

	err = InitializeDBState()
	if err != nil {
		return 1
	}

	return 0
}

func runFlywayTestcontainer(ctx context.Context) error {

	// Set up migrations
	var flywayMigrationsPath = filepath.Join("..", "..", "flyway", "sql")
	var flywayConfigPath = filepath.Join("..", "..", "flyway", "conf")

	var flywayConnectionString = strings.Replace(
		PostgresqlConnectionString,
		PostgresqlUsername+":"+PostgresqlPassword+"@localhost",
		"172.17.0.1",
		1,
	)
	flywayConnectionString = strings.Replace(
		flywayConnectionString,
		PostgresqlDefaultScheme,
		PostgresqlJDBCScheme,
		1,
	)

	glog.Info("Starting flyway testcontainer with connection ", flywayConnectionString)
	var flywayContainerRequest = testcontainers.ContainerRequest{
		Image: FlywayImage,
		Cmd: []string{
			"-url=" + flywayConnectionString,
			"-user=" + PostgresqlUsername,
			"-password=" + PostgresqlPassword,
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

func BuildAndStartApplication() root.App {

	// set up application test server
	var appConnectionString = strings.Replace(
		PostgresqlConnectionString,
		PostgresqlDefaultScheme,
		PostgresqlGoScheme,
		1,
	) + PostgresqlConnectionStringParameters

	var ginServerConfig = infra.GinServerConfiguration{
		Port:    "8081",
		ApiOnly: true,
	}
	var dbConfig = infra.RDBMSConfiguration{
		DriverName: DBDriverName,
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

func ExecuteDBQuery(sql string) error {

	var query = DatabaseConnection.NewQuery(sql)
	_, err := query.Execute()
	if err != nil {
		glog.Errorf("Error executing query: %s", err)
		return err
	}

	return nil
}

func FetchWithDBQuery(sql string, rowMappingFunction func(rows *dbx.Rows) error) error {

	var query = DatabaseConnection.NewQuery(sql)
	rows, err := query.Rows()
	if err != nil {
		glog.Errorf("Error executing query: %s", err)
		return err
	}

	defer func(rows *dbx.Rows) {
		err := rows.Close()
		if err != nil {
			glog.Errorf("Error closing rows: %s", err)
		}
	}(rows)

	for rows.Next() {
		if err := rowMappingFunction(rows); err != nil {
			glog.Errorf("Error mapping row: %s", err)
			return err
		}
	}

	return nil
}

// AttachDatabaseLogsTo attaches a live log consumer so DB logs are printed
// to the provided testing.TB (t.Logf) as they are written by the container.
// It also dumps the buffered history immediately so the test has context.
// A cleanup hook is registered (when available) to auto-detach after the test.
//
// Useful to debug test failures that might be related to DB issues.
//
// Usage:
//
//	inttestinfra.AttachDatabaseLogsTo(t)
//
// Authored by: GitHub Copilot
func AttachDatabaseLogsTo(t testing.TB) {

	t.Helper()

	// Dump current buffer for context
	postgresLogFan.DumpTo(t)

	// Attach real-time sink
	sinkID := postgresLogFan.Attach(t)

	// Auto-detach when the test supports Cleanup
	type cleaner interface{ Cleanup(func()) }
	if c, ok := any(t).(cleaner); ok {
		c.Cleanup(func() { postgresLogFan.Detach(sinkID) })
	}
}
