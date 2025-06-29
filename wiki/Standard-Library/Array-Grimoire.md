# 📋 Array Grimoire - Master of List Manipulation

Welcome to the **Array Grimoire**, one of the most powerful and frequently used grimoires in the Munin standard library! Arrays are the workhorses of data storage and manipulation in Carrion, and the Array grimoire transforms simple lists into powerful magical collections.

---

## 🪄 What is the Array Grimoire?

The **Array Grimoire** is a magical enhancement that wraps around Carrion's basic arrays, providing them with powerful spells for:
- 🔍 **Searching** - Find elements quickly and efficiently
- 🛠️ **Manipulation** - Add, remove, and modify elements
- 📊 **Analysis** - Get insights about your data
- 🔄 **Transformation** - Sort, reverse, and reorganize
- ✨ **Validation** - Check properties and states

---

## 🎭 Creating Array Magic

### Basic Array Creation
```python
// Create arrays the normal way
basic_array = [1, 2, 3, 4, 5]
mixed_array = ["wizard", 42, True, 3.14]
empty_array = []

// Enhance them with Array grimoire magic
magical_numbers = Array([1, 2, 3, 4, 5])
magical_mix = Array(["wizard", 42, True, 3.14])
magical_empty = Array([])

// Arrays also get automatic enhancement!
auto_enhanced = [10, 20, 30]
print(auto_enhanced.length())  // Works automatically! Returns 3
```

### Array Grimoire Constructor
```python
// Various ways to create magical arrays
numbers = Array([1, 2, 3, 4, 5])
words = Array(["hello", "magical", "world"])
from_range = Array(range(1, 11))  // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
copy_array = Array(existing_array)
```

---

## 📊 Information & Analysis Spells

### `.length()` - Measure Your Collection
```python
spells = Array(["Fireball", "Healing", "Teleport", "Shield"])
print(f"Spell count: {spells.length()}")  // Spell count: 4

empty_grimoire = Array([])
print(f"Empty grimoire size: {empty_grimoire.length()}")  // 0

// Also works with auto-enhanced arrays
auto_array = [1, 2, 3]
print(f"Auto array length: {auto_array.length()}")  // 3
```

### `.is_empty()` - Check for Magical Presence
```python
full_inventory = Array(["Sword", "Shield", "Potion"])
empty_bag = Array([])

print(f"Full inventory empty? {full_inventory.is_empty()}")  // False
print(f"Empty bag empty? {empty_bag.is_empty()}")           // True

// Perfect for conditional magic
if not my_spells.is_empty():
    print("Ready to cast magic!")
else:
    print("Need to learn some spells first!")
```

### `.contains(value)` - Detect Magical Presence
```python
magical_elements = Array(["fire", "water", "earth", "air", "spirit"])

print(magical_elements.contains("fire"))     // True
print(magical_elements.contains("ice"))      // False
print(magical_elements.contains("spirit"))   // True

// Case-sensitive checking
creatures = Array(["Dragon", "dragon", "DRAGON"])
print(creatures.contains("Dragon"))  // True
print(creatures.contains("dragon"))  // True
print(creatures.contains("DRAGON"))  // True
```

---

## 🔍 Search & Access Spells

### `.get(index)` - Retrieve Magical Items
```python
treasure_chest = Array(["Gold", "Gems", "Magic Ring", "Ancient Scroll"])

// Positive indexing (0-based)
print(treasure_chest.get(0))   // "Gold"
print(treasure_chest.get(2))   // "Magic Ring"

// Negative indexing (from the end)
print(treasure_chest.get(-1))  // "Ancient Scroll" (last item)
print(treasure_chest.get(-2))  // "Magic Ring" (second to last)

// Safe access (returns None for invalid indices)
print(treasure_chest.get(10))  // None (index too high)
print(treasure_chest.get(-10)) // None (index too low)
```

### `.index_of(value)` - Find Magical Locations
```python
spell_book = Array(["Heal", "Fireball", "Teleport", "Fireball", "Shield"])

first_fireball = spell_book.index_of("Fireball")
print(f"First Fireball at index: {first_fireball}")  // 1

teleport_location = spell_book.index_of("Teleport")
print(f"Teleport at index: {teleport_location}")     // 2

// Returns -1 if not found
lightning_location = spell_book.index_of("Lightning Bolt")
print(f"Lightning Bolt location: {lightning_location}")  // -1

// Useful for conditional logic
if spell_book.index_of("Heal") != -1:
    print("Healing magic available!")
```

