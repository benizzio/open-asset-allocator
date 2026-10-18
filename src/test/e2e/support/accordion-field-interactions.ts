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
