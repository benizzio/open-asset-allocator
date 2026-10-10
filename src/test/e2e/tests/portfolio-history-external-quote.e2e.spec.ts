/**
 * Covers the external quote action and its row-scoped submission contract.
 *
 * @author GPT-6 Luna
 */
import type { Locator, Page, Request, Route } from '@playwright/test';
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';

const DEFAULT_ALLOCATION_STRUCTURE = {
  hierarchy: [
    { name: 'Assets', field: 'assetTicker' },
    { name: 'Classes', field: 'class' },
  ],
} as const;

/** Ordered provider links for the quote-action scenario; the row must carry only the first link.
 *
 * @author GPT-6 Luna
 */
const EXTERNAL_QUOTE_ASSOCIATIONS = [
  { source: 'YAHOO_FINANCE', exchangeId: 'PCX', ticker: 'IAU' },
  { source: 'YAHOO_FINANCE', exchangeId: 'SECOND_EXCHANGE', ticker: 'SECOND_TICKER' },
] as const;
const EXTERNAL_QUOTE_PORTFOLIO_NAME = 'E2E External Quote Portfolio';
const EXTERNAL_QUOTE_ASSET_TICKER = 'E2E:QUOTE-ACTION';
const EXTERNAL_QUOTE_ASSET_NAME = 'E2E External Quote Asset';
const EXTERNAL_QUOTE_TIME_TAG = 'E2E_EXT_QUOTE_001';

/** Seed identifiers for loaded-row and quote-transition scenarios.
 *
 * @author GPT-6 Luna
 */
const EXTERNAL_QUOTE_TRANSITION_PORTFOLIO_NAME = 'E2E External Quote Transition Portfolio';
const EXTERNAL_QUOTE_MISSING_DATA_ASSET_TICKER = 'E2E:QUOTE-NO-DATA';
const EXTERNAL_QUOTE_MISSING_DATA_ASSET_NAME = 'E2E Quote Asset Without Provider Data';
const EXTERNAL_QUOTE_LOADED_OBSERVATION_TAG = 'E2E_QUOTE_LOADED_001';
const EXTERNAL_QUOTE_TRANSITION_TIME_TAG = 'E2E_QUOTE_TRANSITION_001';
const EXTERNAL_QUOTE_ZERO_TIME_TAG = 'E2E_QUOTE_ZERO_001';
const EXTERNAL_QUOTE_ERROR_TIME_TAG = 'E2E_QUOTE_ERROR_001';
const EXTERNAL_QUOTE_STALE_TIME_TAG = 'E2E_QUOTE_STALE_001';
const EXTERNAL_QUOTE_ISOLATION_PORTFOLIO_NAME = 'E2E External Quote Isolation Portfolio';
const EXTERNAL_QUOTE_ISOLATION_TIME_TAG = 'E2E_QUOTE_ISOLATION_001';
const EXTERNAL_QUOTE_FAILURE_TIME_TAG = 'E2E_QUOTE_FAILURE_001';

/** Three independently identified provider associations used to expose cross-row quote leakage.
 *
 * @author GPT-6 Luna
 */
const EXTERNAL_QUOTE_ISOLATION_ASSETS = [
  {
    ticker: 'E2E:QUOTE-ISOLATION-A',
    name: 'E2E Quote Isolation Asset A',
    externalAsset: { source: 'YAHOO_FINANCE', exchangeId: 'PCX', ticker: 'IAU' },
  },
  {
    ticker: 'E2E:QUOTE-ISOLATION-B',
    name: 'E2E Quote Isolation Asset B',
    externalAsset: { source: 'YAHOO_FINANCE', exchangeId: 'PCX', ticker: 'GLD' },
  },
  {
    ticker: 'E2E:QUOTE-ISOLATION-C',
    name: 'E2E Quote Isolation Asset C',
    externalAsset: { source: 'YAHOO_FINANCE', exchangeId: 'PCX', ticker: 'SLV' },
  },
] as const;

/** Provider key triplet attached to one persisted asset.
 *
 * @author GPT-6 Luna
 */
type ExternalAssetAssociation = {
  source: string;
  exchangeId: string;
  ticker: string;
};

/** Deferred response gate used to control mock route completion order.
 *
 * @author GPT-6 Luna
 */
type QuoteResponseGate = {
  started: Promise<void>;
  wait: Promise<void>;
  markStarted: () => void;
  release: () => void;
};

/** Persisted portfolio identity returned by direct SQL test seeding.
 *
 * @author GPT-6 Luna
 */
type SeededPortfolio = {
  id: number;
  name: string;
};

/** Persisted asset identity used by quote-action E2E scenarios.
 *
 * @author GPT-6 Luna
 */
type SeededAsset = {
  id: number;
  name: string;
  ticker: string;
};

/** Persisted observation identity and tag returned by PostgreSQL.
 *
 * @author GPT-6 Luna
 */
type PersistedObservation = {
  id: number;
  observation_timestamp: Date;
  observation_time_tag: string;
};

