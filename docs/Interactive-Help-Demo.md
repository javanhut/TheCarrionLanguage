# Carrion Interactive Help System Demo

## How to Use the Interactive Help

### 1. Start the Carrion REPL
```bash
./src/src
```

### 2. You'll see the enhanced welcome message:
```
🦅 Welcome to the Carrion Programming Language REPL! 🦅

📖 Quick Help:
   • Type 'mimir' for comprehensive function help
   • Type 'help()' for basic information
   • Type 'version()' to see current version
   • Type 'modules()' to list standard library modules

⚡ REPL Commands:
   • 'clear' - Clear screen
   • 'quit', 'exit', 'q', 'qa' - Exit REPL
   • Use Tab for auto-completion

✨ Quick Examples:
   print("Hello, World!")     // Basic output
   x, y = (10, 20)            // Tuple unpacking
   42.to_bin()                // "0b101010"
   "hello".upper()            // "HELLO"

May Mimir guide your coding journey! Type commands below:
>>>
```

### 3. Type 'mimir' to enter interactive help:
```
>>> mimir
```

### 4. Interactive Help Main Menu:
```
═══════════════════════════════════════════════════════════════════
🧙‍♂️ Welcome to Mimir's Interactive Help System 🧙‍♂️
═══════════════════════════════════════════════════════════════════

📚 What would you like help with?

  1️⃣  Built-in Functions    - Core language functions (print, len, type, etc.)
  2️⃣  Standard Library      - Munin modules (Array, String, File, OS, etc.)
  3️⃣  Language Features     - Syntax, control flow, OOP, error handling
  4️⃣  Examples & Demos      - Working code examples and tutorials
  5️⃣  Search Functions      - Find specific functions by name or purpose
  6️⃣  Tips & Tricks         - REPL shortcuts and advanced features

💡 Commands: Type number/name (e.g., '1' or 'builtins'), 'h' for menu, 'q' to quit
🔍 Quick search: Type any function name directly (e.g., 'print', 'Array')

help>
```

## Navigation Examples

### Example 1: Exploring Built-in Functions
```
help> 1
# Shows built-in functions menu with categories

builtins> 1
# Shows type conversion functions with examples

# Press Enter to continue, then 'b' to go back
builtins> b
help>
```

### Example 2: Learning Standard Library
```
help> 2
# Shows standard library modules

stdlib> array
# Shows comprehensive Array module documentation

# Press Enter to continue
stdlib> string
# Shows String module with all methods

stdlib> b
help>
```

### Example 3: Quick Function Search
```
help> print
# Directly shows print function documentation

help> Array
# Directly shows Array module documentation

help> 5
# Enter search mode
search> array sort
# Shows all functions related to array sorting
```

### Example 4: Code Examples
```
help> 4
# Shows examples menu

examples> 1
# Shows Hello World and basic examples

examples> 6
# Shows object-oriented programming examples

examples> b
help>
```

### Example 5: Language Features
```
help> 3
# Shows language features menu

syntax> 2
# Shows control flow documentation

syntax> functions
# Shows function/spell documentation

syntax> b
help>
```

## Interactive Features

### 🔍 Search Functionality
- Type function names directly: `print`, `Array`, `String`
- Search by keywords: `array sort`, `string upper`, `file read`
- Browse by categories
- Get suggestions for misspelled terms

### 📚 Hierarchical Navigation
- Main menu → Categories → Specific topics
- Use numbers or names to navigate
- 'b' to go back, 'q' to quit help system
- 'h' to show menu again

### 💡 Smart Help
- Auto-completion in REPL
- Function signatures with examples
- Copy-paste ready code examples
- Tips and best practices

### ⚡ Quick Access
- Direct function lookup
- Tab completion for help commands
- History navigation with arrow keys
- Clear explanations with visual formatting

## Sample Session Walkthrough

```
>>> mimir
help> 1                    # Built-in functions
builtins> 2               # Utility functions  
# Shows len, type, print, input, range with examples
builtins> print          # Quick lookup for print function
builtins> b              # Back to help menu

help> search             # Function search
search> array           # Search for array functions
# Shows Array module and related functions
search> b               # Back to help menu

help> 4                 # Examples
examples> 2             # Array examples
# Shows complete array manipulation examples
examples> b             # Back to help menu

help> q                 # Quit help system
>>> print("Hello!")     # Back to normal REPL
Hello!
>>> quit               # Exit REPL
```

## Benefits of Interactive Help

✅ **Discoverable**: Easy to find functions and features  
✅ **Organized**: Logical hierarchy and categories  
✅ **Searchable**: Find functions by name or purpose  
✅ **Practical**: Working examples you can copy-paste  
✅ **Complete**: Covers all language features  
✅ **Interactive**: Navigate at your own pace  
✅ **Contextual**: Get help while coding in REPL  

The interactive help system makes learning and using Carrion much more efficient and enjoyable!