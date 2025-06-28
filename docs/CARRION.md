
# The Carrion Programming Language
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
### The Carrion programming language is a dynamically typed, interpreted programming language written in Go. It's designed to be similar to Python with creative syntax modifications and unique functionality. Don't be intimidated - it's meant to be easy to learn and features a fun crow theme! It's named after the Carrion Crow and has Norse mythology inspirations (with the standard library named Munin after Odin's raven).  


## Notes:
Currently the language is in active development (Version 0.1.6-alpha).

## Fun things I'm doing to the language.
2 types of increments are accepted 

* c style prefix and postfix increment: ++i and i++
and
* python style: i += 1

notation. 

It has similar syntax to python but no typing just yet will have optional type hint eventually.

# Comments:

Comments in Carrion use C-style syntax:
- Single-line comments: `//`
- Multi-line comments: `/* */`



"Example syntax":

```python
// This is just a comment
spell foobar(x, y):
    return x + y
```

# Strings
Strings in Carrion follow a similar convention to python
Docstrings, formatted string and single and double quoted strings are support as well.

All of these are valid:
- """This is docstring """
- "Double quoted string"
- 'Single quoted string'
- f"Formatted string"
- f'Formatted single quoted string'
- i"Double String Interpolation"
- i'Single String Interpolation'

* Note: F-strings and interpolated strings currently don't throw errors for unused replacement characters, but proper validation is planned for future versions. String interpolation is generally preferred over concatenation for readability.

### String concatenation
 You can add strings together by using +
E.g.
```python
x = "foo"
y = "bar"

result = x + y
print(result)
```

# Defaults

Spells are Carrion's functions, working similarly to Python methods. Here's how they work:

```python
spell foobar(name="foobar"):
    return name
```

This allows you to set default arguments in the parameters.

# Current Functionality
- Works of a tree walking paradigm
- The carrion language is similar to python but it has some differences i prefer. 
- The interpreter works 
- Some OOP functions
- Munin standard library (limited)
- Programs can be written and parsed/evaluated via the interpreter 
- Working REPL
- OS and File functions

# Running Carrion

## REPL (Interactive Mode)
```bash
carrion
```
Running `carrion` without arguments starts the interactive REPL with command history support.

## Running Scripts
```bash
carrion script.crl
```
Carrion scripts use the `.crl` file extension.

# Data Types

Carrion supports the following data types:

## Primitive Types
- **Integer**: Whole numbers (64-bit)
- **Float**: Decimal numbers (64-bit)
- **String**: Text data with support for single quotes, double quotes, docstrings, and interpolation
- **Boolean**: `True` or `False`
- **None**: Represents absence of value

## Collection Types
- **Array**: Ordered, mutable list of elements `[1, 2, 3]`
- **Hash**: Key-value pairs (dictionary) `{"key": "value"}`
- **Tuple**: Immutable ordered collection `(1, 2, 3)`

# Built-in Functions

## Type Conversion
- `int(value)` - Convert to integer
- `float(value)` - Convert to float
- `str(value)` - Convert to string
- `list(string)` - Convert string to list of characters

## Utility Functions
- `len(object)` - Get length of strings, arrays, or other collections
- `type(object)` - Get the type of an object
- `print(*args)` - Print values to console
- `input(prompt)` - Read user input from terminal
- `range(start, stop[, step])` - Generate a sequence of numbers

## Error Handling
- `Error(message)` - Create a generic error
- `check(condition, message)` - Assert a condition, raise error if false

## OS and File Operations
Carrion provides wrapped Go functions for file and OS operations through the standard library.

# Type Hints
* For fun you can add in type hints for extra clarity i haven't implemented a checker yet but here is the Implementation.

You can set a type hint for variables and parameters. 

```python
x: str = "Foo"

grim Foo:
    init(x:str = x):
        self.x = x

```
As you can see you can set them no issue just doesn't mean much yet until i implement a checker and a vm and jit compiler.


# Loops
* Currently for and while loops are supported 
- For Loops work like python for loops
```python
for x in range(10):
    print(x)
```
- While loops work like python while loops
```python
x = 10
while x < 20:
    print(x)
    x++
```
## Skip/Stop

