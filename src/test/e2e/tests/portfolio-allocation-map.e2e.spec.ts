/**
 * Covers portfolio allocation-map divergence analysis.
 *
 * Scenario 7 seeds two observations and two allocation plans directly in PostgreSQL, compares
 * every history/plan combination through the browser, verifies hierarchy expansion and progress
 * indicators, and confirms that analysis remains read-only.
 *
 * Authored by: OpenCode
 */
import type { Locator, Page } from '@playwright/test';
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';

const DEFAULT_ALLOCATION_STRUCTURE = {
  hierarchy: [
    { name: 'Assets', field: 'assetTicker' },
    { name: 'Classes', field: 'class' },
  ],
} as const;

const PORTFOLIO_NAME = 'E2E Allocation Map Portfolio';
const HISTORY_001_NAME = 'E2E Allocation Map History 001';
const HISTORY_002_NAME = 'E2E Allocation Map History 002';
const PLAN_001_NAME = 'E2E Allocation Map Plan 001';
const PLAN_002_NAME = 'E2E Allocation Map Plan 002';

const ASSET_DATA = [
  { ticker: 'E2E:CORE-A', name: 'Core Alpha' },
  { ticker: 'E2E:CORE-B', name: 'Core Beta' },
  { ticker: 'E2E:CORE-C', name: 'Core Gamma' },
  { ticker: 'E2E:LEGACY-A', name: 'Legacy Alpha' },
  { ticker: 'E2E:NEW-A', name: 'New Alpha' },
] as const;

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

type SeededObservation = {
  id: number;
  timeTag: string;
};

type SeededPlan = {
  createTimestamp: string;
  id: number;
  name: string;
};

type SeededAllocationMapData = {
  assets: Record<string, SeededAsset>;
  observations: {
    history001: SeededObservation;
    history002: SeededObservation;
  };
  plans: {
    plan001: SeededPlan;
    plan002: SeededPlan;
  };
  portfolio: SeededPortfolio;
};

type ExpectedDivergenceNode = {
  barClass: 'bg-danger' | 'bg-success';
  barWidth: number;
  children?: readonly ExpectedDivergenceNode[];
  divergence: string;
  planned: string;
  total: string;
  unit: string;
};

type DatabaseSnapshot = {
  allocations: readonly {
    allocation_plan_id: number;
    asset_id: number | null;
    cash_reserve: boolean;
    hierarchical_id: readonly (string | null)[];
    id: number;
    slice_size_percentage: string;
    total_market_value: string | null;
  }[];
  assets: readonly {
    id: number;
    name: string;
    ticker: string;
  }[];
  counts: {
    allocation_plans: number;
    assets: number;
    observations: number;
    planned_allocations: number;
    portfolio_allocation_facts: number;
    portfolios: number;
  };
  facts: readonly {
    asset_id: number;
    asset_market_price: string;
    asset_quantity: string;
    cash_reserve: boolean;
    class: string;
    observation_time_id: number;
    portfolio_id: number;
    total_market_value: number;
  }[];
  observations: readonly {
    id: number;
    observation_time_tag: string;
    observation_timestamp: string;
  }[];
  plans: readonly {
    create_timestamp: string;
    id: number;
    name: string;
    planned_execution_date: string | null;
    portfolio_id: number;
    type: string;
  }[];
  portfolios: readonly {
    allocation_structure: typeof DEFAULT_ALLOCATION_STRUCTURE;
    id: number;
    name: string;
  }[];
};