### `.first()` & `.last()` - Access Magical Extremes
```python
magical_journey = Array(["Start", "Forest", "Mountain", "Castle", "End"])

beginning = magical_journey.first()
destination = magical_journey.last()

print(f"Journey begins at: {beginning}")      // "Start"
print(f"Journey ends at: {destination}")      // "End"

// Safe for empty arrays
empty_path = Array([])
print(empty_path.first())  // None
print(empty_path.last())   // None
```

---

## 🛠️ Modification Spells

### `.append(value)` - Add Magical Items
```python
spell_arsenal = Array(["Magic Missile"])

spell_arsenal.append("Fireball")
spell_arsenal.append("Healing")
print(spell_arsenal.to_string())  // Array(["Magic Missile", "Fireball", "Healing"])

// Append different types
magical_collection = Array([42])
magical_collection.append("text")
magical_collection.append(True)
magical_collection.append(3.14)
print(magical_collection.to_string())  // Mixed types array
```

### `.set(index, value)` - Modify Magical Items
```python
party_members = Array(["Warrior", "Placeholder", "Healer"])

// Replace the placeholder
party_members.set(1, "Wizard")
print(party_members.to_string())  // Array(["Warrior", "Wizard", "Healer"])

// Update with negative indexing
party_members.set(-1, "Cleric")
print(party_members.to_string())  // Array(["Warrior", "Wizard", "Cleric"])

// Safe setting (does nothing for invalid indices)
party_members.set(10, "Ghost")  // No error, just ignored
```

### `.remove(value)` - Banish Magical Items
```python
magic_ingredients = Array(["Eye of Newt", "Bad Ingredient", "Dragon Scale", "Bad Ingredient"])

// Remove first occurrence
magic_ingredients.remove("Bad Ingredient")
print(magic_ingredients.to_string())  // Array(["Eye of Newt", "Dragon Scale", "Bad Ingredient"])

// Remove again to get the second occurrence
magic_ingredients.remove("Bad Ingredient")
print(magic_ingredients.to_string())  // Array(["Eye of Newt", "Dragon Scale"])

// Safe removal (no error if item not found)
magic_ingredients.remove("Nonexistent Item")  // Just ignored
```

### `.clear()` - Empty Your Magical Container
```python
temporary_storage = Array([1, 2, 3, 4, 5])
print(f"Before clear: {temporary_storage.length()}")  // 5

temporary_storage.clear()
print(f"After clear: {temporary_storage.length()}")   // 0
print(f"Is empty: {temporary_storage.is_empty()}")    // True
```

---

## 🔄 Transformation Spells

### `.reverse()` - Mirror Your Magic
```python
countdown = Array([1, 2, 3, 4, 5])
reversed_countdown = countdown.reverse()

print(f"Original: {countdown.to_string()}")          // [1, 2, 3, 4, 5]
print(f"Reversed: {reversed_countdown.to_string()}")  // [5, 4, 3, 2, 1]

// Original array is unchanged - reverse() creates a new array
print(f"Original still: {countdown.to_string()}")     // [1, 2, 3, 4, 5]

// Works with mixed types
mixed = Array(["first", 42, True, "last"])
reversed_mixed = mixed.reverse()
print(reversed_mixed.to_string())  // ["last", True, 42, "first"]
```

### `.sort()` - Organize Your Magic
```python
// Sort numbers
numbers = Array([3, 1, 4, 1, 5, 9, 2, 6])
sorted_numbers = numbers.sort()
print(f"Sorted numbers: {sorted_numbers.to_string()}")  // [1, 1, 2, 3, 4, 5, 6, 9]

// Sort strings (alphabetical)
spells = Array(["Fireball", "Heal", "Teleport", "Shield"])
sorted_spells = spells.sort()
print(f"Sorted spells: {sorted_spells.to_string()}")  // ["Fireball", "Heal", "Shield", "Teleport"]

// Mixed types sort by string representation
mixed = Array([3, "apple", 1, "zebra"])
sorted_mixed = mixed.sort()
print(f"Sorted mixed: {sorted_mixed.to_string()}")

// Original array remains unchanged
print(f"Original numbers: {numbers.to_string()}")  // Still [3, 1, 4, 1, 5, 9, 2, 6]
```

