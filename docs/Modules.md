# Module System and Imports

Carrion provides a comprehensive module system that allows you to organize code across multiple files and import functionality between them. This enables code reuse, better organization, and collaborative development. The system supports local file imports, project packages, user-specific packages, and system-wide global packages managed by Bifrost package manager.

## Import Resolution System

Carrion uses a smart import resolution system that automatically determines the import type and searches appropriate locations, making imports as simple as possible while maintaining full flexibility.

### Smart Resolution Logic

The import system analyzes your import string and chooses the appropriate resolution strategy:

1. **Relative Imports** (starts with `./` or `../`) → Resolved relative to current directory
2. **Local Files** (simple names like `"filename"`) → Current directory first, then packages
3. **Package Imports** (with `/` like `"package/module"`) → Package directories with version resolution
4. **Explicit Paths** (full paths) → Direct path resolution for backward compatibility

### Search Path Priority

When you import a module, Carrion searches in this order:

1. **Current Directory** - Local files relative to the current working directory
2. **Project Modules** - `./carrion_modules/package/[version]/src/` for project packages
3. **Global Bifrost Modules** - `/usr/bin/carrion_modules/package/[version]/src/` for system Bifrost packages
4. **User Packages** - `~/.carrion/packages/package/[version]/src/` for user-installed packages  
5. **Global Packages** - `/usr/local/share/carrion/lib/package/[version]/src/` for system packages
6. **Standard Library** - Built-in Munin standard library modules

### Automatic Version Resolution

For package imports, Carrion automatically:
- Finds the latest version of a package
- Looks in the `src/` directory within the version folder
- Supports both versioned (`package/1.0.0/src/`) and unversioned structures

### Basic Import Syntax

Carrion provides multiple import patterns designed for ease of use and flexibility:

#### Grimoire-Based Imports (New)
```python
import "GrimoireName"                // Search for grimoire by name in available modules
import "GrimoireName" as MyGrimoire  // Import grimoire with alias
```

#### Local File Imports
```python
import "filename"              // Current directory: ./filename.crl
import "mymodule.ClassName"    // Selective import: ./mymodule.crl -> ClassName
```

#### Simplified Package Imports
```python
import "package/module"              // Auto-resolves to: carrion_modules/package/[version]/src/module.crl
import "package/module.ClassName"    // Selective import from package module
```

#### Relative Path Imports
```python
import "./filename"            // Explicit current directory
import "../parent/module"      // Relative parent directory
import "../../utils/helper"    // Multi-level relative paths
```

#### Legacy Full Path Imports (Backward Compatible)
```python
import "carrion_modules/package/1.0.0/src/module"  // Full explicit path
```

## Grimoire-Based Import System

Carrion's enhanced import system allows you to import grimoires (classes) directly by name without specifying the file path. This provides a more intuitive way to work with classes across your project and installed packages.

### How Grimoire Imports Work

When you use `import "GrimoireName"`, Carrion:

1. **Searches Multiple Locations**: Looks through all available .crl files in the search path
2. **Finds the Grimoire**: Locates the grimoire definition in any available module
3. **Imports Directly**: Makes the grimoire available without importing the entire module
4. **Supports Aliases**: Allows you to rename the grimoire for convenience

### Grimoire Import Examples

#### Basic Grimoire Import

**File: `lib.crl`**
```python
grim Calculator:
    init():
        self.result = 0
    
    spell add(a, b):
        return a + b
    
    spell multiply(a, b):
        return a * b

grim Logger:
    init(name):
        self.name = name
    
    spell log(message):
        print(f"[{self.name}] {message}")
```

**File: `main.crl`**
```python
// Import grimoires directly by name
import "Calculator"
import "Logger" as Log

// Use imported grimoires
calc = Calculator()
result = calc.add(5, 3)
print(f"Result: {result}")

logger = Log("MyApp")
logger.log("Application started")
```

#### Grimoire Imports from Bifrost Packages

**Package Structure:**
```
carrion_modules/
└── hello-world/
    └── 0.1.0/
        └── src/
            └── main.crl
```

**File: `carrion_modules/hello-world/0.1.0/src/main.crl`**
```python
grim HelloWorld:
    init(name):
        self.name = name
    
    spell greet():
        return f"Hello, {self.name}!"
```

**File: `main.crl`**
```python
// Import grimoire from bifrost package
import "HelloWorld" as Hello

// Use the imported grimoire
greeter = Hello("World")
print(greeter.greet())  // Output: Hello, World!
```

#### Global Grimoire Imports

**Global Package Structure:**
```
/usr/bin/carrion_modules/
└── json-parser/
    └── 1.0.0/
        └── src/
            └── main.crl
```

**File: `/usr/bin/carrion_modules/json-parser/1.0.0/src/main.crl`**
```python
grim JSONParser:
    init():
        self.data = {}
    
    spell parse(json_string):
        // JSON parsing logic
        return {"parsed": True}
```

**File: `main.crl`**
```python
// Import from global bifrost modules
import "JSONParser" as JSON

parser = JSON()
result = parser.parse('{"name": "example"}')
print(result)
```

### Search Order for Grimoire Imports

When importing a grimoire by name, Carrion searches in this order:

1. **Current Directory** - Local .crl files
2. **Project Modules** - `./carrion_modules/*/version/src/main.crl`
3. **Global Bifrost Modules** - `/usr/bin/carrion_modules/*/version/src/main.crl`
4. **User Packages** - `~/.carrion/packages/*/version/src/main.crl`
5. **Global Packages** - `/usr/local/share/carrion/lib/*/version/src/main.crl`

