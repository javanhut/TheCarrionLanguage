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
  margin-bottom: 3rem;
`;

const Title = styled.h1`
  font-size: 3rem;
  margin-bottom: 1rem;
  background: ${({ theme }) => theme.gradients.primary};
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
`;

const Subtitle = styled.p`
  font-size: 1.3rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  max-width: 800px;
  margin: 0 auto;
`;

const Navigation = styled.nav`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 15px;
  padding: 2rem;
  margin-bottom: 3rem;
  position: sticky;
  top: 80px;
  z-index: 10;
`;

const NavTitle = styled.h3`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1rem;
`;

const NavGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 0.5rem;
`;

const NavLink = styled.a`
  color: ${({ theme }) => theme.colors.text.primary};
  text-decoration: none;
  padding: 0.5rem;
  border-radius: 5px;
  transition: all ${({ theme }) => theme.transitions.fast};
  
  &:hover {
    background: ${({ theme }) => theme.colors.background.tertiary};
    color: ${({ theme }) => theme.colors.primary};
  }
`;

const Section = styled.section`
  margin-bottom: 4rem;
  scroll-margin-top: 100px;
`;

const SectionTitle = styled.h2`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1.5rem;
  font-size: 2.5rem;
  border-bottom: 2px solid ${({ theme }) => theme.colors.border};
  padding-bottom: 0.5rem;
`;

const SubSectionTitle = styled.h3`
  color: ${({ theme }) => theme.colors.text.primary};
  margin: 2rem 0 1rem;
  font-size: 1.8rem;
`;

const SubSubSectionTitle = styled.h4`
  color: ${({ theme }) => theme.colors.text.primary};
  margin: 1.5rem 0 1rem;
  font-size: 1.3rem;
`;

const CodeBlock = styled.div`
  margin: 1.5rem 0;
  border-radius: 10px;
  overflow: hidden;
`;

const InlineCode = styled.code`
  background: ${({ theme }) => theme.colors.background.secondary};
  color: ${({ theme }) => theme.colors.primary};
  padding: 0.2rem 0.4rem;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 0.9em;
`;

const Table = styled.table`
  width: 100%;
  border-collapse: collapse;
  margin: 1.5rem 0;
  background: ${({ theme }) => theme.colors.background.secondary};
  border-radius: 10px;
  overflow: hidden;

  th, td {
    padding: 1rem;
    text-align: left;
    border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  }

  th {
    background: ${({ theme }) => theme.colors.background.tertiary};
    color: ${({ theme }) => theme.colors.primary};
    font-weight: 600;
  }

  tr:last-child td {
    border-bottom: none;
  }

  code {
    background: ${({ theme }) => theme.colors.background.tertiary};
    padding: 0.2rem 0.4rem;
    border-radius: 3px;
    font-size: 0.9em;
  }
`;

const InfoBox = styled.div`
  background: rgba(0, 204, 153, 0.1);
  border-left: 4px solid ${({ theme }) => theme.colors.primary};
  border-radius: 8px;
  padding: 1.5rem;
  margin: 1.5rem 0;
`;

const WarningBox = styled.div`
  background: rgba(255, 204, 0, 0.1);
  border-left: 4px solid #ffcc00;
  border-radius: 8px;
  padding: 1.5rem;
  margin: 1.5rem 0;
`;

const FeatureCard = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  padding: 1.5rem;
  margin: 1rem 0;
`;

const List = styled.ul`
  margin: 1rem 0;
  padding-left: 2rem;
  line-height: 1.8;
  
  li {
    margin: 0.5rem 0;
  }
