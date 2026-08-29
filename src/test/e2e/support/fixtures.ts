/**
 * Defines Playwright fixtures that isolate every E2E test attempt with PostgreSQL.
 *
 * Import test and expect from this module instead of @playwright/test. Each worker
 * waits for all services, resets data before an attempt, validates cleanup after it,
 * and closes its PostgreSQL pool before exit.
 *
 * Authored by: OpenCode
 */
import { expect, test as base } from '@playwright/test';
import { createE2eDatabase, type E2eDatabase } from './database';
import { waitForE2eReadiness } from './readiness';

type TestFixtures = {
  resetDatabase: void;
};

type WorkerFixtures = {
  database: E2eDatabase;
};

/**
 * Runs E2E tests with worker readiness and test-attempt database isolation.
 *
 * Import this fixture instead of `test` from `@playwright/test`. `BASE_URL`
 * selects the application endpoint and defaults to `http://frontend:8000`.
 * PostgreSQL access requires `PGHOST`, `PGPORT`, `PGDATABASE`, `PGUSER`, and
 * `PGPASSWORD`.
 *
 * Name `database` in a test callback to inject the worker-scoped database
 * fixture. Before the worker serves tests, it waits for the application, API,
 * and PostgreSQL to become ready. Its connection pool closes at worker teardown.
 * The automatic reset fixture resets the disposable database before every test
 * attempt. After each attempt, including a failed attempt, it resets the database
 * in a `finally` block and verifies that no portfolios remain.
 *
 * @example
 * ```ts
 * import { expect, test } from '../support/fixtures';
 *
 * test('creates a portfolio', async ({ page, database }) => {
 *   await page.goto('/');
 *   const portfolios = await database.query('SELECT * FROM public.portfolio');
 *   expect(portfolios).toHaveLength(0);
 * });
 * ```
 *
 * Authored by: OpenCode
 */
export const test = base.extend<TestFixtures, WorkerFixtures>({
  database: [async ({}, use) => {
    const database = createE2eDatabase();

    try {
      await waitForE2eReadiness(process.env.BASE_URL ?? 'http://frontend:8000', database);
      await use(database);
    } finally {
      await database.close();
    }
  }, { scope: 'worker' }],
  resetDatabase: [async ({ database }, use) => {
    await database.reset();

    try {
      await use();
    } finally {
      await database.reset();
      const portfolios = await database.query<{ count: string }>('SELECT COUNT(*)::text AS count FROM public.portfolio');
      expect(portfolios).toEqual([{ count: '0' }]);
    }
  }, { auto: true }],
});

/**
 * Exposes Playwright's standard assertion API for E2E tests and fixture cleanup
 * assertions. Import it with `test` from this module rather than directly from
 * `@playwright/test` so each test uses the database-isolated fixture API.
 *
 * @example
 * ```ts
 * import { expect, test } from '../support/fixtures';
 *
 * test('shows the portfolio page', async ({ page }) => {
 *   await page.goto('/');
 *   await expect(page).toHaveURL(/portfolio/);
 * });
 * ```
 *
 * Authored by: OpenCode
 */
export { expect };
