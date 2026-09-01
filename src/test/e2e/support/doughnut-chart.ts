/**
 * Provides browser-level assertions and pointer helpers for Chart.js doughnut charts.
 *
 * Chart.js renders labels, tooltips, and cash-reserve patterns into a canvas, so these
 * helpers inspect the canvas draw state while keeping pointer coordinates relative to
 * the current visible canvas bounds.
 *
 * Authored by: OpenCode
 */
import type { Locator, Page } from '@playwright/test';
import { expect } from '@playwright/test';

/** Represents a point relative to a canvas element's CSS bounds. */
export type CanvasPoint = {
  x: number;
  y: number;
};

type CanvasTextRecorder = {
  generation: number;
  patternFillCount: number;
  texts: string[];
};

/**
 * Installs a browser-only recorder for text and pattern fills drawn by portfolio doughnut charts.
 *
 * Call this before the first `page.goto()` so Chart.js initialization and every later redraw are
 * captured. For example, `await installCanvasTextRecorder(page)` can be followed by
 * `await expectLatestCanvasTextSet(page, canvas, ['BONDS', 'STOCKS', '40%', '60%'])`.
 */
export async function installCanvasTextRecorder(page: Page): Promise<void> {
  await page.addInitScript(() => {
    const originalClearRect = CanvasRenderingContext2D.prototype.clearRect;
    const originalCreatePattern = CanvasRenderingContext2D.prototype.createPattern;
    const originalFillText = CanvasRenderingContext2D.prototype.fillText;
    const chartIdPrefixes = ['portfolio-chart-', 'allocation-plan-chart-'];
    const patterns = new WeakSet<CanvasPattern>();

    const getRecorder = (): Record<string, CanvasTextRecorder> => {
      const windowWithRecorder = window as Window & {
        __e2eCanvasTextRecorder?: Record<string, CanvasTextRecorder>;
      };
      windowWithRecorder.__e2eCanvasTextRecorder ??= {};
      return windowWithRecorder.__e2eCanvasTextRecorder;
    };

    const isPortfolioChartCanvas = (context: CanvasRenderingContext2D): boolean => {
      return chartIdPrefixes.some((prefix) => context.canvas.id.startsWith(prefix));
    };

    const recordPatternFill = (context: CanvasRenderingContext2D): void => {
      if (!isPortfolioChartCanvas(context) || !patterns.has(context.fillStyle as CanvasPattern)) {
        return;
      }

      const recorder = getRecorder();
      const current = recorder[context.canvas.id] ?? { generation: 0, patternFillCount: 0, texts: [] };
      current.patternFillCount += 1;
      recorder[context.canvas.id] = current;
    };

    CanvasRenderingContext2D.prototype.clearRect = function(x, y, width, height) {
      if (isPortfolioChartCanvas(this)) {
        const recorder = getRecorder();
        const current = recorder[this.canvas.id];
        recorder[this.canvas.id] = {
          generation: (current?.generation ?? 0) + 1,
          patternFillCount: 0,
          texts: [],
        };
      }
      return originalClearRect.call(this, x, y, width, height);
    };

    CanvasRenderingContext2D.prototype.createPattern = function(image, repetition) {
      const createdPattern = originalCreatePattern.call(this, image, repetition);
      if (createdPattern) {
        patterns.add(createdPattern);
      }
      return createdPattern;
    };

    const fillStyleDescriptor = Object.getOwnPropertyDescriptor(CanvasRenderingContext2D.prototype, 'fillStyle');
    if (fillStyleDescriptor?.get && fillStyleDescriptor.set) {
      Object.defineProperty(CanvasRenderingContext2D.prototype, 'fillStyle', {
        configurable: fillStyleDescriptor.configurable,
        enumerable: fillStyleDescriptor.enumerable,
        get() {
          return fillStyleDescriptor.get?.call(this);
        },
        set(value: string | CanvasGradient | CanvasPattern) {
          fillStyleDescriptor.set?.call(this, value);
          recordPatternFill(this);
        },
      });
    }

    CanvasRenderingContext2D.prototype.fillText = function(text, x, y, maxWidth) {
      if (isPortfolioChartCanvas(this)) {
        const recorder = getRecorder();
        const current = recorder[this.canvas.id] ?? { generation: 0, patternFillCount: 0, texts: [] };
        current.texts.push(String(text));
        recorder[this.canvas.id] = current;
      }
      return originalFillText.call(this, text, x, y, maxWidth);
    };
  });
}

/**
 * Asserts that the latest chart render contains exactly the expected labels and percentages.
 *
 * For example, `await expectLatestCanvasTextSet(page, canvas, ['STOCKS', 'BONDS', '60%', '40%'])`
 * waits for a completed Chart.js render and ignores draw-call ordering.
 */
export async function expectLatestCanvasTextSet(
  page: Page,
  canvas: Locator,
  expectedTexts: readonly string[],
): Promise<void> {
  await expect.poll(async () => {
    const actualTexts = await getLatestCanvasTexts(canvas);
    if (actualTexts.length === expectedTexts.length && expectedTexts.every((text) => actualTexts.includes(text))) {
      return true;
    }
    throw new Error(`Unexpected canvas text: ${JSON.stringify(actualTexts)}`);
  }).toBe(true);
}

