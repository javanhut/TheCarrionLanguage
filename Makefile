# Makefile snippet

# Existing variables
USER_NAME ?= username
IMAGE_NAME ?= carrionlanguage
VERSION ?= latest
OS ?= linux

.PHONY: build push run clean install uninstall build-source build-linux build-windows

# 1) Build a tarball of the uncompiled source
build-source:
	git archive --format=tar.gz -o carrion-src.tar.gz HEAD	

# 2) Build the Linux binary + tarball
build-linux:
	GOOS=linux GOARCH=amd64 go build -o carrion ./src
	tar -czf carrion_linux_amd64.tar.gz carrion

# 3) Build the Windows binary + zip
build-windows:
	GOOS=windows GOARCH=amd64 go build -o carrion.exe ./src
	zip carrion_windows_amd64.zip carrion.exe

# Existing Docker image build
build:
	@echo "Building Docker image with tags:"
	@echo "  - $(IMAGE_NAME):$(VERSION)"
	@echo "  - $(IMAGE_NAME):latest"
	docker build \
		-t "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)" \
		-t "$(USER_NAME)/$(IMAGE_NAME):latest" \
		.

# Existing Docker push
push:
	@echo "Pushing Docker image to Docker Hub:"
	@echo "  - $(IMAGE_NAME):$(VERSION)"
	@echo "  - $(IMAGE_NAME):latest"
	docker push "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)"
	docker push "$(USER_NAME)/$(IMAGE_NAME):latest"

# (Optional) run, clean, install, uninstall remain unchanged
run:
	docker run --rm -it "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)" bash

clean:
	docker rmi -f "$(USER_NAME)/$(IMAGE_NAME):$(VERSION)" || true
	docker rmi -f "$(USER_NAME)/$(IMAGE_NAME):latest" || true

install:
	@echo "Installing Carrion Language...."
	@./setup.sh
	@./install/install.sh "$(OS)"

uninstall:
	@echo "Uninstalling Carrion from disk..."
	@./install/uninstall.sh

