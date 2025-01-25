
#!/usr/bin/env bash
set -e

# 1) Build the Go binary from ./src
go build -o carrion ./src

# 2) Move it into /usr/local/bin
mv carrion /usr/local/bin/carrion

# 3) Mark executable
chmod +x /usr/local/bin/carrion

echo "Carrion has been installed successfully!"
echo "You can now run: carrion yourfile.crl"
