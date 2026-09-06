# Open Asset Allocator Project

An Open source project composed of:

- An open, layman-summarized base of knowledge to learn about the subject and methodology Asset Allocation, known Asset
  Allocation strategies and concepts tied to a framework around the concept
  of [Antifragility](https://en.wikipedia.org/wiki/Antifragility)
- An open source tool to apply the knowledge and manage and monitor a long term investment portfolio using asset
  allocation strategies as a first-class citizen.

> [!CAUTION]
> This project is in pre-alpha stage and its tool not ready for production use. It can evolve to a cloud ready service
> in the
> future, but is currently a work in progress without minimum security for anything but local usage.

## What is an asset allocation strategy, and how can I learn more?

![asset allocation flow](/docs/images/asset-allocation-flow.png)

> An asset allocation strategy is a plan for how to distribute your investments across different asset classes, such as
> stocks, bonds, and cash. The goal of an asset allocation strategy is to balance risk and reward by diversifying your
> investments.

Yes, this is the AI generated, statistically predominant, by the book definition. Something similar can also be found in
places like [Investopedia](https://www.investopedia.com/terms/a/assetallocation.asp)
and [Wikipedia](https://en.wikipedia.org/wiki/Asset_allocation). It is bland and not very helpful.

Some definitions I believe are a lot more useful:

- A strategy to manage assets with focus in the long term, with clear fallbacks for short term volatility and
  deviations;
- A technique to store cognitive energy in a plan that allows for intuitive, "automatic" behavior when moving assets,
  helping detach it from emotional influences and deviations;
- A way to time the markets without price and period targets, taking advantage of its cyclic nature (where appliable)
  to buy "cheap" and sell "expensive";
- A process that is historically used to manage large ammounts of assets, enjoying large accumulation zones and
  permanent time in the market;
- A method that allows for higher focus on some very important risk assessment
  and [convex payoff](https://youtu.be/ovEPIQR65hc) topics:
    - Maintaining diversification with
      [uncorrelated assets](https://www.investopedia.com/articles/financial-theory/09/uncorrelated-assets-diversification.asp)
      (without forgetting correlation is fickle, as it breaks and is resstablished dynamically during cycles);
    - Exploring
      [barbell (bimodal) strategies](https://www.investopedia.com/articles/investing/013114/barbell-investment-strategy.asp);
    - Taking care of [hedging tail risks](https://youtu.be/o3Qno1rT-nw).

You can learn more about this and related topics in the
project's [Personal Knowledge Management (PKM) docs](https://www.notion.so/Asset-Allocation-and-Market-Assessment-PKM-228acfc60ff08019ab03e7bc10dc7935).

## What is this open source tool?

An application that allows the continued, long-term portfolio management with the usage of asset allocation strategies
in a "fractal" structure. The allocation needs to follow this bottom-up, repeatable structure, composed of assets on the
lowest level. Those assets can be grouped by any criteria defined by the implemented model, and that will form the top
levels of the hierarchical structure.

> [!CAUTION]
> The examples shown in this document or any part of the project are for mere illustration purposes of the application
> features with real data. They are not, at any point in time, to be considered investment advice. Asset values are also
> fictional.

![class level](/docs/images/portfolio-history.png)

[Learn more...](/docs/readme-features.md)

## Installing and running (pre-alpha)

<!-- Co-authored by: OpenCode and Igor Benicio de Mesquita -->

> [!IMPORTANT]
> The pre-alpha version requires higher technical knowledge about software to use.
> The application workflows require `git`, Docker Engine, Docker Compose, and `make`.
> Verify if there is a provisioning script for your OS in the `provisioning` folder. If not, I welcome PRs to add them.

### Production workflow

Production uses one `open-asset-allocator-monolith:latest` application image. The image contains the Go backend and an
immutable Parcel distribution, which Gin serves with the API from the same process and origin.

1. Clone the repository: `git clone https://github.com/benizzio/open-asset-allocator.git`.
2. Build the monolith image: `make` or `make build`.
3. Start the monolith, PostgreSQL, and Flyway services: `make start`.
4. Open [http://localhost](http://localhost).

Run `make stop` to retain stopped containers or `make destroy` to remove application containers and networks. Production
PostgreSQL data remains under `~/.open-asset-allocator/postgres-data`. Application configuration is read from
[src/main/docker/.env](src/main/docker/.env).

### Development workflow

Development runs Parcel and the backend in separate containers while retaining PostgreSQL and Flyway:

1. Build the frontend and backend development images: `make dev-build`.
2. Start the development stack with Parcel logs attached: `make dev`.
3. Open [http://localhost:8000](http://localhost:8000).

Parcel proxies `/api` requests to the API-only backend container. Backend HTTP and Delve remain available at
`http://localhost:8080` and `localhost:2345`. Edits under `src/main/web-static/websrc` are visible to Parcel without an
image rebuild or Compose restart.

Frontend `node_modules`, `.parcel-cache`, and `dist` directories exist only in the frontend container writable layer.
They are not host bind mounts or persistent Docker volumes. `make stop` retains the stopped container and its temporary
frontend state; `make destroy` removes the container and clears that state. Development PostgreSQL data remains under
`target/postgres-dev-data`.

Host Node.js is optional for the containerized workflow. When it is needed for editor integration or direct frontend
commands, run `nvm install` and `nvm use` from the repository. The exact version in `.nvmrc` is also used by the frontend
development and monolith build images. Run `make validate-node-version` to verify those runtime references.

Renovate owns coordinated Node.js runtime, Playwright toolchain, Node.js type, and NVM installer updates. Application
runtime updates are restricted to LTS releases and update `.nvmrc` and both Node.js Docker images together. Playwright
updates synchronize the E2E npm package and runner image. Dependabot continues to manage every other configured
dependency ecosystem.

#### End-to-end testing workflow

The Playwright E2E workflow requires Docker Engine, Docker Compose, and `make`. It does not install Node.js, Go, ZSH,
Playwright browsers, or npm dependencies on the host. Host Node.js tooling is optional for editor integration; the
containerized commands are the supported test environment.

| Command | Purpose |
|---|---|
| `make e2e` | Run Chromium and Firefox against the source-mounted frontend and standalone backend. |
| `make e2e-ci` | Run Chromium and Firefox against the immutable production monolith used in CI. |
| `make e2e-chromium` | Run the local topology with Chromium only. |
| `make e2e-firefox` | Run the local topology with Firefox only. |
| `make e2e-ui` | Start Playwright UI Mode at [http://localhost:9324](http://localhost:9324). |
| `make e2e-headed` | Start a headed browser through noVNC at [http://localhost:7900](http://localhost:7900). |
| `make e2e-report` | Serve the latest HTML report at [http://localhost:9323](http://localhost:9323). |
| `make e2e-logs` | Follow logs for an active E2E stack. |
| `make e2e-clean` | Recover resources left by an interrupted E2E command while preserving artifacts. |

Pass Playwright file, title, or project filters through `E2E_ARGS`. For example:

```shell
make e2e-chromium E2E_ARGS="tests/portfolio-creation.e2e.spec.ts"
make e2e E2E_ARGS="--grep='creates a portfolio'"
```

Non-interactive suites capture reports and Compose logs under `target/e2e-results` before removing their disposable
PostgreSQL volume. Reserved project names and ports allow production, development, Go integration tests, and one E2E
stack to run concurrently. See the [E2E testing guide](src/test/e2e/README.md) for topology, port overrides, database
isolation, artifacts, dependency updates, and troubleshooting.

> [!NOTE]
> Current pre-alpha version requires data ingestion or manual data insertion on the PostgreSQL database.
> To access the stored portfolio go to `http://localhost:8000/portfolio/<portfolio id>`.

## Roadmap(ish)

#### Backlog - Not planned yet

- none

#### Pre-alpha - Phase 1

**Main stack proof of concept**: Assess that the base technologies and tools selected can solve the problem
consistently.

**Asset allocation management**: Allow the registration and tracking of portfolions and allocations plans. Facilitate
assessing portfolio vs. plan situation through time, helping re-balancing decisions through divergence analysis.

- [x] Portfolio management:
    - [x] new portfolio - add form based data input (beginner-user-friendly - start with fixed base structure class->
      asset)
    - [x] edit portfolio - base fields (e.g. name) only
- Portfolio history
    - [x] add form based data input (beginner-user-friendly)
        - [ ] show total market value for entire form
    - [ ] copy data from different obsevation timestamps
    - [ ] template data from plan
    - [ ] no content message
    - [ ] market prices from external source (e.g. Yahoo Finance)
    - [ ] improve chart navigation
    - [ ] Line chart with portfolio total market value over time - history progress chart and performance
- Allocation Plan
    - [x] add form based data input (beginner-user-friendly)
    - [ ] copy data from other plans
    - [ ] no content message
- Allocation map:
    - DIVERGENCE: last portfolio history from a timeframe, select an allocation plan to analyze
        - [ ] cash reserve marker on table (label & color)
        - [ ] persistence of generated analysis on timeframe for convenience `BLOCKER FOR CONVERGENCE`
        - [x] show divergent value in currency units
        - [x] show divergent value in percentage
        - [ ] show divergence mode used with informative text
        - [ ] add percentage divergence mode config ("aggressive" from level or "conservative" from asset)
- [ ] UI improvements:
    - [ ] Color configuration per hierarchy level record for visual consistency (apply to charts and tables)
    - [ ] Change accordion header font to differentiate from content

#### Alpha - Phase 2

**Portfolio management**: Enhance re-balancing assessments through convergence analysis and planning.

- Allocation map:
    - [ ] CONVERGENCE:
        - [ ] create from a mapped divergence analysis (bound to a divergence analysis and its plan)
        - [ ] fractal execution: create per level and per parent selection when below top
        - [ ] line for each divergent record, add a field for post-convergence goal (maket value) and calculate new
          percentage
        - [ ] allow obtaining inserted external cash inflow (“input” class) and using it for convergence in any asset
            - (?) upper field inside each level of fractal structure to visualize re-balancing and availability deeper
              within the hierachy.
        - (?) show final state after convergence (chart maybe?)
        - [ ] allow creation of multiple steps for segmented convergence plan
            - (?) Copy previous unexecuted plans on new divergence analysis
        - (?) field to input current market value at the time of convergence execution and calculate movement quantity
          (no need to persist)

#### Beta - Phase 3

- Portfolio management:
    - [ ] add structure management (add/remove hierarchy levels)
- Portfolio history
    - [ ] allow inclusion of external cash inflow via UI (as separate "input" class)
    - [ ] ui to handle data ingestion from known sources
- Allocation map:
    - DIVERGENCE:
    - [ ] add divergence bar mode config (change proportion from 2:1)
- [ ] Data ingestion improvements:
    - [ ] equivalent tickers field for assets
    - [ ] UI to call ingestion
    - [ ] Ghostfolio ingestion:
        - [ ] direct integration from API (holdings call)

**Asset allocation intelligence**: Expand management capabilities allowing users to understand how an asset allocation
strategy can be created and modified for personal needs and expectations. Integrate indicators and data sources to
access the ciclic nature of the markets and which phases of a cycle are currently experienced.

**Financial education and research**: Provide links to source material and educational content integrated to the
features and where they are used