### `.slice(start, end)` - Extract Magical Segments
```python
magical_alphabet = Array(["A", "B", "C", "D", "E", "F", "G"])

// Get middle section
middle = magical_alphabet.slice(2, 5)
print(middle.to_string())  // ["C", "D", "E"]

// Get from start
beginning = magical_alphabet.slice(0, 3)
print(beginning.to_string())  // ["A", "B", "C"]

// Get to end (omit end parameter)
end_part = magical_alphabet.slice(4, 7)
print(end_part.to_string())  // ["E", "F", "G"]

// Negative indices work too
last_three = magical_alphabet.slice(-3, 7)
print(last_three.to_string())  // ["E", "F", "G"]

// Original array unchanged
print(f"Original: {magical_alphabet.to_string()}")
```

---

## 🎪 Advanced Array Magic

### Chaining Array Spells
```python
// Build complex transformations by chaining methods
numbers = Array([5, 2, 8, 1, 9, 3])

// Note: Some methods modify the array, others return new arrays
result = numbers.sort()  // Returns new sorted array
print(f"Sorted: {result.to_string()}")

// For methods that modify the original, you can still chain
magical_list = Array([])
magical_list.append("first")
magical_list.append("second")
magical_list.append("third")
print(f"Built list: {magical_list.to_string()}")
```

### Working with Complex Data
```python
// Arrays can hold complex magical data
magical_creatures = Array([
    {"name": "Dragon", "power": 1000, "element": "fire"},
    {"name": "Unicorn", "power": 500, "element": "light"},
    {"name": "Phoenix", "power": 800, "element": "fire"}
])

print(f"Creature count: {magical_creatures.length()}")

// Access complex data
first_creature = magical_creatures.get(0)
print(f"First creature: {first_creature['name']} with {first_creature['power']} power")

// Check for specific creatures
has_dragon = magical_creatures.contains({"name": "Dragon", "power": 1000, "element": "fire"})
print(f"Has dragon: {has_dragon}")  // Note: Object comparison might be tricky
```

### Array Transformation Patterns
```python
// Common magical patterns
spell create_magical_sequence(start, end, step = 1):
    sequence = Array([])
    current = start
    while current < end:
        sequence.append(current)
        current += step
    return sequence

spell filter_by_element(creatures, target_element):
    filtered = Array([])
    for i in range(creatures.length()):
        creature = creatures.get(i)
        if creature["element"] == target_element:
            filtered.append(creature)
    return filtered

spell sum_magical_powers(creatures):
    total = 0
    for i in range(creatures.length()):
        creature = creatures.get(i)
        total += creature["power"]
    return total

// Use the patterns
sequence = create_magical_sequence(0, 10, 2)
print(f"Magical sequence: {sequence.to_string()}")  // [0, 2, 4, 6, 8]

creatures = Array([
    {"name": "Fire Dragon", "element": "fire", "power": 1000},
    {"name": "Ice Dragon", "element": "ice", "power": 900},
    {"name": "Fire Phoenix", "element": "fire", "power": 800}
])

fire_creatures = filter_by_element(creatures, "fire")
print(f"Fire creatures: {fire_creatures.length()}")

total_power = sum_magical_powers(creatures)
print(f"Total magical power: {total_power}")
```

---

## 🎯 Practical Array Magic Examples

### Magical Inventory System
```python
spell create_magical_inventory():
    inventory = Array([])
    
    spell add_item(item_name, quantity = 1):
        // Check if item already exists
        for i in range(inventory.length()):
            item = inventory.get(i)
            if item["name"] == item_name:
                item["quantity"] += quantity
                return f"Added {quantity} {item_name}(s). Total: {item['quantity']}"
        
        // Add new item
        new_item = {"name": item_name, "quantity": quantity}
        inventory.append(new_item)
        return f"Added new item: {quantity} {item_name}(s)"
    
    spell remove_item(item_name, quantity = 1):
        for i in range(inventory.length()):
            item = inventory.get(i)
            if item["name"] == item_name:
                if item["quantity"] >= quantity:
                    item["quantity"] -= quantity
                    if item["quantity"] == 0:
                        inventory.remove(item)
                    return f"Removed {quantity} {item_name}(s)"
                else:
                    return f"Not enough {item_name} (have {item['quantity']}, need {quantity})"
        return f"Item {item_name} not found in inventory"
    
    spell list_items():
        if inventory.is_empty():
            return "Inventory is empty"
        
        items = []
        for i in range(inventory.length()):
            item = inventory.get(i)
            items.append(f"{item['name']}: {item['quantity']}")
        return "Inventory: " + ", ".join(items)
    
    return {
        "add": add_item,
        "remove": remove_item,
        "list": list_items,
        "raw": inventory
    }

// Use the magical inventory
bag = create_magical_inventory()
print(bag["add"]("Health Potion", 5))
print(bag["add"]("Magic Sword", 1))
print(bag["add"]("Health Potion", 2))
print(bag["list"]())
print(bag["remove"]("Health Potion", 3))
print(bag["list"]())
```