test.describe('portfolio allocation map', () => {
  test('scenario 7: analyzes divergence between portfolio history and allocation plans', async ({ database, page }) => {
    test.setTimeout(120_000);
    const seededData = await seedAllocationMapData(database);
    const initialSnapshot = await readDatabaseSnapshot(database);
    expectSeededDatabase(initialSnapshot, seededData);

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
    await expectPortfolioNavigation(page, undefined, false);

    await navigateByClick(
      page,
      `/portfolio/${seededData.portfolio.id}/allocation-map`,
      portfolioNavigationLabel(page, 'Allocation Map'),
    );
    await expectAllocationMapShell(page, seededData);

    const history002 = observationItem(page, seededData.observations.history002);
    const history001 = observationItem(page, seededData.observations.history001);
    await expectPlanControls(page, seededData);

    const history002Plan002 = await generateAnalysis(
      page,
      seededData.portfolio.id,
      history002,
      seededData.observations.history002,
      seededData.plans.plan002,
      history002Plan002Roots(),
    );
    await expandRootAndAssertChildren(history002Plan002, history002Plan002Roots()[0]);
    await expandRootAndAssertChildren(history002Plan002, history002Plan002Roots()[1]);

    const history002Plan001 = await generateAnalysis(
      page,
      seededData.portfolio.id,
      history002,
      seededData.observations.history002,
      seededData.plans.plan001,
      history002Plan001Roots(),
    );
    await expandRootAndAssertChildren(history002Plan001, history002Plan001Roots()[0]);
    await expandRootAndAssertChildren(history002Plan001, history002Plan001Roots()[1]);
    await expandRootAndAssertChildren(history002Plan001, history002Plan001Roots()[2]);

    await history001.locator('xpath=..').getByRole('button', { name: HISTORY_001_NAME, exact: true }).click();
    await expect(history002.locator('xpath=..')).not.toHaveClass(/show/);
    await expect(history001).toHaveClass(/show/);
    await expect(history001.locator(`select[name="${seededData.observations.history001.id}PlanDivergence"]`))
      .toHaveValue('0');
    await expect(history001.locator('table')).toHaveCount(0);

    const history001Plan001 = await generateAnalysis(
      page,
      seededData.portfolio.id,
      history001,
      seededData.observations.history001,
      seededData.plans.plan001,
      history001Plan001Roots(),
    );
    await expandRootAndAssertChildren(history001Plan001, history001Plan001Roots()[0]);
    await expandRootAndAssertChildren(history001Plan001, history001Plan001Roots()[1]);

    const history001Plan002 = await generateAnalysis(
      page,
      seededData.portfolio.id,
      history001,
      seededData.observations.history001,
      seededData.plans.plan002,
      history001Plan002Roots(),
    );
    await expandRootAndAssertChildren(history001Plan002, history001Plan002Roots()[0]);
    await expandRootAndAssertChildren(history001Plan002, history001Plan002Roots()[1]);
    await expandRootAndAssertChildren(history001Plan002, history001Plan002Roots()[2]);

    await history002.locator('xpath=..').getByRole('button', { name: HISTORY_002_NAME, exact: true }).click();
    await expect(history001).not.toHaveClass(/show/);
    await expect(history002).toHaveClass(/show/);
    await expect(history002.locator(`select[name="${seededData.observations.history002.id}PlanDivergence"]`))
      .toHaveValue(seededData.plans.plan001.id.toString());
    await expect(history002.locator('table')).toHaveCount(1);
    await expectRetainedAnalysisTable(history002.locator('table'), history002Plan001Roots());

    await expect(history001.locator(`select[name="${seededData.observations.history001.id}PlanDivergence"]`))
      .toHaveValue(seededData.plans.plan002.id.toString());
    await expect(history001.locator('table')).toHaveCount(1);
    await expectRetainedAnalysisTable(history001.locator('table'), history001Plan002Roots());

    const finalSnapshot = await readDatabaseSnapshot(database);
    expect(finalSnapshot).toEqual(initialSnapshot);
    expect(finalSnapshot.counts).toEqual({
      allocation_plans: 2,
      assets: 5,
      observations: 2,
      planned_allocations: 10,
      portfolio_allocation_facts: 6,
      portfolios: 1,
    });
  });
});

