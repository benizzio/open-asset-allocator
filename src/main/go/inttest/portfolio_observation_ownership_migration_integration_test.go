package inttest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
)

// TODO temporary test to guarantee functionality of a complex migration
// aftrer the code is consolidated it can be removed

// migratedObservation is one persisted observation and owner returned by the migration test queries.
// Authored by: OpenCode
type migratedObservation struct {
	id        int
	portfolio int
	tag       string
	timestamp time.Time
}

// TestPortfolioObservationOwnershipMigration splits shared legacy observations and verifies the migration constraints.
// Authored by: OpenCode
func TestPortfolioObservationOwnershipMigration(t *testing.T) {
	var connectionString = strings.TrimRight(inttestinfra.PostgresqlConnectionString, "?") + "?sslmode=disable"
	database, err := sql.Open(inttestinfra.DBDriverName, connectionString)
	require.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, database.Close()) })

	var schemaName = fmt.Sprintf("pr7_migration_test_%d", os.Getpid())
	_, err = database.ExecContext(context.Background(), "CREATE SCHEMA "+pq.QuoteIdentifier(schemaName))
	require.NoError(t, err)
	t.Cleanup(func() {
		_, cleanupError := database.ExecContext(
			context.Background(),
			"DROP SCHEMA IF EXISTS "+pq.QuoteIdentifier(schemaName)+" CASCADE",
		)
		assert.NoError(t, cleanupError)
	})

	transaction, err := database.BeginTx(context.Background(), nil)
	require.NoError(t, err)
	defer func() { assert.NoError(t, transaction.Rollback()) }()
	_, err = transaction.ExecContext(
		context.Background(),
		"SET LOCAL search_path TO "+pq.QuoteIdentifier(schemaName),
	)
	require.NoError(t, err)
	createLegacyObservationSchema(t, transaction)
	seedLegacyObservations(t, transaction)

	var migrationPath = filepath.Join(
		"..", "..", "..", "..", "src", "main", "flyway", "sql",
		"migration-[16]-portfolio_observation_ownership.sql",
	)
	migrationSQL, err := os.ReadFile(migrationPath)
	require.NoError(t, err)
	for _, statement := range strings.Split(string(migrationSQL), ";") {
		if strings.TrimSpace(statement) == "" {
			continue
		}
		_, err = transaction.ExecContext(context.Background(), statement)
		require.NoError(t, err)
	}

	assertMigratedObservationOwnership(t, transaction)
	assertMigrationConstraints(t, transaction)
	assertObservationSequenceAfterMigration(t, transaction)
}

// createLegacyObservationSchema creates the relevant version-15 schema shape inside the isolated test schema.
// Authored by: OpenCode
func createLegacyObservationSchema(t *testing.T, transaction *sql.Tx) {
	t.Helper()
	var schemaStatements = []string{
		`CREATE TABLE portfolio (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			allocation_structure JSONB NOT NULL
		)`,
		`CREATE TABLE portfolio_allocation_obs_time (
			id SERIAL NOT NULL,
			observation_timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
			observation_time_tag VARCHAR(100) NOT NULL,
			CONSTRAINT portfolio_allocation_obs_time_pk PRIMARY KEY (id),
			CONSTRAINT portfolio_allocation_obs_time_observation_time_tag_unique UNIQUE (observation_time_tag)
		)`,
		`CREATE OR REPLACE FUNCTION set_observation_time_tag()
		RETURNS TRIGGER AS $$
		BEGIN
			IF NEW.observation_time_tag IS NULL THEN
				NEW.observation_time_tag := to_char(NEW.observation_timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"');
			END IF;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql`,
		`CREATE TRIGGER trg_set_observation_time_tag
		BEFORE INSERT OR UPDATE ON portfolio_allocation_obs_time
		FOR EACH ROW EXECUTE FUNCTION set_observation_time_tag()`,
		`CREATE TABLE portfolio_allocation_fact (
			asset_id INTEGER NOT NULL,
			class TEXT NOT NULL,
			cash_reserve BOOLEAN NOT NULL,
			asset_quantity NUMERIC,
			asset_market_price NUMERIC,
			total_market_value BIGINT NOT NULL,
			portfolio_id INTEGER NOT NULL,
			observation_time_id INTEGER NOT NULL,
			CONSTRAINT portfolio_allocation_fact_pk
				PRIMARY KEY (portfolio_id, class, cash_reserve, asset_id, observation_time_id),
			CONSTRAINT portfolio_fk FOREIGN KEY (portfolio_id) REFERENCES portfolio (id)
				ON DELETE RESTRICT ON UPDATE CASCADE,
			CONSTRAINT portfolio_allocation_obs_time_fk FOREIGN KEY (observation_time_id)
				REFERENCES portfolio_allocation_obs_time (id) MATCH FULL
				ON DELETE SET NULL ON UPDATE CASCADE
		)`,
	}
	for _, statement := range schemaStatements {
		_, err := transaction.ExecContext(context.Background(), statement)
		require.NoError(t, err)
	}
}

