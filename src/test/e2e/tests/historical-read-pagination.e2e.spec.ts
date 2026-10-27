/**
 * Verifies cursor pagination, DOM preservation, retry behavior, and terminal controls for history and allocation plans.
 *
 * Scenarios 21.1–21.4 cover the four management and viewing lists using database-seeded records.
 *
 * @author GPT-6 Luna
 */
import type { Locator, Page, Response } from '@playwright/test';
import { expect, test } from '../support/fixtures';
import type { E2eDatabase } from '../support/database';

const ALLOCATION_STRUCTURE = {
  hierarchy: [
    { name: 'Assets', field: 'assetTicker' },
    { name: 'Classes', field: 'class' },
  ],
} as const;

const PAGINATION_RECORD_COUNT = 21;
const PORTFOLIO_NAME = 'E2E Historical Read Pagination';

type SeededPortfolio = {
  id: number;
  name: string;
};

type SeededHistoryRecord = {
  id: number;
  timeTag: string;
};

type SeededPlanRecord = {
  id: number;
  name: string;
};

type PaginatedResponse = {
  items: unknown[];
  nextCursor: string | null;
};

type SeededHistory = {
  portfolio: SeededPortfolio;
  records: readonly SeededHistoryRecord[];
};

type SeededPlans = {
  portfolio: SeededPortfolio;
  records: readonly SeededPlanRecord[];
};

type LateViewResponseIsolationConfig = {
  viewingRoute: string;
  managementRoute: string;
  viewingEndpoint: string;
  managementEndpoint: string;
  viewingOwnerId: string;
  managementOwnerId: string;
  viewingManagementButton: string;
  managementViewingButton: string;
  managementDraftSelector: string;
  draftValue: string;
};

test('scenario 21.1: paginates portfolio history viewing with retry and preserves existing charts', async ({
  database,
  page,
}) => {
  const seeded = await seedHistory(database);
  const endpoint = `/api/portfolio/${seeded.portfolio.id}/history`;
  const initialResponsePromise = waitForPageResponse(page, endpoint, '');

  await page.goto(`/portfolio/${seeded.portfolio.id}/history`);

  const initialResponse = await initialResponsePromise;
  const initialPage = await readPage(initialResponse, endpoint, '');
  const firstCursor = requireContinuationCursor(initialPage);
  const list = page.locator('#accordion-portfolio-history');
  const moreButton = loadMoreButton(page, 'accordion-portfolio-history');
  const status = page.locator('[data-paginated-list-status="accordion-portfolio-history"]');
  const displayedRecords = list.locator('[data-paginated-record-id]');
  const expectedIds = seeded.records.map(({ id }) => String(id));

  await expect(displayedRecords).toHaveCount(10);
  await expect.poll(() => displayedIds(displayedRecords)).toEqual(expectedIds.slice(0, 10));
  await expect(list.locator('[onclick="navigateToPortfolioAllocationManagement()"]')).toHaveCount(1);
  await expect(moreButton).toBeVisible();
  await expectReadyControl(page, 'accordion-portfolio-history');
  await expectPaginationGapAtWidths(page, 'accordion-portfolio-history');

  const firstRecordId = expectedIds[0];
  const firstPanel = page.locator(`#portfolio-allocation-${firstRecordId}`);
  await expect(firstPanel).toHaveClass(/\bshow\b/);
  const originalCanvas = await firstPanel.locator('canvas').elementHandle();
  expect(originalCanvas).not.toBeNull();

  let firstCursorFailure = true;
  let holdRetryResponse = false;
  let continuationRequestCount = 0;
  let releaseRetryResponse: () => void = () => {};
  let signalRetryStarted: () => void = () => {};
  const retryResponseGate = new Promise<void>(resolve => {
    releaseRetryResponse = resolve;
  });
  const retryRequestStarted = new Promise<void>(resolve => {
    signalRetryStarted = resolve;
  });

  await page.route(`**${endpoint}?*`, async route => {
    const requestUrl = new URL(route.request().url());

    if(requestUrl.searchParams.get('pageCursor') === firstCursor) {
      continuationRequestCount++;
    }

    if(requestUrl.searchParams.get('pageCursor') === firstCursor && firstCursorFailure) {
      firstCursorFailure = false;
      await route.abort();
      return;
    }

    if(requestUrl.searchParams.get('pageCursor') === firstCursor && holdRetryResponse) {
      holdRetryResponse = false;
      signalRetryStarted();
      await retryResponseGate;
      await route.continue();
      return;
    }

    await route.continue();
  });

  let retryResponse: Response;
  try {
    await moreButton.click();
    await expect(status).toHaveAttribute('role', 'alert');
    await expect(status).toHaveText('Could not load records. Try again.');
    await expect(displayedRecords).toHaveCount(10);
    await expect(moreButton).toHaveAccessibleName('Retry loading records');
    await expect(moreButton).toHaveAttribute('title', 'Retry loading records');
    await expect(moreButton.locator('[data-paginated-list-retry-icon]')).toBeVisible();
    await expect(moreButton.locator('[data-paginated-list-more-icon]')).toBeHidden();

    holdRetryResponse = true;
    const retryResponsePromise = waitForPageResponse(page, endpoint, firstCursor);
    await moreButton.click();
    await retryRequestStarted;
    await expect(moreButton).toBeDisabled();
    await expect(moreButton).toHaveAccessibleName('Loading records');
    await expect(moreButton).toHaveAttribute('aria-busy', 'true');
    await expect(moreButton.locator('[data-paginated-list-loading-indicator]')).toBeVisible();
    await expect(moreButton.locator('[data-paginated-list-retry-icon]')).toBeHidden();
    await expect(status).toBeEmpty();

    await moreButton.evaluate(button => (button as HTMLButtonElement).click());
    expect(continuationRequestCount).toBe(2);

    releaseRetryResponse();
    retryResponse = await retryResponsePromise;
  } finally {
    releaseRetryResponse();
  }

  const secondPage = await readPage(retryResponse, endpoint, firstCursor);
  const secondCursor = requireContinuationCursor(secondPage);

  await expect(displayedRecords).toHaveCount(20);
  await expect(status).toBeEmpty();
  await expect.poll(() => displayedIds(displayedRecords)).toEqual(expectedIds.slice(0, 20));
  await expectReadyControl(page, 'accordion-portfolio-history');
  await expect(firstPanel).toHaveClass(/\bshow\b/);
  await expect(firstPanel.locator('canvas')).toHaveCount(1);
  expect(await firstPanel.locator('canvas').evaluate((canvas, original) => canvas === original, originalCanvas))
    .toBe(true);
  await expect(list.locator('[onclick="navigateToPortfolioAllocationManagement()"]')).toHaveCount(1);

  const finalResponsePromise = waitForPageResponse(page, endpoint, secondCursor);
  await moreButton.click();
  const finalResponse = await finalResponsePromise;
  await readPage(finalResponse, endpoint, secondCursor);

  await expect(displayedRecords).toHaveCount(PAGINATION_RECORD_COUNT);
  await expect.poll(() => displayedIds(displayedRecords)).toEqual(expectedIds);
  await expect(moreButton).toBeHidden();
  await expect(page.locator('[data-paginated-list-controls="accordion-portfolio-history"]')).toBeHidden();
  await expect(status).toBeEmpty();
  await expect(list.locator('[onclick="navigateToPortfolioAllocationManagement()"]')).toHaveCount(1);
});

