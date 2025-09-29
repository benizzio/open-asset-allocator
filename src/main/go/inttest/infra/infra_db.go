package infra

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/benizzio/open-asset-allocator/infra/util"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// logFanOutConsumer is a thread-safe log consumer that buffers the latest N lines
// and forwards every new line to any attached testing.TB sinks in real-time.
// It implements testcontainers.LogConsumer.
//
// Authored by: GitHub Copilot
type logFanOutConsumer struct {
	mu     sync.Mutex
	cap    int
	lines  []string
	sinks  map[int]testing.TB
	nextID int
}

func newLogFanOutConsumer(capacity int) *logFanOutConsumer {
	if capacity <= 0 {
		capacity = 1000
	}
	return &logFanOutConsumer{
		cap:   capacity,
		lines: make([]string, 0, capacity),
		sinks: make(map[int]testing.TB),
	}
}

// Accept appends a new log line to the buffer and forwards it to all attached sinks.
//
// Authored by: GitHub Copilot
func (c *logFanOutConsumer) Accept(l testcontainers.Log) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var s = strings.TrimSuffix(string(l.Content), "\n")
	if len(c.lines) >= c.cap {
		// drop oldest
		copy(c.lines, c.lines[1:])
		c.lines[len(c.lines)-1] = s
	} else {
		c.lines = append(c.lines, s)
	}

	for _, t := range c.sinks {
		// best-effort logging; tests may have finished
		t.Logf("[postgres] %s", s)
	}
}

// Snapshot returns a copy of buffered lines for read-only use.
//
// Authored by: GitHub Copilot
func (c *logFanOutConsumer) Snapshot() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]string, len(c.lines))
	copy(out, c.lines)
	return out
}

func (c *logFanOutConsumer) DumpTo(t testing.TB) {
	for _, line := range c.Snapshot() {
		t.Logf("[postgres] %s", line)
	}
}

// Attach registers a testing.TB sink to receive live log lines. Returns an id for Detach.
//
// Authored by: GitHub Copilot
func (c *logFanOutConsumer) Attach(t testing.TB) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	id := c.nextID
	c.nextID++
	c.sinks[id] = t
	return id
}

// Detach unregisters a sink by id.
//
// Authored by: GitHub Copilot
func (c *logFanOutConsumer) Detach(id int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.sinks, id)
}

func appendDebugLoggingCommand(req *testcontainers.GenericContainerRequest) error {
	// Ensure logs go to stderr and enable statement logging and connection events
	req.Cmd = append(
		req.Cmd,
		"-c", "log_destination=stderr",
		"-c", "log_statement=all",
		"-c", "log_connections=on",
		"-c", "log_disconnections=on",
	)
	return nil
}

func withPosltgresLogging() testcontainers.CustomizeRequestOption {
	return appendDebugLoggingCommand
}

func buildAndRunPostgresqlTestcontainer(ctx context.Context) (string, error) {

	glog.Info("Starting PostgreSQL testcontainer...")

	postgresContainer, err := postgres.Run(
		ctx, PostgresqlImage,
		postgres.WithDatabase(PostgresqlDatabaseName),
		postgres.WithUsername(PostgresqlUsername),
		postgres.WithPassword(PostgresqlPassword),
		// Apply lightweight logging parameters (no full config override)
		withPosltgresLogging(),
		// Register our log consumer via ContainerRequest so logs are followed automatically
		testcontainers.WithLogConsumers(postgresLogFan),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)

	var deferRegistry = ctx.Value(DeferRegistryKey).(*util.DeferRegistry)

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

func buildDatabaseConnection(ctx context.Context) error {

	var err error
	DatabaseConnection, err = dbx.Open(
		DBDriverName,
		PostgresqlConnectionString+PostgresqlConnectionStringParameters,
	)
	if err != nil {
		glog.Errorf("Error opening DB connection: %s", err)
		return err
	}

	var deferRegistry = ctx.Value(DeferRegistryKey).(*util.DeferRegistry)
	var closeDBConnectionDefer = func() {
		err := DatabaseConnection.Close()
		if err != nil {
			glog.Errorf("Error closing DB connection: %s", err)
		}
	}
	deferRegistry.RegisterDefer(closeDBConnectionDefer)

	return nil
}