/**
 * Asserts that the latest chart render contains the expected text values.
 *
 * This is useful after a Chart.js interaction when the old tooltip text may remain in the
 * current canvas generation. For example, pass `['BOND-B', 'BOND-A', '75%', '25%']` after
 * drilling into a BONDS slice.
 */
export async function expectLatestCanvasTextContains(
  page: Page,
  canvas: Locator,
  expectedTexts: readonly string[],
): Promise<void> {
  await expect.poll(async () => {
    const actualTexts = await getLatestCanvasTexts(canvas);
    return expectedTexts.every((text) => actualTexts.includes(text));
  }).toBe(true);
}

/**
 * Asserts whether the latest chart render applied a Patternomaly pattern to any slice.
 *
 * Use `false` for ordinary class-level or non-cash charts and `true` after rendering a
 * cash-reserve allocation. The assertion intentionally checks presence rather than an exact
 * fill count because responsive Chart.js animations redraw the canvas multiple times.
 */
export async function expectLatestCanvasPatternState(
  page: Page,
  canvas: Locator,
  expectedPattern: boolean,
): Promise<void> {
  await expect.poll(async () => {
    const patternFillCount = await getLatestCanvasPatternFillCount(canvas);
    return expectedPattern ? patternFillCount > 0 : patternFillCount === 0;
  }).toBe(true);
}

/**
 * Hovers a doughnut slice and verifies its rendered tooltip text.
 *
 * `precedingValue`, `value`, and `total` identify the slice geometrically. For example,
 * `expectChartTooltip(page, canvas, 0, 0.6, 1, ['STOCKS', '60%'])` targets the first 60% slice.
 */
export async function expectChartTooltip(
  page: Page,
  canvas: Locator,
  precedingValue: number,
  value: number,
  total: number,
  expectedTexts: readonly string[],
): Promise<void> {
  await canvas.scrollIntoViewIfNeeded();
  const point = await getDoughnutSlicePoint(canvas, precedingValue, value, total);
  const bounds = await canvas.boundingBox();
  if (!bounds) {
    throw new Error('Portfolio chart canvas has no bounding box');
  }
  await page.mouse.move(bounds.x + point.x, bounds.y + point.y);
  await expectLatestCanvasTextContains(page, canvas, expectedTexts);
}

/**
 * Clicks a point calculated from the current rendered canvas geometry.
 *
 * Existing hover state is cleared before the canvas is scrolled into view and the point and
 * bounding box are calculated, avoiding stale coordinates after a tooltip or chart redraw.
 * For example, pass
 * `() => getDoughnutCenterPoint(canvas)` to navigate back to the parent chart level.
 */
export async function clickCanvasPoint(
  page: Page,
  canvas: Locator,
  getPoint: () => Promise<CanvasPoint>,
): Promise<void> {
  await page.mouse.move(0, 0);
  await canvas.scrollIntoViewIfNeeded();
  const point = await getPoint();
  const bounds = await canvas.boundingBox();
  if (!bounds) {
    throw new Error('Portfolio chart canvas has no bounding box');
  }
  await page.mouse.click(bounds.x + point.x, bounds.y + point.y);
}

/**
 * Finds the center of the rendered doughnut from its bright connected pixel components.
 *
 * This excludes the right-side legend and adapts to responsive canvas dimensions, which makes
 * the returned point suitable for `clickCanvasPoint` in both Chromium and Firefox.
 */