test('scenario 21.2: appends history-management records without replacing drafts or editor state', async ({
  database,
  page,
}) => {
  const seeded = await seedHistory(database);
  const endpoint = `/api/portfolio/${seeded.portfolio.id}/history/observation`;
  const initialResponsePromise = waitForPageResponse(page, endpoint, '');

  await page.goto(`/portfolio/${seeded.portfolio.id}/history/manage`);

  const initialResponse = await initialResponsePromise;
  const initialPage = await readPage(initialResponse, endpoint, '');
  const firstCursor = requireContinuationCursor(initialPage);
  const list = page.locator('#accordion-portfolio-history-management');
  const moreButton = loadMoreButton(page, 'accordion-portfolio-history-management');
  const displayedRecords = list.locator('[data-paginated-record-id]');
  const expectedIds = seeded.records.map(({ id }) => String(id));

  await expect(displayedRecords).toHaveCount(11);
  await expect.poll(() => displayedIds(displayedRecords).then(ids => ids.filter(id => id !== '0')))
    .toEqual(expectedIds.slice(0, 10));
  await expect(list.locator(':scope > [data-paginated-record-id="0"]')).toHaveCount(1);
  await expect(list.locator(':scope > [data-paginated-record-id]').first())
    .toHaveAttribute('data-paginated-record-id', '0');
  await expect(page.locator('[data-paginated-list-controls="accordion-portfolio-history"]')).toBeHidden();
  await expect(loadMoreButton(page, 'accordion-portfolio-history')).toBeHidden();
  await expectReadyControl(page, 'accordion-portfolio-history-management');
  await expectPaginationGapAtWidths(page, 'accordion-portfolio-history-management');

  const draftTag = list.locator('.portfolio-history-management-new-time-tag');
  const draftDate = list.locator('#portfolio-history-management-observation-date');
  await draftTag.fill('Unsaved history draft');
  await draftDate.fill('2025-03-04');

  const firstRecord = page.locator(`#portfolio-history-management-container-${expectedIds[0]}`);
  const firstRecordHandle = await firstRecord.elementHandle();
  const detailResponsePromise = page.waitForResponse(response => {
    const url = new URL(response.url());
    return url.pathname === `/api/portfolio/${seeded.portfolio.id}/history`
      && url.searchParams.get('observationTimestampId') === expectedIds[0];
  });

  await firstRecord.locator(`#portfolio-history-management-trigger-${expectedIds[0]} button[data-bs-toggle="collapse"]`)
    .click();
  const detailResponse = await detailResponsePromise;
  expect(new URL(detailResponse.url()).searchParams.has('pageCursor')).toBe(false);
  await expect(firstRecord.locator('.accordion-collapse')).toHaveClass(/\bshow\b/);

  const editorToggle = firstRecord.locator('[data-metadata-editor-trigger]');
  await editorToggle.click();
  const editorTag = firstRecord.locator('[data-observation-editor-tag]');
  await editorTag.fill('Unsaved metadata edit');
  await expect(firstRecord.locator('[data-metadata-editor-active="true"]')).toBeVisible();

  const appendResponsePromise = waitForPageResponse(page, endpoint, firstCursor);
  await moreButton.click();
  const appendResponse = await appendResponsePromise;
  const secondPage = await readPage(appendResponse, endpoint, firstCursor);
  const secondCursor = requireContinuationCursor(secondPage);

  await expect(displayedRecords).toHaveCount(21);
  await expect.poll(() => displayedIds(displayedRecords).then(ids => ids.filter(id => id !== '0')))
    .toEqual(expectedIds.slice(0, 20));
  await expect(list.locator('[data-paginated-record-id="0"]')).toHaveCount(1);
  await expect(draftTag).toHaveValue('UNSAVEDHISTORYDRAFT');
  await expect(draftDate).toHaveValue('2025-03-04');
  await expect(editorTag).toHaveValue('Unsaved metadata edit');
  await expect(firstRecord.locator('[data-metadata-editor-active="true"]')).toBeVisible();
  await expect(firstRecord.locator('.accordion-collapse')).toHaveClass(/\bshow\b/);
  expect(await firstRecord.evaluate((element, original) => element === original, firstRecordHandle)).toBe(true);

  const finalResponsePromise = waitForPageResponse(page, endpoint, secondCursor);
  await moreButton.click();
  await finalResponsePromise;

  await expect(displayedRecords).toHaveCount(PAGINATION_RECORD_COUNT + 1);
  await expect(moreButton).toBeHidden();
  await expect(page.locator('[data-paginated-list-controls="accordion-portfolio-history-management"]'))
    .toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-portfolio-history-management"]')).toBeEmpty();
  await expect(list.locator('[data-paginated-record-id="0"]')).toHaveCount(1);
});