/** Seeds the exact portfolio, history, plans, and allocations required by Scenario 7. */
async function seedAllocationMapData(database: E2eDatabase): Promise<SeededAllocationMapData> {
  const portfolioRows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [PORTFOLIO_NAME, JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );
  const assetRows = await database.query<SeededAsset>(
    `INSERT INTO public.asset (ticker, name)
     VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8), ($9, $10)
     RETURNING id, ticker, name`,
    ASSET_DATA.flatMap(({ ticker, name }) => [ticker, name]),
  );
  const observationRows = await database.query<{ id: number; observation_time_tag: string }>(
    `INSERT INTO public.portfolio_allocation_obs_time (observation_time_tag, observation_timestamp)
     VALUES ($1, $2::timestamptz), ($3, $4::timestamptz)
     RETURNING id, observation_time_tag`,
    [
      HISTORY_001_NAME,
      '2040-01-01T00:00:00Z',
      HISTORY_002_NAME,
      '2040-02-01T00:00:00Z',
    ],
  );

  expect(portfolioRows).toHaveLength(1);
  expect(assetRows).toHaveLength(ASSET_DATA.length);
  expect(observationRows).toHaveLength(2);

  const assets = Object.fromEntries(assetRows.map((asset) => [asset.ticker, asset])) as Record<string, SeededAsset>;
  const observations = Object.fromEntries(
    observationRows.map((observation) => [observation.observation_time_tag, {
      id: observation.id,
      timeTag: observation.observation_time_tag,
    }]),
  ) as Record<string, SeededObservation>;

  await database.query(
    `INSERT INTO public.portfolio_allocation_fact
       (portfolio_id, asset_id, class, cash_reserve, asset_quantity, asset_market_price,
        total_market_value, observation_time_id)
     VALUES
       ($1, $2, 'CORE', false, 2400, 1, 2400, $3),
       ($1, $4, 'CORE', false, 3600, 1, 3600, $3),
       ($1, $5, 'LEGACY', false, 4000, 1, 4000, $3),
       ($1, $6, 'CORE', false, 1000, 1, 1000, $7),
       ($1, $8, 'CORE', false, 3000, 1, 3000, $7),
       ($1, $9, 'NEW', false, 6000, 1, 6000, $7)`,
    [
      portfolioRows[0].id,
      assets['E2E:CORE-A'].id,
      observations[HISTORY_001_NAME].id,
      assets['E2E:CORE-B'].id,
      assets['E2E:LEGACY-A'].id,
      assets['E2E:CORE-A'].id,
      observations[HISTORY_002_NAME].id,
      assets['E2E:CORE-C'].id,
      assets['E2E:NEW-A'].id,
    ],
  );

  const planRows = await database.query<{ create_timestamp: string; id: number; name: string }>(
    `INSERT INTO public.allocation_plan
       (portfolio_id, name, type, planned_execution_date, create_timestamp)
     VALUES
       ($1, $2, 'ALLOCATION_PLAN', NULL, $3::timestamp),
       ($1, $4, 'ALLOCATION_PLAN', NULL, $5::timestamp)
     RETURNING id, name, to_char(create_timestamp, 'YYYY-MM-DD HH24:MI:SS.US') AS create_timestamp`,
    [
      portfolioRows[0].id,
      PLAN_001_NAME,
      '2040-01-10T00:00:00',
      PLAN_002_NAME,
      '2040-02-10T00:00:00',
    ],
  );
  expect(planRows).toHaveLength(2);

  const plans = Object.fromEntries(
    planRows.map((plan) => [plan.name, {
      createTimestamp: plan.create_timestamp,
      id: plan.id,
      name: plan.name,
    }]),
  ) as Record<string, SeededPlan>;

  await database.query(
    `INSERT INTO public.planned_allocation
       (allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
     VALUES
       ($1, ARRAY[NULL::text, 'CORE'], NULL, false, 0.60000, NULL),
       ($1, ARRAY['E2E:CORE-A', 'CORE'], $2, false, 0.40000, NULL),
       ($1, ARRAY['E2E:CORE-B', 'CORE'], $3, false, 0.60000, NULL),
       ($1, ARRAY[NULL::text, 'LEGACY'], NULL, false, 0.40000, NULL),
       ($1, ARRAY['E2E:LEGACY-A', 'LEGACY'], $4, false, 1.00000, NULL),
       ($5, ARRAY[NULL::text, 'CORE'], NULL, false, 0.40000, NULL),
       ($5, ARRAY['E2E:CORE-A', 'CORE'], $2, false, 0.25000, NULL),
       ($5, ARRAY['E2E:CORE-C', 'CORE'], $6, false, 0.75000, NULL),
       ($5, ARRAY[NULL::text, 'NEW'], NULL, false, 0.60000, NULL),
       ($5, ARRAY['E2E:NEW-A', 'NEW'], $7, false, 1.00000, NULL)`,
    [
      plans[PLAN_001_NAME].id,
      assets['E2E:CORE-A'].id,
      assets['E2E:CORE-B'].id,
      assets['E2E:LEGACY-A'].id,
      plans[PLAN_002_NAME].id,
      assets['E2E:CORE-C'].id,
      assets['E2E:NEW-A'].id,
    ],
  );

  return {
    assets,
    observations: {
      history001: observations[HISTORY_001_NAME],
      history002: observations[HISTORY_002_NAME],
    },
    plans: {
      plan001: plans[PLAN_001_NAME],
      plan002: plans[PLAN_002_NAME],
    },
    portfolio: portfolioRows[0],
  };
}