### Practical Usage Patterns

#### Mixed Import Strategies

```python
// Mix grimoire imports with traditional imports
import "Calculator"              // Grimoire import
import "utils/helper"            // Package import  
import "./config"                // Relative import
import "Logger" as AppLogger     // Grimoire import with alias

// Use all imported functionality
calc = Calculator()
helper = HelperClass()
config = load_config()
logger = AppLogger("Main")

result = calc.add(10, 20)
helper.process_data(result)
logger.log(f"Calculated: {result}")
```

#### Conditional Grimoire Imports

```python
// Import different grimoires based on conditions
debug_mode = True

if debug_mode:
    import "DebugLogger" as Logger
else:
    import "ProductionLogger" as Logger

// Same interface, different implementations
logger = Logger("App")
logger.log("Starting application")
```

#### Fallback Grimoire Imports

```python
// Try to import preferred grimoire, fallback to basic one
spell get_database():
    attempt:
        import "AdvancedDatabase" as DB
        return DB()
    ensnare:
        import "BasicDatabase" as DB
        return DB()

database = get_database()
```

## Smart Import Examples

### Example Project Structure
```
my_project/
├── main.crl
├── utils.crl
├── models/
│   ├── user.crl
│   └── product.crl
├── carrion_modules/
│   └── hello-world/
│       └── 0.1.0/
│           └── src/
│               └── main.crl
└── ../shared/
    └── common.crl
```

### All Import Pattern Examples

```python
// 1. Grimoire-based imports (NEW - search by class name)
import "Helper"                   // → Search for Helper grimoire in all locations
import "HelloWorld" as Hello      // → Search for HelloWorld grimoire with alias
import "Logger" as Log            // → Search for Logger grimoire with alias

// 2. Local file imports (current directory)
import "utils"                    // → ./utils.crl
import "utils.Helper"             // → ./utils.crl (import Helper grimoire)

// 3. Simplified package imports (auto-resolves versions)
import "hello-world/main"              // → carrion_modules/hello-world/0.1.0/src/main.crl
import "hello-world/main.HelloWorld"   // → carrion_modules/hello-world/0.1.0/src/main.crl (HelloWorld grimoire)

// 4. Relative path imports
import "./utils"                  // → ./utils.crl (explicit current)
import "./models/user"            // → ./models/user.crl
import "../shared/common"         // → ../shared/common.crl
import "../shared/common.Logger"  // → ../shared/common.crl (Logger grimoire)

// 5. With aliases for convenience
import "hello-world/main.HelloWorld" as Hello
import "../shared/common.Logger" as Log
import "utils.Helper" as MyHelper

// 6. Legacy full paths (still supported)
import "carrion_modules/hello-world/0.1.0/src/main.HelloWorld" as Hello
```

### Usage Examples

**File: `utils.crl`**
```python
grim Helper:
    init():
        ignore
    
    spell format_text(text):
        return f"Formatted: {text}"

grim StringUtils:
    init():
        ignore
    
    spell reverse(text):
        return text[::-1]
```

**File: `main.crl`**
```python
// Smart imports in action - mixing grimoire and path imports
import "Helper" as MyHelper                     // NEW: Grimoire-based import
import "HelloWorld" as Hello                    // NEW: Grimoire-based import
import "../shared/common.Logger" as Log         // Relative selective import

main:
    // Use imported grimoires directly
    helper = MyHelper()
    result = helper.format_text("Hello World")
    print(result)  // → "Formatted: Hello World"
    
    // Use package grimoire (found automatically)
    greeting = Hello()
    greeting.print_greeting()
    
    // Use relative import
    logger = Log()
    logger.info("Application started")
```

## Import Examples

### Basic File Import

**File: `math_utils.crl`**
```python
// Mathematical utility functions
spell add(a, b):
    return a + b

spell multiply(a, b):
    return a * b

spell factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)

grim Calculator:
    init():
        self.history = []
    
    spell calculate(operation, a, b):
        match operation:
            case "add":
                result = add(a, b)
            case "multiply":
                result = multiply(a, b)
            _:
                result = "Unknown operation"
        
        self.history.append(f"{operation}({a}, {b}) = {result}")
        return result
```

**File: `main.crl`**
```python
import "math_utils"

// Use imported functions
result1 = add(5, 3)
result2 = multiply(4, 7)
fact = factorial(5)

print(f"5 + 3 = {result1}")      // → "5 + 3 = 8"
print(f"4 * 7 = {result2}")      // → "4 * 7 = 28"
print(f"5! = {fact}")            // → "5! = 120"

// Use imported grimoire
calc = Calculator()
sum_result = calc.calculate("add", 10, 15)
print(f"Calculator result: {sum_result}")  // → "Calculator result: 25"
```

## Selective Imports

### Importing Specific Grimoires
```python
import "data_structures.Stack"
import "utilities.FileHelper"
```

**File: `data_structures.crl`**
```python
grim Stack:
    init():
        self.items = []
    
    spell push(item):
        self.items.append(item)
    
    spell pop():
        if len(self.items) > 0:
            return self.items.pop()
        return None
    
    spell peek():
        if len(self.items) > 0:
            return self.items[-1]
        return None

grim Queue:
    init():
        self.items = []
    
    spell enqueue(item):
        self.items.append(item)
    
    spell dequeue():
        if len(self.items) > 0:
            return self.items.pop(0)
        return None
```

