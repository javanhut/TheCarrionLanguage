import React from 'react';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import {
  DocLayout,
  Section,
  SectionTitle,
  SubSection,
  SubSectionTitle,
  Paragraph,
  CodeBlock,
  InfoBox,
  InfoTitle,
  InfoText,
  InlineCode,
} from '../../components/docs';

const sections = [
  { id: 'basics', title: 'Basics' },
  { id: 'inheritance', title: 'Inheritance' },
  { id: 'abstract', title: 'Abstract Classes' },
  { id: 'encapsulation', title: 'Encapsulation' },
  { id: 'patterns', title: 'Design Patterns' },
];

const Grimoires: React.FC = () => {
  return (
    <DocLayout
      title="Grimoires (Classes)"
      description="Object-oriented programming with grimoires (classes) and spells (methods)."
      sections={sections}
    >
      <Section id="basics">
        <SectionTitle>Grimoire Basics</SectionTitle>
        <Paragraph>
          In Carrion, classes are called "grimoires" (spellbooks) and methods are called "spells".
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim Person:
    init(name, age):
        self.name = name
        self.age = age

    spell greet():
        return f"Hello, I'm {self.name}"

    spell birthday():
        self.age += 1
        return f"Now {self.age} years old"

// Create and use instances
person = Person("Alice", 30)
print(person.greet())      // "Hello, I'm Alice"
print(person.birthday())   // "Now 31 years old"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Terminology</InfoTitle>
          <InfoText>
            <InlineCode>grim</InlineCode> = class, <InlineCode>spell</InlineCode> = method,
            <InlineCode>init</InlineCode> = constructor, <InlineCode>self</InlineCode> = instance reference
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="inheritance">
        <SectionTitle>Inheritance</SectionTitle>
        <Paragraph>
          Grimoires can inherit from other grimoires using parentheses. Use <InlineCode>super</InlineCode> to
          call parent methods.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Base grimoire
grim Animal:
    init(name, species):
        self.name = name
        self.species = species

    spell speak():
        return f"{self.name} makes a sound"

    spell info():
        return f"{self.name} is a {self.species}"

// Child grimoire
grim Dog(Animal):
    init(name, breed):
        super.init(name, "Dog")
        self.breed = breed

    spell speak():  // Override
        return f"{self.name} barks"

    spell fetch():  // New method
        return f"{self.name} fetches the ball"

// Usage
dog = Dog("Rex", "Golden Retriever")
print(dog.speak())   // "Rex barks" (overridden)
print(dog.info())    // "Rex is a Dog" (inherited)
print(dog.fetch())   // "Rex fetches the ball"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSection>
          <SubSectionTitle>Multiple Inheritance</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim Flyable:
    spell fly():
        return "Flying high"

grim Swimmer:
    spell swim():
        return "Swimming gracefully"

grim Duck(Animal, Flyable, Swimmer):
    init(name):
        super.init(name, "Duck")

    spell speak():
        return f"{self.name} quacks"

duck = Duck("Donald")
print(duck.fly())    // "Flying high"
print(duck.swim())   // "Swimming gracefully"
print(duck.speak())  // "Donald quacks"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="abstract">
        <SectionTitle>Abstract Grimoires</SectionTitle>
        <Paragraph>
          Use <InlineCode>arcane</InlineCode> to define abstract classes and <InlineCode>@arcanespell</InlineCode> for
          abstract methods that must be implemented by child classes.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`arcane grim Shape:
    init(name):
        self.name = name

    @arcanespell
    spell area():
        ignore  // Must be implemented

    @arcanespell
    spell perimeter():
        ignore

    spell description():
        return f"This is a {self.name}"

// Concrete implementations
grim Rectangle(Shape):
    init(width, height):
        super.init("Rectangle")
        self.width = width
        self.height = height

    spell area():
        return self.width * self.height

    spell perimeter():
        return 2 * (self.width + self.height)

grim Circle(Shape):
    init(radius):
        super.init("Circle")
        self.radius = radius

    spell area():
        return 3.14159 * self.radius ** 2

    spell perimeter():
        return 2 * 3.14159 * self.radius

// Usage
rect = Rectangle(5, 3)
circle = Circle(4)
print(f"Rectangle area: {rect.area()}")  // 15`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="encapsulation">
        <SectionTitle>Encapsulation</SectionTitle>
        <Paragraph>
          Use naming conventions to indicate access levels:
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim BankAccount:
    init(account_number, initial_balance):
        self.account_number = account_number  // Public
        self._balance = initial_balance       // Protected
        self.__pin = "1234"                   // Private

    spell get_balance():  // Public getter
        return self._balance

    spell _validate_transaction(amount):  // Protected
        return amount > 0 and amount <= self._balance

    spell __check_pin(pin):  // Private
        return pin == self.__pin

    spell withdraw(amount, pin):
        if self.__check_pin(pin) and self._validate_transaction(amount):
            self._balance -= amount
            return True
        return False

// Usage
account = BankAccount("12345", 1000)
print(account.account_number)  // OK - public
print(account._balance)        // Works but discouraged
// account.__pin                // Error - private`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="patterns">
        <SectionTitle>Design Patterns</SectionTitle>

        <SubSection>
          <SubSectionTitle>Builder Pattern</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim Pizza:
    init():
        self.size = "medium"
        self.toppings = []
        self.crust = "regular"

    spell with_size(size):
        self.size = size
        return self  // Enable chaining

    spell add_topping(topping):
        self.toppings.append(topping)
        return self

    spell with_crust(crust):
        self.crust = crust
        return self

    spell build():
        toppings_str = ", ".join(self.toppings)
        return f"{self.size} {self.crust} pizza with {toppings_str}"

// Method chaining
pizza = Pizza()
    .with_size("large")
    .add_topping("pepperoni")
    .add_topping("mushrooms")
    .with_crust("thin")
    .build()

print(pizza)`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Composition</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim Engine:
    init(horsepower):
        self.horsepower = horsepower

    spell start():
        return "Engine started"

grim Car:  // Has-a relationship
    init(make, model, engine):
        self.make = make
        self.model = model
        self.engine = engine  // Composition

    spell start():
        return f"{self.make} {self.model}: {self.engine.start()}"

// Usage
engine = Engine(300)
car = Car("Tesla", "Model S", engine)
print(car.start())  // "Tesla Model S: Engine started"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <InfoBox>
          <InfoTitle>Favor Composition</InfoTitle>
          <InfoText>
            When possible, use composition (has-a) instead of inheritance (is-a).
            This leads to more flexible and maintainable code.
          </InfoText>
        </InfoBox>
      </Section>
    </DocLayout>
  );
};

export default Grimoires;