/** Reads every persisted record affected by Scenario 7 into a deterministic snapshot. */
async function readDatabaseSnapshot(database: E2eDatabase): Promise<DatabaseSnapshot> {
  const [counts, portfolios, assets, observations, facts, plans, allocations] = await Promise.all([
    database.query<DatabaseSnapshot['counts']>(
      `SELECT
         (SELECT COUNT(*)::int FROM public.portfolio) AS portfolios,
         (SELECT COUNT(*)::int FROM public.asset) AS assets,
         (SELECT COUNT(*)::int FROM public.portfolio_allocation_obs_time) AS observations,
         (SELECT COUNT(*)::int FROM public.portfolio_allocation_fact) AS portfolio_allocation_facts,
         (SELECT COUNT(*)::int FROM public.allocation_plan) AS allocation_plans,
         (SELECT COUNT(*)::int FROM public.planned_allocation) AS planned_allocations`,
    ),
    database.query<DatabaseSnapshot['portfolios'][number]>(
      `SELECT id, name, allocation_structure
       FROM public.portfolio
       ORDER BY id`,
    ),
    database.query<DatabaseSnapshot['assets'][number]>(
      `SELECT id, name, ticker
       FROM public.asset
       ORDER BY id`,
    ),
    database.query<DatabaseSnapshot['observations'][number]>(
      `SELECT id,
              observation_time_tag,
              to_char(observation_timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS observation_timestamp
       FROM public.portfolio_allocation_obs_time
       ORDER BY id`,
    ),
    database.query<DatabaseSnapshot['facts'][number]>(
      `SELECT portfolio_id,
              asset_id,
              class,
              cash_reserve,
              asset_quantity::text AS asset_quantity,
              asset_market_price::text AS asset_market_price,
              total_market_value,
              observation_time_id
       FROM public.portfolio_allocation_fact
       ORDER BY observation_time_id, class, asset_id`,
    ),
    database.query<DatabaseSnapshot['plans'][number]>(
      `SELECT id,
              portfolio_id,
              name,
              type,
              planned_execution_date,
              to_char(create_timestamp, 'YYYY-MM-DD HH24:MI:SS.US') AS create_timestamp
       FROM public.allocation_plan
       ORDER BY id`,
    ),
    database.query<DatabaseSnapshot['allocations'][number]>(
      `SELECT id,
              allocation_plan_id,
              hierarchical_id,
              asset_id,
              cash_reserve,
              slice_size_percentage::numeric(6,5)::text AS slice_size_percentage,
              total_market_value::text AS total_market_value
       FROM public.planned_allocation
       ORDER BY allocation_plan_id, hierarchical_id[2], hierarchical_id[1] NULLS FIRST`,
    ),
  ]);

  return {
    allocations,
    assets,
    counts: counts[0],
    facts,
    observations,
    plans,
    portfolios,
  };
}

/** Verifies the complete Scenario 7 seed and its required row counts. */
function expectSeededDatabase(snapshot: DatabaseSnapshot, seededData: SeededAllocationMapData): void {
  expect(snapshot.counts).toEqual({
    allocation_plans: 2,
    assets: 5,
    observations: 2,
    planned_allocations: 10,
    portfolio_allocation_facts: 6,
    portfolios: 1,
  });
  expect(snapshot.portfolios).toEqual([{
    allocation_structure: DEFAULT_ALLOCATION_STRUCTURE,
    id: seededData.portfolio.id,
    name: PORTFOLIO_NAME,
  }]);
  expect(snapshot.assets).toEqual(ASSET_DATA.map(({ name, ticker }) => ({
    id: seededData.assets[ticker].id,
    name,
    ticker,
  })));
  expect(snapshot.observations.map(({ observation_time_tag, observation_timestamp }) => [
    observation_time_tag,
    observation_timestamp,
  ])).toEqual([
    [HISTORY_001_NAME, '2040-01-01T00:00:00Z'],
    [HISTORY_002_NAME, '2040-02-01T00:00:00Z'],
  ]);
  expect(snapshot.facts).toHaveLength(6);
  expect(snapshot.facts.every(({ asset_market_price, asset_quantity, cash_reserve, portfolio_id }) => {
    return Number(asset_market_price) === 1 && cash_reserve === false && portfolio_id === seededData.portfolio.id
      && [2400, 3600, 4000, 1000, 3000, 6000].includes(Number(asset_quantity));
  })).toBe(true);
  expect(snapshot.plans.map(({ name, type, planned_execution_date, portfolio_id }) => ({
    name,
    planned_execution_date,
    portfolio_id,
    type,
  }))).toEqual([
    {
      name: PLAN_001_NAME,
      planned_execution_date: null,
      portfolio_id: seededData.portfolio.id,
      type: 'ALLOCATION_PLAN',
    },
    {
      name: PLAN_002_NAME,
      planned_execution_date: null,
      portfolio_id: seededData.portfolio.id,
      type: 'ALLOCATION_PLAN',
    },
  ]);
  expect(snapshot.allocations).toHaveLength(10);
  expect(snapshot.allocations.every(({ cash_reserve, total_market_value }) => {
    return cash_reserve === false && total_market_value === null;
  })).toBe(true);
}

