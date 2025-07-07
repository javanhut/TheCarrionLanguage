# Built-in Functions

Carrion provides a rich set of built-in functions for common programming tasks. These functions are available globally without requiring imports.

## Type Conversion Functions

### `int(value)`
Converts a value to an integer.
```python
int("42")      // → 42
int(3.14)      // → 3
int(True)      // → 1
int(False)     // → 0
```

### `to_int(value)`
Alternative syntax for integer conversion.
```python
to_int("42")   // → 42
to_int(3.14)   // → 3
```

### `float(value)`
Converts a value to a floating-point number.
```python
float("3.14")  // → 3.14
float(42)      // → 42.0
float(True)    // → 1.0
```

### `str(value)`
Converts a value to a string representation.
```python
str(42)        // → "42"
str(3.14)      // → "3.14"
str(True)      // → "True"
str([1, 2, 3]) // → "[1, 2, 3]"
```

### `bool(value)`
Converts a value to a boolean.
```python
bool(1)        // → True
bool(0)        // → False
bool("")       // → False
bool("text")   // → True
bool([])       // → False
bool([1])      // → True
```

### `list(iterable)`
Converts an iterable to an array/list.
```python
list("hello")     // → ["h", "e", "l", "l", "o"]
list((1, 2, 3))   // → [1, 2, 3]
```

### `tuple(iterable)`
Converts an iterable to a tuple.
```python
tuple([1, 2, 3])  // → (1, 2, 3)
tuple("hello")    // → ("h", "e", "l", "l", "o")
```

## Utility Functions

### `len(object)`
Returns the length of strings, arrays, hashes, or tuples.
```python
len("hello")           // → 5
len([1, 2, 3, 4])      // → 4
len({"a": 1, "b": 2})  // → 2
len((1, 2, 3))         // → 3
```

### `type(object)`
Returns the type of an object as a string.
```python
type(42)         // → "INTEGER"
type(3.14)       // → "FLOAT"
type("hello")    // → "STRING"
type(True)       // → "BOOLEAN"
type([1, 2, 3])  // → "ARRAY"
type({"a": 1})   // → "HASH"
type((1, 2))     // → "TUPLE"
type(None)       // → "NONE"
```

### `print(*args)`
Prints values to the console with spaces between arguments.
```python
print("Hello")              // → Hello
print("Number:", 42)        // → Number: 42
print(1, 2, 3)             // → 1 2 3
print("Value is", x, "!")  // → Value is 10 !
```

### `input(prompt="")`
Reads user input from the console with an optional prompt.
```python
name = input("Enter your name: ")
age = int(input("Enter your age: "))
```

## Mathematical Functions

### `range(start, stop, step=1)`
Generates a sequence of numbers. Can be called with 1, 2, or 3 arguments.
```python
range(5)         // → [0, 1, 2, 3, 4]
range(2, 8)      // → [2, 3, 4, 5, 6, 7]
range(0, 10, 2)  // → [0, 2, 4, 6, 8]
range(10, 0, -1) // → [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
```

### `max(*args)`
Returns the maximum value from the arguments.
```python
max(1, 5, 3, 2)      // → 5
max([10, 20, 15])    // → 20
max("apple", "zoo")  // → "zoo" (lexicographic)
```

### `abs(value)`
Returns the absolute value of a number.
```python
abs(-42)    // → 42
abs(3.14)   // → 3.14
abs(-2.5)   // → 2.5
```

### `ord(char)`
Returns the ASCII/Unicode code point of a character.
```python
ord("A")    // → 65
ord("a")    // → 97
ord("0")    // → 48
```

### `chr(code)`
Returns the character corresponding to an ASCII/Unicode code point.
```python
chr(65)     // → "A"
chr(97)     // → "a"
chr(48)     // → "0"
```

## Collection Functions

### `enumerate(array)`
Returns an array of (index, value) tuples.
```python
items = ["a", "b", "c"]
for index, value in enumerate(items):
    print(index, value)
// Output:
// 0 a
// 1 b
// 2 c
```

### `pairs(hash, filter="")`
Returns key-value pairs from a hash as an iterable array. The second parameter filters the output.

**Parameters:**
- `hash`: The hash to extract pairs from
- `filter`: Optional filter string:
  - `""` (default): Returns `(key, value)` tuples
  - `"key"` or `"k"`: Returns only keys
  - `"value"` or `"v"`: Returns only values

**Returns:** Array of tuples, keys, or values that can be iterated over.

```python
data = {"name": "John", "age": 30, "city": "NYC"}

// Get all key-value pairs as tuples (default)
for key, value in pairs(data):
    print(f"{key}: {value}")
// Prints: name: John, age: 30, city: NYC

// Get only keys
for key in pairs(data, "key"):
    print(f"Key: {key}")
// Prints: Key: name, Key: age, Key: city

// Alternative key syntax
for key in pairs(data, "k"):
    print(key)

// Get only values  
for value in pairs(data, "value"):
    print(f"Value: {value}")
// Prints: Value: John, Value: 30, Value: NYC

// Alternative value syntax
for value in pairs(data, "v"):
    print(value)

// pairs() returns an array, so you can use it anywhere arrays work
all_pairs = pairs(data)
print(f"Total pairs: {len(all_pairs)}")

// Use with array methods
keys_only = pairs(data, "key")
if "name" in keys_only:
    print("Has name key")
```

### `is_sametype(obj1, obj2)`
Checks if two objects have the same type.
```python
is_sametype(42, 17)        // → True (both integers)
is_sametype(42, 3.14)      // → False (int vs float)
is_sametype("a", "hello")  // → True (both strings)
```

## File and OS Operations

Carrion provides comprehensive file and operating system operations through the File and OS grimoires, which use static methods for clean, consistent access to system resources.

