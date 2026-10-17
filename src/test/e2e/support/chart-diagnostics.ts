/**
 * Provides opt-in failure diagnostics for charts rendered during HTMX navigation.
 *
 * Records native resize deliveries, canvas identity and size changes, and DOM/HTMX
 * lifecycle events. Layout measurements are deferred until failure collection to
 * avoid forcing layout during the timing-sensitive interactions under investigation.
 *
 * @author GPT-6 Astra
 */
import type { ConsoleMessage } from '@playwright/test';
import { test } from './fixtures';

/** Browser-only diagnostic hooks and the existing canvas text recorder.
 * @author GPT-6 Astra
 */
type ChartDiagnosticWindow = Window & {
  __e2eChartDiagnostics?: () => unknown;
  __e2eCanvasTextRecorder?: unknown;
};

/**
 * Runs a database-isolated test with chart diagnostics attached on unexpected failure.
 *
 * Uses the regular page fixture and collects before that page is torn down. The JSON
 * attachment is included in the existing Playwright report and CI artifact upload.
 *
 * @example
 * chartDiagnosticTest('scenario 6: renders allocation charts', async ({ page }) => {
 *   await page.goto('/portfolio/1/allocation');
 * });
 * @author GPT-6 Astra
 */
export const chartDiagnosticTest = test.extend({
  /** Installs diagnostics before navigation and attaches them before page teardown.
   * @author GPT-6 Astra
   */
  page: async ({ page }, use, testInfo) => {
    const errors: Array<{ time: number; kind: string; message: string }> = [];
    let droppedErrors = 0;

    /** Bounds browser error history independently of the in-page event buffer.
     * @author GPT-6 Astra
     */
    const recordError = (kind: string, message: string): void => {
      if (errors.length === 100) {
        errors.shift();
        droppedErrors++;
      }
      errors.push({ time: Date.now(), kind, message });
    };
    const onPageError = (error: Error): void => recordError('pageerror', error.stack ?? error.message);
    const onConsole = (message: ConsoleMessage): void => {
      if (message.type() === 'error') {
        recordError('console', message.text());
      }
    };
    page.on('pageerror', onPageError);
    page.on('console', onConsole);

    try {
      await page.addInitScript(installChartDiagnostics);
      await use(page);
    } finally {
      if (testInfo.status !== testInfo.expectedStatus) {
        let browserState: unknown;
        try {
          browserState = await page.evaluate(() => {
            const capture = (window as ChartDiagnosticWindow).__e2eChartDiagnostics;
            return capture ? capture() : { collectionError: 'Chart diagnostics were not installed in this document.' };
          });
        } catch (error) {
          browserState = { collectionError: String(error) };
        }

        try {
          await testInfo.attach('chart-diagnostics', {
            contentType: 'application/json',
            body: JSON.stringify({
              test: testInfo.title,
              project: testInfo.project.name,
              retry: testInfo.retry,
              status: testInfo.status,
              url: page.url(),
              errors,
              droppedErrors,
              browserState,
            }, null, 2),
          });
        } catch (error) {
          // Keep the original test failure even if artifact collection also fails.
          console.error('Could not attach chart diagnostics:', error);
        }
      }
      page.off('pageerror', onPageError);
      page.off('console', onConsole);
    }
  },
});

/** Installs a self-contained, bounded recorder inside each new browser document.
 * @author GPT-6 Astra
 */
