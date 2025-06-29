# 📖 Classes (Grimoires) - Object-Oriented Spellcrafting

Welcome to the mystical art of **Grimoire crafting**! In Carrion, classes are called **"grimoires"** (magical spellbooks) that contain **"spells"** (methods) and store **magical knowledge** (data). Master this ancient art to create powerful, reusable magical objects!

---

## 🪄 What Are Grimoires?

**Grimoires** (classes) are magical blueprints that define:
- 🏗️ **Structure** - What magical properties objects possess
- ✨ **Spells** - What magical actions objects can perform
- 🎭 **Behavior** - How objects interact with the magical world
- 🧬 **Inheritance** - How magical knowledge passes from master to apprentice

Think of a grimoire as a **template for creating magical beings** with shared characteristics and abilities!

---

## 🎭 Basic Grimoire Creation

### Your First Magical Grimoire
```python
grim MagicalCreature:
    init(name, element):
        self.name = name
        self.element = element
        self.health = 100
        self.mana = 50
    
    spell introduce():
        return f"I am {self.name}, a creature of {self.element}!"
    
    spell cast_elemental_spell():
        if self.mana >= 10:
            self.mana -= 10
            return f"✨ {self.name} casts a {self.element} spell!"
        else:
            return f"💫 {self.name} is out of mana!"

// Create magical beings from the grimoire
dragon = MagicalCreature("Smaug", "fire")
unicorn = MagicalCreature("Stardust", "light")

// Use their magical abilities
print(dragon.introduce())              // I am Smaug, a creature of fire!
print(unicorn.cast_elemental_spell())  // ✨ Stardust casts a light spell!
print(f"Dragon's mana: {dragon.mana}") // Dragon's mana: 40
```

### The `init` Spell (Constructor)
The `init` spell is a special constructor that runs when creating new magical beings:

```python
grim Wizard:
    init(name, school_of_magic, starting_level = 1):
        // Magical properties (attributes)
        self.name = name
        self.school = school_of_magic
        self.level = starting_level
        self.spells_known = []
        self.magical_items = []
        self.experience = 0
    
    spell get_title():
        if self.level >= 20:
            return f"Archmage {self.name}"
        elif self.level >= 10:
            return f"Master {self.name}"
        elif self.level >= 5:
            return f"Mage {self.name}"
        else:
            return f"Apprentice {self.name}"

// Create wizards with different starting conditions
harry = Wizard("Harry Potter", "Defense Against Dark Arts")
gandalf = Wizard("Gandalf", "Light Magic", 50)

print(harry.get_title())    // Apprentice Harry Potter
print(gandalf.get_title())  // Archmage Gandalf
```

---

## 🎪 Grimoire Spells (Methods)

### Instance Spells
These spells work with specific magical beings (instances):

```python
grim MagicalLibrary:
    init():
        self.books = []
        self.visitors = 0
        self.magical_energy = 1000
    
    spell add_book(title, author, magic_level):
        magical_book = {
            "title": title,
            "author": author,
            "magic_level": magic_level,
            "times_borrowed": 0
        }
        self.books.append(magical_book)
        print(f"📚 Added '{title}' by {author} to the library!")
    
    spell find_books_by_magic_level(min_level):
        powerful_books = []
        for book in self.books:
            if book["magic_level"] >= min_level:
                powerful_books.append(book)
        return powerful_books
    
    spell borrow_book(title):
        for book in self.books:
            if book["title"].lower() == title.lower():
                book["times_borrowed"] += 1
                self.visitors += 1
                return f"📖 You borrowed '{book['title']}'. Magic level: {book['magic_level']}"
        return f"❌ '{title}' not found in the magical library!"
    
    spell library_status():
        total_books = len(self.books)
        avg_magic = sum(book["magic_level"] for book in self.books) / total_books if total_books > 0 else 0
        return {
            "total_books": total_books,
            "visitors": self.visitors,
            "average_magic_level": avg_magic,
            "energy": self.magical_energy
        }

// Create and use a magical library
hogwarts_library = MagicalLibrary()
hogwarts_library.add_book("Advanced Transfiguration", "McGonagall", 85)
hogwarts_library.add_book("Potions Mastery", "Snape", 90)
hogwarts_library.add_book("Basic Charms", "Flitwick", 30)

powerful_books = hogwarts_library.find_books_by_magic_level(80)
print(f"Powerful books: {len(powerful_books)}")

print(hogwarts_library.borrow_book("Advanced Transfiguration"))
print(hogwarts_library.library_status())
```

