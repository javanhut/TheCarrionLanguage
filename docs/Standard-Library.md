# Carrion Standard Library (Munin)

The Munin standard library is the core collection of modules and functions for the Carrion programming language. Named after Odin's ravens Huginn and Munin (representing thought and memory), it provides computational capabilities and data management features.

**Current Version**: Carrion 0.1.6, Munin Standard Library 0.1.0

## Core Library Functions

### `help()`
Returns interactive help information about the language.
```python
help()  // Shows language help and available functions
```

### `version()`
Returns version information for Carrion and Munin.
```python
version()  // "Carrion 0.1.6, Munin Standard Library 0.1.0"
```

### `modules()`
Lists all available modules with descriptions.
```python
modules()  // Shows available standard library modules
```

## Array Module

The Array grimoire provides comprehensive array manipulation capabilities.

### Constructor
```python
Array(elements)  // Creates an array wrapper with enhanced methods
```

### Core Methods
```python
arr = Array([1, 2, 3, 4, 5])

// Basic operations
arr.length()        // → 5
arr.append(6)       // Adds 6 to the end
arr.get(0)          // → 1 (supports negative indexing)
arr.set(0, 10)      // Sets first element to 10
arr.is_empty()      // → False
arr.contains(3)     // → True
arr.clear()         // Removes all elements

// Search and access
arr.index_of(3)     // → 2 (index of first occurrence)
arr.remove(3)       // Removes first occurrence of 3
arr.first()         // Gets first element
arr.last()          // Gets last element

// Advanced operations
arr.slice(1, 3)     // Creates new array from range [1, 3)
arr.reverse()       // Creates new array with reversed elements
arr.sort()          // Creates new sorted array using bubble sort
arr.to_string()     // Returns string representation
```

## String Module

The String grimoire provides comprehensive text manipulation.

### Constructor
```python
String(value)  // Creates a string wrapper
```

### Methods
```python
s = String("Hello World")

// Basic properties
s.length()          // → 11
s.to_string()       // → "Hello World"

// Case conversion
s.lower()           // → "hello world"
s.upper()           // → "HELLO WORLD"

// Search and access
s.find("World")     // → 6 (index of first occurrence)
s.contains("Hello") // → True
s.char_at(0)        // → "H" (supports negative indexing)
s.char_at(-1)       // → "d"

// Transformation
s.reverse()         // → "dlroW olleH"
```

## Integer Module

The Integer grimoire provides enhanced integer functionality.

### Constructor
```python
Integer(value = 0)  // Creates an integer wrapper
```

### Methods
```python
i = Integer(42)

// Number base conversions
i.to_bin()          // → "0b101010"
i.to_oct()          // → "0o52"  
i.to_hex()          // → "0x2a"

// Mathematical operations
i.abs()             // → 42
i.pow(2)            // → 1764
i.gcd(18)           // → 6 (greatest common divisor)
i.lcm(18)           // → 126 (least common multiple)

// Utility methods
i.is_even()         // → True
i.is_odd()          // → False
i.is_prime()        // → False
i.to_string()       // → "42"
i.to_float()        // → 42.0
```

### Prime Number Checking
```python
// Automatic primitive wrapping allows:
print(17.is_prime())  // → True
print(42.is_prime())  // → False
```

## Float Module

The Float grimoire provides floating-point operations.

### Constructor
```python
Float(value = 0.0)  // Creates a float wrapper
```

### Methods
```python
f = Float(3.14159)

// Rounding and precision
f.round(2)          // → 3.14
f.floor()           // → 3
f.ceil()            // → 4

// Mathematical operations
f.abs()             // Returns absolute value
f.sqrt()            // Square root using Newton's method
f.pow(2)            // Raises to power
f.sin()             // Sine using Taylor series
f.cos()             // Cosine using Taylor series

// Type checking
f.is_integer()      // → False
f.is_positive()     // → True
f.is_negative()     // → False
f.is_zero()         // → False

// Conversions
f.to_int()          // → 3
f.to_string()       // → "3.14159"
```

## Boolean Module

The Boolean grimoire provides logical operations.

### Constructor
```python
Boolean(value = False)  // Creates boolean with automatic conversion
```

### Methods
```python
b = Boolean(True)

// Conversions
b.to_int()          // → 1 (False → 0)
b.to_string()       // → "True"

// Logical operations
b.negate()          // → False
b.and_with(False)   // → False
b.or_with(False)    // → True
b.xor_with(True)    // → False
b.implies(False)    // → False (logical implication)

// Testing
b.is_true()         // → True
b.is_false()        // → False
```

## File Module

The File grimoire provides file system operations and automatic resource management.