### File Operations

The File grimoire provides static methods for all file operations:

```python
// Basic file operations
content = File.read("data.txt")                // Read entire file
File.write("output.txt", "Hello World")        // Write to file (overwrites)
File.append("log.txt", "New entry\n")          // Append to file
exists = File.exists("config.txt")             // Check if file exists

// File object operations for complex scenarios
autoclose File.open("data.txt", "r") as file:
    content = file.read_content()
    print(content)
```

**Available File Methods:**
- `File.read(path)` - Read entire file content as string
- `File.write(path, content)` - Write content to file (overwrites existing)
- `File.append(path, content)` - Append content to existing file
- `File.exists(path)` - Check if file exists (returns boolean)
- `File.open(path, mode="r")` - Open file for complex operations

### OS Operations

The OS grimoire provides static methods for operating system operations:

```python
// Directory operations
current_dir = OS.cwd()                         // Get current directory
OS.chdir("/home/user")                         // Change directory
files = OS.listdir(".")                        // List directory contents
OS.mkdir("new_folder")                         // Create directory
OS.remove("old_file.txt")                     // Remove file/directory

// Environment variables
home = OS.getenv("HOME")                       // Get environment variable
OS.setenv("DEBUG", "true")                     // Set environment variable
path = OS.expandEnv("$HOME/.config")           // Expand variables in string

// Process operations
OS.run("ls", ["-la"], False)                  // Run command, show output
output = OS.run("pwd", [], True)              // Run command, capture output
OS.sleep(2)                                   // Sleep for 2 seconds
```

**Available OS Methods:**
- `OS.cwd()` - Get current working directory
- `OS.chdir(path)` - Change to specified directory
- `OS.listdir(path=".")` - List directory contents
- `OS.mkdir(path, perm=0755)` - Create directory
- `OS.remove(path)` - Remove file or empty directory
- `OS.getenv(key)` - Get environment variable value
- `OS.setenv(key, value)` - Set environment variable (both arguments must be strings)
- `OS.expandEnv(string)` - Expand environment variables in string
- `OS.run(command, args=[], capture=False)` - Execute system command
- `OS.sleep(seconds)` - Sleep for specified time

### `open(path, mode="r")`
Opens a file and returns a File object for low-level file operations. **Note:** This function is primarily used internally by `File.open()` - for all file operations, use the File grimoire methods instead.

**Parameters:**
- `path` (string): The path to the file
- `mode` (string, optional): The file mode - "r" (read), "w" (write), "a" (append). Defaults to "r".

**Returns:** A File object with methods for file operations.

**Recommended Usage:**
```python
// Recommended: Use File grimoire static methods for simple operations
content = File.read("data.txt")
File.write("output.txt", "hello")
File.append("log.txt", "entry\n")

// Recommended: Use File.open() for complex file operations
autoclose File.open("data.txt", "r") as file:
    content = file.read_content()
    print(content)
```

**File Object Methods (when using File.open() or open()):**
- `read_content()` - Read entire file content from opened file
- `write_content(content)` - Write content to opened file  
- `close()` - Close the file (automatically called with autoclose)

**File Modes:**
- `"r"` - Read mode (file must exist)
- `"w"` - Write mode (creates new file or overwrites existing)
- `"a"` - Append mode (creates file if it doesn't exist)

### Error Functions
- `Error(name, message="")` - Create custom error object

## Module Function Registration

**Technical Note:** File and OS module functions are now properly registered with the evaluator during initialization. This ensures that all module functions (like `osSetEnv`, `fileRead`, etc.) are available as built-in functions and can be called from Carrion standard library grimoires.

The module registration system automatically:
- Loads File module functions (`fileOpen`, `fileRead`, `fileWrite`, etc.) into the global environment
- Loads OS module functions (`osGetEnv`, `osSetEnv`, `osRunCommand`, etc.) into the global environment
- Provides proper error handling with descriptive error messages
- Ensures type validation for function arguments (e.g., `OS.setenv` requires string arguments)

## Examples

### Type Checking and Conversion
```python
value = input("Enter a number: ")
if type(value) == "STRING":
    number = float(value)
    print("You entered:", number)
    print("Absolute value:", abs(number))
```

### Working with Collections
```python
numbers = [1, 5, 3, 9, 2]
print("Length:", len(numbers))
print("Maximum:", max(numbers))

// Convert to string representation
print("As string:", str(numbers))

// Enumerate through items
for index, value in enumerate(numbers):
    print(f"Index {index}: {value}")
```

### Character and String Processing
```python
text = "Hello"
for i in range(len(text)):
    char = text[i]  // Assuming string indexing is supported
    print(f"Character '{char}' has ASCII code {ord(char)}")

// Build string from ASCII codes
codes = [72, 101, 108, 108, 111]  // "Hello"
result = ""
for code in codes:
    result += chr(code)
print(result)  // → Hello
```

### File and System Operations Examples
```python
// Simple file operations using File grimoire
if File.exists("data.txt"):
    content = File.read("data.txt")
    processed = content.upper()
    File.write("output.txt", processed)
    File.append("log.txt", f"Processed file at {OS.cwd()}\n")

// Directory operations using OS grimoire
print("Current directory:", OS.cwd())
files = OS.listdir(".")
for filename in files:
    print(f"Found: {filename}")

// Environment operations
home = OS.getenv("HOME")
config_path = OS.expandEnv("$HOME/.config")
print(f"Config directory: {config_path}")

// Complex file operations with autoclose
autoclose File.open("data.txt", "r") as infile:
    content = infile.read_content()
    lines = content.split("\n")
    
    autoclose File.open("processed.txt", "w") as outfile:
        for line in lines:
            if line.strip():  // Skip empty lines
                outfile.write_content(line.upper() + "\n")
```