test.describe('scenario 16: external asset quote action and submission contract', () => {
  /** Verifies a mocked quote updates the row without saving, then persists the exact first-provider association.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.1: fetches an external quote and submits its provider keys with the explicit save', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedExternalAssetQuoteData(database);
    const [firstExternalAsset] = seededData.externalAssetAssociations;

    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const newObservationItem = page.locator('#portfolio-history-management-container-0');
    const timeTagInput = newObservationItem.getByRole('textbox', { name: 'Time tag' });
    await timeTagInput.fill(EXTERNAL_QUOTE_TIME_TAG);
    await newObservationItem.locator('#portfolio-history-management-trigger-0 > button').click();

    const form = page.locator('#portfolio-history-management-form-0');
    const row = await addAllocationRow(page, form, 0);
    const quoteAction = row.locator('[data-quote-action]');
    const externalAssetKeys = row.locator('[data-external-asset-keys]');
    const externalAssetKeyInputs = externalAssetKeys.locator('input');
    await expect(quoteAction).toBeHidden();
    // Check native fieldset state directly; Playwright's state matcher does not reliably model this fieldset.
    expect(await externalAssetKeys.evaluate((element) => (element as HTMLFieldSetElement).disabled)).toBe(true);
    expect(await externalAssetKeyInputs.evaluateAll((inputs) => {
      return inputs.every((input) => input.matches(':disabled'));
    })).toBe(true);
    await selectAssetFromAutocomplete(page, row, seededData.asset.ticker);
    await expectExistingAsset(row, seededData.asset);

    await expect(quoteAction).toBeVisible();
    await expect(row.getByRole('button', { name: 'Fetch latest closing price', exact: true })).toBeVisible();
    expect(await externalAssetKeys.evaluate((element) => (element as HTMLFieldSetElement).disabled)).toBe(false);
    expect(await externalAssetKeyInputs.evaluateAll((inputs) => {
      return inputs.every((input) => !input.matches(':disabled'));
    })).toBe(true);
    await expect(externalAssetKeys).toBeHidden();
    await expect(externalAssetKeys.locator('[data-external-asset-key]')).toHaveCount(3);
    await expect(row.locator('[data-external-asset-key="source"]')).toHaveValue(firstExternalAsset.source);
    await expect(row.locator('[data-external-asset-key="exchangeId"]')).toHaveValue(firstExternalAsset.exchangeId);
    await expect(row.locator('[data-external-asset-key="ticker"]')).toHaveValue(firstExternalAsset.ticker);

    const externalAssetFormEntries = await form.evaluate((formElement) => {
      return Array.from(new FormData(formElement as HTMLFormElement).entries())
        .filter(([name]) => name.includes('[externalAsset]'));
    });
    expect(externalAssetFormEntries).toEqual([
      ['allocations[0][externalAsset][source]', firstExternalAsset.source],
      ['allocations[0][externalAsset][exchangeId]', firstExternalAsset.exchangeId],
      ['allocations[0][externalAsset][ticker]', firstExternalAsset.ticker],
    ]);

    await row.getByRole('combobox', { name: 'Class' }).fill('BONDS');
    await fillCalculatedValues(row, '3', '10', '30.00');

    const quotePath = `/api/asset/${seededData.asset.id}/external-asset/${firstExternalAsset.source}`
      + `/${firstExternalAsset.exchangeId}/${firstExternalAsset.ticker}/quote`;
    const savePath = `/api/portfolio/${seededData.portfolio.id}/history`;
    const quoteRequests: Request[] = [];
    const interceptedSaveRequests: Request[] = [];
    const managementReloadRequests: string[] = [];
    /** Records management-data fetches to detect an unintended quote-triggered reload.
     *
     * @author GPT-6 Luna
     */
    const trackManagementReload = (request: Request): void => {
      if(request.method() === 'GET'
        && new URL(request.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history/observation`) {
        managementReloadRequests.push(request.url());
      }
    };
    /** Fulfills the quote endpoint locally so this scenario never contacts a live provider.
     *
     * @author GPT-6 Luna
     */
    const fulfillQuote = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          ticker: firstExternalAsset.ticker,
          exchangeId: firstExternalAsset.exchangeId,
          currency: 'USD',
          lastCloseQuote: '125.123456789',
          lastCloseDate: '2024-01-01T00:00:00Z',
        }),
      });
    };
    /** Captures and blocks form submissions if the quote action unexpectedly saves the observation.
     *
     * @author GPT-6 Luna
     */
    const blockUnexpectedSave = async (route: Route): Promise<void> => {
      if(route.request().method() === 'POST') {
        interceptedSaveRequests.push(route.request());
        await route.fulfill({ status: 204 });
        return;
      }

      await route.continue();
    };
    page.on('request', trackManagementReload);
    await page.route(`**${quotePath}`, fulfillQuote);
    await page.route(`**${savePath}`, blockUnexpectedSave);

    try {
      const quoteRequestPromise = page.waitForRequest((request) => {
        return request.method() === 'GET' && new URL(request.url()).pathname === quotePath;
      });
      const quoteResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET' && new URL(response.url()).pathname === quotePath;
      });
      await quoteAction.click();

      const quoteRequest = await quoteRequestPromise;
      const quoteResponse = await quoteResponsePromise;
      expect(quoteRequest.method()).toBe('GET');
      expect(quoteRequest.postData()).toBeNull();
      expect(new URL(quoteRequest.url()).search).toBe('');
      expect(quoteResponse.status()).toBe(200);
      expect(quoteRequests).toHaveLength(1);

      await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue('125.12345679');
      await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue('3');
      await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue('375.00');
      await expect(timeTagInput).toHaveValue(EXTERNAL_QUOTE_TIME_TAG);
      await expect(form).toBeVisible();
      await expect(row).toBeVisible();

      await page.waitForLoadState('networkidle');
      expect(interceptedSaveRequests).toHaveLength(0);
      expect(managementReloadRequests).toHaveLength(0);
      await expect(page.locator('#toast-notification-container .toast[role="alert"]')
        .filter({ hasText: 'Portfolio observation data saved successfully.' })).toHaveCount(0);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
      await expect(database.query(
        'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
        [seededData.portfolio.id],
      )).resolves.toEqual([]);
    } finally {
      page.off('request', trackManagementReload);
      await page.unroute(`**${savePath}`, blockUnexpectedSave);
      await page.unroute(`**${quotePath}`, fulfillQuote);
    }

    const saveRequestPromise = page.waitForRequest((request) => {
      return request.method() === 'POST' && new URL(request.url()).pathname === savePath;
    });
    const saveResponsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'POST' && new URL(response.url()).pathname === savePath;
    });
    await form.locator('tfoot button.btn-primary[type="submit"]').click();

    const saveRequest = await saveRequestPromise;
    const saveResponse = await saveResponsePromise;
    expect(saveResponse.status()).toBe(204);
    const savePayload = JSON.parse(saveRequest.postData() ?? 'null') as {
      allocations?: Array<{ externalAsset?: unknown }>;
    };
    expect(savePayload.allocations).toHaveLength(1);
    expect(savePayload.allocations?.[0]?.externalAsset).toEqual(firstExternalAsset);

    await expectSuccessNotification(page);
    const persistedObservation = await findObservationByTag(
      database,
      seededData.portfolio,
      EXTERNAL_QUOTE_TIME_TAG,
    );
    await expectManagementObservationOrder(page, [persistedObservation]);

    const persistedAllocations = await database.query<{
      asset_id: number;
      asset_market_price: string;
      asset_quantity: string;
      class: string;
      observation_time_tag: string;
      ticker: string;
      total_market_value: string;
    }>(
      `SELECT paf.asset_id,
              COALESCE(paf.asset_market_price, 0)::numeric(18,8)::text AS asset_market_price,
              COALESCE(paf.asset_quantity, 0)::numeric(18,8)::text AS asset_quantity,
              paf.class,
              ot.observation_time_tag,
              a.ticker,
              paf.total_market_value::text AS total_market_value
       FROM public.portfolio_allocation_fact paf
       JOIN public.asset a ON a.id = paf.asset_id
       JOIN public.portfolio_allocation_obs_time ot ON ot.id = paf.observation_time_id
       WHERE paf.portfolio_id = $1
         AND ot.id = $2`,
      [seededData.portfolio.id, persistedObservation.id],
    );
    expect(persistedAllocations).toEqual([{
      asset_id: seededData.asset.id,
      asset_market_price: '125.12345679',
      asset_quantity: '3.00000000',
      class: 'BONDS',
      observation_time_tag: EXTERNAL_QUOTE_TIME_TAG,
      ticker: EXTERNAL_QUOTE_ASSET_TICKER,
      total_market_value: '375',
    }]);

    const persistedAsset = await database.query<{ external_data: string | null }>(
      'SELECT external_data::text AS external_data FROM public.asset WHERE id = $1',
      [seededData.asset.id],
    );
    expect(persistedAsset).toHaveLength(1);
    expect(JSON.parse(persistedAsset[0].external_data ?? 'null')).toEqual({
      data: seededData.externalAssetAssociations,
    });
  });

  /** Covers server-loaded quote keys and eligibility changes through selection, reset, and missing data.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.2: keeps quote eligibility in sync for loaded and newly selected rows', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedExternalQuoteLoadedHistory(database);
    const [firstExternalAsset] = EXTERNAL_QUOTE_ASSOCIATIONS;

    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio, 1);

    const loadedObservationItem = page.locator(
      `#portfolio-history-management-container-${seededData.observation.id}`,
    );
    await loadedObservationItem.getByRole('button', {
      name: EXTERNAL_QUOTE_LOADED_OBSERVATION_TAG,
      exact: true,
    }).click();
    const loadedRow = page.locator(
      `#portfolio-history-management-form-${seededData.observation.id}-row-0`,
    );
    await expect(loadedRow).toBeVisible();
    await expect(loadedRow.locator('td').nth(0)).toContainText(seededData.quoteAsset.ticker);
    await expectExternalQuoteEligibility(loadedRow, firstExternalAsset);

    const newObservationForm = await openNewExternalQuoteForm(
      page,
      seededData.portfolio,
      EXTERNAL_QUOTE_TRANSITION_TIME_TAG,
      1,
    );
    const row = await addAllocationRow(page, newObservationForm, 0);
    await searchForAsset(page, row, seededData.quoteAsset.ticker, 200);
    await expectExistingAsset(row, seededData.quoteAsset);
    await expectExternalQuoteEligibility(row, firstExternalAsset);

    await row.locator('[data-asset-action-button]').click();
    await expectExternalQuoteEligibility(row, null);

    await page.mouse.move(0, 0);
    const assetSearchInput = row.getByRole('combobox', { name: 'Asset', exact: true });
    await assetSearchInput.fill(seededData.quoteAsset.ticker);
    const enterLookup = page.waitForResponse((response) => {
      return response.request().method() === 'GET'
        && new URL(response.url()).pathname === `/api/asset/${seededData.quoteAsset.ticker}`;
    });
    await assetSearchInput.press('Enter');
    expect((await enterLookup).status()).toBe(200);
    await expectExistingAsset(row, seededData.quoteAsset);
    await expectExternalQuoteEligibility(row, firstExternalAsset);

    await row.locator('[data-asset-action-button]').click();
    await expectExternalQuoteEligibility(row, null);
    await selectAssetFromAutocomplete(page, row, seededData.assetWithoutExternalData.ticker);
    await expectExistingAsset(row, seededData.assetWithoutExternalData);
    await expectExternalQuoteEligibility(row, null);

    const persistedObservation = await database.query<{
      asset_id: number;
      asset_market_price: string;
      asset_quantity: string;
      observation_time_tag: string;
      total_market_value: string;
    }>(
      `SELECT paf.asset_id,
              paf.asset_market_price::numeric(18,8)::text AS asset_market_price,
              paf.asset_quantity::numeric(18,8)::text AS asset_quantity,
              ot.observation_time_tag,
              paf.total_market_value::text AS total_market_value
       FROM public.portfolio_allocation_fact paf
       JOIN public.portfolio_allocation_obs_time ot ON ot.id = paf.observation_time_id
       WHERE paf.portfolio_id = $1`,
      [seededData.portfolio.id],
    );
    expect(persistedObservation).toEqual([{
      asset_id: seededData.quoteAsset.id,
      asset_market_price: '25.00000000',
      asset_quantity: '2.00000000',
      observation_time_tag: EXTERNAL_QUOTE_LOADED_OBSERVATION_TAG,
      total_market_value: '50',
    }]);

    const persistedAssets = await database.query<{
      external_data: string | null;
      ticker: string;
    }>(
      'SELECT ticker, external_data::text AS external_data FROM public.asset ORDER BY ticker',
    );
    expect(persistedAssets.map((asset) => ({
      external_data: JSON.parse(asset.external_data ?? 'null'),
      ticker: asset.ticker,
    }))).toEqual([
      {
        external_data: { data: EXTERNAL_QUOTE_ASSOCIATIONS },
        ticker: seededData.quoteAsset.ticker,
      },
      { external_data: null, ticker: seededData.assetWithoutExternalData.ticker },
    ].sort((left, right) => left.ticker.localeCompare(right.ticker)));
  });

  /** Confirms a zero quote updates only the price when quantity is blank, even with an incomplete sibling row.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.3: applies a zero quote without quantity and blocks save for an incomplete sibling', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedExternalAssetQuoteData(database);
    const [firstExternalAsset] = seededData.externalAssetAssociations;
    const form = await openNewExternalQuoteForm(page, seededData.portfolio, EXTERNAL_QUOTE_ZERO_TIME_TAG);
    const quotedRow = await addAllocationRow(page, form, 0);
    await selectAssetFromAutocomplete(page, quotedRow, seededData.asset.ticker);
    await expectExistingAsset(quotedRow, seededData.asset);
    await expectExternalQuoteEligibility(quotedRow, firstExternalAsset);
    await quotedRow.getByRole('combobox', { name: 'Class' }).fill('BONDS');
    await fillDirectTotalMarketValue(quotedRow, '2,000');

    const incompleteRow = await addAllocationRow(page, form, 1);
    await incompleteRow.getByRole('combobox', { name: 'Class' }).fill('BONDS');
    await incompleteRow.getByRole('textbox', { name: 'Total market value' }).fill('100');
    await incompleteRow.getByRole('combobox', { name: 'Asset', exact: true }).fill('E2E:UNRESOLVED-QUOTE');
    await expect(incompleteRow.locator('[data-asset-search-autocomplete]'))
      .toHaveAttribute('data-asset-selection-state', 'search');

    const quotePath = externalAssetQuotePath(seededData.asset.id, firstExternalAsset);
    const savePath = `/api/portfolio/${seededData.portfolio.id}/history`;
    const quoteRequests: Request[] = [];
    const interceptedSaveRequests: Request[] = [];
    /** Returns a zero quote for the eligible row without involving a provider.
     *
     * @author GPT-6 Luna
     */
    const fulfillZeroQuote = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(externalQuoteResponse(firstExternalAsset, '0')),
      });
    };
    /** Intercepts an unexpected save so the invalid-row assertion cannot persist data on failure.
     *
     * @author GPT-6 Luna
     */
    const blockUnexpectedSave = async (route: Route): Promise<void> => {
      if(route.request().method() === 'POST') {
        interceptedSaveRequests.push(route.request());
        await route.fulfill({ status: 204 });
        return;
      }

      await route.continue();
    };
    await page.route(`**${quotePath}`, fulfillZeroQuote);
    await page.route(`**${savePath}`, blockUnexpectedSave);

    try {
      const quoteResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === quotePath;
      });
      await quotedRow.locator('[data-quote-action]').click();
      const quoteResponse = await quoteResponsePromise;
      expect(quoteResponse.status()).toBe(200);
      expect(quoteRequests).toHaveLength(1);
      await expectFinancialValues(quotedRow, '0.00000000', '', '2,000.00');

      await form.locator('tfoot button.btn-primary[type="submit"]').click();
      await expect(incompleteRow.locator('[data-asset-ticker-extra-error-message]'))
        .toHaveText('Reference an existing asset or create a new one');
      expect(interceptedSaveRequests).toHaveLength(0);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
      await expect(database.query(
        'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
        [seededData.portfolio.id],
      )).resolves.toEqual([]);
      await expect(database.query(
        'SELECT id FROM public.asset WHERE ticker = $1',
        ['E2E:UNRESOLVED-QUOTE'],
      )).resolves.toEqual([]);
    } finally {
      await page.unroute(`**${savePath}`, blockUnexpectedSave);
      await page.unroute(`**${quotePath}`, fulfillZeroQuote);
    }
  });

  /** Rejects a mismatched successful response without changing values, then verifies a valid retry recovers.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.4: restores the quote action and recovers after a malformed response', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedExternalAssetQuoteData(database);
    const [firstExternalAsset] = seededData.externalAssetAssociations;
    const form = await openNewExternalQuoteForm(page, seededData.portfolio, EXTERNAL_QUOTE_ERROR_TIME_TAG);
    const row = await addAllocationRow(page, form, 0);
    await selectAssetFromAutocomplete(page, row, seededData.asset.ticker);
    await expectExistingAsset(row, seededData.asset);
    await expectExternalQuoteEligibility(row, firstExternalAsset);
    await row.getByRole('combobox', { name: 'Class' }).fill('BONDS');
    await fillCalculatedValues(row, '3', '10', '30.00');

    const quotePath = externalAssetQuotePath(seededData.asset.id, firstExternalAsset);
    const quoteRequests: Request[] = [];
    /** Returns a mismatched quote first and a valid quote on retry.
     *
     * @author GPT-6 Luna
     */
    const fulfillQuoteAttempts = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      const response = quoteRequests.length === 1
        ? { ...externalQuoteResponse(firstExternalAsset, '999'), ticker: 'MISMATCHED' }
        : externalQuoteResponse(firstExternalAsset, '25');
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(response),
      });
    };
    await page.route(`**${quotePath}`, fulfillQuoteAttempts);

    try {
      const malformedResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === quotePath;
      });
      await row.locator('[data-quote-action]').click();
      expect((await malformedResponsePromise).status()).toBe(200);
      await expect(page.locator('#toast-notification-container .toast.text-bg-danger')
        .filter({ hasText: 'The latest closing price response was invalid.' })).toBeVisible();
      await expectFinancialValues(row, '10.00000000', '3', '30.00');
      await expectQuoteActionIdle(row.locator('[data-quote-action]'));

      const retryResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === quotePath;
      });
      await row.locator('[data-quote-action]').click();
      expect((await retryResponsePromise).status()).toBe(200);
      expect(quoteRequests).toHaveLength(2);
      await expectFinancialValues(row, '25.00000000', '3', '75.00');
      await expectQuoteActionIdle(row.locator('[data-quote-action]'));
      await expect(database.query(
        'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
        [seededData.portfolio.id],
      )).resolves.toEqual([]);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
    } finally {
      await page.unroute(`**${quotePath}`, fulfillQuoteAttempts);
    }
  });

  /** Rejects duplicate in-flight GETs and ignores a valid response after reset and reselection.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.5: deduplicates an in-flight quote and ignores it after row reselection', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedExternalAssetQuoteData(database);
    const [firstExternalAsset] = seededData.externalAssetAssociations;
    const form = await openNewExternalQuoteForm(page, seededData.portfolio, EXTERNAL_QUOTE_STALE_TIME_TAG);
    const row = await addAllocationRow(page, form, 0);
    await selectAssetFromAutocomplete(page, row, seededData.asset.ticker);
    await expectExistingAsset(row, seededData.asset);
    await expectExternalQuoteEligibility(row, firstExternalAsset);
    await row.getByRole('combobox', { name: 'Class' }).fill('BONDS');
    await fillCalculatedValues(row, '3', '10', '30.00');

    const quotePath = externalAssetQuotePath(seededData.asset.id, firstExternalAsset);
    const quoteRequests: Request[] = [];
    let releaseFirstResponse: () => void = () => undefined;
    let notifyFirstRequestStarted: () => void = () => undefined;
    const firstResponseGate = new Promise<void>((resolve) => { releaseFirstResponse = resolve; });
    const firstRequestStarted = new Promise<void>((resolve) => { notifyFirstRequestStarted = resolve; });
    /** Holds the first valid response so duplicate clicks and row reselection occur before completion.
     *
     * @author GPT-6 Luna
     */
    const holdFirstQuoteResponse = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      if(quoteRequests.length === 1) {
        notifyFirstRequestStarted();
        await firstResponseGate;
      }

      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(externalQuoteResponse(firstExternalAsset, '99')),
      });
    };
    await page.route(`**${quotePath}`, holdFirstQuoteResponse);

    try {
      const generationBeforeRequest = await row.getAttribute('data-asset-selection-generation');
      const quoteResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === quotePath;
      });
      await row.locator('[data-quote-action]').click();
      await firstRequestStarted;
      expect(quoteRequests).toHaveLength(1);
      const quoteAction = row.locator('[data-quote-action]');
      await expect(quoteAction).toHaveAttribute('aria-busy', 'true');
      await expect(quoteAction.locator('[data-quote-spinner]')).toBeVisible();

      await quoteAction.evaluate((element) => {
        element.dispatchEvent(new MouseEvent('click', { bubbles: true, cancelable: true }));
      });
      await row.locator('[data-asset-action-button]').click();
      await expectExternalQuoteEligibility(row, null);
      const generationAfterReset = await row.getAttribute('data-asset-selection-generation');
      expect(generationAfterReset).not.toBe(generationBeforeRequest);

      await selectAssetFromAutocomplete(page, row, seededData.asset.ticker);
      await expectExistingAsset(row, seededData.asset);
      await expectExternalQuoteEligibility(row, firstExternalAsset);
      const generationAfterReselection = await row.getAttribute('data-asset-selection-generation');
      expect(generationAfterReselection).not.toBe(generationBeforeRequest);
      expect(generationAfterReselection).not.toBe(generationAfterReset);
      await expectFinancialValues(row, '10.00000000', '3', '30.00');

      releaseFirstResponse();
      expect((await quoteResponsePromise).status()).toBe(200);
      await page.waitForLoadState('networkidle');
      expect(quoteRequests).toHaveLength(1);
      await expectFinancialValues(row, '10.00000000', '3', '30.00');
      await expectQuoteActionIdle(quoteAction);
      await expect(database.query(
        'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
        [seededData.portfolio.id],
      )).resolves.toEqual([]);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
    } finally {
      releaseFirstResponse();
      await page.unroute(`**${quotePath}`, holdFirstQuoteResponse);
    }
  });

  /** Verifies concurrent row isolation and ignores a delayed quote after its row is removed.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.6: isolates concurrent quotes and ignores a response after row removal', async ({ database, page }) => {
    test.setTimeout(90_000);
    const seededData = await seedExternalQuoteIsolationData(database);
    const form = await openNewExternalQuoteForm(page, seededData.portfolio, EXTERNAL_QUOTE_ISOLATION_TIME_TAG);
    const rows: Locator[] = [];

    for (const [index, seededAsset] of seededData.assets.entries()) {
      const row = await addAllocationRow(page, form, index);
      await selectAssetFromAutocomplete(page, row, seededAsset.asset.ticker);
      await expectExistingAsset(row, seededAsset.asset);
      await expectExternalQuoteEligibility(row, seededAsset.externalAsset);
      await row.getByRole('combobox', { name: 'Class' }).fill('BONDS');
      const initialValues = [
        { quantity: '2', marketPrice: '10', total: '20.00' },
        { quantity: '3', marketPrice: '20', total: '60.00' },
        { quantity: '4', marketPrice: '30', total: '120.00' },
      ][index];
      await fillCalculatedValues(row, initialValues.quantity, initialValues.marketPrice, initialValues.total);
      rows.push(row);
    }

    const [removedRow, firstSiblingRow, secondSiblingRow] = rows;
    const [removedSeed, firstSiblingSeed, secondSiblingSeed] = seededData.assets;
    const removedPriceInput = await removedRow.getByRole('textbox', { name: 'Market price' }).elementHandle();
    if (!removedPriceInput) {
      throw new Error('Could not capture the market-price input before removing its row.');
    }
    const removedPriceBeforeQuote = await removedPriceInput.evaluate((input) => (input as HTMLInputElement).value);

    const removedGate = createQuoteResponseGate();
    const firstSiblingGate = createQuoteResponseGate();
    const secondSiblingGate = createQuoteResponseGate();
    const removedQuotePath = externalAssetQuotePath(removedSeed.asset.id, removedSeed.externalAsset);
    const firstSiblingQuotePath = externalAssetQuotePath(firstSiblingSeed.asset.id, firstSiblingSeed.externalAsset);
    const secondSiblingQuotePath = externalAssetQuotePath(secondSiblingSeed.asset.id, secondSiblingSeed.externalAsset);
    const quoteRequests: Request[] = [];
    /** Holds the removed row's response until both surviving rows have completed.
     *
     * @author GPT-6 Luna
     */
    const holdRemovedRowQuote = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      removedGate.markStarted();
      await removedGate.wait;
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(externalQuoteResponse(removedSeed.externalAsset, '999')),
      });
    };
    /** Holds the first surviving row's response to force the other row to finish first.
     *
     * @author GPT-6 Luna
     */
    const holdFirstSiblingQuote = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      firstSiblingGate.markStarted();
      await firstSiblingGate.wait;
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(externalQuoteResponse(firstSiblingSeed.externalAsset, '200')),
      });
    };
    /** Holds the second surviving row's response until both sibling requests are in flight.
     *
     * @author GPT-6 Luna
     */
    const holdSecondSiblingQuote = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      secondSiblingGate.markStarted();
      await secondSiblingGate.wait;
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(externalQuoteResponse(secondSiblingSeed.externalAsset, '300')),
      });
    };
    const quoteRoutes = [
      [removedQuotePath, holdRemovedRowQuote],
      [firstSiblingQuotePath, holdFirstSiblingQuote],
      [secondSiblingQuotePath, holdSecondSiblingQuote],
    ] as const;
    for (const [quotePath, handler] of quoteRoutes) {
      await page.route(`**${quotePath}`, handler);
    }

    try {
      const removedResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === removedQuotePath;
      });
      await removedRow.locator('[data-quote-action]').click();
      await removedGate.started;
      expect(quoteRequests).toHaveLength(1);

      await removedRow.locator('button[onclick^="this.closest"]').click();
      await expect(removedRow).toHaveCount(0);
      expect(await removedPriceInput.evaluate((input) => (input as HTMLInputElement).value))
        .toBe(removedPriceBeforeQuote);

      const firstSiblingResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === firstSiblingQuotePath;
      });
      await firstSiblingRow.locator('[data-quote-action]').click();
      await firstSiblingGate.started;

      const secondSiblingResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === secondSiblingQuotePath;
      });
      await secondSiblingRow.locator('[data-quote-action]').click();
      await secondSiblingGate.started;
      expect(quoteRequests).toHaveLength(3);

      secondSiblingGate.release();
      expect((await secondSiblingResponsePromise).status()).toBe(200);
      await expectFinancialValues(secondSiblingRow, '300.00000000', '4', '1,200.00');
      await expectFinancialValues(firstSiblingRow, '20.00000000', '3', '60.00');
      expect(await removedPriceInput.evaluate((input) => (input as HTMLInputElement).value))
        .toBe(removedPriceBeforeQuote);

      firstSiblingGate.release();
      expect((await firstSiblingResponsePromise).status()).toBe(200);
      await expectFinancialValues(firstSiblingRow, '200.00000000', '3', '600.00');
      await expectFinancialValues(secondSiblingRow, '300.00000000', '4', '1,200.00');

      removedGate.release();
      expect((await removedResponsePromise).status()).toBe(200);
      await page.waitForLoadState('networkidle');
      expect(quoteRequests).toHaveLength(3);
      expect(await removedPriceInput.evaluate((input) => (input as HTMLInputElement).value))
        .toBe(removedPriceBeforeQuote);
      await expectFinancialValues(firstSiblingRow, '200.00000000', '3', '600.00');
      await expectFinancialValues(secondSiblingRow, '300.00000000', '4', '1,200.00');
      await expectQuoteActionIdle(firstSiblingRow.locator('[data-quote-action]'));
      await expectQuoteActionIdle(secondSiblingRow.locator('[data-quote-action]'));
      await expect(database.query(
        'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
        [seededData.portfolio.id],
      )).resolves.toEqual([]);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
    } finally {
      removedGate.release();
      firstSiblingGate.release();
      secondSiblingGate.release();
      for (const [quotePath, handler] of quoteRoutes) {
        await page.unroute(`**${quotePath}`, handler);
      }
    }
  });

  /** Verifies a failed quote GET reports one global error, restores its controls, and permits retry.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.7: restores the quote action after a network failure and allows retry', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedExternalAssetQuoteData(database);
    const [firstExternalAsset] = seededData.externalAssetAssociations;
    const form = await openNewExternalQuoteForm(page, seededData.portfolio, EXTERNAL_QUOTE_FAILURE_TIME_TAG);
    const row = await addAllocationRow(page, form, 0);
    await selectAssetFromAutocomplete(page, row, seededData.asset.ticker);
    await expectExistingAsset(row, seededData.asset);
    await expectExternalQuoteEligibility(row, firstExternalAsset);
    await row.getByRole('combobox', { name: 'Class' }).fill('BONDS');
    await fillCalculatedValues(row, '3', '10', '30.00');

    const quotePath = externalAssetQuotePath(seededData.asset.id, firstExternalAsset);
    const quoteRequests: Request[] = [];
    /** Aborts the first quote GET and returns a valid quote on retry.
     *
     * @author GPT-6 Luna
     */
    const failThenFulfillQuote = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      if(quoteRequests.length === 1) {
        await route.abort('failed');
        return;
      }

      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(externalQuoteResponse(firstExternalAsset, '25')),
      });
    };
    await page.route(`**${quotePath}`, failThenFulfillQuote);

    try {
      const failedRequestPromise = page.waitForEvent('requestfailed', (request) => {
        return request.method() === 'GET' && new URL(request.url()).pathname === quotePath;
      });
      await row.locator('[data-quote-action]').click();
      await failedRequestPromise;

      const errorToasts = page.locator('#toast-notification-container .toast.text-bg-danger[role="alert"]');
      await expect(errorToasts).toHaveCount(1);
      await expect(errorToasts).toContainText('An unexpected error occurred while communicating with the server.');
      await expectFinancialValues(row, '10.00000000', '3', '30.00');
      await expectQuoteActionIdle(row.locator('[data-quote-action]'));

      const retryResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === quotePath;
      });
      await row.locator('[data-quote-action]').click();
      expect((await retryResponsePromise).status()).toBe(200);
      expect(quoteRequests).toHaveLength(2);
      await expectFinancialValues(row, '25.00000000', '3', '75.00');
      await expectQuoteActionIdle(row.locator('[data-quote-action]'));
      await expect(errorToasts).toHaveCount(1);
      await expect(database.query(
        'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
        [seededData.portfolio.id],
      )).resolves.toEqual([]);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
    } finally {
      await page.unroute(`**${quotePath}`, failThenFulfillQuote);
    }
  });

  /** Ignores a delayed quote after the management list reload replaces its loaded observation row.
   *
   * @author GPT-6 Luna
   */
  test('scenario 16.8: ignores a delayed quote after the portfolio-history management reload', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedExternalQuoteLoadedHistory(database);
    const [firstExternalAsset] = EXTERNAL_QUOTE_ASSOCIATIONS;

    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio, 1);

    const observationItem = page.locator(
      `#portfolio-history-management-container-${seededData.observation.id}`,
    );
    const observationHistoryPath = `/api/portfolio/${seededData.portfolio.id}/history`;
    const initialObservationResponsePromise = page.waitForResponse((response) => {
      const url = new URL(response.url());
      return response.request().method() === 'GET'
        && url.pathname === observationHistoryPath
        && url.searchParams.get('observationTimestampId') === String(seededData.observation.id);
    });
    await observationItem.getByRole('button', {
      name: EXTERNAL_QUOTE_LOADED_OBSERVATION_TAG,
      exact: true,
    }).click();
    expect((await initialObservationResponsePromise).status()).toBe(200);

    const rowSelector = `#portfolio-history-management-form-${seededData.observation.id}-row-0`;
    const loadedRow = page.locator(rowSelector);
    await expect(loadedRow).toBeVisible();
    await expectExternalQuoteEligibility(loadedRow, firstExternalAsset);
    await expectFinancialValues(loadedRow, '25.00000000', '2', '50.00');

    const loadedRowHandle = await loadedRow.elementHandle();
    const loadedMarketPriceInputHandle = await loadedRow.getByRole('textbox', { name: 'Market price' }).elementHandle();
    if (!loadedRowHandle || !loadedMarketPriceInputHandle) {
      throw new Error('Could not capture the loaded quote row before the management reload.');
    }
    const detachedInputValueBeforeReload = await loadedMarketPriceInputHandle.evaluate(
      (input) => (input as HTMLInputElement).value,
    );

    const quotePath = externalAssetQuotePath(seededData.quoteAsset.id, firstExternalAsset);
    const quoteGate = createQuoteResponseGate();
    const quoteRequests: Request[] = [];
    /** Holds a valid quote until the loaded observation row has been replaced from server state.
     *
     * @author GPT-6 Luna
     */
    const holdQuoteResponse = async (route: Route): Promise<void> => {
      quoteRequests.push(route.request());
      quoteGate.markStarted();
      await quoteGate.wait;
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(externalQuoteResponse(firstExternalAsset, '999')),
      });
    };
    await page.route(`**${quotePath}`, holdQuoteResponse);

    try {
      const quoteResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === quotePath;
      });
      await loadedRow.locator('[data-quote-action]').click();
      await quoteGate.started;
      expect(quoteRequests).toHaveLength(1);

      const managementContainer = page.locator('#accordion-portfolio-history-management');
      await expect(managementContainer).toHaveAttribute(
        'hx-trigger',
        'load, reload-portfolio-history-management',
      );
      const observationListPath = `/api/portfolio/${seededData.portfolio.id}/history/observation`;
      const observationListResponsePromise = page.waitForResponse((response) => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === observationListPath;
      });
      await managementContainer.evaluate((element) => {
        const htmx = (window as unknown as {
          htmx?: { trigger: (target: Element, eventName: string) => unknown };
        }).htmx;
        if (!htmx) {
          throw new Error('The global HTMX trigger API is unavailable.');
        }
        htmx.trigger(element, 'reload-portfolio-history-management');
      });
      expect((await observationListResponsePromise).status()).toBe(200);

      await expect(loadedRow).toHaveCount(0);
      await expect(managementContainer.locator(':scope > .accordion-item')).toHaveCount(2);
      await expect(managementContainer).not.toHaveClass(/\bhtmx-settling\b/);
      expect(await loadedRowHandle.evaluate((row) => row.isConnected)).toBe(false);
      expect(await loadedMarketPriceInputHandle.evaluate((input) => input.isConnected)).toBe(false);

      const reloadedObservationItem = page.locator(
        `#portfolio-history-management-container-${seededData.observation.id}`,
      );
      const reloadedObservationResponsePromise = page.waitForResponse((response) => {
        const url = new URL(response.url());
        return response.request().method() === 'GET'
          && url.pathname === observationHistoryPath
          && url.searchParams.get('observationTimestampId') === String(seededData.observation.id);
      });
      await reloadedObservationItem.getByRole('button', {
        name: EXTERNAL_QUOTE_LOADED_OBSERVATION_TAG,
        exact: true,
      }).click();
      expect((await reloadedObservationResponsePromise).status()).toBe(200);

      const replacementRow = page.locator(rowSelector);
      await expect(replacementRow).toBeVisible();
      await expectExternalQuoteEligibility(replacementRow, firstExternalAsset);
      await expectFinancialValues(replacementRow, '25.00000000', '2', '50.00');
      expect(await loadedRowHandle.evaluate((previousRow) => {
        const currentRow = document.getElementById(previousRow.id);
        return currentRow !== null && currentRow !== previousRow;
      })).toBe(true);

      quoteGate.release();
      expect((await quoteResponsePromise).status()).toBe(200);
      await page.waitForLoadState('networkidle');
      expect(quoteRequests).toHaveLength(1);
      await expectFinancialValues(replacementRow, '25.00000000', '2', '50.00');
      expect(await loadedMarketPriceInputHandle.evaluate((input) => (input as HTMLInputElement).value))
        .toBe(detachedInputValueBeforeReload);

      const persistedAllocation = await database.query<{
        asset_id: number;
        asset_market_price: string;
        asset_quantity: string;
        total_market_value: string;
      }>(
        `SELECT paf.asset_id,
                paf.asset_market_price::numeric(18,8)::text AS asset_market_price,
                paf.asset_quantity::numeric(18,8)::text AS asset_quantity,
                paf.total_market_value::text AS total_market_value
         FROM public.portfolio_allocation_fact paf
         WHERE paf.portfolio_id = $1 AND paf.observation_time_id = $2`,
        [seededData.portfolio.id, seededData.observation.id],
      );
      expect(persistedAllocation).toEqual([{
        asset_id: seededData.quoteAsset.id,
        asset_market_price: '25.00000000',
        asset_quantity: '2.00000000',
        total_market_value: '50',
      }]);
    } finally {
      quoteGate.release();
      await page.unroute(`**${quotePath}`, holdQuoteResponse);
    }
  });
});

