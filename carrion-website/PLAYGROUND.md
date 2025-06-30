# 🐦‍⬛ Carrion Playground Setup

This guide explains how to set up and run the **real Carrion code execution playground** that uses the actual Carrion interpreter via Docker.

## 🎯 Overview

The playground consists of two components:

1. **Frontend** (React): The web interface with code editor and output display
2. **Backend API** (Node.js): Executes Carrion code securely using Docker containers

## 🔧 Prerequisites

- **Docker**: For running Carrion code in isolated containers
- **Node.js 16+**: For the frontend and backend
- **Internet connection**: To pull the Carrion Docker image

## 🚀 Quick Setup & Start

### Option 1: One-Command Start (Recommended)

```bash
./start-playground.sh
```

This will:
- Check Docker and pull the Carrion image
- Install dependencies if needed
- Start both API server (port 3001) and frontend (port 3000)
- Open playground at http://localhost:3000/playground

### Option 2: Setup Only

```bash
./setup-playground.sh
```

This will only set up dependencies and provide manual start instructions.

## 📋 Manual Setup

### 1. Install Docker and Pull Carrion Image

```bash
# Check Docker is running
docker --version

# Pull the Carrion Docker image
docker pull javanhut/carrionlanguage:latest
```

### 2. Install Dependencies

```bash
# Install frontend dependencies
npm install

# Install API dependencies
cd playground-api
npm install
cd ..
```

## 🏃‍♀️ Running the Playground

### Start the Backend API (Terminal 1)

```bash
cd playground-api
npm start
```

The API will start on `http://localhost:3001`

### Start the Frontend (Terminal 2)

```bash
npm start
```

The website will open at `http://localhost:3000`

## 🔒 Security Features

The playground includes several security measures:

- **Container Isolation**: Each code execution runs in a fresh Docker container
- **Resource Limits**: 64MB memory limit and 50% CPU quota
- **Network Disabled**: Containers have no network access
- **Execution Timeout**: 10-second maximum execution time
- **Input Validation**: Code length limited to 10,000 characters
- **Automatic Cleanup**: Containers are force-removed after execution

## 🐛 Troubleshooting

### "Connection Error: Unexpected token '<'" 

This error means the frontend can't connect to the API server:

1. **Start the API server first**:
   ```bash
   cd playground-api
   npm start
   ```

2. **Check API is running**:
   ```bash
   ./test-api.sh
   ```

3. **Verify API responds**:
   ```bash
   curl http://localhost:3001/health
   ```

### API Not Available

If you see "⚠️ API currently unavailable - using simulation mode":

1. Make sure the backend API is running on port 3001
2. Check Docker is running and has the Carrion image
3. Verify no firewall is blocking port 3001
4. Use `./test-api.sh` to diagnose the issue

### Docker Issues

```bash
# Check Docker status
docker info

# Test Carrion image manually
docker run -it javanhut/carrionlanguage:latest

# Check available images
docker images | grep carrion
```

### Port Conflicts

If port 3001 is in use:

```bash
# Check what's using port 3001
lsof -i :3001

# Or change the port in playground-api/server.js
# Set PORT environment variable
PORT=3002 npm start
```

## 🔧 Development

### API Development

```bash
cd playground-api
npm run dev  # Uses nodemon for auto-restart
```

### Frontend Development

```bash
npm start  # React development server
```

### Testing API Directly

```bash
# Health check
curl http://localhost:3001/health

# Execute code
curl -X POST http://localhost:3001/execute \
  -H "Content-Type: application/json" \
  -d '{"code":"print(\"Hello, Carrion!\")"}'
```

## 📦 Production Deployment

For production deployment:

1. **Build Frontend**:
   ```bash
   npm run build
   ```

2. **Deploy API** with process manager:
   ```bash
   cd playground-api
   npm install pm2 -g
   pm2 start server.js --name carrion-api
   ```

3. **Reverse Proxy** (nginx example):
   ```nginx
   location /api/ {
       proxy_pass http://localhost:3001/;
       proxy_set_header Host $host;
       proxy_set_header X-Real-IP $remote_addr;
   }
   ```

## 🎮 Using the Playground

1. **Write Code**: Use the editor to write Carrion code
2. **Load Examples**: Click example buttons to load pre-written programs
3. **Run Code**: Click "Run" to execute code with the real Carrion interpreter
4. **View Output**: See actual program output, errors, and execution results

### Example Programs Included

- **Hello World**: Basic program structure and print statements
- **Variables & Types**: Working with different data types
- **Grimoires (Classes)**: Object-oriented programming examples
- **Control Flow**: Conditionals and loops with Carrion syntax
- **Error Handling**: `attempt`/`ensnare`/`resolve` blocks
- **Fibonacci**: Recursive function example

## 🤝 Contributing

To contribute to the playground:

1. Fork the repository
2. Make changes to frontend (`src/pages/Playground.tsx`) or backend (`playground-api/`)
3. Test locally using the setup above
4. Submit a pull request

## 📚 Related Documentation

- [Getting Started with Carrion](/docs/getting-started)
- [Language Features](/features)
- [Installation Guide](/docs/installation)
- [Carrion GitHub Repository](https://github.com/javanhut/TheCarrionLanguage)