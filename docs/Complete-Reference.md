# Carrion Programming Language - Complete Reference

**Version**: Carrion 0.1.7, Munin Standard Library 0.1.0

This comprehensive guide covers all aspects of the Carrion programming language, including syntax, built-in functions, standard library modules, and practical examples.

## Table of Contents

1. [Language Overview](#language-overview)
2. [Basic Syntax](#basic-syntax)
3. [Data Types](#data-types)
4. [Variables and Assignments](#variables-and-assignments)
5. [Operators](#operators)
6. [Control Flow](#control-flow)
7. [Functions](#functions)
8. [Grimoires (Classes)](#grimoires-classes)
9. [Error Handling](#error-handling)
10. [Built-in Functions](#built-in-functions)
11. [Standard Library (Munin)](#standard-library-munin)
12. [File Operations](#file-operations)
13. [OS Operations](#os-operations)
14. [Module System](#module-system)
15. [Advanced Features](#advanced-features)
16. [Complete Examples](#complete-examples)

---

## Language Overview

Carrion is a dynamic, interpreted programming language with Norse mythology-inspired naming conventions. It features:

- **Indentation-based syntax** (like Python)
- **Dynamic typing** with automatic type conversion
- **Object-oriented programming** with grimoires (classes) and spells (methods)
- **Built-in error handling** with `attempt/ensnare/resolve` blocks
- **Comprehensive standard library** (Munin) with automatic primitive wrapping
- **Module system** with smart import resolution
- **File and OS operations** through unified grimoire interfaces

### Norse-Inspired Terminology

| Standard Term | Carrion Term | Description |
|---------------|--------------|-------------|
| Class | Grimoire | Object-oriented class definition |
| Method | Spell | Function within a grimoire |
| Constructor | init() | Object initialization method |
| Standard Library | Munin | Core library modules (named after Odin's raven) |

---

## Basic Syntax

### Comments
```python
// Single-line comment
/* Multi-line
   comment */
```

### Print Statements
```python
print("Hello, World!")
print("Number:", 42)
print(1, 2, 3)  // Multiple values separated by spaces
```

### Variables
```python
name = "Alice"
age = 30
height = 5.8
is_student = True
nothing = None
```

### Code Blocks
Carrion uses indentation to define code blocks (like Python):

```python
if age >= 18:
    print("Adult")
    print("Can vote")
else:
    print("Minor")
    print("Cannot vote")
```

---

## Data Types

### Primitive Types

#### Integer
```python
x = 42
negative = -17
binary = 0b1010    // Binary literal (10 in decimal)
octal = 0o52       // Octal literal (42 in decimal)
hex = 0x2A         // Hexadecimal literal (42 in decimal)

// Integer methods (automatic wrapping)
print(42.to_bin())     // "0b101010"
print(42.to_oct())     // "0o52"
print(42.to_hex())     // "0x2a"
print(42.is_even())    // True
print(17.is_prime())   // True
```

#### Float
```python
pi = 3.14159
scientific = 1.23e-4   // Scientific notation
negative_float = -2.5

// Float methods
print(3.14159.round(2))  // 3.14
print(16.0.sqrt())       // 4.0
print(3.14.sin())        // Sine using Taylor series
```

#### String
```python
name = "Alice"
multiline = """This is a
multi-line string"""

// String methods
print("hello".upper())        // "HELLO"
print("WORLD".lower())        // "world"
print("text".length())        // 4
print("hello".contains("ell")) // True
print("text".reverse())       // "txet"
```

#### Boolean
```python
is_true = True
is_false = False

// Boolean methods
print(True.to_int())     // 1
print(False.to_string()) // "False"
print(True.negate())     // False
```

#### None
```python
empty = None
print(type(empty))  // "NONE"
```

### Collection Types

#### Array
```python
numbers = [1, 2, 3, 4, 5]
mixed = [1, "hello", True, 3.14]
empty_array = []

// Array methods (automatic wrapping)
print([1, 2, 3].length())      // 3
print([3, 1, 4].sort())        // [1, 3, 4]
print([1, 2, 3].contains(2))   // True
print([1, 2, 3].reverse())     // [3, 2, 1]

// Array operations
numbers.append(6)              // [1, 2, 3, 4, 5, 6]
numbers[0] = 10               // [10, 2, 3, 4, 5, 6]
first = numbers[0]            // 10
last = numbers[-1]            // 6
```

#### Hash (Dictionary)
```python
person = {
    "name": "Alice",
    "age": 30,
    "city": "New York"
}

// Hash operations
person["email"] = "alice@example.com"
name = person["name"]
print(len(person))            // 4

// Iterate over hash
for key, value in pairs(person):
    print(f"{key}: {value}")
```

#### Tuple
```python
coordinates = (10, 20)
rgb_color = (255, 0, 128)
single_item = (42,)  // Note the comma for single-item tuple

// Tuple operations
x, y = coordinates           // Unpacking
print(coordinates[0])        // 10
print(len(rgb_color))        // 3
```

---

## Variables and Assignments

### Variable Declaration
```python
// Simple assignment
x = 10
name = "Alice"

// Multiple assignment
a, b, c = 1, 2, 3

// Tuple unpacking
point = (10, 20)
x, y = point

// Array unpacking
numbers = [1, 2, 3]
first, second, third = numbers
```

### Variable Scope
```python
global_var = "I'm global"

spell my_function():
    local_var = "I'm local"
    print(global_var)    // Accessible
    print(local_var)     // Accessible

print(global_var)        // Accessible
// print(local_var)      // Error: not accessible outside function
```

---

## Operators

### Arithmetic Operators
```python
a = 10
b = 3

print(a + b)    // Addition: 13
print(a - b)    // Subtraction: 7
print(a * b)    // Multiplication: 30
print(a / b)    // Division: 3.333...
print(a % b)    // Modulo: 1
print(a ** b)   // Exponentiation: 1000
```

### Comparison Operators
```python
x = 10
y = 20

print(x == y)   // Equal: False
print(x != y)   // Not equal: True
print(x < y)    // Less than: True
print(x <= y)   // Less than or equal: True
print(x > y)    // Greater than: False
print(x >= y)   // Greater than or equal: False
```

### Logical Operators
```python
a = True
b = False

print(a and b)  // Logical AND: False
print(a or b)   // Logical OR: True
print(not a)    // Logical NOT: False
```

### Assignment Operators
```python
x = 10
x += 5    // x = x + 5 → 15
x -= 3    // x = x - 3 → 12
x *= 2    // x = x * 2 → 24
x /= 4    // x = x / 4 → 6
x %= 4    // x = x % 4 → 2
```

---

## Control Flow

### Conditional Statements

#### if/otherwise/else
```python
age = 18

if age >= 21:
    print("Can drink alcohol")
otherwise age >= 18:
    print("Can vote")
else:
    print("Minor")

// Nested conditions
score = 85
if score >= 90:
    grade = "A"
otherwise score >= 80:
    grade = "B"
otherwise score >= 70:
    grade = "C"
else:
    grade = "F"
print(f"Grade: {grade}")
```

### Loops

#### for loops
```python
// Iterate over array
numbers = [1, 2, 3, 4, 5]
for num in numbers:
    print(num)

// Iterate over range
for i in range(5):
    print(i)  // 0, 1, 2, 3, 4

for i in range(2, 8):
    print(i)  // 2, 3, 4, 5, 6, 7

for i in range(0, 10, 2):
    print(i)  // 0, 2, 4, 6, 8

// Iterate over hash
person = {"name": "Alice", "age": 30}
for key, value in pairs(person):
    print(f"{key}: {value}")

// Iterate with index
items = ["a", "b", "c"]
for index, value in enumerate(items):
    print(f"{index}: {value}")
```

#### while loops
```python
count = 0
while count < 5:
    print(count)
    count += 1

// Infinite loop with break
while True:
    user_input = input("Enter 'quit' to exit: ")
    if user_input == "quit":
        stop  // break equivalent
    print(f"You entered: {user_input}")
```

#### Loop Control
```python
for i in range(10):
    if i == 3:
        skip  // continue equivalent
    if i == 7:
        stop  // break equivalent
    print(i)  // Prints: 0, 1, 2, 4, 5, 6
```

### Pattern Matching

#### match/case statements
```python
spell process_command(command):
    match command:
        case "start":
            print("Starting application...")
        case "stop":
            print("Stopping application...")
        case "restart":
            print("Restarting application...")
        case "status":
            print("Application is running")
        _:  // Default case
            print("Unknown command")

// Pattern matching with values
spell check_number(num):
    match num:
        case 0:
            return "Zero"
        case 1:
            return "One"
        case n if n > 10:
            return "Large number"
        _:
            return "Small number"
```

---

## Functions

### Function Definition
```python
spell greet():
    print("Hello, World!")

spell add(a, b):
    return a + b

spell greet_person(name, age=25):  // Default parameter
    print(f"Hello {name}, you are {age} years old")

// Function calls
greet()
result = add(5, 3)
greet_person("Alice")
greet_person("Bob", 30)
```

### Variable Arguments
```python
spell sum_all(*numbers):
    total = 0
    for num in numbers:
        total += num
    return total

result = sum_all(1, 2, 3, 4, 5)  // 15
```

### Lambda Functions
```python
// Anonymous functions
square = λ x: x * x
print(square(5))  // 25

add = λ x, y: x + y
print(add(3, 4))  // 7

// Higher-order functions
spell apply_function(func, value):
    return func(value)

result = apply_function(λ x: x * 2, 10)  // 20
```

---

## Grimoires (Classes)

### Basic Grimoire Definition
```python
grim Person:
    init(name, age):
        self.name = name
        self.age = age
        self.friends = []
    
    spell greet():
        return f"Hello, I'm {self.name}"
    
    spell add_friend(friend):
        self.friends.append(friend)
        print(f"{friend} is now a friend of {self.name}")
    
    spell birthday():
        self.age += 1
        print(f"{self.name} is now {self.age} years old")

// Using the grimoire
person = Person("Alice", 25)
print(person.greet())  // "Hello, I'm Alice"
person.add_friend("Bob")
person.birthday()
```

### Grimoire with Class Methods
```python
grim Calculator:
    init():
        self.history = []
    
    spell add(a, b):
        result = a + b
        self.history.append(f"{a} + {b} = {result}")
        return result
    
    spell multiply(a, b):
        result = a * b
        self.history.append(f"{a} * {b} = {result}")
        return result
    
    spell get_history():
        return self.history
    
    spell clear_history():
        self.history = []

calc = Calculator()
print(calc.add(5, 3))      // 8
print(calc.multiply(4, 6)) // 24
print(calc.get_history())  // ["5 + 3 = 8", "4 * 6 = 24"]
```

### Static Methods and Properties
```python
grim MathUtils:
    PI = 3.14159
    E = 2.71828
    
    spell circle_area(radius):
        return MathUtils.PI * radius * radius
    
    spell factorial(n):
        if n <= 1:
            return 1
        return n * MathUtils.factorial(n - 1)

// Using static methods
area = MathUtils.circle_area(5)
fact = MathUtils.factorial(5)
print(f"Area: {area}, Factorial: {fact}")
```

---

## Error Handling

### attempt/ensnare/resolve Blocks
```python
// Basic error handling
attempt:
    result = 10 / 0
    print("This won't print")
ensnare:
    print("Division by zero error caught")

// Error handling with file operations
attempt:
    content = File.read("nonexistent.txt")
    print(content)
ensnare:
    print("File not found or cannot read")
resolve:
    print("This always executes")

// Multiple error types
attempt:
    value = int(input("Enter a number: "))
    result = 100 / value
    print(f"Result: {result}")
ensnare:
    print("Invalid input or division by zero")
resolve:
    print("Operation complete")
```

### Custom Errors
```python
spell validate_age(age):
    if age < 0:
        raise Error("ValidationError", "Age cannot be negative")
    if age > 150:
        raise Error("ValidationError", "Age seems unrealistic")
    return True

attempt:
    validate_age(-5)
ensnare:
    print("Age validation failed")
```

---

## Built-in Functions

### Type Conversion Functions
```python
// String to number conversion
print(int("42"))        // 42
print(float("3.14"))    // 3.14
print(str(123))         // "123"
print(bool(1))          // True

// Collection conversion
print(list("hello"))         // ["h", "e", "l", "l", "o"]
print(tuple([1, 2, 3]))      // (1, 2, 3)
print(list(range(5)))        // [0, 1, 2, 3, 4]
```

### Utility Functions
```python
// Length and type checking
print(len("hello"))          // 5
print(len([1, 2, 3]))        // 3
print(type(42))              // "INTEGER"
print(type("hello"))         // "STRING"

// Mathematical functions
print(abs(-10))              // 10
print(max(1, 5, 3))          // 5
print(min([2, 8, 1]))        // 1

// Character/ASCII functions
print(ord("A"))              // 65
print(chr(65))               // "A"
```

### Input/Output Functions
```python
// User input
name = input("Enter your name: ")
age = int(input("Enter your age: "))

// Formatted output
print(f"Hello {name}, you are {age} years old")
```

### Collection Functions
```python
// Enumeration
items = ["a", "b", "c"]
for index, value in enumerate(items):
    print(f"{index}: {value}")

// Hash iteration
data = {"x": 10, "y": 20}
for key in pairs(data, "key"):
    print(f"Key: {key}")

for value in pairs(data, "value"):
    print(f"Value: {value}")

for key, value in pairs(data):
    print(f"{key} = {value}")
```

---

## Standard Library (Munin)

The Munin standard library provides enhanced functionality through grimoires that wrap primitive types.

### Array Grimoire
```python
// Create enhanced array
arr = Array([3, 1, 4, 1, 5])

// Array operations
print(arr.length())         // 5
arr.append(9)              // [3, 1, 4, 1, 5, 9]
arr.remove(1)              // [3, 4, 1, 5, 9] (removes first occurrence)

// Search and access
print(arr.contains(4))      // True
print(arr.index_of(4))      // 1
print(arr.first())          // 3
print(arr.last())           // 9

// Array transformations
sorted_arr = arr.sort()     // [1, 3, 4, 5, 9]
reversed_arr = arr.reverse() // [9, 5, 1, 4, 3]
slice_arr = arr.slice(1, 4) // [1, 4, 1]

// Array utilities
print(arr.is_empty())       // False
arr.clear()                // []
print(arr.is_empty())       // True
```

### String Grimoire
```python
// Enhanced string operations
text = String("Hello World")

// Basic properties
print(text.length())        // 11
print(text.to_string())     // "Hello World"

// Case conversion
print(text.upper())         // "HELLO WORLD"
print(text.lower())         // "hello world"

// Search operations
print(text.find("World"))   // 6
print(text.contains("ell")) // True
print(text.char_at(0))      // "H"
print(text.char_at(-1))     // "d"

// String transformation
print(text.reverse())       // "dlroW olleH"
```

### Integer Grimoire
```python
// Enhanced integer operations
num = Integer(42)

// Number base conversions
print(num.to_bin())         // "0b101010"
print(num.to_oct())         // "0o52"
print(num.to_hex())         // "0x2a"

// Mathematical operations
print(num.abs())            // 42
print(num.pow(2))           // 1764
print(num.gcd(18))          // 6
print(num.lcm(18))          // 126

// Number properties
print(num.is_even())        // True
print(num.is_odd())         // False
print(num.is_prime())       // False

// Type conversion
print(num.to_float())       // 42.0
print(num.to_string())      // "42"
```

### Float Grimoire
```python
// Enhanced float operations
pi = Float(3.14159)

// Rounding and precision
print(pi.round(2))          // 3.14
print(pi.floor())           // 3
print(pi.ceil())            // 4

// Mathematical functions
print(pi.abs())             // 3.14159
print(pi.sqrt())            // 1.772 (approx)
print(pi.pow(2))            // 9.8696 (approx)
print(pi.sin())             // 0.0 (approx)
print(pi.cos())             // -1.0 (approx)

// Type checking
print(pi.is_positive())     // True
print(pi.is_negative())     // False
print(pi.is_zero())         // False
print(pi.is_integer())      // False

// Conversions
print(pi.to_int())          // 3
print(pi.to_string())       // "3.14159"
```

### Boolean Grimoire
```python
// Enhanced boolean operations
flag = Boolean(True)

// Basic properties
print(flag.to_int())        // 1
print(flag.to_string())     // "True"
print(flag.is_true())       // True
print(flag.is_false())      // False

// Logical operations
print(flag.negate())        // False
print(flag.and_with(False)) // False
print(flag.or_with(False))  // True
print(flag.xor_with(True))  // False
print(flag.implies(False))  // False (logical implication)
```

---

## File Operations

The File grimoire provides comprehensive file system operations.

### Static File Methods (Recommended)
```python
// Basic file operations
File.write("data.txt", "Hello, World!")
content = File.read("data.txt")
print(content)  // "Hello, World!"

// Append to file
File.append("log.txt", "Log entry 1\n")
File.append("log.txt", "Log entry 2\n")

// Check file existence
if File.exists("config.txt"):
    config = File.read("config.txt")
else:
    File.write("config.txt", "default_setting=true")
```

### File Objects for Complex Operations
```python
// Reading files with autoclose
autoclose File.open("data.txt", "r") as file:
    content = file.read_content()
    print(f"File content: {content}")

// Writing files with autoclose
autoclose File.open("output.txt", "w") as file:
    file.write_content("Line 1\n")
    file.write_content("Line 2\n")
    file.write_content("Line 3\n")

// Appending to files
autoclose File.open("log.txt", "a") as file:
    file.write_content("New log entry\n")
    file.write_content("Another entry\n")
```

### File Processing Examples
```python
// Process CSV-like data
File.write("users.csv", "Alice,25,Engineer\nBob,30,Designer\nCharlie,35,Manager")

autoclose File.open("users.csv", "r") as file:
    content = file.read_content()
    lines = content.split("\n")
    
    for line in lines:
        if line.strip():
            parts = line.split(",")
            name, age, job = parts[0], parts[1], parts[2]
            print(f"Name: {name}, Age: {age}, Job: {job}")

// Create report file
autoclose File.open("report.txt", "w") as report:
    report.write_content("Daily Report\n")
    report.write_content("============\n\n")
    
    users = ["Alice", "Bob", "Charlie"]
    for user in users:
        report.write_content(f"User: {user}\n")
        report.write_content(f"Status: Active\n\n")

print("Report generated:", File.exists("report.txt"))
```

---

## OS Operations

The OS grimoire provides operating system interface.

### Directory Operations
```python
// Get current directory
current = OS.cwd()
print(f"Working in: {current}")

// List directory contents
files = OS.listdir(".")
print(f"Found {len(files)} items:")
for filename in files:
    print(f"  - {filename}")

// Create and remove directories
OS.mkdir("temp_dir")
print("Directory created:", File.exists("temp_dir"))

// Create nested structure
OS.mkdir("project/src", 0755)  // Custom permissions
File.write("project/src/main.crl", "print('Hello from project!')")

// Change directory
original_dir = OS.cwd()
OS.chdir("project")
print("Changed to:", OS.cwd())
OS.chdir(original_dir)

// Cleanup
OS.remove("project/src/main.crl")
OS.remove("project/src")
OS.remove("project")
OS.remove("temp_dir")
```

### Environment Variables
```python
// Get environment variables
home = OS.getenv("HOME")
user = OS.getenv("USER")
path = OS.getenv("PATH")

print(f"Home: {home}")
print(f"User: {user}")
print(f"PATH length: {len(path)}")

// Set environment variables (both arguments must be strings)
OS.setenv("MY_APP", "carrion_app")
OS.setenv("DEBUG_MODE", "true")
OS.setenv("PORT", "8080")  // Numbers must be strings

// Verify they were set
app_name = OS.getenv("MY_APP")
debug = OS.getenv("DEBUG_MODE")
port = OS.getenv("PORT")

print(f"App: {app_name}, Debug: {debug}, Port: {port}")

// Expand environment variables
config_path = OS.expandEnv("$HOME/.config/myapp")
log_path = OS.expandEnv("$MY_APP/logs")
print(f"Config path: {config_path}")
```

### System Commands
```python
// Run commands and show output directly
OS.run("echo", ["Hello from system!"], False)

// Capture command output
date_output = OS.run("date", [], True)
print(f"Current date: {date_output}")

// List files using system command
file_list = OS.run("ls", ["-la", "."], True)
print("Directory listing captured")

// Execute with arguments
OS.run("mkdir", ["temp_folder"], False)
OS.run("touch", ["temp_folder/test.txt"], False)
listing = OS.run("ls", ["temp_folder"], True)
print(f"Temp folder contents: {listing}")

// Cleanup
OS.run("rm", ["-rf", "temp_folder"], False)
```

### Process Management
```python
// Sleep/delay operations
print("Starting long process...")
OS.sleep(1)  // Wait 1 second
print("Process step 1 complete")
OS.sleep(0.5)  // Wait 500ms
print("Process step 2 complete")
OS.sleep(2)  // Wait 2 seconds
print("Process complete!")
```

---

## Module System

### Basic Imports
```python
// Import entire module
import "math_utils"
result = add(5, 3)  // Function from math_utils.crl

// Import specific grimoire
import "data_structures.Stack"
stack = Stack()
stack.push(1)
stack.push(2)
print(stack.pop())  // 2

// Import with alias
import "very_long_module_name" as short
import "utilities.FileHelper" as FH
helper = FH()
```

### Package Imports
```python
// Simplified package imports (auto-resolves versions)
import "json-utils/parser"                 // Auto-resolves to latest version
import "json-utils/parser.JSONParser" as JSON
import "http-client/request.HTTPClient" as HTTP

// Use imported functionality
json_parser = JSON()
data = json_parser.parse('{"name": "example"}')

http = HTTP()
response = http.get("https://api.example.com/data")
```

### Module Organization
```python
// File: utils/string_utils.crl
spell capitalize_words(text):
    words = text.split(" ")
    result = []
    for word in words:
        if len(word) > 0:
            result.append(word[0].upper() + word[1:].lower())
    return " ".join(result)

grim TextProcessor:
    init(text):
        self.text = text
    
    spell process():
        return capitalize_words(self.text)

// File: main.crl
import "utils.string_utils"
import "utils.string_utils.TextProcessor" as TP

text = "hello world from carrion"
formatted = capitalize_words(text)
print(formatted)  // "Hello World From Carrion"

processor = TP("another test string")
result = processor.process()
print(result)  // "Another Test String"
```

---

## Advanced Features

### String Interpolation (F-Strings)
```python
name = "Alice"
age = 30
height = 5.8

// Basic interpolation
message = f"Hello {name}!"
print(message)  // "Hello Alice!"

// Complex expressions
info = f"{name} is {age} years old and {height} feet tall"
print(info)

// Calculations in f-strings
x = 10
y = 20
result = f"The sum of {x} and {y} is {x + y}"
print(result)  // "The sum of 10 and 20 is 30"

// Method calls in f-strings
text = "hello"
formatted = f"Uppercase: {text.upper()}, Length: {len(text)}"
print(formatted)  // "Uppercase: HELLO, Length: 5"
```

### Automatic Primitive Wrapping
```python
// Primitives automatically get grimoire methods
print(42.to_bin())          // "0b101010"
print("hello".upper())      // "HELLO"
print(3.14.round(2))        // 3.14
print(True.to_int())        // 1
print([3, 1, 4].sort())     // [1, 3, 4]

// This works because Carrion automatically wraps primitives
number = 17
print(number.is_prime())    // True
print(number.to_hex())      // "0x11"

text = "Hello World"
print(text.contains("World"))  // True
print(text.reverse())          // "dlroW olleH"
```

### Interactive Help System
```python
// Built-in help functions
help()      // Shows general help
version()   // Shows version information
modules()   // Lists available modules
```

---

## Complete Examples

### Example 1: File Processing Application
```python
// File processor that reads, processes, and saves data
grim FileProcessor:
    init(input_file, output_file):
        self.input_file = input_file
        self.output_file = output_file
        self.stats = {"lines": 0, "words": 0, "chars": 0}
    
    spell process():
        attempt:
            // Read input file
            if not File.exists(self.input_file):
                raise Error("FileError", f"Input file {self.input_file} not found")
            
            content = File.read(self.input_file)
            lines = content.split("\n")
            
            // Process content
            processed_lines = []
            for line in lines:
                if line.strip():  // Skip empty lines
                    processed_line = self.process_line(line)
                    processed_lines.append(processed_line)
                    self.update_stats(line)
            
            // Write output
            processed_content = "\n".join(processed_lines)
            File.write(self.output_file, processed_content)
            
            print(f"Processing complete!")
            print(f"Processed {self.stats['lines']} lines")
            print(f"Total words: {self.stats['words']}")
            print(f"Total characters: {self.stats['chars']}")
            
        ensnare:
            print("Error during file processing")
    
    spell process_line(line):
        // Capitalize first letter of each word
        words = line.split(" ")
        processed_words = []
        for word in words:
            if len(word) > 0:
                processed_words.append(word[0].upper() + word[1:].lower())
        return " ".join(processed_words)
    
    spell update_stats(line):
        self.stats["lines"] += 1
        self.stats["words"] += len(line.split(" "))
        self.stats["chars"] += len(line)

// Create sample input file
input_data = """hello world
this is a test file
with multiple lines
some have MIXED case
others are lowercase"""

File.write("input.txt", input_data)

// Process the file
processor = FileProcessor("input.txt", "output.txt")
processor.process()

// Show results
if File.exists("output.txt"):
    result = File.read("output.txt")
    print("\nProcessed content:")
    print(result)

// Cleanup
OS.remove("input.txt")
OS.remove("output.txt")
```

### Example 2: Data Analysis Tool
```python
// Simple data analysis with arrays and hashes
grim DataAnalyzer:
    init():
        self.data = []
    
    spell load_data(filename):
        attempt:
            content = File.read(filename)
            lines = content.split("\n")
            
            for line in lines:
                if line.strip():
                    parts = line.split(",")
                    if len(parts) >= 3:
                        record = {
                            "name": parts[0].strip(),
                            "age": int(parts[1].strip()),
                            "salary": float(parts[2].strip())
                        }
                        self.data.append(record)
        ensnare:
            print("Error loading data")
    
    spell analyze():
        if len(self.data) == 0:
            print("No data to analyze")
            return
        
        // Calculate statistics
        total_people = len(self.data)
        ages = []
        salaries = []
        
        for person in self.data:
            ages.append(person["age"])
            salaries.append(person["salary"])
        
        avg_age = sum(ages) / len(ages)
        avg_salary = sum(salaries) / len(salaries)
        min_age = min(ages)
        max_age = max(ages)
        min_salary = min(salaries)
        max_salary = max(salaries)
        
        // Display results
        print(f"Data Analysis Results:")
        print(f"Total people: {total_people}")
        print(f"Age range: {min_age} - {max_age} (avg: {avg_age:.1f})")
        print(f"Salary range: ${min_salary:.2f} - ${max_salary:.2f} (avg: ${avg_salary:.2f})")
        
        // Find highest paid person
        highest_paid = None
        for person in self.data:
            if highest_paid == None or person["salary"] > highest_paid["salary"]:
                highest_paid = person
        
        print(f"Highest paid: {highest_paid['name']} (${highest_paid['salary']:.2f})")
    
    spell export_summary(filename):
        autoclose File.open(filename, "w") as file:
            file.write_content("Data Analysis Summary\n")
            file.write_content("====================\n\n")
            
            for person in self.data:
                line = f"{person['name']}: Age {person['age']}, Salary ${person['salary']:.2f}\n"
                file.write_content(line)

// Create sample data
sample_data = """Alice,25,55000.00
Bob,30,62000.50
Charlie,35,75000.00
Diana,28,58000.75
Eve,32,68000.25"""

File.write("employees.csv", sample_data)

// Analyze the data
analyzer = DataAnalyzer()
analyzer.load_data("employees.csv")
analyzer.analyze()
analyzer.export_summary("summary.txt")

print("\nSummary exported to summary.txt")

// Cleanup
OS.remove("employees.csv")
OS.remove("summary.txt")
```

### Example 3: System Information Gatherer
```python
// System information and environment analysis
grim SystemInfo:
    init():
        self.info = {}
    
    spell gather_info():
        // Basic system info
        self.info["current_dir"] = OS.cwd()
        self.info["home_dir"] = OS.getenv("HOME")
        self.info["user"] = OS.getenv("USER")
        self.info["path"] = OS.getenv("PATH")
        
        // Directory contents
        files = OS.listdir(".")
        self.info["file_count"] = len(files)
        self.info["files"] = files
        
        // System commands
        attempt:
            self.info["date"] = OS.run("date", [], True).strip()
            self.info["uptime"] = OS.run("uptime", [], True).strip()
        ensnare:
            self.info["date"] = "Unknown"
            self.info["uptime"] = "Unknown"
    
    spell display_info():
        print("System Information Report")
        print("========================")
        print(f"Current Directory: {self.info['current_dir']}")
        print(f"Home Directory: {self.info['home_dir']}")
        print(f"Current User: {self.info['user']}")
        print(f"Date: {self.info['date']}")
        print(f"Files in current directory: {self.info['file_count']}")
        
        if self.info["file_count"] > 0:
            print("Files:")
            for i, filename in enumerate(self.info["files"]):
                if i < 10:  // Show first 10 files
                    print(f"  - {filename}")
            if len(self.info["files"]) > 10:
                print(f"  ... and {len(self.info['files']) - 10} more")
    
    spell save_report(filename):
        autoclose File.open(filename, "w") as file:
            file.write_content("System Information Report\n")
            file.write_content("========================\n\n")
            
            for key, value in pairs(self.info):
                if key != "files":
                    file.write_content(f"{key}: {value}\n")
            
            file.write_content(f"\nDirectory Contents ({len(self.info['files'])} files):\n")
            for filename in self.info["files"]:
                file.write_content(f"  - {filename}\n")

// Generate system report
sys_info = SystemInfo()
sys_info.gather_info()
sys_info.display_info()

print("\nSaving detailed report...")
sys_info.save_report("system_report.txt")
print("Report saved to system_report.txt")

// Show report was created
if File.exists("system_report.txt"):
    size = len(File.read("system_report.txt"))
    print(f"Report file size: {size} characters")

// Cleanup
OS.remove("system_report.txt")
```

---

## Best Practices

### Code Organization
```python
// 1. Use descriptive variable and function names
spell calculate_circle_area(radius):
    return 3.14159 * radius * radius

// 2. Keep functions focused and small
spell process_user_input():
    name = input("Enter name: ")
    age = int(input("Enter age: "))
    return validate_user_data(name, age)

spell validate_user_data(name, age):
    if len(name) == 0:
        return False, "Name cannot be empty"
    if age < 0 or age > 150:
        return False, "Invalid age"
    return True, "Valid"

// 3. Use error handling appropriately
spell safe_file_operation(filename):
    attempt:
        content = File.read(filename)
        return True, content
    ensnare:
        return False, "File operation failed"
```

### Performance Tips
```python
// 1. Use static methods for simple operations
content = File.read("data.txt")  // Preferred
// vs complex file objects for simple reads

// 2. Use appropriate data structures
// Use arrays for ordered collections
numbers = [1, 2, 3, 4, 5]

// Use hashes for key-value mappings
person = {"name": "Alice", "age": 30}

// 3. Minimize file operations
data = []
autoclose File.open("large_file.txt", "r") as file:
    content = file.read_content()
    // Process all data at once
    lines = content.split("\n")
    for line in lines:
        data.append(process_line(line))
```

### Security Considerations
```python
// 1. Validate user input
spell safe_user_input(prompt, input_type):
    attempt:
        user_input = input(prompt)
        if input_type == "int":
            return int(user_input)
        otherwise input_type == "float":
            return float(user_input)
        return user_input
    ensnare:
        print("Invalid input format")
        return None

// 2. Be careful with file paths
spell safe_file_read(filename):
    // Prevent path traversal
    if ".." in filename or filename.startswith("/"):
        print("Invalid file path")
        return None
    
    attempt:
        return File.read(filename)
    ensnare:
        print("File read error")
        return None

// 3. Handle environment variables safely
spell get_config_value(key, default):
    value = OS.getenv(key)
    if value == None or value == "":
        return default
    return value
```

---

This complete reference covers all major aspects of the Carrion programming language. Use it as a comprehensive guide for development, from basic syntax to advanced features and best practices.