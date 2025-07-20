# Type System

The Carrion language supports optional type hints for function parameters, return values, and variable assignments. This system helps with code clarity and enables future type checking capabilities.

## String Concatenation and Type Consistency

Carrion's string concatenation operations now properly maintain type consistency. When performing string concatenation operations (using the `+` operator), the result will always be a properly wrapped String instance that has access to all String grimoire methods.

### String Concatenation Behavior

```carrion
# Regular string concatenation
greeting = "Hello" + " World"
print(greeting.upper())  # Works correctly with String methods

# Concatenation with converted types
number = 42
message = "The answer is: " + str(number)
print(message.length())  # Access to String methods maintained

# Triple-quoted strings in concatenation
html_content = """<html>
<body>
    <h1>Hello World</h1>
</body>
</html>"""

response = "HTTP/1.1 200 OK\r\n\r\n" + html_content
# Result is a proper String instance with method access
```

### Previous Issues (Fixed)

In earlier versions, long string concatenations or concatenations involving triple-quoted strings could result in BUILTIN type objects instead of proper String instances. This has been resolved, ensuring that:

- All string concatenation operations return proper String instances
- String methods remain accessible on concatenated results
- Socket operations and other modules that expect string types work correctly

## Type Hint Syntax

### Function Parameters

Function parameters can include type hints using the colon (`:`) syntax:

```carrion
spell calculate_area(width: int, height: int):
    return width * height
```

Parameters can also have default values:

```carrion
spell greet(name: str = "World"):
    print("Hello, " + name)
```

### Return Type Hints

Functions can specify their return type using the arrow (`->`) syntax:

```carrion
spell add(a: int, b: int) -> int:
    return a + b

spell get_user_info(id: int) -> dict:
    return {"id": id, "name": "John Doe"}
```

### Variable Type Hints

Variables can include type hints in assignment statements:

```carrion
count: int = 0
name: str = "Alice"
scores: list = [85, 92, 78]
```

## Supported Type Annotations

The type system currently supports the following type annotations:

- **Primitive Types**: `int`, `float`, `str`, `bool`
- **Collection Types**: `list`, `dict`, `set`
- **Special Types**: `None`, `any`
- **Custom Types**: Grimoire class names

## Complex Type Examples

### Functions with Multiple Parameters

```carrion
spell process_data(items: list, multiplier: float = 1.0) -> list:
    result = []
    for item in items:
        result.append(item * multiplier)
    return result
```

### Grimoire Methods with Type Hints

```carrion
grim Calculator:
    spell init(self, precision: int = 2):
        self.precision = precision
    
    spell divide(self, a: float, b: float) -> float:
        if b == 0:
            raise ValueError("Division by zero")
        return round(a / b, self.precision)
```

### Nested Functions

```carrion
spell outer_function(x: int) -> spell:
    spell inner_function(y: int) -> int:
        return x + y
    return inner_function
```

## Type Hint Behavior

Currently, type hints in Carrion are:

1. **Optional**: You can write code without any type hints
2. **Documentation**: They serve as documentation for developers
3. **Non-enforcing**: The interpreter does not enforce type checking at runtime
4. **Preparatory**: They lay the groundwork for future static type checking

## Best Practices

1. **Use type hints for public APIs**: Add type hints to functions that will be used by others
2. **Be consistent**: If you start using type hints in a module, use them throughout
3. **Document complex types**: For complex return types, consider adding docstrings
4. **Keep it simple**: Don't over-annotate obvious cases

## Examples

### Basic Function with Type Hints

```carrion
spell factorial(n: int) -> int:
    if n <= 1:
        return 1
    return n * factorial(n - 1)
```

### Grimoire with Typed Methods

```carrion
grim BankAccount:
    spell init(self, initial_balance: float = 0.0):
        self.balance: float = initial_balance
    
    spell deposit(self, amount: float) -> None:
        if amount > 0:
            self.balance += amount
    
    spell withdraw(self, amount: float) -> bool:
        if amount <= self.balance:
            self.balance -= amount
            return True
        return False
    
    spell get_balance(self) -> float:
        return self.balance
```

### Function with Multiple Return Types

```carrion
spell parse_number(value: str) -> any:
    attempt:
        if "." in value:
            return float(value)
        else:
            return int(value)
    ensnare:
        return None
```

## Future Enhancements

The type hint system is designed to be extended in the future with:

- Runtime type checking (opt-in)
- Static type analysis tools
- Generic types (e.g., `list[int]`, `dict[str, float]`)
- Union types (e.g., `int | float`)
- Type aliases
- Protocol/interface definitions

## Integration with Development Tools

Type hints enable better integration with:

- IDE autocompletion
- Static analysis tools
- Documentation generators
- Linting tools

The type system is designed to be gradually adopted, allowing developers to add type hints incrementally to existing codebases without breaking compatibility.