`;

const LanguageReference: React.FC = () => {
  const scrollToSection = (id: string) => {
    const element = document.getElementById(id);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
    }
  };

  return (
    <Container>
      <Header>
        <Title>Carrion Language Reference</Title>
        <Subtitle>
          Complete reference guide covering all aspects of the Carrion programming language,
          from basic syntax to advanced features
        </Subtitle>
      </Header>

      <Navigation>
        <NavTitle>Quick Navigation</NavTitle>
        <NavGrid>
          <NavLink onClick={() => scrollToSection('overview')}>Overview</NavLink>
          <NavLink onClick={() => scrollToSection('lexical')}>Lexical Structure</NavLink>
          <NavLink onClick={() => scrollToSection('types')}>Data Types</NavLink>
          <NavLink onClick={() => scrollToSection('operators')}>Operators</NavLink>
          <NavLink onClick={() => scrollToSection('control')}>Control Flow</NavLink>
          <NavLink onClick={() => scrollToSection('functions')}>Functions</NavLink>
          <NavLink onClick={() => scrollToSection('oop')}>OOP/Grimoires</NavLink>
          <NavLink onClick={() => scrollToSection('errors')}>Error Handling</NavLink>
          <NavLink onClick={() => scrollToSection('types-hints')}>Type Hints</NavLink>
          <NavLink onClick={() => scrollToSection('concurrency')}>Concurrency</NavLink>
          <NavLink onClick={() => scrollToSection('main')}>Main Entry Point</NavLink>
          <NavLink onClick={() => scrollToSection('modules')}>Modules</NavLink>
          <NavLink onClick={() => scrollToSection('builtins')}>Built-in Functions</NavLink>
          <NavLink onClick={() => scrollToSection('stdlib')}>Standard Library</NavLink>
        </NavGrid>
      </Navigation>

      {/* OVERVIEW */}
      <Section id="overview">
        <SectionTitle>Language Overview</SectionTitle>
        
        <p>
          Carrion is a dynamically typed, interpreted programming language with a Norse mythology
          and magical theme. It combines familiar programming concepts with enchanting terminology
          while maintaining readable and practical syntax.
        </p>

        <SubSectionTitle>Key Characteristics</SubSectionTitle>
        <List>
          <li><strong>Dynamic Typing:</strong> Variables don't require explicit type declarations</li>
          <li><strong>Interpreted:</strong> Code is executed directly without compilation</li>
          <li><strong>Object-Oriented:</strong> Full support for classes (grimoires) and inheritance</li>
          <li><strong>Concurrent:</strong> Built-in concurrency with diverge/converge keywords</li>
          <li><strong>Duck Typing:</strong> Objects are characterized by their behavior, not their type</li>
          <li><strong>Memory Managed:</strong> Automatic memory management via Go's garbage collector</li>
        </List>

        <SubSectionTitle>File Extension</SubSectionTitle>
        <p>Carrion source files use the <InlineCode>.crl</InlineCode> extension.</p>

        <SubSectionTitle>Philosophy</SubSectionTitle>
        <p>
          Named after the Carrion Crow and inspired by Norse mythology (with the standard library
          named Munin after Odin's raven), Carrion aims to make programming both powerful and enjoyable
          through creative syntax modifications while maintaining ease of use.
        </p>
      </Section>

      {/* LEXICAL STRUCTURE */}
      <Section id="lexical">
        <SectionTitle>Lexical Structure</SectionTitle>

        <SubSectionTitle>Comments</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Single-line comment
/* Multi-line
   comment */`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Identifiers</SubSectionTitle>
        <p>
          Identifiers must start with a letter or underscore, followed by letters, digits, or underscores:
        </p>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`variable_name
_private_var
MyClass
function123`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Keywords</SubSectionTitle>
        <p>Reserved words that cannot be used as identifiers:</p>
        
        <Table>
          <thead>
            <tr>
              <th>Category</th>
              <th>Keywords</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Control Flow</td>
              <td><code>if</code> <code>otherwise</code> <code>else</code> <code>for</code> <code>in</code> <code>while</code> <code>match</code> <code>case</code></td>
            </tr>
            <tr>
              <td>Loop Control</td>
              <td><code>skip</code> <code>stop</code> <code>return</code></td>
            </tr>
            <tr>
              <td>OOP</td>
              <td><code>grim</code> <code>spell</code> <code>init</code> <code>self</code> <code>super</code> <code>arcane</code> <code>arcanespell</code></td>
            </tr>
            <tr>
              <td>Error Handling</td>
              <td><code>attempt</code> <code>ensnare</code> <code>resolve</code> <code>raise</code> <code>check</code></td>
            </tr>
            <tr>
              <td>Logical</td>
              <td><code>and</code> <code>or</code> <code>not</code> <code>True</code> <code>False</code> <code>None</code></td>
            </tr>
            <tr>
              <td>Modules</td>
              <td><code>import</code> <code>as</code></td>
            </tr>
            <tr>
              <td>Concurrency</td>
              <td><code>diverge</code> <code>converge</code></td>
            </tr>
            <tr>
              <td>Special</td>
              <td><code>main</code> <code>var</code> <code>ignore</code></td>
            </tr>
          </tbody>
        </Table>

        <SubSectionTitle>Literals</SubSectionTitle>
        
        <SubSubSectionTitle>Integer Literals</SubSubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`42          // Decimal
-10         // Negative integer
1000000     // Large number`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSubSectionTitle>Float Literals</SubSubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`3.14
2.718
1.0
.5          // 0.5`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSubSectionTitle>String Literals</SubSubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`"double quotes"
'single quotes'
"""triple quotes for 
   multi-line strings"""
f"formatted {variable}"      // F-strings
i"interpolated {expression}" // Interpolated strings`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSubSectionTitle>Boolean and None Literals</SubSubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`True
False
None`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      {/* DATA TYPES */}
      <Section id="types">
        <SectionTitle>Data Types</SectionTitle>

        <SubSectionTitle>Primitive Types</SubSectionTitle>

        <FeatureCard>
          <SubSubSectionTitle>Integer</SubSubSectionTitle>
          <p>64-bit signed integers</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`age = 25
count = -10
big_number = 1000000`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <FeatureCard>
          <SubSubSectionTitle>Float</SubSubSectionTitle>
          <p>64-bit floating-point numbers</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`pi = 3.14159
temperature = -15.5
rate = 0.075`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <FeatureCard>
          <SubSubSectionTitle>String</SubSubSectionTitle>
          <p>UTF-8 text strings with indexing support</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`name = "Alice"
message = 'Hello, World!'
description = """This is a
multi-line string"""

// String indexing
s = "Hello World"
print(s[0])     // "H"
print(s[-1])    // "d"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <FeatureCard>
          <SubSubSectionTitle>Boolean</SubSubSectionTitle>
          <p>True/False values</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`is_active = True
