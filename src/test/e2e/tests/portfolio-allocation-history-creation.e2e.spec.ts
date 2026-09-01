/**
 * Covers portfolio allocation history creation and direct history navigation.
 *
 * Scenario 3 creates one observation through the allocation-history management
 * form, verifies the rendered history chart, and checks PostgreSQL persistence.
 * Scenario 3.1 verifies the history routes with direct browser URLs without
 * changing persisted data.
 *
 * Authored by: OpenCode
 */
import type { Locator, Page } from '@playwright/test';
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';
import {
  clickCanvasPoint,
  expectChartTooltip,
  expectLatestCanvasTextContains,
  expectLatestCanvasTextSet,
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
const OBSERVATION_TIME_TAG = 'E2E_OBS_001';
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

test.describe('portfolio allocation history creation', () => {
  test('scenario 3: creates and explores a portfolio allocation observation', async ({ database, page }) => {
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
    await newObservationItem.getByRole('button').click();

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

    const saveStartedAt = Date.now();
    const saveResponsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'POST'
        && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/history`;
    });
    await newObservationForm.locator('tfoot button.btn-primary[type="submit"]').click();
    const saveResponse = await saveResponsePromise;
    const saveFinishedAt = Date.now();
    expect(saveResponse.status()).toBe(204);

    const persistedObservation = await expectPersistedPortfolioHistory(
      database,
      seededData.portfolio,
      seededData.assets,
      saveStartedAt,
      saveFinishedAt,
    );

    await expectSuccessNotification(page);
    await expectReloadedManagementItems(page, persistedObservation.id);

    const existingObservationItem = page.locator(
      `#portfolio-history-management-container-${persistedObservation.id}`,
    );
    await existingObservationItem.getByRole('button', { name: OBSERVATION_TIME_TAG, exact: true }).click();

    const existingObservationForm = page.locator(
      `#portfolio-history-management-form-${persistedObservation.id}`,
    );
    await expect(existingObservationForm.locator('tbody tr')).toHaveCount(4);
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
});

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
async function expectPortfolioHistoryManagement(page: Page, portfolio: SeededPortfolio): Promise<void> {
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
  await expect(managementAccordion.locator(':scope > .accordion-item')).toHaveCount(1);

  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  const newObservationButton = newObservationItem.getByRole('button');
  await expect(newObservationButton).toHaveClass(/collapsed/);
  await expect(newObservationItem.getByRole('textbox', { name: 'Time tag' })).toBeVisible();
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
): Promise<void> {
  await selectAssetFromDatalist(page, row, asset.ticker);
  await searchForAsset(page, row, asset.ticker, 200);
  await expectExistingAsset(row, asset);

  await row.getByRole('combobox', { name: 'Class' }).fill(assetClass);
  await fillCalculatedValues(row, quantity, marketPrice, '3,000.00');
}

/** Fills an existing asset row with a direct total market value. */
async function fillExistingAssetDirectValueAllocation(
  page: Page,
  row: Locator,
  asset: SeededAsset,
  assetClass: string,
  totalMarketValue: string,
): Promise<void> {
  await selectAssetFromDatalist(page, row, asset.ticker);
  await searchForAsset(page, row, asset.ticker, 200);
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

/** Selects one known ticker through the populated browser datalist input. */
async function selectAssetFromDatalist(page: Page, row: Locator, ticker: string): Promise<void> {
  await expect(page.locator(`#datalist-assets option[value="${ticker}"]`)).toBeAttached();
  const tickerInput = row.getByRole('combobox', { name: 'Asset ticker' });
  await tickerInput.click();
  await tickerInput.fill(ticker);
  await expect(tickerInput).toHaveValue(ticker);
}

/** Waits for the asset lookup and verifies its HTTP result. */
async function searchForAsset(page: Page, row: Locator, ticker: string, status: number): Promise<void> {
  const tickerInput = row.getByRole('combobox', { name: 'Asset ticker' });
  await tickerInput.fill(ticker);

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
  await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveAttribute('readonly', '');
  await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveValue(asset.ticker);
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

/** Verifies all four rows before saving the new observation. */
async function expectNewObservationRows(form: Locator): Promise<void> {
  const rows = form.locator('tbody tr');
  await expect(rows).toHaveCount(4);

  const expectedRows = [
    { ticker: NEW_BOND_TICKER, name: NEW_BOND_NAME, assetClass: 'BONDS', quantity: '', marketPrice: '', total: '1,000.00' },
    { ticker: EXISTING_BOND_TICKER, name: EXISTING_BOND_NAME, assetClass: 'BONDS', quantity: '30', marketPrice: '100.00000000', total: '3,000.00' },
    { ticker: EXISTING_STOCK_TICKER, name: EXISTING_STOCK_NAME, assetClass: 'STOCKS', quantity: '', marketPrice: '', total: '2,000.00' },
    { ticker: NEW_STOCK_TICKER, name: NEW_STOCK_NAME, assetClass: 'STOCKS', quantity: '40', marketPrice: '100.00000000', total: '4,000.00' },
  ];

  for (const [index, expected] of expectedRows.entries()) {
    const row = rows.nth(index);
    await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveValue(expected.ticker);
    await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveValue(expected.name);
    await expect(row.getByRole('combobox', { name: 'Class' })).toHaveValue(expected.assetClass);
    await expect(row.getByRole('checkbox')).not.toBeChecked();
    await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue(expected.quantity);
    await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue(expected.marketPrice);
    await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue(expected.total);
  }
}

/** Verifies the management list refresh and collapsed new and persisted observation items. */
async function expectReloadedManagementItems(page: Page, observationId: number): Promise<void> {
  const managementAccordion = page.locator('#accordion-portfolio-history-management');
  await expect(managementAccordion.locator(':scope > .accordion-item')).toHaveCount(2);

  const newObservationItem = page.locator('#portfolio-history-management-container-0');
  await expect(newObservationItem.locator('#portfolio-history-management-trigger-0 > button')).toHaveClass(/collapsed/);
  await expect(newObservationItem.getByRole('textbox', { name: 'Time tag' })).toBeVisible();

  const observationItem = page.locator(`#portfolio-history-management-container-${observationId}`);
  const observationButton = observationItem.getByRole('button', { name: OBSERVATION_TIME_TAG, exact: true });
  await expect(observationButton).toBeVisible();
  await expect(observationButton).toHaveClass(/collapsed/);
  await expect(observationItem.locator(`#portfolio-history-management-${observationId}`)).not.toHaveClass(/\bshow\b/);
}

/** Verifies the success notification shown after an observation is persisted. */
async function expectSuccessNotification(page: Page): Promise<void> {
  const toast = page.locator('#toast-notification-container .toast[role="alert"]').filter({
    hasText: 'Portfolio observation data saved successfully.',
  });
  await expect(toast).toBeVisible();
  await expect(toast).toHaveClass(/text-bg-success/);
  await expect(toast.locator('.toast-header strong')).toHaveText('Success');
  await expect(toast.locator('.toast-body')).toHaveText('Portfolio observation data saved successfully.');
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
    await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveCount(0);
    await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveCount(0);
    await expect(row.getByRole('combobox', { name: 'Class' })).toHaveValue(expected.assetClass);
    await expect(row.getByRole('checkbox')).not.toBeChecked();
    await expect(row.getByRole('spinbutton', { name: 'Quantity' })).toHaveValue(expected.quantity);
    await expect(row.getByRole('textbox', { name: 'Market price' })).toHaveValue(expected.marketPrice);
    await expect(row.getByRole('textbox', { name: 'Total market value' })).toHaveValue(expected.total);
  }
}

/** Reads and verifies the observation, facts, assets, and unchanged portfolio in PostgreSQL. */
async function expectPersistedPortfolioHistory(
  database: E2eDatabase,
  portfolio: SeededPortfolio,
  seededAssets: Record<string, SeededAsset>,
  saveStartedAt: number,
  saveFinishedAt: number,
): Promise<PersistedObservation> {
  const observations = await database.query<PersistedObservation>(
    `SELECT DISTINCT ot.id, ot.observation_timestamp, ot.observation_time_tag
     FROM public.portfolio_allocation_obs_time ot
     JOIN public.portfolio_allocation_fact paf ON paf.observation_time_id = ot.id
     WHERE paf.portfolio_id = $1
       AND ot.observation_time_tag = $2`,
    [portfolio.id, OBSERVATION_TIME_TAG],
  );
  expect(observations).toHaveLength(1);

  const observation = observations[0];
  expect(observation.observation_time_tag).toBe(OBSERVATION_TIME_TAG);
  expect(observation.observation_timestamp).toBeInstanceOf(Date);
  expect(observation.observation_timestamp.getTime()).toBeGreaterThanOrEqual(saveStartedAt - 5_000);
  expect(observation.observation_timestamp.getTime()).toBeLessThanOrEqual(saveFinishedAt + 5_000);

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
