.DEFAULT_GOAL := build

.PHONY: frontend-install e2e e2e-ci e2e-chromium e2e-firefox e2e-ui e2e-headed e2e-report e2e-logs e2e-clean

E2E_ARGS ?=


# Runs golangci-lint on the Go source
lint:
	cd src/main/go && golangci-lint run ./...

# Runs golangci-lint formatter (goimports) on the Go source
lint-fmt:
	cd src/main/go && golangci-lint fmt ./...

# Installs the front-end npm dependencies
frontend-install:
	cd src/main/web-static && npm install

# Runs the tests for the application
test:
	./test.sh

# Runs the external integration tests (requires network access to external APIs)
test-ext:
	cd src/main/go && go test -count=1 -tags=extinttest ./extinttest/...

# Builds the application for development mode
dev-build:
	./build-dev.sh

# Runs the entire application in development mode
dev:
	@./dev.sh

# Tails the logs for the development backend
dev-logs:
	@docker logs -f open-asset-allocator-dev-backend-1

# Tails Parcel logs from the development frontend container.
# Co-authored by: OpenCode and Igor Benicio de Mesquita
dev-frontend-logs:
	@POSTGRES_DEV_DATA_DIR="$(CURDIR)/target/postgres-dev-data" docker compose -f src/main/docker/docker-compose-dev.yml logs -f frontend

# Prints the logs for the development migration engine
dev-migration-logs:
	@docker logs open-asset-allocator-dev-migration-engine-1

# Tails the logs for the development database
dev-db-logs:
	@docker logs -f open-asset-allocator-dev-db-1

# Tails the logs for the production monolith.
# Co-authored by: OpenCode and Igor Benicio de Mesquita
logs:
	@POSTGRES_DATA_DIR="$(HOME)/.open-asset-allocator/postgres-data" docker compose -f src/main/docker/docker-compose.yml logs -f monolith

migration-logs:
	@docker logs open-asset-allocator-migration-engine-1

# Stops and removes all docker components
destroy:
	@./destroy.sh

# Builds the application for production usage
build:
	./build.sh

# Starts the application in production mode
start:
	@./start.sh

# Stops the application in production mode
stop:
	@./stop.sh

# Runs local split-application E2E tests in Chromium and Firefox.
# Authored by: OpenCode
e2e:
	@./e2e.sh local $(E2E_ARGS)

# Runs production-monolith E2E tests for CI parity.
# Authored by: OpenCode
e2e-ci:
	@./e2e.sh ci $(E2E_ARGS)

# Runs local E2E tests using only Chromium.
# Authored by: OpenCode
e2e-chromium:
	@./e2e.sh local --project=chromium $(E2E_ARGS)

# Runs local E2E tests using only Firefox.
# Authored by: OpenCode
e2e-firefox:
	@./e2e.sh local --project=firefox $(E2E_ARGS)

# Starts local split services and Playwright UI Mode.
# Authored by: OpenCode
e2e-ui:
	@./e2e.sh ui $(E2E_ARGS)

# Starts local split services and a headed browser available through noVNC.
# Authored by: OpenCode
e2e-headed:
	@./e2e.sh headed $(E2E_ARGS)

# Serves the latest Playwright HTML report.
# Authored by: OpenCode
e2e-report:
	@./e2e.sh report

# Follows logs for active E2E services.
# Authored by: OpenCode
e2e-logs:
	@./e2e.sh logs

# Recovers resources left by interrupted E2E commands while preserving artifacts.
# Authored by: OpenCode
e2e-clean:
	@./e2e.sh clean
