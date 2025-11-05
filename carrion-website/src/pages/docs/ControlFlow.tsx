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
  &:nth-child(5) { animation-delay: 0.4s; }
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

const ControlFlow: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Control Flow</Title>
        <Subtitle>
          Master conditionals, loops, pattern matching, and flow control for precise program execution.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Conditional Statements</SectionTitle>
        
        <SubSectionTitle>Basic If Statement</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Simple if
age = 25
if age >= 18:
    print("You are an adult")

// If-else
score = 85
if score >= 90:
    grade = "A"
else:
    grade = "B"

print(f"Your grade: {grade}")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>If-Otherwise-Else</SubSectionTitle>
        <Text>
          Use <InlineCode>otherwise</InlineCode> for additional conditional branches (like elif in Python).
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
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

print(f"It's {status} outside")

// Grade calculation
score = 87
if score >= 90:
    grade = "A"
otherwise score >= 80:
    grade = "B"
otherwise score >= 70:
    grade = "C"
otherwise score >= 60:
    grade = "D"
else:
    grade = "F"

print(f"Grade: {grade}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Loops</SectionTitle>
        
        <SubSectionTitle>For Loops</SubSectionTitle>
        <Text>
          Iterate over sequences like arrays, strings, and ranges.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Iterate over array
numbers = [1, 2, 3, 4, 5]
for num in numbers:
    print(num)

// Iterate over string
word = "hello"
for char in word:
    print(char)

// Iterate over range
for i in range(5):
    print(f"Count: {i}")  // 0, 1, 2, 3, 4

// Range with parameters
for i in range(2, 10, 2):  // start, stop, step
    print(i)  // 2, 4, 6, 8

// Enumerate for index and value
items = ["apple", "banana", "cherry"]
for index, value in enumerate(items):
    print(f"{index}: {value}")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>For Loop with Else</SubSectionTitle>
        <Text>
          The else clause executes if the loop completes without breaking.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`target = 7
numbers = [1, 3, 5, 9, 11]

for num in numbers:
    if num == target:
        print(f"Found {target}")
        stop  // Break from loop
else:
    print(f"{target} not found")  // Executes if no break`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>While Loops</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Basic while loop
count = 0
while count < 5:
    print(f"Count: {count}")
    count += 1

// While with user input
answer = ""
while answer != "quit":
    answer = input("Enter 'quit' to exit: ")
    if answer != "quit":
        print(f"You entered: {answer}")

// Infinite loop with break
while True:
    user_input = input("Enter a number (or 'exit'): ")
    if user_input == "exit":
        stop  // Break from loop
    
    attempt:
        number = int(user_input)
        print(f"Square: {number ** 2}")
    ensnare:
        print("Invalid number")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Loop Control</SectionTitle>
        
        <SubSectionTitle>skip (Continue)</SubSectionTitle>
        <Text>
          Jump to the next iteration of the loop.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Skip even numbers
for i in range(10):
    if i % 2 == 0:
        skip  // Continue to next iteration
    print(f"Odd number: {i}")

// Skip empty strings
words = ["hello", "", "world", "", "carrion"]
for word in words:
    if word == "":
        skip
    print(word.upper())`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>stop (Break)</SubSectionTitle>
        <Text>
          Exit the loop immediately.
        </Text>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Find first negative number
numbers = [5, 3, 8, -2, 1, 7]
for num in numbers:
    if num < 0:
        print(f"Found negative: {num}")
        stop  // Exit loop
    print(f"Positive: {num}")

// Interactive menu
while True:
    choice = input("Enter choice (1-3, or 'q'): ")
    if choice == "q":
        stop
    elif choice == "1":
        print("Option 1 selected")
    elif choice == "2":
        print("Option 2 selected")
    elif choice == "3":
        print("Option 3 selected")
    else:
        print("Invalid choice")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Pattern Matching</SectionTitle>
        <Text>
          Use <InlineCode>match</InlineCode> statements for elegant pattern-based branching.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Basic pattern matching
status_code = 404
match status_code:
    case 200:
        message = "OK"
    case 404:
        message = "Not Found"
    case 500:
        message = "Internal Server Error"
    _:
        message = "Unknown Status"

print(message)

// Pattern matching with strings
command = "save"
match command:
    case "save":
        print("Saving file...")
    case "load":
        print("Loading file...")
    case "quit":
        print("Goodbye!")
        return
    _:
        print("Unknown command")

// Multiple values
day = "Monday"
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
      </Section>

      <Section>
        <SectionTitle>Nested Control Flow</SectionTitle>
        
        <SubSectionTitle>Nested Loops</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Multiplication table
for i in range(1, 6):
    for j in range(1, 6):
        product = i * j
        print(f"{i} x {j} = {product}")
    print()  // Empty line after each table

// Finding items in nested structure
matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
target = 5
found = False

for row in matrix:
    for item in row:
        if item == target:
            print(f"Found {target}")
            found = True
            stop  // Break from inner loop
    if found:
        stop  // Break from outer loop`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Nested Conditionals</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Grade calculation with criteria
score = 87
attendance = 95

if score >= 90:
    if attendance >= 90:
        grade = "A"
    else:
        grade = "A-"
otherwise score >= 80:
    if attendance >= 90:
        grade = "B+"
    otherwise attendance >= 80:
        grade = "B"
    else:
        grade = "B-"
else:
    if attendance >= 90:
        grade = "C+"
    else:
        grade = "C"

print(f"Final grade: {grade}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Advanced Patterns</SectionTitle>
        
        <SubSectionTitle>Loop with Multiple Conditions</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Process items until condition met
items = [1, 3, 5, 7, 9, 12, 15]
sum_total = 0
index = 0

while index < len(items) and sum_total < 20:
    sum_total += items[index]
    print(f"Added {items[index]}, total: {sum_total}")
    index += 1

print(f"Final total: {sum_total}")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Conditional Expressions</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Ternary-like expressions
age = 20
status = "adult" if age >= 18 else "minor"

// More complex
max_value = a if a > b else b
message = "positive" if number > 0 else "negative" if number < 0 else "zero"

// Use in assignments
price = base_price * (0.9 if is_member else 1.0)
greeting = f"Good {'morning' if hour < 12 else 'afternoon'}"`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Complete Examples</SectionTitle>
        
        <SubSectionTitle>Menu-Driven Program</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell calculator_menu():
    while True:
        print("\\n=== Calculator ===")
        print("1. Add")
        print("2. Subtract")
        print("3. Multiply")
        print("4. Divide")
        print("5. Exit")
        
        choice = input("Choose (1-5): ")
        
        if choice == "5":
            print("Goodbye!")
            stop
        
        if choice not in ["1", "2", "3", "4"]:
            print("Invalid choice!")
            skip
        
        a = float(input("First number: "))
        b = float(input("Second number: "))
        
        match choice:
            case "1":
                print(f"Result: {a + b}")
            case "2":
                print(f"Result: {a - b}")
            case "3":
                print(f"Result: {a * b}")
            case "4":
                if b != 0:
                    print(f"Result: {a / b}")
                else:
                    print("Error: Division by zero!")

calculator_menu()`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Data Validation Loop</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell get_valid_age():
    while True:
        attempt:
            age = int(input("Enter your age (0-150): "))
            
            if age < 0:
                print("Age cannot be negative!")
                skip
            
            if age > 150:
                print("Age seems unrealistic!")
                skip
            
            return age  // Valid age
        
        ensnare:
            print("Invalid input! Enter a number.")

age = get_valid_age()
print(f"Your age: {age}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Best Practices</SectionTitle>
        
        <InfoBox>
          <InfoTitle>Keep Conditions Simple</InfoTitle>
          <InfoText>
            Extract complex conditions into well-named variables or functions. 
            <InlineCode>if is_valid_user()</InlineCode> is clearer than a long compound condition.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Avoid Deep Nesting</InfoTitle>
          <InfoText>
            Use early returns instead of deeply nested if statements. This makes code more readable 
            and easier to maintain.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Use Match for Multiple Values</InfoTitle>
          <InfoText>
            When checking a variable against multiple values, prefer <InlineCode>match</InlineCode> over 
            long if-otherwise chains for better readability.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Be Careful with Infinite Loops</InfoTitle>
          <InfoText>
            When using <InlineCode>while True</InlineCode>, ensure there's a clear exit condition with 
            <InlineCode>stop</InlineCode> to prevent infinite loops.
          </InfoText>
        </InfoBox>
      </Section>
    </Container>
  );
};

export default ControlFlow;
