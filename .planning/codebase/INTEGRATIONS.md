# External Integrations

**Analysis Date:** 2026-01-18

## APIs & External Services

**HTTP Client Integration:**
- Built-in HTTP module - `src/modules/httpmodule.go`
  - Supports GET, POST, PUT, DELETE, HEAD requests
  - Custom timeout handling (30-second default)
  - JSON parsing and stringification
  - Query string building

**External APIs:**
- None required - language interpreter is self-contained

## Data Storage

**Databases:**
- Not applicable - interpreter has no database dependencies

**File Storage:**
- Local file system via File module - `src/modules/file.go`
  - File open/close operations
  - Read/write operations with handle management
  - Directory operations via OS module

**Caching:**
- Package cache: `~/.carrion/cache/` (Bifrost package manager)
- Registry metadata cache: `~/.carrion/registry/`

## Authentication & Identity

**Auth Provider:**
- Bifrost package registry authentication - `bifrost/internal/auth/auth.go`
  - API key management
  - Credentials stored in user config

**OAuth Integrations:**
- Not detected

## Monitoring & Observability

**Error Tracking:**
- Enhanced error system - `src/object/enhanced_errors.go`
- Stack trace generation - `src/object/error_trace.go`
- Error suggestions - `src/utils/error_suggestions.go`

**Analytics:**
- Not detected

**Logs:**
- Console output via print() builtin
- Debug mode via CARRION_DEBUG_WRAPPING env var

## CI/CD & Deployment

**Hosting:**
- Compiled binary distribution
- Docker container support - `Dockerfile`

**CI Pipeline:**
- Not detected (no .github/workflows or similar)

## Environment Configuration

**Development:**
- Required: Go 1.23+ installation
- Optional: Docker for containerized builds
- Optional: Make for build automation

**Package Locations:**
- User packages: `~/.carrion/packages/`
- Global packages: `/usr/local/share/carrion/lib/`
- Project packages: `./carrion_modules/`

**Production:**
- Compiled binary with no runtime dependencies
- Standard library (Munin) embedded via `src/munin/embed.go`

## Webhooks & Callbacks

**Incoming:**
- HTTP server support via sockets module - `src/modules/sockets.go`
  - TCP/UDP socket handling
  - WebSocket support
  - HTTP request parsing and response building

**Outgoing:**
- HTTP client for external API calls - `src/modules/httpmodule.go`

## Network & Socket Integration

**TCP/UDP Sockets:**
- Raw socket support - `src/modules/sockets.go`
  - TCP client/server
  - UDP client/server
  - Unix domain sockets
  - Port allocation with retry logic

**Web Server:**
- Built-in HTTP server capability
- Static file serving
- Route handling via Carrion callbacks

## Operating System Integration

**OS Module:**
- `src/modules/os.go`
  - Command execution (osRunCommand)
  - Environment variables (osGetEnv, osSetEnv)
  - Working directory management (osGetCwd, osChdir)
  - Directory listing (osListDir)
  - File operations (osRemove, osMkdir)

## Time & Scheduling

**Time Module:**
- `src/modules/timemodule.go`
  - Current time (Unix timestamps)
  - Time parsing and formatting
  - Duration calculations
  - Timezone handling
  - Sleep functionality

---

*Integration audit: 2026-01-18*
*Update when adding/removing external services*
