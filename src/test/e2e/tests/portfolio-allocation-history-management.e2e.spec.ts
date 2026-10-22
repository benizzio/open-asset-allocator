/**
 * Covers portfolio allocation history management and direct history navigation.
 *
 * Scenario 3 creates one observation through the allocation-history management
 * form, verifies the rendered history chart, and checks PostgreSQL persistence.
 * Scenario 3.1 verifies the history routes with direct browser URLs without
 * changing persisted data.
 * Scenarios 3.2–3.10 cover date validation and conversion, responsive layout,
 * focus-driven accordion behavior, repeated field interactions, and behavior after refresh.
 * Scenario 17 covers dated headers, stable order, local-timezone formatting, escaped tags,
 * post-submission refreshes, and saved observations outside the current list window.
 * Scenario 19 covers pencil editing, lazy detail loading, exact timestamp preservation, and edited-date ordering.
 * Scenario 5 modifies a seeded observation, creates a second observation, and
 * verifies both browser state and PostgreSQL persistence, including the stored timestamp.
 *
 * @author OpenCode
 * @author GPT-6 Sol
 * @author GPT-6 Luna
 * @author benizzio
 */
import type { Browser, Locator, Page, Request, Route } from '@playwright/test';
import { exerciseAccordionFieldInteractions } from '../support/accordion-field-interactions';
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';
import {
  clickCanvasPoint,
  expectChartTooltip,
  expectLatestCanvasPatternState,
  expectLatestCanvasTextContains,
  expectLatestCanvasTextSet,
  findDoughnutSlicePointByTooltip,
  getDoughnutCenterPoint,
  getDoughnutSlicePoint,
  installCanvasTextRecorder,
} from '../support/doughnut-chart';

const DEFAULT_ALLOCATION_STRUCTURE = {
  hierarchy: [
    { name: 'Assets', field: 'assetTicker' },
    { name: 'Classes', field: 'class' },
  ],
} as const;

const PORTFOLIO_NAME = 'E2E History Portfolio';
const MODIFICATION_PORTFOLIO_NAME = 'E2E History Modification Portfolio';
const OBSERVATION_TIME_TAG = 'E2E_OBS_001';
const SECOND_OBSERVATION_TIME_TAG = 'E2E_OBS_002';
const FIXED_OBSERVATION_TIMESTAMP = '2026-01-15T12:00:00.000Z';
const NEW_BOND_TICKER = 'BOND-A';
const NEW_BOND_NAME = 'Bond Alpha';
const EXISTING_BOND_TICKER = 'BOND-B';
const EXISTING_BOND_NAME = 'Bond Beta';
const NEW_STOCK_TICKER = 'STOCK-B';
const NEW_STOCK_NAME = 'Stock Beta';
const EXISTING_STOCK_TICKER = 'STOCK-A';
const EXISTING_STOCK_NAME = 'Stock Alpha';

const PORTFOLIO_NAVIGATION_OPTIONS = [
  { id: 'nav-radio-portfolio', name: 'Portfolio' },
  { id: 'nav-radio-allocation-plan', name: 'Allocation Plan' },
  { id: 'nav-radio-allocation-map', name: 'Allocation Map' },
] as const;

type NavigationOptionName = typeof PORTFOLIO_NAVIGATION_OPTIONS[number]['name'];

type SeededPortfolio = {
  id: number;
  name: string;
};

type SeededAsset = {
  id: number;
  name: string;
  ticker: string;
};

type PersistedObservation = {
  id: number;
  observation_timestamp: Date;
  observation_time_tag: string;
};

type PersistedAllocation = {
  asset_id: number;
  asset_market_price: string;
  asset_quantity: string;
  cash_reserve: boolean;
  class: string;
  name: string;
  observation_time_id: number;
  ticker: string;
  total_market_value: string;
};

type ExpectedObservationRow = {
  assetClass: string;
  cashReserve: boolean;
  marketPrice: string;
  name: string;
  quantity: string;
  ticker: string;
  total: string;
};

type SeededHistoryModificationData = {
  assets: Record<string, SeededAsset>;
  observation: PersistedObservation;
  portfolio: SeededPortfolio;
};

