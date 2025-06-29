# 👋 Hello World & Basics - Your First Magical Spells

Welcome to your first lesson in Carrion magic! In this guide, you'll learn to cast your first spells and understand the fundamental concepts that make Carrion special.

---

## 🪄 Your Very First Spell

Let's start with the classic "Hello, World!" - a time-honored tradition in the programming realm!

### Method 1: Interactive REPL Magic
```bash
# Start the magical REPL
carrion
```

Then cast your first spell:
```python
>>> print("Hello, Magical World! 🪄")
Hello, Magical World! 🪄
```

**🎉 Congratulations!** You've just cast your first Carrion spell!

### Method 2: Script File Magic
Create a file called `hello_magic.crl`:
```python
// My first Carrion spell!
print("Hello, Magical World! 🪄")
print("Welcome to the realm of Carrion!")
```

Execute it:
```bash
carrion hello_magic.crl
```

---

## 🗣️ The Language of Magic - Comments

In Carrion, we document our spells with comments:

```python
// This is a single-line comment
// Comments help explain your magical intentions

/* 
   This is a multi-line comment
   Perfect for longer explanations
   of complex magical procedures
*/

```
This is also a multi-line comment
using triple backticks
```

print("Comments don't affect the magic!") // Inline comment
```

---

## 📦 Magical Storage - Variables

Variables in Carrion store magical energy (data) for later use:

```python
// Storing different types of magical energy
magical_number = 42
wizard_name = "Merlin"
has_powers = True
spell_count = None

// Display our magical storage
print("Magical Number:", magical_number)
print("Wizard Name:", wizard_name)
print("Has Powers:", has_powers)
print("Spell Count:", spell_count)
```

### ✨ Dynamic Magic
Carrion is dynamically typed - variables can hold different types of magic:

```python
magic_variable = 42              // It's a number!
print(magic_variable)

magic_variable = "Abracadabra!"  // Now it's text!
print(magic_variable)

magic_variable = True            // Now it's a truth value!
print(magic_variable)
```

---

## 🎭 Types of Magical Energy - Data Types

Carrion recognizes several types of magical energy:

### 🔢 Numbers (Integer & Float)
```python
// Whole numbers (Integers)
age = 25
spell_power = 1000
negative_energy = -50

// Decimal numbers (Floats)
magic_ratio = 3.14159
temperature = -15.5
precision = 0.001

print("Age:", age)
print("Magic Ratio:", magic_ratio)
```

### 🔤 Text Magic (Strings)
```python
// Different ways to create text magic
single_quotes = 'Hello, mage!'
double_quotes = "Welcome to Carrion!"
multiline_spell = """
This is a longer incantation
that spans multiple lines
with great power!
"""

// String magic with variables (F-strings)
name = "Luna"
level = 15
greeting = f"Greetings, {name}! You are level {level}."
print(greeting)

// Alternative string interpolation
alt_greeting = i"Hello, ${name}! Your power level is ${level}."
print(alt_greeting)
```

### ✅ Truth Magic (Booleans)
```python
is_magical = True
is_mundane = False
has_completed_quest = True

if is_magical:
    print("You possess magical abilities!")

if not is_mundane:
    print("Magic flows through you!")
```

### 🫥 The Void (None)
```python
unknown_spell = None
empty_potion = None

if unknown_spell is None:
    print("This spell hasn't been learned yet!")
```

---

## 📚 Collections of Magic

### 📋 Arrays (Lists)
Arrays store multiple magical items in order:

```python
// Creating magical collections
spell_names = ["Fireball", "Healing", "Teleport", "Shield"]
magical_numbers = [1, 3, 5, 7, 11, 13]
mixed_magic = [42, "Wand", True, 3.14]

// Accessing magical items (0-based indexing)
print("First spell:", spell_names[0])      // "Fireball"
print("Last spell:", spell_names[-1])      // "Shield"

// Array magic
print("Number of spells:", len(spell_names))
print("All spells:", spell_names)
```

### 🗂️ Hashes (Dictionaries)
Hashes map magical keys to values:

```python
// Creating magical mappings
wizard_stats = {
    "name": "Gandalf",
    "level": 50,
    "class": "White Wizard",
    "health": 100
}

// Accessing magical properties
print("Wizard Name:", wizard_stats["name"])
print("Level:", wizard_stats["level"])

// Magical inventory
spell_inventory = {
    "offensive": ["Fireball", "Lightning Bolt"],
    "defensive": ["Shield", "Barrier"],
    "healing": ["Cure", "Regeneration"]
}
```

### 🎯 Tuples (Immutable Collections)
Tuples are magical collections that cannot be changed:

```python
// Magical coordinates that shouldn't change
dragon_location = (100, 250, 75)  // x, y, z coordinates
spell_components = ("Eye of Newt", "Dragon Scale", "Moonstone")

// Unpacking magical values
x, y, z = dragon_location
print(f"Dragon is at position ({x}, {y}, {z})")

first_component, second_component, third_component = spell_components
print(f"First ingredient: {first_component}")
```

---

## 🎯 Basic Magical Operations

