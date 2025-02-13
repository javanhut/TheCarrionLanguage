# Encapsulation in Carrion

**Encapsulation** in Carrion means controlling access to data and methods within a Spellbook. Carrion adopts naming conventions similar to Python:

- `_private_spell`: Methods or attributes intended for internal use.
- `__protected_spell`: Methods or attributes that are also internal but with a stronger emphasis on restricted access.

## Basic Example

```python
spellbook Foo:
    init(var="foobar"):
        self.var = var

    // A protected method
    spell __protected_spell():
        return self.var

    // A private method
    spell _private_spell():
        return "Private: " + self.var

    // Public method to demonstrate accessing protected/private
    spell reveal():
        return str(self.__protected_spell()) + " | " + str(self._private_spell())
```

In this example:

__protected_spell and _private_spell cannot be called from an object instance directly outside Foo.

reveal() acts as a public interface to those hidden spells.


Key Takeaway: Encapsulation in Carrion helps keep your Spellbook’s data consistent and safe from unintended external changes.



