# 🔤 String Grimoire - Master of Text Magic

Welcome to the **String Grimoire**, the mystical tome of text manipulation in the Munin standard library! Transform simple text into powerful magical incantations with sophisticated string operations, text analysis, and character manipulation spells.

---

## 🪄 What is the String Grimoire?

The **String Grimoire** enhances Carrion's basic strings with powerful text magic:
- 🔍 **Text Analysis** - Search, find, and analyze text content
- 🎭 **Transformation** - Change case, reverse, and modify text
- 📐 **Character Access** - Precise character-level manipulation
- ✨ **Validation** - Check text properties and content
- 🔧 **Text Building** - Construct complex strings efficiently

---

## 🎭 Creating String Magic

### Basic String Enhancement
```python
// Create strings the normal way
basic_text = "Hello, magical world!"
spell_name = "Fireball"
incantation = "Abracadabra"

// Enhance them with String grimoire magic
magical_text = String("Hello, magical world!")
enhanced_spell = String(spell_name)

// Strings also get automatic enhancement!
auto_enhanced = "Automatic magic!"
print(auto_enhanced.upper())  // Works automatically! Returns "AUTOMATIC MAGIC!"
```

### String Grimoire Constructor
```python
// Various ways to create magical strings
from_text = String("Ancient runes")
from_variable = String(some_variable)
from_number = String(42)           // "42"
from_boolean = String(True)        // "True"
empty_spell = String("")
```

---

## 📊 Text Analysis Spells

### `.length()` - Measure Your Text Magic
```python
incantation = String("Expecto Patronum")
spell_power = incantation.length()
print(f"Spell has {spell_power} characters")  // 15 characters

// Also works with auto-enhanced strings
auto_text = "Magical!"
print(f"Length: {auto_text.length()}")  // 8

// Perfect for validation
password = String(input("Enter magical password: "))
if password.length() >= 8:
    print("Password meets minimum length requirement!")
else:
    print(f"Password too short! Need {8 - password.length()} more characters.")
```

### `.contains(substring)` - Detect Text Presence
```python
magical_tome = String("The ancient grimoire contains powerful fire spells")

// Search for magical elements
has_fire = magical_tome.contains("fire")
has_water = magical_tome.contains("water")
has_ancient = magical_tome.contains("ancient")

print(f"Contains fire magic: {has_fire}")      // True
print(f"Contains water magic: {has_water}")    // False
print(f"Contains ancient knowledge: {has_ancient}")  // True

// Case-sensitive searching
spell_book = String("FireBall")
print(spell_book.contains("fire"))     // False (lowercase 'f')
print(spell_book.contains("Fire"))     // True
print(spell_book.contains("FIRE"))     // False (all caps)
```

### `.find(substring)` - Locate Magical Text
```python
magical_recipe = String("Mix dragon scale with phoenix feather and unicorn hair")

// Find ingredient positions
dragon_pos = magical_recipe.find("dragon")
phoenix_pos = magical_recipe.find("phoenix")
unicorn_pos = magical_recipe.find("unicorn")
missing_pos = magical_recipe.find("troll blood")

print(f"Dragon mentioned at position: {dragon_pos}")    // 4
print(f"Phoenix mentioned at position: {phoenix_pos}")  // 22
print(f"Unicorn mentioned at position: {unicorn_pos}")  // 43
print(f"Troll blood position: {missing_pos}")          // -1 (not found)

// Use for conditional magic
if magical_recipe.find("dangerous") != -1:
    print("⚠️ Warning: Recipe contains dangerous ingredients!")
else:
    print("✅ Recipe appears safe to brew")
```

---

## 🔍 Character Access Spells

### `.char_at(index)` - Extract Individual Characters
```python
magic_word = String("ABRACADABRA")

// Positive indexing (0-based)
first_char = magic_word.char_at(0)      // "A"
third_char = magic_word.char_at(2)      // "R"
print(f"First character: {first_char}")
print(f"Third character: {third_char}")

// Negative indexing (from the end)
last_char = magic_word.char_at(-1)      // "A"
second_last = magic_word.char_at(-2)    // "R"
print(f"Last character: {last_char}")
print(f"Second to last: {second_last}")

// Safe access (returns empty string for invalid indices)
beyond = magic_word.char_at(20)         // ""
print(f"Beyond range: '{beyond}'")

// Build character maps
spell analyze_spell_characters(spell_text):
    spell_string = String(spell_text)
    char_frequency = {}
    
    for i in range(spell_string.length()):
        char = spell_string.char_at(i)
        if char in char_frequency:
            char_frequency[char] += 1
        else:
            char_frequency[char] = 1
    
    return char_frequency

char_analysis = analyze_spell_characters("ABRACADABRA")
print(f"Character frequency: {char_analysis}")
// Output: {"A": 5, "B": 2, "R": 2, "C": 1, "D": 1}
```