export async function getDoughnutCenterPoint(canvas: Locator): Promise<CanvasPoint> {
  return canvas.evaluate((element) => {
    const canvasElement = element as HTMLCanvasElement;
    const context = canvasElement.getContext('2d');
    if (!context) {
      throw new Error('Portfolio chart canvas has no 2D context');
    }

    const pixels = context.getImageData(0, 0, canvasElement.width, canvasElement.height);
    const width = canvasElement.width;
    const height = canvasElement.height;
    const isChartPixel = (pixelIndex: number): boolean => {
      const red = pixels.data[pixelIndex];
      const green = pixels.data[pixelIndex + 1];
      const blue = pixels.data[pixelIndex + 2];
      const brightness = (red + green + blue) / 3;
      const backgroundDistance = Math.abs(red - 33) + Math.abs(green - 37) + Math.abs(blue - 41);
      return pixels.data[pixelIndex + 3] > 100 && brightness > 80 && backgroundDistance > 60;
    };
    const visited = new Uint8Array(width * height);
    const components: Array<{ size: number; minX: number; maxX: number; minY: number; maxY: number }> = [];

    for (let start = 0; start < width * height; start++) {
      if (visited[start] || !isChartPixel(start * 4)) {
        continue;
      }

      const queue = [start];
      visited[start] = 1;
      let queueIndex = 0;
      let minX = width;
      let maxX = 0;
      let minY = height;
      let maxY = 0;

      while (queueIndex < queue.length) {
        const current = queue[queueIndex++];
        const x = current % width;
        const y = Math.floor(current / width);
        minX = Math.min(minX, x);
        maxX = Math.max(maxX, x);
        minY = Math.min(minY, y);
        maxY = Math.max(maxY, y);

        const neighbours = [current - 1, current + 1, current - width, current + width];
        for (const neighbour of neighbours) {
          if (neighbour < 0 || neighbour >= width * height || visited[neighbour]) {
            continue;
          }
          const neighbourX = neighbour % width;
          if (Math.abs(neighbourX - x) > 1 || !isChartPixel(neighbour * 4)) {
            continue;
          }
          visited[neighbour] = 1;
          queue.push(neighbour);
        }
      }

      if (queue.length > 1_000) {
        components.push({ size: queue.length, minX, maxX, minY, maxY });
      }
    }

    components.sort((left, right) => right.size - left.size);
    if (components.length === 0) {
      throw new Error('Could not locate the rendered doughnut components');
    }

    const chartComponents = components.slice(0, 4);
    const minX = Math.min(...chartComponents.map(({ minX: componentMinX }) => componentMinX));
    const maxX = Math.max(...chartComponents.map(({ maxX: componentMaxX }) => componentMaxX));
    const minY = Math.min(...chartComponents.map(({ minY: componentMinY }) => componentMinY));
    const maxY = Math.max(...chartComponents.map(({ maxY: componentMaxY }) => componentMaxY));
    const centerX = (minX + maxX) / 2;
    const centerY = (minY + maxY) / 2;
    const bounds = canvasElement.getBoundingClientRect();
    return {
      x: centerX * bounds.width / width,
      y: centerY * bounds.height / height,
    };
  });
}

/** Finds a point halfway through a doughnut slice using its rendered annulus. */
export async function getDoughnutSlicePoint(
  canvas: Locator,
  precedingValue: number,
  value: number,
  total: number,
): Promise<CanvasPoint> {
  const centerPoint = await getDoughnutCenterPoint(canvas);
  return canvas.evaluate((element, data) => {
    const canvasElement = element as HTMLCanvasElement;
    const context = canvasElement.getContext('2d');
    if (!context) {
      throw new Error('Portfolio chart canvas has no 2D context');
    }

    const pixels = context.getImageData(0, 0, canvasElement.width, canvasElement.height);
    const width = canvasElement.width;
    const height = canvasElement.height;
    const bounds = canvasElement.getBoundingClientRect();
    const center = {
      x: data.centerX * width / bounds.width,
      y: data.centerY * height / bounds.height,
    };
    const angle = -Math.PI / 2 + 2 * Math.PI * (data.precedingValue + data.value / 2) / data.total;
    const directionX = Math.cos(angle);
    const directionY = Math.sin(angle);
    let firstPaintedRadius = -1;
    let lastPaintedRadius = -1;

    for (let radius = 0; radius < Math.max(width, height); radius++) {
      const x = Math.round(center.x + directionX * radius);
      const y = Math.round(center.y + directionY * radius);
      if (x < 0 || y < 0 || x >= width || y >= height) {
        break;
      }
      const alpha = pixels.data[(y * width + x) * 4 + 3];
      if (alpha > 20) {
        firstPaintedRadius = firstPaintedRadius < 0 ? radius : firstPaintedRadius;
        lastPaintedRadius = radius;
      } else if (firstPaintedRadius >= 0 && radius > firstPaintedRadius + 5) {
        break;
      }
    }

    if (firstPaintedRadius < 0 || lastPaintedRadius < 0) {
      throw new Error(`Could not locate a rendered doughnut slice: ${JSON.stringify({
        center,
        canvas: { width, height },
        bounds,
        angle,
      })}`);
    }

    const radius = (firstPaintedRadius + lastPaintedRadius) / 2;
    return {
      x: (center.x + directionX * radius) * bounds.width / width,
      y: (center.y + directionY * radius) * bounds.height / height,
    };
  }, { centerX: centerPoint.x, centerY: centerPoint.y, precedingValue, total, value });
}

/** Reads the latest non-empty canvas text values recorded for a chart. */
async function getLatestCanvasTexts(canvas: Locator): Promise<readonly string[]> {
  return canvas.evaluate((element) => {
    const recorder = (window as Window & {
      __e2eCanvasTextRecorder?: Record<string, CanvasTextRecorder>;
    }).__e2eCanvasTextRecorder;
    return [...new Set((recorder?.[element.id]?.texts ?? []).filter((text) => text.length > 0))];
  });
}

/** Reads the number of pattern-backed fills in the latest canvas render. */
async function getLatestCanvasPatternFillCount(canvas: Locator): Promise<number> {
  return canvas.evaluate((element) => {
    const recorder = (window as Window & {
      __e2eCanvasTextRecorder?: Record<string, CanvasTextRecorder>;
    }).__e2eCanvasTextRecorder;
    return recorder?.[element.id]?.patternFillCount ?? 0;
  });
}