**File: `main.crl`**
```python
import "data_structures.Stack"
// Only Stack is imported, Queue is not available

stack = Stack()
stack.push(1)
stack.push(2)
print(stack.pop())  // → 2

// queue = Queue()  // ✗ Error: Queue not imported
```

### Import with Aliases
```python
import "very_long_module_name" as short_name
import "data_structures.Stack" as MyStack
```

**Example:**
```python
import "mathematical_operations" as math_ops
import "string_utilities.StringProcessor" as StrProc

result = math_ops.complex_calculation(10, 20)
processor = StrProc("Hello World")
formatted = processor.format_title()
```

## Package Management with Bifrost

Carrion integrates with the Bifrost package manager to provide seamless package installation and import capabilities.

### Installing Packages

Use Bifrost to install packages for use in Carrion projects:

```bash
# Install a package locally to the project
bifrost install json-utils

# Install a package globally (system-wide)
bifrost install --global http-client

# Initialize a new Carrion package
bifrost init
```

### Package Directory Structure

Packages are organized in versioned directories:

```
/usr/local/share/carrion/lib/          # Global packages
├── json-utils/
│   ├── 1.0.0/
│   │   ├── parser.crl
│   │   ├── formatter.crl
│   │   └── Bifrost.toml
│   └── 1.0.1/
│       ├── parser.crl
│       ├── formatter.crl
│       └── Bifrost.toml
└── http-client/
    └── 2.1.0/
        ├── request.crl
        ├── response.crl
        └── auth.crl

./carrion_modules/                      # Project-local packages
├── test-utils/
│   ├── mock.crl
│   └── assert.crl
└── dev-helpers/
    └── debug.crl
```

### Using Installed Packages

Once installed, packages can be imported using simplified syntax:

```python
# Simplified package imports (auto-resolves to latest version)
import "json-utils/parser"                    // Auto-resolves to src/parser.crl
import "json-utils/parser.JSONParser" as JSON // Selective import with alias
import "http-client/request.HTTPClient" as HTTP

# Use imported functionality
json_parser = JSON()  // Direct use with alias
data = json_parser.parse('{"name": "example"}')

http = HTTP()  // Direct use with alias
response = http.get("https://api.example.com/data")

# Or import entire modules
import "json-utils/parser"
json_parser = JSONParser()  // Use original class name
```

### Version Resolution

Carrion automatically resolves to the latest available version of a package. The import system:

1. Searches for the package in the priority order (local → project → user → global)
2. Finds the latest version directory for the package
3. Imports the requested module from that version

### Environment Variables

You can customize package locations using environment variables:

```bash
# Custom Carrion home directory
export CARRION_HOME=/custom/path/.carrion

# Additional import paths (colon-separated)
export CARRION_IMPORT_PATH=/custom/lib:/another/path
```

### Package Import Examples

```python
# Smart import patterns with automatic resolution

# 1. Local file imports (current directory)
import "helper"                         # → ./helper.crl
import "helper.Helper" as H             # → ./helper.crl (Helper grimoire)

# 2. Simplified package imports (auto-resolves version and src/ path)
import "test-utils/mock"                # → carrion_modules/test-utils/[version]/src/mock.crl
import "test-utils/mock.MockFramework" as Mock  # → selective import with alias

# 3. Global package imports (system-wide, auto-resolves)
import "json-utils/parser"              # → /usr/local/share/carrion/lib/json-utils/[version]/src/parser.crl
import "json-utils/parser.JSONParser" as JSON  # → selective import

# 4. User package imports (auto-resolves)
import "my-lib/utils"                   # → ~/.carrion/packages/my-lib/[version]/src/utils.crl
import "my-lib/utils.Utility" as Util   # → selective import

# 5. Relative imports for shared code
import "../shared/common.Logger" as Log  # → ../shared/common.crl (Logger grimoire)

# Usage examples
mock = Mock()           # Use aliased import
json_parser = JSON()    # Use aliased selective import
helper = Helper()       # Use local import
logger = Log()          # Use relative import
```

## Module Organization Patterns

### Utility Modules

**File: `string_utils.crl`**
```python
// String manipulation utilities
spell reverse_string(text):
    return text[::-1]  // Assuming string slicing is supported

spell capitalize_words(text):
    words = text.split(" ")
    capitalized = []
    for word in words:
        if len(word) > 0:
            capitalized.append(word[0].upper() + word[1:].lower())
    return " ".join(capitalized)

spell count_words(text):
    return len(text.split(" "))

grim TextAnalyzer:
    init(text):
        self.text = text
    
    spell word_count():
        return count_words(self.text)
    
    spell character_count():
        return len(self.text)
    
    spell sentence_count():
        return len(self.text.split("."))
```

### Constants Module

**File: `constants.crl`**
```python
// Application constants
PI = 3.14159265359
E = 2.71828182846
GOLDEN_RATIO = 1.61803398875

// Configuration constants
MAX_RETRY_ATTEMPTS = 3
DEFAULT_TIMEOUT = 30
API_VERSION = "v1.2.0"

// Color constants
COLOR_RED = "#FF0000"
COLOR_GREEN = "#00FF00"
COLOR_BLUE = "#0000FF"

grim Colors:
    RED = "#FF0000"
    GREEN = "#00FF00"
    BLUE = "#0000FF"
    
    spell hex_to_rgb(hex_color):
        // Convert hex to RGB implementation
        return (255, 0, 0)  // Placeholder
```

### Configuration Module

