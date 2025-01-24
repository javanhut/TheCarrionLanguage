# The Carrion Programming Language ver 0.0.1
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
### The Carrion programming language is a python like programming language written in Go. It is meant to be a fun project. it's interpreted and dynamically typed 


## Notes:
Currently the language is in development.

## Fun things I'm doing to the language.
2 types of increments are accepted 

* c style prefix and postfix increment: ++i and i++
and
* python style: i += 1 (Currently a bug i need to fix in parser)

notation. 

It has similar syntax to python but no typing just yet will have optional type hint eventually.

# Comments:

Comments in carrion are either // or /* */ i just like using those so idc if you like # functions i like those and it's my language.



* Methods are defined as spells


"Example syntax":
```python
// This is just a comment
spell foobar(x, y):
    return x + y
```

# Current Functionality
- Works of a tree walking paradigm
- The carrion language is similar to python but it has some differences i prefer. 
- The interpreter works 
- Some OOP functions
- Munin standard library (limited)
- Programs can be written and parsed/evaluated via the interpreter 
- Working REPL
- OS and File functions

# Run REPL
```bash
go build .
./thecarrionlanguage
```
- Note: Run thecarrionlanguage without a file to run REPL

# Data Types Currently supported:
 - Arrays
 - Hashmap
 - Integers
 - Float
 - Strings
 - Tuples

# Builtin Methods

- len() - Gets the length of the object input

- print() - prints out the input of the content

- int() - converts to integer

- float() - converts int to float

- str() - convert to string 

- type() - get the data type of input object

- list() - converts string to list of runes

- input() - takes user input from terminal

- range() - makes a range function from any numbers 

- os and file functions from golang but wrapped in Carrion Lang.


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

*Notes: Currently no support for list comprehensions like in python

# File type:
- .crl

# Classes and Imports
* Currently Classes are implemented as well as Imports
- Classes in Carrion are known as "Spellbooks" and the Methods are spells within those Spellbooks

# Example of how spellbooks are implemented
```python

spellbook Foobar:
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

# OOP part of Spellbooks
Spellbooks are Carrion's classes.
Not all OOP aspects are implemented but some are.
You can declare self from any method and set them in the init.

The construct method can be used as spell init() or init() both work just a preference.
e.g.
init() constructor:

```python
spellbook Foobar:
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

spell init() constructor:

```python
spellbook Foobar:
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


# Built in Go. 
* Needs go to build interpreter

# Example file run.
```bash
./thecarrionlanguage examples/test_file.crl
```
# Standard Library - Munin

## Current Implementation
- OS and File Functionality
- Basic Math Functions


# Future Updates
- Possible list comprehensions
- Fix to self params
- More OOP integration
- Build and alias the carrion language
- Build setup
- JIT Compiler and VM
- Error Handling
- Generic Functions and Abstracts
- String formatter and Defaults
- Type hint

# Author
- Javan Hutchinson


# For issues:
Shove it file a issue i may or may not look at it haha jk i probably will.
