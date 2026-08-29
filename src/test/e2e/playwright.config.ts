/**
 * Configures isolated browser end-to-end testing for the Open Asset Allocator.
 *
 * Set BASE_URL to either the local split frontend origin or CI monolith origin.
 * For example: BASE_URL=http://frontend:8000 npm test.
 *
 * Authored by: OpenCode
 */
import { defineConfig, devices } from '@playwright/test';
import { fileURLToPath } from 'node:url';
import { resolve } from 'node:path';

const packageDirectory = fileURLToPath(new URL('.', import.meta.url));
const defaultArtifactsDirectory = resolve(packageDirectory, '../../../target/e2e-results');

export default defineConfig({
  fullyParallel: false,
  outputDir: resolve(process.env.E2E_ARTIFACTS_DIR ?? defaultArtifactsDirectory, 'test-results'),
  reporter: [
    ['list'],
    ['html', { open: 'never', outputFolder: resolve(process.env.E2E_ARTIFACTS_DIR ?? defaultArtifactsDirectory, 'html-report') }],
    ['junit', { outputFile: resolve(process.env.E2E_ARTIFACTS_DIR ?? defaultArtifactsDirectory, 'junit.xml') }],
  ],
  retries: 0,
  testDir: './tests',
  timeout: 30_000,
  use: {
    baseURL: process.env.BASE_URL ?? 'http://frontend:8000',
    headless: true,
    screenshot: 'only-on-failure',
    trace: 'retain-on-failure',
    video: 'retain-on-failure',
  },
  workers: 1,
  projects: [
    { name: 'chromium', use: { ...devices['Desktop Chrome'] } },
    { name: 'firefox', use: { ...devices['Desktop Firefox'] } },
  ],
});