**File: `config.crl`**
```python
grim Config:
    init():
        self.settings = {
            "debug": False,
            "log_level": "INFO",
            "database_url": "localhost:5432",
            "cache_enabled": True
        }
    
    spell get(key, default = None):
        return self.settings.get(key, default)
    
    spell set(key, value):
        self.settings[key] = value
    
    spell load_from_file(filename):
        // Load configuration from file
        if File.exists(filename):
            content = File.read(filename)
            // Parse and update settings
            print(f"Loaded config from {filename}")
    
    spell save_to_file(filename):
        // Save configuration to file
        content = str(self.settings)  // Simplified serialization
        File.write(filename, content)
        print(f"Saved config to {filename}")

// Global configuration instance
app_config = Config()
```

## Advanced Import Patterns

### Conditional Imports
```python
// Import based on conditions
debug_mode = True

if debug_mode:
    import "debug_utilities"
    enable_debug_logging()
else:
    import "production_utilities"
    enable_performance_monitoring()
```

### Dynamic Module Loading
```python
// Import modules based on runtime decisions
spell load_database_driver(database_type):
    match database_type:
        case "mysql":
            import "drivers.mysql_driver"
            return MySQLDriver()
        case "postgresql":
            import "drivers.postgresql_driver" 
            return PostgreSQLDriver()
        case "sqlite":
            import "drivers.sqlite_driver"
            return SQLiteDriver()
        _:
            raise Error("Database", f"Unsupported database type: {database_type}")

// Usage
db_type = input("Enter database type: ")
driver = load_database_driver(db_type)
```

### Module Initialization
```python
// Module with initialization code
// File: `logger.crl`

// Module-level initialization
print("Logger module loaded")

grim Logger:
    init(name):
        self.name = name
        self.messages = []
    
    spell log(level, message):
        timestamp = get_current_time()  // Assuming this function exists
        formatted_message = f"[{timestamp}] {level}: {message}"
        self.messages.append(formatted_message)
        print(formatted_message)
    
    spell debug(message):
        self.log("DEBUG", message)
    
    spell info(message):
        self.log("INFO", message)
    
    spell error(message):
        self.log("ERROR", message)

// Create default logger
default_logger = Logger("default")

// This runs when module is imported
default_logger.info("Logger module initialized")
```

## File Organization Best Practices

### Project Structure
```
project/
├── main.crl                 // Main application entry point
├── config.crl              // Configuration settings
├── constants.crl           // Application constants
├── utils/
│   ├── string_utils.crl    // String manipulation utilities
│   ├── math_utils.crl      // Mathematical functions
│   └── file_utils.crl      // File operations
├── models/
│   ├── user.crl           // User data model
│   ├── product.crl        // Product data model
│   └── order.crl          // Order data model
├── services/
│   ├── user_service.crl   // User business logic
│   ├── auth_service.crl   // Authentication logic
│   └── data_service.crl   // Data access layer
└── tests/
    ├── test_utils.crl     // Test utilities
    ├── test_models.crl    // Model tests
    └── test_services.crl  // Service tests
```

### Main Application Structure

**File: `main.crl`**
```python
// Main application entry point
import "config"
import "services.user_service"
import "services.auth_service"
import "utils.string_utils"

// Initialize application
app_config.load_from_file("app.config")
auth = AuthService()
user_service = UserService()

spell main():
    print("Welcome to Carrion Application")
    
    username = input("Username: ")
    password = input("Password: ")
    
    if auth.authenticate(username, password):
        user = user_service.get_user(username)
        formatted_name = capitalize_words(user.full_name)
        print(f"Welcome, {formatted_name}!")
        
        // Application logic continues...
    else:
        print("Authentication failed")

// Run application
if __name__ == "__main__":  // Assuming this pattern is supported
    main()
```

## Error Handling with Imports

### Handling Missing Modules
```python
spell safe_import(module_name):
    attempt:
        import module_name
        return True
    ensnare:
        print(f"Failed to import {module_name}")
        return False

// Graceful fallback for optional features
if safe_import("advanced_graphics"):
    use_advanced_graphics = True
    print("Advanced graphics enabled")
else:
    use_advanced_graphics = False
    print("Using basic graphics")
```

### Version Compatibility
```python
// Check for required functionality
import "utils.version_checker"

spell check_dependencies():
    required_modules = ["math_utils", "string_utils", "data_structures"]
    
    for module in required_modules:
        if not safe_import(module):
            print(f"Error: Required module '{module}' not found")
            return False
    
    print("All dependencies satisfied")
    return True

if not check_dependencies():
    print("Cannot start application - missing dependencies")
    exit(1)
```

## Module Documentation

### Self-Documenting Modules
```python
// File: `documented_module.crl`
"""
Math Utilities Module

This module provides essential mathematical functions and utilities
for the Carrion application.

Author: Developer Name
Version: 1.0.0
"""

spell add(a, b):
    """
    Add two numbers together.
    
    Parameters:
        a: First number
        b: Second number
    
    Returns:
        Sum of a and b
    """
    return a + b

spell divide(a, b):
    """
    Divide two numbers with error handling.
    
    Parameters:
        a: Dividend
        b: Divisor
    
    Returns:
        Result of division or error message
    
    Raises:
        Error if division by zero
    """
    if b == 0:
        raise Error("Math", "Division by zero")
    return a / b

grim MathCalculator:
    """
    A calculator class for advanced mathematical operations.
    
    This class maintains calculation history and provides
    methods for complex mathematical operations.
    """
    
    init():
        """Initialize calculator with empty history."""
        self.history = []
    
    spell calculate(operation, *args):
        """
        Perform calculation and store in history.
        
        Parameters:
            operation: String name of operation
            *args: Arguments for the operation
        
        Returns:
            Calculation result
        """
        # Implementation here
        pass
```

