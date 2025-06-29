# ✨ Functions (Spells) - Creating Reusable Magic

In the mystical realm of Carrion, functions are called **"spells"** - powerful incantations that can be cast repeatedly to perform magical operations. Master the art of spellcraft and transform repetitive tasks into elegant, reusable magic!

---

## 🪄 What Are Spells?

**Spells** (functions) are named blocks of code that:
- 🎯 **Perform specific tasks** - Each spell has a clear magical purpose
- 🔄 **Can be reused** - Cast the same spell multiple times
- 📥 **Accept ingredients** - Take parameters to customize their behavior
- 📤 **Return magical results** - Produce outputs for other spells to use
- 🏗️ **Organize code** - Keep your grimoire (program) clean and readable

---

## 🎭 Basic Spell Creation

### The Spell Incantation Syntax
```python
spell spell_name(parameters):
    // Magical operations go here
    return result  // Optional: return magical result
```

### Your First Simple Spell
```python
spell greet():
    print("Hello, fellow mage!")

// Cast the spell
greet()  // Output: Hello, fellow mage!
```

### Spells with Magical Ingredients (Parameters)
```python
spell greet_wizard(name):
    print(f"Greetings, {name}! Welcome to the mystical realm!")

// Cast spells with ingredients
greet_wizard("Merlin")     // Output: Greetings, Merlin! Welcome to the mystical realm!
greet_wizard("Gandalf")    // Output: Greetings, Gandalf! Welcome to the mystical realm!
```

### Spells that Return Magical Energy
```python
spell add_magical_numbers(a, b):
    result = a + b
    return result

// Use the returned magic
total = add_magical_numbers(10, 20)
print(f"The magical sum is: {total}")  // Output: The magical sum is: 30

// Or use directly in other spells
print(f"Double magic: {add_magical_numbers(5, 5) * 2}")
```

---

## 🧙‍♂️ Advanced Spellcrafting

### Spells with Default Magical Ingredients
```python
spell cast_fireball(power = 50, target = "enemy"):
    return f"🔥 Casting fireball with {power} power at {target}!"

// Cast with default values
print(cast_fireball())                    // Uses defaults: power=50, target="enemy"

// Cast with custom power
print(cast_fireball(100))                 // Custom power, default target

// Cast with all custom values
print(cast_fireball(75, "dragon"))        // Custom power and target
```

### Spells with Multiple Return Values
```python
spell analyze_magical_creature(creature_name):
    if creature_name == "dragon":
        return "fire", 1000, True      // element, power, dangerous
    elif creature_name == "unicorn":
        return "light", 500, False
    else:
        return "unknown", 100, False

// Unpack the magical analysis
element, power, is_dangerous = analyze_magical_creature("dragon")
print(f"Element: {element}, Power: {power}, Dangerous: {is_dangerous}")
```

### Spells with Variable Ingredients (*args)
```python
spell combine_spell_components(*components):
    if len(components) == 0:
        return "Empty spell - no magic!"
    
    spell_power = len(components) * 10
    component_list = ", ".join(components)
    return f"Spell with {component_list} has power level {spell_power}!"

// Cast with different numbers of components
print(combine_spell_components("dragon scale"))
print(combine_spell_components("eye of newt", "wing of bat"))
print(combine_spell_components("moonstone", "star dust", "phoenix feather", "mermaid tear"))
```

---

## 🎪 Practical Spell Examples

### Mathematical Magic Spells
```python
spell calculate_spell_damage(base_damage, level, has_magic_weapon = False):
    damage = base_damage * level
    if has_magic_weapon:
        damage = damage * 1.5
    return int(damage)

spell is_prime_magical_number(number):
    if number < 2:
        return False
    for i in range(2, int(number ** 0.5) + 1):
        if number % i == 0:
            return False
    return True

spell factorial_spell(n):
    if n <= 1:
        return 1
    return n * factorial_spell(n - 1)  // Recursive magic!

// Test mathematical spells
print(f"Spell damage: {calculate_spell_damage(20, 5, True)}")    // 150
print(f"Is 17 magical? {is_prime_magical_number(17)}")           // True
print(f"5! = {factorial_spell(5)}")                             // 120
```

### Text Magic Spells
```python
spell create_magical_title(text):
    words = text.split(" ")
    magical_words = []
    for word in words:
        if len(word) > 0:
            magical_word = word[0].upper() + word[1:].lower()
            magical_words.append(magical_word)
    return " ".join(magical_words)

spell reverse_spell_incantation(text):
    return text[::-1]

spell count_magical_letters(text, target_letter):
    count = 0
    for letter in text.lower():
        if letter == target_letter.lower():
            count += 1
    return count

// Test text magic
print(create_magical_title("the ancient spell of power"))    // "The Ancient Spell Of Power"
print(reverse_spell_incantation("abracadabra"))             // "arbadacarba"
print(count_magical_letters("Expelliarmus", "a"))           // 2
```

