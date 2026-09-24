/**
 * Covers scenarios 10–12 for external-asset search and ordered draft persistence in both asset forms.
 * Provider responses are mocked while asset writes are checked directly in PostgreSQL.
 *
 * Authored by: OpenCode
 */
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';

const FIRST = { source: 'YAHOO_FINANCE', ticker: 'IAU', exchangeId: 'PCX' };
const SECOND = { source: 'YAHOO_FINANCE', ticker: 'BIL', exchangeId: 'NMS' };

test('scenario 10: searches, adds, reorders, removes and persists external assets', async ({ page, database }) => {
  await page.route('**/api/external-asset?*', async (route) => {
    const query = new URL(route.request().url()).searchParams.get('query');
    await route.fulfill({ json: query === 'empty' ? [] : [
      { ...FIRST, name: 'Gold Trust', exchangeName: 'NYSE Arca' },
      { ...SECOND, name: 'Treasury Bill', exchangeName: 'NASDAQ' },
    ] });
  });

  await page.goto('/asset/new');
  const query = page.getByRole('searchbox', { name: 'Search external assets' });
  await query.fill('gold');
  const searchResponse = page.waitForResponse(response => response.url().includes('/api/external-asset?'));
  await query.press('Enter');
  expect((await searchResponse).status()).toBe(200);
  await expect(page.getByRole('heading', { name: 'New asset' })).toBeVisible();
  await expect(page.getByRole('table', { name: 'External asset search results' }).getByRole('row')).toHaveCount(3);
  await page.getByRole('button', { name: 'Add IAU from YAHOO_FINANCE on PCX' }).click();
  await page.getByRole('button', { name: 'Add IAU from YAHOO_FINANCE on PCX' }).click();
  await expect(page.locator('[data-external-message]')).toContainText('already added');

  const registered = page.getByRole('table', { name: 'Registered external assets in priority order' });
  await expect(registered.getByRole('row')).toHaveCount(2);
  await page.getByRole('button', { name: 'Add BIL from YAHOO_FINANCE on NMS' }).click();
  await registered.getByRole('button', { name: 'Move BIL up' }).click();
  await expect(registered.locator('tbody tr').first()).toContainText('BIL');
  await expect(registered.getByRole('button', { name: 'Move BIL up' })).toBeDisabled();

  await page.getByRole('textbox', { name: 'Ticker', exact: true }).fill('EXTERNAL-TEST');
  await page.getByRole('textbox', { name: 'Name', exact: true }).fill('External Test');
  const create = page.waitForResponse(response => response.request().method() === 'POST'
    && new URL(response.url()).pathname === '/api/asset');
  await page.getByRole('button', { name: 'Create', exact: true }).click();
  const created = await create;
  expect(created.status()).toBe(201);
  expect(created.request().postDataJSON().externalData).toEqual({ data: [SECOND, FIRST] });
  const id = (await created.json() as { id: number }).id;
  await expect(page).toHaveURL(`/asset/${id}`);
  await expect(registered.locator('tbody tr').first()).toContainText('BIL');
  await expectPersistedOrder(database, id, [SECOND, FIRST]);

  await registered.getByRole('button', { name: 'Remove BIL' }).click();
  await registered.getByRole('button', { name: 'Remove IAU' }).click();
  await expect(page.getByText('No external assets added.')).toBeVisible();
  const update = page.waitForResponse(response => response.request().method() === 'PUT'
    && new URL(response.url()).pathname === '/api/asset');
  await page.getByRole('button', { name: 'Save', exact: true }).click();
  const updated = await update;
  expect(updated.status()).toBe(200);
  expect(updated.request().postDataJSON().externalData).toBeNull();
  await expectPersistedOrder(database, id, null);
});

