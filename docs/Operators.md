# Operators and Expressions

Carrion provides a comprehensive set of operators for mathematical calculations, logical operations, comparisons, and assignments.

## Arithmetic Operators

### Basic Arithmetic
| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `+` | Addition | `5 + 3` | `8` |
| `-` | Subtraction | `5 - 3` | `2` |
| `*` | Multiplication | `5 * 3` | `15` |
| `/` | Division | `15 / 3` | `5.0` |
| `//` | Integer Division | `17 // 3` | `5` |
| `%` | Modulo | `17 % 3` | `2` |
| `**` | Exponentiation | `2 ** 3` | `8` |

### String and Array Multiplication
| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `*` | String repetition | `"hello" * 3` | `"hellohellohello"` |
| `*` | Array repetition | `[1, 2] * 3` | `[1, 2, 1, 2, 1, 2]` |

### Unary Operators
| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `+` | Unary plus | `+5` | `5` |
| `-` | Unary minus | `-5` | `-5` |

```python
// Basic arithmetic
result = 10 + 5 * 2    // → 20 (follows order of operations)
power = 2 ** 3         // → 8
remainder = 17 % 5     // → 2

// String and array multiplication
border = "=" * 50      // → "=================================================="
repeated = [0] * 5     // → [0, 0, 0, 0, 0]
pattern = ["a", "b"] * 3  // → ["a", "b", "a", "b", "a", "b"]

// Multiplication works both ways
hello3 = "hello" * 3   // → "hellohellohello"
hello3_alt = 3 * "hello"  // → "hellohellohello"

// Unary operators
positive = +42         // → 42
negative = -42         // → -42

// Works with variables from tuple unpacking
x, y = (10, 20)
sum = x + y           // → 30
negated = -x          // → -10
```

## Assignment Operators

### Basic Assignment
| Operator | Description | Example |
|----------|-------------|---------|
| `=` | Basic assignment | `x = 5` |

### Compound Assignment
| Operator | Description | Example | Equivalent |
|----------|-------------|---------|------------|
| `+=` | Addition assignment | `x += 3` | `x = x + 3` |
| `-=` | Subtraction assignment | `x -= 3` | `x = x - 3` |
| `*=` | Multiplication assignment | `x *= 3` | `x = x * 3` |
| `/=` | Division assignment | `x /= 3` | `x = x / 3` |

### Increment and Decrement
| Operator | Description | Example | Behavior |
|----------|-------------|---------|----------|
| `++` | Prefix increment | `++x` | Increment then return |
| `++` | Postfix increment | `x++` | Return then increment |
| `--` | Prefix decrement | `--x` | Decrement then return |
| `--` | Postfix decrement | `x--` | Return then decrement |

```python
// Compound assignment
x = 10
x += 5        // x is now 15
x *= 2        // x is now 30
x /= 3        // x is now 10

// Increment/decrement operators
a = 5
result1 = ++a  // a becomes 6, result1 is 6
result2 = a++  // result2 is 6, a becomes 7

b = 10
result3 = --b  // b becomes 9, result3 is 9
result4 = b--  // result4 is 9, b becomes 8
```

## Comparison Operators

| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `==` | Equal | `5 == 5` | `True` |
| `!=` | Not equal | `5 != 3` | `True` |
| `<` | Less than | `3 < 5` | `True` |
| `>` | Greater than | `5 > 3` | `True` |
| `<=` | Less than or equal | `3 <= 5` | `True` |
| `>=` | Greater than or equal | `5 >= 5` | `True` |

```python
// Numeric comparisons
print(10 > 5)     // → True
print(3 <= 3)     // → True
print(7 != 8)     // → True

// String comparisons (lexicographic)
print("apple" < "banana")  // → True
print("hello" == "hello")  // → True

// Chained comparisons
age = 25
valid = 18 <= age < 65     // → True
```

## Logical Operators

| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `and` | Logical AND | `True and False` | `False` |
| `or` | Logical OR | `True or False` | `True` |
| `not` | Logical NOT | `not True` | `False` |

### Short-Circuit Evaluation
Logical operators use short-circuit evaluation:
- `and`: If the left operand is false, the right operand is not evaluated
- `or`: If the left operand is true, the right operand is not evaluated

```python
// Basic logical operations
has_permission = True
is_admin = False

can_edit = has_permission and is_admin     // → False
can_view = has_permission or is_admin      // → True
cannot_edit = not can_edit                 // → True

// Short-circuit evaluation
x = 10
result = x > 5 and x < 20                  // → True
result = x < 5 or x > 15                   // → False

// Practical usage
if age >= 18 and has_id:
    print("Can enter")
```

## Membership Operators

| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `in` | Membership test | `"a" in "apple"` | `True` |
| `not in` | Negative membership | `"z" not in "apple"` | `True` |

The `in` operator works with all iterable types and checks for membership:

