/**
 * Covers portfolio creation through the browser, API, and PostgreSQL boundaries.
 *
 * Run through the containerized E2E runner after its Compose environment supplies
 * BASE_URL and PostgreSQL administrator PG* variables.
 *
 * Authored by: OpenCode
 */
import type { APIResponse, Response } from '@playwright/test';
import { expect, test } from '../support/fixtures';

const DEFAULT_ALLOCATION_STRUCTURE = {
  hierarchy: [
    { name: 'Assets', field: 'assetTicker' },
    { name: 'Classes', field: 'class' },
  ],
};

/** Represents the API response for a portfolio created by the browser flow. */
type Portfolio = {
  allocationStructure: typeof DEFAULT_ALLOCATION_STRUCTURE;
  id: number;
  name: string;
};

/** Represents the persisted portfolio state queried directly from PostgreSQL. */
type PortfolioRow = {
  allocation_structure: typeof DEFAULT_ALLOCATION_STRUCTURE;
  id: number;
  name: string;
};

test('creates a portfolio consistently through the browser, API, and database', async ({ database, page, request }) => {
  const portfolioName = `E2E Portfolio ${Date.now()}`;
  const initialLoad = page.waitForResponse(isPortfolioCollectionRequest);

  await page.goto('/portfolios');

  const initialResponse = await initialLoad;
  expect(initialResponse.status()).toBe(200);
  expect(await initialResponse.json()).toEqual([]);

  const newPortfolioCard = page.locator('[data-navigate-to="/portfolios/new"]');
  await expect(newPortfolioCard).toHaveAttribute('navigate-to-bound', 'true');
  await Promise.all([
    page.waitForURL('**/portfolios/new'),
    page.getByRole('heading', { exact: true, level: 5, name: 'New portfolio' }).click(),
  ]);

  await page.getByRole('textbox', { name: 'Name' }).fill(portfolioName);

  const createResponse = page.waitForResponse(isPortfolioCreationRequest);
  const reloadedPortfolios = page.waitForResponse(isPortfolioCollectionRequest);
  await Promise.all([
    page.waitForURL('**/portfolios'),
    page.getByRole('button', { name: 'Create' }).click(),
  ]);

  const postResponse = await createResponse;
  expect(postResponse.status()).toBe(201);
  expect(postResponse.request().headers()['content-type']).toContain('application/json');
  expect(postResponse.request().postDataJSON()).toEqual({ name: portfolioName });

  const createdPortfolio = await postResponse.json() as Portfolio;
  expect(createdPortfolio).toMatchObject({
    allocationStructure: DEFAULT_ALLOCATION_STRUCTURE,
    id: expect.any(Number),
    name: portfolioName,
  });
  expect(createdPortfolio.id).toBeGreaterThan(0);

  const reloadedResponse = await reloadedPortfolios;
  expect(reloadedResponse.status()).toBe(200);
  expect(await reloadedResponse.json()).toEqual([createdPortfolio]);
  await expect(page.getByRole('heading', { exact: true, level: 5, name: portfolioName })).toBeVisible();

  const databaseRows = await database.query<PortfolioRow>(
    `SELECT id, name, allocation_structure
     FROM public.portfolio
     WHERE name = $1`,
    [portfolioName],
  );
  expect(databaseRows).toEqual([{
    allocation_structure: DEFAULT_ALLOCATION_STRUCTURE,
    id: createdPortfolio.id,
    name: portfolioName,
  }]);

  const apiResponse = await request.get(`/api/portfolio/${createdPortfolio.id}`);
  expect(apiResponse.status()).toBe(200);
  expect(await apiResponse.json()).toEqual(createdPortfolio);
});

/** Matches the same-origin collection request made by the portfolios HTMX component. */
function isPortfolioCollectionRequest(response: Response): boolean {
  return response.request().method() === 'GET' && new URL(response.url()).pathname === '/api/portfolio';
}

/** Matches the same-origin browser request that creates one portfolio. */
function isPortfolioCreationRequest(response: Response): boolean {
  return response.request().method() === 'POST' && new URL(response.url()).pathname === '/api/portfolio';
}