### Class-level Spells (Static-like)
Spells that work with the grimoire itself, not specific instances:

```python
grim MagicalConverter:
    // Class-level magical constants
    FIRE_TO_ICE_RATIO = 0.5
    LIGHT_TO_DARK_RATIO = 1.0
    
    spell convert_fire_to_ice(fire_power):
        return fire_power * MagicalConverter.FIRE_TO_ICE_RATIO
    
    spell convert_light_to_dark(light_power):
        return light_power * MagicalConverter.LIGHT_TO_DARK_RATIO
    
    spell get_conversion_rates():
        return {
            "fire_to_ice": MagicalConverter.FIRE_TO_ICE_RATIO,
            "light_to_dark": MagicalConverter.LIGHT_TO_DARK_RATIO
        }

// Use class-level spells without creating instances
ice_power = MagicalConverter.convert_fire_to_ice(100)
print(f"100 fire power = {ice_power} ice power")

rates = MagicalConverter.get_conversion_rates()
print(f"Conversion rates: {rates}")
```

---

## 🧬 Magical Inheritance

### Basic Inheritance - Apprentice and Master
```python
// Base grimoire - the master's knowledge
grim Mage:
    init(name, magical_school):
        self.name = name
        self.school = magical_school
        self.level = 1
        self.mana = 100
        self.spells = ["Basic Magic Missile"]
    
    spell cast_spell(spell_name):
        if spell_name in self.spells:
            if self.mana >= 10:
                self.mana -= 10
                return f"✨ {self.name} casts {spell_name}!"
            else:
                return f"💫 {self.name} is out of mana!"
        else:
            return f"❌ {self.name} doesn't know {spell_name}!"
    
    spell meditate():
        self.mana = min(self.mana + 20, 100)
        return f"🧘 {self.name} meditates and recovers mana. Current: {self.mana}/100"
    
    spell learn_spell(new_spell):
        if new_spell not in self.spells:
            self.spells.append(new_spell)
            return f"📚 {self.name} learned {new_spell}!"
        else:
            return f"🤔 {self.name} already knows {new_spell}!"

// Specialized grimoire - the apprentice's path
grim FireMage(Mage):
    init(name):
        super.init(name, "School of Fire")  // Call parent constructor
        self.fire_resistance = 50
        self.spells.append("Fireball")      // Add fire-specific spell
    
    spell cast_fireball():
        return self.cast_spell("Fireball")
    
    spell fire_shield():
        self.fire_resistance += 10
        return f"🔥 {self.name} creates a fire shield! Resistance: {self.fire_resistance}"
    
    spell pyroblast():  // Unique fire mage spell
        if self.mana >= 25:
            self.mana -= 25
            return f"🔥🔥🔥 {self.name} unleashes PYROBLAST! Devastating fire magic!"
        else:
            return f"💫 Not enough mana for Pyroblast (need 25, have {self.mana})"

// Another specialized path
grim IceMage(Mage):
    init(name):
        super.init(name, "School of Ice")
        self.ice_resistance = 50
        self.spells.append("Ice Shard")
    
    spell cast_ice_shard():
        return self.cast_spell("Ice Shard")
    
    spell freeze():
        if self.mana >= 15:
            self.mana -= 15
            return f"❄️ {self.name} freezes the target solid!"
        else:
            return f"💫 Not enough mana for Freeze!"
    
    spell blizzard():  // Ultimate ice spell
        if self.mana >= 30:
            self.mana -= 30
            return f"❄️🌨️❄️ {self.name} summons a devastating BLIZZARD!"
        else:
            return f"💫 Not enough mana for Blizzard (need 30, have {self.mana})"

// Create specialized mages
pyromancer = FireMage("Ignitus")
cryomancer = IceMage("Frost")

// Test inherited abilities
print(pyromancer.cast_spell("Basic Magic Missile"))  // Inherited from Mage
print(pyromancer.cast_fireball())                    // Fire-specific spell
print(pyromancer.pyroblast())                        // Unique fire ultimate

print(cryomancer.learn_spell("Teleport"))            // Inherited learning
print(cryomancer.freeze())                           // Ice-specific spell
print(cryomancer.blizzard())                         // Ice ultimate
```