test.describe('portfolio allocation history management', () => {
  test.describe('scenario 14: asset search autocomplete feedback and create action', () => {
    /** Verifies all matching assets remain available in a bounded, scrollable list.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.1: filters asset suggestions and allows scrolling', async ({ database, page }) => {
      const { row, assetSearchInput, suggestions } = await prepareAssetSearchAutocompleteRow(database, page);

      await expect(assetSearchInput).toHaveAttribute('aria-expanded', 'true');
      await expect(suggestions.getByRole('option')).toHaveCount(15);

      await assetSearchInput.fill('auto-1');
      await expect(suggestions.getByRole('option')).toHaveCount(2);
      await expect(suggestions.getByRole('option', { name: 'E2E:AUTO-10 - Autocomplete Asset 10', exact: true })).toBeVisible();
      await expect(suggestions.getByRole('option', { name: 'E2E:AUTO-11 - Autocomplete Asset 11', exact: true })).toBeVisible();

      await assetSearchInput.fill('E2E:');
      await expect(suggestions.getByRole('option')).toHaveCount(12);
      await expect(suggestions.evaluate(element => element.scrollHeight > element.clientHeight)).resolves.toBe(true);

      await assetSearchInput.fill('NO-MATCH');
      await expect(row.getByRole('option', { name: 'No matching assets - create new', exact: true })).toBeVisible();
      await expect(row.getByRole('status')).toHaveText('No matching assets. Select the create option to continue.');
    });

    /** Verifies a refresh of the prefetched catalog removes visible stale options from an open list.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.2: refreshes suggestions after the ticker catalog reloads', async ({ database, page }) => {
      const { suggestions, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);

      await page.locator('#datalist-assets').evaluate(element => {
        const staleOption = document.createElement('option');
        staleOption.value = 'E2E:STALE';
        staleOption.textContent = 'E2E:STALE';
        element.append(staleOption);
      });
      await assetSearchInput.fill('E2E:');
      await expect(suggestions.getByRole('option', { name: 'E2E:STALE', exact: true })).toBeVisible();

      const datalistRefresh = page.waitForResponse(response => {
        return response.request().method() === 'GET' && new URL(response.url()).pathname === '/api/asset';
      });
      await page.evaluate(() => {
        const assetComposedColumnsInput =
          (window as unknown as Record<string, { loadDatalists(): void }>)['AssetComposedColumnsInput'];
        assetComposedColumnsInput.loadDatalists();
      });
      await datalistRefresh;

      await expect(page.locator('#datalist-assets option[value="E2E:STALE"]')).toHaveCount(0);
      await expect(suggestions.getByRole('option', { name: 'E2E:STALE', exact: true })).toHaveCount(0);
      await expect(suggestions.getByRole('option')).toHaveCount(12);
    });

    /** Verifies pointer selection directly starts lookup and populates the existing asset fields.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.3: selecting a suggestion immediately looks up the existing asset', async ({ database, page }) => {
      const { assets, row, assetSearchInput, suggestions } = await prepareAssetSearchAutocompleteRow(database, page);
      const ticker = 'E2E:AUTO-11';
      const lookupRequests: string[] = [];
      const recordLookupRequest = (request: Request): void => {
        if(request.method() === 'GET' && new URL(request.url()).pathname === `/api/asset/${ticker}`) {
          lookupRequests.push(request.url());
        }
      };
      page.on('request', recordLookupRequest);
      try {
        const lookupResponse = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${ticker}`;
        });

        await suggestions.getByRole('option', { name: `${ticker} - ${assets[ticker].name}`, exact: true }).click();

        expect((await lookupResponse).status()).toBe(200);
        expect(lookupRequests).toHaveLength(1);
        await expectExistingAsset(row, assets[ticker]);
        await expect(row.locator('[data-asset-search-input]')).toBeHidden();
        await expect(row.locator('[data-asset-search-input]')).toHaveAttribute('aria-expanded', 'false');
      } finally {
        page.off('request', recordLookupRequest);
      }
    });

    /** Verifies Arrow navigation and Enter select a suggestion and resolve it through the API.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.4: selects the highlighted suggestion with the keyboard', async ({ database, page }) => {
      const { assets, row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const ticker = 'E2E:AUTO-00';

      await assetSearchInput.fill(ticker);
      await assetSearchInput.press('ArrowDown');
      const activeOptionId = await assetSearchInput.getAttribute('aria-activedescendant');
      expect(activeOptionId).toBeTruthy();
      const activeOption = row.locator(`#${activeOptionId}`);
      await expectOptionHighlighted(activeOption);

      const lookupResponse = page.waitForResponse(response => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === `/api/asset/${ticker}`;
      });
      await assetSearchInput.press('Enter');

      expect((await lookupResponse).status()).toBe(200);
      await expectExistingAsset(row, assets[ticker]);
    });

    /** Verifies Enter without a selected suggestion preserves manual lookup and new-asset mode.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.5: keeps manual search for unmatched tickers', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const ticker = 'NO-SUCH-ASSET';

      await assetSearchInput.fill(ticker);
      await expect(row.getByRole('option', { name: 'No matching assets - create new', exact: true })).toBeVisible();
      const lookupResponse = page.waitForResponse(response => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === `/api/asset/${ticker}`;
      });
      await assetSearchInput.press('Enter');

      expect((await lookupResponse).status()).toBe(404);
      await expect(row.getByText('* Creating new asset', { exact: true })).toBeVisible();
      await expect(row.getByRole('textbox', { name: 'Asset name' })).toBeVisible();
    });

    /** Verifies the new-asset reset action clears the row and reopens asset suggestions.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.6: resets a new-asset row to asset search', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const missingAssetTicker = 'NO-SUCH-ASSET';
      const missingAssetResponse = page.waitForResponse(response => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === `/api/asset/${missingAssetTicker}`;
      });

      await assetSearchInput.fill(missingAssetTicker);
      await assetSearchInput.press('Enter');
      expect((await missingAssetResponse).status()).toBe(404);
      await row.locator('[data-asset-action-button]').click();
      await expect(assetSearchInput).toHaveValue('');
      await expect(assetSearchInput).toHaveAttribute('aria-expanded', 'true');
      await assetSearchInput.press('Escape');
      await expect(assetSearchInput).toHaveAttribute('aria-expanded', 'false');
    });

    /** Verifies pointer hover highlights one row without starting an asset lookup.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.7: visibly highlights a hovered asset without selecting it', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const ticker = 'E2E:AUTO-02';
      const lookupTracking = trackAssetLookupRequests(page, ticker);

      try {
        await assetSearchInput.fill(ticker);
        const option = row.getByRole('option', { name: `${ticker} - Autocomplete Asset 2`, exact: true });
        await option.hover();

        await expectOptionHighlighted(option);
        await expect(assetSearchInput).toHaveValue(ticker);
        await expect(row.locator('input[type="hidden"][data-null-if-empty][name$="[assetId]"]')).toHaveValue('');
        expect(lookupTracking.requests).toHaveLength(0);
      } finally {
        lookupTracking.stop();
      }
    });

    /** Verifies selecting the create action resolves the unmatched ticker into new-asset mode.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.8: creates a new-asset draft from the pointer-selected action', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const ticker = 'E2E:NEW-POINTER';
      const lookupTracking = trackAssetLookupRequests(page, ticker);

      try {
        await assetSearchInput.fill(ticker);
        const createOption = row.getByRole('option', { name: 'No matching assets - create new', exact: true });
        await expect(createOption).toBeVisible();
        const lookupResponse = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${ticker}`;
        });

        await createOption.click();

        expect((await lookupResponse).status()).toBe(404);
        expect(lookupTracking.requests).toHaveLength(1);
        await expectNewAssetMode(row, ticker);
        await expect(database.query('SELECT id FROM public.asset WHERE ticker = $1', [ticker])).resolves.toEqual([]);
      } finally {
        lookupTracking.stop();
      }
    });

    /** Verifies keyboard navigation can select the create action and focus the required name field.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.9: selects the create action with the keyboard', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const ticker = 'E2E:NEW-KEYBOARD';
      const lookupTracking = trackAssetLookupRequests(page, ticker);

      try {
        await assetSearchInput.fill(ticker);
        await assetSearchInput.press('ArrowDown');
        const createOption = row.getByRole('option', { name: 'No matching assets - create new', exact: true });
        await expectOptionHighlighted(createOption);
        const lookupResponse = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${ticker}`;
        });

        await assetSearchInput.press('Enter');

        expect((await lookupResponse).status()).toBe(404);
        expect(lookupTracking.requests).toHaveLength(1);
        await expectNewAssetMode(row, ticker);
      } finally {
        lookupTracking.stop();
      }
    });

    /** Verifies a touchscreen press visibly highlights and then selects an existing ticker.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.10: selects an existing ticker using a touchscreen tap', async ({ database, browser, browserName }) => {
      test.skip(browserName !== 'chromium', 'Touchscreen input is exercised with Chromium CDP.');
      const touchPage = await prepareTouchAutocompleteRow(database, browser);
      const ticker = 'E2E:AUTO-03';
      const lookupTracking = trackAssetLookupRequests(touchPage.page, ticker);
      const cdpSession = await touchPage.context.newCDPSession(touchPage.page);

      try {
        const option = touchPage.row.getByRole('option', { name: `${ticker} - Autocomplete Asset 3`, exact: true });
        await option.scrollIntoViewIfNeeded();
        const bounds = await option.boundingBox();
        if(!bounds) {
          throw new Error('The asset option has no visible bounds for touchscreen input.');
        }
        const x = bounds.x + bounds.width / 2;
        const y = bounds.y + bounds.height / 2;
        const lookupResponse = touchPage.page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${ticker}`;
        });

        await cdpSession.send('Input.dispatchTouchEvent', { type: 'touchStart', touchPoints: [{ x, y }] });
        await expectOptionHighlighted(option);
        await cdpSession.send('Input.dispatchTouchEvent', { type: 'touchEnd', touchPoints: [] });

        expect((await lookupResponse).status()).toBe(200);
        expect(lookupTracking.requests).toHaveLength(1);
        await expectExistingAsset(touchPage.row, touchPage.assets[ticker]);
      } finally {
        lookupTracking.stop();
        await cdpSession.detach().catch(() => undefined);
        await touchPage.context.close().catch(() => undefined);
      }
    });

    /** Verifies a touchscreen tap selects the create action and opens the new-asset fields.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.11: selects the create action using a touchscreen tap', async ({ database, browser, browserName }) => {
      test.skip(browserName !== 'chromium', 'Touchscreen taps are tested in Chromium.');
      const touchPage = await prepareTouchAutocompleteRow(database, browser);
      const ticker = 'E2E:NEW-TOUCH';
      const lookupTracking = trackAssetLookupRequests(touchPage.page, ticker);

      try {
        await touchPage.assetSearchInput.fill(ticker);
        const option = touchPage.row.getByRole('option', { name: 'No matching assets - create new', exact: true });
        await expect(option).toBeVisible();
        const lookupResponse = touchPage.page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${ticker}`;
        });

        await option.tap();

        expect((await lookupResponse).status()).toBe(404);
        expect(lookupTracking.requests).toHaveLength(1);
        await expectNewAssetMode(touchPage.row, ticker);
      } finally {
        lookupTracking.stop();
        await touchPage.context.close();
      }
    });

    /** Verifies a touch swipe scrolls the list without activating the option where it began.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.12: scrolls suggestions with touch without selecting a ticker', async ({ database, browser, browserName }) => {
      test.skip(browserName !== 'chromium', 'Native touch scrolling is exercised with Chromium CDP.');
      const touchPage = await prepareTouchAutocompleteRow(database, browser);
      const ticker = 'E2E:AUTO-00';
      const lookupTracking = trackAssetLookupRequests(touchPage.page, ticker);
      const cdpSession = await touchPage.context.newCDPSession(touchPage.page);

      try {
        const listbox = touchPage.row.getByRole('listbox', { name: 'Available assets' });
        const initialOption = touchPage.row.getByRole('option', { name: `${ticker} - Autocomplete Asset 0`, exact: true });
        await initialOption.scrollIntoViewIfNeeded();
        const bounds = await initialOption.boundingBox();
        if(!bounds) {
          throw new Error('The first asset option has no visible bounds for touchscreen scrolling.');
        }
        const x = bounds.x + bounds.width / 2;
        const startY = bounds.y + bounds.height / 2;

        await cdpSession.send('Input.dispatchTouchEvent', {
          type: 'touchStart',
          touchPoints: [{ x, y: startY }],
        });
        await expectOptionHighlighted(initialOption);

        for(let step = 1; step <= 6; step++) {
          await cdpSession.send('Input.dispatchTouchEvent', {
            type: 'touchMove',
            touchPoints: [{ x, y: startY - step * 18 }],
          });
        }

        await cdpSession.send('Input.dispatchTouchEvent', { type: 'touchEnd', touchPoints: [] });

        await expect.poll(() => listbox.evaluate(element => element.scrollTop)).toBeGreaterThan(0);
        await expect(initialOption).toHaveAttribute('aria-selected', 'false');
        await expect(touchPage.assetSearchInput).not.toHaveAttribute('readonly', '');
        expect(lookupTracking.requests).toHaveLength(0);
      } finally {
        lookupTracking.stop();
        await cdpSession.detach().catch(() => undefined);
        await touchPage.context.close().catch(() => undefined);
      }
    });

    /** Verifies an empty ticker does not offer an invalid create action.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 14.13: does not offer asset creation for an empty ticker', async ({ database, page }) => {
      const { row, assetSearchInput, suggestions } = await prepareAssetSearchAutocompleteRow(database, page);

      await assetSearchInput.fill('');

      await expect(suggestions.getByRole('option', { name: 'No matching assets - create new', exact: true }))
        .toHaveCount(0);
      await expect(suggestions.getByRole('option')).toHaveCount(15);
    });
  });

  test.describe('scenario 15: asset name search and committed ticker state', () => {
    /** Verifies ticker/name labels, duplicate-name results, long queries, and committed field serialization.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 15.1: searches full asset labels and commits only the selected ticker', async ({ database, page }) => {
      const { assets, row, assetSearchInput, suggestions } = await prepareAssetSearchAutocompleteRow(database, page);
      const firstTicker = 'SEARCH:NAME-ONE';
      const secondTicker = 'SEARCH:NAME-TWO';
      const label = 'North America: Equity & <Growth>';
      const datalistOption = page.locator(`#datalist-assets option[value="${firstTicker}"]`);

      await expect(datalistOption).toHaveText(`${firstTicker} - ${label}`);
      await expect(datalistOption).toHaveAttribute('value', firstTicker);
      await expect(datalistOption.locator('xpath=./*')).toHaveCount(0);

      await assetSearchInput.fill('aMeRiCa: equity & <gRoWtH>');
      await expect(suggestions.getByRole('option')).toHaveCount(2);
      await expect(suggestions.getByRole('option', { name: `${firstTicker} - ${label}`, exact: true })).toBeVisible();
      await expect(suggestions.getByRole('option', { name: `${secondTicker} - ${label}`, exact: true })).toBeVisible();

      const lookupResponse = page.waitForResponse(response => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === `/api/asset/${firstTicker}`;
      });
      await suggestions.getByRole('option', { name: `${firstTicker} - ${label}`, exact: true }).click();
      expect((await lookupResponse).status()).toBe(200);

      const committedTicker = row.getByRole('textbox', { name: 'Asset ticker', exact: true });
      await expect(committedTicker).toBeVisible();
      await expect(committedTicker).toHaveValue(firstTicker);
      await expect(row.locator('[data-asset-search-input]')).toBeHidden();
      await expect(row.locator('[data-asset-search-input]')).not.toHaveAttribute('name');
      await expect(row.locator(`input[name="allocations[0][assetTicker]"]`)).toHaveCount(1);
      await expect(committedTicker).toHaveValue(assets[firstTicker].ticker);
      const formDataEntries = await row.locator('xpath=ancestor::form').evaluate(form => {
        return Array.from(new FormData(form as HTMLFormElement).entries());
      });
      expect(formDataEntries).toContainEqual(['allocations[0][assetTicker]', firstTicker]);
      expect(formDataEntries.some(([, value]) => value === 'aMeRiCa: equity & <gRoWtH>')).toBe(false);

      await row.locator('[data-asset-action-button]').click();
      await assetSearchInput.fill('An Extremely Long Name for a Global Balanced Portfolio');
      await expect(assetSearchInput).toHaveValue('An Extremely Long Name for a Global Balanced Portfolio');
      await expect(suggestions.getByRole('option', {
        name: 'SEARCH:LONG - An Extremely Long Name for a Global Balanced Portfolio',
        exact: true,
      })).toBeVisible();
    });

    /** Preserves literal Enter lookup when matching suggestions are present but none is active.
     *
     * @author GPT-6 Luna
     * @author benizzio
     * @author GPT-6 Sol
     */
    test('scenario 15.2: keeps Enter lookup literal when suggestions match the query', async ({ database, page }) => {
      const { row, assetSearchInput, suggestions } = await prepareAssetSearchAutocompleteRow(database, page);

      await page.mouse.move(0, 0);
      await assetSearchInput.fill('AUTO');
      await expect(suggestions.getByRole('option')).toHaveCount(12);
      await expect(suggestions.locator('[aria-selected="true"]')).toHaveCount(0);
      await expect(assetSearchInput).not.toHaveAttribute('aria-activedescendant');

      const literalLookup = page.waitForResponse(response => {
        return response.request().method() === 'GET'
          && new URL(response.url()).pathname === '/api/asset/AUTO';
      });
      await assetSearchInput.press('Enter');

      expect((await literalLookup).status()).toBe(404);
      await expectNewAssetMode(row, 'AUTO');
      await expect(database.query('SELECT id FROM public.asset WHERE ticker = $1', ['AUTO'])).resolves.toEqual([]);
    });

    /** Rejects invalid literal ticker candidates without changing the query or issuing a lookup.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 15.3: rejects empty, malformed, and overlength ticker candidates', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const lookupRequests: string[] = [];
      const recordLookup = (request: Request): void => {
        if(request.method() === 'GET' && new URL(request.url()).pathname.startsWith('/api/asset/')) {
          lookupRequests.push(request.url());
        }
      };
      page.on('request', recordLookup);

      try {
        for(const candidate of ['', 'BAD TICKER', 'X'.repeat(41)]) {
          await assetSearchInput.fill(candidate);
          await row.locator('[data-asset-action-button]').click();

          await expect(assetSearchInput).toBeVisible();
          await expect(assetSearchInput).toHaveValue(candidate);
          await expect(row.locator('[data-asset-ticker-input]')).toBeHidden();
          await expect.poll(() => assetSearchInput.evaluate(input => (input as HTMLInputElement).validationMessage))
            .not.toBe('');
        }

        expect(lookupRequests).toHaveLength(0);
      } finally {
        page.off('request', recordLookup);
      }
    });

    /** Ensures an unresolved but valid name/ticker query cannot be submitted as a portfolio allocation.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 15.4: blocks history writes until an asset is committed', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const form = page.locator('#portfolio-history-management-form-0');
      const portfolioId = Number((await page.locator('input[name="portfolioId"]').first().inputValue()));
      const writeRequests: string[] = [];
      const recordWrite = (request: Request): void => {
        if(request.method() === 'POST') {
          writeRequests.push(request.url());
        }
      };
      page.on('request', recordWrite);

      try {
        await row.getByRole('combobox', { name: 'Class' }).fill('BONDS');
        await row.getByRole('textbox', { name: 'Total market value' }).fill('100');
        for(const query of ['', 'SEARCH:NAME-ONE']) {
          await assetSearchInput.fill(query);
          await form.locator('button[type="submit"]').click();

          await expect(row.locator('[data-asset-ticker-extra-error-message]'))
            .toHaveText('Reference an existing asset or create a new one');
          expect(writeRequests).toHaveLength(0);
          await expect(database.query(
            'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
            [portfolioId],
          )).resolves.toEqual([]);
        }
      } finally {
        page.off('request', recordWrite);
      }
    });

    /** Prevents form writes while an asset lookup is awaiting its response.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 15.5: blocks history writes while an asset lookup is pending', async ({ database, page }) => {
      const { row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const form = page.locator('#portfolio-history-management-form-0');
      const lookupTicker = 'E2E:PENDING-LOOKUP';
      const writeRequests: string[] = [];
      let releaseLookup: () => void = () => undefined;
      let notifyLookupStarted: () => void = () => undefined;
      const lookupGate = new Promise<void>(resolve => { releaseLookup = resolve; });
      const lookupStarted = new Promise<void>(resolve => { notifyLookupStarted = resolve; });
      const recordWrite = (request: Request): void => {
        if(request.method() === 'POST') {
          writeRequests.push(request.url());
        }
      };
      const routeLookup = async (route: Route): Promise<void> => {
        notifyLookupStarted();
        await lookupGate;
        await route.fulfill({
          status: 404,
          contentType: 'application/json',
          body: JSON.stringify({ errorMessage: 'Data not found', details: [] }),
        });
      };
      page.on('request', recordWrite);
      await page.route(`**/api/asset/${lookupTicker}`, routeLookup);

      try {
        await row.getByRole('combobox', { name: 'Class' }).fill('BONDS');
        await row.getByRole('textbox', { name: 'Total market value' }).fill('100');
        await assetSearchInput.fill(lookupTicker);
        await row.locator('[data-asset-action-button]').click();
        await lookupStarted;
        await expect(row.locator('[data-asset-search-autocomplete]'))
          .toHaveAttribute('data-asset-selection-state', 'pending');

        await form.locator('button[type="submit"]').click();
        await expect(row.locator('[data-asset-ticker-extra-error-message]'))
          .toHaveText('Asset lookup is still in progress');
        expect(writeRequests).toHaveLength(0);

        releaseLookup();
        await expectNewAssetMode(row, lookupTicker);
        await expect(database.query('SELECT id FROM public.asset WHERE ticker = $1', [lookupTicker]))
          .resolves.toEqual([]);
      } finally {
        releaseLookup();
        page.off('request', recordWrite);
        await page.unroute(`**/api/asset/${lookupTicker}`, routeLookup);
      }
    });

    /** Keeps button lookup literal, resolves numeric IDs to canonical tickers, and permits reset/reselection.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 15.6: uses literal button lookup and commits the ticker returned for an ID', async ({ database, page }) => {
      const { assets, row, assetSearchInput, suggestions } = await prepareAssetSearchAutocompleteRow(database, page);
      const matchingNameQuery = 'AUTO';
      const literalLookupRequests: string[] = [];
      const recordLookup = (request: Request): void => {
        if(request.method() === 'GET' && new URL(request.url()).pathname.startsWith('/api/asset/')) {
          literalLookupRequests.push(new URL(request.url()).pathname);
        }
      };
      page.on('request', recordLookup);

      try {
        await assetSearchInput.fill(matchingNameQuery);
        await expect(suggestions.getByRole('option')).toHaveCount(12);
        const literalLookup = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${matchingNameQuery}`;
        });
        await row.locator('[data-asset-action-button]').click();

        expect((await literalLookup).status()).toBe(404);
        await expectNewAssetMode(row, matchingNameQuery);
        await row.locator('[data-asset-action-button]').click();

        const asset = assets['E2E:AUTO-00'];
        const numericId = asset.id.toString();
        await assetSearchInput.fill(numericId);
        const idLookup = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${numericId}`;
        });
        await row.locator('[data-asset-action-button]').click();

        expect((await idLookup).status()).toBe(200);
        await expectExistingAsset(row, asset);
        await expect(row.locator('input[name="allocations[0][assetId]"]')).toHaveValue(numericId);
        await expect(row.locator('input[name="allocations[0][assetTicker]"]')).toHaveValue(asset.ticker);
        expect(literalLookupRequests).toEqual([`/api/asset/${matchingNameQuery}`, `/api/asset/${numericId}`]);

        await row.locator('[data-asset-action-button]').click();
        await expect(assetSearchInput).toBeVisible();
        await expect(assetSearchInput).toHaveValue('');
        await expect(row.locator('input[name="allocations[0][assetId]"]')).toHaveValue('');
        await expect(row.getByRole('textbox', { name: 'Asset name', exact: true })).toBeHidden();

        await assetSearchInput.fill('Autocomplete Asset 1');
        await expect(suggestions.getByRole('option', {
          name: 'E2E:AUTO-01 - Autocomplete Asset 1',
          exact: true,
        })).toBeVisible();
        const reselection = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === '/api/asset/E2E:AUTO-01';
        });
        await suggestions.getByRole('option', {
          name: 'E2E:AUTO-01 - Autocomplete Asset 1',
          exact: true,
        }).click();

        expect((await reselection).status()).toBe(200);
        await expectExistingAsset(row, assets['E2E:AUTO-01']);
      } finally {
        page.off('request', recordLookup);
      }
    });

    /** Restores the typed query after a failed lookup instead of committing it as a new ticker.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    test('scenario 15.7: restores search state after a non-not-found lookup failure', async ({ database, page }) => {
      const { assets, row, assetSearchInput } = await prepareAssetSearchAutocompleteRow(database, page);
      const failedTicker = 'E2E:AUTO-00';
      const failedLookupRoute = async (route: Route): Promise<void> => {
        await route.fulfill({
          status: 500,
          contentType: 'application/json',
          body: JSON.stringify({ errorMessage: 'Temporary lookup failure', details: [] }),
        });
      };
      await page.route(`**/api/asset/${failedTicker}`, failedLookupRoute);

      try {
        await assetSearchInput.fill(failedTicker);
        const failedLookup = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${failedTicker}`;
        });
        await row.locator('[data-asset-action-button]').click();

        expect((await failedLookup).status()).toBe(500);
        await expect(row.locator('[data-asset-search-autocomplete]'))
          .toHaveAttribute('data-asset-selection-state', 'search');
        await expect(assetSearchInput).toBeVisible();
        await expect(assetSearchInput).toHaveValue(failedTicker);
        await expect(row.locator('[data-asset-ticker-input]')).toBeHidden();
        await expect(row.locator('input[name="allocations[0][assetId]"]')).toHaveValue('');

        await page.unroute(`**/api/asset/${failedTicker}`, failedLookupRoute);
        await assetSearchInput.fill('E2E:AUTO-00');
        const recoveredLookup = page.waitForResponse(response => {
          return response.request().method() === 'GET'
            && new URL(response.url()).pathname === `/api/asset/${failedTicker}`;
        });
        await row.locator('[data-asset-action-button]').click();

        expect((await recoveredLookup).status()).toBe(200);
        await expectExistingAsset(row, assets[failedTicker]);
      } finally {
        await page.unroute(`**/api/asset/${failedTicker}`, failedLookupRoute);
      }
    });
  });

  test('scenario 3: creates and explores a portfolio allocation observation', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedPortfolioHistoryData(database);
    await installCanvasTextRecorder(page);

    await page.goto('/');
    await expectRootShell(page);

    await navigateByClick(page, '/portfolios', page.getByRole('link', { name: 'Portfolios', exact: true }));
    await expectPortfolioList(page, seededData.portfolio);

    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}`,
      portfolioCardFor(page, seededData.portfolio),
    );
    await expectPortfolioContext(page, seededData.portfolio);
    await expectPortfolioNavigation(page, undefined);

    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/history`,
      portfolioNavigationLabel(page, 'Portfolio'),
    );
    await expectPortfolioHistory(page, seededData.portfolio, true);

    const managementButton = page.locator(
      '#accordion-portfolio-history > button[onclick="navigateToPortfolioAllocationManagement()"]',
    );
    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/history/manage`,
      managementButton,
    );
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const newObservationItem = page.locator('#portfolio-history-management-container-0');
    const newObservationTimeTag = newObservationItem.getByRole('textbox', { name: 'Time tag' });
    await newObservationTimeTag.fill(OBSERVATION_TIME_TAG);
    await expect(newObservationTimeTag).toHaveValue(OBSERVATION_TIME_TAG);
    await expect(newObservationItem.locator('#portfolio-history-management-0')).toHaveClass(/\bshow\b/);

    const newObservationForm = page.locator('#portfolio-history-management-form-0');
    await expect(newObservationForm).toBeVisible();
    await expectHistoryManagementForm(newObservationForm, 0);

    const bondAlphaRow = await addAllocationRow(page, newObservationForm, 0);
    await fillNewAssetAllocation(
      page,
      bondAlphaRow,
      NEW_BOND_TICKER,
      NEW_BOND_NAME,
      'BONDS',
      '1,000',
    );

    const bondBetaRow = await addAllocationRow(page, newObservationForm, 1);
    await fillExistingAssetAllocation(
      page,
      bondBetaRow,
      seededData.assets[EXISTING_BOND_TICKER],
      'BONDS',
      '30',
      '100',
      '3,000.00',
    );

    const stockAlphaRow = await addAllocationRow(page, newObservationForm, 2);
    await fillExistingAssetDirectValueAllocation(
      page,
      stockAlphaRow,
      seededData.assets[EXISTING_STOCK_TICKER],
      'STOCKS',
      '2,000',
    );

    const stockBetaRow = await addAllocationRow(page, newObservationForm, 3);
    await fillNewAssetCalculatedAllocation(
      page,
      stockBetaRow,
      NEW_STOCK_TICKER,
      NEW_STOCK_NAME,
      'STOCKS',
      '40',
      '100',
    );

    await expectNewObservationRows(newObservationForm);

    const selectedObservationDate = await newObservationItem.getByLabel('Observation date').inputValue();
    const expectedObservationTimestamp = await localMidnightTimestamp(page, selectedObservationDate);
    const saveRequestPromise = page.waitForRequest((request) => {
      return request.method() === 'POST'
        && new URL(request.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`;
    });
    const saveResponsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'POST'
        && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`;
    });
    await newObservationForm.locator('tfoot button.btn-primary[type="submit"]').click();
    const saveRequest = await saveRequestPromise;
    const saveResponse = await saveResponsePromise;
    expect(saveResponse.status()).toBe(204);
    expect(JSON.parse(saveRequest.postData() ?? 'null').observationTimestamp.timestamp)
      .toBe(expectedObservationTimestamp);

    const persistedObservation = await expectPersistedPortfolioHistory(
      database,
      seededData.portfolio,
      seededData.assets,
      expectedObservationTimestamp,
    );

    await expectSuccessNotification(page);
    await expectReloadedManagementItems(page, persistedObservation);

    const existingObservationItem = page.locator(
      `#portfolio-history-management-container-${persistedObservation.id}`,
    );
    await existingObservationItem.getByRole('button', {
      name: await expectedObservationHeaderName(page, persistedObservation),
      exact: true,
    }).click();

    const existingObservationForm = page.locator(
      `#portfolio-history-management-form-${persistedObservation.id}`,
    );
    await expect(existingObservationForm.locator('tbody tr')).toHaveCount(4);
    expect(new Date(await existingObservationForm.locator('[name="observationTimestamp.timestamp"]').inputValue()))
      .toEqual(persistedObservation.observation_timestamp);
    await expectReloadedObservationRows(existingObservationForm, persistedObservation.id);
    await expectHistoryManagementForm(existingObservationForm, persistedObservation.id);

    const returnButton = page.locator(
      'button[onclick="portfolioHistoryManagement.navigateToPortfolioAllocationViewing()"]',
    );
    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/history`,
      returnButton,
    );

    await expectPortfolioHistory(page, seededData.portfolio, false);
    await expectRenderedPortfolioHistoryChart(page, persistedObservation.id);
    await exercisePortfolioHistoryChart(page, persistedObservation.id);
  });

  test('scenario 5: modifies portfolio allocation history and creates a second observation', async ({ database, page }) => {
    test.setTimeout(60_000);
    const seededData = await seedPortfolioHistoryModificationData(database);
    await installCanvasTextRecorder(page);

    await page.goto('/');
    await expectRootShell(page);

    await navigateByClick(page, '/portfolios', page.getByRole('link', { name: 'Portfolios', exact: true }));
    await expectPortfolioList(page, seededData.portfolio);

    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}`,
      portfolioCardFor(page, seededData.portfolio),
    );
    await expectPortfolioContext(page, seededData.portfolio);
    await expectPortfolioNavigation(page, undefined);

    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/history`,
      portfolioNavigationLabel(page, 'Portfolio'),
    );
    await expectPortfolioHistory(page, seededData.portfolio, false);
    await expectHistoryObservationOrder(page, [seededData.observation], seededData.observation.id);
    await expectRenderedPortfolioHistoryChart(page, seededData.observation.id);
    await exerciseInitialPortfolioHistoryChart(page, seededData.observation.id);

    const managementButton = page.locator(
      '#accordion-portfolio-history > button[onclick="navigateToPortfolioAllocationManagement()"]',
    );
    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/history/manage`,
      managementButton,
    );
    await expectPortfolioHistoryManagement(page, seededData.portfolio, 1);

    const existingObservationItem = page.locator(
      `#portfolio-history-management-container-${seededData.observation.id}`,
    );
    await existingObservationItem.getByRole('button', {
      name: await expectedObservationHeaderName(page, seededData.observation),
      exact: true,
    }).click();

    const existingObservationForm = page.locator(
      `#portfolio-history-management-form-${seededData.observation.id}`,
    );
    await expect(existingObservationForm).toBeVisible();
    await expectHistoryManagementForm(existingObservationForm, seededData.observation.id);
    await expectReadOnlyObservationRows(existingObservationForm, seededData.observation.id, [
      {
        assetClass: 'BONDS',
        cashReserve: false,
        marketPrice: '100.00000000',
        name: EXISTING_BOND_NAME,
        quantity: '30',
        ticker: EXISTING_BOND_TICKER,
        total: '3,000.00',
      },
      {
        assetClass: 'BONDS',
        cashReserve: false,
        marketPrice: '0.00000000',
        name: NEW_BOND_NAME,
        quantity: '0',
        ticker: NEW_BOND_TICKER,
        total: '1,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: false,
        marketPrice: '100.00000000',
        name: NEW_STOCK_NAME,
        quantity: '40',
        ticker: NEW_STOCK_TICKER,
        total: '4,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: false,
        marketPrice: '0.00000000',
        name: EXISTING_STOCK_NAME,
        quantity: '0',
        ticker: EXISTING_STOCK_TICKER,
        total: '2,000.00',
      },
    ]);

    const bondAlphaRow = existingObservationForm.locator('tbody tr').nth(1);
    await selectClassFromDatalist(page, bondAlphaRow, 'STOCKS');

    const bondBetaRow = existingObservationForm.locator('tbody tr').nth(0);
    const bondBetaQuantity = bondBetaRow.getByRole('spinbutton', { name: 'Quantity' });
    await bondBetaQuantity.fill('35');
    await bondBetaQuantity.blur();
    await expect(bondBetaRow.getByRole('textbox', { name: 'Market price' })).toHaveValue('100.00000000');
    await expect(bondBetaRow.getByRole('textbox', { name: 'Total market value' })).toHaveValue('3,500.00');

    const stockBetaRow = existingObservationForm.locator('tbody tr').nth(2);
    await stockBetaRow.getByRole('checkbox').check();

    await expectReadOnlyObservationRows(existingObservationForm, seededData.observation.id, [
      {
        assetClass: 'BONDS',
        cashReserve: false,
        marketPrice: '100.00000000',
        name: EXISTING_BOND_NAME,
        quantity: '35',
        ticker: EXISTING_BOND_TICKER,
        total: '3,500.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: false,
        marketPrice: '0.00000000',
        name: NEW_BOND_NAME,
        quantity: '0',
        ticker: NEW_BOND_TICKER,
        total: '1,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: true,
        marketPrice: '100.00000000',
        name: NEW_STOCK_NAME,
        quantity: '40',
        ticker: NEW_STOCK_TICKER,
        total: '4,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: false,
        marketPrice: '0.00000000',
        name: EXISTING_STOCK_NAME,
        quantity: '0',
        ticker: EXISTING_STOCK_TICKER,
        total: '2,000.00',
      },
    ]);

    const originalObservationTimestamp = await existingObservationForm
      .locator('[name="observationTimestamp.timestamp"]')
      .inputValue();
    expect(new Date(originalObservationTimestamp)).toEqual(seededData.observation.observation_timestamp);
    const updateRequestPromise = page.waitForRequest((request) => {
      return request.method() === 'POST'
        && new URL(request.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`;
    });
    const updateResponsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'POST'
        && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`;
    });
    await existingObservationForm.locator('tfoot button.btn-primary[type="submit"]').click();
    const updateRequest = await updateRequestPromise;
    const updateResponse = await updateResponsePromise;
    expect(updateResponse.status()).toBe(204);
    expect(JSON.parse(updateRequest.postData() ?? 'null').observationTimestamp.timestamp)
      .toBe(originalObservationTimestamp);
    const unchangedObservationTimestamp = await database.query<{ observation_timestamp: Date }>(
      'SELECT observation_timestamp FROM public.portfolio_allocation_obs_time WHERE id = $1',
      [seededData.observation.id],
    );
    expect(unchangedObservationTimestamp).toEqual([{
      observation_timestamp: seededData.observation.observation_timestamp,
    }]);

    await expectRoute(page, `/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectSuccessNotification(page);
    await expect(existingObservationItem.locator(`#portfolio-history-management-${seededData.observation.id}`))
      .toHaveClass(/\bshow\b/);
    await expectReadOnlyObservationRows(existingObservationForm, seededData.observation.id, [
      {
        assetClass: 'BONDS',
        cashReserve: false,
        marketPrice: '100.00000000',
        name: EXISTING_BOND_NAME,
        quantity: '35',
        ticker: EXISTING_BOND_TICKER,
        total: '3,500.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: true,
        marketPrice: '100.00000000',
        name: NEW_STOCK_NAME,
        quantity: '40',
        ticker: NEW_STOCK_TICKER,
        total: '4,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: false,
        marketPrice: '0.00000000',
        name: EXISTING_STOCK_NAME,
        quantity: '0',
        ticker: EXISTING_STOCK_TICKER,
        total: '2,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: false,
        marketPrice: '0.00000000',
        name: NEW_BOND_NAME,
        quantity: '0',
        ticker: NEW_BOND_TICKER,
        total: '1,000.00',
      },
    ]);

    const newObservationItem = page.locator('#portfolio-history-management-container-0');
    const newObservationTimeTag = newObservationItem.getByRole('textbox', { name: 'Time tag' });
    await newObservationTimeTag.fill(SECOND_OBSERVATION_TIME_TAG);
    await expect(newObservationTimeTag).toHaveValue(SECOND_OBSERVATION_TIME_TAG);
    await expect(newObservationItem.locator('#portfolio-history-management-0')).toHaveClass(/\bshow\b/);

    const newObservationForm = page.locator('#portfolio-history-management-form-0');
    await expect(newObservationForm).toBeVisible();
    await expectHistoryManagementForm(newObservationForm, 0);

    const secondBondAlphaRow = await addAllocationRow(page, newObservationForm, 0);
    await fillExistingAssetDirectValueAllocation(
      page,
      secondBondAlphaRow,
      seededData.assets[NEW_BOND_TICKER],
      'BONDS',
      '2,000',
    );
    await selectClassFromDatalist(page, secondBondAlphaRow, 'BONDS');

    const secondBondBetaRow = await addAllocationRow(page, newObservationForm, 1);
    await fillExistingAssetAllocation(
      page,
      secondBondBetaRow,
      seededData.assets[EXISTING_BOND_TICKER],
      'BONDS',
      '50',
      '100',
      '5,000.00',
    );
    await selectClassFromDatalist(page, secondBondBetaRow, 'BONDS');

    const secondStockAlphaRow = await addAllocationRow(page, newObservationForm, 2);
    await fillExistingAssetDirectValueAllocation(
      page,
      secondStockAlphaRow,
      seededData.assets[EXISTING_STOCK_TICKER],
      'STOCKS',
      '1,000',
    );
    await selectClassFromDatalist(page, secondStockAlphaRow, 'STOCKS');

    const secondStockBetaRow = await addAllocationRow(page, newObservationForm, 3);
    await fillExistingAssetAllocation(
      page,
      secondStockBetaRow,
      seededData.assets[NEW_STOCK_TICKER],
      'STOCKS',
      '20',
      '100',
      '2,000.00',
      true,
    );
    await selectClassFromDatalist(page, secondStockBetaRow, 'STOCKS');

    await expectEditableObservationRows(newObservationForm, [
      {
        assetClass: 'BONDS',
        cashReserve: false,
        marketPrice: '',
        name: NEW_BOND_NAME,
        quantity: '',
        ticker: NEW_BOND_TICKER,
        total: '2,000.00',
      },
      {
        assetClass: 'BONDS',
        cashReserve: false,
        marketPrice: '100.00000000',
        name: EXISTING_BOND_NAME,
        quantity: '50',
        ticker: EXISTING_BOND_TICKER,
        total: '5,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: false,
        marketPrice: '',
        name: EXISTING_STOCK_NAME,
        quantity: '',
        ticker: EXISTING_STOCK_TICKER,
        total: '1,000.00',
      },
      {
        assetClass: 'STOCKS',
        cashReserve: true,
        marketPrice: '100.00000000',
        name: NEW_STOCK_NAME,
        quantity: '20',
        ticker: NEW_STOCK_TICKER,
        total: '2,000.00',
      },
    ]);

    const createResponsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'POST'
        && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`;
    });
    await newObservationForm.locator('tfoot button.btn-primary[type="submit"]').click();
    const createResponse = await createResponsePromise;
    expect(createResponse.status()).toBe(204);

    const secondObservation = await findObservationByTag(
      database,
      seededData.portfolio,
      SECOND_OBSERVATION_TIME_TAG,
    );
    await expectRoute(page, `/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectSuccessNotification(page);
    await expectManagementObservationOrder(page, [secondObservation, seededData.observation]);

    const returnButton = page.locator(
      'button[onclick="portfolioHistoryManagement.navigateToPortfolioAllocationViewing()"]',
    );
    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/history`,
      returnButton,
    );

    await expectPortfolioHistory(page, seededData.portfolio, false);
    await expectHistoryObservationOrder(page, [secondObservation, seededData.observation], secondObservation.id);
    await exerciseModifiedPortfolioHistoryChart(page, secondObservation, seededData.observation);
    await expectPersistedModifiedPortfolioHistory(database, seededData, secondObservation);
  });

  test('scenario 3.1: navigates through portfolio history with direct URLs', async ({ database, page }) => {
    const seededData = await seedPortfolioHistoryData(database);

    await page.goto('/');
    await expectRootShell(page);

    await page.goto('/portfolios');
    await expectPortfolioList(page, seededData.portfolio);

    await page.goto(`/portfolio/${seededData.portfolio.id}`);
    await expectPortfolioContext(page, seededData.portfolio);
    await expectPortfolioNavigation(page, undefined);

    await page.goto(`/portfolio/${seededData.portfolio.id}/history`);
    await expectPortfolioHistory(page, seededData.portfolio, true);

    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
    await expect(database.query('SELECT asset_id FROM public.portfolio_allocation_fact')).resolves.toEqual([]);
  });

  test('scenario 3.2: requires an observation date and keeps the selected date after a failed save', async ({
    database,
    page,
  }) => {
    const seededData = await seedPortfolioHistoryData(database);
    const baselineObservation = await seedManagementObservation(
      database,
      seededData.portfolio,
      seededData.assets[EXISTING_BOND_TICKER],
      'E2E_OBS_FAILED_DATE_BASELINE',
      '2026-02-01T12:00:00.000Z',
    );
    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio, 1);
    await expectManagementObservationOrder(page, [baselineObservation]);

    const newObservationItem = page.locator('#portfolio-history-management-container-0');
    const timeTagInput = newObservationItem.getByRole('textbox', { name: 'Time tag' });
    const observationDateInput = newObservationItem.getByLabel('Observation date');
    const accordionToggle = newObservationItem.getByRole('button', { name: 'New observation details' });

    await timeTagInput.click();
    await timeTagInput.fill('E2E_OBS_FAILED_DATE');
    await expect(accordionToggle).toHaveAttribute('aria-expanded', 'true');
    await observationDateInput.click();
    await expect(accordionToggle).toHaveAttribute('aria-expanded', 'true');

    const form = page.locator('#portfolio-history-management-form-0');
    await observationDateInput.fill('');
    expect(await form.evaluate(element => (element as HTMLFormElement).checkValidity())).toBe(false);
    expect(await observationDateInput.evaluate(element => (element as HTMLInputElement).validity.valueMissing))
      .toBe(true);

    const savePath = `/api/portfolio/${seededData.portfolio.id}/history`;
    const saveRequests: Request[] = [];
    const trackSaveRequest = (request: Request): void => {
      if(request.method() === 'POST' && new URL(request.url()).pathname === savePath) {
        saveRequests.push(request);
      }
    };
    const failSave = async (route: Route): Promise<void> => {
      await route.fulfill({ status: 400, contentType: 'application/json', body: '{}' });
    };
    page.on('request', trackSaveRequest);
    try {
      await page.route(`**${savePath}`, failSave);
      await form.locator('tfoot button.btn-primary[type="submit"]').click();
      expect(saveRequests).toHaveLength(0);

      const selectedDate = '2026-04-20';
      await observationDateInput.fill(selectedDate);
      const failedSaveResponse = page.waitForResponse(response => {
        return response.request().method() === 'POST' && new URL(response.url()).pathname === savePath;
      });
      await form.locator('tfoot button.btn-primary[type="submit"]').click();
      expect((await failedSaveResponse).status()).toBe(400);
      expect(saveRequests).toHaveLength(1);
      await expect(observationDateInput).toHaveValue(selectedDate);
      await expectManagementObservationOrder(page, [baselineObservation], false);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([
        { id: baselineObservation.id },
      ]);
    } finally {
      page.off('request', trackSaveRequest);
      await page.unroute(`**${savePath}`, failSave);
    }
  });

  test.describe('scenario 3.3: serializes past and future dates in a positive UTC offset', () => {
    test.use({ timezoneId: 'Asia/Tokyo' });

    test('persists each selected date at Tokyo local midnight', async ({ database, page }) => {
      test.setTimeout(60_000);
      const seededData = await seedPortfolioHistoryData(database);
      await page.goto('/');
      await expectRootShell(page);
      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      await expectPortfolioHistoryManagement(page, seededData.portfolio);

      const dates = [
        { calendarDate: '2024-01-15', expectedTimestamp: '2024-01-14T15:00:00.000Z', timeTag: 'E2E_TOKYO_PAST' },
        { calendarDate: '2030-06-01', expectedTimestamp: '2030-05-31T15:00:00.000Z', timeTag: 'E2E_TOKYO_FUTURE' },
      ];

      for(const selectedDate of dates) {
        const persistedObservation = await submitObservationForDate(
          page,
          database,
          seededData.portfolio,
          seededData.assets[EXISTING_BOND_TICKER],
          selectedDate.calendarDate,
          selectedDate.timeTag,
          selectedDate.expectedTimestamp,
        );
        expect(persistedObservation.observation_timestamp).toEqual(new Date(selectedDate.expectedTimestamp));
      }
    });
  });

  test.describe('scenario 3.4: applies the selected date timezone offset across daylight-saving transitions', () => {
    test.use({ timezoneId: 'America/Los_Angeles' });

    test('persists local midnight before and after both 2026 transitions', async ({ database, page }) => {
      test.setTimeout(60_000);
      const seededData = await seedPortfolioHistoryData(database);
      await page.goto('/');
      await expectRootShell(page);
      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      await expectPortfolioHistoryManagement(page, seededData.portfolio);

      const dates = [
        { calendarDate: '2026-03-08', expectedTimestamp: '2026-03-08T08:00:00.000Z', timeTag: 'E2E_DST_SPRING_BEFORE' },
        { calendarDate: '2026-03-09', expectedTimestamp: '2026-03-09T07:00:00.000Z', timeTag: 'E2E_DST_SPRING_AFTER' },
        { calendarDate: '2026-11-01', expectedTimestamp: '2026-11-01T07:00:00.000Z', timeTag: 'E2E_DST_FALL_BEFORE' },
        { calendarDate: '2026-11-02', expectedTimestamp: '2026-11-02T08:00:00.000Z', timeTag: 'E2E_DST_FALL_AFTER' },
      ];

      for(const selectedDate of dates) {
        await submitObservationForDate(
          page,
          database,
          seededData.portfolio,
          seededData.assets[EXISTING_BOND_TICKER],
          selectedDate.calendarDate,
          selectedDate.timeTag,
          selectedDate.expectedTimestamp,
        );
      }
    });
  });

  test('scenario 3.5: keeps the observation fields beside the toggle until the header needs to wrap', async ({
    database,
    page,
  }) => {
    const seededData = await seedPortfolioHistoryData(database);
    await page.goto('/');
    await expectRootShell(page);
    await page.setViewportSize({ width: 1024, height: 900 });
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const newObservationItem = page.locator('#portfolio-history-management-container-0');
    const timeTagInput = newObservationItem.getByRole('textbox', { name: 'Time tag' });
    const observationDateInput = newObservationItem.getByLabel('Observation date');
    const accordionToggle = newObservationItem.getByRole('button', { name: 'New observation details' });
    const readHeaderLayout = async () => newObservationItem.locator('h2').evaluate(header => {
      const timeTag = header.querySelector<HTMLInputElement>('[aria-label="Time tag"]');
      const observationDate = header.querySelector<HTMLInputElement>('[aria-label="Observation date"]');
      const toggle = header.querySelector<HTMLButtonElement>('.accordion-button');

      if(!timeTag || !observationDate || !toggle) {
        throw new Error('The new-observation header is missing a field or accordion toggle.');
      }

      const timeTagBounds = timeTag.getBoundingClientRect();
      const observationDateBounds = observationDate.getBoundingClientRect();
      const toggleBounds = toggle.getBoundingClientRect();

      return {
        headerClientWidth: header.clientWidth,
        headerScrollWidth: header.scrollWidth,
        timeTag: { left: timeTagBounds.left, right: timeTagBounds.right, top: timeTagBounds.top, width: timeTagBounds.width },
        observationDate: {
          left: observationDateBounds.left,
          right: observationDateBounds.right,
          top: observationDateBounds.top,
          bottom: observationDateBounds.bottom,
          width: observationDateBounds.width,
        },
        toggle: { left: toggleBounds.left, right: toggleBounds.right, top: toggleBounds.top },
        headerBottom: header.getBoundingClientRect().bottom,
      };
    });

    await expect(timeTagInput).toBeVisible();
    await expect(observationDateInput).toBeVisible();

    const expectSideBySide = async (): Promise<void> => {
      const layout = await readHeaderLayout();
      expect(Math.abs(layout.timeTag.top - layout.observationDate.top)).toBeLessThan(2);
      expect(layout.observationDate.left).toBeGreaterThanOrEqual(layout.timeTag.right);
      expect(layout.toggle.left).toBeGreaterThanOrEqual(layout.observationDate.right);
      expect(layout.timeTag.width).toBeGreaterThan(layout.observationDate.width);
      expect(layout.headerScrollWidth).toBeLessThanOrEqual(layout.headerClientWidth);
    };

    await expectSideBySide();
    await page.setViewportSize({ width: 718, height: 900 });
    await expectSideBySide();
    await test.info().attach('observation-header-718px', {
      body: await page.screenshot(),
      contentType: 'image/png',
    });

    const expectFieldsToWrap = async (): Promise<void> => {
      const layout = await readHeaderLayout();
      expect(layout.observationDate.top).toBeGreaterThan(layout.timeTag.top);
      expect(layout.observationDate.bottom).toBeLessThanOrEqual(layout.headerBottom);
      expect(layout.toggle.left).toBeGreaterThan(layout.timeTag.left);
      expect(layout.headerScrollWidth).toBeLessThanOrEqual(layout.headerClientWidth);
    };

    await page.setViewportSize({ width: 375, height: 812 });
    await expectFieldsToWrap();
    await test.info().attach('observation-header-375px', {
      body: await page.screenshot(),
      contentType: 'image/png',
    });
    await page.setViewportSize({ width: 320, height: 740 });
    await expectFieldsToWrap();

    await expect(accordionToggle).toHaveAttribute('aria-expanded', 'false');
  });

  test('scenario 3.6: focusing either new-observation field opens the accordion item', async ({ database, page }) => {
    const seededData = await seedPortfolioHistoryData(database);
    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const item = page.locator('#portfolio-history-management-container-0');
    const timeTag = item.getByRole('textbox', { name: 'Time tag' });
    const observationDate = item.getByLabel('Observation date');
    const toggle = item.getByRole('button', { name: 'New observation details' });
    const panel = item.locator('#portfolio-history-management-0');

    await expect(toggle).toHaveAttribute('aria-expanded', 'false');
    await timeTag.focus();
    await expect(timeTag).toBeFocused();
    await expect(panel).toHaveClass(/\bshow\b/);
    await expect(toggle).toHaveAttribute('aria-expanded', 'true');

    await toggle.click();
    await expect(toggle).toHaveAttribute('aria-expanded', 'false');
    await expect(panel).not.toHaveClass(/\bcollapsing\b/);
    await expect(panel).not.toHaveClass(/\bshow\b/);

    await observationDate.focus();
    await expect(observationDate).toBeFocused();
    await expect(panel).toHaveClass(/\bshow\b/);
    await expect(toggle).toHaveAttribute('aria-expanded', 'true');
  });

  test('scenario 3.7: keeps the item open while keyboard focus moves between header fields', async ({
    database,
    page,
  }) => {
    const seededData = await seedPortfolioHistoryData(database);
    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const item = page.locator('#portfolio-history-management-container-0');
    const timeTag = item.getByRole('textbox', { name: 'Time tag' });
    const observationDate = item.getByLabel('Observation date');
    const toggle = item.getByRole('button', { name: 'New observation details' });
    const panel = item.locator('#portfolio-history-management-0');
    const selectedDate = '2026-04-20';

    await timeTag.focus();
    await timeTag.fill('E2E_FOCUS_NAVIGATION');
    await page.keyboard.press('Tab');
    await expect(observationDate).toBeFocused();
    await observationDate.fill(selectedDate);
    await page.keyboard.press('Shift+Tab');
    await expect(timeTag).toBeFocused();
    await expect(timeTag).toHaveValue('E2E_FOCUS_NAVIGATION');
    await expect(observationDate).toHaveValue(selectedDate);
    await expect(panel).toHaveClass(/\bshow\b/);
    await expect(toggle).toHaveAttribute('aria-expanded', 'true');
  });

  /** Rejects a browser-valid date outside the supported year format without submitting a stale timestamp.
   *
   * @author GPT-6 Sol
   */
  test('scenario 3.8: cancels an unsupported observation date and allows a corrected date', async ({
    database,
    page,
  }) => {
    const seededData = await seedPortfolioHistoryData(database);
    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const newObservationItem = page.locator('#portfolio-history-management-container-0');
    const form = page.locator('#portfolio-history-management-form-0');
    const dateInput = newObservationItem.getByLabel('Observation date');
    const timestampInput = form.locator('input[name="observationTimestamp.timestamp"]');
    const staleTimestamp = '2020-01-01T00:00:00.000Z';
    const savePath = `/api/portfolio/${seededData.portfolio.id}/history`;
    const saveRequests: Request[] = [];
    const trackSaveRequest = (request: Request): void => {
      if(request.method() === 'POST' && new URL(request.url()).pathname === savePath) {
        saveRequests.push(request);
      }
    };
    const failSave = async (route: Route): Promise<void> => {
      await route.fulfill({ status: 400, contentType: 'application/json', body: '{}' });
    };

    await newObservationItem.getByRole('textbox', { name: 'Time tag' }).fill('E2E_UNSUPPORTED_DATE');
    await timestampInput.evaluate((element, value) => {
      (element as HTMLInputElement).value = value;
    }, staleTimestamp);
    await dateInput.evaluate(element => {
      const input = element as HTMLInputElement;
      input.value = '10000-01-01';
      input.addEventListener('invalid', () => {
        input.dataset.reportedValidationMessage = input.validationMessage;
      });
    });
    await expect(dateInput).toHaveValue('10000-01-01');
    expect(await dateInput.evaluate(element => (element as HTMLInputElement).checkValidity())).toBe(true);

    page.on('request', trackSaveRequest);
    try {
      await page.route(`**${savePath}`, failSave);
      await form.locator('tfoot button.btn-primary[type="submit"]').click();
      await expect(dateInput).toHaveAttribute('data-reported-validation-message', 'Observation date is not valid.');
      expect(saveRequests).toHaveLength(0);
      await expect(timestampInput).toHaveValue(staleTimestamp);

      await dateInput.fill('2026-04-20');
      const failedSaveResponse = page.waitForResponse(response => {
        return response.request().method() === 'POST' && new URL(response.url()).pathname === savePath;
      });
      await form.locator('tfoot button.btn-primary[type="submit"]').click();
      expect((await failedSaveResponse).status()).toBe(400);
      expect(saveRequests).toHaveLength(1);
      const payload = JSON.parse(saveRequests[0].postData() ?? 'null') as {
        observationTimestamp?: { timestamp?: string };
      };
      const expectedTimestamp = await page.evaluate(() => new Date('2026-04-20T00:00:00').toISOString());
      expect(payload.observationTimestamp?.timestamp).toBe(expectedTimestamp);
      await expect(timestampInput).toHaveValue(expectedTimestamp);
      await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
    } finally {
      page.off('request', trackSaveRequest);
      await page.unroute(`**${savePath}`, failSave);
    }
  });

  /** Repeats pointer, text, date, and keyboard interaction without closing the new-observation panel.
   * @author GPT-6 Luna
   */
  test('scenario 3.9: keeps the new-observation panel open through repeated field interactions', async ({
    database,
    page,
  }) => {
    const seededData = await seedPortfolioHistoryData(database);
    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const item = page.locator('#portfolio-history-management-container-0');
    await exerciseAccordionFieldInteractions(page, {
      dateInput: item.getByLabel('Observation date'),
      expectedDate: '2026-04-20',
      expectedText: 'E2E_OBS_FIELD_39',
      panel: item.locator('#portfolio-history-management-0'),
      textInput: item.getByRole('textbox', { name: 'Time tag' }),
      toggle: item.getByRole('button', { name: 'New observation details' }),
    });

    await expect(database.query('SELECT id FROM public.portfolio_allocation_obs_time')).resolves.toEqual([]);
    await expect(database.query('SELECT observation_time_id FROM public.portfolio_allocation_fact')).resolves.toEqual([]);
  });

  /** Creates an observation, checks PostgreSQL persistence, then repeats header interactions after refresh.
   * @author GPT-6 Luna
   */
  test('scenario 3.10: preserves accordion interactions after successful observation refresh', async ({
    database,
    page,
  }) => {
    test.setTimeout(60_000);
    const seededData = await seedPortfolioHistoryData(database);
    await page.goto('/');
    await expectRootShell(page);
    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio);

    const initialNewItem = page.locator('#portfolio-history-management-container-0');
    const initialTimeTag = initialNewItem.getByRole('textbox', { name: 'Time tag' });
    await initialTimeTag.focus();
    await expect(initialNewItem.locator('#portfolio-history-management-0')).toHaveClass(/\bshow\b/);
    const initialDate = initialNewItem.getByLabel('Observation date');
    const savedDate = '2026-04-20';
    await initialDate.fill(savedDate);

    const initialForm = page.locator('#portfolio-history-management-form-0');
    const initialFormHandle = await initialForm.elementHandle();
    if(!initialFormHandle) {
      throw new Error('The new-observation form was not rendered.');
    }
    const row = await addAllocationRow(page, initialForm, 0);
    await fillExistingAssetAllocation(
      page,
      row,
      seededData.assets[EXISTING_BOND_TICKER],
      'BONDS',
      '2',
      '100',
      '200.00',
    );

    const timeTag = 'E2E_OBS_AFTER_REFRESH';
    await initialTimeTag.fill(timeTag);
    const expectedTimestamp = await localMidnightTimestamp(page, savedDate);
    const saveResponsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'POST'
        && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`;
    });
    await initialForm.locator('tfoot button.btn-primary[type="submit"]').click();
    expect((await saveResponsePromise).status()).toBe(204);

    const persistedObservation = await findObservationByTag(database, seededData.portfolio, timeTag);
    expect(persistedObservation.observation_timestamp).toEqual(new Date(expectedTimestamp));
    const persistedAllocations = await database.query<PersistedAllocation>(
      `SELECT paf.asset_id,
              COALESCE(paf.asset_market_price, 0)::numeric(18,8)::text AS asset_market_price,
              COALESCE(paf.asset_quantity, 0)::numeric(18,8)::text AS asset_quantity,
              paf.cash_reserve,
              paf.class,
              a.name,
              paf.observation_time_id,
              a.ticker,
              paf.total_market_value::text AS total_market_value
       FROM public.portfolio_allocation_fact paf
       JOIN public.asset a ON a.id = paf.asset_id
       WHERE paf.portfolio_id = $1
         AND paf.observation_time_id = $2`,
      [seededData.portfolio.id, persistedObservation.id],
    );
    expect(persistedAllocations).toEqual([{
      asset_id: seededData.assets[EXISTING_BOND_TICKER].id,
      asset_market_price: '100.00000000',
      asset_quantity: '2.00000000',
      cash_reserve: false,
      class: 'BONDS',
      name: EXISTING_BOND_NAME,
      observation_time_id: persistedObservation.id,
      ticker: EXISTING_BOND_TICKER,
      total_market_value: '200',
    }]);

    await expectSuccessNotification(page);
    await expectReloadedManagementItems(page, persistedObservation);
    expect(await initialFormHandle.evaluate((element) => element.isConnected)).toBe(false);
    await initialFormHandle.dispose();

    const refreshedNewItem = page.locator('#portfolio-history-management-container-0');
    const refreshedTimeTag = refreshedNewItem.getByRole('textbox', { name: 'Time tag' });
    const refreshedForm = page.locator('#portfolio-history-management-form-0');
    await expect(refreshedTimeTag).toHaveValue('');
    await expect(refreshedNewItem.getByLabel('Observation date')).toHaveValue(await browserLocalCalendarDate(page));

    await refreshedTimeTag.focus();
    await expect(refreshedNewItem.locator('#portfolio-history-management-0')).toHaveClass(/\bshow\b/);
    const draftRow = await addAllocationRow(page, refreshedForm, 0);
    await fillExistingAssetDirectValueAllocation(
      page,
      draftRow,
      seededData.assets[EXISTING_STOCK_TICKER],
      'STOCKS',
      '750.00',
    );
    await expect(draftRow.getByRole('textbox', { name: 'Total market value' })).toHaveValue('750.00');

    const newObservationToggle = refreshedNewItem.getByRole('button', { name: 'New observation details' });
    await newObservationToggle.click();
    await expect(newObservationToggle).toHaveAttribute('aria-expanded', 'false');
    await expect(refreshedNewItem.locator('#portfolio-history-management-0')).not.toHaveClass(/\bshow\b/);

    await exerciseAccordionFieldInteractions(page, {
      dateInput: refreshedNewItem.getByLabel('Observation date'),
      expectedDate: '2026-04-21',
      expectedText: 'E2E_OBS_FIELD_310',
      panel: refreshedNewItem.locator('#portfolio-history-management-0'),
      textInput: refreshedTimeTag,
      toggle: newObservationToggle,
    });
    await expect(refreshedForm.locator('[name="observationTimestamp.timeTag"]')).toHaveValue('E2E_OBS_FIELD_310');
    await expect(draftRow.getByRole('textbox', { name: 'Total market value' })).toHaveValue('750.00');

    const savedObservationItem = page.locator(`#portfolio-history-management-container-${persistedObservation.id}`);
    const savedObservationToggle = savedObservationItem.getByRole('button', {
      name: await expectedObservationHeaderName(page, persistedObservation),
      exact: true,
    });
    const savedObservationPanel = savedObservationItem.locator(`#portfolio-history-management-${persistedObservation.id}`);
    const detailRequests: string[] = [];
    const trackDetailRequest = (request: Request): void => {
      const requestUrl = new URL(request.url());
      if(request.method() === 'GET'
        && requestUrl.pathname === `/api/portfolio/${seededData.portfolio.id}/history`
        && requestUrl.searchParams.get('observationTimestampId') === String(persistedObservation.id)) {
        detailRequests.push(request.url());
      }
    };
    page.on('request', trackDetailRequest);
    try {
      const detailResponsePromise = page.waitForResponse((response) => {
        const requestUrl = new URL(response.url());
        return response.request().method() === 'GET'
          && requestUrl.pathname === `/api/portfolio/${seededData.portfolio.id}/history`
          && requestUrl.searchParams.get('observationTimestampId') === String(persistedObservation.id);
      });
      await savedObservationToggle.click();
      expect((await detailResponsePromise).status()).toBe(200);
      await expect(savedObservationPanel).toHaveClass(/\bshow\b/);
      await expect(refreshedNewItem.locator('#portfolio-history-management-0')).not.toHaveClass(/\bshow\b/);
      expect(detailRequests).toHaveLength(1);
    } finally {
      page.off('request', trackDetailRequest);
    }

    await exerciseAccordionFieldInteractions(page, {
      dateInput: refreshedNewItem.getByLabel('Observation date'),
      expectedDate: '2026-04-21',
      expectedText: 'E2E_OBS_FIELD_310',
      panel: refreshedNewItem.locator('#portfolio-history-management-0'),
      textInput: refreshedTimeTag,
      toggle: newObservationToggle,
    });
    await expect(savedObservationPanel).not.toHaveClass(/\bshow\b/);
    await expect(draftRow.getByRole('textbox', { name: 'Total market value' })).toHaveValue('750.00');
  });

  /** Verifies escaped observation tags, date labels, timestamp tie order, and reload stability.
   * @author GPT-6 Luna
   */
  test('scenario 17.1: renders escaped dated observation headers in stable chronological order', async ({
    database,
    page,
  }) => {
    const seededData = await seedPortfolioHistoryData(database);
    const unsafeTag = '<img src=x onerror=alert(1)>';
    const firstEqualTimestamp = await seedManagementObservation(
      database,
      seededData.portfolio,
      seededData.assets[EXISTING_BOND_TICKER],
      unsafeTag,
      '2026-10-13T12:00:00.000Z',
    );
    const olderObservation = await seedManagementObservation(
      database,
      seededData.portfolio,
      seededData.assets[EXISTING_BOND_TICKER],
      'E2E_OBS_OLDER_HEADER',
      '2024-02-29T12:00:00.000Z',
    );
    const secondEqualTimestamp = await seedManagementObservation(
      database,
      seededData.portfolio,
      seededData.assets[EXISTING_BOND_TICKER],
      'E2E_OBS_EQUAL_HEADER',
      '2026-10-13T12:00:00.000Z',
    );

    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    await expectPortfolioHistoryManagement(page, seededData.portfolio, 3);
    await expectManagementObservationOrder(page, [secondEqualTimestamp, firstEqualTimestamp, olderObservation]);

    const unsafeItem = page.locator(`#portfolio-history-management-container-${firstEqualTimestamp.id}`);
    const unsafeHeader = unsafeItem.getByRole('button', {
      name: await expectedObservationHeaderName(page, firstEqualTimestamp),
      exact: true,
    });
    await expect(unsafeHeader).toHaveText(await expectedObservationHeaderName(page, firstEqualTimestamp));
    await expect(unsafeItem.locator('img')).toHaveCount(0);

    await page.reload();
    await expectPortfolioHistoryManagement(page, seededData.portfolio, 3);
    await expectManagementObservationOrder(page, [secondEqualTimestamp, firstEqualTimestamp, olderObservation]);
  });

  test.describe('scenario 17.2: formats observation headers in a positive UTC offset', () => {
    test.use({ timezoneId: 'Asia/Tokyo' });

    /** Verifies persisted instants render as browser-local dates in Tokyo.
     * @author GPT-6 Luna
     * @author benizzio
     */
    test('shows the timestamp’s Tokyo calendar date instead of its UTC date', async ({ database, page }) => {
      const seededData = await seedPortfolioHistoryData(database);
      const observation = await seedManagementObservation(
        database,
        seededData.portfolio,
        seededData.assets[EXISTING_BOND_TICKER],
        'E2E_OBS_TOKYO_HEADER',
        '2026-10-12T15:30:00.000Z',
      );

      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      const item = page.locator(`#portfolio-history-management-container-${observation.id}`);
      await expect(item.getByRole('button', { name: 'E2E_OBS_TOKYO_HEADER - 2026-10-13', exact: true }))
        .toBeVisible();
    });
  });

  test.describe('scenario 17.3: formats observation headers across daylight-saving transitions', () => {
    test.use({ timezoneId: 'America/Los_Angeles' });

    /** Verifies browser-local dates on both sides of the 2026 daylight-saving transitions.
     * @author GPT-6 Luna
     * @author benizzio
     */
    test('uses the browser-local date before and after spring and fall clock changes', async ({ database, page }) => {
      const seededData = await seedPortfolioHistoryData(database);
      const observations: PersistedObservation[] = [];
      const fixtures = [
        { tag: 'E2E_OBS_DST_SPRING_BEFORE', timestamp: '2026-03-08T07:30:00.000Z', date: '2026-03-07' },
        { tag: 'E2E_OBS_DST_SPRING_AFTER', timestamp: '2026-03-09T07:30:00.000Z', date: '2026-03-09' },
        { tag: 'E2E_OBS_DST_FALL_BEFORE', timestamp: '2026-11-01T07:30:00.000Z', date: '2026-11-01' },
        { tag: 'E2E_OBS_DST_FALL_AFTER', timestamp: '2026-11-02T08:30:00.000Z', date: '2026-11-02' },
      ];

      for(const fixture of fixtures) {
        observations.push(await seedManagementObservation(
          database,
          seededData.portfolio,
          seededData.assets[EXISTING_BOND_TICKER],
          fixture.tag,
          fixture.timestamp,
        ));
      }

      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      await expectManagementObservationOrder(page, [...observations].reverse());
      for(const fixture of fixtures) {
        await expect(page.getByRole('button', {
          name: `${fixture.tag} - ${fixture.date}`,
          exact: true,
        })).toBeVisible();
      }
    });
  });

  test.describe('scenario 17.4: refreshes newly created observations in server order', () => {
    test.use({ timezoneId: 'UTC' });

    /** Verifies successful creation refreshes the list using server timestamp order.
     * @author GPT-6 Luna
     * @author benizzio
     */
    test('places backdated and future observations relative to existing records', async ({ database, page }) => {
      test.setTimeout(60_000);
      const seededData = await seedPortfolioHistoryModificationData(database);
      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      await expectPortfolioHistoryManagement(page, seededData.portfolio, 1);

      const backdatedDate = '2024-02-29';
      const backdated = await submitObservationForDate(
        page,
        database,
        seededData.portfolio,
        seededData.assets[EXISTING_BOND_TICKER],
        backdatedDate,
        'E2E_OBS_BACKDATED_HEADER',
        `${backdatedDate}T00:00:00.000Z`,
      );
      await expectManagementObservationOrder(page, [seededData.observation, backdated]);

      const futureDate = '2030-06-01';
      const future = await submitObservationForDate(
        page,
        database,
        seededData.portfolio,
        seededData.assets[EXISTING_BOND_TICKER],
        futureDate,
        'E2E_OBS_FUTURE_HEADER',
        `${futureDate}T00:00:00.000Z`,
      );
      await expectManagementObservationOrder(page, [future, seededData.observation, backdated]);
      await expect(database.query<{ id: number }>(
        `SELECT id FROM public.portfolio_allocation_obs_time WHERE id = ANY($1::bigint[]) ORDER BY id`,
        [[backdated.id, future.id]],
      )).resolves.toHaveLength(2);
    });
  });

  test.describe('scenario 17.5: keeps newly saved observations outside the record window persisted', () => {
    test.use({ timezoneId: 'UTC' });

    /** Verifies a saved observation outside the ten-record window remains persisted but not rendered.
     * @author GPT-6 Luna
     * @author benizzio
     */
    test('does not insert a backdated observation into the refreshed ten-record list', async ({ database, page }) => {
      test.setTimeout(60_000);
      const seededData = await seedPortfolioHistoryData(database);
      const observations: PersistedObservation[] = [];
      for(let index = 0; index < 10; index++) {
        const day = String(index + 1).padStart(2, '0');
        observations.push(await seedManagementObservation(
          database,
          seededData.portfolio,
          seededData.assets[EXISTING_BOND_TICKER],
          `E2E_OBS_WINDOW_${day}`,
          `2026-01-${day}T12:00:00.000Z`,
        ));
      }

      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      await expectPortfolioHistoryManagement(page, seededData.portfolio, 10);
      await expectManagementObservationOrder(page, [...observations].reverse());

      const outsideWindowDate = '2020-01-01';
      const outsideWindowObservation = await submitObservationForDate(
        page,
        database,
        seededData.portfolio,
        seededData.assets[EXISTING_BOND_TICKER],
        outsideWindowDate,
        'E2E_OBS_SAVED_OUTSIDE_WINDOW',
        `${outsideWindowDate}T00:00:00.000Z`,
      );

      await expectManagementObservationOrder(page, [...observations].reverse());
      await expect(page.locator(
        `#portfolio-history-management-container-${outsideWindowObservation.id}`,
      )).toHaveCount(0);
      await expect(database.query<{ count: string }>(
        `SELECT COUNT(*)::text AS count
         FROM public.portfolio_allocation_obs_time ot
         JOIN public.portfolio_allocation_fact paf
           ON paf.portfolio_id = ot.portfolio_id AND paf.observation_time_id = ot.id
         WHERE paf.portfolio_id = $1`,
        [seededData.portfolio.id],
      )).resolves.toEqual([{ count: '11' }]);
    });
  });

  test.describe('scenario 19.1: preserves exact timestamps during observation metadata edits', () => {
    test.use({ timezoneId: 'Asia/Tokyo' });

    /** Edits tag and allocation details while preserving an unchanged non-midnight timestamp exactly.
     * @author GPT-6 Luna
     */
    test('opens the editor accessibly and saves tag and allocation changes without reloading details repeatedly', async ({
      database,
      page,
    }) => {
      const seededData = await seedPortfolioHistoryData(database);
      const originalTimestamp = '2026-10-12T15:30:45.123456Z';
      const observation = await seedManagementObservation(
        database,
        seededData.portfolio,
        seededData.assets[EXISTING_BOND_TICKER],
        'E2E_OBS_PENCIL_ORIGINAL',
        originalTimestamp,
      );
      const detailsRequests: Request[] = [];
      const detailsPath = `/api/portfolio/${seededData.portfolio.id}/history`;
      const matchesDetailsRequest = (url: URL): boolean => url.pathname === detailsPath
        && url.searchParams.get('observationTimestampId') === String(observation.id);
      page.on('request', (request) => {
        const url = new URL(request.url());
        if(request.method() === 'GET' && matchesDetailsRequest(url)) {
          detailsRequests.push(request);
        }
      });

      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      let releaseDetailsRequest: () => void = () => {};
      const detailRequestGate = new Promise<void>(resolve => {
        releaseDetailsRequest = resolve;
      });
      const detailsRoute = (url: URL): boolean => matchesDetailsRequest(url);
      await page.route(detailsRoute, async route => {
        await detailRequestGate;
        await route.continue();
      });
      const item = page.locator(`#portfolio-history-management-container-${observation.id}`);
      const pencil = item.getByRole('button', { name: 'Edit observation E2E_OBS_PENCIL_ORIGINAL', exact: true });
      const detailsRequest = page.waitForRequest(request => request.method() === 'GET'
        && matchesDetailsRequest(new URL(request.url())));
      await pencil.focus();
      await page.keyboard.press('Enter');

      const tagInput = item.getByRole('textbox', { name: 'Observation tag', exact: true });
      const dateInput = item.getByLabel('Observation date', { exact: true });
      const form = item.locator(`#portfolio-history-management-form-${observation.id}`);
      const panel = item.locator(`#portfolio-history-management-${observation.id}`);
      const saveButton = form.locator('button[type="submit"]');
      await expect(tagInput).toBeVisible();
      await expect(tagInput).toBeFocused();
      await expect(tagInput).toHaveValue('E2E_OBS_PENCIL_ORIGINAL');
      await expect(dateInput).toHaveValue('2026-10-13');
      await expect(panel).toHaveClass(/show/);
      await detailsRequest;
      await expect(saveButton).toBeDisabled();
      releaseDetailsRequest();
      await expect(saveButton).toBeEnabled();
      await expect.poll(() => detailsRequests.length).toBe(1);
      await page.unroute(detailsRoute);

      const row = form.locator(`#portfolio-history-management-form-${observation.id}-row-0`);
      await row.getByRole('spinbutton', { name: 'Quantity' }).fill('2');
      await row.getByRole('textbox', { name: 'Market price' }).fill('150');
      await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue('300.00');
      await tagInput.fill('E2E_OBS_PENCIL_UPDATED');
      await dateInput.focus();
      await expect(panel).toHaveClass(/show/);
      await expect.poll(() => detailsRequests.length).toBe(1);

      const saveResponse = page.waitForResponse((response) => response.request().method() === 'POST'
        && new URL(response.url()).pathname === detailsPath);
      await saveButton.click();
      expect((await saveResponse).ok()).toBe(true);

      await expect(item.getByRole('button', {
        name: 'E2E_OBS_PENCIL_UPDATED - 2026-10-13',
        exact: true,
      })).toBeVisible();
      await expect(item.locator('[data-metadata-editor-active]')).toBeHidden();
      await expect.poll(() => detailsRequests.length).toBe(1);
      await expect(database.query<{ id: number; timestamp_preserved: boolean; time_tag: string; quantity_updated: boolean }>(
        `SELECT ot.id,
                ot.observation_timestamp = $2::timestamptz AS timestamp_preserved,
                ot.observation_time_tag AS time_tag,
                paf.asset_quantity = 2 AS quantity_updated
         FROM public.portfolio_allocation_obs_time ot
         JOIN public.portfolio_allocation_fact paf ON paf.observation_time_id = ot.id
         WHERE ot.id = $1 AND paf.portfolio_id = ot.portfolio_id`,
        [observation.id, originalTimestamp],
      )).resolves.toEqual([{
        id: observation.id,
        timestamp_preserved: true,
        time_tag: 'E2E_OBS_PENCIL_UPDATED',
        quantity_updated: true,
      }]);
    });
  });

  test.describe('scenario 19.2: converts edited observation dates across timezone transitions', () => {
    test.use({ timezoneId: 'America/Los_Angeles' });

    /** Changes a record date over the spring daylight-saving boundary and verifies stable identity and ordering.
     * @author GPT-6 Luna
     */
    test('uses the selected date offset and moves the same observation into server order', async ({ database, page }) => {
      const seededData = await seedPortfolioHistoryData(database);
      const editedObservation = await seedManagementObservation(
        database,
        seededData.portfolio,
        seededData.assets[EXISTING_BOND_TICKER],
        'E2E_OBS_PENCIL_DST',
        '2026-03-07T12:00:00.000Z',
      );
      const latestObservation = await seedManagementObservation(
        database,
        seededData.portfolio,
        seededData.assets[EXISTING_STOCK_TICKER],
        'E2E_OBS_PENCIL_LATEST',
        '2026-03-10T12:00:00.000Z',
      );

      await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
      const item = page.locator(`#portfolio-history-management-container-${editedObservation.id}`);
      await item.getByRole('button', { name: 'Edit observation E2E_OBS_PENCIL_DST', exact: true }).click();
      const dateInput = item.getByLabel('Observation date', { exact: true });
      await expect(dateInput).toHaveValue('2026-03-07');
      await expect(item.locator(`#portfolio-history-management-form-${editedObservation.id} button[type="submit"]`))
        .toBeEnabled();
      await dateInput.fill('2026-03-09');

      const saveResponse = page.waitForResponse((response) => response.request().method() === 'POST'
        && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`);
      await item.locator(`#portfolio-history-management-form-${editedObservation.id} button[type="submit"]`).click();
      expect((await saveResponse).ok()).toBe(true);

      const updatedObservation = await findObservationByTag(database, seededData.portfolio, 'E2E_OBS_PENCIL_DST');
      expect(updatedObservation.id).toBe(editedObservation.id);
      await expect(database.query<{ timestamp_matches: boolean }>(
        `SELECT observation_timestamp = $2::timestamptz AS timestamp_matches
         FROM public.portfolio_allocation_obs_time WHERE id = $1`,
        [editedObservation.id, '2026-03-09T07:00:00.000Z'],
      )).resolves.toEqual([{ timestamp_matches: true }]);
      await expectManagementObservationOrder(page, [latestObservation, updatedObservation]);
    });
  });

  /** Confirms a portfolio-scoped duplicate-tag rejection leaves the observation editor draft intact.
   * @author GPT-6 Luna
   */
  test('scenario 19.3: preserves an observation edit after a duplicate-tag rejection', async ({ database, page }) => {
    const seededData = await seedPortfolioHistoryData(database);
    const observation = await seedManagementObservation(
      database,
      seededData.portfolio,
      seededData.assets[EXISTING_BOND_TICKER],
      'E2E_OBS_PENCIL_DUPLICATE_SOURCE',
      '2026-05-01T12:00:00.000Z',
    );
    const duplicate = await seedManagementObservation(
      database,
      seededData.portfolio,
      seededData.assets[EXISTING_STOCK_TICKER],
      'E2E_OBS_PENCIL_DUPLICATE_TARGET',
      '2026-05-02T12:00:00.000Z',
    );

    await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
    const item = page.locator(`#portfolio-history-management-container-${observation.id}`);
    await item.getByRole('button', { name: 'Edit observation E2E_OBS_PENCIL_DUPLICATE_SOURCE', exact: true }).click();
    const form = item.locator(`#portfolio-history-management-form-${observation.id}`);
    const tagInput = item.getByRole('textbox', { name: 'Observation tag', exact: true });
    await expect(form.locator('button[type="submit"]')).toBeEnabled();
    await tagInput.fill(duplicate.observation_time_tag);

    const saveResponse = page.waitForResponse((response) => response.request().method() === 'POST'
      && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`);
    await form.locator('button[type="submit"]').click();
    expect((await saveResponse).ok()).toBe(false);

    await expect(tagInput).toBeVisible();
    await expect(tagInput).toHaveValue(duplicate.observation_time_tag);
    await expect(item.locator(`#portfolio-history-management-form-tbody-${observation.id} > tr`)).toHaveCount(1);
    await expect(database.query<{ id: number; observation_time_tag: string }>(
      `SELECT id, observation_time_tag FROM public.portfolio_allocation_obs_time
       WHERE id = ANY($1::bigint[]) ORDER BY id`,
      [[observation.id, duplicate.id]],
    )).resolves.toEqual([
      { id: observation.id, observation_time_tag: 'E2E_OBS_PENCIL_DUPLICATE_SOURCE' },
      { id: duplicate.id, observation_time_tag: 'E2E_OBS_PENCIL_DUPLICATE_TARGET' },
    ]);
  });
});

