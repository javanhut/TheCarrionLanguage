# Codebase Structure

**Analysis Date:** 2026-01-18

## Directory Layout

```
TheCarrionLanguage/
├── cmd/                    # Companion tool entry points
│   ├── sindri/            # Testing framework
│   └── mimir/             # Documentation tool
├── src/                   # Core language interpreter
│   ├── ast/              # Abstract syntax tree
│   ├── debug/            # Debug configuration
│   ├── evaluator/        # Runtime evaluation
│   ├── lexer/            # Tokenization
│   ├── modules/          # Standard library modules
│   ├── munin/            # Embedded stdlib
│   ├── object/           # Object system
│   ├── parser/           # Parsing
│   ├── repl/             # Interactive shell
│   ├── tests/            # Integration tests
│   ├── token/            # Token definitions
│   ├── utils/            # Utility functions
│   └── main.go           # Primary entry point
├── bifrost/              # Package manager (submodule)
│   ├── cmd/bifrost/     # Entry point
│   ├── internal/        # Implementation
│   └── docs/            # Documentation
├── docker/              # Docker build support
├── docs/                # Language documentation
├── examples/            # Example programs
├── install/             # Installation scripts
├── Changelog/           # Version history
├── Makefile             # Build targets
├── Dockerfile           # Container spec
└── go.mod               # Go module config
```

## Directory Purposes

**cmd/sindri/**
- Purpose: Test framework for Carrion programs
- Contains: `main.go` with test discovery, execution, HTML report generation
- Key files: `main.go` (single file tool)
- Subdirectories: None

**cmd/mimir/**
- Purpose: Interactive documentation and help system
- Contains: `main.go` with function lookup, category browsing
- Key files: `main.go` (single file tool)
- Subdirectories: None

**src/ast/**
- Purpose: Abstract syntax tree node definitions
- Contains: Expression and statement types
- Key files: `ast.go`, `expressions.go`, `statements.go`, `ast_test.go`
- Subdirectories: None

**src/evaluator/**
- Purpose: Core interpreter runtime
- Contains: Tree-walking evaluator, builtins, stdlib loading
- Key files: `evaluator.go` (5,257 lines), `builtins.go`, `stdlib.go`, `evaluator_test.go`
- Subdirectories: None

**src/lexer/**
- Purpose: Source code tokenization
- Contains: Indentation-aware lexer
- Key files: `lexer.go`, `lexer_test.go`
- Subdirectories: None

**src/modules/**
- Purpose: Standard library module implementations
- Contains: File, OS, HTTP, Socket, Time modules
- Key files: `file.go`, `os.go`, `httpmodule.go`, `sockets.go`, `timemodule.go`
- Subdirectories: None

**src/munin/**
- Purpose: Embedded standard library grimoires
- Contains: Stdlib registration and embedding
- Key files: `embed.go`, `munin.go`, `wrapper_grimoire.go`
- Subdirectories: `stdlib/` (embedded .crl files)

**src/object/**
- Purpose: Runtime object type definitions
- Contains: 18 object types, environment, error handling
- Key files: `object.go`, `environment.go`, `enhanced_errors.go`, `error_handling.go`
- Subdirectories: None

**src/parser/**
- Purpose: AST construction from tokens
- Contains: Recursive descent parser
- Key files: `parser.go` (2,626 lines), `parse_test.go`
- Subdirectories: None

**src/repl/**
- Purpose: Interactive read-eval-print loop
- Contains: Line editing, history, debug support
- Key files: `repl.go`
- Subdirectories: None

**bifrost/internal/**
- Purpose: Package manager implementation (private)
- Contains: Auth, config, install, manifest, registry, resolver, version
- Key files: Each subdirectory has `*.go` files
- Subdirectories: `auth/`, `config/`, `install/`, `manifest/`, `registry/`, `resolver/`, `version/`, `archive/`, `integration/`, `uninstall/`

## Key File Locations

**Entry Points:**
- `src/main.go` - Main interpreter entry
- `cmd/sindri/main.go` - Test framework entry
- `cmd/mimir/main.go` - Documentation tool entry
- `bifrost/cmd/bifrost/main.go` - Package manager entry

**Configuration:**
- `go.mod` - Go module configuration
- `bifrost/go.mod` - Bifrost Go module
- `Makefile` - Build automation
- `Dockerfile` - Container build spec

**Core Logic:**
- `src/evaluator/evaluator.go` - Main interpreter
- `src/parser/parser.go` - Language parser
- `src/lexer/lexer.go` - Tokenizer
- `src/object/object.go` - Type system

**Testing:**
- `src/tests/integration_test.go` - Integration tests
- `src/evaluator/evaluator_test.go` - Evaluator tests
- `src/parser/parse_test.go` - Parser tests
- `src/lexer/lexer_test.go` - Lexer tests

**Documentation:**
- `docs/DOCUMENTATION.md` - Complete language documentation
- `docs/README.md` - Documentation index
- `bifrost/docs/ARCHITECTURE.md` - Bifrost architecture

## Naming Conventions

**Files:**
- `snake_case.go` - Go source files (e.g., `error_handling.go`, `enhanced_errors.go`)
- `*_test.go` - Test files (e.g., `evaluator_test.go`)
- `*.crl` - Carrion source files
- `UPPERCASE.md` - Important docs (e.g., `README.md`, `DOCUMENTATION.md`)

**Directories:**
- `lowercase` - All directories use lowercase
- Plural for collections (e.g., `modules/`, `tests/`, `examples/`)
- `internal/` - Private packages (Bifrost)

**Special Patterns:**
- `main.go` - Entry point in each cmd directory
- `*module.go` - Module implementations (e.g., `httpmodule.go`, `timemodule.go`)
- `embed.go` - Go embed directives

## Where to Add New Code

**New Language Feature:**
- Token definition: `src/token/token.go`
- Lexer handling: `src/lexer/lexer.go`
- AST node: `src/ast/expressions.go` or `src/ast/statements.go`
- Evaluation: `src/evaluator/evaluator.go`
- Tests: Corresponding `*_test.go` file

**New Builtin Function:**
- Implementation: `src/evaluator/builtins.go`
- Tests: `src/evaluator/evaluator_test.go`

**New Module:**
- Implementation: `src/modules/<name>.go`
- Registration: `src/evaluator/builtins.go` (add to builtins map)

**New Standard Library Grimoire:**
- Implementation: `src/munin/stdlib/<name>.crl`
- Registration: `src/munin/munin.go`

**New Companion Tool:**
- Entry point: `cmd/<tool>/main.go`
- Build target: Update `Makefile`

**Bifrost Feature:**
- Internal code: `bifrost/internal/<domain>/<name>.go`
- CLI command: `bifrost/cmd/bifrost/main.go`

## Special Directories

**src/munin/stdlib/**
- Purpose: Embedded standard library grimoires
- Source: Carrion source files compiled into binary
- Committed: Yes (source of truth for stdlib)

**examples/**
- Purpose: Example Carrion programs
- Source: Demonstration and testing files
- Committed: Yes

**docker/**
- Purpose: Docker-related build files
- Source: Build support
- Committed: Yes

---

*Structure analysis: 2026-01-18*
*Update when directory structure changes*