/** Asserts the application shell and its global portfolios navigation. */
async function expectRootShell(page: Page): Promise<void> {
  await expectRoute(page, '/');
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
}

/** Asserts the portfolio list contains the seeded portfolio and the new-portfolio card. */
async function expectPortfolioList(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expectRoute(page, '/portfolios');
  const portfolioCards = page.locator('#portfolios .portfolio-card');
  await expect(portfolioCards).toHaveCount(2);

  const existingPortfolioCard = portfolioCardFor(page, portfolio);
  await expect(existingPortfolioCard).toBeVisible();
  await expect(existingPortfolioCard).toHaveAttribute('data-navigate-to', `/portfolio/${portfolio.id}`);
  await expect(existingPortfolioCard).toHaveAttribute('navigate-to-bound', 'true');
  await expect(existingPortfolioCard.getByRole('heading', { level: 5, name: portfolio.name, exact: true }))
    .toBeVisible();

  const newPortfolioCard = portfolioCards.last();
  await expect(newPortfolioCard).toHaveAttribute('data-navigate-to', '/portfolios/new');
  await expect(newPortfolioCard).toHaveAttribute('navigate-to-bound', 'true');
  await expect(newPortfolioCard.getByRole('heading', { level: 5, name: 'New portfolio', exact: true })).toBeVisible();
}

/** Returns the portfolio card used to enter the selected portfolio. */
function portfolioCardFor(page: Page, portfolio: SeededPortfolio): Locator {
  return page.locator('#portfolios .portfolio-card').filter({ hasText: portfolio.name });
}

/** Asserts the selected portfolio context and its edit control. */
async function expectPortfolioContext(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expect(page.locator('#portfolio-context .badge.text-bg-secondary')).toHaveText(portfolio.name);
  await expect(page.locator('#portfolio-context button[data-navigate-to="/portfolio/:portfolioId/edit"]'))
    .toBeVisible();
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
}

/** Asserts the portfolio navigation labels, selection, and selected-button styling. */
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

/** Returns the visible label used to enter a portfolio submenu route. */
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

/** Asserts the allocation-map route, accordion order, and selected navigation state. */
async function expectAllocationMapShell(page: Page, seededData: SeededAllocationMapData): Promise<void> {
  await expectRoute(page, `/portfolio/${seededData.portfolio.id}/allocation-map`);
  await expectPortfolioContext(page, seededData.portfolio);
  await expectPortfolioNavigation(page, 'Allocation Map');

  const accordion = page.locator('#accordion-allocation-map');
  await expect(accordion).toBeVisible();
  const items = accordion.locator(':scope > .accordion-item');
  await expect(items).toHaveCount(2);

  const firstItem = items.nth(0);
  const secondItem = items.nth(1);
  await expect(firstItem.getByRole('button', { name: HISTORY_002_NAME, exact: true })).toBeVisible();
  await expect(secondItem.getByRole('button', { name: HISTORY_001_NAME, exact: true })).toBeVisible();
  await expect(firstItem.locator('.accordion-collapse')).toHaveClass(/show/);
  await expect(secondItem.locator('.accordion-collapse')).not.toHaveClass(/show/);
}

