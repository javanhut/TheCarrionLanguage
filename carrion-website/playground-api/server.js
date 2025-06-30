const express = require('express');
const cors = require('cors');
const { spawn, exec } = require('child_process');
const fs = require('fs').promises;
const path = require('path');
const crypto = require('crypto');

const app = express();
const port = process.env.PORT || 3001;

// Middleware
app.use(cors());
app.use(express.json({ limit: '10mb' }));

// Health check endpoint
app.get('/health', (req, res) => {
  res.json({ status: 'ok', message: 'Carrion Playground API is running' });
});

// Code execution endpoint
app.post('/execute', async (req, res) => {
  const { code } = req.body;
  
  if (!code || typeof code !== 'string') {
    return res.status(400).json({ error: 'Code is required and must be a string' });
  }

  // Security: Limit code length
  if (code.length > 10000) {
    return res.status(400).json({ error: 'Code too long (max 10,000 characters)' });
  }

  try {
    const result = await executeCarrionCode(code);
    res.json(result);
  } catch (error) {
    console.error('Execution error:', error);
    res.status(500).json({ 
      error: 'Execution failed',
      message: error.message,
      output: '',
      stderr: error.message
    });
  }
});

async function executeCarrionCode(code) {
  // Generate unique session ID
  const sessionId = crypto.randomBytes(8).toString('hex');
  const workDir = `/tmp/carrion-${sessionId}`;
  const codePath = path.join(workDir, 'main.crl');

  try {
    // Create temporary directory and write code
    await fs.mkdir(workDir, { recursive: true });
    await fs.writeFile(codePath, code, 'utf8');

    // Check if podman is available
    const podmanCheck = await checkPodman();
    if (!podmanCheck.available) {
      throw new Error(`Podman not available: ${podmanCheck.error}`);
    }

    // Pull image if not exists
    await ensureCarrionImage();

    // Execute code using podman
    const result = await runPodmanContainer(workDir, sessionId);
    
    return result;

  } finally {
    // Cleanup temporary files
    try {
      await fs.rm(workDir, { recursive: true, force: true });
    } catch (err) {
      console.error('Cleanup error:', err);
    }
  }
}

function checkPodman() {
  return new Promise((resolve) => {
    exec('podman --version', (error, stdout, stderr) => {
      if (error) {
        resolve({ available: false, error: `Podman not found: ${error.message}` });
      } else {
        resolve({ available: true, version: stdout.trim() });
      }
    });
  });
}

function ensureCarrionImage() {
  return new Promise((resolve, reject) => {
    // First check if image exists
    exec('podman images docker.io/javanhut/carrionlanguage:latest --format "{{.Repository}}"', (error, stdout) => {
      if (stdout.trim()) {
        // Image exists
        resolve();
      } else {
        // Pull image
        console.log('Pulling Carrion image...');
        exec('podman pull docker.io/javanhut/carrionlanguage:latest', { timeout: 60000 }, (error, stdout, stderr) => {
          if (error) {
            reject(new Error(`Failed to pull image: ${error.message}`));
          } else {
            console.log('Carrion image pulled successfully');
            resolve();
          }
        });
      }
    });
  });
}

function runPodmanContainer(workDir, sessionId) {
  return new Promise((resolve, reject) => {
    const containerName = `carrion-exec-${sessionId}`;
    
    // Podman run command with security restrictions
    const podmanArgs = [
      'run',
      '--rm',                          // Remove container after execution
      '--name', containerName,         // Container name
      '--network', 'none',             // No network access
      '--memory', '64m',               // 64MB memory limit
      '--cpus', '0.5',                 // 50% CPU limit
      '--security-opt', 'no-new-privileges', // No privilege escalation
      '--read-only',                   // Read-only filesystem
      '--tmpfs', '/tmp',               // Writable tmp
      '--volume', `${workDir}:/app:ro`, // Mount code directory read-only
      '--workdir', '/app',             // Set working directory
      '--user', '1001:1001',           // Run as non-root user
      'docker.io/javanhut/carrionlanguage:latest', // Image
      'timeout', '10s',                // 10-second timeout
      'carrion', 'main.crl'            // Command
    ];

    const podman = spawn('podman', podmanArgs);
    
    let stdout = '';
    let stderr = '';
    let killed = false;

    // Set up timeout
    const timeout = setTimeout(() => {
      if (!killed) {
        killed = true;
        podman.kill('SIGKILL');
        reject(new Error('Execution timeout (10 seconds)'));
      }
    }, 12000); // 12 seconds to account for container startup

    podman.stdout.on('data', (data) => {
      stdout += data.toString();
    });

    podman.stderr.on('data', (data) => {
      stderr += data.toString();
    });

    podman.on('close', (code) => {
      clearTimeout(timeout);
      
      if (killed) return; // Already handled by timeout
      
      const success = code === 0;
      
      resolve({
        success,
        output: stdout,
        stderr: success ? '' : stderr,
        exitCode: code
      });
    });

    podman.on('error', (error) => {
      clearTimeout(timeout);
      reject(new Error(`Podman execution failed: ${error.message}`));
    });
  });
}

// Graceful shutdown
process.on('SIGTERM', () => {
  console.log('SIGTERM received, shutting down gracefully');
  process.exit(0);
});

process.on('SIGINT', () => {
  console.log('SIGINT received, shutting down gracefully');
  process.exit(0);
});

app.listen(port, () => {
  console.log(`Carrion Playground API listening on port ${port}`);
  console.log(`Health check: http://localhost:${port}/health`);
  
  // Check Podman availability on startup
  checkPodman().then(result => {
    if (result.available) {
      console.log(`✅ Podman available: ${result.version}`);
    } else {
      console.log(`⚠️  Podman check failed: ${result.error}`);
      console.log('Install Podman: https://podman.io/getting-started/installation');
    }
  });
});