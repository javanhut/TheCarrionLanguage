# Object-Oriented Programming with Grimoires

Carrion embraces object-oriented programming through "grimoires" (spellbooks) - classes that contain "spells" (methods). This magical terminology creates an engaging programming experience while providing full OOP capabilities.

## Basic Grimoire Definition

### Simple Grimoire
```python
grim Person:
    init(name, age):
        self.name = name
        self.age = age
    
    spell greet():
        return f"Hello, I'm {self.name}"
    
    spell get_age():
        return self.age
```

### Creating and Using Instances
```python
# Create an instance
person = Person("Alice", 30)

# Call methods (spells)
greeting = person.greet()
age = person.get_age()

print(greeting)  # → "Hello, I'm Alice"
print(f"Age: {age}")  # → "Age: 30"

# Access properties directly
print(person.name)  # → "Alice"
person.age = 31     # Modify property
```

## Constructor (init spell)

The `init` spell is the constructor method that initializes new instances.

```python
grim BankAccount:
    init(account_number, initial_balance = 0):
        self.account_number = account_number
        self.balance = initial_balance
        self.transaction_history = []
    
    spell deposit(amount):
        if amount > 0:
            self.balance += amount
            self.transaction_history.append(f"Deposit: +{amount}")
            return True
        return False
    
    spell withdraw(amount):
        if amount > 0 and amount <= self.balance:
            self.balance -= amount
            self.transaction_history.append(f"Withdrawal: -{amount}")
            return True
        return False
    
    spell get_balance():
        return self.balance

# Usage
account = BankAccount("12345", 1000)
account.deposit(500)
account.withdraw(200)
print(f"Balance: ${account.get_balance()}")  # → "Balance: $1300"
```

## Inheritance

Grimoires can inherit from other grimoires, gaining their properties and methods.

### Basic Inheritance
```python
grim Animal:
    init(name, species):
        self.name = name
        self.species = species
    
    spell speak():
        return f"{self.name} makes a sound"
    
    spell info():
        return f"{self.name} is a {self.species}"

# Child grimoire inheriting from Animal
grim Dog(Animal):
    init(name, breed):
        super.init(name, "Dog")  # Call parent constructor
        self.breed = breed
    
    spell speak():  # Override parent method
        return f"{self.name} barks"
    
    spell fetch():  # New method specific to Dog
        return f"{self.name} fetches the ball"

# Usage
dog = Dog("Rex", "Golden Retriever")
print(dog.speak())   # → "Rex barks"
print(dog.info())    # → "Rex is a Dog"
print(dog.fetch())   # → "Rex fetches the ball"
```

### Multiple Inheritance Concepts
```python
grim Flyable:
    spell fly():
        return "Flying high"

grim Swimmer:
    spell swim():
        return "Swimming gracefully"

# Bird inherits flying capability
grim Bird(Animal, Flyable):
    init(name, wingspan):
        super.init(name, "Bird")
        self.wingspan = wingspan
    
    spell speak():
        return f"{self.name} chirps"

# Duck can both fly and swim
grim Duck(Bird, Swimmer):
    init(name):
        super.init(name, 24)  # 24 inch wingspan
    
    spell speak():
        return f"{self.name} quacks"

# Usage
duck = Duck("Donald")
print(duck.speak())  # → "Donald quacks"
print(duck.fly())    # → "Flying high"
print(duck.swim())   # → "Swimming gracefully"
```

### Multi-Level Inheritance

Carrion supports deep inheritance hierarchies where classes can inherit from classes that themselves inherit from other classes. The `super` keyword correctly resolves to the immediate parent class at each level.

```python
# Three-level inheritance example
grim Vehicle:
    init(wheels):
        self.wheels = wheels
        print(f"Vehicle created with {wheels} wheels")
    
    spell start():
        return "Vehicle starting"

grim Car(Vehicle):
    init(brand, model):
        super.init(4)  # Cars have 4 wheels
        self.brand = brand
        self.model = model
        print(f"Car created: {brand} {model}")
    
    spell start():
        return super.start() + " - Engine running"

grim ElectricCar(Car):
    init(brand, model, battery_capacity):
        super.init(brand, model)  # Calls Car.init, which calls Vehicle.init
        self.battery_capacity = battery_capacity
        print(f"Electric car with {battery_capacity}kWh battery")
    
    spell start():
        return super.start() + " - Battery powered"
    
    spell charge():
        return f"Charging {self.battery_capacity}kWh battery"

# Usage
tesla = ElectricCar("Tesla", "Model 3", 75)
# Output:
# Vehicle created with 4 wheels
# Car created: Tesla Model 3
# Electric car with 75kWh battery

print(tesla.start())  
# → "Vehicle starting - Engine running - Battery powered"

print(tesla.charge())
# → "Charging 75kWh battery"
```