/** Asserts both independent plan selects, option order, and absent initial analysis tables. */
async function expectPlanControls(page: Page, seededData: SeededAllocationMapData): Promise<void> {
  const expectedOptions = [
    { text: 'Choose an allocation plan', value: '0' },
    { text: PLAN_002_NAME, value: seededData.plans.plan002.id.toString() },
    { text: PLAN_001_NAME, value: seededData.plans.plan001.id.toString() },
  ];

  for (const [index, observation] of [seededData.observations.history002, seededData.observations.history001].entries()) {
    const item = observationItem(page, observation);
    const select = item.locator(`select[name="${observation.id}PlanDivergence"]`);
    await expect(select).toBeAttached();
    if (index === 0) {
      await expect(select).toBeVisible();
    } else {
      await expect(select).not.toBeVisible();
    }
    await expect(select).toHaveValue('0');
    await expect(select.locator('option')).toHaveCount(expectedOptions.length);
    await expect(select.locator('option').evaluateAll((options) => options.map((option) => ({
      text: option.textContent?.trim(),
      value: option.getAttribute('value'),
    })))).resolves.toEqual(expectedOptions);
    await expect(item.locator('.bi-bar-chart-steps')).toHaveCount(1);
    await expect(item.locator('table')).toHaveCount(0);
  }
}

/** Returns one observation accordion body from its persisted observation identity. */
function observationItem(page: Page, observation: SeededObservation): Locator {
  return page.locator(`#time-tag-${observation.id}-allocation-map-container`);
}

/** Selects a plan, triggers the analysis request, and verifies the complete initial table. */
async function generateAnalysis(
  page: Page,
  portfolioId: number,
  observationItemElement: Locator,
  observation: SeededObservation,
  plan: SeededPlan,
  expectedRoots: readonly ExpectedDivergenceNode[],
): Promise<Locator> {
  const select = observationItemElement.locator(`select[name="${observation.id}PlanDivergence"]`);
  await select.selectOption(plan.id.toString());
  await expect(select).toHaveValue(plan.id.toString());

  const analysisPath = `/api/v2/portfolio/${portfolioId}/divergence/${observation.id}/allocation-plan/${plan.id}`;
  const responsePromise = page.waitForResponse((response) => {
    return response.request().method() === 'GET'
      && new URL(response.url()).pathname === analysisPath;
  });
  await observationItemElement.locator('button:has(.bi-bar-chart-steps)').click();
  const response = await responsePromise;
  expect(response.status()).toBe(200);

  const table = observationItemElement.locator('table');
  await expect(table).toHaveCount(1);
  await expectAnalysisTable(table, expectedRoots);
  return table;
}

/** Asserts table metadata, root ordering, root values, child ordering, and initial collapse state. */
async function expectAnalysisTable(table: Locator, expectedRoots: readonly ExpectedDivergenceNode[]): Promise<void> {
  await expect(table.locator('caption')).toHaveText('Total market value: $10,000.00');
  await expect(table.locator('thead th')).toHaveText([
    'Unit',
    'Total market value',
    'Planned market value',
    'Divergence',
  ]);
  await expect(table.locator('thead th').nth(3)).toHaveAttribute('colspan', '2');

  const allRows = table.locator('tbody > tr');
  const expectedChildCount = expectedRoots.reduce((count, root) => count + (root.children?.length ?? 0), 0);
  await expect(allRows).toHaveCount(expectedRoots.length + expectedChildCount);

  const rootRows = table.locator('tbody > tr:not(.collapse)');
  await expect(rootRows).toHaveCount(expectedRoots.length);
  await expect(table.locator('tbody > tr:not(.collapse):visible')).toHaveCount(expectedRoots.length);

  for (const [index, expectedRoot] of expectedRoots.entries()) {
    const rootRow = rootRows.nth(index);
    await expectNodeRow(rootRow, expectedRoot, true);
    const childRows = table.locator(`tbody > tr.collapse[data-divergence-parent="${expectedRoot.unit}"]`);
    await expect(childRows).toHaveCount(expectedRoot.children?.length ?? 0);
    for (let childIndex = 0; childIndex < await childRows.count(); childIndex++) {
      await expect(childRows.nth(childIndex)).not.toBeVisible();
    }
  }
}

/** Verifies a generated table remains in the collapsed or hidden observation body after navigation. */
async function expectRetainedAnalysisTable(
  table: Locator,
  expectedRoots: readonly ExpectedDivergenceNode[],
): Promise<void> {
  await expect(table.locator('caption')).toHaveText('Total market value: $10,000.00');
  const expectedChildCount = expectedRoots.reduce((count, root) => count + (root.children?.length ?? 0), 0);
  await expect(table.locator('tbody > tr')).toHaveCount(expectedRoots.length + expectedChildCount);
  for (const expectedRoot of expectedRoots) {
    await expect(table).toContainText(expectedRoot.unit);
    await expect(table).toContainText(expectedRoot.total);
    await expect(table).toContainText(expectedRoot.planned);
    await expect(table).toContainText(expectedRoot.divergence);
  }
}

