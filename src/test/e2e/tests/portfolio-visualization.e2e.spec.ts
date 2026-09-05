/**
 * Covers empty portfolio visualization through click navigation and direct URL navigation.
 *
 * The scenarios seed their portfolios directly in PostgreSQL and validate the rendered
 * application state without creating or changing domain data through the application API.
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

const EMPTY_PORTFOLIO_NAMES = [
  'E2E Empty Portfolio Alpha',
  'E2E Empty Portfolio Beta',
] as const;

const PORTFOLIO_NAVIGATION_OPTIONS = [
  { id: 'nav-radio-portfolio', name: 'Portfolio' },
  { id: 'nav-radio-allocation-plan', name: 'Allocation Plan' },
  { id: 'nav-radio-allocation-map', name: 'Allocation Map' },
] as const;

type SeededPortfolio = {
  id: number;
  name: string;
};

type NavigationOptionName = typeof PORTFOLIO_NAVIGATION_OPTIONS[number]['name'];

test.describe('empty portfolio visualization', () => {
  test('scenario 1: navigates through empty portfolio views with clicks', async ({ database, page }) => {
    await runEmptyPortfolioVisualizationScenario(page, database, false);
  });

  test('scenario 1.1: navigates through empty portfolio views with direct URLs', async ({ database, page }) => {
    await runEmptyPortfolioVisualizationScenario(page, database, true);
  });
});

/** Runs scenario 1 or scenario 1.1, depending on whether direct navigation is enabled. */
async function runEmptyPortfolioVisualizationScenario(
  page: Page,
  database: E2eDatabase,
  useDirectNavigation: boolean,
): Promise<void> {
  const seededPortfolios = await seedEmptyPortfolios(database);

  await page.goto('/');
  await expectRootShell(page);

  if (useDirectNavigation) {
    await page.goto('/portfolios');
  } else {
    await navigateByClick(page, '/portfolios', page.getByRole('link', { name: 'Portfolios', exact: true }));
  }

  const selectedPortfolio = await expectPortfolioList(page, seededPortfolios);

  if (useDirectNavigation) {
    await page.goto(`/portfolio/${selectedPortfolio.id}`);
  } else {
    const selectedPortfolioCard = page
      .locator('#portfolios .portfolio-card')
      .filter({ hasText: selectedPortfolio.name });
    await navigateByClick(page, `/portfolio/${selectedPortfolio.id}`, selectedPortfolioCard);
  }

  await expectPortfolioShell(page, selectedPortfolio, undefined);

  if (useDirectNavigation) {
    await page.goto(`/portfolio/${selectedPortfolio.id}/history`);
  } else {
    await navigateByClick(page, `/portfolio/${selectedPortfolio.id}/history`, portfolioNavigationLabel(page, 'Portfolio'));
  }

  await expectPortfolioHistory(page, selectedPortfolio);

  if (useDirectNavigation) {
    await page.goto(`/portfolio/${selectedPortfolio.id}/allocation`);
  } else {
    await navigateByClick(
      page,
      `/portfolio/${selectedPortfolio.id}/allocation`,
      portfolioNavigationLabel(page, 'Allocation Plan'),
    );
  }

  await expectAllocationPlan(page, selectedPortfolio);

  if (useDirectNavigation) {
    await page.goto(`/portfolio/${selectedPortfolio.id}/allocation-map`);
  } else {
    await navigateByClick(
      page,
      `/portfolio/${selectedPortfolio.id}/allocation-map`,
      portfolioNavigationLabel(page, 'Allocation Map'),
    );
  }

  await expectAllocationMap(page, selectedPortfolio);

  if (useDirectNavigation) {
    await page.goto('/portfolios');
  } else {
    await navigateByClick(page, '/portfolios', page.getByRole('link', { name: 'Portfolios', exact: true }));
  }

  await expectPortfolioList(page, seededPortfolios);
}

