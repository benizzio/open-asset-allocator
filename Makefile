.DEFAULT_GOAL := build


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

# Prints the logs for the development migration engine
dev-migration-logs:
	@docker logs open-asset-allocator-dev-migration-engine-1

# Tails the logs for the development database
dev-db-logs:
	@docker logs -f open-asset-allocator-dev-db-1

# Tails the logs for the backend
logs:
	@docker logs -f open-asset-allocator-backend-1

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
