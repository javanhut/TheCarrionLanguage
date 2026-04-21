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

# Carrion binary version metadata injected via -ldflags so the built binary can
# report exactly what source it came from. CARRION_VERSION is parsed from the
# single source of truth in src/version/version.go.
# CARRION_CHANNEL defaults to "release" because Makefile targets produce
# distributable artifacts (tarballs consumed by `carrion update`). Override
# with `CARRION_CHANNEL=dev` for local throwaway builds.
CARRION_VERSION := $(shell awk -F'"' '/^var Version =/{print $$2; exit}' src/version/version.go)
CARRION_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
CARRION_CHANNEL ?= release
CARRION_BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

CARRION_LDFLAGS := -X github.com/javanhut/TheCarrionLanguage/src/version.Version=$(CARRION_VERSION) \
                   -X github.com/javanhut/TheCarrionLanguage/src/version.Commit=$(CARRION_COMMIT) \
                   -X github.com/javanhut/TheCarrionLanguage/src/version.Channel=$(CARRION_CHANNEL) \
                   -X github.com/javanhut/TheCarrionLanguage/src/version.BuildDate=$(CARRION_BUILD_DATE)

.PHONY: build push run clean install uninstall build-source build-linux build-linux-arm64 build-windows build-mac build-mac-amd64 build-mac-arm64 build-release bifrost-update tidy bench test sync-version version-check

# 1) Build a tarball of the uncompiled source
build-source:
	git archive --format=tar.gz -o carrion-src.tar.gz HEAD	

# 2) Build the Linux binary + tarball (amd64)
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(CARRION_LDFLAGS)" -o carrion ./src
	cd cmd/sindri && GOOS=linux GOARCH=amd64 go build -o sindri .
	cd cmd/mimir && GOOS=linux GOARCH=amd64 go build -o mimir .
	tar -czf carrion_linux_amd64.tar.gz carrion cmd/sindri/sindri cmd/mimir/mimir

# 2b) Linux arm64 variant (for Raspberry Pi, Graviton, arm64 CI runners)
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "$(CARRION_LDFLAGS)" -o carrion ./src
	cd cmd/sindri && GOOS=linux GOARCH=arm64 go build -o sindri .
	cd cmd/mimir && GOOS=linux GOARCH=arm64 go build -o mimir .
	tar -czf carrion_linux_arm64.tar.gz carrion cmd/sindri/sindri cmd/mimir/mimir

# 3) Build the Windows binary + zip
build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "$(CARRION_LDFLAGS)" -o carrion.exe ./src
	cd cmd/sindri && GOOS=windows GOARCH=amd64 go build -o sindri.exe .
	cd cmd/mimir && GOOS=windows GOARCH=amd64 go build -o mimir.exe .
	zip carrion_windows_amd64.zip carrion.exe cmd/sindri/sindri.exe cmd/mimir/mimir.exe

# 4) macOS tarballs (both Intel and Apple Silicon)
build-mac: build-mac-amd64 build-mac-arm64

build-mac-amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(CARRION_LDFLAGS)" -o carrion ./src
	cd cmd/sindri && GOOS=darwin GOARCH=amd64 go build -o sindri .
	cd cmd/mimir && GOOS=darwin GOARCH=amd64 go build -o mimir .
	tar -czf carrion_darwin_amd64.tar.gz carrion cmd/sindri/sindri cmd/mimir/mimir

build-mac-arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(CARRION_LDFLAGS)" -o carrion ./src
	cd cmd/sindri && GOOS=darwin GOARCH=arm64 go build -o sindri .
	cd cmd/mimir && GOOS=darwin GOARCH=arm64 go build -o mimir .
	tar -czf carrion_darwin_arm64.tar.gz carrion cmd/sindri/sindri cmd/mimir/mimir

# 4b) Convenience: build every release artifact in one shot
build-release: build-source build-linux build-linux-arm64 build-windows build-mac
	@echo "Release artifacts:"
	@ls -1 carrion_*_*.tar.gz carrion_*_*.zip carrion-src.tar.gz 2>/dev/null || true

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
	@echo "Installing Carrion Language, Sindri Testing Framework, Mimir Documentation Tool, and Bifrost Package Manager for $(OS)...."
	@./setup.sh
	@./install/install.sh "$(OS)"
	@if [ -d "bifrost" ]; then \
		echo "Installing Bifrost Package Manager..."; \
		cd bifrost && make install; \
	else \
		echo "Bifrost directory not found, skipping Bifrost installation"; \
	fi

uninstall:
	@echo "Uninstalling Carrion, Sindri, Mimir, and Bifrost from disk..."
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

bench:
	@echo "Running evaluator benchmarks..."
	go test -bench=. -benchmem -count=1 ./src/evaluator/

test:
	@echo "Running all tests..."
	go test ./src/...

# Keep version references in documentation aligned with the single source of
# truth in src/version/version.go. Run sync-version to rewrite stale mentions;
# run version-check in CI to fail the build when docs drift.
sync-version:
	@go run ./cmd/versionsync

version-check:
	@go run ./cmd/versionsync --check

tidy:
	@echo "Running go mod tidy on all modules..."
	@go mod tidy & \
	cd cmd/sindri && go mod tidy & \
	cd cmd/mimir && go mod tidy & \
	cd bifrost && go mod tidy & \
	wait
	@echo "All modules tidied."
