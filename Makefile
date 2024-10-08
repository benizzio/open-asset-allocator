.DEFAULT_GOAL := build

dev:
	@./dev.sh

dev-logs:
	@docker logs -f open-asset-allocator-dev-backend-1

dev-migration-logs:
	@docker logs open-asset-allocator-dev-migration-engine-1

destroy:
	@./destroy.sh

build:
	./build.sh

start:
	@./start.sh

stop:
	@./stop.sh