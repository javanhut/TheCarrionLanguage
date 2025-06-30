#!/bin/bash

echo "🧪 Testing Carrion Playground API..."

# Test if API is running
echo "1. Checking API health..."
response=$(curl -s -w "%{http_code}" http://localhost:3001/health -o /tmp/health_response.json)
http_code=${response: -3}

if [ "$http_code" = "200" ]; then
    echo "✅ API is running!"
    cat /tmp/health_response.json
    echo ""
else
    echo "❌ API not responding (HTTP $http_code)"
    echo "Make sure to start the API first:"
    echo "  cd playground-api && npm start"
    exit 1
fi

# Test code execution
echo ""
echo "2. Testing code execution..."
test_code='print("Hello from API test!")'
response=$(curl -s -X POST http://localhost:3001/execute \
  -H "Content-Type: application/json" \
  -d "{\"code\":\"$test_code\"}" \
  -w "%{http_code}" \
  -o /tmp/execute_response.json)

http_code=${response: -3}

if [ "$http_code" = "200" ]; then
    echo "✅ Code execution works!"
    echo "Response:"
    cat /tmp/execute_response.json | jq .
else
    echo "❌ Code execution failed (HTTP $http_code)"
    echo "Response:"
    cat /tmp/execute_response.json
fi

# Cleanup
rm -f /tmp/health_response.json /tmp/execute_response.json

echo ""
echo "🔗 Frontend should connect to: http://localhost:3001"
echo "🌐 Open playground at: http://localhost:3000/playground"