### Magical Battle Queue System
```python
grim BattleMage:
    init(name, speed, health):
        self.name = name
        self.speed = speed
        self.health = health
        self.is_alive = True
    
    spell to_string():
        status = "💚" if self.is_alive else "💀"
        return f"{status} {self.name} (Speed: {self.speed}, HP: {self.health})"

spell create_battle_queue():
    battle_queue = Array([])
    
    spell add_combatant(mage):
        // Insert based on speed (highest first)
        inserted = False
        for i in range(battle_queue.length()):
            if mage.speed > battle_queue.get(i).speed:
                battle_queue.set(i, mage)  // This replaces, not inserts
                inserted = True
                stop
        
        if not inserted:
            battle_queue.append(mage)
        
        return f"Added {mage.name} to battle queue"
    
    spell get_next_combatant():
        if not battle_queue.is_empty():
            next_mage = battle_queue.get(0)
            battle_queue.remove(next_mage)
            return next_mage
        return None
    
    spell show_queue():
        if battle_queue.is_empty():
            return "Battle queue is empty"
        
        queue_display = "Battle Queue:\n"
        for i in range(battle_queue.length()):
            mage = battle_queue.get(i)
            queue_display += f"{i + 1}. {mage.to_string()}\n"
        return queue_display
    
    return {
        "add": add_combatant,
        "next": get_next_combatant,
        "show": show_queue,
        "queue": battle_queue
    }

// Create battle system
battle = create_battle_queue()
mage1 = BattleMage("Fire Wizard", 85, 100)
mage2 = BattleMage("Ice Sorceress", 90, 95)
mage3 = BattleMage("Lightning Mage", 95, 90)

battle["add"](mage1)
battle["add"](mage2)
battle["add"](mage3)

print(battle["show"]())

next_fighter = battle["next"]()
print(f"Next to act: {next_fighter.to_string()}")
print(battle["show"]())
```

---

## 🏆 Array Magic Challenges

### 🥉 Beginner Challenges
1. **Spell List Manager**: Create an array of spells and implement add/remove/search functionality
2. **Magic Number Calculator**: Use arrays to store and calculate magical number sequences
3. **Ingredient Sorter**: Sort magical ingredients alphabetically and by rarity

### 🥈 Intermediate Challenges
1. **Magical Deck Builder**: Create a card game deck with shuffling and dealing
2. **Adventure Path Finder**: Use arrays to represent and navigate magical paths
3. **Spell Combo System**: Build combinations of spells stored in arrays

### 🥇 Advanced Challenges
1. **Magical Matrix Operations**: Implement matrix operations using nested arrays
2. **Spell Pattern Recognition**: Find patterns in spell sequences using array analysis
3. **Dynamic Quest System**: Create branching quest paths using arrays

---

## 🔗 Related Magic

Explore other grimoires that work well with arrays:

- **[String Grimoire](String-Grimoire.md)** - For text arrays and string manipulation
- **[Number Grimoires](Number-Grimoires.md)** - For numerical array operations
- **[File & OS Grimoires](File-OS-Grimoires.md)** - For reading/writing array data
- **[Builtin Functions](../Reference/Builtin-Functions.md)** - Core functions that work with arrays

---

## 💡 Array Mastery Tips

1. **🔄 Remember Mutability**: Some methods modify the array, others return new ones
2. **🔍 Use `.contains()` for Existence**: Better than manually searching
3. **📐 Leverage Negative Indexing**: `.get(-1)` for last element
4. **🛡️ Handle Empty Arrays**: Always check `.is_empty()` when needed
5. **⚡ Chain Operations**: Combine methods for powerful transformations
6. **📊 Use Arrays for Complex Data**: Store objects and hashes in arrays

---

*Master the Array Grimoire, and you'll have the power to organize, search, and manipulate any collection of magical data! Arrays are the foundation of many advanced spells and enchantments. 🪄✨*

> "An organized array is a powerful spell waiting to be cast." - *Master Array Sorceress*