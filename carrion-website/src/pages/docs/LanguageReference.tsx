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
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableCell,
  CardGrid,
  Card,
  CardTitle,
  CardDescription,
  InlineCode,
  List,
  ListItem,
} from '../../components/docs';

const sections = [
  { id: 'overview', title: 'Overview' },
  { id: 'lexical', title: 'Lexical Structure' },
  { id: 'types', title: 'Data Types' },
  { id: 'type-annotations', title: 'Type Annotations' },
  { id: 'operators', title: 'Operators' },
  { id: 'control-flow', title: 'Control Flow' },
  { id: 'functions', title: 'Functions' },
  { id: 'oop', title: 'Classes' },
  { id: 'error-handling', title: 'Error Handling' },
  { id: 'concurrency', title: 'Concurrency' },
  { id: 'modules', title: 'Modules' },
];

const LanguageReference: React.FC = () => {
  return (
    <DocLayout
      title="Language Reference"
      description="Complete reference guide to the Carrion programming language syntax and features."
      sections={sections}
    >
      <Section id="overview">
        <SectionTitle>Overview</SectionTitle>
        <Lead>
          Carrion is a dynamically typed, interpreted language with optional strict typing.
          Classes are "grimoires", methods are "spells", and the standard library is named
          Munin after Odin's raven.
        </Lead>

        <List>
          <ListItem><strong>Dynamic Typing</strong> - No type declarations required by default</ListItem>
          <ListItem><strong>Optional Type Hints</strong> - Lock variables to specific types when needed</ListItem>
          <ListItem><strong>Object-Oriented</strong> - Full support for classes and inheritance</ListItem>
          <ListItem><strong>Concurrent</strong> - Built-in concurrency with diverge/converge</ListItem>
          <ListItem><strong>File Extension</strong> - <InlineCode>.crl</InlineCode></ListItem>
        </List>
      </Section>

      <Section id="lexical">
        <SectionTitle>Lexical Structure</SectionTitle>

        <SubSection>
          <SubSectionTitle>Comments</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Single-line comment
/* Multi-line
   comment */
# Also single-line`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Keywords</SubSectionTitle>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Category</TableHead>
                <TableHead>Keywords</TableHead>
              </TableRow>
            </TableHeader>
            <tbody>
              <TableRow>
                <TableCell>Control Flow</TableCell>
                <TableCell><InlineCode>if</InlineCode> <InlineCode>otherwise</InlineCode> <InlineCode>else</InlineCode> <InlineCode>for</InlineCode> <InlineCode>while</InlineCode> <InlineCode>match</InlineCode> <InlineCode>case</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Loop Control</TableCell>
                <TableCell><InlineCode>skip</InlineCode> <InlineCode>stop</InlineCode> <InlineCode>return</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>OOP</TableCell>
                <TableCell><InlineCode>grim</InlineCode> <InlineCode>spell</InlineCode> <InlineCode>init</InlineCode> <InlineCode>self</InlineCode> <InlineCode>super</InlineCode> <InlineCode>arcane</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Error Handling</TableCell>
                <TableCell><InlineCode>attempt</InlineCode> <InlineCode>ensnare</InlineCode> <InlineCode>resolve</InlineCode> <InlineCode>raise</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Logical</TableCell>
                <TableCell><InlineCode>and</InlineCode> <InlineCode>or</InlineCode> <InlineCode>not</InlineCode> <InlineCode>True</InlineCode> <InlineCode>False</InlineCode> <InlineCode>None</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Modules</TableCell>
                <TableCell><InlineCode>import</InlineCode> <InlineCode>as</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Concurrency</TableCell>
                <TableCell><InlineCode>diverge</InlineCode> <InlineCode>converge</InlineCode></TableCell>
              </TableRow>
            </tbody>
          </Table>
        </SubSection>
      </Section>

      <Section id="types">
        <SectionTitle>Data Types</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Type</TableHead>
              <TableHead>Description</TableHead>
              <TableHead>Example</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell>Integer</TableCell>
              <TableCell>64-bit signed integer</TableCell>
              <TableCell><InlineCode>42</InlineCode>, <InlineCode>-10</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Float</TableCell>
              <TableCell>64-bit floating-point</TableCell>
              <TableCell><InlineCode>3.14</InlineCode>, <InlineCode>-2.5</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell>String</TableCell>
              <TableCell>UTF-8 text</TableCell>
              <TableCell><InlineCode>"hello"</InlineCode>, <InlineCode>'world'</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Boolean</TableCell>
              <TableCell>True or False</TableCell>
              <TableCell><InlineCode>True</InlineCode>, <InlineCode>False</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Array</TableCell>
              <TableCell>Ordered collection</TableCell>
              <TableCell><InlineCode>[1, 2, 3]</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Hash</TableCell>
              <TableCell>Key-value pairs</TableCell>
              <TableCell><InlineCode>&#123;"a": 1&#125;</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell>None</TableCell>
              <TableCell>Null value</TableCell>
              <TableCell><InlineCode>None</InlineCode></TableCell>
            </TableRow>
          </tbody>
        </Table>

        <SubSection>
          <SubSectionTitle>String Interpolation</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`name = "Alice"
age = 25

// F-strings for interpolation
greeting = f"Hello, {name}! You are {age}."

// Multi-line strings
text = """This spans
multiple lines"""`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="type-annotations">
        <SectionTitle>Type Annotations</SectionTitle>
        <Paragraph>
          Carrion supports optional type hints for variables and function parameters. When a type
          hint is provided, the variable is locked to that type and cannot be reassigned to a
          value of a different type.
        </Paragraph>

        <SubSection>
          <SubSectionTitle>Basic Type Hints</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Variable with type hint - locked to int
x: int = 10
x = 20       // OK: still an int
x = "hello"  // Error: cannot assign String to variable with type hint int

// Without type hint - can be reassigned freely
y = 10
y = "hello"  // OK: no type restriction`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Function Parameter Types</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell greet(name: str, count: int):
    for i in range(count):
        print(f"Hello, {name}!")

greet("Alice", 3)  // OK
greet(123, 3)      // Error: type mismatch`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Available Types</SubSectionTitle>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Type</TableHead>
                <TableHead>Description</TableHead>
                <TableHead>Example</TableHead>
              </TableRow>
            </TableHeader>
            <tbody>
              <TableRow>
                <TableCell><InlineCode>int</InlineCode></TableCell>
                <TableCell>Integer numbers</TableCell>
                <TableCell><InlineCode>count: int = 0</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>float</InlineCode></TableCell>
                <TableCell>Floating-point numbers</TableCell>
                <TableCell><InlineCode>price: float = 9.99</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>str</InlineCode></TableCell>
                <TableCell>Text strings</TableCell>
                <TableCell><InlineCode>name: str = "Bob"</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>bool</InlineCode></TableCell>
                <TableCell>Boolean values</TableCell>
                <TableCell><InlineCode>active: bool = True</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>array</InlineCode></TableCell>
                <TableCell>Arrays/lists</TableCell>
                <TableCell><InlineCode>items: array = []</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>hash</InlineCode></TableCell>
                <TableCell>Hash maps/dictionaries</TableCell>
                <TableCell><InlineCode>data: hash = &#123;&#125;</InlineCode></TableCell>
              </TableRow>
            </tbody>
          </Table>
        </SubSection>

        <InfoBox>
          <InfoTitle>Gradual Typing</InfoTitle>
          <InfoText>
            Type hints are optional. Variables without type hints behave dynamically and can
            be reassigned to any type. Add type hints where you want to enforce type safety.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="operators">
        <SectionTitle>Operators</SectionTitle>

        <SubSection>
          <SubSectionTitle>Quick Reference</SubSectionTitle>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Type</TableHead>
                <TableHead>Operators</TableHead>
              </TableRow>
            </TableHeader>
            <tbody>
              <TableRow>
                <TableCell>Arithmetic</TableCell>
                <TableCell><InlineCode>+</InlineCode> <InlineCode>-</InlineCode> <InlineCode>*</InlineCode> <InlineCode>/</InlineCode> <InlineCode>%</InlineCode> <InlineCode>**</InlineCode> <InlineCode>{'//'}</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Comparison</TableCell>
                <TableCell><InlineCode>==</InlineCode> <InlineCode>!=</InlineCode> <InlineCode>&lt;</InlineCode> <InlineCode>&gt;</InlineCode> <InlineCode>&lt;=</InlineCode> <InlineCode>&gt;=</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Logical</TableCell>
                <TableCell><InlineCode>and</InlineCode> <InlineCode>or</InlineCode> <InlineCode>not</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Assignment</TableCell>
                <TableCell><InlineCode>=</InlineCode> <InlineCode>+=</InlineCode> <InlineCode>-=</InlineCode> <InlineCode>*=</InlineCode> <InlineCode>/=</InlineCode></TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Membership</TableCell>
                <TableCell><InlineCode>in</InlineCode> <InlineCode>not in</InlineCode></TableCell>
              </TableRow>
            </tbody>
          </Table>
        </SubSection>

        <InfoBox>
          <InfoTitle>Learn More</InfoTitle>
          <InfoText>
            See the <Link to="/docs/operators">Operators</Link> page for detailed examples of each operator.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="control-flow">
        <SectionTitle>Control Flow</SectionTitle>

        <SubSection>
          <SubSectionTitle>Conditionals</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`if score >= 90:
    grade = "A"
otherwise score >= 80:    // like elif
    grade = "B"
else:
    grade = "C"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Loops</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// For loop
for item in [1, 2, 3]:
    print(item)

// While loop
while count < 10:
    count += 1

// Loop control
for i in range(10):
    if i == 3:
        skip     // continue
    if i == 7:
        stop     // break`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Pattern Matching</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`match status_code:
    case 200:
        message = "OK"
    case 404:
        message = "Not Found"
    _:
        message = "Unknown"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <InfoBox>
          <InfoTitle>Learn More</InfoTitle>
          <InfoText>
            See the <Link to="/docs/control-flow">Control Flow</Link> page for comprehensive examples.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="functions">
        <SectionTitle>Functions (Spells)</SectionTitle>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Define a function
spell greet(name, greeting="Hello"):
    return f"{greeting}, {name}!"

// Call it
print(greet("Alice"))        // Hello, Alice!
print(greet("Bob", "Hi"))    // Hi, Bob!

// Lambda expressions
double = (x) -> x * 2
numbers = [1, 2, 3].map((x) -> x * 2)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="oop">
        <SectionTitle>Classes (Grimoires)</SectionTitle>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Define a class
grim Person:
    init(name, age):
        self.name = name
        self.age = age

    spell greet():
        return f"I'm {self.name}"

// Inheritance
grim Student(Person):
    init(name, age, school):
        super.init(name, age)
        self.school = school

// Create instances
alice = Person("Alice", 30)
bob = Student("Bob", 20, "MIT")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Learn More</InfoTitle>
          <InfoText>
            See the <Link to="/docs/grimoires">Grimoires</Link> page for inheritance, abstract classes, and design patterns.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="error-handling">
        <SectionTitle>Error Handling</SectionTitle>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell divide(a, b):
    attempt:
        return a / b
    ensnare (ZeroDivisionError):
        print("Division by zero!")
        return None
    resolve:
        print("Done")

// Raise errors
raise Error("Something went wrong")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Carrion</TableHead>
              <TableHead>Python/JS Equivalent</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>attempt</InlineCode></TableCell>
              <TableCell><InlineCode>try</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>ensnare</InlineCode></TableCell>
              <TableCell><InlineCode>except</InlineCode> / <InlineCode>catch</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>resolve</InlineCode></TableCell>
              <TableCell><InlineCode>finally</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>raise</InlineCode></TableCell>
              <TableCell><InlineCode>raise</InlineCode> / <InlineCode>throw</InlineCode></TableCell>
            </TableRow>
          </tbody>
        </Table>
      </Section>

      <Section id="concurrency">
        <SectionTitle>Concurrency</SectionTitle>
        <Paragraph>
          Carrion supports concurrent execution using <InlineCode>diverge</InlineCode> and <InlineCode>converge</InlineCode>.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Start concurrent tasks
diverge:
    task1()
    task2()
    task3()

// Wait for all tasks
converge

// Diverge returns array of results
results = diverge:
    fetch_data(url1)
    fetch_data(url2)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="modules">
        <SectionTitle>Modules</SectionTitle>
        <Paragraph>
          Import code from other files using the <InlineCode>import</InlineCode> statement.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Import a module
import math

// Import with alias
import utils as u

// Use imported code
result = math.sqrt(16)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Learn More</InfoTitle>
          <InfoText>
            See the <Link to="/docs/modules">Modules</Link> page for details on creating and organizing modules.
          </InfoText>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>Deep Dive Topics</SectionTitle>
        <CardGrid>
          <Card as={Link} to="/docs/operators" style={{ textDecoration: 'none' }}>
            <CardTitle>Operators</CardTitle>
            <CardDescription>Arithmetic, comparison, logical, and special operators.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/control-flow" style={{ textDecoration: 'none' }}>
            <CardTitle>Control Flow</CardTitle>
            <CardDescription>Conditionals, loops, and pattern matching in depth.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/grimoires" style={{ textDecoration: 'none' }}>
            <CardTitle>Grimoires (OOP)</CardTitle>
            <CardDescription>Classes, inheritance, and object-oriented patterns.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/error-handling" style={{ textDecoration: 'none' }}>
            <CardTitle>Error Handling</CardTitle>
            <CardDescription>Attempt, ensnare, resolve, and custom errors.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/standard-library" style={{ textDecoration: 'none' }}>
            <CardTitle>Standard Library</CardTitle>
            <CardDescription>Munin standard library modules and functions.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/builtin-functions" style={{ textDecoration: 'none' }}>
            <CardTitle>Built-in Functions</CardTitle>
            <CardDescription>All globally available functions.</CardDescription>
          </Card>
        </CardGrid>
      </Section>
    </DocLayout>
  );
};

export default LanguageReference;
