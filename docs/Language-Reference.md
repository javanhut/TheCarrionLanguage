# Carrion Language Reference

This comprehensive reference guide covers all aspects of the Carrion programming language, from basic syntax to advanced features.

## Table of Contents

1. [Language Overview](#language-overview)
2. [Lexical Structure](#lexical-structure)
3. [Data Types](#data-types)
4. [Operators](#operators)
5. [Expressions](#expressions)
6. [Statements](#statements)
7. [Functions (Spells)](#functions-spells)
8. [Object-Oriented Programming (Grimoires)](#object-oriented-programming-grimoires)
9. [Error Handling](#error-handling)
10. [Modules and Imports](#modules-and-imports)
11. [Built-in Functions](#built-in-functions)
12. [Standard Library (Munin)](#standard-library-munin)

## Language Overview

Carrion is a dynamically typed, interpreted programming language with a Norse mythology and magical theme. It combines familiar programming concepts with enchanting terminology while maintaining readable and practical syntax.

### Key Characteristics
- **Dynamic Typing**: Variables don't require explicit type declarations
- **Interpreted**: Code is executed directly without compilation
- **Object-Oriented**: Full support for classes (grimoires) and inheritance
- **Duck Typing**: Objects are characterized by their behavior, not their type
- **Memory Managed**: Automatic memory management via Go's garbage collector

### File Extension
Carrion source files use the `.crl` extension.

## Lexical Structure

### Comments
```python
"/#" Single-line comment
``` Multi-line
   comment```
```

### Identifiers
Identifiers must start with a letter or underscore, followed by letters, digits, or underscores:
```python
variable_name
_private_var
MyClass
function123
```

### Keywords
Reserved words that cannot be used as identifiers:
```
attempt      ensnare     resolve     raise       check
if           otherwise   else        for         in
not          while       skip        stop        return
match        case        and         or          True
False        None        grim        spell       init
self         super       arcane      arcanespell import
as           var         ignore      autoclose
```

### Literals

#### Integer Literals
```python
42          // Decimal
0b1010      // Binary (not in source, but supported in display)
0o52        // Octal (not in source, but supported in display)
0x2A        // Hexadecimal (not in source, but supported in display)
```

#### Float Literals
```python
3.14
2.718
1.0
.5          // 0.5
2e10        // Scientific notation (if supported)
```

#### String Literals
```python
"double quotes"
'single quotes'
"""triple quotes for 
   multi-line strings"""
f"formatted {variable}"      // F-strings
i"interpolated {expression}" // Interpolated strings
```

#### Boolean Literals
```python
True
False
```

#### None Literal
```python
None
```

## Data Types

### Primitive Types

#### Integer
64-bit signed integers:
```python
age = 25
count = -10
big_number = 1000000
```

#### Float
64-bit floating-point numbers:
```python
pi = 3.14159
temperature = -15.5
rate = 0.075
```

#### String
UTF-8 text strings:
```python
name = "Alice"
message = 'Hello, World!'
description = """This is a
multi-line string"""
```

#### Boolean
True/False values:
```python
is_active = True
has_permission = False
```

#### None
Represents absence of value:
```python
result = None
optional_param = None
```

### Collection Types

#### Array
Ordered, mutable sequences:
```python
numbers = [1, 2, 3, 4, 5]
mixed = [1, "hello", True, 3.14]
empty = []
```

#### Map
Key-value mappings (dictionaries) with support for multiple key types:
```python
person = {"name": "Alice", "age": 30}  // String keys
config = {"debug": True, "timeout": 30}  // String keys
mixed_map = {"name": "Alice", 42: "answer", 3.14: "pi", True: "enabled"}  // Mixed key types
empty_map = {}
```

**Supported Key Types:**
- **String**: `"key"`, `'key'`
- **Integer**: `42`, `-17`
- **Float**: `3.14`, `-2.5`
- **Boolean**: `True`, `False`

#### Tuple
Immutable ordered sequences:
```python
coordinates = (10, 20)
rgb = (255, 128, 0)
single = (42,)  // Single-element tuple
```

### Type Checking
```python
value = 42
print(type(value))  // → "INTEGER"

if type(value) == "INTEGER":
    print("It's an integer")
```

## Operators

### Arithmetic Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `+` | Addition | `5 + 3` → `8` |
| `-` | Subtraction | `5 - 3` → `2` |
| `*` | Multiplication | `5 * 3` → `15` |
| `/` | Division | `15 / 3` → `5.0` |
| `//` | Integer Division | `17 // 3` → `5` |
| `%` | Modulo | `17 % 3` → `2` |
| `**` | Exponentiation | `2 ** 3` → `8` |

### Assignment Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `=` | Assignment | `x = 5` |
| `+=` | Add and assign | `x += 3` |
| `-=` | Subtract and assign | `x -= 3` |
| `*=` | Multiply and assign | `x *= 3` |
| `/=` | Divide and assign | `x /= 3` |
| `++` | Increment | `x++` or `++x` |
| `--` | Decrement | `x--` or `--x` |

### Comparison Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `==` | Equal | `5 == 5` → `True` |
| `!=` | Not equal | `5 != 3` → `True` |
| `<` | Less than | `3 < 5` → `True` |
| `>` | Greater than | `5 > 3` → `True` |
| `<=` | Less than or equal | `3 <= 5` → `True` |
| `>=` | Greater than or equal | `5 >= 5` → `True` |

### Logical Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `and` | Logical AND | `True and False` → `False` |
| `or` | Logical OR | `True or False` → `True` |
| `not` | Logical NOT | `not True` → `False` |

### Membership Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `in` | Membership test | `"a" in "apple"` → `True` |
| `not in` | Negative membership | `"z" not in "apple"` → `True` |

### Bitwise Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `&` | Bitwise AND | `5 & 3` → `1` |
| `\|` | Bitwise OR | `5 \| 3` → `7` |
| `^` | Bitwise XOR | `5 ^ 3` → `6` |
| `~` | Bitwise NOT | `~5` → `-6` |
| `<<` | Left shift | `5 << 1` → `10` |
| `>>` | Right shift | `10 >> 1` → `5` |

## Expressions

### Operator Precedence
1. Parentheses: `()`
2. Exponentiation: `**`
3. Unary: `+`, `-`, `not`, `~`
4. Multiplicative: `*`, `/`, `//`, `%`
5. Additive: `+`, `-`
6. Shift: `<<`, `>>`
7. Bitwise AND: `&`
8. Bitwise XOR: `^`
9. Bitwise OR: `|`
10. Comparison: `<`, `<=`, `>`, `>=`, `==`, `!=`, `in`, `not in`
11. Logical NOT: `not`
12. Logical AND: `and`
13. Logical OR: `or`
14. Assignment: `=`, `+=`, `-=`, `*=`, `/=`

### Expression Examples
```python
// Arithmetic expressions
result = 2 + 3 * 4        // → 14
result = (2 + 3) * 4      // → 20

// Boolean expressions
valid = age >= 18 and has_license
can_proceed = user.is_admin() or user.has_permission("write")

// String expressions
full_name = first_name + " " + last_name
greeting = f"Hello, {name}!"

// Collection expressions
squares = [x ** 2 for x in range(5)]  // If list comprehension is supported
```

## Statements

### Assignment Statements
```python
// Simple assignment
x = 42
name = "Alice"

// Multiple assignment (tuple unpacking)
x, y = (10, 20)
a, b, c = [1, 2, 3]

// Compound assignment
x += 5
count *= 2
```

### Expression Statements
```python
// Function calls
print("Hello")
result = calculate(10, 20)

// Method calls
text.upper()
array.append(item)
```

### Control Flow Statements

#### If Statements
```python
if condition:
    // code
otherwise another_condition:
    // code
else:
    // code
```

#### For Loops
```python
for item in iterable:
    // code
else:
    // optional else clause
```

#### While Loops
```python
while condition:
    // code
```

#### Match Statements
```python
match value:
    case pattern1:
        // code
    case pattern2:
        // code
    _:
        // default case
```

#### Loop Control
```python
skip  // Continue to next iteration
stop  // Break from loop
```

#### Return Statements
```python
return            // Return None
return value      // Return specific value
```

#### Autoclose Statements
Automatic resource management for objects that need cleanup:
```python
autoclose expression as variable:
    // code block
    // variable.close() is called automatically
```

The `autoclose` statement ensures that resources are properly cleaned up when the block exits, even if an error occurs. It works with any object that has a `close()` method.

Examples:
```python
// File operations with File grimoire (recommended)
autoclose File.open("data.txt", "r") as file:
    content = file.read_content()
    print(content)

// Writing files with File grimoire
autoclose File.open("output.txt", "w") as file:
    file.write_content("Hello, World!")

// Appending to files with File grimoire
autoclose File.open("log.txt", "a") as file:
    file.write_content("New entry\n")

// Alternative: using open() builtin (less preferred)
autoclose open("data.txt", "r") as file:
    content = file.read_content()
    print(content)
```

## Functions (Spells)

### Function Definition
```python
spell function_name(parameters):
    // function body
    return value  // optional
```

### Parameters
```python
// Basic parameters
spell greet(name):
    return f"Hello, {name}!"

// Default parameters
spell power(base, exponent = 2):
    return base ** exponent

// Variable arguments (if supported)
spell sum_all(*numbers):
    total = 0
    for num in numbers:
        total += num
    return total
```

### Function Examples
```python
spell factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)

spell is_prime(number):
    if number < 2:
        return False
    for i in range(2, int(number ** 0.5) + 1):
        if number % i == 0:
            return False
    return True
```

## Object-Oriented Programming (Grimoires)

### Class Definition
```python
grim ClassName:
    init(parameters):
        // constructor
    
    spell method_name(parameters):
        // method implementation
```

### Inheritance
```python
grim Child(Parent):
    init(parameters):
        super.init(parent_parameters)
        // child-specific initialization
    
    spell method_name(parameters):
        // override or new method
```

### Abstract Classes
```python
arcane grim AbstractClass:
    @arcanespell
    spell abstract_method():
        ignore  // no implementation
```

### Access Modifiers
- Public: `self.public_attribute`
- Protected: `self._protected_attribute`
- Private: `self.__private_attribute`

## Error Handling

### Basic Syntax
```python
attempt:
    // risky code
ensnare (ErrorType):
    // handle specific error
ensnare:
    // handle any error
resolve:
    // finally block
```

### Raising Errors
```python
raise Error("ErrorType", "Error message")
```

### Assertions
```python
check(condition, "Error message")
```

## Modules and Imports

### Import Syntax
```python
# Smart import patterns
import "GrimoireName"                # NEW: Grimoire-based import (searches all locations)
import "GrimoireName" as alias       # NEW: Grimoire import with alias
import "filename"                    # Local file
import "filename.ClassName"          # Selective import
import "package/module"              # Simplified package import
import "package/module.ClassName"    # Package selective import
import "./filename"                  # Relative current directory
import "../filename"                 # Relative parent directory
import "module" as alias             # Import with alias
import "module.ClassName" as alias   # Selective import with alias
```

### Module Structure
Each `.crl` file is a module that can export functions and grimoires (classes). The import system automatically resolves packages with version management and searches appropriate directories.

### Grimoire-Based Imports
The enhanced import system allows importing grimoires (classes) directly by name:
- Searches current directory, project modules, global bifrost modules, and system packages
- Automatically locates grimoire definitions without requiring file paths
- Supports aliasing for convenient naming
- Works with both local files and installed packages

## Built-in Functions

### Type Conversion
- `int(value)` - Convert to integer
- `float(value)` - Convert to float
- `str(value)` - Convert to string
- `bool(value)` - Convert to boolean
- `list(iterable)` - Convert to array
- `tuple(iterable)` - Convert to tuple

### Utility Functions
- `len(object)` - Get length
- `type(object)` - Get type
- `print(*args)` - Print values
- `input(prompt)` - Get user input
- `range(start, stop, step)` - Generate number sequence
- `max(*args)` - Find maximum
- `abs(value)` - Absolute value
- `ord(char)` - Get ASCII code
- `chr(code)` - Get character from ASCII

### Collection Functions
- `enumerate(array)` - Get indexed pairs
- `pairs(map, filter)` - Get key-value pairs (enhanced: supports "key"/"k" and "value"/"v" filters)
- `is_sametype(obj1, obj2)` - Compare types

### JSON Parsing Functions
- `parseHash(jsonString)` - Parse JSON object string into a Carrion map (dictionaries only)
  - Only accepts valid JSON objects `{...}`
  - Returns an error for arrays, primitives, or invalid JSON
  - Example: `parseHash('{"name": "Alice", "age": 30}')`

### Iteration Support
All collection types support the `in` operator and iteration with `for` loops:
- **Strings**: Character-by-character iteration and substring checking
- **Arrays**: Element iteration and membership testing  
- **Maps**: Key iteration by default, key membership testing
- **Ranges**: Number sequence iteration

## Standard Library (Munin)

### Core Modules
- **Array** - Enhanced array operations with methods like `append()`, `sort()`, `reverse()`
- **String** - String manipulation with methods like `upper()`, `lower()`, `find()`
- **Integer** - Integer utilities and conversions with methods like `to_bin()`, `is_prime()`, `gcd()`
- **Float** - Floating-point operations with methods like `round()`, `sqrt()`, `sin()`, `cos()`
- **Boolean** - Boolean logic operations with methods like `to_int()`, `negate()`
- **File** - File I/O operations using static methods like `File.read()`, `File.write()`, `File.open()`
- **OS** - Operating system interface using static methods like `OS.cwd()`, `OS.listdir()`, `OS.run()`
- **HTTP** - HTTP client operations with JSON support:
  - `httpGet(url, headers)` - Make GET request
  - `httpPost(url, body, headers)` - Make POST request
  - `httpParseJSON(jsonString)` - Parse any JSON string into Carrion objects
  - `httpStringifyJSON(object)` - Convert Carrion objects to JSON strings

### Standard Functions
- `help()` - Get help information
- `version()` - Version information
- `modules()` - List available modules

### File and OS Operations
The File and OS grimoires provide comprehensive system operations:

```python
// File operations with static methods (recommended)
content = File.read("config.txt")
File.write("output.txt", "Hello")
File.append("log.txt", "New entry\n")
exists = File.exists("data.txt")

// OS operations with static methods
current_dir = OS.cwd()
OS.chdir("/home/user")
files = OS.listdir(".")
OS.mkdir("new_folder")
home = OS.getenv("HOME")
OS.run("ls", ["-la"], False)
```

### JSON Handling

Carrion provides comprehensive JSON support for parsing and stringifying data:

#### JSON to Carrion Type Conversion
- `null` → `None`
- `true/false` → `Boolean`
- Numbers → `Integer` (whole) or `Float` (decimal)
- Strings → `String`
- Arrays → `Array`
- Objects → `Map` (with string keys)

#### Parsing JSON
```python
// Parse JSON objects with parseHash (built-in)
data = parseHash('{"name": "Alice", "age": 30, "active": true}')
print(data["name"])   // → "Alice"
print(data["age"])    // → 30
print(data["active"]) // → True

// Parse any JSON with httpParseJSON (more flexible)
// Arrays
array_data = httpParseJSON('[1, 2, 3, "hello", true, null]')
print(array_data[3])  // → "hello"
print(array_data[5])  // → None

// Nested structures
complex_data = httpParseJSON('{"users": [{"id": 1, "name": "Bob"}, {"id": 2, "name": "Carol"}]}')
print(complex_data["users"][0]["name"])  // → "Bob"

// Error handling
attempt:
    invalid = parseHash('[1, 2, 3]')  // Error: parseHash only accepts objects
ensnare:
    print("parseHash requires a JSON object, not an array")
```

#### Converting to JSON
```python
// Convert Carrion objects to JSON strings
user = {
    "name": "Alice",
    "age": 30,
    "skills": ["Python", "Carrion", "Go"],
    "active": True,
    "metadata": None
}

json_string = httpStringifyJSON(user)
print(json_string)
// → '{"name":"Alice","age":30,"skills":["Python","Carrion","Go"],"active":true,"metadata":null}'

// Works with arrays too
numbers = [1, 2.5, 3]
json_array = httpStringifyJSON(numbers)
print(json_array)  // → '[1,2.5,3]'
```

#### Practical JSON Examples
```python
// API Integration
response = httpGet("https://api.example.com/users/123")
user_data = httpParseJSON(response)
print(f"User: {user_data['name']}, Email: {user_data['email']}")

// Configuration files
config_json = File.read("config.json")
config = httpParseJSON(config_json)
debug_mode = config.get("debug", False)

// Creating API requests
request_data = {
    "action": "create_user",
    "params": {
        "username": "newuser",
        "email": "user@example.com"
    }
}
json_body = httpStringifyJSON(request_data)
response = httpPost("https://api.example.com/users", json_body)
```

## Language Grammar (Simplified)

```
program         := statement*

statement       := assignment_stmt
                |  expression_stmt
                |  if_stmt
                |  for_stmt
                |  while_stmt
                |  match_stmt
                |  function_def
                |  class_def
                |  import_stmt
                |  error_stmt
                |  return_stmt
                |  skip_stmt
                |  stop_stmt
                |  autoclose_stmt

assignment_stmt := identifier '=' expression
                |  identifier compound_op expression
                |  tuple_literal '=' expression

expression      := logical_or

logical_or      := logical_and ('or' logical_and)*
logical_and     := equality ('and' equality)*
equality        := comparison (('==' | '!=') comparison)*
comparison      := bitwise_or (('<' | '>' | '<=' | '>=') bitwise_or)*
bitwise_or      := bitwise_xor ('|' bitwise_xor)*
bitwise_xor     := bitwise_and ('^' bitwise_and)*
bitwise_and     := shift ('&' shift)*
shift           := addition (('<<' | '>>') addition)*
addition        := multiplication (('+' | '-') multiplication)*
multiplication  := exponentiation (('*' | '/' | '//' | '%') exponentiation)*
exponentiation  := unary ('**' unary)*
unary           := ('not' | '-' | '+' | '~') unary | postfix
postfix         := primary ('++' | '--')*
primary         := identifier | literal | '(' expression ')' | call

call            := primary '(' arguments? ')'
                |  primary '.' identifier
                |  primary '[' expression ']'

function_def    := 'spell' identifier '(' parameters? ')' ':' block
class_def       := 'grim' identifier ('(' identifier ')')? ':' class_body

if_stmt         := 'if' expression ':' block ('otherwise' expression ':' block)* ('else' ':' block)?
for_stmt        := 'for' identifier 'in' expression ':' block ('else' ':' block)?
while_stmt      := 'while' expression ':' block
match_stmt      := 'match' expression ':' case_stmt*

error_stmt      := 'attempt' ':' block ('ensnare' '(' identifier ')' ':' block)* ('ensnare' ':' block)? ('resolve' ':' block)?
autoclose_stmt  := 'autoclose' expression 'as' identifier ':' block

literal         := INTEGER | FLOAT | STRING | BOOLEAN | NONE | array_literal | hash_literal | tuple_literal
```

## Reserved Word Reference

### Control Flow
- `if`, `otherwise`, `else` - Conditional statements
- `for`, `in`, `while` - Loops
- `skip`, `stop` - Loop control
- `match`, `case` - Pattern matching
- `return` - Function return

### Object-Oriented
- `grim` - Class definition
- `spell` - Method/function definition
- `init` - Constructor
- `self` - Instance reference
- `super` - Parent class reference
- `arcane` - Abstract class modifier
- `arcanespell` - Abstract method decorator

### Error Handling
- `attempt` - Try block
- `ensnare` - Catch block
- `resolve` - Finally block
- `raise` - Throw error
- `check` - Assertion

### Logical
- `and`, `or`, `not` - Logical operators
- `True`, `False` - Boolean literals
- `None` - Null value

### Module System
- `import` - Import statement
- `as` - Import alias

### Resource Management
- `autoclose` - Automatic resource cleanup

### Miscellaneous
- `var` - Variable declaration (optional)
- `ignore` - Empty statement placeholder

This reference provides a comprehensive overview of the Carrion programming language syntax, semantics, and standard library.
