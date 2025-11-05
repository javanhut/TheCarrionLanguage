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

const Grid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin: 2rem 0;
`;

const Card = styled.div`
  background: ${({ theme }) => theme.colors.surface};
  padding: 2rem;
  border-radius: ${({ theme }) => theme.borderRadius.large};
  box-shadow: ${({ theme }) => theme.shadows.medium};
  transition: all ${({ theme }) => theme.transitions.standard};
  border: 1px solid rgba(6, 182, 212, 0.1);
  &:hover {
    transform: translateY(-5px);
    box-shadow: ${({ theme }) => theme.shadows.large};
    border-color: ${({ theme }) => theme.colors.primary};
  }
`;

const CardTitle = styled.h4`
  color: ${({ theme }) => theme.colors.primary};
  font-size: 1.4rem;
  margin-bottom: 1rem;
  font-weight: 600;
`;

const CardText = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  font-size: 1rem;
  line-height: 1.6;
`;

const FunctionTable = styled.table`
  width: 100%;
  border-collapse: collapse;
  margin: 2rem 0;
  background: ${({ theme }) => theme.colors.surface};
  border-radius: ${({ theme }) => theme.borderRadius.medium};
  overflow: hidden;
  box-shadow: ${({ theme }) => theme.shadows.medium};
`;

const TableHeader = styled.th`
  background: ${({ theme }) => theme.gradients.primary};
  color: ${({ theme }) => theme.colors.background};
  padding: 1rem;
  text-align: left;
  font-weight: 600;
`;

const TableRow = styled.tr`
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  &:hover {
    background: rgba(6, 182, 212, 0.05);
  }
`;

const TableCell = styled.td<{ code?: boolean }>`
  padding: 1rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  font-family: ${props => props.code ? "'Monaco', 'Courier New', monospace" : 'inherit'};
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

const BuiltinFunctions: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Builtin Functions</Title>
        <Subtitle>
          Comprehensive reference for Carrion's built-in functions. These functions are globally available without imports.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Type Conversion Functions</SectionTitle>
        <Text>
          Convert values between different types with these essential functions.
        </Text>

        <FunctionTable>
          <thead>
            <tr>
              <TableHeader>Function</TableHeader>
              <TableHeader>Description</TableHeader>
              <TableHeader>Example</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell code>int(value)</TableCell>
              <TableCell>Convert to integer</TableCell>
              <TableCell code>int("42") → 42</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>float(value)</TableCell>
              <TableCell>Convert to float</TableCell>
              <TableCell code>float("3.14") → 3.14</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>str(value)</TableCell>
              <TableCell>Convert to string</TableCell>
              <TableCell code>str(42) → "42"</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>bool(value)</TableCell>
              <TableCell>Convert to boolean</TableCell>
              <TableCell code>bool(1) → True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>list(iterable)</TableCell>
              <TableCell>Convert to list/array</TableCell>
              <TableCell code>list("hi") → ["h","i"]</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>tuple(iterable)</TableCell>
              <TableCell>Convert to tuple</TableCell>
              <TableCell code>tuple([1,2]) → (1,2)</TableCell>
            </TableRow>
          </tbody>
        </FunctionTable>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Type conversion examples
num_str = "42"
num = int(num_str)          // 42
decimal = float("3.14")     // 3.14

// Boolean conversions
bool(1)                     // True
bool(0)                     // False
bool("")                    // False
bool("text")                // True

// Collection conversions
chars = list("hello")       // ["h", "e", "l", "l", "o"]
coords = tuple([10, 20])    // (10, 20)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>I/O Functions</SectionTitle>
        
        <SubSectionTitle>print(*args)</SubSectionTitle>
        <Text>
          Output values to console with automatic spacing between arguments.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Basic printing
print("Hello, World!")              // Hello, World!
print("The answer is", 42)          // The answer is 42
print(1, 2, 3, 4, 5)               // 1 2 3 4 5

// Printing expressions
x = 10
print("Value:", x, "Squared:", x ** 2)  // Value: 10 Squared: 100`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>input(prompt="")</SubSectionTitle>
        <Text>
          Read user input from the console with an optional prompt message.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Basic input
name = input("Enter your name: ")
print(f"Hello, {name}!")

// Input with type conversion
age_str = input("Enter your age: ")
age = int(age_str)
print(f"You are {age} years old")

// Interactive program
choice = input("Choose (1-3): ")
if choice == "1":
    print("Option 1 selected")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Utility Functions</SectionTitle>
        
        <Grid>
          <Card>
            <CardTitle>len(object)</CardTitle>
            <CardText>
              Returns the length of strings, arrays, hashes, or tuples.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`len("hello")           // 5
len([1, 2, 3])         // 3
len({"a": 1, "b": 2})  // 2`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>

          <Card>
            <CardTitle>type(object)</CardTitle>
            <CardText>
              Returns the type of an object as a string.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`type(42)          // "INTEGER"
type(3.14)        // "FLOAT"
type("hello")     // "STRING"`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>

          <Card>
            <CardTitle>is_sametype(a, b)</CardTitle>
            <CardText>
              Check if two objects have the same type.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`is_sametype(42, 17)      // True
is_sametype(42, 3.14)    // False
is_sametype("a", "b")    // True`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>
        </Grid>
      </Section>

      <Section>
        <SectionTitle>Mathematical Functions</SectionTitle>
        
        <SubSectionTitle>range(start, stop, step)</SubSectionTitle>
        <Text>
          Generate sequences of numbers. Can be called with 1, 2, or 3 arguments.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Single argument (stop)