---

## 🎭 Text Transformation Spells

### `.upper()` & `.lower()` - Change Magical Case
```python
quiet_spell = String("whispered healing charm")
loud_spell = String("THUNDEROUS LIGHTNING BOLT")

// Transform case
amplified = quiet_spell.upper()
hushed = loud_spell.lower()

print(f"Amplified: {amplified}")  // "WHISPERED HEALING CHARM"
print(f"Hushed: {hushed}")        // "thunderous lightning bolt"

// Perfect for user input normalization
user_command = String(input("Enter spell name: "))
normalized_command = user_command.lower()

if normalized_command == "fireball":
    print("🔥 Casting Fireball!")
elif normalized_command == "heal":
    print("💚 Casting Healing Magic!")
else:
    print("❓ Unknown spell command")

// Create title case manually
spell create_title_case(text):
    title_string = String(text.lower())
    words = title_string.to_string().split(" ")
    title_words = []
    
    for word in words:
        if len(word) > 0:
            word_string = String(word)
            first_char = word_string.char_at(0).upper()
            rest_chars = word_string.to_string()[1:]
            title_words.append(first_char + rest_chars)
    
    return " ".join(title_words)

title = create_title_case("the ancient book of fire spells")
print(f"Title case: {title}")  // "The Ancient Book Of Fire Spells"
```

### `.reverse()` - Mirror Your Text Magic
```python
magical_code = String("ABRACADABRA")
reversed_code = magical_code.reverse()

print(f"Original: {magical_code.to_string()}")   // "ABRACADABRA"
print(f"Reversed: {reversed_code}")              // "ARBADACARBA"

// Check for palindromic spells
spell is_palindromic_spell(spell_text):
    spell_string = String(spell_text.lower())
    reversed_string = spell_string.reverse()
    return spell_string.to_string() == reversed_string

print(is_palindromic_spell("ABRACADABRA"))  // False
print(is_palindromic_spell("racecar"))      // True
print(is_palindromic_spell("A man a plan a canal Panama"))  // False (spaces matter)

// Create mirror spells
spell create_mirror_incantation(base_spell):
    base_string = String(base_spell)
    mirror = base_string.reverse()
    return f"{base_spell} {mirror}"

mirror_spell = create_mirror_incantation("MAGIC")
print(f"Mirror spell: {mirror_spell}")  // "MAGIC CIGAM"
```

---

## 🔧 String Building & Manipulation

### `.to_string()` - Access Raw Text
```python
enhanced_text = String("Magical Enhancement")
raw_text = enhanced_text.to_string()

print(f"Enhanced: {enhanced_text}")  // The String object
print(f"Raw: {raw_text}")           // "Magical Enhancement"

// Use for string concatenation and manipulation
spell combine_spell_components(*components):
    combined = ""
    for component in components:
        component_string = String(component)
        combined += component_string.to_string() + " + "
    
    // Remove the last " + "
    if len(combined) >= 3:
        combined = combined[:-3]
    
    return f"Spell formula: {combined}"

formula = combine_spell_components("Fire Essence", "Wind Spirit", "Earth Power")
print(formula)  // "Spell formula: Fire Essence + Wind Spirit + Earth Power"
```

---

## 🎪 Advanced String Magic

### Text Pattern Analysis
```python
spell analyze_magical_text(text):
    text_string = String(text)
    
    analysis = {
        "length": text_string.length(),
        "uppercase_count": 0,
        "lowercase_count": 0,
        "digit_count": 0,
        "space_count": 0,
        "vowel_count": 0,
        "consonant_count": 0
    }
    
    vowels = "aeiouAEIOU"
    
    for i in range(text_string.length()):
        char = text_string.char_at(i)
        
        if char.isupper():
            analysis["uppercase_count"] += 1
        elif char.islower():
            analysis["lowercase_count"] += 1
        elif char.isdigit():
            analysis["digit_count"] += 1
        elif char == " ":
            analysis["space_count"] += 1
        
        if char in vowels:
            analysis["vowel_count"] += 1
        elif char.isalpha():
            analysis["consonant_count"] += 1
    
    return analysis

spell_analysis = analyze_magical_text("The Quick Brown Fox Jumps Over 123 Lazy Dogs!")
print("Magical text analysis:")
for key, value in spell_analysis.items():
    print(f"  {key}: {value}")
```

