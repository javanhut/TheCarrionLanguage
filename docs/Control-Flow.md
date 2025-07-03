# Control Flow Structures

Carrion provides comprehensive control flow constructs for conditional execution, loops, pattern matching, and function control.

## Conditional Statements

### Basic If Statement
```python
if condition:
    // code block
```

### If-Else Statement
```python
if condition:
    // code if true
else:
    // code if false
```

### If-Otherwise-Else Statement
The `otherwise` keyword provides additional conditional branches (similar to `elif` in Python):

```python
if condition1:
    // code for condition1
otherwise condition2:
    // code for condition2
otherwise condition3:
    // code for condition3
else:
    // default code
```

### Examples
```python
// Simple if statement
age = 25
if age >= 18:
    print("You are an adult")

// If-else statement
score = 85
if score >= 90:
    grade = "A"
else:
    grade = "B"

// Multiple conditions with otherwise
temperature = 75
if temperature < 32:
    status = "Freezing"
otherwise temperature < 50:
    status = "Cold"
otherwise temperature < 70:
    status = "Cool"
otherwise temperature < 85:
    status = "Warm"
else:
    status = "Hot"
    
print(f"It's {status} outside")
```

## Loops

### For Loops
For loops iterate over sequences like arrays, strings, ranges, or any iterable.

#### Basic For Loop
```python
for item in iterable:
    // code block
```

#### For Loop with Else
The `else` clause executes if the loop completes normally (not broken):
```python
for item in iterable:
    // code block
else:
    // executed if loop completes without break
```

#### Examples
```python
// Iterate over array
numbers = [1, 2, 3, 4, 5]
for num in numbers:
    print(num)

// Iterate over string
word = "hello"
for char in word:
    print(char)

// Iterate over range
for i in range(5):
    print(f"Count: {i}")

// Iterate with range parameters
for i in range(2, 10, 2):  // start=2, stop=10, step=2
    print(i)  // Prints: 2, 4, 6, 8

// For loop with else clause
target = 7
numbers = [1, 3, 5, 9, 11]
for num in numbers:
    if num == target:
        print(f"Found {target}")
        stop  // Break from loop
else:
    print(f"{target} not found")

// Enumerate with index
items = ["apple", "banana", "cherry"]
for index, value in enumerate(items):
    print(f"{index}: {value}")
```

### While Loops
While loops continue executing as long as a condition is true.

```python
while condition:
    // code block
```

#### Examples
```python
// Basic while loop
count = 0
while count < 5:
    print(f"Count: {count}")
    count += 1

// While loop with user input
answer = ""
while answer != "quit":
    answer = input("Enter 'quit' to exit: ")
    if answer != "quit":
        print(f"You entered: {answer}")

// Infinite loop with break
while True:
    user_input = input("Enter a number (or 'exit'): ")
    if user_input == "exit":
        stop  // Break from loop
    
    try:
        number = int(user_input)
        print(f"Square: {number ** 2}")
    except:
        print("Invalid number")
```

## Loop Control Statements

### `skip` (Continue)
The `skip` statement jumps to the next iteration of the loop.

```python
// Skip even numbers
for i in range(10):
    if i % 2 == 0:
        skip  // Continue to next iteration
    print(f"Odd number: {i}")

// Skip empty strings
words = ["hello", "", "world", "", "carrion"]
for word in words:
    if word == "":
        skip
    print(word.upper())
```

### `stop` (Break)
The `stop` statement exits the loop immediately.

```python
// Find first negative number
numbers = [5, 3, 8, -2, 1, 7]
for num in numbers:
    if num < 0:
        print(f"Found negative: {num}")
        stop  // Exit loop
    print(f"Positive: {num}")

// Interactive menu
while True:
    choice = input("Enter choice (1-3, or 'q' to quit): ")
    if choice == "q":
        stop
    elif choice == "1":
        print("Option 1 selected")
    elif choice == "2":
        print("Option 2 selected")
    elif choice == "3":
        print("Option 3 selected")
    else:
        print("Invalid choice")
```

## Pattern Matching

### Match Statement
The `match` statement provides pattern matching capabilities.

```python
match value:
    case pattern1:
        // code for pattern1
    case pattern2:
        // code for pattern2
    _:  // default case
        // default code
```

#### Examples
```python
// Basic pattern matching
status_code = 404
match status_code:
    case 200:
        message = "OK"
    case 404:
        message = "Not Found"
    case 500:
        message = "Internal Server Error"
    _:
        message = "Unknown Status"

print(message)

// Pattern matching with strings
command = "save"
match command:
    case "save":
        print("Saving file...")
    case "load":
        print("Loading file...")
    case "quit":
        print("Goodbye!")
        return
    _:
        print("Unknown command")

// Pattern matching with multiple values
day = "Monday"
weather = "sunny"
match (day, weather):
    case ("Monday", "sunny"):
        activity = "Go for a walk"
    case ("Monday", "rainy"):
        activity = "Work indoors"
    case ("Friday", _):  // Any weather on Friday
        activity = "Plan weekend"
    _:
        activity = "Normal routine"

print(f"Today's activity: {activity}")
```

## Resource Management

### Autoclose Statement
The `autoclose` statement provides automatic resource management for objects that need cleanup, such as files, network connections, or database handles.

```python
autoclose expression as variable:
    // code block
    // variable.close() is called automatically
```