### Method Overriding - Personalizing Magic
```python
grim Animal:
    init(name, species):
        self.name = name
        self.species = species
    
    spell make_sound():
        return f"{self.name} makes a generic animal sound."
    
    spell info():
        return f"{self.name} is a {self.species}."

grim MagicalDragon(Animal):
    init(name, element):
        super.init(name, "Dragon")
        self.element = element
        self.hoard_size = 0
    
    spell make_sound():  // Override parent method
        return f"🐉 {self.name} ROARS with {self.element} magic!"
    
    spell breathe_fire():
        return f"🔥 {self.name} breathes {self.element} breath!"
    
    spell add_to_hoard(treasure):
        self.hoard_size += 1
        return f"💎 {self.name} adds {treasure} to their hoard! Total: {self.hoard_size}"

grim MagicalUnicorn(Animal):
    init(name):
        super.init(name, "Unicorn")
        self.purity_level = 100
    
    spell make_sound():  // Override with unicorn-specific sound
        return f"🦄 {self.name} whinnies melodically, spreading magical harmony!"
    
    spell heal_others():
        if self.purity_level >= 20:
            self.purity_level -= 20
            return f"💚 {self.name} uses pure magic to heal! Purity: {self.purity_level}/100"
        else:
            return f"✨ {self.name} needs to restore purity before healing again."

// Test method overriding
smaug = MagicalDragon("Smaug", "fire")
stardust = MagicalUnicorn("Stardust")

print(smaug.make_sound())      // Dragon's roar
print(stardust.make_sound())   // Unicorn's whinny
print(smaug.info())            // Inherited info method
print(stardust.heal_others())  // Unicorn-specific magic
```

---

## 🔒 Access Control & Encapsulation

### Public, Protected, and Private Magic
```python
grim SecretSpellBook:
    init(owner_name):
        self.owner = owner_name                    // Public: Anyone can access
        self._protected_spells = []                // Protected: Subclasses can access
        self.__secret_spells = ["Ultimate Power"]  // Private: Only this class
        self._spell_count = 0
    
    spell add_public_spell(spell_name):  // Public method
        self._protected_spells.append(spell_name)
        self._spell_count += 1
        return f"📖 Added public spell: {spell_name}"
    
    spell _protected_method(self):  // Protected method
        return f"🔐 Protected knowledge accessible to {self.owner}"
    
    spell __private_method(self):  // Private method
        return f"🚫 Secret magic known only to the original grimoire"
    
    spell get_spell_info(self):
        // Public method that uses private functionality
        secret_count = len(self.__secret_spells)
        return {
            "owner": self.owner,
            "public_spells": len(self._protected_spells),
            "secret_spells": secret_count,  // Can access private data internally
            "total": self._spell_count + secret_count
        }
    
    spell reveal_secret():  // Public method accessing private
        return self.__private_method()

grim AdvancedSpellBook(SecretSpellBook):
    init(owner_name, specialty):
        super.init(owner_name)
        self.specialty = specialty
    
    spell access_protected_knowledge():
        // Can access protected members from parent
        count = len(self._protected_spells)
        return f"🔓 {self.specialty} mage {self.owner} knows {count} protected spells"
    
    spell try_access_private():
        // Cannot access private members from parent
        // return self.__private_method()  // ❌ This would cause an error
        return "❌ Cannot access private methods from parent grimoire"

// Test encapsulation
basic_book = SecretSpellBook("Merlin")
advanced_book = AdvancedSpellBook("Gandalf", "Light Magic")

print(basic_book.add_public_spell("Lightning Bolt"))
print(basic_book.get_spell_info())

print(advanced_book.access_protected_knowledge())
print(advanced_book.try_access_private())

// Direct access examples
print(f"Public access: {basic_book.owner}")              // ✅ Works
print(f"Protected access: {len(basic_book._protected_spells)}")  // ⚠️ Works but discouraged
// print(basic_book.__secret_spells)  // ❌ Error: Private access not allowed
```

---

## 🌟 Abstract Grimoires - Defining Magical Contracts

