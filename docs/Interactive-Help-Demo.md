# Carrion Interactive Help System Demo

## How to Use the Interactive Help

### 1. Start the Carrion REPL
```bash
carrion
```

### 2. You'll see the enhanced welcome message:
```
Welcome to the Carrion Programming Language REPL!

Quick Help:
   • Run 'mimir' for comprehensive interactive documentation
   • Run 'mimir scry <function>' for specific function help
   • Type 'help()' for basic information
   • Type 'version()' to see current version
   • Type 'modules()' to list standard library modules

REPL Commands:
   • 'clear' - Clear screen
   • 'quit', 'exit', 'q', 'qa' - Exit REPL
   • Use Tab for auto-completion

Quick Examples:
   print("Hello, World!")     // Basic output
   x, y = (10, 20)            // Tuple unpacking
   x = 42
	x.to_bin()                // "0b101010"
   "hello".upper()            // "HELLO"

May Mimir guide your coding journey! Type commands below:
>>>
```

### 3. Use the external Mimir tool for comprehensive help:
```bash
# In a separate terminal or after exiting REPL
mimir
```

### 4. Or get specific function help:
```bash
mimir scry print
mimir scry Array
```

### 5. Mimir Interactive Help Main Menu:
```
    MIMIR - The All-Seeing Helper
   ═══════════════════════════════════════
    Interactive Carrion Documentation
   ═══════════════════════════════════════
   "Knowledge is the greatest treasure"

What knowledge do you seek?

  1.  Built-in Functions    - Core language functions (print, len, type, etc.)
  2.  Standard Library      - Munin modules (Array, String, File, OS, etc.)
  3.  Language Features     - Syntax, control flow, OOP, error handling
  4.  Examples & Demos      - Working code examples and tutorials
  5.  Search Functions      - Find specific functions by name or purpose
  6.  Tips & Tricks         - REPL shortcuts and advanced features

Commands: Type number/name (e.g., '1' or 'builtins'), 'h' for menu, 'q' to quit
Quick search: Type any function name directly (e.g., 'print', 'Array')

mimir>
```

## Navigation Examples

### Example 1: Exploring Built-in Functions
```
mimir> 1
# Shows built-in functions menu with categories

builtins> 1
# Shows type conversion functions with examples

# Press Enter to continue, then 'b' to go back
builtins> b
mimir>
```

### Example 2: Learning Standard Library
```
mimir> 2
# Shows standard library modules

stdlib> array
# Shows comprehensive Array module documentation

# Press Enter to continue
stdlib> string
# Shows String module with all methods

stdlib> b
mimir>
```

### Example 3: Quick Function Search
```
mimir> print
# Directly shows print function documentation

mimir> Array
# Directly shows Array module documentation

mimir> 5
# Enter search mode
search> array sort
# Shows all functions related to array sorting
```

### Example 4: Code Examples
```
mimir> 4
# Shows examples menu

examples> 1
# Shows Hello World and basic examples

examples> 6
# Shows object-oriented programming examples

examples> b
mimir>
```

### Example 5: Language Features
```
mimir> 3
# Shows language features menu

syntax> 2
# Shows control flow documentation

syntax> functions
# Shows function/spell documentation

syntax> b
mimir>
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

```bash
# Start Mimir documentation tool
$ mimir

mimir> 1                    # Built-in functions
builtins> 2                # Utility functions  
# Shows len, type, print, input, range with examples
builtins> print           # Quick lookup for print function
builtins> b               # Back to help menu

mimir> search             # Function search
search> array            # Search for array functions
# Shows Array module and related functions
search> b                # Back to help menu

mimir> 4                 # Examples
examples> 2              # Array examples
# Shows complete array manipulation examples
examples> b              # Back to help menu

mimir> q                 # Quit help system

# Or use command-line for quick lookup
$ mimir scry print
# Shows print function documentation

# Then use the REPL for coding
$ carrion
>>> print("Hello!")      # Use what you learned
Hello!
>>> quit                # Exit REPL
```

## Benefits of Interactive Help

✅ **Discoverable**: Easy to find functions and features  
✅ **Organized**: Logical hierarchy and categories  
✅ **Searchable**: Find functions by name or purpose  
✅ **Practical**: Working examples you can copy-paste  
✅ **Complete**: Covers all language features  
✅ **Interactive**: Navigate at your own pace  
✅ **Contextual**: Get help while coding in REPL  

The Mimir documentation system and REPL integration make learning and using Carrion much more efficient and enjoyable!