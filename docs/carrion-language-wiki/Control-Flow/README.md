# Carrion Language Documentation
## Control Flow

So you want to manipulate how your program runs? Carrion’s got you covered with **if-else** statements, **while** loops, **for** loops, and **match case** statements. If you’ve dabbled in other languages, some of this will look familiar—**but** we’ve sprinkled a bit of our own Carrion charm in. Let’s dive in!

---

## IF/ELSE

Carrion’s **if** statement is basically what you’d expect, but with a teeny twist in how you handle additional conditions. Instead of `elif` (Python) or `else if` (C/C++/Java/...bleh) we use **`otherwise`**:

```python
if condition:
    # do something
otherwise some_other_condition:
    # do something else
else:
    # the fallback if none of the above conditions are true
```

**Translation**:
- **`if`**: your main condition check.
- **`otherwise`**: Carrion’s stylish stand-in for `elif`.
- **`else`**: the final “catch-all.”

This should read nicely, even if you’re new to programming!

---

## Loops

### While Loops

A **while** loop will run as long as its condition is true.

```python
x = 0

while x < 5:
    x++
```

**Yup, that’s `x++`**. We support `++` out of sheer nostalgia for C/C++—plus, it’s kind of fun. If you prefer the more Pythonic approach:

```python
x = 0

while x < 5:
    x += 1
```

We won’t judge. Both ways work. Aren’t we magnanimous?

---

### For Loops

Carrion’s **for** loops follow a Python-esque style: `for element in collection` or `for i in range(...)`. 

```python
arr = [1, 2, 3, 4, 5, 6]

for i in arr:
    print(i)
```

Or iterate over a range:

```python
for i in range(10):
    print(i)
```

**Note**: Because we can reuse that beloved `range` function from Python-like constructs, you can easily generate sequences for your loops.

---

## Match Case

Think of **match case** like a more elegant switch statement. If you’re accustomed to Python’s pattern matching or a switch-case from other languages, this’ll feel right at home.

```python
foo = 10

match foo:
    case 10:
        # do something
    case 20:
        # do something else
    _:
        print("This is the default statement")
```

Breakdown:
- `match foo`: We’re matching the variable `foo`.
- `case 10`: If `foo` is `10`, do that block of code.
- `case 20`: If `foo` is `20`, do something else.
- `_`: The “wildcard” case—like a default or else block for everything not listed above.

**Bonus**: You can match on all sorts of data types—strings, tuples, maybe even complex pattern matching if the language evolves further. If you don’t match anything, `_` is there to catch you.

---

## Extra Goodies: break & continue

Although not explicitly shown above, **Carrion** also provides the classic loop control keywords (if you’re coming from other languages, you’ll feel right at home):

- **`stop`**: Exits the nearest loop immediately.
- **`skip`**: Skips the rest of the current loop iteration and jumps straight to the next iteration check.

```python
for i in range(10):
    if i == 5:
        stop       # stops the loop entirely
    print(i)
```

```python
for i in range(10):
    if i % 2 == 0:
        skip    # skip even numbers
    print(i)        # prints only odd numbers
```

So if you’re feeling dramatic and want to bail on a loop halfway through, or skip a few iterations, these will come in handy. Plus the synatax makes sense right?

---

## Summary

Carrion’s control flow statements will look pretty familiar to anyone with Python, C, or even some other language experience. That said, we’ve sprinkled a bit of playful flavor:

- **if / otherwise / else**: Because we can.  
- **while** loop with either `x++` or `x += 1` (flexibility is magical).  
- **for** loops with a Python-esque `for x in something:` approach.  
- **match case** for clean and readable branching—plus a default `_` case.  

Mix and match them however you please—**just don’t summon any infinite loops** unless you’re in for a crash course in killing processes.

**Have fun controlling the flow of your code—like a wizard controlling chaotic energies!**  