/** Creates a minimal observation for one calendar date and verifies the request and PostgreSQL timestamp.
 *
 * @author GPT-6 Luna
 */
async function submitObservationForDate(
  page: Page,
  database: E2eDatabase,
  portfolio: SeededPortfolio,
  asset: SeededAsset,
  calendarDate: string,
  timeTag: string,
  expectedTimestamp: string,
): Promise<PersistedObservation> {
  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  await newObservationItem.getByRole('textbox', { name: 'Time tag' }).fill(timeTag);
  await newObservationItem.getByLabel('Observation date').fill(calendarDate);
  await expect(newObservationItem.locator('#portfolio-history-management-0')).toHaveClass(/\bshow\b/);

  const form = page.locator('#portfolio-history-management-form-0');
  const row = await addAllocationRow(page, form, 0);
  await fillExistingAssetDirectValueAllocation(page, row, asset, 'BONDS', '100.00');

  const savePath = `/api/portfolio/${portfolio.id}/history`;
  const saveRequestPromise = page.waitForRequest(request => {
    return request.method() === 'POST' && new URL(request.url()).pathname === savePath;
  });
  const saveResponsePromise = page.waitForResponse(response => {
    return response.request().method() === 'POST' && new URL(response.url()).pathname === savePath;
  });
  await form.locator('tfoot button.btn-primary[type="submit"]').click();

  const saveRequest = await saveRequestPromise;
  const saveResponse = await saveResponsePromise;
  expect(saveResponse.status()).toBe(204);
  const payload = JSON.parse(saveRequest.postData() ?? 'null') as {
    observationTimestamp?: { timestamp?: string };
  };
  expect(payload.observationTimestamp?.timestamp).toBe(expectedTimestamp);

  const persistedObservation = await findObservationByTag(database, portfolio, timeTag);
  expect(persistedObservation.observation_timestamp).toEqual(new Date(expectedTimestamp));
  await expect(newObservationItem.getByLabel('Observation date'))
    .toHaveValue(await browserLocalCalendarDate(page));

  return persistedObservation;
}

