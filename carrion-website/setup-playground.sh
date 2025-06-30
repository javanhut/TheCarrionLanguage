#!/bin/bash

echo "🐦‍⬛ Setting up Carrion Playground..."

# Check if Podman is installed
if ! command -v podman &> /dev/null; then
    echo "❌ Podman is not installed. Please install Podman first."
    echo "Visit: https://podman.io/getting-started/installation"
    echo "On Ubuntu/Debian: sudo apt-get install podman"
    echo "On RHEL/CentOS: sudo dnf install podman"
    echo "On macOS: brew install podman"
    exit 1
fi

echo "✅ Podman is available"

# Pull the Carrion image
echo "📦 Pulling Carrion image..."
if podman pull javanhut/carrionlanguage:latest; then
    echo "✅ Carrion image pulled successfully"
else
    echo "❌ Failed to pull Carrion image"
    echo "Make sure you have internet access and the image exists"
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 16+ first."
    echo "Visit: https://nodejs.org/"
    exit 1
fi

echo "✅ Node.js is installed"

# Install API dependencies
echo "📦 Installing API dependencies..."
cd playground-api
if npm install; then
    echo "✅ API dependencies installed"
else
    echo "❌ Failed to install API dependencies"
    exit 1
fi

# Return to root directory
cd ..

# Install frontend dependencies (if not already done)
echo "📦 Installing frontend dependencies..."
if npm install; then
    echo "✅ Frontend dependencies installed"
else
    echo "❌ Failed to install frontend dependencies"
    exit 1
fi

echo ""
echo "🎉 Setup complete! To start the playground:"
echo ""
echo "1. Start the API server:"
echo "   cd playground-api && npm start"
echo ""
echo "2. In a new terminal, start the frontend:"
echo "   npm start"
echo ""
echo "3. Open http://localhost:3000 in your browser"
echo ""
echo "🔒 Security features:"
echo "   - Isolated Podman containers"
echo "   - 64MB memory limit"
echo "   - 10-second execution timeout"
echo "   - No network access"
echo "   - Read-only filesystem"
echo ""
echo "Happy coding with Carrion! 🐦‍⬛✨"