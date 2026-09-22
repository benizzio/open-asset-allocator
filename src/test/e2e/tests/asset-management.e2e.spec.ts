/**
 * Covers the first assets page version through browser, HTTP API, and PostgreSQL boundaries.
 *
 * Assets are seeded directly in PostgreSQL so each scenario remains independent from other API
 * flows. The update scenario verifies that deferred external-data editing does not clear the
 * existing persisted payload.
 *
 * Authored by: OpenCode
 */
import type { Response } from '@playwright/test';
import { expect, test } from '../support/fixtures';

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

type Asset = {
  id: number;
  name: string;
  ticker: string;
  externalData?: typeof PERSISTED_EXTERNAL_DATA;
};

type AssetRow = {
  external_data: string | null;
  id: number;
  name: string;
  ticker: string;
};

test.describe('asset management', () => {
  test('lists, creates, edits, and cancels asset changes', async ({ database, page }) => {
    await page.goto('/assets');
    await expectAssetTable(page);

    await page.getByRole('button', { name: 'New asset' }).click();
    await expect(page).toHaveURL(/\/assets\/new$/);

    await page.getByRole('textbox', { name: 'Ticker' }).fill(DRAFT_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(DRAFT_NAME);

    const cancelledCreateListResponsePromise = page.waitForResponse(isAssetCollectionRequest);
    await page.getByRole('button', { name: 'Cancel' }).click();
    await expect(page).toHaveURL(/\/assets$/);
    await cancelledCreateListResponsePromise;
    await expect(page.getByText('No assets have been registered yet.')).toBeVisible();

    await page.getByRole('button', { name: 'New asset' }).click();
    await expect(page).toHaveURL(/\/assets\/new$/);

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
    expect(updateRequest).not.toHaveProperty('externalData');
    await expectAssetEditor(page, { id: createdAsset.id, name: UPDATED_NAME, ticker: UPDATED_TICKER });

    await page.getByRole('textbox', { name: 'Ticker' }).fill(DRAFT_TICKER);
    await page.getByRole('textbox', { name: 'Name' }).fill(DRAFT_NAME);

    const listResponsePromise = page.waitForResponse(isAssetCollectionRequest);
    await page.getByRole('button', { name: 'Cancel' }).click();
    await expect(page).toHaveURL(/\/assets$/);
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
  });

  test('preserves deferred external data when saving basic asset fields', async ({ database, page }) => {
    const asset = await seedAsset(database, ORIGINAL_TICKER, ORIGINAL_NAME, PERSISTED_EXTERNAL_DATA);

    await page.goto(`/asset/${asset.id}`);
    await expectAssetEditor(page, { ...asset, externalData: PERSISTED_EXTERNAL_DATA });

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
    await expect(page).toHaveURL(/\/assets$/);
    await expect(page.getByRole('cell', { name: UPDATED_TICKER, exact: true })).toBeVisible();
  });
});

/** Seeds one asset directly in PostgreSQL for a browser scenario. */
async function seedAsset(
  database: { query<Row extends object>(text: string, values?: unknown[]): Promise<readonly Row[]> },
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

/** Asserts the list route, basic columns, and empty-state behavior. */
async function expectAssetTable(page: import('@playwright/test').Page): Promise<void> {
  await expect(page).toHaveURL(/\/assets$/);
  await expect(page.getByRole('link', { name: 'Assets', exact: true })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Assets', exact: true })).toBeVisible();
  await expect(page.getByRole('columnheader', { name: 'Ticker', exact: true })).toBeVisible();
  await expect(page.getByRole('columnheader', { name: 'Name', exact: true })).toBeVisible();
  await expect(page.getByText('No assets have been registered yet.')).toBeVisible();
}

/** Asserts the edit route contains the expected basic asset fields. */
async function expectAssetEditor(page: import('@playwright/test').Page, asset: Asset): Promise<void> {
  await expect(page).toHaveURL(`/asset/${asset.id}`);
  await expect(page.getByRole('heading', { name: 'Edit asset', exact: true })).toBeVisible();
  await expect(page.getByRole('textbox', { name: 'Ticker' })).toHaveValue(asset.ticker);
  await expect(page.getByRole('textbox', { name: 'Name' })).toHaveValue(asset.name);
}

/** Matches the same-origin request that loads the asset collection. */
function isAssetCollectionRequest(response: Response): boolean {
  return response.request().method() === 'GET' && new URL(response.url()).pathname === '/api/asset';
}

/** Matches the same-origin request that creates one asset. */
function isAssetCreationRequest(response: Response): boolean {
  return response.request().method() === 'POST' && new URL(response.url()).pathname === '/api/asset';
}

/** Matches the same-origin request that updates one asset. */
function isAssetUpdateRequest(response: Response): boolean {
  return response.request().method() === 'PUT' && new URL(response.url()).pathname === '/api/asset';
}

/** Matches an asset detail request, optionally constrained to one identifier. */
function isAssetDetailRequest(response: Response, identifier: number | undefined): boolean {
  const request = response.request();
  const pathname = new URL(response.url()).pathname;
  const matchesPath = /^\/api\/asset\/[^/]+$/.test(pathname);
  const matchesIdentifier = identifier === undefined || pathname === `/api/asset/${identifier}`;
  return request.method() === 'GET' && matchesPath && matchesIdentifier;
}