/** Seeds the portfolio and persisted asset associations used by the external quote action scenario.
 *
 * @param database - Isolated test database used to create the required rows directly in PostgreSQL.
 * @returns The persisted portfolio, asset, and ordered external-provider associations.
 *
 * @author GPT-6 Luna
 */
async function seedExternalAssetQuoteData(database: E2eDatabase): Promise<{
  asset: SeededAsset;
  externalAssetAssociations: typeof EXTERNAL_QUOTE_ASSOCIATIONS;
  portfolio: SeededPortfolio;
}> {
  const portfolioRows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [EXTERNAL_QUOTE_PORTFOLIO_NAME, JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );
  const assetRows = await database.query<SeededAsset>(
    `INSERT INTO public.asset (ticker, name, external_data)
     VALUES ($1, $2, $3::jsonb)
     RETURNING id, ticker, name`,
    [
      EXTERNAL_QUOTE_ASSET_TICKER,
      EXTERNAL_QUOTE_ASSET_NAME,
      JSON.stringify({ data: EXTERNAL_QUOTE_ASSOCIATIONS }),
    ],
  );

  expect(portfolioRows).toHaveLength(1);
  expect(assetRows).toHaveLength(1);

  return {
    asset: assetRows[0],
    externalAssetAssociations: EXTERNAL_QUOTE_ASSOCIATIONS,
    portfolio: portfolioRows[0],
  };
}