## Sockets Module

Carrion includes a powerful sockets module that simplifies network programming by wrapping Go's net and net/http packages with an easy-to-use Carrion interface. The sockets module supports TCP, UDP, Unix domain sockets, and HTTP/Web servers with built-in timeout and connection management.

### Socket Creation

#### Basic Socket Creation
```python
// Create different types of sockets
tcp_socket = new_socket("tcp", "tcp", "localhost:8080", 30)
udp_socket = new_socket("udp", "udp", "localhost:9090", 30)
web_socket = new_socket("web", "http", "localhost:8000", 60)
unix_socket = new_socket("unix", "unix", "/tmp/my.sock", 30)
```

#### Simplified Socket Creation
```python
// Using defaults - creates TCP socket on localhost:8080 with 30s timeout
socket_id = new_socket("tcp")

// With custom address
socket_id = new_socket("tcp", "tcp", "192.168.1.100:9000")

// With custom timeout (in seconds)
socket_id = new_socket("tcp", "tcp", "localhost:8080", 60)
```

### Client Connections

#### TCP Client
```python
// Connect to a TCP server
client_id = client("tcp", "localhost:8080", 30)

// Send data
bytes_sent = socket_send(client_id, "Hello Server!")

// Receive data
response = socket_receive(client_id, 1024)
print(f"Server response: {response}")

// Close connection
socket_close(client_id)
```

#### UDP Client
```python
// Connect to UDP endpoint
udp_client = client("udp", "localhost:9090", 30)

// Send UDP datagram
socket_send(udp_client, "UDP message")

// Receive response
response = socket_receive(udp_client, 1024)
print(f"UDP response: {response}")

socket_close(udp_client)
```

#### Unix Domain Socket Client
```python
// Connect to Unix socket
unix_client = client("unix", "/tmp/server.sock", 30)
socket_send(unix_client, "Unix socket message")
response = socket_receive(unix_client)
socket_close(unix_client)
```

### Server Creation

#### TCP Server
```python
// Start TCP server
server_id = server("tcp", "localhost:8080", 30)

// Listen for connections
listener_id = socket_listen(server_id)

// Accept client connections
while True:
    client_conn = socket_accept(listener_id)
    
    // Handle client
    message = socket_receive(client_conn, 1024)
    print(f"Client says: {message}")
    
    socket_send(client_conn, "Hello Client!")
    socket_close(client_conn)
```

#### UDP Server
```python
// Start UDP server
udp_server = server("udp", "localhost:9090", 30)

// Receive and respond to datagrams
while True:
    data = socket_receive(udp_server, 1024)
    print(f"Received UDP: {data}")
    
    socket_send(udp_server, "UDP response")
```

#### Web/HTTP Server
```python
// Start HTTP server
web_server = server("web", "localhost:8000", 60)

// The HTTP server runs in the background
// Routes and handlers would be configured separately
print("HTTP server started on localhost:8000")

// Server continues running until closed
// socket_close(web_server)  // Stop the server
```

#### Unix Domain Socket Server
```python
// Start Unix socket server
unix_server = server("unix", "/tmp/server.sock", 30)
listener = socket_listen(unix_server)

while True:
    client = socket_accept(listener)
    data = socket_receive(client, 1024)
    socket_send(client, f"Echo: {data}")
    socket_close(client)
```

### Socket Management

#### Setting Timeouts
```python
socket_id = new_socket("tcp", "tcp", "localhost:8080")

// Set timeout to 45 seconds
socket_set_timeout(socket_id, 45)

// Timeout applies to all subsequent operations
client_id = client("tcp", "localhost:8080")
socket_set_timeout(client_id, 10)  // 10 second timeout for this client
```

#### Getting Socket Information
```python
socket_id = new_socket("tcp", "tcp", "localhost:8080", 30)

info = socket_get_info(socket_id)
print(f"Socket type: {info['type']}")
print(f"Address: {info['address']}")
print(f"Timeout: {info['timeout']} seconds")
```

### Complete Examples

#### Simple Echo Server
```python
// echo_server.crl
import "sockets"

spell start_echo_server():
    server_id = server("tcp", "localhost:8080", 30)
    listener = socket_listen(server_id)
    
    print("Echo server started on localhost:8080")
    
    while True:
        attempt:
            client = socket_accept(listener)
            print("Client connected")
            
            while True:
                data = socket_receive(client, 1024)
                if len(data) == 0:
                    break
                
                print(f"Echoing: {data}")
                socket_send(client, f"Echo: {data}")
            
            socket_close(client)
            print("Client disconnected")
            
        ensnare error:
            print(f"Error handling client: {error}")

start_echo_server()
```

#### HTTP-like Client
```python
// http_client.crl
import "sockets"

spell make_http_request(host, port, path):
    client_id = client("tcp", f"{host}:{port}", 30)
    
    // Send HTTP request
    request = f"GET {path} HTTP/1.1\r\nHost: {host}\r\n\r\n"
    socket_send(client_id, request)
    
    // Receive response
    response = socket_receive(client_id, 4096)
    socket_close(client_id)
    
    return response

// Usage
response = make_http_request("httpbin.org", "80", "/get")
print(response)
```