has_permission = False`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <FeatureCard>
          <SubSubSectionTitle>None</SubSubSectionTitle>
          <p>Represents absence of value</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`result = None
optional_param = None`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <SubSectionTitle>Collection Types</SubSectionTitle>

        <FeatureCard>
          <SubSubSectionTitle>Array</SubSubSectionTitle>
          <p>Ordered, mutable sequences</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`numbers = [1, 2, 3, 4, 5]
mixed = [1, "hello", True, 3.14]
empty = []

// Array access
numbers[0]     // 1
numbers[-1]    // 5`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <FeatureCard>
          <SubSubSectionTitle>Hash</SubSubSectionTitle>
          <p>Key-value mappings (dictionaries)</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`person = {"name": "Alice", "age": 30}
config = {"debug": True, "timeout": 30}
empty_hash = {}`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <FeatureCard>
          <SubSubSectionTitle>Tuple</SubSubSectionTitle>
          <p>Immutable ordered sequences</p>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`coordinates = (10, 20)
rgb = (255, 128, 0)
single = (42,)  // Single-element tuple

// Tuple unpacking
x, y = (10, 20)`}
            </SyntaxHighlighter>
          </CodeBlock>
        </FeatureCard>

        <SubSectionTitle>Type Checking</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`value = 42
print(type(value))  // "INTEGER"

if type(value) == "INTEGER":
    print("It's an integer")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      {/* OPERATORS */}
      <Section id="operators">
        <SectionTitle>Operators</SectionTitle>

        <SubSectionTitle>Arithmetic Operators</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Description</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>+</code></td><td>Addition</td><td><code>5 + 3</code></td><td><code>8</code></td></tr>
            <tr><td><code>-</code></td><td>Subtraction</td><td><code>5 - 3</code></td><td><code>2</code></td></tr>
            <tr><td><code>*</code></td><td>Multiplication</td><td><code>5 * 3</code></td><td><code>15</code></td></tr>
            <tr><td><code>/</code></td><td>Division</td><td><code>15 / 3</code></td><td><code>5.0</code></td></tr>
            <tr><td><code>{'//'}</code></td><td>Integer Division</td><td><code>17 {'//'} 3</code></td><td><code>5</code></td></tr>
            <tr><td><code>%</code></td><td>Modulo</td><td><code>17 % 3</code></td><td><code>2</code></td></tr>
            <tr><td><code>**</code></td><td>Exponentiation</td><td><code>2 ** 3</code></td><td><code>8</code></td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Assignment Operators</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Description</th>
              <th>Example</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>=</code></td><td>Assignment</td><td><code>x = 5</code></td></tr>
            <tr><td><code>+=</code></td><td>Add and assign</td><td><code>x += 3</code></td></tr>
            <tr><td><code>-=</code></td><td>Subtract and assign</td><td><code>x -= 3</code></td></tr>
            <tr><td><code>*=</code></td><td>Multiply and assign</td><td><code>x *= 3</code></td></tr>
            <tr><td><code>/=</code></td><td>Divide and assign</td><td><code>x /= 3</code></td></tr>
            <tr><td><code>++</code></td><td>Increment (prefix/postfix)</td><td><code>x++</code> or <code>++x</code></td></tr>
            <tr><td><code>--</code></td><td>Decrement (prefix/postfix)</td><td><code>x--</code> or <code>--x</code></td></tr>
          </tbody>
        </Table>

        <InfoBox>
          <strong>Note:</strong> Carrion supports both C-style increments (<code>++i</code>, <code>i++</code>) 
          and Python-style increments (<code>i += 1</code>).
        </InfoBox>

        <SubSectionTitle>Comparison Operators</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Description</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>==</code></td><td>Equal</td><td><code>5 == 5</code></td><td><code>True</code></td></tr>
            <tr><td><code>!=</code></td><td>Not equal</td><td><code>5 != 3</code></td><td><code>True</code></td></tr>
            <tr><td><code>&lt;</code></td><td>Less than</td><td><code>3 &lt; 5</code></td><td><code>True</code></td></tr>
            <tr><td><code>&gt;</code></td><td>Greater than</td><td><code>5 &gt; 3</code></td><td><code>True</code></td></tr>
            <tr><td><code>&lt;=</code></td><td>Less than or equal</td><td><code>3 &lt;= 5</code></td><td><code>True</code></td></tr>
            <tr><td><code>&gt;=</code></td><td>Greater than or equal</td><td><code>5 &gt;= 5</code></td><td><code>True</code></td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Logical Operators</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Description</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>and</code></td><td>Logical AND</td><td><code>True and False</code></td><td><code>False</code></td></tr>
            <tr><td><code>or</code></td><td>Logical OR</td><td><code>True or False</code></td><td><code>True</code></td></tr>
            <tr><td><code>not</code></td><td>Logical NOT</td><td><code>not True</code></td><td><code>False</code></td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Membership Operators</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Description</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>in</code></td><td>Membership test</td><td><code>"a" in "apple"</code></td><td><code>True</code></td></tr>
            <tr><td><code>not in</code></td><td>Negative membership</td><td><code>"z" not in "apple"</code></td><td><code>True</code></td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Bitwise Operators</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Description</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>&amp;</code></td><td>Bitwise AND</td><td><code>5 &amp; 3</code></td><td><code>1</code></td></tr>
            <tr><td><code>|</code></td><td>Bitwise OR</td><td><code>5 | 3</code></td><td><code>7</code></td></tr>
            <tr><td><code>^</code></td><td>Bitwise XOR</td><td><code>5 ^ 3</code></td><td><code>6</code></td></tr>
            <tr><td><code>~</code></td><td>Bitwise NOT</td><td><code>~5</code></td><td><code>-6</code></td></tr>
            <tr><td><code>&lt;&lt;</code></td><td>Left shift</td><td><code>5 &lt;&lt; 1</code></td><td><code>10</code></td></tr>
            <tr><td><code>&gt;&gt;</code></td><td>Right shift</td><td><code>10 &gt;&gt; 1</code></td><td><code>5</code></td></tr>
          </tbody>
        </Table>
      </Section>

      {/* CONTROL FLOW */}
      <Section id="control">
        <SectionTitle>Control Flow</SectionTitle>

        <SubSectionTitle>If Statements</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`if condition:
    // code