### Magical Text Validation
```python
spell validate_spell_name(name):
    name_string = String(name)
    errors = []
    
    // Check length
    if name_string.length() < 3:
        errors.append("Spell name too short (minimum 3 characters)")
    elif name_string.length() > 50:
        errors.append("Spell name too long (maximum 50 characters)")
    
    // Check for forbidden characters
    forbidden = "!@#$%^&*()=+[]{}|\\:;\"'<>?,."
    for i in range(name_string.length()):
        char = name_string.char_at(i)
        if forbidden.find(char) != -1:
            errors.append(f"Forbidden character found: '{char}'")
            stop
    
    // Check that it starts with a letter
    if name_string.length() > 0:
        first_char = name_string.char_at(0)
        if not (first_char.upper() >= "A" and first_char.upper() <= "Z"):
            errors.append("Spell name must start with a letter")
    
    // Check for required elements
    if not (name_string.contains("Fire") or name_string.contains("Ice") or 
            name_string.contains("Lightning") or name_string.contains("Heal")):
        errors.append("Spell name must contain a magical element")
    
    return {
        "valid": len(errors) == 0,
        "errors": errors,
        "name": name_string.to_string()
    }

// Test spell name validation
test_names = ["Fire", "Lightning Bolt", "Mega Fire Blast", "Invalid@Spell", "A"]

for name in test_names:
    result = validate_spell_name(name)
    if result["valid"]:
        print(f"✅ '{result['name']}' is a valid spell name")
    else:
        print(f"❌ '{result['name']}' is invalid:")
        for error in result["errors"]:
            print(f"   - {error}")
```

### Magical Text Search & Replace
```python
spell magical_text_processor(text):
    text_string = String(text)
    
    spell find_all_occurrences(search_text):
        positions = []
        start_pos = 0
        
        while True:
            pos = text_string.find(search_text)
            if pos == -1:
                stop
            positions.append(pos)
            // Note: This is simplified - real implementation would need
            // to search from different starting positions
            stop
        
        return positions
    
    spell replace_magical_words():
        // Replace common magical terms with enhanced versions
        enhanced_text = text_string.to_string()
        
        replacements = {
            "magic": "MYSTICAL POWER",
            "spell": "ARCANE INCANTATION",
            "wand": "ENCHANTED STAFF",
            "potion": "MAGICAL ELIXIR"
        }
        
        for old_word, new_word in replacements.items():
            enhanced_text = enhanced_text.replace(old_word, new_word)
        
        return String(enhanced_text)
    
    spell count_magical_words():
        magical_words = ["magic", "spell", "enchant", "mystical", "arcane"]
        count = 0
        
        for word in magical_words:
            if text_string.contains(word):
                count += 1
        
        return count
    
    return {
        "find": find_all_occurrences,
        "enhance": replace_magical_words,
        "count_magic": count_magical_words,
        "original": text_string
    }

// Use the magical text processor
processor = magical_text_processor("This magic spell uses a wand to create a potion")
enhanced = processor["enhance"]()
magic_count = processor["count_magic"]()

print(f"Original: {processor['original'].to_string()}")
print(f"Enhanced: {enhanced.to_string()}")
print(f"Magical words found: {magic_count}")
```

---

## 🎯 Practical String Magic Examples

### Magical Password Generator
```python
spell generate_magical_password(length = 12, include_symbols = True):
    if length < 4:
        return "Password too short for magical security!"
    
    // Magical character sets
    lowercase = "abcdefghijklmnopqrstuvwxyz"
    uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numbers = "0123456789"
    magical_symbols = "*@#$%&!?"
    
    all_chars = lowercase + uppercase + numbers
    if include_symbols:
        all_chars += magical_symbols
    
    password_chars = []
    
    // Ensure at least one of each type
    password_chars.append(lowercase[random_index(len(lowercase))])
    password_chars.append(uppercase[random_index(len(uppercase))])
    password_chars.append(numbers[random_index(len(numbers))])
    
    if include_symbols:
        password_chars.append(magical_symbols[random_index(len(magical_symbols))])
    
    // Fill remaining length with random characters
    for i in range(len(password_chars), length):
        password_chars.append(all_chars[random_index(len(all_chars))])
    
    // Shuffle the password (simplified)
    // In real implementation, you'd want a proper shuffle
    password = "".join(password_chars)
    return String(password)

// Note: This is a simplified example
// Real implementation would need proper random number generation
```

