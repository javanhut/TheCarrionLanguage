# Abstraction in Carrion

**Abstraction** in Carrion is achieved by marking a Spellbook as `arcane` (similar to "abstract" in other languages). You can also declare certain spells as `@arcanespell` (abstract methods), which must be implemented by child Spellbooks.

## Arcane Spellbooks and Arcanespells

```python
arcane spellbook BaseSpellbook:
    @arcanespell
    spell some_method():
        ignore  // No implementation

spellbook Implementation(BaseSpellbook):
    spell some_method():
        return "I have implemented the abstract method!"
```


1. arcane spellbook indicates no instances of this Spellbook should be created directly.


2. @arcanespell means the spell has no implementation in the arcane Spellbook—child Spellbooks must define it.


3. ignore tells Carrion the method is intentionally empty.



Benefit: Abstraction ensures that derived Spellbooks implement critical methods without providing direct implementation in the base Spellbook.