/** Seeds the portfolio and asset catalog needed to exercise asset search filtering and scrolling.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
async function seedAssetSearchAutocompleteData(database: E2eDatabase): Promise<{
  assets: Record<string, SeededAsset>;
  portfolio: SeededPortfolio;
}> {
  const portfolioRows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    ['E2E Ticker Autocomplete Portfolio', JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );
  const autoAssets = Array.from({ length: 12 }, (_, index) => ({
    ticker: `E2E:AUTO-${index.toString().padStart(2, '0')}`,
    name: `Autocomplete Asset ${index}`,
  }));
  const assetsToSeed = [
    ...autoAssets,
    { ticker: 'SEARCH:NAME-ONE', name: 'North America: Equity & <Growth>' },
    { ticker: 'SEARCH:NAME-TWO', name: 'North America: Equity & <Growth>' },
    { ticker: 'SEARCH:LONG', name: 'An Extremely Long Name for a Global Balanced Portfolio' },
  ];
  const values = assetsToSeed.map((_, index) => `($${index * 2 + 1}, $${index * 2 + 2})`);
  const assetParameters = assetsToSeed.flatMap(({ ticker, name }) => [ticker, name]);
  const assetRows = await database.query<SeededAsset>(
    `INSERT INTO public.asset (ticker, name)
     VALUES ${values.join(', ')}
     RETURNING id, ticker, name`,
    assetParameters,
  );

  expect(portfolioRows).toHaveLength(1);
  expect(assetRows).toHaveLength(assetsToSeed.length);

  return {
    assets: Object.fromEntries(assetRows.map(asset => [asset.ticker, asset])) as Record<string, SeededAsset>,
    portfolio: portfolioRows[0],
  };
}

/** Seeds the catalog and opens one editable asset row for a focused asset search scenario.
 *
 * @param database - Isolated test database used to seed the portfolio and assets.
 * @param page - Browser page used to open the portfolio history form.
 * @returns The seeded assets and the rendered row controls used by the scenario.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
async function prepareAssetSearchAutocompleteRow(database: E2eDatabase, page: Page): Promise<{
  assets: Record<string, SeededAsset>;
  row: Locator;
  assetSearchInput: Locator;
  suggestions: Locator;
}> {
  const seededData = await seedAssetSearchAutocompleteData(database);

  await page.goto('/');
  await expectRootShell(page);
  await page.goto(`/portfolio/${seededData.portfolio.id}/history/manage`);
  await expectPortfolioHistoryManagement(page, seededData.portfolio);

  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  await newObservationItem.getByRole('textbox', { name: 'Time tag' }).fill('AUTOCOMPLETE');
  const accordionBody = newObservationItem.locator('#portfolio-history-management-0');
  await expect(accordionBody).toHaveClass(/\bshow\b/);
  await expect(accordionBody).not.toHaveClass(/\bcollapsing\b/);

  const form = page.locator('#portfolio-history-management-form-0');
  const row = await addAllocationRow(page, form, 0);

  return {
    assets: seededData.assets,
    row,
    assetSearchInput: row.getByRole('combobox', { name: 'Asset', exact: true }),
    suggestions: row.getByRole('listbox', { name: 'Available assets' }),
  };
}

/** Creates a touch-enabled Chromium page and prepares one editable asset search row.
 *
 * @param database - Isolated test database used to seed the portfolio and assets.
 * @param browser - Chromium browser used to create a touch-capable context.
 * @returns The isolated context, page, and controls for one autocomplete row.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
async function prepareTouchAutocompleteRow(database: E2eDatabase, browser: Browser): Promise<{
  assets: Record<string, SeededAsset>;
  context: Awaited<ReturnType<Browser['newContext']>>;
  page: Page;
  row: Locator;
  assetSearchInput: Locator;
  suggestions: Locator;
}> {
  const context = await browser.newContext({
    deviceScaleFactor: 1,
    hasTouch: true,
    viewport: { width: 1280, height: 900 },
  });

  try {
    const page = await context.newPage();
    const preparedRow = await prepareAssetSearchAutocompleteRow(database, page);

    return { ...preparedRow, context, page };
  } catch(error) {
    await context.close();
    throw error;
  }
}

/** Asserts that an option has synchronized accessibility and theme-aware visible highlight state.
 *
 * @param option - Rendered suggestion expected to be the active option.
 *
 * @author GPT-6 Luna
 */