range(5)              // [0, 1, 2, 3, 4]

// Two arguments (start, stop)
range(2, 8)           // [2, 3, 4, 5, 6, 7]

// Three arguments (start, stop, step)
range(0, 10, 2)       // [0, 2, 4, 6, 8]
range(10, 0, -1)      // [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

// Use in loops
for i in range(5):
    print(i)          // Prints 0 through 4`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>max(*args) and abs(value)</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Maximum value
max(1, 5, 3, 2)           // 5
max([10, 20, 15])         // 20
max("apple", "zoo")       // "zoo" (lexicographic)

// Absolute value
abs(-42)                  // 42
abs(3.14)                 // 3.14
abs(-2.5)                 // 2.5`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Character Functions</SectionTitle>
        
        <Grid>
          <Card>
            <CardTitle>ord(char)</CardTitle>
            <CardText>
              Returns the ASCII/Unicode code point of a character.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`ord("A")    // 65
ord("a")    // 97
ord("0")    // 48`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>

          <Card>
            <CardTitle>chr(code)</CardTitle>
            <CardText>
              Returns the character for an ASCII/Unicode code point.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`chr(65)     // "A"
chr(97)     // "a"
chr(48)     // "0"`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>
        </Grid>

        <SubSectionTitle>Character Processing Example</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Build string from ASCII codes
codes = [72, 101, 108, 108, 111]
result = ""
for code in codes:
    result += chr(code)
print(result)  // "Hello"

// Get ASCII codes from string
text = "ABC"
for i in range(len(text)):
    char = text[i]
    print(f"{char} = {ord(char)}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Collection Functions</SectionTitle>
        
        <SubSectionTitle>enumerate(array)</SubSectionTitle>
        <Text>
          Returns an array of (index, value) tuples for iteration.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`items = ["apple", "banana", "cherry"]

for index, value in enumerate(items):
    print(f"{index}: {value}")

// Output:
// 0: apple
// 1: banana
// 2: cherry`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>pairs(hash, filter="")</SubSectionTitle>
        <Text>
          Returns key-value pairs from a hash, optionally filtered by prefix.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`data = {"name": "John", "age": 30, "city": "NYC"}

// Iterate over all pairs
for key, value in pairs(data):
    print(f"{key}: {value}")

// Output:
// name: John
// age: 30
// city: NYC`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>System Functions</SectionTitle>
        
        <Text>
          These functions interface with the operating system and file system. Use the OS and File grimoires for more comprehensive operations.
        </Text>

        <FunctionTable>
          <thead>
            <tr>
              <TableHeader>Category</TableHeader>
              <TableHeader>Functions</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell>OS Operations</TableCell>
              <TableCell code>osGetEnv, osSetEnv, osGetCwd, osChdir, osSleep, osListDir, osRemove, osMkdir</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>File Operations</TableCell>
              <TableCell code>fileRead, fileWrite, fileAppend, fileExists</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Process Control</TableCell>
              <TableCell code>osRunCommand, osExpandEnv</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Error Handling</TableCell>
              <TableCell code>Error(name, message)</TableCell>
            </TableRow>
          </tbody>
        </FunctionTable>

        <InfoBox>
          <InfoTitle>Prefer Grimoires</InfoTitle>
          <InfoText>
            For most system operations, use the <InlineCode>OS()</InlineCode> and <InlineCode>File()</InlineCode> 
            grimoires instead of the raw functions. They provide better error handling and a cleaner interface.
          </InfoText>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>Complete Examples</SectionTitle>
        
        <SubSectionTitle>Interactive Calculator</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell calculator():
    while True:
        print("\\nSimple Calculator")
        print("1. Add")
        print("2. Subtract")
        print("3. Multiply")
        print("4. Divide")
        print("5. Exit")
        
        choice = input("Choose operation (1-5): ")
        
        if choice == "5":
            print("Goodbye!")
            stop
        
        if choice not in ["1", "2", "3", "4"]:
            print("Invalid choice!")
            skip
        
        a = float(input("Enter first number: "))
        b = float(input("Enter second number: "))
        
        if choice == "1":
            print(f"Result: {a + b}")
        otherwise choice == "2":
            print(f"Result: {a - b}")
        otherwise choice == "3":
            print(f"Result: {a * b}")
        otherwise choice == "4":
            if b != 0:
                print(f"Result: {a / b}")
            else:
                print("Error: Division by zero!")

calculator()`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Data Processing</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Process and analyze data
data = [45, 23, 67, 89, 12, 34, 56, 78, 90, 11]

// Calculate statistics
total = 0
for num in data:
    total += num

average = total / len(data)
maximum = max(data)

print(f"Count: {len(data)}")
print(f"Total: {total}")
print(f"Average: {average}")
print(f"Maximum: {maximum}")

// Count values above average
above_avg = 0
for num in data:
    if num > average:
        above_avg += 1

print(f"Above average: {above_avg}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Best Practices</SectionTitle>
        
        <InfoBox>
          <InfoTitle>Type Checking Before Conversion</InfoTitle>
          <InfoText>
            Always validate input before type conversion to prevent errors. Use <InlineCode>attempt-ensnare</InlineCode> 
            blocks when converting user input.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Use Appropriate Functions</InfoTitle>
          <InfoText>
            Choose the right function for the task. For example, use <InlineCode>len()</InlineCode> instead of manually 
            counting elements, and <InlineCode>max()</InlineCode> instead of implementing your own maximum finder.
          </InfoText>
        </InfoBox>
      </Section>
    </Container>
  );
};

export default BuiltinFunctions;