otherwise another_condition:
    // code
else:
    // code

// Example
age = 18
if age >= 18:
    print("Adult")
otherwise age >= 13:
    print("Teenager")
else:
    print("Child")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>For Loops</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`for item in iterable:
    // code

// Examples
for i in range(10):
    print(i)

for name in ["Alice", "Bob", "Charlie"]:
    print(name)

for key, value in pairs({"a": 1, "b": 2}):
    print(key, value)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>While Loops</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`while condition:
    // code

// Example
x = 0
while x < 10:
    print(x)
    x++`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Match Statements</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`match value:
    case pattern1:
        // code
    case pattern2:
        // code
    _:
        // default case

// Example
status = "success"
match status:
    case "success":
        print("Operation successful")
    case "error":
        print("Operation failed")
    _:
        print("Unknown status")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Loop Control</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`skip  // Continue to next iteration (like Python's continue)
stop  // Break from loop (like Python's break)

// Example
for i in range(10):
    if i % 2 == 0:
        skip  // Skip even numbers
    if i > 7:
        stop  // Stop at 7
    print(i)  // Prints: 1, 3, 5, 7`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      {/* FUNCTIONS */}
      <Section id="functions">
        <SectionTitle>Functions (Spells)</SectionTitle>

        <InfoBox>
          <strong>Terminology:</strong> In Carrion, functions are called "spells" and use 
          the <InlineCode>spell</InlineCode> keyword.
        </InfoBox>

        <SubSectionTitle>Function Definition</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell function_name(parameters):
    // function body
    return value  // optional`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Basic Examples</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell greet(name):
    return f"Hello, {name}!"

spell add(a, b):
    return a + b

spell print_message():
    print("This spell has no return value")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Default Parameters</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell power(base, exponent = 2):
    return base ** exponent

print(power(5))      // 25 (5^2)
print(power(5, 3))   // 125 (5^3)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Recursion</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)

print(factorial(5))  // 120`}
          </SyntaxHighlighter>
        </CodeBlock>

        <WarningBox>
          <strong>Important:</strong> Carrion does NOT support lambda/anonymous functions. 
          All functions must be named using the <InlineCode>spell</InlineCode> keyword.
        </WarningBox>
      </Section>

      {/* OOP/GRIMOIRES */}
      <Section id="oop">
        <SectionTitle>Object-Oriented Programming (Grimoires)</SectionTitle>

        <InfoBox>
          <strong>Terminology:</strong> In Carrion, classes are called "grimoires" (or "spellbooks") 
          and use the <InlineCode>grim</InlineCode> keyword.
        </InfoBox>

        <SubSectionTitle>Class Definition</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim ClassName:
    init(parameters):
        // constructor
    
    spell method_name(parameters):
        // method implementation`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Basic Example</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim Person:
    init(name, age):
        self.name = name
        self.age = age
    
    spell introduce():
        return f"I am {self.name}, {self.age} years old"
    
    spell birthday():
        self.age += 1

// Create instance
person = Person("Alice", 30)
print(person.introduce())  // "I am Alice, 30 years old"
person.birthday()
print(person.age)  // 31`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Inheritance</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim Animal:
    init(name):
        self.name = name
    
    spell speak():
        return "Some sound"

grim Dog(Animal):
    init(name, breed):
        super.init(name)
        self.breed = breed
    
    spell speak():
        return "Woof!"
    
    spell get_info():
        return f"{self.name} is a {self.breed}"

