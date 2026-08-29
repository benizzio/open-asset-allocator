/**
 * Provides PostgreSQL administration helpers for isolated E2E test attempts.
 *
 * The runner connects with PGHOST, PGPORT, PGDATABASE, PGUSER, and PGPASSWORD.
 * For example, set PGHOST=db and PGUSER=postgres in the E2E Compose service.
 *
 * Authored by: OpenCode
 */
import { Pool, type QueryResultRow } from 'pg';

const FLYWAY_HISTORY_TABLES = ['flyway_schema_history', 'schema_history'];
const PUBLIC_SCHEMA = 'public';

/**
 * Executes queries and resets the disposable E2E PostgreSQL database.
 *
 * Create one instance per Playwright worker with createE2eDatabase(), then call
 * reset() before and after every test attempt and close() when the worker exits.
 *
 * Authored by: OpenCode
 */
export class E2eDatabase {
  /** Creates an E2E database helper around an administrator PostgreSQL pool. */
  public constructor(private readonly pool: Pool) {}

  /**
   * Runs a parameterized SQL query against the disposable E2E database.
   *
   * Example: await database.query('SELECT id FROM public.portfolio WHERE name = $1', [name]).
   */
  public async query<Row extends QueryResultRow>(text: string, values: unknown[] = []): Promise<readonly Row[]> {
    const result = await this.pool.query<Row>(text, values);
    return result.rows;
  }

  /**
   * Clears application tables and resets identities without modifying Flyway history.
   *
   * Example: await database.reset() before a scenario to begin from an empty schema.
   */
  public async reset(): Promise<void> {
    const tables = await this.applicationTables();

    if (tables.length === 0) {
      return;
    }

    const tableReferences = tables.map((table) => `${quoteIdentifier(PUBLIC_SCHEMA)}.${quoteIdentifier(table)}`);
    await this.pool.query(`TRUNCATE TABLE ${tableReferences.join(', ')} RESTART IDENTITY CASCADE`);
  }

  /**
   * Closes all PostgreSQL connections held by this helper.
   *
   * Example: await database.close() during Playwright worker teardown.
   */
  public async close(): Promise<void> {
    await this.pool.end();
  }

  /** Discovers mutable application base tables while preserving Flyway metadata. */
  private async applicationTables(): Promise<readonly string[]> {
    const rows = await this.query<{ table_name: string }>(
      `SELECT table_name
       FROM information_schema.tables
       WHERE table_schema = $1
         AND table_type = 'BASE TABLE'
         AND NOT (table_name = ANY($2::text[]))
       ORDER BY table_name`,
      [PUBLIC_SCHEMA, FLYWAY_HISTORY_TABLES],
    );

    return rows.map((row) => row.table_name);
  }
}

/**
 * Creates an administrator-backed helper for the disposable E2E PostgreSQL service.
 *
 * The standard PG* variables keep this helper independent of host-published ports.
 * Example: const database = createE2eDatabase() in a runner attached to the db network.
 *
 * Authored by: OpenCode
 */
export function createE2eDatabase(): E2eDatabase {
  return new E2eDatabase(new Pool({
    database: requiredEnvironment('PGDATABASE'),
    host: requiredEnvironment('PGHOST'),
    password: requiredEnvironment('PGPASSWORD'),
    port: parsePort(requiredEnvironment('PGPORT')),
    user: requiredEnvironment('PGUSER'),
  }));
}

/** Safely quotes a PostgreSQL identifier discovered from the database catalog. */
function quoteIdentifier(identifier: string): string {
  return `"${identifier.replaceAll('"', '""')}"`;
}

/** Reads a required E2E runtime environment variable. */
function requiredEnvironment(name: string): string {
  const value = process.env[name];

  if (value === undefined || value.length === 0) {
    throw new Error(`Missing required E2E database environment variable: ${name}`);
  }

  return value;
}

/** Parses the PostgreSQL port while rejecting malformed environment configuration. */
function parsePort(value: string): number {
  const port = Number(value);

  if (!Number.isSafeInteger(port) || port < 1 || port > 65_535) {
    throw new Error(`Invalid E2E PostgreSQL port: ${value}`);
  }

  return port;
}