// seedLegacyObservations inserts exclusive, shared, and orphan observations with distinguishable fact values.
// Authored by: OpenCode
func seedLegacyObservations(t *testing.T, transaction *sql.Tx) {
	t.Helper()
	var seedSQL = `
		INSERT INTO portfolio (id, name, allocation_structure) VALUES
			(1, 'Portfolio 1', '{}'::jsonb),
			(2, 'Portfolio 2', '{}'::jsonb),
			(3, 'Portfolio 3', '{}'::jsonb);
		INSERT INTO portfolio_allocation_obs_time (id, observation_time_tag, observation_timestamp) VALUES
			(10, 'shared-observation', '2042-01-02T03:04:05.123456Z'),
			(11, 'exclusive-observation', '2042-02-03T04:05:06.654321Z'),
			(12, 'orphan-observation', '2042-03-04T05:06:07Z');
		INSERT INTO portfolio_allocation_fact (
			asset_id, class, cash_reserve, asset_quantity, asset_market_price, total_market_value,
			portfolio_id, observation_time_id
		) VALUES
			(101, 'BONDS', FALSE, 1.1, 100, 100, 1, 10),
			(101, 'BONDS', FALSE, 2.2, 100, 200, 2, 10),
			(101, 'BONDS', FALSE, 3.3, 100, 300, 3, 10),
			(102, 'STOCKS', TRUE, 4.4, 100, 400, 2, 11);
	`
	_, err := transaction.ExecContext(context.Background(), seedSQL)
	require.NoError(t, err)
}

// assertMigratedObservationOwnership checks deterministic ID retention, metadata copying, orphan deletion, and fact integrity.
// Authored by: OpenCode
func assertMigratedObservationOwnership(t *testing.T, transaction *sql.Tx) {
	t.Helper()
	var sharedObservations = make([]migratedObservation, 0, 3)
	rows, err := transaction.QueryContext(
		context.Background(),
		`SELECT id, portfolio_id, observation_time_tag, observation_timestamp
		 FROM portfolio_allocation_obs_time WHERE observation_time_tag = 'shared-observation'
		 ORDER BY portfolio_id`,
	)
	require.NoError(t, err)
	for rows.Next() {
		var observation migratedObservation
		assert.NoError(t, rows.Scan(
			&observation.id,
			&observation.portfolio,
			&observation.tag,
			&observation.timestamp,
		))
		sharedObservations = append(sharedObservations, observation)
	}
	assert.NoError(t, rows.Err())
	assert.NoError(t, rows.Close())
	require.Len(t, sharedObservations, 3)
	assert.Equal(t, []int{1, 2, 3}, []int{
		sharedObservations[0].portfolio,
		sharedObservations[1].portfolio,
		sharedObservations[2].portfolio,
	})
	assert.Equal(t, 10, sharedObservations[0].id, "lowest portfolio ID retains the legacy observation ID")
	assert.NotEqual(t, 10, sharedObservations[1].id)
	assert.NotEqual(t, 10, sharedObservations[2].id)
	assert.NotEqual(t, sharedObservations[1].id, sharedObservations[2].id)
	for _, observation := range sharedObservations {
		assert.Equal(t, "shared-observation", observation.tag)
		assert.Equal(t, time.Date(2042, 1, 2, 3, 4, 5, 123456000, time.UTC), observation.timestamp)
	}

	var exclusiveId int
	err = transaction.QueryRowContext(
		context.Background(),
		`SELECT id FROM portfolio_allocation_obs_time
		 WHERE portfolio_id = 2 AND observation_time_tag = 'exclusive-observation'`,
	).Scan(&exclusiveId)
	require.NoError(t, err)
	assert.Equal(t, 11, exclusiveId)

	var orphanCount int
	err = transaction.QueryRowContext(
		context.Background(),
		`SELECT COUNT(*) FROM portfolio_allocation_obs_time WHERE observation_time_tag = 'orphan-observation'`,
	).Scan(&orphanCount)
	require.NoError(t, err)
	assert.Zero(t, orphanCount)

	var factCount int
	var factValue int64
	var mismatchedOwnerCount int
	err = transaction.QueryRowContext(
		context.Background(),
		`SELECT COUNT(*), SUM(total_market_value),
		        COUNT(*) FILTER (WHERE fact.portfolio_id <> observation.portfolio_id)
		 FROM portfolio_allocation_fact fact
		 JOIN portfolio_allocation_obs_time observation ON observation.id = fact.observation_time_id`,
	).Scan(&factCount, &factValue, &mismatchedOwnerCount)
	require.NoError(t, err)
	assert.Equal(t, 4, factCount)
	assert.EqualValues(t, 1000, factValue)
	assert.Zero(t, mismatchedOwnerCount)
}

