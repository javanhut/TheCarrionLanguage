# Error Handling

Carrion provides a comprehensive error handling system with magical terminology that makes exception management both powerful and intuitive.

## Basic Error Handling Syntax

### Attempt-Ensnare-Resolve Structure
```python
attempt:
    // Risky code that might fail
ensnare (ErrorType):
    // Handle specific error type
ensnare:
    // Handle any other error
resolve:
    // Code that always runs (finally block)
```

## Basic Error Handling

### Simple Try-Catch
```python
attempt:
    result = 10 / 0
ensnare:
    print("An error occurred!")
    result = 0

print(f"Result: {result}")
```

### Multiple Exception Types
```python
attempt:
    number = int(input("Enter a number: "))
    result = 100 / number
    print(f"Result: {result}")
ensnare (ValueError):
    print("Invalid number format!")
ensnare (ZeroDivisionError):
    print("Cannot divide by zero!")
ensnare:
    print("An unexpected error occurred!")
```

### With Finally (Resolve) Block
```python
file = None
attempt:
    file = File()
    content = file.read("important_data.txt")
    process_data(content)
ensnare (FileNotFoundError):
    print("File not found!")
ensnare:
    print("Error processing file!")
resolve:
    // This always runs
    if file is not None:
        print("Cleanup completed")
```

## Raising Errors

### Basic Error Raising
```python
spell validate_age(age):
    if age < 0:
        raise Error("Validation", "Age cannot be negative")
    if age > 150:
        raise Error("Validation", "Age seems unrealistic")
    return True

// Usage
attempt:
    validate_age(-5)
ensnare (Error):
    print("Validation failed!")
```

### Custom Error Types
```python
// Create custom error objects
ValidationError = Error("ValidationError", "Input validation failed")
NetworkError = Error("NetworkError", "Network operation failed")

spell connect_to_server(url):
    if not url.startswith("http"):
        raise ValidationError
    
    // Simulate network connection
    if url == "http://broken-server.com":
        raise NetworkError
    
    return "Connected successfully"

// Usage
attempt:
    result = connect_to_server("invalid-url")
ensnare (ValidationError):
    print("Invalid URL format")
ensnare (NetworkError):
    print("Network connection failed")
```

## Advanced Error Handling Patterns

### Error Information Access
```python
attempt:
    risky_operation()
ensnare (error):
    print(f"Error type: {type(error)}")
    print(f"Error message: {error.message}")
    // Handle based on error details
```

### Nested Error Handling
```python
spell process_user_data(user_data):
    attempt:
        // Validate data structure
        attempt:
            validate_structure(user_data)
        ensnare (ValidationError):
            print("Data structure validation failed")
            raise Error("Processing", "Invalid data structure")
        
        // Process individual fields
        for field in user_data:
            attempt:
                process_field(field)
            ensnare:
                print(f"Warning: Failed to process field {field}")
                skip  // Continue with next field
    
    ensnare (Error):
        print("Critical error in data processing")
        return False
    
    return True
```

### Error Recovery and Retry
```python
spell connect_with_retry(url, max_attempts = 3):
    attempts = 0
    
    while attempts < max_attempts:
        attempt:
            return connect_to_server(url)
        ensnare (NetworkError):
            attempts += 1
            print(f"Connection attempt {attempts} failed")
            
            if attempts < max_attempts:
                print("Retrying in 2 seconds...")
                os = OS()
                os.sleep(2)
            else:
                print("Max retry attempts reached")
                raise Error("Connection", "Failed to connect after retries")

// Usage
attempt:
    connection = connect_with_retry("http://unreliable-server.com")
    print("Successfully connected!")
ensnare:
    print("Connection failed permanently")
```

## Error Handling in Functions

### Function-Level Error Handling
```python
spell safe_divide(a, b):
    """Safely divide two numbers with error handling."""
    attempt:
        if type(a) not in ["INTEGER", "FLOAT"] or type(b) not in ["INTEGER", "FLOAT"]:
            raise Error("TypeError", "Arguments must be numbers")
        
        if b == 0:
            raise Error("ZeroDivision", "Cannot divide by zero")
        
        return a / b
    
    ensnare (Error):
        print(f"Division error: {error.message}")
        return None

// Usage
results = [
    safe_divide(10, 2),    // → 5.0
    safe_divide(10, 0),    // → None (with error message)
    safe_divide("10", 2),  // → None (with error message)
]

for result in results:
    if result is not None:
        print(f"Result: {result}")
```

### Error Propagation
```python
spell read_config_file(filename):
    """Read configuration with proper error propagation."""
    attempt:
        file = File()
        if not file.exists(filename):
            raise Error("FileNotFound", f"Config file '{filename}' not found")
        
        content = file.read(filename)
        if len(content) == 0:
            raise Error("EmptyFile", "Configuration file is empty")
        
        return parse_config(content)  // May also raise errors
    
    ensnare:
        print(f"Failed to read config: {error.message}")
        raise  // Re-raise the error

spell initialize_application():
    """Initialize app with error handling."""
    attempt:
        config = read_config_file("app.config")
        setup_database(config)
        start_services(config)
        return True
    
    ensnare:
        print("Application initialization failed")
        return False

// Usage
if not initialize_application():
    print("Cannot start application")
    exit(1)
```

## Error Handling in Object-Oriented Code

