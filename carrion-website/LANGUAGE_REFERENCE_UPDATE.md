# Language Reference Documentation Update

## Summary

Updated the Carrion website's Language Reference page from a placeholder to a comprehensive, fully-documented reference guide based on the `repl_fix` branch of TheCarrionLanguage repository.

## File Updated

- **File**: `src/pages/docs/LanguageReference.tsx`
- **Lines**: 1,611 lines of code
- **Type**: React TypeScript component with styled-components

## Features Documented

### 1. Language Overview
- Key characteristics (dynamic typing, interpreted, OOP, concurrent, duck typing)
- File extension (.crl)
- Philosophy and naming inspiration

### 2. Lexical Structure
- Comments (single-line `//`, multi-line `/* */`)
- Identifiers and naming rules
- Complete keyword table organized by category
- All literal types (integer, float, string, boolean, None)

### 3. Data Types
Complete coverage of:
- **Primitives**: Integer, Float, String, Boolean, None
- **Collections**: Array, Hash, Tuple
- Type checking with `type()` function
- String indexing support

### 4. Operators
Full operator reference tables:
- Arithmetic (`+`, `-`, `*`, `/`, `//`, `%`, `**`)
- Assignment (`=`, `+=`, `-=`, `*=`, `/=`, `++`, `--`)
- Comparison (`==`, `!=`, `<`, `>`, `<=`, `>=`)
- Logical (`and`, `or`, `not`)
- Membership (`in`, `not in`)
- Bitwise (`&`, `|`, `^`, `~`, `<<`, `>>`)

### 5. Control Flow
- If/otherwise/else statements
- For loops with examples
- While loops
- Match/case pattern matching
- Loop control (skip/stop)

### 6. Functions (Spells)
- Function definition with `spell` keyword
- Default parameters
- Recursion examples
- **Important**: Explicitly notes NO lambda/anonymous functions

### 7. OOP (Grimoires)
- Class definition with `grim` keyword
- Constructors with `init`
- Inheritance with `super`
- Abstract classes with `arcane` and `@arcanespell`
- Encapsulation (private `_`, protected `__`)
- Polymorphism examples

### 8. Error Handling
- attempt/ensnare/resolve blocks
- Raising errors with `raise`
- Assertions with `check()`
- Custom error classes

### 9. Type Hints (NEW)
- Optional, non-enforcing type annotations
- Parameter type hints: `spell add(a: int, b: int):`
- Return type hints: `-> int`
- Variable type hints: `count: int = 0`
- Supported types: int, float, str, bool, list, dict, set, None, any, custom grimoires
- Clear warning that they are documentation-only

### 10. Concurrency (NEW)
- **diverge** keyword for creating goroutines
  - Anonymous: `diverge: <code>`
  - Named: `diverge name: <code>`
- **converge** keyword for waiting
  - All: `converge`
  - Specific: `converge name1, name2`
- Examples: basic concurrency, producer-consumer, parallel computation
- Features: thread-safe, isolated environments, automatic cleanup

### 11. Main Entry Point (NEW)
- **main:** keyword for program entry point
- How it affects top-level code execution
- Examples with and without concurrency
- When to use main:

### 12. Modules and Imports
- Grimoire-based imports
- Local file imports
- Package imports with version resolution
- Relative imports
- Selective imports (grimoires and spells)
- Import resolution order
- Complete examples

### 13. Built-in Functions
Comprehensive tables:
- Type conversion (int, float, str, bool, list, tuple)
- Utility functions (len, type, print, input, range, max, abs, ord, chr)
- Collection functions (enumerate, pairs, is_sametype)
- Meta functions (help, version, modules, mimir)

### 14. Standard Library (Munin)
Complete module listing:

**Core Type Modules (8):**
- Array, String, Integer, Float, Boolean, Primitive, Math, Time

**I/O and System (2):**
- File, OS

**Networking (2):**
- ApiRequest (HTTP client)
- Servers (TCP, UDP, Unix, HTTP, WebServer)

**Data Structures (2):**
- DataStructures (Stack, Queue, Heap, BTree with iterators)
- Iterable (abstract base)

**Utilities (2):**
- Debug, BuiltinErrors

## Visual Features

### Styled Components
- Consistent styling with other documentation pages
- Responsive design
- Syntax highlighting for all code examples
- Color-coded tables
- Info boxes and warning boxes
- Sticky navigation for quick section access

### Interactive Elements
- Quick navigation bar with 14 section links
- Smooth scrolling to sections
- Hover effects on links and cards
- Organized tables for easy reference

### Code Examples
- Over 50 code examples throughout
- Syntax highlighting using react-syntax-highlighter
- Real-world usage patterns
- Clear, commented examples

## Key Differences from Placeholder

**Before**: 
- 12 lines
- Just a title and placeholder text
- No actual documentation

**After**:
- 1,611 lines
- Comprehensive reference covering all language features
- 14 major sections
- 50+ code examples
- Multiple reference tables
- Styled, interactive, and professional

## Important Notes

### Features NOT Included (as they don't exist in Carrion)
- Lambda/anonymous functions (explicitly removed in repl_fix)
- List comprehensions (marked as future feature)
- Generator expressions
- Async/await

### New Features from repl_fix Branch
- Type hints (optional, non-enforcing)
- Concurrency (diverge/converge)
- Main entry point (main:)
- Enhanced module system
- Expanded standard library (16+ modules)

## Testing

The component is ready for:
1. React build compilation
2. Integration with existing routing
3. Deployment to GitHub Pages

## Next Steps

1. Install dependencies: `npm install`
2. Test build: `npm run build`
3. Test locally: `npm start`
4. Commit changes
5. Deploy to GitHub Pages

## Resources Referenced

- `/docs/Language-Reference.md` from main branch
- `/docs/CARRION.md` from repl_fix branch
- `/docs/Concurrency.md` from repl_fix branch
- `/docs/Type-System.md` from repl_fix branch
- `/docs/Modules.md` from repl_fix branch
- `/docs/Standard-Library.md` from repl_fix branch
- `/src/munin/` directory for module verification

## Maintainability

The documentation is:
- Well-structured with clear section IDs
- Easy to update with new features
- Follows existing styling patterns
- Uses reusable styled components
- Properly typed with TypeScript