### File Object Creation
```python
// Create File objects using the open() builtin function
file = open("filename.txt", "r")      // Read mode
file = open("filename.txt", "w")      // Write mode  
file = open("filename.txt", "a")      // Append mode
```

### File Object Methods
```python
// Reading
content = file.read()                 // Read entire file content

// Writing
file.write("Hello World")             // Write content to file

// Closing
file.close()                          // Close file handle
```

### Automatic Resource Management
The File grimoire works seamlessly with the `autoclose` statement to ensure proper resource cleanup:

```python
// Automatic file closing - recommended approach
autoclose open("data.txt", "r") as file:
    content = file.read()
    print(content)
// file.close() is called automatically

// Writing with autoclose
autoclose open("output.txt", "w") as file:
    file.write("Hello, World!")
    file.write("Second line")

// Appending with autoclose  
autoclose open("log.txt", "a") as file:
    file.write("New log entry\n")
```

### Static File Operations
For simple file operations, static methods are still available:

```python
// Static methods (legacy approach)
f = File()
content = f.read("input.txt")
f.write("output.txt", "Hello World")  // Overwrites file
f.append("log.txt", "New entry\n")    // Appends to file

// File existence check
if f.exists("config.txt"):
    config = f.read("config.txt")
```

### File Modes
- `"r"` - Read mode (default): Opens file for reading
- `"w"` - Write mode: Creates new file or overwrites existing file  
- `"a"` - Append mode: Opens file for appending, creates if doesn't exist

### Error Handling
File operations handle common errors gracefully:
```python
autoclose open("nonexistent.txt", "r") as file:
    content = file.read()  // May raise file not found error

// Use with error handling
attempt:
    autoclose open("data.txt", "r") as file:
        content = file.read()
        print(content)
ensnare:
    print("Error reading file")
```

## OS Module

The OS grimoire provides operating system interface.

### Constructor and Methods
```python
os = OS()

// Directory operations
current_dir = os.cwd()              // Get current directory
os.chdir("/path/to/directory")      // Change directory
files = os.listdir(".")             // List directory contents
os.mkdir("new_folder", 0755)        // Create directory
os.remove("file_or_folder")         // Remove file/directory

// Environment variables
home = os.getenv("HOME")            // Get environment variable
os.setenv("MY_VAR", "value")        // Set environment variable
expanded = os.expandEnv("$HOME/docs") // Expand variables

// Process management
os.run("ls", ["-la"], False)        // Execute command
os.sleep(2)                         // Sleep for 2 seconds
```

## Automatic Primitive Wrapping

Carrion automatically wraps primitive types with their corresponding grimoire objects, allowing method calls directly on basic values:

```python
// These work automatically due to primitive wrapping:
print(10.to_bin())        // "0b1010"
print("Hello".upper())    // "HELLO"
print(3.14.round(1))      // 3.1
print(True.to_int())      // 1
print(17.is_prime())      // True

// Arrays also get automatic enhancement
numbers = [3, 1, 4, 1, 5]
print(numbers.sort())     // Creates sorted copy
print(numbers.contains(4)) // True
```

## Usage Examples

### Array Processing
```python
numbers = [10, 5, 8, 3, 7]
arr = Array(numbers)

// Find and manipulate data
max_index = arr.index_of(max(numbers))
arr.append(12)
sorted_copy = arr.sort()

print("Original:", arr.to_string())
print("Sorted:", sorted_copy.to_string())
```

### String Processing
```python
text = String("Hello, World!")
print("Length:", text.length())
print("Uppercase:", text.upper())
print("Contains 'World':", text.contains("World"))
print("Reversed:", text.reverse())
```

### Mathematical Operations
```python
// Integer operations
num = 42
print("Binary:", num.to_bin())
print("Is prime:", num.is_prime())
print("GCD with 18:", num.gcd(18))

// Float operations  
pi = 3.14159
print("Rounded:", pi.round(2))
print("Square root:", pi.sqrt())
```

### File and System Operations
```python
// Modern file operations with autoclose
autoclose open("data.txt", "r") as input_file:
    content = input_file.read()
    processed = content.upper()
    
    autoclose open("output.txt", "w") as output_file:
        output_file.write(processed)

// System operations
os = OS()
print("Current directory:", os.cwd())
files = os.listdir(".")
for filename in files:
    print("Found file:", filename)

// Legacy file operations (still supported)
file = File()
if file.exists("data.txt"):
    content = file.read("data.txt")
    processed = content.upper()
    file.write("output.txt", processed)
```

## Module Organization

The standard library follows a modular design where each module is defined as a grimoire (class) with related functionality. All modules are automatically loaded when the Carrion interpreter starts, making functions immediately available without explicit imports.

For advanced usage and additional functions, see the individual module documentation and the built-in functions reference.