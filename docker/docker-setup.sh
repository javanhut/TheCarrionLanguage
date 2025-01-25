
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
chmod +x docker/docker-install.sh

# 3) Run install.sh
./docker/docker-install.sh