test('scenario 11: limits search results and keeps edit drafts until save or cancel', async ({ page, database }) => {
  const rows = await database.query<{ id: number }>(
    `INSERT INTO public.asset (ticker, name, external_data)
     VALUES ('EXTERNAL-EDIT', 'External Edit', $1::jsonb) RETURNING id`,
    [JSON.stringify({ data: [FIRST, SECOND] })],
  );
  const id = rows[0].id;
  await database.query("INSERT INTO public.asset (ticker, name) VALUES ('TAKEN', 'Taken')");
  await page.route('**/api/external-asset?*', async (route) => {
    const query = new URL(route.request().url()).searchParams.get('query');
    if(query === 'failure') {
      await route.fulfill({ status: 500, json: { errorMessage: 'Search failed' } });
    } else {
      await route.fulfill({ json: query === 'empty' ? [] : Array.from({ length: 12 }, (_, index) => ({
        ...FIRST, ticker: index === 0 ? FIRST.ticker : `EXTERNAL-${index}`,
      })) });
    }
  });

  await page.goto(`/asset/${id}`);
  const registered = page.getByRole('table', { name: 'Registered external assets in priority order' });
  await expect(registered.locator('tbody tr')).toHaveCount(2);
  await page.getByRole('searchbox', { name: 'Search external assets' }).fill('gold');
  await page.getByRole('button', { name: 'Search external assets' }).click();
  const results = page.getByRole('table', { name: 'External asset search results' });
  await expect(results.locator('tbody tr')).toHaveCount(10);
  await results.getByRole('button', { name: 'Add IAU from YAHOO_FINANCE on PCX' }).click();
  await expect(page.locator('[data-external-message]')).toContainText('already added');
  await registered.getByRole('button', { name: 'Move BIL up' }).click();

  await page.getByRole('textbox', { name: 'Ticker', exact: true }).fill('TAKEN');
  const failedSave = page.waitForResponse(response => response.request().method() === 'PUT'
    && new URL(response.url()).pathname === '/api/asset');
  await page.getByRole('button', { name: 'Save', exact: true }).click();
  expect((await failedSave).ok()).toBe(false);
  await expect(registered.locator('tbody tr').first()).toContainText('BIL');
  await expectPersistedOrder(database, id, [FIRST, SECOND]);

  await page.getByRole('searchbox', { name: 'Search external assets' }).fill('empty');
  await page.getByRole('button', { name: 'Search external assets' }).click();
  await expect(page.locator('[data-external-message]')).toContainText('No external assets found.');
  await page.getByRole('searchbox', { name: 'Search external assets' }).fill('failure');
  await page.getByRole('button', { name: 'Search external assets' }).click();
  await expect(page.locator('[data-external-message]')).toContainText('could not be searched');

  await page.getByRole('button', { name: 'Cancel', exact: true }).click();
  await page.goto(`/asset/${id}`);
  await expect(registered.locator('tbody tr').first()).toContainText('IAU');
});

test('scenario 12: ignores an older search response when a newer query is pending', async ({ page }) => {
  let releaseSlow: () => void = () => {};
  const slowGate = new Promise<void>(resolve => { releaseSlow = resolve; });
  await page.route('**/api/external-asset?*', async (route) => {
    const query = new URL(route.request().url()).searchParams.get('query');
    if(query === 'slow') {
      await slowGate;
    }
    await route.fulfill({ json: [{ ...FIRST, ticker: query === 'slow' ? 'OLD' : 'NEW' }] });
  });

  await page.goto('/asset/new');
  const query = page.getByRole('searchbox', { name: 'Search external assets' });
  await query.fill('slow');
  const slowRequest = page.waitForRequest(request => new URL(request.url()).searchParams.get('query') === 'slow');
  await query.press('Enter');
  await slowRequest;
  await query.fill('new');
  await query.press('Enter');
  releaseSlow();

  const results = page.getByRole('table', { name: 'External asset search results' });
  await expect(results).toContainText('NEW');
  await expect(results).not.toContainText('OLD');
});

/** Asserts exactly the identifiers and priority order stored by the asset write. Authored by: OpenCode. */
async function expectPersistedOrder(database: E2eDatabase, assetId: number, expected: typeof FIRST[] | null) {
  const rows = await database.query<{ external_data: string | null }>(
    'SELECT external_data::text AS external_data FROM public.asset WHERE id = $1', [assetId],
  );
  expect(JSON.parse(rows[0].external_data ?? 'null')).toEqual(expected === null ? null : { data: expected });
}