#### Chat Server
```python
// chat_server.crl
import "sockets"

clients = []

spell handle_client(client_id):
    while True:
        attempt:
            message = socket_receive(client_id, 1024)
            if len(message) == 0:
                break
            
            // Broadcast to all clients
            broadcast_message = f"Client says: {message}"
            for other_client in clients:
                if other_client != client_id:
                    socket_send(other_client, broadcast_message)
                    
        ensnare:
            break
    
    // Remove client
    clients.remove(client_id)
    socket_close(client_id)

spell start_chat_server():
    server_id = server("tcp", "localhost:8080", 30)
    listener = socket_listen(server_id)
    
    print("Chat server started on localhost:8080")
    
    while True:
        client = socket_accept(listener)
        clients.append(client)
        print(f"Client connected. Total clients: {len(clients)}")
        
        // Handle client in background (simplified - would use threads/async)
        handle_client(client)

start_chat_server()
```

### Socket Functions Reference

| Function | Description | Parameters | Returns |
|----------|-------------|------------|---------|
| `new_socket(type, [protocol], [address], [timeout])` | Create new socket | type, protocol, address, timeout | socket handle |
| `client(type, address, [timeout])` | Connect as client | type, address, timeout | client handle |
| `server(type, address, [timeout])` | Start server | type, address, timeout | server handle |
| `socket_send(handle, data)` | Send data | handle, data string | bytes sent |
| `socket_receive(handle, [buffer_size])` | Receive data | handle, buffer size | received string |
| `socket_close(handle)` | Close socket | handle | none |
| `socket_listen(handle)` | Listen for connections | server handle | listener handle |
| `socket_accept(handle)` | Accept client connection | listener handle | client handle |
| `socket_set_timeout(handle, seconds)` | Set timeout | handle, timeout seconds | none |
| `socket_get_info(handle)` | Get socket info | handle | info hash |

### Supported Socket Types

- **TCP**: Reliable, connection-oriented protocol
- **UDP**: Fast, connectionless protocol  
- **Web/HTTP**: HTTP server functionality
- **Unix**: Unix domain sockets for local IPC

### Automatic Port Allocation and Conflict Resolution

The sockets module includes intelligent port allocation with automatic conflict resolution:

#### Mutex-Protected Port Allocation
- All port allocation is protected by a global mutex for thread safety
- Prevents race conditions when multiple servers try to bind to the same port simultaneously
- Tracks allocated ports to avoid conflicts

#### Automatic Port Incrementing
When creating a server and the requested port is already in use:
- The system automatically tries the next available port (increments by 1)
- Continues trying up to 100 ports to find an available one
- Displays an informative message about the port change
- Works for TCP, UDP, and Web/HTTP servers

#### Unix Socket Path Resolution
For Unix domain sockets, if the socket file path is already in use:
- Automatically appends a number to the filename (e.g., `/tmp/server.sock.1`)
- Continues incrementing until an available path is found
- Displays an informative message about the path change

#### Example of Automatic Port Allocation
```python
// First server gets the requested port
server1 = server("tcp", "localhost:8080", 30)
// Output: Server started on localhost:8080

// Second server automatically gets incremented port
server2 = server("tcp", "localhost:8080", 30)  
// Output: Port 8080 already allocated, incremented to port 8081

// Third server gets next available port
server3 = server("tcp", "localhost:8080", 30)
// Output: Port 8080 already allocated, incremented to port 8082

// When servers are closed, ports are automatically released
socket_close(server1)  // Port 8080 is now available again
```

#### Port Release and Cleanup
- Ports are automatically released when sockets are closed with `socket_close()`
- Ensures efficient port reuse and prevents port exhaustion
- Maintains accurate tracking of allocated vs. available ports

The sockets module provides a unified interface for all network programming needs in Carrion, abstracting away the complexity of Go's networking APIs while maintaining full functionality and performance.

The module system in Carrion provides flexible code organization capabilities while maintaining simplicity and readability. It supports both simple imports for small projects and sophisticated module hierarchies for larger applications.

## Sockets Module

Carrion includes a powerful sockets module that simplifies network programming by wrapping Go's net and net/http packages with an easy-to-use Carrion interface. The sockets module supports TCP, UDP, Unix domain sockets, and HTTP/Web servers with built-in timeout and connection management.

### Socket Creation

#### Basic Socket Creation
```python
// Create different types of sockets
tcp_socket = new_socket("tcp", "tcp", "localhost:8080", 30)
udp_socket = new_socket("udp", "udp", "localhost:9090", 30)
web_socket = new_socket("web", "http", "localhost:8000", 60)
unix_socket = new_socket("unix", "unix", "/tmp/my.sock", 30)
```

#### Simplified Socket Creation
```python
// Using defaults - creates TCP socket on localhost:8080 with 30s timeout
socket_id = new_socket("tcp")

// With custom address
socket_id = new_socket("tcp", "tcp", "192.168.1.100:9000")

// With custom timeout (in seconds)
socket_id = new_socket("tcp", "tcp", "localhost:8080", 60)
```

### Client Connections

#### TCP Client
```python
// Connect to a TCP server
client_id = client("tcp", "localhost:8080", 30)

// Send data
bytes_sent = socket_send(client_id, "Hello Server!")

// Receive data
response = socket_receive(client_id, 1024)
print(f"Server response: {response}")

// Close connection
socket_close(client_id)
```

#### UDP Client
```python
// Connect to UDP endpoint
udp_client = client("udp", "localhost:9090", 30)

// Send UDP datagram
socket_send(udp_client, "UDP message")

// Receive response
response = socket_receive(udp_client, 1024)
print(f"UDP response: {response}")

socket_close(udp_client)
```

