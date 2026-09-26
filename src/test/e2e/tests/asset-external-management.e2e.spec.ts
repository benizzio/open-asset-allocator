/**
 * Covers scenarios 10–13 for external-asset search and ordered draft persistence in both asset forms.
 * Provider responses are mocked while asset writes are checked directly in PostgreSQL.
 *
 * @author OpenCode
 */
import { expect, test } from '../support/fixtures';
import type { Page } from '@playwright/test';
import type { E2eDatabase } from '../support/database';

const FIRST = { source: 'YAHOO_FINANCE', ticker: 'IAU', exchangeId: 'PCX' };
const SECOND = { source: 'YAHOO_FINANCE', ticker: 'BIL', exchangeId: 'NMS' };

test('scenario 10: searches, adds, reorders, removes and persists external assets', async ({ page, database }) => {
  const selectorErrors: string[] = [];
  let searchRequests = 0;
  page.on('console', (message) => {
    if(message.type() === 'error' && message.text().includes('hx-disabled-elt')) {
      selectorErrors.push(message.text());
    }
  });
  await page.route('**/api/external-asset?*', async (route) => {
    searchRequests++;
    const query = new URL(route.request().url()).searchParams.get('query');
    await route.fulfill({ json: query === 'empty' ? [] : [
      { ...FIRST, name: 'Gold Trust', exchangeName: 'NYSE Arca' },
      { ...SECOND, name: 'Treasury Bill', exchangeName: 'NASDAQ' },
      ...Array.from({ length: 5 }, (_, index) => ({
        ...FIRST, ticker: `GOLD-${index + 1}`, name: `Gold asset ${index + 1}`, exchangeName: 'NYSE Arca',
      })),
    ] });
  });

  await page.goto('/asset/new');
  const query = page.getByRole('searchbox', { name: 'Search external assets' });
  await query.fill('   ');
  await page.getByRole('button', { name: 'Search external assets' }).click();
  const invalidSearchToast = page.locator('.toast.text-bg-warning').filter({ hasText: 'Enter a search term' });
  await expect(invalidSearchToast).toBeVisible();
  expect(searchRequests).toBe(0);
  await query.evaluate((element) => { (element as HTMLInputElement).value = 'x'.repeat(101); });
  await query.press('Enter');
  await expect(page.locator('.toast.text-bg-warning').filter({ hasText: '1 to 100 characters' }).last()).toBeVisible();
  expect(searchRequests).toBe(0);

  await query.fill('gold');
  const searchResponse = page.waitForResponse(response => response.url().includes('/api/external-asset?'));
  await query.press('Enter');
  const searched = await searchResponse;
  expect(searched.status()).toBe(200);
  expect(await searched.json()).toHaveLength(7);
  await expect(page.getByRole('heading', { name: 'New asset' })).toBeVisible();
  await expect(page.getByRole('table', { name: 'External asset search results' }).locator('tbody tr')).toHaveCount(7);
  await expect(page.getByText('The server returned an invalid asset.')).toHaveCount(0);
  expect(selectorErrors).toEqual([]);
  await page.getByRole('button', { name: 'Add IAU from YAHOO_FINANCE on PCX' }).click();
  await expect(query).toHaveValue('');
  await expect(page.getByRole('table', { name: 'External asset search results' })).toHaveCount(0);

  await query.fill('gold');
  const duplicateSearch = page.waitForResponse(response => response.url().includes('/api/external-asset?'));
  await query.press('Enter');
  expect((await duplicateSearch).status()).toBe(200);
  await page.getByRole('button', { name: 'Add IAU from YAHOO_FINANCE on PCX' }).click();
  const duplicateToast = page.locator('.toast.text-bg-warning').filter({ hasText: 'already added' });
  await expect(duplicateToast).toBeVisible();
  await expect(page.getByRole('table', { name: 'External asset search results' }).locator('tbody tr')).toHaveCount(7);

  const registered = page.getByRole('table', { name: 'Registered external assets in priority order' });
  await expect(registered.getByRole('row')).toHaveCount(2);
  await page.getByRole('button', { name: 'Add BIL from YAHOO_FINANCE on NMS' }).click();
  await expect(query).toHaveValue('');
  await expect(page.getByRole('table', { name: 'External asset search results' })).toHaveCount(0);
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
  const selectorErrors: string[] = [];
  let searchRequests = 0;
  page.on('console', (message) => {
    if(message.type() === 'error' && message.text().includes('hx-disabled-elt')) {
      selectorErrors.push(message.text());
    }
  });
  const rows = await database.query<{ id: number }>(
    `INSERT INTO public.asset (ticker, name, external_data)
     VALUES ('EXTERNAL-EDIT', 'External Edit', $1::jsonb) RETURNING id`,
    [JSON.stringify({ data: [FIRST, SECOND] })],
  );
  const id = rows[0].id;
  await database.query("INSERT INTO public.asset (ticker, name) VALUES ('TAKEN', 'Taken')");
  await page.route('**/api/external-asset?*', async (route) => {
    searchRequests++;
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
  const query = page.getByRole('searchbox', { name: 'Search external assets' });
  await query.fill('   ');
  await query.press('Enter');
  const invalidSearchToast = page.locator('.toast.text-bg-warning').filter({ hasText: 'Enter a search term' });
  await expect(invalidSearchToast).toBeVisible();
  expect(searchRequests).toBe(0);

  await query.fill('gold');
  await page.getByRole('button', { name: 'Search external assets' }).click();
  const results = page.getByRole('table', { name: 'External asset search results' });
  await expect(results.locator('tbody tr')).toHaveCount(10);
  await expect(page.getByText('The server returned an invalid asset.')).toHaveCount(0);
  expect(selectorErrors).toEqual([]);
  await results.getByRole('button', { name: 'Add IAU from YAHOO_FINANCE on PCX' }).click();
  await expect(page.locator('.toast.text-bg-warning').filter({ hasText: 'already added' })).toBeVisible();
  await expect(results.locator('tbody tr')).toHaveCount(10);
  await results.getByRole('button', { name: 'Add EXTERNAL-1 from YAHOO_FINANCE on PCX' }).click();
  await expect(query).toHaveValue('');
  await expect(page.getByRole('table', { name: 'External asset search results' })).toHaveCount(0);
  await expect(registered.locator('tbody tr')).toHaveCount(3);
  await registered.getByRole('button', { name: 'Move BIL up' }).click();

  await page.getByRole('textbox', { name: 'Ticker', exact: true }).fill('TAKEN');
  const failedSave = page.waitForResponse(response => response.request().method() === 'PUT'
    && new URL(response.url()).pathname === '/api/asset');
  await page.getByRole('button', { name: 'Save', exact: true }).click();
  expect((await failedSave).ok()).toBe(false);
  await expect(registered.locator('tbody tr').first()).toContainText('BIL');
  await expectPersistedOrder(database, id, [FIRST, SECOND]);

  await query.fill('empty');
  await page.getByRole('button', { name: 'Search external assets' }).click();
  await expect(page.locator('.toast.text-bg-primary').filter({ hasText: 'No external assets found.' })).toBeVisible();
  await expect(page.locator('[data-external-message]')).toHaveText('');
  await query.fill('failure');
  await page.getByRole('button', { name: 'Search external assets' }).click();
  await expect(page.locator('.toast.text-bg-danger').filter({ hasText: 'could not be searched' })).toBeVisible();
  await expect(page.locator('[data-external-message]')).toHaveText('');

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

test('scenario 13: fills only empty asset fields when an external result is added', async ({ page, database }) => {
  const rows = await database.query<{ id: number }>(
    "INSERT INTO public.asset (ticker, name) VALUES ('AUTOFILL-EDIT', 'Autofill Edit') RETURNING id",
  );
  const editAssetId = rows[0].id;
  const additionalResult = { source: 'YAHOO_FINANCE', ticker: 'GLD', exchangeId: 'PCX' };

  await page.route('**/api/external-asset?*', async (route) => {
    await route.fulfill({ json: [
      { ...FIRST, name: 'Gold Trust', exchangeName: 'NYSE Arca' },
      { ...SECOND, name: 'Treasury Bill', exchangeName: 'NASDAQ' },
      { ...additionalResult, name: 'SPDR Gold Shares', exchangeName: 'NYSE Arca' },
    ] });
  });

  for(const view of [{ path: '/asset/new', isEdit: false }, { path: `/asset/${editAssetId}`, isEdit: true }]) {
    await page.goto(view.path);
    const ticker = page.getByRole('textbox', { name: 'Ticker', exact: true });
    const name = page.getByRole('textbox', { name: 'Name', exact: true });

    if(view.isEdit) {
      await expect(ticker).toBeVisible();
      await expect(page.locator('#asset-edit-content')).not.toHaveClass(/\bhtmx-settling\b/);
      await ticker.fill('');
      await name.fill('');
    }

    await searchAndAdd(page, FIRST.ticker, FIRST.exchangeId);
    await expect(ticker).toHaveValue('PCX:IAU');
    await expect(name).toHaveValue('Gold Trust');

    await name.fill('Keep this name');
    await ticker.fill('');
    await searchAndAdd(page, SECOND.ticker, SECOND.exchangeId);
    await expect(ticker).toHaveValue('NMS:BIL');
    await expect(name).toHaveValue('Keep this name');

    await ticker.fill('Keep this ticker');
    await name.fill('');
    await searchAndAdd(page, additionalResult.ticker, additionalResult.exchangeId);
    await expect(ticker).toHaveValue('Keep this ticker');
    await expect(name).toHaveValue('SPDR Gold Shares');
  }

  await page.getByRole('button', { name: 'Cancel', exact: true }).click();
  const persisted = await database.query<{ ticker: string; name: string; external_data: string | null }>(
    'SELECT ticker, name, external_data::text AS external_data FROM public.asset WHERE id = $1', [editAssetId],
  );
  expect(persisted).toEqual([{ ticker: 'AUTOFILL-EDIT', name: 'Autofill Edit', external_data: null }]);
});

/**
 * Searches for and adds one mocked provider result to the asset form.
 * @author OpenCode
 */
async function searchAndAdd(page: Page, ticker: string, exchangeId: string): Promise<void> {
  const query = page.getByRole('searchbox', { name: 'Search external assets' });

  await query.fill('gold');
  await query.press('Enter');
  await page.getByRole('button', {
    name: `Add ${ticker} from YAHOO_FINANCE on ${exchangeId}`,
  }).click();
}

/**
 * Asserts the identifiers and priority order stored by the asset write.
 * @author OpenCode
 */
async function expectPersistedOrder(database: E2eDatabase, assetId: number, expected: typeof FIRST[] | null) {
  const rows = await database.query<{ external_data: string | null }>(
    'SELECT external_data::text AS external_data FROM public.asset WHERE id = $1', [assetId],
  );
  expect(JSON.parse(rows[0].external_data ?? 'null')).toEqual(expected === null ? null : { data: expected });
}
