/** Provides geometry and alignment assertions shared by management-header E2E tests.
 *
 * Import the exported assertions to compare persisted metadata editors with new-record headers and verify the resting
 * pencil's alignment. Example: `await expectManagementMetadataHeaderLayout(persisted, newRecord)` checks the layout.
 * @author GPT-6 Luna
 */

import { expect, type Locator } from '@playwright/test';

/** Describes the controls whose rendered geometry makes up one metadata accordion header.
 * @author GPT-6 Luna
 */
type MetadataHeaderControls = {
  header: Locator;
  textInput: Locator;
  dateInput: Locator;
  toggle: Locator;
};

/** Verifies that a persisted editor matches its new-record header at the current viewport size.
 *
 * The helper compares field and chevron geometry, checks that each control stays inside its header, and rejects
 * overlapping text/date fields. Use it after revealing the persisted editor at both desktop and narrow viewports.
 * Example: `await expectManagementMetadataHeaderLayout(persistedControls, newRecordControls)` validates both layouts.
 * @param persisted - Header and controls for the persisted record being edited.
 * @param newRecord - Header and controls for the corresponding new-record entry.
 * @author GPT-6 Luna
 */
export async function expectManagementMetadataHeaderLayout(
  persisted: MetadataHeaderControls,
  newRecord: MetadataHeaderControls,
): Promise<void> {
  const [persistedHeader, persistedText, persistedDate, persistedToggle,
    newHeader, newText, newDate, newToggle] = await Promise.all([
    persisted.header.boundingBox(),
    persisted.textInput.boundingBox(),
    persisted.dateInput.boundingBox(),
    persisted.toggle.boundingBox(),
    newRecord.header.boundingBox(),
    newRecord.textInput.boundingBox(),
    newRecord.dateInput.boundingBox(),
    newRecord.toggle.boundingBox(),
  ]);

  if(!persistedHeader || !persistedText || !persistedDate || !persistedToggle
    || !newHeader || !newText || !newDate || !newToggle) {
    throw new Error('Metadata header geometry requires all compared controls to be rendered.');
  }

  expect(Math.abs(persistedText.width - newText.width)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedDate.width - newDate.width)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedToggle.width - newToggle.width)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedHeader.height - newHeader.height)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedText.height - newText.height)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedDate.height - newDate.height)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedToggle.height - newToggle.height)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedText.x - persistedHeader.x - newText.x + newHeader.x)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedDate.x - persistedHeader.x - newDate.x + newHeader.x)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedToggle.x - persistedHeader.x - newToggle.x + newHeader.x)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedText.y - persistedHeader.y - newText.y + newHeader.y)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedDate.y - persistedHeader.y - newDate.y + newHeader.y)).toBeLessThanOrEqual(1);
  expect(Math.abs(persistedToggle.y - persistedHeader.y - newToggle.y + newHeader.y)).toBeLessThanOrEqual(1);

  assertContained(persistedHeader, [persistedText, persistedDate, persistedToggle]);
  assertContained(newHeader, [newText, newDate, newToggle]);
  expect(areOrderedWithoutOverlap(persistedText, persistedDate)).toBe(true);
  expect(areOrderedWithoutOverlap(newText, newDate)).toBe(true);
  expect(areOrderedWithoutOverlap(persistedText, persistedToggle)).toBe(true);
  expect(areOrderedWithoutOverlap(persistedDate, persistedToggle)).toBe(true);
  expect(areOrderedWithoutOverlap(newText, newToggle)).toBe(true);
  expect(areOrderedWithoutOverlap(newDate, newToggle)).toBe(true);
}

/** Verifies a read-only pencil is vertically centered with its adjacent label.
 * @param pencil - Visible edit control in the persisted header.
 * @param label - Visible read-only metadata label beside the pencil.
 * @example `await expectManagementMetadataPencilAlignment(pencil, label)` checks resting alignment.
 * @author GPT-6 Luna
 */
export async function expectManagementMetadataPencilAlignment(
  pencil: Locator,
  label: Locator,
): Promise<void> {
  const [pencilBox, labelBox] = await Promise.all([pencil.boundingBox(), label.boundingBox()]);

  if(!pencilBox || !labelBox) {
    throw new Error('Pencil alignment requires the pencil and read-only label to be rendered.');
  }

  const pencilCenter = pencilBox.y + pencilBox.height / 2;
  const labelCenter = labelBox.y + labelBox.height / 2;
  expect(Math.abs(pencilCenter - labelCenter)).toBeLessThanOrEqual(1);
}

/** Checks that rendered controls fit inside the owning accordion header.
 * @author GPT-6 Luna
 */
function assertContained(header: BoundingBox, controls: BoundingBox[]): void {
  controls.forEach(control => {
    expect(control.x).toBeGreaterThanOrEqual(header.x - 1);
    expect(control.x + control.width).toBeLessThanOrEqual(header.x + header.width + 1);
    expect(control.y).toBeGreaterThanOrEqual(header.y - 1);
    expect(control.y + control.height).toBeLessThanOrEqual(header.y + header.height + 1);
  });
}

/** Returns whether two header controls are separated horizontally or vertically.
 * @author GPT-6 Luna
 */
function areOrderedWithoutOverlap(first: BoundingBox, second: BoundingBox): boolean {
  return first.x + first.width <= second.x + 1 || first.y + first.height <= second.y + 1;
}

type BoundingBox = NonNullable<Awaited<ReturnType<Locator['boundingBox']>>>;
