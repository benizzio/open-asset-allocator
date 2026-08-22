# Best maintained and community-adopted E2E testing tool for the Open Asset Allocator project

Generated from 14 structured research results.

## Table of Contents

1. [Playwright Test](#playwright-test) - Latest Release: v1.62.1, published July 30, 2026. | npm Downloads: @playwright/test received 46,346,538 downloads during the exact seven-day npm window 2026-08-14 through 2026-08-20.
2. [Cypress](#cypress) - Latest Release: 15.21.0, released 2026-08-18. | npm Downloads: The npm downloads API reported 6,155,552 downloads of the exact cypress package for the complete seven-day window 2026-08-14 through 2026-08-20.
3. [WebdriverIO](#webdriverio) - Latest Release: WebdriverIO v9.31.2, published 2026-08-21. | npm Downloads: The official npm downloads API recorded 1,544,931 downloads of @wdio/cli for the complete seven-day window 2026-08-15 through 2026-08-21, and 6,463,515 for 2026-07-22 through 20...
4. [Selenium WebDriver with a JavaScript test runner](#selenium-webdriver-with-a-javascript-test-runner) - Latest Release: selenium-webdriver 4.47.0, released 2026-08-10. | npm Downloads: The official npm downloads API reported 1,820,837 direct downloads of selenium-webdriver for 2026-08-14 through 2026-08-20.
5. [Nightwatch.js](#nightwatch-js) - Latest Release: 3.16.0, published 2026-05-25 on npm and GitHub. | npm Downloads: The npm downloads API recorded 137,213 downloads for the nightwatch package from 2026-08-15 through 2026-08-21 and 642,513 downloads from 2026-07-22 through 2026-08-21.
6. [TestCafe](#testcafe) - Latest Release: 3.7.6, published on 2026-07-07 to npm and GitHub. | npm Downloads: The npm downloads API recorded 176,264 downloads for the exact seven-day window 2026-08-14 through 2026-08-20 and 904,319 downloads for 2026-07-22 through 2026-08-21 for the tes...
7. [Puppeteer with Vitest or Jest](#puppeteer-with-vitest-or-jest) - Latest Release: Puppeteer 25.8.0 was released on 2026-08-17. | npm Downloads: For the exact npm window 2026-08-14 through 2026-08-20, puppeteer received 9,775,888 downloads, vitest received 78,083,394, and jest received 38,752,856.
8. [Rod with Go testing](#rod-with-go-testing) - Latest Release: v0.116.2, published July 12, 2024. | npm Downloads: Not applicable.
9. [chromedp with Go testing](#chromedp-with-go-testing) - Latest Release: v0.16.0, the latest non-prerelease Go module tag, published July 14, 2026. | npm Downloads: Not applicable. chromedp is distributed as the Go module github.com/chromedp/chromedp and has no official npm package, so an npm download count would be misleading and is record...
10. [CodeceptJS](#codeceptjs) - Latest Release: 4.1.0, published 2026-07-30 on GitHub and npm. | npm Downloads: The npm downloads API recorded 572,087 downloads of codeceptjs from 2026-08-15 through 2026-08-21 and 2,355,761 downloads from 2026-07-22 through 2026-08-21.
11. [Testplane](#testplane) - Latest Release: Testplane 9.1.1, published to npm on 2026-08-21 from commit 9828757. | npm Downloads: The npm Downloads API recorded 14,620 downloads of the exact `testplane` package for the complete seven-day window 2026-08-15 through 2026-08-21 and 60,812 downloads for 2026-07...
12. [Cucumber.js with Playwright or Playwright-BDD](#cucumber-js-with-playwright-or-playwright-bdd) - Latest Release: Observed 2026-08-22: @cucumber/cucumber 13.2.1, released 2026-08-04; playwright-bdd 9.2.0, released 2026-06-18; and upstream @playwright/test 1.62.1, released 2026-07-30. playwr... | npm Downloads: For the complete seven-day window 2026-08-15 through 2026-08-21, the npm API recorded 2,457,586 downloads of @cucumber/cucumber and 540,138 downloads of playwright-bdd.
13. [Robot Framework Browser](#robot-framework-browser) - Latest Release: 20.4.0, published on August 19, 2026. | npm Downloads: Not applicable as an npm adoption metric: users install the E2E framework as the PyPI package `robotframework-browser`, while npm packages are internal wrapper dependencies.
14. [Grafana k6 Browser](#grafana-k6-browser) - Latest Release: v2.2.0, published August 10, 2026, was the latest principal stable release observed on August 22, 2026. | npm Downloads: Not applicable.

## Detailed Results

<a id="playwright-test"></a>
## 1. Playwright Test

Source result: `Playwright_Test.json`

### Project And Compatibility

#### Implementation Language

> Playwright and Playwright Test are primarily implemented in TypeScript.<br>
> Tests can be authored in TypeScript or JavaScript and run on Node.js.<br>
> The evaluated npm package is @playwright/test.

#### Operating System Support

> Officially supported Linux targets include Debian 12/13 and Ubuntu 22.04/24.04/26.04 on x86-64 and arm64.<br>
> These cover local supported Linux installations and GitHub-hosted Ubuntu x86-64 runners.<br>
> Official Ubuntu-based Docker images are available.<br>
> Firefox and WebKit browser builds require glibc, so Alpine and other musl distributions are not supported for those engines.

#### License And Governance

> Apache License 2.0.<br>
> The project is owned and commercially stewarded by Microsoft in the microsoft/playwright repository, developed in public, and accepts community contributions under repository contribution rules.<br>
> The permissive license allows use, modification, and redistribution in this repository subject to Apache-2.0 notice and patent terms; no paid license is required.

#### Installation Model

> Add @playwright/test as a development dependency, normally with npm install -D @playwright/test, then install version-matched browsers with npx playwright install or a selected engine.<br>
> On Linux CI, npx playwright install --with-deps installs browser binaries and operating-system libraries.<br>
> Versioned Microsoft Container Registry images such as mcr.microsoft.com/playwright:v1.62.0-noble include browsers and system dependencies but not the project's test package.<br>
> Browser binaries are large external artifacts stored in Playwright's cache unless a custom PLAYWRIGHT_BROWSERS_PATH is used.

#### Candidate Scope And Layer

> Complete E2E framework: integrated test runner, assertions, fixtures, isolation, parallel workers, browser management, API requests, reports, retries, traces, screenshots, video, code generation, and debugging tools.

#### Authoring And Async Model

> Native TypeScript/JavaScript async and await with test, expect, fixtures, hooks, projects, and configuration APIs.<br>
> Browser operations return promises; there is no hidden command queue or required Gherkin/keyword DSL.<br>
> Web-first locator assertions retry asynchronously and must be awaited.

#### Build Pipeline Coupling

> Black-box tests can target either the production image built by Parcel and served with the Gin backend or the Parcel development server.<br>
> Playwright does not require Vite, application source transformation, browser instrumentation, or replacement of Parcel.<br>
> In production, this repository serves the frontend and /api from one Gin origin on port 80.<br>
> In development, Parcel serves port 8000 and .proxyrc.js proxies /api to Gin on port 8080; baseURL can point to either arrangement.

#### Testability Instrumentation Required

> No production instrumentation, special route, injected script, relaxed content-security policy, or build hook is required.<br>
> Existing semantic roles, labels, visible text, IDs, and attributes can be used.<br>
> Adding stable data-testid attributes is optional where the current dynamic templates do not expose a unique user-facing locator.<br>
> Test-only API fixture/reset endpoints would be optional custom application work, not a Playwright requirement; database setup can instead use existing APIs or direct fixture code outside production paths.

### Maintenance Health

#### Latest Stable Release

> v1.62.1, published July 30, 2026.<br>
> It is the npm latest version observed for @playwright/test on August 22, 2026.

#### Release Cadence

> Active and frequent.<br>
> The ten most recent stable GitHub releases ran from v1.58.0 on January 23, 2026 through v1.62.1 on July 30, 2026, with feature releases 1.58, 1.59, 1.60, 1.61, and 1.62 plus patches.<br>
> Main-branch code commits also produce daily canary packages on the next npm tag, while beta builds are published around release branching.

#### Repository Activity

> Very high current activity.<br>
> The main branch was pushed on August 22, 2026.<br>
> The ten newest observed commits span August 20-22 and include runner reporting, tracing, network, browser-roll, accessibility, and documentation work from several maintainers and automation.<br>
> Multiple pull requests were merged within roughly a day of creation in the same period.

#### Wrapper Upstream Lag

> Not applicable as a wrapper-lag risk: Playwright Test, the Playwright library, and the managed browser revisions are released from the same repository and @playwright/test 1.62.1 pins playwright 1.62.1.<br>
> Browser engine revisions are rolled continuously on main.<br>
> This integration reduces wrapper lag but intentionally ties browser updates to Playwright package upgrades.

### Community Adoption

#### Npm Downloads

> @playwright/test received 46,346,538 downloads during the exact seven-day npm window 2026-08-14 through 2026-08-20.

#### Ecosystem Usage

> The project has official JavaScript/TypeScript, Python, Java, and .NET bindings, an official VS Code extension, official Docker images, integrations documented for major CI systems, axe-core integration guidance, and many third-party reporters and services.<br>
> The large npm and GitHub usage signals are consistent with broad production adoption, although package downloads include CI and transitive activity and do not identify organizations.

#### Community Support

> Extensive versioned official documentation covers concepts, API details, debugging, CI, migration-related release notes, and examples.<br>
> Support is available through the active GitHub issue tracker, Stack Overflow's playwright tag, and broad third-party material.<br>
> The main repository does not expose GitHub Discussions, so issue-based support and external forums carry more of the community-help load.

#### Adoption Trend

> Strongly growing on comparable npm evidence. @playwright/test rose from 10,361,283 downloads in 2025-08-14 through 2025-08-20 to 46,346,538 in 2026-08-14 through 2026-08-20, about 4.47 times the year-earlier level.<br>
> The repository also had 94,936 stars and same-day development activity.<br>
> These are adoption proxies, not counts of active E2E teams.

#### Adoption Metric Normalization

> The npm metric is the package @playwright/test, all versions, registry-wide, for the exact seven-day period 2026-08-14 through 2026-08-20.<br>
> It directly represents installation traffic for the integrated runner but can include repeated CI installs, cache misses, bots, mirrors, and transitive dependency installs; it is not unique users.<br>
> The year-over-year comparison uses the identical package and a matching seven-day calendar window.<br>
> GitHub metrics cover the broader Playwright monorepo, including language bindings and automation use beyond Playwright Test E2E suites.

### Browser And Runtime Coverage

#### Browser Engines

> Managed Chromium, patched Firefox, and patched WebKit are first-class projects.<br>
> Chromium mode can also run installed Google Chrome and Microsoft Edge channels, including stable, beta, dev, and canary where available; branded browsers are not installed by default.<br>
> Playwright does not drive stock Firefox or Safari because its Firefox and WebKit support relies on patches.<br>
> Electron support exists but is not relevant to this web application.

#### Browser Protocol

> Normal operation uses Playwright's own high-fidelity client/server protocol and patched browser integrations rather than W3C WebDriver or WebDriver BiDi.<br>
> Chromium-only CDP sessions and connectOverCDP are available, but official documentation describes CDP attachment as lower fidelity than the Playwright protocol.<br>
> This provides deep control and consistent APIs at the cost of protocol and managed-browser lock-in.

#### Headless And Headed Modes

> Headless execution is the default and fits CI.<br>
> Headed mode is supported locally with headless: false or --headed; Linux CI can use Xvfb, which is present in official images.<br>
> Debug mode launches headed browsers and the Inspector.

#### Browser Version Management

> Each Playwright release is paired with tested browser revisions downloaded by npx playwright install.<br>
> Updating @playwright/test generally requires reinstalling matching browsers.<br>
> Teams can install only selected engines, set the browser cache path, use installed Chrome/Edge channels, or pin an official image tag.<br>
> Arbitrary executablePath use is explicitly unsupported-risky.<br>
> Linux system dependencies require install --with-deps or a matching image.

#### Parallel Browser Support

> Configuration projects define browser, device, and environment matrices.<br>
> Projects run the same tests across selected engines; --project selects subsets.<br>
> Worker count controls local parallelism, --shard distributes tests across jobs, and blob reports can be merged.<br>
> Database-sensitive suites can use one worker or serial groups while independent browser-only tests retain parallel execution.

#### Mobile Emulation

> Built-in device descriptors configure viewport, screen, user agent, touch, device scale factor, mobile mode, and browser context.<br>
> Pixel and iPhone-like projects are supported for responsive checks.<br>
> This is emulation in desktop browser engines, not execution on physical Android or iOS devices.

#### Real Browser Fidelity

> Chromium projects can use managed builds or installed Chrome/Edge channels.<br>
> Firefox is a patched build near recent Firefox Stable, not branded Firefox.<br>
> WebKit is built from upstream WebKit and is not Apple Safari; Safari packaging, OS integration, codecs, and release timing differ.<br>
> Official guidance says macOS WebKit is closer to Safari for platform-sensitive features such as media, but true Safari coverage requires a separate Safari-capable system/tool.

#### Environment Determinism Controls

> Per-project or per-test controls include viewport, screen, device scale, user agent, touch/mobile mode, locale, timezone, geolocation, permissions, color scheme, contrast, forced colors, and reduced motion.<br>
> Clock controls Date, timers, animation frames, performance, and event timestamps.<br>
> Screenshot assertions disable animations by default and can mask or style volatile areas.<br>
> Playwright does not provide a general seeded replacement for Math.random or backend randomness; tests must inject or control those separately.

### Application Fit

#### Dynamic Dom Synchronization

> Strong fit for HTMX swaps and lazy Handlebars rendering.<br>
> Locators resolve an up-to-date DOM element before each action, retry if elements detach during actionability checks, and web-first assertions poll until expected state.<br>
> This directly addresses this repository's hx-get, hx-trigger, after-settle, and template replacement patterns.<br>
> Tests should assert a meaningful post-swap state or await a specific response rather than use networkidle or fixed sleeps.

#### Routing Support

> Playwright can navigate directly to deep links, assert URLs, wait for URL changes, use goBack/goForward, and observe history-driven DOM changes.<br>
> Gin's fallback serves root HTML for non-API GET routes, which supports direct Navigo deep links, while the frontend uses Navigo strategy ALL.<br>
> Route tests should use page.waitForURL or URL assertions when one action can trigger multiple history changes.<br>
> Playwright disables BFCache by default for deterministic navigation; explicitly enabling BFCache is unsupported for reliable lifecycle tracking.

#### Locator Model

> Role, label, placeholder, text, alt text, title, CSS, and configurable test-ID locators are available.<br>
> Locators are strict for single-element actions and fail on ambiguous matches, which surfaces template ambiguity.<br>
> They reacquire nodes after HTMX replacement.<br>
> Prefer roles/labels/text, with existing stable IDs or optional test IDs for dynamic rows where accessible identity is insufficient.

#### Form Interaction

> fill, clear, press, pressSequentially, focus, blur, check, selectOption, keyboard, mouse, and dispatchEvent cover realistic form flows.<br>
> Actions perform visibility, stability, event-receivability, enabled, and editable checks as applicable, and native validation can be asserted through DOM state and messages.<br>
> This can exercise the repository's input masks, hx-validate, blur/focus behavior, hidden values, and dynamically added allocation/history rows without bypassing browser behavior.

#### Canvas And Download Support

> Chart.js canvas can be validated through application-visible legends/labels and source data, page.evaluate inspection of Chart state or canvas pixels, targeted screenshots, and visual comparisons.<br>
> Trace canvas rendering is not fully reliable and current release notes say canvas trace display is disabled by default, so traces alone should not prove chart correctness.<br>
> Download events expose suggested filenames, streams, paths, failure state, and saveAs; payload contents can be parsed and asserted.<br>
> No current browser-side download implementation was found in the inspected frontend, so a future download scenario needs its concrete format defined.

#### Network And Api Access

> APIRequestContext supports REST fixture seeding and postcondition checks.<br>
> Browser request/response events, waitForResponse, page/context routing, fulfill/abort/continue/fetch, HAR recording/replay, and WebSocket inspection are available.<br>
> Yahoo Finance calls can be replaced at the browser network layer if browser-originated, or through a mock service/configured backend fixture if Gin performs them server-side.<br>
> Existing Go integration-test mock patterns can remain separate.

#### Same Origin Support

> Direct support with no CORS change.<br>
> Production can be tested through the consolidated Gin origin where static files and /api share host and port.<br>
> Development can target Parcel at port 8000, whose proxy forwards /api to localhost:8080.<br>
> A baseURL per project selects either topology.

#### Test Isolation

> Every test receives a fresh BrowserContext equivalent to a new browser profile, isolating cookies, local storage, IndexedDB, permissions, and pages.<br>
> API request contexts can also have isolated cookie stores.<br>
> Database state is external to browser isolation; use unique IDs per worker, API/direct-database cleanup fixtures, a disposable PostgreSQL instance, or workers: 1 for shared mutable fixtures.<br>
> Fixture teardown should restore state regardless of test outcome.

#### External Server Model

> Tests can run against any independently started application by setting baseURL from configuration or environment. webServer is optional, so local port 80 production Compose, Parcel port 8000, a CI service URL, or an externally managed environment can use the same black-box tests.

#### Application Lifecycle

> webServer can launch one or multiple commands, poll a URL for readiness, apply startup timeouts, reuse a local server, and terminate its process group.<br>
> Project dependencies or setup/teardown fixtures can run builds, Docker Compose, migrations, fixture setup, and cleanup with report/trace visibility.<br>
> For this repository, a small custom wrapper is still needed because build.sh, start.sh, Flyway completion, PostgreSQL disposal, and destroy.sh are multi-service concerns; use explicit health/readiness checks and gracefulShutdown SIGTERM for Docker-related commands.

#### Visual Regression Workflow

> Core screenshot baselines are generated on first run, committed, reviewed, and updated with --update-snapshots.<br>
> Per-project/platform paths separate browser baselines.<br>
> Assertions wait for consecutive stable screenshots and support masks, custom styles, disabled animations, threshold, maxDiffPixels, maxDiffPixelRatio, CSS/device scaling, and volatile-region suppression.<br>
> Fonts, OS, browser versions, data, viewport, time, and Chart.js animation must be pinned.<br>
> Canvas is suitable for targeted visual checks only after deterministic chart data and animation completion; semantic/data assertions should remain primary.

#### Accessibility Audit Integration

> Official guidance integrates the open-source @axe-core/playwright package through reusable fixtures and test attachments.<br>
> Core assertions cover accessible roles, names, descriptions, and ARIA YAML snapshots.<br>
> This provides open-source reporting through standard attachments/JSON/JUnit/HTML, but automated axe checks detect only a subset of accessibility issues and must be complemented by manual assessment.

### Reliability

#### Waiting Model

> Actions auto-wait for strict single-element resolution and applicable visibility, stability, receives-events, enabled, editable, and attachment conditions.<br>
> Locator assertions retry for up to their expectation timeout.<br>
> Navigation and URL waits, event waits, request/response waits, locator waits, and expect.poll/toPass support explicit domain synchronization.<br>
> Fixed waitForTimeout sleeps and blanket networkidle waits should be avoided for HTMX; assert the desired DOM or response instead.

#### Flake Controls

> Independent test, assertion, action, navigation, fixture, and global timeout scopes are configurable.<br>
> Retries can be CI-only; v1.62 adds immediate or isolated retry strategy.<br>
> Failed workers and browsers are discarded before subsequent work. --repeat-each, --last-failed, --workers=1, serial groups, trace retention modes, and failOnFlakyTests support reproduction and policy enforcement.<br>
> Retries should expose rather than hide persistent shared-database defects.

#### Isolation Model

> Fresh BrowserContext and page fixtures are created per test, while worker processes isolate JavaScript globals and each worker launches a browser.<br>
> A worker is replaced after failure.<br>
> Serial groups and beforeAll-created shared pages are available for genuinely dependent scenarios, but the official recommendation is isolated tests.<br>
> PostgreSQL must be isolated or reset separately because it sits outside the browser context.

#### Parallelism Controls

> Test files run in parallel by default while tests within a file run in declaration order. workers accepts an integer or CPU percentage; fullyParallel controls intra-file parallelism, serial mode protects dependent groups, project dependencies order setup, and shards distribute work across jobs.<br>
> For the current shared PostgreSQL topology, start with workers: 1 for state-mutating E2E tests or allocate a unique database/schema and test-data namespace from testInfo.parallelIndex before increasing workers.

#### Flake Observability

> Reports classify passed-on-first-attempt, flaky-passed-on-retry, and failed tests; failOnFlakyTests can fail CI. retry and repeatEach expose recurrence, and JSON/JUnit/blob/custom reporters can export retry status for local trend processing.<br>
> Traces can be retained from the first failed attempt or retries.<br>
> There is no built-in long-term quarantine registry or historical dashboard in the open-source runner; quarantine tags and trend storage require repository conventions, a custom reporter, or an optional external service.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Configurable screenshots, videos, execution traces, attachments, HTML reports, and HAR files are available with on, failure-only, first-failure, or retry retention modes.<br>
> Trace Viewer includes actions, before/after DOM snapshots, source, actionability logs, errors, console, and network request/response details.<br>
> Browser console and page errors can also be attached through event listeners.<br>
> Backend and Docker logs are not captured automatically.

#### Debugging Tools

> Inspector supports pause, play, step, locator editing/picking, and actionability logs.<br>
> UI Mode provides watch mode and time-travel-like exploration over trace snapshots, source, console, and network.<br>
> Headed mode, slowMo, page.pause, browser DevTools, Node debuggers, and the official VS Code extension support interactive and breakpoint debugging.

#### Test Generation

> Official codegen records clicks and fills, creates assertions, and generates role/text/test-id oriented locators through the Inspector or VS Code.<br>
> Generated code is a starting point; dynamic HTMX waits, domain fixture setup, cleanup, and reusable page abstractions still require review and maintenance.

#### Reporters

> Built-in list, line, dot, HTML, blob, JSON, JUnit, GitHub, and null reporters can be combined.<br>
> Blob output supports shard merging; GitHub output creates annotations; custom reporters implement the reporter API.<br>
> HTML reports link errors, steps, traces, and attachments.

#### Documentation Quality

> High.<br>
> Official versioned documentation covers installation, concepts, complete API references, examples, CI providers, Docker, migration-relevant release notes, debugging, authentication, accessibility, visual testing, network mocking, and common best practices.<br>
> Fast release cadence means configuration should be checked against the pinned version rather than copied from next documentation.

#### Local Workflow

> CLI filters support a file, line, title grep, tag, project/browser, last failures, headed mode, one worker, repeats, and debug mode.<br>
> UI Mode runs or watches individual tests and groups interactively.<br>
> The VS Code extension discovers, runs, debugs, and generates tests. npx playwright show-report and show-trace inspect prior artifacts without a hosted service.

#### Failure Log Correlation

> Playwright timestamps and correlates browser actions, console entries, network traffic, assertions, and attachments inside the trace/report.<br>
> Event listeners can attach failed-response and browser-error logs to testInfo.<br>
> Gin, PostgreSQL, Flyway, and Docker Compose logs require custom setup/teardown or GitHub Actions steps that collect and attach/upload service logs, preferably with UTC timestamps and the run/test identifier.<br>
> There is no automatic cross-process correlation.

#### Artifact Data Exposure

> High sensitivity must be assumed.<br>
> Official CI guidance warns that traces, reports, console logs, and screenshots can expose credentials, access tokens, test data, and source.<br>
> HAR files may include headers, cookies, bodies, and portfolio records; storageState can contain impersonation-capable cookies, local storage, IndexedDB, and passkey private keys.<br>
> Use synthetic data, omit/filter HAR content, sanitize custom attachments/logs, never commit auth state, restrict artifact access and retention, and encrypt external sharing.<br>
> Trace Viewer itself processes traces locally in the browser and does not transmit them.

### Github Actions Fit

#### Official Ci Support

> Official GitHub Actions guidance uses actions/setup-node, npm ci, npx playwright install --with-deps, npx playwright test, and actions/upload-artifact.<br>
> Versioned Ubuntu Docker images are maintained.<br>
> No framework-specific GitHub Action is necessary.<br>
> This repository currently has Go test/lint workflows only, so a Node 24 E2E job must be added.

#### Browser Caching

> Official documentation does not recommend caching Playwright browser binaries because restoring the cache is often comparable to downloading them and Linux system packages are not cacheable. npm dependencies can use setup-node's npm cache.<br>
> Pin @playwright/test and install matching browsers, or pin the official Docker image; install only required engines to reduce time.

#### Artifact Integration

> Low effort: configure HTML plus blob/JSON/JUnit/GitHub reporters and upload playwright-report and selected test-results with actions/upload-artifact under if: !cancelled() or always() according to policy.<br>
> Traces, screenshots, videos, and custom backend logs are ordinary files/attachments.<br>
> Apply restricted access and short retention because artifacts may contain sensitive portfolio data.

#### Sharding And Matrix Support

> Native --shard=i/n works with a GitHub matrix.<br>
> Each shard can emit a blob report, and a dependent job can download and merge blobs into one HTML report.<br>
> Browser projects can run inside one job or as a browser matrix.<br>
> Retries occur within a shard; changing shard count or fullyParallel affects distribution.<br>
> Database-mutating shards require independent databases/schemas or must remain unsharded.

#### Container Compatibility

> Can run directly on ubuntu-latest with install --with-deps or in a version-matched mcr.microsoft.com/playwright image.<br>
> GitHub-hosted runners can also start this repository's Docker Compose services.<br>
> If Playwright itself runs in a container, host/container networking must expose the application using a reachable hostname such as a host-gateway mapping; localhost refers to the test container.<br>
> Official Docker guidance recommends --init and --ipc=host for Chromium stability.

#### Failure Cleanup

> Browser contexts, pages, worker fixtures, and webServer processes have runner-managed teardown; explicit context close ensures video/HAR flushing.<br>
> Project teardown and global teardown can clean test data and services, and webServer can send SIGTERM before SIGKILL.<br>
> Docker Compose, PostgreSQL volumes, and generated files remain repository responsibilities: GitHub Actions should run docker compose logs and docker compose down -v in an if: always() cleanup step, with unique project/volume names, because cancellation can bypass ordinary test teardown.

### Cost And Risk

#### Open Source Completeness

> All required local and GitHub Actions capabilities evaluated here are available in the Apache-2.0 @playwright/test package and managed open-source browser builds: runner, browsers, fixtures, assertions, parallelism, sharding, API/network controls, screenshots, visual comparison, video, traces, HTML/JSON/JUnit/blob/GitHub reports, Inspector, UI Mode, and codegen.<br>
> No paid service is required.

#### Optional Cloud Dependency

> None for execution, reports, trace viewing, visual baselines, or sharding.<br>
> Third-party hosted dashboards, browser grids, and visual-testing services are optional.<br>
> Equivalent core workflows can use local HTML/trace tools, committed baselines, GitHub artifacts, JSON/JUnit output, and repository-owned history processing, though open-source core does not provide a long-term hosted analytics dashboard.

#### Migration Cost

> Moderate framework lock-in.<br>
> Tests use Playwright-specific locators, fixtures, expect matchers, projects, routing APIs, traces, and browser protocol; these would require rewriting for WebDriver/Cypress.<br>
> Risk is reduced by standard TypeScript, async/await, accessible locator concepts, HTTP APIs, and black-box application boundaries.<br>
> Keep domain fixture helpers and test data independent from page APIs where practical.

#### Security And Supply Chain

> The package is Apache-2.0 and owned by Microsoft. @playwright/test 1.62.1 has a small direct npm surface, pinning playwright 1.62.1; npm metadata exposes integrity hashes, registry signatures, a SLSA provenance attestation link, and publication by a trusted GitHub Actions OIDC publisher.<br>
> Browser binaries and Linux dependencies are downloaded during installation unless supplied by a pinned image/cache, so versions, lockfiles, image digests, egress, and Dependabot/security updates should be controlled.<br>
> The browser process executes untrusted web content and should run with CI isolation; official Docker guidance distinguishes trusted E2E sites from untrusted scraping and recommends a non-root user/seccomp for the latter.

#### Custom Harness Burden

> Moderate and primarily application-specific.<br>
> Playwright supplies browser, fixture, API, readiness, artifact, and teardown primitives.<br>
> This repository still needs package/config files, an E2E workflow, Docker Compose startup/readiness, Flyway completion checks, deterministic PostgreSQL seed/reset and cleanup, Yahoo Finance replacement at the correct server/browser boundary, backend/service log attachment, and sensitive-artifact policy.<br>
> Existing build/start/destroy scripts help but production start persists PostgreSQL data, so disposable CI volumes and explicit cleanup are necessary.

#### Capability Delivery Tier

> Core: TypeScript runner, browser management, locators, auto-waiting, assertions, isolation, API requests, network interception, projects, workers, retries, sharding, visual snapshots, traces, video, screenshots, downloads, reports, Inspector, UI Mode, and codegen.<br>
> Official companion/integration: VS Code extension, Microsoft Docker images, and documented @axe-core/playwright use.<br>
> Custom repository code: Compose/Flyway/PostgreSQL lifecycle, database disposal, backend log collection, Yahoo Finance backend mocking when server-side, artifact redaction policy, and long-term flake trends.<br>
> Paid cloud: none required; hosted dashboards/grids are optional third-party capabilities.

#### Ai Execution Boundary

> The deterministic test runner, browser actions, assertions, retries, and CI do not require an LLM, network egress to an AI service, AI credentials, or per-run model cost.<br>
> Codegen, locator picking, and ordinary debugging are local non-LLM features.<br>
> Optional Playwright CLI/MCP or third-party AI authoring tools should remain outside required CI, receive least-privilege credentials and synthetic data, and produce reviewed TypeScript tests that continue to run without AI.

### Evidence And Decision

#### Sources

- Official Playwright installation and system requirements, https://playwright.dev/docs/intro, accessed 2026-08-22.
- Official Playwright release notes, https://playwright.dev/docs/release-notes, accessed 2026-08-22.
- Official Playwright browser documentation, https://playwright.dev/docs/browsers, accessed 2026-08-22.
- Official Playwright Docker documentation, https://playwright.dev/docs/docker, accessed 2026-08-22.
- Official Playwright locator and auto-waiting documentation, https://playwright.dev/docs/locators and https://playwright.dev/docs/actionability, accessed 2026-08-22.
- Official Playwright test configuration, projects, parallelism, retries, sharding, and webServer documentation, https://playwright.dev/docs/test-configuration, https://playwright.dev/docs/test-projects, https://playwright.dev/docs/test-parallel, https://playwright.dev/docs/test-retries, https://playwright.dev/docs/test-sharding, and https://playwright.dev/docs/test-webserver, accessed 2026-08-22.
- Official Playwright network, API testing, downloads, and authentication documentation, https://playwright.dev/docs/network, https://playwright.dev/docs/api-testing, https://playwright.dev/docs/downloads, and https://playwright.dev/docs/auth, accessed 2026-08-22.
- Official Playwright visual, accessibility, emulation, and clock documentation, https://playwright.dev/docs/test-snapshots, https://playwright.dev/docs/accessibility-testing, https://playwright.dev/docs/aria-snapshots, https://playwright.dev/docs/emulation, and https://playwright.dev/docs/clock, accessed 2026-08-22.
- Official Playwright reporters, trace, debugging, UI Mode, codegen, VS Code, and CI documentation, https://playwright.dev/docs/test-reporters, https://playwright.dev/docs/trace-viewer, https://playwright.dev/docs/debug, https://playwright.dev/docs/test-ui-mode, https://playwright.dev/docs/codegen, https://playwright.dev/docs/getting-started-vscode, and https://playwright.dev/docs/ci, accessed 2026-08-22.
- Microsoft Playwright GitHub repository and releases, https://github.com/microsoft/playwright and https://github.com/microsoft/playwright/releases, observed 2026-08-22 through GitHub API data.
- npm registry metadata for @playwright/test 1.62.1, https://registry.npmjs.org/%40playwright%2Ftest/latest, observed 2026-08-22.
- npm downloads API for @playwright/test, https://api.npmjs.org/downloads/point/last-week/%40playwright%2Ftest and matching 2025 window, observed 2026-08-22.
- Independent repository commit index, https://commits.ecosyste.ms/hosts/GitHub/repositories/microsoft/playwright, observed through web search 2026-08-22.
- Independent historical issue pulse, https://git-pulse.github.io/snapshots/microsoft-playwright-2025-07-03-pulse.html, used only for labeled 2025 issue-health context.
- Open Asset Allocator repository evidence: .nvmrc, src/main/web-static/package.json, tsconfig.json, .proxyrc.js, root.html, routing/HTMX/Chart modules, Gin server, Docker Compose files, lifecycle scripts, Makefile, and GitHub Actions workflows, inspected 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence in license, release/version, npm downloads, GitHub snapshot, Node support, browser/runtime capabilities, runner behavior, and repository architecture because these use first-party documentation, registry/API records, and local source.<br>
> Medium confidence in community-support breadth, issue health, maintainer bus factor, and the weighted fit score because some inputs are proxies or independent aggregates.<br>
> Low confidence in repository-specific flake/resource performance and end-to-end application behavior until the required spike runs.

#### Deal Breakers

> No confirmed exclusion-level incompatibility.<br>
> Material limitations are that managed WebKit is not native Safari, stock Firefox/Safari are not driven, Alpine/musl cannot run managed Firefox/WebKit, database and Compose lifecycle need a custom harness, artifacts can expose sensitive portfolio data, and managed browsers/protocols create framework lock-in.<br>
> Any requirement for native Safari, WebDriver-standard remote-grid portability, or zero browser-download footprint would become a deal breaker.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `roadmap_and_deprecation_risk`
- `dependency_currency`
- `maintainer_concentration`
- `github_metrics`
- `resource_usage`
- `hard_gate_result`
- `project_fit_score`
- `recommendation`
- `empirical_project_spike_result`

<a id="cypress"></a>
## 2. Cypress

Source result: `Cypress.json`

### Project And Compatibility

#### Implementation Language

> Cypress App is primarily implemented in TypeScript and JavaScript.<br>
> Tests, support code, plugins, and configuration can be authored in TypeScript or JavaScript and run with the Cypress binary plus a supported Node.js runtime.<br>
> [S1][S2][S8]

#### Node And Typescript Compatibility

> Verified for this repository's target stack.<br>
> Cypress 15.21.0 declares Node.js ^20.1.0, ^22.0.0, or >=24.0.0 and the official documentation supports npm >=10.1.0.<br>
> Cypress 15.14.0 added TypeScript 6.0 support; current documentation supports TypeScript 5.x, 6.x, and 7.x.<br>
> A separate cypress/tsconfig.json is advisable so Cypress's global Mocha, Chai, and jQuery types do not pollute the Parcel application's ESNext/bundler tsconfig.<br>
> Cypress configuration supports .ts, .mts, and .cts with Node module rules, so the repository may retain its current package.json without a type field and use cypress.config.ts as CommonJS, or use cypress.config.mts for explicit ESM.<br>
> [S1][S2][S8][S21]

#### Operating System Support

> Official support includes Linux x64 and arm64 on Ubuntu >=22.04, Debian >=11, and Fedora >=43.<br>
> GitHub-hosted Ubuntu runners fit this support window.<br>
> Native installation requires GTK, GBM, NSS, X11/Xvfb, audio, and related libraries; official Debian-based Cypress Docker images supply them.<br>
> On arm64, official browser-image availability is narrower: Chrome is available only from version 151, Firefox from 136, and Edge is unavailable.<br>
> [S1][S9]

#### License And Governance

> The Cypress App repository and npm package use the permissive MIT License and are owned and led by Cypress.io.<br>
> Development occurs publicly in cypress-io/cypress, while Cypress.io commercially stewards the optional proprietary Cypress Cloud service.<br>
> The license permits private, commercial, modified, and redistributed repository use subject to preserving the license notice.<br>
> This is company-led open source rather than neutral-foundation governance.<br>
> [S8][S19][S20]

#### Installation Model

> Install cypress as a dev dependency with npm install --save-dev cypress.<br>
> Its postinstall downloads a platform-specific Cypress binary into a global cache, normally ~/.cache/Cypress on Linux; npx cypress install can perform that step explicitly.<br>
> Electron is bundled in Cypress 15, while Chrome, Edge, and Firefox are system-installed, GitHub-runner-provided, or supplied through cypress/browsers images.<br>
> Experimental WebKit additionally requires playwright-webkit and Linux WebKit dependencies.<br>
> Official cypress/base, cypress/browsers, cypress/included, and cypress/factory images support pinned environments.<br>
> [S1][S3][S9]

#### Candidate Scope And Layer

> Complete end-to-end browser testing framework: integrated Mocha-based runner, assertions, browser launcher, command log, retry model, network proxy/interception, fixtures, screenshots, videos, reporters, plugin/node-event API, and interactive authoring application.<br>
> It is not merely a browser driver or an abstraction over another runner.<br>
> [S11][S12][S13][S14][S15]

#### Authoring And Async Model

> Tests use Cypress's framework-specific command queue and Chainable DSL, for example cy.get(...).should(...), rather than native async/await for browser commands.<br>
> Commands enqueue asynchronously; queries and assertions retry, while actions execute once after actionability checks.<br>
> Values are consumed through .then(), aliases, or cy.wrap().<br>
> Ordinary promises can be used in hooks and Node tasks, but mixing them with queued commands requires care.<br>
> This model is productive but creates meaningful Cypress lock-in and a learning boundary for developers expecting Promise-returning browser APIs.<br>
> [S12]

#### Build Pipeline Coupling

> E2E tests can black-box the existing application with baseUrl and cy.visit() and do not require Cypress to transform the Parcel source, substitute Vite, or instrument the production bundle.<br>
> The production build can be served by the Gin container on port 80; development can use Parcel on port 8000 with the existing /api proxy to Gin on port 8080.<br>
> Cypress separately preprocesses test/configuration TypeScript with its own tooling, so a dedicated Cypress tsconfig avoids coupling that test compilation to the application's TypeScript configuration.<br>
> [S2][S21]

#### Testability Instrumentation Required

> No special application route, injected script, relaxed CORS rule, or alternate build is mandatory for basic E2E operation.<br>
> Reliable suites would benefit from adding stable data-cy attributes to high-value controls because the current templates primarily expose dynamic IDs, visible labels, and ARIA labels; Cypress explicitly recommends dedicated data attributes.<br>
> Test-only database reset/seeding through an authenticated test API or a cy.task Node handler is also advisable.<br>
> These are testability improvements, not Cypress runtime requirements.<br>
> Cypress proxies and rewrites application traffic internally, so unusually strict CSP or SRI may require Cypress configuration, but no such project constraint was identified.<br>
> [S13][S21]

### Maintenance Health

#### Latest Stable Release

> 15.21.0, released 2026-08-18.<br>
> The npm latest metadata, GitHub release, and official changelog agree.<br>
> [S4][S5][S8]

#### Release Cadence

> High and consistent.<br>
> The ten most recent stable releases observed span 2026-04-29 through 2026-08-18, usually at roughly two-week intervals with patch releases between some minors.<br>
> Cypress's security page states that customers can expect minor releases approximately every two weeks.<br>
> Releases include browser/runtime support, performance work, dependency security updates, features, and bug fixes rather than version-only churn.<br>
> [S4][S5][S20]

#### Repository Activity

> Active as observed on 2026-08-22.<br>
> The default develop branch was pushed on 2026-08-22; sampled commits from 2026-08-17 through 2026-08-21 include browser compatibility, network-path, Studio, dependency, and release work.<br>
> Multiple pull requests were merged on 2026-08-20 and 2026-08-21, with active work from several Cypress members and external contributors.<br>
> [S6][S22]

#### Dependency Currency

> Current for the relevant environment.<br>
> Cypress 15.21 supports Node 24, npm 10+, TypeScript 6 and 7, and the latest three Chrome, Edge, and Firefox major versions.<br>
> Firefox automation moved to WebDriver BiDi with Firefox >=140 required.<br>
> Recent releases upgraded webdriver/geckodriver, tsx, esbuild, tar, ws, and undici, including explicit CVE responses.<br>
> Go and PostgreSQL versions are outside Cypress's runtime contract because Cypress talks to the application over HTTP.<br>
> [S1][S2][S3][S4][S8]

### Community Adoption

#### Npm Downloads

> The npm downloads API reported 6,155,552 downloads of the exact cypress package for the complete seven-day window 2026-08-14 through 2026-08-20.<br>
> [S7]

#### Ecosystem Usage

> Strong ecosystem evidence includes the maintained official GitHub Action, official multi-architecture Docker images, Real World App, example recipes, Cypress Testing Library, cypress-axe, Mocha reporter compatibility, multiple actively maintained visual-diff plugins, and integrations from major visual-testing vendors.<br>
> Cypress's site lists enterprise users, but those vendor claims are supporting rather than independent evidence.<br>
> [S9][S10][S15][S16][S17][S24]

#### Community Support

> Extensive, frequently updated official documentation covers concepts, APIs, CI providers, migration, troubleshooting, recipes, and best practices.<br>
> Public GitHub issues and discussions, Discord, learning material, videos, and a large body of community plugins and third-party answers provide multiple support paths.<br>
> Free users receive community support; paid Cloud tiers add email or premium support.<br>
> [S1][S10][S12][S15][S24]

#### Adoption Metric Normalization

> The primary registry metric is direct downloads of npm package cypress, not all Cypress plugins, Docker pulls, Cloud users, or unique projects, over the explicitly fixed 2026-08-14 to 2026-08-20 window.<br>
> Repeated CI installs and cache misses can count multiple times, while cached binaries can reduce downloads.<br>
> The package includes E2E and component-testing capabilities, so downloads cannot be attributed solely to E2E.<br>
> Stars and forks are cumulative interest signals, not current production usage.<br>
> The vendor's 1.5M+ dependent-repository claim is rounded and may include direct and transitive package relationships.<br>
> [S6][S7][S8][S24]

### Browser And Runtime Coverage

#### Browser Engines

> Stable support covers Chromium-family browsers through bundled Electron 15.x, installed Chrome/Chrome for Testing/Chromium, and Microsoft Edge, plus installed Firefox.<br>
> Cypress officially supports the latest three major Chrome, Edge, and Firefox versions.<br>
> WebKit is experimental, requires experimentalWebKitSupport plus playwright-webkit, and has documented gaps including no cy.origin(), no Test Replay, disabled forceNetworkError, and input/stack/video limitations.<br>
> Native Safari and mobile browsers are not controlled.<br>
> [S1][S3]

#### Browser Protocol

> Cypress uses a framework-specific architecture combining code in the browser, a local Node process, and an HTTP proxy.<br>
> Chromium launch/control uses Chrome DevTools Protocol; current Firefox automation uses WebDriver BiDi and geckodriver for launch, with Firefox >=140 required; experimental WebKit is launched through Playwright's WebKit integration.<br>
> The mixed implementation enables network rewriting and time-travel UX but is not a portable WebDriver client and inherits browser-specific constraints.<br>
> [S3][S13]

#### Headless And Headed Modes

> cypress run is headless by default for Electron, Chrome-family, Firefox, and experimental WebKit and accepts --headed for visible execution. cypress open is always headed and provides the interactive runner.<br>
> Headless Chrome uses the modern headless implementation; headless Linux is supported directly or through Xvfb where required.<br>
> [S3][S9]

#### Browser Version Management

> Each Cypress 15 release bundles a fixed Electron/Chromium runtime.<br>
> Cypress detects installed Chrome-family and Firefox browsers; Chrome for Testing or long-form cypress/browsers image tags can pin exact versions and avoid auto-update drift.<br>
> Official images pin Node and browser combinations.<br>
> Firefox may download the latest geckodriver through its wrapper when no cached driver exists, creating an air-gapped/network and possible drift concern unless a current image or explicit cache/path is used.<br>
> Experimental WebKit version follows the installed playwright-webkit package.<br>
> [S3][S9]

#### Parallel Browser Support

> One Cypress process selects one browser per invocation; Cypress has no Playwright-style projects array that runs a browser matrix in one command.<br>
> Run separate CI matrix jobs for Chrome, Firefox, Edge, Electron, or WebKit.<br>
> Specs run sequentially in one Cypress process; multi-machine spec distribution with --parallel and dynamic load balancing requires recording to Cypress Cloud, while fully local parallelism requires explicit spec partitioning or community/custom tooling.<br>
> [S9][S10]

#### Mobile Emulation

> cy.viewport() supports explicit dimensions, orientation, and named phone/tablet viewport presets, and configuration can override userAgent.<br>
> This is responsive-layout simulation, not full device emulation: devicePixelRatio is not simulated, and Cypress does not provide a unified core device profile covering touch hardware, sensors, locale, timezone, network, and mobile browser binaries.<br>
> Real iOS Safari and Android Chrome require external infrastructure rather than core Cypress.<br>
> [S25]

#### Real Browser Fidelity

> Installed Chrome, Chromium, Edge, and Firefox runs use real desktop browser binaries in Cypress-controlled isolated profiles.<br>
> Bundled Electron is Chromium-based but is not an end-user browser, and Cypress recommends another browser for representative coverage.<br>
> Experimental WebKit uses Playwright's WebKit build; it is not Apple's Safari application and cannot validate Safari-only UI chrome, Apple platform integration, codecs, keychain, or OS rendering.<br>
> Viewport presets likewise do not equal physical mobile-device testing.<br>
> [S3][S25]

#### Environment Determinism Controls

> Core controls include cy.clock()/cy.tick() for Date and browser timers, fixed viewport dimensions/orientation, isolated browser profiles, browser launch flags/preferences, network stubs/delays/errors, fixed fixtures, and screenshot animation handling.<br>
> Headless runs default to 1280x720 screen and DPR 1, while the AUT viewport defaults to 1000x660 unless configured.<br>
> Locale, timezone, geolocation, permissions, DPR, reduced motion, and randomness lack one unified cross-browser device/project profile and generally require environment variables, browser launch arguments, stubs, CSS, or application-level setup. cy.clock() affects only the top window, not embedded iframes.<br>
> [S3][S13][S25][S26]

### Application Fit

#### Dynamic Dom Synchronization

> Good fit for HTMX swaps and asynchronous Handlebars rendering when tests use fresh query chains and observable end states.<br>
> Cypress retries linked DOM queries/assertions and re-runs queries before an action until actionability passes; cy.intercept aliases can explicitly await HTMX XHR/fetch responses.<br>
> Because a passed mid-chain assertion locks its subject, an HTMX replacement can still produce a detached element for later chained queries.<br>
> Split chains and re-query after actions/swaps rather than retaining jQuery elements or using non-retried .then() callbacks.<br>
> [S12][S13][S21]

#### Routing Support

> Navigo history/path routing can be exercised through links and browser actions, while retrying cy.location('pathname'), cy.url(), and cy.go() assertions verify route transitions.<br>
> Direct deep links work if the independently served application returns root.html or the intended route fallback; this must be verified for the Gin production server.<br>
> Test isolation visits about:blank before each test, so each test should visit its route and not depend on prior route state.<br>
> [S21][S27]

#### Locator Model

> Core Cypress offers CSS selectors through cy.get(), scoped queries, visible-text matching through cy.contains(), and custom queries.<br>
> Accessible role/name locators require the maintained community Cypress Testing Library plugin rather than a core locator API.<br>
> Core queries are not globally strict: they can yield collections, although many actions reject multiple targets unless configured.<br>
> The project has useful IDs and ARIA labels but no observed data-cy attributes.<br>
> Stable data-cy selectors are Cypress's recommendation.<br>
> Fresh query chains reacquire HTMX-replaced elements; aliases of DOM queries can re-run, but captured raw elements and subjects after assertion boundaries can become stale.<br>
> [S12][S17][S21]

#### Form Interaction

> Cypress supports focus, blur, clear, type, select, check, uncheck, submit, native Tab through cy.press(), and assertions on values, validity, attributes, and focused elements.<br>
> Separate cy.get() calls between focus/type/blur actions are recommended when application handlers can replace a row.<br>
> This can cover the repository's input masks, hidden-value synchronization, dynamic allocation rows, Bootstrap validation classes, and native validity APIs.<br>
> Input events are browser-dispatched by Cypress rather than operating-system-level keystrokes; WebKit documents additional type-event differences.<br>
> [S3][S12][S17][S21]

#### Canvas And Download Support

> Chart.js canvases can be validated without relying only on pixels by asserting canvas size/presence, reading the 2D context or toDataURL(), exposing/inspecting the Chart instance or its application data, and checking accessible/table equivalents; Cypress visual plugins can add pixel baselines for a small stable set.<br>
> Browser downloads are redirected to cypress/downloads without a native Save As prompt and can be checked with cy.readFile() or cy.task().<br>
> Cypress does not provide built-in screenshot comparison, and project code may need to expose chart semantic state if it is not otherwise reachable.<br>
> [S3][S16][S21]

#### Network And Api Access

> Strong fit. cy.request() can seed/query REST endpoints without the UI, while cy.intercept() observes, aliases, mutates, delays, fails, or fixtures XHR/fetch requests.<br>
> Interception can let the Gin API run normally while replacing only Yahoo Finance-bound browser traffic if that traffic originates in the browser.<br>
> If Yahoo Finance is called exclusively by the Go backend, browser interception cannot see it; use backend dependency injection, a fake upstream endpoint, or seeded API/database state.<br>
> Request and response bodies, headers, methods, URLs, and statuses are assertable.<br>
> [S13][S21]

#### Same Origin Support

> The consolidated production server exposes the application and /api from one origin at localhost:80, and Parcel's development server exposes localhost:8000 while proxying /api to localhost:8080.<br>
> Cypress baseUrl can target either public origin with relative visits and requests, so no CORS change is needed.<br>
> Cross-origin restrictions and cy.origin() are not involved in these normal paths.<br>
> [S13][S21]

#### Test Isolation

> Cypress clears the page, cookies, localStorage, sessionStorage, aliases, clocks, intercepts, spies, stubs, and viewport changes before each E2E test when testIsolation is enabled.<br>
> It does not create a new OS process or independent browser context per test, and IndexedDB or other storage must be cleared manually.<br>
> PostgreSQL is external to browser isolation; reliable parallel suites need a reset/seed API or cy.task plus unique test namespaces, transactions, schemas, databases, or serialized DB-mutating specs.<br>
> [S18]

#### External Server Model

> Supported and conventional.<br>
> Set e2e.baseUrl or CYPRESS_BASE_URL and run tests against any independently started Gin/Parcel/Docker Compose application.<br>
> Cypress verifies baseUrl availability for visits but expects the server to be started outside test code.<br>
> Different local and CI ports can be selected through configuration/environment overrides.<br>
> [S9][S21]

#### Application Lifecycle

> Cypress itself does not provide a declarative multi-service webServer lifecycle.<br>
> The official GitHub Action has build, start, and wait-on options, and start-server-and-test or custom Node orchestration can start one or more commands.<br>
> This repository needs custom CI steps to build the Parcel/Gin image, create disposable PostgreSQL storage, run Docker Compose, wait for database health, Flyway completion and HTTP/API readiness, execute Cypress, capture service logs, and docker compose down --volumes under an always-run cleanup step.<br>
> Existing Compose dependency conditions help, but start.sh persists data under the home directory and stop.sh stops rather than removes services/volumes, so they are not sufficient as an isolated CI harness.<br>
> [S9][S10][S21]

#### Visual Regression Workflow

> Cypress core captures screenshots but does not compare them.<br>
> Official documentation lists actively maintained open-source image-diff plugins and commercial services.<br>
> An open-source project workflow must own baseline files, approval review, masking/threshold configuration, per-browser baselines, and CI diff uploads.<br>
> Pin the browser/OS/fonts in a Cypress Docker image, fix viewport/time/data, disable animations, and wait for Chart.js completion.<br>
> Chart canvases are visually testable, but semantic dataset assertions should remain primary to reduce pixel flake.<br>
> [S14][S16]

#### Accessibility Audit Integration

> Compatible through the community cypress-axe plugin using axe-core, with other open-source reporters/plugins available.<br>
> Cypress Testing Library can add role/label locators, and cy.press() supports native Tab-focused keyboard tests.<br>
> Core Cypress does not supply ARIA snapshot assertions.<br>
> Cypress Accessibility performs automatic step-level analysis in Cloud but is a paid premium solution; it is unnecessary for basic open-source axe scans and explicit semantic/keyboard assertions.<br>
> [S17]

### Reliability

#### Waiting Model

> Linked queries and assertions retry from the top until their timeout; action commands re-run preceding queries and wait for existence, visibility, enablement, coverage, stability, and scroll/actionability before acting once. cy.visit() waits for page load, cy.location()/cy.url() poll SPA state, and cy.intercept()+cy.wait('@alias') provides explicit request/response synchronization.<br>
> Non-query actions, .then() callbacks, and assertions chained directly to some yielded one-shot values are not automatically replayed, so tests must respect retry boundaries.<br>
> [S12][S13][S27]

#### Flake Controls

> Configurable retries exist globally, separately for run/open mode, and per suite/test; default is zero.<br>
> Every retry reruns beforeEach/afterEach and creates attempt-specific failure screenshots.<br>
> Experimental retry strategies can require passes and classify flake.<br>
> Timeouts are scoped per command/configuration, deterministic waits can use network aliases or state assertions, and fixed sleeps are discouraged.<br>
> Cypress lacks a core repeat-each stress flag and automatic quarantine in the local runner; repeat loops and quarantine policy require CI/custom tooling or Cloud features.<br>
> [S11][S12]

#### Isolation Model

> Default E2E isolation resets the page and common browser storage before each test, but one runner process does not provide Playwright-style independent browser contexts.<br>
> IndexedDB persists unless explicitly deleted.<br>
> Suites may opt out at describe level for serial scenarios, but this increases order dependence.<br>
> Database state is never reset by Cypress, so serial DB scenarios should be deliberately grouped while independent tests reset or namespace server state.<br>
> [S18]

#### Parallelism Controls

> The safest initial configuration is one DB-mutating Cypress invocation and one browser, with spec-level serialization.<br>
> Later, GitHub matrices can split read-only specs or use isolated PostgreSQL databases/schemas per job.<br>
> Cypress Cloud can dynamically load-balance whole specs across machines but cannot make shared database writes safe; explicit groups, separate base URLs/data namespaces, and CI concurrency limits remain necessary.<br>
> Without Cloud, assign disjoint --spec globs to jobs or use a maintained splitter and merge reports explicitly.<br>
> [S9][S10]

#### Flake Observability

> The open runner shows retry attempts and preserves attempt-specific screenshots; videos can include every attempt, and JUnit/JSON-capable community reporters can expose attempt results for CI processing.<br>
> It does not provide a built-in local historical flake database, quarantine workflow, trend dashboard, or replay trace.<br>
> Cypress Cloud adds flaky badges, flake rate/history, alerts, Test Replay, analytics, and test-level rerun features, with some flake analytics requiring paid tiers.<br>
> A service-free solution must ingest reporter output and retain artifacts/history itself.<br>
> [S11][S15][S24]

### Diagnostics And Developer Experience

#### Failure Artifacts

> Core local artifacts include automatic failure screenshots in cypress run, manual screenshots, optional per-spec MP4 video, CLI stack/error output, and reporter files.<br>
> Interactive Command Log snapshots show DOM state and XHR/fetch routes.<br>
> Cypress does not create a portable local execution trace, HAR, or complete browser-console/network archive comparable to a trace viewer; console capture and durable network logs need event hooks/plugins.<br>
> Cypress Cloud adds Test Replay with DOM, network, console, JavaScript errors, and rendering context.<br>
> [S13][S14][S15]

#### Debugging Tools

> cypress open provides a headed interactive runner, live Command Log, DOM snapshots/time travel, browser DevTools console output, reruns, .pause(), .debug(), and selector inspection. cypress run supports --headed and --no-exit for reproduction.<br>
> IDE type declarations/JSDoc and source maps improve navigation.<br>
> Cloud Test Replay supplies remote time travel but is optional.<br>
> [S2][S12][S14]

#### Test Generation

> Cypress Studio records click, type, check, uncheck, and select interactions, generates selectors, supports manual assertions, and edits test source.<br>
> Non-AI Studio does not require a Cloud account, but current documentation says it requires internet access and source maps.<br>
> Studio AI is opt-in, Cloud-connected, beta functionality that proposes assertions.<br>
> Generated selector quality follows a priority headed by data-cy/data-test/data-testid, then progressively more fragile attributes, so generated code still requires review.<br>
> [S23]

#### Reporters

> The built-in default spec reporter writes to stdout; JUnit and TeamCity are bundled, and any compatible Mocha reporter can be installed, including JSON/HTML options such as Mochawesome.<br>
> Multiple reporters require a package such as cypress-multi-reporters.<br>
> Every spec is processed separately, so report filenames need a hash and JUnit/Mochawesome outputs need an explicit merge step.<br>
> GitHub annotations can be produced through suitable reporters or workflow processing; Cloud provides its own hosted reports.<br>
> [S15]

#### Documentation Quality

> High breadth and currentness.<br>
> Official pages separately document installation, TypeScript, command semantics, browser limits, CI, retries, isolation, networking, artifacts, accessibility, visual testing, security, migrations, changelogs, and common failure modes, often with JavaScript and TypeScript examples.<br>
> Several pages used here were updated in July or August 2026.<br>
> The documentation also states important limitations, including experimental WebKit, non-cleared IndexedDB, non-strict retry boundaries, and Cloud-only features.<br>
> [S1][S2][S3][S9][S12][S18]

#### Local Workflow

> Use npx cypress open for browser selection, spec selection, interactive reruns, time travel, Studio, and DevTools.<br>
> Use npx cypress run --spec <glob> --browser chrome for a focused CI-like run, optionally --headed --no-exit to retain the final state.<br>
> Separate runMode/openMode retries avoid hiding failures during authoring.<br>
> The repository must first start its Gin/Parcel/Compose services; Cypress should not start them from inside a test.<br>
> [S3][S9][S11][S23]

#### Failure Log Correlation

> Cypress timestamps and orders browser commands, request aliases, errors, screenshots, and videos, but it does not automatically ingest Gin, PostgreSQL, Flyway, or Docker Compose logs into one local timeline.<br>
> The workflow should record test start/end times, stream service logs to separate files, label cy.intercept aliases, capture browser uncaught exceptions/console messages, and upload compose ps/logs plus Flyway output on failure.<br>
> Cloud Replay correlates browser-side evidence only; backend correlation remains custom.<br>
> [S9][S13][S21]

#### Artifact Data Exposure

> Screenshots, video, HTML/JSON/JUnit reports, Command Log snapshots, network bodies, console output, downloads, and Cloud Test Replay can expose portfolio values, credentials, API responses, or tokens. cy.env() hides values only while retrieving them; later assertions/logs can print them.<br>
> Use synthetic data, Cypress/GitHub secrets, {log:false} for sensitive commands/intercepts, screenshot blackout where practical, selective artifact retention, private CI artifacts, and no storage-state dumps.<br>
> Cloud uploads test content to US multi-tenant storage; public projects make content public.<br>
> No universal automatic redaction for all local artifacts was identified.<br>
> [S14][S20][S28]

### Github Actions Fit

#### Official Ci Support

> Strong.<br>
> Cypress maintains cypress-io/github-action, with v7 recommended and v7.4.3 released 2026-08-20.<br>
> Official documentation targets ubuntu-24.04 and covers dependency installation, start/wait, browser selection, caching, matrices, Cloud recording, and Docker images.<br>
> GitHub-hosted Ubuntu includes supported browsers, while official images offer pinning.<br>
> [S10][S22]

#### Browser Caching

> The official action automatically caches/restores package-manager and Cypress dependencies.<br>
> For manual workflows, cache the package-manager cache and ~/.cache/Cypress after npm ci; do not cache node_modules.<br>
> Use strict cache keys to avoid accumulating old Cypress binaries.<br>
> Long-form Docker tags avoid runtime browser download but shift caching to container layers.<br>
> Firefox driver downloads should also be made deterministic through a current image/cache for offline CI.<br>
> [S9][S10]

#### Artifact Integration

> Straightforward but not automatic: use actions/upload-artifact with if: always() for cypress/screenshots, cypress/videos, JUnit/Mochawesome output, downloads as appropriate, and collected application/Compose logs.<br>
> Configure screenshot/video retention and compression to control size.<br>
> Cypress Cloud can host screenshots, videos, Test Replay, and reports when --record is used, but local GitHub artifacts remain sufficient for core OSS execution.<br>
> [S10][S14][S15]

#### Sharding And Matrix Support

> GitHub matrix jobs can cover browsers, viewports, and manually assigned spec partitions.<br>
> Native Cypress --parallel dynamic spec balancing, run grouping, merged hosted results, and balancing require Cypress Cloud recording.<br>
> Service-free sharding needs fixed --spec partitions or community tooling and explicit JUnit/Mochawesome merge jobs.<br>
> Retry attempts occur within the worker that owns a spec; rerunning a failed GitHub job is separate from Cypress test retries.<br>
> Browser versions must match across Cloud parallel workers, so pinned Docker images are recommended.<br>
> [S10][S15]

#### Container Compatibility

> Cypress runs well on Ubuntu hosts or in official Debian containers.<br>
> This repository itself needs Docker Compose for Gin, PostgreSQL, and Flyway.<br>
> The least complex GitHub setup is to run Cypress directly on the Ubuntu host, start Compose on the same host, and browse exposed localhost ports.<br>
> Running the whole job inside cypress/browsers complicates access to the host Docker daemon and sibling Compose networking unless the workflow deliberately mounts/configures Docker.<br>
> Pin Chrome for Testing or an image where reproducibility outweighs that complexity.<br>
> [S9][S10][S21]

#### Failure Cleanup

> Cypress normally closes its browser/Xvfb process, but repository services and data volumes require explicit cleanup.<br>
> Add an if: always() step that captures docker compose ps/logs and runs docker compose down --volumes --remove-orphans, and use job-level timeouts/concurrency cancellation.<br>
> Use disposable workspace-owned PostgreSQL storage rather than start.sh's persistent home-directory path.<br>
> GitHub cancellation can interrupt ordinary commands, so cleanup should be idempotent and the next run should remove stale project-scoped resources before startup.<br>
> [S9][S21]

### Cost And Risk

#### Open Source Completeness

> All capabilities required to author and execute deterministic local/CI E2E tests are available under MIT or compatible open-source tooling: runner, Chrome/Firefox control, network interception, API calls, retries, screenshots, videos, JUnit, Docker images, axe-core integration, and community visual diffing.<br>
> Stable WebKit is not available, and advanced orchestration/observability are not open-source core capabilities.<br>
> [S14][S15][S16][S17][S19]

#### Optional Cloud Dependency

> No Cloud account is required for ordinary tests, retries, screenshots, local videos, network control, reports, or open mode.<br>
> Cypress Cloud is required for Cypress's native --parallel load balancing and supplies hosted Test Replay, retained runs/artifacts, analytics, integration comments/statuses, flake management, prioritization, cancellation, UI Coverage, Accessibility, and AI features according to tier.<br>
> The free plan listed 500 recorded results/month; Team started at USD 799/year on 2026-08-22.<br>
> Equivalent local pieces exist for artifacts, reports, fixed sharding, axe, and visual diffs, but no equally integrated core local replay/analytics service is bundled.<br>
> [S9][S10][S17][S24]

#### Migration Cost

> Moderate to high for a mature suite.<br>
> Test intent, CSS/data selectors, fixtures, and API concepts are portable, but cy.* command queues, Chainable custom commands, aliases, retry semantics, cy.intercept, Node events/tasks, Cypress configuration, Studio output, and Cloud run metadata are framework-specific.<br>
> Moving to native async/await frameworks requires structural rewrites rather than mechanical import changes.<br>
> Avoid overusing custom commands and Cloud-only control flow to limit lock-in.<br>
> [S12][S13][S15]

#### Security And Supply Chain

> The npm package is signed, publishes integrity hashes, and release 15.21 includes a downloadable SBOM artifact.<br>
> The MIT source, public changelog, SECURITY.md disclosure route, frequent dependency upgrades, and recent CVE-driven tar/esbuild/undici updates are positive.<br>
> Risks include a broad dependency tree, a postinstall binary download, optional geckodriver download, official Docker-image trust, privileged browser/proxy behavior, and optional telemetry/Cloud egress.<br>
> Pin package/image versions and lockfiles, cache verified binaries, scan dependencies/images/SBOMs, use CYPRESS_DISABLE_GUEST_TELEMETRY where required, and keep secrets out of bundled specs and artifacts.<br>
> [S4][S5][S8][S20]

#### Custom Harness Burden

> Moderate.<br>
> Cypress supplies browser/network/test lifecycle but the repository must add Cypress configuration/specs, stable selectors, REST/database fixtures, a PostgreSQL reset or namespace strategy, Yahoo Finance replacement at the correct backend boundary, Compose/Flyway readiness, deterministic ports, log collection, report/artifact upload, and cleanup.<br>
> The official action reduces Node/cache/start-wait work, but the multi-service disposable environment and backend external-service seam remain project-owned.<br>
> [S9][S10][S13][S21]

#### Capability Delivery Tier

> Core OSS: E2E runner, TypeScript, Chrome/Edge/Firefox/Electron 15, headless/headed mode, command retries, optional test retries, test isolation, cy.request, cy.intercept, screenshots, video, downloads, spec/JUnit reporters, and open mode.<br>
> Experimental core plus external package: WebKit via playwright-webkit.<br>
> Official OSS companion: GitHub Action and Docker images.<br>
> Community OSS: Testing Library accessible locators, cypress-axe, visual image diffing, richer/multiple reporters, and fixed sharding helpers.<br>
> Custom project code: Docker Compose/Flyway readiness, DB reset, backend Yahoo stub, service-log correlation, deterministic cleanup, and semantic Chart.js assertions.<br>
> Cloud/free-or-paid SaaS: smart parallelization/load balancing, Test Replay, hosted artifacts/history/analytics; paid/premium tiers add richer flake analytics, prioritization/cancellation, UI Coverage, Accessibility, and enterprise controls.<br>
> [S3][S9][S10][S15][S16][S17][S24]

#### Ai Execution Boundary

> Deterministic Cypress tests and CI execution do not require an LLM.<br>
> Studio works without AI or a Cloud account, although current Studio requires internet access; Studio AI and cy.prompt are optional Cloud/AI-assisted authoring or execution features with usage limits and egress.<br>
> Do not commit cy.prompt into required CI paths if deterministic non-AI operation is a hard requirement.<br>
> Keep generated assertions under human review, disable Studio AI/guest telemetry where policy requires, and use ordinary explicit cy commands as the complete fallback.<br>
> Cypress states AI data is session-bound and not used for model training, but use synthetic data and avoid exposing portfolio/customer content regardless.<br>
> [S20][S23][S24]

### Evidence And Decision

#### Sources

- [S1] Cypress Installation and System Requirements, https://docs.cypress.io/app/get-started/install-cypress, accessed 2026-08-22.
- [S2] Cypress TypeScript Support, https://docs.cypress.io/app/tooling/typescript-support, accessed 2026-08-22.
- [S3] Cypress Launching Browsers, https://docs.cypress.io/app/references/launching-browsers, accessed 2026-08-22.
- [S4] Cypress App Changelog, https://docs.cypress.io/app/references/changelog, accessed 2026-08-22.
- [S5] cypress-io/cypress GitHub releases, including v15.21.0, https://github.com/cypress-io/cypress/releases/tag/v15.21.0, observed 2026-08-22.
- [S6] GitHub repository metadata for cypress-io/cypress, https://github.com/cypress-io/cypress, observed 2026-08-22.
- [S7] npm Downloads API for cypress, https://api.npmjs.org/downloads/point/last-week/cypress and fixed 2025/2026 comparison windows, retrieved 2026-08-22.
- [S8] npm registry metadata for cypress 15.21.0, https://registry.npmjs.org/cypress/latest, retrieved 2026-08-22.
- [S9] Cypress Continuous Integration Overview, https://docs.cypress.io/app/continuous-integration/overview, accessed 2026-08-22.
- [S10] Cypress GitHub Actions Guide, https://docs.cypress.io/app/continuous-integration/github-actions, accessed 2026-08-22.
- [S11] Cypress Test Retries, https://docs.cypress.io/app/guides/test-retries, accessed 2026-08-22.
- [S12] Cypress Retry-ability, https://docs.cypress.io/app/core-concepts/retry-ability, accessed 2026-08-22.
- [S13] Cypress Network Requests, https://docs.cypress.io/app/guides/network-requests, accessed 2026-08-22.
- [S14] Cypress Screenshots and Videos, https://docs.cypress.io/app/guides/screenshots-and-videos, accessed 2026-08-22.
- [S15] Cypress Reporters, https://docs.cypress.io/app/tooling/reporters, accessed 2026-08-22.
- [S16] Cypress Visual Testing, https://docs.cypress.io/app/tooling/visual-testing, accessed 2026-08-22.
- [S17] Cypress Accessibility Testing, https://docs.cypress.io/app/guides/accessibility-testing, accessed 2026-08-22.
- [S18] Cypress Test Isolation, https://docs.cypress.io/app/core-concepts/test-isolation, accessed 2026-08-22.
- [S19] Cypress MIT License at v15.21.0, https://github.com/cypress-io/cypress/blob/v15.21.0/LICENSE, accessed 2026-08-22.
- [S20] Cypress Security and Compliance and repository SECURITY.md, https://www.cypress.io/security and https://github.com/cypress-io/cypress/blob/v15.21.0/SECURITY.md, accessed 2026-08-22.
- [S21] Open Asset Allocator repository files: src/main/web-static/package.json, tsconfig.json, .proxyrc.js, HTMX/Navigo/Handlebars/Chart/form code, build/start/dev/stop scripts, and Docker Compose definitions, inspected 2026-08-22.
- [S22] GitHub API observations for cypress-io/cypress commits, issues, pull requests, releases, cypress-io/github-action v7.4.3, and official companion repositories, observed 2026-08-22.
- [S23] Cypress Studio AI, https://docs.cypress.io/app/guides/cypress-studio, accessed 2026-08-22.
- [S24] Cypress Pricing, Cloud feature tiers, and self-reported adoption, https://www.cypress.io/pricing, accessed 2026-08-22.
- [S25] Cypress cy.viewport API, https://docs.cypress.io/api/commands/viewport, accessed 2026-08-22.
- [S26] Cypress cy.clock API, https://docs.cypress.io/api/commands/clock, accessed 2026-08-22.
- [S27] Cypress cy.location API, https://docs.cypress.io/api/commands/location, accessed 2026-08-22.
- [S28] Cypress Environment Variables and Secrets, https://docs.cypress.io/app/guides/environment-variables, accessed 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence: release/version, Node 24, TypeScript 6, OS/browser support, license, npm downloads, GitHub stars/forks, core command/network/isolation/artifact behavior, CI guidance, and repository topology, because these come from current official documentation, registry metadata, APIs, and source files.<br>
> Moderate confidence: application fit, harness burden, ecosystem depth, and weighted score, because they are reasoned from documented behavior without executing Cypress here.<br>
> Low confidence: issue trend/response distributions, upstream lag duration, maintainer bus factor, comparative rank, and project performance/flake results, because no longitudinal analysis or empirical spike was performed.

#### Deal Breakers

> No confirmed project-specific exclusion was found.<br>
> Potential exclusion conditions are: stable Safari/WebKit coverage is mandatory; native built-in service-free trace/replay and historical flake analytics are mandatory; dynamic sharding must work without Cypress Cloud or custom tooling; native async/await authoring is required; or the team will not accept a framework-specific command queue.<br>
> None of these constraints is stated as mandatory in the supplied scope.<br>
> The unperformed spike could still reveal HTMX timing, Gin deep-link, Chart.js, download, or database-reset problems.

### Uncertain Fields

- `issue_health`
- `roadmap_and_deprecation_risk`
- `wrapper_upstream_lag`
- `maintainer_concentration`
- `github_metrics`
- `adoption_trend`
- `resource_usage`
- `hard_gate_result`
- `project_fit_score`
- `recommendation`
- `empirical_project_spike_result`

<a id="webdriverio"></a>
## 3. WebdriverIO

Source result: `WebdriverIO.json`

### Project And Compatibility

#### Implementation Language

> The framework is implemented primarily in TypeScript and published as ESM and CommonJS-compatible Node.js packages.<br>
> Tests can be authored in TypeScript or JavaScript and run with Mocha, Jasmine, or Cucumber adapters.<br>
> Node.js is the required runtime; Go is not required for test authoring.

#### Operating System Support

> The Node.js client and local runner support Linux, macOS, and Windows.<br>
> Linux x64 is a normal local and GitHub-hosted Ubuntu target for headless Chrome/Chromium and Firefox.<br>
> Linux arm64 viability depends on browser and driver binary availability; Safari requires macOS or a remote provider, and Edge browser setup is manual or remote because automatic browser installation does not support Edge.

#### License And Governance

> Core WebdriverIO is MIT licensed.<br>
> It is an OpenJS Foundation project governed by a Technical Steering Committee under a published charter and consensus process, with an employer-concentration limit of one third of TSC seats.<br>
> Funding is disclosed through Open Collective, GitHub Sponsors, and Tidelift; commercial browser vendors sponsor the project but do not own the code.<br>
> The MIT license imposes no material constraint on repository use beyond retaining copyright and license notices.

#### Installation Model

> Recommended bootstrap is `npm init wdio@latest .`; manual setup starts with `npm install --save-dev @wdio/cli` plus a framework adapter, local runner, reporter, and `tsx` for TypeScript.<br>
> Since v8.14, WebdriverIO can locate or download Chrome/Chromium and Firefox and their drivers, or use supplied binaries.<br>
> Automatic browser installation excludes Edge, and SafariDriver ships with macOS.<br>
> Browser system libraries still need to exist on bare Linux; alternatively use Selenium standalone browser images.<br>
> Optional capabilities add packages such as @wdio/visual-service, @axe-core/webdriverio, and @wdio/devtools-service.

#### Candidate Scope And Layer

> Complete Node.js E2E framework: integrated CLI, configurable test runner, assertions, services, reporters, parallel workers, sharding, and direct WebDriver/WebDriver BiDi automation.<br>
> It can also operate as a standalone browser client and connect to Selenium Grid, Appium, or commercial browser clouds.

#### Authoring And Async Model

> Native asynchronous JavaScript and TypeScript using async/await.<br>
> The default setup uses Mocha and a WebdriverIO command API with expect-webdriverio assertions; Jasmine and Cucumber are supported through adapters.<br>
> There is no Cypress-style command queue.<br>
> Globals are optional and can be replaced with explicit imports from @wdio/globals.

#### Build Pipeline Coupling

> Black-box E2E tests use an independently built and served application through `baseUrl`; they do not transform application source and therefore can exercise the existing Parcel output, Gin server, or Parcel development proxy unchanged.<br>
> Only WebdriverIO's separate browser runner for unit/component tests uses Vite, so that runner should not be selected for this repository's E2E suite.

#### Testability Instrumentation Required

> No production build hook, source transform, injected runtime, alternate route, or relaxed content-security policy is required for ordinary WebDriver E2E execution.<br>
> Existing accessible names, text, IDs, and semantic controls are usable.<br>
> Adding stable `data-testid` attributes is optional where the current DOM has no durable user-facing locator.<br>
> Network interception, fake clock, and API emulation install test-session behavior and may depend on WebDriver BiDi or CDP, but do not require shipped application code changes.

### Maintenance Health

#### Latest Stable Release

> WebdriverIO v9.31.2, published 2026-08-21.<br>
> GitHub marks it non-draft and non-prerelease, and the npm registry reports @wdio/cli 9.31.2 as latest.

#### Release Cadence

> Frequent and sustained.<br>
> The v9 changelog contains 19 stable tags from 2026-01-03 through 2026-08-21, including feature and patch releases, with occasional same-day corrective releases and longer gaps of several weeks.<br>
> The project also issued v8 maintenance releases in June and July 2026.

#### Repository Activity

> Very active as observed on 2026-08-22: the main branch was pushed on 2026-08-21, the latest ten commits span multiple human contributors plus the release bot, 40 pull requests were merged from 2026-07-22 through 2026-08-21, and 22 issues were closed in the same period.<br>
> Release 9.31.2 credits five committers.

#### Issue Health

> The GitHub issue search showed 232 open issues on 2026-08-22, a substantial backlog, while 22 issues closed in the preceding 31-day window and several August reports were fixed and merged within days.<br>
> Maintainers and contributors actively reproduce and patch issues.<br>
> Current open reports include lost BiDi command responses, a session-creation retry gap, stale cross-context shadow-root state, and a ChromeDriver deadlock with network interception, so active triage does not remove near-term protocol reliability risk.<br>
> A repository-wide median response or resolution time was not calculated.

#### Roadmap And Deprecation Risk

> A published TSC-owned roadmap tracks component testing, core initiatives, network recording, frontend support, debugging, a VS Code extension, and related work.<br>
> WebdriverIO v9 makes WebDriver BiDi the default and permits fallback to classic WebDriver with `wdio:enforceWebDriverClassic`; BiDi-only features are unavailable on unsupported remote environments and Safari.<br>
> The core v9 line is stable, but rapid BiDi evolution and recent protocol issues create transition risk.<br>
> The separately versioned visual service v10 changed its comparison engine from ResembleJS to Pixelmatch, requiring visual baselines to be reaccepted even though its public API stayed stable.

#### Wrapper Upstream Lag

> WebdriverIO is not a thin Selenium-language wrapper: it ships its own JavaScript WebDriver and BiDi protocol client and follows W3C protocol work directly.<br>
> Driver and browser behavior can still lag upstream specifications, especially BiDi network and emulation primitives; Safari lacks required BiDi support for several features, and cloud grids must expose BiDi or CDP.<br>
> Recent releases show fixes landing within days for current-browser regressions, but no guaranteed upstream-adoption service level exists.

### Community Adoption

#### Npm Downloads

> The official npm downloads API recorded 1,544,931 downloads of @wdio/cli for the complete seven-day window 2026-08-15 through 2026-08-21, and 6,463,515 for 2026-07-22 through 2026-08-21.

#### Ecosystem Usage

> The monorepo maintains core runner, framework adapters, assertions, CLI, official reporters, and cloud services.<br>
> Documented integrations cover Appium, Selenium Grid, BrowserStack, Sauce Labs, TestingBot, TestMu AI, visual testing, axe-core, Allure, JUnit, JSON, Docker, and many community services/reporters.<br>
> An awesome-webdriverio catalog and multiple boilerplates add depth.<br>
> Sponsorship from BrowserStack, TestMu AI, TestingBot, SAP, and others demonstrates vendor investment, but sponsorship is not treated as proof that each sponsor runs WebdriverIO in production.

#### Community Support

> Versioned official documentation includes concepts, API references, migration material, examples, recipes, a blog, an official YouTube channel, and boilerplates.<br>
> Support channels include Discord, GitHub Issues and Discussions, the Selenium Slack community, and the Stack Overflow `webdriver-io` tag.<br>
> Community plugin documentation is surfaced in the official site but varies in freshness and quality.

#### Adoption Trend

> The comparable @wdio/cli 31-day npm window increased from 3,309,216 downloads in 2025-07-22 through 2025-08-21 to 6,463,515 in 2026-07-22 through 2026-08-21, approximately 95.3 percent growth.<br>
> GitHub activity and frequent releases are also current.<br>
> This supports a growing rather than declining adoption signal, while download growth may include CI reinstalls and transitive tooling use and is not a count of organizations.

#### Adoption Metric Normalization

> The primary metric is the exact @wdio/cli npm package, not the broader `webdriverio` automation-client package and not the sum of the many @wdio packages.<br>
> Windows are complete UTC registry periods: 2026-08-15..2026-08-21 for seven-day current usage and 2025-07-22..2025-08-21 versus 2026-07-22..2026-08-21 for year-over-year comparison.<br>
> Downloads include repeated CI installs and may include non-E2E uses; they neither equal unique users nor isolate this framework from all transitive installation.

### Browser And Runtime Coverage

#### Browser Engines

> WebdriverIO drives real Chrome/Chromium and Edge (Blink), Firefox (Gecko), and Safari (WebKit) through standards-based browser drivers, with mobile browsers and native apps available through Appium.<br>
> Chrome/Chromium and Firefox can be installed automatically.<br>
> Edge requires an installed browser or remote infrastructure, and native Safari requires macOS or a cloud grid.<br>
> There is no bundled Linux WebKit browser equivalent to Playwright WebKit, so Ubuntu cannot provide local Safari coverage.

#### Browser Protocol

> W3C WebDriver and WebDriver BiDi are first-class transports; v9 attempts a BiDi session by default and can force classic WebDriver.<br>
> CDP is used for Chromium-specific capabilities and can be reached through Puppeteer.<br>
> BiDi unlocks events, richer navigation, logs, mocks, and emulation, but support differs by browser and grid.<br>
> Safari does not yet support the BiDi emulation features documented by WebdriverIO, and current BiDi/ChromeDriver issue reports warrant testing both BiDi and classic paths.

#### Headless And Headed Modes

> Chrome/Chromium, Firefox, and Edge support headed local debugging and headless CI through browser capabilities.<br>
> The runner also supports Xvfb where needed.<br>
> Native Safari is normally run on macOS or remote infrastructure and does not provide an equivalent Ubuntu headless path.<br>
> Cloud vendors can supply additional headed or recorded sessions.

#### Browser Version Management

> WebdriverIO uses @puppeteer/browsers and driver packages to locate or download requested Chrome/Chromium and Firefox versions and compatible drivers. `browserVersion` pins a browser, driver-specific options can pin or supply driver binaries, and `WEBDRIVER_CACHE_DIR` or `cacheDir` controls the driver cache.<br>
> Edge browser installation is not automated; SafariDriver is provided by macOS.<br>
> Supplying only one of a custom browser and driver binary can create mismatches, so both should be pinned together.

#### Parallel Browser Support

> A capabilities array defines browser/version/platform combinations, and each capability can run specs in separate worker processes. `maxInstances`, `maxInstancesPerCapability`, and `wdio:maxInstances` cap global and per-browser concurrency.<br>
> CLI sharding (`--shard=x/y`) distributes specs across machines, and multiremote controls several sessions in one scenario.<br>
> Browser matrices and shard report aggregation remain configuration and CI responsibilities.

#### Mobile Emulation

> The BiDi `emulate('device', ...)` helper changes viewport, device scale factor, and user agent from maintained descriptors; specific geolocation, color scheme, and user-agent emulation are also available.<br>
> WebdriverIO explicitly warns that desktop-device emulation is not real mobile testing because mobile engines, UI, GPU, and hardware APIs differ.<br>
> Real Android/iOS browser coverage is available through Appium and local or cloud devices.

#### Real Browser Fidelity

> WebDriver controls installed vendor browsers rather than patched or embedded browser builds, which provides high fidelity for Chrome, Firefox, Edge, and native Safari.<br>
> Safari execution means SafariDriver on macOS or a real cloud Safari session; it is not inferred from a generic WebKit build.<br>
> This is a coverage advantage when macOS/cloud infrastructure is available, but makes fully local Ubuntu-only cross-engine testing impossible.

#### Environment Determinism Controls

> Core controls include window/viewport size, device scale and user agent through device emulation, geolocation, color scheme, online state, network throttling, and a fake clock covering Date and timer APIs.<br>
> These newer emulation helpers require BiDi and are unavailable in Safari.<br>
> Locale, timezone, permissions, random-number seeding, CSS animation suppression, font installation, and reduced-motion normalization are not one uniform cross-browser fixture API; use browser capabilities, preload scripts, application-independent CSS, container images, and repository helpers as applicable.

### Application Fit

#### Dynamic Dom Synchronization

> Element interaction commands automatically wait for visibility and interactability, and expect-webdriverio assertions retry until timeout.<br>
> This fits HTMX requests, lazy partials, and Handlebars rendering when assertions target the post-swap state.<br>
> A resolved WebDriver element can become stale when HTMX replaces its node, so page objects should expose selector getters or re-run `$`/`$$` after swaps rather than retain element IDs.<br>
> Explicit `waitForDisplayed`, `waitForExist`, `waitUntil`, and request-spy `waitForResponse` cover application-specific completion signals.

#### Routing Support

> Navigo is framework-agnostic to WebdriverIO.<br>
> Tests can navigate through `browser.url`, click links, inspect `getUrl`, use browser back/forward/refresh, and wait for URL or DOM state after History API changes.<br>
> Direct deep links work when the independently started Gin or Parcel server supplies the SPA fallback.<br>
> Route listener cleanup and application state must be asserted by tests; WebdriverIO has no Navigo-specific instrumentation.

#### Locator Model

> Supports CSS, XPath, visible-text, partial-text, accessible-name (`aria/...`), role, stable ID through CSS, `data-testid`, chained, custom, and shadow-DOM selectors.<br>
> Documentation favors user-visible text and accessible names. `$` returns the first match rather than enforcing strict uniqueness, so tests must assert cardinality where ambiguity matters.<br>
> Re-query after an HTMX replacement to avoid stale element references; selector getters make that inexpensive.

#### Form Interaction

> Standards-based click, setValue, addValue, clearValue, keys/actions, select, focus, and JavaScript execution exercise native inputs and event sequences.<br>
> Tests can press Tab to trigger blur, inspect focus, assert hidden synchronized fields, add/remove dynamic rows, and call or inspect native validity APIs.<br>
> Browser-driver differences and application-specific custom widgets require targeted helpers; script assignment should be avoided when real keyboard/focus behavior is under test.

#### Canvas And Download Support

> Chart.js can be validated semantically by executing browser JavaScript to inspect chart data, canvas dimensions, accessibility fallback, and exported data, with @wdio/visual-service used for a limited stabilized image check.<br>
> File downloads are configured per Chrome, Firefox, and Edge capability and verified by polling a worker-specific filesystem directory; Chromium can also use CDP download behavior.<br>
> Download completion, filename/content checks, cleanup, and cross-browser path differences require test harness code rather than an integrated download object.

#### Network And Api Access

> Node-side test hooks can seed REST fixtures with fetch or repository clients. `browser.mock` supplies request spies, response replacement, aborts, and `waitForResponse`; official docs still warn that full cross-browser mocking depends on CDP/BiDi primitives and provider support.<br>
> Yahoo Finance can be replaced reliably with an application-side stub server or Chromium interception, but Firefox/Safari interception must be proven before relying on one implementation.<br>
> Recent open Chrome/BiDi interception issues make a server-side stub the safer project-wide default.

#### Same Origin Support

> Tests operate normally against the consolidated Gin production server through one `baseUrl`, requiring no CORS changes.<br>
> They can also target Parcel's development server and existing proxy.<br>
> WebdriverIO itself does not alter origin policy; any cross-origin behavior remains real browser behavior unless a mock, proxy, or test server is configured.

#### Test Isolation

> The local runner starts a worker process and browser session per spec file/capability, not a fresh browser context per test.<br>
> Tests sharing one spec worker therefore share cookies, storage, and the session unless hooks clear state or use the slower `reloadSession`.<br>
> Unique data/schema identifiers and a disposable PostgreSQL database per suite or worker are compatible but custom.<br>
> Grouping specs in an array intentionally serializes them in one worker for stateful scenarios.

#### External Server Model

> `baseUrl` points tests at any independently started HTTP server and relative `browser.url` calls use it.<br>
> WebDriver endpoints, application host, ports, proxy, and cloud-grid connection details are configurable, so the same suite can target a locally composed stack, Parcel development server, or deployed Gin server.

#### Application Lifecycle

> WebdriverIO has lifecycle hooks and an official static-server service, but no integrated generic `webServer` declaration that builds Parcel, starts Docker Compose, waits for Flyway and API health, and always tears down the stack.<br>
> This repository should keep those steps in Make/GitHub Actions or add a small custom service/onPrepare-onComplete harness.<br>
> Readiness should poll a health endpoint, and teardown should run from an `always()` CI step; database reset and compose logs remain repository responsibilities.

#### Visual Regression Workflow

> The WebdriverIO-organization @wdio/visual-service v10, documented as a third-party package, creates and compares element, viewport, and full-page baselines across desktop and mobile targets.<br>
> It supports blocked regions, hidden text, thresholds, per-capability naming, actual/diff output, automatic baseline creation, and `--update-visual-baseline`; a visual HTML reporter is available.<br>
> Pin fonts, viewport, browser image, timezone, animations, and Chart.js animation completion, and maintain per-browser baselines.<br>
> The v10 Pixelmatch migration changes mismatch percentages and requires baseline review.

#### Accessibility Audit Integration

> The open-source @axe-core/webdriverio adapter from Deque works in runner or standalone mode and returns standard axe results for local assertions and reporting.<br>
> Accessible-name selectors and role/attribute assertions cover targeted semantics, and the visual service can inspect tab order.<br>
> There is no core Playwright-style ARIA snapshot facility; accessibility-tree snapshots and consolidated open-source reports require axe result serialization or another reporter.

### Reliability

#### Waiting Model

> Actions automatically wait for an element to be displayed and interactable; element `waitFor*`, general `waitUntil`, polling expect-webdriverio assertions, WebDriver page-load/script timeouts, and framework test timeouts provide explicit scopes.<br>
> BiDi events and request spies can wait for logs or responses.<br>
> There is no universal automatic wait for all HTMX network-idle and application-render completion, so tests should wait on a stable user-visible result or a specific request rather than sleep.

#### Flake Controls

> Controls include per-test retries from Mocha/Jasmine/Cucumber, `specFileRetries` with delay/deferred behavior, command connection retries, scoped timeouts, worker limits, deterministic emulation, watch mode, and CLI sharding.<br>
> Reproduction uses one spec/browser with `maxInstances: 1`, logs, DevTools traces, or repeated external invocation.<br>
> A current open issue reports that spec-file retries can be skipped when session creation itself fails, and recent BiDi issues show that retries cannot substitute for protocol stability.

#### Isolation Model

> One browser session belongs to each worker/spec grouping.<br>
> Fresh sessions between files provide coarse isolation; tests inside a file need cookie/storage cleanup, navigation reset, or `reloadSession`.<br>
> Stateful database scenarios can be grouped and run serially while unrelated suites run in other workers, but database transactions/schemas and cleanup are outside the framework.

#### Parallelism Controls

> Global, per-capability, and capability-local instance limits can set `maxInstances: 1` for database-sensitive suites while allowing other capabilities or future shards to run concurrently.<br>
> Spec arrays keep selected files in one worker; separate configuration files or suites can assign serial and parallel policies.<br>
> Cross-worker PostgreSQL safety still needs worker-specific data/schema identifiers or an external lock and deterministic reset.

#### Flake Observability

> Core and official reporters record failures and retries, JSON results can be merged within a run, and DevTools can preserve/rerun failures and compare command timelines.<br>
> WebdriverIO has no core quarantine registry, first-attempt-failure trend store, or long-term local dashboard.<br>
> Those require Allure/history tooling, JSON post-processing, labels/tags, or a community service.<br>
> A non-hosted repository implementation is possible but not turnkey.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Core commands and hooks can save screenshots and driver logs.<br>
> The WebdriverIO DevTools service adds timestamped command screenshots, browser preview, console and network logs, per-session WebM video, and portable trace.zip replay; other video/timeline/HTML reporters exist.<br>
> Artifact-on-failure policy, output directories, retention, and CI upload must be configured.<br>
> Backend and Docker logs are separate artifacts.

#### Debugging Tools

> Headed runs, `browser.debug()` REPL, standalone REPL, Node inspector, VS Code launch configurations, log levels, and single-worker/single-spec execution are documented.<br>
> The DevTools UI adds live previews, command timelines, network/console inspection, selective reruns, preserve-and-compare, source navigation, and offline traces.<br>
> This is strong but service-based rather than all enabled by the base runner.

#### Test Generation

> Chrome DevTools Recorder can export a flow through the WebdriverIO Chrome Recorder extension or @wdio/chrome-recorder.<br>
> Generated scripts are executable but commonly contain long structural CSS selectors, and official guidance says to replace them with resilient text, accessible-name, or stable-attribute selectors.<br>
> Generation is Chrome-oriented rather than a framework-wide cross-browser recorder.

#### Reporters

> Official/core packages cover dot, spec, concise, JUnit XML, JSON, Allure, and custom reporters; JSON output from parallel sessions has a merge utility.<br>
> HTML, timeline, video, ReportPortal, and additional GitHub-compatible workflows are available through official-adjacent or community packages.<br>
> Plain GitHub annotations require a compatible reporter or workflow conversion; report aggregation across shards is not automatic.

#### Documentation Quality

> The current v9 site is extensive, versioned, searchable, multilingual, and includes architecture concepts, configuration/API references, examples, recipes, migration guides, security guidance, and a published roadmap.<br>
> Documentation was actively improved in August 2026 with machine-readable output.<br>
> Quality is lower at ecosystem edges: many service/reporter pages mirror third-party READMEs, some examples are old, and protocol support notes can lag rapidly changing BiDi behavior.

#### Local Workflow

> Run all tests with `npx wdio run ./wdio.conf.ts`, one file with `--spec`, or a configured suite with `--suite`; select capabilities through config/CLI overrides.<br>
> Watch mode reruns changed specs, `browser.debug()` opens a REPL, headed capabilities support visual debugging, and `maxInstances: 1` plus Node inspector supports breakpoints.<br>
> Repeating only failures needs framework retries, DevTools selective rerun, a community rerun service, or an external command loop.

#### Failure Log Correlation

> WebdriverIO logs commands and driver traffic with configurable levels and timestamps; BiDi can capture browser console/network events, and DevTools aligns command, screenshot, console, and network evidence.<br>
> Gin, Flyway, PostgreSQL, and Docker Compose output are not ingested automatically.<br>
> The CI harness must preserve service logs with timestamps/run IDs and upload them beside WDIO artifacts to correlate backend and browser failures.

#### Artifact Data Exposure

> Screenshots, video, trace archives, network bodies, browser console, driver logs, reports, and downloaded files may expose portfolio values, API responses, session tokens, or credentials.<br>
> WebdriverIO supports `maskingPatterns`/`WDIO_LOG_MASKING_PATTERNS` and masked setValue/addValue logging, and visual comparisons can hide regions/text.<br>
> Generic screenshots, video, HAR-like network evidence, and trace content are not automatically redacted; use synthetic data, least-privilege credentials, restricted artifact access, short retention, and avoid recording secrets.

### Github Actions Fit

#### Official Ci Support

> Official documentation and maintained boilerplates show GitHub Actions and Ubuntu execution, including a matrix sharding example.<br>
> There is no dedicated official WebdriverIO GitHub Action or WebdriverIO-owned all-browser container image; workflows use Node/npm, installed browsers or Selenium images, and standard GitHub actions.<br>
> This is adequate but less turnkey than a framework-supplied CI image/action pair.

#### Browser Caching

> Use the npm cache for packages and set `WEBDRIVER_CACHE_DIR`/driver `cacheDir` to a stable path for downloaded drivers.<br>
> GitHub-hosted Ubuntu often provides Chrome and Firefox, reducing browser downloads; explicit versions or Selenium images improve determinism.<br>
> If caching downloaded browsers/drivers, key by OS, architecture, browser and driver version, and refresh deliberately.<br>
> WebdriverIO does not provide a dedicated cache action.

#### Artifact Integration

> Write screenshots, traces, videos, JSON/JUnit/Allure/HTML output, and logs beneath known directories, then upload them with `actions/upload-artifact` under `if: failure()` or `if: always()`.<br>
> Official sharding documentation demonstrates per-shard log upload.<br>
> Naming and retention policy, sensitive-data controls, and cross-job report assembly are workflow configuration.

#### Sharding And Matrix Support

> Native `--shard=x/y` partitions specs across jobs and official docs pair it with a GitHub strategy matrix.<br>
> Browser capabilities can be another matrix axis or run within each WDIO process.<br>
> Retries happen inside a shard; re-running a failed GitHub job reruns that shard.<br>
> JSON files can be merged, but gathering artifacts and producing one report across jobs requires an aggregation job and reporter-specific tooling.

#### Container Compatibility

> Official guidance runs WDIO in Selenium standalone browser images, and community Docker services can start containerized dependencies.<br>
> It can run beside this repository's Docker Compose stack on GitHub-hosted Ubuntu.<br>
> The workflow must choose whether the browser runs on the host or a Selenium container, expose the application on a reachable interface, use service/container DNS rather than assuming localhost across containers, pin matching browser/driver versions, and allocate shared memory.

#### Failure Cleanup

> The runner normally deletes sessions and stops locally managed drivers.<br>
> Application, Flyway, PostgreSQL, and Compose teardown are not owned by WebdriverIO; use workflow steps guarded by `if: always()`, timeouts, unique Compose project names, and volume cleanup.<br>
> GitHub cancellation can prevent ordinary process hooks from completing, so job-level container cleanup and disposable runners/databases are safer than relying only on `onComplete`.

### Cost And Risk

#### Open Source Completeness

> Core runner, assertions, local browsers, reporters, sharding, axe integration, visual service, and local DevTools diagnostics are available as open-source packages.<br>
> A paid cloud is not required for Chrome/Firefox E2E on Ubuntu.<br>
> Native Safari, broad OS/device matrices, and hosted history dashboards require owned infrastructure or optional providers, but are not prerequisites for the repository's base local/CI capability.

#### Optional Cloud Dependency

> BrowserStack, Sauce Labs, TestingBot, TestMu AI, and other clouds are optional for device/Safari/browser-version breadth, tunnels, provider video, and hosted analytics.<br>
> Equivalent basic local execution, screenshots, traces, reports, sharding, visual comparison, and axe audits are available without them.<br>
> Native Safari still requires macOS somewhere, so an Ubuntu-only organization needs a macOS runner or cloud for that gate.

#### Migration Cost

> Medium.<br>
> Tests use JavaScript/TypeScript, async/await, standards-oriented selectors, and WebDriver concepts that transfer conceptually, but `browser`, `$`, expect-webdriverio matchers, WDIO configuration, hooks, services, and reporter APIs are framework-specific.<br>
> Page objects can isolate much of the command API.<br>
> Cloud capabilities and visual baselines increase lock-in, while plain REST fixtures and semantic locators reduce it.

#### Security And Supply Chain

> The project publishes a security policy and threat model, requests private reports by email with a 48-hour response target, and offers OpenJS CNA escalation. @wdio/cli 9.31.2 has npm registry signatures and SLSA provenance attestation from a GitHub OIDC trusted publisher.<br>
> Risk remains from a large modular dependency/plugin surface, dynamically loaded services/reporters, runtime browser/driver downloads, and test code with host/network access.<br>
> Pin versions and lockfiles, use `npm ci`, verify provenance, scan dependencies, restrict CI permissions/egress, and prefer reviewed @wdio packages over unvetted plugins.

#### Custom Harness Burden

> Moderate to high for this repository.<br>
> Browser lifecycle, runner, waiting, reporters, and sharding are supplied.<br>
> Custom work remains for Parcel/Gin build/startup, Docker Compose and Flyway readiness, disposable PostgreSQL data, API seeding, Yahoo Finance replacement that works beyond Chromium, worker-safe download directories, browser/backend log collection, and guaranteed teardown.<br>
> Hooks and custom services provide extension points but do not remove that code.

#### Capability Delivery Tier

> Core: WebDriver/BiDi automation, CLI/local runner, async commands, auto-wait actions, assertions, capabilities, worker parallelism, sharding, retries, screenshots, driver logging, baseUrl, browser/driver management, request spying/mocking where protocols support it, and emulation.<br>
> Official package: framework adapters, JUnit/JSON/Allure reporters, static server, cloud services, shared store, and DevTools trace/live debugging.<br>
> WebdriverIO-organization package documented as third party: @wdio/visual-service.<br>
> External open-source integration: @axe-core/webdriverio.<br>
> Community/custom: generic Docker Compose lifecycle, database reset, cross-browser Yahoo stub server, HTML/video alternatives, quarantine/trend reporting, backend-log correlation, and final artifact aggregation.<br>
> Paid cloud: optional real-device/native-Safari breadth and hosted dashboards, not required for base CI.

#### Ai Execution Boundary

> Core test authoring and CI execution are deterministic and do not call an LLM.<br>
> Recorder, configuration wizard, normal selectors, and all required assertions have non-AI paths.<br>
> WebdriverIO's optional MCP/AI assistance or documentation copilot should remain outside CI, with no production credentials or portfolio data sent to an external model.<br>
> Disabling AI features does not remove browser automation, reports, traces, or retries.

### Evidence And Decision

#### Sources

- Official getting started and Node requirements: https://webdriver.io/docs/gettingstarted/
- Official TypeScript setup: https://webdriver.io/docs/typescript/
- Official capabilities and driver management: https://webdriver.io/docs/capabilities/
- Official driver binaries guide: https://webdriver.io/docs/driverbinaries/
- Official runner architecture: https://webdriver.io/docs/runner/
- Official auto-waiting guide: https://webdriver.io/docs/autowait/
- Official selectors guide: https://webdriver.io/docs/selectors/
- Official request mocks and spies: https://webdriver.io/docs/mocksandspies/
- Official emulation guide: https://webdriver.io/docs/emulation/
- Official file-download guidance: https://webdriver.io/docs/best-practices/file-download/
- Official browser-log guidance: https://webdriver.io/docs/best-practices/browser-logs/
- Official visual-testing guide: https://webdriver.io/docs/visual-testing/
- Official axe-core integration guide: https://webdriver.io/docs/accessibility-testing/axe-core/
- Official DevTools diagnostics guide: https://webdriver.io/docs/devtools/
- Official sharding guide: https://webdriver.io/docs/sharding/
- Official GitHub Actions guide: https://webdriver.io/docs/githubactions/
- Official Docker guide: https://webdriver.io/docs/docker/
- Official security policy and threat model: https://github.com/webdriverio/webdriverio/blob/main/.github/SECURITY.md
- Official governance: https://github.com/webdriverio/webdriverio/blob/main/GOVERNANCE.md
- Official roadmap: https://github.com/webdriverio/webdriverio/blob/main/ROADMAP.md
- Official team list: https://github.com/webdriverio/webdriverio/blob/main/AUTHORS.md
- Latest release v9.31.2: https://github.com/webdriverio/webdriverio/releases/tag/v9.31.2
- Repository metadata and activity: https://github.com/webdriverio/webdriverio
- npm latest @wdio/cli metadata: https://registry.npmjs.org/%40wdio%2Fcli/latest
- npm downloads, seven days: https://api.npmjs.org/downloads/point/2026-08-15:2026-08-21/%40wdio%2Fcli
- npm downloads, current 31 days: https://api.npmjs.org/downloads/point/2026-07-22:2026-08-21/%40wdio%2Fcli
- npm downloads, prior-year 31 days: https://api.npmjs.org/downloads/point/2025-07-22:2025-08-21/%40wdio%2Fcli
- GitHub open issue search observed 2026-08-22: https://api.github.com/search/issues?q=repo%3Awebdriverio%2Fwebdriverio%20type%3Aissue%20state%3Aopen
- GitHub recently closed issue search: https://api.github.com/search/issues?q=repo%3Awebdriverio%2Fwebdriverio%20type%3Aissue%20closed%3A2026-07-22..2026-08-21
- GitHub recently merged pull-request search: https://api.github.com/search/issues?q=repo%3Awebdriverio%2Fwebdriverio%20type%3Apr%20merged%3A2026-07-22..2026-08-21

#### Observed At

2026-08-22

#### Confidence

> High confidence for release, license, governance, npm, repository, installation, protocol, runner, and documented capability facts because they come from official documentation and live registry/GitHub APIs.<br>
> Medium confidence for application-fit and CI judgments because they are architectural inferences from the repository requirements.<br>
> Low confidence for TypeScript 6 edge compatibility, practical maintainer concentration, resource cost, and the final comparative score because no project spike or benchmark was run.

#### Deal Breakers

> No confirmed exclusion-level incompatibility for the Open Asset Allocator stack.<br>
> Potential deal breakers are: a requirement for local Safari/WebKit on Ubuntu; a requirement for equally reliable cross-browser request interception without a stub server; refusal to maintain custom Compose/PostgreSQL lifecycle code; or a project spike reproducing the current BiDi session/interception failures.<br>
> TypeScript 6 incompatibility would also exclude it, but none has been demonstrated.

#### Recommendation

> Viable alternative.<br>
> WebdriverIO is actively maintained, widely downloaded, standards-based, capable of real Safari through macOS/cloud, and able to test the complete HTMX/Navigo/Gin application without changing Parcel.<br>
> It should not be selected as the default until a project spike proves Node 24 plus TypeScript 6, HTMX element replacement, Navigo deep links, Yahoo replacement, PostgreSQL reset, downloads, and failure traces on Ubuntu.<br>
> Prefer a local Yahoo stub over protocol interception and retain the ability to force classic WebDriver while current BiDi reliability issues are evaluated.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `dependency_currency`
- `maintainer_concentration`
- `github_metrics`
- `resource_usage`
- `hard_gate_result`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="selenium-webdriver-with-a-javascript-test-runner"></a>
## 4. Selenium WebDriver with a JavaScript test runner

Source result: `Selenium_WebDriver_with_a_JavaScript_test_runner.json`

### Project And Compatibility

#### Implementation Language

> The browser client is the selenium-webdriver JavaScript package running on Node.js; tests can be authored in JavaScript or TypeScript and use native promises with async/await.<br>
> The broader Selenium monorepo is cross-language, Selenium Manager is implemented in Rust, browser drivers are vendor executables, and the selected runner baseline is Mocha on Node.js.

#### Operating System Support

> Supported locally on Windows, macOS, and Linux.<br>
> The default x64 GitHub-hosted Ubuntu runner is suitable for Chrome/Chromium, Firefox, and Edge sessions.<br>
> Selenium Manager ships Windows, Linux, and macOS binaries, but its official Linux ARM64 binary/browser management remains limited; official docker-selenium images support amd64 broadly and ARM64 for Chromium and Firefox, not Google Chrome or Edge.<br>
> Native Safari requires macOS and cannot run on Ubuntu.

#### License And Governance

> Selenium and selenium-webdriver use Apache-2.0; Mocha uses MIT.<br>
> Both permit private and open-source repository use without a paid service.<br>
> Selenium is a community project under the Software Freedom Conservancy with published Project Leadership and Technical Leadership committees, committer roles, public governance, sponsors, and open contribution processes.<br>
> Commercial browser-cloud vendors sponsor or integrate with Selenium but do not control access to the local core stack.

#### Installation Model

> A minimal local stack is npm install --save-dev selenium-webdriver mocha, plus an assertion library if Node's assert is not used and likely @types/selenium-webdriver for TypeScript.<br>
> Selenium Manager is bundled and resolves, downloads, and caches compatible drivers when they are absent; it can also manage requested browser versions.<br>
> Browsers and their Linux system libraries must still be installed or downloaded.<br>
> An optional remote setup uses a pinned official selenium/standalone-chrome, standalone-firefox, or Grid image from Docker Hub or GHCR; Grid adds Java and container resources but is not required for local execution.

#### Candidate Scope And Layer

> A standards-based composed stack: Selenium WebDriver is a browser automation client and Grid is optional remote infrastructure, while Mocha supplies test discovery, hooks, retries, parallel workers, and reporting.<br>
> It is not an integrated E2E framework and requires a project-owned harness for lifecycle, fixtures, assertions, artifacts, and configuration.

#### Authoring And Async Model

> Tests use ordinary JavaScript or TypeScript, native async/await, WebDriver objects, locators, and Mocha describe/it plus before/after hooks.<br>
> There is no command queue or required DSL.<br>
> Every asynchronous WebDriver operation must be awaited, and project helpers or page objects are advisable to centralize waits, reacquisition, browser options, and diagnostics.

#### Build Pipeline Coupling

> Black-box Selenium sessions navigate to an HTTP URL and do not transform or instrument frontend source.<br>
> They can test the Parcel development server through its existing proxy or the production Gin server serving the built Parcel output.<br>
> E2E code can be a separate Node test package or dev-dependency set, so no Vite migration, Parcel plugin, or alternate frontend build is required.

#### Testability Instrumentation Required

> No production script injection, relaxed content-security policy, framework-specific build hook, or special browser runtime is required.<br>
> Stable unique IDs or data-test attributes would materially improve locator durability but are not mandatory.<br>
> API fixture/reset endpoints, deterministic test data, and an external-market-data stub may be needed for reliable application tests; these are application testability facilities rather than Selenium requirements and should be disabled or protected outside test environments.

### Maintenance Health

#### Latest Stable Release

> selenium-webdriver 4.47.0, released 2026-08-10.<br>
> The selected runner baseline observed was Mocha 11.8.0.

#### Release Cadence

> Selenium JavaScript releases were frequent and regular: 4.40.0 on 2026-01-18, 4.41.0 on 2026-02-20, 4.42.0 on 2026-04-09, 4.43.0 on 2026-04-10, 4.44.0 on 2026-05-12, 4.45.0 on 2026-06-16, 4.46.0 on 2026-07-11, and 4.47.0 on 2026-08-10.<br>
> Releases include browser/protocol compatibility, Selenium Manager, Grid, and language-binding work rather than only documentation changes.

#### Repository Activity

> SeleniumHQ/selenium was pushed on 2026-08-22.<br>
> The ten most recent commits inspected were dated 2026-08-21 or 2026-08-22 and covered build/release reliability, BiDi schemas, browser extensions, Ruby, Python, and flake recording.<br>
> Multiple pull requests were integrated on those dates.<br>
> This is high current activity in the shared monorepo, although not every commit affects the JavaScript binding.

#### Roadmap And Deprecation Risk

> The project is actively moving functionality from WebDriver Classic and temporary CDP access toward the W3C WebDriver BiDi standard while attempting backward compatibility.<br>
> Selenium 4.47 published a Selenium 5 release charter and a BiDi implementation-boundary ADR.<br>
> Browser-vendor protocol transitions can introduce binding changes, and CDP APIs are explicitly temporary and version-sensitive.<br>
> The standards-based WebDriver surface lowers long-term lock-in, but advanced BiDi APIs are still expanding and need upgrade testing.

#### Dependency Currency

> selenium-webdriver 4.47.0 was published using Node 24.14.1, requires Node >=22, and has four direct npm dependencies.<br>
> It follows current Node LTS/Current support policy and released ten days before observation.<br>
> Selenium Manager tracks browser and driver compatibility, while official Grid images publish pinned browser/driver combinations.<br>
> TypeScript declarations are maintained separately on DefinitelyTyped, which is the main currency concern for typed authoring.

#### Wrapper Upstream Lag

> The JavaScript binding is released from the Selenium monorepo, so there is no third-party Selenium wrapper lag.<br>
> Selenium Manager resolves vendor drivers independently and can match installed/requested browsers.<br>
> CDP support is limited to recent Chrome versions and can require Selenium updates; WebDriver BiDi is intended to remove that vendor-version coupling.<br>
> A measurable wrapper gap exists for TypeScript declarations: @types/selenium-webdriver was 4.35.6 while the runtime was 4.47.0 at observation.

### Community Adoption

#### Npm Downloads

> The official npm downloads API reported 1,820,837 direct downloads of selenium-webdriver for 2026-08-14 through 2026-08-20.

#### Ecosystem Usage

> Selenium is the established cross-language WebDriver baseline with official bindings, Grid, docker-selenium, browser-vendor driver support, hosted-grid integrations, page-object guidance, and extensive third-party reporters and utilities.<br>
> Maintained integrations include official Selenium Docker images and Deque's @axe-core/webdriverjs.<br>
> Its broad use spans functional testing and automation, so ecosystem size is strong even though a minimal JavaScript stack must select and maintain its own companion packages.

#### Community Support

> Official documentation covers WebDriver, JavaScript APIs, Selenium Manager, Grid, BiDi, browser options, troubleshooting, and test practices.<br>
> Support channels include GitHub issues, Selenium Chat/community channels, conferences, and a large body of Stack Overflow and vendor material.<br>
> Documentation quality varies by language tab, and current advanced examples are not always equally complete for JavaScript.

#### Adoption Metric Normalization

> The download metric is for the direct npm package selenium-webdriver over the exact seven-day window 2026-08-14 through 2026-08-20.<br>
> It includes E2E tests, libraries, transitive installs, bots, CI, and non-test browser automation and does not identify Mocha usage.<br>
> GitHub metrics are for the SeleniumHQ/selenium monorepo across Java, JavaScript, Python, Ruby, .NET, Grid, Manager, and shared code; npm dependents are package-level and are not equivalent to production repositories or users.

### Browser And Runtime Coverage

#### Browser Engines

> Selenium drives installed Google Chrome and Microsoft Edge (Chromium), Mozilla Firefox (Gecko), and native Apple Safari through their vendor drivers; it also supports Chromium variants through browser-specific options.<br>
> There is no bundled generic WebKit engine.<br>
> Linux/GitHub Ubuntu can cover Chrome/Chromium, Edge, and Firefox; genuine Safari coverage requires a macOS machine or remote macOS service.<br>
> Legacy Internet Explorer support exists but is irrelevant to this project.

#### Browser Protocol

> Core automation uses the W3C WebDriver protocol through vendor drivers, locally or through Grid.<br>
> WebDriver BiDi adds a standards-based WebSocket channel for console, script, browsing-context, and network events.<br>
> Selenium also exposes temporary Chrome DevTools Protocol access for Chromium-specific gaps, but CDP is explicitly version-sensitive and not cross-browser.<br>
> Advanced capabilities must be tested against each target browser's current BiDi implementation.

#### Headless And Headed Modes

> Chrome/Edge and Firefox support both normal headed sessions for local debugging and vendor headless modes for CI, configured through browser arguments such as --headless=new for Chrome.<br>
> Safari execution is headed/native and platform constrained.<br>
> The same tests can select mode through project configuration; Selenium itself does not provide an interactive runner UI.

#### Browser Version Management

> Selenium Manager discovers installed browsers, obtains compatible drivers, downloads requested browser versions where supported, and caches artifacts.<br>
> Browser versions can be requested through options, while official Docker image tags pin browser, driver, Grid, Java, and OS combinations.<br>
> Reproducibility requires pinning selenium-webdriver and image/browser versions; relying on preinstalled or auto-updated host browsers allows environmental drift.<br>
> Linux ARM64 has reduced browser availability.

#### Parallel Browser Support

> Each WebDriver object represents one browser session.<br>
> Mocha can run test files in a worker-process pool with --parallel and --jobs, GitHub Actions can provide a browser matrix, and Selenium Grid can schedule concurrent sessions across browser/version/platform nodes.<br>
> Browser projects, sharding policy, database-safe grouping, and result merging are custom configuration rather than first-class Selenium JavaScript concepts.

#### Mobile Emulation

> Chromium options can set device metrics, viewport, touch, and user agent for responsive/mobile emulation, and generic window sizing covers viewport tests.<br>
> Coverage is browser-specific and does not equal a real mobile browser or device.<br>
> Real Android/iOS testing normally adds Appium or a device cloud, which is outside this candidate's core stack.

#### Real Browser Fidelity

> Selenium's principal strength is controlling installed vendor browsers through their native drivers rather than a framework-bundled approximation.<br>
> Chrome, Edge, Firefox, and Safari results therefore reflect those actual products, subject to automation flags and headless differences.<br>
> Native Safari on macOS is materially different from a WebKit build on Linux; Selenium correctly requires the former instead of claiming Linux WebKit as Safari coverage.

#### Environment Determinism Controls

> Window size, device metrics, user agent, language arguments/preferences, browser profiles, certificates, page-load strategy, and timeouts can be configured.<br>
> Permissions, geolocation, locale, timezone, and network behavior are available to varying degrees through browser options, CDP, or expanding BiDi APIs, but support is not uniform across JavaScript/browser combinations.<br>
> There is no integrated cross-browser clock, randomness, animation, or reduced-motion fixture; tests must use browser preferences, application APIs, executeScript, CSS injection, environment variables, or server-side controls and document browser-specific gaps.

### Application Fit

#### Dynamic Dom Synchronization

> Selenium waits for configured document readiness on direct navigation and performs basic interactability checks during element actions, but it does not automatically retry every locator/action/assertion until an HTMX swap settles.<br>
> HTMX requests, lazy partials, and asynchronous Handlebars rendering require explicit conditions such as elementLocated, elementIsVisible, text/url predicates, or custom waits.<br>
> Stored element references become stale when HTMX replaces nodes and must be reacquired, making locator-returning helpers important.

#### Routing Support

> WebDriver can navigate directly to deep links, read URLs, wait for URL/title/DOM conditions, use back/forward/refresh, and inspect window history through script execution.<br>
> Navigo pushState/popstate flows are normal browser behavior.<br>
> Reliability depends on a custom route-ready condition after client navigation and cleanup that opens a fresh session or resets route/storage between tests; Selenium has no Navigo-specific integration.

#### Locator Model

> Built-in strategies include ID, name, CSS, XPath, link text, tag name, and relative locators.<br>
> Stable unique IDs and compact CSS selectors are officially encouraged.<br>
> Unlike modern integrated runners, core Selenium has no first-class strict getByRole, getByLabel, or resilient visible-text locator API; accessible locators require CSS/XPath conventions or helper libraries, duplicate matches are not rejected by default, and HTMX-replaced elements must be located again.

#### Form Interaction

> WebDriver click, clear, sendKeys, active-element switching, and the W3C Actions API exercise browser input, focus, keyboard, pointer, wheel, and touch-like behavior.<br>
> Blur can be produced with Tab or a real click, and native validation state/properties can be asserted.<br>
> Hidden inputs are intentionally not user-interactable; their synchronized values can be read as DOM properties, while executeScript should be reserved for inspection because setting values by script bypasses user fidelity.<br>
> Dynamic rows require explicit waits and fresh locators.

#### Canvas And Download Support

> Chart.js canvas can be validated through application-visible labels/data, DOM/ARIA summaries, Chart.js state exposed to executeScript, and targeted screenshots; Selenium only supplies screenshots, not semantic Chart.js or pixel-diff assertions.<br>
> Downloads can be triggered, configured to a directory, and handled through Grid's downloadable-files capability, but official guidance notes that core WebDriver exposes no progress wait and recommends using an HTTP client with session cookies for file-content verification.<br>
> Robust completion checks, checksums, cleanup, and visual baselines are custom.

#### Network And Api Access

> Tests can seed and reset REST fixtures with Node fetch or another HTTP library alongside browser steps.<br>
> WebDriver BiDi can observe requests and console events and provides growing request/response interception and authentication handling; CDP remains a Chromium-only fallback.<br>
> A controlled Yahoo Finance response can be supplied by an application-side stub/proxy or BiDi interception where target browsers support it.<br>
> There is no integrated APIRequest fixture, route registry, HAR mocking layer, or guaranteed identical interception behavior across all browsers.

#### Same Origin Support

> Because Selenium operates outside the page and navigates to a configured URL, it can test the consolidated Gin production origin or Parcel's development origin/proxy without injecting code or changing CORS.<br>
> Tests should use the externally visible origin consistently.<br>
> Remote browser containers must address the application by a reachable host/service name rather than localhost inside the browser container.

#### Test Isolation

> Selenium recommends a new WebDriver session per test; vendor drivers normally create a clean browser profile, providing cookies/storage isolation at the cost of startup time.<br>
> It has no lightweight browser-context primitive, so sharing a session is faster but risks leakage.<br>
> Disposable PostgreSQL state, unique fixture identifiers, transaction/reset APIs, and worker-to-database allocation must be implemented by the project; Mocha parallel workers do not coordinate database access.

#### External Server Model

> The driver naturally targets an independently started application through any HTTP base URL.<br>
> A BASE_URL environment variable and helper are sufficient for local, production-build, Docker Compose, or remote deployments.<br>
> Selenium and Mocha do not include a built-in webServer process manager, port allocator, readiness probe, or baseURL fixture.

#### Application Lifecycle

> Mocha global fixtures/hooks or repository scripts can run the build, invoke Docker Compose, wait for Flyway/backend readiness, and tear services down.<br>
> The repository's Compose base already gates the backend on a healthy PostgreSQL service and successful migration service, which helps.<br>
> HTTP readiness polling, per-worker database disposal, signal handling, log capture, and guaranteed docker compose down --volumes in CI remain custom harness/workflow code.

#### Visual Regression Workflow

> Core Selenium captures page and element screenshots but has no baseline store, update/review command, masking, tolerance engine, per-browser snapshot convention, or HTML diff viewer.<br>
> An open-source image-diff library or a separate tool such as BackstopJS can be composed; commercial services such as Applitools are optional.<br>
> The project must stabilize fonts, viewport, device scale, animations, data, and Chart.js rendering and maintain OS/browser-specific baselines.

#### Accessibility Audit Integration

> Deque's maintained open-source @axe-core/webdriverjs accepts a Selenium WebDriver instance and can audit documents and frames, so axe violations can be asserted and emitted as JSON.<br>
> This is a strong community integration, not Selenium core.<br>
> Accessibility-tree/ARIA snapshot assertions and polished open-source report formatting require additional libraries or custom output, and automated axe checks do not replace manual accessibility testing.

### Reliability

#### Waiting Model

> Direct navigation waits for a selected document.readyState strategy, element actions perform limited interactability checks, implicit waits affect element lookup, and explicit driver.wait conditions can poll visibility, presence, staleness, text, title, URL, alerts, or custom predicates.<br>
> The default implicit wait is zero.<br>
> Selenium explicitly warns not to mix implicit and explicit waits.<br>
> There is no locator-wide auto-wait plus retrying assertion layer, so project helpers must encode HTMX request/swap and application-ready conditions.

#### Flake Controls

> Mocha provides suite/test timeouts, retries, serial-by-default execution, grep filters, and repeated invocation through scripts.<br>
> Selenium provides explicit synchronization primitives, page/script/implicit timeouts, fresh-session guidance, and Grid/browser logs.<br>
> There is no built-in trace replay, automatic action/assertion retry, deterministic clock, retry classification, or repeat-each option; fixed sleeps should be prohibited and flaky reproduction loops must be scripted.

#### Isolation Model

> The safest model is a fresh browser session and profile per test, paired with unique or reset database state.<br>
> Serial suites can intentionally share domain state, but browser cookies, localStorage, IndexedDB, windows, and input state must then be cleared explicitly.<br>
> Mocha hooks can enforce quit in finally/afterEach, while database serial groups should run outside Mocha parallel mode or use an explicit lock/allocation strategy.

#### Parallelism Controls

> Mocha is serial by default and offers --parallel plus a --jobs limit at test-file granularity; this makes jobs=1 a simple database-safe default.<br>
> Browser/version matrices and future sharding can be implemented in GitHub Actions or Grid, with unique database/schema allocation per job.<br>
> Parallel mode has reporter and .only limitations and shares process-level state between files reused by a worker, so safe grouping and cleanup are project responsibilities.

#### Flake Observability

> Mocha can expose currentRetry and reporter events, allowing first-attempt failures and eventual passes to be recorded, and Selenium artifacts can be captured on each failed attempt.<br>
> However, neither core package supplies retry classification, quarantine metadata, historical trend storage, or a local flake dashboard.<br>
> Repeat runs, JSON/JUnit enrichment, first-failure artifact preservation, and trend aggregation must be added to the harness or CI.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Core WebDriver provides screenshots, element screenshots, page source/script inspection, driver logs, and BiDi console/JavaScript/network events where supported.<br>
> Official docker-selenium can add session video, but it requires a separate FFmpeg container, substantial resources, and does not support headless video recording.<br>
> Automatic per-failure collection, HAR generation, trace timelines, artifact naming, retention, and attachment to a test result are not integrated and must be implemented.

#### Debugging Tools

> Headed mode, browser developer tools, breakpoints through Node's inspector, Mocha inspect, longer/pause waits, and VNC/noVNC for official browser containers are available.<br>
> A developer can run a single test and inspect the live browser.<br>
> There is no Selenium-native step inspector, time-travel replay, trace viewer, or automatic pause-on-failure experience comparable to an integrated E2E runner.

#### Test Generation

> Selenium WebDriver core has no recorder.<br>
> The separate Selenium IDE browser extension can record browser interactions and export code, but generated locators, waits, fixtures, and page objects require review and refactoring before production use.

#### Reporters

> Mocha includes console and machine-readable reporters such as spec, dot, TAP, JSON, JSON stream, and XUnit-style output and supports custom reporters.<br>
> Common community additions include mocha-junit-reporter for JUnit and Mochawesome for HTML.<br>
> GitHub annotations, merged shard reports, screenshots, logs, and retry metadata require reporter selection or custom integration; some third-party reporters are incompatible with Mocha parallel mode.

#### Documentation Quality

> Selenium has extensive maintained documentation, generated JavaScript API references, troubleshooting pages, examples, release notes, and architectural guidance; Mocha documents hooks, TypeScript, ESM, retries, parallel mode, and CLI behavior.<br>
> The main weakness is fragmentation across language tabs, WebDriver Classic, CDP, BiDi, Grid, Docker, and runner-specific material, with some advanced capabilities better documented in languages other than JavaScript.

#### Local Workflow

> Developers can run one file with npx mocha path/to/test, filter with --grep, select a browser through project config or SELENIUM_BROWSER, run headed by default, enable headless arguments for CI, and debug with mocha inspect or Node --inspect-brk.<br>
> Mocha watch mode exists but has native ESM limitations.<br>
> Repeating only failures and preserving a failed browser require scripts/hooks; there is no interactive test list or dedicated UI.

#### Failure Log Correlation

> BiDi can stream browser console, JavaScript, and network events with test-controlled timestamps, while Mocha hooks can attach the current test ID.<br>
> Backend, Flyway, PostgreSQL, and Docker Compose logs can be captured with timestamps in an always-running post step.<br>
> Selenium does not correlate these sources automatically; the harness must define a run/test correlation ID, clock basis, log boundaries, file naming, and collection after startup or test failure.

#### Artifact Data Exposure

> Screenshots, page source, video, network bodies/headers, console logs, downloadable files, browser profiles, and serialized storage can expose portfolio data, API responses, cookies, authorization headers, and credentials.<br>
> Selenium offers selective event handling but no universal artifact redaction pipeline.<br>
> CI must avoid persisting profiles/storage state, redact headers and bodies before writing logs, mask sensitive UI regions where possible, use synthetic data, limit retention/access, and never upload unreviewed HAR or full traces by default.

### Github Actions Fit

#### Official Ci Support

> Selenium documents Grid as suitable for CI/CD, maintains official Selenium Manager binaries and official Docker/Grid images on Docker Hub and GHCR, and supplies headless browser examples.<br>
> There is no required or comprehensive Selenium-specific GitHub Action; a normal workflow uses setup-node, npm ci, app startup/readiness, and npm test, optionally with a pinned Grid service/container.<br>
> GitHub-hosted Ubuntu x64 is a supported practical target.

#### Browser Caching

> npm's cache can cover JavaScript packages, Selenium Manager caches drivers and managed browsers under its cache, and pinned Docker images benefit from runner/container-layer caches where available.<br>
> The project must choose cache keys that include OS, architecture, Selenium version, browser channel/version, and lockfile.<br>
> Blindly restoring stale drivers risks browser mismatch; downloading through Selenium Manager on a cold runner is simpler but adds network time and egress dependency.

#### Artifact Integration

> Screenshots, JSON/JUnit/HTML reports, page source, BiDi logs, and Compose logs can be written to a known artifacts directory and uploaded with actions/upload-artifact under an if: always() condition.<br>
> Docker video volumes can be uploaded similarly.<br>
> Selenium does not generate or index the bundle automatically, so per-test naming, retention, compression, secret review, and report links require workflow and harness code.

#### Sharding And Matrix Support

> GitHub Actions can matrix over Chrome/Firefox/Edge and a project-defined shard index; Mocha can divide work by files through parallel workers but has no deterministic cross-job sharding or native report merger.<br>
> Selenium Grid distributes browser sessions but not test ownership.<br>
> The project must generate shard manifests, prevent duplicate retry execution, assign isolated databases, and merge JUnit/JSON/HTML output with custom or community tooling.

#### Container Compatibility

> Official Selenium standalone/Grid containers can run alongside the repository's Docker Compose application on GitHub-hosted Linux.<br>
> The easiest baseline is a host-run Node test process controlling a host browser against the published application port.<br>
> With a remote browser container, localhost refers to that container, so it must join a shared network or use a reachable host/service address; shm_size, image pinning, health checks, and Compose-project naming also require configuration.

#### Failure Cleanup

> Tests must always await driver.quit in finally or afterEach, close BiDi subscriptions, and remove temporary downloads/profiles.<br>
> GitHub Actions should place docker compose down --volumes --remove-orphans and log collection in if: always() steps.<br>
> Mocha hooks may not run after a hard cancellation or killed process, so CI-level cleanup is still necessary.<br>
> Selenium provides session deletion but no application/database cleanup coordinator.

### Cost And Risk

#### Open Source Completeness

> Browser automation, local/remote execution, Grid, Selenium Manager, Mocha, screenshots, BiDi logging/network access, Docker images, JUnit/JSON reporting, image comparison libraries, and axe-core integration are available under open-source licenses.<br>
> A paid cloud is not required, although the project must assemble several packages and custom facilities to reach an integrated E2E experience.

#### Optional Cloud Dependency

> BrowserStack, Sauce Labs, TestMu/LambdaTest, Applitools, and other hosted products can provide remote browsers, native Safari/device coverage, visual review, analytics, and dashboards.<br>
> They are optional.<br>
> Local browsers/Grid, open-source screenshot diffing, local reports, and axe-core cover the mandatory baseline, but native Safari in Linux CI requires macOS capacity or a remote service rather than Selenium alone.

#### Migration Cost

> WebDriver concepts and the W3C protocol are portable across languages, vendors, Grid providers, and many frameworks, which limits infrastructure lock-in.<br>
> Test code still binds directly to Selenium's asynchronous API, explicit waits, page objects, Mocha hooks, and a project-specific fixture/artifact harness.<br>
> Migrating to an integrated runner would require rewriting locators, waits, fixtures, and diagnostics but can preserve scenario intent and backend fixture APIs.

#### Security And Supply Chain

> selenium-webdriver 4.47.0 has four direct npm dependencies, registry signatures, and an npm provenance attestation; version 4.47 also reports a generated Selenium Manager SBOM and NOTICE.<br>
> Mocha 11.8.0 likewise has registry signatures/provenance but a larger dependency tree.<br>
> Cold setup may download browser/driver binaries from external vendor endpoints, and optional Docker images greatly expand the image footprint.<br>
> Pin lockfiles and image digests/tags, verify npm integrity/provenance, scan images and dependencies, control Selenium Manager egress/cache/telemetry settings, and update promptly for browser security releases.

#### Custom Harness Burden

> High.<br>
> The project must own browser factories, typed configuration, base URL and readiness, explicit HTMX/Navigo waits, page objects, REST fixtures, Yahoo Finance replacement, PostgreSQL reset/allocation, Compose lifecycle, retries, screenshot/log capture, report attachments, visual comparison, secret redaction, matrices/shards, and cancellation-safe cleanup.<br>
> Selenium and Mocha provide flexible primitives but deliberately do not integrate these concerns.

#### Capability Delivery Tier

- **Core:** W3C browser control, real input, navigation/history, DOM locators, explicit waits, screenshots, JavaScript execution, browser options, local/remote sessions, Grid, Selenium Manager, and expanding BiDi events/interception.
- **Runner:** Mocha test discovery, async hooks, serial execution, timeouts, retries, grep, parallel workers, and basic reporters.
- **Official Adjacent:** Official docker-selenium images, Grid observability, VNC/noVNC, and optional FFmpeg video containers.
- **Maintained Third Party:** @axe-core/webdriverjs, JUnit/HTML reporters, assertion libraries, and image-diff tools.
- **Custom Code:** Application lifecycle/readiness, API/database fixtures, HTMX synchronization helpers, route mocks, download completion/content checks, artifact correlation/redaction, visual workflow, deterministic sharding, and flake trends.
- **Optional Paid Cloud:** Hosted browser/device/Safari capacity, visual review, and analytics dashboards.

#### Ai Execution Boundary

> No AI or LLM is required to author, execute, debug, or report Selenium/Mocha tests.<br>
> Optional AI assistance can be limited to offline authoring or code review; committed deterministic tests remain the CI source of truth.<br>
> CI therefore needs no model egress, model credentials, nondeterministic runtime decisions, or per-test AI cost.

### Evidence And Decision

#### Sources

- Title: selenium-webdriver npm metadata 4.47.0 | Url: https://registry.npmjs.org/selenium-webdriver/latest | Evidence: Version, Node >=22 engine, license, four direct dependencies, npm maintainers, publish runtime, package size, signatures, provenance, and SBOM-related release context.
- Title: npm downloads API for selenium-webdriver | Url: https://api.npmjs.org/downloads/point/last-week/selenium-webdriver | Evidence: 1,820,837 downloads for 2026-08-14 through 2026-08-20.
- Title: Selenium 4.47 release | Url: https://www.selenium.dev/blog/2026/selenium-4-47-released/ | Evidence: Release date, cross-language release scope, BiDi work, Selenium 5 charter, Manager SBOM/NOTICE, and build updates.
- Title: Selenium downloads | Url: https://www.selenium.dev/downloads/ | Evidence: Current stable Selenium release.
- Title: Selenium WebDriver JavaScript API | Url: https://www.selenium.dev/selenium/docs/api/javascript/ | Evidence: Node support policy, Builder configuration, Selenium Manager integration, and remote-server support.
- Title: Selenium Manager documentation | Url: https://www.selenium.dev/documentation/selenium_manager/ | Evidence: Bundled automated driver/browser management, caching, platforms, configuration, and Linux ARM64 limitation.
- Title: Selenium waiting strategies | Url: https://www.selenium.dev/documentation/webdriver/waits/ | Evidence: Navigation readiness, implicit and explicit waits, JavaScript expected conditions, and warning against mixed waits.
- Title: Selenium common errors | Url: https://www.selenium.dev/documentation/webdriver/troubleshooting/errors/ | Evidence: Stale element behavior and the need to relocate references after dynamic DOM replacement.
- Title: Selenium WebDriver BiDi | Url: https://www.selenium.dev/documentation/webdriver/bidi/ | Evidence: Standards-based bidirectional protocol, console/network/script events, migration direction, and evolving implementation.
- Title: Selenium BiDi network | Url: https://www.selenium.dev/documentation/webdriver/bidi/w3c/network/ | Evidence: Current JavaScript network interception and authentication examples.
- Title: Selenium browser options and sessions | Url: https://www.selenium.dev/documentation/webdriver/drivers/options/ | Evidence: Browser/version capabilities, page-load strategies, and browser-specific configuration.
- Title: Selenium file downloads guidance | Url: https://www.selenium.dev/documentation/test_practices/discouraged/file_downloads/ | Evidence: No download-progress API and recommendation to use an HTTP client for content verification.
- Title: Selenium remote WebDriver | Url: https://www.selenium.dev/documentation/webdriver/drivers/remote_webdriver/ | Evidence: Remote file upload/download handling and Grid downloads capability.
- Title: Selenium project structure and governance | Url: https://www.selenium.dev/project/ | Evidence: Community governance, leadership committees, roles, sponsors, and open contribution model.
- Title: SeleniumHQ/selenium GitHub repository | Url: https://github.com/SeleniumHQ/selenium | Evidence: 34,385 stars, 8,716 forks, 186 open issue/PR count, Apache-2.0, current commits, pull requests, and issue activity observed through GitHub APIs on 2026-08-22.
- Title: Official docker-selenium repository | Url: https://github.com/SeleniumHQ/docker-selenium | Evidence: Official images, Grid topology, amd64/ARM64 browser limitations, resource sizing, VNC, and video behavior.
- Title: Official Grid images on GHCR | Url: https://www.selenium.dev/blog/2026/selenium-grid-docker-images-are-now-mirrored-to-ghcr/ | Evidence: Official GHCR mirror and pinned-image usage for CI.
- Title: Mocha npm metadata 11.8.0 | Url: https://registry.npmjs.org/mocha/latest | Evidence: Version, Node engine, MIT license, dependency and provenance metadata.
- Title: Mocha TypeScript, CLI, and parallel-mode documentation | Url: https://mochajs.org/explainers/typescript/ | Evidence: Native TypeScript stripping, ESM, async hooks, retries, timeouts, reporters, filters, parallel workers, and limitations; related pages are https://mochajs.org/running/cli/ and https://mochajs.org/features/parallel-mode/.
- Title: @types/selenium-webdriver npm package | Url: https://www.npmjs.com/package/@types/selenium-webdriver | Evidence: Separate TypeScript declarations at 4.35.6, last updated 2026-05-28.
- Title: Deque axe-core WebDriverJS integration | Url: https://www.npmjs.com/package/@axe-core/webdriverjs | Evidence: Maintained Selenium WebDriverJS accessibility integration and frame-aware audit API.
- Title: Open Asset Allocator frontend package | Url: src/main/web-static/package.json | Evidence: TypeScript 6.0.3, Parcel 2.16.4, HTMX 2.0.10, Navigo 8.11.1, Handlebars 4.7.9, Chart.js 4.5.1, and current npm module setup.
- Title: Open Asset Allocator build and Compose lifecycle | Url: src/main/docker/docker-compose-base.yml | Evidence: PostgreSQL health check, Flyway completion dependency, backend dependency ordering, and external ports; also build.sh, start.sh, destroy.sh, and Makefile.

#### Observed At

2026-08-22

#### Confidence

- **High:** Release/version, Node engine, license, weekly npm downloads, repository stars/forks, browser/protocol model, installation model, explicit-wait behavior, Docker availability, and absence of integrated runner/lifecycle/artifact facilities are supported by primary sources.
- **Moderate:** Project-specific fit, BiDi cross-browser practicality, adoption direction, issue health, typed authoring, diagnostics assembly effort, and weighted score combine current documentation with engineering analysis and need a repository spike.
- **Low:** Open Asset Allocator runtime resource use and the complete empirical spike result were not measured.

#### Hard Gate Result

- **Overall:** PASS for inclusion as the standards-based baseline, with substantial custom-harness conditions.
- **Node 24:** PASS: selenium-webdriver requires Node >=22 and Mocha supports Node 24.
- **Typescript 6:** PASS WITH RISK: JavaScript execution is compatible; typed tests need a tsc spike because declarations are separately versioned.
- **Linux And Github Ubuntu:** PASS on x64 for Chrome/Chromium, Firefox, and Edge; native Safari is unavailable.
- **Parcel Black Box Build:** PASS: testing is URL-based and requires no frontend transform or instrumentation.
- **Htmx Handlebars Dynamic Dom:** PASS WITH CUSTOM WAITS: explicit conditions and stale-element reacquisition are required.
- **Navigo Deep Links And History:** PASS WITH CUSTOM READY CONDITIONS.
- **Forms Blur Focus And Validation:** PASS through WebDriver interactions and Actions API.
- **Chartjs Canvas:** PASS WITH CUSTOM SEMANTIC OR SCREENSHOT ASSERTIONS.
- **Downloads:** PASS WITH CUSTOM COMPLETION AND CONTENT VERIFICATION; core has no progress event.
- **Api Fixtures And Yahoo Replacement:** PASS WITH NODE HTTP HELPERS, APP STUB/PROXY, OR BIDI WHERE SUPPORTED.
- **Postgresql Isolation:** PASS WITH PROJECT-OWNED RESET/ALLOCATION AND SERIALIZATION.
- **External Server And Docker Compose:** PASS WITH CUSTOM STARTUP, READINESS, NETWORKING, AND TEARDOWN.
- **Failure Diagnostics:** PASS AT MINIMUM FOR SCREENSHOTS/LOGS, BUT TRACES, CORRELATION, AND AUTOMATIC BUNDLING ARE CUSTOM.
- **Open Source Local Ci:** PASS without a paid cloud.

#### Deal Breakers

> No absolute incompatibility prevents use as a baseline.<br>
> Material reasons not to select it as the primary integrated E2E solution are the absence of locator/action/assertion auto-waiting, no lightweight browser contexts, no built-in application lifecycle or API fixture client, no integrated trace/video/HTML artifact workflow, no first-class visual regression, no Linux Safari/WebKit target, separate and lagging TypeScript declarations, and high custom-harness burden.<br>
> Native Safari additionally needs macOS or optional remote capacity.

#### Recommendation

> Viable alternative, and retain it as the established cross-language standards baseline rather than the leading recommendation for this repository.<br>
> It can cover the required user journeys on real Chrome/Edge/Firefox and native Safari infrastructure, works against the existing Parcel/Gin build without instrumentation, and is exceptionally mature and adopted.<br>
> Choose it primarily when W3C portability, existing Grid/cloud infrastructure, or cross-language consistency outweigh integrated JavaScript E2E ergonomics; otherwise the high harness and diagnostics burden is a significant disadvantage.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `maintainer_concentration`
- `github_metrics`
- `adoption_trend`
- `resource_usage`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="nightwatch-js"></a>
## 5. Nightwatch.js

Source result: `Nightwatchjs.json`

### Project And Compatibility

#### Implementation Language

> Nightwatch 3 is implemented primarily in JavaScript for Node.js and ships TypeScript declarations.<br>
> Tests can be authored in JavaScript or TypeScript and require a Node.js runtime; WebDriver browser drivers or a remote WebDriver endpoint provide browser control.

#### License And Governance

> Nightwatch and its core repository use the MIT License.<br>
> The project was created independently and has been part of BrowserStack since 2021, with development attributed to BrowserStack's open-source program.<br>
> Contributions are accepted through the public nightwatchjs organization.<br>
> BrowserStack commercial services are promoted but are not required for local WebDriver execution.

#### Installation Model

> Install the npm package with npm install --save-dev nightwatch or initialize with npm init nightwatch@latest.<br>
> Version 3 can use Selenium Manager through its Selenium dependency to resolve drivers, while explicit chromedriver, geckodriver, SafariDriver, Selenium Server, Appium, or a remote grid can also be configured.<br>
> Actual browser applications and their Linux system libraries are not bundled.<br>
> Optional capabilities require packages such as @nightwatch/vrt or @nightwatch/apitesting.<br>
> The repository should install Nightwatch in src/main/web-static or a dedicated E2E package and install Chrome/Firefox on CI.

#### Candidate Scope And Layer

> Complete Node.js E2E framework: integrated CLI runner, assertions, page objects, WebDriver session management, environments, workers, retries, reports, debugging, and optional testing plugins.<br>
> It is not merely a browser driver or a composed test-runner stack.

#### Authoring And Async Model

> Supports the built-in Nightwatch runner with exports or describe/it syntax, plus Mocha and CucumberJS integration.<br>
> Commands can be chained through Nightwatch's managed command sequence and modern APIs can be awaited.<br>
> Promise-returning callbacks are awaited.<br>
> This is less purely native-async than a browser API designed only around promises, so mixing chains, callbacks, and async/await requires consistency.

#### Build Pipeline Coupling

> Black-box E2E tests can target either the production Gin-served build or Parcel at http://localhost:8000.<br>
> Nightwatch drives the externally served page and does not require source transformation, Vite, code instrumentation, or replacement of Parcel.<br>
> Nightwatch's separate component-testing path uses Vite, but that path is not needed for this repository's E2E scope.

#### Testability Instrumentation Required

> No production build hook, injected runtime, relaxed content-security policy, or special test route is required for normal WebDriver testing.<br>
> Stable data-testid or semantic attributes would improve locator reliability but are optional application markup changes.<br>
> Deterministic PostgreSQL reset, fixture APIs, and Yahoo Finance replacement may require test-only backend or harness facilities; those requirements come from application state and external services rather than Nightwatch itself. axeInject and DOM-history collection inject or capture data only during tests.

### Maintenance Health

#### Latest Stable Release

3.16.0, published 2026-05-25 on npm and GitHub.

#### Release Cadence

> Ten stable releases were published from 2025-01-16 through 2026-05-25.<br>
> Cadence was active but irregular: several releases were weeks apart, with longer gaps from 2025-06-06 to 2025-10-28 and from 2026-01-21 to 2026-05-25.<br>
> The latest release was 89 days old on the observation date.

#### Repository Activity

> The main repository's latest commit and push were on 2026-05-25.<br>
> The latest 20 commits span 2025-12-09 to 2026-05-25 and include merged fixes for Node 24 CI, ESM, TypeScript declarations, security, and parallel WebDriver logs.<br>
> No newer core commit or merged pull request was visible by 2026-08-22, indicating materially reduced recent activity despite a current 2026 release.

#### Dependency Currency

> Mixed.<br>
> Version 3.16.0 raised the Node floor to 18.20.5, tests Node 24, fixed modern ESM handling, updated minimatch for ReDoS, and bundles current-enough axe support through nightwatch-axe-verbose.<br>
> However, the release pins selenium-webdriver 4.27.0 while npm's current version was 4.47.0 on 2026-08-22, and several direct dependencies remain on older major versions.<br>
> Browser-driver matching can be delegated to Selenium Manager, but Nightwatch's protocol API remains tied to its pinned Selenium library.

#### Wrapper Upstream Lag

> Nightwatch 3.16.0 pins selenium-webdriver 4.27.0; npm reported selenium-webdriver 4.47.0 on 2026-08-22.<br>
> This twenty-minor-version gap can delay newer WebDriver BiDi APIs and browser fixes even though Selenium Manager can obtain current driver binaries.<br>
> The official VRT plugin was still 3.5.2, published from a Node 18-era toolchain, which also shows plugin lag.

### Community Adoption

#### Npm Downloads

> The npm downloads API recorded 137,213 downloads for the nightwatch package from 2026-08-15 through 2026-08-21 and 642,513 downloads from 2026-07-22 through 2026-08-21.

#### Github Metrics

> Observed 2026-08-22 for nightwatchjs/nightwatch: 11,957 stars, 1,396 forks, 334 open issues and pull requests in the repository summary, and 264 open issues in the issue query.<br>
> GitHub-derived search metadata reported about 140 contributors. npm showed 163 dependents, while Nightwatch's marketing site claims a much larger repository usage figure that is not directly comparable.

#### Ecosystem Usage

> Evidence includes long-running usage since 2014, current adoption in Vue Router's E2E suite, integrations for Mocha, CucumberJS, Selenium Grid, Appium, BrowserStack and other WebDriver clouds, plus official or documented plugins for visual regression, API testing, accessibility, BrowserStack, reporters, and Chrome Recorder export.<br>
> Some optional official plugins have not released since 2024, so ecosystem breadth is stronger than ecosystem freshness.

#### Community Support

> The project provides a large versioned guide, API reference, GitHub issues and discussions, Discord, Stack Overflow material, examples, and public plugin repositories.<br>
> Support material covers common WebDriver and Nightwatch tasks.<br>
> Quality is uneven: some pages retain old paths or examples, and the current GitHub Actions guide still shows obsolete Node 12/14/16 and actions v3 configuration.

#### Adoption Trend

> Current usage appears broadly stable rather than clearly growing.<br>
> The comparable seven-day period fell from 146,772 downloads in 2025 to 137,213 in 2026, about 6.5 percent lower, while the comparable 31-day period rose from 634,171 to 642,513, about 1.3 percent.<br>
> Stars remain substantial, but flat download comparisons and lower recent repository activity do not support a strong growth claim.

#### Adoption Metric Normalization

> All download figures refer only to the direct npm package nightwatch and exact stated UTC date ranges; they may include CI, caching misses, upgrades, and transitive installation and do not count active projects or users.<br>
> GitHub stars and forks refer to nightwatchjs/nightwatch.<br>
> The npm dependent count is registry-specific.<br>
> Marketing claims such as over 100,000 repositories use a different and undocumented measurement, so they were not combined with npm or GitHub metrics.

### Browser And Runtime Coverage

#### Browser Engines

> Supports installed Chrome/Chromium, Firefox, Microsoft Edge, and Safari through their W3C WebDriver implementations, plus remote Selenium/Appium grids.<br>
> Chromium and Firefox work on Linux.<br>
> Safari execution is native Safari on macOS or a remote real-browser service; Nightwatch does not bundle a portable WebKit engine for Linux.<br>
> Legacy or additional browsers depend on an available conforming driver or grid.

#### Headless And Headed Modes

> The CLI supports --headless for Chrome and Firefox and normal headed execution for local debugging.<br>
> Browser-specific capabilities can configure headless Edge or remote sessions.<br>
> Headed Linux CI may use Xvfb, while modern headless browser modes avoid a virtual display for typical runs.

#### Browser Version Management

> Nightwatch does not bundle browser builds.<br>
> Browsers are installed by the operating system, CI image, Docker image, or cloud provider.<br>
> Selenium Manager can locate or download compatible driver binaries, and explicit npm driver packages or paths can be pinned.<br>
> Deterministic runs therefore require pinning the browser or container image as well as package-lock.json; Selenium Manager's cache should be preserved deliberately if reproducibility and download cost matter.

#### Parallel Browser Support

> Named test environments can run in parallel with --env firefox,chrome, and file-level test workers are controlled by test_workers or --workers.<br>
> Filter and exclude rules can split suites by environment.<br>
> There is no documented core browser-project abstraction or cross-job shard index and merge workflow comparable to purpose-built CI sharding; GitHub Actions matrices and explicit file groups must provide job-level distribution.

#### Mobile Emulation

> Nightwatch provides Chromium-only setDeviceDimensions controls for width, height, device scale factor, and the mobile rendering flag.<br>
> Browser capabilities can set user agents and other options.<br>
> Real and virtual Android/iOS mobile web execution is supported through Appium or cloud grids, but desktop emulation is not equivalent to a real device and cross-browser emulation controls are not uniform.

#### Real Browser Fidelity

> Nightwatch normally automates installed vendor browsers rather than framework-patched engines, which gives strong fidelity to Chrome, Firefox, Edge, and native Safari versions actually installed.<br>
> There is no Linux WebKit substitute.<br>
> Testing a generic WebKit build would not establish native Safari behavior, and native Safari coverage requires macOS/iOS hardware, simulators, or a cloud provider.

#### Environment Determinism Controls

> Core configuration covers window size, capabilities, profiles, browser arguments, timeouts, and fresh driver sessions.<br>
> Chromium-only helpers cover geolocation and device dimensions.<br>
> Locale, timezone, permissions, user agent, download directories, reduced motion, and animation settings can be passed through browser-specific capabilities, preferences, command-line arguments, injected CSS, or scripts.<br>
> No unified first-class controls were found for clock freezing, randomness, all permissions, animations, or cross-browser locale/timezone, so the harness must standardize these explicitly.

### Application Fit

#### Routing Support

> Navigo routes, History API transitions, back/forward navigation, URL assertions, reloads, and direct deep links are ordinary browser behavior under WebDriver.<br>
> Tests can navigate directly to Parcel or Gin URLs and assert URL and rendered state.<br>
> Route completion should be synchronized on a route-specific DOM condition rather than fixed pauses, and each test should begin from a known URL and clear relevant storage.

#### Locator Model

> Supports CSS and XPath, Selenium locator objects, page-object selectors, nested element searches, text lookup, accessibility-name retrieval, and configurable retry intervals.<br>
> It does not provide Playwright-style role locators with strict single-match enforcement as the primary model; selector results normally default to index zero unless configured.<br>
> Stable IDs, names, visible text, or test IDs are advisable, and selector-based commands should be preferred over retaining element IDs across HTMX swaps.

#### Form Interaction

> WebDriver click, setValue, sendKeys, Actions API, focus, blur through keyboard or script, select, clear, and native form submission support realistic interactions.<br>
> Nightwatch distinguishes sendKeys from setValue and can drive dynamic rows and browser validation.<br>
> Hidden-value synchronization and framework-specific blur behavior need explicit assertions and, where required, Tab or script-triggered blur rather than assuming setValue emits every application event.

#### Canvas And Download Support

> Chart.js canvas can be validated through browser.execute by reading the Chart instance, dataset, accessible fallback, dimensions, or selected state; element screenshots and @nightwatch/vrt can cover rendering, but pixel-only checks should not replace data assertions.<br>
> Nightwatch has no documented Playwright-style download event artifact.<br>
> Downloads are conventionally configured with browser download preferences, triggered through WebDriver, and verified with Node filesystem polling and cleanup, requiring custom harness code.

#### Network And Api Access

> The official @nightwatch/apitesting plugin provides SuperTest requests and an Express/Sinon mock server, and Node code in hooks can call REST fixture endpoints.<br>
> Nightwatch can capture requests and mock strict URL responses in Chromium through CDP.<br>
> That interception is not documented for Firefox or Safari.<br>
> Reliable Yahoo Finance replacement across browsers should therefore occur behind the application's HTTP boundary, DNS/proxy layer, or backend dependency injection rather than relying only on mockNetworkResponse.

#### Same Origin Support

> Nightwatch runs outside the page and can test the consolidated production server without CORS changes.<br>
> The repository's Parcel server proxies /api to localhost:8080, preserving browser-visible same-origin requests on port 8000.<br>
> Nightwatch can set launch_url or a configurable environment base URL for either topology.

#### External Server Model

> A configurable launch_url/base URL and remote WebDriver settings allow tests to target an independently started Gin/Parcel or production Docker Compose application on any reachable port.<br>
> Nightwatch does not require ownership of the application process.

#### Application Lifecycle

> Nightwatch manages WebDriver processes but has no documented built-in webServer block for arbitrary build, Docker Compose, Flyway readiness, and teardown.<br>
> Global hooks or an outer CI script must run npm build/dev, docker compose up, wait for PostgreSQL health, Flyway completion, and an HTTP readiness endpoint, then always tear services down.<br>
> This repository already has Compose health and migration dependencies, but dev.sh leaves Parcel foregrounded and is not directly a complete CI lifecycle command.

#### Visual Regression Workflow

> The MIT @nightwatch/vrt plugin captures element screenshots, creates baseline/latest/diff directories, performs Jimp pixel comparison, supports thresholds, generates a review report, and updates baselines through --update-screenshots.<br>
> Per-browser paths can be arranged in configuration.<br>
> No documented built-in masking, font readiness, animation stabilization, or semantic handling of canvas was found; those need test CSS, waits, deterministic fonts/data, and separate baselines.<br>
> Chart.js should combine data assertions with limited visual checks.

#### Accessibility Audit Integration

> Nightwatch includes axe-based axeInject and axeRun commands through nightwatch-axe-verbose, with rule selection, exclusions, and detailed violation output.<br>
> The bundled package currently uses axe-core 4.11.1.<br>
> Accessibility-name retrieval is available, but no native ARIA snapshot format comparable to an accessibility-tree snapshot was found.<br>
> Open-source JSON/JUnit/custom reporting can retain audit failures.

### Reliability

#### Waiting Model

> Element lookup and assertions poll with configurable timeout and retry interval; explicit waitForElement* and generic waitUntil cover application-specific conditions.<br>
> WebDriver enforces protocol-level element interactability, but Nightwatch does not document the same broad pre-action checks and locator auto-wait semantics as newer browser-native frameworks.<br>
> Navigation and HTMX completion should be tied to URL, DOM, or API state rather than pause or a generic network-idle assumption.

#### Flake Controls

> Provides command/assertion timeouts, element command retries, test-case --retries, suite --suiteRetries, fail-fast, explicit waits, serial mode, and file-level worker limits.<br>
> Retried tests also rerun relevant hooks.<br>
> There is no documented core repeat-each or repeat-until-failure option, so reproduction loops and random-seed control require scripts.<br>
> Retries should be reported rather than used to hide first-attempt failures.

#### Isolation Model

> The practical isolation unit is a WebDriver session associated with a suite or worker, not a cheap per-test browser context. beforeEach/afterEach hooks can reset cookies, storage, URL, and server data.<br>
> Serial scenarios can stay in one file with workers disabled or constrained, while independent files can run in parallel against isolated database data.

#### Parallelism Controls

> Workers are file-level and can be set to an exact number or disabled with --serial.<br>
> Multiple browser environments can run concurrently, and filters can partition files.<br>
> For this PostgreSQL-backed application, keep stateful suites in one serial file or set workers to one until per-worker databases or unique data namespaces exist; later CI jobs can partition explicit file groups.<br>
> Core documentation does not show a shard index or automatic cross-job report merge.

#### Flake Observability

> JUnit, JSON, HTML, timestamps, screenshots, DOM history, and per-worker WebDriver logs can expose failed attempts, but Nightwatch does not provide a documented local quarantine registry, first-attempt-versus-retry trend dashboard, or historical flake classification.<br>
> JUnit/JSON can be processed by custom scripts or CI analytics.<br>
> BrowserStack Test Observability adds hosted analysis but is optional and commercial.

### Diagnostics And Developer Experience

#### Debugging Tools

> Provides headed mode, pause, browser.debug REPL, Nightwatch Inspector, selector recommendations, --devtools, DOM History, Chrome DevTools Recorder integration, and a BrowserStack-maintained VS Code extension.<br>
> It lacks deterministic time-travel execution; DOM History is post-run HTML snapshot inspection rather than replaying browser execution.

#### Test Generation

> Chrome DevTools Recorder flows can be exported through the Nightwatch Recorder extension or converted with @nightwatch/chrome-recorder.<br>
> The Inspector recommends selectors.<br>
> Generated tests are Chrome-recording based and still require review, semantic locators, assertions, fixture setup, and removal of timing assumptions.

#### Reporters

> Built-in HTML, JUnit XML, and JSON reporters can run together, and custom reporters are supported.<br>
> Mocha mode enables its reporter ecosystem such as Mochawesome; documentation also covers Allure, TeamCity, and Slack integrations.<br>
> JUnit is suitable for GitHub test annotations through third-party actions, while HTML and JSON can be uploaded as ordinary artifacts.

#### Documentation Quality

> Broad guide and API coverage exists for setup, commands, environments, workers, retries, reporters, network helpers, visual tests, accessibility, and CI.<br>
> Current release navigation correctly identifies 3.16.0.<br>
> Quality is inconsistent: search results expose legacy expr pages, some guidance still describes explicit old drivers, and the current GitHub Actions page contains obsolete Node 12/14/16, actions v3, and an unresolved $undefined template placeholder.<br>
> Upgrade and troubleshooting confidence is therefore moderate rather than high.

#### Local Workflow

> The CLI can run one path with --test/-t, select one or more environments with --env, use --headless or headed mode, pause/debug interactively, open the HTML report, filter groups/tags, force --serial, and set retries or workers.<br>
> A developer can point launch_url at the existing Parcel dev server.<br>
> Repeated failing-test execution requires a shell/package script because no core repeat-each switch was found.

#### Failure Log Correlation

> Nightwatch can timestamp runner output, retain WebDriver logs per worker, collect Chromium console/errors/network events, and put assertion/HTTP details in reports.<br>
> It does not automatically ingest Gin, Flyway, PostgreSQL, or Docker Compose logs.<br>
> CI must timestamp and upload those service logs under a shared run ID, preserve Nightwatch worker/environment names, and collect them in an always-running failure step.

#### Artifact Data Exposure

> Screenshots, DOM History HTML, verbose WebDriver logs, network callbacks, JSON/HTML reports, VRT images, and any custom videos can expose portfolio values, API payloads, URLs, cookies, tokens, or entered credentials.<br>
> No comprehensive built-in redaction policy was found.<br>
> Disable unnecessary trace/network capture, avoid logging screenshot base64 and secrets, use synthetic accounts/data, redact custom logs, restrict artifact permissions and retention, and never upload browser profiles or storage state unless explicitly sanitized.

### Github Actions Fit

#### Official Ci Support

> Nightwatch has an official GitHub Actions guide and its own repository uses GitHub Actions, but there is no dedicated Nightwatch action or current official framework container image required for use.<br>
> Tests run on stock Ubuntu with setup-node, npm ci, installed browsers, and npx nightwatch.<br>
> The published guide is materially stale and must not be copied unchanged.

#### Browser Caching

> Cache npm through setup-node using package-lock.json.<br>
> If Selenium Manager downloads drivers, its cache can be cached only with a key tied to operating system, architecture, browser version, and driver policy; stale driver restoration is a compatibility risk.<br>
> GitHub Ubuntu images already contain browsers, but their versions change.<br>
> Pin a browser/container when deterministic baselines matter rather than relying only on caching.

#### Artifact Integration

> Use actions/upload-artifact in an always() step for tests_output, screenshots, VRT baseline/latest/diff output, DOM History HTML, WebDriver logs, and collected Compose logs.<br>
> JUnit can be consumed by a test-report action.<br>
> Nightwatch does not upload or redact these artifacts itself, so paths, retention, missing-file behavior, and secret handling are workflow responsibilities.

#### Sharding And Matrix Support

> GitHub matrices can run Nightwatch environments or explicit test groups in separate jobs, while --workers provides within-job file parallelism.<br>
> Core Nightwatch has no documented CI shard-number/total algorithm or built-in multi-job HTML report merger.<br>
> Avoid combining automatic workers with an uncoordinated file matrix, and expect custom partitioning plus JUnit aggregation.<br>
> Retries occur inside each invocation and must not be mistaken for GitHub job retries.

#### Container Compatibility

> Nightwatch can run on the GitHub host alongside this repository's Docker Compose services and access published ports 8000/8080/80 and 5432/5433 as configured.<br>
> Running the browser inside a container requires Chrome system libraries, adequate /dev/shm or --ipc=host, and possibly --no-sandbox under the documented container model.<br>
> A host-runner browser is simpler while Compose owns Gin, PostgreSQL, and Flyway.<br>
> Service readiness remains custom.

#### Failure Cleanup

> Nightwatch ends managed WebDriver sessions under normal completion, but application and database cleanup are external.<br>
> The workflow should use if: always() for artifact collection followed by docker compose down --volumes --remove-orphans, terminate Parcel/background processes, and use job timeouts.<br>
> Cancellation can interrupt ordinary hooks, so disposable Compose volumes and GitHub runner teardown are safer than relying only on after hooks.

### Cost And Risk

#### Open Source Completeness

> The core runner, local Chrome/Firefox execution, workers, retries, screenshots, HTML/JUnit/JSON reports, DOM History, axe integration, API plugin, and VRT plugin are available under open-source licenses.<br>
> Native Safari still requires Apple infrastructure, as it does for WebDriver generally.<br>
> No paid service is mandatory for the repository's Linux Chrome/Firefox E2E baseline.

#### Optional Cloud Dependency

> BrowserStack, Sauce Labs, and other grids are optional for hosted real-browser/device coverage.<br>
> BrowserStack Test Observability, Percy, AI test management, and broad device labs can add paid capabilities, but local reports, JUnit, @nightwatch/vrt, and axe provide non-cloud alternatives.<br>
> BrowserStack's natural-language AI execution is a separate beta cloud capability and is not required by Nightwatch.

#### Migration Cost

> Moderate.<br>
> Standard JavaScript/TypeScript, WebDriver concepts, CSS/XPath, and direct Selenium driver access are portable, but Nightwatch command chaining, browser API names, page objects, custom commands/assertions, configuration environments, hooks, and reporter/plugin APIs are framework-specific.<br>
> Tests heavily using Chromium-only Nightwatch CDP helpers or BrowserStack capabilities increase lock-in.

#### Custom Harness Burden

> Moderate to high for this application.<br>
> Nightwatch supplies browser lifecycle, waits, reports, retries, and workers, but the repository must add Compose startup/teardown, readiness checks, Flyway verification, PostgreSQL reset or namespacing, API fixture helpers, cross-browser Yahoo Finance replacement, download polling, backend log capture, and artifact upload/redaction.<br>
> This burden is greater than frameworks with a built-in web-server lifecycle and browser-context fixtures.

#### Capability Delivery Tier

> Core: WebDriver E2E, runner, assertions, hooks, page objects, Chrome/Firefox/Edge/Safari environments, headless mode, workers, retries, screenshots, HTML/JUnit/JSON, DOM History, Inspector, and WebDriver logs.<br>
> Bundled/community dependency: axe commands through nightwatch-axe-verbose.<br>
> Official optional plugin: @nightwatch/vrt and @nightwatch/apitesting.<br>
> Browser-specific core helper: Chromium network capture/mock, console/error capture, geolocation, and device dimensions.<br>
> Custom repository code: Compose/Flyway/PostgreSQL lifecycle, cross-browser service replacement, downloads, deterministic clock/animations, log correlation, and CI artifact policy.<br>
> Paid optional service: hosted devices, Test Observability, Percy, and BrowserStack AI.

#### Ai Execution Boundary

> Normal Nightwatch tests are deterministic code and need no LLM, AI token, model egress, or AI credential in CI.<br>
> Chrome Recorder and Inspector are non-LLM authoring aids.<br>
> BrowserStack's AI test-management and natural-language execution features are separate optional cloud products; the natural-language agent is beta and can interpret steps at runtime.<br>
> Keep aiAuthoring disabled, do not install cloud AI credentials, and retain ordinary Nightwatch code as the CI fallback and source of truth.

### Evidence And Decision

#### Sources

- Nightwatch 3.16 official guide and API: https://nightwatchjs.org/guide/overview/what-is-nightwatch.html
- Nightwatch release history: https://github.com/nightwatchjs/nightwatch/releases
- Nightwatch 3.16.0 package metadata at tag v3.16.0: https://github.com/nightwatchjs/nightwatch/blob/v3.16.0/package.json
- Nightwatch npm package: https://www.npmjs.com/package/nightwatch
- npm downloads API, exact 2026 week: https://api.npmjs.org/downloads/point/2026-08-15:2026-08-21/nightwatch
- npm downloads API, exact 2026 31-day period: https://api.npmjs.org/downloads/point/2026-07-22:2026-08-21/nightwatch
- npm downloads API comparison periods: https://api.npmjs.org/downloads/point/2025-08-16:2025-08-22/nightwatch and https://api.npmjs.org/downloads/point/2025-07-22:2025-08-21/nightwatch
- Core repository metrics, commits, issues, and pull requests: https://github.com/nightwatchjs/nightwatch
- Node 24 CI change: https://github.com/nightwatchjs/nightwatch/pull/4428
- Nightwatch parallel execution guide: https://nightwatchjs.org/guide/running-tests/parallel-running.html
- Nightwatch GitHub Actions guide: https://nightwatchjs.org/guide/ci-integrations/run-nightwatch-on-github-actions.html
- Nightwatch network capture and mocking guides: https://nightwatchjs.org/guide/network-requests/capture-network-calls.html and https://nightwatchjs.org/guide/network-requests/mock-network-response.html
- Nightwatch visual regression guide: https://nightwatchjs.org/guide/writing-tests/visual-regression-testing.html
- Nightwatch accessibility guide: https://nightwatchjs.org/guide/using-nightwatch/accessibility-testing.html
- Nightwatch DOM History guide: https://nightwatchjs.org/guide/reporters/dom-history.html
- Current selenium-webdriver registry metadata: https://registry.npmjs.org/selenium-webdriver/latest
- Open Asset Allocator package and TypeScript configuration: src/main/web-static/package.json and src/main/web-static/tsconfig.json
- Open Asset Allocator Parcel proxy and lifecycle: src/main/web-static/.proxyrc.js, dev.sh, and src/main/docker/docker-compose-base.yml

#### Observed At

2026-08-22

#### Confidence

> High confidence in release, package, license, Node 24, repository, npm download, protocol baseline, browser, runner, reporter, and documented feature facts.<br>
> Medium confidence in project-fit inferences, maintenance trajectory, plugin freshness implications, isolation details, and comparative developer experience.<br>
> Low confidence in TypeScript 6 compatibility, architecture coverage, exact resource use, security status, and behavior against this application because no project spike was run.

#### Deal Breakers

> No unconditional deal-breaker exists for a JavaScript-based Chrome/Firefox E2E suite.<br>
> Exclude Nightwatch if mandatory requirements include native Safari or WebKit on Linux, first-class cross-browser request interception, isolated browser contexts per test, built-in application-server/Compose lifecycle, current Selenium APIs without wrapper lag, or verified TypeScript 6 support without a compatibility spike.<br>
> The stale official CI example must not be adopted directly.

#### Recommendation

> Viable alternative, not recommended as the primary choice.<br>
> Nightwatch is an established MIT complete framework with current Node 24 coverage, substantial stable npm usage, real-browser WebDriver reach, workers, retries, reports, DOM History, visual and accessibility options, and no mandatory cloud.<br>
> It ranks below a primary recommendation for this repository because of low recent core activity, a large Selenium binding lag, unverified TypeScript 6 compatibility, Chromium-only network controls, no isolated contexts or built-in application lifecycle, stale CI guidance, and significant custom harness work for PostgreSQL, Flyway, Yahoo Finance, downloads, and diagnostics.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `operating_system_support`
- `issue_health`
- `roadmap_and_deprecation_risk`
- `maintainer_concentration`
- `browser_protocol`
- `dynamic_dom_synchronization`
- `test_isolation`
- `resource_usage`
- `failure_artifacts`
- `security_and_supply_chain`
- `hard_gate_result`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="testcafe"></a>
## 6. TestCafe

Source result: `TestCafe.json`

### Project And Compatibility

#### Implementation Language

> TestCafe is implemented primarily in JavaScript with a substantial TypeScript portion and bundled TypeScript declarations.<br>
> Tests can be authored in JavaScript or TypeScript and run on Node.js.<br>
> Browser-side operations use TestCafe's Selector and TestController APIs.<br>
> [S1][S3]

#### License And Governance

> TestCafe is MIT-licensed and owned and commercially stewarded by Developer Express Inc.<br>
> Development and issue tracking are public in DevExpress/testcafe.<br>
> DevExpress also sells TestCafe Studio, but the open-source runner does not require a commercial license or account.<br>
> The governance model is company-led rather than foundation-governed.<br>
> [S1][S3][S20]

#### Installation Model

> Install locally with npm install --save-dev testcafe and run with npx testcafe; the official documentation also shows a global install.<br>
> TestCafe does not bundle browser binaries, so Chrome/Chromium, Firefox, Edge, or Safari must be installed on the host.<br>
> Video recording additionally requires FFmpeg.<br>
> The official testcafe/testcafe:3.7.6 image is about 500 MB compressed, amd64-only, Alpine-based, and includes Chromium, Firefox, Node, npm, Xvfb, and fonts.<br>
> Pin the npm lockfile and an image digest or host browser policy for deterministic CI.<br>
> [S1][S9][S11][S23]

#### Candidate Scope And Layer

> Complete Node.js E2E framework with a self-contained CLI and programmatic runner, fixture/test DSL, selectors, assertions, waits, request hooks, API requests, concurrency, quarantine retries, screenshots, video, reporters, browser launching, and debugging tools.<br>
> It is not merely a browser driver or a composed runner stack.<br>
> [S7][S8][S9]

#### Authoring And Async Model

> Tests use native JavaScript or TypeScript async functions and await a framework-specific TestController chain such as await t.click(...).expect(...).<br>
> Selectors are lazy query objects and assertions can re-evaluate them.<br>
> Fixtures, tests, hooks, Roles, RequestHooks, and ClientFunctions form a TestCafe-specific DSL, but browser commands are promise-compatible rather than a hidden Cypress-style command queue.<br>
> [S7][S13]

#### Build Pipeline Coupling

> Black-box tests can target either the built Parcel application served by Gin or the Parcel development server through baseUrl.<br>
> TestCafe does not require Vite, source instrumentation, or replacement of Parcel.<br>
> Chromium native automation uses CDP without rewriting the application.<br>
> Firefox, Safari, remote, and cloud execution disables native automation and uses the Hammerhead reverse proxy, which injects and rewrites client content at runtime but still does not alter the source build pipeline.<br>
> [S6][S19]

#### Testability Instrumentation Required

> No production build hook, special route, or test runtime is required for ordinary Chromium E2E tests.<br>
> Stable IDs or data attributes are optional but advisable because TestCafe lacks first-class role/label locators.<br>
> Proxy-mode browsers inject TestCafe scripts and may expose compatibility differences with strict security policies or modern client code.<br>
> Deterministic PostgreSQL reset, fixture seeding, and server-side Yahoo Finance replacement require test harness or backend seams because those concerns are outside the browser runner.<br>
> [S6][S8][S19]

### Maintenance Health

#### Latest Stable Release

> 3.7.6, published on 2026-07-07 to npm and GitHub.<br>
> The release only updated the package lock/install command and version metadata.<br>
> [S1][S2]

#### Release Cadence

> Slow and irregular.<br>
> Seven stable releases were published from 2024-11-04 through 2026-07-07.<br>
> Gaps included about ten months between 3.7.2 and 3.7.3 and about five months between 3.7.4 and 3.7.5.<br>
> Releases 3.7.3 through 3.7.6 were largely fixes, dependency/security maintenance, lockfile work, and release metadata; the last documented feature release was 3.7.0 in November 2024.<br>
> [S2][S12]

#### Repository Activity

> Low but not dormant.<br>
> The latest default-branch commit observed was the 3.7.6 release on 2026-07-07; a search from 2026-07-01 returned only that commit.<br>
> The repository was updated on 2026-08-20 by open Dependabot activity, but no later merged core change was found.<br>
> Recent merged work was concentrated on dependency vulnerabilities, installation, CI maintenance, and releases rather than new E2E capabilities.<br>
> [S3][S21]

### Community Adoption

#### Npm Downloads

> The npm downloads API recorded 176,264 downloads for the exact seven-day window 2026-08-14 through 2026-08-20 and 904,319 downloads for 2026-07-22 through 2026-08-21 for the testcafe package.<br>
> [S4]

#### Ecosystem Usage

> TestCafe has a long-lived ecosystem with BrowserStack, Sauce Labs, LambdaTest, remote-browser providers, reporter plugins, IDE plugins, Cucumber adapters, an official Docker image, and commercial TestCafe Studio compatibility.<br>
> Current ecosystem freshness is weak: the official GitHub Action is archived and explicitly unmaintained, the README-listed axe-testcafe helper has not been pushed since 2020, and sampled visual-regression integrations are small, stale, or archived.<br>
> This is stronger evidence for an installed legacy base than for current greenfield adoption.<br>
> [S10][S11][S16][S17][S20]

#### Community Support

> Official documentation is broad and Stack Overflow has about 874 cumulative questions under the testcafe tag.<br>
> GitHub issues remain available, but the repository has no Discussions forum.<br>
> Recent Stack Overflow activity is sparse, and important official pages contain obsolete Node versions, action versions, links, and product status.<br>
> Existing troubleshooting material is substantial, but its age increases verification effort for current environments.<br>
> [S3][S10][S18]

#### Adoption Trend

> Declining on directly comparable npm evidence.<br>
> Weekly downloads fell from 227,383 in 2025-08-14 through 2025-08-20 to 176,264 in the matching 2026 window, about 22.5 percent.<br>
> The matching 31-day comparison fell from 1,085,504 to 904,319, about 16.7 percent.<br>
> Slow feature delivery, archived official integration tooling, and stale plugins agree with a contraction trajectory, although downloads remain large enough to show continued use.<br>
> [S4][S10][S16]

#### Adoption Metric Normalization

> All registry counts refer only to the direct npm package testcafe, all versions, over the exact stated UTC windows.<br>
> They can include repeat CI installations, bots, mirrors, and cache misses and are not unique projects or users.<br>
> The package represents the complete E2E runner, so its scope is more directly comparable than a low-level driver package.<br>
> GitHub stars/forks are cumulative interest signals, and npm/GitHub dependent counts can include transitive use; none should be treated as active production-suite counts.<br>
> [S3][S4][S15]

### Browser And Runtime Coverage

#### Browser Engines

> TestCafe supports installed Chromium, Google Chrome, Chrome Canary, Chromium-based Edge, Opera, Firefox, and Safari.<br>
> Chrome and Firefox support headless execution on Linux.<br>
> Native Safari requires macOS or a remote/cloud environment.<br>
> TestCafe does not ship a Linux WebKit browser and does not automate legacy Edge or Internet Explorer.<br>
> Remote/mobile/cloud support requires proxy mode.<br>
> [S6]

#### Browser Protocol

> Chromium-based browsers use Chrome DevTools Protocol native automation by default.<br>
> Firefox, Safari, remote, mobile, and cloud browsers disable native automation and use TestCafe's Hammerhead reverse proxy and injected client scripts to emulate and coordinate interactions; Firefox launch control can also use a Marionette port.<br>
> TestCafe does not provide a W3C WebDriver or WebDriver BiDi execution path.<br>
> This split creates browser-specific behavior and feature limitations.<br>
> [S6]

#### Headless And Headed Modes

> Chrome/Chromium and Firefox can run headless with the :headless browser option and headed for local debugging.<br>
> Screenshots and window resizing are supported in headless mode, subject to documented Chromium window-size quirks.<br>
> Other installed browsers run headed unless their own provider supports a headless mode.<br>
> Xvfb is included in the official Linux image.<br>
> [S6][S11]

#### Browser Version Management

> TestCafe detects and launches system-installed browsers or an explicitly supplied executable path.<br>
> It does not download or pin browser versions.<br>
> Deterministic CI therefore requires a pinned host image, browser package policy, or Docker digest.<br>
> The official 3.7.6 container bundles Alpine repository versions of Chromium and Firefox at image-build time, but tag reuse and unpinned system packages should be avoided by pinning the observed image digest.<br>
> [S6][S11][S23]

#### Parallel Browser Support

> A single run can target multiple browser aliases.<br>
> The concurrency factor launches N instances of every selected browser and distributes tests between them; a fixture can disable concurrency.<br>
> There is no projects abstraction, worker fixture system, native shard index/total, or cross-job result merger.<br>
> Browser matrices and CI shards require separate commands and explicit test partitioning.<br>
> [S7]

#### Mobile Emulation

> Chromium emulation can set a named device or width, height, device pixel ratio, orientation, touch, and user agent through browser options.<br>
> TestCafe can connect remote mobile browsers through a URL/QR workflow or cloud provider, but native automation must be disabled.<br>
> Desktop emulation is not a physical device, and cross-browser device controls are not unified.<br>
> [S6]

#### Real Browser Fidelity

> TestCafe uses installed vendor browsers, so Chromium CDP runs exercise real Chrome/Edge/Chromium binaries.<br>
> Firefox and native Safari can also be actual installed browsers, but TestCafe's proxy rewrites page content and emulates events in those engines, reducing fidelity and introducing framework-specific behavior.<br>
> There is no generic WebKit substitute on Ubuntu, and neither Chromium emulation nor a WebKit build would establish native Safari behavior.<br>
> [S6]

#### Environment Determinism Controls

> Core controls include an empty browser profile, window size, Chromium device emulation, request mocks, cookies, and geolocation responses through native-dialog handling.<br>
> Browser flags and ClientFunctions can adjust some settings.<br>
> There is no unified first-class API for clock freezing, timezone, locale, all permissions, seeded randomness, animation disabling, reduced motion, or cross-browser device profiles.<br>
> These require browser flags, injected scripts/CSS, application fixtures, or custom helpers and may differ between CDP and proxy modes.<br>
> [S6][S8][S13]

### Application Fit

#### Routing Support

> Navigo navigation can be exercised through real links, t.navigateTo, URL ClientFunctions, and direct fixture/base URLs.<br>
> Each test can begin at a known deep link, while browser history can be invoked and asserted through ClientFunction code.<br>
> Gin or Parcel must serve the SPA shell for direct routes.<br>
> TestCafe has no dedicated waitForURL or back/forward assertion API, so route completion should be synchronized on location and route-specific DOM assertions rather than sleeps.<br>
> [S13][S19]

#### Locator Model

> Selectors support CSS, functions, text filtering, attributes, descendants, siblings, index selection, and custom DOM properties/methods.<br>
> They are lazy and normally reacquire nodes after HTMX replacement.<br>
> TestCafe has no core role, accessible-name, label, placeholder, or strict-locator API; a selector that matches multiple elements commonly resolves by index rather than failing for ambiguity.<br>
> Prefer semantic IDs, visible text, explicit label relationships, or optional data-testid attributes, and avoid storing Selector snapshots across swaps.<br>
> [S13]

#### Form Interaction

> TestController actions cover click, typeText, selectText, pressKey, focus through interaction, file upload, drag, hover, and native-dialog handling.<br>
> Blur can be produced with Tab or a click on another element.<br>
> Chromium native automation fires browser-native protocol events; proxy browsers emulate events.<br>
> This can test input masks, hidden-value synchronization, dynamic rows, and native validity, but tests must explicitly assert blur-driven state and use ClientFunction for properties not exposed through Selector snapshots.<br>
> [S6][S13][S19]

#### Canvas And Download Support

> Chart.js can be checked semantically through ClientFunction or t.eval by reading chart data, canvas dimensions, or application state, with targeted element screenshots as secondary evidence.<br>
> Core TestCafe captures screenshots but has no visual comparator.<br>
> It also has no first-class download event, stream, suggested-filename, or saveAs API.<br>
> RequestLogger can capture a download response body, while true browser downloads require configured browser preferences and Node filesystem polling/cleanup.<br>
> [S8][S9][S19]

#### Network And Api Access

> t.request can seed and query REST fixtures.<br>
> RequestLogger observes browser and t.request traffic, and RequestMock or custom RequestHook can match, log, mock, and modify HTTP requests/responses.<br>
> WebSockets cannot be intercepted, and native CDP mode limits request mutation fields.<br>
> Browser hooks cannot replace Yahoo Finance traffic initiated exclusively by the Go backend; that requires backend dependency injection, a fake upstream endpoint, proxy/DNS control, or seeded data.<br>
> [S8][S19]

#### Same Origin Support

> TestCafe can target the consolidated Gin production origin or Parcel on port 8000, where the repository proxy forwards /api to Gin on port 8080. baseUrl supports relative test pages and Roles.<br>
> No application CORS change is needed for these topologies, although Firefox/Safari proxy mode introduces TestCafe's own reverse-proxy origin mechanics internally.<br>
> [S8][S19]

#### External Server Model

> Supported.<br>
> Tests can use fixture.page, test.page, baseUrl, or environment-selected URLs to target independently started Parcel, Gin, Docker Compose, or remote deployments on arbitrary reachable ports.<br>
> TestCafe does not require ownership of the application process.<br>
> [S13][S19]

#### Application Lifecycle

> TestCafe can execute one application command before the run and stop its child process afterward, using a fixed initialization delay.<br>
> It does not provide declarative multi-service orchestration, URL readiness polling, Flyway state checks, or Docker Compose health integration.<br>
> This repository therefore needs an outer script or run hook to build, start Compose, wait for PostgreSQL, Flyway, Gin, and Parcel readiness, then collect logs and tear down services reliably.<br>
> [S13][S19]

#### Visual Regression Workflow

> Core TestCafe provides page/element screenshots, full-page capture, failure capture, and configurable paths, but no baseline creation, pixel comparison, masking, threshold policy, review UI, or report merger.<br>
> Search results for TestCafe visual-regression projects were predominantly small, stale, or archived.<br>
> A project-owned image comparator could be built, but it must pin browser/image/fonts, disable animations, wait for Chart.js, maintain per-browser baselines, and upload diffs.<br>
> Semantic chart assertions should remain primary.<br>
> [S9][S17]

#### Accessibility Audit Integration

> TestCafe can inject a current axe-core bundle through client scripts and execute it with ClientFunction/custom helpers.<br>
> The README-listed axe-testcafe community helper was last pushed in 2020 and should not be adopted without a security and compatibility review.<br>
> TestCafe lacks core role locators, accessibility-tree assertions, and ARIA snapshots.<br>
> Open-source axe JSON can be attached to custom JSON/xUnit reporting, but this is custom integration work.<br>
> [S16][S20]

### Reliability

#### Waiting Model

> Selectors wait for DOM presence; actions wait for a visible target; smart assertions retry Selector properties and ClientFunctions; redirects receive a fixed wait; and TestCafe waits for XHR/fetch up to the AJAX timeout before continuing.<br>
> Important limitations are that smart assertions do not themselves wait for element existence, native automation ignores page/AJAX timeout settings, and an overlapped action target may cause TestCafe to interact with the topmost element after timeout instead of failing.<br>
> HTMX tests should assert explicit post-action DOM/API state.<br>
> [S6][S7]

#### Flake Controls

> Selector/assertion timeouts, explicit state assertions, request logging/mocking, fixture-level concurrency disabling, and Quarantine Mode are available.<br>
> Quarantine reruns a failed test until it reaches a configurable success threshold or attempt limit; defaults are three successes within five attempts and it classifies alternating results as unstable.<br>
> This is stronger than a blind retry but can multiply runtime and database side effects.<br>
> No core repeat-each stress option, random seed control, or automatic per-test rollback exists.<br>
> [S7]

#### Isolation Model

> The principal process isolation unit is a browser instance created by the run/concurrency setting, not a fresh browser context per test.<br>
> Fixtures and hooks can reset browser and server state, and fixtures can disable concurrency for serial scenarios.<br>
> Shared PostgreSQL state requires repository-owned cleanup or per-worker namespacing.<br>
> Quarantine repeats execute the scenario multiple times, so fixture setup and cleanup must be idempotent.<br>
> [S7][S19]

#### Parallelism Controls

> Concurrency can be set globally to an exact browser-instance count and disabled for selected fixtures.<br>
> Multiple target browsers multiply that count.<br>
> For this shared PostgreSQL application, start at concurrency 1 or mark stateful fixtures disableConcurrency; increase only after assigning independent databases/schemas and unique test data.<br>
> There is no native CI sharding algorithm, worker index fixture, or merged report pipeline, so future distribution requires explicit file/metadata partitions.<br>
> [S7][S19]

#### Flake Observability

> Quarantine Mode labels stable versus unstable outcomes and artifact path patterns include the quarantine attempt.<br>
> JSON/xUnit/custom reporters can retain per-attempt data, screenshots, and video.<br>
> TestCafe has no integrated open-source historical flake database, quarantine registry, first-attempt trend dashboard, or replay trace.<br>
> Long-term classification requires repository-owned report ingestion and artifact retention.<br>
> [S7][S9][S14]

### Diagnostics And Developer Experience

#### Failure Artifacts

> TestCafe can capture automatic failure screenshots, manual page/element screenshots, per-test or single-file MP4 video, test/report output, browser JavaScript errors, browser console messages through APIs, and selected HTTP data through RequestLogger.<br>
> It does not produce a replayable execution trace, DOM timeline, HAR, or automatically correlated complete network/console archive.<br>
> Remote browsers cannot use core screenshot/video capture, and local video requires FFmpeg.<br>
> [S8][S9][S14]

#### Debugging Tools

> Headed runs, t.debug breakpoints, browser developer tools, a Visual Selector Debugger with element picking, debug-on-fail, and Live Mode for automatic reruns are available.<br>
> The tools can inspect current DOM and selector matches but do not provide time travel, trace replay, or a maintained first-party VS Code testing experience comparable to newer frameworks.<br>
> [S13][S14]

#### Test Generation

> Open-source TestCafe can experimentally replay Chrome Recorder user-flow JSON and the Visual Selector Debugger can generate selector expressions.<br>
> It does not include a maintained core code recorder that generates complete tests.<br>
> Visual recording, automatic selector generation, and an interactive editor are TestCafe Studio commercial features.<br>
> Generated selectors or recordings still require assertions, deterministic fixtures, and maintenance review.<br>
> [S12][S14][S20]

#### Reporters

> The package includes spec, list, minimal, JSON, and xUnit reporters, supports multiple reporters, output files, custom reporter plugins, and output hooks. xUnit can feed GitHub-compatible test-report actions; JSON can support custom analytics.<br>
> HTML, Allure, and GitHub annotation experiences require community reporters or custom processing, and there is no core shard-report merger.<br>
> [S1][S14]

#### Documentation Quality

> Broad but materially stale.<br>
> Official guides cover selectors, waits, requests, browsers, concurrency, quarantine, artifacts, CI, Docker, and extension APIs.<br>
> However, the install README still says Node 16 while 3.7.6 requires Node >=20, TypeScript guidance remains on version 4 while the repository uses TypeScript 6, release notes stop at 3.7.0, the roadmap is mostly 2021-2022 work, and the GitHub Actions guide recommends an archived action plus setup-node/checkout v1 and Node 8/10/12.<br>
> Every current setup claim requires source/package verification.<br>
> [S1][S5][S10][S12]

#### Local Workflow

> Developers can run one file/glob, browser, metadata filter, concurrency level, or quarantine configuration from the CLI; headed mode, t.debug, debug-on-fail, and Live Mode support focused iteration. create-testcafe scaffolds a project.<br>
> There is no open-source TestCafe UI runner, watch dashboard, trace viewer, or built-in last-failed/repeat-each workflow, so repeated failure reproduction needs filtering and a package/script loop.<br>
> [S7][S13][S14]

#### Failure Log Correlation

> TestCafe reporters can timestamp test results, RequestLogger can retain selected browser/API requests, and screenshots/videos can encode test/browser/attempt names.<br>
> It does not ingest Gin, PostgreSQL, Flyway, or Docker Compose logs or build a shared event timeline.<br>
> CI must assign a run identifier, use UTC timestamps, save browser console/request evidence, collect compose ps/logs and Flyway output, and upload all files in an always-running artifact step.<br>
> [S8][S9][S19]

#### Artifact Data Exposure

> Screenshots, MP4 video, request/response bodies and headers, browser console messages, JSON/xUnit reports, and downloaded files can expose portfolio data, credentials, cookies, tokens, and external API data.<br>
> TestCafe offers selective RequestLogger body/header capture and artifact path controls but no universal redaction layer.<br>
> Use synthetic data, avoid logging secrets, sanitize custom reporters/hooks, restrict GitHub artifact access and retention, and do not persist browser profiles or raw sensitive payloads.<br>
> [S8][S9]

### Github Actions Fit

#### Official Ci Support

> Weak.<br>
> TestCafe has an official GitHub Actions guide, but it directs users to DevExpress/testcafe-action, whose repository is archived and explicitly says the team no longer maintains it.<br>
> The guide itself uses obsolete checkout/setup-node v1 and Node 8/10/12 examples.<br>
> A viable workflow should instead use current actions/checkout and actions/setup-node, npm ci, installed Ubuntu browsers, and npx testcafe.<br>
> The official Docker image remains current at 3.7.6.<br>
> [S10][S11][S23]

#### Browser Caching

> TestCafe does not download managed browser binaries, so cache npm through setup-node and rely on GitHub runner browsers or a pinned Docker image.<br>
> Host browser versions change with runner images and are not controlled by npm caching.<br>
> Pulling testcafe/testcafe can use the runner's Docker layer cache, but image reproducibility requires a version and digest; the observed image is about 500 MB.<br>
> [S6][S11][S23]

#### Artifact Integration

> Use actions/upload-artifact under always() for screenshots, MP4 videos, JSON/xUnit output, custom request/console logs, downloads when safe, and Compose/Flyway/backend logs.<br>
> TestCafe writes artifacts to configurable paths but does not upload, redact, annotate, or merge them. xUnit annotations and HTML conversion need separate maintained actions/tools.<br>
> [S9][S14]

#### Sharding And Matrix Support

> GitHub matrices can run separate browsers or manually assigned file/metadata groups, while TestCafe concurrency parallelizes within a job.<br>
> There is no native shard i/n switch, deterministic balancing, worker-index fixture, blob format, or cross-job HTML merge.<br>
> Database-mutating matrices require independent PostgreSQL instances or must stay serial.<br>
> Quarantine attempts happen inside one invocation and should not be confused with GitHub job retries.<br>
> [S7][S10]

#### Container Compatibility

> The simplest topology is TestCafe on the Ubuntu host with this repository's Gin/PostgreSQL/Flyway services in Docker Compose and published localhost ports.<br>
> The official amd64 TestCafe image includes Chromium/Firefox and Xvfb, but a containerized browser must reach the application by service name, host gateway, or host networking; localhost points to the browser container.<br>
> Official guidance uses --net=host for some cases and warns that it weakens isolation.<br>
> [S11][S19][S23]

#### Failure Cleanup

> TestCafe normally closes launched browsers and can terminate an appCommand child process, but Docker Compose, PostgreSQL volumes, Flyway state, and independent Parcel/Gin processes remain external.<br>
> The workflow should collect artifacts and service logs, then run docker compose down --volumes --remove-orphans under always(), with job timeouts and unique Compose project names.<br>
> Cancellation may bypass TestCafe hooks, so disposable resources and idempotent pre-run cleanup are required.<br>
> [S10][S19]

### Cost And Risk

#### Open Source Completeness

> The MIT core supplies the runner, installed-browser automation, selectors, assertions, waits, HTTP/API controls, concurrency, quarantine, screenshots, video, JSON/xUnit reporting, and debugging needed for a basic local and CI suite.<br>
> No paid service is mandatory.<br>
> Modern visual regression, maintained accessibility integration, trace replay, CI sharding/merge, and application/database lifecycle are not complete first-party open-source capabilities and require custom or third-party tooling.<br>
> [S8][S9][S14][S16][S17]

#### Optional Cloud Dependency

> BrowserStack, Sauce Labs, LambdaTest, and other grids are optional for hosted browsers and physical devices.<br>
> TestCafe Studio is an optional paid IDE/recorder; open-source TestCafe can run compatible Studio tests in CI.<br>
> Core execution, screenshots, video, quarantine, requests, and reporters need no cloud account.<br>
> Local alternatives exist for most requirements, but not as an integrated trace/dashboard product.<br>
> [S6][S20]

#### Migration Cost

> Moderate.<br>
> JavaScript/TypeScript, async/await, CSS selectors, page objects, REST fixtures, and test data are portable concepts.<br>
> TestCafe fixture/test globals, TestController chains, Selector filters/snapshots, Roles, ClientFunctions, RequestHooks, quarantine semantics, reporter APIs, and configuration are framework-specific and require rewrites.<br>
> Keeping domain fixtures and database helpers independent from TestCafe lowers future migration cost.<br>
> [S7][S8][S13]

#### Custom Harness Burden

> High for this application.<br>
> TestCafe supplies browser lifecycle, selectors, waits, HTTP hooks, API calls, concurrency, retries, and artifacts, but the repository must build and start Compose, poll PostgreSQL/Flyway/Gin/Parcel readiness, reset or namespace database data, replace server-side Yahoo Finance, handle downloads, integrate current axe-core and visual comparison, collect backend logs, split/merge CI work, redact artifacts, and guarantee teardown.<br>
> Stale official CI guidance adds maintenance burden.<br>
> [S8][S10][S19]

#### Capability Delivery Tier

> Core OSS: runner, JavaScript/TypeScript DSL, installed Chrome/Edge/Firefox/Safari support, Chromium CDP, proxy fallback, selectors, waits, API requests, request hooks/mocks, concurrency, quarantine, screenshots, video, JSON/xUnit, debug tools, and Live Mode.<br>
> Official OSS companion: amd64 Docker image.<br>
> Deprecated official companion: archived GitHub Action.<br>
> Stale/community: axe-testcafe, visual comparison, richer HTML/Allure reporters, IDE and Cucumber plugins.<br>
> Custom repository code: Compose/Flyway/PostgreSQL lifecycle, backend Yahoo replacement, downloads, deterministic environment, visual/accessibility maintenance, sharding/merge, log correlation, and artifact policy.<br>
> Paid optional: TestCafe Studio and external browser clouds.<br>
> [S10][S11][S14][S16][S17][S20]

#### Ai Execution Boundary

> Normal TestCafe tests, assertions, selectors, request hooks, quarantine runs, and CI execution require no LLM, model credential, AI egress, or per-run AI cost.<br>
> Visual Selector Debugger and Chrome user-flow replay are non-LLM aids.<br>
> Keep any third-party AI authoring outside required CI, provide synthetic data and least-privilege access, and commit reviewed JavaScript/TypeScript tests that run fully without AI.<br>
> [S12][S14]

### Evidence And Decision

#### Sources

- [S1] TestCafe 3.7.6 package metadata and tagged package.json, https://registry.npmjs.org/testcafe/latest and https://github.com/DevExpress/testcafe/blob/v3.7.6/package.json, accessed 2026-08-22.
- [S2] TestCafe GitHub releases v3.7.0 through v3.7.6, https://github.com/DevExpress/testcafe/releases, observed 2026-08-22.
- [S3] DevExpress/testcafe repository metadata, commits, issues, pull requests, and contributor search results, https://github.com/DevExpress/testcafe, observed through GitHub APIs on 2026-08-22.
- [S4] npm Downloads API for testcafe exact windows 2026-08-14:2026-08-20, 2025-08-14:2025-08-20, 2026-07-22:2026-08-21, and 2025-07-22:2025-08-21, retrieved 2026-08-22.
- [S5] Official TestCafe TypeScript and CoffeeScript guide, https://testcafe.io/documentation/402824/guides/intermediate-guides/typescript-and-coffeescript, accessed 2026-08-22.
- [S6] Official TestCafe Browsers and Native Automation guides, https://testcafe.io/documentation/402828/guides/intermediate-guides/browsers and https://testcafe.io/documentation/404237/guides/intermediate-guides/native-automation-mode, accessed 2026-08-22.
- [S7] Official TestCafe Built-In Wait Mechanisms, Quarantine Mode, and Concurrent Runs guides, https://testcafe.io/documentation/402827/guides/advanced-guides/built-in-wait-mechanisms, https://testcafe.io/documentation/403841/guides/intermediate-guides/quarantine-mode, and https://testcafe.io/documentation/403626/guides/intermediate-guides/run-tests-concurrently, accessed 2026-08-22.
- [S8] Official TestCafe HTTP interception and API testing documentation, https://testcafe.io/documentation/402842/guides/intermediate-guides/intercept-http-requests and https://testcafe.io/documentation/403971/guides/intermediate-guides/api-testing, accessed 2026-08-22.
- [S9] Official TestCafe Screenshots and Videos guide, https://testcafe.io/documentation/402840/guides/intermediate-guides/screenshots-and-videos, accessed 2026-08-22.
- [S10] Official TestCafe GitHub Actions guide and archived testcafe-action repository, https://testcafe.io/documentation/402817/guides/continuous-integration/github-actions and https://github.com/DevExpress/testcafe-action, accessed 2026-08-22.
- [S11] Official TestCafe Docker guide, tagged Dockerfile, and Docker Hub image metadata, https://testcafe.io/documentation/402838/guides/advanced-guides/use-testcafe-docker-image, https://github.com/DevExpress/testcafe/blob/v3.7.6/docker/Dockerfile, and https://hub.docker.com/r/testcafe/testcafe, accessed 2026-08-22.
- [S12] Official TestCafe roadmap, release notes, changelog, and Chrome replay guide, https://testcafe.io/402949/roadmap, https://testcafe.io/release-notes, https://github.com/DevExpress/testcafe/blob/v3.7.6/CHANGELOG.md, and https://testcafe.io/documentation/403998/guides/advanced-guides/chrome-replay-support, accessed 2026-08-22.
- [S13] Official TestCafe selector, action, test structure, ClientFunction, Role, lifecycle, and CLI/API references under https://testcafe.io/documentation/402632/api and https://testcafe.io/documentation/402635/guides/overview/getting-started, accessed 2026-08-22.
- [S14] Official TestCafe reporters, debug, Visual Selector Debugger, and Live Mode documentation under https://testcafe.io/documentation/402825/guides/intermediate-guides/reporters and https://testcafe.io/documentation/402835/guides/basic-guides/debug-tests, accessed 2026-08-22.
- [S15] npm package page for testcafe, including dependent/package summary, https://www.npmjs.com/package/testcafe, observed 2026-08-22.
- [S16] axe-testcafe repository metadata, https://github.com/helen-dikareva/axe-testcafe, observed 2026-08-22.
- [S17] GitHub repository search for TestCafe visual-regression integrations and their activity/archive metadata, observed 2026-08-22.
- [S18] Stack Overflow testcafe tag summary and newest-question listing, https://stackoverflow.com/questions/tagged/testcafe, observed 2026-08-22.
- [S19] Open Asset Allocator repository evidence: .nvmrc, frontend package.json and tsconfig.json, Parcel proxy, HTMX/Navigo/Handlebars/Chart.js code, Gin server, Docker Compose definitions, lifecycle scripts, Makefile, and GitHub Actions workflows, inspected 2026-08-22.
- [S20] TestCafe 3.7.6 README and TestCafe Studio comparison, https://github.com/DevExpress/testcafe/blob/v3.7.6/README.md, accessed 2026-08-22.
- [S21] Sampled 2026 TestCafe issues and pull requests, including #8549, #8546, #8494, #8493, and releases/security PRs, https://github.com/DevExpress/testcafe/issues and https://github.com/DevExpress/testcafe/pulls, observed 2026-08-22.
- [S22] Context7 index for /devexpress/testcafe, used to cross-check official repository documentation, queried 2026-08-22.
- [S23] Docker Hub API metadata for testcafe/testcafe:3.7.6 and latest, including digest, size, architecture, and push date, retrieved 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence in package version, Node engine, bundled TypeScript version, license, release history, npm downloads, GitHub snapshot, official-action deprecation, Docker architecture/size, protocol split, and documented core APIs because these come from tagged source, registries, APIs, and first-party documentation.<br>
> Medium confidence in application fit, ecosystem freshness, maintenance trajectory, and harness burden because they combine documented behavior with repository-specific reasoning.<br>
> Low confidence in TypeScript 6 behavior, arm64 support, issue-response distributions, resource/flake performance, and complete application behavior because no project spike or benchmark was run.

#### Deal Breakers

> For a new suite selected specifically for current maintenance and community adoption, the decisive risks are declining npm usage, slow feature cadence, stale roadmap/documentation, an archived official GitHub Action, TypeScript 4.9.5 bundled against this repository's TypeScript 6, and incomplete native automation outside Chromium.<br>
> Exclude TestCafe if first-class TypeScript 6, Linux WebKit, native cross-engine protocols, per-test browser contexts, trace replay, maintained accessible locators, built-in visual comparison, or native CI sharding/report merge are mandatory.<br>
> It remains serviceable for an existing stable TestCafe suite that does not need those capabilities.<br>
> [S1-S18]

### Uncertain Fields

- `node_and_typescript_compatibility`
- `operating_system_support`
- `issue_health`
- `roadmap_and_deprecation_risk`
- `dependency_currency`
- `wrapper_upstream_lag`
- `maintainer_concentration`
- `github_metrics`
- `dynamic_dom_synchronization`
- `test_isolation`
- `resource_usage`
- `security_and_supply_chain`
- `hard_gate_result`
- `project_fit_score`
- `recommendation`
- `empirical_project_spike_result`

<a id="puppeteer-with-vitest-or-jest"></a>
## 7. Puppeteer with Vitest or Jest

Source result: `Puppeteer_with_Vitest_or_Jest.json`

### Project And Compatibility

#### Implementation Language

> Puppeteer, Vitest, and Jest are implemented primarily in TypeScript and run on Node.js.<br>
> E2E tests can be authored in TypeScript or JavaScript with Puppeteer's bundled type declarations.<br>
> This evaluation uses Puppeteer with Vitest as the preferred composition and treats Jest as a viable runner alternative.

#### Operating System Support

> Puppeteer supports Windows x64, macOS x64/arm64, and Debian/Ubuntu or Fedora/openSUSE Linux x64 for Chrome for Testing; Firefox follows Mozilla platform requirements.<br>
> Local x64 Linux and GitHub-hosted Ubuntu x64 are suitable.<br>
> Official Chrome for Testing Linux support is x64 in Puppeteer's system requirements, and an open project task still tracked Linux arm64 documentation, so arm64 Linux should not be assumed for the standard downloaded Chrome artifact.

#### License And Governance

> Puppeteer is Apache-2.0 and is led by Google's Chrome browser-automation team in a public GitHub organization.<br>
> Vitest is MIT, community-developed with public team and OpenCollective information, and its site identifies VoidZero Inc. and contributors.<br>
> Jest is MIT and governed under the OpenJS Foundation.<br>
> These licenses permit repository and CI use without a paid service; normal license and notice obligations remain.

#### Installation Model

> Preferred baseline: npm install --save-dev puppeteer vitest.<br>
> Puppeteer postinstall downloads a version-matched Chrome for Testing of about 282 MB on Linux and chrome-headless-shell into $HOME/.cache/puppeteer; Firefox is skipped by default and can be enabled in Puppeteer configuration followed by npx puppeteer browsers install.<br>
> Linux browser libraries and unzip are required, and Firefox extraction needs xz or bzip2.<br>
> The alternative is puppeteer-core with a separately managed browser.<br>
> Jest would replace Vitest and may add @jest/globals plus a TypeScript transformer.<br>
> A versioned ghcr.io/puppeteer/puppeteer image is available.

#### Candidate Scope And Layer

> A composed browser-automation stack, not a complete E2E framework.<br>
> Puppeteer supplies browser control, locators, contexts, network events, screenshots, video, and low-level tracing; Vitest or Jest supplies discovery, assertions, hooks, retries, workers, and reports.<br>
> Application lifecycle, database fixtures, visual baselines, and integrated failure bundles remain project-owned.

#### Authoring And Async Model

> Tests use ordinary TypeScript or JavaScript with native async/await.<br>
> The preferred API combines Vitest describe/test/expect and typed fixtures with Puppeteer Browser, BrowserContext, Page, and Locator objects.<br>
> There is no command queue or required DSL.<br>
> Every Puppeteer promise must be awaited, and helper fixtures should centralize browser/context creation, page cleanup, synchronization, and failure artifacts.

#### Build Pipeline Coupling

> Puppeteer is black-box browser automation and can navigate to the Parcel development server or the production Parcel output served by Gin without transforming application source.<br>
> Vitest uses Vite to transform E2E test modules, but that Vite dependency is confined to the test runner and does not replace or instrument the application's Parcel pipeline.<br>
> Jest can avoid Vite but needs its own TypeScript/ESM transform configuration.<br>
> Neither option requires a production frontend build change.

#### Testability Instrumentation Required

> No injected runtime, special browser script, relaxed content-security policy, or frontend build hook is required.<br>
> Existing labels, roles, text, IDs, and CSS can drive the application.<br>
> Stable data-test attributes are optional where templates lack unique semantic selectors.<br>
> Deterministic API seed/reset facilities and a controlled Yahoo Finance boundary may still be needed for reliable tests, but those are application testability concerns rather than Puppeteer requirements.

### Maintenance Health

#### Latest Stable Release

> Puppeteer 25.8.0 was released on 2026-08-17.<br>
> The preferred runner Vitest was 4.1.11, released on 2026-08-18; the alternative Jest was 30.4.2, released on 2026-05-09.

#### Release Cadence

> Puppeteer releases are very frequent and track browser changes: 25.5.0 on August 4, 25.6.0 on August 11, 25.7.0 on August 13, and 25.8.0 on August 17, 2026.<br>
> Recent releases included Chrome/Firefox rolls, browser-management fixes, and API work.<br>
> Vitest also shipped a current 4.1.11 patch on August 18.<br>
> Jest's stable cadence is slower, although its main branch remained active after the May 30.4.2 release.

#### Repository Activity

> puppeteer/puppeteer was pushed on 2026-08-21.<br>
> The ten newest inspected commits spanned August 19-21 and included a Chrome 152 roll, Firefox 154 roll, frame behavior, WebMCP, Windows reliability, tests, and dependency updates from multiple humans and bots.<br>
> Vitest and Jest also had human-authored commits on August 21-22, showing that both runner choices remained active.

#### Dependency Currency

> Puppeteer 25.8.0 requires Node >=22.12, was published with Node 24.15, and pins puppeteer-core 25.8.0, @puppeteer/browsers 3.2.1, chromium-bidi 17.0.2, and current protocol data.<br>
> Its repository rolled Chrome 152 and Firefox 154 immediately before observation.<br>
> Vitest 4.1.11 directly supports Node 24 and Vite 6-8; Jest 30.4.2 declares Node 24 support.<br>
> This is current enough for the repository's Node 24.12 and TypeScript 6 baseline.

#### Wrapper Upstream Lag

> There is no Puppeteer wrapper lag when Vitest or Jest calls Puppeteer directly; the browser library and its managed revisions are released together, while the runner is protocol-agnostic.<br>
> The optional jest-puppeteer preset is a separate community wrapper: version 11.0.0 was published in December 2024 with a broad puppeteer >=19 peer range.<br>
> Avoiding that preset and owning a small fixture removes its release-lag risk.<br>
> Vitest requires no Puppeteer adapter.

### Community Adoption

#### Npm Downloads

> For the exact npm window 2026-08-14 through 2026-08-20, puppeteer received 9,775,888 downloads, vitest received 78,083,394, and jest received 38,752,856.<br>
> These are separate package counts and must not be added together.

#### Ecosystem Usage

> Puppeteer has extensive use in browser automation, testing, scraping, rendering, and tooling; official Chrome DevTools Recorder workflows and Puppeteer Replay can produce or replay automation.<br>
> Jest publishes an official Puppeteer integration guide, while Vitest's generic typed fixtures compose directly without an adapter.<br>
> Maintained adjacent packages include @axe-core/puppeteer 4.13.0, image comparison libraries, CI reporters, and integrations from browser-cloud vendors.<br>
> These integrations are useful but do not constitute one unified E2E lifecycle.

#### Community Support

> Versioned official Puppeteer documentation covers installation, browser management, locators, network interception, debugging, Docker, protocols, and APIs.<br>
> Support material is available through an active GitHub tracker, the large Stack Overflow puppeteer tag, Chrome Developers channels, and extensive third-party examples.<br>
> Vitest and Jest both provide mature runner documentation and community channels.<br>
> Puppeteer's GitHub repository does not enable Discussions, so issue and external-forum support are more important.

#### Adoption Trend

> Strong positive package-install trend on a matched seven-day comparison. puppeteer rose from 5,135,936 downloads in 2025-08-14 through 2025-08-20 to 9,775,888 in the matching 2026 window, about 1.90 times.<br>
> Vitest rose about 5.74 times and Jest about 1.31 times over their matching windows.<br>
> Same-week development and nearly 95,500 Puppeteer stars reinforce continued interest, while downloads remain installation traffic rather than unique teams.

#### Adoption Metric Normalization

> The primary browser-layer metric is the direct npm package puppeteer, all versions, for 2026-08-14 through 2026-08-20, compared with the identical package and calendar dates in 2025.<br>
> It includes CI reinstalls, transitive dependencies, bots, scraping, PDF generation, and non-test automation, so it overstates E2E adoption.<br>
> Vitest and Jest counts are reported separately and mostly represent unit/integration testing unrelated to Puppeteer.<br>
> GitHub Puppeteer metrics cover its whole browser-automation project, not this exact stack.

### Browser And Runtime Coverage

#### Browser Engines

> Official first-class targets are Chrome for Testing and stable Firefox.<br>
> Chrome/Chromium installations can be selected through executablePath or supported channels, but compatibility is guaranteed for the revision paired with the Puppeteer release.<br>
> There is no WebKit or native Safari target and no first-class Edge project.<br>
> This is broader than historical Chrome-only Puppeteer but materially narrower than a Chromium/Firefox/WebKit framework.

#### Browser Protocol

> Chrome uses Chrome DevTools Protocol by default.<br>
> Firefox uses WebDriver BiDi by default, and Chrome can opt into WebDriver BiDi.<br>
> Puppeteer documents that Chrome remains on CDP by default because not every CDP feature is available through BiDi; unsupported BiDi operations throw UnsupportedOperation.<br>
> CDP provides deep Chrome control but is browser/version-coupled, while BiDi improves cross-browser standardization as its implementation matures.

#### Headless And Headed Modes

> Chrome for Testing and Firefox can run headless in CI or headed for local diagnosis.<br>
> Chrome for Testing uses the same modern browser path for headed and default headless modes; the older chrome-headless-shell is separately available with headless set to shell.<br>
> Puppeteer defaults to headless and supports headless false, slowMo, and DevTools for interactive debugging.

#### Browser Version Management

> The puppeteer package downloads a tested Chrome for Testing revision and chrome-headless-shell into a configurable global cache; Firefox can be enabled and downloaded through configuration.<br>
> Versions are mapped to Puppeteer releases, and updating Puppeteer may require a new browser download. puppeteer-core skips downloads for system or remote browser management.<br>
> Deterministic CI should pin the package/lockfile and image or browser cache rather than use latest or an auto-updated executable.

#### Parallel Browser Support

> Puppeteer can launch multiple browsers, pages, or contexts, but it has no browser-project matrix or test scheduler.<br>
> Vitest projects, workers, maxWorkers, fileParallelism, and native file sharding can supply those controls; Jest offers projects, workers, and sharding.<br>
> The harness must map each runner project to Chrome or Firefox configuration, name artifacts by browser, cap concurrent browser processes, and allocate database state safely.

#### Mobile Emulation

> Puppeteer provides KnownDevices and Page.emulate for device viewport and user-agent presets, plus viewport, touch, device scale, geolocation, orientation-related metrics, and low-level CDP controls.<br>
> This supports responsive and mobile-like Chrome testing, not physical Android/iOS execution.<br>
> Some emulation functions are protocol/browser-specific and should not be assumed identical under Firefox BiDi.

#### Real Browser Fidelity

> The managed targets are real Chrome for Testing and stable Firefox binaries rather than generic in-process DOM simulators.<br>
> Chrome for Testing is designed for automation and tracks Chrome code, and current Firefox support uses stable releases.<br>
> No execution here represents Apple Safari: neither Chrome nor Firefox is WebKit, and even a generic WebKit build would not be equivalent to Safari's macOS packaging and integrations.

#### Environment Determinism Controls

> Puppeteer exposes viewport/device descriptors, user agent, locale, timezone, geolocation, permissions, media type/features such as reduced motion, color scheme, network/offline and CPU conditions, cache control, and script/style injection.<br>
> BrowserContext isolates cookies and local storage.<br>
> There is no integrated cross-browser test clock or seeded randomness fixture equivalent to a complete runner's deterministic environment; clock, Math.random, animations, backend time, and server randomness require evaluateOnNewDocument, CDP, CSS, application configuration, or custom fixtures.

### Application Fit

#### Dynamic Dom Synchronization

> Puppeteer Locators are a good but incomplete fit for HTMX swaps.<br>
> They retain a query rather than a stale ElementHandle, wait for presence, visibility, enabled state, viewport placement, and a stable bounding box, and can retry an action.<br>
> Lower-level waitForSelector and stored handles do not retry detached actions.<br>
> Puppeteer supplies no web-first assertion library, so tests must pair locators with waitForFunction, DOM/response predicates, or Vitest expect.poll and assert a specific post-swap state rather than rely on sleeps or blanket network idle.

#### Routing Support

> Page.goto, url, goBack, goForward, reload, waitForNavigation, and waitForFunction can exercise direct Navigo deep links and history transitions.<br>
> Because pushState routing may not create a document navigation, the harness should wait for the expected URL and a route-specific DOM condition instead of always using waitForNavigation.<br>
> A fresh BrowserContext/page and explicit route cleanup prevent history and storage leakage between tests.

#### Locator Model

> Recommended locators support CSS plus Puppeteer text, XPath, ARIA name/role, shadow-DOM piercing, filters, and custom query handlers.<br>
> ARIA selectors use the browser accessibility tree and are preferable for labels and controls.<br>
> Puppeteer does not provide the same dedicated getByRole/getByLabel/test-id vocabulary or strict single-match failure behavior as leading integrated E2E runners; selectors that match multiple elements need an explicit uniqueness convention, and ElementHandles must not be retained across HTMX replacement.

#### Form Interaction

> Locator fill, click, hover and scroll plus Page/ElementHandle focus, select, type, keyboard, mouse, touchscreen, and real Tab/click interactions cover focus, blur, input, change, native validation, and dynamic rows.<br>
> Locator fill chooses behavior by input type.<br>
> Hidden synchronized values can be inspected through evaluate, while script-setting fields should be avoided because it bypasses user behavior.<br>
> Dynamic rows require locator reacquisition and an explicit DOM-ready assertion.

#### Canvas And Download Support

> Chart.js canvas can be checked through visible legends/labels, application data, page.evaluate inspection of Chart.js state or pixel data, screenshots, and a project-selected image diff library; semantic/data assertions should remain primary.<br>
> Puppeteer's official files guide states that downloads cannot currently be handled programmatically.<br>
> Chrome download behavior or raw CDP plus filesystem polling can direct and detect files, but filename, completion, payload verification, cleanup, and cross-browser behavior are custom and materially weaker than an integrated Download object.

#### Network And Api Access

> Puppeteer exposes request, response, request-failed, console, and page-error events; waitForRequest/waitForResponse; and request interception with abort, continue, or respond.<br>
> Enabling interception stalls every request until handled and asynchronous handlers require careful resolution checks.<br>
> Node fetch can seed/reset REST fixtures.<br>
> Browser-originated Yahoo requests can be intercepted, but server-side Gin Yahoo access needs an application-side stub, proxy, or configuration because browser interception cannot observe backend egress.

#### Same Origin Support

> The browser can target the consolidated production Gin origin or Parcel's development origin and existing API proxy without source injection or CORS changes.<br>
> A BASE_URL fixture is sufficient.<br>
> If tests or browsers run in containers, localhost is container-relative, so the harness must use published host ports, a host-gateway mapping, or shared Compose networking.

#### Test Isolation

> Puppeteer BrowserContexts isolate cookies and local storage and can be created per test; pages and contexts must be closed manually in fixture cleanup.<br>
> Vitest 4.1 typed fixtures can create test-, file-, or worker-scoped resources and register cleanup.<br>
> Browser isolation does not isolate PostgreSQL, so state-mutating tests still need unique data, a reset API, separate schemas/databases, or serial execution.<br>
> A browser can be worker-scoped while each test receives a fresh context and page.

#### External Server Model

> Puppeteer naturally targets any independently running HTTP server.<br>
> A configurable base URL can point at Parcel development, the production Compose stack, or a deployed environment.<br>
> Neither Puppeteer nor Vitest/Jest provides a Playwright-style webServer coordinator, so process launch, dynamic ports, HTTP readiness, and service teardown need scripts, global setup, or CI steps.

#### Application Lifecycle

> Vitest global setup and typed fixtures, or Jest globalSetup/globalTeardown and a custom environment, can invoke the build, start Docker Compose, wait for Flyway/backend readiness, and tear down.<br>
> The repository's Compose dependencies already gate the backend on healthy PostgreSQL and successful migration.<br>
> The stack still needs a repository-owned process wrapper for readiness polling, disposable volumes, fixture seeding, cancellation signals, log capture, and guaranteed docker compose down --volumes; Puppeteer only owns browser lifecycle.

#### Visual Regression Workflow

> Puppeteer provides page/element screenshots but no baseline store, update command, masking, tolerance policy, per-browser snapshot paths, diff viewer, or CI review flow.<br>
> Vitest/Jest snapshots do not turn screenshots into a complete visual workflow without a matcher and image-diff library.<br>
> The project must add tools such as pixelmatch or an image-snapshot matcher, stabilize fonts/OS/browser/viewport/data/time/animations, retain expected/actual/diff images, and keep separate Chrome/Firefox baselines.<br>
> Deterministic Chart.js data and completed animations are mandatory.

#### Accessibility Audit Integration

> Deque's maintained open-source @axe-core/puppeteer 4.13.0 injects axe into frames and returns violations suitable for Vitest/Jest assertions and reports.<br>
> Puppeteer also exposes accessibility-tree inspection and ARIA locators.<br>
> There is no integrated ARIA-snapshot assertion/reporting workflow in this composition, so reusable audit fixtures, exclusions, attachments, and human-readable output require harness code.<br>
> Automated axe checks still cover only part of accessibility.

### Reliability

#### Waiting Model

> Locator actions wait for presence and applicable viewport, visibility, enabled, and stable-bounding-box conditions and may retry before acting.<br>
> Explicit primitives include locator.wait, waitForSelector, waitForFunction, waitForNavigation, waitForRequest/Response, and waitForNetworkIdle.<br>
> The lower-level APIs do not automatically retry failed actions, and ordinary Vitest/Jest expect assertions are immediate unless wrapped in Vitest expect.poll or a custom eventually helper.<br>
> Reliable HTMX tests therefore need domain-specific postcondition waits and must prohibit fixed sleeps.

#### Flake Controls

> Puppeteer supplies configurable default/action/navigation timeouts and deterministic event/condition waits.<br>
> Vitest adds test/hook/teardown timeouts, retries with count/delay/error conditions, worker controls, bail, filters, and fixtures; Jest offers per-test timeouts and jest.retryTimes.<br>
> Repeated whole-suite execution requires a script or CI loop.<br>
> Retries should start fresh contexts and preserve first-attempt evidence instead of sharing a failed page, and database state must be reset before each attempt.

#### Isolation Model

> The recommended model is one worker-scoped Browser with a fresh BrowserContext and Page per test, closed through fixture cleanup, combined with unique or reset database state.<br>
> A hard browser failure should recreate the worker/browser.<br>
> Serial scenarios can intentionally share a context or database fixture, but that increases leakage risk.<br>
> Vitest process/file isolation does not automatically isolate a separately shared Puppeteer Browser or PostgreSQL, so ownership must be explicit.

#### Parallelism Controls

> Vitest runs files in parallel by default and supports maxWorkers, fileParallelism false, isolated pools, separate projects, and file-level --shard; projects can separate serial database tests from independent browser tests.<br>
> Jest has maxWorkers, runInBand, projects, and sharding.<br>
> Start the state-mutating E2E project with fileParallelism false or one worker, then increase only after assigning independent databases/schemas and artifact directories.<br>
> Puppeteer itself does not schedule sessions or protect shared state.

#### Flake Observability

> Vitest's GitHub Actions summary identifies tests that passed only after retries, and JSON/JUnit/blob/custom reporters can retain attempt status; test context hooks can attach first-failure screenshots and logs.<br>
> Jest reporters and retry hooks can expose similar data with more custom work.<br>
> Neither Puppeteer nor these local runners supplies a long-term quarantine registry, historical flake dashboard, or root-cause classification.<br>
> The repository must preserve attempt numbers, first-failure artifacts, repeat-run results, and trend data.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Puppeteer core can capture page/element screenshots, HTML content, console/page errors, request/response failures, Chrome performance traces, heap/coverage data, and WebM page screencasts when ffmpeg is installed.<br>
> Its trace is a DevTools performance trace, not an action-by-action DOM snapshot trace.<br>
> Nothing automatically starts, stops, or attaches these on test failure.<br>
> Vitest annotations/file attachments or Jest custom reporters can associate files, but the fixture must collect them per attempt and ensure recorder/context closure.

#### Debugging Tools

> Developers can run headed, add slowMo, open DevTools, forward browser stdio with dumpio, log protocol traffic, inspect pending protocol calls, and debug Node code with --inspect-brk.<br>
> Vitest adds file/name filters, watch mode, UI/HTML output, and IDE integration; Jest has mature watch and debugger workflows.<br>
> There is no Puppeteer-specific step inspector, locator picker, automatic pause-on-failure, or Playwright-style time-travel trace viewer.

#### Test Generation

> Puppeteer core has no integrated recorder in the runner.<br>
> Chrome DevTools Recorder can record browser flows and export/replay Puppeteer scripts through adjacent Puppeteer Replay tooling.<br>
> Generated scripts are useful scaffolding but need conversion into Vitest/Jest tests, resilient locators, domain waits, fixtures, cleanup, and assertions before production use.

#### Reporters

> Vitest provides default/tree/verbose/dot/minimal, JSON, JUnit, TAP, HTML through @vitest/ui, GitHub Actions, blob, and custom reporters; blob reports support shard merging and attachments are handled separately.<br>
> Jest provides console, JSON, custom reporters, and a large reporter ecosystem.<br>
> Puppeteer contributes no reporter.<br>
> The preferred Vitest stack can emit GitHub annotations and a job summary, but screenshots, traces, videos, and service logs must be explicitly attached and linked.

#### Documentation Quality

> Puppeteer has comprehensive, versioned API documentation and practical guides for installation, locators, browsers, protocols, Docker, request interception, and debugging.<br>
> The docs accurately expose limitations such as no programmatic download handling and the difference between CDP and BiDi.<br>
> Vitest and Jest runner docs are extensive.<br>
> The composition lacks one authoritative E2E architecture guide, so lifecycle, fixture, retry-artifact, and database-isolation decisions remain spread across three projects and repository code.

#### Local Workflow

> With Vitest, developers can run one file or title filter, use watch mode or the UI, select a project/browser through configuration, disable file parallelism, and debug Node plus a headed browser.<br>
> Puppeteer supports slow motion and DevTools.<br>
> A project fixture can expose BASE_URL and browser choice.<br>
> Keeping a failed browser open, replaying only a retry attempt, and viewing a unified browser/backend timeline require custom flags or hooks rather than a dedicated E2E UI.

#### Failure Log Correlation

> Page events can timestamp console messages, uncaught errors, requests, responses, and failures; Vitest task metadata supplies test identity and attachments.<br>
> Gin, PostgreSQL, Flyway, and Docker Compose logs can be captured under the same run/test identifier in always-run workflow steps.<br>
> Puppeteer and the runner do not correlate these processes automatically.<br>
> The harness must establish UTC clocks, test/attempt IDs, log boundaries, artifact names, and capture ordering, especially for parallel workers.

#### Artifact Data Exposure

> Screenshots, videos, HTML, console output, response bodies, protocol logs, traces, downloads, and browser profiles can expose portfolio records, cookies, authorization headers, Yahoo responses, and credentials.<br>
> NODE_DEBUG protocol logs are explicitly warned as sensitive.<br>
> There is no universal redaction layer.<br>
> Use synthetic data, fresh contexts, short restricted retention, selective network-body capture, header redaction, UI masking before screenshots where possible, and never upload cache/profile or unreviewed trace/video data by default.

### Github Actions Fit

#### Official Ci Support

> Puppeteer maintains installation/troubleshooting guidance and a versioned official GHCR image containing Chrome for Testing, dependencies, and Puppeteer.<br>
> A normal GitHub-hosted Ubuntu job can use setup-node, npm ci, browser/system dependency installation, and vitest run.<br>
> Vitest has documented GitHub Actions sharding/report merging and a built-in GitHub reporter.<br>
> No dedicated Puppeteer GitHub Action is required, and application startup remains repository-specific.

#### Browser Caching

> Puppeteer caches managed browsers under $HOME/.cache/puppeteer by default, so GitHub Actions can cache that directory alongside npm's package cache.<br>
> Keys should include OS, architecture, lockfile/Puppeteer version, and selected browser set.<br>
> A restored browser must match the pinned Puppeteer revision; otherwise npx puppeteer browsers install should repair/populate it.<br>
> The default Chrome payload is large, but caching must not preserve browser profiles or downloaded test data.

#### Artifact Integration

> Vitest can write JSON, JUnit, HTML, blob reports, annotations, and file attachments; Puppeteer artifacts can be placed in per-test attachment directories.<br>
> GitHub actions/upload-artifact can upload reports, screenshots, traces, videos, diffs, downloads, and Compose logs under an always or not-cancelled policy.<br>
> This is mechanically straightforward but not automatic: fixture hooks must create and annotate the files, and retention/access need a sensitive-data policy.

#### Sharding And Matrix Support

> Vitest natively shards test files with --shard, emits blob reports, and merges them after GitHub matrix jobs; its documentation includes a complete Actions pattern.<br>
> Browser projects or a matrix can select Chrome and Firefox.<br>
> Puppeteer does not define projects or merge artifacts.<br>
> Each shard/browser needs a unique database or must avoid state mutation, a unique artifact path, and deterministic file assignment.<br>
> Jest also supports sharding, but this evaluation prefers Vitest's documented blob merge flow.

#### Container Compatibility

> Host-run Node/Puppeteer can test the repository's Docker Compose services through published ports on GitHub Ubuntu.<br>
> The official Puppeteer image is versioned and includes Chrome dependencies, but its documented sandboxed invocation requests SYS_ADMIN and --init, so it should not be adopted blindly as a least-privilege CI default.<br>
> If tests run in a container, localhost cannot reach host/other-container services without explicit networking.<br>
> Shared memory, init, user/sandbox settings, image pinning, and service health checks require configuration.

#### Failure Cleanup

> Fixtures must stop screencasts/traces, close pages/contexts, and close or recreate the browser in finally/onCleanup hooks.<br>
> Vitest's AbortSignal can react to test timeout or cancellation, but hard job cancellation can bypass Node teardown.<br>
> GitHub Actions should collect logs first and then run docker compose down --volumes --remove-orphans under if: always(), delete downloads/artifacts according to retention, and use unique Compose project/volume names.<br>
> Puppeteer has no application/database cleanup coordinator.

### Cost And Risk

#### Open Source Completeness

> Puppeteer, Vitest/Jest, Chrome for Testing, Firefox, screenshots, screencasts, tracing, reporters, axe integration, and open-source image-diff libraries can provide the mandatory local and Linux CI capability without licensing fees.<br>
> No hosted service is required.<br>
> Completeness is achieved by composition and custom code rather than an integrated open-source E2E product, and native Safari is outside the stack.

#### Optional Cloud Dependency

> None for Chrome/Firefox execution, local reports, screenshots, video, accessibility checks, or sharding.<br>
> Browser clouds and visual-review dashboards are optional for broader browser/device capacity, native Safari, analytics, and hosted baseline review.<br>
> Local browser binaries, GitHub artifacts, committed image baselines, and Vitest/Jest reports provide non-cloud alternatives for the evaluated baseline, but not native Safari on Ubuntu.

#### Migration Cost

> Moderate.<br>
> Tests use standard TypeScript, async/await, Node fetch, and familiar runner assertions, but page actions, locators, contexts, interception, traces, and browser setup bind to Puppeteer.<br>
> Runner-neutral domain fixture APIs and page-object/helper boundaries can preserve scenario logic if moving to Playwright or WebDriver.<br>
> Visual/lifecycle code is project-owned and may be reusable, but that ownership is also ongoing maintenance cost.

#### Security And Supply Chain

> Puppeteer 25.8.0 has six direct npm dependencies in registry metadata, integrity signatures, and a SLSA provenance attestation published by GitHub Actions; Vitest 4.1.11 also exposes signatures/provenance, while Jest is an additional large dependency tree if selected.<br>
> Installation downloads executable browser artifacts and Linux packages, introducing runtime egress and binary-update risk.<br>
> Pin lockfiles and image digests/tags, verify npm integrity/provenance, scan dependencies/images, preserve the browser sandbox, restrict external URLs, and update quickly for browser security releases.

#### Custom Harness Burden

> High.<br>
> The project must own Vitest/Jest configuration; browser/context/page fixtures; browser matrix; HTMX/Navigo waits; assertion polling; REST and PostgreSQL seed/reset; server-side Yahoo replacement; Compose/Flyway readiness and cleanup; programmatic download completion; visual baselines/diffs; retry-safe screenshot/video/trace collection; service-log correlation; artifact redaction; and database-safe parallelism.<br>
> Puppeteer reduces browser-management and actionability work compared with raw WebDriver but does not remove integrated-runner lifecycle work.

#### Capability Delivery Tier

- **Core Puppeteer:** Chrome and Firefox control, CDP/BiDi transports, locators and actionability waits, browser contexts, input/navigation, device/environment emulation, request events/interception, screenshots, DevTools traces, screencasts, and browser downloads/management.
- **Runner Vitest Or Jest:** Discovery, assertions, hooks, fixtures or environments, timeouts, retries, worker limits, filters, projects, sharding, and machine-readable reports. Vitest is preferred for typed fixture scopes, attachments, blob merge, and GitHub reporting.
- **Official Or Maintained Adjacent:** Official Puppeteer Docker image and Chrome DevTools Recorder/Puppeteer Replay; maintained @axe-core/puppeteer accessibility audits; community image matchers and report viewers.
- **Custom Repository Code:** Application readiness and teardown, disposable PostgreSQL state, server-side external-service mocking, HTMX route postconditions, download completion/content checks, visual workflow, artifact capture/redaction/correlation, and safe browser/database sharding.
- **Optional Paid Cloud:** Native Safari/device capacity, hosted visual review, remote browser grids, and historical analytics; none is required for the Linux Chrome/Firefox baseline.

#### Ai Execution Boundary

> Puppeteer and Vitest/Jest execute deterministic committed code without an LLM, AI credentials, model egress, nondeterministic model choices, or per-run AI cost.<br>
> Chrome DevTools Recorder is local deterministic recording rather than required AI.<br>
> Optional AI authoring or MCP tools should remain outside required CI, use synthetic data and least privilege, and produce reviewed tests that continue to run when all AI services are unavailable.

### Evidence And Decision

#### Sources

- Puppeteer 25.8.0 system requirements, installation, configuration, supported browsers, WebDriver BiDi, page interactions, browser management, debugging, network interception, files, screenshots, Docker, tracing, screencast, and API documentation at https://pptr.dev, accessed 2026-08-22.
- Puppeteer npm metadata for 25.8.0 at https://registry.npmjs.org/puppeteer/latest, observed 2026-08-22.
- Puppeteer releases and repository activity at https://github.com/puppeteer/puppeteer and https://github.com/puppeteer/puppeteer/releases, observed through GitHub API data on 2026-08-22.
- npm downloads API for puppeteer at https://api.npmjs.org/downloads/point/2026-08-14:2026-08-20/puppeteer and the matching 2025 window, observed 2026-08-22.
- Vitest 4.1.11 documentation for requirements, fixtures, retries, parallelism, sharding, attachments, CLI, UI, and reporters at https://vitest.dev, accessed 2026-08-22.
- Vitest npm metadata, downloads, release, repository metrics, and commits for 4.1.11 and vitest-dev/vitest, observed 2026-08-22.
- Jest 30.4 documentation for Puppeteer, TypeScript, ESM, setup/teardown, retry, workers, sharding, and reporters at https://jestjs.io/docs, accessed 2026-08-22.
- Jest npm metadata, downloads, release, repository metrics, and commits for 30.4.2 and jestjs/jest, observed 2026-08-22.
- @axe-core/puppeteer 4.13.0 and jest-puppeteer 11.0.0 npm registry metadata, observed 2026-08-22.
- Open Asset Allocator local evidence: .nvmrc, src/main/web-static/package.json, Parcel scripts/proxy, frontend HTMX/Navigo/Handlebars/Chart.js modules, Gin server, Docker Compose PostgreSQL/Flyway dependencies, lifecycle scripts, Makefile, and existing workflows, inspected 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

- **High:** Versions, release dates, license, declared Node/TypeScript ranges, exact npm windows, repository stars/forks, Chrome/Firefox and protocol support, locator behavior, installation size/model, official download limitation, runner controls, and absence of integrated application lifecycle are supported by primary documentation and APIs.
- **Moderate:** Project-specific synchronization design, issue health, maintainer resilience, cross-protocol feature consistency, diagnostics effort, and the weighted score combine documented primitives with engineering analysis.
- **Low:** Open Asset Allocator runtime performance, database-safe concurrency limits, and the complete empirical application proof remain unmeasured.

#### Deal Breakers

> No confirmed incompatibility blocks Chrome/Firefox E2E use.<br>
> Reasons not to choose this as the primary stack are the absence of WebKit/Safari, no programmatic Download object, no integrated retrying assertions, no built-in web-server/Compose lifecycle, no first-class visual-regression workflow, no automatic action trace or failure artifact bundle, and the high custom harness burden.<br>
> A mandatory WebKit/Safari target, minimal repository-owned harness, or integrated trace-driven diagnosis would be an exclusion-level requirement.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `roadmap_and_deprecation_risk`
- `maintainer_concentration`
- `github_metrics`
- `resource_usage`
- `hard_gate_result`
- `project_fit_score`
- `recommendation`
- `empirical_project_spike_result`

<a id="rod-with-go-testing"></a>
## 8. Rod with Go testing

Source result: `Rod_with_Go_testing.json`

### Project And Compatibility

#### Implementation Language

> Rod is implemented in Go and tests are ordinary Go code executed by the standard testing package.<br>
> It embeds a Go CDP client and launches or connects to a separate Chrome/Chromium process.<br>
> Node.js, TypeScript, and an npm test runtime are not required for Rod itself.

#### Operating System Support

> Rod's own workflows test ubuntu-latest, macos-latest, and windows-latest.<br>
> The launcher has browser discovery for Linux, macOS, Windows, and OpenBSD; automatic snapshot downloads explicitly cover common desktop combinations, including Linux amd64, and use a Playwright-hosted Chromium build for Linux arm64.<br>
> GitHub-hosted Ubuntu x64 is therefore a supported practical target.<br>
> The official documentation warns that automatic browser downloads may not work on every Go-supported platform and recommends its Docker launcher for servers.

#### License And Governance

> MIT License.<br>
> The public repository belongs to the go-rod GitHub organization and is overwhelmingly authored by Yad Smood (ysmood).<br>
> The contribution guide describes a community membership and maintainer-election model, but the observed activity does not demonstrate a currently distributed governance team.<br>
> TestMu AI is displayed as a sponsor, not as an owner or support guarantor.<br>
> The MIT terms permit repository and CI use without a paid license.

#### Installation Model

> Add the pinned Go module with go get github.com/go-rod/rod@v0.116.2 or a go.mod requirement.<br>
> On first launch, Rod searches for a browser and otherwise downloads its fixed Chromium revision into $HOME/.cache/rod/browser; Linux still needs Chromium runtime libraries.<br>
> A system Chrome/Chromium/Edge binary can be selected explicitly.<br>
> The project also publishes ghcr.io/go-rod/rod, an Ubuntu Noble remote launcher image containing Chromium, fonts, Xvfb, and dumb-init.<br>
> The application build remains a separate npm/Parcel step.

#### Candidate Scope And Layer

> Browser automation stack, not a complete E2E framework.<br>
> Rod supplies a high-level CDP driver, browser launcher, element and event helpers, request interception, screenshots, downloads, emulation, and pools.<br>
> Go testing supplies discovery and execution; the repository must provide assertions, fixtures, application and database lifecycle, retries, reporting, visual baselines, and artifact policy.

#### Authoring And Async Model

> Tests use normal Go functions such as func TestName(t *testing.T), synchronous-looking Rod calls, contexts, goroutines, channels, and two-step event waits.<br>
> Methods returning errors support explicit Go error handling; Must-prefixed methods panic by default and WithPanic can redirect failures to a test-aware handler that stops the current goroutine.<br>
> There is no command queue, Gherkin layer, keyword DSL, or JavaScript async/await runtime.

#### Build Pipeline Coupling

> Rod navigates to a URL and does not transform, instrument, or bundle frontend source.<br>
> It can exercise the production Parcel build served by Gin or the Parcel development server and its existing API proxy.<br>
> It does not require Vite, a Parcel plugin, injected test runtime, or alternate module settings.<br>
> A Go TestMain or CI job must ensure that the selected frontend and backend topology is built and reachable first.

#### Testability Instrumentation Required

> No Rod-specific production script, route, source transform, or relaxed content-security policy is required.<br>
> Stable IDs and accessible HTML improve selectors but are optional.<br>
> Reliable project tests still need deterministic data setup, the existing Yahoo Finance mock/configuration path, a non-API-only server that serves built frontend assets, readiness checks, and database cleanup.<br>
> Those can remain test infrastructure rather than production browser instrumentation.

### Maintenance Health

#### Latest Stable Release

> v0.116.2, published July 12, 2024.<br>
> It is the latest tagged non-prerelease observed on August 22, 2026; pkg.go.dev also identifies v0.116.2 as the latest module version.<br>
> The module remains below semantic version 1.0.

#### Release Cadence

> Release cadence has stopped.<br>
> Rod published roughly monthly releases from August 2023 through July 2024, including seven releases from January through July 2024, but no tagged release followed during the next 25 months through observation.<br>
> This is a material maintenance-health failure for a browser driver whose protocol and browser security baseline move frequently.

#### Repository Activity

> The default branch was pushed on August 11, 2026, but the 2025-2026 commits inspected were sponsor or documentation edits.<br>
> The newest observed merged code changes were in December 2024.<br>
> Community pull requests continued to arrive, including dependency and memory-state fixes in 2026, while 29 pull requests remained open.<br>
> Recent activity therefore demonstrates user demand but not active integration of browser-driver maintenance.

#### Wrapper Upstream Lag

> Rod is not a wrapper over WebDriver or Playwright, but it generates CDP types and pins a Chromium revision.<br>
> Official compatibility documentation guarantees each Rod version only with its launcher.DefaultRevision.<br>
> That revision and the latest release are from the 2024 maintenance period while users are now filing Chrome 150-related work.<br>
> This creates roughly two years of browser/protocol integration lag; using a newer system Chrome may work but lies outside the documented guarantee.

#### Maintainer Concentration

> Very high bus-factor risk.<br>
> GitHub contributor data gives ysmood 715 commits, while the next observed contributor has 9 and the next four have 7 or 6.<br>
> Recent sponsor/documentation commits are also from ysmood, while community code pull requests remain open.<br>
> The contribution guide describes distributed governance in principle, but no organizational owner, funded maintenance team, published succession plan, or active multi-maintainer review distribution was evidenced.

### Community Adoption

#### Npm Downloads

> Not applicable.<br>
> Rod has no official npm package and requires no Node runtime.<br>
> The comparable registry signal is pkg.go.dev's 1,021 known importing Go modules, observed August 22, 2026.

#### Ecosystem Usage

> pkg.go.dev records 1,021 known importing modules.<br>
> The maintainer's adoption discussion lists recognizable users including charmbracelet/vhs, Authelia, ProjectDiscovery Nuclei, Katana and httpx, Storj, Botkube, Sigstore, GoNB, and gowitness.<br>
> Community comments include E2E testing with Testcontainers as well as scraping, screenshots, PDF generation, authentication, and automation.<br>
> The evidence is meaningful for Go browser automation, but much usage is scraping or rendering rather than E2E test suites.

#### Community Support

> The project provides a docs site, Go API reference, examples, repository tests, FAQ, GitHub issues and discussions, and a Discord link.<br>
> The documentation explains launch, selectors, network interception, emulation, debugging flags, and low-level CDP use.<br>
> Troubleshooting often points users to source tests and issue search.<br>
> Community questions continue, but current maintainer response is weak and the docs do not provide an integrated E2E harness guide comparable to complete test frameworks.

#### Adoption Metric Normalization

> There is no npm measurement.<br>
> The 1,021 figure is pkg.go.dev's current count of known modules importing github.com/go-rod/rod; it includes direct and indexed public imports for scraping, PDF/screenshot generation, security tools, automation, and E2E testing, and does not count unique users, private modules, downloads, or active deployments.<br>
> GitHub stars and forks are repository-wide lifetime vanity signals and do not isolate the Go testing composition evaluated here.

### Browser And Runtime Coverage

#### Browser Engines

> Practical first-class coverage is Chromium through the Chrome DevTools Protocol.<br>
> Rod can launch downloaded Chromium or discover installed Chrome and Microsoft Edge.<br>
> It does not provide production-ready Firefox, WebKit, or native Safari projects.<br>
> The official compatibility page says Firefox was supporting CDP and Safari had no plan to do so, but Rod gives no guaranteed Firefox target and current Firefox automation uses other protocols.<br>
> This candidate therefore cannot provide a cross-engine Linux matrix.

#### Browser Protocol

> Rod communicates directly with the browser's Chrome DevTools Protocol endpoint over WebSocket/JSON-RPC and exposes generated low-level proto calls beneath its high-level API.<br>
> It does not use W3C WebDriver or WebDriver BiDi.<br>
> Deep Chromium control enables network and emulation features but ties compatibility to Chrome's evolving CDP and Rod's pinned generated revision.

#### Headless And Headed Modes

> Headless is the launcher default and is suitable for CI.<br>
> Headed mode is enabled with launcher.Headless(false), the -rod=show option, or equivalent defaults.<br>
> The launcher supports the new Chromium headless switch and Xvfb on Linux; the official image installs Xvfb for headed execution.<br>
> DevTools can be opened automatically in headed mode.

#### Browser Version Management

> Rod searches for a local browser or downloads a statically selected Chromium snapshot to its cache.<br>
> A specific revision or binary can be configured, and a remote launcher image can centralize the browser.<br>
> This supports pinning, but only the default revision is guaranteed compatible with a Rod release.<br>
> Since that default has not been released forward since 2024, teams must choose between an old guaranteed browser, an unguaranteed current system browser, or maintaining a fork and compatibility tests.

#### Parallel Browser Support

> Rod operations are documented as thread-safe and it includes browser and page pools.<br>
> Go goroutines and testing parallelism can run multiple pages, incognito contexts, or browser processes.<br>
> There is no declarative browser-project matrix, selected-engine CLI, test shard planner, or merged result format; browser versions and job distribution are project-owned.<br>
> Firefox/WebKit matrices are unavailable.

#### Mobile Emulation

> Page.Emulate and predefined device descriptors set viewport, screen size, device scale factor, user agent, touch/mobile capability, and orientation.<br>
> Additional CDP calls can set related metrics.<br>
> The bundled profiles are dated, including many Chrome 114 user agents and old named devices, and represent Chromium desktop emulation rather than physical Android, iOS, Mobile Safari, or mobile Firefox.

#### Real Browser Fidelity

> Rod can drive installed branded Chrome or Edge, which provides those real Chromium products, or a downloaded Chromium snapshot.<br>
> Automation flags and headless rendering can still differ from an interactive user session.<br>
> There is no WebKit result that can be mistaken for Safari: native Safari is simply unsupported.<br>
> Device emulation does not turn Chromium into Mobile Safari or a physical device.

#### Environment Determinism Controls

> High-level APIs cover viewport/device emulation, user agent, headers, cookies, network conditions, and browser launch environment such as TZ.<br>
> Generated CDP proto calls can additionally control locale, timezone, geolocation, permissions, virtual time, media features, and device metrics where the pinned Chromium supports them.<br>
> Rod provides no unified test fixture for clock, seeded randomness, animation suppression, fonts, or reduced motion; these require low-level calls, JavaScript/CSS injection, launch configuration, and repository conventions.

### Application Fit

#### Dynamic Dom Synchronization

> Rod's single-element queries poll until an element appears, actions can wait for visibility/interactability/stability, and Page.WaitDOMStable, WaitStable, WaitRequestIdle, WaitEvent, Race, and JavaScript conditions provide explicit synchronization.<br>
> This can handle HTMX requests and Handlebars rendering without fixed sleeps.<br>
> Unlike a locator abstraction, an Element is a concrete remote object and is not automatically reacquired after an HTMX replacement; tests must store selector intent and query again before later actions or assertions.<br>
> Assertions from Go libraries are not automatically polled.

#### Routing Support

> Direct navigation, current URL inspection, back/forward/reload, navigation history, and the v0.116.2 ResetNavigationHistory API cover normal browser routing.<br>
> Navigo pushState changes may not emit a full page lifecycle event, so tests need a URL or route-specific DOM polling helper rather than relying only on WaitNavigation.<br>
> Fresh incognito contexts or explicit route reset are needed between tests.<br>
> Gin must continue serving the SPA shell for direct deep links.

#### Locator Model

> Rod supports CSS, XPath, JavaScript, visible-text regular expressions, tree traversal, and search through nested frames and shadow DOM.<br>
> A single-element query waits for a match, but it returns a concrete element, does not enforce strict uniqueness, and has no first-class getByRole, getByLabel, or test-id policy.<br>
> Accessible locators can be encoded with CSS attributes or JavaScript helpers.<br>
> Stable IDs and selector helper functions are important for reacquiring HTMX-replaced nodes.

#### Form Interaction

> Element APIs cover focus, blur, click, input, keyboard typing, key actions, selection, file upload, color/time inputs, touch, and property inspection.<br>
> These can exercise native focus/blur, validation, dynamic rows, hidden-value synchronization, and input event handling rather than mutating application state directly.<br>
> Tests should use user actions for behavior and reserve Eval/Property for observing hidden values or validation state.<br>
> Post-action waits and assertions remain custom.

#### Canvas And Download Support

> Rod can capture full-page or element screenshots and convert a canvas element directly to image bytes with CanvasToImage/MustCanvasToImage, enabling pixel or image decoding in Go.<br>
> Browser.WaitDownload and MustWaitDownload register a two-step wait and return downloaded bytes, after which tests can inspect payload content.<br>
> Rod does not manage visual baselines, masks, tolerances, review, or per-browser diffs; Chart.js semantic access, animation completion, and image comparison must be implemented by the project.

#### Network And Api Access

> Browser- or page-scoped HijackRequests can observe, continue, fail, or modify requests and responses; CDP events and WaitRequestIdle cover observation and synchronization.<br>
> Go's net/http client can seed and reset REST fixtures in the same tests.<br>
> Browser-originated Yahoo Finance traffic can be intercepted, but this application performs external integration server-side, so the existing Go Yahoo Finance mock and overridden backend configuration are the more reliable boundary.<br>
> HAR replay and an integrated API fixture context are absent.

#### Same Origin Support

> Rod controls a normal external browser and can navigate to either the consolidated Gin production origin or Parcel's development origin and API proxy.<br>
> No CORS change is needed when the existing topology keeps browser traffic same-origin.<br>
> If Rod or the browser runs in its remote Docker image, the base URL must use a network-reachable host or service name rather than localhost inside the browser container.

#### Test Isolation

> Browser.Incognito creates an isolated browser context for cookies and storage, and a fresh browser per test is also possible.<br>
> Go t.Cleanup can close pages and contexts.<br>
> Database state remains outside browser isolation.<br>
> This repository already has Testcontainers PostgreSQL/Flyway setup, initialized fixtures, Yahoo Finance mocking, and per-test SQL cleanup builders, which can be reused or factored for E2E tests; unsafe t.Parallel use must be prevented until each worker has independent data or a database.

#### External Server Model

> Rod naturally tests an independently started server: code passes a configured base URL to Browser.Page or Page.Navigate.<br>
> Local Compose, a TestMain-started Gin server, Parcel port 8000, or a CI deployment are all possible.<br>
> Rod has no baseURL fixture, port reservation, or server health manager, so environment parsing and readiness belong in the Go harness.

#### Application Lifecycle

> Rod does not build Parcel, start Gin, run Docker Compose or Testcontainers, wait for Flyway, or tear down PostgreSQL.<br>
> Standard Go TestMain, contexts, t.Cleanup, os/exec, net/http polling, and Testcontainers can supply those concerns.<br>
> This is a relative advantage because the repository already runs PostgreSQL and Flyway through Testcontainers and starts an API-only Gin server in inttest, but E2E work must add built static assets or a Parcel server, browser lifecycle, HTTP readiness, and cancellation-safe cleanup.

#### Visual Regression Workflow

> Core support stops at page, full-page, element, scroll, and canvas image capture.<br>
> Baseline naming/storage, update approval, image diffing, masks, tolerances, font and animation stabilization, CI diff reports, and platform-specific baselines require a Go image library and custom harness.<br>
> The official image includes common fonts, which can improve consistency.<br>
> Chart.js tests should assert deterministic data and canvas readiness before a narrowly scoped image comparison.

#### Accessibility Audit Integration

> Rod can inject and evaluate axe-core JavaScript or call Chromium's generated CDP Accessibility domain, so open-source audits are technically possible.<br>
> There is no official axe integration, accessible locator API, ARIA snapshot assertion, violation formatter, or accessibility report pipeline.<br>
> Script loading, rules, frame handling, result conversion, and artifact output are custom, and automated checks still require manual accessibility review.

### Reliability

#### Waiting Model

> Element lookup retries until presence, and Rod provides visibility, enabled, writable, interactable, stable, DOM-stable, page-stable, load, navigation, request-idle, repaint, event, and arbitrary JavaScript waits.<br>
> Its two-step event APIs register before the triggering action to avoid missed events.<br>
> Timeouts are context-derived and can be applied to browser, page, or element chains.<br>
> There are no web-first assertion matchers, and retained Element objects are not locator-like, so postconditions and HTMX node reacquisition need explicit polling helpers.

#### Flake Controls

> Go testing supplies per-test deadlines through contexts and selective execution but has no built-in retries.<br>
> Rod supplies configurable sleepers/backoff, chained timeouts, event waits, Race selectors, and deterministic explicit waits; the documentation discourages time.Sleep.<br>
> Retry policy, repeat-until-failure scripts, first-attempt preservation, and failure classification must be added around go test.<br>
> Browser/CDP version staleness is an additional flake and breakage risk that test-level retries cannot correct.

#### Isolation Model

> The preferred model is a fresh incognito context or browser per test plus t.Cleanup, with unique or restored PostgreSQL data.<br>
> A browser can be shared at suite scope for speed while contexts isolate cookies and storage.<br>
> Serial stateful scenarios can intentionally share a context, but ordinary tests should not.<br>
> Go package-level TestMain and the repository's existing SQL cleanup facilities fit this model; browser isolation does not serialize or reset database state.

#### Parallelism Controls

> Rod is thread-safe and offers browser/page pools; Go testing exposes -parallel, package parallelism, t.Parallel, and test-name filters.<br>
> Keeping database-mutating tests serial is simple by not calling t.Parallel and using one E2E package, while independent read-only scenarios can use bounded goroutines or tests.<br>
> Cross-job shards, browser matrices, database allocation, and result merging are not Rod features and require GitHub Actions and harness conventions.

#### Flake Observability

> Go test -json exposes test start, output, pass, fail, skip, elapsed time, and rerun results, and custom wrappers can preserve a first failure before retry.<br>
> Rod itself has no retry classification, quarantine registry, flaky-test status, historical database, or local trend dashboard.<br>
> Repeat runs, attempt identifiers, JSON enrichment, and trend aggregation must be implemented or composed from open-source Go tooling.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Rod can produce page, full-page, scroll, element, and canvas screenshots; log CDP traffic; subscribe to console, exception, network, and lifecycle events; and save page/DOM data through APIs.<br>
> Its own CI stores tmp/cdp-log artifacts.<br>
> It does not automatically capture these on every test failure and has no integrated video, HAR, DOM/action trace archive, retention modes, or HTML artifact index.<br>
> The test harness must register t.Cleanup and failure hooks and collect backend/container logs.

#### Debugging Tools

> Headed mode, automatic DevTools opening, slow motion, visual input trace overlays, CDP logging, and remote browser monitoring are built in.<br>
> Normal Go breakpoints and IDE debugging apply because tests are Go.<br>
> There is no test inspector, step runner, time-travel DOM viewer, watch UI, or pause-on-failure workflow comparable to an integrated browser test runner.

#### Test Generation

> Rod has no official browser recorder, scenario code generator, or accessible locator generator.<br>
> Developers select CSS/text/XPath/JavaScript queries with browser DevTools and write Go code manually.<br>
> Third-party generation would not remove the need to design HTMX waits, database fixtures, cleanup, and diagnostics.

#### Reporters

> The standard runner provides human-readable go test output and structured go test -json events.<br>
> JUnit conversion, HTML reports, GitHub annotations, attachment indexing, retry metadata, and shard merging require third-party Go tools or custom code.<br>
> Rod has no reporter API of its own, so browser artifacts must be linked to test names through harness conventions.

#### Documentation Quality

> The API reference is extensive and the docs site covers launch, selectors, network, emulation, debugging, examples, FAQ, and compatibility.<br>
> Source tests provide many practical examples.<br>
> Weaknesses are that documentation is not versioned as a rendered release site, some compatibility language is dated, device/browser references are old, E2E runner architecture is not covered, and current maintenance status is undocumented.

#### Local Workflow

> Developers can run one test with go test -run, run a package, use standard IDE test actions, switch to headed mode with -rod=show, add devtools/slow/trace flags, and debug with Go breakpoints.<br>
> Browser downloads are automatic and cacheable.<br>
> Re-running only failures, watching files, interactively selecting scenarios, and opening a unified report require scripts or IDE facilities rather than Rod.

#### Failure Log Correlation

> Rod loggers and event handlers can timestamp CDP, browser console, exception, and network events, while Go tests can include test names and UTC timestamps.<br>
> The existing integration infrastructure already fans PostgreSQL logs into testing.TB and prints Flyway logs.<br>
> Correlating those with Gin, Parcel, browser, and screenshot artifacts still requires a shared run/test ID, synchronized timestamps, bounded log capture, and custom attachment naming; Rod does not correlate processes automatically.

#### Artifact Data Exposure

> Screenshots, DOM/page content, download bytes, request and response bodies, headers, cookies, console output, CDP logs, and browser user-data directories can expose portfolio records, credentials, and Yahoo responses.<br>
> Rod provides selective event handling but no artifact redaction pipeline.<br>
> Use synthetic fixtures and incognito profiles, avoid persistent user data, filter headers/bodies before writing, restrict artifact access and retention, and do not upload raw CDP traffic by default.

### Github Actions Fit

#### Official Ci Support

> Rod maintains example workflows that run Go tests on ubuntu-latest and upload CDP logs, plus macOS and Windows jobs.<br>
> It publishes an Ubuntu-based remote launcher image.<br>
> There is no Rod-specific GitHub Action or application-test workflow template; a normal job uses setup-go, builds Parcel, starts Testcontainers or Compose, runs go test, and uploads custom artifacts.

#### Browser Caching

> Cache the Go module/build caches and $HOME/.cache/rod/browser using keys that include OS, architecture, Rod version, and Chromium revision, or pull a pinned remote launcher image.<br>
> Caching avoids each cold runner downloading Chromium.<br>
> Because the official image tag and upstream Dockerfile bases are not digest-pinned, production CI should pin an image digest and explicitly control upgrades rather than rely on latest mutable state.

#### Artifact Integration

> Screenshots, go test JSON/JUnit conversions, CDP logs, downloads, and Compose service logs are ordinary files and can be uploaded with actions/upload-artifact under an always condition.<br>
> Rod does not create a report directory or attachment manifest, so failure hooks, naming, secret filtering, retention, and links between tests and files are custom.<br>
> Integration effort is moderate to high.

#### Sharding And Matrix Support

> GitHub Actions can matrix over shard indexes, configurations, or installed Chromium channels, and go test can select package/test subsets.<br>
> Rod has no native shard calculation, blob report, report merger, or declarative browser project.<br>
> A deterministic test manifest and JUnit/JSON merge step are required.<br>
> Shared PostgreSQL tests should remain in one serial job until each shard receives an independent database and fixture namespace.

#### Container Compatibility

> A host-run Go test can control a host-installed/downloaded browser while the application uses Docker Compose or Testcontainers, which avoids nested container complexity.<br>
> Alternatively, ghcr.io/go-rod/rod exposes a remote launcher on port 7317 and can join the application network.<br>
> The harness must configure reachable hostnames, ports, shared networks, health checks, shm/resources, and a pinned image; localhost inside the browser container is not the runner host.

#### Failure Cleanup

> Browser.Close, page/context cleanup, launcher.Cleanup, and t.Cleanup handle orderly exits; launcher leakless mode is designed to kill browser processes if the Go process crashes.<br>
> TestMain can stop Gin and terminate Testcontainers as the existing integration suite does.<br>
> GitHub Actions must still collect logs and run Compose down with volume/orphan cleanup under always, because cancellation or runner termination can bypass Go cleanup and leak external resources.

### Cost And Risk

#### Open Source Completeness

> Rod, Go testing, Testcontainers, image comparison libraries, axe-core injection, and common JUnit/HTML converters can provide a fully local open-source stack.<br>
> No paid service is required.<br>
> However, Rod alone is not feature-complete as an E2E framework: reporters, retries, visual baselines, accessibility integration, videos/traces, lifecycle, and flake analytics must be composed or built, and Firefox/WebKit coverage cannot be added through Rod.

#### Optional Cloud Dependency

> No cloud is required for authoring, Chromium execution, screenshots, downloads, request interception, or CI.<br>
> A remote browser service, visual review platform, or analytics dashboard is optional.<br>
> Cloud services could add current browsers or managed artifacts, but they do not repair Rod's stale generated CDP API, and native Safari or Firefox would require a different driver rather than merely a Rod dashboard.

#### Migration Cost

> Tests bind to Go, Rod's Must/error APIs, concrete Element handles, CDP-specific events/proto types, and custom lifecycle/reporting helpers.<br>
> Scenario intent, REST fixtures, Testcontainers code, and database cleanup can survive a migration, but browser interactions, waits, locators, interception, emulation, and artifacts would need rewriting for Playwright, WebDriver, or Cypress.<br>
> Low-level proto usage increases lock-in.

#### Security And Supply Chain

> The Go module is MIT-licensed and has a relatively small direct dependency set, but all six direct dependencies are under the primary maintainer's ysmood namespace.<br>
> The August 17, 2026 OpenSSF Scorecard was 3.5/10: no security policy, SAST, fuzzing, signed releases, pinned workflow/container dependencies, or least-privilege workflow tokens were detected.<br>
> Browser setup downloads executable archives from external hosts; the guaranteed default browser is old, creating browser-vulnerability risk.<br>
> Pin Go sums and image digests, scan dependencies/images, control egress, sandbox Chromium, and avoid an obsolete browser baseline.

#### Custom Harness Burden

> High, though lower than a new Go stack because this repository already has TestMain, Testcontainers PostgreSQL/Flyway, Yahoo Finance mocking, SQL cleanup, and log fan-out.<br>
> New work still includes a non-API-only application server or Parcel lifecycle, base URL/readiness, browser/context fixtures, selector and polling helpers for HTMX/Navigo, assertion strategy, retries/repeats, screenshots and browser events on failure, JUnit/HTML output, visual baselines, accessibility audits, shard/database policy, artifact redaction, and cancellation-safe cleanup.

#### Capability Delivery Tier

> Core Rod: Chromium CDP control, launcher/download, headless/headed modes, contexts, pages, selectors, explicit waits, input, navigation, network interception, emulation, screenshots, canvas bytes, downloads, debugging trace/monitoring, and pools.<br>
> Standard Go/testing: discovery, subtests, cleanup, filters, JSON events, parallel controls, HTTP fixture clients, and assertions selected by the project.<br>
> Existing repository infrastructure: PostgreSQL/Flyway Testcontainers, Yahoo mock, fixture state, cleanup SQL, and database logs.<br>
> Community/custom: JUnit/HTML reporting, retries, visual regression, axe integration, lifecycle/readiness, log correlation, redaction, sharding, and flake trends.<br>
> Paid cloud: none required.<br>
> Unavailable in this stack: first-class Firefox, WebKit, and Safari execution.

#### Ai Execution Boundary

> Rod and Go testing are deterministic conventional software and require no LLM, AI credentials, model egress, or per-test inference cost.<br>
> Any optional AI-assisted authoring can remain outside CI and produce reviewed Go tests.<br>
> Browser execution, assertions, fixture setup, and failure reporting retain a complete non-AI path.

### Evidence And Decision

#### Sources

- Rod repository README and source, https://github.com/go-rod/rod, observed through GitHub APIs on 2026-08-22.
- Rod v0.116.2 release and release history, https://github.com/go-rod/rod/releases/tag/v0.116.2 and https://github.com/go-rod/rod/releases, observed 2026-08-22.
- Rod Go package registry entry, https://pkg.go.dev/github.com/go-rod/rod, version, license, API, and 1,021 known importers observed 2026-08-22.
- Official Rod documentation overview and getting started, https://go-rod.github.io/ and https://go-rod.github.io/#/get-started/README, accessed 2026-08-22.
- Official Rod compatibility documentation, https://go-rod.github.io/#/compatibility, accessed 2026-08-22.
- Official Rod custom launch, selectors, network, emulation, and browser/page documentation, https://go-rod.github.io/#/custom-launch, https://go-rod.github.io/#/selectors/README, https://go-rod.github.io/#/network/README, https://go-rod.github.io/#/emulation, and https://go-rod.github.io/#/browsers-pages, accessed 2026-08-22.
- Rod source API reference and examples, https://pkg.go.dev/github.com/go-rod/rod and https://github.com/go-rod/rod/blob/main/examples_test.go, accessed 2026-08-22.
- Rod module, launcher, pinned revision, device profiles, Dockerfile, and workflows, https://github.com/go-rod/rod/blob/main/go.mod, lib/launcher/revision.go, lib/launcher/browser.go, lib/launcher/launcher.go, lib/devices/list.go, lib/docker/Dockerfile, and .github/workflows, inspected 2026-08-22.
- Rod maintenance-mode question and community responses, https://github.com/go-rod/rod/issues/1238, observed 2026-08-22.
- Rod issues, commits, and pull requests, https://github.com/go-rod/rod/issues, https://github.com/go-rod/rod/commits/main, and https://github.com/go-rod/rod/pulls, observed through GitHub APIs on 2026-08-22.
- Rod use-case and project adoption discussion, https://github.com/go-rod/rod/discussions/412, accessed 2026-08-22.
- OpenSSF Scorecard API result for go-rod/rod dated 2026-08-17, https://api.securityscorecards.dev/projects/github.com/go-rod/rod, accessed 2026-08-22.
- Open Asset Allocator local repository evidence: Makefile, test.sh, src/main/go/go.mod, src/main/go/inttest/infra_for_test.go, src/main/go/inttest/infra, and src/main/web-static/package.json, inspected 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence in release age, repository activity, license, stars/forks/issues, registry importers, protocol/browser limits, documented APIs, and local repository integration architecture because these come from first-party source, APIs, and local files.<br>
> Medium confidence in ecosystem breadth, issue trend, future maintenance, and the weighted fit score because adoption metrics are broad proxies and maintainers have not stated project status.<br>
> Low confidence in current Chrome/Go runtime behavior and application-specific flake/resource performance until a spike is executed.

#### Hard Gate Result

> FAIL overall for primary selection.<br>
> PASS: open-source Go authoring, Linux/GitHub Ubuntu x64, black-box Parcel/Gin operation, Chromium headless/headed execution, HTMX-capable explicit waits, Navigo history primitives, realistic forms, canvas/download APIs, network interception, external server access, and reuse of existing Go Testcontainers fixtures.<br>
> CONDITIONAL: robust DOM replacement, retries, reports, visual testing, accessibility, diagnostics, and database-safe parallelism require custom code.<br>
> FAIL: actively maintained/current browser baseline, current guaranteed CDP compatibility, and first-class Firefox/WebKit/Safari coverage.<br>
> The defined empirical project spike is also not complete, but the maintenance and engine failures are independently sufficient for exclusion.

#### Deal Breakers

> The latest release is more than two years old, substantive merges stopped in 2024, and the only documented browser compatibility guarantee is tied to an old pinned Chromium/CDP revision.<br>
> Rod is effectively Chromium-only and cannot supply Firefox, WebKit, or native Safari coverage.<br>
> It also lacks an integrated runner-level retry, reporter, trace/video, visual-regression, accessibility, and shard workflow.<br>
> These are exclusion-level risks for selecting the best-maintained primary E2E tool, even though Go integration and Chromium automation are technically viable.

#### Recommendation

> Excluded as the primary E2E choice.<br>
> Rod is a credible Go-native Chromium automation library and could reuse this repository's Testcontainers, Yahoo mock, SQL cleanup, and testing conventions, but its July 2024 final release, old guaranteed CDP/browser baseline, concentrated maintenance, open dependency work, Chromium-only scope, and missing integrated diagnostics/reporting make it inconsistent with the research goal of a best-maintained and broadly community-adopted E2E tool.<br>
> Retain it only as a Go-specific comparison or for a narrowly Chromium-only internal harness if the project accepts owning a maintained fork.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `roadmap_and_deprecation_risk`
- `dependency_currency`
- `github_metrics`
- `adoption_trend`
- `resource_usage`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="chromedp-with-go-testing"></a>
## 9. chromedp with Go testing

Source result: `chromedp_with_Go_testing.json`

### Project And Compatibility

#### Implementation Language

> chromedp is implemented in Go and tests are authored in Go, normally with the standard testing package.<br>
> Version v0.16.0 declares Go 1.26 and uses generated Go bindings from github.com/chromedp/cdproto.<br>
> No Node.js runtime, WebDriver server, or Java runtime is required for the test process, but an external Chrome, Chromium, Edge, or compatible headless-shell executable is required at runtime.

#### Node And Typescript Compatibility

> Node.js 24 and TypeScript 6 are not chromedp runtimes.<br>
> This repository can continue to use Node 24.12, TypeScript 6.0.3, ESNext modules, and Parcel to build or serve the frontend while Go tests drive the resulting HTTP application as a black box.<br>
> The browser tests therefore avoid npm and TypeScript compatibility coupling, but they also cannot reuse TypeScript test utilities or type-safe frontend models without a language boundary.

#### Operating System Support

> The Go package and Chrome-family browsers can run on Linux, macOS, and Windows.<br>
> Local Linux and GitHub-hosted Ubuntu x86-64 are practical targets, and the project's own workflow uses ubuntu-latest.<br>
> The official chromedp/headless-shell image is multi-architecture.<br>
> Actual availability depends on installing a compatible Chrome-family binary and its system libraries; native Safari and Firefox are outside the protocol and cannot be added by changing the operating system.

#### License And Governance

> MIT License, copyright Kenneth Shaw.<br>
> The code is maintained publicly under the chromedp GitHub organization and accepts issues and contributions.<br>
> No commercial license or hosted service is required.<br>
> No published foundation governance, maintainer charter, funding model, or succession policy was found in the reviewed repository, and recent release work is highly concentrated in Kenneth Shaw.

#### Installation Model

> Add github.com/chromedp/chromedp v0.16.0 to the Go module and run tests with Go 1.26 or newer.<br>
> Install or supply Chrome/Chromium separately, select an exact executable with chromedp.ExecPath when determinism matters, or use the multi-arch chromedp/headless-shell container.<br>
> The official container publishes daily stable, beta, and dev tags plus exact browser-version tags. chromedp discovers common local Chrome paths but does not download, patch, or pin a browser for the project.

#### Candidate Scope And Layer

> Browser automation stack: chromedp is a high-level Chrome DevTools Protocol client and process allocator, while Go testing supplies test discovery, assertions through project-selected libraries, filtering, and basic output.<br>
> It is not a complete E2E framework and does not integrate browser matrices, retries, fixtures, application lifecycle, visual baselines, reports, videos, or test-oriented traces.

#### Authoring And Async Model

> Tests use ordinary Go functions, testing.T, contexts, defer or t.Cleanup, and ordered chromedp.Action values executed by chromedp.Run.<br>
> Context cancellation supplies deadlines and browser cleanup.<br>
> Browser and target events are delivered through synchronous ListenBrowser or ListenTarget callbacks, where blocking work must be moved to another goroutine and CDP commands must not be called directly from the callback.<br>
> There is no hidden command queue, Gherkin syntax, assertion DSL, or native async/await.

#### Build Pipeline Coupling

> The tests can navigate to the existing Parcel development server and its /api proxy or to the production Parcel output served by Gin. chromedp does not require Vite, source transformation, code coverage instrumentation, injected framework scripts, or changes to the frontend module graph.<br>
> Builds and servers must be started by Make, shell scripts, Docker Compose, TestMain, testcontainers, or CI steps outside chromedp.

#### Testability Instrumentation Required

> No production browser instrumentation, relaxed content-security policy, or framework-specific build hook is required.<br>
> Reliable tests will benefit from stable IDs or data-test attributes because core locators are CSS, XPath/search, JavaScript paths, or node IDs rather than accessible-role locators.<br>
> Deterministic database reset, fixture APIs, a Yahoo Finance substitute, and readiness signals are repository responsibilities; these may require test-only infrastructure but are not chromedp runtime requirements.

### Maintenance Health

#### Latest Stable Release

> v0.16.0, the latest non-prerelease Go module tag, published July 14, 2026.<br>
> The Go proxy and pkg.go.dev identify it as latest.<br>
> GitHub's latest Release object still reports v0.15.1 from April 1, 2026, so release-page metadata lags the module tag.

#### Release Cadence

> Irregular but current.<br>
> Observed tags were v0.14.0 on July 26, 2025, v0.14.1 on August 5, v0.14.2 on October 7, v0.15.0 on March 21, 2026, v0.15.1 on March 23, and v0.16.0 on July 14.<br>
> Recent versions mainly update cdproto and dependencies or make small compatibility fixes; v0.16.0 contains two commits beyond v0.15.1 and has no GitHub Release notes.

#### Repository Activity

> Current but low-volume and maintainer-driven.<br>
> The default branch was last pushed July 14, 2026.<br>
> The ten newest commits observed span June 2025 through July 2026; the six newest were authored by Kenneth Shaw, and the two v0.16.0 commits update cdproto/Go code and run modernization.<br>
> A targeted GitHub search found no merged pull requests from August 22, 2025 through August 22, 2026, while several contributor pull requests remain open.

#### Dependency Currency

> v0.16.0 requires Go 1.26, exactly matching this repository's Go 1.26.5 baseline.<br>
> Its cdproto pseudo-version was generated on the same day as the release, and its go.mod contains five direct and four indirect modules.<br>
> The release therefore aligns with current Go and CDP code, but browser compatibility remains tied to continued cdproto refreshes and Chrome's evolving protocol.<br>
> Node.js and TypeScript versions do not affect the Go test runtime.

### Community Adoption

#### Npm Downloads

> Not applicable. chromedp is distributed as the Go module github.com/chromedp/chromedp and has no official npm package, so an npm download count would be misleading and is recorded as N/A.

#### Github Metrics

> Observed August 22, 2026: chromedp/chromedp had 13,263 stars, 885 forks, 167 open issues, 177 combined open issues/pull requests, and 45 named entries from GitHub's contributor endpoint; the default branch was last pushed July 14, 2026. pkg.go.dev reported 2,179 known importing packages for v0.16.0.<br>
> GitHub repository metrics cover scraping, PDF generation, profiling, and general automation as well as E2E testing.

#### Ecosystem Usage

> The project provides generated cdproto access, a maintained examples repository, device descriptors, keyboard helpers, a CDP logging proxy, and a multi-arch headless-shell image. pkg.go.dev's 2,179 known importers and Debian packaging provide concrete ecosystem evidence.<br>
> Third-party projects build higher-level test frameworks and automation products on chromedp, but there is no large official plugin ecosystem for reporters, visual testing, fixtures, or accessibility comparable to integrated E2E frameworks.

#### Community Support

> The official README, pkg.go.dev API reference, self-contained package examples, separate examples repository, GitHub issue tracker, and Chrome CDP documentation are available.<br>
> Documentation is strongest at API and protocol-example level and weak on production E2E architecture, reliable synchronization, report assembly, and upgrade troubleshooting.<br>
> GitHub Discussions is disabled, and the issue tracker contains unanswered reports, including the Go test compatibility report with no comments since September 2025.

#### Adoption Metric Normalization

> There is no official Go module download count equivalent to the npm seven-day metric.<br>
> The recorded pkg.go.dev Imported By value is 2,179 known public importing packages at observation time; it excludes private modules and can include indirect libraries, non-E2E scraping, PDF, profiling, and automation use.<br>
> GitHub stars, forks, and contributors describe the full chromedp project rather than teams running browser tests, and they are not unique active users.

### Browser And Runtime Coverage

#### Browser Engines

> Chromium/Blink only. chromedp drives browsers that expose the Chrome DevTools Protocol, including Google Chrome, Chromium, Chrome headless-shell, and potentially Chromium-based Edge when an executable or remote endpoint is supplied.<br>
> It cannot drive Firefox/Gecko, WebKit, or native Safari, and it has no standards-based fallback.<br>
> This is a confirmed cross-browser coverage failure compared with complete three-engine E2E frameworks.

#### Browser Protocol

> Direct Chrome DevTools Protocol over WebSocket, implemented through chromedp and generated cdproto domain bindings.<br>
> This gives low-level access to DOM, Runtime, Page, Network, Fetch, Emulation, Accessibility, Tracing, Storage, and other Chrome domains without WebDriver.<br>
> Chrome's official protocol documentation warns that tip-of-tree CDP changes frequently and has no backward-compatibility guarantee, creating Chromium-specific version coupling.

#### Headless And Headed Modes

> Headless is the default and is suitable for Linux CI.<br>
> Headed execution is enabled by overriding the allocator's headless flag and is useful for local observation.<br>
> The official headless-shell image is intentionally headless; headed Linux containers require display/Xvfb configuration that chromedp does not provide.

#### Browser Version Management

> chromedp locates an installed Chrome-family executable or accepts an explicit ExecPath/remote endpoint; it does not download or pin browsers.<br>
> Reproducible CI therefore requires the project to install a selected Chrome for Testing build, pin an exact chromedp/headless-shell version tag or digest, and coordinate upgrades with chromedp/cdproto.<br>
> Floating image channels are updated daily and trade reproducibility for currency.<br>
> Browser and Go module caches are managed separately.

#### Parallel Browser Support

> Go tests can create independent browser processes, multiple targets in one process, or explicit isolated BrowserContexts, and t.Parallel or CI matrices can schedule them. chromedp has no browser-project abstraction, cross-engine matrix, worker scheduler, deterministic sharder, or resource-aware concurrency policy.<br>
> The harness must allocate browsers, contexts, downloads, ports, artifacts, and database state for each parallel test or worker.

#### Mobile Emulation

> The device package and Emulate APIs configure viewport, screen orientation, device scale, mobile mode, touch, and user agent for Chrome.<br>
> Raw Emulation domain calls expose additional Chromium controls.<br>
> This is desktop Chromium emulation, not a real Android Chrome or iOS Safari device, and it provides no Firefox/WebKit mobile coverage.

#### Real Browser Fidelity

> Local execution can drive an installed, vendor-built Google Chrome or Chromium binary, so it exercises a real Blink browser rather than a DOM simulator.<br>
> The official headless-shell is a reduced Chrome build and its image modifies the reported user agent, so it is not identical to full headed Chrome.<br>
> Chromium-based Edge may be driven through its executable but is not a separately managed project.<br>
> No WebKit result can be interpreted as Safari coverage because no WebKit target exists.

#### Environment Determinism Controls

> Core options set executable path, user-data directory, window size, user agent, proxy, environment, and Chrome flags. chromedp exposes viewport/device emulation, and raw cdproto can set timezone, locale, geolocation, permissions, virtual time, media features, and network conditions where current Chrome supports them.<br>
> These are low-level protocol calls rather than test fixtures.<br>
> There is no integrated seeded randomness, automatic animation disabling, reduced-motion project, or reset policy; the harness must apply and restore controls consistently.

### Application Fit

#### Dynamic Dom Synchronization

> chromedp query actions poll until their configured node condition succeeds, and WaitVisible, WaitReady, WaitEnabled, WaitNotPresent, Poll, mutation polling, and context deadlines can synchronize with HTMX and asynchronous Handlebars work.<br>
> A selector action queries the current DOM, which avoids retaining a stale node when tests use selectors for each step.<br>
> It does not provide locator-wide actionability or assertion auto-retry, and WaitReady(body) does not imply that later HTMX content is complete.<br>
> Tests must wait for a domain-specific post-swap element, value, event, or response after every asynchronous transition.

#### Routing Support

> Navigate, Location, NavigateBack, NavigateForward, navigation entries, raw Page lifecycle events, and JavaScript evaluation cover direct Navigo deep links and history changes.<br>
> PushState routing may not produce a full navigation event, so tests need explicit URL plus route-content polling.<br>
> Gin must continue serving the SPA shell for deep links, and each test must reset route, targets, and browser state through custom setup.

#### Locator Model

> Selectors support CSS queries, XPath/text-capable DOM search, JavaScript paths, node IDs, custom query functions, subtree roots, and minimum match counts.<br>
> There are no first-class getByRole, getByLabel, test-ID, visible-text, or ARIA snapshot locators, and matching is not strict by default; an action can use the first of multiple matches.<br>
> Stable IDs or carefully maintained CSS selectors are therefore important.<br>
> Each selector action can reacquire a replacement node, but retained cdp.Node values can become stale after HTMX swaps.

#### Form Interaction

> Click and DoubleClick dispatch mouse input at element coordinates; SendKeys and keyboard helpers dispatch key events; Focus, Blur, Clear, SetValue, SetUploadFiles, and raw Input/DOM calls cover forms.<br>
> Tests can inspect values, hidden synchronized fields, validity, and validation messages through attributes, properties, or Evaluate.<br>
> Realistic focus/blur and native validation require using pointer/keyboard actions rather than direct SetValue or JavaScript mutation, and select, checkbox, dynamic-row, occlusion, and event-order behavior need explicit helpers and assertions.

#### Canvas And Download Support

> Chart.js can be checked by evaluating its application data/state, reading accessible fallback content, or capturing element/full-page screenshots. chromedp has no visual baseline manager or public Chart.js assertion.<br>
> Downloads can be enabled with browser.SetDownloadBehavior, observed with browser.EventDownloadProgress, written to a test directory, and parsed by Go.<br>
> File naming, completion channels, payload assertions, cleanup, and protection from cross-test collisions are custom harness work.

#### Network And Api Access

> Go's net/http client can seed and reset REST fixtures outside the browser.<br>
> ListenTarget exposes Network and Fetch events, and raw cdproto can inspect requests/responses, continue, fail, fulfill, or modify intercepted requests in Chromium.<br>
> This can replace browser-originated Yahoo Finance traffic; server-side Gin calls still require the existing Go HTTP mock/service substitution because browser interception cannot see backend egress.<br>
> There is no integrated API fixture, route registry, HAR replay, or cross-browser mock layer.

#### Same Origin Support

> Tests operate outside the page and can navigate to the consolidated Gin production origin or Parcel port 8000, where the existing proxy forwards /api to Gin.<br>
> No CORS relaxation is required.<br>
> When the browser runs in a separate container, localhost names that browser container, so the base URL must use a shared Compose service name, published host address, or host-gateway mapping.

#### Test Isolation

> A new browser process with a temporary user-data directory gives strong isolation but has the highest startup cost.<br>
> Multiple child contexts normally create targets in the same browser profile and can share cookies/storage.<br>
> WithNewBrowserContext creates an isolated incognito BrowserContext that is disposed on cancellation, but it requires an already initialized browser and explicit harness setup.<br>
> PostgreSQL remains outside browser isolation and needs unique fixtures, reset APIs, per-worker databases/schemas, or serialized tests.

#### External Server Model

> chromedp naturally targets any independently started HTTP application by navigating to a configurable URL.<br>
> The same Go test can use Parcel development, a locally published Compose port, a production image, or a remote environment.<br>
> The library has no built-in baseURL fixture, server command, port allocator, health probe, or reuse-existing-server policy.

#### Application Lifecycle

> Go TestMain, t.Cleanup, os/exec, the repository's Make/scripts, or existing testcontainers dependencies can build the frontend, start Docker Compose, wait for PostgreSQL/Flyway/Gin readiness, seed data, collect logs, and tear down services. chromedp only manages the browser allocator and contexts.<br>
> The project must implement cancellation-safe readiness and teardown across Compose, Flyway, Gin, browser processes, download directories, and database volumes.

#### Visual Regression Workflow

> chromedp captures viewport, full-page, and element PNGs, but it provides no test-facing baseline creation/update command, masking, tolerance policy, per-browser naming, CI diff report, or review UI.<br>
> A Go image comparison library and project conventions can supply these features.<br>
> Fonts, browser version, viewport, scale, data, time, and Chart.js animation must be stabilized manually; semantic chart-data assertions should remain primary because only Chromium screenshots are available.

#### Accessibility Audit Integration

> Raw CDP Accessibility domain calls can inspect Chrome's accessibility tree, and tests can inject an axe-core script then evaluate it, but chromedp has no official axe integration, ARIA snapshot matcher, violation reporter, or cross-engine accessibility coverage.<br>
> The repository would need to vendor or obtain an audited axe asset, execute it under the page's security policy, deserialize results into Go, attach reports, and complement automation with manual accessibility testing.

### Reliability

#### Waiting Model

> Navigation waits for the relevant page load behavior, selector actions poll for node conditions, Poll/PollFunction can use interval, mutation, or requestAnimationFrame polling, and context deadlines bound operations.<br>
> Click defaults to waiting for a visible node, but there is no complete actionability model covering stability, enabled state, event reception, occlusion, and retry after detachment, and ordinary Go assertions do not poll.<br>
> Reliable HTMX tests need reusable domain waits and must prohibit fixed Sleep except for diagnostics or bounded animation cases.

#### Flake Controls

> Go contexts provide suite/action deadlines, go test provides a package timeout and -count repetition, and tests can isolate browser processes or contexts.<br>
> Standard Go testing has no retry policy, flaky-pass classification, last-failed mode, automatic action/assertion retry, or trace-on-retry.<br>
> Retries require a third-party runner or custom loop and can hide shared PostgreSQL defects unless first-failure evidence and cleanup are preserved.<br>
> Upstream v0.16.0 also has open reports for a cancellation race and go test allocator cancellation.

#### Isolation Model

> The safest model is one isolated BrowserContext or browser process per test plus unique/reset PostgreSQL fixtures.<br>
> Sharing one browser process can reduce cost, but default child targets share profile state unless WithNewBrowserContext is used.<br>
> Context cancellation disposes explicit browser contexts and targets; tests must still remove downloads, reset permissions/network hooks, and clean database state.<br>
> Serial scenarios can share domain state deliberately but require package-level coordination and reliable cleanup.

#### Parallelism Controls

> go test controls package concurrency with -p, test-level concurrency with t.Parallel, and CPU scheduling with GOMAXPROCS.<br>
> The harness can use semaphores for browser slots and choose shared browsers or per-test processes.<br>
> There is no chromedp worker pool, serial group, shard coordinator, or database lock.<br>
> Start state-mutating E2E tests serially, then allocate unique database/schema and artifact namespaces before enabling package, test, or CI parallelism.

#### Flake Observability

> go test -json identifies test attempts only when the harness models them, and custom listeners can capture evidence for each attempt.<br>
> Neither chromedp nor standard testing classifies flaky retries, stores history, manages quarantine, or produces a local trend dashboard.<br>
> The project must preserve first-attempt failures, label retries, emit JSON/JUnit through additional tooling, and aggregate trends outside the browser library.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Core primitives can capture page or element screenshots, full-page screenshots, DOM/HTML dumps, browser process output, console events, JavaScript exceptions, network events, response metadata/bodies, and optional raw CDP tracing data.<br>
> They are not automatically collected on test failure or assembled into a navigable action trace.<br>
> Browser video remains an open feature request.<br>
> File naming, retention modes, HAR-like serialization, attachment, and backend/Compose logs are custom.

#### Debugging Tools

> Developers can run headed Chrome, use Chrome DevTools/CDP, enable chromedp debug/log/error callbacks, inspect browser combined output, pause with a debugger or breakpoint, and debug Go with Delve or an IDE.<br>
> There is no test inspector, locator picker, step mode, time-travel DOM viewer, automatic pause-on-failure, or official editor extension.<br>
> Remote headless debugging requires exposing and securing the CDP endpoint.

#### Test Generation

> No recorder, code generator, or locator generator is provided for tests.<br>
> Chrome DevTools Protocol Monitor can help discover raw commands, but test scenarios, selectors, waits, fixtures, and assertions must be authored and reviewed manually.

#### Reporters

> Standard go test provides text output and -json event streams. go tool test2json is available, and community tools such as gotestsum can produce JUnit, but chromedp has no HTML, JUnit, JSON, GitHub annotation, blob, shard-merging, screenshot-attachment, or trace reporter.<br>
> A project-owned reporter/helper layer must connect browser artifacts to testing.T names and CI results.

#### Documentation Quality

> The API reference is complete enough to enumerate actions and low-level extension points, and official examples cover navigation, forms, screenshots, downloads, remote browsers, devices, tabs, and events.<br>
> The README is concise but its 'without external dependencies' wording applies to the Go CDP client, not the required browser runtime.<br>
> Guidance is limited for E2E architecture, auto-wait alternatives, isolation, CI artifacts, security, browser pinning, and current known failures; users often need cdproto and Chrome protocol documentation.

#### Local Workflow

> Developers can run one package or test with go test ./path -run TestName -v, repeat with -count, set Go timeouts, and switch to headed mode through allocator options.<br>
> This aligns with the repository's existing Go integration tests and IDE tooling.<br>
> Browser installation, base URL, application startup, artifact directories, and a convenient debug profile must be scripted.<br>
> The unresolved Go 1.25+ allocator report means execution under this repository's Go 1.26.5 needs an explicit local spike.

#### Failure Log Correlation

> ListenTarget can timestamp console, exception, and network events with a test/run identifier, while chromedp logging and CombinedOutput capture client/browser diagnostics.<br>
> Gin, PostgreSQL, Flyway, Docker Compose, and Yahoo mock logs are separate streams.<br>
> The harness must use UTC clocks, correlation IDs, per-test boundaries, and always-run log collection to relate browser events to backend failures; chromedp supplies no cross-process correlation.

#### Artifact Data Exposure

> Screenshots, DOM dumps, console messages, response bodies, headers, downloads, traces, browser profiles, cookies, local storage, and database/service logs can expose credentials and portfolio data.<br>
> Raw CDP interception can observe authorization and cookie headers. chromedp has no artifact redaction pipeline.<br>
> Use synthetic fixtures, isolated temporary profiles, selective body capture, header redaction, short artifact retention, restricted GitHub access, and cleanup that never uploads user-data directories or unreviewed network payloads.

### Github Actions Fit

#### Official Ci Support

> The repository maintains a GitHub Actions workflow using actions/setup-go, ubuntu-latest, go test, and a headless-shell Docker test, and it publishes the chromedp/headless-shell image.<br>
> There is no dedicated chromedp GitHub Action or complete application-test workflow.<br>
> Materially, all 16 retained workflow runs listed by the GitHub Actions API from July 26, 2025 through July 14, 2026 concluded failure; the v0.16.0 stable-Go job timed out after 600 seconds in TestPoll.<br>
> This weakens confidence in the current official CI signal.

#### Browser Caching

> actions/setup-go can cache Go modules/build output.<br>
> Chrome is external: use the runner's installed browser with drift risk, cache a downloaded Chrome for Testing artifact under an OS/architecture/version key, or pin the official headless-shell image tag/digest.<br>
> The library provides no browser cache manager.<br>
> Cache restoration must not mix incompatible Chrome, chromedp, and cdproto versions.

#### Artifact Integration

> Tests can write screenshots, DOM, JSON events, downloads, and custom logs to a known directory and upload them with actions/upload-artifact under an always condition.<br>
> No report or manifest is produced automatically, so the harness must map files to test names, include first-failure evidence, collect Compose logs before teardown, enforce retention/access, and avoid uploading secrets.

#### Sharding And Matrix Support

> GitHub Actions can matrix over a pinned Chrome version, headless/headed topology, or custom shard index, while go test can split by package or -run pattern. chromedp cannot create Firefox/WebKit matrix entries, deterministically allocate tests to cross-job shards, merge reports, or coordinate retries.<br>
> Shard manifests, JUnit/JSON merging, database allocation, and duplicate-safe retries are custom work.

#### Failure Cleanup

> Cancel chromedp target, browser-context, browser, and allocator contexts and register cleanup immediately.<br>
> On Linux, the local allocator uses parent-death handling to reduce leaked Chrome children; remote browsers remain separately managed, and the official image recommends --init to reap zombies.<br>
> CI must still collect logs and run docker compose down --volumes --remove-orphans under an always condition because Go cleanup may not execute after hard cancellation or runner termination.

### Cost And Risk

#### Open Source Completeness

> The MIT-licensed library, generated CDP bindings, Go testing, screenshots, events, emulation, and official headless image are sufficient for open-source Chromium automation without a paid service.<br>
> They are not a complete open-source E2E product: runner features, reports, visual baselines, accessibility reporting, retries, application lifecycle, and flake history require additional open-source dependencies or custom code, and Firefox/WebKit coverage cannot be delivered by chromedp at all.

#### Optional Cloud Dependency

> No cloud is required for local or GitHub Actions Chromium execution.<br>
> Optional remote-CDP/browser services can outsource browser capacity, and optional visual/reporting services can reduce custom work, but they are not required and do not change chromedp's protocol limits unless tests are rewritten for another client.<br>
> Repository-owned files and GitHub artifacts can replace hosted dashboards at higher harness cost.

#### Migration Cost

> High API and harness lock-in despite ordinary Go syntax.<br>
> Tests bind to chromedp actions, context lifecycle, cdproto events/domains, CSS/XPath selectors, custom polling, and project-specific fixtures/reporting.<br>
> Domain setup and HTTP clients can be preserved, but moving to Playwright, WebDriver, or another language requires rewriting browser interactions, waits, isolation, artifacts, and matrices.<br>
> Direct CDP calls increase the migration surface.

#### Security And Supply Chain

> v0.16.0 is MIT-licensed, available through the checksum-backed Go module ecosystem, and deps.dev reported no known advisories for the module at observation time.<br>
> Its go.mod has five direct and four indirect modules, including a generated cdproto pseudo-version. deps.dev reported no SLSA provenance or attestations, and the repository root has no SECURITY.md.<br>
> Chrome/headless-shell is a large separately updated executable or image with external download and browser-sandbox concerns.<br>
> Pin module checksums and image/browser versions or digests, scan both Go and container dependencies, restrict CDP ports, avoid --no-sandbox where possible, and update quickly for Chrome security fixes.

#### Custom Harness Burden

> Very high.<br>
> The repository must implement browser discovery/pinning, allocator/context fixtures, accessible locator conventions, HTMX/Navigo waits, assertions, REST/database fixtures, Yahoo Finance replacement, Compose/Flyway/Gin readiness, PostgreSQL reset and worker allocation, retries and first-failure tracking, screenshots/DOM/network/console capture, visual baselines, axe injection/reporting, JSON/JUnit/HTML output, artifact redaction, sharding, and cancellation-safe cleanup.<br>
> Existing Go integration infrastructure helps with HTTP mocking, testcontainers, and database setup but does not supply browser-test ergonomics.

#### Capability Delivery Tier

- **Core:** Chrome-family process allocation or remote attachment, CDP actions/domains, contexts/targets, CSS/XPath/search queries, explicit waits and polling, mouse/keyboard/form actions, navigation/history, JavaScript evaluation, screenshots, downloads, network events/interception, emulation, and raw accessibility/tracing access.
- **Standard Go Testing:** Test discovery, subtests, basic assertions through testing.T failures, filters, repetition, package/test parallelism, timeouts, cleanup registration, and JSON event output.
- **Official Adjacent:** Generated cdproto bindings, examples repository, device and keyboard packages, headless-shell container, and chromedp-proxy.
- **Community Or Custom:** JUnit/HTML reports, assertion polling, retries, visual diffing, axe injection, application lifecycle, API/database fixtures, backend mocking, artifact bundles/redaction, sharding, and flake history.
- **Unavailable In Stack:** Firefox, WebKit, and native Safari automation through chromedp.
- **Optional Paid Cloud:** Remote Chromium capacity, visual review, and analytics only; none is mandatory.

#### Ai Execution Boundary

> No AI or LLM is required to author, execute, assert, debug, or report Go/chromedp tests.<br>
> CI can remain deterministic with no model credentials, AI egress, token cost, or runtime model decisions.<br>
> Optional AI authoring should remain outside required execution, use synthetic data, and produce reviewed Go code that runs without AI.

### Evidence And Decision

#### Sources

- Title: chromedp README and Go package reference | Url: https://pkg.go.dev/github.com/chromedp/chromedp | Evidence: v0.16.0, Go 1.26, MIT license, 2,179 known importers, installation, actions, queries, polling, contexts, allocators, emulation, screenshots, events, and official examples.
- Title: Go module proxy metadata for chromedp | Url: https://proxy.golang.org/github.com/chromedp/chromedp/@latest | Evidence: Latest module v0.16.0 and publication time 2026-07-14T21:56:16Z.
- Title: chromedp v0.16.0 go.mod | Url: https://github.com/chromedp/chromedp/blob/v0.16.0/go.mod | Evidence: Go 1.26 requirement, same-day cdproto pseudo-version, and direct/indirect dependency set.
- Title: chromedp GitHub repository, tags, commits, pull requests, issues, and contributors | Url: https://github.com/chromedp/chromedp | Evidence: 13,263 stars, 885 forks, 167 open issues, 45 named contributors, release cadence, sparse recent commits, no merged pull requests in the sampled year, maintainer concentration, and current issue evidence observed through GitHub APIs on 2026-08-22.
- Title: chromedp v0.15.1 GitHub Release and v0.16.0 tag | Url: https://github.com/chromedp/chromedp/releases | Evidence: GitHub Release metadata lag: v0.15.1 is the latest Release object while Go module/tag v0.16.0 is newer.
- Title: chromedp official GitHub Actions workflow and runs | Url: https://github.com/chromedp/chromedp/actions/workflows/test.yml | Evidence: ubuntu-latest Go test and headless-shell workflow; all 16 retained runs failed, including v0.16.0 timing out after 600 seconds in TestPoll.
- Title: Open chromedp Go test allocator issue 1591 | Url: https://github.com/chromedp/chromedp/issues/1591 | Evidence: Unanswered report of context cancellation under go test with Go 1.25.1 on Ubuntu and Debian while equivalent go run succeeds.
- Title: Open chromedp v0.16.0 Context.Target race issue 1638 | Url: https://github.com/chromedp/chromedp/issues/1638 | Evidence: Reproduced race during target attachment and browser cancellation under Go 1.26.5 on Windows.
- Title: Open chromedp RemoteAllocator issue 1601 | Url: https://github.com/chromedp/chromedp/issues/1601 | Evidence: Unresolved report that v0.14.2 broke the official remote allocator example.
- Title: Official chromedp headless-shell documentation | Url: https://github.com/chromedp/docker-headless-shell | Evidence: Multi-arch image, daily stable/beta/dev and exact-version tags, remote CDP usage, shared-memory recommendation, non-root example, and init/zombie guidance.
- Title: Chrome DevTools Protocol documentation | Url: https://chromedevtools.github.io/devtools-protocol/ | Evidence: Chromium/Blink scope, domain capabilities, WebSocket transport, tip-of-tree change frequency, and absence of backward-compatibility guarantees.
- Title: Open Source Insights metadata for chromedp v0.16.0 | Url: https://deps.dev/go/github.com%2Fchromedp%2Fchromedp/v0.16.0 | Evidence: MIT license, no recorded advisories, no deprecation, and no SLSA provenance or attestations at observation time.
- Title: Official chromedp examples repository | Url: https://github.com/chromedp/examples | Evidence: Examples for forms, downloads, device emulation, remote execution, screenshots, uploads, and selectors, plus explicit warning that examples/selectors may break.
- Title: Independent chromedp controlled test review | Url: https://thunderbit.com/blog/chromedp-review | Evidence: Secondary evidence for the required external browser, domain-specific waits, release-object mismatch, and reproduced go test issue; quantitative macOS measurements were not generalized to this project.
- Title: Open Asset Allocator local repository evidence | Url: src/main/go/go.mod and src/main/web-static/package.json | Evidence: Go 1.26.5, existing testing/testcontainers/httpmock stack, Node 24.12, TypeScript 6.0.3, Parcel, HTMX, Navigo, Handlebars, Chart.js, proxy configuration, and PostgreSQL/Flyway Compose dependencies inspected 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

- **High:** Language/runtime requirements, module version, license, Chromium-only protocol scope, installation model, core API capabilities, GitHub snapshot, official workflow failures, and missing integrated framework features are supported by primary source/API evidence.
- **Moderate:** Application-specific fit, issue responsiveness, long-term maintenance risk, isolation architecture, security process, and custom harness estimates combine primary capabilities with engineering analysis.
- **Low:** Adoption direction, resource performance in this repository, and complete end-to-end reliability remain unmeasured.

#### Hard Gate Result

- **Overall:** FAIL for selection as the primary E2E framework.
- **Go Ecosystem Alignment:** PASS: v0.16.0 requires Go 1.26 and this repository uses Go 1.26.5 with existing integration-test infrastructure.
- **Node 24 And Typescript 6:** PASS BY SEPARATION: chromedp tests the built/served application and does not execute through Node or TypeScript.
- **Linux And Github Ubuntu:** CAPABILITY PASS, HEALTH CONDITION: local Chrome/headless-shell and ubuntu-latest are supported patterns, but all retained upstream CI runs currently fail and the project-specific run was not executed.
- **Cross Browser Coverage:** FAIL: Chromium/Blink only; Firefox, WebKit, and native Safari cannot be driven.
- **Parcel Black Box Build:** PASS: URL-level automation requires no Parcel replacement or frontend instrumentation.
- **Htmx Handlebars Dynamic Dom:** CONDITIONAL PASS: explicit domain polling and node reacquisition helpers are required; there is no comprehensive auto-waiting.
- **Navigo Deep Links And History:** CONDITIONAL PASS: navigation/history primitives exist, but pushState readiness conditions are custom.
- **Forms Blur Focus And Validation:** PASS WITH CUSTOM HELPERS through mouse, keyboard, Focus, Blur, DOM inspection, and raw CDP.
- **Chartjs Canvas:** PASS WITH CUSTOM semantic evaluation and screenshot comparison; no visual workflow is integrated.
- **Downloads:** PASS WITH CUSTOM download behavior, event handling, file parsing, naming, and cleanup.
- **Api Fixtures And Yahoo Replacement:** PASS WITH Go HTTP helpers and browser Fetch interception or backend-side mock substitution.
- **Postgresql Isolation:** PASS WITH PROJECT-OWNED reset/allocation and serial execution; browser contexts do not isolate the database.
- **External Server And Docker Compose:** PASS WITH CUSTOM startup, readiness, networking, and teardown.
- **Failure Diagnostics:** FAIL AS AN INTEGRATED REQUIREMENT: primitives exist for screenshots and logs, but automatic trace/video/report/artifact correlation is absent and must be built.
- **Open Source Local Ci:** PASS for Chromium automation without paid services.

#### Deal Breakers

> Confirmed exclusion factors are Chromium-only coverage with no Firefox/WebKit/Safari path, absence of an integrated E2E runner and diagnostic artifact workflow, very high custom synchronization/lifecycle/reporting burden, and a currently failing upstream CI history.<br>
> Additional serious risks are a v0.16.0 cancellation race report, an unanswered Go test allocator report relevant to this repository's Go 1.26.5 runtime, and unresolved RemoteAllocator behavior.<br>
> These do not prevent a focused Chromium smoke suite or low-level debugging tool, but they prevent selecting this stack as the best broadly maintained E2E solution.

#### Recommendation

> Excluded as the primary E2E testing choice.<br>
> Retain chromedp only as a specialized Go-native Chromium option when direct CDP access or reuse of Go integration fixtures is more important than cross-browser coverage and integrated diagnostics.<br>
> For the project-wide E2E stack, prefer a maintained complete framework with Chromium, Firefox, and WebKit projects, locator/action auto-waiting, isolated contexts, application lifecycle, retries, reports, traces, screenshots/video, visual baselines, and first-class CI sharding.

### Uncertain Fields

- `issue_health`
- `roadmap_and_deprecation_risk`
- `wrapper_upstream_lag`
- `maintainer_concentration`
- `adoption_trend`
- `resource_usage`
- `container_compatibility`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="codeceptjs"></a>
## 10. CodeceptJS

Source result: `CodeceptJS.json`

### Project And Compatibility

#### Implementation Language

> CodeceptJS 4 is implemented primarily in JavaScript and runs on Node.js.<br>
> It ships TypeScript declarations, supports JavaScript and TypeScript test authoring, and delegates browser control to separately installed Playwright, WebdriverIO, Puppeteer, or Appium packages.

#### Operating System Support

> Local Linux and GitHub-hosted Ubuntu are first-class targets in the official CI guide.<br>
> The Playwright helper supports Chromium, Firefox, and WebKit on Linux; the WebDriver helper can use installed or downloaded Chrome, Chromium, and Firefox.<br>
> The official Docker image is built for linux/amd64 and linux/arm64 as of July 2026.<br>
> Native Safari remains limited to Apple platforms or a remote browser provider.

#### License And Governance

> The framework and core repository use the MIT License, permitting repository and CI use without a commercial license.<br>
> The public codeceptjs GitHub organization owns the repository; Michael Bodnarchuk (DavertMik) is the creator and dominant historical contributor, with four npm maintainers listed for 4.1.0.<br>
> Testomat.io products and reporting are promoted and share contributor relationships, but local execution does not require Testomat.io or another paid service.<br>
> No foundation charter or formal public governance and succession policy was found.

#### Installation Model

> For the preferred backend, install codeceptjs, playwright, and optionally tsx as development dependencies, initialize configuration, then run npx playwright install --with-deps to download browsers and Linux libraries.<br>
> CodeceptJS itself does not bundle the Playwright package.<br>
> WebDriver uses a separately installed WebdriverIO backend and can automatically obtain supported browsers and drivers; Puppeteer installs its own browser according to Puppeteer's package policy.<br>
> The official codeceptjs/codeceptjs image is an optional prebuilt environment.<br>
> Lock the npm graph and browser versions for CI reproducibility.

#### Candidate Scope And Layer

> E2E meta-framework and scenario-oriented abstraction layer.<br>
> It includes a Mocha-based runner, actor DSL, assertions, hooks, dependency injection, workers, sharding, retries, plugins, reports, and diagnostics while delegating browser automation to Playwright, WebDriver, or Puppeteer.

#### Authoring And Async Model

> Tests use Feature and Scenario declarations and an injected I actor with linear BDD-style commands; Gherkin is also supported.<br>
> CodeceptJS normally schedules I.* commands through its recorder queue, so action steps can be written without await while grabbers and returned values require await.<br>
> The fullPromiseBased option generates typings that require await for every I command.<br>
> Native async functions, hooks, effects, helper code, and direct Playwright calls are supported, but mixing queued commands and native promises requires understanding the recorder model.

#### Build Pipeline Coupling

> Black-box scenarios can target the existing Parcel server on port 8000 or the production Gin-served build through the helper's configurable base URL.<br>
> CodeceptJS does not require Vite, source transforms, application instrumentation, or replacement of Parcel.<br>
> Its TypeScript loader and native ESM requirements apply only to test code, so keeping E2E configuration in a dedicated package avoids coupling to the frontend compilation pipeline.

#### Testability Instrumentation Required

> Normal browser testing requires no injected application runtime, build hook, relaxed content security policy, or framework-specific source transform.<br>
> Semantic HTML, accessible names, and stable test IDs improve locator reliability but are optional.<br>
> Deterministic PostgreSQL reset, fixture seeding, Yahoo Finance substitution, and readiness endpoints may require test-only backend or harness facilities; these are application-state requirements rather than CodeceptJS runtime requirements.<br>
> Browser interception can avoid application changes when the Playwright helper is used.

### Maintenance Health

#### Latest Stable Release

4.1.0, published 2026-07-30 on GitHub and npm.

#### Release Cadence

> CodeceptJS 4.0.0 was released on 2026-05-21, followed by nine 4.0 patch releases through 2026-07-07 and 4.1.0 on 2026-07-30.<br>
> The concentrated cadence reflects an actively stabilized major release.<br>
> The preceding 3.7 line also received releases through April 2026, while numerous 4.x beta and release-candidate builds were published before general availability.

#### Repository Activity

> The 4.x branch had three human-authored fixes in August 2026, with its latest commit on 2026-08-12.<br>
> An open feature pull request was updated on 2026-08-21.<br>
> Recent merged work covers worker startup, ESM loading, TypeScript, Playwright compatibility, JUnit reporting, Docker arm64, security updates, and recorder error paths.<br>
> Recent contributions and reviews involve DavertMik, kobenguyent, DenysKuchma, mirao, and external contributors rather than only automation.

#### Roadmap And Deprecation Risk

> CodeceptJS 4 is the published direction and is described as a foundation for agent-assisted testing.<br>
> It is a substantial transition to native ESM, Node 20-oriented setup, helper-agnostic elements, strict locators, MCP tooling, and revised plugins.<br>
> Version 4 removed Nightmare, Protractor, TestCafe, and AI helpers, CodeceptUI, the bundled Allure and HTML reporters, and several legacy APIs.<br>
> Existing 3.x projects therefore face real migration work.<br>
> No dated support lifetime or formal roadmap beyond the 4.x direction was found, but active post-release stabilization reduces immediate abandonment risk.

#### Dependency Currency

> Version 4.1.0 was published with Node 24.18.0 and current major dependencies including Mocha 11.7.5, Axios 1.16.1, Zod 4, Cucumber Gherkin 38, and the MCP SDK 1.26.<br>
> Its development matrix used Playwright ^1.59.0, Puppeteer 24.36.0, WebdriverIO 9.23.0, and TypeScript 5.9.3.<br>
> Browser helpers are user-installed, so applications can update them independently, but CodeceptJS wrapper behavior still needs compatibility fixes when upstream APIs or output formats change.

### Community Adoption

#### Npm Downloads

> The npm downloads API recorded 572,087 downloads of codeceptjs from 2026-08-15 through 2026-08-21 and 2,355,761 downloads from 2026-07-22 through 2026-08-21.

#### Github Metrics

> Observed 2026-08-22 for codeceptjs/CodeceptJS: 4,240 stars, 758 forks, 95 subscribers, 186 open issues, 218 open issues plus pull requests in the repository summary, and approximately 390 contributors. npm listed 64 dependent packages.<br>
> The repository was pushed on 2026-08-21 and the latest 4.x commit was dated 2026-08-12.

#### Ecosystem Usage

> The project has operated since 2015 and provides maintained helpers for Playwright, WebDriver, Puppeteer, REST, GraphQL, Appium, and Detox, plus integrations for cloud grids and test management.<br>
> Community packages cover visual comparison, accessibility, Lighthouse, ReportPortal, Xray, TestRail, and other services.<br>
> Registry dependents and download volume demonstrate material ecosystem use, but no independently verified current list of production organizations was found, and community helper maintenance varies.

#### Community Support

> Support is available through extensive official guides and API references, GitHub issues and discussions, a community forum, Slack, examples, and long-running third-party articles and Stack Overflow material.<br>
> The 4.x documentation now covers migration, ESM, TypeScript, locators, CI, debugging, MCP, and helper-specific behavior.<br>
> Some older pages and examples remain discoverable and can conflict with 4.x guidance, so version awareness is necessary.

#### Adoption Trend

> Direct npm usage increased sharply year over year around the 4.x release.<br>
> The comparable seven-day count rose from 163,703 in 2025 to 572,087 in 2026, about 249.5 percent, and the comparable 31-day count rose from 732,771 to 2,355,761, about 221.5 percent.<br>
> This indicates strong current package acquisition, although release automation, CI installs, caching behavior, and the major-version migration can inflate download growth relative to active-user growth.

#### Adoption Metric Normalization

> All download values refer to the direct npm package codeceptjs and exact UTC date ranges.<br>
> They may include repeated CI installs, mirrors, upgrades, automated analysis, and transitive installs and do not count distinct users or projects.<br>
> GitHub stars, forks, issues, and contributors refer only to codeceptjs/CodeceptJS. npm's 64 dependents count is registry-specific.<br>
> Helper packages such as playwright and webdriverio have much broader usage that must not be attributed to CodeceptJS.

### Browser And Runtime Coverage

#### Browser Engines

> With its recommended Playwright helper, CodeceptJS runs Chromium, Firefox, and Playwright WebKit on Linux and can launch branded Chrome or Edge channels where installed.<br>
> The WebDriver helper extends reach to installed Chrome, Chromium, Firefox, Edge, native Safari, Selenium Grid, and commercial grids.<br>
> Puppeteer primarily provides Chromium-family coverage.<br>
> Backend differences mean one scenario API does not guarantee identical capability or behavior across helpers.

#### Browser Protocol

> The transport depends on the selected helper.<br>
> Playwright uses Playwright's browser server and framework-specific protocol, Puppeteer primarily uses Chrome DevTools Protocol, and WebdriverIO uses W3C WebDriver with WebDriver BiDi or vendor transports according to browser and WebdriverIO support.<br>
> CodeceptJS adds its recorder and actor abstraction above these transports, so low-level features and error semantics are not uniform.

#### Headless And Headed Modes

> Playwright, Puppeteer, and WebDriver helpers support normal headed local sessions and headless CI execution.<br>
> CodeceptJS configuration and the browser plugin can switch show/headless behavior without changing scenarios; generated CI setup detects CI and runs headless.<br>
> Xvfb remains available when a headed Linux process is required.

#### Browser Version Management

> The Playwright installation command downloads browser builds matched to the installed Playwright version and can install Linux dependencies; package-lock.json and a pinned image or Playwright version provide determinism.<br>
> Puppeteer follows its own managed download policy.<br>
> Current WebdriverIO can download matching Chrome, Chromium, or Firefox and drivers, while remote grids manage versions externally.<br>
> CodeceptJS itself does not impose one cross-helper browser lock, so the repository must pin the chosen helper and browser strategy.

#### Parallel Browser Support

> CodeceptJS supplies worker threads, dynamic pool distribution, file sharding, a browser configuration plugin, and a programmatic Workers API that can assign groups to different browser configurations.<br>
> GitHub Actions matrices can run Chromium, Firefox, and WebKit independently.<br>
> Worker output is merged within one run, while cross-job report aggregation depends on JUnit processing or a reporting service.

#### Mobile Emulation

> The Playwright helper accepts contextOptions and emulate settings, supports Playwright device descriptors in additional sessions, and can control viewport, touch, mobile flags, user agent, and device scale.<br>
> Browser and window-size overrides can be applied from the CLI.<br>
> Appium is available for real or virtual mobile devices.<br>
> Desktop emulation is not equivalent to testing a real mobile browser or device.

#### Real Browser Fidelity

> Playwright's bundled Chromium, Firefox, and WebKit builds are controlled and reproducible but are not every vendor browser release.<br>
> Playwright WebKit on Ubuntu is not native macOS or iOS Safari.<br>
> The WebDriver helper can automate installed vendor Chrome, Edge, Firefox, and native Safari or a remote real-device grid, trading some determinism and speed for vendor fidelity.

#### Environment Determinism Controls

> The Playwright helper forwards browser launch and context options, including timezone, locale, viewport, user agent, geolocation, permissions, color scheme, reduced motion, device scale, storage state, and service-worker policy.<br>
> Direct browserContext and page access enables clock APIs, request routing, offline mode, animation disabling, and seeded application setup.<br>
> CodeceptJS does not provide one helper-neutral API for all controls; WebDriver and Puppeteer configurations require backend-specific capabilities or custom helper code.

### Application Fit

#### Dynamic Dom Synchronization

> The Playwright helper combines upstream locator waiting with CodeceptJS automatic element waits, failed-step retry, explicit waitForElement, waitForVisible, waitForText, waitForFunction, and URL wait commands.<br>
> Selector-based actor commands locate elements for each step, which is suitable for HTMX node replacement.<br>
> CodeceptJS has no HTMX-specific completion signal, and its default 100 ms waitForAction is not a correctness condition; tests should wait on the post-swap DOM or API state and avoid fixed delays and retained native element handles.

#### Routing Support

> Navigo History API navigation, URL changes, reloads, browser back/forward, and direct deep links are ordinary browser behaviors supported by all browser helpers.<br>
> The base URL can point to Parcel or Gin, and scenarios can assert the current URL and route-specific DOM.<br>
> Direct deep-link reliability still depends on the application's server fallback and should be synchronized on a route-specific visible state rather than navigation timing alone.

#### Locator Model

> CodeceptJS 4 supports semantic text and label locators, ARIA role and accessible-name locators, CSS, XPath, id, name, accessibility id, shadow DOM, Playwright-native selectors, custom test-id schemes, contexts, and a locate builder.<br>
> Strict mode or per-step exact mode rejects ambiguous matches; otherwise the first match is used.<br>
> Locator commands are resolved per action, helping after HTMX swaps.<br>
> ARIA plus a stable context is a strong default for this application.

#### Form Interaction

> The unified API covers click, fillField, appendField, clearField, selectOption, checkOption, attachFile, focus-sensitive type and pressKey, and native browser submission.<br>
> Playwright can provide realistic keyboard, pointer, focus, blur, and validation behavior through direct access where the abstraction is insufficient.<br>
> Tests should trigger blur with Tab or an explicit native action and assert both visible validation and hidden-value synchronization for dynamic rows.

#### Canvas And Download Support

> Chart.js can be checked semantically by using usePlaywrightTo or injected page access to inspect chart configuration, datasets, canvas dimensions, accessible fallback, or interaction state; screenshots can provide secondary rendering evidence.<br>
> The Playwright helper supports download handling, and the FileSystem helper supplies waitForFile and filename checks for browser downloads.<br>
> Robust tests need per-test download directories and cleanup.<br>
> Neither canvas internals nor download contents are validated automatically by the actor DSL.

#### Network And Api Access

> The REST helper can seed and clean fixtures through Axios, while the Playwright helper can make API requests with the browser session's cookies.<br>
> Playwright mockRoute, request/response waiters, and traffic recording support request observation, interception, modification, and controlled Yahoo Finance responses.<br>
> Puppeteer has analogous interception; WebDriver mocking relies on the separate Polly-based MockRequest helper and has navigation limitations.<br>
> The Playwright backend is therefore the strongest fit for deterministic external-service replacement.

#### Same Origin Support

> CodeceptJS operates outside the application page and requires no CORS changes for normal E2E use.<br>
> It can target the consolidated production Gin server or Parcel on port 8000, where the existing /api proxy keeps browser requests same-origin while forwarding to port 8080.<br>
> Helper REST endpoints can be configured separately when fixture calls intentionally bypass the browser origin.

#### Test Isolation

> Configure the Playwright helper to restart the browser context between scenarios for cookie, local-storage, cache, and permission isolation while retaining the browser process.<br>
> Named session blocks create additional contexts for multi-user cases.<br>
> Each worker is an independent CodeceptJS instance.<br>
> Database state remains external: parallel workers need unique portfolio data or disposable PostgreSQL databases or schemas, and scenarios sharing one database state must run serially.

#### External Server Model

> Every browser helper accepts a configurable base URL and can test an independently started application on any reachable host and port.<br>
> CodeceptJS does not require ownership of the server process, so the same scenarios can target Parcel development, host-run Docker Compose, or a production-like Gin container.

#### Application Lifecycle

> Configuration supports async bootstrap, teardown, bootstrapAll, and teardownAll around serial or worker runs, which can invoke repository lifecycle utilities.<br>
> There is no Playwright-Test-style built-in webServer block that automatically builds Parcel, starts Docker Compose, waits for PostgreSQL and Flyway, polls Gin and frontend readiness, and guarantees teardown.<br>
> This repository needs an outer workflow or custom bootstrap module with try/finally cleanup; cancellation-safe Compose teardown should remain a CI step.

#### Visual Regression Workflow

> Core CodeceptJS captures screenshots but does not provide a first-party baseline review and pixel-diff system in 4.1.<br>
> The official guide points to community Resemble, Pixelmatch, and visual helpers or commercial Applitools.<br>
> Community options can create baseline, current, and diff folders with mismatch tolerances, but maintenance, masking, per-browser baselines, font readiness, animation stabilization, and CI review policy must be evaluated and configured separately.<br>
> Chart.js should primarily use data assertions, with narrowly scoped visual checks.

#### Accessibility Audit Integration

> CodeceptJS 4 offers ARIA role locators and grabAriaSnapshot, but those are not standards audits. axe-core can be integrated through community codeceptjs-a11y-helper, playwright-axe, or direct axe-core execution from a custom helper.<br>
> This remains open-source and can emit ordinary test failures and JSON artifacts, but there is no bundled first-party axe audit runner or dedicated accessibility report in core.

### Reliability

#### Waiting Model

> With Playwright, actions inherit upstream locator waiting and CodeceptJS adds helper waits, configurable waitForNavigation, a default post-action delay, assertion commands, explicit wait* methods, and retryFailedStep.<br>
> CodeceptJS documentation describes automatic waiting, but helper abstraction and step retries can hide which condition succeeded.<br>
> HTMX and Navigo tests should express state-based DOM, URL, response, or function waits and should not use generic network idle or I.wait as the primary completion signal.

#### Flake Controls

> Controls include per-step timeout and retry options, retryFailedStep, scenario retries through Mocha, retryTo blocks, run-rerun with required success counts, shuffle, test timeout, workers, serial execution, and CI sharding.<br>
> Playwright contexts and route mocking can remove shared browser and network state.<br>
> Retries must be reported as instability rather than treated as success; open 4.x issues around retry configuration and occasional hangs justify a compatibility spike before relying on defaults.

#### Isolation Model

> The preferred model is one fresh Playwright browser context per scenario, with explicit session contexts only when multi-user interaction is required.<br>
> Worker threads run independent framework and helper instances.<br>
> Serial suites can deliberately retain coordinated database state, but before/after hooks must reset browser and server fixtures.<br>
> CodeceptJS does not create PostgreSQL databases, schemas, or transaction boundaries.

#### Parallelism Controls

> run-workers accepts an exact worker count and pool or suite distribution; --shard divides test files across CI machines; the Workers API supports custom groups and per-worker browser config.<br>
> Keep the count at one for database-mutating tests until unique data or per-worker databases exist, then parallelize independent files and browser matrices. bootstrapAll and teardownAll support once-per-run resources, while worker hooks handle isolated resources.

#### Flake Observability

> run-rerun can execute repeated success criteria and JUnit, local HTML, JSON-compatible Mocha reporting, screenshots, pageInfo, Playwright traces, and aiTrace expose failures.<br>
> Core has no persistent local quarantine registry or historical first-attempt-versus-retry dashboard.<br>
> Testomat.io adds hosted flaky-test history, parallel merge, and reruns, but equivalent local trend analysis requires retaining JUnit or JSON results and custom processing.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Default and bundled plugins can save failure screenshots, page URL and HTML context, console logs, browser logs, ARIA snapshots, HTTP records, and step-oriented aiTrace markdown.<br>
> The Playwright helper can retain failed-test video and Playwright trace ZIP files containing screenshots, DOM snapshots, console, network, and commands; screencast adds WebM and optional subtitles.<br>
> HAR or specialized network logs require direct Playwright configuration or custom capture.<br>
> All artifacts are written under the configured output tree for CI upload.

#### Debugging Tools

> CodeceptJS provides headed mode, verbose recorder logging, pause and pause-on-failure, an interactive REPL, step-by-step screenshots, direct helper access, Playwright Trace Viewer, and an in-process MCP server that can run or pause a scenario and inspect the live browser.<br>
> It can also run one scenario step by step through an agent.<br>
> Playwright traces provide timeline inspection, while CodeceptJS itself does not offer deterministic time-travel execution.

#### Test Generation

> The CLI scaffolds scenarios, page objects, helpers, and TypeScript definitions.<br>
> Version 4 adds optional MCP and official agent skills that can explore a live page and write verified CodeceptJS commands.<br>
> CodeceptUI was removed, and no current core non-AI recorder equivalent to Playwright codegen was documented.<br>
> Generated or agent-authored scenarios still require review for assertions, fixture isolation, semantic locators, and deterministic waits.

#### Reporters

> Core prints console, step, debug, and verbose output and includes a JUnit reporter plugin.<br>
> Mocha reporter configuration supports JSON and other ecosystem reporters.<br>
> The separately installed open-source @testomatio/reporter supplies a self-contained local HTML report, GitHub pull-request comments, and optional hosted Testomat.io results.<br>
> Bundled Allure, Mochawesome, and htmlReporter integrations were removed in 4.x, so existing projects must migrate reporting configuration.

#### Documentation Quality

> The current site has broad 4.x guides for installation, ESM migration, TypeScript, Playwright, WebDriver, locators, waits, effects, workers, sharding, CI, reports, artifacts, debugging, and AI boundaries.<br>
> API references are detailed and recent release work updated many examples.<br>
> Quality is reduced by old 3.x pages and snippets still being indexed, occasional contradictory Node engine guidance, and helper-neutral claims that need backend-specific qualification.

#### Local Workflow

> Developers can run one file or grep one scenario, select browser and viewport from the CLI, switch headed mode, pause at a step or failure, use a live REPL or MCP session, view Playwright traces, and use run-rerun for intermittent failures. npx codeceptjs check validates configuration and opens the browser before a run.<br>
> A dedicated E2E package would keep this workflow isolated from the frontend's CommonJS Parcel proxy configuration.

#### Failure Log Correlation

> CodeceptJS timestamps step and reporter events and can collect browser console, HTTP, HTML, ARIA, screenshot, video, and trace evidence under one scenario.<br>
> It does not automatically ingest Gin, PostgreSQL, Flyway, or Docker Compose logs.<br>
> The workflow must assign a common run and worker identifier, timestamp service output, preserve the Codecept scenario name and shard, and upload backend and browser artifacts together in an always-running failure step.

#### Artifact Data Exposure

> Screenshots, videos, traces, aiTrace HTML and ARIA, console JSON, request records, local HTML reports, storage snapshots, and downloaded files can expose portfolio values, credentials, cookies, local storage, API bodies, and external-service URLs.<br>
> CodeceptJS offers secret values and maskSensitiveData for step and console output, but these do not comprehensively redact binary or browser artifacts.<br>
> Use synthetic accounts, restrict captured traffic and storage, redact custom logs, disable AI and unnecessary traces, and apply short artifact retention and access controls.

### Github Actions Fit

#### Official Ci Support

> The maintained official CI guide includes GitHub Actions examples for Playwright and WebDriver, workers, sharded matrices, browser matrices, checks, and artifact upload.<br>
> No dedicated CodeceptJS action is required.<br>
> Stock ubuntu-latest plus setup-node, npm ci, Playwright browser installation, and npx codeceptjs run is supported; an official multi-architecture Docker image is also available.

#### Browser Caching

> Use setup-node's npm cache with the lockfile.<br>
> Playwright browsers live in a separate cache and can be restored only with keys tied to operating system, architecture, and the exact Playwright version; the official GitHub example installs them rather than prescribing a browser cache.<br>
> A pinned Playwright Docker image avoids repeated browser downloads but adds image-pull cost.<br>
> WebdriverIO-managed browser and driver caches likewise need version-aware keys.

#### Artifact Integration

> All CodeceptJS and Playwright evidence can be placed under output and uploaded with actions/upload-artifact in an if: failure() or if: always() step.<br>
> JUnit can feed a test-report action, and local HTML can be uploaded or linked from a job summary.<br>
> Compose and backend logs need explicit collection.<br>
> The workflow, not CodeceptJS, controls retention, permissions, missing-file behavior, and redaction.

#### Sharding And Matrix Support

> The built-in --shard index/total option splits files across GitHub matrix jobs, while run-workers handles within-job concurrency and the browser plugin supports a browser and viewport matrix.<br>
> Worker results merge in one process; independent shards do not have a core multi-job HTML merger.<br>
> JUnit aggregation or Testomat.io reporting is needed for a unified cross-job view, and retries remain local to each invocation.

#### Container Compatibility

> CodeceptJS can run on the GitHub host against this repository's Docker Compose published ports, which is the simplest topology.<br>
> Alternatively, the official codeceptjs/codeceptjs image contains Playwright, Puppeteer, and WebDriver support and now targets amd64 and arm64.<br>
> Browser containers require adequate shared memory and network access to host or sibling services.<br>
> Pin the CodeceptJS, helper, and image versions together to avoid browser executable mismatch.

#### Failure Cleanup

> Normal helper and teardown hooks close browser contexts and processes, while teardownAll can release once-per-run resources.<br>
> GitHub cancellation can interrupt JavaScript hooks, so Compose cleanup must also run as a separate if: always() step using docker compose down --volumes --remove-orphans.<br>
> Use job timeouts, disposable volumes, unique project names, and explicit termination of Parcel or other background processes rather than relying only on framework teardown.

### Cost And Risk

#### Open Source Completeness

> The MIT core, Playwright helper integration, local Chromium/Firefox/WebKit execution, workers, sharding, retries, screenshots, traces, videos, JUnit, REST fixtures, request mocking, and custom helpers are sufficient for the required Linux E2E baseline without payment.<br>
> Local HTML reporting and axe or visual checks require additional open-source packages.<br>
> Native Safari and hosted real devices require Apple infrastructure or an optional grid, as with competing tools.

#### Optional Cloud Dependency

> No cloud is required for deterministic test execution or ordinary diagnostics.<br>
> Testomat.io can add hosted history, flake analytics, report merging, and reruns; BrowserStack and similar providers add real-browser and device coverage; Applitools adds hosted visual review.<br>
> Local JUnit, HTML, traces, community pixel diff, and axe-core alternatives exist, although they require more repository configuration.

#### Migration Cost

> Moderate to high.<br>
> Scenarios are concise JavaScript or TypeScript, and CSS, ARIA, HTTP, and Playwright concepts transfer, but Feature, Scenario, the I actor, recorder queue, dependency injection, effects, helpers, page objects, configuration, and plugins are CodeceptJS-specific.<br>
> Direct usePlaywrightTo calls are more portable than custom actor steps but reduce helper neutrality.<br>
> Version 4's ESM and plugin removals demonstrate additional framework-level migration risk.

#### Custom Harness Burden

> Moderate.<br>
> CodeceptJS supplies runner lifecycle hooks, browser isolation, REST seeding, Playwright request mocking, workers, retries, downloads, traces, and artifact paths.<br>
> The repository still needs Compose and Parcel startup, readiness checks, Flyway verification, PostgreSQL reset or namespacing, Yahoo Finance fixture policy, backend-log collection, cancellation-safe teardown, artifact redaction, and likely small Chart.js and axe helpers.<br>
> Core visual baseline review and cross-job report merging are not included.

#### Capability Delivery Tier

> Core CodeceptJS: scenario runner, actor API, assertions, hooks, contexts, workers, sharding, retries, JUnit, screenshots, pageInfo, aiTrace, REST helper, and CLI debugging.<br>
> Upstream helper: Chromium/Firefox/WebKit, actionability, browser contexts, network interception, API requests, downloads, device and environment control, trace, and video through Playwright.<br>
> Community or separate open-source package: local HTML reporting, axe audits, pixel visual comparison, and some service reporters.<br>
> Custom repository code: Compose/Flyway/PostgreSQL lifecycle, readiness, log correlation, deterministic Chart.js checks, cross-job merge, and artifact policy.<br>
> Paid optional tier: hosted browser/device grids, Testomat.io analytics, and commercial visual services.

#### Ai Execution Boundary

> Ordinary CodeceptJS scenarios, workers, retries, Playwright controls, and reports execute without an LLM, model credential, inference cost, or model network egress.<br>
> MCP and agent skills are optional authoring and debugging interfaces.<br>
> AI healing and failure analysis require explicit provider configuration and the --ai flag and can send HTML, errors, and screenshots to a provider.<br>
> Keep --ai disabled in CI, omit provider credentials, disable heal and analyze, and retain deterministic authored scenarios as the source of truth.

### Evidence And Decision

#### Sources

- CodeceptJS official documentation overview and basics: https://codecept.io/basics/
- CodeceptJS 4 installation and TypeScript setup: https://codecept.io/installation/
- CodeceptJS 4 release and migration direction: https://codecept.io/blog/codeceptjs-4/ and https://codecept.io/migration-4/
- CodeceptJS Playwright guide and helper API: https://codecept.io/playwright/ and https://codecept.io/helpers/playwright
- CodeceptJS WebDriver guide and helper API: https://codecept.io/webdriver/ and https://codecept.io/helpers/web-driver/
- CodeceptJS locator and strict-selection guides: https://codecept.io/locators/ and https://codecept.io/element-selection/
- CodeceptJS configuration and lifecycle hooks: https://codecept.io/configuration/
- CodeceptJS parallel execution and sharding: https://codecept.io/parallel/
- CodeceptJS current CI guide: https://codecept.io/continuous-integration/
- CodeceptJS reports and plugins: https://codecept.io/reports/ and https://codecept.io/plugins/
- CodeceptJS debugging, MCP, and AI boundaries: https://codecept.io/debugging/, https://codecept.io/mcp/, and https://codecept.io/ai/
- CodeceptJS visual and community helper documentation: https://codecept.io/visual/ and https://codecept.io/community-helpers/
- CodeceptJS REST and FileSystem helper APIs: https://codecept.io/helpers/rest/ and https://codecept.io/helpers/FileSystem/
- CodeceptJS releases and 4.1.0 notes: https://github.com/codeceptjs/CodeceptJS/releases and https://github.com/codeceptjs/CodeceptJS/releases/tag/4.1.0
- CodeceptJS 4.1.0 npm package metadata and provenance: https://registry.npmjs.org/codeceptjs/latest
- npm downloads API, 2026 periods: https://api.npmjs.org/downloads/point/2026-08-15:2026-08-21/codeceptjs and https://api.npmjs.org/downloads/point/2026-07-22:2026-08-21/codeceptjs
- npm downloads API, comparison periods: https://api.npmjs.org/downloads/point/2025-08-16:2025-08-22/codeceptjs and https://api.npmjs.org/downloads/point/2025-07-22:2025-08-21/codeceptjs
- Core repository metadata, commits, pull requests, issues, and contributors: https://github.com/codeceptjs/CodeceptJS and https://api.github.com/repos/codeceptjs/CodeceptJS
- TypeScript 6 future-proofing change: https://github.com/codeceptjs/CodeceptJS/commit/e27f2a7693151eeb775d20c17201dd7fc245f3e0
- Playwright 1.61 and multi-architecture Docker change: https://github.com/codeceptjs/CodeceptJS/commit/7a5c4f9416cc71b9b8dc21b82c99204dbf2e19e5
- Open Asset Allocator frontend runtime and compiler configuration: src/main/web-static/package.json, src/main/web-static/tsconfig.json, and src/main/web-static/.proxyrc.js
- Open Asset Allocator development and service lifecycle: dev.sh and src/main/docker/docker-compose-base.yml

#### Observed At

2026-08-22

#### Confidence

> High confidence in release, license, npm, repository, installation, ESM, runner, helper, browser, worker, sharding, diagnostic, and documented API facts.<br>
> Medium confidence in project-fit conclusions, issue health, maintainer risk, adoption interpretation, wrapper lag, and custom-harness estimates.<br>
> Low confidence in exact TypeScript 6 behavior, resource cost, supply-chain status, and application behavior because no repository-specific CodeceptJS spike or benchmark was run.

#### Deal Breakers

> No unconditional deal-breaker exists if CodeceptJS is placed in a dedicated E2E package and the Playwright helper is used.<br>
> Exclude it if the project requires zero abstraction above Playwright, one authoring and waiting model, a first-party visual and axe workflow, native Safari on Ubuntu, no additional application lifecycle harness, or verified TypeScript 6 compatibility without a spike.<br>
> Installing 4.x directly into the current frontend package without resolving its CommonJS Parcel proxy is also unacceptable.

#### Recommendation

> Viable alternative, not the primary recommendation.<br>
> CodeceptJS 4.1 is actively maintained, has strong and growing npm use, provides a concise scenario API, supports multiple automation backends, and exposes Playwright's browser coverage, isolation, interception, downloads, traces, and environment controls while adding workers, sharding, retries, REST fixtures, rich failure evidence, and optional agent tooling.<br>
> It ranks below direct Playwright Test for this repository because its recorder and helper abstraction add another synchronization and migration layer, 4.x is still stabilizing, ESM setup conflicts with the current frontend package unless separated, TypeScript 6 is unverified, and visual, accessibility, application lifecycle, database isolation, and report merging need additional packages or harness code.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `wrapper_upstream_lag`
- `maintainer_concentration`
- `resource_usage`
- `security_and_supply_chain`
- `hard_gate_result`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="testplane"></a>
## 11. Testplane

Source result: `Testplane.json`

### Project And Compatibility

#### Implementation Language

> Testplane is implemented primarily in TypeScript and runs on Node.js.<br>
> Tests are authored in TypeScript or JavaScript with Mocha syntax, async functions, WebdriverIO-style browser commands, and expect-webdriverio assertions.<br>
> Go is not required for the test suite.<br>
> [S1][S2][S5]

#### License And Governance

> Testplane core is MIT licensed.<br>
> Package metadata states that the source is published and distributed by Yandex LLC as owner, while development occurs publicly under the gemini-testing organization and seven npm maintainers are listed.<br>
> Governance is maintainer/company-led: no foundation charter, elected steering model, public succession plan, or repository security policy was found.<br>
> The MIT license imposes no material repository-use restriction beyond preserving its notice.<br>
> [S1][S2][S16]

#### Installation Model

> Install with `npm install -D testplane@9` or scaffold with `npm init testplane@latest`.<br>
> The wizard can also install `@testplane/testing-library` and `html-reporter`. `npx testplane install-deps` downloads configured Chrome and Firefox versions, their drivers, and supported Ubuntu packages into `~/.testplane` or `TESTPLANE_BROWSERS_PATH`; missing dependencies can also be downloaded on first run.<br>
> Edge and Safari browsers are not auto-downloaded.<br>
> Selenoid/browser images or commercial grids are optional.<br>
> [S5][S6][S15]

#### Candidate Scope And Layer

> Complete Node.js E2E framework: integrated Mocha-based runner, WebDriver client, assertions, browser management, retries, worker parallelism, browser sets, screenshots and visual comparison, failure history, hooks, plugins, and an interactive report.<br>
> Sharding and several lifecycle features are official companion plugins rather than core.<br>
> [S2][S8][S13][S14]

#### Authoring And Async Model

> Native JavaScript and TypeScript async/await with Mocha-style `describe` and `it`, a WebdriverIO-derived browser and element API, polling `expect` matchers, and optional Testing Library queries.<br>
> Testplane 9 favors explicit imports but `@testplane/globals` can restore global typings.<br>
> There is no Cypress-style command queue; every asynchronous browser command must be awaited.<br>
> [S5][S7][S9]

#### Build Pipeline Coupling

> Node-environment black-box E2E tests navigate to `baseUrl` and can exercise the existing Parcel build served by Gin or the Parcel development proxy without transforming application source.<br>
> Testplane's separate browser/component-test mode uses bundled Vite, but that mode is not required for E2E.<br>
> Visual checks, CDP helpers, and Time Travel inject test-session behavior into the page without replacing the production build pipeline.<br>
> [S5][S8][S14][S18]

#### Testability Instrumentation Required

> Ordinary WebDriver E2E execution requires no production source transform, special route, injected shipped runtime, or relaxed browser security policy.<br>
> Existing accessible names, visible text, IDs, and CSS selectors are usable; stable `data-testid` attributes are optional where semantic locators are unavailable.<br>
> A health route is useful for readiness, and deterministic database reset plus server-side Yahoo Finance replacement require test-only harness seams outside Testplane.<br>
> Time Travel and CDP tooling inject scripts only into controlled sessions.<br>
> [S9][S11][S12][S18]

### Maintenance Health

#### Latest Stable Release

> Testplane 9.1.1, published to npm on 2026-08-21 from commit 9828757.<br>
> It is the registry's latest version; the changelog records a selectivity navigation fix.<br>
> GitHub's Releases page had not created a separate v9.1.1 release entry when observed.<br>
> [S1][S2][S3]

#### Release Cadence

> Very frequent.<br>
> The v9 line published 16 stable versions from 9.0.0 on 2026-06-16 through 9.1.1 on 2026-08-21, while the v8 maintenance branch also received releases through 8.47.5.<br>
> Some tags contain dependency or metadata changes rather than user-facing features, but the cadence includes a major visual-engine release and repeated browser, isolation, Time Travel, and selectivity fixes.<br>
> [S2][S3][S7]

#### Repository Activity

> Active as observed on 2026-08-22.<br>
> The default branch was pushed on 2026-08-21, recent commits came from Roman Kuznetsov, Sergey Cherebedov, Nikolai Markov, Dmitriy Dudkevich, and the y-infra account, and at least nine pull requests were merged from 2026-07-28 through 2026-08-21.<br>
> Open work on 2026-08-22 included profiler, iframe pixel-ratio, vulnerability, and external-contributor changes.<br>
> [S3][S21]

#### Roadmap And Deprecation Risk

> Testplane is not deprecated.<br>
> Version 9 removed the unreliable `automationProtocol: devtools` mode, requires Node 22+, defaults local execution to WebDriver, enables its WebSocket/BiDi driver path, changes global-helper typing, and introduced a rewritten screenshot engine that may change about 5% of baselines.<br>
> Both v8 and v9 are maintained, but no formal long-term roadmap or support-period policy was found.<br>
> Ongoing open work on BiDi, profiler, browser images, and TypeScript typing shows direction but also architectural transition risk.<br>
> [S2][S7][S21]

### Community Adoption

#### Npm Downloads

> The npm Downloads API recorded 14,620 downloads of the exact `testplane` package for the complete seven-day window 2026-08-15 through 2026-08-21 and 60,812 downloads for 2026-07-22 through 2026-08-21.<br>
> [S4]

#### Ecosystem Usage

> The gemini-testing organization maintains the core, html-reporter, looks-same, Testing Library adapter, GitHub Action, chunks, Storybook, global-hook, JSON/stat reporters, test repeater, VS Code extension, browser mocks, OAuth, profiling, and retry plugins.<br>
> Remote Selenium, Selenoid, Appium capabilities, BrowserStack, and Sauce Labs are documented.<br>
> The framework inherits a long Hermione history and explicit Yandex ownership, but independently verifiable named production users outside that ecosystem were not found.<br>
> [S2][S6][S13][S16]

#### Community Support

> Current versioned English and Russian documentation covers setup, commands, configuration, visual testing, CI, migrations, and plugins.<br>
> Support channels include GitHub Issues and Discussions, a Telegram community link, Stack Overflow's `testplane` tag, and public docs contributions.<br>
> The available knowledge base is much smaller than Playwright, Cypress, Selenium, or upstream WebdriverIO, and many searches still surface former Hermione terminology.<br>
> [S2][S14][S16]

#### Adoption Trend

> Strongly growing for the renamed package, with an important comparability caveat.<br>
> Direct 31-day downloads rose from 6,826 in 2025-07-22 through 2025-08-21 to 60,812 in 2026-07-22 through 2026-08-21, about 790.9%.<br>
> The matching seven-day count rose from 1,836 to 14,620.<br>
> Version 9, new documentation, an official action, and maintained companion tools support a growth signal, but migration from Hermione and repeated CI installs can create growth without an equivalent increase in unique organizations.<br>
> [S4][S15][S16]

#### Adoption Metric Normalization

> All download values use the exact `testplane` npm package and complete UTC windows.<br>
> They exclude downloads still attributed to the historical `hermione` package and companion packages such as `html-reporter`, so they undercount the total installed ecosystem while still including repeat CI installs, bots, mirrors, and cache misses.<br>
> The package represents the complete runner, making it more comparable than a low-level driver.<br>
> Stars and dependents are cumulative/incomplete signals, not active production-suite counts.<br>
> [S1][S4]

### Browser And Runtime Coverage

#### Browser Engines

> Through W3C WebDriver, Testplane can drive installed or remote Chrome/Chromium and Edge, Firefox, and native Safari, plus Appium targets and older browsers supported by their drivers.<br>
> Local auto-download covers Chrome and Firefox; Edge's driver but not browser can be downloaded, and Safari requires macOS.<br>
> Ubuntu has no local native Safari/WebKit equivalent.<br>
> Browser-specific CDP features work only on Chromium.<br>
> [S6][S7][S8]

#### Browser Protocol

> Testplane 9 uses W3C WebDriver as its only primary automation protocol.<br>
> WebDriver BiDi/WebSocket support is enabled for compatible Chrome, Edge, and Firefox sessions; Safari did not have documented BiDi support.<br>
> Chromium CDP remains available through `getPuppeteer` or experimental CDP APIs for network mocks, accessibility snapshots, and emulation, but the former all-CDP `devtools` automation mode was removed because of hard-to-reproduce failures.<br>
> [S7][S8][S11][S17]

#### Headless And Headed Modes

> Chrome/Chromium, Firefox, and Edge can run headed for local debugging and headless for CI through configuration.<br>
> The official GitHub guide requires headless browser execution on hosted Ubuntu and documents Chrome sandbox flags.<br>
> Remote grids can expose VNC, video, and logs.<br>
> Native Safari execution is macOS or cloud based rather than an Ubuntu headless path.<br>
> [S6][S15]

#### Browser Version Management

> `install-deps` can download exact configured Chrome and Firefox versions plus compatible Chrome, Firefox, and Edge drivers, with the cache rooted at `~/.testplane` or `TESTPLANE_BROWSERS_PATH`.<br>
> Browser versions are set through WebDriver capabilities.<br>
> Edge and Safari binaries must be supplied by the host or grid.<br>
> Official documentation warns that ready-made Selenoid images may not contain current browsers; deterministic CI should pin browser versions, cache keys, and container digests.<br>
> [S6][S15]

#### Parallel Browser Support

> Browser IDs and sets define selected browser combinations. `sessionsPerBrowser`, worker count, process `parallelLimit`, and per-session reuse control local concurrency.<br>
> The official `@testplane/chunks` plugin deterministically partitions tests across CI jobs, and html-reporter can merge shard reports.<br>
> There is no single Playwright-style projects object, but the set/browser/plugin model covers matrices and sharding.<br>
> [S8][S13]

#### Mobile Emulation

> WebDriver capabilities can configure Chromium mobile emulation, viewport, device scale, user agent, touch, and orientation; documented examples use an iPhone 12 Pro profile.<br>
> Appium/Selenium grids and commercial clouds can supply Android, iOS simulators, or physical devices, and Testplane 9 added iOS 26 screenshot handling.<br>
> Desktop emulation is not a physical mobile device, and many emulation controls are Chromium/CDP-specific.<br>
> [S6][S7]

#### Real Browser Fidelity

> WebDriver controls installed vendor browsers, so Chrome, Firefox, Edge, and Safari sessions are real browser products rather than a JavaScript DOM simulation.<br>
> Native Safari means SafariDriver on macOS or a real cloud device; Chromium emulation and an unrelated WebKit build do not establish Safari behavior.<br>
> Selenoid/cloud fidelity depends on the selected image, OS, version, and provider, while CDP-only helpers reduce feature parity outside Chromium.<br>
> [S6][S7]

#### Environment Determinism Controls

> Core configuration controls browser/version, window size, orientation, headless mode, browser capabilities, context isolation, cursor reset, screenshot delay, animation disabling during `assertView`, and network throttling.<br>
> Puppeteer/CDP can additionally control CPU and Chromium features.<br>
> There is no unified cross-browser fixture API for clock, timezone, locale, permissions, geolocation, random seeds, all animations, fonts, or reduced motion.<br>
> Those require capabilities, preload/execute helpers, fixed images/fonts, and repository fixtures.<br>
> [S8][S10]

### Application Fit

#### Routing Support

> Navigo is framework-agnostic to Testplane.<br>
> Tests can use `browser.url`, relative `baseUrl`, real links, WebDriver back/forward/refresh commands, URL assertions, and route-specific DOM waits.<br>
> Direct SPA deep links work if Gin or Parcel supplies the application shell for the route. `openAndWait` can wait for selectors, predicates, and Chromium network idle, but route cleanup and History API correctness must be asserted by application-specific tests.<br>
> [S8][S9][S18]

#### Locator Model

> Core WebdriverIO-compatible selectors include CSS, XPath, link/text, accessibility `aria/...`, IDs, and shadow DOM.<br>
> The official `@testplane/testing-library` adapter adds role, label, text, placeholder, display-value, and test-id queries and is included by the project wizard. `$` returns the first match rather than enforcing strict uniqueness, so cardinality must be asserted where ambiguity matters; re-query after HTMX replacements.<br>
> [S5][S9]

#### Form Interaction

> WebDriver element commands cover click, set/add/clear value, keyboard actions, focus-sensitive interactions, selects, files, drag, touch, and JavaScript execution.<br>
> Tests can press Tab or click elsewhere to trigger blur, inspect focus/value/validity, assert synchronized hidden fields, and manipulate dynamic rows through real browser events.<br>
> Driver differences and custom widgets still require targeted assertions; direct script assignment should not replace user-level interactions under test.<br>
> [S9]

#### Canvas And Download Support

> Chart.js can be validated semantically through `browser.execute` by inspecting the chart instance, canvas dimensions, labels, datasets, and exported values, with stabilized `assertView` checks as secondary evidence.<br>
> Testplane has screenshot/PDF and upload commands but no documented first-class download event/save stream.<br>
> True downloads therefore require browser profile preferences plus a worker-specific filesystem directory and polling, or Chromium Puppeteer/CDP code.<br>
> Filename, completion, content, and cleanup remain harness work.<br>
> [S9][S10][S18]

#### Network And Api Access

> Node-side tests can use Node 24 `fetch` or repository clients to seed REST fixtures.<br>
> Core `browser.mock` can observe, respond, redirect, or abort browser requests, but official documentation says it works only with CDP, so it is Chromium-specific despite BiDi support.<br>
> Puppeteer and raw CDP provide deeper Chromium interception.<br>
> Browser interception cannot replace Yahoo Finance calls made by the Go backend; use a fake upstream service/configuration seam or seeded data for cross-browser determinism.<br>
> [S11][S18]

#### Same Origin Support

> Tests can target the consolidated Gin production server or Parcel's development server/proxy through `baseUrl` with no Testplane-specific CORS changes.<br>
> WebDriver preserves normal browser origin behavior.<br>
> Cross-origin test flows remain possible through WebDriver window/navigation APIs, while request mocking and some diagnostics may be Chromium-only.<br>
> [S8][S18]

#### Test Isolation

> Chrome 93+ defaults to a fresh isolated browser context for tests while reusing the underlying session; other browsers default to no context isolation. `clearSession`, cookies, local/session storage commands, state save/restore, and session recreation provide explicit cleanup.<br>
> IndexedDB is not included in saved state, and PostgreSQL remains external.<br>
> Firefox/Safari safety requires explicit browser-state cleanup or low `testsPerSession`, while all browsers need unique server data or disposable databases.<br>
> [S8]

#### External Server Model

> Supported.<br>
> Each browser has a configurable `baseUrl` and WebDriver grid URL, so the same test suite can target an independently started Parcel/Gin stack, Docker Compose, a local host process, a remote Selenium grid, or a deployed environment on arbitrary reachable ports.<br>
> Testplane need not own the application process.<br>
> [S6][S8]

#### Application Lifecycle

> Core `devServer` starts one command, streams its logs, can reuse an existing server, supports an async or HTTP readiness probe, and stops its child process with the runner; `beforeAll` and `afterAll` support additional setup/cleanup.<br>
> A Compose command could wrap the stack, but Testplane does not natively model PostgreSQL, Flyway, Gin, Parcel, health dependencies, volumes, and log collection as separate services.<br>
> This repository should retain Make/GitHub Actions orchestration and use always-running teardown.<br>
> [S12][S18]

#### Visual Regression Workflow

> Visual regression is a core strength. `assertView` captures element, viewport, or full-page/composite images, auto-waits for the target, waits for static resources, disables animation, ignores selected elements/caret/antialiasing, supports pixel and color tolerances, and stores per-browser references.<br>
> Testplane 9 improves long/scrollable elements, sticky overlays, hover suppression, pseudo-elements, and iOS cropping.<br>
> Open-source html-reporter offers six diff views, one-click review/update, `--update-refs`, static CI reports, and report merging.<br>
> Pin OS/browser/fonts/viewport and await Chart.js animation before a limited visual assertion.<br>
> [S7][S10][S14]

#### Accessibility Audit Integration

> Testing Library queries and computed role/label commands support accessible authoring.<br>
> Official accessibility guidance uses Puppeteer's Chromium accessibility-tree snapshot and is explicitly Chrome/CDP-only.<br>
> Current `axe-core` can be injected or called from a custom Testplane helper and serialized to open-source reports, but Testplane has no first-party axe adapter, cross-browser ARIA snapshot matcher, or integrated audit report.<br>
> [S9][S17]

### Reliability

#### Waiting Model

> Element lookup waits for existence; WebDriver actions and explicit `waitFor*` commands cover visibility, enabled, and clickable states; `expect` assertions poll; `waitUntil` handles application predicates; and navigation has page-load timeouts. `openAndWait` can combine selector, predicate, network-idle, and failed-resource checks, but network-idle/error features are CDP-only.<br>
> HTMX tests should wait for a post-swap user-visible condition or known request rather than use `pause`.<br>
> [S8][S9][S11]

#### Flake Controls

> Core controls include per-browser retry counts, conditional `shouldRetry`, test/page/WebDriver/wait timeouts, session rejection patterns, command history, last-failed reruns, process/session recycling, tags, and selective execution.<br>
> Official plugins add command retry, progressive retry, retry limits, repeat runs, and deterministic chunks.<br>
> Visual checks disable animations and wait for static resources.<br>
> Retries can conceal database side effects, so setup/reset must be idempotent and first-attempt failures retained.<br>
> [S8][S13][S14][S16]

#### Isolation Model

> Chrome 93+ can create an isolated browser context per test while keeping a session warm.<br>
> Other browsers reuse the session by default and need `clearSession`, hooks, `testsPerSession: 1`, or `reloadSession` for stronger isolation.<br>
> Node workers and browser sessions are independently bounded.<br>
> Stateful PostgreSQL scenarios can run serially, but database schemas, transactions, reset, and retry-safe fixture cleanup are repository responsibilities.<br>
> [S8][S12][S13]

#### Parallelism Controls

> `sessionsPerBrowser`, worker count, `parallelLimit`, browser sets, and `testsPerSession` give independent concurrency controls.<br>
> Database-sensitive runs can use one browser session and one worker or a separate serial set, while unrelated read-only suites use greater concurrency. `@testplane/chunks` distributes future CI work.<br>
> Safe PostgreSQL parallelism still requires worker/shard-specific databases or data namespaces and cannot be inferred from browser-context isolation.<br>
> [S8][S13][S18]

#### Flake Observability

> HTML reports retain every retry, failure, error, command history, metadata, screenshot diff, and filtering for retried tests.<br>
> Time Travel modes can record all runs, retries only, or the last failed run; last-failed JSON and repeat plugins help reproduction.<br>
> JSON/stat reporters enable local post-processing.<br>
> There is no core long-term quarantine registry, first-attempt trend database, or historical dashboard, so durable flake rates and ownership need repository-owned ingestion or external tooling.<br>
> [S8][S14][S16]

### Diagnostics And Developer Experience

#### Failure Artifacts

> Core failure artifacts include automatic full-page or viewport screenshots, visual expected/actual/diff images, stack traces, source snippets, command history, test metadata, JSONL output, and optional Time Travel DOM snapshots. html-reporter stores results in SQLite and exposes retries and logs.<br>
> Selenoid/cloud grids can add video, VNC, driver, network, and console logs.<br>
> Complete HAR/network archives and video are not uniform core artifacts, and backend/Compose logs remain separate.<br>
> [S8][S14]

#### Debugging Tools

> `npx testplane gui` provides live execution, screenshots, diffs, command history, targeted reruns, and CI-report inspection.<br>
> Headed mode, `browser.debug`, Node inspector, REPL before/on failure, `switchToRepl`, keep-browser-on-fail, a VS Code extension, and Time Travel DOM replay cover interactive and postmortem debugging.<br>
> Time Travel replays DOM state rather than a video and does not yet provide a complete network timeline.<br>
> [S14][S16]

#### Test Generation

> The project wizard creates configuration and an example test, while Testing Library helpers improve locator authoring.<br>
> No maintained deterministic core recorder/code generator comparable to Playwright codegen was documented.<br>
> Optional Testplane AI/CLI and agent tooling can inspect pages and reports, but generated tests still require human review and AI is not necessary for authoring or execution.<br>
> [S5][S14][S16]

#### Reporters

> Core terminal reporters include flat, plain, and JSONL with file output and multiple-reporter support.<br>
> Official organization packages provide JSON, statistics, TeamCity, and the feature-rich html-reporter; custom reporters can subscribe to runner events.<br>
> HTML reports can merge SQLite/data sources across shards.<br>
> No first-party JUnit reporter was found in the current official catalog, so GitHub test annotations or JUnit XML require a community/custom converter.<br>
> [S13][S14][S16]

#### Documentation Quality

> The current v9 site is extensive, versioned, searchable, bilingual, and contains guides, command/config references, migration material, examples, a blog, and source edit links.<br>
> It is actively maintained in 2026.<br>
> Accuracy is uneven at edges: some current CI examples still use unsupported Node 20 or older action versions, the browser reference displays removed `devtools` values before correcting them, and tolerance defaults conflict within one page.<br>
> Setup claims should be checked against v9 package metadata and migration notes.<br>
> [S5][S7][S8][S15]

#### Local Workflow

> Run all tests with `npx testplane`, select a path, test name with `--grep`, tag, set, or browser, use `--last-failed-only`, and update references with `--update-refs`.<br>
> GUI/headed mode supports visual iteration, REPL and inspector support breakpoints, keep-browser preserves a failed session, and official repeat/last-failed tools aid reproduction.<br>
> Focused local workflow is strong, although plugin installation and configuration add more moving parts than a single-package runner.<br>
> [S8][S14]

#### Failure Log Correlation

> Testplane adds a unique X-Request-ID containing per-test/retry and per-WebDriver-request identifiers to grid traffic and records test PID, browser/session metadata, timestamps, command history, and reports.<br>
> Dev-server logs are prefixed and streamed.<br>
> Browser application requests do not automatically receive that WebDriver header, and Gin, Flyway, PostgreSQL, and Compose logs are not ingested into one timeline.<br>
> The CI harness should propagate its own run/test ID and upload UTC-stamped service logs beside reports.<br>
> [S8][S12][S14]

#### Artifact Data Exposure

> Screenshots, visual baselines/diffs, Time Travel DOM snapshots, SQLite/JSON reports, command arguments, browser/grid logs, state files, videos, and network mock calls can expose portfolio data, credentials, cookies, tokens, or downloaded statements.<br>
> Official authorization guidance warns not to commit state files, and a 2025 fix removed tokens from unhandled-rejection logs.<br>
> There is no universal artifact-redaction layer.<br>
> Use synthetic data, selective logging, restricted artifacts, short retention, secret masking, and disable unnecessary Time Travel/body capture.<br>
> [S2][S8][S14]

### Github Actions Fit

#### Official Ci Support

> The maintained `gemini-testing/gh-actions-testplane@v1` action runs Testplane, caches local browsers, writes failed-test statistics to the job summary, and returns HTML-report path and exit-code outputs.<br>
> Official docs target GitHub-hosted Ubuntu and standard checkout/setup-node actions.<br>
> The action itself uses Node 20 internally, which GitHub Actions supports, but current documentation examples specifying application Node 20 conflict with Testplane 9's Node >=22 requirement and must be changed to Node 24.<br>
> [S1][S15][S20]

#### Browser Caching

> The official action automatically caches local browsers, while npm dependencies use setup-node's npm cache.<br>
> Manual workflows can cache `TESTPLANE_BROWSERS_PATH` keyed by OS, architecture, browser/version, Testplane lockfile, and system image.<br>
> Browser system packages are runner-image state rather than portable cache content.<br>
> Pinning explicit versions is safer than reusing an unversioned cache, and Edge/Safari remain externally managed.<br>
> [S6][S15][S20]

#### Artifact Integration

> The official action exposes `html-report-path`; documentation shows `actions/upload-artifact@v4` under `if: always()` and optional GitHub Pages publishing.<br>
> Screenshots, diffs, SQLite, Time Travel snapshots, JSONL, browser/grid logs, and Compose logs can be uploaded as standard artifacts.<br>
> Privacy, retention, naming, and non-report service-log collection remain workflow configuration.<br>
> [S14][S15][S20]

#### Sharding And Matrix Support

> Use a GitHub matrix with the official `@testplane/chunks` count/run parameters to partition tests, optionally add browser/set axes, upload each HTML report, and run `html-reporter merge-reports` in an aggregation job.<br>
> Retries occur within a shard; a GitHub job retry reruns the shard.<br>
> Official examples use S3, but GitHub artifacts can replace it with custom download/merge steps.<br>
> Database-mutating shards require independent PostgreSQL instances or serial execution.<br>
> [S13][S15]

#### Container Compatibility

> Testplane can run on the Ubuntu host beside this repository's Docker Compose stack or connect to Selenoid/Selenium browser containers.<br>
> Host execution avoids container-to-host browser networking complexity.<br>
> If browsers are containerized, the application must bind to a reachable interface and use service/host-gateway DNS because container localhost is not the Compose host.<br>
> Ready-made Selenoid images may lag current browsers, and the project still has an open issue to build current browser images.<br>
> [S6][S18][S21]

#### Failure Cleanup

> Testplane normally closes sessions and its `devServer` child and provides `afterAll`, but Compose services, PostgreSQL volumes, Flyway state, downloaded files, and independently started processes need workflow cleanup guarded by `if: always()`.<br>
> Cancellation can bypass process hooks, so use disposable runners/databases, unique Compose project names, timeouts, idempotent pre-run cleanup, and an always-running `docker compose down --volumes --remove-orphans` step after log collection.<br>
> [S12][S15][S18]

### Cost And Risk

#### Open Source Completeness

> Core runner, local Chrome/Firefox management, WebDriver execution, retries, parallelism, visual comparison, screenshots, command history, state tools, and Time Travel are MIT-licensed.<br>
> Official open-source companions provide Testing Library, HTML/JSON reports, GitHub Action, chunks, repeat/retry plugins, and VS Code integration.<br>
> No paid service is required for the repository's Ubuntu Chrome/Firefox baseline, but Safari/real devices need owned infrastructure or a provider.<br>
> [S1][S13][S14][S16]

#### Optional Cloud Dependency

> BrowserStack, Sauce Labs, and other Selenium/Appium grids are optional for native Safari, physical devices, OS/version breadth, VNC, provider video, and high concurrency.<br>
> Equivalent basic local Chrome/Firefox execution, visual baselines, reports, retries, sharding, and DOM replay are available without a cloud.<br>
> Hosted long-term analytics and broad real-device coverage have no complete local equivalent, but they are not required for base CI.<br>
> [S6][S13][S14]

#### Migration Cost

> Medium to high.<br>
> Async JavaScript/TypeScript, CSS/ARIA/Testing Library selectors, Mocha structure, REST fixtures, and WebDriver concepts transfer.<br>
> Testplane browser helpers, `assertView`, config inheritance, context injection, plugin events, html-reporter SQLite, retries, sets/chunks, and visual baselines are framework-specific.<br>
> Keeping page/domain fixtures, API clients, and database lifecycle independent from Testplane reduces lock-in.<br>
> [S5][S8][S10][S13]

#### Custom Harness Burden

> Moderate.<br>
> Testplane supplies browser management, waits, retries, Chrome isolation, visual regression, dev-server readiness, workers, sharding, reports, and diagnostics.<br>
> Repository code remains necessary for multi-service Compose/Flyway/Gin readiness, disposable PostgreSQL state, worker-safe test data and downloads, server-side Yahoo Finance replacement, cross-browser state cleanup, axe integration, backend-log correlation, artifact redaction, and guaranteed teardown.<br>
> [S8][S12][S13][S18]

#### Capability Delivery Tier

> Core OSS: Mocha runner, WebDriver/BiDi sessions, TypeScript transpilation, local browser management, waits/assertions, retries, Chrome contexts, browser sets/workers, screenshots, `assertView`, state commands, command history, last-failed mode, and Time Travel recording.<br>
> Official OSS companions: Testing Library locators, html/json/stat reporters, GitHub Action, chunks, global hooks, repeat/retry tools, Storybook, VS Code, and AI/CLI tooling.<br>
> Chromium/CDP only: request mocking, deep network/CPU emulation, and accessibility snapshots.<br>
> Custom: Compose/Flyway/PostgreSQL lifecycle, backend Yahoo stub, downloads, axe reporting, non-Chromium isolation cleanup, JUnit conversion, and service-log correlation.<br>
> Paid/owned infrastructure: optional real devices, native Safari, and large remote grids.<br>
> [S6][S11][S13][S14][S16][S17]

#### Ai Execution Boundary

> Normal Testplane authoring, browser execution, retries, visual comparison, reporting, and CI require no LLM, model credential, AI egress, or per-run AI fee. `testplane-ai`, Testplane Skill, MCP, and CLI agent features are optional authoring/debugging aids.<br>
> Keep them outside required CI, expose only synthetic data and least-privilege credentials, and commit reviewed deterministic tests that retain non-AI GUI, REPL, selector, and report workflows.<br>
> [S14][S16]

### Evidence And Decision

#### Sources

- [S1] Testplane 9.1.1 npm metadata, dependency versions, Node engine, maintainers, signatures, and provenance, https://registry.npmjs.org/testplane/latest, accessed 2026-08-22.
- [S2] Testplane repository package.json, README, changelog, license, and history, https://github.com/gemini-testing/testplane, accessed 2026-08-22.
- [S3] Testplane GitHub releases, commits, pull requests, and repository metadata, https://github.com/gemini-testing/testplane/releases and GitHub APIs, observed 2026-08-22.
- [S4] npm Downloads API for `testplane`, exact windows 2026-08-15:2026-08-21, 2025-08-15:2025-08-21, 2026-07-22:2026-08-21, and 2025-07-22:2025-08-21, retrieved 2026-08-22.
- [S5] Official installation and TypeScript/ESM documentation, https://testplane.io/docs/quickstart/ and https://testplane.io/docs/basic-guides/typescript-esm/, accessed 2026-08-22.
- [S6] Official browser management and browser-source documentation, https://testplane.io/docs/basic-guides/browsers-overview/, accessed 2026-08-22.
- [S7] Official Testplane 9 announcement, migration guide, and BiDi announcement, https://testplane.io/blog/testplane-9/, https://testplane.io/docs/migrations/how-to-upgrade-testplane-to-9/, and https://testplane.io/blog/support-bidi-protocol/, accessed 2026-08-22.
- [S8] Official browser/system/config, isolation, retries, state, last-failed, and command-history references, https://testplane.io/docs/reference/config/browsers/, https://testplane.io/docs/reference/config/system/, and https://testplane.io/docs/reference/config/last-failed/, accessed 2026-08-22.
- [S9] Official writing, selectors, Testing Library, assertions, waits, and browser command documentation, https://testplane.io/docs/quickstart/writing-tests/ and https://testplane.io/docs/basic-guides/selectors/, accessed 2026-08-22.
- [S10] Official visual testing and assertView configuration documentation, https://testplane.io/docs/basic-guides/visual-testing/ and https://testplane.io/docs/reference/config/browsers/#taking-and-comparing-screenshots, accessed 2026-08-22.
- [S11] Official openAndWait and network-mocking documentation, https://testplane.io/docs/commands/browser/openAndWait/, https://testplane.io/docs/commands/browser/mock/, and https://testplane.io/docs/guides/how-to-intercept-requests-and-responses/, accessed 2026-08-22.
- [S12] Official devServer, setup, and teardown documentation, https://testplane.io/docs/reference/config/dev-server/ and https://testplane.io/docs/basic-guides/setup-and-teardown/, accessed 2026-08-22.
- [S13] Official parallelism, chunks, sharding, and report-merging documentation, https://testplane.io/docs/basic-guides/parallelism/, accessed 2026-08-22.
- [S14] Official Testplane UI, reporters, Time Travel, and debugging documentation, https://testplane.io/docs/html-reporter/overview/, https://testplane.io/docs/basic-guides/reporters/, https://testplane.io/docs/basic-guides/time-travel/, and https://testplane.io/docs/quickstart/running-tests/, accessed 2026-08-22.
- [S15] Official GitHub Actions guide, https://testplane.io/docs/guides/how-to-run-on-github/, accessed 2026-08-22.
- [S16] gemini-testing organization repository catalog and Testplane contributor API, https://github.com/orgs/gemini-testing/repositories and https://api.github.com/repos/gemini-testing/testplane/contributors, observed 2026-08-22.
- [S17] Official accessibility-tree guidance, https://testplane.io/docs/guides/how-to-check-accessibility/, accessed 2026-08-22.
- [S18] Open Asset Allocator repository stack, Makefile, Docker Compose, frontend, and existing integration-test requirements, local repository inspection, observed 2026-08-22.
- [S19] Testplane npm attestation endpoint, https://registry.npmjs.org/-/npm/v1/attestations/testplane@9.1.1, observed through npm metadata 2026-08-22.
- [S20] Official Testplane GitHub Action source and action.yml, https://github.com/gemini-testing/gh-actions-testplane, accessed 2026-08-22.
- [S21] Testplane open/closed issues and pull requests through GitHub APIs, https://api.github.com/repos/gemini-testing/testplane/issues and https://api.github.com/repos/gemini-testing/testplane/pulls, observed 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence for version, package metadata, Node support, maintenance activity, npm counts, browser/config behavior, visual tooling, and CI integration because they come from current official docs and live npm/GitHub APIs.<br>
> Medium confidence for application fit, isolation behavior outside Chrome, ecosystem production use, and supply-chain judgment because they require project-specific execution or unavailable private adoption data.<br>
> Low confidence for TypeScript 6 compatibility, architecture support, resource cost, and the weighted score because no project spike or benchmark was run.

#### Deal Breakers

> No confirmed exclusion-level incompatibility.<br>
> It should be excluded if the project requires tested TypeScript 6 compatibility without a separate E2E tsconfig, local Safari/WebKit on Ubuntu, first-class download objects, uniform cross-browser request interception, fresh browser contexts for every Firefox/Safari test, built-in JUnit/axe output, or zero custom Compose/PostgreSQL harness.<br>
> A project spike exposing stale-element failures, v9 visual instability, Node 24/TS6 type errors, or unsafe database retry behavior would also be a deal breaker.

#### Recommendation

> Viable alternative.<br>
> Testplane is current, feature-rich, and particularly strong where large screenshot suites, retries, sharding, and post-failure visual diagnosis dominate.<br>
> It can test the Open Asset Allocator's HTMX/Navigo/Gin application without changing Parcel.<br>
> It is not the recommended default because current adoption and independent support are much smaller than leading candidates, core behavior depends on a separately maintained WebdriverIO fork, and several project gates remain untested.<br>
> Retain it for a focused spike if integrated open-source visual regression is valued more highly than ecosystem breadth.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `operating_system_support`
- `issue_health`
- `dependency_currency`
- `wrapper_upstream_lag`
- `maintainer_concentration`
- `github_metrics`
- `dynamic_dom_synchronization`
- `resource_usage`
- `security_and_supply_chain`
- `hard_gate_result`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="cucumber-js-with-playwright-or-playwright-bdd"></a>
## 12. Cucumber.js with Playwright or Playwright-BDD

Source result: `Cucumberjs_with_Playwright_or_Playwright-BDD.json`

### Project And Compatibility

#### Implementation Language

> Cucumber.js and playwright-bdd are implemented in TypeScript/JavaScript and run on Node.js.<br>
> Gherkin feature files contain the executable specifications, while TypeScript or JavaScript step definitions call Playwright APIs.<br>
> The direct composition uses @cucumber/cucumber plus Playwright as a browser library; the stronger evaluated composition uses playwright-bdd to generate native @playwright/test cases.

#### Operating System Support

> The Node-based composition runs on Linux, macOS, and Windows.<br>
> With current Playwright, local supported Linux and GitHub-hosted Ubuntu x86-64 are first-class targets; Playwright also documents supported Ubuntu/Debian arm64 targets.<br>
> Official Playwright Ubuntu container images are available.<br>
> Managed Firefox and WebKit require glibc, so Alpine/musl is not a complete three-engine target.

#### License And Governance

> @cucumber/cucumber, playwright-bdd, and the relevant Playwright packages are MIT or Apache-2.0 licensed and require no commercial license.<br>
> Cucumber returned to community ownership and now publishes a collaborative governance process with five core team members, consensus or majority decisions, public funding channels, and David Goss as JavaScript core maintainer. playwright-bdd is a personal repository owned and overwhelmingly maintained by Vitaliy Potapov (vitalets), funded through GitHub Sponsors, with no published multi-maintainer governance or succession policy.<br>
> Microsoft governs and commercially stewards upstream Playwright.<br>
> All licenses permit repository and CI use subject to their notice terms.

#### Installation Model

> Preferred path: npm install -D @playwright/test playwright-bdd, then npx playwright install --with-deps or install selected engines.<br>
> Configure defineBddConfig, generate native tests with npx bddgen, and run npx playwright test.<br>
> Direct Cucumber path: install @cucumber/cucumber, Playwright or @playwright/test, and a TypeScript loader such as tsx; implement World and hooks to launch a browser, create a fresh context/page, attach diagnostics, and close resources.<br>
> The Playwright browser binaries and Linux libraries remain large separate artifacts.<br>
> A version-matched mcr.microsoft.com/playwright image is optional and does not include the project's npm packages.

#### Candidate Scope And Layer

> BDD composition and abstraction layer, not an independent browser engine. playwright-bdd translates Gherkin scenarios into generated Playwright Test files and retains the integrated Playwright runner.<br>
> Direct Cucumber.js instead uses Cucumber as the runner and Playwright only as the browser automation library, requiring more lifecycle and artifact code.

#### Authoring And Async Model

> Stakeholder-facing behavior is expressed as Given/When/Then Gherkin scenarios, scenario outlines, tables, doc strings, tags, and backgrounds.<br>
> Async TypeScript step definitions map text through Cucumber Expressions or regular expressions to awaited Playwright actions and assertions. playwright-bdd supports Playwright-style fixture arguments, Cucumber-style World objects, scoped definitions, scenario/step/worker hooks, and decorators.<br>
> This introduces a three-layer model: feature text, step matching and shared definitions, then Playwright fixtures/actions.<br>
> It improves readable behavioral specifications only when product stakeholders review and own the Gherkin; UI-script-like steps merely add indirection and duplicate lifecycle concepts.

#### Build Pipeline Coupling

> Black-box tests can target the existing Parcel development server or the production Parcel build served by Gin.<br>
> Neither composition requires Vite, frontend source transformation, browser-side instrumentation, or a replacement build pipeline. playwright-bdd adds a test-only generation phase that writes .features-gen files before Playwright Test starts.<br>
> Stable 9.2.0 requires explicit generation; the unreleased 9.3 line adds watch mode, source maps, and generation locks.<br>
> Generated files should be ignored rather than fed into the application build.

#### Testability Instrumentation Required

> No production source transform, injected runtime, special route, or relaxed security policy is required.<br>
> Playwright can use roles, labels, visible text, IDs, and optional stable test IDs after HTMX swaps.<br>
> Application-specific deterministic fixture/reset APIs, health endpoints, and a Yahoo Finance substitute may still be justified, but they are not BDD requirements.<br>
> Gherkin does require a maintained step vocabulary and fixture/World layer in test code; this is significant test-harness instrumentation even though it does not ship in production.

### Maintenance Health

#### Latest Stable Release

> Observed 2026-08-22: @cucumber/cucumber 13.2.1, released 2026-08-04; playwright-bdd 9.2.0, released 2026-06-18; and upstream @playwright/test 1.62.1, released 2026-07-30. playwright-bdd main identified itself as 9.3.2-beta.0, so its watch mode and source-map work were not treated as stable capabilities.

#### Release Cadence

> Cucumber.js is highly active: ten stable releases were published from 2026-04-12 through 2026-08-04, including the 13.0 worker-thread transition and subsequent fixes. playwright-bdd published 8.5.0 in March, 8.5.1 in May, then 9.0.0, 9.1.0, and 9.2.0 in June 2026.<br>
> Its main branch remained active in August, but the interval since the last stable release was about two months and current work was still beta.<br>
> Upstream Playwright releases frequently, so the composed stack requires coordinated package and browser upgrades.

#### Repository Activity

> Both repositories were active on 2026-08-22. cucumber/cucumber-js was pushed that day; its newest commits were largely Renovate/Dependabot updates after human-authored runtime and release work in July and early August. vitalets/playwright-bdd was pushed on 2026-08-21 with watch, lock, source-map, test, and audit work in August, but all ten newest sampled commits were authored by vitalets.<br>
> The wrapper accepts outside contributions, yet no merged pull request was found in the sampled 2026-07-22 through 2026-08-22 window.

#### Dependency Currency

> Current dependency evidence is strong.<br>
> Cucumber.js 13.2.1 supports Node 24 and 26, uses current Cucumber parser/message packages, and develops against TypeScript 6. playwright-bdd 9.2.0 uses TypeScript 6.0.3 in development, declares Node >=20 and @playwright/test >=1.44, explicitly added Playwright 1.60 and 1.61 support, and passed its suite against 1.62 according to the maintainer.<br>
> Stable 9.2 still embeds older Cucumber component generations than Cucumber.js 13, because it translates messages rather than running the current Cucumber.js runtime.

### Community Adoption

#### Npm Downloads

> For the complete seven-day window 2026-08-15 through 2026-08-21, the npm API recorded 2,457,586 downloads of @cucumber/cucumber and 540,138 downloads of playwright-bdd.

#### Ecosystem Usage

> Cucumber has a mature cross-language ecosystem, a common Gherkin/message model, official editors and reports, formatters, tag expressions, and long-standing integrations with test management and CI products. playwright-bdd supplies working examples, Cucumber HTML/JSON/JUnit/message adapters, custom formatter support, Playwright projects/fixtures, BrowserStack/Sauce/Currents/TestDino guides, Allure guidance, and a Discord channel.<br>
> The wrapper's npm volume demonstrates meaningful use, but no independently verified list of organizations running this exact Playwright-BDD composition in production was found.<br>
> Direct Cucumber-Playwright examples are common but vary substantially in lifecycle quality.

#### Community Support

> Cucumber provides extensive official Gherkin and Cucumber.js documentation, GitHub Issues and Discussions, Discord, training material, and a large body of Stack Overflow and third-party guidance. playwright-bdd provides a focused documentation site, examples, GitHub Issues, Discord, migration guides, and Playwright/Cucumber editor integration instructions.<br>
> Support for the exact composition is narrower than either upstream community, and advice often mixes the direct Cucumber runner with the generated Playwright-BDD runner, so version and architecture must be identified before applying examples.

#### Adoption Trend

> Both direct package signals are growing. @cucumber/cucumber rose from 1,485,670 downloads in the aligned Saturday-Friday window 2025-08-16 through 2025-08-22 to 2,457,586 in 2026-08-15 through 2026-08-21, about 65.4 percent. playwright-bdd rose from 97,897 to 540,138 in those aligned seven-day windows, about 451.7 percent, while its repository remained active.<br>
> This is a strong growth signal for package acquisition, not proof of active stakeholder participation or unique production projects.

#### Adoption Metric Normalization

> The metrics are separate and must not be added. @cucumber/cucumber measures the broad JavaScript Cucumber runner across API, integration, and browser-testing uses; it does not identify Playwright compositions. playwright-bdd more directly measures the evaluated generated-runner layer, but downloads can include repeated CI installs, upgrades, mirrors, bots, and transitive use.<br>
> Both comparisons use complete aligned Saturday-Friday UTC windows one year apart.<br>
> Playwright's much larger upstream usage is not attributed to the BDD wrapper.

### Browser And Runtime Coverage

#### Browser Engines

> With Playwright, both compositions can drive managed Chromium, patched Firefox, and patched WebKit.<br>
> Chromium projects can also target installed Google Chrome and Microsoft Edge channels.<br>
> Stock Firefox and native Safari are not driven by Playwright.<br>
> The Gherkin layer does not expand browser coverage; playwright-bdd maps scenarios to normal Playwright projects, while direct Cucumber code must implement its own browser matrix and launch policy.

#### Browser Protocol

> Browser control uses Playwright's client/server protocol and patched browser integrations, not WebDriver or WebDriver BiDi.<br>
> Chromium-only CDP attachment is available but lower fidelity than the normal Playwright connection.<br>
> Cucumber and playwright-bdd operate above this layer and do not make tests protocol-portable.

#### Headless And Headed Modes

> Playwright supports default headless CI execution and headed local debugging for all managed engines. playwright-bdd inherits --headed, --debug, UI Mode, slow motion, Inspector, and Xvfb behavior from Playwright Test.<br>
> Direct Cucumber must expose launch options through profiles, environment variables, or World parameters and wire pause/debug behavior itself.

#### Browser Version Management

> Playwright package versions select tested browser revisions installed by npx playwright install; upgrades normally require reinstalling matching browsers.<br>
> Teams can install only selected engines, set the browser cache path, use branded Chromium channels, or pin a version-matched official image.<br>
> Cucumber.js and playwright-bdd do not manage browsers independently, so package-lock and browser/image versions must be coordinated across three layers.

#### Parallel Browser Support

> playwright-bdd inherits Playwright projects, project selection, workers, fullyParallel, and --shard, and its Cucumber report adapter represents project/browser metadata.<br>
> Direct Cucumber.js supports worker-thread parallelism and native sharding, but browser matrices and report merging require profiles or CI matrices plus custom setup.<br>
> Database-sensitive scenarios can be kept at one Playwright worker or constrained through Cucumber's custom parallel assignment, while browser projects run separately.

#### Mobile Emulation

> Playwright device descriptors and context options provide viewport, screen, user-agent, touch, mobile mode, device scale, locale, and related emulation. playwright-bdd fixtures can derive those options from projects or tags.<br>
> This is responsive emulation in desktop browser engines, not testing on physical Android/iOS devices or native mobile Safari.

#### Real Browser Fidelity

> Managed Chromium can be supplemented with installed Chrome/Edge channels.<br>
> Playwright Firefox is patched and Playwright WebKit is an upstream WebKit build, not branded Firefox or Apple Safari.<br>
> Ubuntu WebKit does not reproduce Safari packaging, codecs, operating-system integration, or release timing.<br>
> Native Safari remains outside this stack and requires a separate Safari/WebDriver or cloud path if it becomes mandatory.

#### Environment Determinism Controls

> The preferred playwright-bdd path exposes Playwright fixtures for locale, timezone, viewport, screen, device scale, user agent, geolocation, permissions, color scheme, contrast, reduced motion, offline state, storage state, and clock control.<br>
> Screenshot assertions can disable animations and mask/style volatile regions.<br>
> Gherkin tags can select fixtures but should not duplicate configuration logic.<br>
> Random-number and backend-time determinism remain application or fixture responsibilities.

### Application Fit

#### Dynamic Dom Synchronization

> Strong browser-layer fit.<br>
> Playwright locators reacquire DOM nodes before each action, actionability checks retry, and web-first assertions poll, which handles HTMX replacement and lazy Handlebars rendering when step definitions assert the post-swap state.<br>
> Gherkin adds no automatic HTMX awareness; step authors must avoid fixed sleeps, generic prose such as 'the page is ready', and retained element handles.<br>
> Reusable business-level steps can centralize correct state-based synchronization, but overly generic steps can hide it.

#### Routing Support

> Navigo History API routing is ordinary browser behavior for Playwright.<br>
> Steps can navigate directly to Gin-served deep links, assert URLs and route-specific DOM, and exercise back/forward/reload. playwright-bdd retains page.waitForURL and web-first URL assertions.<br>
> Feature text should describe route-visible behavior rather than implementation details; cleanup and listener behavior still require explicit scenarios and assertions.

#### Locator Model

> Step definitions have the complete Playwright locator model: role, label, text, placeholder, alt text, title, CSS, and configurable test IDs, with strict single-target actions and reacquisition after HTMX swaps.<br>
> Feature files should not contain CSS selectors or unstable IDs.<br>
> A small domain vocabulary mapped to accessible locators improves stakeholder readability; a large generic click/fill step catalog transfers brittle selectors into Gherkin and weakens both readability and type safety.

#### Form Interaction

> Playwright fill, clear, press, pressSequentially, focus, blur, check, select, keyboard, pointer, and native validation APIs remain available in async steps.<br>
> This can exercise input masks, HTMX validation, focus/blur, hidden-value synchronization, and dynamic allocation/history rows realistically.<br>
> Scenario outlines and data tables express representative input combinations, but complex tables can become a second data DSL and should not replace typed fixture builders for low-level setup.

#### Canvas And Download Support

> Chart.js can be checked through visible labels/legends, source data or chart state evaluated in the page, canvas dimensions, and a narrowly stabilized screenshot assertion. playwright-bdd inherits Playwright's snapshot workflow and download events, including filename, stream/path, saveAs, failure, and payload parsing.<br>
> Direct Cucumber must configure snapshot comparison and download directories itself.<br>
> Gherkin should assert business meaning and exported payload content; pixel-level details belong in step implementation and reports.

#### Network And Api Access

> playwright-bdd steps and fixtures can use APIRequestContext for REST seeding/cleanup and Playwright routing, HAR, request/response events, waitForResponse, fulfill, abort, continue, and fetch.<br>
> Browser-originated Yahoo Finance traffic can be intercepted; server-side Gin calls require a configured backend stub or mock service.<br>
> Direct Cucumber can use the same library APIs but must construct request contexts and teardown manually.<br>
> Feature files should state external-service behavior without embedding fixture endpoints.

#### Same Origin Support

> Both compositions can test the consolidated production Gin origin or Parcel port 8000 with its existing /api proxy and need no CORS relaxation. baseURL is a normal Playwright project/fixture option in playwright-bdd.<br>
> Direct Cucumber can pass the URL through World parameters or environment configuration.

#### Test Isolation

> The preferred composition inherits a fresh Playwright BrowserContext and page per generated scenario/test, isolating cookies, storage, IndexedDB, permissions, and tabs.<br>
> Cucumber-style World can itself be a test-scoped Playwright fixture.<br>
> Direct Cucumber requires explicit Before/After hooks to create and close one context per scenario.<br>
> PostgreSQL remains external to browser isolation; use unique per-worker data, disposable databases/schemas, deterministic reset, or one worker for shared state.

#### External Server Model

> Both paths can point at any independently started application through baseURL, World parameters, or environment variables. playwright-bdd does not require Playwright webServer ownership, so the same features can target Parcel development, production-like Docker Compose, or another reachable deployment.<br>
> Environment-specific behavior should remain in fixtures/configuration rather than conditional Gherkin.

#### Application Lifecycle

> playwright-bdd inherits Playwright webServer, project dependencies, global setup/teardown, and worker/test fixtures, but adds a required bddgen phase before discovery.<br>
> It still cannot declaratively complete this repository's full build, Docker Compose, Flyway, Gin readiness, PostgreSQL reset, backend-log collection, and volume disposal without repository code.<br>
> Direct Cucumber provides BeforeAll/AfterAll and coordinator hooks but needs even more custom process management.<br>
> Cancellation-safe Docker cleanup belongs in an always-running CI step, not only a BDD hook.

#### Visual Regression Workflow

> playwright-bdd retains Playwright Test's first-party screenshot baselines, update workflow, per-project paths, masks, styles, animation disabling, thresholds, pixel limits, and CI diffs.<br>
> Gherkin can expose a concise outcome such as 'the allocation chart matches the approved view', while baseline paths and tolerances remain code/configuration.<br>
> Pin browsers, fonts, viewport, timezone, data, and Chart.js animation completion.<br>
> Direct Cucumber does not receive this integrated snapshot lifecycle from the Playwright library alone and would need a separate comparator/harness.

#### Accessibility Audit Integration

> The preferred path is compatible with the documented open-source @axe-core/playwright integration and Playwright role/name assertions and ARIA snapshots.<br>
> Reusable fixtures or steps can attach axe results to Playwright and Cucumber reports.<br>
> Accessibility scenarios can be stakeholder-readable, but broad 'page is accessible' steps should retain detailed violation evidence and cannot replace manual assessment.<br>
> Direct Cucumber can call axe-core but must assemble reporting itself.

### Reliability

#### Waiting Model

> playwright-bdd preserves Playwright actionability checks, locator assertion polling, navigation/URL waits, event and network waits, expect.poll, and fixture/test timeouts inside each step.<br>
> The BDD layer adds step and hook boundaries but not additional browser waiting.<br>
> Direct Cucumber has step timeouts and hooks but relies on Playwright library calls for waiting.<br>
> Correct steps should await user-visible state or a specific response; fixed delays and broad network-idle waits remain unreliable for HTMX.

#### Flake Controls

> The preferred composition inherits Playwright retries, immediate/isolated retry strategies, assertion/action/navigation/test timeouts, --repeat-each, --last-failed, worker replacement, trace-on-retry, and failOnFlakyTests.<br>
> Tags can apply special timeout/retry behavior.<br>
> Cucumber.js independently supports scenario retries with fresh Worlds, retry tag filters, rerun files, random order, and custom parallel assignment.<br>
> Do not combine both runners' retry models in one path. playwright-bdd issue 395 records that Cucumber HTML formatters still await upstream retry-information support, so Playwright reports should remain the authoritative flake view.

#### Isolation Model

> Each generated Playwright scenario is a normal test with a fresh BrowserContext; fixtures and World are test scoped, while browser and worker fixtures are reused by scope.<br>
> BeforeAll/AfterAll aliases are worker hooks, not once-per-suite hooks.<br>
> The wrapper recommends another package for exactly-once cross-worker caching, which should not be used as a substitute for database isolation.<br>
> Direct Cucumber also creates a fresh World per scenario/retry but needs explicit context teardown.

#### Parallelism Controls

> playwright-bdd uses Playwright worker count, fullyParallel, serial groups, projects, dependencies, and shards.<br>
> Generated scenarios can be tagged/configured, but PostgreSQL safety is still external; start state-mutating E2E execution with one worker, then use per-worker databases or unique data before increasing concurrency.<br>
> Direct Cucumber has an additional setParallelCanAssign API that can prevent simultaneous scenarios by tag, plus worker IDs and native shards.<br>
> The direct path offers finer BDD-specific assignment but requires custom browser and report orchestration.

#### Flake Observability

> Playwright HTML/JSON/JUnit/blob/custom reports classify first-attempt pass, retry pass/flaky, and final failure; traces can retain the first failure or retries and failOnFlakyTests can enforce policy. playwright-bdd maps scenarios and steps into those reports and can also emit Cucumber formats, but Cucumber HTML retry representation is currently incomplete.<br>
> There is no open-source built-in long-term quarantine registry or historical dashboard; tags, JSON/JUnit retention, and repository-owned trend processing are required.

### Diagnostics And Developer Experience

#### Failure Artifacts

> playwright-bdd automatically carries Playwright screenshots, video, traces, console/network trace evidence, attachments, and HTML output into generated scenario results; its Cucumber HTML reporter can embed or externally store screenshots, video, and traces.<br>
> Cucumber JSON skips large attachments by default.<br>
> Direct Cucumber can call this.attach and Playwright tracing/screenshot/video APIs but must implement capture and flushing in hooks.<br>
> Neither path automatically captures Gin, PostgreSQL, Flyway, or Docker logs.

#### Debugging Tools

> The preferred path inherits headed mode, Inspector, page.pause, UI Mode, Trace Viewer, browser DevTools, Node debugging, and the official Playwright VS Code extension.<br>
> Stable 9.2 exposes generated .features-gen tests in the extension; direct run/debug from original .feature gutter locations and source-mapped HTML locations are unreleased 9.3 capabilities as of observation.<br>
> Direct Cucumber supports Node debugging and feature-line execution but lacks Playwright Test UI/fixture integration unless custom tooling is added.

#### Test Generation

> bddgen deterministically generates Playwright tests from existing Gherkin and step definitions; it is a compiler step, not a recorder that discovers meaningful behavior.<br>
> Missing-step snippets and bddgen export help maintain the vocabulary.<br>
> Playwright codegen can record browser actions for implementation, but cannot design stakeholder-quality Gherkin or reusable domain steps. playwright-bdd offers an optional agent skill and prompt-generation feature; generated specifications and code still require human review.

#### Reporters

> playwright-bdd supports all normal Playwright reporters, including console, HTML, blob, JSON, JUnit, GitHub, and custom reporters.<br>
> Its adapter additionally emits Cucumber HTML, JSON, JUnit, message/NDJSON, and some custom Cucumber formatters, with project metadata and automatic attachments.<br>
> Four Cucumber message types are not yet emitted, and some custom formatter constructor objects are fake.<br>
> Shard blob reports can be merged into Cucumber output.<br>
> Direct Cucumber supplies progress/summary/HTML/JSON/JUnit/message/rerun/custom formatters but does not natively supply Playwright's trace-oriented report integration.

#### Documentation Quality

> Cucumber.js documentation is extensive and current for configuration, Gherkin support code, World, hooks, parallelism, sharding, retries, formatters, attachments, ESM, and TypeScript. playwright-bdd documentation is unusually thorough for a community wrapper and covers fixtures, hooks, generation, tags, reports, projects, authentication, IDEs, migration, and examples.<br>
> The stable docs branch and main/upcoming docs can diverge: watch mode and source maps are documented on main but were still 9.3 beta, so users must verify the pinned version.<br>
> Upstream Playwright documentation remains the source for most browser behavior.

#### Local Workflow

> Stable workflow is npx bddgen followed by npx playwright test.<br>
> Developers can select generated tests by file/title/grep/tag/project, use headed/debug/UI mode, inspect reports/traces, and run bddgen --tags.<br>
> The official Playwright extension discovers generated tests, while Cucumber extensions provide Gherkin completion/navigation.<br>
> Stable 9.2 requires regeneration after feature changes and generally exposes generated files for run/debug; 9.3 beta adds watch, locks, source maps, and original-feature gutter actions.<br>
> This is measurably more ceremony than direct Playwright Test.

#### Failure Log Correlation

> Playwright timestamps actions, BDD steps, assertions, console and network entries, and attachments in one scenario/test trace; Cucumber reports retain feature/scenario/step identity.<br>
> Custom fixtures can attach browser errors and failed responses with testInfo.<br>
> The workflow must separately timestamp and collect Gin, Flyway, PostgreSQL, and Compose logs using a shared run, shard, worker, and scenario identifier.<br>
> Generated-file paths and original feature paths should both be preserved until stable source maps are adopted.

#### Artifact Data Exposure

> Screenshots, videos, traces, HTML/JSON/message reports, HAR/network bodies, console output, storage state, downloads, feature example tables, and World attachments can expose credentials and portfolio data. playwright-bdd Cucumber reports can embed binary artifacts; external attachments reduce report size, not sensitivity.<br>
> Its optional Fix with AI creates a prompt attachment containing the error, Gherkin, code, and ARIA snapshot but does not itself call a model.<br>
> Use synthetic data, skip unnecessary attachments, sanitize logs, restrict retention/access, never commit auth state, and keep AI prompts and external model use disabled in CI.

### Github Actions Fit

#### Official Ci Support

> There is no dedicated official Cucumber-Playwright or playwright-bdd GitHub Action.<br>
> The preferred path follows official Playwright guidance: setup Node, npm ci, install version-matched browsers with --with-deps or use a pinned image, run bddgen, then run Playwright Test and upload reports. playwright-bdd maintains its own Linux/Windows CI and passed Playwright 1.62 in GitHub Actions.<br>
> Cucumber documents CI sharding generically.<br>
> This is viable on ubuntu-latest but adds one generation step and another package compatibility boundary.

#### Browser Caching

> Use setup-node's npm cache for lockfile dependencies.<br>
> Playwright does not generally recommend caching browser binaries because restoration can cost as much as download and Linux packages are separate; installing selected engines or pinning the official image is simpler.<br>
> If cached, key by OS, architecture, and exact Playwright version.<br>
> Gherkin and generated test files are small and should normally be regenerated, not shared as mutable cross-job cache state.

#### Artifact Integration

> Low effort on playwright-bdd: configure Playwright/Cucumber reporters and capture paths, then upload playwright-report, cucumber-report, blob reports, and selected test-results with actions/upload-artifact under an always/failure policy.<br>
> Cucumber HTML can externalize large attachments and blob reports can merge across shards.<br>
> Backend and Compose logs still need explicit workflow collection.<br>
> Apply short retention and restricted access because both feature examples and browser artifacts may contain portfolio data.

#### Sharding And Matrix Support

> playwright-bdd inherits Playwright --shard and browser projects/matrices.<br>
> Each shard emits a blob report; an aggregation job downloads blobs and runs playwright merge-reports with the Playwright config to produce Playwright or Cucumber reports. bddgen must use the same config and should run deterministically before each shard.<br>
> Direct Cucumber 13 supports shard INDEX/TOTAL plus in-process parallel workers, but cross-job report aggregation is less integrated.<br>
> Database-mutating shards require separate databases/schemas or must remain unsharded.

#### Container Compatibility

> The stack can run directly on ubuntu-latest beside this repository's Docker Compose services, or in a version-matched Playwright Ubuntu image.<br>
> If the test runner is containerized, the application must be reachable through service DNS or a host-gateway mapping because localhost is container-local; Chromium benefits from --init and adequate shared memory.<br>
> Cucumber/playwright-bdd add no browser container.<br>
> Pin Node packages, image/browser versions, and architecture together.

#### Failure Cleanup

> Playwright Test fixtures close pages, contexts, browsers, and owned webServer processes; direct Cucumber must close them in After/AfterAll even after failures.<br>
> Hard timeout or cancellation can interrupt BDD hooks.<br>
> Docker Compose, PostgreSQL volumes, generated files, and backend processes therefore need job-level cleanup guarded by if: always(), unique Compose project names, disposable volumes, and explicit process termination.<br>
> Preserve logs before docker compose down.

### Cost And Risk

#### Open Source Completeness

> All required baseline capabilities are open source: Gherkin parsing and execution, Playwright browser automation/Test runner, Chromium/Firefox/WebKit builds, fixtures, waiting, retries, workers, sharding, screenshots, visual snapshots, video, traces, HTML/JSON/JUnit/message reports, axe integration, and local debugging.<br>
> No paid service is required.<br>
> The direct Cucumber path needs more custom code to reach the same diagnostics, while playwright-bdd delivers them through Playwright Test.

#### Optional Cloud Dependency

> No cloud is mandatory.<br>
> Cucumber Reports publishing, BrowserStack/Sauce grids, Currents/TestDino dashboards, Allure services, commercial visual review, and model providers are optional.<br>
> Local Playwright/Cucumber HTML, trace viewing, committed visual baselines, GitHub artifacts, JUnit/JSON, and repository-owned trend processing provide non-cloud alternatives, though they do not supply a turnkey long-term analytics dashboard.

#### Migration Cost

> Moderate to high.<br>
> Feature text is portable in principle, especially when it describes domain behavior, and TypeScript/Playwright actions remain familiar.<br>
> In practice the suite depends on Gherkin structure, step-expression matching, shared vocabulary, tags, generated tests, playwright-bdd fixtures/hooks/report translation, Playwright locators, and visual baselines.<br>
> Moving to direct Playwright requires flattening scenarios into tests; moving to direct Cucumber requires replacing Playwright Test fixtures and artifacts.<br>
> Generic reusable steps can create hidden coupling and make migration or refactoring harder than duplicated explicit tests.

#### Custom Harness Burden

> High relative to direct Playwright Test. playwright-bdd supplies generation, Gherkin matching, Playwright fixtures, browser lifecycle, reports, and attachments, but the repository must maintain feature conventions, step definitions, fixture/World typing, generation scripts, source/generated-file policy, version coordination, and BDD-specific IDE guidance.<br>
> It also still needs Compose/Flyway/Gin readiness, deterministic PostgreSQL setup/reset, Yahoo Finance replacement, backend log collection, artifact redaction, and cancellation-safe teardown.<br>
> Direct Cucumber adds custom browser/context lifecycle, traces, video, visual baselines, and report wiring on top of that.

#### Capability Delivery Tier

> Core upstream Cucumber: Gherkin, expressions, World, hooks, tags, retries, parallelism, sharding, attachments, and Cucumber reports.<br>
> Community wrapper playwright-bdd: generated Playwright tests, fixture integration, BDD hooks/tags, project mapping, and Cucumber report adapters.<br>
> Core Playwright Test: browser management, locators, waiting, isolation, API/network controls, projects, retries, workers, sharding, snapshots, downloads, traces, video, screenshots, reports, UI Mode, Inspector, and codegen.<br>
> Separate open-source integration: axe-core and optional Allure/global-cache tools.<br>
> Custom repository code: application/database lifecycle, service mocking, log correlation, deterministic Chart.js checks, artifact policy, and stakeholder vocabulary governance.<br>
> Paid cloud: none required.

#### Ai Execution Boundary

> Deterministic Gherkin parsing, bddgen generation, Playwright execution, assertions, retries, and CI require no LLM, AI credential, inference cost, or model egress. playwright-bdd's agent skill is optional authoring assistance.<br>
> Fix with AI only generates a local report attachment until a human sends it to a model, but that prompt can contain errors, source, Gherkin, and an ARIA snapshot.<br>
> Keep AI options disabled in required CI and require all feature and step changes to run without AI.

### Evidence And Decision

#### Sources

- Official Cucumber.js repository, releases, changelog, package metadata, issues, and activity: https://github.com/cucumber/cucumber-js, observed 2026-08-22.
- Official Cucumber.js installation, configuration, profiles, ESM, and transpiling documentation: https://github.com/cucumber/cucumber-js/tree/main/docs, accessed 2026-08-22.
- Official Cucumber.js World, hooks, step definitions, attachments, and timeout documentation: https://github.com/cucumber/cucumber-js/tree/main/docs/support_files, accessed 2026-08-22.
- Official Cucumber.js parallelism, sharding, retry, rerun, and formatter documentation: https://github.com/cucumber/cucumber-js/blob/main/docs/parallel.md, https://github.com/cucumber/cucumber-js/blob/main/docs/sharding.md, https://github.com/cucumber/cucumber-js/blob/main/docs/retry.md, and https://github.com/cucumber/cucumber-js/blob/main/docs/formatters.md, accessed 2026-08-22.
- Official Cucumber.js deprecation policy: https://github.com/cucumber/cucumber-js/blob/main/docs/deprecations.md, accessed 2026-08-22.
- Cucumber community governance: https://github.com/cucumber/governance, accessed 2026-08-22.
- Cucumber 2025 year in review and community-ownership/funding statement: https://news.cucumber.io/archive/cucumber-in-2025-year-in-review/, accessed 2026-08-22.
- Official playwright-bdd repository, package, changelog, releases, issues, contributors, and activity: https://github.com/vitalets/playwright-bdd, observed 2026-08-22.
- Official playwright-bdd overview, installation, CLI, and configuration: https://vitalets.github.io/playwright-bdd/ and corresponding docs in https://github.com/vitalets/playwright-bdd/tree/main/docs, accessed 2026-08-22.
- Official playwright-bdd fixtures, Cucumber-style steps, and hook lifecycle documentation: https://github.com/vitalets/playwright-bdd/tree/main/docs/writing-steps, accessed 2026-08-22.
- Official playwright-bdd Playwright and Cucumber reporter documentation: https://github.com/vitalets/playwright-bdd/tree/main/docs/reporters, accessed 2026-08-22.
- Official playwright-bdd IDE, UI Mode, authentication, source-map, watch-mode, and AI prompt guides: https://github.com/vitalets/playwright-bdd/tree/main/docs/guides, accessed 2026-08-22; source maps/watch mode were treated as unreleased 9.3 functionality.
- playwright-bdd issue 396 and maintainer confirmation of Playwright 1.62 CI compatibility: https://github.com/vitalets/playwright-bdd/issues/396, observed 2026-08-22.
- playwright-bdd issue 395 concerning Cucumber HTML retry information: https://github.com/vitalets/playwright-bdd/issues/395, observed 2026-08-22.
- npm registry metadata for @cucumber/cucumber 13.2.1 and playwright-bdd 9.2.0: https://registry.npmjs.org/%40cucumber%2Fcucumber/latest and https://registry.npmjs.org/playwright-bdd/latest, observed 2026-08-22.
- npm downloads API for 2026-08-15 through 2026-08-21 and aligned 2025 windows for @cucumber/cucumber and playwright-bdd: https://api.npmjs.org/downloads/, observed 2026-08-22.
- Official Playwright documentation for browsers, locators, actionability, fixtures, projects, isolation, retries, sharding, reports, traces, snapshots, downloads, network/API controls, emulation, clock, Docker, and CI: https://playwright.dev/docs/, accessed 2026-08-22.
- Microsoft Playwright repository and @playwright/test 1.62.1 release evidence: https://github.com/microsoft/playwright and https://registry.npmjs.org/%40playwright%2Ftest/latest, observed 2026-08-22.
- Open Asset Allocator repository evidence: .nvmrc, src/main/web-static/package.json, tsconfig.json, .proxyrc.js, HTMX/Navigo/Handlebars/Chart.js frontend code, Gin server, Docker Compose, Flyway, Makefile, lifecycle scripts, and GitHub Actions workflows, inspected in the companion Playwright Test research on 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence in licenses, versions, package requirements, Node/TypeScript development versions, npm downloads, GitHub snapshots, Gherkin lifecycle, generation architecture, Playwright inheritance, report formats, and documented CI capabilities because they use first-party repositories, documentation, registry data, and live APIs.<br>
> Medium confidence in application-fit, adoption interpretation, issue health, wrapper lag, and governance risk.<br>
> Low confidence in resource use, complete supply-chain status, final weighted score, and repository-specific behavior because no spike or benchmark was executed.

#### Deal Breakers

> No confirmed browser or runtime incompatibility.<br>
> Exclude the composition if stakeholders will not actively review or co-own Gherkin, because then step indirection and duplicate lifecycle concepts have cost without the stated benefit.<br>
> Other exclusion triggers are refusal to depend on a high-bus-factor wrapper, a requirement for native Safari on Ubuntu, zero generated test files, one-layer debugging, or zero custom application/database lifecycle code.<br>
> A failed Node 24/TypeScript 6/application spike or unacceptable wrapper lag on a pinned Playwright release would also be a deal breaker.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `roadmap_and_deprecation_risk`
- `wrapper_upstream_lag`
- `maintainer_concentration`
- `github_metrics`
- `resource_usage`
- `security_and_supply_chain`
- `hard_gate_result`
- `project_fit_score`
- `recommendation`
- `empirical_project_spike_result`

<a id="robot-framework-browser"></a>
## 13. Robot Framework Browser

Source result: `Robot_Framework_Browser.json`

### Project And Compatibility

#### Implementation Language

> Robot Framework Browser 20.4.0 has a Python library front end and a TypeScript/Node.js wrapper that communicates over local gRPC and drives the Playwright Node module.<br>
> Tests are normally authored as Robot Framework keyword tables in `.robot` and `.resource` files; reusable extensions can be written in Python or JavaScript.<br>
> Required runtimes are Python 3.10 or newer and Robot Framework 7.1.1 or newer, plus either supported Node.js 22, 24, or 26 or the optional BrowserBatteries package with bundled Node.js 24.19.0.

#### Operating System Support

> Local Linux and GitHub-hosted Ubuntu are supported.<br>
> The upstream CI runs Browser on `ubuntu-latest`, including Python 3.10-3.14 and Node 22, 24, and 26 combinations, and tests BrowserBatteries on x86-64 and Ubuntu ARM64.<br>
> BrowserBatteries wheels cover Linux x64 and arm64 with glibc 2.28 or newer, Windows x64, and macOS x64 and arm64; there is no musl/Alpine or Windows ARM wheel.<br>
> The plain Node.js installation can cover additional platforms where Python, a supported Node version, and Playwright work.<br>
> Official Docker images are published for linux/amd64 and linux/arm64/v8.

#### License And Governance

> Apache License 2.0, permitting use, modification, and redistribution in this repository subject to the license notice and patent terms.<br>
> MarketSquare, the Robot Framework community organization, owns the public repository.<br>
> The official community page names Tatu Aalto and Rene Rohner as the two current core maintainers, lists three core alumni, and reports funding from Robocorp, the Robot Framework Foundation, and imbus.<br>
> Development, issues, and releases occur publicly; this is a community project with funded maintenance rather than a vendor-backed commercial product, and no paid license is required.

#### Installation Model

> Recommended minimal-runtime route: create a Python virtual environment, install `robotframework-browser[bb]`, and run `rfbrowser install` to install version-matched browsers.<br>
> The alternative is `pip install robotframework-browser` followed by mandatory `rfbrowser init`, which uses npm to install the locked Node dependencies and Playwright browsers.<br>
> All three engines are installed by default, normally adding more than 700 MB, but `chromium`, `firefox`, or `webkit` can be selected. `--skip-browsers` plus `PLAYWRIGHT_BROWSERS_PATH` supports externally managed binaries.<br>
> Version-pinned Docker images include Python, Node, Robot Framework, Browser, system libraries, and browsers.

#### Candidate Scope And Layer

> Keyword-driven E2E framework composition: Robot Framework supplies the runner, suite model, keywords, setup/teardown, reports, and extension APIs; Browser supplies the Playwright-backed web automation library.<br>
> It is a higher-level abstraction over Playwright API, not Playwright Test, WebDriver, or only a browser driver.

#### Authoring And Async Model

> Tests use synchronous-looking Robot Framework keyword tables, variables, user keywords, templates, setup/teardown, and tags.<br>
> Browser crosses Python/gRPC/Node internally and hides JavaScript promises. `Promise To` and `Wait For` explicitly overlap any Browser keyword when a listener must start before an action, such as request, navigation, download, or dialog waiting.<br>
> Python and JavaScript plugins can add keywords.<br>
> This model is business-readable but is framework-specific and does not reuse the repository's TypeScript types or native async/await conventions.

#### Build Pipeline Coupling

> Black-box suites can target the existing Parcel development server on port 8000 or the production Parcel output served by Gin.<br>
> Browser does not require Vite, frontend transforms, source instrumentation, or a replacement build pipeline.<br>
> Its Python and Node wrapper dependencies can live in a dedicated E2E environment.<br>
> The application can therefore keep the current npm package, TypeScript module configuration, Parcel proxy, and production build unchanged.

#### Testability Instrumentation Required

> No Browser runtime injection, build hook, relaxed security policy, or source transform is required.<br>
> Role, accessible-name, visible-text, ID, CSS, XPath, and test-ID selectors work against ordinary HTML; stable test IDs are optional where dynamic rows lack unique semantic identity.<br>
> Deterministic database reset, readiness, backend-side Yahoo Finance replacement, and fixture seeding may require test harness facilities or test-only APIs.<br>
> Browser-side request mocking is not a 20.4.0 core keyword and requires a JavaScript extension or custom plugin.

### Maintenance Health

#### Latest Stable Release

> 20.4.0, published on August 19, 2026.<br>
> The release supports Python 3.10+, Node.js 22/24/26, Robot Framework 7.1.1+, and was tested with Playwright 1.62.1; BrowserBatteries ships Node.js 24.19.0.

#### Release Cadence

> Very active.<br>
> Stable releases 20.1.0, 20.2.0, 20.3.0, and 20.4.0 were published on July 22, August 1, August 7, and August 19, 2026, following 20.0.0 on June 7 and many 19.x releases earlier in 2026.<br>
> Green main-branch commits also produce replaceable nightly wheels.<br>
> The cadence delivers browser and security updates quickly, but frequent dependency and minimum-version changes require pinned upgrades and release-note review.

#### Repository Activity

> Current activity is strong.<br>
> The main repository was pushed on August 21, 2026, one day before observation.<br>
> The ten newest observed commits dated August 19-21 include the 20.4.0 release, documentation, GitHub Actions reporting, and Python-extension work from Tatu Aalto and Rene Rohner.<br>
> Recent issues were closed into the release, and the upstream CI runs a broad Python, Node, Robot Framework, OS, architecture, package-install, and Docker matrix.

#### Dependency Currency

> Current.<br>
> Browser 20.4.0 pins Playwright 1.62.1, uses TypeScript 6.0.3 in development, supports Node 22/24/26, supports Python 3.10+, and constrains Robot Framework to 7.1.1 through below 9.<br>
> The CI directly tests Node 24 and current Python versions on Ubuntu, Windows, and macOS.<br>
> Recent releases updated gRPC, protobuf, tar, esbuild, Playwright, and the supported Node lines in response to browser, compatibility, and security changes.

### Community Adoption

#### Npm Downloads

> Not applicable as an npm adoption metric: users install the E2E framework as the PyPI package `robotframework-browser`, while npm packages are internal wrapper dependencies.<br>
> The comparable registry metric is 163,806 PyPI downloads without mirrors during the exact seven-day period August 15-21, 2026; PyPI Stats also reported 872,835 downloads in its rolling last-month measure on August 22.

#### Ecosystem Usage

> Material adoption exists within the Robot Framework ecosystem: more than 160,000 weekly PyPI downloads, a six-year release history, Pabot parallel execution compatibility, Robot Framework's API/database/process libraries, official Docker images, and a MarketSquare extensions repository.<br>
> The extension collection includes axe-core checks, basic image comparison, network throttling, request mocking/HAR replay, highlighting, and direct Playwright page methods.<br>
> Robocorp, the Robot Framework Foundation, and imbus have funded development, but funding is not the same as verified current production adoption and no comprehensive current organization list was found.

#### Community Support

> The project provides a newly redesigned official guide, generated versioned keyword reference, release notes, examples, Slack `#browser`, the searchable Robot Framework forum, and an active GitHub issue tracker.<br>
> Robot Framework itself adds a large body of general documentation and community material.<br>
> The dedicated Browser repository has no GitHub Discussions, and search results can still surface older installation documentation with obsolete Python and Node requirements, so users must verify the pinned version.

#### Adoption Metric Normalization

> The primary adoption count is the direct PyPI package `robotframework-browser`, all versions, downloads without mirrors, for August 15-21, 2026.<br>
> It covers the Browser library rather than Robot Framework, Playwright, BrowserBatteries, Pabot, or extension packages.<br>
> It may include repeated CI installs, upgrades, caches, bots, and automated analysis and does not count distinct people or projects.<br>
> GitHub metrics cover Browser source and development, while the official 207-person all-contributors list includes ideas, reports, support, funding, documentation, testing, and code rather than only commit authors.

### Browser And Runtime Coverage

#### Browser Engines

> First-class managed Chromium, Firefox, and WebKit through Playwright. `New Browser` can launch the three engines, and the `channel` option can select installed Chromium-family channels such as Google Chrome or Microsoft Edge where supported.<br>
> Playwright-managed Firefox and WebKit are patched/tested builds rather than arbitrary stock Firefox or Apple Safari.<br>
> A browser server or Chromium CDP endpoint can also be connected.

#### Browser Protocol

> Robot Framework calls the Python Browser library, which sends gRPC messages to a Node.js wrapper that drives Playwright through Playwright's framework-specific transport.<br>
> It does not use W3C WebDriver or WebDriver BiDi. `Connect To Browser` supports a Playwright WebSocket endpoint and Chromium CDP attachment, but normal execution uses Playwright's higher-fidelity protocol and therefore inherits Playwright version and browser-build coupling.

#### Headless And Headed Modes

> `New Browser` is headless by default for CI and supports `headless=False` for local observation. `Open Browser` is intentionally headed for experiments.<br>
> Slow motion, DevTools for Chromium, headed Docker execution through Xvfb, and ordinary headless official-image execution are supported.

#### Browser Version Management

> `rfbrowser init` or `rfbrowser install` installs browser revisions matched to the Playwright package bundled with the Browser release.<br>
> Selected engines can be installed, and `PLAYWRIGHT_BROWSERS_PATH` can move or share the cache.<br>
> Upgrades require cleaning/reinitializing the Node side and reinstalling matching browsers; Browser and BrowserBatteries versions are intended to match.<br>
> Exact Docker tags should be pinned.<br>
> The 20.4.0 official image uses a Playwright 1.62.0 base while Browser pins 1.62.1, and the official docs explicitly warn that this manually maintained patch-level relationship should be checked.

#### Parallel Browser Support

> Multiple browsers, contexts, and pages can exist in one Robot run and can be switched by ID.<br>
> Browser does not provide Playwright Test projects or native browser matrices.<br>
> Browser selection can be parameterized with Robot variables and GitHub matrices, while suite/test process parallelism comes from the separate Pabot package.<br>
> Each Pabot worker normally starts its own Node/Playwright process; sharing one Node process is documented as experimental.

#### Mobile Emulation

> `Get Device` exposes Playwright device descriptors, and `New Context` accepts viewport, screen, user agent, device scale factor, mobile mode, touch, and default browser type. `Tap`, geolocation, permissions, orientation-sized viewports, locale, timezone, color scheme, and reduced-motion checks support responsive scenarios.<br>
> This is desktop-engine emulation, not physical Android/iOS hardware, and there is no native pinch or multi-touch support.

#### Real Browser Fidelity

> Managed engines provide reproducible cross-engine behavior, and installed Chrome/Edge channels improve branded Chromium fidelity.<br>
> Managed Firefox is not arbitrary stock Firefox.<br>
> Playwright WebKit shares an engine family with Safari but is not macOS or iOS Safari and does not reproduce Safari packaging, OS integration, release timing, codecs, on-screen keyboards, or device performance.<br>
> Native Safari or real-device requirements need a separate Apple or remote-grid solution.

#### Environment Determinism Controls

> Context controls include viewport, screen, device scale, user agent, locale, timezone, geolocation, permissions, color scheme, forced colors, reduced motion, offline mode, service-worker policy, headers, proxy, storage state, and JavaScript enablement.<br>
> Core clock keywords set, pause, resume, or advance page time and timers.<br>
> Tests can inject CSS or JavaScript to disable animations and control application state.<br>
> Browser does not supply a general seeded replacement for application/backend randomness or `Math.random`; those remain application fixture responsibilities.

### Application Fit

#### Dynamic Dom Synchronization

> Strong fit for HTMX swaps and asynchronous Handlebars rendering.<br>
> Playwright actionability waits before interactions, Browser getter assertions re-read values until they pass, `Wait For Condition` covers DOM/value state, and ordinary selectors and returned element references are re-resolved on use after node replacement.<br>
> Tests should wait for route-specific text, element state, response, or an application readiness condition rather than sleep.<br>
> A `Wait For Function` selector is resolved once, so it should not retain a node that HTMX will replace.

#### Routing Support

> Direct navigation, URL assertions, `Go Back`, `Go Forward`, page history, and `Wait For Condition Url` cover Navigo routes and history changes. `Wait For Navigation` handles full navigations but explicitly does not detect fragment-only changes, so hash/SPA transitions need retrying `Get Url` assertions or route-DOM conditions.<br>
> Deep links work when the selected Gin or Parcel topology serves the SPA fallback; that server behavior must be part of the application spike.

#### Locator Model

> Role and accessible-name selectors, text, test IDs, IDs, CSS, XPath, filters, chained selectors, iframes, and open shadow DOM are supported.<br>
> Strict mode is enabled by default and rejects ambiguous matches.<br>
> Role selectors and `Get Element By Role` support accessible UI targeting.<br>
> Element references are selector strings that re-resolve rather than stale DOM handles, which is useful after HTMX replacement.<br>
> Prefer role/label-like selectors and visible text, using stable test IDs for dynamic rows where semantic identity is insufficient.

#### Form Interaction

> Core keywords cover fill, type, clear, secrets, keyboard keys, key sequences, check/uncheck, option selection, focus, mouse, touch, file upload, and dialogs.<br>
> Actionability checks preserve realistic enabled, visible, stable, and event-receiving behavior.<br>
> Tab or focusing another element can trigger blur, and `Get Element States` can assert focused/defocused state.<br>
> JavaScript/property getters can inspect hidden-value synchronization and native validation when visible messages are insufficient.<br>
> This is suitable for the repository's masks, dynamic rows, HTMX validation, and focus/blur behavior.

#### Canvas And Download Support

> Chart.js canvas can be checked through visible legends/data, `Evaluate JavaScript`, canvas dimensions/pixels, and `Take Screenshot`; `Wait For Function` can poll for drawn pixels.<br>
> Core has no semantic Chart.js adapter or visual baseline assertion, so data-level custom keywords should remain primary.<br>
> Downloads have dedicated promise/wait, state, cancellation, save-path, URL download, and upload keywords, allowing filename and payload checks through Robot/Python filesystem libraries.

#### Network And Api Access

> Core `HTTP` sends fetch requests in the active page/context and can seed or inspect same-origin REST fixtures; `Wait For Request` and `Wait For Response` observe matching traffic and return metadata/body.<br>
> Context HAR recording is available.<br>
> Version 20.4.0 does not expose core route fulfill/abort/modify keywords; issue 5120 requests that family.<br>
> The community `mockUrl` JavaScript extension provides fulfill, block, record-HAR, and HAR replay, but its compatibility and maintenance must be validated.<br>
> Browser-originated Yahoo traffic can be replaced there; Gin-originated Yahoo Finance calls require a backend mock service or application fixture because the browser cannot intercept server-to-server traffic.

#### Same Origin Support

> Normal E2E operation needs no CORS relaxation.<br>
> Tests can use the consolidated production Gin origin where static content and `/api` share host/port, or target Parcel port 8000 where the existing proxy keeps browser API traffic same-origin.<br>
> The context `baseURL` resolves relative navigation and request/response matchers.<br>
> The `HTTP` keyword uses page fetch and therefore follows browser same-origin/CORS rules by design.

#### Test Isolation

> Browser contexts isolate cookies, local storage, permissions, cache, and pages, and the default `auto_closing_level=TEST` closes contexts/pages created by each test.<br>
> Explicit `New Context` per test provides a cheap clean session; saved storage state can be reused within a run but contains credentials.<br>
> Browser isolation does not isolate PostgreSQL.<br>
> State-mutating parallel workers need unique data or disposable databases/schemas, while dependent flows need serial ordering and unconditional cleanup.

#### External Server Model

> Suites can target any independently started application through variables and `New Context baseURL` or absolute URLs.<br>
> Browser does not require ownership of the server process.<br>
> The same keyword suite can therefore address Parcel development, host-run Docker Compose, a production-like Gin container, or another reachable test environment.

#### Application Lifecycle

> Robot Framework provides suite/test setup and teardown and can call Python, Process, OperatingSystem, or custom keywords, but Browser has no built-in equivalent of Playwright Test `webServer`.<br>
> This repository needs an outer script or resource library to build Parcel, start Docker Compose, wait for PostgreSQL and Flyway, poll Gin/frontend readiness, seed/reset data, capture logs, and always tear down services and volumes.<br>
> GitHub cancellation-safe cleanup must also exist outside ordinary Robot teardown.

#### Visual Regression Workflow

> Core Browser takes page/element screenshots with path, format, quality, clip, full-page, scale, and style controls, but it has no first-party baseline creation/update/review, masking, threshold policy, per-browser snapshots, or CI diff reporter.<br>
> The MarketSquare extension repository contains a basic Pixelmatch JavaScript example that writes a diff and accepts a threshold, but it requires additional npm modules and custom Robot logic and was not verified against Browser 20.4.0.<br>
> A production workflow would need custom baseline naming/review, deterministic fonts/data/animations, artifact upload, and selective Chart.js use.

#### Accessibility Audit Integration

> Core role locators and `Get Aria Snapshot` support accessible-tree inspection, but an ARIA snapshot is not a WCAG audit and it is excluded from retrying `Wait For Condition`.<br>
> The MarketSquare extensions repository contains an axe-core/Playwright example and HTML reporter under Apache-2.0, but its manifest still pins old axe package ranges and current 20.4.0 compatibility was not demonstrated.<br>
> A maintained project-owned extension or Python/JavaScript keyword can run current axe-core and attach JSON/HTML results without a paid service.

### Reliability

#### Waiting Model

> Actions inherit Playwright checks for attachment, visibility, stability, event reception, enabled state, and related actionability.<br>
> Browser getter assertions poll values for `retry_assertions_for`, bounded by the general element `timeout`. `Wait For Condition`, `Wait For Elements State`, `Wait For Function`, load-state, navigation, request/response, and promise keywords provide explicit synchronization.<br>
> For HTMX and Navigo, state-based DOM/URL/response assertions are preferable to fixed sleeps and broad network-idle waits.

#### Flake Controls

> Timeout and assertion-retry settings have global, suite, and test scopes.<br>
> Robot Framework can rerun failed tests with `--rerunfailed` and merge results with Rebot; the separate RetryFailed listener can retry automatically, and `Wait Until Keyword Succeeds` retries a keyword.<br>
> Tags, random ordering, repeated CI commands, Pabot controls, and traces aid reproduction.<br>
> These mechanisms are less integrated than Playwright Test retry strategy: the repository must define retry policy, preserve first-attempt evidence, and prevent retries from hiding shared-database defects.

#### Isolation Model

> The preferred model is one fresh context per test inside a reused browser, with default test-level automatic closing.<br>
> Contexts separate browser session state, while multiple pages in a context intentionally share it.<br>
> Suite setup can deliberately create longer-lived state.<br>
> Pabot workers are separate Robot and normally separate Node/browser processes.<br>
> PostgreSQL and external services remain shared unless the project allocates per-worker resources or serializes stateful suites.

#### Parallelism Controls

> Robot Framework core runs serially; the separate Pabot package supplies process counts, suite/test splitting, ordering, locks through PabotLib, and merged output.<br>
> GitHub browser/job matrices provide coarse sharding.<br>
> Browser can share a standalone Node process experimentally, but normal worker isolation is safer.<br>
> Start database-mutating E2E work at one process, or use Pabot ordering/locks, until each worker receives a unique PostgreSQL database/schema and fixture namespace.

#### Flake Observability

> Robot output XML/JSON, log HTML, report HTML, Pabot worker files, debug logs, and merged rerun results can be retained and processed locally.<br>
> The framework can identify final pass/fail and preserve retry runs if the workflow keeps both outputs, but it has no built-in first-attempt-versus-retry trend dashboard, quarantine registry, historical flake store, or hosted-free analytics UI.<br>
> Repository scripts or a reporting package must classify retries and maintain trends.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Robot Framework produces `log.html`, `report.html`, and output XML or JSON.<br>
> Browser adds configurable on-failure screenshots, `playwright-log.txt`, Playwright traces with DOM snapshots, context video, HAR recording, console/page-error getters, and ordinary attachments/files in the output directory.<br>
> Traces can be enabled globally and passed traces auto-deleted, but explicit context creation and closing are required.<br>
> Backend, PostgreSQL, Flyway, and Docker logs are not captured automatically.

#### Debugging Tools

> Available tools include headed mode, slow motion, Chromium DevTools, detailed Robot and Playwright logging, `Open Browser`, `Record Selector`, element highlighting, presenter mode, Node inspector flags, and `rfbrowser show-trace`/Playwright Trace Viewer.<br>
> This is useful but lacks Playwright Test Inspector, UI Mode, fixture view, watch runner, and integrated source-step debugging.<br>
> Python/Robot IDE support comes from external tools rather than Browser core.

#### Test Generation

> `Record Selector` helps interactively identify an element and return a selector, and Robot Framework tooling can scaffold suites, but Browser has no core recorder that generates maintainable end-to-end Robot tests comparable to Playwright codegen.<br>
> Tests, domain keywords, assertions, HTMX waits, fixtures, and cleanup must be authored and reviewed manually.

#### Reporters

> Robot Framework core supplies console output, detailed self-contained HTML log/report, machine-readable output XML and JSON, xUnit/JUnit-compatible output, listener APIs, and Rebot combination/merge.<br>
> Pabot merges process results.<br>
> The separate `robotframework-ghareports` tool can create GitHub job summaries or pull-request comments, and other Robot ecosystem reporters exist.<br>
> There is no Browser-native blob/shard reporter or Playwright Test HTML report.

#### Documentation Quality

> Good and improving.<br>
> The new official site provides current architecture, installation, waiting, assertions, selectors, mobile, Docker, and debugging guidance; the generated 20.4.0 reference documents 152 keywords and versioned releases.<br>
> Source links and executable examples are extensive.<br>
> Risks are that older Robot Framework guides and old generated Browser pages remain searchable, some extension examples pin old dependencies, and the fast release cadence requires checking docs against the pinned release.

#### Local Workflow

> The `robot` CLI can select one test, suite, or tag, set variables such as browser/base URL, control output and log level, randomize execution, and rerun failures.<br>
> Headed mode and slow motion are configured through Browser keywords/variables, while traces and selector recording support investigation.<br>
> Pabot handles local parallel runs.<br>
> The workflow is straightforward for Robot users but introduces virtualenv, pip, Robot syntax, and Python tooling to a repository currently centered on Go and TypeScript.

#### Failure Log Correlation

> Robot logs preserve suite/test/keyword hierarchy and timestamps, and recent Browser releases added suite/test context to `playwright-log.txt`; traces correlate browser API calls, snapshots, and network details.<br>
> Console and page errors must be explicitly read or attached.<br>
> Gin, PostgreSQL, Flyway, and Compose output require a common UTC run/worker ID and custom collection into the same artifact set.<br>
> There is no automatic cross-process correlation.

#### Artifact Data Exposure

> Treat all artifacts as sensitive.<br>
> Screenshots, videos, traces, HARs, page source, console logs, storage-state files, downloads, and Robot logs can expose portfolio data, cookies, headers, API bodies, and credentials.<br>
> Official docs warn that Playwright debug logging records fill/type values in clear text and that values can also appear in traces even when `Fill Secret` suppresses Robot logging.<br>
> Use synthetic data, avoid publishing storage state, disable unnecessary debug/HAR/video capture, sanitize custom logs, restrict artifact access, and set short retention.

### Github Actions Fit

#### Official Ci Support

> Official Browser documentation provides GitHub Actions job-container examples, and the project publishes versioned Docker Hub and GHCR images for amd64/arm64.<br>
> Upstream CI itself runs on GitHub-hosted Ubuntu and uploads Robot artifacts/job summaries.<br>
> A host job can instead use `actions/setup-python`, install the pinned requirements, run `rfbrowser install --with-deps`, and execute `robot`.<br>
> No dedicated Browser action is required.

#### Browser Caching

> The official architecture guide documents a shared external browser directory through `PLAYWRIGHT_BROWSERS_PATH`, using Browser's pinned wrapper to install binaries and setting the same path during execution.<br>
> Cache keys must include OS, architecture, Browser/Playwright version, and selected engines; the default location inside the Python environment may be lost with environment recreation.<br>
> Pip/npm caches do not replace Linux browser system libraries.<br>
> A pinned official image or installing only Chromium is simpler than an unsafe stale browser cache.

#### Artifact Integration

> Low to moderate effort. `actions/upload-artifact` can upload the Robot HTML/XML/JSON outputs, Browser screenshots, videos, traces, HARs, and Node logs under `if: always()`. `robotframework-ghareports` can write a GitHub job summary after serial or Pabot-merged execution.<br>
> Compose/backend logs need explicit collection.<br>
> Missing-file handling, retention, permissions, and redaction remain workflow responsibilities.

#### Sharding And Matrix Support

> GitHub matrices can vary browser, viewport, Python version, and shard.<br>
> Pabot provides within-job parallel suite/test distribution and merges worker outputs; Rebot combines results from independent jobs if downloaded into an aggregation job.<br>
> This is external composition rather than native Browser sharding, and cross-job HTML/artifact naming requires custom workflow logic.<br>
> Database-mutating shards must receive isolated databases or remain serial.

#### Container Compatibility

> The simplest topology is to run Robot on the GitHub host against Compose-published ports.<br>
> The official `marketsquare/robotframework-browser:20.4.0` or GHCR equivalent can also run as a job container and includes browsers/system libraries.<br>
> Container runs should use `pwuser`, `--ipc=host` or adequate shared memory, and the recommended seccomp profile.<br>
> If the application is on the host or sibling containers, base URLs and Compose networks must use reachable host/service names because container `localhost` is not the host.

#### Failure Cleanup

> Default automatic closing handles pages, contexts, browsers at execution end, and the Node process; `KEEP` must not be used in CI.<br>
> Suite teardown can release fixtures, but job cancellation can bypass Robot teardown.<br>
> GitHub Actions therefore needs independent `if: always()` steps to capture logs and run `docker compose down --volumes --remove-orphans`, plus unique Compose project/volume names and timeouts.<br>
> A manually shared experimental Node process also requires explicit termination.

### Cost And Risk

#### Open Source Completeness

> The Apache-2.0 Browser and Robot Framework packages provide the required keyword runner, three managed engines, locators, waits/assertions, contexts, forms, downloads, screenshots, video, traces, HAR recording, HTTP requests, reports, and extension APIs without payment.<br>
> Pabot, axe-core, Pixelmatch, request-mocking extensions, and GitHub reporting are also available as open source.<br>
> However, visual baseline governance, current axe integration, browser request interception, Compose/database lifecycle, and long-term flake analytics are not complete core capabilities and require extra packages or custom code.

#### Optional Cloud Dependency

> No cloud service is required for local or GitHub Actions execution, browser binaries, HTML reports, traces, visual diffs, accessibility checks, or Pabot parallelism.<br>
> Commercial grids, device farms, report portals, and visual services are optional.<br>
> Local alternatives exist but require more repository ownership for baseline review, history, report aggregation, and artifact retention.

#### Migration Cost

> High relative to direct Playwright Test for this repository.<br>
> Test cases, user keywords, variables, resources, listeners, setup/teardown, Pabot controls, and reports use Robot-specific syntax and conventions, while custom libraries use Python and Browser plugins may use JavaScript.<br>
> Existing TypeScript models and helpers are not directly reusable.<br>
> Browser concepts and selectors transfer back to Playwright, but suites would be rewritten.<br>
> Keeping data and lifecycle APIs behind small domain keywords can reduce lock-in.

#### Custom Harness Burden

> Moderate to high.<br>
> Browser supplies browser lifecycle, context isolation, waits, HTTP fetch, network observation, downloads, traces, screenshots, and reports.<br>
> The repository still needs Python dependency management, Robot resources, Compose/Parcel startup, readiness and Flyway checks, PostgreSQL seed/reset or per-worker allocation, backend-side Yahoo Finance replacement, core-quality request mocking, Chart.js assertions, axe integration, visual baseline review, service-log correlation, cancellation-safe teardown, and artifact redaction.<br>
> Pabot and report aggregation add configuration if parallelism is enabled.

#### Capability Delivery Tier

> Core Browser/Robot: keyword runner, Chromium/Firefox/WebKit, locators, actionability, retrying assertions, contexts, forms, HTTP fetch, request/response observation, downloads, screenshots, video, traces, HAR capture, clock, HTML/XML/JSON/xUnit reports, and setup/teardown.<br>
> Official project distribution: BrowserBatteries and multi-architecture Docker images.<br>
> Community/separate open source: Pabot sharding, RetryFailed, GitHub summaries, axe audits, basic Pixelmatch comparison, and route/HAR request mocking.<br>
> Custom repository code: application lifecycle/readiness, PostgreSQL isolation, backend Yahoo mock, production-grade visual/axe policies, cross-job aggregation, logs, and artifact security.<br>
> Paid cloud: none required.

#### Ai Execution Boundary

> Browser, Robot Framework, Pabot, assertions, retries, reports, and CI are deterministic and do not require an LLM, AI credential, inference cost, or model egress.<br>
> Selector recording and trace viewing are local non-AI tools.<br>
> Any optional AI-assisted authoring should remain outside required CI, use synthetic data and least privilege, and produce reviewed `.robot`/resource/Python code that runs unchanged without AI.

### Evidence And Decision

#### Sources

- Robot Framework Browser official architecture and installation guide, https://robotframework-browser.org/docs/concepts/architecture, accessed 2026-08-22.
- Robot Framework Browser 20.4.0 keyword reference, https://robotframework-browser.org/keywords/20.4.0, accessed 2026-08-22.
- Official Browser guides for contexts, selectors, assertions, waiting, and logging, https://robotframework-browser.org/docs/concepts/browser-context-page, https://robotframework-browser.org/docs/concepts/selectors, https://robotframework-browser.org/docs/concepts/assertions, https://robotframework-browser.org/docs/concepts/waiting, and https://robotframework-browser.org/docs/concepts/logging, accessed 2026-08-22.
- Official Browser mobile guides, https://robotframework-browser.org/docs/mobile/responsive, https://robotframework-browser.org/docs/mobile/devices, and https://robotframework-browser.org/docs/mobile/touch, accessed 2026-08-22.
- Official Browser Docker and image guides, https://robotframework-browser.org/docs/operations/docker and https://robotframework-browser.org/docs/operations/docker-images, accessed 2026-08-22.
- Official Browser Node process and environment guides, https://robotframework-browser.org/docs/operations/node-process and https://robotframework-browser.org/docs/operations/environment-variables, accessed 2026-08-22.
- Official Browser community and Playwright Test comparison pages, https://robotframework-browser.org/community and https://robotframework-browser.org/why/vs-playwright, accessed 2026-08-22.
- Robot Framework Browser 20.4.0 release, https://github.com/MarketSquare/robotframework-browser/releases/tag/v20.4.0, observed 2026-08-22.
- MarketSquare/robotframework-browser repository metadata, commits, issues, releases, package.json, pyproject.toml, network/clock keyword source, and GitHub Actions workflows, https://github.com/MarketSquare/robotframework-browser, observed 2026-08-22.
- robotframework-browser PyPI metadata, https://pypi.org/project/robotframework-browser/20.4.0/, observed 2026-08-22.
- PyPI Stats recent-download API for robotframework-browser, https://pypistats.org/api/packages/robotframework-browser/recent, observed 2026-08-22.
- MarketSquare Browser extensions repository and axe, image comparison, and mockUrl examples, https://github.com/MarketSquare/robotframework-browser-extensions, observed 2026-08-22.
- Robot Framework 7.4.2 User Guide, https://robotframework.org/robotframework/latest/RobotFrameworkUserGuide.html, accessed 2026-08-22.
- Robot Framework retry and GitHub Actions guides, https://docs.robotframework.org/docs/flaky_tests and https://docs.robotframework.org/docs/using_rf_in_ci_systems/ci/github-actions, accessed 2026-08-22.
- Pabot repository documentation, https://github.com/mkorpela/pabot, accessed through current web evidence 2026-08-22.
- Open Asset Allocator repository evidence: `.nvmrc`, `src/main/web-static/package.json`, Parcel/HTMX/Navigo/Handlebars/Chart modules, Makefile, lifecycle scripts, Gin code, and Docker Compose/Flyway/PostgreSQL configuration, inspected 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence in release, license, current runtime/dependency support, architecture, installation, browser coverage, keywords, waiting, contexts, Docker, PyPI downloads, and repository snapshot because these come from primary documentation, source, workflows, registry data, and local files.<br>
> Medium confidence in issue health, community breadth, wrapper lag, maintainer resilience, plugin suitability, and harness burden.<br>
> Low confidence in resource cost, exact repository behavior, full supply-chain posture, and the weighted score until a project-specific spike and benchmark are run.

#### Deal Breakers

> No confirmed application incompatibility if the project accepts Python and Robot Framework.<br>
> Exclude this candidate if tests must remain TypeScript-native, if adding Python/pip/Robot dependencies is prohibited, if first-party Playwright Test projects/sharding/fixtures/UI Mode are mandatory, if core production-grade visual/axe/request-mocking workflows are required, or if native Safari on Ubuntu is expected.<br>
> Its own official comparison recommends Playwright Test when the team writes TypeScript, tests only web applications, and does not require business-readable keyword reports; that description substantially matches this repository unless non-developer test authorship is a separate requirement.

#### Recommendation

> Viable alternative, not the primary recommendation.<br>
> Robot Framework Browser 20.4.0 is actively maintained, current with Playwright 1.62.1, materially adopted in its ecosystem, and technically capable of reliable HTMX/Navigo/Handlebars/Chart.js black-box testing with strong keyword reports.<br>
> Choose it when business-readable keyword suites, non-TypeScript authors, or cross-technology Robot libraries provide enough value to justify Python and the DSL.<br>
> For this existing Go/TypeScript web repository, direct Playwright Test remains a better default because it uses the current runtime and author skills, supplies integrated projects/sharding/fixtures/reporting, exposes new Playwright capabilities without wrapper lag, and needs fewer community extensions.<br>
> If Browser is selected, begin Chromium-first and serial against a disposable database, pin 20.4.0 and images, and close the stated spike before adoption.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `roadmap_and_deprecation_risk`
- `wrapper_upstream_lag`
- `maintainer_concentration`
- `github_metrics`
- `adoption_trend`
- `resource_usage`
- `security_and_supply_chain`
- `hard_gate_result`
- `project_fit_score`
- `empirical_project_spike_result`

<a id="grafana-k6-browser"></a>
## 14. Grafana k6 Browser

Source result: `Grafana_k6_Browser.json`

### Project And Compatibility

#### Implementation Language

> k6 is implemented primarily in Go and embeds the Sobek JavaScript runtime.<br>
> Browser tests are authored in JavaScript or partially supported TypeScript and execute in the k6 binary, not in Node.js.<br>
> Chromium automation is implemented in the core k6/browser module through CDP.

#### Operating System Support

> k6 publishes Linux amd64 and arm64 binaries and supports Debian/Ubuntu package installation, so local Linux and GitHub-hosted Ubuntu amd64 are covered.<br>
> Browser tests additionally require a compatible Chromium-based browser.<br>
> Official with-browser container images bundle Chromium; the documented Apple Silicon container path may require amd64 emulation, but this does not affect the repository's expected Ubuntu amd64 CI path.

#### License And Governance

> GNU Affero General Public License v3.0 for grafana/k6, owned and commercially stewarded by Grafana Labs with public development, contribution guidance, and a public roadmap.<br>
> Running an unmodified k6 binary and keeping test scripts in this repository does not require a paid license.<br>
> Distribution or network deployment of modified k6 itself must satisfy AGPL source-offer obligations, so custom Go extensions or modified binaries need license review.

#### Installation Model

> Install the standalone k6 binary from the Grafana apt repository or a pinned GitHub release, then install a Chromium-based browser and system libraries.<br>
> Alternatively use the official grafana/k6:*with-browser image that bundles k6 and Chromium.<br>
> K6_BROWSER_EXECUTABLE_PATH can select an installed executable. npm is optional except for editor types or an external bundling workflow; no browser download command comparable to Playwright's installer is provided by the native binary.

#### Candidate Scope And Layer

> Performance-first framework and browser automation module, not a complete E2E test framework.<br>
> It combines Chromium user journeys, protocol-level load generation, Web Vitals, thresholds, and metrics, but official migration guidance states that k6/browser has no test runner or fixture system.

#### Authoring And Async Model

> Browser interactions use native-looking JavaScript/TypeScript async and await inside exported k6 lifecycle or scenario functions.<br>
> Work is scheduled as virtual users and iterations rather than named test cases.<br>
> Core checks record metrics and do not fail the process unless paired with thresholds; execution.test.fail, thrown errors, or the Grafana-maintained k6-testing assertion library can provide functional failure semantics.

#### Build Pipeline Coupling

> Black-box browser scripts can target the Parcel development origin or the production application built by Parcel and served by Gin. k6 does not require source transformation, Vite, application instrumentation, or replacement of the frontend build.<br>
> Test-side TypeScript transpilation is independent of the application build; importing npm helpers may introduce a separate bundling step because k6 does not resolve node_modules.

#### Testability Instrumentation Required

> No production build hook, injected runtime, special route, or relaxed browser security is required for ordinary UI interaction.<br>
> Current APIs provide role, label, text, title, test-ID, CSS, and XPath locators.<br>
> Stable data-testid attributes remain optional for ambiguous dynamic templates.<br>
> Deterministic database reset, server-side Yahoo Finance replacement, and fixture endpoints would be repository harness concerns rather than k6 requirements.<br>
> Disabling web security or the Chromium sandbox is optional and should not be used except where the execution environment requires it.

### Maintenance Health

#### Latest Stable Release

> v2.2.0, published August 10, 2026, was the latest principal stable release observed on August 22, 2026.<br>
> A v1.8.1 maintenance-line release was published August 12, 2026.

#### Release Cadence

> Active and frequent.<br>
> The ten newest observed releases span v1.6.0 on February 10, 2026 through v2.2.0 on August 10 and v1.8.1 on August 12, with feature releases on both v1 and v2 lines plus patches.<br>
> Grafana's stated historical cadence is approximately every two months, and the 2026 record is consistent with that cadence.

#### Repository Activity

> Very high current activity. grafana/k6 was pushed on August 22, 2026; the ten newest observed master commits span August 12-21 and include browser concurrency fixes, browser-image smoke testing, tracing, dependency security updates, and external contributions.<br>
> Open pull requests were updated on August 22, and multiple August commits correspond to recently merged pull requests.

#### Roadmap And Deprecation Risk

> k6 has a public Grafana roadmap, published release notes, and explicit migration material. k6/browser became a stable core module in v0.52 and was declared production-ready at v1.0. v2.0 deliberately removed deprecated APIs and updated Web Vitals, including removal of deprecated FID, so major upgrades can require script changes.<br>
> No deprecation or reduced-maintenance signal for k6/browser was observed.

#### Dependency Currency

> v2.2.0 release maintenance included current Go toolchain, Go dependency, container base-image, and GitHub Actions updates, including security-labeled updates. k6 ships as a self-contained Go binary and therefore does not depend on the repository's Node or Go toolchain at runtime.<br>
> Browser compatibility follows CDP and the separately installed or container-bundled Chromium, making browser/image pinning important for reproducibility.

#### Wrapper Upstream Lag

> k6/browser is a core CDP implementation, not a wrapper around Playwright or WebDriver, so conventional wrapper release lag does not apply.<br>
> However, it intentionally mirrors only part of Playwright's API.<br>
> The official parity table still lists major unsupported areas such as downloads, video, Playwright tracing, storageState, HAR replay, dialogs, and concurrent contexts; feature parity therefore trails Playwright by design even when k6 itself is current.

### Community Adoption

#### Npm Downloads

> Not applicable.<br>
> Grafana k6 and k6/browser are distributed as a Go binary, OS package, and container image, not as an official npm runtime package. @types/k6 or unrelated npm packages would not measure browser-module execution and are therefore not used as an adoption proxy.

#### Ecosystem Usage

> k6 has broad performance-testing adoption, official Grafana Cloud integration, an official Kubernetes operator, official setup-k6 and run-k6 GitHub Actions, numerous output and protocol extensions, community forums, and mature Grafana/Prometheus/InfluxDB result paths. k6 Studio and the Grafana-maintained k6-testing library extend browser authoring.<br>
> These signals cover the broad k6 ecosystem; public evidence does not isolate how many organizations use k6/browser as their functional E2E suite.

#### Community Support

> Extensive official documentation covers browser APIs, migration from Playwright, Web Vitals, CI concepts, containers, distributed execution, troubleshooting, and API parity.<br>
> Support is available through the active GitHub issue tracker and Grafana community forum. k6 Studio has dedicated documentation and current development.<br>
> The core repository does not expose GitHub Discussions, and browser-specific troubleshooting material is smaller than that of leading E2E-first frameworks.

#### Adoption Metric Normalization

> The recorded GitHub snapshot covers the entire grafana/k6 repository across protocol load testing, browser testing, CLI, cloud integration, and extensions.<br>
> It is not a count of active browser E2E teams.<br>
> No npm count is reported because there is no official npm runtime package.<br>
> Release-asset and container pulls are also omitted because they combine CI, upgrades, mirrors, architectures, and non-browser use and were not collected over a consistent comparison window.

### Browser And Runtime Coverage

#### Browser Engines

> Chromium only.<br>
> The required scenario browser type is chromium, and documentation requires a Chromium-based browser such as Google Chrome.<br>
> Firefox, WebKit, native Safari, and their corresponding engine projects are not supported.<br>
> This is the primary browser-coverage limitation for use as the repository's sole E2E framework.

#### Browser Protocol

> Chrome DevTools Protocol. k6 launches or connects to Chromium and communicates through CDP; v2.2.0 added chromium.connectOverCDP for an already-running browser.<br>
> It does not use W3C WebDriver or WebDriver BiDi, which limits standards-based remote-grid portability and confines execution to Chromium-compatible targets.

#### Headless And Headed Modes

> Headless mode is the default for local and CI runs.<br>
> K6_BROWSER_HEADLESS=false launches a visible browser for local debugging when a display is available.<br>
> Official browser containers can run headless; headed container execution requires display support and is not the normal GitHub Actions path.

#### Parallel Browser Support

> k6 scenarios, virtual users, and iterations can run multiple Chromium journeys concurrently, and execution segments can partition load across processes.<br>
> This is performance-workload parallelism, not a multi-engine project matrix.<br>
> Only one browser context can be active at a time per browser API instance, though that context may own multiple pages.<br>
> Database-mutating journeys should begin with one VU.

#### Mobile Emulation

> The k6/browser devices collection supplies predefined device settings, and newContext can configure viewport, screen, user agent, device scale, touch, and mobile behavior.<br>
> This supports responsive and mobile-like Chromium testing.<br>
> It is desktop-browser emulation, not execution on physical Android/iOS devices or Mobile Safari.

#### Real Browser Fidelity

> k6 drives a real installed or container-bundled Chromium-based browser rather than a DOM simulator.<br>
> Chrome/Chromium rendering, JavaScript, network behavior, and Web Vitals are representative of that executable and environment.<br>
> There is no Firefox or WebKit approximation, and Chromium execution provides no evidence about native Safari behavior.

#### Environment Determinism Controls

> Documented context controls include viewport/screen/device settings, user agent, locale, timezoneID, geolocation, permissions, offline mode, touch/mobile behavior, color scheme, reduced motion, HTTP credentials, proxying, and HTTPS-error handling.<br>
> Page APIs can throttle CPU/network and emulate media. addInitScript can inject controlled values, but no first-class fake clock or seeded randomness API comparable to a full E2E framework was found; backend time, random data, fonts, GPU/SwiftShader behavior, and Chart.js animation still require explicit harness control.

### Application Fit

#### Dynamic Dom Synchronization

> A good documented fit for HTMX replacement at the automation layer.<br>
> Locator resolves a fresh element after navigation or DOM changes, strict single-target methods reject ambiguity, actions wait on element conditions, and retrying k6-testing assertions can poll.<br>
> Tests still need explicit locator/assertion or response waits after HTMX swaps; ordinary checks do not retry automatically, and the official examples still use some fixed waits.<br>
> Repository-specific behavior remains unproven until a spike runs.

#### Routing Support

> Page APIs include goto, waitForURL, goBack, goForward, waitForNavigation, waitForFunction, and URL inspection, which can exercise Navigo direct links and browser history.<br>
> The repository's Gin production origin and Navigo strategy ALL permit black-box deep-link testing.<br>
> Correct route cleanup and HTMX trigger ordering must be asserted explicitly because k6 has no router-aware fixture.

#### Locator Model

> Current Page, Frame, and Locator APIs expose getByRole, getByLabel, getByText, getByTestId, getByPlaceholder, getByAltText, getByTitle, CSS, and XPath.<br>
> Locator methods reacquire dynamic nodes and enforce strictness when one target is expected.<br>
> This is materially improved in current v2-era APIs, although an older official selector guidance page still says text selectors are unsupported, showing documentation drift that should be checked against the pinned API version.

#### Form Interaction

> Locator/Page methods cover fill, clear, press, pressSequentially/type, focus, check/uncheck, setChecked, selectOption, setInputFiles, tap, keyboard, mouse, touchscreen, and DOM event dispatch.<br>
> These can exercise native validation, focus/blur flows, hidden-value synchronization, and dynamically replaced rows. k6 lacks a complete Playwright fixture/test layer, so reusable setup, validation-message helpers, and cleanup must be custom.

#### Canvas And Download Support

> Chart.js canvas can be checked through page.evaluate against application/Chart state, deterministic source data, element/page screenshots, and browser performance metrics. k6 has no integrated screenshot-baseline matcher, so pixel comparison requires an external tool.<br>
> The official Playwright-parity table marks the entire Download API unsupported, including events, suggested filename, path, stream, failure, and saveAs.<br>
> A browser-side download requirement therefore fails without a fragile custom CDP/filesystem workaround or a separate E2E tool.

#### Network And Api Access

> Strong protocol/browser combination. k6/http can seed and inspect REST fixtures, while page request/response events, waitForRequest/waitForResponse, and page.route can continue, abort, modify, or fulfill browser requests.<br>
> HAR replay, route.fetch/fallback, WebSocket routing, and context-wide routing are absent.<br>
> Browser-originated Yahoo calls can be fulfilled; Gin's server-side Yahoo Finance traffic requires the existing Go mock pattern, a mock service, or backend configuration.

#### Same Origin Support

> Directly supports both repository topologies without CORS changes: the production image exposes Gin and static assets on one origin at port 80, while Parcel development on port 8000 proxies /api to Gin.<br>
> A BASE_URL environment variable can select either.<br>
> Chromium's same-origin policy remains active unless explicitly disabled, which is the correct production-fidelity default.

#### Test Isolation

> BrowserContext separates pages, cache, and cookies, and browser.newPage creates an incognito context. k6 resets VU cookies between iterations, but k6/browser permits only one active context at a time per browser instance and lacks Playwright storageState/fixture abstractions.<br>
> PostgreSQL is external state and must be isolated through unique data, API/SQL cleanup, disposable volumes, or one-VU serialization.

#### External Server Model

> Natural black-box model.<br>
> Scripts take a URL from __ENV and run against any independently started local, containerized, or remote application. k6 does not require ownership of the server process and can perform an HTTP readiness probe in setup before browser VUs begin.

#### Application Lifecycle

> k6 provides init, setup, VU, teardown, and handleSummary stages, but no built-in webServer manager. setup can poll an API and seed data; teardown can request cleanup, but official documentation warns that teardown is not called if setup fails.<br>
> Building the application, starting Docker Compose, waiting for Flyway/Gin, collecting logs, and guaranteed down -v cleanup require GitHub Actions or a wrapper script with an unconditional external cleanup trap.

#### Visual Regression Workflow

> Not built in.<br>
> Page and element screenshots support PNG/JPEG capture, but k6 provides no baseline creation/update command, image matcher, masking, tolerances, per-browser baseline management, diff viewer, or CI review workflow.<br>
> These require an external image-diff package/process and custom artifact conventions.<br>
> Chart.js screenshots also require pinned Chromium, fonts, viewport, data, clock, and disabled animation.

### Reliability

#### Waiting Model

> Locator actions use strict, dynamic locators and actionability-related waits; explicit primitives include locator.waitFor, waitForSelector, waitForURL, waitForNavigation, waitForLoadState, waitForFunction, waitForRequest, waitForResponse, and event waits.<br>
> The companion k6-testing library offers retrying assertions.<br>
> Core check() is immediate and fixed waitForTimeout/sleep calls remain easy to misuse, so HTMX tests must standardize on state- or response-based waits.

#### Flake Controls

> Configurable page/context/navigation timeouts, scenario duration controls, repeated iterations, deterministic thresholds, and explicit failure APIs are available. k6 has no E2E-runner-level retry policy, retry-only-failed-tests mode, last-failed filter, or first-retry trace retention.<br>
> Re-running an iteration can be scripted but risks duplicating database mutations and performance samples, so retries and idempotency would be custom harness policy.

#### Isolation Model

> Each VU has separate JavaScript state, and incognito BrowserContext state separates cookies/cache, but only one context can be active at once within a browser instance.<br>
> Scenarios can separate flows, while one VU and one iteration can serialize stateful journeys.<br>
> There is no per-test fixture object or automatic database transaction/reset; browser closure in finally and application-level cleanup are mandatory.

#### Parallelism Controls

> Scenarios control VUs, iterations, start times, and separate functions; execution segments and k6 Operator can partition a workload.<br>
> Setting vus: 1 protects a shared PostgreSQL database, and independent scenarios can be delayed with startTime.<br>
> This is less test-oriented than workers/serial suites: there are no named-test shards, fixture-aware scheduling, or automatic merged E2E reports, and distributed instances evaluate thresholds/report metrics separately unless external aggregation is added.

#### Flake Observability

> Checks, threshold failures, per-scenario tags, custom metrics, JSON/CSV/time-series outputs, and browser spans can show recurring failures.<br>
> There is no core classification of first-attempt versus retry pass, quarantine registry, flaky-test report, or historical local dashboard because there is no integrated retrying test runner.<br>
> Long-term trends require a Grafana/Prometheus/InfluxDB backend, Grafana Cloud, or repository-owned processing.

### Diagnostics And Developer Experience

#### Failure Artifacts

> Screenshots are available but must be requested in script, including in catch/finally logic.<br>
> Browser console, request, response, requestfailed, metrics, and CDP debug logs can be collected through events/environment settings. k6 browser can emit OpenTelemetry-style spans, but the Playwright startTracing/stopTracing APIs, action DOM-snapshot trace viewer, video, HAR recording/replay, and Download artifacts are absent.<br>
> Backend and Compose logs remain external.

#### Debugging Tools

> Headed Chromium, K6_BROWSER_DEBUG CDP logs, browser events, page screenshots, standard JavaScript logging, and browser DevTools provide basic debugging. k6 Studio Validator can run one iteration and display interactions, requests, console messages, DOM, and a replay/preview.<br>
> The CLI has no Playwright-like inspector, pause/step UI, local time-travel trace viewer, or integrated IDE test debugger.

#### Test Generation

> k6 new --template browser creates a starter script.<br>
> The separate open-source k6 Studio desktop app can record browser events, build browser actions/assertions visually, validate one iteration, and export a k6 script.<br>
> Generated output still requires review for stable locators, HTMX synchronization, database fixtures, cleanup, and performance thresholds.

#### Reporters

> Core output is a terminal summary plus JSON/CSV and multiple time-series output integrations. handleSummary can emit arbitrary JSON, HTML, XML, or text; official examples use jslib helpers for JUnit, while community HTML reporters exist.<br>
> There is no integrated test-case HTML report, GitHub annotation reporter, trace report, or shard-report merger equivalent to an E2E-first runner; these require custom summary logic or external services.

#### Documentation Quality

> Broad, current, and generally detailed official documentation covers installation, browser APIs, Web Vitals, migration, CDP options, containers, lifecycle, outputs, and limitations through a Playwright parity table.<br>
> Quality is reduced by some stale or contradictory pages, such as selector guidance saying text selectors are unsupported while current Page/Locator APIs document getByText and other semantic locators.<br>
> Documentation should be read against the pinned k6 version.

#### Local Workflow

> A script runs with k6 run file.js; environment variables select BASE_URL, executable, headless/headed mode, timeouts, and debug logging.<br>
> Scenarios organize flows, and k6 Studio provides a one-iteration visual validator.<br>
> The CLI lacks named-test discovery, file/line/title filtering, watch mode, last-failed selection, and an interactive E2E UI, so running one functional journey usually requires a separate script/config or environment-conditioned scenario selection.

#### Failure Log Correlation

> k6 can timestamp script logs, export metrics/time series, emit browser spans, and capture console plus request/response events.<br>
> Scenario, VU, iteration, and custom tags can provide correlation IDs.<br>
> Gin, PostgreSQL, Flyway, and Docker Compose logs must be captured by the surrounding workflow with UTC timestamps and run identifiers; there is no automatic cross-process trace/report that joins browser actions with those service logs.

#### Artifact Data Exposure

> Screenshots, console/network logs, JSON outputs, custom summaries, browser spans, and k6 Studio recordings can expose portfolio data, request bodies, URLs, headers, or credentials.<br>
> Grafana Cloud execution additionally sends configured metrics/logs to a hosted service; v2.2.0 added an opt-out for cloud log streaming during local-execution.<br>
> Use synthetic data, sanitize event handlers and summaries, avoid logging secrets, restrict artifact retention/access, disable unintended cloud outputs and anonymous usage reporting where policy requires, and never persist reusable credentials in scripts.

### Github Actions Fit

#### Official Ci Support

> Grafana maintains setup-k6-action and run-k6-action, both active in July 2026, and publishes official Linux packages, binaries, and with-browser images.<br>
> A GitHub-hosted Ubuntu job can install pinned k6 and use its installed Chrome or run the browser image.<br>
> Browser-specific setup, application lifecycle, and artifacts still need workflow steps; the repository currently has only Go test/lint workflows.

#### Artifact Integration

> GitHub's upload-artifact action can upload manually captured screenshots, JSON/CSV summaries, JUnit generated by handleSummary, browser debug logs, and Compose/backend logs.<br>
> This is conventional but not automatic: scripts must capture failure screenshots and events, and workflow steps must collect files under always() while enforcing sensitivity and retention policy.<br>
> There are no videos or Playwright-style trace bundles to upload.

#### Sharding And Matrix Support

> A GitHub matrix can vary scripts, environments, or execution segments, and k6 execution segments partition VUs/iterations. k6 Operator or Grafana Cloud supplies broader distributed execution.<br>
> Open-source distributed runs have no coordinating primary; each instance applies thresholds and reports metrics independently, so aggregation is external.<br>
> There is only one browser engine, no native named-test shard command, and no integrated report merger.<br>
> Database mutations require isolated databases or no sharding.

#### Container Compatibility

> The official grafana/k6:latest-with-browser family bundles k6 and Chromium and can run beside this repository's Docker Compose services on ubuntu-latest.<br>
> If k6 runs in its own container, localhost refers to that container, so the application must be exposed through the host or a shared Compose network.<br>
> Official docs warn that no-sandbox mode is used in example images and that Chromium needs writable temporary storage and sufficient CPU/memory; a seccomp profile and trusted target are preferable.

#### Failure Cleanup

> Pages and contexts must be closed in finally blocks. k6 teardown can clean fixtures after normal execution and execution.test.abort still permits teardown, but teardown does not run when setup fails and job cancellation can bypass in-process cleanup.<br>
> GitHub Actions therefore needs unconditional external steps to collect logs and run docker compose down -v with unique volumes/project names.<br>
> The current destroy.sh omits -v, while production Compose persists PostgreSQL through a host-mounted directory, so a CI-specific disposable database path is required.

### Cost And Risk

#### Open Source Completeness

> The k6 binary, core browser module, local execution, Chromium automation, metrics, thresholds, screenshots, routing, outputs, and k6 Studio are open source under AGPL-3.0; k6-testing is Apache-2.0.<br>
> No paid service is required to run Chromium performance journeys.<br>
> However, the open-source stack is not complete for this project's primary E2E requirements because cross-engine execution, downloads, integrated retries, visual baselines, video/action traces, accessibility audits, and a fixture-aware test runner are missing rather than paywalled.

#### Optional Cloud Dependency

> Grafana Cloud k6 is optional for hosted/distributed execution, dashboards, collaboration, synthetic scheduling, and managed result retention.<br>
> Local k6 can emit terminal, JSON/CSV, custom summaries, and metrics to self-hosted Grafana-compatible stores.<br>
> Equivalent browser action trace/video and full E2E reporting do not exist locally or merely become available by enabling the core cloud output; cloud adoption also introduces account, egress, credential, retention, and cost controls.

#### Migration Cost

> High if used as the sole functional E2E framework.<br>
> Scripts use k6 lifecycle functions, VUs, scenarios, checks/thresholds, custom summaries, and a partial Playwright-like API running on Sobek.<br>
> Moving to or from an E2E-first runner requires restructuring test cases, fixtures, retries, reports, and artifacts, although async browser actions and semantic locator concepts are partially portable to Playwright.

#### Security And Supply Chain

> Official releases provide checksums and SPDX SBOM assets, and the actively maintained Go project performs frequent dependency/security updates.<br>
> Pin the k6 binary or image digest, Chromium version, remote jslib URLs, and lock any external bundler dependencies.<br>
> Runtime remote imports and optional extension auto-resolution add egress/supply-chain risk; custom xk6 extensions produce a larger custom binary.<br>
> Container examples that disable Chromium's sandbox materially weaken isolation, and cloud/log outputs plus anonymous usage reporting require explicit policy review.

#### Custom Harness Burden

> High for functional E2E use.<br>
> The repository must add test organization, fail-fast/assertion conventions, per-journey selection, retries or deliberate no-retry policy, Compose/Flyway/Gin readiness, disposable PostgreSQL seed/reset, server-side Yahoo mocking, guaranteed cleanup, download verification workaround or second tool, screenshot diffing, accessibility injection, browser/service log collection, report generation, and sensitive-artifact handling. k6 does supply useful HTTP fixture, browser, metrics, and lifecycle primitives.

#### Capability Delivery Tier

> Core: Chromium/CDP automation, semantic and CSS/XPath locators, waits, forms, contexts/pages, screenshots, network events/routing, HTTP fixtures, scenarios/VUs, Web Vitals, thresholds, outputs, lifecycle, and browser spans.<br>
> Official companion: k6 Studio, k6-testing assertions, setup-k6-action, run-k6-action, with-browser images, k6 Operator, and Grafana integrations.<br>
> Community/custom: HTML/JUnit presentation beyond basic helpers, visual diffs, axe audit injection, retries, database/Compose harness, backend mocking/log correlation, and flake trends.<br>
> Unsupported: Firefox/WebKit/Safari and the documented browser Download API.<br>
> Paid cloud: hosted execution, dashboards, collaboration, and managed analytics.

#### Ai Execution Boundary

> Core k6 scripts, browser actions, thresholds, assertions, metrics, and CI execution are deterministic and do not require an LLM, AI credentials, model egress, or per-run inference cost.<br>
> Grafana's newer AI-assisted authoring/MCP workflows are optional and should remain outside required CI.<br>
> Any generated script must be reviewed, committed, and runnable with the ordinary open-source k6 binary when AI is unavailable.

### Evidence And Decision

#### Sources

- Official Grafana k6 browser overview and API, https://grafana.com/docs/k6/latest/using-k6-browser/ and https://grafana.com/docs/k6/latest/javascript-api/k6-browser/, accessed 2026-08-22.
- Official k6 browser Page and Locator APIs, https://grafana.com/docs/k6/latest/javascript-api/k6-browser/page/ and https://grafana.com/docs/k6/latest/javascript-api/k6-browser/locator/, accessed 2026-08-22.
- Official Playwright API parity and migration guidance, https://grafana.com/docs/k6/latest/using-k6-browser/playwright-apis-in-k6/ and https://grafana.com/docs/k6/latest/using-k6-browser/migrate-from-playwright-to-k6/, accessed 2026-08-22.
- Official browser metrics and Web Vitals documentation, https://grafana.com/docs/k6/latest/using-k6-browser/metrics/, accessed 2026-08-22.
- Official browser options, execution, element-selection, and Docker guidance, https://grafana.com/docs/k6/latest/using-k6-browser/options/, https://grafana.com/docs/k6/latest/using-k6-browser/running-browser-tests/, and https://grafana.com/docs/k6/latest/using-k6-browser/recommended-practices/select-elements/, accessed 2026-08-22.
- Official k6 TypeScript/module compatibility documentation, https://grafana.com/docs/k6/latest/using-k6/javascript-typescript-compatibility-mode/ and https://grafana.com/docs/k6/latest/using-k6/modules/, accessed 2026-08-22.
- Official k6 checks, scenarios, lifecycle, result output, custom summary, and distributed execution documentation, https://grafana.com/docs/k6/latest/using-k6/checks/, https://grafana.com/docs/k6/latest/using-k6/scenarios/, https://grafana.com/docs/k6/latest/using-k6/test-lifecycle/, https://grafana.com/docs/k6/latest/get-started/results-output/, https://grafana.com/docs/k6/latest/results-output/end-of-test/custom-summary/, and https://grafana.com/docs/k6/latest/testing-guides/running-large-tests/, accessed 2026-08-22.
- Official k6 installation and distributed browser container guidance, https://grafana.com/docs/k6/latest/set-up/install-k6/ and https://grafana.com/docs/k6/latest/set-up/set-up-distributed-k6/browser-tests/, accessed 2026-08-22.
- Official k6 Studio browser editor and validator documentation, https://grafana.com/docs/k6/latest/k6-studio/components/browser-editor/ and https://grafana.com/docs/k6/latest/k6-studio/components/validator/, accessed 2026-08-22.
- grafana/k6 GitHub repository and releases, https://github.com/grafana/k6 and https://github.com/grafana/k6/releases/tag/v2.2.0, observed through GitHub API data on 2026-08-22.
- Grafana-maintained companion repositories, https://github.com/grafana/k6-studio, https://github.com/grafana/k6-jslib-testing, https://github.com/grafana/setup-k6-action, and https://github.com/grafana/run-k6-action, observed 2026-08-22.
- Grafana k6 v1.0 and v2.0 release material, https://github.com/grafana/k6/releases/tag/v1.0.0 and https://grafana.com/blog/k6-2-0-release/, accessed 2026-08-22.
- Independent repository health snapshot, https://inspect.software/software/grafana/k6, observed 2026-08-22 and used only for labeled contributor-concentration context.
- Open Asset Allocator repository evidence: .nvmrc, frontend package.json and tsconfig.json, Parcel proxy, HTMX templates, Navigo routing, Chart.js modules, Gin server, Docker Compose files, lifecycle scripts, Go integration fixtures/Yahoo mock, Makefile, and GitHub Actions workflows, inspected 2026-08-22.

#### Observed At

2026-08-22

#### Confidence

> High confidence in license, release/activity snapshot, Chromium-only/CDP architecture, TypeScript runtime model, Web Vitals, missing Playwright APIs, lifecycle, and documented browser capabilities because these rely on first-party docs and GitHub API data.<br>
> Medium confidence in issue health, adoption direction, maintainer concentration, and resource sizing because broad k6 metrics do not isolate browser use and one source is independent.<br>
> Low confidence in repository-specific reliability, exact CI cost, axe injection, and HTMX/Navigo/Chart behavior until an empirical spike runs.

#### Hard Gate Result

> FAIL.<br>
> Passes: active open-source maintenance; JavaScript/partial TypeScript authoring; Linux and GitHub Ubuntu operation; black-box Parcel/Gin testing; Chromium automation; HTMX-capable dynamic locators and waits; Navigo history APIs; realistic forms; REST fixture setup; browser request observation/mocking; screenshots; Web Vitals; and external-server execution.<br>
> Fails mandatory primary-E2E coverage: no Firefox or WebKit/Safari engine, no supported browser Download API, no complete test runner/fixtures, no integrated functional retries, no action trace/video, and no built-in visual-regression or accessibility workflow.<br>
> The absent cross-engine and download capabilities cannot be closed with ordinary configuration.

#### Deal Breakers

> Confirmed exclusion-level gaps for the sole E2E role are Chromium-only execution and the entirely unsupported Download API.<br>
> Additional material gaps are lack of a fixture-aware test runner, first-class retries, automatic failure screenshots, video/action traces, visual baseline tooling, and accessibility audits.<br>
> The AGPL license is manageable for unmodified tool use but is less permissive than Apache/MIT for distributed modifications.<br>
> None of these prevents using k6 as a supplementary Chromium performance/Web Vitals suite.

#### Recommendation

> Excluded as the repository's primary E2E framework.<br>
> The documented lack of Firefox/WebKit and browser download support fails core coverage, while missing runner, fixture, retry, trace/video, visual, and accessibility workflows would create high custom-harness cost.<br>
> Retain Grafana k6 Browser as a viable supplementary performance-first tool after a complete E2E framework is selected: use one or a few deterministic Chromium journeys to enforce Web Vital and frontend-performance thresholds under protocol-generated load.

### Uncertain Fields

- `node_and_typescript_compatibility`
- `issue_health`
- `maintainer_concentration`
- `github_metrics`
- `adoption_trend`
- `browser_version_management`
- `accessibility_audit_integration`
- `resource_usage`
- `browser_caching`
- `project_fit_score`
- `empirical_project_spike_result`
