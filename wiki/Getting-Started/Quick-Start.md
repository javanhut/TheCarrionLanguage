# ⚡ Quick Start Tutorial - Learn Carrion in 30 Minutes

Ready to master the basics of Carrion magic quickly? This tutorial will have you casting spells and creating grimoires in just 30 minutes! Perfect for experienced programmers who want to learn Carrion's unique magical approach.

---

## 🎯 What You'll Learn

- ✨ **Basic Syntax** - Variables, operators, and data types
- 🪄 **Functions (Spells)** - Creating reusable magical code
- 📖 **Classes (Grimoires)** - Object-oriented magic
- 🔍 **Standard Library** - Using Munin's powerful grimoires
- 🎮 **Interactive Features** - REPL and error handling

---

## 🚀 Let's Start Coding!

### Step 1: Your First Magic (5 minutes)

Open your Carrion REPL:
```bash
carrion
```

Try these basic spells:

```python
// Variables and basic operations
name = "Aspiring Mage"
power_level = 42
is_magical = True

print(f"Hello, {name}! Your power level is {power_level}")

// Dynamic typing - variables can change type
magic_value = 100        // Integer
magic_value = "Powerful" // Now a string
magic_value = True       // Now a boolean

// Collections
spell_list = ["Fireball", "Heal", "Teleport"]
wizard_stats = {"health": 100, "mana": 50, "level": 5}

print(f"Known spells: {spell_list}")
print(f"Current health: {wizard_stats['health']}")
```

**Key Takeaway**: Carrion is dynamically typed with familiar Python-like syntax, but uses magical terminology!

### Step 2: Spell Creation (Functions) (8 minutes)

Functions in Carrion are called "spells":

```python
// Basic spell
spell greet(wizard_name):
    return f"Greetings, {wizard_name}! Welcome to the magical realm!"

// Spell with default parameters
spell cast_damage_spell(base_damage, multiplier = 1.5):
    total_damage = base_damage * multiplier
    return f"💥 Dealing {total_damage} magical damage!"

// Advanced spell with multiple returns
spell analyze_creature(creature_type):
    if creature_type == "dragon":
        return "fire", 1000, True  // element, power, dangerous
    elif creature_type == "unicorn":
        return "light", 500, False
    else:
        return "unknown", 100, False

// Test your spells
print(greet("Merlin"))
print(cast_damage_spell(50))
print(cast_damage_spell(50, 2.0))  // Custom multiplier

element, power, dangerous = analyze_creature("dragon")
print(f"Dragon: {element} element, {power} power, dangerous: {dangerous}")
```

**Key Takeaway**: Spells (functions) use the `spell` keyword and support default parameters, multiple returns, and unpacking!

### Step 3: Grimoire Crafting (Classes) (10 minutes)

Classes in Carrion are called "grimoires" (magical spellbooks):

```python
// Basic grimoire
grim MagicalCreature:
    init(name, element, health = 100):
        self.name = name
        self.element = element
        self.health = health
        self.max_health = health
    
    spell introduce():
        return f"I am {self.name}, a creature of {self.element}!"
    
    spell take_damage(damage):
        self.health = max(0, self.health - damage)
        if self.health == 0:
            return f"💀 {self.name} has been defeated!"
        else:
            return f"💔 {self.name} takes {damage} damage! Health: {self.health}/{self.max_health}"
    
    spell heal(amount):
        old_health = self.health
        self.health = min(self.max_health, self.health + amount)
        healed = self.health - old_health
        return f"💚 {self.name} heals for {healed} HP! Health: {self.health}/{self.max_health}"

// Inheritance - specialized grimoires
grim Dragon(MagicalCreature):
    init(name, color):
        super.init(name, "fire", 200)  // Dragons have more health
        self.color = color
        self.breath_attack_power = 75
    
    spell breathe_fire():
        return f"🔥 {self.color} dragon {self.name} breathes fire for {self.breath_attack_power} damage!"
    
    spell intimidate():
        return f"😱 {self.name} roars menacingly! All enemies are frightened!"

// Create and use magical creatures
unicorn = MagicalCreature("Stardust", "light", 80)
dragon = Dragon("Smaug", "Red")

print(unicorn.introduce())
print(dragon.introduce())
print(dragon.breathe_fire())
print(unicorn.take_damage(30))
print(unicorn.heal(20))
```

**Key Takeaway**: Grimoires use `grim` keyword, methods are "spells", and inheritance works like Python with `super`!

### Step 4: Standard Library Magic (5 minutes)

Carrion comes with Munin, a powerful standard library:

```python
// Array grimoire magic
magical_items = ["Wand", "Staff", "Orb", "Crystal"]
enhanced_array = Array(magical_items)

enhanced_array.append("Tome")
print(f"Items count: {enhanced_array.length()}")
print(f"Contains Wand: {enhanced_array.contains('Wand')}")
print(f"Sorted items: {enhanced_array.sort().to_string()}")

// String grimoire magic
spell_name = String("FIREBALL SUPREME")
print(f"Lowercase: {spell_name.lower()}")
print(f"Length: {spell_name.length()}")
print(f"Contains 'FIRE': {spell_name.contains('FIRE')}")
print(f"Reversed: {spell_name.reverse()}")

// Number grimoire magic (automatic enhancement)
magic_number = 42
print(f"Binary: {magic_number.to_bin()}")   // Automatic Integer grimoire!
print(f"Is prime: {magic_number.is_prime()}")

pi_approximation = 3.14159
print(f"Rounded: {pi_approximation.round(2)}")  // Automatic Float grimoire!
```

**Key Takeaway**: Munin provides enhanced grimoires for arrays, strings, numbers, and more - with automatic primitive wrapping!

