# Technology Stack

**Analysis Date:** 2026-01-18

## Languages

**Primary:**
- Go 1.23.0 - All interpreter and tooling implementation (`go.mod`)

**Secondary:**
- Carrion (.crl files) - The interpreted language being implemented

## Runtime

**Environment:**
- Go 1.23.0 with go1.24.2 toolchain
- Targets Linux, macOS, Windows
- Docker containerized builds supported (`Dockerfile`)

**Package Manager:**
- Go modules
- Lockfiles: `go.sum`, `bifrost/go.sum`
- Custom package manager: Bifrost (`bifrost/Bifrost.toml`)

## Frameworks

**Core:**
- None (vanilla Go CLI application)

**Testing:**
- Go built-in `testing` package
- Custom test framework: Sindri (`cmd/sindri/main.go`)

**Build/Dev:**
- Make for build orchestration (`Makefile`, `bifrost/Makefile`)
- Docker for containerized builds (`Dockerfile`)

## Key Dependencies

**Critical:**
- `github.com/peterh/liner` - Interactive REPL line editing and history (`go.mod`)
- `github.com/spf13/cobra` - CLI framework for Bifrost package manager (`bifrost/go.mod`)
- `github.com/BurntSushi/toml` - TOML config parsing for Bifrost (`bifrost/go.mod`)

**Infrastructure:**
- Go standard library - fs, path, net/http, encoding/json for core functionality
- `golang.org/x/term` - Terminal handling for Bifrost (`bifrost/go.mod`)

## Configuration

**Environment:**
- `.env` files supported (`.gitignore` excludes `.env`)
- `CARRION_DEBUG_WRAPPING` environment variable for debug mode
- Command-line flags: `--version`, `-v`, `--idebug`, `--lexer`, `--parser`, `--evaluator`, `--all`

**Build:**
- `go.mod` - Go module configuration
- `Makefile` - Build targets with OS auto-detection
- `Dockerfile` - Container build specification

## Platform Requirements

**Development:**
- Any platform with Go 1.23+ installed
- Make for build automation
- Docker optional for containerized builds

**Production:**
- Compiled binary distribution
- Targets: Linux (amd64, arm), macOS (darwin), Windows
- Installation via `make install` or Docker

---

*Stack analysis: 2026-01-18*
*Update after major dependency changes*
