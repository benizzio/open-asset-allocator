// This test infrastructure records the SQL and positional arguments sent by real
// repository reads, without exposing database internals through production APIs.
// Authored by: GPT-6.1 Sol
package inttest

import (
	"context"
	"database/sql"
	"strings"
	"sync"
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
)

const paginationCaptureDriverName = "pagination-plan-postgres"

// capturedPaginationQuery retains executable SQL and its original bound arguments.
// Authored by: GPT-6.1 Sol
type capturedPaginationQuery struct {
	sql  string
	args []any
}

// paginationQueryCaptureExecutor delegates execution unchanged and records reads.
// Only the query-plan test uses this driver alias; it must not run parallel captures.
// Authored by: GPT-6.1 Sol
type paginationQueryCaptureExecutor struct {
	dbx.Executor
}

var paginationQueryCapture = struct {
	sync.Mutex
	queries []capturedPaginationQuery
}{}

// init registers an isolated PostgreSQL alias before application goroutines start.
// Authored by: GPT-6.1 Sol
func init() {
	sql.Register(paginationCaptureDriverName, &pq.Driver{})
	dbx.BuilderFuncMap[paginationCaptureDriverName] = newPaginationCaptureBuilder
}

// newPaginationCaptureBuilder preserves PostgreSQL SQL construction and wraps only execution.
// Authored by: GPT-6.1 Sol
func newPaginationCaptureBuilder(database *dbx.DB, executor dbx.Executor) dbx.Builder {
	return dbx.NewPgsqlBuilder(database, &paginationQueryCaptureExecutor{Executor: executor})
}

// Query records and executes an unchanged repository query, for example Query(sql, id).
// Authored by: GPT-6.1 Sol
func (capture *paginationQueryCaptureExecutor) Query(query string, args ...any) (*sql.Rows, error) {
	return capture.QueryContext(context.Background(), query, args...)
}

// QueryContext records SQL and arguments before delegating to the real database.
// For example, QueryContext(ctx, sql, id) retains the exact statement sent to PostgreSQL.
// Authored by: GPT-6.1 Sol
func (capture *paginationQueryCaptureExecutor) QueryContext(
	ctx context.Context,
	query string,
	args ...any,
) (*sql.Rows, error) {
	paginationQueryCapture.Lock()
	paginationQueryCapture.queries = append(paginationQueryCapture.queries, capturedPaginationQuery{
		sql: query, args: append([]any(nil), args...),
	})
	paginationQueryCapture.Unlock()
	return capture.Executor.QueryContext(ctx, query, args...)
}

// takePaginationQuery requires exactly one production read and clears it for the next case.
// Authored by: GPT-6.1 Sol
func takePaginationQuery(t *testing.T) capturedPaginationQuery {
	t.Helper()
	paginationQueryCapture.Lock()
	defer paginationQueryCapture.Unlock()
	require.Len(t, paginationQueryCapture.queries, 1)
	var query = paginationQueryCapture.queries[0]
	paginationQueryCapture.queries = nil
	return query
}

// buildPaginationCaptureAdapter connects the real repositories to the local test database.
// Authored by: GPT-6.1 Sol
func buildPaginationCaptureAdapter(t *testing.T) *rdbms.Adapter {
	t.Helper()
	var adapter = rdbms.BuildDatabaseAdapter(&infra.Configuration{
		RdbmsConfig: infra.RDBMSConfiguration{
			DriverName: paginationCaptureDriverName,
			RdbmsURL: strings.Replace(inttestinfra.PostgresqlConnectionString,
				inttestinfra.PostgresqlDefaultScheme, inttestinfra.PostgresqlGoScheme, 1) +
				inttestinfra.PostgresqlConnectionStringParameters,
		},
	})
	adapter.Init()
	t.Cleanup(func() {
		adapter.Stop()
		paginationQueryCapture.Lock()
		paginationQueryCapture.queries = nil
		paginationQueryCapture.Unlock()
	})
	return adapter
}