### Magical Text Adventure Parser
```python
spell create_text_adventure_parser():
    spell parse_command(user_input):
        input_string = String(user_input.lower().strip())
        
        // Define command patterns
        move_words = ["go", "move", "walk", "run", "travel"]
        take_words = ["take", "get", "grab", "pick", "collect"]
        use_words = ["use", "cast", "activate", "invoke"]
        
        words = input_string.to_string().split(" ")
        if len(words) == 0:
            return {"action": "unknown", "object": "", "direction": ""}
        
        action_word = words[0]
        object_word = words[1] if len(words) > 1 else ""
        
        // Determine action type
        if action_word in move_words:
            return {
                "action": "move",
                "direction": object_word,
                "object": ""
            }
        elif action_word in take_words:
            return {
                "action": "take",
                "object": object_word,
                "direction": ""
            }
        elif action_word in use_words:
            return {
                "action": "use",
                "object": object_word,
                "direction": ""
            }
        else:
            return {
                "action": "unknown",
                "object": object_word,
                "direction": ""
            }
    
    spell format_response(action_data):
        action = action_data["action"]
        
        if action == "move":
            direction = action_data["direction"]
            if direction in ["north", "south", "east", "west"]:
                return f"🚶 Moving {direction}..."
            else:
                return "❓ Where do you want to go? (north/south/east/west)"
        
        elif action == "take":
            object_name = action_data["object"]
            if object_name:
                return f"👋 Attempting to take {object_name}..."
            else:
                return "❓ What do you want to take?"
        
        elif action == "use":
            object_name = action_data["object"]
            if object_name:
                return f"✨ Attempting to use {object_name}..."
            else:
                return "❓ What do you want to use?"
        
        else:
            return "❓ I don't understand that command. Try: go [direction], take [item], or use [item]"
    
    return {
        "parse": parse_command,
        "respond": format_response
    }

// Use the text adventure parser
parser = create_text_adventure_parser()

test_commands = [
    "go north",
    "take sword",
    "use magic wand",
    "dance wildly",
    "pick up treasure"
]

for command in test_commands:
    parsed = parser["parse"](command)
    response = parser["respond"](parsed)
    print(f"Command: '{command}' -> {response}")
```

---

## 🏆 String Magic Challenges

### 🥉 Beginner Challenges
1. **Spell Name Generator**: Create random magical spell names using string manipulation
2. **Text Encoder**: Create a simple magical cipher for encoding messages
3. **Word Counter**: Count specific words in magical texts

### 🥈 Intermediate Challenges
1. **Magical Anagram Detector**: Find if two spell names are anagrams
2. **Text Formatter**: Format magical texts with proper spacing and capitalization
3. **Spell Pattern Matcher**: Match text patterns for spell recognition

### 🥇 Advanced Challenges
1. **Magical Language Translator**: Create a fantasy language translator
2. **Advanced Text Parser**: Build a command parser for magical interfaces
3. **Spell Completion System**: Create autocomplete for magical spell names

---

## 🔗 Related Magic

Explore other grimoires that work well with strings:

- **[Array Grimoire](Array-Grimoire.md)** - For working with arrays of strings
- **[File & OS Grimoires](File-OS-Grimoires.md)** - For reading/writing text files
- **[Builtin Functions](../Reference/Builtin-Functions.md)** - Core functions for string operations
- **[Functions (Spells)](../Language-Fundamentals/Functions.md)** - Create custom string processing spells

---

## 💡 String Mastery Tips

1. **🔍 Use `.find()` for Positions**: Better than manual searching for substrings
2. **📐 Leverage Negative Indexing**: `.char_at(-1)` for last character access
3. **🎭 Normalize for Comparisons**: Use `.lower()` for case-insensitive operations
4. **🛡️ Validate Input**: Always check string properties before processing
5. **⚡ Chain String Operations**: Combine multiple transformations efficiently
6. **📊 Analyze Before Processing**: Use length and content checks for safety

---

*Master the String Grimoire, and you'll wield the power to manipulate text like ancient runic magic! Every character becomes a tool, every word a component in your magical arsenal. 🪄✨*

> "In the realm of text, every character holds power, and every string tells a story waiting to be transformed." - *Master String Enchanter*