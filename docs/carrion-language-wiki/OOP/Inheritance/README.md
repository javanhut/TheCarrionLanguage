# Inheritance in Carrion

**Inheritance** lets one Spellbook derive from another, sharing or extending functionality. Carrion’s syntax is similar to Python, using parentheses after a child’s Spellbook name to specify the parent.

## Basic Example

```python
spellbook Parent:
    spell parent_method():
        return "I am the Parent!"

spellbook Child(Parent):
    spell child_method():
        return "I am the Child!"

child_instance = Child()
print(child_instance.parent_method())  // Inherited from Parent
print(child_instance.child_method())   // Defined in Child
```

Here, Child can use all spells declared in Parent, making it easy to build upon existing functionality.

Overriding the Parent’s init

If you need a custom constructor, you can declare an init in the child, just like in Python:
```python
spellbook Parent:
    init(name):
        self.name = name

spellbook Child(Parent):
    init(name, age):
        self.name = name
        self.age = age
        // Optionally call parent's init or do something extra
```

Inheritance helps you create a clear hierarchy of Spellbooks in Carrion.