### 🧮 Arithmetic Magic
```python
// Basic magical calculations
power_level = 100
experience = 250
bonus = 50

total_power = power_level + bonus
remaining_exp = experience - 100
double_power = power_level * 2
half_power = power_level / 2

print(f"Total Power: {total_power}")
print(f"Double Power: {double_power}")

// Advanced magical operations
squared_power = power_level ** 2        // Exponentiation
magic_remainder = experience % 7        // Modulo
floor_division = experience // 3        // Integer division

print(f"Squared Power: {squared_power}")
print(f"Magic Remainder: {magic_remainder}")
```

### 🔤 Text Magic Operations
```python
// Combining magical texts
greeting = "Hello"
target = "Fellow Mage"
full_greeting = greeting + ", " + target + "!"
print(full_greeting)

// Magical text repetition
magical_chant = "Abracadabra! " * 3
print(magical_chant)

// Text magic properties
spell_name = "Teleportation"
print(f"Spell length: {len(spell_name)}")
print(f"Uppercase: {spell_name.upper()}")
print(f"Lowercase: {spell_name.lower()}")
```

---

## 🎪 Interactive Magic - Input & Output

### 📤 Displaying Magic (Output)
```python
// Basic magical output
print("Welcome to the magic academy!")

// Displaying multiple magical values
name = "Merlin"
age = 1000
print("Wizard:", name, "Age:", age)

// Formatted magical output
print(f"The great {name} is {age} years old!")

// Different ways to display
print("First line")
print("Second line")
print("Same line:", end=" ")
print("continued here!")
```

### 📥 Gathering Magic (Input)
```python
// Asking for magical input
wizard_name = input("What is your wizard name? ")
print(f"Welcome, {wizard_name}!")

// Converting magical input
age_text = input("How old are you? ")
age_number = int(age_text)
print(f"You are {age_number} years old!")

// Or do it in one magical step
spell_power = int(input("Enter your spell power level: "))
print(f"Your magic power is {spell_power}!")
```

---

## 🎮 Your First Interactive Spell

Let's create a simple interactive magic program:

```python
// Create a file called: magical_greeting.crl

print("🪄 Welcome to the Magical Greeting Spell! 🪄")
print()

// Gather magical information
wizard_name = input("Enter your wizard name: ")
favorite_spell = input("What's your favorite spell? ")
power_level = int(input("What's your power level (1-100)? "))

// Process the magical data
if power_level >= 80:
    rank = "Archmage"
    symbol = "🌟"
elif power_level >= 50:
    rank = "Master Wizard"
    symbol = "✨"
elif power_level >= 20:
    rank = "Apprentice"
    symbol = "🪄"
else:
    rank = "Novice"
    symbol = "📚"

// Display magical results
print()
print("=" * 40)
print(f"{symbol} Magical Profile {symbol}")
print("=" * 40)
print(f"Name: {wizard_name}")
print(f"Rank: {rank}")
print(f"Favorite Spell: {favorite_spell}")
print(f"Power Level: {power_level}/100")
print()
print(f"May your {favorite_spell} spell serve you well, {rank} {wizard_name}!")
print("=" * 40)
```

Run it with:
```bash
carrion magical_greeting.crl
```

---

## 🔍 Exploring Your Magic - Type Checking

Carrion provides magical functions to examine your data:

```python
// Magical inspection spells
mysterious_value = 42

print("Value:", mysterious_value)
print("Type:", type(mysterious_value))        // "INTEGER"
print("Length:", len("Hello"))                // 5

// Check if values are the same type
print(is_sametype(42, 17))                   // True
print(is_sametype(42, 3.14))                 // False
print(is_sametype("hello", "world"))         // True
```

---

## 🎯 Practice Challenges

Try these challenges to master the basics:

### 🏅 Challenge 1: Personal Magic Calculator
Create a program that:
1. Asks for your name and age
2. Calculates how many days you've been alive
3. Displays a magical message with the result

### 🏅 Challenge 2: Spell Component Mixer
Create a program that:
1. Asks for three spell components
2. Combines them into a magical spell name
3. Calculates the "power level" based on the total length

### 🏅 Challenge 3: Magic Number Guesser
Create a simple number guessing game using basic input/output.

---

## 🔗 What's Next?

Congratulations! You've mastered the basics of Carrion magic. Here's your next quest:

1. **⚡ [Quick Start Tutorial](Quick-Start.md)** - Learn more advanced basics
2. **🎮 [REPL Guide](REPL-Guide.md)** - Master the interactive environment
3. **📚 [Data Types Deep Dive](../Language-Fundamentals/Data-Types.md)** - Understand magical data types
4. **✨ [Functions (Spells)](../Language-Fundamentals/Functions.md)** - Create reusable magic

---

## 💡 Tips for New Mages

1. **🎪 Use the REPL**: It's perfect for experimenting with small spells
2. **📖 Read Error Messages**: They often contain helpful guidance
3. **🔍 Use `type()`**: When unsure about data types
4. **💬 Use Comments**: Document your magical intentions
5. **🎯 Start Small**: Begin with simple spells and build complexity

---

*Remember: Every great wizard started with "Hello, World!" Your magical journey has just begun! 🪄✨*

> "The first spell is always the hardest to cast, but also the most magical to remember." - *Ancient Programming Wisdom*