/**
 * Covers portfolio allocation-plan creation and direct allocation-plan navigation.
 *
 * Scenario 4 creates one hierarchical allocation plan through the management form, verifies
 * its refreshed form, rendered chart, and PostgreSQL persistence. Scenario 4.1 verifies the
 * allocation-plan routes with direct browser URLs without changing persisted data.
 *
 * Authored by: OpenCode
 */
import type { Locator, Page } from '@playwright/test';
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';
import {
  clickCanvasPoint,
  expectChartTooltip,
  expectLatestCanvasPatternState,
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

const PORTFOLIO_NAME = 'E2E Allocation Plan Portfolio';
const ALLOCATION_PLAN_NAME = 'E2E Allocation Plan 001';
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

type SeededData = {
  assets: Record<string, SeededAsset>;
  portfolio: SeededPortfolio;
};

type PersistedAllocationPlan = {
  create_timestamp: Date;
  id: number;
  name: string;
  planned_execution_date: string | null;
  portfolio_id: number;
  type: string;
};

type PersistedPlannedAllocation = {
  asset_id: number | null;
  asset_name: string | null;
  asset_ticker: string | null;
  cash_reserve: boolean;
  hierarchical_id: Array<string | null>;
  slice_size_percentage: string;
  total_market_value: string | null;
};

type ExpectedAllocationRow = {
  assetName?: string;
  assetTicker?: string;
  cashReserve: boolean;
  className: string;
  index: number;
  percentage: string;
  percentageDecimal: string;
  parent?: boolean;
};

test.describe('portfolio allocation plan creation', () => {
  test('scenario 4: creates and explores a portfolio allocation plan', async ({ database, page }) => {
    const seededData = await seedAllocationPlanData(database);
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
      `/portfolio/${seededData.portfolio.id}/allocation`,
      portfolioNavigationLabel(page, 'Allocation Plan'),
    );
    await expectAllocationPlan(page, seededData.portfolio, true);

    const managementButton = page.locator(
      '#accordion-allocation-plan > button[onclick="navigateToPortfolioAllocationPlanManagement()"]',
    );
    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/allocation/manage`,
      managementButton,
    );
    await expectAllocationPlanManagement(page, seededData.portfolio);

    const newPlanItem = page.locator('#allocation-plan-management-container-0');
    const newPlanNameInput = newPlanItem.getByRole('textbox', { name: 'New allocation plan name' });
    await newPlanNameInput.fill(ALLOCATION_PLAN_NAME);
    await expect(newPlanItem.locator('#allocation-plan-management-form-0 input[name="name"]')).toHaveValue(
      ALLOCATION_PLAN_NAME,
    );
    await newPlanItem.locator('#allocation-plan-management-trigger-0 > button').click();

    const newPlanForm = page.locator('#allocation-plan-management-form-0');
    await expect(newPlanItem.locator('#allocation-plan-management-0')).toHaveClass(/\bshow\b/);
    await expect(newPlanForm).toBeVisible();
    await expectAllocationPlanManagementForm(newPlanForm, 0);

    const bondsClassRow = await addClassAllocationRow(page, newPlanForm, 1);
    await fillClassAllocation(bondsClassRow, 1, 'BONDS', '40', '0.4');

    const bondAlphaRow = await addAssetAllocationRow(page, bondsClassRow, 2, 1);
    await fillNewAssetAllocation(
      page,
      bondAlphaRow,
      2,
      NEW_BOND_TICKER,
      NEW_BOND_NAME,
      'BONDS',
      '25',
      '0.25',
      false,
    );

    const bondBetaRow = await addAssetAllocationRow(page, bondsClassRow, 3, 1);
    await fillExistingAssetAllocation(
      page,
      bondBetaRow,
      seededData.assets[EXISTING_BOND_TICKER],
      3,
      'BONDS',
      '75',
      '0.75',
      false,
    );

    const stocksClassRow = await addClassAllocationRow(page, newPlanForm, 4);
    await fillClassAllocation(stocksClassRow, 4, 'STOCKS', '60', '0.6');

    const stockAlphaRow = await addAssetAllocationRow(page, stocksClassRow, 5, 4);
    await fillExistingAssetAllocation(
      page,
      stockAlphaRow,
      seededData.assets[EXISTING_STOCK_TICKER],
      5,
      'STOCKS',
      '33.333',
      '0.33333',
      false,
    );

    const stockBetaRow = await addAssetAllocationRow(page, stocksClassRow, 6, 4);
    await fillNewAssetAllocation(
      page,
      stockBetaRow,
      6,
      NEW_STOCK_TICKER,
      NEW_STOCK_NAME,
      'STOCKS',
      '66.667',
      '0.66667',
      true,
    );

    await expectDraftAllocationRows(newPlanForm);

    const saveResponsePromise = page.waitForResponse((response) => {
      return response.request().method() === 'POST'
        && new URL(response.url()).pathname === `/api/portfolio/${seededData.portfolio.id}/allocation-plan`;
    });
    await newPlanForm.locator('tfoot button.btn-primary[type="button"]').click();
    const saveResponse = await saveResponsePromise;
    expect(saveResponse.status()).toBe(204);

    const persistedPlan = await expectPersistedAllocationPlan(database, seededData);
    await expectAllocationPlanSuccessNotification(page);
    await expectReloadedManagementItems(page, persistedPlan.id);

    const savedPlanItem = page.locator(`#allocation-plan-management-container-${persistedPlan.id}`);
    await savedPlanItem.getByRole('button', { name: ALLOCATION_PLAN_NAME, exact: true }).click();

    const savedPlanForm = page.locator(`#allocation-plan-management-form-${persistedPlan.id}`);
    await expect(savedPlanItem.locator(`#allocation-plan-management-${persistedPlan.id}`)).toHaveClass(/\bshow\b/);
    await expect(savedPlanForm).toBeVisible();
    await expectSavedAllocationPlanForm(savedPlanForm);

    const returnButton = page.locator(
      'button[onclick="allocationPlanManagement.navigateToAllocationPlansViewing()"]',
    );
    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/allocation`,
      returnButton,
    );

    await expectRenderedAllocationPlan(page, seededData.portfolio, persistedPlan.id);
    await exerciseAllocationPlanChart(page, persistedPlan.id);
  });

  test('scenario 4.1: navigates to allocation-plan pages with direct URLs', async ({ database, page }) => {
    const seededData = await seedAllocationPlanData(database);

    await page.goto('/');
    await expectRootShell(page);

    await page.goto('/portfolios');
    await expectPortfolioList(page, seededData.portfolio);

    await page.goto(`/portfolio/${seededData.portfolio.id}`);
    await expectPortfolioContext(page, seededData.portfolio);
    await expectPortfolioNavigation(page, undefined);

    await page.goto(`/portfolio/${seededData.portfolio.id}/allocation`);
    await expectAllocationPlan(page, seededData.portfolio, true);

    await page.goto(`/portfolio/${seededData.portfolio.id}/allocation/manage`);
    await expectAllocationPlanManagement(page, seededData.portfolio);

    await expect(database.query('SELECT id FROM public.allocation_plan WHERE portfolio_id = $1', [seededData.portfolio.id]))
      .resolves.toEqual([]);
    await expect(database.query(
      `SELECT pa.id
       FROM public.planned_allocation pa
       JOIN public.allocation_plan ap ON ap.id = pa.allocation_plan_id
       WHERE ap.portfolio_id = $1`,
      [seededData.portfolio.id],
    )).resolves.toEqual([]);
    await expect(database.query(
      'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
      [seededData.portfolio.id],
    )).resolves.toEqual([]);
    await expect(database.query(
      `SELECT a.id, a.name, a.ticker
       FROM public.asset a
       ORDER BY a.ticker ASC`,
    )).resolves.toEqual([
      { id: seededData.assets[EXISTING_BOND_TICKER].id, name: EXISTING_BOND_NAME, ticker: EXISTING_BOND_TICKER },
      { id: seededData.assets[EXISTING_STOCK_TICKER].id, name: EXISTING_STOCK_NAME, ticker: EXISTING_STOCK_TICKER },
    ]);
    await expect(database.query(
      `SELECT id, name, allocation_structure
       FROM public.portfolio
       WHERE id = $1`,
      [seededData.portfolio.id],
    )).resolves.toEqual([{
      allocation_structure: DEFAULT_ALLOCATION_STRUCTURE,
      id: seededData.portfolio.id,
      name: PORTFOLIO_NAME,
    }]);
  });
});

/** Seeds one portfolio and the two assets scenario 4 must resolve as existing. */
async function seedAllocationPlanData(database: E2eDatabase): Promise<SeededData> {
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

/** Asserts the portfolio list contains one seeded portfolio and the New portfolio card. */
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

/** Asserts the empty allocation-plan view and its management control. */
async function expectAllocationPlan(
  page: Page,
  portfolio: SeededPortfolio,
  expectEmpty: boolean,
): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}/allocation`);
  await expectPortfolioContext(page, portfolio);
  await expectPortfolioNavigation(page, 'Allocation Plan');

  const allocationPlan = page.locator('#accordion-allocation-plan');
  await expect(allocationPlan).toBeVisible();
  if (expectEmpty) {
    await expect(allocationPlan.locator(':scope > *')).toHaveCount(1);
  }

  const managementButton = allocationPlan.locator(
    '> button[onclick="navigateToPortfolioAllocationPlanManagement()"]',
  );
  await expect(managementButton).toBeVisible();
  await expect(managementButton.locator('.bi-database-gear')).toBeVisible();
}

/** Asserts the allocation-plan management shell and its collapsed new-plan item. */
async function expectAllocationPlanManagement(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}/allocation/manage`);
  await expectPortfolioContext(page, portfolio);
  await expectPortfolioNavigation(page, undefined, false);

  const returnButton = page.locator(
    'button[onclick="allocationPlanManagement.navigateToAllocationPlansViewing()"]',
  );
  await expect(returnButton).toBeVisible();
  await expect(returnButton.locator('.bi-pie-chart-fill')).toBeVisible();

  const managementCard = page.locator('.card').filter({ hasText: 'Manage allocation plan data' });
  await expect(managementCard).toBeVisible();
  await expect(managementCard.getByText('Manage allocation plan data', { exact: true })).toBeVisible();

  const managementAccordion = page.locator('#accordion-allocation-plan-management');
  await expect(managementAccordion).toBeVisible();
  await expect(managementAccordion.locator(':scope > .accordion-item')).toHaveCount(1);

  const newPlanItem = page.locator('#allocation-plan-management-container-0');
  await expect(newPlanItem.locator('#allocation-plan-management-trigger-0 > button')).toHaveClass(/collapsed/);
  await expect(newPlanItem.getByRole('textbox', { name: 'New allocation plan name' })).toBeVisible();
}

/** Asserts the allocation-plan management table headers, row count, and icon controls. */
async function expectAllocationPlanManagementForm(
  form: Locator,
  rowCount: number,
): Promise<void> {
  const table = form.getByRole('table');
  await expect(table).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Classes', exact: true })).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Assets', exact: true })).toHaveAttribute('colspan', '2');
  await expect(table.getByRole('columnheader', { name: 'Cash reserve?', exact: true })).toBeVisible();
  await expect(table.getByRole('columnheader', { name: 'Slice size (%)', exact: true })).toBeVisible();
  await expect(form.locator('tbody tr')).toHaveCount(rowCount);

  const addButton = form.locator('tfoot button.btn-secondary[type="button"]');
  await expect(addButton).toBeVisible();
  await expect(addButton.locator('.bi-plus-circle')).toBeVisible();

  const saveButton = form.locator('tfoot button.btn-primary[type="button"]');
  await expect(saveButton).toBeVisible();
  await expect(saveButton.locator('.bi-save-fill')).toBeVisible();
}

/** Adds a top-level class row and waits for its percentage binding. */
async function addClassAllocationRow(page: Page, form: Locator, index: number): Promise<Locator> {
  await form.locator('tfoot button.btn-secondary[type="button"]').click();
  const row = page.locator(`#allocation-plan-management-form-0-row-${index}`);
  await expect(row).toBeVisible();
  await expectPercentageBinding(row);
  return row;
}

/** Adds an asset child row beneath a class row and verifies inherited hierarchy metadata. */
async function addAssetAllocationRow(
  page: Page,
  parentRow: Locator,
  index: number,
  parentIndex: number,
): Promise<Locator> {
  const addButton = parentRow.locator('button.btn-secondary');
  await expect(addButton).toHaveCount(1);
  await addButton.click();

  const row = page.locator(`#allocation-plan-management-form-0-row-${index}`);
  await expect(row).toBeVisible();
  await expect(row).toHaveAttribute(
    'data-parent-row-id',
    `allocation-plan-management-form-0-row-${parentIndex}`,
  );
  await expectPercentageBinding(row);
  return row;
}

/** Verifies that a dynamically inserted percentage field has been initialized. */
async function expectPercentageBinding(row: Locator): Promise<void> {
  await expect(row.getByRole('spinbutton', { name: 'Slice size percentage' })).toHaveAttribute(
    'data-percentage-input-bound',
    'true',
  );
}

/** Fills a class row and verifies its visible and decimal percentage values. */
async function fillClassAllocation(
  row: Locator,
  index: number,
  className: string,
  percentage: string,
  percentageDecimal: string,
): Promise<void> {
  const classInput = row.getByRole('combobox', { name: 'Class' });
  await classInput.fill(className);
  await expect(classInput).toHaveValue(className);
  await fillPercentage(row, index, percentage, percentageDecimal);
  await expect(row.getByRole('checkbox')).not.toBeChecked();
}

/** Fills a new asset row and verifies the creating-asset state and inherited class. */
async function fillNewAssetAllocation(
  page: Page,
  row: Locator,
  index: number,
  ticker: string,
  name: string,
  className: string,
  percentage: string,
  percentageDecimal: string,
  cashReserve: boolean,
): Promise<void> {
  await searchForAsset(page, row, ticker, 404);
  await expect(row.getByText('* Creating new asset', { exact: true })).toBeVisible();
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toBeVisible();

  const tickerInput = row.getByRole('combobox', { name: 'Asset ticker' });
  await expect(tickerInput).toHaveValue(ticker);
  await row.getByRole('textbox', { name: 'Asset name' }).fill(name);
  await expect(row.locator(`input[name="details[${index}][hierarchicalId][1]"]`)).toHaveValue(className);
  await fillPercentage(row, index, percentage, percentageDecimal);
  await setCashReserve(row, cashReserve);
}

/** Fills an existing asset row and verifies the resolved asset and inherited class. */
async function fillExistingAssetAllocation(
  page: Page,
  row: Locator,
  asset: SeededAsset,
  index: number,
  className: string,
  percentage: string,
  percentageDecimal: string,
  cashReserve: boolean,
): Promise<void> {
  await selectAssetFromDatalist(page, row, asset.ticker);
  await searchForAsset(page, row, asset.ticker, 200);
  await expectExistingAsset(row, asset);
  await expect(row.locator(`input[name="details[${index}][hierarchicalId][1]"]`)).toHaveValue(className);
  await fillPercentage(row, index, percentage, percentageDecimal);
  await setCashReserve(row, cashReserve);
}

/** Selects a known ticker from the populated browser datalist. */
async function selectAssetFromDatalist(page: Page, row: Locator, ticker: string): Promise<void> {
  await expect(page.locator(`#datalist-assets option[value="${ticker}"]`)).toBeAttached();
  const tickerInput = row.getByRole('combobox', { name: 'Asset ticker' });
  await tickerInput.click();
  await tickerInput.fill(ticker);
  await expect(tickerInput).toHaveValue(ticker);
}

/** Waits for an asset lookup request and verifies the expected HTTP result. */
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

/** Verifies the read-only fields produced by an existing asset lookup. */
async function expectExistingAsset(row: Locator, asset: SeededAsset): Promise<void> {
  await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveAttribute('readonly', '');
  await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveValue(asset.ticker);
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toBeVisible();
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveAttribute('readonly', '');
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveValue(asset.name);
  await expect(row.locator('input[name$="[asset][id]"]')).toHaveValue(asset.id.toString());
  await expect(row.getByText('* Creating new asset', { exact: true })).toBeHidden();
  await expect(row.locator('[data-asset-action-button] .bi-x-circle')).toBeVisible();
}

/** Fills and verifies one percentage input and its hidden decimal field. */
async function fillPercentage(
  row: Locator,
  index: number,
  percentage: string,
  percentageDecimal: string,
): Promise<void> {
  const percentageInput = row.getByRole('spinbutton', { name: 'Slice size percentage' });
  await percentageInput.fill(percentage);
  await percentageInput.blur();
  await expect(percentageInput).toHaveValue(percentage);
  await expect(row.locator(`input[type="hidden"][name="details[${index}][sliceSizePercentage]"]`))
    .toHaveValue(percentageDecimal);
}

/** Sets and verifies the row's cash-reserve checkbox. */
async function setCashReserve(row: Locator, cashReserve: boolean): Promise<void> {
  const checkbox = row.getByRole('checkbox');
  if (cashReserve) {
    await checkbox.check();
  } else {
    await checkbox.uncheck();
  }
  await expect(checkbox).toBeChecked({ checked: cashReserve });
}

/** Verifies every draft row, including hierarchy, percentages, and cash-reserve state. */
async function expectDraftAllocationRows(form: Locator): Promise<void> {
  await expect(form.locator('tbody tr')).toHaveCount(6);

  const expectedRows: ExpectedAllocationRow[] = [
    { className: 'BONDS', cashReserve: false, index: 1, percentage: '40', percentageDecimal: '0.4', parent: true },
    {
      assetName: NEW_BOND_NAME,
      assetTicker: NEW_BOND_TICKER,
      className: 'BONDS',
      cashReserve: false,
      index: 2,
      percentage: '25',
      percentageDecimal: '0.25',
    },
    {
      assetName: EXISTING_BOND_NAME,
      assetTicker: EXISTING_BOND_TICKER,
      className: 'BONDS',
      cashReserve: false,
      index: 3,
      percentage: '75',
      percentageDecimal: '0.75',
    },
    { className: 'STOCKS', cashReserve: false, index: 4, percentage: '60', percentageDecimal: '0.6', parent: true },
    {
      assetName: EXISTING_STOCK_NAME,
      assetTicker: EXISTING_STOCK_TICKER,
      className: 'STOCKS',
      cashReserve: false,
      index: 5,
      percentage: '33.333',
      percentageDecimal: '0.33333',
    },
    {
      assetName: NEW_STOCK_NAME,
      assetTicker: NEW_STOCK_TICKER,
      className: 'STOCKS',
      cashReserve: true,
      index: 6,
      percentage: '66.667',
      percentageDecimal: '0.66667',
    },
  ];

  for (const expected of expectedRows) {
    await expectDraftAllocationRow(form, expected);
  }
}

/** Verifies one draft allocation row by its stable dynamic index. */
async function expectDraftAllocationRow(form: Locator, expected: ExpectedAllocationRow): Promise<void> {
  const row = form.locator(`#allocation-plan-management-form-0-row-${expected.index}`);
  await expect(row).toBeVisible();
  const classInput = row.locator(`input[name="details[${expected.index}][hierarchicalId][1]"]`);
  await expect(classInput).toHaveValue(expected.className);
  await expect(row.getByRole('checkbox')).toBeChecked({ checked: expected.cashReserve });
  await expect(row.getByRole('spinbutton', { name: 'Slice size percentage' })).toHaveValue(expected.percentage);
  await expect(row.locator(`input[type="hidden"][name="details[${expected.index}][sliceSizePercentage]"]`))
    .toHaveValue(expected.percentageDecimal);

  if (expected.parent) {
    await expect(row.getByRole('combobox', { name: 'Class' })).toHaveValue(expected.className);
    await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveCount(0);
    return;
  }

  if (expected.assetTicker === undefined || expected.assetName === undefined) {
    throw new Error(`Expected asset details are missing for allocation row ${expected.index}`);
  }

  await expect(row.locator(`input[name="details[${expected.index}][hierarchicalId][1]"]`)).toHaveValue(expected.className);
  await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveValue(expected.assetTicker);
  await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveValue(expected.assetName);
}

/** Verifies the refreshed management list and its collapsed new and saved items. */
async function expectReloadedManagementItems(page: Page, allocationPlanId: number): Promise<void> {
  const managementAccordion = page.locator('#accordion-allocation-plan-management');
  await expect(managementAccordion.locator(':scope > .accordion-item')).toHaveCount(2);

  const newPlanItem = page.locator('#allocation-plan-management-container-0');
  await expect(newPlanItem.locator('#allocation-plan-management-trigger-0 > button')).toHaveClass(/collapsed/);
  await expect(newPlanItem.getByRole('textbox', { name: 'New allocation plan name' })).toBeVisible();

  const savedPlanItem = page.locator(`#allocation-plan-management-container-${allocationPlanId}`);
  await expect(savedPlanItem.getByRole('button', { name: ALLOCATION_PLAN_NAME, exact: true })).toBeVisible();
  await expect(savedPlanItem.getByRole('button', { name: ALLOCATION_PLAN_NAME, exact: true })).toHaveClass(/collapsed/);
}

/** Verifies the success toast shown after an allocation plan is persisted. */
async function expectAllocationPlanSuccessNotification(page: Page): Promise<void> {
  const toast = page.locator('#toast-notification-container .toast[role="alert"]').filter({
    hasText: 'Allocation plan saved successfully',
  });
  await expect(toast).toBeVisible();
  await expect(toast).toHaveClass(/text-bg-success/);
  await expect(toast.locator('.toast-header strong')).toHaveText('Success');
  await expect(toast.locator('.toast-body')).toHaveText('Allocation plan saved successfully');
}

/** Verifies the saved form's depth-first rows, editability, and row controls. */
async function expectSavedAllocationPlanForm(form: Locator): Promise<void> {
  await expectAllocationPlanManagementForm(form, 6);

  const expectedRows: ExpectedAllocationRow[] = [
    { className: 'STOCKS', cashReserve: false, index: 0, percentage: '60', percentageDecimal: '0.6', parent: true },
    {
      assetName: NEW_STOCK_NAME,
      assetTicker: NEW_STOCK_TICKER,
      className: 'STOCKS',
      cashReserve: true,
      index: 1,
      percentage: '66.667',
      percentageDecimal: '0.66667',
    },
    {
      assetName: EXISTING_STOCK_NAME,
      assetTicker: EXISTING_STOCK_TICKER,
      className: 'STOCKS',
      cashReserve: false,
      index: 2,
      percentage: '33.333',
      percentageDecimal: '0.33333',
    },
    { className: 'BONDS', cashReserve: false, index: 3, percentage: '40', percentageDecimal: '0.4', parent: true },
    {
      assetName: EXISTING_BOND_NAME,
      assetTicker: EXISTING_BOND_TICKER,
      className: 'BONDS',
      cashReserve: false,
      index: 4,
      percentage: '75',
      percentageDecimal: '0.75',
    },
    {
      assetName: NEW_BOND_NAME,
      assetTicker: NEW_BOND_TICKER,
      className: 'BONDS',
      cashReserve: false,
      index: 5,
      percentage: '25',
      percentageDecimal: '0.25',
    },
  ];

  const rows = form.locator('tbody tr');
  for (const expected of expectedRows) {
    const row = rows.nth(expected.index);
    await expect(row.locator(`input[name="details[${expected.index}][hierarchicalId][1]"]`)).toHaveValue(
      expected.className,
    );
    await expect(row.getByRole('checkbox')).toBeChecked({ checked: expected.cashReserve });
    await expect(row.getByRole('spinbutton', { name: 'Slice size percentage' })).toBeEditable();
    await expect(row.getByRole('spinbutton', { name: 'Slice size percentage' })).toHaveValue(expected.percentage);
    await expect(row.locator(`input[type="hidden"][name="details[${expected.index}][sliceSizePercentage]"]`))
      .toHaveValue(expected.percentageDecimal);

    if (expected.parent) {
      await expect(row.getByRole('combobox', { name: 'Class' })).toBeEditable();
      await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveCount(0);
      continue;
    }

    if (expected.assetTicker === undefined || expected.assetName === undefined) {
      throw new Error(`Expected asset details are missing for allocation row ${expected.index}`);
    }

    await expect(row.getByRole('textbox', { name: 'Class' })).toHaveCount(0);
    await expect(row.getByRole('combobox', { name: 'Asset ticker' })).toHaveCount(0);
    await expect(row.getByRole('textbox', { name: 'Asset name' })).toHaveCount(0);
    await expect(row.locator('td').nth(1)).toHaveText(expected.assetTicker);
    await expect(row.locator('td').nth(2)).toHaveText(expected.assetName);
  }

  await expect(form.locator('tfoot button.btn-secondary[type="button"]')).toHaveCount(1);
  await expect(form.locator('tbody button.btn-secondary[type="button"]')).toHaveCount(2);
  await expect(form.locator('tbody button.btn-danger[type="button"]')).toHaveCount(6);
  await expect(form.locator('tfoot button.btn-primary[type="button"]')).toHaveCount(1);
  await expect(rows.nth(0).locator('button.btn-secondary')).toHaveCount(1);
  await expect(rows.nth(3).locator('button.btn-secondary')).toHaveCount(1);
  await expect(rows.nth(1).locator('button.btn-secondary')).toHaveCount(0);
  await expect(rows.nth(2).locator('button.btn-secondary')).toHaveCount(0);
  await expect(rows.nth(4).locator('button.btn-secondary')).toHaveCount(0);
  await expect(rows.nth(5).locator('button.btn-secondary')).toHaveCount(0);
}

/** Reads the saved allocation plan and verifies its metadata and six persisted details. */
async function expectPersistedAllocationPlan(
  database: E2eDatabase,
  seededData: SeededData,
): Promise<PersistedAllocationPlan> {
  const plans = await database.query<PersistedAllocationPlan>(
    `SELECT id, name, type, planned_execution_date, create_timestamp, portfolio_id
     FROM public.allocation_plan
     WHERE portfolio_id = $1`,
    [seededData.portfolio.id],
  );
  expect(plans).toHaveLength(1);

  const plan = plans[0];
  expect(plan).toEqual(expect.objectContaining({
    id: expect.any(Number),
    name: ALLOCATION_PLAN_NAME,
    planned_execution_date: null,
    portfolio_id: seededData.portfolio.id,
    type: 'ALLOCATION_PLAN',
  }));
  expect(plan.create_timestamp).toBeInstanceOf(Date);

  const allocations = await database.query<PersistedPlannedAllocation>(
    `SELECT pa.hierarchical_id,
            pa.asset_id,
            a.name AS asset_name,
            a.ticker AS asset_ticker,
            pa.cash_reserve,
            pa.slice_size_percentage::numeric(6,5)::text AS slice_size_percentage,
            pa.total_market_value::text AS total_market_value
     FROM public.planned_allocation pa
     LEFT JOIN public.asset a ON a.id = pa.asset_id
     WHERE pa.allocation_plan_id = $1
     ORDER BY pa.hierarchical_id[2] ASC, pa.hierarchical_id[1] NULLS FIRST`,
    [plan.id],
  );
  expect(allocations).toHaveLength(6);
  expect(allocations).toEqual([
    {
      asset_id: null,
      asset_name: null,
      asset_ticker: null,
      cash_reserve: false,
      hierarchical_id: [null, 'BONDS'],
      slice_size_percentage: '0.40000',
      total_market_value: null,
    },
    expect.objectContaining({
      asset_id: expect.any(Number),
      asset_name: NEW_BOND_NAME,
      asset_ticker: NEW_BOND_TICKER,
      cash_reserve: false,
      hierarchical_id: [NEW_BOND_TICKER, 'BONDS'],
      slice_size_percentage: '0.25000',
      total_market_value: null,
    }),
    {
      asset_id: seededData.assets[EXISTING_BOND_TICKER].id,
      asset_name: EXISTING_BOND_NAME,
      asset_ticker: EXISTING_BOND_TICKER,
      cash_reserve: false,
      hierarchical_id: [EXISTING_BOND_TICKER, 'BONDS'],
      slice_size_percentage: '0.75000',
      total_market_value: null,
    },
    {
      asset_id: null,
      asset_name: null,
      asset_ticker: null,
      cash_reserve: false,
      hierarchical_id: [null, 'STOCKS'],
      slice_size_percentage: '0.60000',
      total_market_value: null,
    },
    {
      asset_id: seededData.assets[EXISTING_STOCK_TICKER].id,
      asset_name: EXISTING_STOCK_NAME,
      asset_ticker: EXISTING_STOCK_TICKER,
      cash_reserve: false,
      hierarchical_id: [EXISTING_STOCK_TICKER, 'STOCKS'],
      slice_size_percentage: '0.33333',
      total_market_value: null,
    },
    expect.objectContaining({
      asset_id: expect.any(Number),
      asset_name: NEW_STOCK_NAME,
      asset_ticker: NEW_STOCK_TICKER,
      cash_reserve: true,
      hierarchical_id: [NEW_STOCK_TICKER, 'STOCKS'],
      slice_size_percentage: '0.66667',
      total_market_value: null,
    }),
  ]);

  const assets = await database.query<{ id: number; name: string; ticker: string }>(
    `SELECT id, name, ticker
     FROM public.asset
     ORDER BY ticker ASC`,
  );
  expect(assets).toHaveLength(4);
  expect(assets).toEqual(expect.arrayContaining([
    { id: seededData.assets[EXISTING_BOND_TICKER].id, name: EXISTING_BOND_NAME, ticker: EXISTING_BOND_TICKER },
    { id: seededData.assets[EXISTING_STOCK_TICKER].id, name: EXISTING_STOCK_NAME, ticker: EXISTING_STOCK_TICKER },
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
    [seededData.portfolio.id],
  );
  expect(portfolios).toEqual([{
    allocation_structure: DEFAULT_ALLOCATION_STRUCTURE,
    id: seededData.portfolio.id,
    name: PORTFOLIO_NAME,
  }]);

  await expect(database.query(
    `SELECT DISTINCT ot.id
     FROM public.portfolio_allocation_obs_time ot
     JOIN public.portfolio_allocation_fact paf ON paf.observation_time_id = ot.id
     WHERE paf.portfolio_id = $1`,
    [seededData.portfolio.id],
  )).resolves.toEqual([]);
  await expect(database.query(
    'SELECT portfolio_id FROM public.portfolio_allocation_fact WHERE portfolio_id = $1',
    [seededData.portfolio.id],
  )).resolves.toEqual([]);

  return plan;
}

/** Asserts the allocation-plan view, its selected navigation, and the rendered chart shell. */
async function expectRenderedAllocationPlan(
  page: Page,
  portfolio: SeededPortfolio,
  allocationPlanId: number,
): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}/allocation`);
  await expectPortfolioContext(page, portfolio);
  await expectPortfolioNavigation(page, 'Allocation Plan');

  const allocationPlanItem = page.locator(`#allocation-plan-${allocationPlanId}`);
  await expect(allocationPlanItem).toHaveClass(/\bshow\b/);
  await expect(page.locator(`#hierarchy-level-allocation-plan-chart-${allocationPlanId}`)).toHaveText('Classes');

  const canvas = page.locator(`#allocation-plan-chart-${allocationPlanId}`);
  await expect(canvas).toBeVisible();
  await expect.poll(async () => canvas.evaluate((element) => {
    const bounds = element.getBoundingClientRect();
    return bounds.width > 0 && bounds.height > 0;
  })).toBe(true);
}

/** Verifies class and asset chart values, drill-down boundaries, and cash-reserve pattern rendering. */
async function exerciseAllocationPlanChart(page: Page, allocationPlanId: number): Promise<void> {
  const canvas = page.locator(`#allocation-plan-chart-${allocationPlanId}`);
  const levelLabel = page.locator(`#hierarchy-level-allocation-plan-chart-${allocationPlanId}`);

  await expectLatestCanvasTextSet(page, canvas, ['STOCKS', 'BONDS', '60%', '40%']);
  await expectLatestCanvasPatternState(page, canvas, false);
  await expectChartTooltip(page, canvas, 0.6, 0.4, 1, ['BONDS', '40%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 0.6, 0.4, 1));
  await expect(levelLabel).toHaveText('Assets for BONDS');
  await expectLatestCanvasTextContains(page, canvas, ['BOND-B', 'BOND-A', '75%', '25%']);
  await expectLatestCanvasPatternState(page, canvas, false);
  await page.mouse.move(0, 0);
  await expectChartTooltip(page, canvas, 0, 0.75, 1, ['BOND-B', '75%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 0, 0.75, 1));
  await expect(levelLabel).toHaveText('Assets for BONDS');
  await expectLatestCanvasTextContains(page, canvas, ['BOND-B', 'BOND-A', '75%', '25%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, canvas, ['STOCKS', 'BONDS', '60%', '40%']);
  await expectLatestCanvasPatternState(page, canvas, false);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, canvas, ['STOCKS', 'BONDS', '60%', '40%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutSlicePoint(canvas, 0, 0.6, 1));
  await expect(levelLabel).toHaveText('Assets for STOCKS');
  await expectLatestCanvasTextContains(page, canvas, ['STOCK-B', 'STOCK-A', '66.67%', '33.33%']);
  await expectLatestCanvasPatternState(page, canvas, true);
  await page.mouse.move(0, 0);
  await expectChartTooltip(page, canvas, 0, 0.66667, 1, ['STOCK-B', '66.67%']);

  await clickCanvasPoint(page, canvas, () => getDoughnutCenterPoint(canvas));
  await expect(levelLabel).toHaveText('Classes');
  await expectLatestCanvasTextContains(page, canvas, ['STOCKS', 'BONDS', '60%', '40%']);
  await expectLatestCanvasPatternState(page, canvas, false);
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
