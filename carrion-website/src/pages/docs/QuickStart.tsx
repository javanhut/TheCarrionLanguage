import React from 'react';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import { Link } from 'react-router-dom';
import {
  DocLayout,
  Section,
  SectionTitle,
  SubSection,
  SubSectionTitle,
  Paragraph,
  Lead,
  CodeBlock,
  InfoBox,
  InfoTitle,
  InfoText,
  TipBox,
  CardGrid,
  Card,
  CardTitle,
  CardDescription,
  InlineCode,
  List,
  ListItem,
} from '../../components/docs';

const sections = [
  { id: 'installation', title: 'Installation' },
  { id: 'first-program', title: 'First Program' },
  { id: 'basics', title: 'Language Basics' },
  { id: 'functions', title: 'Functions' },
  { id: 'control-flow', title: 'Control Flow' },
  { id: 'oop', title: 'Object-Oriented' },
  { id: 'error-handling', title: 'Error Handling' },
  { id: 'next-steps', title: 'Next Steps' },
];

const QuickStart: React.FC = () => {
  return (
    <DocLayout
      title="Quick Start Guide"
      description="Get up and running with Carrion in minutes. Learn the fundamentals and write your first spells."
      sections={sections}
    >
      <Section id="installation">
        <SectionTitle>Installation</SectionTitle>
        <Paragraph>
          Install Carrion on your system using one of the methods below.
        </Paragraph>

        <SubSection>
          <SubSectionTitle>Linux/macOS Quick Install</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`curl -sSL https://raw.githubusercontent.com/javanhut/TheCarrionLanguage/main/install/install.sh | bash`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Build from Source</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage
make install`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <InfoBox>
          <InfoTitle>Verify Installation</InfoTitle>
          <InfoText>
            After installation, verify that Carrion is working: <InlineCode>carrion --version</InlineCode>
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="first-program">
        <SectionTitle>Your First Program</SectionTitle>
        <Paragraph>
          Create a file called <InlineCode>hello.crl</InlineCode>:
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// hello.crl - Your first Carrion program
print("Hello, World!")
print("Welcome to Carrion!")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <Paragraph>Run your program:</Paragraph>
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`carrion hello.crl`}
          </SyntaxHighlighter>
        </CodeBlock>

        <TipBox>
          <InfoTitle>Interactive REPL</InfoTitle>
          <InfoText>
            Start an interactive session by running <InlineCode>carrion</InlineCode> without arguments.
          </InfoText>
        </TipBox>
      </Section>

      <Section id="basics">
        <SectionTitle>Language Basics</SectionTitle>

        <SubSection>
          <SubSectionTitle>Variables and Data Types</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Numbers
age = 25
price = 19.99

// Strings
name = "Alice"
greeting = "Hello, World!"

// Booleans
is_active = True
is_empty = False

// Arrays
numbers = [1, 2, 3, 4, 5]
names = ["Alice", "Bob", "Charlie"]

// Hashes (dictionaries)
person = {"name": "Alice", "age": 30}

// None
result = None`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Basic Operations</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Arithmetic
sum = 10 + 5        // 15
product = 10 * 5    // 50
power = 2 ** 3      // 8

// String concatenation
full_name = "John" + " " + "Doe"

// Array operations
numbers.append(6)
first = numbers[0]
length = len(numbers)`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="functions">
        <SectionTitle>Functions (Spells)</SectionTitle>
        <Paragraph>
          In Carrion, functions are called "spells" and are defined using the <InlineCode>spell</InlineCode> keyword.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Define a simple spell
spell greet(name):
    return f"Hello, {name}!"

// Call the spell
message = greet("Alice")
print(message)  // Hello, Alice!

// Spell with default parameters
spell introduce(name, age = 25):
    return f"My name is {name} and I'm {age}"

print(introduce("Bob"))        // My name is Bob and I'm 25
print(introduce("Alice", 30))  // My name is Alice and I'm 30

// Lambda expressions
double = (x) -> x * 2
print(double(5))  // 10`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="control-flow">
        <SectionTitle>Control Flow</SectionTitle>

        <SubSection>
          <SubSectionTitle>Conditionals</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// If statement
if age >= 18:
    print("Adult")

// If-otherwise-else (like elif)
if temperature < 32:
    status = "Freezing"
otherwise temperature < 70:
    status = "Cool"
otherwise temperature < 85:
    status = "Warm"
else:
    status = "Hot"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Loops</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// For loop
for num in [1, 2, 3]:
    print(num)

// Range loop
for i in range(5):
    print(f"Count: {i}")

// While loop
count = 0
while count < 5:
    print(count)
    count += 1

// Loop control
for i in range(10):
    if i == 3:
        skip  // continue
    if i == 7:
        stop  // break
    print(i)`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="oop">
        <SectionTitle>Object-Oriented Programming</SectionTitle>
        <Paragraph>
          Create classes (called "grimoires") to organize your code.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Define a grimoire (class)
grim Person:
    init(name, age):
        self.name = name
        self.age = age

    spell greet():
        return f"Hello, I'm {self.name}"

    spell birthday():
        self.age += 1
        return f"Happy birthday! Now {self.age}"

// Create an instance
person = Person("Alice", 30)
print(person.greet())      // Hello, I'm Alice
print(person.birthday())   // Happy birthday! Now 31`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="error-handling">
        <SectionTitle>Error Handling</SectionTitle>
        <Paragraph>
          Handle errors gracefully using the <InlineCode>attempt-ensnare</InlineCode> pattern.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`attempt:
    number = int(input("Enter a number: "))
    result = 100 / number
    print(f"Result: {result}")
ensnare:
    print("Error: Invalid input!")

// Specific error handling
attempt:
    file = File()
    content = file.read("data.txt")
ensnare (FileNotFoundError):
    print("File not found!")
ensnare:
    print("An unexpected error!")
resolve:
    print("Cleanup completed")  // Always runs`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="next-steps">
        <SectionTitle>Next Steps</SectionTitle>
        <Lead>
          Now that you understand the basics, explore these resources to deepen your knowledge.
        </Lead>

        <CardGrid>
          <Card as={Link} to="/docs/language-reference" style={{ textDecoration: 'none' }}>
            <CardTitle>Language Reference</CardTitle>
            <CardDescription>Comprehensive guide to Carrion's syntax and features.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/standard-library" style={{ textDecoration: 'none' }}>
            <CardTitle>Standard Library</CardTitle>
            <CardDescription>Explore the Munin standard library documentation.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/repl-guide" style={{ textDecoration: 'none' }}>
            <CardTitle>REPL Guide</CardTitle>
            <CardDescription>Master the interactive REPL for rapid prototyping.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/grimoires" style={{ textDecoration: 'none' }}>
            <CardTitle>Grimoires (OOP)</CardTitle>
            <CardDescription>Deep dive into object-oriented programming.</CardDescription>
          </Card>
        </CardGrid>
      </Section>
    </DocLayout>
  );
};

export default QuickStart;