### List Magic Spells
```python
spell find_magical_items(inventory, item_type):
    magical_items = []
    for item in inventory:
        if item_type.lower() in item.lower():
            magical_items.append(item)
    return magical_items

spell sort_by_magical_power(creatures):
    // Simple bubble sort for magical creatures
    n = len(creatures)
    for i in range(n):
        for j in range(0, n - i - 1):
            if creatures[j][1] > creatures[j + 1][1]:  // Compare power levels
                creatures[j], creatures[j + 1] = creatures[j + 1], creatures[j]
    return creatures

// Test list magic
inventory = ["Fire Wand", "Ice Shield", "Lightning Staff", "Healing Potion", "Fire Ring"]
fire_items = find_magical_items(inventory, "fire")
print(f"Fire items: {fire_items}")

creatures = [("Dragon", 1000), ("Unicorn", 500), ("Phoenix", 800), ("Griffin", 600)]
sorted_creatures = sort_by_magical_power(creatures)
print(f"Sorted by power: {sorted_creatures}")
```

---

## 🎯 Spell Scope & Magic Variables

### Local Magic (Local Variables)
```python
spell local_magic_example():
    local_magic = "This magic only exists within this spell"
    magic_power = 100
    print(f"Local magic: {local_magic}")
    return magic_power

local_magic_example()
// print(local_magic)  // ❌ Error! local_magic doesn't exist outside the spell
```

### Global Magic (Global Variables)
```python
global_magic_power = 1000  // Global magical energy

spell use_global_magic():
    print(f"Accessing global magic: {global_magic_power}")
    local_bonus = 50
    return global_magic_power + local_bonus

spell modify_global_magic(new_power):
    global global_magic_power  // Access global magic
    global_magic_power = new_power
    print(f"Global magic changed to: {global_magic_power}")

// Test global magic
print(f"Initial power: {global_magic_power}")
print(f"Used power: {use_global_magic()}")
modify_global_magic(1500)
print(f"Final power: {global_magic_power}")
```

---

## 🧩 Nested Spells & Higher-Order Magic

### Spells within Spells
```python
spell create_spell_factory(spell_type):
    spell fire_spell():
        return "🔥 Fireball cast!"
    
    spell ice_spell():
        return "❄️ Ice shard launched!"
    
    spell healing_spell():
        return "💚 Healing magic activated!"
    
    // Return the appropriate inner spell
    if spell_type == "fire":
        return fire_spell()
    elif spell_type == "ice":
        return ice_spell()
    elif spell_type == "healing":
        return healing_spell()
    else:
        return "🔮 Unknown magic attempted!"

// Use the spell factory
print(create_spell_factory("fire"))     // 🔥 Fireball cast!
print(create_spell_factory("healing"))  // 💚 Healing magic activated!
```

### Spells as Magical Ingredients
```python
spell apply_magical_operation(numbers, operation_spell):
    results = []
    for number in numbers:
        result = operation_spell(number)
        results.append(result)
    return results

spell double_magic(x):
    return x * 2

spell square_magic(x):
    return x ** 2

spell magical_cube(x):
    return x ** 3

// Use spells as ingredients for other spells
numbers = [1, 2, 3, 4, 5]
doubled = apply_magical_operation(numbers, double_magic)
squared = apply_magical_operation(numbers, square_magic)
cubed = apply_magical_operation(numbers, magical_cube)

print(f"Original: {numbers}")
print(f"Doubled: {doubled}")
print(f"Squared: {squared}")
print(f"Cubed: {cubed}")
```

---

## 🎨 Spell Patterns & Best Practices

### The Single Responsibility Spell Principle
```python
// ❌ Bad: Spell does too many things
spell bad_character_manager(name, action, value):
    if action == "create":
        // create character logic
        return f"Created {name}"
    elif action == "level_up":
        // level up logic
        return f"{name} leveled up"
    elif action == "calculate_damage":
        // damage calculation logic
        return value * 10
    // ... too many responsibilities

// ✅ Good: Separate spells for separate responsibilities
spell create_character(name, class_type):
    character = {
        "name": name,
        "class": class_type,
        "level": 1,
        "health": 100
    }
    return character

spell level_up_character(character):
    character["level"] += 1
    character["health"] += 20
    return character

spell calculate_spell_damage(base_damage, level, spell_modifier = 1.0):
    return base_damage * level * spell_modifier
```