/** Seeds a saved observation with one quote-enabled asset and one asset without provider data.
 *
 * @param database - Isolated test database used to create all scenario records directly in PostgreSQL.
 * @returns The portfolio, both assets, and the persisted observation used to render a loaded row.
 *
 * @author GPT-6 Luna
 */
async function seedExternalQuoteLoadedHistory(database: E2eDatabase): Promise<{
  assetWithoutExternalData: SeededAsset;
  observation: PersistedObservation;
  portfolio: SeededPortfolio;
  quoteAsset: SeededAsset;
}> {
  const portfolioRows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [EXTERNAL_QUOTE_TRANSITION_PORTFOLIO_NAME, JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );
  const assetRows = await database.query<SeededAsset>(
    `INSERT INTO public.asset (ticker, name, external_data)
     VALUES ($1, $2, $3::jsonb), ($4, $5, NULL)
     RETURNING id, ticker, name`,
    [
      EXTERNAL_QUOTE_ASSET_TICKER,
      EXTERNAL_QUOTE_ASSET_NAME,
      JSON.stringify({ data: EXTERNAL_QUOTE_ASSOCIATIONS }),
      EXTERNAL_QUOTE_MISSING_DATA_ASSET_TICKER,
      EXTERNAL_QUOTE_MISSING_DATA_ASSET_NAME,
    ],
  );

  expect(portfolioRows).toHaveLength(1);
  expect(assetRows).toHaveLength(2);
  const assetsByTicker = Object.fromEntries(
    assetRows.map((asset) => [asset.ticker, asset]),
  ) as Record<string, SeededAsset>;
  const observationRows = await database.query<PersistedObservation>(
    `INSERT INTO public.portfolio_allocation_obs_time (observation_time_tag, observation_timestamp)
     VALUES ($1, CURRENT_TIMESTAMP)
     RETURNING id, observation_timestamp, observation_time_tag`,
    [EXTERNAL_QUOTE_LOADED_OBSERVATION_TAG],
  );
  expect(observationRows).toHaveLength(1);

  const observation = observationRows[0];
  const allocationRows = await database.query<{ asset_id: number }>(
    `INSERT INTO public.portfolio_allocation_fact
       (portfolio_id, asset_id, class, cash_reserve, asset_quantity, asset_market_price,
        total_market_value, observation_time_id)
     VALUES ($1, $2, 'BONDS', false, 2, 25, 50, $3)
     RETURNING asset_id`,
    [portfolioRows[0].id, assetsByTicker[EXTERNAL_QUOTE_ASSET_TICKER].id, observation.id],
  );
  expect(allocationRows).toEqual([{ asset_id: assetsByTicker[EXTERNAL_QUOTE_ASSET_TICKER].id }]);

  return {
    assetWithoutExternalData: assetsByTicker[EXTERNAL_QUOTE_MISSING_DATA_ASSET_TICKER],
    observation,
    portfolio: portfolioRows[0],
    quoteAsset: assetsByTicker[EXTERNAL_QUOTE_ASSET_TICKER],
  };
}