### Step 5: Control Flow & Error Handling (2 minutes)

```python
// Control flow with magical keywords
for i in range(5):
    if i % 2 == 0:
        print(f"Even magic: {i}")
    else:
        skip  // "skip" instead of "continue"

// Enhanced conditionals
magic_level = 75
if magic_level >= 90:
    rank = "Archmage"
otherwise magic_level >= 50:  // "otherwise" instead of "elif"
    rank = "Master"
else:
    rank = "Apprentice"

print(f"Your rank: {rank}")

// Error handling with magical syntax
attempt:
    risky_spell = 10 / 0
ensnare:  // "ensnare" instead of "except"
    print("🚫 Magical error prevented!")
resolve:  // "resolve" instead of "finally"
    print("✨ Spell casting complete")
```

---

## 🎪 Build Something Cool - Magical Battle System

Let's combine everything into a fun project:

```python
grim BattleMage:
    init(name, element):
        self.name = name
        self.element = element
        self.health = 100
        self.mana = 50
        self.spells = Array(["Magic Missile"])
    
    spell learn_spell(spell_name):
        if not self.spells.contains(spell_name):
            self.spells.append(spell_name)
            return f"📚 {self.name} learned {spell_name}!"
        return f"🤔 {self.name} already knows {spell_name}"
    
    spell cast_spell(spell_name, target):
        if not self.spells.contains(spell_name):
            return f"❌ {self.name} doesn't know {spell_name}!"
        
        if self.mana < 10:
            return f"💫 {self.name} is out of mana!"
        
        self.mana -= 10
        damage = 25
        
        if spell_name == "Fireball" and self.element == "fire":
            damage *= 2  // Elemental bonus
        
        target.health = max(0, target.health - damage)
        return f"✨ {self.name} casts {spell_name} on {target.name} for {damage} damage!"
    
    spell status():
        spell_list = ", ".join([self.spells.get(i) for i in range(self.spells.length())])
        return f"{self.name} ({self.element}): HP {self.health}/100, Mana {self.mana}/50, Spells: {spell_list}"

// Create magical battle
fire_mage = BattleMage("Ignis", "fire")
ice_mage = BattleMage("Frost", "ice")

// Learning phase
print(fire_mage.learn_spell("Fireball"))
print(ice_mage.learn_spell("Ice Shard"))

// Battle phase
print("\n🔥 MAGICAL BATTLE BEGINS! ❄️")
print(fire_mage.status())
print(ice_mage.status())

print("\n⚔️ Round 1:")
print(fire_mage.cast_spell("Fireball", ice_mage))
print(ice_mage.cast_spell("Ice Shard", fire_mage))

print("\n📊 After Round 1:")
print(fire_mage.status())
print(ice_mage.status())
```

---

## 🎯 Interactive Features

### REPL Magic
```python
// In the REPL, try these magical commands:
help()      // Basic help
version()   // Version information
modules()   // Available modules
mimir       // Interactive help system! (Type this in REPL)

// Tab completion works for:
print  // <TAB> - shows print function
Array  // <TAB> - shows Array grimoire
"hello".  // <TAB> - shows string methods
```

### Automatic Enhancements
```python
// These work automatically without importing:
[1, 2, 3].sort()           // Array methods
"hello".upper()            // String methods
42.to_bin()               // Integer methods
3.14.round(1)             // Float methods
True.to_int()             // Boolean methods
```

---

## 🚀 What's Next?

Congratulations! You've learned Carrion basics in 30 minutes. Here's your magical progression path:

### Immediate Next Steps:
1. **🎮 [REPL Guide](REPL-Guide.md)** - Master the interactive environment
2. **📚 [Language Fundamentals](../Language-Fundamentals/Syntax-and-Terminology.md)** - Deeper dive into syntax
3. **🏰 [Advanced Grimoires](../Advanced-Features/Classes-Grimoires.md)** - Advanced OOP patterns

### Build Projects:
1. **🎲 Dice Rolling Game** - Practice spells and randomization
2. **📚 Spell Book Manager** - Use arrays and file operations
3. **⚔️ Text Adventure** - Combine grimoires and control flow
4. **🧮 Magical Calculator** - Advanced mathematical operations

### Explore Standard Library:
1. **[📋 Array Grimoire](../Standard-Library/Array-Grimoire.md)** - Master list manipulation
2. **[🔤 String Grimoire](../Standard-Library/String-Grimoire.md)** - Text processing magic
3. **[📁 File & OS](../Standard-Library/File-OS-Grimoires.md)** - System interaction

---

## 💡 Quick Reference Card

| Python | Carrion | Purpose |
|--------|---------|---------|
| `def` | `spell` | Define function |
| `class` | `grim` | Define class |
| `elif` | `otherwise` | Else-if condition |
| `continue` | `skip` | Skip iteration |
| `break` | `stop` | Break loop |
| `try/except/finally` | `attempt/ensnare/resolve` | Error handling |
| `self.__init__` | `init` | Constructor |
| `super()` | `super` | Parent class access |

---

## 🎊 You're Now a Carrion Mage!

You've successfully learned:
- ✅ **Variables & Data Types** - Storing magical energy
- ✅ **Spells (Functions)** - Creating reusable magic
- ✅ **Grimoires (Classes)** - Object-oriented spellcrafting
- ✅ **Standard Library** - Using Munin's power
- ✅ **Control Flow** - Directing magical energy
- ✅ **Error Handling** - Protecting against magical mishaps

*Welcome to the mystical world of Carrion programming! May your code be bug-free and your spells compile cleanly! 🪄✨*

> "The journey of a thousand spells begins with a single `print('Hello, World!')` incantation." - *Ancient Carrion Wisdom*