#!/usr/bin/env bash
#
# setup.sh - Checks for Go, sets up permission on install scripts, then runs install.sh
#

set -e  # Exit on any error

# 1) Check if Go is installed
if ! command -v go &>/dev/null; then
    echo "Error: Go is not installed or not on your PATH."
    echo "Please install Go and try again."
    exit 1
fi

# 2) Make install scripts executable
chmod +x install/install.sh
chmod +x install/uninstall.sh

# 3) Initialize and update Bifrost submodule (if in git repository)
if command -v git &>/dev/null && [ -d ".git" ]; then
    echo "Initializing Bifrost submodule..."
    git submodule update --init --recursive
else
    echo "Skipping git submodule update (not in a git repository or git not installed)"
fi