/** Expands one root hierarchy row and verifies all child values and visual indicators. */
async function expandRootAndAssertChildren(
  table: Locator,
  expectedRoot: ExpectedDivergenceNode,
): Promise<void> {
  const rootRow = await findRootRow(table, expectedRoot.unit);
  const childRows = table.locator(`tbody > tr.collapse[data-divergence-parent="${expectedRoot.unit}"]`);
  const expectedChildren = expectedRoot.children ?? [];
  await expect(childRows).toHaveCount(expectedChildren.length);

  await rootRow.locator('span.badge[data-bs-toggle="collapse"]').click();
  await expect(childRows).toHaveCount(expectedChildren.length);
  for (let childIndex = 0; childIndex < await childRows.count(); childIndex++) {
    await expect(childRows.nth(childIndex)).toBeVisible();
  }

  for (const [index, expectedChild] of expectedChildren.entries()) {
    await expectNodeRow(childRows.nth(index), expectedChild, true);
  }
}

/** Finds a root row by its exact rendered hierarchy label. */
async function findRootRow(table: Locator, unit: string): Promise<Locator> {
  const rootRows = table.locator('tbody > tr:not(.collapse)');
  for (let index = 0; index < await rootRows.count(); index++) {
    const row = rootRows.nth(index);
    const label = row.locator('td').nth(0).locator('span').last();
    if ((await label.textContent())?.trim() === unit) {
      return row;
    }
  }
  throw new Error(`Could not find divergence root row ${unit}`);
}

/** Asserts one divergence row's values, icon, and configured progress-bar state. */
async function expectNodeRow(
  row: Locator,
  expected: ExpectedDivergenceNode,
  visible: boolean,
): Promise<void> {
  if (visible) {
    await expect(row).toBeVisible();
  } else {
    await expect(row).not.toBeVisible();
  }

  await expect(row.locator('td').nth(0).locator('span').last()).toHaveText(expected.unit);
  await expect(row.locator('td').nth(1)).toContainText(expected.total);
  await expect(row.locator('td').nth(2)).toContainText(expected.planned);
  await expect(row.locator('td').nth(3)).toContainText(expected.divergence);

  if (expected.children && expected.children.length > 0) {
    await expect(row.locator('span.badge[data-bs-toggle="collapse"] .bi-arrow-90deg-down')).toBeVisible();
  } else {
    await expect(row.locator('.bi-arrow-return-right')).toBeVisible();
  }
  await expectProgressBar(row, expected.barClass, expected.barWidth);
}

/** Verifies a striped divergence bar's style and clipping behavior. */
async function expectProgressBar(row: Locator, barClass: ExpectedDivergenceNode['barClass'], width: number): Promise<void> {
  const progress = row.locator('td').nth(4).locator('.progress');
  const bar = progress.locator('.progress-bar');
  await expect(progress).toHaveCount(1);
  await expect(progress).toHaveCSS('overflow', 'hidden');
  await expect(bar).toHaveClass(new RegExp(`\\bprogress-bar-striped\\b`));
  await expect(bar).toHaveClass(new RegExp(`\\b${barClass}\\b`));
  await expect(bar).toHaveAttribute('style', new RegExp(`width:\\s*${width}%`));
}

/** Builds the convergent History 002 versus Plan 002 roots. */
function history002Plan002Roots(): readonly ExpectedDivergenceNode[] {
  return [
    node('CORE', '$4,000.00 (40%)', '$4,000.00 (40%)', '$0.00 (0%)', 'bg-success', 0, [
      node('E2E:CORE-C', '$3,000.00 (75%)', '$3,000.00 (75%)', '$0.00 (0%)', 'bg-success', 0),
      node('E2E:CORE-A', '$1,000.00 (25%)', '$1,000.00 (25%)', '$0.00 (0%)', 'bg-success', 0),
    ]),
    node('NEW', '$6,000.00 (60%)', '$6,000.00 (60%)', '$0.00 (0%)', 'bg-success', 0, [
      node('E2E:NEW-A', '$6,000.00 (100%)', '$6,000.00 (100%)', '$0.00 (0%)', 'bg-success', 0),
    ]),
  ];
}