function installChartDiagnostics(): void {
  const diagnosticWindow = window as ChartDiagnosticWindow;
  if (diagnosticWindow.__e2eChartDiagnostics) {
    return;
  }

  const selector = 'canvas[id^="allocation-plan-chart-"], canvas[id^="portfolio-chart-"]';
  const events: unknown[] = [];
  const nodeIds = new WeakMap<Element, number>();
  let nextNodeId = 0;
  let nextObserverId = 0;
  let droppedEvents = 0;

  /** Records timing without reading computed style or layout.
   * @author GPT-6 Astra
   */
  const record = (kind: string, detail: unknown): void => {
    if (events.length === 1_000) {
      events.shift();
      droppedEvents++;
    }
    events.push({ time: performance.now(), route: location.pathname, kind, detail });
  };

  /** Identifies DOM instances separately from their reusable HTML IDs.
   * @author GPT-6 Astra
   */
  const describe = (element: Element) => {
    if (!nodeIds.has(element)) {
      nodeIds.set(element, ++nextNodeId);
    }
    return {
      nodeId: nodeIds.get(element),
      tag: element.tagName,
      id: element.id,
      connected: element.isConnected,
      class: element.getAttribute('class'),
      style: element.getAttribute('style'),
      hidden: element.hasAttribute('hidden'),
    };
  };

  /** Finds chart canvases inside a mutation or resize target, including the target itself.
   * @author GPT-6 Astra
   */
  const chartsWithin = (node: Node): Element[] => {
    if (!(node instanceof Element)) {
      return [];
    }
    return [...(node.matches(selector) ? [node] : []), ...node.querySelectorAll(selector)];
  };

  /** Measures current layout only when the failed test collects its attachment.
   * @author GPT-6 Astra
   */
  const measure = (element: Element) => {
    const style = getComputedStyle(element);
    return {
      ...describe(element),
      bounds: element.getBoundingClientRect().toJSON(),
      clientWidth: element.clientWidth,
      clientHeight: element.clientHeight,
      display: style.display,
      visibility: style.visibility,
      position: style.position,
      width: style.width,
      height: style.height,
      maxWidth: style.maxWidth,
      maxHeight: style.maxHeight,
      overflow: style.overflow,
    };
  };

  diagnosticWindow.__e2eChartDiagnostics = () => ({
    timeOrigin: performance.timeOrigin,
    capturedAt: performance.now(),
    visibilityState: document.visibilityState,
    viewport: { width: innerWidth, height: innerHeight, devicePixelRatio },
    scroll: { x: scrollX, y: scrollY },
    fonts: document.fonts.status,
    droppedEvents,
    events: [...events],
    canvasTextRecorder: diagnosticWindow.__e2eCanvasTextRecorder,
    canvases: [...document.querySelectorAll<HTMLCanvasElement>(selector)].map(canvas => {
      const ancestors = [];
      for (let parent = canvas.parentElement; parent; parent = parent.parentElement) {
        ancestors.push(measure(parent));
      }
      return { ...measure(canvas), buffer: { width: canvas.width, height: canvas.height }, ancestors };
    }),
  });

  for (const property of ['width', 'height'] as const) {
    const descriptor = Object.getOwnPropertyDescriptor(HTMLCanvasElement.prototype, property);
    const originalSet = descriptor?.set;
    if (!originalSet) {
      record('unsupported-canvas-property', property);
      continue;
    }
    Object.defineProperty(HTMLCanvasElement.prototype, property, {
      ...descriptor,
      /** Preserves native sizing while recording the setter's caller and DOM identity.
       * @author GPT-6 Astra
       */
      set(this: HTMLCanvasElement, value: number) {
        originalSet.call(this, value);
        if (this.matches(selector)) {
          record('canvas-size', {
            canvas: describe(this), property, value,
            width: this.width, height: this.height,
            stack: new Error().stack,
          });
        }
      },
    });
  }

  const NativeResizeObserver = window.ResizeObserver;
  /** Observes the application's existing subscriptions without adding resize observers.
   * @author GPT-6 Astra
   */
  window.ResizeObserver = class extends NativeResizeObserver {
    private readonly diagnosticId: number;

    /** Preserves native callback delivery and records its supplied content rectangles.
     * @author GPT-6 Astra
     */
    constructor(callback: ResizeObserverCallback) {
      const id = ++nextObserverId;
      super((entries, observer) => {
        for (const entry of entries) {
          const canvases = chartsWithin(entry.target);
          if (canvases.length) {
            record('resize-delivery', {
              observerId: id, target: describe(entry.target),
              contentRect: entry.contentRect.toJSON(), canvases: canvases.map(describe),
            });
          }
        }
        callback.call(observer, entries, observer);
      });
      this.diagnosticId = id;
    }

    /** Records native observation registration for chart containers.
     * @author GPT-6 Astra
     */
    override observe(target: Element, options?: ResizeObserverOptions): void {
      super.observe(target, options);
      if (chartsWithin(target).length) {
        record('resize-observe', { observerId: this.diagnosticId, target: describe(target) });
      }
    }

    /** Records removal of individual subscriptions.
     * @author GPT-6 Astra
     */
    override unobserve(target: Element): void {
      super.unobserve(target);
      record('resize-unobserve', { observerId: this.diagnosticId, target: describe(target) });
    }

    /** Records observer disconnection, including detached chart cleanup.
     * @author GPT-6 Astra
     */
    override disconnect(): void {
      super.disconnect();
      record('resize-disconnect', { observerId: this.diagnosticId });
    }
  };

  new MutationObserver(mutations => {
    for (const mutation of mutations) {
      if (mutation.type === 'attributes' && chartsWithin(mutation.target).length) {
        const target = mutation.target as Element;
        record('attribute', {
          target: describe(target), name: mutation.attributeName, oldValue: mutation.oldValue,
          value: target.getAttribute(mutation.attributeName!),
        });
      } else if (mutation.type === 'childList') {
        for (const [kind, nodes] of [['canvas-added', mutation.addedNodes], ['canvas-removed', mutation.removedNodes]] as const) {
          for (const node of nodes) {
            for (const canvas of chartsWithin(node)) {
              record(kind, describe(canvas));
            }
          }
        }
      }
    }
  }).observe(document, {
    childList: true, subtree: true, attributes: true, attributeOldValue: true,
    attributeFilter: ['style', 'class', 'hidden', 'width', 'height'],
  });

  for (const name of ['htmx:beforeSwap', 'htmx:afterSwap', 'htmx:afterSettle', 'htmx:beforeCleanupElement']) {
    document.addEventListener(name, event => {
      const target = event.target;
      // Cleanup fires for every removed node; retain only chart-related cleanup.
      if (name === 'htmx:beforeCleanupElement' && (!(target instanceof Node) || !chartsWithin(target).length)) {
        return;
      }
      record(name, target instanceof Element ? describe(target) : null);
    }, true);
  }
  record('installed', { userAgent: navigator.userAgent });
}