### Grimoire with Error Handling
```python
grim BankAccount:
    init(account_number, initial_balance = 0):
        attempt:
            if initial_balance < 0:
                raise Error("InvalidBalance", "Initial balance cannot be negative")
            
            self.account_number = account_number
            self.balance = initial_balance
            self.transaction_history = []
        
        ensnare:
            print(f"Account creation failed: {error.message}")
            raise  // Re-raise to prevent invalid object creation
    
    spell withdraw(amount):
        attempt:
            if amount <= 0:
                raise Error("InvalidAmount", "Withdrawal amount must be positive")
            
            if amount > self.balance:
                raise Error("InsufficientFunds", "Not enough balance")
            
            self.balance -= amount
            self.transaction_history.append(f"Withdrawal: -{amount}")
            return True
        
        ensnare (Error):
            print(f"Withdrawal failed: {error.message}")
            return False
    
    spell deposit(amount):
        attempt:
            if amount <= 0:
                raise Error("InvalidAmount", "Deposit amount must be positive")
            
            self.balance += amount
            self.transaction_history.append(f"Deposit: +{amount}")
            return True
        
        ensnare (Error):
            print(f"Deposit failed: {error.message}")
            return False

// Usage
attempt:
    account = BankAccount("12345", 1000)
    
    if account.withdraw(500):
        print("Withdrawal successful")
    
    if not account.withdraw(2000):
        print("Large withdrawal blocked")
    
ensnare:
    print("Account operations failed")
```

### Error Handling in Inheritance
```python
grim Shape:
    init(name):
        if name is None or len(name) == 0:
            raise Error("InvalidName", "Shape name cannot be empty")
        self.name = name
    
    spell area():
        raise Error("NotImplemented", "Subclasses must implement area()")

grim Rectangle(Shape):
    init(width, height):
        attempt:
            super.init("Rectangle")
            
            if width <= 0 or height <= 0:
                raise Error("InvalidDimensions", "Width and height must be positive")
            
            self.width = width
            self.height = height
        
        ensnare:
            print(f"Rectangle creation failed: {error.message}")
            raise
    
    spell area():
        return self.width * self.height

// Usage
attempt:
    rect = Rectangle(5, 3)
    print(f"Area: {rect.area()}")
    
    invalid_rect = Rectangle(-5, 3)  // Will raise error
    
ensnare:
    print("Shape creation or calculation failed")
```

## Assertions and Debugging

### Using Check Statements
```python
spell factorial(n):
    check(n >= 0, "Factorial is not defined for negative numbers")
    check(type(n) == "INTEGER", "Factorial requires an integer argument")
    
    if n <= 1:
        return 1
    return n * factorial(n - 1)

// Usage
attempt:
    result = factorial(5)    // ✓ Works fine
    result = factorial(-3)   // ✗ Triggers check error
ensnare:
    print("Factorial calculation failed")
```

### Debug Assertions
```python
DEBUG_MODE = True

spell debug_assert(condition, message):
    if DEBUG_MODE and not condition:
        raise Error("AssertionError", f"Debug assertion failed: {message}")

spell process_array(arr):
    debug_assert(arr is not None, "Array cannot be None")
    debug_assert(len(arr) > 0, "Array cannot be empty")
    debug_assert(type(arr) == "ARRAY", "Input must be an array")
    
    // Process array safely
    total = 0
    for item in arr:
        debug_assert(type(item) in ["INTEGER", "FLOAT"], "Array items must be numbers")
        total += item
    
    return total
```

## Error Handling Best Practices

### Specific Error Messages
```python
// Good: Specific, actionable error messages
spell parse_user_input(input_string):
    if input_string is None:
        raise Error("InvalidInput", "Input cannot be None")
    
    if len(input_string) == 0:
        raise Error("InvalidInput", "Input cannot be empty")
    
    if not input_string.contains("="):
        raise Error("FormatError", "Input must contain '=' separator (format: key=value)")
    
    parts = input_string.split("=")
    if len(parts) != 2:
        raise Error("FormatError", "Input must have exactly one '=' separator")
    
    return parts[0].strip(), parts[1].strip()
```

### Graceful Degradation
```python
spell load_user_preferences():
    """Load user preferences with fallback to defaults."""
    attempt:
        file = File()
        if file.exists("user_prefs.json"):
            return parse_json(file.read("user_prefs.json"))
    ensnare:
        print("Could not load user preferences, using defaults")
    
    // Return default preferences
    return {
        "theme": "light",
        "language": "en",
        "auto_save": True
    }

spell save_user_preferences(prefs):
    """Save preferences with error logging."""
    attempt:
        file = File()
        json_data = to_json(prefs)
        file.write("user_prefs.json", json_data)
        return True
    ensnare:
        print("Warning: Could not save user preferences")
        return False
```

### Resource Management
```python
spell process_large_file(filename):
    """Process file with proper resource management."""
    file = None
    
    attempt:
        file = File()
        
        if not file.exists(filename):
            raise Error("FileNotFound", f"File '{filename}' does not exist")
        
        content = file.read(filename)
        result = expensive_processing(content)
        
        return result
    
    ensnare (Error):
        print(f"File processing failed: {error.message}")
        return None
    
    resolve:
        // Cleanup always happens
        if file is not None:
            print("File processing cleanup completed")
```

### Error Context
```python
spell execute_batch_operations(operations):
    """Execute multiple operations with detailed error context."""
    results = []
    errors = []
    
    for i, operation in enumerate(operations):
        attempt:
            result = execute_operation(operation)
            results.append(result)
        
        ensnare (Error):
            error_context = {
                "operation_index": i,
                "operation": operation,
                "error": error.message
            }
            errors.append(error_context)
            print(f"Operation {i} failed: {error.message}")
    
    return {
        "results": results,
        "errors": errors,
        "success_count": len(results),
        "error_count": len(errors)
    }
```

Error handling in Carrion provides robust mechanisms for creating reliable applications while maintaining the language's magical theme and readable syntax.