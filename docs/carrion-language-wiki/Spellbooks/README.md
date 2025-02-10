# Carrion Language Documentation
## Spellbooks

> **Classes? Ugh.** Let’s call them **spellbooks** instead. Other languages like C++, Python, or ... **Java** (yeah, let’s just move on) refer to these as “classes.” But we have a theme going, and we’re not breaking it. End of story.

---

### What’s a Spellbook?

A **spellbook** in Carrion is a collection of **spells** grouped under one magical entity. Think of it as a neat binder that keeps all related spells together. This makes your code organized and your spells easier to find—no rummaging through code like a stray cat in a dumpster.

---

### Declaring a Spellbook

In Carrion, you create a spellbook with the keyword **`spellbook`** followed by the spellbook’s name:

```python
spellbook Example:
    """This is an example spellbook"""
    # ...
```

#### Example Structure

```python
spellbook Example:
    """This is an example spellbook"""

    init constructor_name(parameters):
        self.parameters = parameters  # store parameters on "self" for instance use

    spell other_related_spell():
        # do some magical stuff here
        print("This spell is part of the Example spellbook!")
```

1. **Keyword**: `spellbook` – Our magic incantation to declare a new spellbook.
2. **Name**: `Example` – A unique identifier for your spellbook.
3. **Body**: Everything **indented** under the spellbook definition.

---

### The `init` Spell (Constructor)

Inside a spellbook, you can define an **`init`** spell (think “constructor”) to handle any setup when an instance of this spellbook is created. It’s like preparing your magical arsenal before you actually start slinging spells:

```python
init constructor_name(parameters):
    self.parameters = parameters
```

- **`init constructor_name(...)`**: The syntax you use to define the constructor.  
- **`self.parameters = parameters`**: Stores the passed-in parameters on the instance (so you can use them later in your spells).

**Note**: The `self` reference works similarly to how `this` or `self` works in other OOP languages, but we’re keeping it mystical here.

---

### Spells Within a Spellbook

Any spells you define inside a spellbook are associated with that spellbook’s instances. In Python-speak, they’re “instance methods.” In Carrion-speak, they’re just **spells** living inside a **spellbook**:

```python
spell other_related_spell():
    # conditions or logic here
```

#### A Note on `self` in Spells
If you need to access instance-level data (stuff stored in `init`), include a `self` parameter in the spell definition. If you don’t need it, don’t bother with `self` at all—Carrion doesn’t force you to, which is **amazing**.

---

### Usage

Here’s a quick how-to on using your **spellbook**:

```python
# Creating an instance of the Example spellbook
my_example = Example("some parameter")

# Calling the spells defined in the spellbook
my_example.other_related_spell()
```

Under the hood, `init constructor_name("some parameter")` runs when you create `my_example`, setting up any instance data you need.

---

### Further Reading

For more advanced uses, including inheritance and other OOP goodies, check out our [Object Oriented Programming](../OOP/README.md) section of the wiki.

**Remember**: The magic is in your hands. Use it wisely... or irresponsibly—who are we to judge?

