# Carrion Language Changelog

## Version 0.1.6 - String Indexing & Standard Library Enhancement

### 🎉 Major Features

#### String Indexing Support
- **Implemented string indexing for primitive strings**
  - Supports positive indices: `s[0]`, `s[1]`, `s[6]`
  - Supports negative indices: `s[-1]`, `s[-2]` (Python-style)
  - Proper bounds checking with descriptive error messages
  - Returns single-character strings
  - **Location**: `src/evaluator/evaluator.go:1498-1616`

#### String Grimoire (Standard Library)
- **Created comprehensive String grimoire** in Munin standard library
  - `String(value)` - constructor for string objects
  - `length()` - get string length
  - `upper()` - convert to uppercase
  - `lower()` - convert to lowercase  
  - `reverse()` - reverse string order
  - `find(substring)` - find substring position (returns -1 if not found)
  - `contains(substring)` - check if string contains substring
  - `char_at(index)` - safe character access with bounds checking
  - **Location**: `src/munin/string.crl`

### 🔧 New Builtin Functions

#### Character/ASCII Functions
- **`ord(char)`** - Convert single character to ASCII code
  - Example: `ord("A")` returns `65`
  - **Location**: `src/evaluator/builtins.go:767-780`

- **`chr(code)`** - Convert ASCII code to character  
  - Example: `chr(65)` returns `"A"`
  - Supports range 0-255
  - **Location**: `src/evaluator/builtins.go:782-795`

### 🗂️ Project Organization

#### File Structure Improvements
- **Moved all test files to proper locations**
  - All `test_*.crl` files → `src/examples/` (66 files)
  - All debug files → `debug/` directory
  - Cleaned up root directory structure

#### Code Quality
- **Enhanced error handling** for string operations
- **Improved bounds checking** with clear error messages
- **Maintained backward compatibility** with existing code

### 🧪 Testing & Verification

#### Functionality Verified
- ✅ String indexing with positive indices
- ✅ String indexing with negative indices  
- ✅ Bounds checking and error handling
- ✅ Integration with existing recursion system
- ✅ String grimoire instantiation and basic operations
- ✅ New builtin functions (`ord`, `chr`)
- ✅ Existing functionality preserved

#### Example Usage
```carrion
// String indexing
s = "hello world"
print(s[0])    // "h"
print(s[6])    // "w" 
print(s[-1])   // "d"
print(s[-2])   // "l"

// String grimoire
sg = String("Hello World")
print(sg.length())           // 11
print(sg.upper())           // "HELLO WORLD"
print(sg.find("World"))     // 6
print(sg.contains("Hello")) // True

// New builtins
print(ord("A"))  // 65
print(chr(65))   // "A"
```

### 🔄 Recursive String Operations
- **Enhanced recursive function support** with string indexing
- **Example**: Recursive string reversal now possible
```carrion
spell reverse_string(s, index):
    if index < 0:
        return ""
    return s[index] + reverse_string(s, index - 1)

spell reverse(s):
    return reverse_string(s, len(s) - 1)

print(reverse("Carrion"))  // "noirraC"
```

### 🏗️ Technical Implementation

#### Core Changes
- **Modified `evalIndexExpression`** to handle `STRING_OBJ` type
- **Added `evalStringIndexExpression`** for string-specific indexing logic
- **Enhanced builtin function registry** with character conversion functions
- **Implemented String grimoire** following Carrion's grimoire patterns

#### Error Handling
- **Descriptive error messages** for out-of-bounds access
- **Type safety** for index operations (must be INTEGER)
- **Graceful handling** of negative indices
- **Clear stack traces** for debugging

### 📝 Notes
- **Backward Compatible**: All existing code continues to work
- **Performance**: String indexing is O(1) operation
- **Memory Safe**: Proper bounds checking prevents crashes
- **Consistent**: Follows Python-style negative indexing conventions

---

**Contributors**: Claude Code Assistant  
**Date**: June 7, 2025  
**Commit Range**: Latest development commits  
**Files Changed**: 2 core files, 1 new grimoire, 66+ files reorganized