#### Unix Domain Socket Client
```python
// Connect to Unix socket
unix_client = client("unix", "/tmp/server.sock", 30)
socket_send(unix_client, "Unix socket message")
response = socket_receive(unix_client)
socket_close(unix_client)
```

### Server Creation

#### TCP Server
```python
// Start TCP server
server_id = server("tcp", "localhost:8080", 30)

// Listen for connections
listener_id = socket_listen(server_id)

// Accept client connections
while True:
    client_conn = socket_accept(listener_id)
    
    // Handle client
    message = socket_receive(client_conn, 1024)
    print(f"Client says: {message}")
    
    socket_send(client_conn, "Hello Client!")
    socket_close(client_conn)
```

#### UDP Server
```python
// Start UDP server
udp_server = server("udp", "localhost:9090", 30)

// Receive and respond to datagrams
while True:
    data = socket_receive(udp_server, 1024)
    print(f"Received UDP: {data}")
    
    socket_send(udp_server, "UDP response")
```

#### Web/HTTP Server
```python
// Start HTTP server
web_server = server("web", "localhost:8000", 60)

// The HTTP server runs in the background
// Routes and handlers would be configured separately
print("HTTP server started on localhost:8000")

// Server continues running until closed
// socket_close(web_server)  // Stop the server
```

#### Unix Domain Socket Server
```python
// Start Unix socket server
unix_server = server("unix", "/tmp/server.sock", 30)
listener = socket_listen(unix_server)

while True:
    client = socket_accept(listener)
    data = socket_receive(client, 1024)
    socket_send(client, f"Echo: {data}")
    socket_close(client)
```

### Socket Management

#### Setting Timeouts
```python
socket_id = new_socket("tcp", "tcp", "localhost:8080")

// Set timeout to 45 seconds
socket_set_timeout(socket_id, 45)

// Timeout applies to all subsequent operations
client_id = client("tcp", "localhost:8080")
socket_set_timeout(client_id, 10)  // 10 second timeout for this client
```

#### Getting Socket Information
```python
socket_id = new_socket("tcp", "tcp", "localhost:8080", 30)

info = socket_get_info(socket_id)
print(f"Socket type: {info['type']}")
print(f"Address: {info['address']}")
print(f"Timeout: {info['timeout']} seconds")
```

### Complete Examples

#### Simple Echo Server
```python
// echo_server.crl
import "sockets"

spell start_echo_server():
    server_id = server("tcp", "localhost:8080", 30)
    listener = socket_listen(server_id)
    
    print("Echo server started on localhost:8080")
    
    while True:
        attempt:
            client = socket_accept(listener)
            print("Client connected")
            
            while True:
                data = socket_receive(client, 1024)
                if len(data) == 0:
                    break
                
                print(f"Echoing: {data}")
                socket_send(client, f"Echo: {data}")
            
            socket_close(client)
            print("Client disconnected")
            
        ensnare error:
            print(f"Error handling client: {error}")

start_echo_server()
```

#### HTTP-like Client
```python
// http_client.crl
import "sockets"

spell make_http_request(host, port, path):
    client_id = client("tcp", f"{host}:{port}", 30)
    
    // Send HTTP request
    request = f"GET {path} HTTP/1.1\r\nHost: {host}\r\n\r\n"
    socket_send(client_id, request)
    
    // Receive response
    response = socket_receive(client_id, 4096)
    socket_close(client_id)
    
    return response

// Usage
response = make_http_request("httpbin.org", "80", "/get")
print(response)
```

#### Chat Server
```python
// chat_server.crl
import "sockets"

clients = []

spell handle_client(client_id):
    while True:
        attempt:
            message = socket_receive(client_id, 1024)
            if len(message) == 0:
                break
            
            // Broadcast to all clients
            broadcast_message = f"Client says: {message}"
            for other_client in clients:
                if other_client != client_id:
                    socket_send(other_client, broadcast_message)
                    
        ensnare:
            break
    
    // Remove client
    clients.remove(client_id)
    socket_close(client_id)

spell start_chat_server():
    server_id = server("tcp", "localhost:8080", 30)
    listener = socket_listen(server_id)
    
    print("Chat server started on localhost:8080")
    
    while True:
        client = socket_accept(listener)
        clients.append(client)
        print(f"Client connected. Total clients: {len(clients)}")
        
        // Handle client in background (simplified - would use threads/async)
        handle_client(client)

start_chat_server()
```

### Socket Functions Reference

| Function | Description | Parameters | Returns |
|----------|-------------|------------|---------|
| `new_socket(type, [protocol], [address], [timeout])` | Create new socket | type, protocol, address, timeout | socket handle |
| `client(type, address, [timeout])` | Connect as client | type, address, timeout | client handle |
| `server(type, address, [timeout])` | Start server | type, address, timeout | server handle |
| `socket_send(handle, data)` | Send data | handle, data string | bytes sent |
| `socket_receive(handle, [buffer_size])` | Receive data | handle, buffer size | received string |
| `socket_close(handle)` | Close socket | handle | none |
| `socket_listen(handle)` | Listen for connections | server handle | listener handle |
| `socket_accept(handle)` | Accept client connection | listener handle | client handle |
| `socket_set_timeout(handle, seconds)` | Set timeout | handle, timeout seconds | none |
| `socket_get_info(handle)` | Get socket info | handle | info hash |

### Supported Socket Types

- **TCP**: Reliable, connection-oriented protocol
- **UDP**: Fast, connectionless protocol  
- **Web/HTTP**: HTTP server functionality
- **Unix**: Unix domain sockets for local IPC