### Arcane Grimoires (Abstract Classes)
```python
arcane grim MagicalWeapon:
    init(name, damage):
        self.name = name
        self.damage = damage
        self.durability = 100
    
    @arcanespell  // Abstract method - must be implemented by children
    spell attack():
        ignore  // No implementation in abstract class
    
    @arcanespell
    spell special_ability():
        ignore
    
    spell repair():  // Concrete method - can be used by all weapons
        self.durability = min(self.durability + 20, 100)
        return f"🔧 {self.name} repaired! Durability: {self.durability}/100"
    
    spell get_weapon_info():  // Another concrete method
        return f"Weapon: {self.name}, Damage: {self.damage}, Durability: {self.durability}/100"

// Concrete implementation of the abstract weapon
grim MagicalSword(MagicalWeapon):
    init(name, damage, element):
        super.init(name, damage)
        self.element = element
    
    spell attack():  // Must implement abstract method
        self.durability -= 5
        return f"⚔️ {self.name} slashes with {self.element} energy! Damage: {self.damage}"
    
    spell special_ability():  // Must implement abstract method
        if self.durability >= 20:
            self.durability -= 20
            special_damage = self.damage * 2
            return f"✨ {self.name} unleashes {self.element} wave! Special damage: {special_damage}"
        else:
            return f"💔 {self.name} is too damaged for special attacks!"

grim MagicalBow(MagicalWeapon):
    init(name, damage, arrow_type):
        super.init(name, damage)
        self.arrow_type = arrow_type
        self.arrows = 30
    
    spell attack():  // Must implement abstract method
        if self.arrows > 0:
            self.arrows -= 1
            self.durability -= 2
            return f"🏹 {self.name} fires a {self.arrow_type} arrow! Damage: {self.damage}, Arrows left: {self.arrows}"
        else:
            return f"🎯 {self.name} is out of arrows!"
    
    spell special_ability():  // Must implement abstract method
        if self.arrows >= 3:
            self.arrows -= 3
            self.durability -= 10
            triple_damage = self.damage * 3
            return f"🎯🎯🎯 {self.name} fires triple {self.arrow_type} shot! Damage: {triple_damage}"
        else:
            return f"🏹 Not enough arrows for special attack (need 3, have {self.arrows})"

// Use the concrete weapons
excalibur = MagicalSword("Excalibur", 50, "holy")
elven_bow = MagicalBow("Elven Longbow", 35, "piercing")

print(excalibur.attack())
print(excalibur.special_ability())
print(excalibur.get_weapon_info())

print(elven_bow.attack())
print(elven_bow.special_ability())
print(elven_bow.repair())
```

---

## 🎪 Advanced Grimoire Patterns

### Builder Pattern - Crafting Complex Objects
```python
grim MagicalPotion:
    init():
        self.name = "Unknown Potion"
        self.ingredients = []
        self.brewing_time = 0
        self.potency = 1.0
        self.effects = []
        self.side_effects = []
    
    spell add_ingredient(ingredient, amount = 1):
        self.ingredients.append({"name": ingredient, "amount": amount})
        return self  // Return self for method chaining
    
    spell set_name(name):
        self.name = name
        return self
    
    spell set_brewing_time(minutes):
        self.brewing_time = minutes
        self.potency = min(1.0 + (minutes / 60), 3.0)  // Longer brewing = more potent
        return self
    
    spell add_effect(effect):
        self.effects.append(effect)
        return self
    
    spell add_side_effect(side_effect):
        self.side_effects.append(side_effect)
        return self
    
    spell brew():
        ingredient_list = [f"{ing['amount']}x {ing['name']}" for ing in self.ingredients]
        result = {
            "name": self.name,
            "ingredients": ingredient_list,
            "brewing_time": f"{self.brewing_time} minutes",
            "potency": f"{self.potency:.1f}x",
            "effects": self.effects,
            "side_effects": self.side_effects
        }
        return f"🧪 Successfully brewed: {self.name}! Recipe: {result}"

// Use builder pattern with method chaining
healing_potion = MagicalPotion()
    .set_name("Greater Healing Potion")
    .add_ingredient("Phoenix Feather", 2)
    .add_ingredient("Unicorn Hair", 1)
    .add_ingredient("Spring Water", 3)
    .set_brewing_time(120)
    .add_effect("Restore 100 HP")
    .add_effect("Remove poison")
    .add_side_effect("Temporary drowsiness")
    .brew()

print(healing_potion)
```

