
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
### The Carrion programming language is a dynamically typed programming language written in Go. It's supposed to be similar to python with some syntax changes and functionality that i like. Don't beintimidated. It's meant to be easy to learn it just has a fun crow theme! It's named after the Carrion Crow. And has some norse inspirations.  


## Notes:
Currently the language is in development.

## Fun things I'm doing to the language.
2 types of increments are accepted 

* c style prefix and postfix increment: ++i and i++
and
* python style: i += 1

notation. 

It has similar syntax to python but no typing just yet will have optional type hint eventually.

# Comments:

Comments in carrion are either // or /* */ i just like using those so idc if you like # functions i like those and it's my language.



"Example syntax":

```python
// This is just a comment
spell foobar(x, y):
    return x + y
```

# Defaults

Spells work just like methods in python if you're familar with python if not here's an example

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

# Run REPL
```bash
carrion
```
- Note: Run carrion without a file to run REPL

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

# With init: init() constructor:

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

# With spell init: spell init() constructor:

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

# OOP- Object Oriented Programming
Finally i know you're wondering is this functional or object oriented. Big reveal it's object oriented no surprise.
So inspired by python it's no surprise.

## Abstraction
In Carrion Language there is abstract classes but they are labeled by the keyword arcane.

To declare an arcane spellbook you declare it as follows:

```python
arcane spellbook Foo:
    init();
        ignore

```

The arcane keyword is used to declare abstract classes.

For abstracted method they are labeled as arcane spells but uses the @ symbol and a decorator format as follows:


```python
arcane spellbook Foo:
    @arcanespell
    spell foobar():
        ignore
```

You're probably wondering about the ignore. Yes ignore is the keyword that indicates a empty method and that it doesn't need a body statement. You're welcome


## Inheritance
Inheritance is pretty similar to python. Currently the super method isn't implemented but you can share resources with self and inheriting from parent class.

```python
spellbook Parent:
    spell parent_method():
        ignore

spellbook Child(Parent):
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
spellbook Foo:
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
spellbook Foo:
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
spellbook Parent:
    init(name="parent"):
        self.name = name
    spell parent_method():
        return self.name


spellbook Child(Parent):
    init(name="child"):
        self.name = name
    spell parent_method(child_age="age"):
        return "Name: " + str(self.name) + "/t Age: " + child_age
```

Yay i'm a genius.

Anyways


# Error Handling
Yeah but it's only partially implemented:
You have 3 keywords for errors **attempt**, **ensnare**, **resolve**

Attempt is the original case want to work. If it doesn't you can ensnare the error.
This allows you to catch a raised error and handle it accordingly.
Finally resolve just finishes the functionality that if specified will always run after error.

```python

// Custom error
 // Define a custom error
spellbook ValueError:
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



# Example file run.
```bash
carrion examples/test_file.crl
```
# Standard Library - Munin

## Current Implementation
- OS and File Functionality
- Basic Math Functions

