import React from 'react';
import styled from 'styled-components';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';

const Container = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  padding: 6rem 2rem 4rem;
  min-height: 100vh;
`;

const Header = styled.div`
  text-align: center;
  margin-bottom: 4rem;
  animation: fadeIn 0.8s ease;

  @keyframes fadeIn {
    from { opacity: 0; transform: translateY(-20px); }
    to { opacity: 1; transform: translateY(0); }
  }
`;

const Title = styled.h1`
  font-size: 3.5rem;
  margin-bottom: 1.5rem;
  background: ${({ theme }) => theme.gradients.primary};
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  font-weight: 800;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    font-size: 2.5rem;
  }
`;

const Subtitle = styled.p`
  font-size: 1.4rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  max-width: 700px;
  margin: 0 auto;
  line-height: 1.8;
`;

const Section = styled.section`
  margin-bottom: 4rem;
  animation: fadeInUp 0.6s ease;
  animation-fill-mode: both;

  &:nth-child(2) { animation-delay: 0.1s; }
  &:nth-child(3) { animation-delay: 0.2s; }
  &:nth-child(4) { animation-delay: 0.3s; }

  @keyframes fadeInUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
  }
`;

const SectionTitle = styled.h2`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 2rem;
  font-size: 2.5rem;
  font-weight: 700;
  position: relative;
  padding-left: 1rem;

  &::before {
    content: '';
    position: absolute;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    width: 4px;
    height: 70%;
    background: ${({ theme }) => theme.gradients.primary};
    border-radius: 2px;
  }
`;

const SubSectionTitle = styled.h3`
  color: ${({ theme }) => theme.colors.text.primary};
  margin: 2rem 0 1.5rem;
  font-size: 1.8rem;
  font-weight: 600;
`;

const Text = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  font-size: 1.1rem;
  line-height: 1.8;
  margin-bottom: 1.5rem;
`;

const CodeBlock = styled.div`
  margin: 2rem 0;
  border-radius: ${({ theme }) => theme.borderRadius.large};
  overflow: hidden;
  box-shadow: ${({ theme }) => theme.shadows.large};
  transition: transform ${({ theme }) => theme.transitions.standard};

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.3);
  }
`;

const InfoBox = styled.div`
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.1), rgba(139, 92, 246, 0.1));
  padding: 1.5rem;
  border-radius: ${({ theme }) => theme.borderRadius.medium};
  border-left: 4px solid ${({ theme }) => theme.colors.primary};
  margin: 2rem 0;
`;

const InfoTitle = styled.h4`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 0.5rem;
  font-weight: 600;
`;

const InfoText = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.6;
`;

const InlineCode = styled.code`
  background: rgba(6, 182, 212, 0.1);
  color: ${({ theme }) => theme.colors.primary};
  padding: 0.2rem 0.5rem;
  border-radius: 4px;
  font-family: 'Monaco', 'Courier New', monospace;
  font-size: 0.95em;
`;

const Grimoires: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Grimoires - OOP in Carrion</Title>
        <Subtitle>
          Master object-oriented programming with grimoires (classes) and spells (methods). Create elegant, reusable code with magical syntax.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>What are Grimoires?</SectionTitle>
        <Text>
          In Carrion, classes are called "grimoires" - spellbooks that contain "spells" (methods). This magical terminology 
          makes OOP engaging while providing full object-oriented capabilities including inheritance, encapsulation, 
          polymorphism, and abstraction.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Define a grimoire (class)
grim Person:
    init(name, age):
        self.name = name
        self.age = age
    
    spell greet():
        return f"Hello, I'm {self.name}"
    
    spell birthday():
        self.age += 1
        return f"Happy birthday! Now {self.age} years old"

// Create and use instances
person = Person("Alice", 30)
print(person.greet())      // "Hello, I'm Alice"
print(person.birthday())   // "Happy birthday! Now 31 years old"`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Inheritance</SectionTitle>
        <Text>
          Grimoires can inherit from other grimoires, gaining their properties and methods. Use <InlineCode>super</InlineCode> 
          to call parent methods.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Base grimoire
grim Animal:
    init(name, species):
        self.name = name
        self.species = species
    
    spell speak():
        return f"{self.name} makes a sound"
    
    spell info():
        return f"{self.name} is a {self.species}"

// Child grimoire with inheritance
grim Dog(Animal):
    init(name, breed):
        super.init(name, "Dog")
        self.breed = breed
    
    spell speak():  // Override parent method
        return f"{self.name} barks"
    
    spell fetch():  // New method
        return f"{self.name} fetches the ball"

// Usage
dog = Dog("Rex", "Golden Retriever")
print(dog.speak())   // "Rex barks" (overridden)
print(dog.info())    // "Rex is a Dog" (inherited)
print(dog.fetch())   // "Rex fetches the ball" (new)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Abstract Grimoires</SectionTitle>
        <Text>
          Abstract grimoires define interfaces that child grimoires must implement. Use the <InlineCode>arcane</InlineCode> keyword 
          and <InlineCode>@arcanespell</InlineCode> decorator for abstract methods.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Abstract base grimoire
