/**
 * Exercises accordion header fields and records Bootstrap collapse events for E2E assertions.
 *
 * @author GPT-6 Luna
 */
import type { Locator, Page, Request } from '@playwright/test';
import { expect } from './fixtures';

/** Identifies the editable fields and accordion controls exercised by the shared interaction test.
 * @author GPT-6 Luna
 */
type AccordionFieldControls = {
  dateInput: Locator;
  expectedDate: string;
  expectedText: string;
  panel: Locator;
  textInput: Locator;
  toggle: Locator;
};

/** Counts the collapse transitions observed during a field-interaction phase.
 * @author GPT-6 Luna
 */
type CollapseEventCounts = {
  hide: number;
  show: number;
};

/** Counts the Bootstrap collapse transitions observed for one accordion panel.
 * @author GPT-6 Luna
 */
export type AccordionTransitionCounts = CollapseEventCounts;

/** Holds transition listeners and their counts until the interaction test finishes.
 * @author GPT-6 Luna
 */
type CollapseEventRecorder = {
  counts: CollapseEventCounts;
  hideListener: EventListener;
  showListener: EventListener;
};

/** Adds test-only collapse recorder storage to the page window.
 * @author GPT-6 Luna
 */
type AccordionTestWindow = Window & {
  __accordionFieldInteractionRecorders?: Map<string, CollapseEventRecorder>;
};

let recorderSequence = 0;

/** Installs counters for the next Bootstrap show and hide events on one panel.
 *
 * Use before exercising an accordion header to detect duplicate toggles that could be hidden by checking only its final
 * state. Example: `await installAccordionTransitionCounter(panel)` starts counting that panel's transitions.
 * @param panel - Accordion collapse panel whose Bootstrap events should be counted.
 * @author GPT-6 Luna
 */
export async function installAccordionTransitionCounter(panel: Locator): Promise<void> {
  await panel.evaluate((element) => {
    element.dataset.e2eCollapseShowCount = '0';
    element.dataset.e2eCollapseHideCount = '0';
    element.addEventListener('show.bs.collapse', () => {
      element.dataset.e2eCollapseShowCount = String(Number(element.dataset.e2eCollapseShowCount) + 1);
    });
    element.addEventListener('hide.bs.collapse', () => {
      element.dataset.e2eCollapseHideCount = String(Number(element.dataset.e2eCollapseHideCount) + 1);
    });
  });
}

/** Reads Bootstrap show/hide event totals recorded on one panel.
 *
 * Example: `await readAccordionTransitionCounts(panel)` returns `{ show: 1, hide: 0 }` after one opening.
 * @param panel - Accordion collapse panel with an installed transition counter.
 * @returns The number of show and hide events observed so far.
 * @author GPT-6 Luna
 */
export async function readAccordionTransitionCounts(panel: Locator): Promise<AccordionTransitionCounts> {
  return panel.evaluate((element) => ({
    show: Number(element.dataset.e2eCollapseShowCount),
    hide: Number(element.dataset.e2eCollapseHideCount),
  }));
}

/** Verifies a read-only metadata label and its header padding each toggle the panel exactly once.
 *
 * The helper starts from either expanded state and returns the panel to that same state. Example:
 * `await exerciseReadOnlyManagementHeaderToggle(header, label, panel)` checks the read-only header's click surfaces.
 * @param header - Persisted metadata accordion header.
 * @param label - Visible read-only label inside the header.
 * @param panel - The header's Bootstrap collapse panel.
 * @author GPT-6 Luna
 */
export async function exerciseReadOnlyManagementHeaderToggle(
  header: Locator,
  label: Locator,
  panel: Locator,
): Promise<void> {
  await expect(header).toHaveAttribute('data-metadata-editor-state', 'read-only');
  await expect(label).toBeVisible();
  await expect(panel).not.toHaveClass(/\bcollapsing\b/);

  const initiallyExpanded = await panel.evaluate(element => element.classList.contains('show'));
  await installAccordionTransitionCounter(panel);
  await label.click();
  if(initiallyExpanded) {
    await expect(panel).not.toHaveClass(/\bshow\b/);
  }
  else {
    await expect(panel).toHaveClass(/\bshow\b/);
  }
  await expect(panel).not.toHaveClass(/\bcollapsing\b/);

  const headerBounds = await header.boundingBox();
  if(!headerBounds) {
    throw new Error('The read-only management header must have rendered geometry.');
  }

  await header.click({ position: { x: 2, y: headerBounds.height / 2 } });
  if(initiallyExpanded) {
    await expect(panel).toHaveClass(/\bshow\b/);
  }
  else {
    await expect(panel).not.toHaveClass(/\bshow\b/);
  }
  await expect(panel).not.toHaveClass(/\bcollapsing\b/);
  expect(await readAccordionTransitionCounts(panel)).toEqual({ show: 1, hide: 1 });
}

/**
 * Repeats pointer, text, date, and keyboard interactions while verifying that the owning panel stays expanded.
 * The assertion observes the panel's `hide.bs.collapse` event, so a transient close and reopen fails even if the
 * panel is open at the end. It also checks for API requests caused by field interaction, preserves field values,
 * verifies focus movement, and checks that the external toggle can close and reopen the panel.
 *
 * @param page - Playwright page emitting browser requests and keyboard input.
 * @param controls - The new-record header controls and expected values after interaction.
 * @example
 * ```ts
 * await exerciseAccordionFieldInteractions(page, {
 *   dateInput: item.getByLabel('Planned date'),
 *   expectedDate: '2026-04-20',
 *   expectedText: 'E2E_PLAN_DRAFT',
 *   panel: item.locator('#allocation-plan-management-0'),
 *   textInput: item.getByRole('textbox', { name: 'New allocation plan name' }),
 *   toggle: item.getByRole('button', { name: 'New allocation plan details' }),
 * });
 * ```
 * @author GPT-6 Luna
 */
