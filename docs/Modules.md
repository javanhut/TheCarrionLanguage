# Module System and Imports

Carrion provides a comprehensive module system that allows you to organize code across multiple files and import functionality between them. This enables code reuse, better organization, and collaborative development. The system supports local file imports, project packages, user-specific packages, and system-wide global packages managed by Bifrost package manager.

## Import Resolution System

Carrion uses a sophisticated import resolution system that searches for modules in multiple locations, allowing for flexible project organization and package management.

### Search Path Priority

When you import a module, Carrion searches in this order:

1. **Current Directory** - Local files relative to the current working directory
2. **Project Modules** - `./carrion_modules/` directory for project-specific packages
3. **User Packages** - `~/.carrion/packages/` for user-installed packages  
4. **Global Packages** - `/usr/local/share/carrion/lib/` for system-wide packages (managed by Bifrost)
5. **Standard Library** - Built-in Munin standard library modules

### Basic Import Syntax

#### Simple Import
```python
import "filename"
```

This imports all public functions and grimoires from the specified file.

#### Import with File Extension
```python
import "utilities.crl"  // Explicit .crl extension
import "math_functions"  // Extension optional for .crl files
```

#### Package Imports
```python
import "json-utils/parser"     // Import from package with slash notation
import "http-client/request"   // Package/module structure
import "math-lib/advanced"     // Nested package imports
```

#### Local and Relative Imports
```python
import "utils/helpers"         // Local subdirectory
import "lib/data_structures"   // Local library folder
import "../shared/common"      // Relative paths supported
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

Once installed, packages can be imported using their package name and module path:

```python
# Import from globally installed packages
import "json-utils/parser"
import "http-client/request"

# Use imported functionality
json_parser = JSONParser()
data = json_parser.parse('{"name": "example"}')

http = HTTPClient()
response = http.get("https://api.example.com/data")
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
# Import from different package locations

# 1. Local file (current directory)
import "helper"                    # ./helper.crl

# 2. Project package (carrion_modules)
import "test-utils/mock"          # ./carrion_modules/test-utils/mock.crl

# 3. Global package (system-wide)
import "json-utils/parser"        # /usr/local/share/carrion/lib/json-utils/1.0.0/parser.crl

# 4. User package (~/.carrion/packages)
import "my-lib/utils"             # ~/.carrion/packages/my-lib/1.0.0/utils.crl

# Use imported modules
mock = MockFramework()
json_parser = JSONParser()
local_helper = HelperClass()
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
        file = File()
        if file.exists(filename):
            content = file.read(filename)
            // Parse and update settings
            print(f"Loaded config from {filename}")
    
    spell save_to_file(filename):
        // Save configuration to file
        file = File()
        content = str(self.settings)  // Simplified serialization
        file.write(filename, content)
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

The module system in Carrion provides flexible code organization capabilities while maintaining simplicity and readability. It supports both simple imports for small projects and sophisticated module hierarchies for larger applications.