/** Builds the divergent History 002 versus Plan 001 roots. */
function history002Plan001Roots(): readonly ExpectedDivergenceNode[] {
  return [
    node('CORE', '$4,000.00 (40%)', '$6,000.00 (60%)', '-$2,000.00 (-20%)', 'bg-success', 40, [
      node('E2E:CORE-C', '$3,000.00 (75%)', '$0.00 (0%)', '$3,000.00 (75%)', 'bg-danger', 150),
      node('E2E:CORE-A', '$1,000.00 (25%)', '$1,600.00 (40%)', '-$600.00 (-15%)', 'bg-success', 30),
      node('E2E:CORE-B', '$0.00 (0%)', '$2,400.00 (60%)', '-$2,400.00 (-60%)', 'bg-success', 120),
    ]),
    node('NEW', '$6,000.00 (60%)', '$0.00 (0%)', '$6,000.00 (60%)', 'bg-danger', 120, [
      node('E2E:NEW-A', '$6,000.00 (100%)', '$0.00 (0%)', '$6,000.00 (100%)', 'bg-danger', 200),
    ]),
    node('LEGACY', '$0.00 (0%)', '$4,000.00 (40%)', '-$4,000.00 (-40%)', 'bg-success', 80, [
      node('E2E:LEGACY-A', '$0.00', '$0.00', '$0.00', 'bg-success', 0),
    ]),
  ];
}

/** Builds the convergent History 001 versus Plan 001 roots. */
function history001Plan001Roots(): readonly ExpectedDivergenceNode[] {
  return [
    node('CORE', '$6,000.00 (60%)', '$6,000.00 (60%)', '$0.00 (0%)', 'bg-success', 0, [
      node('E2E:CORE-B', '$3,600.00 (60%)', '$3,600.00 (60%)', '$0.00 (0%)', 'bg-success', 0),
      node('E2E:CORE-A', '$2,400.00 (40%)', '$2,400.00 (40%)', '$0.00 (0%)', 'bg-success', 0),
    ]),
    node('LEGACY', '$4,000.00 (40%)', '$4,000.00 (40%)', '$0.00 (0%)', 'bg-success', 0, [
      node('E2E:LEGACY-A', '$4,000.00 (100%)', '$4,000.00 (100%)', '$0.00 (0%)', 'bg-success', 0),
    ]),
  ];
}

/** Builds the divergent History 001 versus Plan 002 roots. */
function history001Plan002Roots(): readonly ExpectedDivergenceNode[] {
  return [
    node('CORE', '$6,000.00 (60%)', '$4,000.00 (40%)', '$2,000.00 (20%)', 'bg-danger', 40, [
      node('E2E:CORE-B', '$3,600.00 (60%)', '$0.00 (0%)', '$3,600.00 (60%)', 'bg-danger', 120),
      node('E2E:CORE-A', '$2,400.00 (40%)', '$1,500.00 (25%)', '$900.00 (15%)', 'bg-danger', 30),
      node('E2E:CORE-C', '$0.00 (0%)', '$4,500.00 (75%)', '-$4,500.00 (-75%)', 'bg-success', 150),
    ]),
    node('LEGACY', '$4,000.00 (40%)', '$0.00 (0%)', '$4,000.00 (40%)', 'bg-danger', 80, [
      node('E2E:LEGACY-A', '$4,000.00 (100%)', '$0.00 (0%)', '$4,000.00 (100%)', 'bg-danger', 200),
    ]),
    node('NEW', '$0.00 (0%)', '$6,000.00 (60%)', '-$6,000.00 (-60%)', 'bg-success', 120, [
      node('E2E:NEW-A', '$0.00', '$0.00', '$0.00', 'bg-success', 0),
    ]),
  ];
}

/** Creates one expected divergence node while keeping scenario data readable. */
function node(
  unit: string,
  total: string,
  planned: string,
  divergence: string,
  barClass: ExpectedDivergenceNode['barClass'],
  barWidth: number,
  children?: readonly ExpectedDivergenceNode[],
): ExpectedDivergenceNode {
  return { barClass, barWidth, children, divergence, planned, total, unit };
}

/** Waits until the browser location has reached the expected application path. */
async function expectRoute(page: Page, path: string): Promise<void> {
  await expect(page).toHaveURL(routePattern(path));
}

/** Builds a URL matcher that permits an optional trailing slash and no other path. */
function routePattern(path: string): RegExp {
  const escapedPath = path.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  return new RegExp(`${escapedPath}/?$`);
}
