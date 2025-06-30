#!/bin/bash

echo "🐦‍⬛ Starting Carrion Playground..."

# Function to cleanup background processes
cleanup() {
    echo ""
    echo "🛑 Shutting down playground..."
    if [ ! -z "$API_PID" ]; then
        kill $API_PID 2>/dev/null
        echo "✅ API server stopped"
    fi
    if [ ! -z "$FRONTEND_PID" ]; then
        kill $FRONTEND_PID 2>/dev/null
        echo "✅ Frontend server stopped"
    fi
    exit 0
}

# Set up cleanup on script exit
trap cleanup SIGINT SIGTERM EXIT

# Check if Podman is available
if ! command -v podman &> /dev/null; then
    echo "❌ Podman is not installed. Please install Podman first."
    echo "Visit: https://podman.io/getting-started/installation"
    exit 1
fi

echo "✅ Podman is available"

# Check if Carrion image exists
if ! podman images | grep -q "javanhut/carrionlanguage"; then
    echo "📦 Pulling Carrion image..."
    podman pull javanhut/carrionlanguage:latest
fi

# Check if API dependencies are installed
if [ ! -d "playground-api/node_modules" ]; then
    echo "📦 Installing API dependencies..."
    cd playground-api && npm install && cd ..
fi

# Check if frontend dependencies are installed
if [ ! -d "node_modules" ]; then
    echo "📦 Installing frontend dependencies..."
    npm install
fi

echo "🚀 Starting API server..."
cd playground-api
npm start &
API_PID=$!
cd ..

# Wait a moment for the API to start
sleep 3

echo "🌐 Starting frontend server..."
npm start &
FRONTEND_PID=$!

echo ""
echo "✅ Playground is starting up!"
echo ""
echo "📍 Frontend: http://localhost:3000"
echo "📍 API: http://localhost:3001"
echo "📍 Playground: http://localhost:3000/playground"
echo ""
echo "🔒 Security features active:"
echo "   - Podman container isolation"
echo "   - 64MB memory limit"
echo "   - 10-second timeout"
echo "   - No network access"
echo "   - Read-only filesystem"
echo ""
echo "Press Ctrl+C to stop both servers"
echo ""

# Wait for both processes
wait