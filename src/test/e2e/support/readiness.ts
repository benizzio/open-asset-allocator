/**
 * Waits for the E2E application and database services to become usable.
 *
 * Use waitForE2eReadiness(baseUrl, database) once per Playwright worker after
 * Compose starts its services. It validates static content, the same-origin API,
 * and a direct PostgreSQL connection.
 *
 * Authored by: OpenCode
 */
import type { E2eDatabase } from './database';

const READINESS_TIMEOUT_MS = 30_000;
const POLL_INTERVAL_MS = 250;

/**
 * Waits until application static content, API access, and PostgreSQL are ready.
 *
 * Example: await waitForE2eReadiness('http://monolith:8080', database).
 */
export async function waitForE2eReadiness(baseUrl: string, database: E2eDatabase): Promise<void> {
  await waitFor('application static content', async () => hasSuccessfulResponse(new URL('/', baseUrl)));
  await waitFor('application API', async () => hasSuccessfulResponse(new URL('/api/portfolio', baseUrl)));
  await waitFor('PostgreSQL', async () => {
    await database.query('SELECT 1');
    return true;
  });
}

/** Polls a readiness condition with a bounded timeout and actionable failure detail. */
async function waitFor(description: string, condition: () => Promise<boolean>): Promise<void> {
  const deadline = Date.now() + READINESS_TIMEOUT_MS;
  let lastError: unknown;

  while (Date.now() < deadline) {
    try {
      if (await condition()) {
        return;
      }
    } catch (error) {
      lastError = error;
    }

    await delay(POLL_INTERVAL_MS);
  }

  const detail = lastError instanceof Error ? `: ${lastError.message}` : '';
  throw new Error(`Timed out waiting for ${description}${detail}`);
}

/** Returns whether an application endpoint responds with a successful status. */
async function hasSuccessfulResponse(url: URL): Promise<boolean> {
  const response = await fetch(url, { signal: AbortSignal.timeout(POLL_INTERVAL_MS) });
  return response.ok;
}

/** Delays the next readiness-poll attempt. */
function delay(milliseconds: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}