```python
// String membership (substring and character checking)
letter = "a"
word = "banana"
print(letter in word)      // → True
print("ana" in word)       // → True (substring)
print("z" not in word)     // → True

// Array membership (element checking)
numbers = [1, 2, 3, 4, 5]
mixed = [1, "hello", 3.14, True]

print(3 in numbers)        // → True
print(6 not in numbers)    // → True
print("hello" in mixed)    // → True
print(3.14 in mixed)       // → True

// Hash membership (key checking)
data = {"name": "John", "age": 30, "active": True}
print("name" in data)      // → True
print("email" not in data) // → True
print("John" in data)      // → False (values are not checked)

// Nested structure membership
matrix = [[1, 2], [3, 4], [5, 6]]
print([3, 4] in matrix)    // → True

// Works with all data types
booleans = [True, False, True]
print(False in booleans)   // → True
```

## Bitwise Operators

| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `&` | Bitwise AND | `5 & 3` | `1` |
| `\|` | Bitwise OR | `5 \| 3` | `7` |
| `^` | Bitwise XOR | `5 ^ 3` | `6` |
| `~` | Bitwise NOT | `~5` | `-6` |
| `<<` | Left shift | `5 << 1` | `10` |
| `>>` | Right shift | `10 >> 1` | `5` |

```python
// Bitwise operations on integers
a = 5    // Binary: 101
b = 3    // Binary: 011

and_result = a & b    // → 1 (Binary: 001)
or_result = a | b     // → 7 (Binary: 111)
xor_result = a ^ b    // → 6 (Binary: 110)
not_result = ~a       // → -6 (Two's complement)

// Bit shifting
left_shift = a << 2   // → 20 (Binary: 10100)
right_shift = a >> 1  // → 2 (Binary: 10)
```

## Operator Precedence

Operators are evaluated in the following order (highest to lowest precedence):

1. **Parentheses**: `()`
2. **Exponentiation**: `**`
3. **Unary operators**: `+`, `-`, `not`, `~`
4. **Multiplicative**: `*`, `/`, `//`, `%`
5. **Additive**: `+`, `-`
6. **Shift**: `<<`, `>>`
7. **Bitwise AND**: `&`
8. **Bitwise XOR**: `^`
9. **Bitwise OR**: `|`
10. **Comparison**: `<`, `<=`, `>`, `>=`, `==`, `!=`, `in`, `not in`
11. **Logical NOT**: `not`
12. **Logical AND**: `and`
13. **Logical OR**: `or`
14. **Assignment**: `=`, `+=`, `-=`, `*=`, `/=`

```python
// Operator precedence examples
result = 2 + 3 * 4        // → 14 (not 20)
result = (2 + 3) * 4      // → 20
result = 2 ** 3 * 4       // → 32 (exponentiation first)
result = not False and True  // → True
```

## Special Operators

### Dot Operator (Member Access)
| Operator | Description | Example |
|----------|-------------|---------|
| `.` | Member access | `object.method()` |

```python
// Method calls on objects
text = "hello"
uppercase = text.upper()     // → "HELLO"

numbers = [1, 2, 3]
length = numbers.length()    // → 3

// Chained method calls
result = "  hello  ".strip().upper()  // → "HELLO"
```

### Hash Operator
| Operator | Description | Usage |
|----------|-------------|-------|
| `#` | Hash operator | Special operations |

### At Symbol (Decorators)
| Operator | Description | Usage |
|----------|-------------|-------|
| `@` | Decorator symbol | `@arcanespell` |

```python
// Decorator usage in abstract methods
arcane grim AbstractClass:
    @arcanespell
    spell abstract_method():
        ignore
```

## Expression Examples

### Complex Expressions
```python
// Mathematical expressions
area = 3.14159 * radius ** 2
discriminant = b ** 2 - 4 * a * c
average = (sum_total / count) if count > 0 else 0

// Boolean expressions
is_valid = age >= 18 and has_license and not is_suspended
can_proceed = (user.is_admin() or user.has_permission("write")) and not system.is_locked()

// Mixed expressions with method calls
formatted = f"Result: {(value * 1.1).round(2)}"
processed = text.lower().strip().replace(" ", "_")
```

### Type Checking in Expressions
```python
// Using type() in expressions
if type(value) == "INTEGER" and value > 0:
    result = value.to_bin()

// Combining operators with built-ins
max_length = max(len(s) for s in string_list)
all_positive = all(num > 0 for num in numbers)
```

### Tuple Unpacking in Expressions
```python
// Multiple assignment
x, y = (10, 20)
a, b, c = range(3)
first, *rest = [1, 2, 3, 4, 5]

// Using unpacked values in expressions
distance = ((x2 - x1) ** 2 + (y2 - y1) ** 2) ** 0.5
swapped = (y, x)  // Swap values
```

All operators work correctly with both primitive types and Instance-wrapped values from tuple unpacking, providing consistent behavior throughout the language.