// assertMigrationConstraints verifies portfolio-local tags, composite ownership, and required ownership at the database boundary.
// Authored by: OpenCode
func assertMigrationConstraints(t *testing.T, transaction *sql.Tx) {
	t.Helper()
	_, err := transaction.ExecContext(
		context.Background(),
		`INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		 VALUES (1, 'cross-owner-tag', '2043-01-01T00:00:00Z'),
		        (2, 'cross-owner-tag', '2043-01-01T00:00:00Z')`,
	)
	require.NoError(t, err)

	assertExpectedMigrationFailure(
		t,
		transaction,
		"duplicate_observation_tag",
		`INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		 VALUES (1, 'cross-owner-tag', '2043-01-02T00:00:00Z')`,
		"23505",
		"portfolio_allocation_obs_time_portfolio_tag_uk",
	)
	assertExpectedMigrationFailure(
		t,
		transaction,
		"cross_portfolio_fact_observation",
		`INSERT INTO portfolio_allocation_fact (
			asset_id, class, cash_reserve, asset_quantity, asset_market_price, total_market_value,
			portfolio_id, observation_time_id
		) VALUES (103, 'TEST', FALSE, 1, 1, 1, 1,
			(SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 2 AND observation_time_tag = 'shared-observation'))`,
		"23503",
		"portfolio_allocation_fact_observation_owner_fk",
	)
	assertExpectedMigrationFailure(
		t,
		transaction,
		"missing_observation_owner",
		`INSERT INTO portfolio_allocation_obs_time (observation_time_tag, observation_timestamp)
		 VALUES ('missing-owner', '2043-01-03T00:00:00Z')`,
		"23502",
		"",
	)
}

// assertExpectedMigrationFailure runs a rejected insert inside a savepoint and checks its PostgreSQL error identity.
// Authored by: OpenCode
func assertExpectedMigrationFailure(
	t *testing.T,
	transaction *sql.Tx,
	savepoint, statement, expectedCode, expectedConstraint string,
) {
	t.Helper()
	_, err := transaction.ExecContext(context.Background(), "SAVEPOINT "+pq.QuoteIdentifier(savepoint))
	require.NoError(t, err)
	_, err = transaction.ExecContext(context.Background(), statement)
	require.Error(t, err)
	var postgresError *pq.Error
	require.True(t, errors.As(err, &postgresError))
	assert.Equal(t, expectedCode, string(postgresError.Code))
	if expectedConstraint != "" {
		assert.Equal(t, expectedConstraint, postgresError.Constraint)
	}
	_, err = transaction.ExecContext(context.Background(), "ROLLBACK TO SAVEPOINT "+pq.QuoteIdentifier(savepoint))
	require.NoError(t, err)
	_, err = transaction.ExecContext(context.Background(), "RELEASE SAVEPOINT "+pq.QuoteIdentifier(savepoint))
	require.NoError(t, err)
}

// assertObservationSequenceAfterMigration confirms sequence-generated IDs remain unique after explicit legacy IDs and copies.
// Authored by: OpenCode
func assertObservationSequenceAfterMigration(t *testing.T, transaction *sql.Tx) {
	t.Helper()
	var generatedId int
	err := transaction.QueryRowContext(
		context.Background(),
		`INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		 VALUES (1, 'post-migration-sequence', '2044-01-01T00:00:00Z') RETURNING id`,
	).Scan(&generatedId)
	require.NoError(t, err)
	assert.Greater(t, generatedId, 14)
}