/** Seeds three persisted assets with distinct provider tickers for concurrent row-isolation coverage.
 *
 * @param database - Isolated test database used to create the portfolio and assets directly in PostgreSQL.
 * @returns The portfolio and assets paired with their independent quote identifiers.
 *
 * @author GPT-6 Luna
 */
async function seedExternalQuoteIsolationData(database: E2eDatabase): Promise<{
  assets: Array<{ asset: SeededAsset; externalAsset: ExternalAssetAssociation }>;
  portfolio: SeededPortfolio;
}> {
  const portfolioRows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [EXTERNAL_QUOTE_ISOLATION_PORTFOLIO_NAME, JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );
  expect(portfolioRows).toHaveLength(1);

  const valuesClause = EXTERNAL_QUOTE_ISOLATION_ASSETS.map((_, index) => {
    const firstParameter = index * 3 + 1;
    return `($${firstParameter}, $${firstParameter + 1}, $${firstParameter + 2}::jsonb)`;
  }).join(', ');
  const assetParameters = EXTERNAL_QUOTE_ISOLATION_ASSETS.flatMap((asset) => [
    asset.ticker,
    asset.name,
    JSON.stringify({ data: [asset.externalAsset] }),
  ]);
  const assetRows = await database.query<SeededAsset>(
    `INSERT INTO public.asset (ticker, name, external_data)
     VALUES ${valuesClause}
     RETURNING id, ticker, name`,
    assetParameters,
  );
  expect(assetRows).toHaveLength(EXTERNAL_QUOTE_ISOLATION_ASSETS.length);
  const assetsByTicker = new Map(assetRows.map((asset) => [asset.ticker, asset]));

  return {
    assets: EXTERNAL_QUOTE_ISOLATION_ASSETS.map((assetSeed) => {
      const asset = assetsByTicker.get(assetSeed.ticker);
      if (!asset) {
        throw new Error(`Expected seeded quote asset ${assetSeed.ticker} to be returned from PostgreSQL.`);
      }
      return { asset, externalAsset: assetSeed.externalAsset };
    }),
    portfolio: portfolioRows[0],
  };
}