### Observer Pattern - Magical Event System
```python
grim MagicalEvent:
    init():
        self.observers = []  // List of grimoires watching for events
    
    spell add_observer(observer):
        if observer not in self.observers:
            self.observers.append(observer)
            return f"👁️ {observer.name} is now watching for magical events"
    
    spell remove_observer(observer):
        if observer in self.observers:
            self.observers.remove(observer)
            return f"👁️ {observer.name} stopped watching for events"
    
    spell notify_observers(event_type, event_data):
        notifications = []
        for observer in self.observers:
            response = observer.handle_event(event_type, event_data)
            notifications.append(response)
        return notifications

grim MagicalGuard:
    init(name, location):
        self.name = name
        self.location = location
        self.alert_level = "normal"
    
    spell handle_event(event_type, event_data):
        if event_type == "intruder_detected":
            self.alert_level = "high"
            return f"🚨 Guard {self.name} at {self.location}: INTRUDER ALERT! {event_data}"
        elif event_type == "magical_disturbance":
            return f"⚡ Guard {self.name}: Magical disturbance detected - {event_data}"
        elif event_type == "all_clear":
            self.alert_level = "normal"
            return f"✅ Guard {self.name}: All clear signal received"
        else:
            return f"📢 Guard {self.name}: Unknown event - {event_type}"

grim MagicalAlarm:
    init(name):
        self.name = name
        self.activation_count = 0
    
    spell handle_event(event_type, event_data):
        if event_type == "intruder_detected":
            self.activation_count += 1
            return f"🔔 ALARM {self.name}: INTRUDER! Count: {self.activation_count}"
        else:
            return f"🔕 Alarm {self.name}: Monitoring..."

// Set up the magical security system
security_system = MagicalEvent()
guard1 = MagicalGuard("Alaric", "North Tower")
guard2 = MagicalGuard("Lyra", "Main Gate")
alarm = MagicalAlarm("Central Security")

// Add observers
print(security_system.add_observer(guard1))
print(security_system.add_observer(guard2))
print(security_system.add_observer(alarm))

// Trigger events
print("\n--- Intruder Detected ---")
responses = security_system.notify_observers("intruder_detected", "Unauthorized mage in courtyard")
for response in responses:
    print(response)

print("\n--- All Clear ---")
responses = security_system.notify_observers("all_clear", "Threat neutralized")
for response in responses:
    print(response)
```

---

## 🏆 Practice Challenges

### 🥉 Beginner Grimoire Challenges
1. **Magical Pet**: Create a `MagicalPet` grimoire with feeding, playing, and status methods
2. **Spell Component**: Design a `SpellComponent` grimoire with rarity, power, and combination methods
3. **Magic Shop**: Build a `MagicShop` grimoire to buy/sell magical items

### 🥈 Intermediate Grimoire Challenges
1. **RPG Character System**: Create a full character system with different classes
2. **Magical Inventory**: Design an inventory system with sorting and filtering
3. **Spell Crafting System**: Build a system to combine spell components into new spells

### 🥇 Advanced Grimoire Challenges
1. **Magical Academy**: Create a complete academy system with students, teachers, and courses
2. **Battle System**: Design a turn-based magical combat system
3. **Magical Economy**: Build a complex trading system for magical goods

---

## 🔗 What's Next?

Master the deeper mysteries of grimoire crafting:

1. **🌟 [Inheritance & Polymorphism](Inheritance.md)** - Advanced magical hierarchies
2. **⚔️ [Error Handling](Error-Handling.md)** - Protect your grimoires from magical mishaps
3. **📦 [Modules & Imports](Modules.md)** - Organize your magical libraries
4. **📚 [Standard Library](../Standard-Library/Munin-Overview.md)** - Use pre-built magical grimoires

---

## 💡 Grimoire Crafting Wisdom

1. **🎯 Single Responsibility**: Each grimoire should have one clear purpose
2. **🏗️ Composition over Inheritance**: Prefer "has-a" relationships over "is-a"
3. **🔒 Encapsulation**: Keep internal magic private, expose only what's needed
4. **📝 Clear Interfaces**: Make it obvious how to use your grimoires
5. **🧪 Test Your Magic**: Ensure your grimoires work in all situations
6. **📚 Document Thoroughly**: Explain how to use your magical creations

---

*Remember: Great grimoires are like ancient spellbooks - they contain wisdom that can be passed down through generations of mages. Craft yours with care and pride! 🪄✨*

> "A well-crafted grimoire is not just code, but a legacy of magical knowledge that empowers all who wield it." - *Master Grimoire Architect*