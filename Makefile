.DEFAULT_GOAL := build

dev:
	@./dev.sh

destroy:
	@./destroy.sh

build:
	@echo "Preparing build folders..."
	@rm -rf target/build/*
	@[ -d "target/build/" ] || mkdir "target/build/"
	@[ -d "target/build/src-go" ] || mkdir "target/build/src-go"
	@[ -d "target/build/dist-web-static" ] || mkdir "target/build/dist-web-static"

	@echo "Building frontend..."
	@cd src/main/web-static && npm run build

	@echo "Copying frontend resources for build..."
	@cp -rf src/main/web-static/dist/* target/build/dist-web-static

	@echo "Copying backend resources for build..."
	@cp -rf src/main/go/* target/build/src-go

	@echo "Copying docker resources for build..."
	@cp src/main/docker/backend/Dockerfile target/build

	@echo "Building docker image..."
	@cd target/build && docker build --tag open-asset-allocator .