/** Opens a portfolio's new-observation form using a unique time tag.
 *
 * @param page - Browser page that opens and expands the management form.
 * @param portfolio - Portfolio whose history form will be opened.
 * @param timeTag - Unique time tag entered into the new-observation panel.
 * @returns The expanded form ready for allocation rows.
 *
 * @author GPT-6 Luna
 */
async function openNewExternalQuoteForm(
  page: Page,
  portfolio: SeededPortfolio,
  timeTag: string,
  observationCount = 0,
): Promise<Locator> {
  await page.goto('/');
  await expectRootShell(page);
  await page.goto(`/portfolio/${portfolio.id}/history/manage`);
  await expectPortfolioHistoryManagement(page, portfolio, observationCount);

  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  await newObservationItem.getByRole('textbox', { name: 'Time tag' }).fill(timeTag);
  await newObservationItem.locator('#portfolio-history-management-trigger-0 > button').click();

  const form = page.locator('#portfolio-history-management-form-0');
  await expect(form).toBeVisible();
  return form;
}

/** Asserts row-scoped quote eligibility, native fieldset disablement, and the first provider key triplet.
 *
 * @param row - Allocation row whose quote state is checked.
 * @param association - Expected first provider association, or null when quote data must be disabled.
 *
 * @author GPT-6 Luna
 */
