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



* Methods are defined as spells


"Example syntax":
```python
spell foobar(x, y):
    return x + y
```

# Current Functionality
-  The carrion language is similar to python but it has some differences i prefer. 
- The interpreter works but OOP features haven't been implemented yet
- No standard library yet but will be implemented with OOP
- Programs can be written and parsed/evaluated via the interpreter 
- Working Interactive REPL

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

# Loops
* Currently for and while loops are supported 
- For Loops work like python for loops 
- While loops work like python while loops

*Notes: Currently no support for list comprehensions like in python

File type:
- .crl

# Classes and Imports
* Currently Classes are implemented as well as Imports
- Classes in Carrion are known as "Spellbooks" and the Methods are spells within those Spellbooks

# Example of how spellbooks are implemented
```python

spellbook Example:
    init(example):
        self.example = example
    spell example_method():
        return self.example

example  = Example("example value")
print(example.example_method())
```

# Imports also work but are based on file name

```python
import "example/example"

example = Example("example value")

example.example_method()
```

* Notes: The init method is used as a contructor however Currently there is a bug with setting multiple self params i'm trying to fix it

# Built in Go. 
* Needs go to build interpreter

# Example file run.
```bash
./thecarrionlanguage examples/test_file.crl
```

# Future Updates
- Possible list comprehensions
- Fix to self params
- More OOP integration
- Standard Munin library
- Build and alias the carrion language
- File I/O
- Build setup


# Author
- Javan Hutchinson