dog = Dog("Buddy", "Golden Retriever")
print(dog.speak())      // "Woof!"
print(dog.get_info())   // "Buddy is a Golden Retriever"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Abstract Classes</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`arcane grim Shape:
    init():
        ignore
    
    @arcanespell
    spell area():
        ignore
    
    @arcanespell
    spell perimeter():
        ignore

grim Circle(Shape):
    init(radius):
        self.radius = radius
    
    spell area():
        return 3.14159 * self.radius ** 2
    
    spell perimeter():
        return 2 * 3.14159 * self.radius`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <List>
            <li><InlineCode>arcane</InlineCode> - Declares an abstract class</li>
            <li><InlineCode>@arcanespell</InlineCode> - Decorator for abstract methods</li>
            <li><InlineCode>ignore</InlineCode> - Empty placeholder for method body</li>
          </List>
        </InfoBox>

        <SubSectionTitle>Encapsulation</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim BankAccount:
    init(balance):
        self.__balance = balance  // Protected (double underscore)
        self._account_id = 12345  // Private (single underscore)
    
    spell get_balance():
        return self.__balance
    
    spell deposit(amount):
        self.__balance += amount
    
    spell __internal_audit():  // Protected method
        return self.__balance * 0.01

account = BankAccount(1000)
print(account.get_balance())  // 1000
account.deposit(500)
// account.__balance  // Error: Cannot access protected attribute
// account.__internal_audit()  // Error: Cannot access protected method`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <strong>Access Modifiers:</strong>
          <List>
            <li><InlineCode>_attribute</InlineCode> - Private (single underscore)</li>
            <li><InlineCode>__attribute</InlineCode> - Protected (double underscore)</li>
            <li><InlineCode>attribute</InlineCode> - Public (no underscore)</li>
          </List>
        </InfoBox>
      </Section>

      {/* ERROR HANDLING */}
      <Section id="errors">
        <SectionTitle>Error Handling</SectionTitle>

        <InfoBox>
          <strong>Terminology:</strong>
          <List>
            <li><InlineCode>attempt</InlineCode> - try block</li>
            <li><InlineCode>ensnare</InlineCode> - catch/except block</li>
            <li><InlineCode>resolve</InlineCode> - finally block</li>
          </List>
        </InfoBox>

        <SubSectionTitle>Basic Syntax</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`attempt:
    // risky code
ensnare (ErrorType):
    // handle specific error
ensnare:
    // handle any error
resolve:
    // finally block (always runs)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Example</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell divide(a, b):
    attempt:
        result = a / b
        return result
    ensnare (ZeroDivisionError):
        print("Cannot divide by zero!")
        return None
    ensnare:
        print("An unexpected error occurred")
        return None
    resolve:
        print("Division operation completed")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Raising Errors</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell validate_age(age):
    if age < 0:
        raise Error("ValueError", "Age cannot be negative")
    if age > 150:
        raise Error("ValueError", "Age is unrealistic")
    return True`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Assertions</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`check(condition, "Error message")

// Examples
x = 10
check(x == 10, f"x should equal 10 got: {x}")  // Passes

y = 5
check(y == 10, f"y should equal 10 got: {y}")  // Raises assertion error`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Custom Error Classes</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim ValueError:
    init(message):
        self.message = message

attempt:
    raise ValueError("Invalid value")
ensnare (ValueError):
    print("Caught ValueError!")
resolve:
    print("Cleanup complete")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      {/* TYPE HINTS */}
      <Section id="types-hints">
        <SectionTitle>Type Hints (Optional)</SectionTitle>

        <InfoBox>
          <strong>Important:</strong> Type hints in Carrion are optional, non-enforcing, and serve 
          as documentation only. They do not provide runtime type checking.
        </InfoBox>

        <SubSectionTitle>Function Parameter Type Hints</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell calculate_area(width: int, height: int):
    return width * height

spell greet(name: str = "World"):
    print("Hello, " + name)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Return Type Hints</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell add(a: int, b: int) -> int:
    return a + b

spell get_user_info(id: int) -> dict:
    return {"id": id, "name": "John Doe"}

spell process_data(items: list) -> list:
    return [item * 2 for item in items]`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Variable Type Hints</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`count: int = 0