async function expectOptionHighlighted(option: Locator): Promise<void> {
  await expect(option).toHaveAttribute('aria-selected', 'true');

  const colors = await option.evaluate(element => {
    const expectedColors = document.createElement('span');
    expectedColors.style.backgroundColor = 'var(--bs-primary-bg-subtle)';
    expectedColors.style.color = 'var(--bs-primary-text-emphasis)';
    element.append(expectedColors);
    const expectedBackgroundColor = getComputedStyle(expectedColors).backgroundColor;
    const expectedColor = getComputedStyle(expectedColors).color;
    expectedColors.remove();

    const actualStyle = getComputedStyle(element);
    return {
      actualBackgroundColor: actualStyle.backgroundColor,
      actualColor: actualStyle.color,
      boxShadow: actualStyle.boxShadow,
      expectedBackgroundColor,
      expectedColor,
    };
  });

  expect(colors.actualBackgroundColor).toBe(colors.expectedBackgroundColor);
  expect(colors.actualColor).toBe(colors.expectedColor);
  expect(colors.boxShadow).toContain('inset');
}

/** Asserts the unmatched ticker has entered editable, focused new-asset mode.
 *
 * @param row - Allocation row containing the ticker and asset-name inputs.
 * @param ticker - Unmatched ticker that must remain in the row.
 *
 * @author GPT-6 Luna
 */