The sockets module provides a unified interface for all network programming needs in Carrion, abstracting away the complexity of Go's networking APIs while maintaining full functionality and performance.

## Server Framework (Grimoires)

Carrion provides an object-oriented server framework through grimoire classes that simplify server creation and management. These classes provide auto-close functionality, inheritance hierarchies, and standardized interfaces for different server types.

### Base Server Class

The `Server` grimoire provides the foundation for all server types:

```python
# Create base server (usually not used directly)
base_server = Server("tcp", "localhost:8080", 30)

# Basic server operations
base_server.start()              // Start the server
base_server.shutdown()           // Stop the server  
base_server.is_active()          // Check if running
base_server.get_info()           // Get server information

# Context management
base_server.set_context("key", "value")
value = base_server.get_context("key", "default")
```

### TCP Server

The `TCPServer` grimoire provides TCP socket server functionality:

```python
# Create TCP server
tcp_server = TCPServer("127.0.0.1:8080", 30)

# Start server and accept connections
tcp_server.tcp_start()
client = tcp_server.accept_client()

# Client communication
tcp_server.send_to_client(client, "Hello Client!")
response = tcp_server.receive_from_client(client, 1024)

# Server management
print(f"Connected clients: {tcp_server.client_count()}")
tcp_server.tcp_shutdown()
```

### UDP Server  

The `UDPServer` grimoire provides UDP datagram communication:

```python
# Create UDP server
udp_server = UDPServer("127.0.0.1:8081", 30)

# Start server
udp_server.udp_start()

# Send and receive datagrams
udp_server.send_datagram("Hello!", "127.0.0.1:9000")
data_and_sender = udp_server.receive_datagram(1024)

# Message buffering
udp_server.buffer_message("Test message", "sender_address")
buffered_messages = udp_server.get_buffered_messages()
message_count = udp_server.message_count()

# Cleanup
udp_server.udp_shutdown()
```

### Unix Domain Socket Server

The `UnixServer` grimoire provides Unix domain socket communication:

```python
# Create Unix socket server
unix_server = UnixServer("/tmp/my_socket", 30)

# Start server
unix_server.unix_start()

# Accept and handle clients
client = unix_server.accept_client()
unix_server.send_to_client(client, "Hello Unix Client!")
response = unix_server.receive_from_client(client, 1024)

# Broadcast to all clients  
sent_count = unix_server.broadcast_to_clients("Broadcast message")
print(f"Clients connected: {unix_server.client_count()}")

# Cleanup (automatically removes socket file)
unix_server.unix_shutdown()
```

### HTTP Server

The `HTTPServer` grimoire provides HTTP protocol server with routing:

```python
# Create HTTP server
http_server = HTTPServer("127.0.0.1", 8082, 30)

# Add routes (handlers would be defined as spell functions)
http_server.add_route("GET", "/api/users", user_handler)
http_server.add_route("POST", "/api/data", data_handler)

# Add middleware 
http_server.add_middleware(auth_middleware)
http_server.add_middleware(logging_middleware)

# Static file serving
http_server.add_static_path("/static", "/var/www/static")

# Start server
http_server.http_start()

# Process requests (typically in a loop)
response = http_server.handle_request(request_data)

# Server info
routes = http_server.get_routes()
route_count = http_server.route_count()

# Cleanup
http_server.http_shutdown()
```

### Web Server

The `WebServer` grimoire extends HTTPServer for static file serving:

```python
# Create web server with document root
web_server = WebServer("127.0.0.1", 8083, 30, "/var/www/html")

# Configure MIME types
web_server.add_mime_type("crl", "text/x-carrion")

# Set default index pages
web_server.set_default_pages(["index.html", "home.html", "main.html"])

# Start web server (automatically serves static files)
web_server.web_start()

# Serve specific files
response = web_server.serve_file("about.html")
dir_listing = web_server.serve_directory("docs/")

# Cleanup
web_server.web_shutdown()
```

### Server Framework Features

#### Auto-Close Registration
All servers automatically register for cleanup when created:
```python
# Servers are automatically added to _auto_close_servers
tcp_server = TCPServer("localhost:8080", 30)
// Server will be cleaned up automatically on program exit
```

#### Inheritance and Polymorphism
```python
# All server types inherit from base Server class
servers = [
    TCPServer("localhost:8080", 30),
    UDPServer("localhost:8081", 30), 
    UnixServer("/tmp/socket", 30),
    HTTPServer("localhost", 8082, 30),
    WebServer("localhost", 8083, 30, "/var/www")
]

# Common interface across all server types
for server in servers:
    server.set_context("created", "now")
    print(f"Server active: {server.is_active()}")
```

#### Context Management
All servers support context storage for application data:
```python
server.set_context("database_url", "postgresql://...")
server.set_context("max_connections", 100)
server.set_context("ssl_enabled", True)

// Retrieve with defaults
db_url = server.get_context("database_url", "sqlite://default.db")
max_conn = server.get_context("max_connections", 50)
```

### Server Types Summary

| Server Type | Use Case | Key Features |
|-------------|----------|--------------|
| `Server` | Base class | Context management, auto-close |
| `TCPServer` | TCP connections | Client management, reliable communication |
| `UDPServer` | UDP datagrams | Message buffering, connectionless |
| `UnixServer` | Local IPC | Unix sockets, client broadcasting |
| `HTTPServer` | HTTP protocol | Routing, middleware, static files |
| `WebServer` | Static websites | Document root, MIME types, directory listing |

The server framework provides a modern, object-oriented approach to network programming in Carrion while maintaining the simplicity and performance of the underlying socket system.