name: str = "Alice"
scores: list = [85, 92, 78]
config: dict = {"debug": True}
active: bool = True`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Supported Type Annotations</SubSectionTitle>
        <List>
          <li><strong>Primitive Types:</strong> <InlineCode>int</InlineCode>, <InlineCode>float</InlineCode>, <InlineCode>str</InlineCode>, <InlineCode>bool</InlineCode></li>
          <li><strong>Collection Types:</strong> <InlineCode>list</InlineCode>, <InlineCode>dict</InlineCode>, <InlineCode>set</InlineCode></li>
          <li><strong>Special Types:</strong> <InlineCode>None</InlineCode>, <InlineCode>any</InlineCode></li>
          <li><strong>Custom Types:</strong> Grimoire class names</li>
        </List>

        <SubSectionTitle>Complex Example</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`grim Calculator:
    init(precision: int = 2):
        self.precision: int = precision
    
    spell divide(a: float, b: float) -> float:
        if b == 0:
            raise ValueError("Division by zero")
        return round(a / b, self.precision)
    
    spell add_numbers(numbers: list) -> float:
        total: float = 0.0
        for num in numbers:
            total += num
        return total`}
          </SyntaxHighlighter>
        </CodeBlock>

        <WarningBox>
          <strong>Current Behavior:</strong>
          <List>
            <li>Optional - not required</li>
            <li>Documentation only</li>
            <li>No runtime type checking</li>
            <li>Prepared for future static analysis</li>
          </List>
        </WarningBox>
      </Section>

      {/* CONCURRENCY */}
      <Section id="concurrency">
        <SectionTitle>Concurrency (diverge/converge)</SectionTitle>

        <InfoBox>
          <strong>New Feature:</strong> Carrion provides built-in concurrency support through 
          the <InlineCode>diverge</InlineCode> and <InlineCode>converge</InlineCode> keywords, 
          built on top of Go's goroutines.
        </InfoBox>

        <SubSectionTitle>The diverge Keyword</SubSectionTitle>
        <p>Creates a new goroutine that executes code concurrently:</p>
        
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Anonymous goroutine
diverge:
    print("Running concurrently")
    // code block

// Named goroutine
diverge worker_name:
    print("Named goroutine")
    // code block`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>The converge Keyword</SubSectionTitle>
        <p>Waits for goroutines to complete:</p>
        
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Wait for all goroutines
converge

// Wait for specific named goroutines
converge worker1, worker2, worker3`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Basic Example</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`main:
    print("Starting concurrent processing")
    
    diverge task1:
        print("Task 1: Processing...")
        sleep(2000)
        print("Task 1: Complete")
    
    diverge task2:
        print("Task 2: Calculating...")
        sleep(1500)
        print("Task 2: Complete")
    
    print("Main thread continues...")
    
    converge task1, task2
    print("All tasks completed!")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Producer-Consumer Pattern</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`main:
    diverge producer:
        for i in range(10):
            print("Producing item", i)
            sleep(200)
    
    diverge consumer:
        for i in range(10):
            print("Consuming item", i)
            sleep(300)
    
    converge producer, consumer
    print("Processing complete")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Parallel Computation</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`main:
    // Launch multiple workers
    for i in range(4):
        diverge:
            worker_id = i
            print("Worker", worker_id, "starting")
            // Simulate work
            total = 0
            for j in range(1000000):
                total += j
            print("Worker", worker_id, "result:", total)
    
    // Wait for all workers
    converge
    print("All workers completed")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Features</SubSectionTitle>
        <List>
          <li><strong>Thread-safe:</strong> Proper mutex protection for goroutine management</li>
          <li><strong>Isolated Environment:</strong> Each goroutine gets its own enclosed environment</li>
          <li><strong>Automatic Cleanup:</strong> Resources are cleaned up when goroutines complete</li>
          <li><strong>Error Handling:</strong> Errors within goroutines are contained and reported</li>
          <li><strong>Sequential Support:</strong> Multiple sequential diverge/converge operations work reliably</li>
        </List>

        <InfoBox>
          <strong>Best Practices:</strong>
          <List>
            <li>Use named goroutines for complex logic</li>
            <li>Handle errors within goroutines</li>
            <li>Use converge strategically to wait for critical tasks</li>
            <li>Avoid long-running anonymous goroutines</li>
          </List>
        </InfoBox>
      </Section>

      {/* MAIN ENTRY POINT */}
      <Section id="main">
        <SectionTitle>Main Entry Point</SectionTitle>

        <InfoBox>
          <strong>New Feature:</strong> The <InlineCode>main:</InlineCode> keyword defines the 
          program's entry point, similar to Python's <InlineCode>if __name__ == "__main__":</InlineCode>
        </InfoBox>

        <SubSectionTitle>Syntax</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`main:
    // program entry point
    // indented code block`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>How It Works</SubSectionTitle>
        <p>When a <InlineCode>main:</InlineCode> block exists in your program:</p>
        <List>
          <li>Only definitions (functions, classes, assignments) run at the top level</li>
          <li>Executable statements are skipped at the top level</li>
          <li>Code inside <InlineCode>main:</InlineCode> executes as the program entry point</li>
        </List>

        <SubSectionTitle>Example</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Define functions and classes
spell helper(text):
    return f"Helper: {text}"

grim Calculator:
    spell add(a, b):
        return a + b

// This won't run because main: exists
// print("This is skipped")

// Main entry point
main:
    print("Program starts here")
    result = helper("test")
    print(result)
    
    calc = Calculator()
    sum = calc.add(10, 20)
    print("Sum:", sum)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>With Concurrency</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell process_data(id):
    print("Processing", id)
    sleep(1000)
    return id * 2

