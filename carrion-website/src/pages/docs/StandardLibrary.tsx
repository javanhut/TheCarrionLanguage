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
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
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
  margin-bottom: 1rem;
`;

const MethodList = styled.ul`
  color: ${({ theme }) => theme.colors.text.secondary};
  font-size: 0.95rem;
  line-height: 1.6;
  list-style: none;
  padding: 0;
`;

const MethodItem = styled.li`
  padding: 0.5rem 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  font-family: 'Monaco', 'Courier New', monospace;
  
  &:last-child {
    border-bottom: none;
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

const StandardLibrary: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Standard Library - Munin</Title>
        <Subtitle>
          The Munin standard library provides comprehensive functionality for arrays, strings, math, files, and system operations. Named after Odin's raven of memory.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Overview</SectionTitle>
        <Text>
          The Munin standard library is automatically loaded when Carrion starts, making all grimoires and functions 
          immediately available. It provides both primitive type enhancements through automatic wrapping and utility grimoires 
          for common tasks.
        </Text>

        <InfoBox>
          <InfoTitle>Automatic Primitive Wrapping</InfoTitle>
          <InfoText>
            Carrion automatically wraps primitive types (integers, floats, strings, booleans, arrays) with their corresponding 
            grimoire objects, allowing you to call methods directly on values: <InlineCode>42.to_bin()</InlineCode>, 
            <InlineCode>"hello".upper()</InlineCode>
          </InfoText>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>Core Library Functions</SectionTitle>
        <Text>
          These functions are available globally for system information and help:
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Get version information
version()  // "Carrion 0.1.6, Munin Standard Library 0.1.0"

// Interactive help system
help()     // Shows language help and available functions

// List all available modules
modules()  // Shows available standard library modules`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Primitive Type Grimoires</SectionTitle>
        
        <Grid>
          <Card>
            <CardTitle>Array Grimoire</CardTitle>
            <CardText>
              Enhanced array manipulation with sorting, searching, and transformation methods.
            </CardText>
            <MethodList>
              <MethodItem>length() - Get array size</MethodItem>
              <MethodItem>append(item) - Add element</MethodItem>
              <MethodItem>get(index) - Access with negative indexing</MethodItem>
              <MethodItem>set(index, value) - Modify element</MethodItem>
              <MethodItem>contains(value) - Check membership</MethodItem>
              <MethodItem>index_of(value) - Find index</MethodItem>
              <MethodItem>remove(value) - Remove element</MethodItem>
              <MethodItem>slice(start, end) - Extract range</MethodItem>
              <MethodItem>reverse() - Reverse copy</MethodItem>
              <MethodItem>sort() - Sorted copy</MethodItem>
            </MethodList>
          </Card>

          <Card>
            <CardTitle>String Grimoire</CardTitle>
            <CardText>
              Text manipulation with case conversion, searching, and transformation.
            </CardText>
            <MethodList>
              <MethodItem>length() - String length</MethodItem>
              <MethodItem>upper() - Convert to uppercase</MethodItem>
              <MethodItem>lower() - Convert to lowercase</MethodItem>
              <MethodItem>find(sub) - Find substring</MethodItem>
              <MethodItem>contains(sub) - Check substring</MethodItem>
              <MethodItem>char_at(index) - Get character</MethodItem>
              <MethodItem>reverse() - Reverse string</MethodItem>
              <MethodItem>to_string() - String representation</MethodItem>
            </MethodList>
          </Card>

          <Card>
            <CardTitle>Integer Grimoire</CardTitle>
            <CardText>
              Mathematical operations and number base conversions.
            </CardText>
            <MethodList>
              <MethodItem>to_bin() - Binary representation</MethodItem>
              <MethodItem>to_oct() - Octal representation</MethodItem>
              <MethodItem>to_hex() - Hexadecimal representation</MethodItem>
              <MethodItem>abs() - Absolute value</MethodItem>
              <MethodItem>pow(exp) - Power operation</MethodItem>
              <MethodItem>gcd(other) - Greatest common divisor</MethodItem>
              <MethodItem>lcm(other) - Least common multiple</MethodItem>
              <MethodItem>is_even() - Check if even</MethodItem>
              <MethodItem>is_odd() - Check if odd</MethodItem>
              <MethodItem>is_prime() - Prime number test</MethodItem>
            </MethodList>
          </Card>

          <Card>
            <CardTitle>Float Grimoire</CardTitle>
            <CardText>
              Floating-point operations with rounding, trigonometry, and conversions.
            </CardText>
            <MethodList>
              <MethodItem>round(places) - Round to decimals</MethodItem>
              <MethodItem>floor() - Round down</MethodItem>
              <MethodItem>ceil() - Round up</MethodItem>
              <MethodItem>abs() - Absolute value</MethodItem>
              <MethodItem>sqrt() - Square root</MethodItem>
              <MethodItem>pow(exp) - Power operation</MethodItem>
              <MethodItem>sin() - Sine (Taylor series)</MethodItem>
              <MethodItem>cos() - Cosine (Taylor series)</MethodItem>
              <MethodItem>is_integer() - Check if whole number</MethodItem>
              <MethodItem>to_int() - Convert to integer</MethodItem>
            </MethodList>
          </Card>

          <Card>
            <CardTitle>Boolean Grimoire</CardTitle>
            <CardText>
              Logical operations and boolean conversions.
            </CardText>
            <MethodList>
              <MethodItem>to_int() - Convert to 0 or 1</MethodItem>
              <MethodItem>to_string() - String representation</MethodItem>
              <MethodItem>negate() - Logical NOT</MethodItem>
              <MethodItem>and_with(other) - Logical AND</MethodItem>
              <MethodItem>or_with(other) - Logical OR</MethodItem>
              <MethodItem>xor_with(other) - Logical XOR</MethodItem>
              <MethodItem>implies(other) - Logical implication</MethodItem>
              <MethodItem>is_true() - Check if true</MethodItem>
              <MethodItem>is_false() - Check if false</MethodItem>
            </MethodList>
          </Card>
        </Grid>
      </Section>

      <Section>
        <SectionTitle>Array Operations Examples</SectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Create and manipulate arrays
numbers = [10, 5, 8, 3, 7]
arr = Array(numbers)

// Basic operations
print(arr.length())        // 5
arr.append(12)
print(arr.contains(8))     // True
print(arr.index_of(5))     // 1

// Advanced operations
sorted_arr = arr.sort()
reversed_arr = arr.reverse()
slice = arr.slice(1, 4)

// Works directly on array literals via automatic wrapping
print([3, 1, 4, 1, 5].sort())     // [1, 1, 3, 4, 5]
print([1, 2, 3].contains(2))      // True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>String Operations Examples</SectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// String manipulation
text = String("Hello, World!")

// Case conversion
print(text.upper())          // "HELLO, WORLD!"
print(text.lower())          // "hello, world!"

// Search operations
print(text.find("World"))    // 7
print(text.contains("Hello")) // True
print(text.char_at(0))       // "H"
print(text.char_at(-1))      // "!" (negative indexing)

// Transformation
print(text.reverse())        // "!dlroW ,olleH"

// Works directly on string literals
print("carrion".upper())     // "CARRION"
print("hello".length())      // 5`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Mathematical Operations</SectionTitle>
        
        <SubSectionTitle>Integer Operations</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`num = Integer(42)

// Number base conversions
print(num.to_bin())          // "0b101010"
print(num.to_oct())          // "0o52"
print(num.to_hex())          // "0x2a"

// Mathematical operations
print(num.abs())             // 42
print(num.pow(2))            // 1764
print(num.gcd(18))           // 6
print(num.lcm(18))           // 126

// Properties
print(num.is_even())         // True
print(num.is_odd())          // False
print(num.is_prime())        // False

// Works directly on integers
print(17.is_prime())         // True
print(10.to_bin())           // "0b1010"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Float Operations</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`f = Float(3.14159)

// Rounding operations
print(f.round(2))            // 3.14
print(f.floor())             // 3
print(f.ceil())              // 4

// Mathematical operations
print(f.abs())               // 3.14159
print(f.sqrt())              // Square root using Newton's method
print(f.pow(2))              // 9.8696...

// Trigonometry
print(f.sin())               // Sine using Taylor series
print(f.cos())               // Cosine using Taylor series

// Type checking
print(f.is_integer())        // False
print(f.is_positive())       // True

// Works directly on floats
print(3.14.round(1))         // 3.1`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>File Grimoire</SectionTitle>
        <Text>
          The File grimoire provides comprehensive file system operations for reading, writing, and managing files.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`file = File()

// Read entire file
content = file.read("input.txt")
print(content)

// Write to file (overwrites)
file.write("output.txt", "Hello World")

// Append to file
file.append("log.txt", "New log entry\\n")

// Check file existence
if file.exists("config.txt"):
    config = file.read("config.txt")
    print("Config loaded")
else:
    print("Config file not found")

// Error handling with files
attempt:
    data = file.read("important.txt")
    process(data)
ensnare (FileNotFoundError):
    print("File not found!")
ensnare:
    print("Error reading file!")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>OS Grimoire</SectionTitle>
        <Text>
          The OS grimoire provides operating system interface for directory operations, environment variables, and process management.
        </Text>

        <SubSectionTitle>Directory Operations</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`os = OS()

// Directory navigation
current = os.cwd()
print(f"Current directory: {current}")

os.chdir("/path/to/directory")

// List directory contents
files = os.listdir(".")
for filename in files:
    print(f"Found: {filename}")

// Create directory
os.mkdir("new_folder", 0755)

// Remove file or directory
os.remove("old_file.txt")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Environment Variables</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`os = OS()

// Get environment variable
home = os.getenv("HOME")
print(f"Home directory: {home}")

// Set environment variable
os.setenv("MY_VAR", "my_value")

// Expand environment variables
path = os.expandEnv("$HOME/documents")
print(path)  // Expands to actual home path`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Process Management</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`os = OS()

// Run external commands
os.run("ls", ["-la"], False)
output = os.run("echo", ["Hello"], True)  // Capture output

// Sleep/delay
print("Waiting...")
os.sleep(2)  // Sleep for 2 seconds
print("Done!")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Complete Usage Examples</SectionTitle>
        
        <SubSectionTitle>File Processing Pipeline</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Read, process, and write file with error handling
file = File()

attempt:
    // Check if input exists
    if not file.exists("input.txt"):
        raise Error("FileNotFound", "Input file missing")
    
    // Read and process
    content = file.read("input.txt")
    lines = content.split("\\n")
    
    // Process each line
    processed = []
    for line in lines:
        if line.length() > 0:
            processed.append(line.upper())
    
    // Write result
    result = "\\n".join(processed)
    file.write("output.txt", result)
    
    print("Processing complete!")

ensnare (Error):
    print(f"Error: {error.message}")
ensnare:
    print("Unexpected error occurred")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>System Information Gatherer</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`os = OS()
file = File()

// Gather system information
info = []
info.append(f"Current Directory: {os.cwd()}")
info.append(f"Home: {os.getenv('HOME')}")
info.append(f"User: {os.getenv('USER')}")

// List files
info.append("\\nFiles in current directory:")
files = os.listdir(".")
for filename in files:
    info.append(f"  - {filename}")

// Save to file
report = "\\n".join(info)
file.write("system_info.txt", report)
print("System information saved!")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Data Analysis with Arrays</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Analyze numeric data
data = [45, 23, 67, 89, 12, 34, 56, 78, 90, 11]
arr = Array(data)

// Statistics
sorted_data = arr.sort()
print(f"Minimum: {sorted_data.first()}")
print(f"Maximum: {sorted_data.last()}")

// Calculate average
total = 0
for num in data:
    total += num
average = total / arr.length()
print(f"Average: {average}")

// Find specific values
target = 67
if arr.contains(target):
    index = arr.index_of(target)
    print(f"Found {target} at index {index}")

// Filter above average
above_avg = []
for num in data:
    if num > average:
        above_avg.append(num)
print(f"Above average: {above_avg}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Best Practices</SectionTitle>
        
        <InfoBox>
          <InfoTitle>Always Handle Errors</InfoTitle>
          <InfoText>
            When working with files and system operations, always use <InlineCode>attempt-ensnare</InlineCode> blocks 
            to handle potential errors gracefully.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Check Before Operations</InfoTitle>
          <InfoText>
            Use <InlineCode>file.exists()</InlineCode> to check for file existence before reading, and validate paths 
            before directory operations to prevent runtime errors.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Use Method Chaining</InfoTitle>
          <InfoText>
            Take advantage of automatic primitive wrapping to write concise code: 
            <InlineCode>"hello world".upper().reverse()</InlineCode>
          </InfoText>
        </InfoBox>
      </Section>
    </Container>
  );
};

export default StandardLibrary;