async function expectExternalQuoteEligibility(
  row: Locator,
  association: ExternalAssetAssociation | null,
): Promise<void> {
  const fieldset = row.locator('[data-external-asset-keys]');
  const keyInputs = fieldset.locator('[data-external-asset-key]');
  const isDisabled = association === null;
  expect(await fieldset.evaluate((element) => (element as HTMLFieldSetElement).disabled)).toBe(isDisabled);
  await expect(fieldset).toBeHidden();
  await expect(keyInputs).toHaveCount(3);
  expect(await keyInputs.evaluateAll((inputs) => {
    return inputs.every((input) => input.matches(':disabled'));
  })).toBe(isDisabled);
  await expect(row.locator('[data-external-asset-key="source"]')).toHaveValue(association?.source ?? '');
  await expect(row.locator('[data-external-asset-key="exchangeId"]')).toHaveValue(association?.exchangeId ?? '');
  await expect(row.locator('[data-external-asset-key="ticker"]')).toHaveValue(association?.ticker ?? '');

  if(isDisabled) {
    await expect(row.locator('[data-quote-action]')).toBeHidden();
  } else {
    await expect(row.getByRole('button', { name: 'Fetch latest closing price', exact: true })).toBeVisible();
  }
}

/** Builds the encoded quote endpoint path for one persisted asset-provider association.
 *
 * @param assetId - Persisted asset ID required by the quote route.
 * @param association - Provider identifiers appended to the quote route.
 * @returns The exact GET path used by the quote action.
 *
 * @author GPT-6 Luna
 */
function externalAssetQuotePath(
  assetId: number,
  association: ExternalAssetAssociation,
): string {
  return `/api/asset/${encodeURIComponent(assetId)}/external-asset/${encodeURIComponent(association.source)}`
    + `/${encodeURIComponent(association.exchangeId)}/${encodeURIComponent(association.ticker)}/quote`;
}

/** Creates a valid mocked quote response for one persisted provider association.
 *
 * @param association - Provider ticker and exchange the quote response must match.
 * @param lastCloseQuote - Decimal quote string returned by the provider route.
 * @returns A valid response body suitable for a mocked quote GET.
 *
 * @author GPT-6 Luna
 */
function externalQuoteResponse(
  association: ExternalAssetAssociation,
  lastCloseQuote: string,
): {
  currency: string;
  exchangeId: string;
  lastCloseDate: string;
  lastCloseQuote: string;
  ticker: string;
} {
  return {
    ticker: association.ticker,
    exchangeId: association.exchangeId,
    currency: 'USD',
    lastCloseQuote,
    lastCloseDate: '2024-01-01T00:00:00Z',
  };
}

/** Asserts the displayed market price, quantity, and total for one allocation row.
 *
 * @param row - Allocation row containing the financial inputs.
 * @param marketPrice - Expected normalized market-price display value.
 * @param quantity - Expected quantity display value.
 * @param total - Expected total-market-value display value.
 *
 * @author GPT-6 Luna
 */
async function expectFinancialValues(
  row: Locator,
  marketPrice: string,
  quantity: string,
  total: string,
): Promise<void> {
  await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue(marketPrice);
  await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue(quantity);
  await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue(total);
}

/** Confirms the quote button has returned to its idle, accessible state.
 *
 * @param quoteAction - Quote action whose button, icon, and spinner states are verified.
 *
 * @author GPT-6 Luna
 */
async function expectQuoteActionIdle(quoteAction: Locator): Promise<void> {
  await expect(quoteAction).toBeEnabled();
  await expect(quoteAction).not.toHaveAttribute('aria-busy', 'true');
  await expect(quoteAction.locator('[data-quote-icon]')).toBeVisible();
  await expect(quoteAction.locator('[data-quote-spinner]')).toBeHidden();
}

/** Opens the root shell and confirms the shared portfolio navigation is ready.
 *
 * @param page - Browser page that hosts the SPA.
 *
 * @author GPT-6 Luna
 */
async function expectRootShell(page: Page): Promise<void> {
  await expect(page).toHaveURL(/\/$/);
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
}

/** Verifies the portfolio history management shell before opening an observation form.
 *
 * @param page - Browser page displaying portfolio history management.
 * @param portfolio - Portfolio whose context should be selected.
 * @param observationCount - Number of saved observations expected beside the new-observation item.
 *
 * @author GPT-6 Luna
 */
async function expectPortfolioHistoryManagement(
  page: Page,
  portfolio: SeededPortfolio,
  observationCount = 0,
): Promise<void> {
  await expect(page).toHaveURL(new RegExp(`/portfolio/${portfolio.id}/history/manage$`));
  await expect(page.locator('#portfolio-context .badge.text-bg-secondary')).toHaveText(portfolio.name);
  const returnButton = page.locator(
    'button[onclick="portfolioHistoryManagement.navigateToPortfolioAllocationViewing()"]',
  );
  await expect(returnButton).toBeVisible();
  await expect(returnButton.locator('.bi-pie-chart-fill')).toBeVisible();

  const managementCard = page.locator('.card').filter({ hasText: 'Manage portfolio allocation data' });
  await expect(managementCard).toBeVisible();
  await expect(managementCard.getByText('Manage portfolio allocation data', { exact: true })).toBeVisible();

  const accordion = page.locator('#accordion-portfolio-history-management');
  await expect(accordion).toBeVisible();
  await expect(accordion.locator(':scope > .accordion-item'))
    .toHaveCount(1 + observationCount);
  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  await expect(newObservationItem.getByRole('button')).toHaveClass(/collapsed/);
  await expect(newObservationItem.getByRole('textbox', { name: 'Time tag' }))
    .toBeVisible();
}

/** Adds one allocation row to the new-observation form and waits for its financial inputs.
 *
 * @param page - Browser page that contains the new observation form.
 * @param form - Expanded observation form receiving the row.
 * @param index - Zero-based allocation row index.
 * @returns Locator for the newly added row.
 *
 * @author GPT-6 Luna
 */
