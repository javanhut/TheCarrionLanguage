# 🏠 Installation Guide - Setting Up Your Magical Environment

Welcome, future Carrion mage! This guide will help you set up your mystical coding environment and prepare you to cast your first spells with the Carrion programming language.

---

## 🎯 Prerequisites

Before you begin your magical journey, ensure you have these tools ready:

### Required
- **Go 1.19+** - Carrion is built with Go's robust foundation
- **Git** - To clone the magical repository
- **Terminal/Command Line** - Your spellcasting interface

### Optional but Recommended
- **Code Editor** with syntax highlighting (VS Code, Vim, Emacs, etc.)
- **Docker** - For containerized magic (optional)

---

## 🚀 Installation Methods

Choose your preferred installation method, young mage!

### 🌟 Method 1: Quick Install Script (Linux/macOS)

The fastest way to get Carrion up and running!

```bash
# Clone the mystical repository
git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage

# Run the magical installation script
./install/install.sh
```

**What this script does:**
- ✅ Checks for Go installation
- ✅ Builds the Carrion interpreter
- ✅ Adds Carrion to your PATH
- ✅ Verifies the installation
- ✅ Shows you how to get started

### 🔧 Method 2: Manual Installation

For those who prefer hands-on control of their magic!

#### Step 1: Clone the Repository
```bash
git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage
```

#### Step 2: Build Carrion
```bash
# Using the Makefile (recommended)
make build

# Or build directly with Go
cd src && go build -o ../carrion main.go
```

#### Step 3: Add to PATH (Linux/macOS)
```bash
# Add to your shell profile (~/.bashrc, ~/.zshrc, etc.)
echo 'export PATH="$PATH:/path/to/TheCarrionLanguage"' >> ~/.bashrc
source ~/.bashrc

# Or create a symlink
sudo ln -s /path/to/TheCarrionLanguage/carrion /usr/local/bin/carrion
```

#### Step 4: Verify Installation
```bash
carrion --version
```

### 🐳 Method 3: Docker Installation

For isolated magical environments!

#### Build Docker Image
```bash
# Clone and build
git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage

# Build the magical container
docker build -t carrion .
```

#### Run Carrion in Docker
```bash
# Interactive REPL
docker run -it carrion

# Run a specific script
docker run -v $(pwd):/workspace carrion /workspace/your_script.crl

# Mount your project directory
docker run -it -v $(pwd):/workspace -w /workspace carrion
```

---

## 🧪 Testing Your Installation

Let's make sure everything is working with some quick tests!

### 1. Check Version
```bash
carrion --version
```
*Expected output: Something like "Carrion 0.1.6, Munin Standard Library 0.1.0"*

### 2. Test REPL
```bash
carrion
```
You should see something like:
```
🐦‍⬛ Welcome to Carrion 0.1.6-alpha
   The Mystical Programming Language
   
   Type 'mimir' for interactive help
   Type 'quit' or 'exit' to leave
   
>>> 
```

### 3. Cast Your First Spell
In the REPL, try:
```python
>>> print("Hello, Magical World! 🪄")
Hello, Magical World! 🪄

>>> spell greet(name):
...     return f"Greetings, {name}!"
... 
>>> greet("Aspiring Mage")
Greetings, Aspiring Mage!
```

### 4. Test Script Execution
Create a test file:
```bash
echo 'print("Carrion magic works! ✨")' > test_magic.crl
carrion test_magic.crl
```

---

## 🎮 Post-Installation Setup

### Enable REPL Enhancements

#### Command History
Your REPL comes with built-in command history! Use:
- **↑/↓ arrows** - Navigate command history
- **Tab** - Auto-complete keywords and functions

#### Interactive Help System
```python
>>> mimir
# This launches the comprehensive help system!
```

### IDE Integration

#### VS Code (Recommended)
1. Install a generic syntax highlighter
2. Set file association for `.crl` files
3. Configure Go tools for development

```json
{
    "files.associations": {
        "*.crl": "python"
    }
}
```

#### Vim
Add to your `.vimrc`:
```vim
au BufRead,BufNewFile *.crl set filetype=python
```

---

## 🔧 Troubleshooting

### Common Issues and Solutions

#### ❌ "Go not found" Error
```bash
# Install Go from https://golang.org/dl/
# Or using package managers:

# Ubuntu/Debian
sudo apt install golang-go

# macOS (Homebrew)
brew install go

# CentOS/RHEL
sudo yum install golang
```

#### ❌ "Permission denied" Error
```bash
# Make the binary executable
chmod +x carrion

# Or run with proper permissions
sudo ./install/install.sh
```

#### ❌ "Command not found: carrion"
```bash
# Check if PATH is set correctly
echo $PATH | grep carrion

# Reload your shell configuration
source ~/.bashrc  # or ~/.zshrc

# Or run with full path
/path/to/TheCarrionLanguage/carrion
```

#### ❌ Build Errors
```bash
# Ensure you have the latest Go version
go version

# Clean and rebuild
make clean
make build

# Or manually
cd src && go mod tidy && go build -o ../carrion main.go
```

### Getting Help

If you encounter issues:

1. **📖 Check the Documentation** - Many answers are in this wiki
2. **🐛 Search Issues** - [GitHub Issues](https://github.com/javanhut/TheCarrionLanguage/issues)
3. **💬 Ask the Community** - [GitHub Discussions](https://github.com/javanhut/TheCarrionLanguage/discussions)
4. **📧 Contact Support** - javanhut@carrionlang.com

---

## 🎊 What's Next?

Congratulations! You've successfully set up your magical Carrion environment. Here's what to do next:

1. **🎪 Explore the REPL**: Start with `carrion` and try `mimir` for guided help
2. **👋 Hello World**: Move on to [Hello World & Basics](Hello-World.md)
3. **⚡ Quick Tutorial**: Try our [Quick Start Tutorial](Quick-Start.md)
4. **📚 Learn Fundamentals**: Dive into [Language Fundamentals](../Language-Fundamentals/Syntax-and-Terminology.md)

---

## 🌟 Development Setup (Optional)

If you want to contribute to Carrion development:

### Clone for Development
```bash
git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage

# Create a development branch
git checkout -b your-feature-branch
```

### Development Tools
```bash
# Run tests
cd src && go test ./...

# Format code
go fmt ./...

# Static analysis
go vet ./...
```

### Build Variants
```bash
# Debug build
go build -tags debug -o carrion-debug src/main.go

# Release build
go build -ldflags "-s -w" -o carrion src/main.go
```

---

*May your installation be smooth and your first spells successful! Welcome to the magical world of Carrion! 🪄✨*

> "Every great mage started with a single installation." - *Carrion Proverb*