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
  InlineCode,
} from '../../components/docs';

const sections = [
  { id: 'overview', title: 'Overview' },
  { id: 'array', title: 'Array Grimoire' },
  { id: 'string', title: 'String Grimoire' },
  { id: 'integer', title: 'Integer Grimoire' },
  { id: 'float', title: 'Float Grimoire' },
  { id: 'boolean', title: 'Boolean Grimoire' },
  { id: 'file', title: 'File Grimoire' },
  { id: 'os', title: 'OS Grimoire' },
];

const StandardLibrary: React.FC = () => {
  return (
    <DocLayout
      title="Standard Library - Munin"
      description="The Munin standard library provides comprehensive functionality for arrays, strings, math, files, and system operations."
      sections={sections}
    >
      <Section id="overview">
        <SectionTitle>Overview</SectionTitle>
        <Lead>
          The Munin standard library is named after Odin's raven of memory. It is automatically loaded when
          Carrion starts, making all grimoires and functions immediately available.
        </Lead>

        <InfoBox>
          <InfoTitle>Automatic Primitive Wrapping</InfoTitle>
          <InfoText>
            Carrion automatically wraps primitive types with their corresponding grimoire objects.
            This allows calling methods directly on values: <InlineCode>42.to_bin()</InlineCode>,
            <InlineCode>"hello".upper()</InlineCode>
          </InfoText>
        </InfoBox>

        <SubSection>
          <SubSectionTitle>Global Functions</SubSectionTitle>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Function</TableHead>
                <TableHead>Description</TableHead>
              </TableRow>
            </TableHeader>
            <tbody>
              <TableRow>
                <TableCell><InlineCode>version()</InlineCode></TableCell>
                <TableCell>Get Carrion and Munin version info</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>help()</InlineCode></TableCell>
                <TableCell>Show language help and available functions</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>modules()</InlineCode></TableCell>
                <TableCell>List all available standard library modules</TableCell>
              </TableRow>
            </tbody>
          </Table>
        </SubSection>
      </Section>

      <Section id="array">
        <SectionTitle>Array Grimoire</SectionTitle>
        <Paragraph>
          The Array grimoire provides enhanced array manipulation with sorting, searching, and
          transformation methods.
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Method</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>length()</InlineCode></TableCell>
              <TableCell>Get array size</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>append(item)</InlineCode></TableCell>
              <TableCell>Add element to end</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>get(index)</InlineCode></TableCell>
              <TableCell>Access element (supports negative indexing)</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>set(index, value)</InlineCode></TableCell>
              <TableCell>Modify element at index</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>contains(value)</InlineCode></TableCell>
              <TableCell>Check if array contains value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>index_of(value)</InlineCode></TableCell>
              <TableCell>Find index of value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>remove(value)</InlineCode></TableCell>
              <TableCell>Remove first occurrence of value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>slice(start, end)</InlineCode></TableCell>
              <TableCell>Extract a range of elements</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>reverse()</InlineCode></TableCell>
              <TableCell>Return reversed copy</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>sort()</InlineCode></TableCell>
              <TableCell>Return sorted copy</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Create and manipulate arrays
numbers = [10, 5, 8, 3, 7]

// Basic operations
print(numbers.length())        // 5
numbers.append(12)
print(numbers.contains(8))     // True
print(numbers.index_of(5))     // 1

// Advanced operations
sorted_arr = numbers.sort()    // [3, 5, 7, 8, 10, 12]
reversed_arr = numbers.reverse()
slice = numbers.slice(1, 4)    // [5, 8, 3]

// Works directly on array literals
print([3, 1, 4, 1, 5].sort())  // [1, 1, 3, 4, 5]`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="string">
        <SectionTitle>String Grimoire</SectionTitle>
        <Paragraph>
          Text manipulation with case conversion, searching, and transformation methods.
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Method</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>length()</InlineCode></TableCell>
              <TableCell>Get string length</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>upper()</InlineCode></TableCell>
              <TableCell>Convert to uppercase</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>lower()</InlineCode></TableCell>
              <TableCell>Convert to lowercase</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>find(sub)</InlineCode></TableCell>
              <TableCell>Find substring position</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>contains(sub)</InlineCode></TableCell>
              <TableCell>Check if contains substring</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>char_at(index)</InlineCode></TableCell>
              <TableCell>Get character at index</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>reverse()</InlineCode></TableCell>
              <TableCell>Reverse string</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// String manipulation
text = "Hello, World!"

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
print("carrion".upper())     // "CARRION"`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="integer">
        <SectionTitle>Integer Grimoire</SectionTitle>
        <Paragraph>
          Mathematical operations and number base conversions for integers.
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Method</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>to_bin()</InlineCode></TableCell>
              <TableCell>Binary representation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>to_oct()</InlineCode></TableCell>
              <TableCell>Octal representation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>to_hex()</InlineCode></TableCell>
              <TableCell>Hexadecimal representation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>abs()</InlineCode></TableCell>
              <TableCell>Absolute value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>pow(exp)</InlineCode></TableCell>
              <TableCell>Power operation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>gcd(other)</InlineCode></TableCell>
              <TableCell>Greatest common divisor</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>lcm(other)</InlineCode></TableCell>
              <TableCell>Least common multiple</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>is_even()</InlineCode></TableCell>
              <TableCell>Check if even</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>is_odd()</InlineCode></TableCell>
              <TableCell>Check if odd</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>is_prime()</InlineCode></TableCell>
              <TableCell>Prime number test</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`num = 42

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
print(17.is_prime())         // True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="float">
        <SectionTitle>Float Grimoire</SectionTitle>
        <Paragraph>
          Floating-point operations with rounding, trigonometry, and conversions.
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Method</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>round(places)</InlineCode></TableCell>
              <TableCell>Round to decimal places</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>floor()</InlineCode></TableCell>
              <TableCell>Round down</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>ceil()</InlineCode></TableCell>
              <TableCell>Round up</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>abs()</InlineCode></TableCell>
              <TableCell>Absolute value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>sqrt()</InlineCode></TableCell>
              <TableCell>Square root</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>pow(exp)</InlineCode></TableCell>
              <TableCell>Power operation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>sin()</InlineCode></TableCell>
              <TableCell>Sine (Taylor series)</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>cos()</InlineCode></TableCell>
              <TableCell>Cosine (Taylor series)</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>is_integer()</InlineCode></TableCell>
              <TableCell>Check if whole number</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>to_int()</InlineCode></TableCell>
              <TableCell>Convert to integer</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`f = 3.14159

// Rounding operations
print(f.round(2))            // 3.14
print(f.floor())             // 3
print(f.ceil())              // 4

// Mathematical operations
print(f.abs())               // 3.14159
print(f.sqrt())              // Square root
print(f.pow(2))              // 9.8696...

// Trigonometry
print(f.sin())               // Sine
print(f.cos())               // Cosine

// Type checking
print(f.is_integer())        // False
print(3.14.round(1))         // 3.1`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="boolean">
        <SectionTitle>Boolean Grimoire</SectionTitle>
        <Paragraph>
          Logical operations and boolean conversions.
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Method</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>to_int()</InlineCode></TableCell>
              <TableCell>Convert to 0 or 1</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>to_string()</InlineCode></TableCell>
              <TableCell>String representation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>negate()</InlineCode></TableCell>
              <TableCell>Logical NOT</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>and_with(other)</InlineCode></TableCell>
              <TableCell>Logical AND</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>or_with(other)</InlineCode></TableCell>
              <TableCell>Logical OR</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>xor_with(other)</InlineCode></TableCell>
              <TableCell>Logical XOR</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>implies(other)</InlineCode></TableCell>
              <TableCell>Logical implication</TableCell>
            </TableRow>
          </tbody>
        </Table>
      </Section>

      <Section id="file">
        <SectionTitle>File Grimoire</SectionTitle>
        <Paragraph>
          Comprehensive file system operations for reading, writing, and managing files.
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Method</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>read(path)</InlineCode></TableCell>
              <TableCell>Read entire file contents</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>write(path, content)</InlineCode></TableCell>
              <TableCell>Write content to file (overwrites)</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>append(path, content)</InlineCode></TableCell>
              <TableCell>Append content to file</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>exists(path)</InlineCode></TableCell>
              <TableCell>Check if file exists</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`file = File()

// Read entire file
content = file.read("input.txt")

// Write to file (overwrites)
file.write("output.txt", "Hello World")

// Append to file
file.append("log.txt", "New log entry\\n")

// Check file existence
if file.exists("config.txt"):
    config = file.read("config.txt")
    print("Config loaded")

// Error handling with files
attempt:
    data = file.read("important.txt")
ensnare (FileNotFoundError):
    print("File not found!")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="os">
        <SectionTitle>OS Grimoire</SectionTitle>
        <Paragraph>
          Operating system interface for directory operations, environment variables, and process management.
        </Paragraph>

        <SubSection>
          <SubSectionTitle>Directory Operations</SubSectionTitle>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Method</TableHead>
                <TableHead>Description</TableHead>
              </TableRow>
            </TableHeader>
            <tbody>
              <TableRow>
                <TableCell><InlineCode>cwd()</InlineCode></TableCell>
                <TableCell>Get current working directory</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>chdir(path)</InlineCode></TableCell>
                <TableCell>Change directory</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>listdir(path)</InlineCode></TableCell>
                <TableCell>List directory contents</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>mkdir(path, mode)</InlineCode></TableCell>
                <TableCell>Create directory</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>remove(path)</InlineCode></TableCell>
                <TableCell>Remove file or directory</TableCell>
              </TableRow>
            </tbody>
          </Table>

          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`os = OS()

// Directory navigation
current = os.cwd()
print(f"Current directory: {current}")

// List directory contents
files = os.listdir(".")
for filename in files:
    print(f"Found: {filename}")

// Create directory
os.mkdir("new_folder", 0755)`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Environment Variables</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`os = OS()

// Get environment variable
home = os.getenv("HOME")
print(f"Home directory: {home}")

// Set environment variable
os.setenv("MY_VAR", "my_value")

// Expand environment variables
path = os.expandEnv("$HOME/documents")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Process Management</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
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
        </SubSection>

        <InfoBox>
          <InfoTitle>Best Practices</InfoTitle>
          <InfoText>
            Always use <InlineCode>attempt-ensnare</InlineCode> blocks when working with files
            and system operations. Check file existence with <InlineCode>file.exists()</InlineCode>
            before reading.
          </InfoText>
        </InfoBox>
      </Section>
    </DocLayout>
  );
};

export default StandardLibrary;