### Magical Error Handling in Spells
```python
spell safe_magical_division(numerator, denominator):
    attempt:
        if denominator == 0:
            return "Error: Cannot divide by zero in magical calculations!"
        return numerator / denominator
    ensnare:
        return "Error: Something went wrong with the magical division!"

spell validate_spell_components(components):
    if not components:
        return False, "No spell components provided!"
    
    required_components = ["base", "modifier", "catalyst"]
    for required in required_components:
        if required not in components:
            return False, f"Missing required component: {required}"
    
    return True, "All spell components are present!"

// Test error handling
print(safe_magical_division(10, 2))    // 5.0
print(safe_magical_division(10, 0))    // Error message

components = {"base": "fire", "modifier": "wind"}
is_valid, message = validate_spell_components(components)
print(f"Valid: {is_valid}, Message: {message}")
```

### Documentation Spells (Docstrings)
```python
spell calculate_magical_distance(x1, y1, x2, y2):
    """
    Calculate the magical distance between two points in the mystical realm.
    
    This spell uses the ancient Pythagorean theorem to determine the straight-line
    distance between two magical locations.
    
    Parameters:
        x1, y1: Coordinates of the first magical location
        x2, y2: Coordinates of the second magical location
    
    Returns:
        The magical distance as a floating-point number
    
    Example:
        distance = calculate_magical_distance(0, 0, 3, 4)
        # Returns 5.0 (the classic 3-4-5 triangle)
    """
    dx = x2 - x1
    dy = y2 - y1
    return (dx ** 2 + dy ** 2) ** 0.5

// The docstring can be accessed (if supported by the interpreter)
print(calculate_magical_distance(0, 0, 3, 4))  // 5.0
```

---

## 🎮 Interactive Spell Examples

### Magical Menu System
```python
spell display_magical_menu():
    print("\n🪄 Magical Menu 🪄")
    print("1. Cast Fireball")
    print("2. Cast Healing Spell")
    print("3. Check Mana")
    print("4. Exit")
    return input("Choose your magic (1-4): ")

spell cast_fireball():
    return "🔥 WHOOSH! Fireball explodes in brilliant flames!"

spell cast_healing_spell():
    return "💚 Gentle light surrounds you, healing your wounds."

spell check_mana():
    return "✨ Your mana reserves are at 85/100."

spell magical_menu_system():
    mana = 100
    while True:
        choice = display_magical_menu()
        
        if choice == "1":
            if mana >= 20:
                print(cast_fireball())
                mana -= 20
            else:
                print("❌ Not enough mana for fireball!")
        elif choice == "2":
            if mana >= 10:
                print(cast_healing_spell())
                mana -= 10
            else:
                print("❌ Not enough mana for healing!")
        elif choice == "3":
            print(f"✨ Current mana: {mana}/100")
        elif choice == "4":
            print("👋 Farewell, brave mage!")
            stop
        else:
            print("❓ Unknown magic. Please choose 1-4.")

// Uncomment to run the magical menu
// magical_menu_system()
```

---

## 🏆 Practice Challenges

### 🥉 Beginner Spells
1. **Magical Calculator**: Create spells for +, -, *, / operations
2. **Name Formatter**: Create a spell that formats names properly
3. **Temperature Converter**: Convert between Celsius and Fahrenheit

### 🥈 Intermediate Spells
1. **Magical Password Generator**: Generate secure magical passwords
2. **Spell Component Analyzer**: Analyze ingredient lists for completeness
3. **Magical Number Guesser**: Create a number guessing game with spells

### 🥇 Advanced Spells
1. **Magical Inventory System**: Complete item management with spells
2. **Spell Damage Calculator**: Complex damage calculation with multiple factors
3. **Magical Text Adventure**: Create a mini text adventure using spells

---

## 🔗 What's Next?

Now that you've mastered the art of spellcrafting, expand your magical knowledge:

1. **🏰 [Classes (Grimoires)](../Advanced-Features/Classes-Grimoires.md)** - Create magical objects and classes
2. **⚔️ [Error Handling](../Advanced-Features/Error-Handling.md)** - Protect your spells from magical mishaps
3. **📦 [Modules & Imports](../Advanced-Features/Modules.md)** - Organize your spell libraries
4. **📚 [Standard Library](../Standard-Library/Munin-Overview.md)** - Use pre-built magical spells

---

## 💡 Spellcrafting Wisdom

1. **🎯 Keep spells focused**: Each spell should do one thing well
2. **📖 Use descriptive names**: `calculate_damage()` is better than `calc()`
3. **🛡️ Handle magical errors**: Always consider what could go wrong
4. **📝 Document your magic**: Future you (and other mages) will thank you
5. **🧪 Test your spells**: Make sure they work in all magical conditions
6. **🔄 Avoid repetition**: If you're copying code, create a spell instead

---

*Remember: The most powerful mages are those who master the art of creating reusable, elegant spells. Your magical library grows with every spell you craft! 🪄✨*

> "A spell well-crafted is a gift to your future self and all who follow in your magical footsteps." - *Master Spellcrafter's Wisdom*