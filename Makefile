# Makefile for building and pushing a Docker image to Docker Hub
# with a user-supplied version/tag (no hard-coding).
#
# Also supports an OS parameter (OS=windows, OS=mac, OS=linux, etc.)
# Default is OS=linux if none is provided.

# --- Variables --- #
# Default OS if none is provided
OS ?= linux

# Docker Hub repo: e.g. "username/carrionlanguage"
USER_NAME ?= username
IMAGE_NAME ?= carrionlanguage

# Default Docker image tag/version
VERSION ?= latest

# --- Phony Targets --- #
.PHONY: build push run clean install uninstall

## Build the Docker image
build:
	@echo "-------------------------------------------"
	@echo "Building Docker image for OS=$(OS)"
	@echo "Version Tag: $(VERSION)"
	@echo "-------------------------------------------"
	docker build \
		--build-arg TARGET_OS=$(OS) \
		-t "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)" \
		-t "$(USER_NAME)/$(IMAGE_NAME):latest" \
		.

## Push the image to Docker Hub
push:
	@echo "-------------------------------------------"
	@echo "Pushing Docker image for OS=$(OS)"
	@echo "Version Tag: $(VERSION)"
	@echo "-------------------------------------------"
	docker push "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)"
	docker push "$(USER_NAME)/$(IMAGE_NAME):latest"

## Run the container (example: open a bash shell)
run:
	@echo "-------------------------------------------"
	@echo "Running Docker image for OS=$(OS) with tag $(VERSION)"
	@echo "-------------------------------------------"
	docker run --rm -it "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)" bash

## Remove local images (if you want to reclaim space)
clean:
	@echo "-------------------------------------------"
	@echo "Cleaning local Docker images for OS=$(OS) with tag $(VERSION)"
	@echo "-------------------------------------------"
	docker rmi -f "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)" || true
	docker rmi -f "$(USER_NAME)/$(IMAGE_NAME):latest" || true

## Install Carrion locally (calls your install script)
install:
	@echo "-------------------------------------------"
	@echo "Installing Carrion on OS=$(OS)"
	@echo "-------------------------------------------"
	# Adjust paths/scripts to match your project layout.
	@./install/install.sh $(OS)

## Uninstall Carrion locally (calls your uninstall script)
uninstall:
	@echo "-------------------------------------------"
	@echo "Uninstalling Carrion on OS=$(OS)"
	@echo "-------------------------------------------"
	@./install/uninstall.sh $(OS)

