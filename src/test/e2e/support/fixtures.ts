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
 * Example: test('creates a portfolio', async ({ page, database }) => { ... }).
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

export { expect };
