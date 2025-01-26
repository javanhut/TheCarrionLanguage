# Makefile for building and pushing a Docker image to Docker Hub
# with a user-supplied version/tag (no hard-coding).

# Docker Hub repo name, e.g. "<username>/carrionlanguage"
IMAGE_NAME ?= carrionlanguage

# VERSION is provided at run time:
#   make build VERSION=0.1.0
#   make push VERSION=0.1.0
# If you don't want a default, you could also do:
#   ifndef VERSION
#   $(error "VERSION not set! Usage: make build VERSION=x.x.x")
#   endif
VERSION ?= dev

.PHONY: build push run clean

## Build the Docker image with two tags:
## - <repo>:<VERSION> (e.g. carrionlanguage:0.1.0)
## - <repo>:latest
build:
	@echo "Building Docker image with tags:"
	@echo "  - $(IMAGE_NAME):$(VERSION)"
	@echo "  - $(IMAGE_NAME):latest"
	docker build \
		-t "$(IMAGE_NAME):$(VERSION)" \
		-t "$(IMAGE_NAME):latest" \
		.

## Push the image to Docker Hub under both tags
push:
	@echo "Pushing Docker image to Docker Hub:"
	@echo "  - $(IMAGE_NAME):$(VERSION)"
	@echo "  - $(IMAGE_NAME):latest"
	docker push "$(IMAGE_NAME):$(VERSION)"
	docker push "$(IMAGE_NAME):latest"

## Optional: run the container interactively from the specific tag
run:
	docker run --rm -it "$(IMAGE_NAME):$(VERSION)" bash

## Clean your local Docker images (optional)
clean:
	docker rmi "$(IMAGE_NAME):$(VERSION)" || true
	docker rmi "$(IMAGE_NAME):latest" || true