**Important Notes:**
- Each `super.init()` call invokes the immediate parent's constructor
- The inheritance chain is properly maintained through multiple levels
- All parent constructors are called in the correct order
- Instance variables from all levels of the hierarchy are accessible

### Deep Inheritance Hierarchies

Carrion supports inheritance hierarchies of any depth:

```python
grim Level1:
    init(x):
        self.l1 = x
        print(f"Level1: {x}")

grim Level2(Level1):
    init(x):
        super.init(x)
        self.l2 = x * 2
        print(f"Level2: {self.l2}")

grim Level3(Level2):
    init(x):
        super.init(x)
        self.l3 = x * 3
        print(f"Level3: {self.l3}")

grim Level4(Level3):
    init(x):
        super.init(x)
        self.l4 = x * 4
        print(f"Level4: {self.l4}")

grim Level5(Level4):
    init(x):
        super.init(x)
        self.l5 = x * 5
        print(f"Level5: {self.l5}")

# Create a Level5 instance
obj = Level5(10)
# Output:
# Level1: 10
# Level2: 20
# Level3: 30
# Level4: 40
# Level5: 50

# Access all inherited properties
print(f"Values: {obj.l1}, {obj.l2}, {obj.l3}, {obj.l4}, {obj.l5}")
# → "Values: 10, 20, 30, 40, 50"
```

## Abstract Grimoires

Abstract grimoires define interfaces that child grimoires must implement.

### Abstract Grimoire Definition
```python
arcane grim Shape:
    init(name):
        self.name = name
    
    @arcanespell
    spell area():
        ignore  # Abstract method - no implementation
    
    @arcanespell
    spell perimeter():
        ignore  # Abstract method - no implementation
    
    spell description():
        return f"This is a {self.name}"
```

### Implementing Abstract Grimoires
```python
grim Rectangle(Shape):
    init(width, height):
        super.init("Rectangle")
        self.width = width
        self.height = height
    
    spell area():  # Must implement abstract method
        return self.width * self.height
    
    spell perimeter():  # Must implement abstract method
        return 2 * (self.width + self.height)

grim Circle(Shape):
    init(radius):
        super.init("Circle")
        self.radius = radius
    
    spell area():
        return 3.14159 * self.radius ** 2
    
    spell perimeter():
        return 2 * 3.14159 * self.radius

# Usage
rectangle = Rectangle(5, 3)
circle = Circle(4)

print(f"Rectangle area: {rectangle.area()}")      # → "Rectangle area: 15"
print(f"Circle perimeter: {circle.perimeter()}")  # → "Circle perimeter: 25.13272"
```

## Access Modifiers

Carrion supports access control through naming conventions.

### Public, Protected, and Private
```python
grim BankAccount:
    init(account_number, initial_balance):
        self.account_number = account_number    # Public
        self._balance = initial_balance         # Protected (single underscore)
        self.__pin = "1234"                    # Private (double underscore)
    
    spell get_balance():  # Public method
        return self._balance
    
    spell _validate_transaction(amount):  # Protected method
        return amount > 0 and amount <= self._balance
    
    spell __check_pin(entered_pin):  # Private method
        return entered_pin == self.__pin
    
    spell withdraw(amount, pin):
        if self.__check_pin(pin) and self._validate_transaction(amount):
            self._balance -= amount
            return True
        return False

# Usage
account = BankAccount("12345", 1000)
print(account.account_number)    # Public access - allowed
print(account._balance)          # Protected - accessible but discouraged
# print(account.__pin)          # Private - not accessible
```

## Method Types

### Instance Methods
Regular methods that operate on instance data.

```python
grim Counter:
    init(start = 0):
        self.count = start
    
    spell increment():       # Instance method
        self.count += 1
    
    spell decrement():       # Instance method
        self.count -= 1
    
    spell get_value():       # Instance method
        return self.count
```

