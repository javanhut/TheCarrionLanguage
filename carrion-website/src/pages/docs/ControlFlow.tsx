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
  { id: 'conditionals', title: 'Conditionals' },
  { id: 'for-loops', title: 'For Loops' },
  { id: 'while-loops', title: 'While Loops' },
  { id: 'loop-control', title: 'Loop Control' },
  { id: 'pattern-matching', title: 'Pattern Matching' },
];

const ControlFlow: React.FC = () => {
  return (
    <DocLayout
      title="Control Flow"
      description="Conditionals, loops, and pattern matching for controlling program execution."
      sections={sections}
    >
      <Section id="conditionals">
        <SectionTitle>Conditional Statements</SectionTitle>

        <SubSection>
          <SubSectionTitle>If Statement</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`age = 25
if age >= 18:
    print("You are an adult")

// If-else
score = 85
if score >= 90:
    grade = "A"
else:
    grade = "B"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>If-Otherwise-Else</SubSectionTitle>
          <Paragraph>
            Use <InlineCode>otherwise</InlineCode> for additional branches (like elif in Python):
          </Paragraph>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`temperature = 75

if temperature < 32:
    status = "Freezing"
otherwise temperature < 50:
    status = "Cold"
otherwise temperature < 70:
    status = "Cool"
otherwise temperature < 85:
    status = "Warm"
else:
    status = "Hot"

print(f"It's {status} outside")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Ternary Expression</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`age = 20
status = "adult" if age >= 18 else "minor"

max_value = a if a > b else b

message = "positive" if n > 0 else "negative" if n < 0 else "zero"`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="for-loops">
        <SectionTitle>For Loops</SectionTitle>

        <SubSection>
          <SubSectionTitle>Iterating Over Collections</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Array iteration
numbers = [1, 2, 3, 4, 5]
for num in numbers:
    print(num)

// String iteration
word = "hello"
for char in word:
    print(char)

// Hash iteration
data = {"a": 1, "b": 2}
for key in data:
    print(f"{key}: {data[key]}")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Range-Based Loops</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Basic range (0 to 4)
for i in range(5):
    print(i)  // 0, 1, 2, 3, 4

// Range with start and end
for i in range(2, 10):
    print(i)  // 2, 3, 4, 5, 6, 7, 8, 9

// Range with step
for i in range(0, 10, 2):
    print(i)  // 0, 2, 4, 6, 8`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Enumerate</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`items = ["apple", "banana", "cherry"]
for index, value in enumerate(items):
    print(f"{index}: {value}")
// Output:
// 0: apple
// 1: banana
// 2: cherry`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>For-Else</SubSectionTitle>
          <Paragraph>
            The <InlineCode>else</InlineCode> block executes if the loop completes without <InlineCode>stop</InlineCode>:
          </Paragraph>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`target = 7
numbers = [1, 3, 5, 9, 11]

for num in numbers:
    if num == target:
        print(f"Found {target}")
        stop
else:
    print(f"{target} not found")  // This runs`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="while-loops">
        <SectionTitle>While Loops</SectionTitle>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Basic while loop
count = 0
while count < 5:
    print(f"Count: {count}")
    count += 1

// While with condition
answer = ""
while answer != "quit":
    answer = input("Enter 'quit' to exit: ")
    if answer != "quit":
        print(f"You entered: {answer}")

// Infinite loop with break
while True:
    user_input = input("Enter a number (or 'exit'): ")
    if user_input == "exit":
        stop
    print(f"You entered: {user_input}")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Infinite Loops</InfoTitle>
          <InfoText>
            When using <InlineCode>while True</InlineCode>, always ensure there's a clear exit
            condition with <InlineCode>stop</InlineCode> to prevent infinite loops.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="loop-control">
        <SectionTitle>Loop Control</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Carrion</TableHead>
              <TableHead>Python/JS</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>skip</InlineCode></TableCell>
              <TableCell><InlineCode>continue</InlineCode></TableCell>
              <TableCell>Skip to next iteration</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>stop</InlineCode></TableCell>
              <TableCell><InlineCode>break</InlineCode></TableCell>
              <TableCell>Exit the loop</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <SubSection>
          <SubSectionTitle>skip (Continue)</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Skip even numbers
for i in range(10):
    if i % 2 == 0:
        skip
    print(f"Odd: {i}")  // 1, 3, 5, 7, 9

// Skip empty strings
words = ["hello", "", "world", ""]
for word in words:
    if word == "":
        skip
    print(word.upper())`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>stop (Break)</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Find first negative
numbers = [5, 3, 8, -2, 1, 7]
for num in numbers:
    if num < 0:
        print(f"Found negative: {num}")
        stop
    print(f"Positive: {num}")

// Interactive menu
while True:
    choice = input("Choice (1-3, q): ")
    if choice == "q":
        stop
    print(f"Selected: {choice}")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="pattern-matching">
        <SectionTitle>Pattern Matching</SectionTitle>
        <Paragraph>
          Use <InlineCode>match</InlineCode> for elegant pattern-based branching:
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Basic pattern matching
status_code = 404
match status_code:
    case 200:
        message = "OK"
    case 404:
        message = "Not Found"
    case 500:
        message = "Server Error"
    _:
        message = "Unknown"

print(message)  // "Not Found"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSection>
          <SubSectionTitle>String Matching</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`command = "save"
match command:
    case "save":
        print("Saving file...")
    case "load":
        print("Loading file...")
    case "quit":
        print("Goodbye!")
        return
    _:
        print("Unknown command")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Tuple Matching</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`day = "Monday"
weather = "sunny"

match (day, weather):
    case ("Monday", "sunny"):
        activity = "Go for a walk"
    case ("Monday", "rainy"):
        activity = "Work indoors"
    case ("Friday", _):  // Any weather on Friday
        activity = "Plan weekend"
    _:
        activity = "Normal routine"

print(f"Activity: {activity}")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <InfoBox>
          <InfoTitle>Default Case</InfoTitle>
          <InfoText>
            Use <InlineCode>_:</InlineCode> as the default case to handle any unmatched values.
            It should be the last case in the match statement.
          </InfoText>
        </InfoBox>
      </Section>
    </DocLayout>
  );
};

export default ControlFlow;
