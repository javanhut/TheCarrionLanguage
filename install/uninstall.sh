#!/usr/bin/env bash
# uninstall.sh - Remove the carrion binary from /usr/local/bin
set -e  # Exit script on error
if [ -f /usr/local/bin/carrion ]; then
    sudo rm /usr/local/bin/carrion
    echo "Carrion uninstalled successfully."
else
    echo "Carrion does not appear to be installed in /usr/local/bin."
fi
