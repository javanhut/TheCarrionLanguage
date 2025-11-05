<meta name="google-site-verification" content="7L-IkFjwJUUVamMg0bj1PwYOVcpowQyomYVhYM4e6lk" />
<meta name="description" content="Carrion Programming Language is a modern, dynamically typed, interpreted language inspired by Python and built in Go. Discover its fun crow theme, unique syntax, and powerful features for fast development and learning." />The Carrion Programming Language

Latest Version: 0.1.8 [![Release](https://img.shields.io/badge/version-0.1.8-blue.svg)](https://github.com/javanhut/TheCarrionLanguage/releases)

```bash


⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⡟⠋⢻⣷⣄⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣾⣿⣷⣿⣿⣿⣿⣿⣶⣾⣿⣿⠿⠿⠿⠶⠄⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠉⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀
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

## What's New in v0.1.8

### Comprehensive Tooling Ecosystem

Carrion v0.1.8 introduces a complete development toolchain that transforms how you write, test, and document code:

#### Sindri Testing Framework
- Automatic test discovery using "appraise" naming convention
- Colored terminal output (green for pass, red for fail)
- Flexible assertions with `check()` function
- Support for both boolean and value comparison assertions
- Multiple output modes (summary and detailed)
- Directory and file-based testing

#### Mimir Documentation System
- Interactive documentation browser with menu navigation
- Command-line function lookup (`mimir scry <function>`)
- Comprehensive coverage of built-in functions and standard library
- Category-based browsing and search functionality
- Seamless REPL integration

#### Bifrost Package Manager
- Git submodule integration for package management
- Install, manage, and distribute Carrion packages
- Foundation for growing Carrion ecosystem

### API Improvements

#### Static Method Support
Grimoires now support static method calls for cleaner, more organized APIs:

```python
# New unified grimoire API
content = File.read("data.txt")
File.write("output.txt", "hello")
current_dir = OS.cwd()
files = OS.listdir(".")
```

#### File Operations
- `File.read(path)` - Read file content
- `File.write(path, content)` - Write to file
- `File.append(path, content)` - Append to file
- `File.exists(path)` - Check if file exists
- `File.open(path, mode)` - Open file for complex operations

#### OS Operations
- `OS.cwd()` - Get current directory
- `OS.chdir(path)` - Change directory
- `OS.listdir(path)` - List directory contents
- `OS.getenv(key)` - Get environment variable
- `OS.setenv(key, value)` - Set environment variable
- `OS.remove(path)` - Remove file/directory
- `OS.mkdir(path, perm)` - Create directory
- `OS.run(cmd, args, capture)` - Execute system commands
- `OS.sleep(seconds)` - Sleep for specified time

### Critical Bug Fixes

#### Multi-Level Inheritance
Fixed infinite recursion bug in 3+ level inheritance chains. Now supports deep inheritance hierarchies:

```python
grim Level1:
    init(x):
        self.l1 = x

grim Level2(Level1):
    init(x):
        super.init(x)  # Now correctly calls Level1.init()
        self.l2 = x

grim Level3(Level2):
    init(x):
        super.init(x)  # Correctly calls Level2.init()
        self.l3 = x
```

#### String Concatenation
Fixed critical bug where string concatenation returned incorrect types, affecting socket operations and other string-heavy code:

```python
# Now works correctly
headers = "HTTP/1.1 200 OK\r\n\r\n"
html = """<!DOCTYPE html><html></html>"""
response = headers + html  # Returns proper String instance
print(response.length())   # String methods work correctly
```

#### Variable Resolution
Fixed environment variable precedence - user-defined variables now correctly override builtin function names.

### Enhanced Import System

#### Selective Imports
Import specific grimoires or spells from modules using dot notation:

```python
# Import specific grimoire from a module
import "utils.StringHelper"
helper = StringHelper()

# Import specific spell/function from a module
import "math.advanced_calc"
result = advanced_calc(10, 20)

# Traditional full module import still works
import "utils"
```

This feature allows for cleaner namespace management and more explicit dependencies in your code.

### REPL Improvements

#### Cleaner Output Display
The REPL now provides a more streamlined experience:

```python
# Assignment statements don't print output
x = 10                    # (no output)
y = "hello"               # (no output)

# Function/grimoire definitions don't print
spell test():             # (no output)
    return 42

grim MyClass:             # (no output)
    init():
        self.value = 0

# Only expression results are shown
x + 5                     # 15
len("hello")              # 5
test()                    # 42
```

This reduces clutter and makes the REPL feel more intuitive for interactive development.

### Tuple Handling Refinements

Improved tuple behavior in function calls and assignments for more predictable operations:

```python
# Tuples are now handled more explicitly
coords = (10, 20, 30)
spell process_coords(point):
    # point is the tuple itself, not auto-unpacked
    return point

result = process_coords(coords)  # Returns the tuple as-is
```

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
- **[Operators](Operators.md)** - Arithmetic, logical, and comparison operators
- **[Modules](Modules.md)** - Import system and module organization

### Object-Oriented Programming

- **[Grimoires (Classes)](Grimoires.md)** - Object-oriented programming with grimoires

### Standard Library

- **[Standard Library (Munin)](Standard-Library.md)** - Built-in functions and modules
- **[Builtin Functions](Builtin-Functions.md)** - Core functions available in every program

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

- Current Version: 0.1.8
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