async function expectNewAssetMode(row: Locator, ticker: string): Promise<void> {
  const tickerInput = row.getByRole('textbox', { name: 'Asset ticker', exact: true });
  const assetNameInput = row.getByRole('textbox', { name: 'Asset name' });

  await expect(tickerInput).toHaveAttribute('readonly', '');
  await expect(tickerInput).toHaveValue(ticker);
  await expect(row.getByText('* Creating new asset', { exact: true })).toBeVisible();
  await expect(assetNameInput).toBeVisible();
  await expect(assetNameInput).not.toHaveAttribute('readonly', '');
  await expect(assetNameInput).toHaveAttribute('required', '');
  await expect(assetNameInput).toBeFocused();
  await expect(row.locator('[data-asset-search-input]')).toBeHidden();
  await expect(row.locator('[data-asset-search-input]')).toHaveAttribute('aria-expanded', 'false');
  await expect(tickerInput).toBeVisible();
}

/** Records asset endpoint requests for one ticker and returns a cleanup callback.
 *
 * @param page - Page emitting the network requests.
 * @param ticker - Exact ticker path to record.
 * @returns Captured request URLs and a method that unregisters the listener.
 *
 * @author GPT-6 Luna
 */
function trackAssetLookupRequests(page: Page, ticker: string): { requests: string[]; stop: () => void } {
  const requests: string[] = [];
  const listener = (request: Request): void => {
    if(request.method() === 'GET' && new URL(request.url()).pathname === `/api/asset/${ticker}`) {
      requests.push(request.url());
    }
  };
  page.on('request', listener);

  return {
    requests,
    stop: () => page.off('request', listener),
  };
}

/** Seeds one portfolio and the two assets that scenario 3 must resolve as existing. */
async function seedPortfolioHistoryData(database: E2eDatabase): Promise<{
  assets: Record<string, SeededAsset>;
  portfolio: SeededPortfolio;
}> {
  const portfolioRows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [PORTFOLIO_NAME, JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );
  const assetRows = await database.query<SeededAsset>(
    `INSERT INTO public.asset (ticker, name)
     VALUES ($1, $2), ($3, $4)
     RETURNING id, ticker, name`,
    [EXISTING_BOND_TICKER, EXISTING_BOND_NAME, EXISTING_STOCK_TICKER, EXISTING_STOCK_NAME],
  );

  expect(portfolioRows).toHaveLength(1);
  expect(assetRows).toHaveLength(2);
  expect(portfolioRows[0]).toEqual(expect.objectContaining({ id: expect.any(Number), name: PORTFOLIO_NAME }));
  expect(assetRows.map(({ ticker }) => ticker).sort()).toEqual([EXISTING_BOND_TICKER, EXISTING_STOCK_TICKER]);

  const assets = Object.fromEntries(assetRows.map((asset) => [asset.ticker, asset])) as Record<string, SeededAsset>;
  return { assets, portfolio: portfolioRows[0] };
}

/** Seeds one persisted observation and a fact so it is eligible for management-list reads.
 *
 * @param database - Isolated E2E database receiving the fixture rows.
 * @param portfolio - Portfolio that owns the observation allocation fact.
 * @param asset - Existing asset referenced by the seeded fact.
 * @param timeTag - Unique management-header tag for this observation.
 * @param timestamp - Explicit observation instant stored in PostgreSQL.
 * @returns The observation identity and values read back from PostgreSQL.
 * @author GPT-6 Luna
 * @author benizzio
 * @author OpenCode
  */
async function seedManagementObservation(
  database: E2eDatabase,
  portfolio: SeededPortfolio,
  asset: SeededAsset,
  timeTag: string,
  timestamp: string,
): Promise<PersistedObservation> {
  const observations = await database.query<PersistedObservation>(
    `INSERT INTO public.portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
     VALUES ($1, $2, $3::timestamptz)
     RETURNING id, observation_timestamp, observation_time_tag`,
    [portfolio.id, timeTag, timestamp],
  );
  expect(observations).toHaveLength(1);

  const observation = observations[0];
  await database.query(
    `INSERT INTO public.portfolio_allocation_fact
       (portfolio_id, asset_id, class, cash_reserve, asset_quantity, asset_market_price,
        total_market_value, observation_time_id)
     VALUES ($1, $2, 'BONDS', false, 1, 100, 100, $3)`,
    [portfolio.id, asset.id, observation.id],
  );

  return observation;
}

/** Seeds the fixed four-asset observation required by scenario 5.
 *
 * @author OpenCode
 */
async function seedPortfolioHistoryModificationData(
  database: E2eDatabase,
): Promise<SeededHistoryModificationData> {
  const portfolioRows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [MODIFICATION_PORTFOLIO_NAME, JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );
  const assetRows = await database.query<SeededAsset>(
    `INSERT INTO public.asset (ticker, name)
     VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8)
     RETURNING id, ticker, name`,
    [
      NEW_BOND_TICKER,
      NEW_BOND_NAME,
      EXISTING_BOND_TICKER,
      EXISTING_BOND_NAME,
      EXISTING_STOCK_TICKER,
      EXISTING_STOCK_NAME,
      NEW_STOCK_TICKER,
      NEW_STOCK_NAME,
    ],
  );

  expect(portfolioRows).toHaveLength(1);
  expect(assetRows).toHaveLength(4);
  expect(portfolioRows[0]).toEqual(expect.objectContaining({
    id: expect.any(Number),
    name: MODIFICATION_PORTFOLIO_NAME,
  }));

  const assets = Object.fromEntries(assetRows.map((asset) => [asset.ticker, asset])) as Record<string, SeededAsset>;
  const observationRows = await database.query<PersistedObservation>(
    `INSERT INTO public.portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
     VALUES ($1, $2, $3::timestamptz)
     RETURNING id, observation_timestamp, observation_time_tag`,
    [portfolioRows[0].id, OBSERVATION_TIME_TAG, FIXED_OBSERVATION_TIMESTAMP],
  );
  expect(observationRows).toHaveLength(1);

  const observation = observationRows[0];
  await database.query(
    `INSERT INTO public.portfolio_allocation_fact
       (portfolio_id, asset_id, class, cash_reserve, asset_quantity, asset_market_price,
        total_market_value, observation_time_id)
     VALUES
       ($1, $2, 'BONDS', false, 0, 0, 1000, $3),
       ($1, $4, 'BONDS', false, 30, 100, 3000, $3),
       ($1, $5, 'STOCKS', false, 0, 0, 2000, $3),
       ($1, $6, 'STOCKS', false, 40, 100, 4000, $3)`,
    [
      portfolioRows[0].id,
      assets[NEW_BOND_TICKER].id,
      observation.id,
      assets[EXISTING_BOND_TICKER].id,
      assets[EXISTING_STOCK_TICKER].id,
      assets[NEW_STOCK_TICKER].id,
    ],
  );

  expect(observation.observation_time_tag).toBe(OBSERVATION_TIME_TAG);
  expect(observation.observation_timestamp).toEqual(new Date(FIXED_OBSERVATION_TIMESTAMP));
  return { assets, observation, portfolio: portfolioRows[0] };
}

/** Asserts the global shell and its portfolios navigation link. */
async function expectRootShell(page: Page): Promise<void> {
  await expectRoute(page, '/');
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
}

/** Asserts the portfolio list contains the seeded portfolio and New portfolio card. */
async function expectPortfolioList(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expectRoute(page, '/portfolios');

  const portfolioCards = page.locator('#portfolios .portfolio-card');
  await expect(portfolioCards).toHaveCount(2);

  const existingPortfolioCard = portfolioCardFor(page, portfolio);
  await expect(existingPortfolioCard).toBeVisible();
  await expect(existingPortfolioCard).toHaveAttribute('data-navigate-to', `/portfolio/${portfolio.id}`);
  await expect(existingPortfolioCard).toHaveAttribute('navigate-to-bound', 'true');
  await expect(
    existingPortfolioCard.getByRole('heading', { level: 5, name: portfolio.name, exact: true }),
  ).toBeVisible();

  const newPortfolioCard = portfolioCards.last();
  await expect(newPortfolioCard).toHaveAttribute('data-navigate-to', '/portfolios/new');
  await expect(newPortfolioCard).toHaveAttribute('navigate-to-bound', 'true');
  await expect(
    newPortfolioCard.getByRole('heading', { level: 5, name: 'New portfolio', exact: true }),
  ).toBeVisible();
}

/** Returns the seeded portfolio card used for navigation. */
function portfolioCardFor(page: Page, portfolio: SeededPortfolio): Locator {
  return page.locator('#portfolios .portfolio-card').filter({ hasText: portfolio.name });
}

/** Asserts the selected portfolio context and edit control. */
async function expectPortfolioContext(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expect(page.locator('#portfolio-context .badge.text-bg-secondary')).toHaveText(portfolio.name);

  const editButton = page.locator('#portfolio-context button[data-navigate-to="/portfolio/:portfolioId/edit"]');
  await expect(editButton).toBeVisible();
  await expect(editButton.locator('.bi-pen')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
}

/** Asserts submenu labels, selected state, and outlined or filled button styling. */
async function expectPortfolioNavigation(
  page: Page,
  selectedNavigation: NavigationOptionName | undefined,
  assertSelection = true,
): Promise<void> {
  const navigation = page.getByRole('group', { name: 'Navigation' });
  await expect(navigation).toBeVisible();

  for (const option of PORTFOLIO_NAVIGATION_OPTIONS) {
    const radio = navigation.getByRole('radio', { name: option.name, exact: true });
    const label = navigation.locator(`label[for="${option.id}"]`);

    await expect(radio).toBeAttached();
    await expect(label).toHaveText(option.name);
    await expect(label).toHaveClass(/btn-outline-primary/);

    if (!assertSelection) {
      continue;
    }

    if (option.name === selectedNavigation) {
      await expect(radio).toBeChecked();
      const borderColor = await label.evaluate((element) => getComputedStyle(element).borderTopColor);
      await expect(label).toHaveCSS('background-color', borderColor);
    } else {
      await expect(radio).not.toBeChecked();
      await expect(label).toHaveCSS('background-color', 'rgba(0, 0, 0, 0)');
    }
  }
}

/** Asserts the empty history page and its history-management control. */
async function expectPortfolioHistory(
  page: Page,
  portfolio: SeededPortfolio,
  expectEmpty: boolean,
): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}/history`);
  await expectPortfolioContext(page, portfolio);
  await expectPortfolioNavigation(page, 'Portfolio');

  const history = page.locator('#accordion-portfolio-history');
  await expect(history).toBeVisible();
  if (expectEmpty) {
    await expect(history.locator(':scope > *')).toHaveCount(1);
  }

  const managementButton = history.locator('> button[onclick="navigateToPortfolioAllocationManagement()"]');
  await expect(managementButton).toBeVisible();
  await expect(managementButton.locator('.bi-database-gear')).toBeVisible();
}

/** Asserts the history-management shell and its collapsed new-observation item. */
async function expectPortfolioHistoryManagement(
  page: Page,
  portfolio: SeededPortfolio,
  observationCount = 0,
): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}/history/manage`);
  await expectPortfolioContext(page, portfolio);
  await expectPortfolioNavigation(page, undefined, false);

  const returnButton = page.locator(
    'button[onclick="portfolioHistoryManagement.navigateToPortfolioAllocationViewing()"]',
  );
  await expect(returnButton).toBeVisible();
  await expect(returnButton.locator('.bi-pie-chart-fill')).toBeVisible();

  const managementCard = page.locator('.card').filter({ hasText: 'Manage portfolio allocation data' });
  await expect(managementCard).toBeVisible();
  await expect(managementCard.getByText('Manage portfolio allocation data', { exact: true })).toBeVisible();

  const managementAccordion = page.locator('#accordion-portfolio-history-management');
  await expect(managementAccordion).toBeVisible();
  await expect(managementAccordion.locator(':scope > .accordion-item')).toHaveCount(1 + observationCount);

  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  const newObservationButton = newObservationItem.getByRole('button');
  await expect(newObservationButton).toHaveClass(/collapsed/);
  await expect(newObservationItem.getByRole('textbox', { name: 'Time tag' })).toBeVisible();
  const observationDateInput = newObservationItem.getByLabel('Observation date');
  await expect(observationDateInput).toHaveAttribute('required', '');
  await expect(observationDateInput).toHaveAttribute('form', 'portfolio-history-management-form-0');
  await expect(observationDateInput).toHaveValue(await browserLocalCalendarDate(page));
}

/** Asserts the history management table headers and icon controls. */
async function expectHistoryManagementForm(form: Locator, observationId: number): Promise<void> {
  const table = form.getByRole('table');
  await expect(table).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Asset', exact: true })).toHaveAttribute('colspan', '2');
  await expect(table.getByRole('columnheader', { name: 'Class', exact: true })).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Cash reserve?', exact: true })).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Quantity', exact: true })).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Market price', exact: true })).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Total market value', exact: true })).toBeVisible();

  const addButton = form.locator('tfoot button.btn-secondary[type="button"]');
  await expect(addButton).toBeVisible();
  await expect(addButton.locator('.bi-plus-circle')).toBeVisible();

  const saveButton = form.locator('tfoot button.btn-primary[type="submit"]');
  await expect(saveButton).toBeVisible();
  await expect(saveButton.locator('.bi-save-fill')).toBeVisible();
  await expect(form.locator(`tbody#portfolio-history-management-form-tbody-${observationId} tr`)).toHaveCount(
    observationId === 0 ? 0 : 4,
  );
}

/** Adds one allocation row and waits for its financial input bindings. */
async function addAllocationRow(page: Page, form: Locator, index: number): Promise<Locator> {
  await form.locator('tfoot button.btn-secondary[type="button"]').click();

  const row = page.locator(`#portfolio-history-management-form-0-row-${index}`);
  await expect(row).toBeVisible();
  await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveAttribute(
    'data-financial-input-bound',
    'true',
  );
  await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveAttribute(
    'data-financial-input-bound',
    'true',
  );

  return row;
}

/** Fills a new asset row with a direct total market value. */
async function fillNewAssetAllocation(
  page: Page,
  row: Locator,
  ticker: string,
  name: string,
  assetClass: string,
  totalMarketValue: string,
): Promise<void> {
  await searchForAsset(page, row, ticker, 404);
  await expect(row.getByText('* Creating new asset', { exact: true })).toBeVisible();

  await row.getByRole('textbox', { name: 'Asset name' }).fill(name);
  await row.getByRole('combobox', { name: 'Class' }).fill(assetClass);
  await fillDirectTotalMarketValue(row, totalMarketValue);
  await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue('');
  await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue('');
}

/** Fills an existing asset row with quantity and market price. */
async function fillExistingAssetAllocation(
  page: Page,
  row: Locator,
  asset: SeededAsset,
  assetClass: string,
  quantity: string,
  marketPrice: string,
  totalMarketValue: string,
  cashReserve = false,
): Promise<void> {
  await selectAssetFromAutocomplete(page, row, asset.ticker);
  await expectExistingAsset(row, asset);

  await row.getByRole('combobox', { name: 'Class' }).fill(assetClass);
  await fillCalculatedValues(row, quantity, marketPrice, totalMarketValue);
  await setCashReserve(row, cashReserve);
}

/** Fills an existing asset row with a direct total market value. */
async function fillExistingAssetDirectValueAllocation(
  page: Page,
  row: Locator,
  asset: SeededAsset,
  assetClass: string,
  totalMarketValue: string,
): Promise<void> {
  await selectAssetFromAutocomplete(page, row, asset.ticker);
  await expectExistingAsset(row, asset);

  await row.getByRole('combobox', { name: 'Class' }).fill(assetClass);
  await fillDirectTotalMarketValue(row, totalMarketValue);
  await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue('');
  await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue('');
}

/** Fills a new asset row with quantity and market price. */
async function fillNewAssetCalculatedAllocation(
  page: Page,
  row: Locator,
  ticker: string,
  name: string,
  assetClass: string,
  quantity: string,
  marketPrice: string,
): Promise<void> {
  await searchForAsset(page, row, ticker, 404);
  await expect(row.getByText('* Creating new asset', { exact: true })).toBeVisible();

  await row.getByRole('textbox', { name: 'Asset name' }).fill(name);
  await row.getByRole('combobox', { name: 'Class' }).fill(assetClass);
  await fillCalculatedValues(row, quantity, marketPrice, '4,000.00');
}

/**
 * Selects a known autocomplete suggestion and waits for its immediate asset lookup.
 *
 * @param page - Browser page that receives the lookup request.
 * @param row - Portfolio-history allocation row containing the asset search input.
 * @param ticker - Seeded ticker expected in the autocomplete suggestions.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
async function selectAssetFromAutocomplete(page: Page, row: Locator, ticker: string): Promise<void> {
  await expect(page.locator(`#datalist-assets option[value="${ticker}"]`)).toBeAttached();
  const assetSearchInput = row.getByRole('combobox', { name: 'Asset', exact: true });
  await assetSearchInput.click();
  await assetSearchInput.fill(ticker);

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

/** Selects one known class through the populated browser datalist input. */
async function selectClassFromDatalist(page: Page, row: Locator, assetClass: string): Promise<void> {
  await expect(page.locator(`#datalist-classes option[value="${assetClass}"]`)).toBeAttached();
  const classInput = row.getByRole('combobox', { name: 'Class' });
  await classInput.click();
  await classInput.fill(assetClass);
  await expect(classInput).toHaveValue(assetClass);
}

/** Waits for the asset lookup and verifies its HTTP result.
 *
 * @author GPT-6 Sol
 */
async function searchForAsset(page: Page, row: Locator, ticker: string, status: number): Promise<void> {
  const assetSearchInput = row.getByRole('combobox', { name: 'Asset', exact: true });
  await assetSearchInput.fill(ticker);

  const responsePromise = page.waitForResponse((response) => {
    return response.request().method() === 'GET'
      && new URL(response.url()).pathname === `/api/asset/${ticker}`;
  });
  await row.locator('[data-asset-action-button]').click();
  const response = await responsePromise;
  expect(response.status()).toBe(status);
}

/** Verifies that an asset lookup selected the expected existing asset. */
async function expectExistingAsset(row: Locator, asset: SeededAsset): Promise<void> {
  await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveAttribute('readonly', '');
  await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveValue(asset.ticker);
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toBeVisible();
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveAttribute('readonly', '');
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveValue(asset.name);
  await expect(row.locator('input[type="hidden"]').first()).toHaveValue(asset.id.toString());
  await expect(row.getByText('* Creating new asset', { exact: true })).toBeHidden();
  await expect(row.locator('[data-asset-action-button] .bi-x-circle')).toBeVisible();
}