async function addAllocationRow(page: Page, form: Locator, index: number): Promise<Locator> {
  await form.locator('tfoot button.btn-secondary[type="button"]').click();
  const row = page.locator(`#portfolio-history-management-form-0-row-${index}`);
  await expect(row).toBeVisible();
  await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveAttribute('data-financial-input-bound', 'true');
  await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveAttribute('data-financial-input-bound', 'true');
  return row;
}

/** Selects one known asset from the shared suggestions and confirms its GET lookup succeeds.
 *
 * @param page - Browser page receiving the lookup.
 * @param row - Allocation row containing the asset autocomplete.
 * @param ticker - Exact seeded ticker to select.
 *
 * @author GPT-6 Luna
 */
async function selectAssetFromAutocomplete(page: Page, row: Locator, ticker: string): Promise<void> {
  await expect(page.locator(`#datalist-assets option[value="${ticker}"]`)).toBeAttached();
  const searchInput = row.getByRole('combobox', { name: 'Asset', exact: true });
  await searchInput.click();
  await searchInput.fill(ticker);

  const lookupRequests: string[] = [];
  const recordLookupRequest = (request: Request): void => {
    if(request.method() === 'GET' && new URL(request.url()).pathname === `/api/asset/${ticker}`) {
      lookupRequests.push(request.url());
    }
  };
  page.on('request', recordLookupRequest);
  try {
    const responsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'GET'
        && new URL(response.url()).pathname === `/api/asset/${ticker}`;
    });
    const optionLabel = await page.locator(`#datalist-assets option[value="${ticker}"]`).textContent();
    await row.getByRole('option', { name: optionLabel?.trim(), exact: true }).click();
    const response = await responsePromise;
    expect(response.status()).toBe(200);
    expect(lookupRequests).toHaveLength(1);
    await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveValue(ticker);
  } finally {
    page.off('request', recordLookupRequest);
  }
}

/** Uses the row's literal search action and checks the expected lookup status.
 *
 * @param page - Browser page receiving the lookup.
 * @param row - Allocation row containing the asset search field and action button.
 * @param ticker - Ticker sent to the literal lookup endpoint.
 * @param status - Expected HTTP response status.
 *
 * @author GPT-6 Luna
 */
async function searchForAsset(page: Page, row: Locator, ticker: string, status: number): Promise<void> {
  await row.getByRole('combobox', { name: 'Asset', exact: true }).fill(ticker);
  const responsePromise = page.waitForResponse((response) => {
    return response.request().method() === 'GET'
      && new URL(response.url()).pathname === `/api/asset/${ticker}`;
  });
  await row.locator('[data-asset-action-button]').click();
  expect((await responsePromise).status()).toBe(status);
}

/** Asserts the committed asset identity rendered in an allocation row.
 *
 * @param row - Allocation row whose selected asset is checked.
 * @param asset - Persisted asset expected in the row.
 *
 * @author GPT-6 Luna
 */
async function expectExistingAsset(row: Locator, asset: SeededAsset): Promise<void> {
  await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveAttribute('readonly', '');
  await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveValue(asset.ticker);
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveAttribute('readonly', '');
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveValue(asset.name);
  await expect(row.locator('input[type="hidden"]').first()).toHaveValue(asset.id.toString());
  await expect(row.getByText('* Creating new asset', { exact: true })).toBeHidden();
  await expect(row.locator('[data-asset-action-button] .bi-x-circle')).toBeVisible();
}

/** Fills an allocation using quantity and market price and verifies the calculated total.
 *
 * @param row - Allocation row whose financial values will be filled.
 * @param quantity - Quantity to enter.
 * @param marketPrice - Market price to enter.
 * @param totalMarketValue - Expected formatted total after calculation.
 *
 * @author GPT-6 Luna
 */
async function fillCalculatedValues(
  row: Locator,
  quantity: string,
  marketPrice: string,
  totalMarketValue: string,
): Promise<void> {
  await row.getByRole('spinbutton', { name: 'Quantity' }).fill(quantity);
  const marketPriceInput = row.getByRole('textbox', { name: 'Market price' });
  await marketPriceInput.fill(marketPrice);
  await marketPriceInput.blur();
  await expect(marketPriceInput).toHaveValue(`${marketPrice}.00000000`);
  await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue(totalMarketValue);
}

/** Fills a direct total while leaving quantity and market price blank.
 *
 * @param row - Allocation row whose total will be filled.
 * @param totalMarketValue - Total-market-value input value.
 *
 * @author GPT-6 Luna
 */
async function fillDirectTotalMarketValue(row: Locator, totalMarketValue: string): Promise<void> {
  const totalInput = row.getByRole('textbox', { name: 'Total market value' });
  await totalInput.fill(totalMarketValue);
  await totalInput.blur();
  await expect(totalInput).toHaveValue(totalMarketValue.includes(',') ? `${totalMarketValue}.00` : totalMarketValue);
}

/** Verifies the save-success toast for the explicit observation submission.
 *
 * @param page - Browser page showing the saved observation notification.
 *
 * @author GPT-6 Luna
 */
async function expectSuccessNotification(page: Page): Promise<void> {
  const toast = page.locator('#toast-notification-container .toast[role="alert"]').filter({
    hasText: 'Portfolio observation data saved successfully.',
  }).last();
  await expect(toast).toBeVisible();
  await expect(toast).toHaveClass(/text-bg-success/);
  await expect(toast.locator('.toast-header strong')).toHaveText('Success');
  await expect(toast.locator('.toast-body')).toHaveText('Portfolio observation data saved successfully.');
}

/** Finds one persisted observation from its portfolio-scoped time tag.
 *
 * @param database - Isolated test database.
 * @param portfolio - Portfolio containing the observation.
 * @param timeTag - Unique observation tag to find.
 * @returns The matching persisted observation.
 *
 * @author GPT-6 Luna
 */
async function findObservationByTag(
  database: E2eDatabase,
  portfolio: SeededPortfolio,
  timeTag: string,
): Promise<PersistedObservation> {
  const observations = await database.query<PersistedObservation>(
    `SELECT DISTINCT ot.id, ot.observation_timestamp, ot.observation_time_tag
     FROM public.portfolio_allocation_obs_time ot
     JOIN public.portfolio_allocation_fact paf ON paf.observation_time_id = ot.id
     WHERE paf.portfolio_id = $1 AND ot.observation_time_tag = $2`,
    [portfolio.id, timeTag],
  );
  expect(observations).toHaveLength(1);
  return observations[0];
}

/** Verifies the new item followed by saved observation items in their displayed order.
 *
 * @param page - Browser page showing the management accordion.
 * @param observations - Persisted observations expected after the new-observation item.
 *
 * @author GPT-6 Luna
 */
async function expectManagementObservationOrder(
  page: Page,
  observations: readonly PersistedObservation[],
): Promise<void> {
  const accordion = page.locator('#accordion-portfolio-history-management');
  await expect(accordion.locator(':scope > .accordion-item')).toHaveCount(1 + observations.length);
  await expect(page.locator('#portfolio-history-management-container-0 #portfolio-history-management-trigger-0 > button'))
    .toHaveClass(/collapsed/);
  for (const [index, observation] of observations.entries()) {
    const item = accordion.locator(':scope > .accordion-item').nth(index + 1);
    await expect(item).toHaveAttribute('id', `portfolio-history-management-container-${observation.id}`);
    const button = item.getByRole('button', { name: observation.observation_time_tag, exact: true });
    await expect(button).toBeVisible();
    await expect(button).toHaveClass(/collapsed/);
    await expect(item.locator(`#portfolio-history-management-${observation.id}`)).not.toHaveClass(/\bshow\b/);
  }
}

/** Creates a one-shot gate that lets a test control when its mocked quote response is released.
 *
 * @returns A signal for request start and a promise controlled by the release callback.
 *
 * @author GPT-6 Luna
 */
function createQuoteResponseGate(): QuoteResponseGate {
  let signalStarted = (): void => undefined;
  let releaseResponse = (): void => undefined;
  const started = new Promise<void>((resolve) => { signalStarted = resolve; });
  const wait = new Promise<void>((resolve) => { releaseResponse = resolve; });

  return {
    started,
    wait,
    markStarted: signalStarted,
    release: releaseResponse,
  };
}