arcane grim Shape:
    init(name):
        self.name = name
    
    @arcanespell
    spell area():
        ignore  // Must be implemented by children
    
    @arcanespell
    spell perimeter():
        ignore  // Must be implemented by children
    
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

print(f"Rectangle area: {rect.area()}")           // 15
print(f"Circle perimeter: {circle.perimeter()}")  // 25.13`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Encapsulation</SectionTitle>
        <Text>
          Control access to data using naming conventions: public (no underscore), protected (single underscore), 
          and private (double underscore).
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`grim BankAccount:
    init(account_number, initial_balance):
        self.account_number = account_number    // Public
        self._balance = initial_balance         // Protected
        self.__pin = "1234"                    // Private
    
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
print(account.account_number)    // OK - public
print(account._balance)          // Works but discouraged - protected
// print(account.__pin)          // Error - private`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Practical Examples</SectionTitle>
        
        <SubSectionTitle>Builder Pattern</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
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
        return f"{self.size} pizza with {self.crust} crust and {toppings_str}"

// Method chaining
pizza = Pizza()
    .with_size("large")
    .add_topping("pepperoni")
    .add_topping("mushrooms")
    .with_crust("thin")
    .build()

print(pizza)  // "large pizza with thin crust and pepperoni, mushrooms"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Observer Pattern</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`grim Observable:
    init():
        self.observers = []
    
    spell add_observer(observer):
        self.observers.append(observer)
    
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
        print(f"{self.name} received: {news}")

// Usage
publisher = NewsPublisher()
sub1 = NewsSubscriber("Alice")
sub2 = NewsSubscriber("Bob")

publisher.add_observer(sub1)
publisher.add_observer(sub2)

publisher.publish_news("Carrion released!")
// Alice received: Carrion released!
// Bob received: Carrion released!`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Composition Over Inheritance</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`grim Engine:
    init(horsepower):
        self.horsepower = horsepower
    
    spell start():
        return "Engine started"
    
    spell stop():
        return "Engine stopped"

grim Car:  // Has-a relationship
    init(make, model, engine):
        self.make = make
        self.model = model
        self.engine = engine  // Composition
    
    spell start():
        return f"{self.make} {self.model}: {self.engine.start()}"
    
    spell info():
        return f"{self.make} {self.model} with {self.engine.horsepower}HP"

// Usage
engine = Engine(300)
car = Car("Tesla", "Model S", engine)
print(car.start())  // "Tesla Model S: Engine started"
print(car.info())   // "Tesla Model S with 300HP"`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Advanced Features</SectionTitle>
        
        <SubSectionTitle>Custom String Representation</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`grim Point:
    init(x, y):
        self.x = x
        self.y = y
    
    spell to_string():
        return f"Point({self.x}, {self.y})"
    
    spell distance_to(other):
        dx = self.x - other.x
        dy = self.y - other.y
        return (dx ** 2 + dy ** 2) ** 0.5
    
    spell move(dx, dy):
        return Point(self.x + dx, self.y + dy)

p1 = Point(3, 4)
p2 = Point(6, 8)
print(p1.to_string())             // "Point(3, 4)"
print(f"Distance: {p1.distance_to(p2)}")  // "Distance: 5.0"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Multiple Inheritance</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`grim Flyable:
    spell fly():
        return "Flying high"

grim Swimmer:
    spell swim():
        return "Swimming gracefully"

grim Bird(Animal, Flyable):
    init(name, wingspan):
        super.init(name, "Bird")
        self.wingspan = wingspan
    
    spell speak():
        return f"{self.name} chirps"

grim Duck(Bird, Swimmer):
    init(name):
        super.init(name, 24)
    
    spell speak():
        return f"{self.name} quacks"

duck = Duck("Donald")
print(duck.speak())  // "Donald quacks"
print(duck.fly())    // "Flying high"
print(duck.swim())   // "Swimming gracefully"`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Best Practices</SectionTitle>
        
        <InfoBox>
          <InfoTitle>Use Descriptive Names</InfoTitle>
          <InfoText>
            Choose clear, meaningful names for grimoires and spells. The grimoire name should describe what it represents, 
            and spell names should describe what they do.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Keep Grimoires Focused</InfoTitle>
          <InfoText>
            Each grimoire should have a single, well-defined responsibility. If a grimoire is doing too much, 
            consider splitting it into multiple grimoires.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Favor Composition</InfoTitle>
          <InfoText>
            When possible, use composition (has-a relationships) instead of inheritance (is-a relationships). 
            This leads to more flexible and maintainable code.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Document Abstract Methods</InfoTitle>
          <InfoText>
            When creating abstract grimoires, clearly document what child grimoires must implement and what the 
            expected behavior should be.
          </InfoText>
        </InfoBox>
      </Section>
    </Container>
  );
};

export default Grimoires;
