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
    echo "Uninstalling Carrion, Sindri, and Mimir from Linux..."
    removed_count=0
    if [ -f /usr/local/bin/carrion ]; then
      sudo rm /usr/local/bin/carrion
      echo "Carrion uninstalled successfully from /usr/local/bin."
      removed_count=$((removed_count + 1))
    else
      echo "Carrion does not appear to be installed in /usr/local/bin."
    fi
    if [ -f /usr/local/bin/sindri ]; then
      sudo rm /usr/local/bin/sindri
      echo "Sindri uninstalled successfully from /usr/local/bin."
      removed_count=$((removed_count + 1))
    else
      echo "Sindri does not appear to be installed in /usr/local/bin."
    fi
    if [ -f /usr/local/bin/mimir ]; then
      sudo rm /usr/local/bin/mimir
      echo "Mimir uninstalled successfully from /usr/local/bin."
      removed_count=$((removed_count + 1))
    else
      echo "Mimir does not appear to be installed in /usr/local/bin."
    fi
    if [ $removed_count -eq 0 ]; then
      echo "No Carrion components found to uninstall."
    fi
    ;;

  windows)
    echo "Uninstalling Carrion, Sindri, and Mimir from Windows..."
    removed_count=0
    if [ -f "carrion.exe" ]; then
      rm "carrion.exe"
      echo "Carrion.exe removed from the current directory."
      removed_count=$((removed_count + 1))
    else
      echo "Carrion.exe not found in the current directory."
    fi
    if [ -f "sindri.exe" ]; then
      rm "sindri.exe"
      echo "Sindri.exe removed from the current directory."
      removed_count=$((removed_count + 1))
    else
      echo "Sindri.exe not found in the current directory."
    fi
    if [ -f "mimir.exe" ]; then
      rm "mimir.exe"
      echo "Mimir.exe removed from the current directory."
      removed_count=$((removed_count + 1))
    else
      echo "Mimir.exe not found in the current directory."
    fi
    if [ $removed_count -eq 0 ]; then
      echo "No Carrion components found in current directory."
      echo "If they were placed elsewhere, please remove them manually."
    fi
    ;;

  darwin)
    # Example logic for macOS if you have it
    echo "Uninstalling Carrion, Sindri, and Mimir from macOS..."
    removed_count=0
    if [ -f /usr/local/bin/carrion ]; then
      sudo rm /usr/local/bin/carrion
      echo "Carrion uninstalled successfully from /usr/local/bin."
      removed_count=$((removed_count + 1))
    else
      echo "Carrion does not appear to be installed in /usr/local/bin."
    fi
    if [ -f /usr/local/bin/sindri ]; then
      sudo rm /usr/local/bin/sindri
      echo "Sindri uninstalled successfully from /usr/local/bin."
      removed_count=$((removed_count + 1))
    else
      echo "Sindri does not appear to be installed in /usr/local/bin."
    fi
    if [ -f /usr/local/bin/mimir ]; then
      sudo rm /usr/local/bin/mimir
      echo "Mimir uninstalled successfully from /usr/local/bin."
      removed_count=$((removed_count + 1))
    else
      echo "Mimir does not appear to be installed in /usr/local/bin."
    fi
    if [ $removed_count -eq 0 ]; then
      echo "No Carrion components found to uninstall."
    fi
    ;;

  *)
    echo "Error: OS '$OS' is not supported by this uninstall script."
    exit 1
    ;;
esac

echo "Uninstall script completed."