/** Seeds exactly two empty portfolios with the allocation hierarchy required by the views. */
async function seedEmptyPortfolios(
  database: E2eDatabase,
): Promise<readonly SeededPortfolio[]> {
  const rows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $3::jsonb), ($2, $3::jsonb)
     RETURNING id, name`,
    [EMPTY_PORTFOLIO_NAMES[0], EMPTY_PORTFOLIO_NAMES[1], JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );

  expect(rows).toHaveLength(EMPTY_PORTFOLIO_NAMES.length);
  expect(rows.map(({ name }) => name).sort()).toEqual([...EMPTY_PORTFOLIO_NAMES].sort());

  return rows;
}

/** Asserts that the application shell exposes the global portfolios navigation. */
async function expectRootShell(page: Page): Promise<void> {
  await expectRoute(page, '/');
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
}

/** Asserts the portfolio cards and returns the first rendered portfolio for subsequent navigation. */
async function expectPortfolioList(
  page: Page,
  seededPortfolios: readonly SeededPortfolio[],
): Promise<SeededPortfolio> {
  await expectRoute(page, '/portfolios');

  const portfolios = page.locator('#portfolios');
  const portfolioCards = portfolios.locator('.portfolio-card');
  await expect(portfolioCards).toHaveCount(seededPortfolios.length + 1);

  for (const seededPortfolio of seededPortfolios) {
    const portfolioCard = portfolioCards.filter({ hasText: seededPortfolio.name });
    await expect(portfolioCard).toHaveCount(1);
    await expect(portfolioCard).toBeVisible();
    await expect(portfolioCard).toHaveAttribute('data-navigate-to', `/portfolio/${seededPortfolio.id}`);
    await expect(portfolioCard).toHaveAttribute('navigate-to-bound', 'true');
    await expect(
      portfolioCard.getByRole('heading', { level: 5, name: seededPortfolio.name, exact: true }),
    ).toBeVisible();
  }

  const newPortfolioCard = portfolioCards.last();
  await expect(newPortfolioCard).toHaveAttribute('data-navigate-to', '/portfolios/new');
  await expect(newPortfolioCard).toHaveAttribute('navigate-to-bound', 'true');
  await expect(
    newPortfolioCard.getByRole('heading', { level: 5, name: 'New portfolio', exact: true }),
  ).toBeVisible();

  const firstPortfolioCard = portfolioCards.first();
  const firstPortfolioName = (await firstPortfolioCard.getByRole('heading', { level: 5 }).innerText()).trim();
  const selectedPortfolio = seededPortfolios.find(({ name }) => name === firstPortfolioName);

  if (!selectedPortfolio) {
    throw new Error(`The first portfolio card is not one of the seeded portfolios: ${firstPortfolioName}`);
  }

  return selectedPortfolio;
}

/** Asserts the selected portfolio header, edit control, submenu, and global navigation. */
async function expectPortfolioShell(
  page: Page,
  portfolio: SeededPortfolio,
  selectedNavigation: NavigationOptionName | undefined,
): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}` + (selectedNavigation ? getRouteSuffix(selectedNavigation) : ''));
  await expect(page.locator('#portfolio-context .badge.text-bg-secondary')).toHaveText(portfolio.name);

  const editButton = page.locator('#portfolio-context button[data-navigate-to="/portfolio/:portfolioId/edit"]');
  await expect(editButton).toBeVisible();
  await expect(editButton.locator('.bi-pen')).toBeVisible();

  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
  await expectPortfolioNavigation(page, selectedNavigation);
}

/** Asserts the empty history view and its portfolio-history management control. */
async function expectPortfolioHistory(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expectPortfolioShell(page, portfolio, 'Portfolio');

  const managementButton = page.locator(
    '#accordion-portfolio-history > button[onclick="navigateToPortfolioAllocationManagement()"]',
  );
  await expect(managementButton).toBeVisible();
  await expect(managementButton.locator('.bi-database-gear')).toBeVisible();
  await expect(page.locator('#accordion-portfolio-history > *')).toHaveCount(1);
}

/** Asserts the empty allocation-plan view and its management control. */
async function expectAllocationPlan(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expectPortfolioShell(page, portfolio, 'Allocation Plan');

  const managementButton = page.locator(
    '#accordion-allocation-plan > button[onclick="navigateToPortfolioAllocationPlanManagement()"]',
  );
  await expect(managementButton).toBeVisible();
  await expect(managementButton.locator('.bi-database-gear')).toBeVisible();
  await expect(page.locator('#accordion-allocation-plan > *')).toHaveCount(1);
}

/** Asserts the empty allocation-map partial and its selected submenu state. */
async function expectAllocationMap(page: Page, portfolio: SeededPortfolio): Promise<void> {
  await expectPortfolioShell(page, portfolio, 'Allocation Map');
  await expect(page.locator('#accordion-allocation-map')).toBeAttached();
  await expect(page.locator('#accordion-allocation-map > *')).toHaveCount(0);
}

/** Asserts submenu labels, radio selection, and outlined or filled button styling. */
async function expectPortfolioNavigation(
  page: Page,
  selectedNavigation: NavigationOptionName | undefined,
): Promise<void> {
  const navigation = page.getByRole('group', { name: 'Navigation' });
  await expect(navigation).toBeVisible();

  for (const option of PORTFOLIO_NAVIGATION_OPTIONS) {
    const radio = navigation.getByRole('radio', { name: option.name, exact: true });
    const label = navigation.locator(`label[for="${option.id}"]`);

    await expect(label).toHaveText(option.name);
    await expect(label).toHaveClass(/btn-outline-primary/);

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

/** Maps a submenu option to its selected portfolio route suffix. */
function getRouteSuffix(optionName: NavigationOptionName): string {
  const suffixes: Record<NavigationOptionName, string> = {
    Portfolio: '/history',
    'Allocation Plan': '/allocation',
    'Allocation Map': '/allocation-map',
  };

  return suffixes[optionName];
}
