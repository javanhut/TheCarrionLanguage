# Makefile snippet

# Existing variables
USER_NAME ?= username
IMAGE_NAME ?= carrionlanguage
VERSION ?= latest

# Auto-detect OS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	DETECTED_OS ?= linux
endif
ifeq ($(UNAME_S),Darwin)
	DETECTED_OS ?= mac
endif
ifneq ($(filter MINGW%,$(UNAME_S)),)
	DETECTED_OS ?= windows
endif

# Set default or error if OS detection failed
ifndef DETECTED_OS
	DETECTED_OS := unsupported
endif

OS ?= $(DETECTED_OS)

# Validate OS value
ifeq ($(OS),unsupported)
	$(error Unsupported operating system: $(UNAME_S). Please set OS manually to one of: linux, mac, windows)
endif

.PHONY: build push run clean install uninstall build-source build-linux build-windows bifrost-update

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
	@echo "Installing Carrion Language and Bifrost Package Manager for $(OS)...."
	@./setup.sh
	@./install/install.sh "$(OS)"
	@if [ -d "bifrost" ]; then \
		echo "Installing Bifrost Package Manager..."; \
		cd bifrost && make install; \
	else \
		echo "Bifrost directory not found, skipping Bifrost installation"; \
	fi

uninstall:
	@echo "Uninstalling Carrion and Bifrost from disk..."
	@./install/uninstall.sh
	@if [ -d "bifrost" ]; then \
		echo "Uninstalling Bifrost Package Manager..."; \
		$(MAKE) -C bifrost uninstall; \
	else \
		echo "Bifrost directory not found, skipping Bifrost uninstallation"; \
	fi

bifrost-update:
	@echo "Updating Bifrost submodule..."
	@git submodule update --init --recursive
	@git submodule update --remote

