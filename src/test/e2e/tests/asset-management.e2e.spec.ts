/**
 * Covers scenarios 8, 8.1, 8.2, and 9 for asset management through browser, API, and PostgreSQL boundaries.
 *
 * Assets are seeded directly in PostgreSQL so each scenario remains independent from other API
 * flows. The write scenarios verify that forms submit and preserve their complete external data.
 *
 * @author OpenCode
 */
import type { Locator, Page, Response } from '@playwright/test';
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';

const ORIGINAL_TICKER = 'E2E:ORIGINAL';
const ORIGINAL_NAME = 'E2E Original Asset';
const CREATED_TICKER = 'E2E:CREATED';
const CREATED_NAME = 'E2E Created Asset';
const UPDATED_TICKER = 'E2E:UPDATED';
const UPDATED_NAME = 'E2E Updated Asset';
const DRAFT_TICKER = 'E2E:DRAFT';
const DRAFT_NAME = 'E2E Draft Asset';

const PERSISTED_EXTERNAL_DATA = {
  data: [
    {
      source: 'YAHOO_FINANCE',
      ticker: 'ORIGINAL-EXTERNAL-TICKER',
      exchangeId: 'ORIGINAL-EXCHANGE',
    },
  ],
} as const;

/**
 * Represents the asset shape returned by the API in asset-management scenarios.
 * @author OpenCode
 */
type Asset = {
  id: number;
  name: string;
  ticker: string;
  externalData?: typeof PERSISTED_EXTERNAL_DATA;
};

/**
 * Represents the asset columns read directly from PostgreSQL.
 * @author OpenCode
 */
type AssetRow = {
  external_data: string | null;
  id: number;
  name: string;
  ticker: string;
};