### Class-like Methods (Static Functionality)
```python
grim MathUtils:
    spell add(a, b):         # Utility method
        return a + b
    
    spell multiply(a, b):    # Utility method
        return a * b
    
    spell factorial(n):      # Utility method
        if n <= 1:
            return 1
        return n * MathUtils.factorial(n - 1)

# Usage (can be called without instance)
result = MathUtils.add(5, 3)
fact = MathUtils.factorial(5)
```

## Properties and Encapsulation

### Getter and Setter Patterns
```python
grim Temperature:
    init(celsius = 0):
        self._celsius = celsius
    
    spell get_celsius():
        return self._celsius
    
    spell set_celsius(value):
        if value < -273.15:
            raise Error("Temperature", "Cannot be below absolute zero")
        self._celsius = value
    
    spell get_fahrenheit():
        return (self._celsius * 9/5) + 32
    
    spell set_fahrenheit(value):
        celsius = (value - 32) * 5/9
        self.set_celsius(celsius)
    
    spell get_kelvin():
        return self._celsius + 273.15

# Usage
temp = Temperature(25)
print(f"Celsius: {temp.get_celsius()}")      # → "Celsius: 25"
print(f"Fahrenheit: {temp.get_fahrenheit()}") # → "Fahrenheit: 77"
print(f"Kelvin: {temp.get_kelvin()}")        # → "Kelvin: 298.15"

temp.set_fahrenheit(100)
print(f"Celsius: {temp.get_celsius()}")      # → "Celsius: 37.777..."
```

## Advanced OOP Patterns

### Builder Pattern
```python
grim Pizza:
    init():
        self.size = "medium"
        self.toppings = []
        self.crust = "regular"
    
    spell with_size(size):
        self.size = size
        return self  # Return self for chaining
    
    spell add_topping(topping):
        self.toppings.append(topping)
        return self  # Return self for chaining
    
    spell with_crust(crust):
        self.crust = crust
        return self  # Return self for chaining
    
    spell build():
        toppings_str = ", ".join(self.toppings)
        return f"{self.size} pizza with {self.crust} crust and {toppings_str}"

# Usage with method chaining
pizza = Pizza()
    .with_size("large")
    .add_topping("pepperoni")
    .add_topping("mushrooms")
    .with_crust("thin")
    .build()

print(pizza)  # → "large pizza with thin crust and pepperoni, mushrooms"
```

### Observer Pattern
```python
grim Observable:
    init():
        self.observers = []
    
    spell add_observer(observer):
        self.observers.append(observer)
    
    spell remove_observer(observer):
        if observer in self.observers:
            self.observers.remove(observer)
    
    spell notify_observers(data):
        for observer in self.observers:
            observer.update(data)

grim NewsPublisher(Observable):
    init():
        super.init()
        self.news = ""
    
    spell publish_news(news):
        self.news = news
        self.notify_observers(news)

grim NewsSubscriber:
    init(name):
        self.name = name
    
    spell update(news):
        print(f"{self.name} received news: {news}")

# Usage
publisher = NewsPublisher()
subscriber1 = NewsSubscriber("Alice")
subscriber2 = NewsSubscriber("Bob")

publisher.add_observer(subscriber1)
publisher.add_observer(subscriber2)

publisher.publish_news("Breaking: Carrion language released!")
# Output:
# Alice received news: Breaking: Carrion language released!
# Bob received news: Breaking: Carrion language released!
```

### Singleton-like Pattern
```python
grim Config:
    init():
        if hasattr(Config, "_instance"):
            raise Error("Singleton", "Config already exists")
        
        Config._instance = self
        self.settings = {}
    
    spell get_instance():
        if not hasattr(Config, "_instance"):
            Config._instance = Config()
        return Config._instance
    
    spell set_setting(key, value):
        self.settings[key] = value
    
    spell get_setting(key, default = None):
        return self.settings.get(key, default)

# Usage
config1 = Config.get_instance()
config1.set_setting("debug", True)

config2 = Config.get_instance()
print(config2.get_setting("debug"))  # → True (same instance)
```

## Operator Overloading Concepts