test('scenario 21.3: paginates allocation-plan viewing records in backend order', async ({
  database,
  page,
}) => {
  const seeded = await seedPlans(database);
  const endpoint = `/api/portfolio/${seeded.portfolio.id}/allocation-plan`;
  const initialResponsePromise = waitForPageResponse(page, endpoint, '');

  await page.goto(`/portfolio/${seeded.portfolio.id}/allocation`);

  const initialResponse = await initialResponsePromise;
  const initialPage = await readPage(initialResponse, endpoint, '');
  const firstCursor = requireContinuationCursor(initialPage);
  const list = page.locator('#accordion-allocation-plan');
  const moreButton = loadMoreButton(page, 'accordion-allocation-plan');
  const displayedRecords = list.locator('[data-paginated-record-id]');
  const expectedIds = seeded.records.map(({ id }) => String(id));

  await expect(displayedRecords).toHaveCount(10);
  await expect.poll(() => displayedIds(displayedRecords)).toEqual(expectedIds.slice(0, 10));
  await expect(list.locator('[onclick="navigateToPortfolioAllocationPlanManagement()"]')).toHaveCount(1);
  await expectReadyControl(page, 'accordion-allocation-plan');
  await expectPaginationGapAtWidths(page, 'accordion-allocation-plan');

  const firstPanel = page.locator(`#allocation-plan-${expectedIds[0]}`);
  await expect(firstPanel).toHaveClass(/\bshow\b/);
  const originalCanvas = await firstPanel.locator('canvas').elementHandle();
  expect(originalCanvas).not.toBeNull();

  const appendResponsePromise = waitForPageResponse(page, endpoint, firstCursor);
  await moreButton.click();
  const appendResponse = await appendResponsePromise;
  const secondPage = await readPage(appendResponse, endpoint, firstCursor);
  const secondCursor = requireContinuationCursor(secondPage);

  await expect(displayedRecords).toHaveCount(20);
  await expect.poll(() => displayedIds(displayedRecords)).toEqual(expectedIds.slice(0, 20));
  await expect(firstPanel).toHaveClass(/\bshow\b/);
  expect(await firstPanel.locator('canvas').evaluate((canvas, original) => canvas === original, originalCanvas))
    .toBe(true);
  await expect(page.locator(`#allocation-plan-${expectedIds[10]}`)).not.toHaveClass(/\bshow\b/);
  await expect(list.locator('[onclick="navigateToPortfolioAllocationPlanManagement()"]')).toHaveCount(1);

  const finalResponsePromise = waitForPageResponse(page, endpoint, secondCursor);
  await moreButton.click();
  await finalResponsePromise;

  await expect(displayedRecords).toHaveCount(PAGINATION_RECORD_COUNT);
  await expect.poll(() => displayedIds(displayedRecords)).toEqual(expectedIds);
  await expect(moreButton).toBeHidden();
  await expect(page.locator('[data-paginated-list-controls="accordion-allocation-plan"]')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-allocation-plan"]')).toBeEmpty();
});

test('scenario 21.4: appends allocation-plan management records while keeping one draft form and active edits', async ({
  database,
  page,
}) => {
  const seeded = await seedPlans(database);
  const endpoint = `/api/portfolio/${seeded.portfolio.id}/allocation-plan`;
  const initialResponsePromise = waitForPageResponse(page, endpoint, '');

  await page.goto(`/portfolio/${seeded.portfolio.id}/allocation/manage`);

  const initialResponse = await initialResponsePromise;
  const initialPage = await readPage(initialResponse, endpoint, '');
  const firstCursor = requireContinuationCursor(initialPage);
  const list = page.locator('#accordion-allocation-plan-management');
  const moreButton = loadMoreButton(page, 'accordion-allocation-plan-management');
  const displayedRecords = list.locator('[data-paginated-record-id]');
  const expectedIds = seeded.records.map(({ id }) => String(id));

  await expect(displayedRecords).toHaveCount(11);
  await expect.poll(() => displayedIds(displayedRecords).then(ids => ids.filter(id => id !== '0')))
    .toEqual(expectedIds.slice(0, 10));
  await expect(list.locator('[data-paginated-record-id="0"]')).toHaveCount(1);
  await expect(list.locator(':scope > [data-paginated-record-id]').first())
    .toHaveAttribute('data-paginated-record-id', '0');
  await expect(page.locator('[data-paginated-list-controls="accordion-allocation-plan"]')).toBeHidden();
  await expect(loadMoreButton(page, 'accordion-allocation-plan')).toBeHidden();
  await expectReadyControl(page, 'accordion-allocation-plan-management');
  await expectPaginationGapAtWidths(page, 'accordion-allocation-plan-management');

  const draftName = list.locator('.allocation-plan-management-new-name');
  const draftDate = list.locator('#allocation-plan-management-planned-date');
  await draftName.fill('Unsaved plan draft');
  await draftDate.fill('2025-03-04');

  const firstRecord = page.locator(`#allocation-plan-management-container-${expectedIds[0]}`);
  const firstRecordHandle = await firstRecord.elementHandle();
  const editButton = firstRecord.locator('[data-metadata-editor-trigger]');
  await editButton.click();
  const editorName = firstRecord.locator('[data-allocation-plan-editor-name]');
  await editorName.fill('Unsaved plan metadata edit');
  await expect(firstRecord.locator('[data-metadata-editor-active="true"]')).toBeVisible();
  const originalAllocationRow = await firstRecord.locator('tbody tr').first().elementHandle();

  const appendResponsePromise = waitForPageResponse(page, endpoint, firstCursor);
  await moreButton.click();
  const appendResponse = await appendResponsePromise;
  const secondPage = await readPage(appendResponse, endpoint, firstCursor);
  const secondCursor = requireContinuationCursor(secondPage);

  await expect(displayedRecords).toHaveCount(21);
  await expect.poll(() => displayedIds(displayedRecords).then(ids => ids.filter(id => id !== '0')))
    .toEqual(expectedIds.slice(0, 20));
  await expect(list.locator('[data-paginated-record-id="0"]')).toHaveCount(1);
  await expect(draftName).toHaveValue('Unsaved plan draft');
  await expect(draftDate).toHaveValue('2025-03-04');
  await expect(editorName).toHaveValue('Unsaved plan metadata edit');
  await expect(firstRecord.locator('[data-metadata-editor-active="true"]')).toBeVisible();
  expect(await firstRecord.evaluate((element, original) => element === original, firstRecordHandle)).toBe(true);
  expect(await firstRecord.locator('tbody tr').first().evaluate((row, original) => row === original, originalAllocationRow))
    .toBe(true);

  const appendedRecord = page.locator(`#allocation-plan-management-container-${expectedIds[10]}`);
  await expect(appendedRecord.locator('.accordion-collapse')).not.toHaveClass(/\bshow\b/);
  await appendedRecord.locator('button[data-bs-toggle="collapse"]').click();
  await expect(appendedRecord.locator('.accordion-collapse')).toHaveClass(/\bshow\b/);
  const appendedRows = appendedRecord.locator('tbody tr');
  await expect(appendedRows).toHaveCount(2);
  await appendedRecord.locator('tfoot button[type="button"]').first().click();
  await expect(appendedRows).toHaveCount(3);

  const finalResponsePromise = waitForPageResponse(page, endpoint, secondCursor);
  await moreButton.click();
  await finalResponsePromise;

  await expect(displayedRecords).toHaveCount(PAGINATION_RECORD_COUNT + 1);
  await expect.poll(() => displayedIds(displayedRecords).then(ids => ids.filter(id => id !== '0')))
    .toEqual(expectedIds);
  await expect(moreButton).toBeHidden();
  await expect(page.locator('[data-paginated-list-controls="accordion-allocation-plan-management"]'))
    .toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-allocation-plan-management"]')).toBeEmpty();
  await expect(list.locator('[data-paginated-record-id="0"]')).toHaveCount(1);
});

test('scenario 21.5: handles empty history and allocation-plan lists without continuation controls', async ({
  database,
  page,
}) => {
  const emptyHistory = await seedHistory(database, 0);
  const historyEndpoint = `/api/portfolio/${emptyHistory.portfolio.id}/history`;

  let responsePromise = waitForPageResponse(page, historyEndpoint, '');
  await page.goto(`/portfolio/${emptyHistory.portfolio.id}/history`);
  let response = await responsePromise;
  let responsePage = await readPage(response, historyEndpoint, '');
  expect(responsePage.items).toHaveLength(0);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-portfolio-history [data-paginated-record-id]')).toHaveCount(0);
  await expect(loadMoreButton(page, 'accordion-portfolio-history')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-portfolio-history"]'))
    .toHaveText('No records found.');

  const historyManagementEndpoint = `/api/portfolio/${emptyHistory.portfolio.id}/history/observation`;
  const historyManagementStatus = page.locator('[data-paginated-list-status="accordion-portfolio-history-management"]');
  responsePromise = waitForPageResponse(page, historyManagementEndpoint, '');
  await page.locator('#accordion-portfolio-history > button[onclick="navigateToPortfolioAllocationManagement()"]')
    .click();
  response = await responsePromise;
  responsePage = await readPage(response, historyManagementEndpoint, '');
  expect(responsePage.items).toHaveLength(0);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-portfolio-history-management [data-paginated-record-id]'))
    .toHaveCount(1);
  await expect(page.locator('#accordion-portfolio-history-management [data-paginated-record-id="0"]'))
    .toHaveCount(1);
  await expect(loadMoreButton(page, 'accordion-portfolio-history-management')).toBeHidden();
  await expect(historyManagementStatus).toHaveText('No records found.');
  await expect(page.locator('[data-paginated-list-status="accordion-portfolio-history"]')).toBeHidden();

  const emptyPlans = await seedPlans(database, 0);
  const plansEndpoint = `/api/portfolio/${emptyPlans.portfolio.id}/allocation-plan`;

  responsePromise = waitForPageResponse(page, plansEndpoint, '');
  await page.goto(`/portfolio/${emptyPlans.portfolio.id}/allocation`);
  response = await responsePromise;
  responsePage = await readPage(response, plansEndpoint, '');
  expect(responsePage.items).toHaveLength(0);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-allocation-plan [data-paginated-record-id]')).toHaveCount(0);
  await expect(loadMoreButton(page, 'accordion-allocation-plan')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-allocation-plan"]'))
    .toHaveText('No records found.');

  responsePromise = waitForPageResponse(page, plansEndpoint, '');
  await page.locator('#accordion-allocation-plan > button[onclick="navigateToPortfolioAllocationPlanManagement()"]')
    .click();
  response = await responsePromise;
  responsePage = await readPage(response, plansEndpoint, '');
  expect(responsePage.items).toHaveLength(0);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-allocation-plan-management [data-paginated-record-id]'))
    .toHaveCount(1);
  await expect(page.locator('#accordion-allocation-plan-management [data-paginated-record-id="0"]'))
    .toHaveCount(1);
  await expect(loadMoreButton(page, 'accordion-allocation-plan-management')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-allocation-plan-management"]'))
    .toHaveText('No records found.');
  await expect(page.locator('[data-paginated-list-status="accordion-allocation-plan"]')).toBeHidden();
});

test('scenario 21.6: treats exactly ten records as terminal on all four lists', async ({ database, page }) => {
  const history = await seedHistory(database, 10);
  const historyEndpoint = `/api/portfolio/${history.portfolio.id}/history`;

  let responsePromise = waitForPageResponse(page, historyEndpoint, '');
  await page.goto(`/portfolio/${history.portfolio.id}/history`);
  let responsePage = await readPage(await responsePromise, historyEndpoint, '');
  expect(responsePage.items).toHaveLength(10);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-portfolio-history [data-paginated-record-id]')).toHaveCount(10);
  await expect(loadMoreButton(page, 'accordion-portfolio-history')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-portfolio-history"]'))
    .toBeEmpty();
  await expect(page.locator('[data-paginated-list-controls="accordion-portfolio-history"]')).toBeHidden();

  const historyManagementEndpoint = `/api/portfolio/${history.portfolio.id}/history/observation`;
  responsePromise = waitForPageResponse(page, historyManagementEndpoint, '');
  await page.goto(`/portfolio/${history.portfolio.id}/history/manage`);
  responsePage = await readPage(await responsePromise, historyManagementEndpoint, '');
  expect(responsePage.items).toHaveLength(10);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-portfolio-history-management [data-paginated-record-id]'))
    .toHaveCount(11);
  await expect(page.locator('#accordion-portfolio-history-management [data-paginated-record-id="0"]'))
    .toHaveCount(1);
  await expect(loadMoreButton(page, 'accordion-portfolio-history-management')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-portfolio-history-management"]'))
    .toBeEmpty();
  await expect(page.locator('[data-paginated-list-controls="accordion-portfolio-history-management"]'))
    .toBeHidden();

  const plans = await seedPlans(database, 10);
  const plansEndpoint = `/api/portfolio/${plans.portfolio.id}/allocation-plan`;

  responsePromise = waitForPageResponse(page, plansEndpoint, '');
  await page.goto(`/portfolio/${plans.portfolio.id}/allocation`);
  responsePage = await readPage(await responsePromise, plansEndpoint, '');
  expect(responsePage.items).toHaveLength(10);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-allocation-plan [data-paginated-record-id]')).toHaveCount(10);
  await expect(loadMoreButton(page, 'accordion-allocation-plan')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-allocation-plan"]'))
    .toBeEmpty();
  await expect(page.locator('[data-paginated-list-controls="accordion-allocation-plan"]')).toBeHidden();

  responsePromise = waitForPageResponse(page, plansEndpoint, '');
  await page.goto(`/portfolio/${plans.portfolio.id}/allocation/manage`);
  responsePage = await readPage(await responsePromise, plansEndpoint, '');
  expect(responsePage.items).toHaveLength(10);
  expect(responsePage.nextCursor).toBeNull();
  await expect(page.locator('#accordion-allocation-plan-management [data-paginated-record-id]'))
    .toHaveCount(11);
  await expect(page.locator('#accordion-allocation-plan-management [data-paginated-record-id="0"]'))
    .toHaveCount(1);
  await expect(loadMoreButton(page, 'accordion-allocation-plan-management')).toBeHidden();
  await expect(page.locator('[data-paginated-list-status="accordion-allocation-plan-management"]'))
    .toBeEmpty();
  await expect(page.locator('[data-paginated-list-controls="accordion-allocation-plan-management"]'))
    .toBeHidden();
});

test('scenario 21.7: isolates history controls when a hidden viewing page settles late', async ({ database, page }) => {
  const seeded = await seedHistory(database);

  await verifyLateViewResponseIsolation(page, {
    viewingRoute: `/portfolio/${seeded.portfolio.id}/history`,
    managementRoute: `/portfolio/${seeded.portfolio.id}/history/manage`,
    viewingEndpoint: `/api/portfolio/${seeded.portfolio.id}/history`,
    managementEndpoint: `/api/portfolio/${seeded.portfolio.id}/history/observation`,
    viewingOwnerId: 'accordion-portfolio-history',
    managementOwnerId: 'accordion-portfolio-history-management',
    viewingManagementButton: '#accordion-portfolio-history > button[onclick="navigateToPortfolioAllocationManagement()"]',
    managementViewingButton: 'button[onclick="portfolioHistoryManagement.navigateToPortfolioAllocationViewing()"]',
    managementDraftSelector: '.portfolio-history-management-new-time-tag',
    draftValue: 'Late history response draft',
  });
});

test('scenario 21.8: isolates allocation-plan controls when a hidden viewing page settles late', async ({
  database,
  page,
}) => {
  const seeded = await seedPlans(database);

  await verifyLateViewResponseIsolation(page, {
    viewingRoute: `/portfolio/${seeded.portfolio.id}/allocation`,
    managementRoute: `/portfolio/${seeded.portfolio.id}/allocation/manage`,
    viewingEndpoint: `/api/portfolio/${seeded.portfolio.id}/allocation-plan`,
    managementEndpoint: `/api/portfolio/${seeded.portfolio.id}/allocation-plan`,
    viewingOwnerId: 'accordion-allocation-plan',
    managementOwnerId: 'accordion-allocation-plan-management',
    viewingManagementButton: '#accordion-allocation-plan > button[onclick="navigateToPortfolioAllocationPlanManagement()"]',
    managementViewingButton: 'button[onclick="allocationPlanManagement.navigateToAllocationPlansViewing()"]',
    managementDraftSelector: '.allocation-plan-management-new-name',
    draftValue: 'Late plan response draft',
  });
});

test('scenario 21.9: moves keyboard focus to the newly appended record after terminal pagination', async ({
  database,
  page,
}) => {
  const seeded = await seedHistory(database, 20);
  const endpoint = `/api/portfolio/${seeded.portfolio.id}/history`;
  const initialResponsePromise = waitForPageResponse(page, endpoint, '');

  await page.goto(`/portfolio/${seeded.portfolio.id}/history`);

  const initialResponse = await initialResponsePromise;
  const firstPage = await readPage(initialResponse, endpoint, '');
  const finalCursor = requireContinuationCursor(firstPage);
  const ownerId = 'accordion-portfolio-history';
  const moreButton = loadMoreButton(page, ownerId);
  await expectReadyControl(page, ownerId);

  const finalResponsePromise = waitForPageResponse(page, endpoint, finalCursor);
  await moreButton.focus();
  await expect(moreButton).toBeFocused();
  await page.keyboard.press('Enter');

  const finalResponse = await finalResponsePromise;
  const finalPage = await readPage(finalResponse, endpoint, finalCursor);
  expect(finalPage.nextCursor).toBeNull();
  await expect(page.locator(`#${ownerId} [data-paginated-record-id]`)).toHaveCount(20);
  await expect(moreButton).toBeHidden();
  await expect(page.locator(
    `#${ownerId} [data-paginated-record-id="${seeded.records[10].id}"] .accordion-button`,
  )).toBeFocused();
  await expect(page.locator(`[data-paginated-list-controls="${ownerId}"]`)).toBeHidden();
});

test('scenario 21.10: keeps an initial-load retry visible and busy without a cursor', async ({ database, page }) => {
  const seeded = await seedHistory(database);
  const endpoint = `/api/portfolio/${seeded.portfolio.id}/history`;
  const ownerId = 'accordion-portfolio-history';
  const status = page.locator(`[data-paginated-list-status="${ownerId}"]`);
  const moreButton = loadMoreButton(page, ownerId);
  let failInitialRequest = true;
  let holdRetryResponse = false;
  let releaseRetryResponse: () => void = () => {};
  let signalRetryStarted: () => void = () => {};
  const retryResponseGate = new Promise<void>(resolve => {
    releaseRetryResponse = resolve;
  });
  const retryRequestStarted = new Promise<void>(resolve => {
    signalRetryStarted = resolve;
  });

  await page.route(`**${endpoint}?*`, async route => {
    const requestUrl = new URL(route.request().url());

    if(requestUrl.searchParams.get('pageCursor') !== '') {
      await route.continue();
      return;
    }

    if(failInitialRequest) {
      failInitialRequest = false;
      await route.abort();
      return;
    }

    if(holdRetryResponse) {
      holdRetryResponse = false;
      signalRetryStarted();
      await retryResponseGate;
    }

    await route.continue();
  });

  try {
    await page.goto(`/portfolio/${seeded.portfolio.id}/history`);
    await expect(status).toHaveAttribute('role', 'alert');
    await expect(status).toHaveText('Could not load records. Try again.');
    await expect(moreButton).toBeVisible();
    await expect(moreButton).toHaveAccessibleName('Retry loading records');
    await expect(moreButton.locator('[data-paginated-list-retry-icon]')).toBeVisible();

    holdRetryResponse = true;
    const retryResponsePromise = waitForPageResponse(page, endpoint, '');
    await moreButton.click();
    await retryRequestStarted;
    await expect(moreButton).toBeVisible();
    await expect(moreButton).toBeDisabled();
    await expect(moreButton).toHaveAccessibleName('Loading records');
    await expect(moreButton).toHaveAttribute('aria-busy', 'true');
    await expect(moreButton.locator('[data-paginated-list-loading-indicator]')).toBeVisible();

    releaseRetryResponse();
    const retryResponse = await retryResponsePromise;
    const retryPage = await readPage(retryResponse, endpoint, '');
    expect(retryPage.nextCursor).not.toBeNull();
    await expectReadyControl(page, ownerId);
    await expect(status).toBeEmpty();
    await expect(page.locator(`#${ownerId} [data-paginated-record-id]`)).toHaveCount(10);
  } finally {
    releaseRetryResponse();
  }
});

/** Verifies icon-only button state, ownership, accessibility, and empty child markup. */
async function expectReadyControl(page: Page, ownerId: string): Promise<void> {
  const button = loadMoreButton(page, ownerId);
  const controls = page.locator(`[data-paginated-list-controls="${ownerId}"]`);

  await expect(controls).toBeVisible();
  await expect(button).toBeVisible();
  await expect(button).toBeEnabled();
  await expect(button).toHaveClass(/\bbtn-primary\b/);
  await expect(button).not.toHaveClass(/btn-outline/);
  await expect(button).toHaveAccessibleName('Show more');
  await expect(button).toHaveAttribute('title', 'Show more');
  await expect(button).toHaveAttribute('aria-busy', 'false');
  await expect(button).toBeEmpty();
  await expect(button.locator('[data-paginated-list-more-icon]')).toBeVisible();
  await expect(button.locator('[data-paginated-list-retry-icon]')).toBeHidden();
  await expect(button.locator('[data-paginated-list-loading-indicator]')).toBeHidden();
}

/** Verifies the configured 1.5rem list-to-control gap at desktop and narrow widths. */
async function expectPaginationGapAtWidths(page: Page, ownerId: string): Promise<void> {
  const originalViewport = page.viewportSize() ?? { width: 1280, height: 900 };

  try {
    for(const viewport of [{ width: 1280, height: 900 }, { width: 390, height: 844 }]) {
      await page.setViewportSize(viewport);
      await expect.poll(async () => page.evaluate((id) => {
        const accordion = document.getElementById(id);
        const button = document.querySelector<HTMLElement>(`[data-paginated-list-more="${id}"]`);

        if(!accordion || !button || button.getClientRects().length === 0) {
          return Number.POSITIVE_INFINITY;
        }

        const gap = button.getBoundingClientRect().top - accordion.getBoundingClientRect().bottom;
        const expectedGap = parseFloat(getComputedStyle(document.documentElement).fontSize) * 1.5;
        return Math.abs(gap - expectedGap);
      }, ownerId)).toBeLessThanOrEqual(1);
    }
  } finally {
    await page.setViewportSize(originalViewport);
  }
}

/** Exercises a delayed viewing append across route changes without leaking controls or stealing focus. */
async function verifyLateViewResponseIsolation(
  page: Page,
  config: LateViewResponseIsolationConfig,
): Promise<void> {
  const initialResponsePromise = waitForPageResponse(page, config.viewingEndpoint, '');
  await page.goto(config.viewingRoute);

  const initialResponse = await initialResponsePromise;
  const initialPage = await readPage(initialResponse, config.viewingEndpoint, '');
  const continuationCursor = requireContinuationCursor(initialPage);
  const moreButton = loadMoreButton(page, config.viewingOwnerId);
  const viewingControls = page.locator(`[data-paginated-list-controls="${config.viewingOwnerId}"]`);
  const managementControls = page.locator(`[data-paginated-list-controls="${config.managementOwnerId}"]`);
  const viewingStatus = page.locator(`[data-paginated-list-status="${config.viewingOwnerId}"]`);
  const draft = page.locator(config.managementDraftSelector);

  await expectReadyControl(page, config.viewingOwnerId);
  await expectPaginationGapAtWidths(page, config.viewingOwnerId);

  let releaseDelayedResponse: () => void = () => {};
  let signalDelayedRequest: () => void = () => {};
  let delayContinuation = true;
  const delayedResponseGate = new Promise<void>(resolve => {
    releaseDelayedResponse = resolve;
  });
  const delayedRequestStarted = new Promise<void>(resolve => {
    signalDelayedRequest = resolve;
  });
  const routePattern = `**${config.viewingEndpoint}?*`;

  await page.route(routePattern, async route => {
    const requestUrl = new URL(route.request().url());

    if(delayContinuation && requestUrl.searchParams.get('pageCursor') === continuationCursor) {
      delayContinuation = false;
      signalDelayedRequest();
      await delayedResponseGate;
    }

    await route.continue();
  });

  try {
    const appendResponsePromise = waitForPageResponse(page, config.viewingEndpoint, continuationCursor);
    await moreButton.click();
    await delayedRequestStarted;

    const managementResponsePromise = waitForPageResponse(page, config.managementEndpoint, '');
    await page.locator(config.viewingManagementButton).click();
    const managementResponse = await managementResponsePromise;
    const managementPage = await readPage(managementResponse, config.managementEndpoint, '');
    expect(managementPage.nextCursor).not.toBeNull();
    await expect.poll(() => new URL(page.url()).pathname).toBe(config.managementRoute);
    await expect(viewingControls).toBeHidden();
    await expect(viewingStatus).toBeHidden();
    await expect(managementControls).toBeVisible();
    await expectReadyControl(page, config.managementOwnerId);
    await expectPaginationGapAtWidths(page, config.managementOwnerId);

    await draft.fill(config.draftValue);
    const enteredDraftValue = await draft.inputValue();
    await expect(draft).toBeFocused();

    releaseDelayedResponse();
    const appendResponse = await appendResponsePromise;
    const appendedPage = await readPage(appendResponse, config.viewingEndpoint, continuationCursor);
    expect(appendedPage.items).toHaveLength(10);
    await expect(page.locator(`#${config.viewingOwnerId} [data-paginated-record-id]`)).toHaveCount(20);
    await expect(viewingControls).toBeHidden();
    await expect(viewingStatus).toBeHidden();
    await expect(draft).toHaveValue(enteredDraftValue);
    await expect(draft).toBeFocused();

    await page.locator(config.managementViewingButton).click();
    await expect.poll(() => new URL(page.url()).pathname).toBe(config.viewingRoute);
    await expect(viewingControls).toBeVisible();
    await expect(moreButton).toBeVisible();
    await expect(managementControls).toBeHidden();
  } finally {
    releaseDelayedResponse();
    await page.unroute(routePattern);
  }
}

/** Seeds the requested number of observations, including a boundary timestamp tie when enough exist. */
async function seedHistory(
  database: E2eDatabase,
  recordCount = PAGINATION_RECORD_COUNT,
): Promise<SeededHistory> {
  const portfolio = await seedPortfolio(database);
  const assets = await database.query<{ id: number }>(
    `INSERT INTO public.asset (ticker, name)
     VALUES ('PAGINATION-ASSET', 'Pagination asset')
     RETURNING id`,
  );
  const records: SeededHistoryRecord[] = [];

  for(let index = 0; index < recordCount; index++) {
    const timeTag = `Pagination history ${String(index).padStart(2, '0')}`;
    const timestamp = historyTimestamp(index);
    const observations = await database.query<SeededHistoryRecord>(
      `INSERT INTO public.portfolio_allocation_obs_time
         (portfolio_id, observation_time_tag, observation_timestamp)
       VALUES ($1, $2, $3::timestamptz)
       RETURNING id, observation_time_tag AS "timeTag"`,
      [portfolio.id, timeTag, timestamp],
    );
    const observation = observations[0];

    await database.query(
      `INSERT INTO public.portfolio_allocation_fact
         (portfolio_id, asset_id, class, cash_reserve, asset_quantity, asset_market_price,
          total_market_value, observation_time_id)
       VALUES ($1, $2, 'BONDS', false, 1, 100, 100, $3)`,
      [portfolio.id, assets[0].id, observation.id],
    );
    records.push(observation);
  }

  const orderedRecords = await database.query<SeededHistoryRecord>(
    `SELECT id, observation_time_tag AS "timeTag"
     FROM public.portfolio_allocation_obs_time
     WHERE portfolio_id = $1
     ORDER BY observation_timestamp DESC, id DESC`,
    [portfolio.id],
  );

  return { portfolio, records: orderedRecords };
}

/** Seeds the requested number of complete plans, including a boundary date tie when enough exist. */
async function seedPlans(database: E2eDatabase, recordCount = PAGINATION_RECORD_COUNT): Promise<SeededPlans> {
  const portfolio = await seedPortfolio(database);
  const assets = await database.query<{ id: number; ticker: string }>(
    `INSERT INTO public.asset (ticker, name)
     VALUES ('PAGINATION-PLAN-ASSET', 'Pagination plan asset')
     RETURNING id, ticker`,
  );

  for(let index = 0; index < recordCount; index++) {
    const name = `Pagination plan ${String(index).padStart(2, '0')}`;
    const plans = await database.query<SeededPlanRecord>(
      `INSERT INTO public.allocation_plan
         (portfolio_id, name, type, planned_execution_date, planned_date, create_timestamp)
       VALUES ($1, $2, 'ALLOCATION_PLAN', NULL, $3::date, CURRENT_TIMESTAMP)
       RETURNING id, name`,
      [portfolio.id, name, planDate(index)],
    );

    await database.query(
      `INSERT INTO public.planned_allocation
         (allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
       VALUES
         ($1, ARRAY[NULL::text, 'BONDS'], NULL, false, 1.00000, NULL),
         ($1, ARRAY[$2::text, 'BONDS'], $3, false, 1.00000, NULL)`,
      [plans[0].id, assets[0].ticker, assets[0].id],
    );
  }

  const orderedRecords = await database.query<SeededPlanRecord>(
    `SELECT id, name
     FROM public.allocation_plan
     WHERE portfolio_id = $1
     ORDER BY planned_date DESC, id DESC`,
    [portfolio.id],
  );

  return { portfolio, records: orderedRecords };
}

/** Seeds a portfolio with the hierarchy shared by history and plan fixtures. */
async function seedPortfolio(database: E2eDatabase): Promise<SeededPortfolio> {
  const portfolios = await database.query<SeededPortfolio>(
    `INSERT INTO public.portfolio (name, allocation_structure)
     VALUES ($1, $2::jsonb)
     RETURNING id, name`,
    [PORTFOLIO_NAME, JSON.stringify(ALLOCATION_STRUCTURE)],
  );

  return portfolios[0];
}

/** Returns a descending timestamp while deliberately tying records nine and ten. */
function historyTimestamp(index: number): string {
  const tiedIndex = index === 10 ? 9 : index;
  return new Date(Date.UTC(2026, 0, 31 - tiedIndex, 12)).toISOString();
}

/** Returns a descending calendar date while deliberately tying records nine and ten. */
function planDate(index: number): string {
  const tiedIndex = index === 10 ? 9 : index;
  return new Date(Date.UTC(2026, 0, 31 - tiedIndex)).toISOString().slice(0, 10);
}

/** Waits for a pagination response from one endpoint and cursor value. */
function waitForPageResponse(page: Page, endpoint: string, cursor: string): Promise<Response> {
  return page.waitForResponse(response => {
    const url = new URL(response.url());
    return url.pathname === endpoint
      && url.searchParams.has('pageCursor')
      && url.searchParams.get('pageCursor') === cursor;
  });
}

/** Validates the page envelope and confirms HTMX serialized its opaque cursor unchanged. */
async function readPage(response: Response, endpoint: string, expectedCursor: string): Promise<PaginatedResponse> {
  expect(response.status()).toBe(200);
  const url = new URL(response.url());
  expect(url.pathname).toBe(endpoint);
  expect(url.searchParams.has('pageCursor')).toBe(true);
  expect(url.searchParams.get('pageCursor')).toBe(expectedCursor);

  const page = await response.json() as PaginatedResponse;
  expect(Array.isArray(page.items)).toBe(true);
  expect(page.nextCursor === null || typeof page.nextCursor === 'string').toBe(true);

  return page;
}

/** Reads rendered persisted IDs in DOM order. */
async function displayedIds(records: Locator): Promise<string[]> {
  return records.evaluateAll(elements => elements.map(element => element.getAttribute('data-paginated-record-id') ?? ''));
}

/** Resolves one list's native continuation control independently of other mounted list views. */
function loadMoreButton(page: Page, ownerId: string): Locator {
  return page.locator(`[data-paginated-list-more="${ownerId}"]`);
}

/** Returns a continuation token and fails clearly when a fixture unexpectedly has no next page. */
function requireContinuationCursor(page: PaginatedResponse): string {
  if(page.nextCursor === null) {
    throw new Error('Expected the seeded page to include a continuation cursor.');
  }

  return page.nextCursor;
}
