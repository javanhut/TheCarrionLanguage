# Polymorphism in Carrion

**Polymorphism** is the ability to override or redefine spells from a parent Spellbook, letting each child Spellbook have its own version of a spell.

## Polymorphism Example

```python
spellbook Parent:
    init(name="parent"):
        self.name = name

    spell describe():
        return f"Parent: {self.name}"

spellbook Child(Parent):
    init(name="child"):
        self.name = name
    
    // Overriding the parent's 'describe' spell
    spell describe(child_age=0):
        return f"Name: {self.name}\nAge: {child_age}"
```

When you call describe() on a Child instance, it will use the child’s version of describe, instead of the Parent’s. This allows the same method name to work differently depending on the actual Spellbook instance.

Key Point: Polymorphism provides flexibility and maintainability, making code less repetitive and easier to extend.