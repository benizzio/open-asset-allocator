# Containerized Playwright E2E tests

<!-- Authored by: OpenCode -->

This package tests the complete application through browser, HTTP API, and PostgreSQL boundaries. Docker Compose starts
the application topology, Flyway migrations, a disposable PostgreSQL database, and a Playwright runner. The host only
orchestrates those containers and stores generated diagnostics.

## Requirements

Run E2E commands from the repository root. The required host tools are:

- Docker Engine;
- Docker Compose through `docker compose`;
- `make`.

The E2E commands do not require host installations of Node.js, npm dependencies, Playwright browsers, Go, or ZSH. A host
Node.js installation may be useful for editor integration, but the container image and lockfile define the supported
runtime.

## Commands

| Command | Topology and behavior |
|---|---|
| `make e2e` | Build the local split topology, run Chromium and Firefox headlessly, capture logs, and remove E2E resources. |
| `make e2e-ci` | Build the immutable monolith topology used by GitHub Actions, run both browsers headlessly, capture logs, and remove E2E resources. |
| `make e2e-chromium` | Run only the local Chromium project. |
| `make e2e-firefox` | Run only the local Firefox project. |
| `make e2e-ui` | Keep the local topology active while Playwright UI Mode is available at `http://localhost:9324`. |
| `make e2e-headed` | Keep the local topology active while a headed browser is available through noVNC at `http://localhost:7900`. |
| `make e2e-report` | Serve the most recent HTML report at `http://localhost:9323` without starting the application. |
| `make e2e-logs` | Follow service logs for an active E2E project from another terminal. |
| `make e2e-clean` | Remove E2E containers, networks, and volumes left by an interrupted command. Generated artifacts remain. |

Stop UI, headed, and report sessions with Ctrl-C. Their signal traps remove resources owned by the selected E2E project.
Stopping `make e2e-logs` with Ctrl-C only stops following logs; it leaves the active stack unchanged.

### Focused runs

`E2E_ARGS` forwards arguments to Playwright without changing the container lifecycle. Examples:

```shell
make e2e-chromium E2E_ARGS="tests/portfolio-creation.e2e.spec.ts"
make e2e E2E_ARGS="--grep='creates a portfolio'"
make e2e-ci E2E_ARGS="--project=firefox"
```

Playwright uses one worker because every test attempt shares the same mutable database. Chromium and Firefox therefore
run serially.

### Build controls

Docker's build cache is enabled by default. Use these opt-in controls when diagnosing image state:

```shell
E2E_NO_CACHE=1 make e2e
E2E_PULL=1 make e2e-ci
```

`E2E_NO_CACHE=1` rebuilds all selected image layers. `E2E_PULL=1` checks for newer versions of referenced base images.

## Topologies

Test and interactive modes combine `src/main/docker/docker-compose-e2e.yml` with one topology-specific override. The
shared file defines PostgreSQL, Flyway, the runner, the E2E network, the disposable database volume, and the artifact
bind mount.

### Local split application

`make e2e`, focused runs, UI Mode, and headed mode use `docker-compose-e2e-local.yml`:

- `frontend` runs the development Parcel image with selected frontend source files mounted read-only;
- `backend` runs the standalone API image;
- Parcel proxies same-origin `/api` requests from `frontend:8000` to `backend:8080`;
- `e2e-runner` or `e2e-debug` accesses the application through `http://frontend:8000`;
- test configuration, support code, and specifications are mounted read-only for local iteration;
- the debug runner adds Xvfb, x11vnc, and noVNC only for headed inspection.

Frontend and E2E source edits are visible without restarting Compose. Dependency file or image changes still require a
build; the orchestration command invokes a cache-aware build on every run.

### CI monolith

`make e2e-ci` and `.github/workflows/e2e.yml` use `docker-compose-e2e-ci.yml`:

- `monolith` is built from the production runtime target with immutable frontend assets;
- Gin serves both static SPA content and API routes from `http://monolith:8080`;
- source directories are not mounted into the monolith or runner;
- application and database ports are not published to the GitHub-hosted runner;
- GitHub Actions assigns a run-specific Compose project name, uploads `target/e2e-results` under `if: always()`, and
  invokes emergency E2E-only cleanup under `if: always()`.

This topology tests the production image boundary rather than the development Parcel proxy.

## Local ports

Local bindings default to loopback. Override a port or bind address by setting the corresponding environment variable
before the Make command.

