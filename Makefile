# Makefile for building and running a Docker image named CarrionLang

# You can change these if you want a specific version/tag
IMAGE_NAME := carrionlanguage
TAG := latest
IMAGE := $(IMAGE_NAME):$(TAG)

.PHONY: all build run push clean

# Default target: build the image
all: build

build:
	@echo "Building Docker image $(IMAGE)..."
	docker build -t $(IMAGE) .

run:
	@echo "Running Docker image $(IMAGE) interactively..."
	docker run --rm -it $(IMAGE) bash

push:
	@echo "Pushing Docker image $(IMAGE) to registry..."
	# Make sure you're logged in (docker login) to the registry you intend to push to
	docker push $(IMAGE)

clean:
	@echo "Removing Docker image $(IMAGE)..."
	docker rmi -f $(IMAGE) || true