test.describe('asset management', () => {
  test('scenario 8: lists, creates, edits, and cancels asset changes with clicks', async ({ database, page }) => {
    await page.goto('/');
    await page.getByRole('link', { name: 'Assets', exact: true }).click();
    await expectAssetTable(page);

    await page.getByRole('button', { name: 'New asset' }).click();
    await expectNewAssetForm(page);

    await page.getByRole('textbox', { name: 'Ticker' }).fill(DRAFT_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(DRAFT_NAME);

    const cancelledCreateListResponsePromise = page.waitForResponse(isAssetCollectionRequest);
    await page.getByRole('button', { name: 'Cancel' }).click();
    await expect(page).toHaveURL(/\/asset$/);
    await cancelledCreateListResponsePromise;
    await expect(page.getByText('No assets have been registered yet.')).toBeVisible();

    await page.getByRole('button', { name: 'New asset' }).click();
    await expectNewAssetForm(page);

    await page.getByRole('textbox', { name: 'Ticker' }).fill(CREATED_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(CREATED_NAME);

    const createResponsePromise = page.waitForResponse(isAssetCreationRequest);
    const createDetailResponsePromise = page.waitForResponse((response) => {
      return isAssetDetailRequest(response, undefined);
    });

    await page.getByRole('button', { name: 'Create' }).click();

    const createResponse = await createResponsePromise;
    expect(createResponse.status()).toBe(201);
    expect(createResponse.request().headers()['content-type']).toContain('application/json');
    expect(createResponse.request().postDataJSON()).toEqual({
      externalData: null,
      name: CREATED_NAME,
      ticker: CREATED_TICKER,
    });

    const createdAsset = await createResponse.json() as Asset;
    expect(createdAsset).toMatchObject({
      id: expect.any(Number),
      name: CREATED_NAME,
      ticker: CREATED_TICKER,
    });
    expect(createdAsset.id).toBeGreaterThan(0);

    await expect(page).toHaveURL(`/asset/${createdAsset.id}`);
    const createDetailResponse = await createDetailResponsePromise;
    expect(createDetailResponse.status()).toBe(200);
    expect(await createDetailResponse.json()).toEqual(createdAsset);
    await expectAssetEditor(page, createdAsset);

    await page.getByRole('textbox', { name: 'Ticker' }).fill(UPDATED_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(UPDATED_NAME);

    const updateResponsePromise = page.waitForResponse(isAssetUpdateRequest);
    await page.getByRole('button', { name: 'Save' }).click();

    const updateResponse = await updateResponsePromise;
    expect(updateResponse.status()).toBe(200);
    const updateRequest = updateResponse.request().postDataJSON() as Record<string, unknown>;
    expect(updateRequest).toMatchObject({
      name: UPDATED_NAME,
      ticker: UPDATED_TICKER,
    });
    expect(Number(updateRequest.id)).toBe(createdAsset.id);
    expect(updateRequest.externalData).toBeNull();
    await expectAssetEditor(page, { id: createdAsset.id, name: UPDATED_NAME, ticker: UPDATED_TICKER });

    await page.getByRole('textbox', { name: 'Ticker' }).fill(DRAFT_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(DRAFT_NAME);

    const listResponsePromise = page.waitForResponse(isAssetCollectionRequest);
    await page.getByRole('button', { name: 'Cancel' }).click();
    await expect(page).toHaveURL(/\/asset$/);
    await listResponsePromise;
    await expect(page.getByRole('cell', { name: UPDATED_TICKER, exact: true })).toBeVisible();
    await expect(page.getByRole('cell', { name: DRAFT_TICKER, exact: true })).toHaveCount(0);

    const persistedAssets = await database.query<AssetRow>(
      `SELECT id, name, ticker, external_data::text AS external_data
       FROM public.asset
       WHERE id = $1`,
      [createdAsset.id],
    );
    expect(persistedAssets).toEqual([{
      external_data: null,
      id: createdAsset.id,
      name: UPDATED_NAME,
      ticker: UPDATED_TICKER,
    }]);

    await page.getByRole('button', { name: 'New asset' }).click();
    await expectNewAssetForm(page);
    await page.getByRole('textbox', { name: 'Ticker' }).fill(UPDATED_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(DRAFT_NAME);

    const duplicateResponsePromise = page.waitForResponse(isAssetCreationRequest);
    await page.getByRole('button', { name: 'Create' }).click();
    expect((await duplicateResponsePromise).ok()).toBe(false);
    await expect(page.locator('.toast.text-bg-danger')).toContainText('Error');
    await expect(page.getByRole('textbox', { name: 'Ticker' })).toHaveValue(UPDATED_TICKER);
    await expect(page.getByRole('textbox', { name: 'Name' })).toHaveValue(DRAFT_NAME);
    await expect(page).toHaveURL(/\/asset\/new$/);

    await page.getByRole('button', { name: 'Cancel' }).click();
    await expect(page.getByRole('cell', { name: DRAFT_NAME, exact: true })).toHaveCount(0);
  });

  test('scenario 8.1: opens asset pages through direct browser URLs', async ({ database, page }) => {
    const asset = await seedAsset(database, ORIGINAL_TICKER, ORIGINAL_NAME);
    const otherAsset = await seedAsset(database, CREATED_TICKER, CREATED_NAME);
    const newAssetDetailRequests: string[] = [];
    page.on('request', (request) => {
      if(new URL(request.url()).pathname === '/api/asset/new') {
        newAssetDetailRequests.push(request.url());
      }
    });

    await page.goto('/asset');
    await expectAssetTable(page, asset);
    await page.reload();
    await expectAssetTable(page, asset);

    await page.goto('/asset/new');
    await expectNewAssetForm(page);
    await page.reload();
    await expectNewAssetForm(page);
    expect(newAssetDetailRequests).toHaveLength(0);

    await page.goto(`/asset/${asset.id}`);
    await expectAssetEditor(page, asset);
    await page.reload();
    await expectAssetEditor(page, asset);

    await navigateWithinApp(page, '/asset/new');
    await expectNewAssetForm(page);
    await page.goBack();
    await expectAssetEditor(page, asset);
    await page.goForward();
    await expectNewAssetForm(page);

    await navigateWithinApp(page, `/asset/${asset.id}`);
    await expectAssetEditor(page, asset);

    await navigateWithinApp(page, `/asset/${otherAsset.id}`);
    await expectAssetEditor(page, otherAsset);
    await page.goBack();
    await expectAssetEditor(page, asset);

    await page.route(`**/api/asset/${asset.id}`, async (route) => {
      const response = await route.fetch();
      const loadedAsset = await response.json() as Asset;
      await route.fulfill({ response, json: { ...loadedAsset, id: asset.id + 100 } });
    });
    await page.reload();
    await expect(page.getByRole('heading', { name: 'Asset could not be loaded' })).toBeVisible();
    await expect(page.getByRole('button', { name: 'Save', exact: true })).toHaveCount(0);
    await expectIconOnlyButton(
      page.locator('#asset-edit-error').getByRole('button', { name: 'Back to assets', exact: true }),
      'arrow-left',
    );
    expect(newAssetDetailRequests).toHaveLength(0);

    await expectPersistedAsset(database, asset, null);
    await expectPersistedAsset(database, otherAsset, null);
  });

  test('scenario 8.2: opens an asset once by keyboard without Space scrolling', async ({ database, page }) => {
    const asset = await seedAsset(database, ORIGINAL_TICKER, ORIGINAL_NAME);
    await database.query(
      `INSERT INTO public.asset (ticker, name)
       SELECT 'E2E:ZZ-FILLER:' || series_number, 'E2E filler asset ' || series_number
       FROM generate_series(1, 20) AS series_number`,
    );
    const detailRequests: string[] = [];
    page.on('request', (request) => {
      if(new URL(request.url()).pathname === `/api/asset/${asset.id}`) {
        detailRequests.push(request.url());
      }
    });

    await page.goto('/asset');
    await expectAssetTable(page, asset);

    await page.getByRole('link', { name: ORIGINAL_TICKER, exact: true }).press('Enter');
    await expectAssetEditor(page, asset);
    expect(detailRequests).toHaveLength(1);

    await navigateWithinApp(page, '/asset');
    await expectAssetTable(page, asset);

    const assetTableScroll = page.locator('.asset-table-scroll');
    const assetRow = page.getByRole('row').filter({
      has: page.getByRole('link', { name: ORIGINAL_TICKER, exact: true }),
    });
    await expect(assetTableScroll).toBeVisible();
    expect(await assetTableScroll.evaluate((element) => element.scrollHeight)).toBeGreaterThan(
      await assetTableScroll.evaluate((element) => element.clientHeight),
    );
    await assetTableScroll.evaluate((element) => {
      element.scrollTop = 0;
    });
    await assetRow.focus();
    const initialScrollTop = await assetTableScroll.evaluate((element) => element.scrollTop);
    const spaceDefaultPreventedPromise = page.evaluate(() => new Promise<boolean>((resolve) => {
      document.addEventListener('keypress', (event) => {
        if(event.key === ' ') {
          resolve(event.defaultPrevented);
        }
      }, { once: true });
    }));

    await assetRow.press('Space');

    expect(await spaceDefaultPreventedPromise).toBe(true);
    await expectAssetEditor(page, asset);
    expect(await assetTableScroll.evaluate((element) => element.scrollTop)).toBe(initialScrollTop);
    expect(detailRequests).toHaveLength(2);
  });

  test('scenario 9: submits and preserves loaded external data when saving asset fields', async ({ database, page }) => {
    const asset = await seedAsset(database, ORIGINAL_TICKER, ORIGINAL_NAME, PERSISTED_EXTERNAL_DATA);
    await seedAsset(database, CREATED_TICKER, CREATED_NAME);

    await page.goto(`/asset/${asset.id}`);
    await expectAssetEditor(page, { ...asset, externalData: PERSISTED_EXTERNAL_DATA });

    await page.getByRole('textbox', { name: 'Ticker' }).fill(CREATED_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(DRAFT_NAME);

    const duplicateUpdatePromise = page.waitForResponse(isAssetUpdateRequest);
    await page.getByRole('button', { name: 'Save' }).click();
    expect((await duplicateUpdatePromise).ok()).toBe(false);
    await expect(page.locator('.toast.text-bg-danger')).toContainText('Error');
    await expect(page.getByRole('textbox', { name: 'Ticker' })).toHaveValue(CREATED_TICKER);
    await expect(page.getByRole('textbox', { name: 'Name' })).toHaveValue(DRAFT_NAME);
    await expectPersistedAsset(database, asset, PERSISTED_EXTERNAL_DATA);

    await page.getByRole('textbox', { name: 'Ticker' }).fill(UPDATED_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(UPDATED_NAME);

    const updateResponsePromise = page.waitForResponse(isAssetUpdateRequest);
    await page.getByRole('button', { name: 'Save' }).click();

    const updateResponse = await updateResponsePromise;
    expect(updateResponse.status()).toBe(200);
    const updateRequest = updateResponse.request().postDataJSON() as Record<string, unknown>;
    expect(updateRequest).toMatchObject({
      name: UPDATED_NAME,
      ticker: UPDATED_TICKER,
    });
    expect(Number(updateRequest.id)).toBe(asset.id);
    expect(updateRequest.externalData).toEqual(PERSISTED_EXTERNAL_DATA);

    const persistedAssets = await database.query<AssetRow>(
      `SELECT id, name, ticker, external_data::text AS external_data
       FROM public.asset
       WHERE id = $1`,
      [asset.id],
    );
    expect(persistedAssets).toHaveLength(1);
    expect(persistedAssets[0]).toMatchObject({
      id: asset.id,
      name: UPDATED_NAME,
      ticker: UPDATED_TICKER,
    });
    expect(JSON.parse(persistedAssets[0].external_data ?? 'null')).toEqual(PERSISTED_EXTERNAL_DATA);

    await page.getByRole('button', { name: 'Cancel' }).click();
    await expect(page).toHaveURL(/\/asset$/);
    await expect(page.getByRole('cell', { name: UPDATED_TICKER, exact: true })).toBeVisible();
  });
});

/**
 * Navigates within the mounted SPA to exercise route handlers without reloading the document.
 * @author OpenCode
 */
async function navigateWithinApp(page: Page, path: string): Promise<void> {
  await page.evaluate((destination) => {
    (window as unknown as { navigateTo: (path: string) => void }).navigateTo(destination);
  }, path);
}

/**
 * Seeds one asset directly in PostgreSQL for a browser scenario.
 * @author OpenCode
 */
async function seedAsset(
  database: E2eDatabase,
  ticker: string,
  name: string,
  externalData?: object,
): Promise<{ id: number; name: string; ticker: string }> {
  const rows = await database.query<{ id: number; name: string; ticker: string }>(
    `INSERT INTO public.asset (ticker, name, external_data)
     VALUES ($1, $2, $3::jsonb)
     RETURNING id, name, ticker`,
    [ticker, name, externalData ? JSON.stringify(externalData) : null],
  );

  expect(rows).toHaveLength(1);
  return rows[0];
}

/**
 * Asserts the list route, basic columns, and expected asset or empty state.
 * @author OpenCode
 */
async function expectAssetTable(page: Page, asset?: Asset): Promise<void> {
  await expect(page).toHaveURL(/\/asset$/);
  await expect(page.getByRole('link', { name: 'Assets', exact: true })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Assets', exact: true })).toBeVisible();
  await expect(page.getByRole('columnheader', { name: 'Ticker', exact: true })).toBeVisible();
  await expect(page.getByRole('columnheader', { name: 'Name', exact: true })).toBeVisible();

  if(asset) {
    await expect(page.getByRole('cell', { name: asset.ticker, exact: true })).toBeVisible();
    await expect(page.getByRole('cell', { name: asset.name, exact: true })).toBeVisible();
  } else {
    await expect(page.getByText('No assets have been registered yet.')).toBeVisible();
  }
}

/**
 * Asserts the shared form is configured for creating an asset.
 * @author OpenCode
 */
async function expectNewAssetForm(page: Page): Promise<void> {
  await expect(page).toHaveURL(/\/asset\/new$/);
  await expect(page.getByRole('heading', { name: 'New asset', exact: true })).toBeVisible();
  await expect(page.getByRole('textbox', { name: 'Ticker' })).toHaveValue('');
  await expect(page.getByRole('textbox', { name: 'Name' })).toHaveValue('');
  await expectIconOnlyButton(page.getByRole('button', { name: 'Back to assets', exact: true }), 'arrow-left');
  await expectIconOnlyButton(page.getByRole('button', { name: 'Cancel', exact: true }), 'x-circle');
  await expectIconOnlyButton(page.getByRole('button', { name: 'Create', exact: true }), 'plus-circle');
}

/**
 * Asserts the edit route contains the expected basic asset fields.
 * @author OpenCode
 */
async function expectAssetEditor(page: Page, asset: Asset): Promise<void> {
  await expect(page).toHaveURL(`/asset/${asset.id}`);
  await expect(page.getByRole('heading', { name: 'Edit asset', exact: true })).toBeVisible();
  await expect(page.getByRole('textbox', { name: 'Ticker' })).toHaveValue(asset.ticker);
  await expect(page.getByRole('textbox', { name: 'Name' })).toHaveValue(asset.name);
  await expectIconOnlyButton(page.getByRole('button', { name: 'Back to assets', exact: true }), 'arrow-left');
  await expectIconOnlyButton(page.getByRole('button', { name: 'Cancel', exact: true }), 'x-circle');
  await expectIconOnlyButton(page.getByRole('button', { name: 'Save', exact: true }), 'save-fill');
}

/**
 * Verifies an action button has an accessible label and icon without visible text.
 * @author OpenCode
 */
async function expectIconOnlyButton(button: Locator, iconName: string): Promise<void> {
  await expect(button).toBeVisible();
  await expect(button).toHaveText('');
  await expect(button).toHaveAttribute('title', /.+/);
  await expect(button.locator(`span.bi-${iconName}`)).toHaveAttribute('aria-hidden', 'true');
}

/**
 * Asserts one asset's complete persisted state directly in PostgreSQL.
 * @author OpenCode
 */
async function expectPersistedAsset(
  database: E2eDatabase,
  asset: Asset,
  externalData: object | null,
): Promise<void> {
  const rows = await database.query<AssetRow>(
    `SELECT id, name, ticker, external_data::text AS external_data
     FROM public.asset
     WHERE id = $1`,
    [asset.id],
  );

  expect(rows).toHaveLength(1);
  expect(rows[0]).toMatchObject({
    id: asset.id,
    name: asset.name,
    ticker: asset.ticker,
  });
  expect(JSON.parse(rows[0].external_data ?? 'null')).toEqual(externalData);
}

/**
 * Matches the same-origin request that loads the asset collection.
 * @author OpenCode
 */
function isAssetCollectionRequest(response: Response): boolean {
  return response.request().method() === 'GET' && new URL(response.url()).pathname === '/api/asset';
}

/**
 * Matches the same-origin request that creates one asset.
 * @author OpenCode
 */
function isAssetCreationRequest(response: Response): boolean {
  return response.request().method() === 'POST' && new URL(response.url()).pathname === '/api/asset';
}

/**
 * Matches the same-origin request that updates one asset.
 * @author OpenCode
 */
function isAssetUpdateRequest(response: Response): boolean {
  return response.request().method() === 'PUT' && new URL(response.url()).pathname === '/api/asset';
}

/**
 * Matches an asset detail request, optionally constrained to one identifier.
 * @author OpenCode
 */
function isAssetDetailRequest(response: Response, identifier: number | undefined): boolean {
  const request = response.request();
  const pathname = new URL(response.url()).pathname;
  const matchesPath = /^\/api\/asset\/[^/]+$/.test(pathname);
  const matchesIdentifier = identifier === undefined || pathname === `/api/asset/${identifier}`;
  return request.method() === 'GET' && matchesPath && matchesIdentifier;
}