/** Fills a direct total value and verifies that quantity and market price remain empty. */
async function fillDirectTotalMarketValue(row: Locator, totalMarketValue: string): Promise<void> {
  const totalInput = row.getByRole('textbox', { name: 'Total market value' });
  await totalInput.fill(totalMarketValue);
  await totalInput.blur();
  await expect(totalInput).toHaveValue(totalMarketValue.includes(',') ? `${totalMarketValue}.00` : totalMarketValue);
}

/** Fills quantity and price and verifies the calculated formatted total. */
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

/** Sets and verifies one observation row's cash-reserve state. */
async function setCashReserve(row: Locator, cashReserve: boolean): Promise<void> {
  const checkbox = row.getByRole('checkbox');
  if (cashReserve) {
    await checkbox.check();
  } else {
    await checkbox.uncheck();
  }
  await expect(checkbox).toBeChecked({ checked: cashReserve });
}

/** Verifies all four rows before saving the new observation. */
async function expectNewObservationRows(form: Locator): Promise<void> {
  await expectEditableObservationRows(form, [
    {
      assetClass: 'BONDS',
      cashReserve: false,
      marketPrice: '',
      name: NEW_BOND_NAME,
      quantity: '',
      ticker: NEW_BOND_TICKER,
      total: '1,000.00',
    },
    {
      assetClass: 'BONDS',
      cashReserve: false,
      marketPrice: '100.00000000',
      name: EXISTING_BOND_NAME,
      quantity: '30',
      ticker: EXISTING_BOND_TICKER,
      total: '3,000.00',
    },
    {
      assetClass: 'STOCKS',
      cashReserve: false,
      marketPrice: '',
      name: EXISTING_STOCK_NAME,
      quantity: '',
      ticker: EXISTING_STOCK_TICKER,
      total: '2,000.00',
    },
    {
      assetClass: 'STOCKS',
      cashReserve: false,
      marketPrice: '100.00000000',
      name: NEW_STOCK_NAME,
      quantity: '40',
      ticker: NEW_STOCK_TICKER,
      total: '4,000.00',
    },
  ]);
}

/** Verifies editable asset, class, quantity, price, total, and cash-reserve fields. */
async function expectEditableObservationRows(
  form: Locator,
  expectedRows: readonly ExpectedObservationRow[],
): Promise<void> {
  const rows = form.locator('tbody tr');
  await expect(rows).toHaveCount(expectedRows.length);

  for (const [index, expected] of expectedRows.entries()) {
    const row = rows.nth(index);
    await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveValue(expected.ticker);
    await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveValue(expected.name);
    await expect(row.getByRole('combobox', { name: 'Class' })).toHaveValue(expected.assetClass);
    await expect(row.getByRole('checkbox')).toBeChecked({ checked: expected.cashReserve });
    await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue(expected.quantity);
    await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue(expected.marketPrice);
    await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue(expected.total);
  }
}

/** Verifies persisted rows whose asset columns are rendered as read-only text. */
async function expectReadOnlyObservationRows(
  form: Locator,
  observationId: number,
  expectedRows: readonly ExpectedObservationRow[],
): Promise<void> {
  const rows = form.locator('tbody tr');
  await expect(rows).toHaveCount(expectedRows.length);

  for (const [index, expected] of expectedRows.entries()) {
    const row = rows.nth(index);
    await expect(row).toHaveAttribute(
      'id',
      `portfolio-history-management-form-${observationId}-row-${index}`,
    );
    await expect(row.locator('td').nth(0)).toContainText(expected.ticker);
    await expect(row.locator('td').nth(1)).toContainText(expected.name);
    await expect(row.getByRole('combobox', { name: 'Asset', exact: true })).toHaveCount(0);
    await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveCount(0);
    await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveCount(0);
    await expect(row.getByRole('combobox', { name: 'Class' })).toHaveValue(expected.assetClass);
    await expect(row.getByRole('checkbox')).toBeChecked({ checked: expected.cashReserve });
    await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue(expected.quantity);
    await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue(expected.marketPrice);
    await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue(expected.total);
  }
}

/** Verifies the management list refresh and collapsed new and persisted observation items.
 * @author GPT-6 Luna
 * @author benizzio
 */
async function expectReloadedManagementItems(page: Page, observation: PersistedObservation): Promise<void> {
  const managementAccordion = page.locator('#accordion-portfolio-history-management');
  await expect(managementAccordion.locator(':scope > .accordion-item')).toHaveCount(2);

  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  await expect(newObservationItem.locator('#portfolio-history-management-trigger-0 > button')).toHaveClass(/collapsed/);
  await expect(newObservationItem.getByRole('textbox', { name: 'Time tag' })).toBeVisible();
  await expect(newObservationItem.getByLabel('Observation date')).toHaveValue(await browserLocalCalendarDate(page));

  const observationItem = page.locator(`#portfolio-history-management-container-${observation.id}`);
  const observationButton = observationItem.getByRole('button', {
    name: await expectedObservationHeaderName(page, observation),
    exact: true,
  });
  await expect(observationButton).toBeVisible();
  await expect(observationButton).toHaveClass(/collapsed/);
  await expect(observationItem.locator(`#portfolio-history-management-${observation.id}`)).not.toHaveClass(/\bshow\b/);
}

/** Verifies the synthetic new item followed by persisted observations in timestamp order.
 * @author GPT-6 Luna
 * @author benizzio
 */
async function expectManagementObservationOrder(
  page: Page,
  observations: readonly PersistedObservation[],
  expectNewItemCollapsed = true,
): Promise<void> {
  const managementAccordion = page.locator('#accordion-portfolio-history-management');
  await expect(managementAccordion.locator(':scope > .accordion-item')).toHaveCount(1 + observations.length);

  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  const newItemToggle = newObservationItem.locator('#portfolio-history-management-trigger-0 > button');
  if(expectNewItemCollapsed) {
    await expect(newItemToggle).toHaveClass(/collapsed/);
  }

  for (const [index, observation] of observations.entries()) {
    const item = managementAccordion.locator(':scope > .accordion-item').nth(index + 1);
    await expect(item).toHaveAttribute('id', `portfolio-history-management-container-${observation.id}`);
    const button = item.getByRole('button', {
      name: await expectedObservationHeaderName(page, observation),
      exact: true,
    });
    await expect(button).toBeVisible();
    await expect(button).toHaveClass(/collapsed/);
    await expect(item.locator(`#portfolio-history-management-${observation.id}`)).not.toHaveClass(/\bshow\b/);
  }
}

/** Builds the exact observation header using the browser's configured local timezone.
 *
 * @param page - Browser page with the timezone used to render the management header.
 * @param observation - Persisted observation supplying the tag and timestamp.
 * @returns The expected `<tag> - <YYYY-MM-DD>` accessible name.
 * @author GPT-6 Luna
 * @author benizzio
 */
async function expectedObservationHeaderName(page: Page, observation: PersistedObservation): Promise<string> {
  const localDate = await page.evaluate((timestamp) => {
    const date = new Date(timestamp);
    const year = String(date.getFullYear()).padStart(4, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }, observation.observation_timestamp.toISOString());

  return `${observation.observation_time_tag} - ${localDate}`;
}

/** Verifies the success notification shown after an observation is persisted. */
async function expectSuccessNotification(page: Page): Promise<void> {
  const toast = page.locator('#toast-notification-container .toast[role="alert"]').filter({
    hasText: 'Portfolio observation data saved successfully.',
  }).last();
  await expect(toast).toBeVisible();
  await expect(toast).toHaveClass(/text-bg-success/);
  await expect(toast.locator('.toast-header strong')).toHaveText('Success');
  await expect(toast.locator('.toast-body')).toHaveText('Portfolio observation data saved successfully.');
}

/** Finds one persisted observation by its portfolio-scoped time tag.
 *
 * @author OpenCode
 */
async function findObservationByTag(
  database: E2eDatabase,
  portfolio: SeededPortfolio,
  timeTag: string,
): Promise<PersistedObservation> {
  const observations = await database.query<PersistedObservation>(
    `SELECT DISTINCT ot.id, ot.observation_timestamp, ot.observation_time_tag
     FROM public.portfolio_allocation_obs_time ot
     JOIN public.portfolio_allocation_fact paf
       ON paf.portfolio_id = ot.portfolio_id AND paf.observation_time_id = ot.id
     WHERE paf.portfolio_id = $1
       AND ot.observation_time_tag = $2`,
    [portfolio.id, timeTag],
  );
  expect(observations).toHaveLength(1);
  return observations[0];
}

/** Verifies the four reloaded rows in persistence-query order and their read-only assets. */
async function expectReloadedObservationRows(form: Locator, observationId: number): Promise<void> {
  const expectedRows = [
    { ticker: EXISTING_BOND_TICKER, name: EXISTING_BOND_NAME, assetClass: 'BONDS', quantity: '30', marketPrice: '100.00000000', total: '3,000.00' },
    { ticker: NEW_BOND_TICKER, name: NEW_BOND_NAME, assetClass: 'BONDS', quantity: '0', marketPrice: '0.00000000', total: '1,000.00' },
    { ticker: NEW_STOCK_TICKER, name: NEW_STOCK_NAME, assetClass: 'STOCKS', quantity: '40', marketPrice: '100.00000000', total: '4,000.00' },
    { ticker: EXISTING_STOCK_TICKER, name: EXISTING_STOCK_NAME, assetClass: 'STOCKS', quantity: '0', marketPrice: '0.00000000', total: '2,000.00' },
  ];

  for (const [index, expected] of expectedRows.entries()) {
    const row = form.locator(`#portfolio-history-management-form-${observationId}-row-${index}`);
    await expect(row).toBeVisible();
    await expect(row.locator('td').nth(0)).toContainText(expected.ticker);
    await expect(row.locator('td').nth(1)).toContainText(expected.name);
    await expect(row.getByRole('combobox', { name: 'Asset', exact: true })).toHaveCount(0);
    await expect(row.getByRole('textbox', { name: 'Asset ticker', exact: true })).toHaveCount(0);
    await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveCount(0);
    await expect(row.getByRole('combobox', { name: 'Class' })).toHaveValue(expected.assetClass);
    await expect(row.getByRole('checkbox')).not.toBeChecked();
    await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue(expected.quantity);
    await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue(expected.marketPrice);
    await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue(expected.total);
  }
}

/** Reads and verifies the observation, facts, assets, and unchanged portfolio in PostgreSQL.
 *
 * @author OpenCode
 */
async function expectPersistedPortfolioHistory(
  database: E2eDatabase,
  portfolio: SeededPortfolio,
  seededAssets: Record<string, SeededAsset>,
  expectedTimestamp: string,
): Promise<PersistedObservation> {
  const observations = await database.query<PersistedObservation>(
    `SELECT DISTINCT ot.id, ot.observation_timestamp, ot.observation_time_tag
     FROM public.portfolio_allocation_obs_time ot
     JOIN public.portfolio_allocation_fact paf
       ON paf.portfolio_id = ot.portfolio_id AND paf.observation_time_id = ot.id
     WHERE paf.portfolio_id = $1
       AND ot.observation_time_tag = $2`,
    [portfolio.id, OBSERVATION_TIME_TAG],
  );
  expect(observations).toHaveLength(1);

  const observation = observations[0];
  expect(observation.observation_time_tag).toBe(OBSERVATION_TIME_TAG);
  expect(observation.observation_timestamp).toEqual(new Date(expectedTimestamp));

  const allocations = await database.query<PersistedAllocation>(
    `SELECT paf.asset_id,
            COALESCE(paf.asset_market_price, 0)::numeric(18,8)::text AS asset_market_price,
            COALESCE(paf.asset_quantity, 0)::numeric(18,8)::text AS asset_quantity,
            paf.cash_reserve,
            paf.class,
            a.name,
            paf.observation_time_id,
            a.ticker,
            paf.total_market_value::text AS total_market_value
     FROM public.portfolio_allocation_fact paf
     JOIN public.asset a ON a.id = paf.asset_id
     WHERE paf.portfolio_id = $1
       AND paf.observation_time_id = $2
     ORDER BY paf.class ASC, paf.total_market_value DESC, a.ticker ASC`,
    [portfolio.id, observation.id],
  );
  expect(allocations).toHaveLength(4);
  expect(allocations).toEqual([
    {
      asset_id: seededAssets[EXISTING_BOND_TICKER].id,
      asset_market_price: '100.00000000',
      asset_quantity: '30.00000000',
      cash_reserve: false,
      class: 'BONDS',
      name: EXISTING_BOND_NAME,
      observation_time_id: observation.id,
      ticker: EXISTING_BOND_TICKER,
      total_market_value: '3000',
    },
    expect.objectContaining({
      asset_market_price: '0.00000000',
      asset_quantity: '0.00000000',
      cash_reserve: false,
      class: 'BONDS',
      name: NEW_BOND_NAME,
      observation_time_id: observation.id,
      ticker: NEW_BOND_TICKER,
      total_market_value: '1000',
    }),
    expect.objectContaining({
      asset_market_price: '100.00000000',
      asset_quantity: '40.00000000',
      cash_reserve: false,
      class: 'STOCKS',
      name: NEW_STOCK_NAME,
      observation_time_id: observation.id,
      ticker: NEW_STOCK_TICKER,
      total_market_value: '4000',
    }),
    {
      asset_id: seededAssets[EXISTING_STOCK_TICKER].id,
      asset_market_price: '0.00000000',
      asset_quantity: '0.00000000',
      cash_reserve: false,
      class: 'STOCKS',
      name: EXISTING_STOCK_NAME,
      observation_time_id: observation.id,
      ticker: EXISTING_STOCK_TICKER,
      total_market_value: '2000',
    },
  ]);

  const totalMarketValue = await database.query<{ total: string }>(
    `SELECT COALESCE(SUM(total_market_value), 0)::text AS total
     FROM public.portfolio_allocation_fact
     WHERE portfolio_id = $1
       AND observation_time_id = $2`,
    [portfolio.id, observation.id],
  );
  expect(totalMarketValue).toEqual([{ total: '10000' }]);

  const assets = await database.query<{ id: number; name: string; ticker: string }>(
    `SELECT id, name, ticker
     FROM public.asset
     ORDER BY ticker ASC`,
  );
  expect(assets).toHaveLength(4);
  expect(assets).toEqual(expect.arrayContaining([
    { id: seededAssets[EXISTING_BOND_TICKER].id, name: EXISTING_BOND_NAME, ticker: EXISTING_BOND_TICKER },
    { id: seededAssets[EXISTING_STOCK_TICKER].id, name: EXISTING_STOCK_NAME, ticker: EXISTING_STOCK_TICKER },
    expect.objectContaining({ name: NEW_BOND_NAME, ticker: NEW_BOND_TICKER }),
    expect.objectContaining({ name: NEW_STOCK_NAME, ticker: NEW_STOCK_TICKER }),
  ]));

  const portfolios = await database.query<{
    allocation_structure: typeof DEFAULT_ALLOCATION_STRUCTURE;
    id: number;
    name: string;
  }>(
    `SELECT id, name, allocation_structure
     FROM public.portfolio
     WHERE id = $1`,
    [portfolio.id],
  );
  expect(portfolios).toEqual([{
    allocation_structure: DEFAULT_ALLOCATION_STRUCTURE,
    id: portfolio.id,
    name: PORTFOLIO_NAME,
  }]);

  return observation;
}

/** Returns the date currently displayed by the browser in its configured local timezone.
 *
 * @author GPT-6 Luna
 */
async function browserLocalCalendarDate(page: Page): Promise<string> {
  return page.evaluate(() => {
    const today = new Date();
    const month = String(today.getMonth() + 1).padStart(2, '0');
    const day = String(today.getDate()).padStart(2, '0');
    return `${today.getFullYear()}-${month}-${day}`;
  });
}

/** Computes the expected UTC instant for browser-local midnight on a supplied calendar date.
 *
 * @author GPT-6 Luna
 */
async function localMidnightTimestamp(page: Page, calendarDate: string): Promise<string> {
  return page.evaluate(date => new Date(`${date}T00:00:00`).toISOString(), calendarDate);
}

/** Verifies history items, their timestamp order, and the expanded observation. */
async function expectHistoryObservationOrder(
  page: Page,
  observations: readonly PersistedObservation[],
  expandedObservationId: number,
): Promise<void> {
  const history = page.locator('#accordion-portfolio-history');
  await expect(history.locator(':scope > .accordion-item')).toHaveCount(observations.length);

  for (const [index, observation] of observations.entries()) {
    const item = history.locator(':scope > .accordion-item').nth(index);
    await expect(item.locator(`#portfolio-allocation-${observation.id}`)).toHaveAttribute(
      'id',
      `portfolio-allocation-${observation.id}`,
    );
    await expect(item.getByRole('button', { name: observation.observation_time_tag, exact: true })).toBeVisible();

    const body = item.locator(`#portfolio-allocation-${observation.id}`);
    if (observation.id === expandedObservationId) {
      await expect(body).toHaveClass(/\bshow\b/);
    } else {
      await expect(body).not.toHaveClass(/\bshow\b/);
    }
  }
}

/** Verifies the history chart values and exercises slice, lowest-level, and center navigation. */
async function exercisePortfolioHistoryChart(page: Page, observationId: number): Promise<void> {
  const canvas = page.locator(`#portfolio-chart-${observationId}`);
  const levelLabel = page.locator(`#hierarchy-level-portfolio-chart-${observationId}`);

  await expectLatestCanvasTextSet(page, canvas, ['BONDS', 'STOCKS', '40%', '60%']);
  await expectChartTooltip(page, canvas, 0, 4_000, 10_000, ['BONDS', '$4,000.00']);

  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 0, 4_000, 10_000));
  await expect(levelLabel).toHaveText('Assets for BONDS');
  await expectLatestCanvasTextContains(page, canvas, ['BOND-B', 'BOND-A', '75%', '25%']);
  await expectChartTooltip(page, canvas, 0, 3_000, 4_000, ['BOND-B', '$3,000.00']);

  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 0, 3_000, 4_000));
  await expect(levelLabel).toHaveText('Assets for BONDS');
  await expectLatestCanvasTextContains(page, canvas, ['BOND-B', 'BOND-A', '75%', '25%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, canvas, ['BONDS', 'STOCKS', '40%', '60%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, canvas, ['BONDS', 'STOCKS', '40%', '60%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 4_000, 6_000, 10_000));
  await expect(levelLabel).toHaveText('Assets for STOCKS');
  await page.mouse.move(0, 0);
  await expectLatestCanvasTextContains(page, canvas, ['STOCK-B', 'STOCK-A', '66.67%', '33.33%']);
  await expectChartTooltip(page, canvas, 0, 4_000, 6_000, ['STOCK-B', '$4,000.00']);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, canvas, ['BONDS', 'STOCKS', '40%', '60%']);
}

/** Verifies the seeded history chart without repeating the lowest-level boundary interaction. */
async function exerciseInitialPortfolioHistoryChart(page: Page, observationId: number): Promise<void> {
  const canvas = page.locator(`#portfolio-chart-${observationId}`);
  const levelLabel = page.locator(`#hierarchy-level-portfolio-chart-${observationId}`);

  await expectLatestCanvasTextSet(page, canvas, ['BONDS', 'STOCKS', '40%', '60%']);
  await expectLatestCanvasPatternState(page, canvas, false);
  await expectChartTooltip(page, canvas, 0, 4_000, 10_000, ['BONDS', '$4,000.00']);

  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 0, 4_000, 10_000));
  await expect(levelLabel).toHaveText('Assets for BONDS');
  await page.mouse.move(0, 0);
  await expectLatestCanvasTextContains(page, canvas, ['BOND-B', 'BOND-A', '75%', '25%']);
  await expectLatestCanvasPatternState(page, canvas, false);
  await expectChartTooltip(page, canvas, 0, 3_000, 4_000, ['BOND-B', '$3,000.00']);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, canvas, ['BONDS', 'STOCKS', '40%', '60%']);
  await expectChartTooltip(page, canvas, 4_000, 6_000, 10_000, ['STOCKS', '$6,000.00']);
  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 4_000, 6_000, 10_000));
  await expect(levelLabel).toHaveText('Assets for STOCKS');
  await page.mouse.move(0, 0);
  await expectLatestCanvasTextContains(page, canvas, ['STOCK-B', 'STOCK-A', '66.67%', '33.33%']);
  await expectLatestCanvasPatternState(page, canvas, false);
  await expectChartTooltip(page, canvas, 0, 4_000, 6_000, ['STOCK-B', '$4,000.00']);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
}