| Purpose | URL or address | Override |
|---|---|---|
| Frontend application | `http://localhost:8082` | `E2E_FRONTEND_PORT` |
| Backend API | `http://localhost:8083` | `E2E_BACKEND_PORT` |
| PostgreSQL | `localhost:5434` | `E2E_POSTGRES_PORT` |
| HTML report | `http://localhost:9323` | `E2E_REPORT_PORT` |
| Playwright UI Mode | `http://localhost:9324` | `E2E_UI_PORT` |
| noVNC headed browser | `http://localhost:7900` | `E2E_NOVNC_PORT` |
| Bind address for every published port | `127.0.0.1` | `E2E_HOST_IP` |

For example:

```shell
E2E_FRONTEND_PORT=9082 E2E_POSTGRES_PORT=6434 make e2e-ui
```

The E2E project defaults to `open-asset-allocator-e2e`. A custom `COMPOSE_PROJECT_NAME` must equal that name or start
with `open-asset-allocator-e2e-`. Reuse the same value when cleaning a custom project:

```shell
COMPOSE_PROJECT_NAME=open-asset-allocator-e2e-manual make e2e-ci
COMPOSE_PROJECT_NAME=open-asset-allocator-e2e-manual make e2e-clean
```

Production, development, Go integration tests, and E2E use distinct Compose projects and reserved host ports. They can
run concurrently. Multiple E2E stacks also require unique project names and non-conflicting local port overrides.

## Database isolation

The runner connects as a PostgreSQL administrator only inside the disposable E2E network. Before every test attempt, the
automatic fixture discovers base tables in the `public` schema and executes a safely quoted
`TRUNCATE ... RESTART IDENTITY CASCADE`. Flyway history tables are excluded.

The fixture repeats the reset after each attempt, verifies that portfolio state was removed, and closes its connection
pool when the worker exits. Reset-before-attempt is authoritative because process termination can skip teardown. A
non-interactive suite then removes the complete named PostgreSQL volume as the final isolation boundary.

Tests that need database access must import `test` and `expect` from `support/fixtures.ts`, not directly from
`@playwright/test`.

## Artifacts

The runner writes all diagnostics to `target/e2e-results`, which is bind-mounted from the host and preserved after
cleanup:

| Path | Contents |
|---|---|
| `html-report/` | Playwright HTML report for the latest suite. |
| `junit.xml` | JUnit result report. |
| `test-results/` | Per-test traces, screenshots, and videos retained on failure. |
| `compose.log` | Timestamped PostgreSQL, Flyway, and application service logs captured before cleanup. |

Use `make e2e-report` to serve `html-report/`. GitHub Actions uploads the complete directory for every workflow outcome
and retains the workflow artifact for 14 days.

## Dependency updates

Dependabot checks npm dependencies under `/src/test/e2e` and the Playwright Docker image under
`/src/main/docker/e2e` weekly. Playwright updates must keep these sources on one exact version:

- `@playwright/test` in `package.json`;
- Playwright packages in `package-lock.json`;
- `mcr.microsoft.com/playwright` in `src/main/docker/e2e/Dockerfile`.

The runner build executes `npm run validate:playwright` and fails if package, lockfile, Dockerfile, expected image, or
installed image versions differ. A dependency update is not complete until all Playwright version sources agree.

## Troubleshooting

### Ports or stale resources conflict

Run `make e2e-clean`, then retry. If another legitimate process uses a reserved port, set the relevant override from
the ports table. Cleanup addresses only the selected E2E project.

### A suite fails before the application is ready

Inspect `target/e2e-results/compose.log`. During an interactive run, use `make e2e-logs` from another terminal. The
runner waits up to 30 seconds for static content, the same-origin API, and PostgreSQL before reporting which boundary did
not become ready.

### The HTML report is missing

Run a headless suite first. `make e2e-report` requires `target/e2e-results/html-report/index.html` and does not start the
application stack.

### A dependency update reports a Playwright version mismatch

Align `package.json`, `package-lock.json`, and the literal Playwright image version in the runner Dockerfile. Do not
disable the validation or use version ranges.

### Docker appears to reuse obsolete image content

Confirm that the edited path belongs to the selected build context. Retry once with `E2E_NO_CACHE=1`. Use `E2E_PULL=1`
separately when the intended change is a refreshed base image.

### An interrupted run remains active

Run `make e2e-clean` with the same `COMPOSE_PROJECT_NAME` used by the interrupted command. Cleanup removes containers,
networks, and database volumes but preserves `target/e2e-results` for diagnosis.
