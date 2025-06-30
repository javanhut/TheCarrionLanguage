# Carrion Playground API

A secure backend API for executing Carrion code in a sandboxed Docker environment.

## Features

- **Secure Execution**: Runs code in isolated Docker containers
- **Resource Limits**: Memory (64MB) and CPU (50%) constraints
- **Timeout Protection**: 10-second execution limit
- **Network Isolation**: No network access for security
- **Automatic Cleanup**: Containers are automatically removed after execution

## Setup

1. **Install Dependencies**:
   ```bash
   cd playground-api
   npm install
   ```

2. **Start Docker** (required):
   ```bash
   # Make sure Docker is running
   docker --version
   
   # Pull the Carrion image
   docker pull javanhut/carrionlanguage:latest
   ```

3. **Start the API**:
   ```bash
   # Development mode
   npm run dev
   
   # Production mode
   npm start
   ```

## API Endpoints

### Health Check
```
GET /health
```
Response:
```json
{
  "status": "ok",
  "message": "Carrion Playground API is running"
}
```

### Execute Code
```
POST /execute
Content-Type: application/json

{
  "code": "print(\"Hello, Carrion!\")"
}
```

Response (Success):
```json
{
  "success": true,
  "output": "Hello, Carrion!\n",
  "stderr": "",
  "exitCode": 0
}
```

Response (Error):
```json
{
  "success": false,
  "output": "",
  "stderr": "Error: Syntax error on line 1\n",
  "exitCode": 1
}
```

## Security Features

- **Container Isolation**: Each execution runs in a fresh Docker container
- **Resource Limits**: Memory and CPU constraints prevent resource exhaustion
- **Network Disabled**: Containers have no network access
- **Timeout Protection**: Executions are limited to 10 seconds
- **Input Validation**: Code length limited to 10,000 characters
- **Automatic Cleanup**: Containers are force-removed after execution

## Environment Variables

- `PORT`: API port (default: 3001)
- `DOCKER_HOST`: Docker daemon host (default: local socket)

## Requirements

- Node.js 16+
- Docker with `javanhut/carrionlanguage:latest` image
- Docker daemon accessible to the API process

## Production Deployment

For production deployment, consider:

1. **Reverse Proxy**: Use nginx or similar for HTTPS and load balancing
2. **Process Manager**: Use PM2 or similar for process management
3. **Monitoring**: Add logging and monitoring for container executions
4. **Rate Limiting**: Implement rate limiting to prevent abuse
5. **Authentication**: Add authentication if needed

Example nginx configuration:
```nginx
location /api/ {
    proxy_pass http://localhost:3001/;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
}
```