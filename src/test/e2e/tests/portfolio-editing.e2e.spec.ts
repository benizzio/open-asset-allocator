/**
 * Covers portfolio editing through click navigation and direct URL navigation.
 *
 * Scenario 2 verifies the Save and Cancel flows for one portfolio and checks the
 * resulting state directly in PostgreSQL. Scenario 2.1 verifies the equivalent
 * read-only navigation using direct browser URLs.
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

const ORIGINAL_PORTFOLIO_NAME = 'E2E Original Portfolio';
const UPDATED_PORTFOLIO_NAME = 'E2E Updated Portfolio';
const CANCELLED_PORTFOLIO_NAME = 'E2E Cancelled Portfolio';

type SeededPortfolio = {
  id: number;
  name: string;
};

type PortfolioRow = SeededPortfolio & {
  allocation_structure: typeof DEFAULT_ALLOCATION_STRUCTURE;
};

type ExpectedPortfolio = SeededPortfolio;

type NavigationOptionName = typeof PORTFOLIO_NAVIGATION_OPTIONS[number]['name'];

const PORTFOLIO_NAVIGATION_OPTIONS = [
  { id: 'nav-radio-portfolio', name: 'Portfolio' },
  { id: 'nav-radio-allocation-plan', name: 'Allocation Plan' },
  { id: 'nav-radio-allocation-map', name: 'Allocation Map' },
] as const;

test.describe('portfolio editing', () => {
  test('scenario 2: edits a portfolio with clicks and verifies Save and Cancel', async ({ database, page }) => {
    const portfolio = await seedPortfolio(database, ORIGINAL_PORTFOLIO_NAME);

    await page.goto('/');
    await expectRootShell(page);

    await navigateByClick(page, '/portfolios', page.getByRole('link', { name: 'Portfolios', exact: true }));
    await expectPortfolioList(page, portfolio);

    const portfolioCard = portfolioCardFor(page, portfolio);
    await navigateByClick(page, `/portfolio/${portfolio.id}`, portfolioCard);
    await expectPortfolioShell(page, portfolio, undefined);

    const editButton = page.locator('#portfolio-context button[data-navigate-to="/portfolio/:portfolioId/edit"]');
    await navigateByClick(page, `/portfolio/${portfolio.id}/edit`, editButton);
    await expectEditPortfolio(page, portfolio);

    await page.getByRole('textbox', { name: 'Name' }).fill(UPDATED_PORTFOLIO_NAME);
    await navigateByClick(page, `/portfolio/${portfolio.id}`, page.getByRole('button', { name: 'Save' }));

    const updatedPortfolio: ExpectedPortfolio = { id: portfolio.id, name: UPDATED_PORTFOLIO_NAME };
    await expectPortfolioShell(page, updatedPortfolio, undefined);
    await expectPersistedPortfolio(database, updatedPortfolio);

    await navigateByClick(page, '/portfolios', page.getByRole('link', { name: 'Portfolios', exact: true }));
    await expectPortfolioList(page, updatedPortfolio);

    await navigateByClick(page, `/portfolio/${updatedPortfolio.id}`, portfolioCardFor(page, updatedPortfolio));
    await expectPortfolioShell(page, updatedPortfolio, undefined);

    await navigateByClick(
      page,
      `/portfolio/${updatedPortfolio.id}/edit`,
      page.locator('#portfolio-context button[data-navigate-to="/portfolio/:portfolioId/edit"]'),
    );
    await expectEditPortfolio(page, updatedPortfolio);

    await page.getByRole('textbox', { name: 'Name' }).fill(CANCELLED_PORTFOLIO_NAME);
    await navigateByClick(page, `/portfolio/${updatedPortfolio.id}`, page.getByRole('button', { name: 'Cancel' }));

    await expectPortfolioShell(page, updatedPortfolio, undefined);
    await expectPersistedPortfolio(database, updatedPortfolio);
  });

  test('scenario 2.1: navigates to portfolio editing pages with direct URLs', async ({ database, page }) => {
    const portfolio = await seedPortfolio(database, ORIGINAL_PORTFOLIO_NAME);

    await page.goto('/');
    await expectRootShell(page);

    await page.goto('/portfolios');
    await expectPortfolioList(page, portfolio);

    await page.goto(`/portfolio/${portfolio.id}`);
    await expectPortfolioShell(page, portfolio, undefined);

    await page.goto(`/portfolio/${portfolio.id}/edit`);
    await expectEditPortfolio(page, portfolio);
    await expectPersistedPortfolio(database, portfolio);
  });
});

/** Seeds one empty portfolio with the allocation structure expected by the page. */
async function seedPortfolio(database: E2eDatabase, name: string): Promise<SeededPortfolio> {
  const rows = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [name, JSON.stringify(DEFAULT_ALLOCATION_STRUCTURE)],
  );

  expect(rows).toHaveLength(1);
  expect(rows[0]).toEqual(expect.objectContaining({ id: expect.any(Number), name }));

  return rows[0];
}

