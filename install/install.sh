#!/usr/bin/env bash
# install.sh - Build and install the Carrion Language runtime as 'carrion'
# so you can run 'carrion file.crl' from any directory.
set -e  # Exit on any error

# 1) Build the Go binary
go build -o carrion ./src

# 2) Move it into /usr/local/bin, so 'carrion' is on the PATH
sudo mv carrion /usr/local/bin/carrion
# 3) Ensure it's executable (it usually is already, but just to be safe)
sudo chmod +x /usr/local/bin/carrion

echo "Carrion has been installed successfully!"
echo "You can now run: carrion yourfile.crl"
