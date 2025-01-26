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
    
    echo "Moving binary to /usr/local/bin..."
    sudo mv carrion /usr/local/bin/carrion
    sudo chmod +x /usr/local/bin/carrion
    
    echo "The Carrion Programming Language has been installed successfully on Linux!"
    echo "You can now run the interactive REPL by typing: carrion"
    ;;

  windows)
    echo "Building Carrion for Windows..."
    # Cross-compile for Windows on amd64. Adjust GOARCH if needed (e.g., arm64).
    GOOS=windows GOARCH=amd64 go build -o carrion.exe ./src
    
    echo "A 'carrion.exe' file has been created."
    echo "On Windows, place carrion.exe in a directory on your PATH (e.g., C:\\Windows\\System32)"
    echo "or simply run it directly in your terminal:"
    echo "  .\\carrion.exe"
    ;;

  mac)
    echo "Building Carrion for macOS..."
    # Cross-compile for Darwin on amd64. Adjust GOARCH if you're on Apple Silicon (e.g., arm64).
    GOOS=darwin GOARCH=amd64 go build -o carrion ./src
    
    echo "Moving binary to /usr/local/bin..."
    sudo mv carrion /usr/local/bin/carrion
    sudo chmod +x /usr/local/bin/carrion
    
    echo "The Carrion Programming Language has been installed successfully on macOS!"
    echo "You can now run the interactive REPL by typing: carrion"
    ;;

  *)
    echo "Unsupported OS: $TARGET_OS"
    echo "Valid options are 'linux', 'windows', or 'mac'."
    exit 1
    ;;
esac