* For conditions inside a loop, you might want to skip over something or stop execution based on a rule.
This is where the keywords `skip` and `stop` come in:
- `skip`: Skips to the next iteration (equivalent to Python's `continue`)
- `stop`: Stops the loop execution (equivalent to Python's `break`)

e.g.
```python
x = 0
while x < 10:
    if x % 2 == 0:
        skip
    else:
        print(x)
    x++
// should return all odds
i = 0
while i < 10:
    if x == 3:
        stop
    else:
        print(i)
// Should output 0,1,2

```


# Match/Case
Match case works similar to python you declare a match and a case for and use an underscore as a default.
e.g.
```python 
foo = "foobar"

match foo:
    case "foo":
        print("foo")
    case "bar":
        print("bar")
    _:
        print("foobar")

```

*Note: List comprehensions are not currently supported but are planned for future versions.

# Classes and Imports
* Currently Classes are implemented as well as Imports
- Classes in Carrion are conceptually called "Spellbooks" but use the keyword `grim` (short for grimoire)
- Methods within classes are called "spells"

# Example of how classes (grimoires/spellbooks) are implemented
```python

grim Foobar:
    init(foo):
        self.foo = foo
    spell foobar_method():
        return self.foo

foo  = Foobar("foobar")
print(foo.foobar_method())
```

# Imports also work but are based on file name

```python
import "example/example"
example = Example("example value")

example.example_method()
```

* Note: Once you import a file you have access to it's methods by calling in the class name.

# OOP with Grimoires (Spellbooks)
Grimoires (conceptually called Spellbooks) are Carrion's classes. Use the `grim` keyword to define them.
Not all OOP aspects are implemented but some are.
You can declare self from any method and set them in the init.

The construct method can be used as spell init() or init() both work just a preference.
e.g.

# With init: init() constructor:

```python
grim Foobar:
    init(foo, bar):
        self.foo = foo
        self.bar = bar
    spell print_foo():
        return self.foo
    spell print_bar():
        return self.bar

foo = Foobar("foo","bar")
foo.print_foo()
foo.print_bar()
```

# With spell init: spell init() constructor:

```python
grim Foobar:
    spell init(foo, bar):
        self.foo = foo
        self.bar = bar
    spell print_foo():
        return self.foo
    spell print_bar():
        return self.bar

foo = Foobar("foo","bar")
foo.print_foo()
foo.print_bar()
```

# OOP- Object Oriented Programming
Finally i know you're wondering is this functional or object oriented. Big reveal it's object oriented no surprise.
So inspired by python it's no surprise.

## Abstraction
In Carrion Language there is abstract classes but they are labeled by the keyword arcane.

To declare an arcane grimoire (abstract class) you declare it as follows:

```python
arcane grim Foo:
    init();
        ignore

```

The arcane keyword is used to declare abstract classes.

For abstracted method they are labeled as arcane spells but uses the @ symbol and a decorator format as follows:


```python
arcane grim Foo:
    @arcanespell
    spell foobar():
        ignore
```

You're probably wondering about the ignore. Yes ignore is the keyword that indicates a empty method and that it doesn't need a body statement. You're welcome


## Inheritance
Inheritance is pretty similar to python. Currently the super method isn't implemented but you can share resources with self and inheriting from parent class.

```python
grim Parent:
    spell parent_method():
        ignore

grim Child(Parent):
    init():
        ignore

child = Child()
child.parent_method()
```

Pretty simple no?

## Encapsulation
For Encapsulation i took inspiration from python again i love me some dunder method.
Protected are declare by '__' double underscore while private are declared by '_' a singular underscore:

e.g Protected:
```python
grim Foo:
    init(var="foobar"):
        self.var = var

    spell __protected_spell():
        return self.var

    spell return_protected():
        return str(self.__protected_spell())

foobar = Foo()
foobar.return_protected()

```
* Accessing the protected spell outside of class will cause error.


e.g Private:
```python
grim Foo:
    init(var="foobar"):
        self.var = var

    spell _private_spell():
        return self.var

    spell return_private():
        return str(self._private_spell())

foobar = Foo()
foobar.return_private()

```
* Accessing the private spell outside of class will cause error.

## Polymorphism

Last OOP concept.
You can overwrite spells from inherited parents.

```python
grim Parent:
    init(name="parent"):
        self.name = name
    spell parent_method():
        return self.name


grim Child(Parent):
    init(name="child"):
        self.name = name
    spell parent_method(child_age="age"):
        return "Name: " + str(self.name) + "/t Age: " + child_age
```



# Error Handling
Yeah but it's only partially implemented:
You have 3 keywords for errors **attempt**, **ensnare**, **resolve**

Attempt is the original case want to work. If it doesn't you can ensnare the error.
This allows you to catch a raised error and handle it accordingly.
Finally resolve just finishes the functionality that if specified will always run after error.

```python

// Custom error
 // Define a custom error
grim ValueError:
    spell init(message):
        self.message = message

// Raise a custom error
attempt:
    raise ValueError("Invalid value")
ensnare (ValueError):
    print("Caught ValueError!")
ensnare:
    print("Base case generic error")
resolve:
    print("This will finish method")
```

Raise is the keyword to throw an error because i love you and its easy.


now There is some checks you can do here for error handling. You have literally a check statement. This statement will  check if a condition is true and return an error otherwise
```python
x = 10
// Should pass
check (x == 10, f"x should equal 10 got: {x}")

// Should fail and raise assertion check error
check (x == 12, f"x should equal 12 got: {x}")
```
See i really do love you makes sense right?



# Running Example Files
```bash
carrion examples/test_file.crl
```

The `examples/` directory contains numerous test files demonstrating various language features.
# Standard Library - Munin

The Munin standard library (named after Odin's raven) provides essential functionality for Carrion programs.

## Available Modules

### Core Modules
- **Array**: Enhanced array manipulation with the `Array` grimoire
- **String**: String operations and utilities
- **Integer**: Integer-specific operations
- **Float**: Float-specific operations
- **Boolean**: Boolean logic operations

### System Modules
- **OS**: Operating system interface
- **File**: File I/O operations
- **Math**: Mathematical functions and constants

### Development Tools
- **Debug**: Debugging utilities
- **Primitive**: Basic type operations

## Using the Standard Library

The standard library is automatically loaded when Carrion starts. You can check the loaded modules with:

```python
modules()  // List all available modules
version()  // Check Carrion and Munin versions
help()     // Get help information
```

## Example: Using the Array Grimoire

```python
// Create an enhanced array
arr = Array([1, 2, 3])
arr.append(4)
print(arr.length())  // Output: 4
print(arr.get(0))    // Output: 1
```

# Language Reference

## Keywords

### Control Flow
- `if`, `else`, `otherwise` - Conditional statements
- `for`, `while` - Loop constructs
- `skip` - Continue to next iteration
- `stop` - Break from loop
- `return` - Return from function
- `match`, `case` - Pattern matching

### Object-Oriented Programming
- `grim` - Class definition keyword (short for grimoire/spellbook, can be preceded by `arcane` for abstract classes)
- `spell` - Method definition
- `init` - Constructor method
- `self` - Reference to current instance
- `super` - Reference to parent class

### Error Handling
- `attempt` - Try block
- `ensnare` - Catch block
- `resolve` - Finally block
- `raise` - Throw an error
- `check` - Assert a condition

### Special Keywords
- `import` - Import modules
- `as` - Alias in imports
- `in` - Membership test
- `not in` - Negative membership test
- `and`, `or`, `not` - Boolean operators
- `True`, `False` - Boolean literals
- `None` - Null value
- `ignore` - Empty statement placeholder

### Decorators
- `@arcanespell` - Mark abstract methods

## Operators

### Arithmetic
- `+` - Addition
- `-` - Subtraction
- `*` - Multiplication
- `/` - Division
- `%` - Modulo
- `**` - Exponentiation

### Assignment
- `=` - Basic assignment
- `+=` - Addition assignment
- `-=` - Subtraction assignment
- `*=` - Multiplication assignment
- `/=` - Division assignment
- `++` - Increment (prefix/postfix)
- `--` - Decrement (prefix/postfix)

### Comparison
- `==` - Equal
- `!=` - Not equal
- `<` - Less than
- `>` - Greater than
- `<=` - Less than or equal
- `>=` - Greater than or equal

### Bitwise
- `&` - Bitwise AND
- `|` - Bitwise OR
- `^` - Bitwise XOR
- `~` - Bitwise NOT
- `<<` - Left shift
- `>>` - Right shift

## Type Hints (Optional)

Carrion supports optional type hints for clarity:

```python
x: int = 42
name: str = "Carrion"

spell calculate(a: int, b: int) -> int:
    return a + b
```

## Best Practices

1. **Naming Conventions**
   - Use snake_case for variables and functions
   - Use PascalCase for grimoires/classes (defined with `grim`)
   - Private methods: prefix with `_`
   - Protected methods: prefix with `__`

2. **Error Handling**
   - Always use attempt/ensnare for operations that might fail
   - Create custom error classes for specific error types
   - Use `check` for precondition validation

3. **Code Organization**
   - One grimoire/class per file for large classes
   - Group related functions together
   - Use the standard library modules when available

## Limitations and Future Features

### Current Limitations
- No list comprehensions (planned)
- No generator expressions
- Limited metaclass support
- No async/await syntax

### Planned Features
- JIT compilation for performance
- Virtual machine implementation
- Enhanced type system with static checking
- List comprehensions
- More comprehensive standard library
- Better interoperability with other languages

