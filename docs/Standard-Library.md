# Carrion Standard Library (Munin)

The Munin standard library is the core collection of modules and functions for the Carrion programming language. Named after Odin's ravens Huginn and Munin (representing thought and memory), it provides computational capabilities and data management features.

**Current Version**: Carrion 0.1.7, Munin Standard Library 0.1.0

## Core Library Functions

### `help()`
Returns interactive help information about the language.
```python
help()  // Shows language help and available functions
```

### `version()`
Returns version information for Carrion and Munin.
```python
version()  // "Carrion 0.1.7, Munin Standard Library 0.1.0"
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

The File grimoire provides comprehensive file system operations with both static methods for simple operations and file objects for complex scenarios.

### Static File Operations (Recommended for Simple Tasks)

The File grimoire provides static methods that are the preferred way to perform basic file operations:

```python
// Read entire file content
content = File.read("data.txt")
print(f"File contains: {content}")

// Write content to file (overwrites existing content)
File.write("output.txt", "Hello from Carrion!")

// Append content to existing file
File.append("log.txt", "New log entry\n")

// Check if file exists
if File.exists("config.txt"):
    config = File.read("config.txt")
    print("Config loaded")
else:
    print("Config file not found")
```

**Available Static Methods:**
- `File.read(path)` - Read entire file content as string
- `File.write(path, content)` - Write content to file (overwrites)
- `File.append(path, content)` - Append content to file
- `File.exists(path)` - Check if file exists (returns boolean)
- `File.open(path, mode="r")` - Create File object for complex operations

### File Objects for Complex Operations

When you need more control over file operations, use File objects with the `autoclose` statement:

```python
// Reading files with File objects
autoclose File.open("data.txt", "r") as file:
    content = file.read_content()
    print(f"Read: {content}")
    // file.close() called automatically

// Writing files with File objects
autoclose File.open("output.txt", "w") as file:
    file.write_content("Line 1\n")
    file.write_content("Line 2\n")
    file.write_content("Line 3\n")

// Appending to files with File objects
autoclose File.open("log.txt", "a") as file:
    file.write_content("New log entry\n")
    file.write_content("Another entry\n")
```

**File Object Methods:**
- `read_content()` - Read entire file content (for files opened in "r" mode)
- `write_content(content)` - Write content to file (for files opened in "w" or "a" mode)
- `close()` - Close the file handle (automatically called with `autoclose`)

### File Modes
- `"r"` - **Read mode** (default): Opens file for reading. File must exist.
- `"w"` - **Write mode**: Creates new file or overwrites existing file completely.
- `"a"` - **Append mode**: Opens file for appending. Creates file if it doesn't exist.

### When to Use Which Approach

**Use Static Methods for:**
- Simple read/write operations
- One-time file operations
- Quick file existence checks
- Scripts and small programs

**Use File Objects for:**
- Multiple operations on the same file
- When you need precise control over file handles
- Complex file processing workflows
- When using with error handling blocks

### Complete Example

```python
// Static methods for simple operations
if not File.exists("users.txt"):
    File.write("users.txt", "admin\nguest\n")

// File objects for complex operations
autoclose File.open("users.txt", "a") as userfile:
    userfile.write_content("newuser\n")
    userfile.write_content("testuser\n")

// Read and process
users = File.read("users.txt")
for user in users.split("\n"):
    if user.strip():
        print(f"User: {user}")
```

### Error Handling with Files

```python
// Error handling with static methods
attempt:
    content = File.read("missing_file.txt")
    print(content)
ensnare:
    print("Could not read file")

// Error handling with file objects
attempt:
    autoclose File.open("sensitive_file.txt", "r") as file:
        data = file.read_content()
        print(f"Data: {data}")
ensnare:
    print("File access error")
```

## OS Module

The OS grimoire provides a comprehensive interface to operating system services through static methods, offering clean and consistent access to system operations.

### Directory Operations

```python
// Get current working directory
current_dir = OS.cwd()
print(f"Working in: {current_dir}")

// Change to different directory
OS.chdir("/home/user/projects")
print(f"Now in: {OS.cwd()}")

// List directory contents
files = OS.listdir(".")                   // List current directory
files = OS.listdir("/home/user")          // List specific directory
print(f"Found {len(files)} items")

for filename in files:
    print(f"File: {filename}")

// Create new directory
OS.mkdir("new_project")                   // Default permissions (0755)
OS.mkdir("secure_dir", 0700)              // Custom permissions

