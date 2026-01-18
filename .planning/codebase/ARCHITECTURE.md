# Architecture

**Analysis Date:** 2026-01-18

## Pattern Overview

**Overall:** Multi-tier Interpreter with Classic Compilation Pipeline

**Key Characteristics:**
- Tree-walking interpreter (Lexer → Parser → AST → Evaluator → Objects)
- Monolithic core with modular subsystems
- Companion tools as separate entry points (Sindri, Mimir, Bifrost)
- Indentation-sensitive parsing (Python-like syntax)

## Layers

**Token Layer:**
- Purpose: Define all language tokens and keywords
- Contains: Token type definitions (118 types), keyword lookup
- Location: `src/token/token.go`
- Depends on: Nothing
- Used by: Lexer

**Lexer Layer:**
- Purpose: Tokenization with indentation awareness
- Contains: String processing, indent/dedent handling, comment parsing
- Location: `src/lexer/lexer.go`
- Depends on: Token layer
- Used by: Parser

**Parser Layer:**
- Purpose: Build AST from token stream
- Contains: Recursive descent parser, expression precedence
- Location: `src/parser/parser.go`
- Depends on: Lexer, AST, Token layers
- Used by: Evaluator

**AST Layer:**
- Purpose: Abstract syntax tree node definitions
- Contains: Expression and statement nodes
- Location: `src/ast/ast.go`, `src/ast/expressions.go`, `src/ast/statements.go`
- Depends on: Token layer (for positions)
- Used by: Parser, Evaluator

**Object Layer:**
- Purpose: Runtime value representations
- Contains: 18 object types, environment scoping
- Location: `src/object/object.go`, `src/object/environment.go`
- Depends on: Nothing
- Used by: Evaluator, Builtins, Modules

**Evaluator Layer:**
- Purpose: AST interpretation and execution
- Contains: Tree-walking interpreter, call context, goroutine management
- Location: `src/evaluator/evaluator.go`
- Depends on: All other layers
- Used by: REPL, Main, Sindri

**Module Layer:**
- Purpose: Standard library implementations
- Contains: File, OS, HTTP, Socket, Time modules
- Location: `src/modules/*.go`
- Depends on: Object layer
- Used by: Evaluator (via builtins registration)

## Data Flow

**Script Execution:**

1. User runs: `carrion script.crl`
2. Main reads file content (`src/main.go`)
3. Lexer tokenizes source with indentation tracking
4. Parser builds AST from tokens
5. Evaluator traverses AST in shared Environment
6. Objects created/manipulated during evaluation
7. Results output via print() or return value
8. Process exits with status code

**REPL Mode:**

1. REPL starts with liner for line editing (`src/repl/repl.go`)
2. Loads Munin stdlib via LoadMuninStdlib()
3. For each input line:
   a. Lexer tokenizes input
   b. Parser builds AST
   c. Evaluator executes in persistent Environment
   d. Result displayed or error shown
4. History maintained across inputs

**State Management:**
- Global Environment persists across REPL inputs
- Function calls create new child Environments (closure support)
- Imported files cached in global `importedFiles` map
- Goroutines tracked via global GoroutineManager

## Key Abstractions

**Object:**
- Purpose: Runtime value representation
- Examples: `*object.Integer`, `*object.String`, `*object.Function`, `*object.Grimoire`
- Pattern: Interface with Type() and Inspect() methods
- Location: `src/object/object.go`

**Environment:**
- Purpose: Variable scoping and lookup
- Examples: Global env, function scope, grimoire instance env
- Pattern: Parent-child chain for closure support
- Location: `src/object/environment.go`

**CallContext:**
- Purpose: Track function call chain and recursion depth
- Examples: Used in Eval() to prevent stack overflow
- Pattern: Linked list of call frames
- Location: `src/evaluator/evaluator.go`

**Grimoire (Class):**
- Purpose: Object-oriented class definition
- Examples: User-defined classes, primitive wrappers
- Pattern: Constructor (init), instance methods, static methods, inheritance
- Location: `src/object/object.go`, evaluated in `src/evaluator/evaluator.go`

**Builtin:**
- Purpose: Go functions callable from Carrion
- Examples: print(), len(), type(), range()
- Pattern: Registered function map with variadic Object parameters
- Location: `src/evaluator/builtins.go`

## Entry Points

**Main Interpreter:**
- Location: `src/main.go`
- Triggers: `carrion <file.crl>` or `carrion` (REPL mode)
- Responsibilities: Parse flags, load file or start REPL, execute code

**REPL:**
- Location: `src/repl/repl.go`
- Triggers: Running carrion without file argument
- Responsibilities: Line editing, history, multi-line input, interactive evaluation

**Sindri (Test Framework):**
- Location: `cmd/sindri/main.go`
- Triggers: `sindri appraise [path]`
- Responsibilities: Discover tests, execute with timeout, generate reports

**Mimir (Documentation):**
- Location: `cmd/mimir/main.go`
- Triggers: `mimir` or `mimir interactive`
- Responsibilities: Function lookup, category browsing, help system

**Bifrost (Package Manager):**
- Location: `bifrost/cmd/bifrost/main.go`
- Triggers: `bifrost <command>`
- Responsibilities: Package installation, dependency resolution, registry access

## Error Handling

**Strategy:** Throw Error objects, catch at evaluation boundaries

**Patterns:**
- Parser collects syntax errors in errors slice
- Evaluator wraps runtime errors in `*object.Error`
- Enhanced errors include stack traces (`src/object/enhanced_errors.go`)
- Error suggestions provide fix hints (`src/utils/error_suggestions.go`)
- try/ensnare (attempt/resolve) for user-level exception handling

## Cross-Cutting Concerns

**Logging:**
- Console output via print() builtin
- Debug mode via environment variable and --idebug flag
- No structured logging framework

**Validation:**
- Type checking in builtin functions
- Parser validates syntax during AST construction
- Runtime type assertions in evaluator

**Concurrency:**
- Goroutine support via `diverge` keyword
- Global GoroutineManager tracks all spawned goroutines
- Environment uses sync.RWMutex for thread safety
- Cleanup via context channels on program exit

---

*Architecture analysis: 2026-01-18*
*Update when major patterns change*
