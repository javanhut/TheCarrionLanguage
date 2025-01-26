#!/usr/bin/env bash
#
# uninstall.sh - Uninstall the Carrion Programming Language
#
# Usage:
#   ./uninstall.sh <os>
#   Example: ./uninstall.sh linux
#            ./uninstall.sh windows
#
# If no <os> argument is provided, script attempts to detect OS (basic approach).
#
set -e  # Exit script on error

# --- 1) OS Detection Helpers (if no argument is given) ---
detect_os() {
  # Basic detection. Adjust to your environment as needed.
  case "$(uname -s)" in
    Linux*)   echo "linux" ;;
    Darwin*)  echo "darwin" ;;  # For macOS (example)
    CYGWIN*|MINGW*|MSYS*) echo "windows" ;;
    *)        echo "unknown" ;;
  esac
}

# --- 2) Parse Argument or Auto-detect ---
if [ $# -eq 1 ]; then
  OS=$1
elif [ $# -eq 0 ]; then
  # Attempt to auto-detect if no argument is passed
  OS=$(detect_os)
  echo "No <os> argument provided. Detected OS: $OS"
else
  echo "Usage: $0 [os]"
  echo "  where [os] is 'linux', 'windows', etc."
  exit 1
fi

# --- 3) Uninstall Logic based on OS ---
case "$OS" in

  linux)
    echo "Uninstalling Carrion from Linux..."
    if [ -f /usr/local/bin/carrion ]; then
      sudo rm /usr/local/bin/carrion
      echo "Carrion uninstalled successfully from /usr/local/bin."
    else
      echo "Carrion does not appear to be installed in /usr/local/bin."
    fi
    ;;

  windows)
    echo "Uninstalling Carrion from Windows..."
    # Typically, on Windows, you might not have used `sudo` or `/usr/local/bin`.
    # Here’s one example approach:
    if [ -f "carrion.exe" ]; then
      rm "carrion.exe"
      echo "Carrion.exe removed from the current directory."
    else
      echo "Carrion.exe not found in the current directory."
      echo "If Carrion was placed elsewhere, please remove it manually."
    fi
    ;;

  darwin)
    # Example logic for macOS if you have it
    echo "Uninstalling Carrion from macOS..."
    if [ -f /usr/local/bin/carrion ]; then
      sudo rm /usr/local/bin/carrion
      echo "Carrion uninstalled successfully from /usr/local/bin."
    else
      echo "Carrion does not appear to be installed in /usr/local/bin."
    fi
    ;;

  *)
    echo "Error: OS '$OS' is not supported by this uninstall script."
    exit 1
    ;;
esac

echo "Uninstall script completed."