// Remove files and directories  
OS.remove("old_file.txt")                 // Remove file
OS.remove("empty_directory")              // Remove empty directory
```

### Environment Variables

```python
// Get environment variables
home = OS.getenv("HOME")
path = OS.getenv("PATH")
user = OS.getenv("USER")

print(f"Home directory: {home}")
print(f"Current user: {user}")

// Set environment variables (for current process)
// Note: All arguments must be strings
OS.setenv("MY_APP_CONFIG", "/etc/myapp")
OS.setenv("DEBUG_MODE", "true")
OS.setenv("PORT_NUMBER", "8080")  // Numbers must be converted to strings

// Expand environment variables in strings
config_path = OS.expandEnv("$HOME/.config/myapp")
log_path = OS.expandEnv("$MY_APP_CONFIG/logs")
print(f"Config will be at: {config_path}")

// Error handling for environment variables
attempt:
    OS.setenv("VALID_KEY", "valid_value")
    print("Environment variable set successfully")
ensnare:
    print("Failed to set environment variable")
```

### Process Management

```python
// Execute system commands
OS.run("ls", ["-la"], False)              // Run and show output directly
OS.run("mkdir", ["temp_folder"], False)   // Create directory via command

// Capture command output
output = OS.run("ls", ["-1"], True)       // Capture output as string
lines = output.split("\n")
print(f"Directory has {len(lines)} items")

// Simple commands with no arguments
OS.run("pwd", [], True)                   // Print working directory

// Sleep/delay execution
print("Starting process...")
OS.sleep(2)                               // Wait 2 seconds
print("Process complete")

OS.sleep(0.5)                             // Wait 500 milliseconds
```

**Available OS Methods:**
- `OS.cwd()` - Get current working directory
- `OS.chdir(path)` - Change to specified directory
- `OS.listdir(path=".")` - List directory contents (defaults to current dir)
- `OS.mkdir(path, perm=0755)` - Create directory with optional permissions
- `OS.remove(path)` - Remove file or empty directory
- `OS.getenv(key)` - Get environment variable value
- `OS.setenv(key, value)` - Set environment variable for current process (both arguments must be strings)
- `OS.expandEnv(string)` - Expand environment variables in string
- `OS.run(command, args=[], capture=False)` - Execute system command
- `OS.sleep(seconds)` - Sleep for specified time (supports decimals)

### Complete OS Example

```python
// System information gathering
print("=== System Information ===")
print(f"Current directory: {OS.cwd()}")
print(f"Home directory: {OS.getenv('HOME')}")
print(f"Current user: {OS.getenv('USER')}")

// Directory management
print("\n=== Directory Operations ===")
if not File.exists("temp_work"):
    OS.mkdir("temp_work")
    print("Created temp_work directory")

OS.chdir("temp_work")
print(f"Changed to: {OS.cwd()}")

// File operations in new directory
File.write("status.txt", "Working in temporary directory")
File.write("config.txt", "debug=true\nverbose=false")

// List and display contents
files = OS.listdir(".")
print(f"\nCreated {len(files)} files:")
for file in files:
    if File.exists(file):
        content = File.read(file)
        print(f"{file}: {content}")

// Cleanup
OS.chdir("..")
for file in ["temp_work/status.txt", "temp_work/config.txt"]:
    OS.remove(file)
OS.remove("temp_work")
print("Cleanup complete")
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
// Modern file operations with static methods (recommended)
content = File.read("data.txt")
processed = content.upper()
File.write("output.txt", processed)

// File operations with autoclose for complex scenarios
autoclose File.open("data.txt", "r") as input_file:
    content = input_file.read_content()
    processed = content.upper()
    
    autoclose File.open("output.txt", "w") as output_file:
        output_file.write_content(processed)

// System operations with static methods
print("Current directory:", OS.cwd())
files = OS.listdir(".")
for filename in files:
    print("Found file:", filename)

// Complete file processing workflow
if File.exists("config.txt"):
    config = File.read("config.txt")
    print("Configuration loaded")
else:
    File.write("config.txt", "default_setting=true\n")
    print("Created default configuration")

// Environment-aware file operations
backup_dir = OS.expandEnv("$HOME/backups")
if not File.exists(backup_dir):
    OS.mkdir(backup_dir)

backup_file = f"{backup_dir}/data_backup.txt"
File.write(backup_file, File.read("data.txt"))
print(f"Backup created at: {backup_file}")
```

## Module Organization

The standard library follows a modular design where each module is defined as a grimoire (class) with related functionality. All modules are automatically loaded when the Carrion interpreter starts, making functions immediately available without explicit imports.

For advanced usage and additional functions, see the individual module documentation and the built-in functions reference.