main:
    print("Starting concurrent tasks")
    
    diverge worker1:
        result = process_data(1)
        print("Worker 1 result:", result)
    
    diverge worker2:
        result = process_data(2)
        print("Worker 2 result:", result)
    
    converge worker1, worker2
    print("All processing complete")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <strong>When to Use:</strong>
          <List>
            <li>When you want clear separation between definitions and execution</li>
            <li>For larger programs with multiple functions and classes</li>
            <li>When using concurrency features</li>
            <li>To prevent code execution during module imports</li>
          </List>
        </InfoBox>
      </Section>

      {/* MODULES */}
      <Section id="modules">
        <SectionTitle>Modules and Imports</SectionTitle>

        <SubSectionTitle>Basic Import Syntax</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Grimoire-based imports
import "GrimoireName"
import "GrimoireName" as MyGrimoire

// Local file imports
import "filename"                 // Imports all definitions
import "mymodule.ClassName"       // Selective import
import "mymodule.spell_name"      // Import specific spell
import "mymodule.spell_name" as fn // Import with alias

// Package imports
import "package/module"
import "package/module.ClassName"
import "package/module.function_name"

// Relative imports
import "./filename"
import "../parent/module"
import "../../utils/helper"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Import Resolution</SubSectionTitle>
        <p>Carrion searches for modules in this order:</p>
        <List>
          <li>Current directory</li>
          <li>Project modules (<InlineCode>./carrion_modules/</InlineCode>)</li>
          <li>Global Bifrost modules (<InlineCode>/usr/bin/carrion_modules/</InlineCode>)</li>
          <li>User packages (<InlineCode>~/.carrion/packages/</InlineCode>)</li>
          <li>System packages (<InlineCode>/usr/local/share/carrion/lib/</InlineCode>)</li>
          <li>Standard library (Munin)</li>
        </List>

        <SubSectionTitle>Selective Imports</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Import specific grimoires
import "data_structures.Stack"
import "data_structures.Queue" as Q

// Import specific spells
import "math_helpers.add"
import "math_helpers.multiply" as mult

// Mixed imports
import "utilities.Logger"
import "utilities.format_date"
import "utilities.validate_email" as check_email

main:
    // Use imported items
    stack = Stack()
    queue = Q()
    sum = add(5, 3)
    product = mult(4, 7)
    logger = Logger("App")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Module Example</SubSectionTitle>
        <p><strong>File: math_utils.crl</strong></p>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell add(a, b):
    return a + b

spell multiply(a, b):
    return a * b