export async function exerciseAccordionFieldInteractions(
  page: Page,
  controls: AccordionFieldControls,
): Promise<void> {
  const recorderId = `accordion-field-${recorderSequence++}`;
  const apiRequests: string[] = [];
  const trackApiRequest = (request: Request): void => {
    const requestUrl = new URL(request.url());
    if(requestUrl.pathname.startsWith('/api/')) {
      apiRequests.push(`${request.method()} ${requestUrl.pathname}`);
    }
  };

  await installCollapseEventRecorder(controls.panel, recorderId);
  page.on('request', trackApiRequest);

  try {
    await expect(controls.toggle).toHaveAttribute('aria-expanded', 'false');
    await expect(controls.panel).not.toHaveClass(/\bshow\b/);

    await controls.textInput.click();
    await expectPanelExpanded(controls);
    await controls.textInput.pressSequentially(`${controls.expectedText}_DRAFT`);
    await controls.textInput.click();
    await controls.textInput.selectText();
    await controls.textInput.pressSequentially(controls.expectedText);

    const dateBeforeKeyboardInteraction = await controls.dateInput.inputValue();
    await controls.dateInput.click();
    await controls.dateInput.press('ArrowUp');
    await expect(controls.dateInput).not.toHaveValue(dateBeforeKeyboardInteraction);
    await controls.dateInput.fill(controls.expectedDate);
    await controls.dateInput.click();
    await page.keyboard.press('Escape');

    await controls.textInput.focus();
    await page.keyboard.press('Tab');
    await expect(controls.dateInput).toBeFocused();
    await page.keyboard.press('Shift+Tab');
    await expect(controls.textInput).toBeFocused();

    await expect(controls.textInput).toHaveValue(controls.expectedText);
    await expect(controls.dateInput).toHaveValue(controls.expectedDate);
    await expectPanelExpanded(controls);
    expect(await readCollapseEventCounts(controls.panel, recorderId)).toEqual({ hide: 0, show: 1 });
    expect(apiRequests).toEqual([]);

    await controls.toggle.click();
    await expect(controls.toggle).toHaveAttribute('aria-expanded', 'false');
    await expect(controls.panel).not.toHaveClass(/\bshow\b/);
    await expect(controls.panel).not.toHaveClass(/\bcollapsing\b/);
    expect(await readCollapseEventCounts(controls.panel, recorderId)).toEqual({ hide: 1, show: 1 });

    await controls.dateInput.click();
    await expect(controls.dateInput).toBeFocused();
    await expectPanelExpanded(controls);
    await expect(controls.textInput).toHaveValue(controls.expectedText);
    await expect(controls.dateInput).toHaveValue(controls.expectedDate);
    expect(await readCollapseEventCounts(controls.panel, recorderId)).toEqual({ hide: 1, show: 2 });
    expect(apiRequests).toEqual([]);
  } finally {
    page.off('request', trackApiRequest);
    await removeCollapseEventRecorder(controls.panel, recorderId);
  }
}

/** Registers `show` and `hide` listeners on one panel for this interaction check.
 * @author GPT-6 Luna
 */
async function installCollapseEventRecorder(panel: Locator, recorderId: string): Promise<void> {
  await panel.evaluate((element, key) => {
    const testWindow = window as AccordionTestWindow;
    const recorders = testWindow.__accordionFieldInteractionRecorders ?? new Map();
    const counts = { hide: 0, show: 0 };
    const showListener: EventListener = () => counts.show++;
    const hideListener: EventListener = () => counts.hide++;

    element.addEventListener('show.bs.collapse', showListener);
    element.addEventListener('hide.bs.collapse', hideListener);
    recorders.set(key, { counts, hideListener, showListener });
    testWindow.__accordionFieldInteractionRecorders = recorders;
  }, recorderId);
}

/** Reads the recorded collapse-event counts for a panel.
 * @author GPT-6 Luna
 */
async function readCollapseEventCounts(panel: Locator, recorderId: string): Promise<CollapseEventCounts> {
  return panel.evaluate((element, key) => {
    const recorders = (window as AccordionTestWindow).__accordionFieldInteractionRecorders;
    const recorder = recorders?.get(key);
    if(!recorder) {
      throw new Error(`No collapse-event recorder exists for ${key}.`);
    }
    return { ...recorder.counts };
  }, recorderId);
}

/** Removes recorder listeners so they cannot affect later accordion interactions.
 * @author GPT-6 Luna
 */
async function removeCollapseEventRecorder(panel: Locator, recorderId: string): Promise<void> {
  await panel.evaluate((element, key) => {
    const testWindow = window as AccordionTestWindow;
    const recorders = testWindow.__accordionFieldInteractionRecorders;
    const recorder = recorders?.get(key);
    if(!recorders || !recorder) {
      return;
    }

    element.removeEventListener('show.bs.collapse', recorder.showListener);
    element.removeEventListener('hide.bs.collapse', recorder.hideListener);
    recorders.delete(key);
    if(recorders.size === 0) {
      delete testWindow.__accordionFieldInteractionRecorders;
    }
  }, recorderId);
}

/** Waits for the panel's opening transition and checks its accessible expanded state.
 * @author GPT-6 Luna
 */
async function expectPanelExpanded(controls: AccordionFieldControls): Promise<void> {
  await expect(controls.panel).toHaveClass(/\bshow\b/);
  await expect(controls.panel).not.toHaveClass(/\bcollapsing\b/);
  await expect(controls.toggle).toHaveAttribute('aria-expanded', 'true');
}
