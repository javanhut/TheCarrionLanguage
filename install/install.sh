#!/usr/bin/env bash
#
# install.sh - Install the Carrion Programming Language
#
# Usage:
#   ./install.sh <os>
#   Example: ./install.sh linux
#            ./install.sh windows
#            ./install.sh mac
#
# This script expects a single argument that specifies the target OS.
# It requires Go to be installed and accessible in your PATH.
#
set -e  # Exit the script if any command fails

# --- 1) Check Arguments ---
if [ $# -ne 1 ]; then
  echo "Usage: $0 <os>"
  echo "  <os> can be 'linux', 'windows', or 'mac'"
  exit 1
fi

TARGET_OS=$1

# --- 2) Check if Go is Installed ---
if ! command -v go &> /dev/null; then
  echo "Error: 'go' is not installed or not in your PATH."
  echo "Please install Go and rerun this script."
  exit 1
fi

# --- 3) Build & Install Logic ---
case "$TARGET_OS" in
  linux)
    echo "Building Carrion for Linux..."
    # Cross-compile for Linux on amd64. Adjust GOARCH if needed (e.g., arm64).
    GOOS=linux GOARCH=amd64 go build -o carrion ./src
    
    echo "Building Sindri Testing Framework for Linux..."
    cd cmd/sindri
    GOOS=linux GOARCH=amd64 go build -o sindri .
    cd ../..
    
    echo "Building Mimir Documentation Tool for Linux..."
    cd cmd/mimir
    GOOS=linux GOARCH=amd64 go build -o mimir .
    cd ../..

    echo "Building Bifrost Package Manager for Linux..."
    if [ -d "bifrost" ]; then
      cd bifrost
      make build
      cd ..
    else
      echo "Warning: Bifrost submodule not found. Skipping Bifrost installation."
      echo "Run 'git submodule update --init --recursive' to get Bifrost."
    fi

    echo "Moving binaries to /usr/local/bin..."
    sudo mv carrion /usr/local/bin/carrion
    sudo chmod +x /usr/local/bin/carrion
    sudo mv cmd/sindri/sindri /usr/local/bin/sindri
    sudo chmod +x /usr/local/bin/sindri
    sudo mv cmd/mimir/mimir /usr/local/bin/mimir
    sudo chmod +x /usr/local/bin/mimir

    if [ -f "bifrost/build/bifrost" ]; then
      sudo mv bifrost/build/bifrost /usr/local/bin/bifrost
      sudo chmod +x /usr/local/bin/bifrost
      echo "The Carrion Programming Language, Sindri Testing Framework, Mimir Documentation Tool, and Bifrost Package Manager have been installed successfully on Linux!"
      echo "You can now run:"
      echo "  - Interactive REPL: carrion"
      echo "  - Test runner: sindri appraise test_file.crl"
      echo "  - Documentation: mimir"
      echo "  - Package manager: bifrost"
    else
      echo "The Carrion Programming Language, Sindri Testing Framework, and Mimir Documentation Tool have been installed successfully on Linux!"
      echo "You can now run:"
      echo "  - Interactive REPL: carrion"
      echo "  - Test runner: sindri appraise test_file.crl"
      echo "  - Documentation: mimir"
    fi
    ;;

  windows)
    echo "Building Carrion for Windows..."
    # Cross-compile for Windows on amd64. Adjust GOARCH if needed (e.g., arm64).
    GOOS=windows GOARCH=amd64 go build -o carrion.exe ./src
    
    echo "Building Sindri Testing Framework for Windows..."
    cd cmd/sindri
    GOOS=windows GOARCH=amd64 go build -o sindri.exe .
    cd ../..
    
    echo "Building Mimir Documentation Tool for Windows..."
    cd cmd/mimir
    GOOS=windows GOARCH=amd64 go build -o mimir.exe .
    cd ../..

    echo "Building Bifrost Package Manager for Windows..."
    if [ -d "bifrost" ]; then
      cd bifrost
      GOOS=windows GOARCH=amd64 go build -o build/bifrost.exe ./cmd/bifrost
      cd ..
    else
      echo "Warning: Bifrost submodule not found. Skipping Bifrost build."
      echo "Run 'git submodule update --init --recursive' to get Bifrost."
    fi

    if [ -f "bifrost/build/bifrost.exe" ]; then
      echo "Binaries 'carrion.exe', 'sindri.exe', 'mimir.exe', and 'bifrost.exe' have been created."
      echo "On Windows, place all files in a directory on your PATH (e.g., C:\\Windows\\System32)"
      echo "or simply run them directly in your terminal:"
      echo "  .\\carrion.exe"
      echo "  .\\sindri.exe appraise test_file.crl"
      echo "  .\\mimir.exe"
      echo "  .\\bifrost.exe"
    else
      echo "Binaries 'carrion.exe', 'sindri.exe', and 'mimir.exe' have been created."
      echo "On Windows, place all files in a directory on your PATH (e.g., C:\\Windows\\System32)"
      echo "or simply run them directly in your terminal:"
      echo "  .\\carrion.exe"
      echo "  .\\sindri.exe appraise test_file.crl"
      echo "  .\\mimir.exe"
    fi
    ;;

  mac)
    echo "Building Carrion for macOS..."
    # Cross-compile for Darwin on amd64. Adjust GOARCH if you're on Apple Silicon (e.g., arm64).
    GOOS=darwin GOARCH=amd64 go build -o carrion ./src
    
    echo "Building Sindri Testing Framework for macOS..."
    cd cmd/sindri
    GOOS=darwin GOARCH=amd64 go build -o sindri .
    cd ../..
    
    echo "Building Mimir Documentation Tool for macOS..."
    cd cmd/mimir
    GOOS=darwin GOARCH=amd64 go build -o mimir .
    cd ../..

    echo "Building Bifrost Package Manager for macOS..."
    if [ -d "bifrost" ]; then
      cd bifrost
      make build
      cd ..
    else
      echo "Warning: Bifrost submodule not found. Skipping Bifrost installation."
      echo "Run 'git submodule update --init --recursive' to get Bifrost."
    fi

    echo "Moving binaries to /usr/local/bin..."
    sudo mv carrion /usr/local/bin/carrion
    sudo chmod +x /usr/local/bin/carrion
    sudo mv cmd/sindri/sindri /usr/local/bin/sindri
    sudo chmod +x /usr/local/bin/sindri
    sudo mv cmd/mimir/mimir /usr/local/bin/mimir
    sudo chmod +x /usr/local/bin/mimir

    if [ -f "bifrost/build/bifrost" ]; then
      sudo mv bifrost/build/bifrost /usr/local/bin/bifrost
      sudo chmod +x /usr/local/bin/bifrost
      echo "The Carrion Programming Language, Sindri Testing Framework, Mimir Documentation Tool, and Bifrost Package Manager have been installed successfully on macOS!"
      echo "You can now run:"
      echo "  - Interactive REPL: carrion"
      echo "  - Test runner: sindri appraise test_file.crl"
      echo "  - Documentation: mimir"
      echo "  - Package manager: bifrost"
    else
      echo "The Carrion Programming Language, Sindri Testing Framework, and Mimir Documentation Tool have been installed successfully on macOS!"
      echo "You can now run:"
      echo "  - Interactive REPL: carrion"
      echo "  - Test runner: sindri appraise test_file.crl"
      echo "  - Documentation: mimir"
    fi
    ;;

  *)
    echo "Unsupported OS: $TARGET_OS"
    echo "Valid options are 'linux', 'windows', or 'mac'."
    exit 1
    ;;
esac

