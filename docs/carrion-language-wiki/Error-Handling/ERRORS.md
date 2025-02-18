# Error Handling

For Handling errors in Carrion is relatively simple.

There are a few parts to handling errors accurately so i'll go through the simple syntax first

## Creating Generic errors
Generic errors can be declared by the builtin Error function.

Ex.
```python
Error("Error Name", "Error description")
```

This is for a Generic error however there are times where you want to pass a specific Error in that case you can use a predefined error in the Munin Standard Library or Make a User Defined Error. More information on the Munin Standard Library can be found here -> (Munin StdLib)[../Munin-Standard-Library/README.md].

## Declaring Custom errors

To declare custom errors you just import the Exception Library and declare a custom Error.

```python

spellbook CustomError(Exception):
    spell init(message):
        self.message = message 

raise CustomError("Custom error description")
```

Pretty Simple so far.

Now onto the traditional Error throwing and catching.

In Carrion the syntax is alittle different You can create a Attempt, ensnare, resolve for errors.

Attempt is the intended code, Ensnare is the way to catch the error if it's throw ans can also be used as a generic if no specific error is ensnared. Resolve is the final condition that will run regardless of if the error caught or not.

ex.
```python


attempt:
    """Attempted code"""
ensnare(CustomError):
    """Custom error action"""
ensnare:
    """Generic Error """
resolve:
    """Final method to be used regardless if error throw or not"""
```


the Raise function in Carrion works to throw an error.

```python

raise CustomError("Error description")

```

The check function allows you to make a comparsion in the function and throw an error if assertion check fails.


```python
x = 9

check(x == 10, "Value Error: Expected x to equal 10")
```

Thats the simple application in Carrion.