grim Calculator:
    init():
        self.result = 0
    
    spell calculate(a, b, op):
        if op == "+":
            return add(a, b)
        otherwise op == "*":
            return multiply(a, b)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <p><strong>File: main.crl</strong></p>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`import "math_utils.add"
import "math_utils.Calculator"

main:
    sum = add(10, 5)
    calc = Calculator()
    result = calc.calculate(4, 7, "*")
    print(sum, result)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      {/* BUILT-IN FUNCTIONS */}
      <Section id="builtins">
        <SectionTitle>Built-in Functions</SectionTitle>

        <SubSectionTitle>Type Conversion</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Description</th>
              <th>Example</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>int(value)</code></td><td>Convert to integer</td><td><code>int("42")</code> → <code>42</code></td></tr>
            <tr><td><code>float(value)</code></td><td>Convert to float</td><td><code>float("3.14")</code> → <code>3.14</code></td></tr>
            <tr><td><code>str(value)</code></td><td>Convert to string</td><td><code>str(42)</code> → <code>"42"</code></td></tr>
            <tr><td><code>bool(value)</code></td><td>Convert to boolean</td><td><code>bool(1)</code> → <code>True</code></td></tr>
            <tr><td><code>list(iterable)</code></td><td>Convert to array</td><td><code>list("abc")</code> → <code>["a","b","c"]</code></td></tr>
            <tr><td><code>tuple(iterable)</code></td><td>Convert to tuple</td><td><code>tuple([1,2,3])</code> → <code>(1,2,3)</code></td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Utility Functions</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Description</th>
              <th>Example</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>len(object)</code></td><td>Get length</td><td><code>len([1,2,3])</code> → <code>3</code></td></tr>
            <tr><td><code>type(object)</code></td><td>Get type</td><td><code>type(42)</code> → <code>"INTEGER"</code></td></tr>
            <tr><td><code>print(*args)</code></td><td>Print values</td><td><code>print("Hello", 123)</code></td></tr>
            <tr><td><code>input(prompt)</code></td><td>Get user input</td><td><code>input("Name: ")</code></td></tr>
            <tr><td><code>range(start, stop, step)</code></td><td>Generate sequence</td><td><code>range(5)</code> → <code>0..4</code></td></tr>
            <tr><td><code>max(*args)</code></td><td>Find maximum</td><td><code>max(1,5,3)</code> → <code>5</code></td></tr>
            <tr><td><code>abs(value)</code></td><td>Absolute value</td><td><code>abs(-5)</code> → <code>5</code></td></tr>
            <tr><td><code>ord(char)</code></td><td>Get ASCII code</td><td><code>ord("A")</code> → <code>65</code></td></tr>
            <tr><td><code>chr(code)</code></td><td>Get character</td><td><code>chr(65)</code> → <code>"A"</code></td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Collection Functions</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Description</th>
              <th>Example</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>enumerate(array)</code></td><td>Get indexed pairs</td><td><code>enumerate(["a","b"])</code> → <code>[(0,"a"),(1,"b")]</code></td></tr>
            <tr><td><code>pairs(hash, filter)</code></td><td>Get key-value pairs</td><td><code>pairs({'{'}&#34;a&#34;:1{'}'})</code></td></tr>
            <tr><td><code>is_sametype(obj1, obj2)</code></td><td>Compare types</td><td><code>is_sametype(1, 2)</code> → <code>True</code></td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Meta Functions</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>help()</code></td><td>Get help information</td></tr>
            <tr><td><code>version()</code></td><td>Show version information</td></tr>
            <tr><td><code>modules()</code></td><td>List available modules</td></tr>
            <tr><td><code>mimir</code></td><td>Launch interactive help (REPL only)</td></tr>
          </tbody>
        </Table>
      </Section>

      {/* STANDARD LIBRARY */}
      <Section id="stdlib">
        <SectionTitle>Standard Library (Munin)</SectionTitle>

        <InfoBox>
          <strong>Munin</strong> - The Carrion standard library, named after Odin's raven. 
          Provides essential functionality automatically loaded when Carrion starts.
        </InfoBox>

        <SubSectionTitle>Core Type Modules</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Module</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>Array</code></td><td>Array manipulation (append, sort, reverse, contains, slice)</td></tr>
            <tr><td><code>String</code></td><td>String operations (upper, lower, find, char_at, reverse)</td></tr>
            <tr><td><code>Integer</code></td><td>Integer utilities (to_bin, to_hex, is_prime, gcd, lcm)</td></tr>
            <tr><td><code>Float</code></td><td>Float operations (round, sqrt, sin, cos, abs)</td></tr>
            <tr><td><code>Boolean</code></td><td>Boolean logic (to_int, negate, and_with, or_with, xor_with)</td></tr>
            <tr><td><code>Primitive</code></td><td>Basic type operations</td></tr>
            <tr><td><code>Math</code></td><td>Mathematical functions and constants</td></tr>
            <tr><td><code>Time</code></td><td>Time and date operations</td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>I/O and System Modules</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Module</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>File</code></td><td>File I/O operations (read, write, append, exists)</td></tr>
            <tr><td><code>OS</code></td><td>Operating system interface (cwd, listdir, getenv, run, sleep)</td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Networking and Servers</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Module</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>ApiRequest</code></td><td>HTTP client (GET, POST, PUT, DELETE, HEAD, JSON, auth, retry)</td></tr>
            <tr><td><code>Servers</code></td><td>Server implementations (TCP, UDP, Unix, HTTP, WebServer)</td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Data Structures</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Module</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>DataStructures</code></td><td>Advanced structures (Stack, Queue, Heap, BTree with iterators)</td></tr>
            <tr><td><code>Iterable</code></td><td>Abstract base class for iterables</td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Utility Modules</SubSectionTitle>
        <Table>
          <thead>
            <tr>
              <th>Module</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>Debug</code></td><td>Debugging utilities</td></tr>
            <tr><td><code>BuiltinErrors</code></td><td>Error handling types</td></tr>
          </tbody>
        </Table>

        <SubSectionTitle>Usage Examples</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Array operations
arr = Array([3, 1, 4, 1, 5])
arr.append(9)
print(arr.sort())        // [1, 1, 3, 4, 5, 9]
print(arr.contains(3))   // True

// String manipulation
s = "Hello World"
print(s[0])              // "H"
sg = String(s)
print(sg.upper())        // "HELLO WORLD"

// Integer utilities
num = Integer(42)
print(num.to_bin())      // "0b101010"
print(num.is_prime())    // False

// Character conversion
print(ord("A"))          // 65
print(chr(65))           // "A"

// HTTP requests
api = ApiRequest()
response = api.get_json("https://api.example.com/data")
if response["success"]:
    print(response["data"])

// Data structures
stack = Stack()
stack.push(1)
stack.push(2)
print(stack.pop())       // 2`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Resources</SectionTitle>
        <List>
          <li>GitHub Repository: <a href="https://github.com/javanhut/TheCarrionLanguage" target="_blank" rel="noopener noreferrer">TheCarrionLanguage</a></li>
          <li>Documentation: <a href="https://github.com/javanhut/TheCarrionLanguage/tree/main/docs" target="_blank" rel="noopener noreferrer">Official Docs</a></li>
          <li>Standard Library Reference: <a href="https://github.com/javanhut/TheCarrionLanguage/blob/main/docs/Standard-Library.md" target="_blank" rel="noopener noreferrer">Munin Documentation</a></li>
          <li>Examples: <a href="https://github.com/javanhut/TheCarrionLanguage/tree/main/src/examples" target="_blank" rel="noopener noreferrer">Code Examples</a></li>
        </List>
      </Section>
    </Container>
  );
};

export default LanguageReference;
