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
  TipTitle,
  CardGrid,
  Card,
  CardTitle,
  CardDescription,
  InlineCode,
  ComparisonTable,
  ComparisonItem,
  ComparisonLabel,
  ComparisonValue,
  List,
  ListItem,
} from '../../components/docs';

const sections = [
  { id: 'hello-world', title: 'Hello World' },
  { id: 'variables', title: 'Variables' },
  { id: 'data-types', title: 'Data Types' },
  { id: 'functions', title: 'Functions' },
  { id: 'classes', title: 'Grimoires (Classes)' },
  { id: 'error-handling', title: 'Error Handling' },
  { id: 'repl', title: 'Using the REPL' },
];

const GettingStarted: React.FC = () => {
  return (
    <DocLayout
      title="Introduction"
      description="Welcome to Carrion! Learn the fundamentals and start writing magical code."
      sections={sections}
    >
      <Lead>
        Carrion is a dynamically-typed programming language with a magical theme. Classes are called
        "grimoires", methods are "spells", and error handling uses "attempt/ensnare". This guide
        covers the basics you need to get started.
      </Lead>

      <Section id="hello-world">
        <SectionTitle>Hello World</SectionTitle>
        <Paragraph>
          Create a file called <InlineCode>hello.crl</InlineCode> with the following content:
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
            {'print("Hello, Carrion!")'}
          </SyntaxHighlighter>
        </CodeBlock>

        <Paragraph>Run it:</Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
            {'carrion hello.crl'}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>File Extension</InfoTitle>
          <InfoText>
            Carrion source files use the <InlineCode>.crl</InlineCode> extension.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="variables">
        <SectionTitle>Variables</SectionTitle>
        <Paragraph>
          Variables in Carrion are dynamically typed. No type declarations needed.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`name = "Alice"
age = 25
height = 5.9
is_student = True

# String interpolation with f-strings
greeting = f"Hello, {name}! You are {age} years old."
print(greeting)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <TipBox>
          <TipTitle>String Interpolation</TipTitle>
          <InfoText>
            Use f-strings for embedding variables: <InlineCode>f"Hello, &#123;name&#125;!"</InlineCode>
          </InfoText>
        </TipBox>
      </Section>

      <Section id="data-types">
        <SectionTitle>Data Types</SectionTitle>
        <Paragraph>Carrion supports the following data types:</Paragraph>

        <ComparisonTable>
          <ComparisonItem>
            <ComparisonLabel>Integer</ComparisonLabel>
            <ComparisonValue>42, -10</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>Float</ComparisonLabel>
            <ComparisonValue>3.14, -2.5</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>String</ComparisonLabel>
            <ComparisonValue>"hello", 'world'</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>Boolean</ComparisonLabel>
            <ComparisonValue>True, False</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>Array</ComparisonLabel>
            <ComparisonValue>[1, 2, 3]</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>Hash</ComparisonLabel>
            <ComparisonValue>&#123;"key": "value"&#125;</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>None</ComparisonLabel>
            <ComparisonValue>None</ComparisonValue>
          </ComparisonItem>
        </ComparisonTable>

        <SubSection>
          <SubSectionTitle>Arrays</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`numbers = [1, 2, 3, 4, 5]
mixed = ["text", 42, True, [1, 2]]

# Access elements
first = numbers[0]       # 1
last = numbers[-1]       # 5 (negative indexing)

# Array methods
numbers.append(6)
print(numbers.length())  # 6`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Hashes (Dictionaries)</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`person = {
    "name": "Alice",
    "age": 25,
    "city": "NYC"
}

# Access values
print(person["name"])  # Alice

# Add/update
person["email"] = "alice@example.com"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="functions">
        <SectionTitle>Functions</SectionTitle>
        <Paragraph>
          Functions in Carrion are called "spells". Define them with the <InlineCode>spell</InlineCode> keyword.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell greet(name):
    return f"Hello, {name}!"

spell add(a, b):
    return a + b

# Call functions
message = greet("Alice")
print(message)  # Hello, Alice!

result = add(5, 3)
print(result)   # 8`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSection>
          <SubSectionTitle>Default Parameters</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell greet(name, greeting="Hello"):
    return f"{greeting}, {name}!"

print(greet("Alice"))           # Hello, Alice!
print(greet("Bob", "Hi"))       # Hi, Bob!`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="classes">
        <SectionTitle>Grimoires (Classes)</SectionTitle>
        <Paragraph>
          Classes in Carrion are called "grimoires" - spellbooks that contain methods (spells).
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

# Create an instance
alice = Person("Alice", 30)
print(alice.greet())     # Hello, I'm Alice
print(alice.birthday())  # Now 31 years old`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Magical Terminology</InfoTitle>
          <InfoText>
            <InlineCode>grim</InlineCode> = class, <InlineCode>spell</InlineCode> = method/function, <InlineCode>init</InlineCode> = constructor
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="error-handling">
        <SectionTitle>Error Handling</SectionTitle>
        <Paragraph>
          Use <InlineCode>attempt</InlineCode>/<InlineCode>ensnare</InlineCode>/<InlineCode>resolve</InlineCode> for error handling (like try/catch/finally).
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell divide(a, b):
    attempt:
        result = a / b
        return result
    ensnare (ZeroDivisionError):
        print("Cannot divide by zero!")
        return None
    resolve:
        print("Division complete")

print(divide(10, 2))  # 5.0
print(divide(10, 0))  # Cannot divide by zero! -> None`}
          </SyntaxHighlighter>
        </CodeBlock>

        <ComparisonTable>
          <ComparisonItem>
            <ComparisonLabel>attempt</ComparisonLabel>
            <ComparisonValue>try</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>ensnare</ComparisonLabel>
            <ComparisonValue>catch/except</ComparisonValue>
          </ComparisonItem>
          <ComparisonItem>
            <ComparisonLabel>resolve</ComparisonLabel>
            <ComparisonValue>finally</ComparisonValue>
          </ComparisonItem>
        </ComparisonTable>
      </Section>

      <Section id="repl">
        <SectionTitle>Using the REPL</SectionTitle>
        <Paragraph>
          Start the interactive REPL by running <InlineCode>carrion</InlineCode> without arguments.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`$ carrion
Welcome to Carrion REPL v0.1.9

>>> name = "Coder"
>>> f"Hello, {name}!"
'Hello, Coder!'

>>> version()
Carrion v0.1.9

>>> mimir     # Get help
>>> quit      # Exit`}
          </SyntaxHighlighter>
        </CodeBlock>

        <Paragraph>REPL features:</Paragraph>
        <List>
          <ListItem>Tab completion for keywords and functions</ListItem>
          <ListItem>Command history with arrow keys</ListItem>
          <ListItem>Type <InlineCode>mimir</InlineCode> for interactive help</ListItem>
          <ListItem>Type <InlineCode>version()</InlineCode> to check the version</ListItem>
        </List>
      </Section>

      <Section>
        <SectionTitle>Next Steps</SectionTitle>
        <CardGrid>
          <Card as={Link} to="/docs/quick-start" style={{ textDecoration: 'none' }}>
            <CardTitle>Quick Start Tutorial</CardTitle>
            <CardDescription>Build your first Carrion project step by step.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/language-reference" style={{ textDecoration: 'none' }}>
            <CardTitle>Language Reference</CardTitle>
            <CardDescription>Complete guide to Carrion's syntax and features.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/grimoires" style={{ textDecoration: 'none' }}>
            <CardTitle>Grimoires (OOP)</CardTitle>
            <CardDescription>Learn about classes, inheritance, and more.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/standard-library" style={{ textDecoration: 'none' }}>
            <CardTitle>Standard Library</CardTitle>
            <CardDescription>Explore Munin, the standard library.</CardDescription>
          </Card>
        </CardGrid>
      </Section>
    </DocLayout>
  );
};

export default GettingStarted;
