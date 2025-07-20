# Mimir Documentation Tool

Mimir is Carrion's interactive documentation and help system, named after Odin's wise raven. It provides comprehensive documentation access both interactively and through command-line queries.

## Overview

Mimir serves as the primary documentation interface for the Carrion Programming Language, offering:

- Interactive documentation browsing
- Function-specific help lookup
- Comprehensive function and module listings
- Category-based browsing
- Search functionality

## Installation

Mimir is automatically installed when you install Carrion using:

```bash
make install
```

## Usage

### Interactive Mode

Start Mimir in interactive mode for browsing documentation:

```bash
mimir
```

This launches an interactive session where you can:

1. **Browse Built-in Functions** - Core language functions (print, len, type, etc.)
2. **Explore Standard Library** - Munin modules (Array, String, File, OS, etc.)
3. **Learn Language Features** - Syntax, control flow, OOP, error handling
4. **View Examples & Demos** - Working code examples and tutorials
5. **Search Functions** - Find specific functions by name or purpose
6. **Get Tips & Tricks** - REPL shortcuts and advanced features

### Command-Line Usage

#### Get Help for Specific Functions

```bash
# Get help for a specific function
mimir scry print
mimir scry Array
mimir scry os

# Alternative short form
mimir s print
```

#### List All Functions

```bash
# Show all available functions and modules
mimir list

# Alternative short form
mimir l
```

#### Browse by Categories

```bash
# Show function categories
mimir categories

# Alternative short forms
mimir cat
mimir c
```

#### Help and Usage

```bash
# Show usage information
mimir help
mimir --help
mimir -h
```

## Examples

### Interactive Session Example

```bash
$ mimir
    MIMIR - The All-Seeing Helper
   ═══════════════════════════════════════
    Interactive Carrion Documentation
   ═══════════════════════════════════════
   "Knowledge is the greatest treasure"

What knowledge do you seek?

  1.  Built-in Functions    - Core language functions (print, len, type, etc.)
  2.  Standard Library      - Munin modules (Array, String, File, OS, etc.)
  3.  Language Features     - Syntax, control flow, OOP, error handling
  4.  Examples & Demos      - Working code examples and tutorials
  5.  Search Functions      - Find specific functions by name or purpose
  6.  Tips & Tricks         - REPL shortcuts and advanced features

Commands: Type number/name (e.g., '1' or 'builtins'), 'h' for menu, 'q' to quit
Quick search: Type any function name directly (e.g., 'print', 'Array')

mimir> 1
```

### Function Lookup Example

```bash
$ mimir scry print
═══════════════════════════════════════════════════════════════════
SCRYING: PRINT
═══════════════════════════════════════════════════════════════════

BUILT-IN FUNCTION: PRINT
═════════════════════════════
print(*args) - Print values to console
   Example: print("Hello", 42, True)
```

### List Functions Example

```bash
$ mimir list
ALL AVAILABLE FUNCTIONS AND MODULES
═══════════════════════════════════════════════════════════════════

BUILT-IN FUNCTIONS:
Type Conversion: int, float, str, bool, list, tuple
Utility:        print, len, type, input, range, max, abs
System:         help, version, modules, open, parseHash
Mathematical:   max, abs, ord, chr
Collections:    enumerate, pairs, is_sametype

STANDARD LIBRARY MODULES:
Array:    Enhanced array operations and methods
String:   String manipulation and processing
Integer:  Integer utilities and conversions
Float:    Floating-point operations
Boolean:  Boolean logic operations
File:     File I/O operations
OS:       Operating system interface
HTTP:     Web requests and JSON processing
Time:     Date and time utilities

Use 'mimir scry <function>' for detailed help on any item above
```

## Integration with Carrion REPL

The Carrion REPL now directs users to Mimir for comprehensive documentation:

```
>>> help()
For comprehensive interactive documentation, run: mimir
For specific function help, use: mimir scry <function>
```

## Command Reference

| Command | Short Form | Description |
|---------|------------|-------------|
| `mimir` | - | Start interactive mode |
| `mimir interactive` | `mimir i` | Start interactive mode explicitly |
| `mimir scry <function>` | `mimir s <function>` | Get help for specific function |
| `mimir list` | `mimir l` | List all functions and modules |
| `mimir categories` | `mimir cat`, `mimir c` | Show function categories |
| `mimir help` | `mimir -h`, `mimir --help` | Show usage information |

## Documentation Coverage

Mimir provides documentation for:

### Built-in Functions
- Type conversion functions (int, float, str, bool, list, tuple)
- Utility functions (print, len, type, input, range, max, abs)
- System functions (help, version, modules, open, parseHash)
- Mathematical functions (max, abs, ord, chr)
- Collection functions (enumerate, pairs, is_sametype)

### Standard Library Modules
- **Array**: Enhanced array operations and methods
- **String**: String manipulation and processing
- **Integer**: Integer utilities and conversions
- **Float**: Floating-point operations
- **Boolean**: Boolean logic operations
- **File**: File I/O operations
- **OS**: Operating system interface
- **HTTP**: Web requests and JSON processing
- **Time**: Date and time utilities

### Language Features
- Syntax and grammar
- Control flow structures
- Object-oriented programming
- Error handling
- Module system

## Tips for Effective Use

1. **Start with Interactive Mode**: Use `mimir` to explore available documentation categories
2. **Use Direct Lookup**: When you know what you're looking for, use `mimir scry <function>`
3. **Browse by Category**: Use `mimir categories` to find functions by purpose
4. **Quick Reference**: Use `mimir list` for a complete overview
5. **Integration**: Mimir complements the REPL - use both for efficient development

## Future Enhancements

Planned improvements for Mimir include:

- Enhanced search capabilities with fuzzy matching
- Code examples for all functions
- Interactive tutorials
- Offline documentation caching
- Custom documentation plugins

## See Also

- [Interactive Help Demo](Interactive-Help-Demo.md) - REPL features and interactive development
- [Sindri Testing Framework](Sindri.md) - Testing and benchmarking tool
- [Standard Library](Standard-Library.md) - Complete standard library reference
- [Builtin Functions](Builtin-Functions.md) - Core function documentation