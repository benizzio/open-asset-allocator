.DEFAULT_GOAL := build

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