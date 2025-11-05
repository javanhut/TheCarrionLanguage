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
    from {
      opacity: 0;
      transform: translateY(-20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
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
  &:nth-child(5) { animation-delay: 0.4s; }
  &:nth-child(6) { animation-delay: 0.5s; }

  @keyframes fadeInUp {
    from {
      opacity: 0;
      transform: translateY(30px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
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

const StepList = styled.ol`
  color: ${({ theme }) => theme.colors.text.secondary};
  font-size: 1.1rem;
  line-height: 1.8;
  margin: 1.5rem 0 1.5rem 2rem;
`;

const StepItem = styled.li`
  margin-bottom: 1rem;
`;

const InlineCode = styled.code`
  background: rgba(6, 182, 212, 0.1);
  color: ${({ theme }) => theme.colors.primary};
  padding: 0.2rem 0.5rem;
  border-radius: 4px;
  font-family: 'Monaco', 'Courier New', monospace;
  font-size: 0.95em;
`;

const QuickStart: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Quick Start Guide</Title>
        <Subtitle>
          Get up and running with Carrion in minutes. Learn the fundamentals and write your first spells.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Installation</SectionTitle>
        <Text>
          First, install Carrion on your system. Choose the installation method that works best for you:
        </Text>
        
        <SubSectionTitle>Linux/macOS Quick Install</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`curl -sSL https://raw.githubusercontent.com/YourRepo/carrion/main/install/install.sh | bash`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Build from Source</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`git clone https://github.com/YourRepo/carrion.git
cd carrion
make install`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Verify Installation</InfoTitle>
          <InfoText>
            After installation, verify that Carrion is working by running: <InlineCode>carrion --version</InlineCode>
          </InfoText>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>Your First Program</SectionTitle>
        <Text>
          Let's write your first Carrion program. Create a file called <InlineCode>hello.crl</InlineCode>:
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// hello.crl - Your first Carrion program
print("Hello, World!")
print("Welcome to Carrion!")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <Text>Run your program:</Text>
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`carrion hello.crl`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Interactive REPL</InfoTitle>
          <InfoText>
            You can also start an interactive session by running <InlineCode>carrion</InlineCode> without arguments. 
            This is great for experimenting!
          </InfoText>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>Language Basics</SectionTitle>
        
        <SubSectionTitle>Variables and Data Types</SubSectionTitle>
        <Text>Carrion has dynamic typing with several built-in types:</Text>
        
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
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
person = {"name": "Alice", "age": 30, "city": "NYC"}

// None
result = None`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Basic Operations</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Arithmetic
sum = 10 + 5        // 15
difference = 10 - 5 // 5
product = 10 * 5    // 50
quotient = 10 / 5   // 2.0
power = 2 ** 3      // 8

// String concatenation
full_name = "John" + " " + "Doe"

// Array operations
numbers.append(6)
first = numbers[0]
length = len(numbers)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Functions (Spells)</SectionTitle>
        <Text>
          In Carrion, functions are called "spells" and are defined using the <InlineCode>spell</InlineCode> keyword:
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Define a simple spell
spell greet(name):
    return f"Hello, {name}!"

// Call the spell
message = greet("Alice")
print(message)  // Output: Hello, Alice!

// Spell with default parameters
spell introduce(name, age = 25):
    return f"My name is {name} and I'm {age} years old"

print(introduce("Bob"))        // My name is Bob and I'm 25 years old
print(introduce("Alice", 30))  // My name is Alice and I'm 30 years old

// Spell with multiple return values
spell get_coordinates():
    return 10, 20

x, y = get_coordinates()
print(f"X: {x}, Y: {y}")  // X: 10, Y: 20`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Control Flow</SectionTitle>
        
        <SubSectionTitle>Conditionals</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// If statement
age = 18
if age >= 18:
    print("You are an adult")

// If-else
score = 85
if score >= 90:
    grade = "A"
else:
    grade = "B"

// If-otherwise-else (like elif)
temperature = 75
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

        <SubSectionTitle>Loops</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// For loop
numbers = [1, 2, 3, 4, 5]
for num in numbers:
    print(num)

// Range loop
for i in range(5):
    print(f"Count: {i}")

// While loop
count = 0
while count < 5:
    print(f"Count: {count}")
    count += 1

// Loop control
for i in range(10):
    if i == 3:
        skip  // Continue to next iteration
    if i == 7:
        stop  // Break from loop
    print(i)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Object-Oriented Programming</SectionTitle>
        <Text>
          Create classes (called "grimoires") to organize your code:
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

// Create an instance
person = Person("Alice", 30)
print(person.greet())      // Hello, I'm Alice
print(person.birthday())   // Happy birthday! Now 31 years old`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Error Handling</SectionTitle>
        <Text>
          Handle errors gracefully using the <InlineCode>attempt-ensnare</InlineCode> pattern:
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`attempt:
    number = int(input("Enter a number: "))
    result = 100 / number
    print(f"Result: {result}")
ensnare:
    print("Error: Invalid input or division by zero!")

// Specific error handling
attempt:
    file = File()
    content = file.read("data.txt")
    print(content)
ensnare (FileNotFoundError):
    print("File not found!")
ensnare:
    print("An unexpected error occurred!")
resolve:
    print("Cleanup completed")  // Always runs`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Using the Standard Library</SectionTitle>
        <Text>
          Carrion comes with a powerful standard library called Munin:
        </Text>

        <Grid>
          <Card>
            <CardTitle>String Operations</CardTitle>
            <CardText>
              Transform and manipulate text with the String grimoire.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`text = "hello"
print(text.upper())    // HELLO
print(text.length())   // 5`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>

          <Card>
            <CardTitle>Array Methods</CardTitle>
            <CardText>
              Work with collections using enhanced array functionality.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`arr = [3, 1, 4, 1, 5]
print(arr.sort())      // [1, 1, 3, 4, 5]
print(arr.contains(4)) // True`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>

          <Card>
            <CardTitle>File Operations</CardTitle>
            <CardText>
              Read and write files with the File grimoire.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`file = File()
content = file.read("data.txt")
file.write("out.txt", content)`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>

          <Card>
            <CardTitle>Math Functions</CardTitle>
            <CardText>
              Perform calculations with Integer and Float grimoires.
            </CardText>
            <CodeBlock>
              <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1rem', fontSize: '0.9rem' }}>
{`num = 42
print(num.to_bin())    // 0b101010
print(num.is_prime())  // False`}
              </SyntaxHighlighter>
            </CodeBlock>
          </Card>
        </Grid>
      </Section>

      <Section>
        <SectionTitle>Next Steps</SectionTitle>
        <Text>
          Now that you understand the basics, explore these resources to deepen your knowledge:
        </Text>

        <StepList>
          <StepItem>
            <strong>Language Reference:</strong> Comprehensive guide to Carrion's syntax and features
          </StepItem>
          <StepItem>
            <strong>Standard Library:</strong> Explore the full Munin standard library documentation
          </StepItem>
          <StepItem>
            <strong>REPL Guide:</strong> Master the interactive REPL for rapid prototyping
          </StepItem>
          <StepItem>
            <strong>Grimoires (OOP):</strong> Deep dive into object-oriented programming with Carrion
          </StepItem>
          <StepItem>
            <strong>Error Handling:</strong> Learn advanced error handling patterns
          </StepItem>
        </StepList>

        <InfoBox>
          <InfoTitle>Join the Community</InfoTitle>
          <InfoText>
            Have questions? Join our community to connect with other Carrion developers, share your projects, 
            and get help when you need it.
          </InfoText>
        </InfoBox>
      </Section>
    </Container>
  );
};

export default QuickStart;
