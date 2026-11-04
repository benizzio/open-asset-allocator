/**
 * Verifies that feature-owned HTML fragments are served at their nested URLs.
 * The same contract runs against the development frontend and production monolith.
 * @author GPT-6.1 Terra
 */
import { expect, test } from '../support/fixtures';

test('scenario 22: serves nested portfolio component templates', async ({ page, request }) => {
  const templates = [
    ['/websrc/components/allocation-plan/allocation-plan.html', 'accordion-allocation-plan'],
    ['/websrc/components/allocation-plan/allocation-plan-management.html', 'accordion-allocation-plan-management'],
    ['/websrc/components/portfolio-history/portfolio-history.html', 'accordion-portfolio-history'],
    ['/websrc/components/portfolio-history/portfolio-history-management.html', 'accordion-portfolio-history-management'],
  ] as const;

  for (const [url, marker] of templates) {
    const response = await request.get(url, { headers: { 'HX-Request': 'true' } });
    expect(response.status(), url).toBe(200);
    const html = await response.text();
    const hasMarker = await page.evaluate(({ html, marker }) => {
      const document = new DOMParser().parseFromString(html, 'text/html');
      return document.getElementById(marker) !== null;
    }, { html, marker });
    expect(hasMarker, url).toBe(true);
  }
});