/** Verifies both modified history charts and their class and asset drill-downs. */
async function exerciseModifiedPortfolioHistoryChart(
  page: Page,
  secondObservation: PersistedObservation,
  firstObservation: PersistedObservation,
): Promise<void> {
  await expectModifiedObservationChart(
    page,
    secondObservation,
    '10,000.00',
    ['BONDS', 'STOCKS', '70%', '30%'],
  );

  const secondCanvas = page.locator(`#portfolio-chart-${secondObservation.id}`);
  const secondLevelLabel = page.locator(`#hierarchy-level-portfolio-chart-${secondObservation.id}`);
  await expectChartTooltip(page, secondCanvas, 0, 7_000, 10_000, ['BONDS', '$7,000.00']);
  await clickCanvasPoint(page, secondCanvas, () => getDoughnutSlicePoint(secondCanvas, 0, 7_000, 10_000));
  await expect(secondLevelLabel).toHaveText('Assets for BONDS');
  await expectLatestCanvasTextContains(page, secondCanvas, ['BOND-B', 'BOND-A', '71.43%', '28.57%']);
  await expectLatestCanvasPatternState(page, secondCanvas, false);
  await expectChartTooltip(page, secondCanvas, 0, 5_000, 7_000, ['BOND-B', '$5,000.00']);

  await clickCanvasPoint(page, secondCanvas, () => getDoughnutCenterPoint(secondCanvas));
  await expect(secondLevelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, secondCanvas, ['BONDS', 'STOCKS', '70%', '30%']);
  const secondStocksPoint = await findDoughnutSlicePointByTooltip(page, secondCanvas, ['STOCKS', '$3,000.00']);
  await clickCanvasPoint(page, secondCanvas, () => Promise.resolve(secondStocksPoint));
  await expect(secondLevelLabel).toHaveText('Assets for STOCKS');
  await expectLatestCanvasTextContains(page, secondCanvas, [
    'STOCK-B (cash reserve)',
    'STOCK-A',
    '66.67%',
    '33.33%',
  ]);
  await expectLatestCanvasPatternState(page, secondCanvas, true);
  await expectChartTooltip(page, secondCanvas, 0, 2_000, 3_000, ['STOCK-B (cash reserve)', '$2,000.00']);
  await clickCanvasPoint(page, secondCanvas, () => getDoughnutCenterPoint(secondCanvas));
  await expect(secondLevelLabel).toHaveText('Classes');

  const firstCanvas = page.locator(`#portfolio-chart-${firstObservation.id}`);
  const firstLevelLabel = page.locator(`#hierarchy-level-portfolio-chart-${firstObservation.id}`);
  await page.locator(`#portfolio-allocation-${firstObservation.id}`).locator('xpath=..')
    .getByRole('button', { name: OBSERVATION_TIME_TAG, exact: true }).click();
  await expect(page.locator(`#portfolio-allocation-${firstObservation.id}`)).toHaveClass(/\bshow\b/);
  await expect(page.locator(`#portfolio-allocation-${secondObservation.id}`)).not.toHaveClass(/\bshow\b/);
  await expect(page.locator(`#portfolio-allocation-${firstObservation.id}`)
    .getByText('Total market value: $10,500.00', { exact: true })).toBeVisible();
  await expect(firstLevelLabel).toHaveText('Classes');
  await expectLatestCanvasTextSet(page, firstCanvas, ['BONDS', 'STOCKS', '33.33%', '66.67%']);
  await expectLatestCanvasPatternState(page, firstCanvas, false);
  await expectChartTooltip(page, firstCanvas, 0, 3_500, 10_500, ['BONDS', '$3,500.00']);

  await clickCanvasPoint(page, firstCanvas, () => getDoughnutSlicePoint(firstCanvas, 0, 3_500, 10_500));
  await expect(firstLevelLabel).toHaveText('Assets for BONDS');
  await expectLatestCanvasTextContains(page, firstCanvas, ['BOND-B', '100%']);
  await expectLatestCanvasPatternState(page, firstCanvas, false);
  await expectChartTooltip(page, firstCanvas, 0, 3_500, 3_500, ['BOND-B', '$3,500.00']);

  await clickCanvasPoint(page, firstCanvas, () => getDoughnutCenterPoint(firstCanvas));
  await expectLatestCanvasTextContains(page, firstCanvas, ['BONDS', 'STOCKS', '33.33%', '66.67%']);
  const firstStocksPoint = await findDoughnutSlicePointByTooltip(page, firstCanvas, ['STOCKS', '$7,000.00']);
  await clickCanvasPoint(page, firstCanvas, () => Promise.resolve(firstStocksPoint));
  await expect(firstLevelLabel).toHaveText('Assets for STOCKS');
  await expectLatestCanvasTextContains(page, firstCanvas, [
    'STOCK-B (cash reserve)',
    'STOCK-A',
    'BOND-A',
    '57.14%',
    '28.57%',
    '14.29%',
  ]);
  await expectLatestCanvasPatternState(page, firstCanvas, true);
  await expectChartTooltip(page, firstCanvas, 0, 4_000, 7_000, ['STOCK-B (cash reserve)', '$4,000.00']);
  await clickCanvasPoint(page, firstCanvas, () => getDoughnutCenterPoint(firstCanvas));
  await expect(firstLevelLabel).toHaveText('Classes');
}

/** Verifies one observation summary, class-level labels, and chart pattern state. */
async function expectModifiedObservationChart(
  page: Page,
  observation: PersistedObservation,
  totalMarketValue: string,
  chartTexts: readonly string[],
): Promise<void> {
  const item = page.locator(`#portfolio-allocation-${observation.id}`);
  await expect(item).toHaveClass(/\bshow\b/);
  await expect(item.getByText(`Total market value: $${totalMarketValue}`, { exact: true })).toBeVisible();
  await expect(page.locator(`#hierarchy-level-portfolio-chart-${observation.id}`)).toHaveText('Classes');

  const canvas = page.locator(`#portfolio-chart-${observation.id}`);
  await expect(canvas).toBeVisible();
  await expectLatestCanvasTextSet(page, canvas, chartTexts);
  await expectLatestCanvasPatternState(page, canvas, false);
  await expect(canvas).toHaveAttribute('id', `portfolio-chart-${observation.id}`);
  await expect.poll(async () => canvas.evaluate((element) => {
    const bounds = element.getBoundingClientRect();
    return bounds.width > 0 && bounds.height > 0;
  })).toBe(true);
}

/** Verifies final observation identities, allocation facts, assets, portfolio, and plan absence.
 *
 * @author OpenCode
 */
async function expectPersistedModifiedPortfolioHistory(
  database: E2eDatabase,
  seededData: SeededHistoryModificationData,
  secondObservation: PersistedObservation,
): Promise<void> {
  const observations = await database.query<PersistedObservation>(
    `SELECT DISTINCT ot.id, ot.observation_timestamp, ot.observation_time_tag
     FROM public.portfolio_allocation_obs_time ot
     JOIN public.portfolio_allocation_fact paf
       ON paf.portfolio_id = ot.portfolio_id AND paf.observation_time_id = ot.id
     WHERE paf.portfolio_id = $1
     ORDER BY ot.observation_timestamp ASC`,
    [seededData.portfolio.id],
  );
  expect(observations).toHaveLength(2);

  const firstObservation = observations.find(({ id }) => id === seededData.observation.id);
  expect(firstObservation).toEqual(seededData.observation);
  expect(firstObservation?.observation_time_tag).toBe(OBSERVATION_TIME_TAG);
  expect(firstObservation?.observation_timestamp).toEqual(new Date(FIXED_OBSERVATION_TIMESTAMP));

  const persistedSecondObservation = observations.find(({ id }) => id === secondObservation.id);
  expect(persistedSecondObservation).toEqual(secondObservation);
  expect(persistedSecondObservation?.id).not.toBe(seededData.observation.id);
  expect(persistedSecondObservation?.observation_time_tag).toBe(SECOND_OBSERVATION_TIME_TAG);
  expect(persistedSecondObservation?.observation_timestamp).toBeInstanceOf(Date);
  expect(persistedSecondObservation?.observation_timestamp.getTime())
    .toBeGreaterThan(seededData.observation.observation_timestamp.getTime());

  await expectPersistedAllocations(database, seededData.portfolio, seededData.observation, [
    {
      asset_id: seededData.assets[EXISTING_BOND_TICKER].id,
      asset_market_price: '100.00000000',
      asset_quantity: '35.00000000',
      cash_reserve: false,
      class: 'BONDS',
      name: EXISTING_BOND_NAME,
      observation_time_id: seededData.observation.id,
      ticker: EXISTING_BOND_TICKER,
      total_market_value: '3500',
    },
    {
      asset_id: seededData.assets[NEW_STOCK_TICKER].id,
      asset_market_price: '100.00000000',
      asset_quantity: '40.00000000',
      cash_reserve: true,
      class: 'STOCKS',
      name: NEW_STOCK_NAME,
      observation_time_id: seededData.observation.id,
      ticker: NEW_STOCK_TICKER,
      total_market_value: '4000',
    },
    {
      asset_id: seededData.assets[EXISTING_STOCK_TICKER].id,
      asset_market_price: '0.00000000',
      asset_quantity: '0.00000000',
      cash_reserve: false,
      class: 'STOCKS',
      name: EXISTING_STOCK_NAME,
      observation_time_id: seededData.observation.id,
      ticker: EXISTING_STOCK_TICKER,
      total_market_value: '2000',
    },
    {
      asset_id: seededData.assets[NEW_BOND_TICKER].id,
      asset_market_price: '0.00000000',
      asset_quantity: '0.00000000',
      cash_reserve: false,
      class: 'STOCKS',
      name: NEW_BOND_NAME,
      observation_time_id: seededData.observation.id,
      ticker: NEW_BOND_TICKER,
      total_market_value: '1000',
    },
  ]);

  await expectPersistedAllocations(database, seededData.portfolio, secondObservation, [
    {
      asset_id: seededData.assets[EXISTING_BOND_TICKER].id,
      asset_market_price: '100.00000000',
      asset_quantity: '50.00000000',
      cash_reserve: false,
      class: 'BONDS',
      name: EXISTING_BOND_NAME,
      observation_time_id: secondObservation.id,
      ticker: EXISTING_BOND_TICKER,
      total_market_value: '5000',
    },
    {
      asset_id: seededData.assets[NEW_BOND_TICKER].id,
      asset_market_price: '0.00000000',
      asset_quantity: '0.00000000',
      cash_reserve: false,
      class: 'BONDS',
      name: NEW_BOND_NAME,
      observation_time_id: secondObservation.id,
      ticker: NEW_BOND_TICKER,
      total_market_value: '2000',
    },
    {
      asset_id: seededData.assets[NEW_STOCK_TICKER].id,
      asset_market_price: '100.00000000',
      asset_quantity: '20.00000000',
      cash_reserve: true,
      class: 'STOCKS',
      name: NEW_STOCK_NAME,
      observation_time_id: secondObservation.id,
      ticker: NEW_STOCK_TICKER,
      total_market_value: '2000',
    },
    {
      asset_id: seededData.assets[EXISTING_STOCK_TICKER].id,
      asset_market_price: '0.00000000',
      asset_quantity: '0.00000000',
      cash_reserve: false,
      class: 'STOCKS',
      name: EXISTING_STOCK_NAME,
      observation_time_id: secondObservation.id,
      ticker: EXISTING_STOCK_TICKER,
      total_market_value: '1000',
    },
  ]);

  await expect(database.query<{ count: number }>(
    `SELECT COUNT(*)::int AS count
     FROM public.portfolio_allocation_fact
     WHERE portfolio_id = $1`,
    [seededData.portfolio.id],
  )).resolves.toEqual([{ count: 8 }]);
  await expect(database.query<{ count: number }>(
    `SELECT COUNT(DISTINCT observation_time_id)::int AS count
     FROM public.portfolio_allocation_fact
     WHERE portfolio_id = $1`,
    [seededData.portfolio.id],
  )).resolves.toEqual([{ count: 2 }]);

  await expect(database.query(
    `SELECT asset_id, class, cash_reserve
     FROM public.portfolio_allocation_fact
     WHERE portfolio_id = $1
       AND observation_time_id = $2
       AND ((asset_id = $3 AND class = $4 AND cash_reserve = $5)
         OR (asset_id = $6 AND class = $7 AND cash_reserve = $8))`,
    [
      seededData.portfolio.id,
      seededData.observation.id,
      seededData.assets[NEW_BOND_TICKER].id,
      'BONDS',
      false,
      seededData.assets[NEW_STOCK_TICKER].id,
      'STOCKS',
      false,
    ],
  )).resolves.toEqual([]);

  await expect(database.query<{ id: number; name: string; ticker: string }>(
    `SELECT id, name, ticker
     FROM public.asset
     ORDER BY ticker ASC`,
  )).resolves.toEqual([
    seededData.assets[NEW_BOND_TICKER],
    seededData.assets[EXISTING_BOND_TICKER],
    seededData.assets[EXISTING_STOCK_TICKER],
    seededData.assets[NEW_STOCK_TICKER],
  ]);
  await expect(database.query(
    `SELECT id, name, allocation_structure
     FROM public.portfolio
     WHERE id = $1`,
    [seededData.portfolio.id],
  )).resolves.toEqual([{
    allocation_structure: DEFAULT_ALLOCATION_STRUCTURE,
    id: seededData.portfolio.id,
    name: MODIFICATION_PORTFOLIO_NAME,
  }]);
  await expect(database.query(
    'SELECT id FROM public.allocation_plan WHERE portfolio_id = $1',
    [seededData.portfolio.id],
  )).resolves.toEqual([]);
}

/** Loads one observation's facts in the same deterministic order used by the application. */
async function expectPersistedAllocations(
  database: E2eDatabase,
  portfolio: SeededPortfolio,
  observation: PersistedObservation,
  expectedAllocations: readonly PersistedAllocation[],
): Promise<void> {
  const allocations = await database.query<PersistedAllocation>(
    `SELECT paf.asset_id,
            COALESCE(paf.asset_market_price, 0)::numeric(18,8)::text AS asset_market_price,
            COALESCE(paf.asset_quantity, 0)::numeric(18,8)::text AS asset_quantity,
            paf.cash_reserve,
            paf.class,
            a.name,
            paf.observation_time_id,
            a.ticker,
            paf.total_market_value::text AS total_market_value
     FROM public.portfolio_allocation_fact paf
     JOIN public.asset a ON a.id = paf.asset_id
     WHERE paf.portfolio_id = $1
       AND paf.observation_time_id = $2
     ORDER BY paf.class ASC, paf.cash_reserve DESC, paf.total_market_value DESC, a.ticker ASC`,
    [portfolio.id, observation.id],
  );
  expect(allocations).toEqual(expectedAllocations);
  expect(allocations.reduce((total, allocation) => total + Number(allocation.total_market_value), 0))
    .toBe(expectedAllocations.reduce((total, allocation) => total + Number(allocation.total_market_value), 0));
}

/** Verifies the initial history chart is rendered with its total and Classes hierarchy. */
async function expectRenderedPortfolioHistoryChart(page: Page, observationId: number): Promise<void> {
  const historyItem = page.locator(`#portfolio-allocation-${observationId}`);
  await expect(historyItem).toHaveClass(/\bshow\b/);
  await expect(historyItem.locator('xpath=..').getByRole('button', { name: OBSERVATION_TIME_TAG, exact: true })).toBeVisible();
  await expect(historyItem.getByText('Total market value: $10,000.00', { exact: true })).toBeVisible();
  await expect(page.locator(`#hierarchy-level-portfolio-chart-${observationId}`)).toHaveText('Classes');

  const canvas = page.locator(`#portfolio-chart-${observationId}`);
  await expect(canvas).toBeVisible();
  await expect.poll(async () => canvas.evaluate((element) => {
    const bounds = element.getBoundingClientRect();
    return bounds.width > 0 && bounds.height > 0;
  })).toBe(true);
}

/** Returns the visible label used to click a portfolio submenu option. */
function portfolioNavigationLabel(page: Page, optionName: NavigationOptionName): Locator {
  const option = PORTFOLIO_NAVIGATION_OPTIONS.find(({ name }) => name === optionName);
  if (!option) {
    throw new Error(`Unknown portfolio navigation option: ${optionName}`);
  }
  return page.getByRole('group', { name: 'Navigation' }).locator(`label[for="${option.id}"]`);
}

/** Navigates through a bound clickable element and waits for the expected route. */
async function navigateByClick(page: Page, path: string, target: Locator): Promise<void> {
  await expect(target).toBeVisible();
  await Promise.all([
    page.waitForURL(routePattern(path)),
    target.click(),
  ]);
}

/** Waits until the browser location has reached the expected application path. */
async function expectRoute(page: Page, path: string): Promise<void> {
  await expect(page).toHaveURL(routePattern(path));
}

/** Builds a URL matcher that permits an optional trailing slash but no other path. */
function routePattern(path: string): RegExp {
  const escapedPath = path.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  return new RegExp(`${escapedPath}/?$`);
}
