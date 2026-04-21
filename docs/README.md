<meta name="google-site-verification" content="7L-IkFjwJUUVamMg0bj1PwYOVcpowQyomYVhYM4e6lk" />
<meta name="description" content="Carrion Programming Language is a modern, dynamically typed, interpreted language inspired by Python and built in Go. Discover its fun crow theme, unique syntax, and powerful features for fast development and learning." />The Carrion Programming Language

Latest Version: 0.1.9 [![Release](https://img.shields.io/badge/version-0.1.9-blue.svg)](https://github.com/javanhut/TheCarrionLanguage/releases)

```bash


⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⡟⠋⢻⣷⣄⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣾⣿⣷⣿⣿⣿⣿⣿⣶⣾⣿⣿⠿⠿⠿⠶⠄⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠉⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⣿⣿⣿⣿⠟⠻⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣆⣤⠿⢶⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠑⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠸⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠙⠛⠋⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀
```

## Overview

Carrion Programming Language is a modern, dynamically typed, interpreted language designed for both beginners and experienced developers. Inspired by Python and built in Go, Carrion offers a unique combination of readability, flexibility, and performance. Its engaging crow-themed aesthetic and innovative syntax enhancements set it apart as a fun, educational, and efficient language for rapid prototyping and software development.

## Key Features

### Dynamic Typing & Interpretation

Experience the benefits of runtime type checking and an interpreter that supports a rapid development cycle.

### Python-Inspired Syntax with Creative Enhancements

- Familiar programming constructs with unique modifications
- Crow-themed keywords: `spell` (function), `grim` (class - short for grimoire/spellbook)
- Enhanced loop control: `skip` (continue) and `stop` (break)
- Powerful error handling: `attempt`/`ensnare`/`resolve`

### Built in Go for Performance

- Leverages Go's robustness and efficiency
- Tree-walking interpreter with planned JIT compilation
- Fast execution and minimal resource usage

### Comprehensive Standard Library (Munin)

- Named after Odin's raven, representing memory
- Built-in modules for arrays, strings, math, OS operations, and more
- Embedded directly into the language runtime

### Object-Oriented Programming Support

- Full OOP with classes (grimoires), inheritance, and encapsulation
- Abstract classes with `arcane grim`
- Private (`_`) and protected (`__`) methods
- Method overriding and polymorphism

## Installation

### Quick Install (Linux/macOS/Windows)

```bash
# Clone the repository
git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage

# Install Carrion Language, Sindri Testing Framework, Mimir Documentation Tool, and Bifrost package manager (auto-detects OS)
make install
```

**Note**: Installing Carrion automatically installs the **Sindri Testing Framework**, **Mimir Documentation Tool**, and **Bifrost Package Manager** (v1.0.0) as well. All tools are installed together and can be uninstalled with `make uninstall`.

### Manual Installation

1. Ensure Go 1.19+ is installed
2. Clone the repository
3. Run `make build` or `go build -o carrion src/main.go`
4. Add the binary to your PATH

### Docker Installation

```bash
# Build the Docker image
docker build -t carrion .

# Run Carrion in a container
docker run -it carrion
```

## Updating Carrion

Once installed, Carrion can update itself in place.

### Stable releases

```bash
carrion update            # Prompts before installing the latest tagged release
carrion update --check    # Report status without installing
carrion update -y         # Skip the confirmation prompt
```

`carrion update` fetches the latest release from GitHub, downloads the prebuilt asset for your OS/arch when available (`carrion_linux_amd64.tar.gz`, `carrion_darwin_amd64.tar.gz`, `carrion_windows_amd64.zip`), and falls back to a source build at the release tag if no asset is present. Stable updates always take precedence over experimental builds.

### Experimental (latest `main` commit)

```bash
carrion update --experimental          # Prompts before building from source
carrion update --experimental --check  # Report status without installing
```

Experimental updates track the `main` branch: Carrion fetches the latest commit SHA, clones the source, and rebuilds with the commit hash baked into the version string. Requires `git` and a Go 1.24+ toolchain in `PATH`.

### Version format

```bash
carrion --version
# Carrion Language version v0.1.10            — tagged release
# Carrion Language version v0.1.10-a1b2c3d    — experimental (main@a1b2c3d)
```

Stable releases show plain `v{major}.{minor}.{patch}`; experimental builds append `-{short-sha}` so you always know exactly what's running.

### Permissions

If Carrion is installed in a system directory like `/usr/local/bin`, `carrion update` will tell you to re-run with `sudo`. The binary is replaced atomically — no downtime for long-running processes that already have it open.

## Build & Test Targets

The Makefile exposes the following targets for contributors:

- `make build` — Build the Carrion binary for the host platform
- `make install` — Install Carrion, Sindri, Mimir, and Bifrost to the system
- `make uninstall` — Remove installed binaries
- `make build-linux` / `make build-linux-arm64` — Cross-compile a Linux amd64/arm64 binary tarball
- `make build-windows` — Cross-compile a Windows amd64 binary zip
- `make build-mac` (`build-mac-amd64`, `build-mac-arm64`) — Cross-compile macOS tarballs
- `make build-release` — Build every release artifact in one shot
- `make build-source` — Produce a source tarball
- `make test` — Run the full Go test suite (`go test ./src/...`)
- `make bench` — Run evaluator benchmarks with memory profiling
- `make sync-version` — Rewrite version references in docs from `src/version/version.go`
- `make version-check` — Report whether any docs are out of sync (exits 1 if so; CI-friendly)
- `make tidy` — Run `go mod tidy` across all modules
- `make bifrost-update` — Initialize and update the Bifrost git submodule

## Package Management

Carrion integrates with **Bifrost**, the official package manager, for seamless dependency management. Bifrost is automatically installed when you install Carrion.

```bash
# Initialize a new Carrion package
bifrost init

# Install packages for your project
bifrost install json-utils
bifrost install --global http-client

# Use installed packages in Carrion
import "json-utils/parser"
import "http-client/request"
```

For detailed Bifrost documentation, see the [Bifrost Repository](https://github.com/javanhut/bifrost) and [Bifrost Documentation](https://github.com/javanhut/bifrost/blob/main/README.md).

### Package Import Resolution

Carrion automatically resolves imports from multiple locations:

- Local files (current directory)
- Project packages (`./carrion_modules/`)
- User packages (`~/.carrion/packages/`)
- Global packages (`/usr/local/share/carrion/lib/`)
- Standard library (Munin)

See **[Modules Documentation](Modules.md)** for detailed import and package management information.

## Documentation

### Core Documentation

- **[Language Documentation](CARRION.md)** - Comprehensive guide to Carrion syntax and features
- **[Language Overview](Language-Overview.md)** - High-level introduction to Carrion's design principles
- **[Language Reference](Language-Reference.md)** - Complete language specification and reference

### Language Features

- **[Control Flow](Control-Flow.md)** - Loops, conditionals, and flow control structures
- **[Error Handling](Error-Handling.md)** - Exception handling with attempt/ensnare/resolve
- **[Enhanced Error System](ENHANCED_ERROR_SYSTEM.md)** - Detailed error messages with suggestions
- **[Operators](Operators.md)** - Arithmetic, logical, and comparison operators
- **[Modules](Modules.md)** - Import system and module organization
- **[Type System](Type-System.md)** - Type hints and static type checking
- **[Indentation](Indentation.md)** - Indentation rules and best practices

### Object-Oriented Programming

- **[Grimoires (Classes)](Grimoires.md)** - Object-oriented programming with grimoires

### Standard Library

- **[Standard Library (Munin)](Standard-Library.md)** - Built-in functions and modules
- **[Builtin Functions](Builtin-Functions.md)** - Core functions available in every program
- **[Data Structures](Data-Structures.md)** - Stack, Queue, Heap, and Binary Search Tree
- **[Concurrency](Concurrency.md)** - Goroutines with converge/diverge patterns
- **[HTTP Server](HTTP-Server-Enhancement.md)** - Building web applications and REST APIs
- **[Time Functions](TimeFunctions.md)** - Date, time, and duration operations

### Tools & Interactive Features

- **[Interactive Help](Interactive-Help-Demo.md)** - REPL and interactive development features
- **[Sindri Testing Framework](Sindri.md)** - Comprehensive testing and benchmarking tool
- **[Mimir Documentation Tool](Mimir.md)** - Interactive documentation and help system

### Additional Resources

- **[Examples](src/examples/)** - Sample programs demonstrating various language features
- **[Changelog](Changelog/README.md)** - Version history and updates

### Quick Start

```python
// Hello World in Carrion
print("Hello, World!")

// Define a function (spell)
spell greet(name):
    return f"Hello, {name}!"

// Create a class (grimoire)
grim Crow:
    init(name):
        self.name = name

    spell caw():
        print(f"{self.name} says: Caw!")

// Use the class
my_crow = Crow("Munin")
my_crow.caw()
```

## File Extension

Carrion source code files use the .crl extension, making it simple to identify and manage your projects.

## Future Enhancements

Carrion is an evolving language with exciting future updates planned:

- **List Comprehensions**: Simplify collection processing with Python-like concise syntax
- **JIT Compiler & Virtual Machine**: Enhance runtime performance with just-in-time compilation
- **Generic Functions & Abstract Data Types**: Improve code reuse and flexibility
- **Enhanced Standard Library**: Expand built-in functionalities with a richer set of tools
- **Improved Type System**: Optional static type checking for better code safety
- **Async/Await Support**: Modern concurrency patterns
- **Language Interoperability**: Integrate with other programming languages

## Build and Release Status

- Current Version: 0.1.9
- Standard Library (Munin): 0.1.0
- Status: Active Development

## About the Author

Carrion was created by Javan Hutchinson, a dedicated developer committed to exploring innovative programming paradigms and enhancing developer experiences.

## Contributing

We welcome contributions! Please:

1. Fork the repository
2. Create a feature branch
3. Submit a pull request

## Contact & Support

- **Email**: javanhut@carrionlang.com
- **Issues**: [GitHub Issues](https://github.com/javanhut/TheCarrionLanguage/issues)
- **Discussions**: Join our community discussions on GitHub

Your feedback and contributions help shape the future of Carrion!

> **Note**: While Carrion embraces a playful crow theme, it's a serious programming language built on Go's robust foundation, designed for real-world applications and educational purposes.

## License

Carrion is open-source software. See the [LICENSE](LICENSE) file for details.
