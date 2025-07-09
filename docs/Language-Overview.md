# Carrion Programming Language - Overview

## Introduction

Carrion is a dynamically typed, interpreted programming language with a Norse mythology and crow-themed syntax. It combines familiar programming concepts with magical terminology, creating an engaging and unique development experience.

## Key Features

- **Dynamic Typing**: Variables don't require explicit type declarations
- **Magical Syntax**: Classes are "grimoires", methods are "spells"
- **Norse Mythology Theme**: Standard library named "Munin" (Odin's raven)
- **Python-Inspired**: Familiar syntax with creative modifications
- **Object-Oriented**: Full support for classes, inheritance, and encapsulation
- **Interactive REPL**: Built-in read-eval-print loop
- **Error Handling**: Comprehensive try-catch system with custom syntax
- **Module System**: Import and organize code across files
- **Package Management**: Integrated Bifrost package manager for dependency management
- **Universal Iteration**: Strings, arrays, and maps are all iterable
- **Membership Testing**: Enhanced `in` operator works with all collection types
- **Multiplication Operators**: String and array repetition with `*` operator

## Basic Syntax

### Hello World
```python
print("Hello, World!")
```

### Variable Assignment
```python
# Basic assignment
name = "Carrion"
age = 42
is_magical = True

# Tuple unpacking
x, y = (10, 20)
```

### Comments
```python
// Single-line comment
/* 
   Multi-line
   comment 
*/
```

## Data Types

### Primitive Types
- **Integer**: 64-bit signed integers (`42`, `-17`)
- **Float**: 64-bit floating-point numbers (`3.14`, `-2.718`)
- **String**: UTF-8 text (`"hello"`, `'world'`)
- **Boolean**: True/False values (`True`, `False`)
- **None**: Null/empty value (`None`)

### Collection Types
- **Array**: Ordered, mutable lists (`[1, 2, 3]`)
- **Map**: Key-value dictionaries (`{"key": "value"}`) with support for string, integer, float, and boolean keys
- **Tuple**: Immutable ordered collections (`(1, 2, 3)`)

## Getting Started

### Installation
```bash
# Clone the repository
git clone https://github.com/javanhut/TheCarrionLanguage
cd TheCarrionLanguage

# Build the interpreter
cd src && go build

# Run a Carrion file
./src filename.crl

# Start interactive REPL
./src
```

### File Extension
Carrion source files use the `.crl` extension.

### Basic Program Structure
```python
// Import statements
import "math"

// Variable declarations
name = "World"

// Function definitions
spell greet(name):
    return f"Hello, {name}!"

// Main program logic
message = greet(name)
print(message)
```

## Philosophy

Carrion embraces the mystical and magical, transforming mundane programming concepts into enchanting terminology:

- **Classes** become **grimoires** (spellbooks)
- **Methods** become **spells**
- **Standard library** is **Munin** (Odin's wise raven)
- **Error handling** uses **attempt/ensnare/resolve** instead of try/catch/finally

This creative approach makes programming feel more like crafting magic while maintaining familiar, readable syntax patterns.

## Package Management

Carrion includes **Bifrost**, the official package manager, which is automatically installed when you install Carrion. Bifrost provides seamless dependency management and module distribution.

### Quick Start with Bifrost

```bash
# Initialize a new Carrion project
bifrost init

# Install packages
bifrost install json-utils
bifrost install --global http-client

# Use packages in your code
import "json-utils/parser"
import "http-client/request"
```

For complete Bifrost documentation, visit the [Bifrost Repository](https://github.com/javanhut/bifrost).

## Next Steps

Explore the detailed documentation:
- [Built-in Functions](Builtin-Functions.md)
- [Operators and Expressions](Operators.md)
- [Control Flow](Control-Flow.md)
- [Object-Oriented Programming](Grimoires.md)
- [Standard Library](Standard-Library.md)
- [Error Handling](Error-Handling.md)
- [Module System](Modules.md)