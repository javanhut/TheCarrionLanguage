# Carrion Language Documentation
## Spells

> **Because coding is magic, right?** In Carrion, **spells** are your basic function definitions. Why “spells,” you ask? Well, coding is pretty magical if you think about it—conjuring solutions out of thin air with a few lines of logic.

---

### Casting a Spell (Function Definition)

A **spell** in Carrion defines a piece of functionality that you can reuse. This is particularly helpful when you need to perform the same task multiple times with different values. Here’s the syntax:

```python
spell function_name(parameters):
    # Indented lines make up the body of the spell
    # Perform some magical operations here
```

Just like in Python, Carrion relies on **indentation** to determine the scope (body) of the spell. That means every line inside the spell must be indented at the same level.

#### Example:

```python
spell say_hello(name):
    print("Hello,", name)
```

Here’s a simple spell, `say_hello`, that takes in a `name` parameter and prints out a friendly greeting.

---

### The Anatomy of a Spell

1. **Keyword**: `spell` – Our incantation that tells Carrion we’re defining a function.
2. **Name**: `function_name` – A unique identifier for your spell. Remember, your code might blow up if two spells share the same name (unless you’re trying some advanced wizardry).
3. **Parameters**: `(parameters)` – A comma-separated list of variables your spell can use. They’re optional, so if your spell doesn’t need inputs, you can leave ’em out.
4. **Body**: Everything **indented** under the spell definition. These lines describe what the spell does.

---

### Indentation Magic

Because indentation is how Carrion (and Python) figures out what belongs to the spell, be careful with your spaces or tabs. One slip-up in indentation, and your code might not work as you expect—or at all.

**Tip**: Use a consistent indentation style, like four spaces or a single tab. Just don’t mix them up, or you’ll summon the dreaded indentation errors!

---

### Spells Inside and Outside Spellbooks

Spells can live:

- **Independently** in your code, ready to be called at any time.
- **Inside spellbooks**, which are Carrion’s way of grouping related spells (kind of like modules or libraries in other languages).

You can import a [Spellbook](../Spellbooks/README.md) wherever you need it to harness its arcane power—er, functionality.  

---

### Example: A Spell with a Return

Although not explicitly required, you might want your spell to return something. Here’s what that looks like in Carrion (not too different from what you’d expect in Python):

```python
spell add_two_numbers(x, y):
    result = x + y
    return result
```

**Note**: The `return` statement conjures up a value back to wherever the spell was called.

---

### Wrapping Up

So there you have it—Carrion spells in a nutshell. They’re just function definitions, wrapped in a bit of playful mysticism to keep coding fun and whimsical. As you continue your Carrion journey, you’ll find spells invaluable for repeating tasks, organizing your code, and casting them at will from anywhere in your project or spellbooks.

**Go forth and conjure powerful spells!**  

