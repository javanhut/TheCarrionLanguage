# Carrion Language Documentation
## Data Types & Structures

> **Hey there!** Buckle up for a quick tour through the types and data structures Carrion currently supports. We’ll peek at examples using Python’s code fencing just so it’s nicely highlighted. Remember though, these examples are for Carrion!

---

### 1. Integer
Integers represent whole numbers (positive, negative, and zero).  
Examples: `-1`, `0`, `42`.

```python
# Assigning an integer
my_int = 42

# Checking integer type (type checks are mostly for fun right now in Carrion)
my_int: int = 100
```

**Fun fact**: In Carrion, you don’t have to declare types, but you can if you want to help future you figure out what you were trying to do in your code.  

---

### 2. Float
Floats represent decimal numbers (like 3.14 for pi lovers).  
Examples: `-2.5`, `0.0`, `99.99`.

```python
# Assigning a float
my_float = 3.14

# Type hint for future performance improvements
my_float: float = -2.71828
```

**Remember**: Carrion automatically distinguishes between integers and floats when it sees that decimal point.

---

### 3. Strings
Strings can be defined using **single** or **double quotes** even **triple quotes** —whichever you prefer. No biggie.

```python
my_string = "Hello, world!"
another_string = 'I can use single quotes too!'
doctring = """This is doctring supported"""
```

Also just in case you were wondering yes fstrings are supported out of the box.
If you're coming from python no .format method required!

```python
fstr = 'some formatted strings'
f_string = f"I love me {fstr}"
```

You can use both single and double quoted f-strings!



Just like in other languages, you can concatenate strings or access their characters by index. But we’ll leave those details for the “Operators” and “String Manipulation” sections.

---

### 4. Arrays
Arrays in Carrion (akin to lists in Python) are dynamic, meaning you can stuff them with different types of elements. That’s right—mix and match if you like:

```python
my_array = [1, 2, 3]
mixed_array = [1, "two", 3.0]

# Arrays can grow dynamically
my_array.append(4)
print(my_array)  # [1, 2, 3, 4]
```

You can slice, dice, and iterate over arrays. Think of them as your friendly, flexible containers.

---

### 5. Hashmaps
Hashmaps (like Python dictionaries) store **key-value** pairs. Accessing data is a breeze thanks to hashing under the hood.

```python
my_hashmap = {
    "name": "Carrion",
    "version": 0.1,
    "fun": True
}

print(my_hashmap["name"])  # "Carrion"
```

**Tip**: Keep your keys unique or you’ll overwrite your data. (Unless your data is secretly meant to be overwritten—hey, no judgment here!)

---

### 6. Tuples
Tuples in Carrion are like tuples in Python: **immutable** sequences of items. Once defined, the contents cannot be changed (that’s the difference from arrays).

```python
my_tuple = (1, "hello", 3.14)
# Trying to alter a tuple’s contents will make Carrion give you the side eye.
```

Tuples are great when you have data that shouldn’t change, like config settings or min/max boundaries.

---

## Summary & Next Steps

That’s all for now on the data types and structures in Carrion. As the language evolves, you can expect more data structures and refined features. In the meantime, feel free to experiment with arrays, hashmaps, tuples, and all sorts of combinations!

Looking for more info? Check out our [Basic Syntax Documentation](../Syntax/README.md) for operators, variable assignment, and other language fundamentals.  

**Happy coding with Carrion!**
