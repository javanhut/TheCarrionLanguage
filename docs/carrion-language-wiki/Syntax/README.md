# Carrion Language Documentation

## Basic Syntax

Carrion’s syntax takes inspiration from various programming languages—particularly Python—so it should feel somewhat familiar. However, there are unique aspects to Carrion that set it apart. If you have not done so already, you can read more about [Data Types in Carrion](../Data-Types/README.md).

Below is a quick overview of Carrion’s basic syntax, including variable handling, numeric literals, and operators.

---

### Number Handling – Integer Literals and Float Literals

In Carrion, numerical values are automatically recognized as either **integer literals** or **float literals**:

- **Integer Literals** represent whole numbers (positive, negative, and zero).  
  Examples: `-1`, `0`, `1`
- **Float Literals** represent decimal numbers (again, positive, negative, and zero).  
  Examples: `-1.00`, `0.11`, `1.000`

You can simply type these values into the REPL or in your code, and the language will correctly interpret them.

---

### Variables

Carrion does not require explicit type definitions for variables. Variables are **dynamically typed**, which means you can reassign them to different types without declaring a type:

```python
variable = 10
variable = "name"
```

Although this offers flexibility, it can lead to potential confusion regarding variable types. If you prefer, you can use **type hints** to clarify or enforce a variable’s type:

```python
variable: int = 10
variable: str = "string"
```

At this time, type hints are primarily syntactic, but they will play a role in performance optimizations once the JIT compiler or VM is fully implemented.

---

### Basic Operators

Carrion supports the following standard arithmetic operators:

- `+`  : Addition
- `-`  : Subtraction
- `*`  : Multiplication
- `/`  : Division
- `%`  : Modulus

#### Assignment Operators

- `=`   : Basic assignment
- `+=`  : Increment assignment (adds and then assigns)
- `-=`  : Decrement assignment (subtracts and then assigns)
- `*=`  : Multiplication assignment (multiplies and then assigns)
/ `=`   : Division assignment (divides and then assigns)

#### Comparison Operators

- `>`   : Greater Than
- `<`   : Less Than
- `>=`  : Greater Than or Equal
- `<=`  : Less Than or Equal
- `!=`  : Not Equal

*Note: The logical operators `and` and `or` are scheduled for the minor update **0.1.5** and may not be available in earlier versions.*

---

**Happy coding with Carrion!**
