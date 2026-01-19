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
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableCell,
  InlineCode,
} from '../../components/docs';

const sections = [
  { id: 'io', title: 'Input/Output' },
  { id: 'type-conversion', title: 'Type Conversion' },
  { id: 'collections', title: 'Collections' },
  { id: 'math', title: 'Math' },
  { id: 'utility', title: 'Utility' },
];

const BuiltinFunctions: React.FC = () => {
  return (
    <DocLayout
      title="Built-in Functions"
      description="Reference for all globally available built-in functions in Carrion."
      sections={sections}
    >
      <Section id="io">
        <SectionTitle>Input/Output</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Function</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>print(value)</InlineCode></TableCell>
              <TableCell>Print value to stdout with newline</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>printn(value)</InlineCode></TableCell>
              <TableCell>Print value without newline</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>input(prompt)</InlineCode></TableCell>
              <TableCell>Read user input with optional prompt</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`print("Hello, World!")
printn("Enter name: ")
name = input()

// Or combined
name = input("Enter name: ")
print(f"Hello, {name}!")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="type-conversion">
        <SectionTitle>Type Conversion</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Function</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>int(value)</InlineCode></TableCell>
              <TableCell>Convert to integer</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>float(value)</InlineCode></TableCell>
              <TableCell>Convert to float</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>str(value)</InlineCode></TableCell>
              <TableCell>Convert to string</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>bool(value)</InlineCode></TableCell>
              <TableCell>Convert to boolean</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>type(value)</InlineCode></TableCell>
              <TableCell>Get type name as string</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`num = int("42")        // 42
pi = float("3.14")     // 3.14
text = str(123)        // "123"
flag = bool(1)         // True

print(type(42))        // "INTEGER"
print(type("hello"))   // "STRING"
print(type([1, 2]))    // "ARRAY"`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="collections">
        <SectionTitle>Collection Functions</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Function</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>len(collection)</InlineCode></TableCell>
              <TableCell>Get length of array, string, or hash</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>range(start, end, step)</InlineCode></TableCell>
              <TableCell>Generate array of numbers</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>enumerate(array)</InlineCode></TableCell>
              <TableCell>Get index-value pairs</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>zip(arr1, arr2)</InlineCode></TableCell>
              <TableCell>Combine arrays into pairs</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>sorted(array)</InlineCode></TableCell>
              <TableCell>Return sorted copy of array</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>reversed(array)</InlineCode></TableCell>
              <TableCell>Return reversed copy of array</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`arr = [3, 1, 4, 1, 5]

print(len(arr))           // 5
print(len("hello"))       // 5

// Range
for i in range(5):
    print(i)  // 0, 1, 2, 3, 4

for i in range(2, 10, 2):
    print(i)  // 2, 4, 6, 8

// Enumerate
for i, val in enumerate(["a", "b", "c"]):
    print(f"{i}: {val}")

// Sorted and reversed
print(sorted(arr))        // [1, 1, 3, 4, 5]
print(reversed(arr))      // [5, 1, 4, 1, 3]`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="math">
        <SectionTitle>Math Functions</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Function</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>abs(n)</InlineCode></TableCell>
              <TableCell>Absolute value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>min(a, b, ...)</InlineCode></TableCell>
              <TableCell>Minimum value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>max(a, b, ...)</InlineCode></TableCell>
              <TableCell>Maximum value</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>sum(array)</InlineCode></TableCell>
              <TableCell>Sum of array elements</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>round(n, places)</InlineCode></TableCell>
              <TableCell>Round to decimal places</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`print(abs(-5))            // 5
print(min(3, 1, 4))       // 1
print(max(3, 1, 4))       // 4
print(sum([1, 2, 3, 4]))  // 10
print(round(3.14159, 2))  // 3.14`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="utility">
        <SectionTitle>Utility Functions</SectionTitle>

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
              <TableCell>Get Carrion version info</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>help()</InlineCode></TableCell>
              <TableCell>Show help information</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>modules()</InlineCode></TableCell>
              <TableCell>List available modules</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>mimir</InlineCode></TableCell>
              <TableCell>Interactive help system (REPL)</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// In REPL
>>> version()
Carrion v0.1.9, Munin Standard Library 0.1.0

>>> help()
// Shows language help

>>> modules()
// Lists available standard library modules

>>> mimir
// Starts interactive help system`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>
    </DocLayout>
  );
};

export default BuiltinFunctions;
