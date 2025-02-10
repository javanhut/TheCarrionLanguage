```markdown
# Functions

Carrion comes with **two main varieties** of functions:

1. **Spells** (user-defined functions)
2. **Built-In Functions** (pre-packaged magic)

For more information on **Spells** (how to define and use them), check out our [Spell Definitions](../Spells/README.md).

---

## Built-In Functions

Built-In Functions are baked right into the language itself and **are not** part of the Munin Standard Library. They handle core tasks like I/O, type conversion, file operations, and more. If you need to get work done quickly—reading files, converting data, or even running system commands—these are your go-to tools.

Below is a quick overview of the built-in functions available:

- **`len`** – Get the length of strings or arrays.
- **`print`** – Print arguments to standard output.
- **`input`** – Prompt for user input (optional prompt).
- **`type`** – Return the type of an object.
- **`int`** – Convert a string/float to an integer.
- **`float`** – Convert a string/int to a float.
- **`str`** – Convert just about anything into a string.
- **`list`** – Convert a string or tuple into an array.
- **`tuple`** – Convert an array into a tuple.
- **`range`** – Generate a sequence of integers (like Python’s `range`).
- **`osRunCommand`** – Execute external commands with optional args/capture.
- **`osGetEnv`** – Retrieve an environment variable.
- **`osSetEnv`** – Set an environment variable.
- **`osGetCwd`** – Get current working directory.
- **`osChdir`** – Change the current working directory.
- **`osSleep`** – Pause execution for X seconds.
- **`osListDir`** – List the contents of a directory.
- **`osRemove`** – Remove a file or directory.
- **`osMkdir`** – Create a new directory (with optional permissions).
- **`osExpandEnv`** – Expand environment variables in a string.
- **`fileRead`** – Read file content into a string.
- **`fileWrite`** – Write/overwrite file content.
- **`fileAppend`** – Append content to a file.
- **`fileExists`** – Check if a file or directory exists.
- **`Error`** – Construct a custom error object.

---

# Notes to keep in mind:
The builtin functions for OS and I/O are wrapped in the Munin Standard Library so they just need to be called in. No need to remeber them all.
 Link to Munin Standard Library -> [Standard Library](../Munin-Standard-Library/README.md)


**Happy casting!**
```