### Custom String Representation
```python
grim Point:
    init(x, y):
        self.x = x
        self.y = y
    
    spell to_string():  # Custom string representation
        return f"Point({self.x}, {self.y})"
    
    spell distance_to(other):
        dx = self.x - other.x
        dy = self.y - other.y
        return (dx ** 2 + dy ** 2) ** 0.5
    
    spell move(dx, dy):
        return Point(self.x + dx, self.y + dy)

# Usage
p1 = Point(3, 4)
p2 = Point(6, 8)

print(p1.to_string())                # → "Point(3, 4)"
print(f"Distance: {p1.distance_to(p2)}")  # → "Distance: 5.0"

p3 = p1.move(2, 3)
print(p3.to_string())                # → "Point(5, 7)"
```

## Best Practices

### Composition over Inheritance
```python
grim Engine:
    init(horsepower):
        self.horsepower = horsepower
    
    spell start():
        return "Engine started"
    
    spell stop():
        return "Engine stopped"

grim Car:  # Composition instead of inheritance
    init(make, model, engine):
        self.make = make
        self.model = model
        self.engine = engine  # Has-a relationship
    
    spell start():
        return f"{self.make} {self.model}: {self.engine.start()}"
    
    spell info():
        return f"{self.make} {self.model} with {self.engine.horsepower}HP engine"

# Usage
engine = Engine(300)
car = Car("Tesla", "Model S", engine)
print(car.start())  # → "Tesla Model S: Engine started"
```

### Interface Segregation
```python
# Better: Small, focused interfaces
arcane grim Readable:
    @arcanespell
    spell read():
        ignore

arcane grim Writable:
    @arcanespell
    spell write(data):
        ignore

# Classes implement only what they need
grim FileReader(Readable):
    init(filename):
        self.filename = filename
    
    spell read():
        return f"Reading from {self.filename}"

grim FileWriter(Writable):
    init(filename):
        self.filename = filename
    
    spell write(data):
        return f"Writing '{data}' to {self.filename}"

grim FileManager(Readable, Writable):  # Implements both
    init(filename):
        self.filename = filename
    
    spell read():
        return f"Reading from {self.filename}"
    
    spell write(data):
        return f"Writing '{data}' to {self.filename}"
```

### Inheritance Best Practices

When using inheritance in Carrion, follow these guidelines for maintainable and efficient code:

#### 1. Always Call Super in Constructors
When overriding `init` in a child class, always call `super.init()` to ensure proper initialization:

```python
grim Parent:
    init(name):
        self.name = name
        self.created_at = current_time()

grim Child(Parent):
    init(name, age):
        super.init(name)  # Essential - initializes parent properties
        self.age = age
```

#### 2. Limit Inheritance Depth
While Carrion supports deep inheritance hierarchies, keep them shallow for maintainability:

```python
# Good: 2-3 levels
grim Animal -> Mammal -> Dog

# Avoid: Too deep
grim Entity -> LivingThing -> Animal -> Vertebrate -> Mammal -> Canine -> Dog
```

#### 3. Use Abstract Classes for Contracts
Define clear contracts with abstract grimoires:

```python
arcane grim Processor:
    @arcanespell
    spell process(data):
        ignore
    
    spell validate(data):  # Concrete method with default implementation
        return data is not None

grim JsonProcessor(Processor):
    spell process(data):
        # Must implement abstract method
        return parse_json(data)
```

#### 4. Prefer Composition for Complex Behaviors
Use inheritance for "is-a" relationships and composition for "has-a" relationships:

```python
# Inheritance: Dog IS AN Animal
grim Dog(Animal):
    init(name):
        super.init(name, "Canine")

# Composition: Car HAS AN Engine
grim Car:
    init(engine):
        self.engine = engine  # Composition
```

#### 5. Override Methods Consistently
When overriding methods, maintain the same signature and return type:

```python
grim Shape:
    spell area():
        return 0.0

grim Circle(Shape):
    spell area():  # Same signature as parent
        return 3.14159 * self.radius * self.radius
```

#### 6. Document Inheritance Hierarchies
Always document the purpose and relationships in complex hierarchies:

```python
# Base class for all game entities
grim Entity:
    init(x, y):
        self.x = x
        self.y = y

# Represents any object that can move
grim Movable(Entity):
    spell move(dx, dy):
        self.x += dx
        self.y += dy

# Player character with movement and health
grim Player(Movable):
    init(x, y, health):
        super.init(x, y)
        self.health = health
```

Grimoires in Carrion provide powerful object-oriented programming capabilities with an engaging magical theme, supporting inheritance, abstraction, encapsulation, and polymorphism while maintaining clear, readable syntax.