/** Asserts the global shell and its portfolios navigation link. */
async function expectRootShell(page: Page): Promise<void> {
  await expectRoute(page, '/');
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
}

/** Asserts the portfolio list contains the seeded portfolio and New portfolio card. */
async function expectPortfolioList(page: Page, portfolio: ExpectedPortfolio): Promise<void> {
  await expectRoute(page, '/portfolios');

  const portfolios = page.locator('#portfolios');
  const portfolioCards = portfolios.locator('.portfolio-card');
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
function portfolioCardFor(page: Page, portfolio: ExpectedPortfolio): Locator {
  return page.locator('#portfolios .portfolio-card').filter({ hasText: portfolio.name });
}

/** Asserts the selected portfolio header, edit control, submenu, and global navigation. */
async function expectPortfolioShell(
  page: Page,
  portfolio: ExpectedPortfolio,
  selectedNavigation: NavigationOptionName | undefined,
): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}`);
  await expect(page.locator('#portfolio-context .badge.text-bg-secondary')).toHaveText(portfolio.name);

  const editButton = page.locator('#portfolio-context button[data-navigate-to="/portfolio/:portfolioId/edit"]');
  await expect(editButton).toBeVisible();
  await expect(editButton.locator('.bi-pen')).toBeVisible();

  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
  await expectPortfolioNavigation(page, selectedNavigation);
}

/** Asserts the edit page form, pre-filled name, controls, and hidden submenu. */
async function expectEditPortfolio(page: Page, portfolio: ExpectedPortfolio): Promise<void> {
  await expectRoute(page, `/portfolio/${portfolio.id}/edit`);
  await expect(page.locator('#portfolio-context .badge.text-bg-secondary')).toHaveText(portfolio.name);

  const editButton = page.locator('#portfolio-context button[data-navigate-to="/portfolio/:portfolioId/edit"]');
  await expect(editButton).toBeVisible();
  await expect(editButton.locator('.bi-pen')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Portfolios', exact: true })).toBeVisible();
  await expect(page.getByRole('group', { name: 'Navigation' })).not.toBeVisible();

  const editCard = page.locator('.card:has(#edit-portfolio-form)');
  await expect(editCard.getByText('Edit portfolio', { exact: true })).toBeVisible();
  await expect(page.getByRole('textbox', { name: 'Name' })).toHaveValue(portfolio.name);
  await expect(editCard.getByRole('button', { name: 'Cancel' })).toHaveClass(/btn-secondary/);
  await expect(editCard.getByRole('button', { name: 'Save' })).toHaveClass(/btn-primary/);
}

/** Asserts that the portfolio submenu options have the expected selected state. */
async function expectPortfolioNavigation(
  page: Page,
  selectedNavigation: NavigationOptionName | undefined,
): Promise<void> {
  const navigation = page.getByRole('group', { name: 'Navigation' });
  await expect(navigation).toBeVisible();

  for (const option of PORTFOLIO_NAVIGATION_OPTIONS) {
    const radio = navigation.getByRole('radio', { name: option.name, exact: true });
    const label = navigation.locator(`label[for="${option.id}"]`);

    await expect(radio).toBeAttached();
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

/** Verifies the portfolio identity and unchanged allocation structure in PostgreSQL. */
async function expectPersistedPortfolio(database: E2eDatabase, portfolio: ExpectedPortfolio): Promise<void> {
  const rows = await database.query<PortfolioRow>(
    `SELECT id, name, allocation_structure
     FROM public.portfolio
     WHERE id = $1`,
    [portfolio.id],
  );

  expect(rows).toEqual([{
    allocation_structure: DEFAULT_ALLOCATION_STRUCTURE,
    id: portfolio.id,
    name: portfolio.name,
  }]);
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