#### File Operations with Autoclose
```python
// Reading files - file is automatically closed
autoclose open("data.txt", "r") as file:
    content = file.read()
    lines = content.split("\n")
    for line in lines:
        print(line)

// Writing files - file is automatically closed
autoclose open("output.txt", "w") as file:
    file.write("Header\n")
    for i in range(5):
        file.write(f"Line {i}\n")

// Appending to files - file is automatically closed
autoclose open("log.txt", "a") as file:
    timestamp = get_current_time()
    file.write(f"[{timestamp}] Application started\n")
```

#### Autoclose with Error Handling
The `autoclose` statement ensures resources are cleaned up even when errors occur:

```python
attempt:
    autoclose open("important_data.txt", "r") as file:
        content = file.read()
        processed = process_complex_data(content)
        return processed
ensnare:
    print("Error processing file")
    return None
// file.close() is called automatically even if an error occurs
```

#### Nested Autoclose Operations
```python
// Copy file contents with automatic cleanup
autoclose open("source.txt", "r") as input_file:
    autoclose open("destination.txt", "w") as output_file:
        while True:
            chunk = input_file.read_chunk(1024)  // If supported
            if not chunk:
                stop
            output_file.write(chunk)
```

#### Custom Resources with Autoclose
Any object with a `close()` method can be used with `autoclose`:

```python
// Hypothetical database connection
autoclose database.connect("localhost", "mydb") as conn:
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users")
    results = cursor.fetchall()
    for row in results:
        print(row)
// conn.close() is called automatically

// Hypothetical network connection
autoclose network.connect("api.example.com", 443) as conn:
    response = conn.send_request("/api/data")
    data = response.get_json()
    process_api_data(data)
// conn.close() is called automatically
```

## Function Control Flow

### Return Statement
The `return` statement exits a function and optionally returns a value.

```python
spell function_name(parameters):
    // function body
    return value  // optional

// Examples
spell add(a, b):
    return a + b

spell greet(name):
    return f"Hello, {name}!"

spell is_positive(number):
    if number > 0:
        return True
    else:
        return False

// Early return
spell process_data(data):
    if data is None:
        return "Error: No data provided"
    
    if len(data) == 0:
        return "Error: Empty data"
    
    // Process data
    result = data.upper()
    return f"Processed: {result}"

// Function using autoclose for file operations
spell read_config(filename):
    autoclose open(filename, "r") as file:
        content = file.read()
        return parse_config(content)
```

## Nested Control Flow

### Nested Loops
```python
// Multiplication table
for i in range(1, 6):
    for j in range(1, 6):
        product = i * j
        print(f"{i} x {j} = {product}")
    print()  // Empty line after each table

// Finding items in nested structure
matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
target = 5
found = False

for row in matrix:
    for item in row:
        if item == target:
            print(f"Found {target}")
            found = True
            stop  // Break from inner loop
    if found:
        stop  // Break from outer loop
```

### Nested Conditionals
```python
// Grade calculation with multiple criteria
score = 87
attendance = 95

if score >= 90:
    if attendance >= 90:
        grade = "A"
    else:
        grade = "A-"
otherwise score >= 80:
    if attendance >= 90:
        grade = "B+"
    otherwise attendance >= 80:
        grade = "B"
    else:
        grade = "B-"
else:
    if attendance >= 90:
        grade = "C+"
    else:
        grade = "C"

print(f"Final grade: {grade}")
```

## Advanced Control Flow Patterns

### Loop with Multiple Conditions
```python
// Process items until condition met
items = [1, 3, 5, 7, 9, 12, 15]
sum_total = 0
index = 0

while index < len(items) and sum_total < 20:
    sum_total += items[index]
    print(f"Added {items[index]}, total: {sum_total}")
    index += 1

print(f"Final total: {sum_total}")
```

### Conditional Expressions (Ternary)
```python
// Ternary-like expressions using if-else
age = 20
status = "adult" if age >= 18 else "minor"

// More complex conditional expressions
max_value = a if a > b else b
message = "positive" if number > 0 else "negative" if number < 0 else "zero"
```

### Exception-like Flow with Match
```python
// Simulating exception handling with match
spell divide(a, b):
    match b:
        case 0:
            return "Error: Division by zero"
        _:
            return a / b

result = divide(10, 0)
match result:
    case str if "Error" in result:
        print(f"Operation failed: {result}")
    _:
        print(f"Result: {result}")
```

## Best Practices

### Readable Conditions
```python
// Good: Clear variable names
is_valid_age = age >= 18 and age <= 65
has_permission = user.is_admin() or user.has_role("editor")

if is_valid_age and has_permission:
    process_request()

// Good: Extract complex conditions
spell is_business_hours():
    current_hour = get_current_hour()
    return 9 <= current_hour <= 17

if is_business_hours():
    handle_business_request()
```

### Avoid Deep Nesting
```python
// Avoid: Deep nesting
if user is not None:
    if user.is_active():
        if user.has_permission("read"):
            if document.is_available():
                return document.content

// Better: Early returns
if user is None:
    return "User not found"

if not user.is_active():
    return "User inactive"

if not user.has_permission("read"):
    return "Permission denied"

if not document.is_available():
    return "Document unavailable"

return document.content
```

### Use Match for Multiple Conditions
```python
// Better than long if-elif chains
match user_type:
    case "admin":
        return admin_dashboard()
    case "editor":
        return editor_dashboard()
    case "viewer":
        return viewer_dashboard()
    _:
        return login_page()
```

Control flow in Carrion provides powerful and flexible ways to direct program execution, with clear syntax and familiar patterns that make code readable and maintainable.