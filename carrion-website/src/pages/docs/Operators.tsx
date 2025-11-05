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

const OperatorTable = styled.table`
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

const Operators: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Operators & Expressions</Title>
        <Subtitle>
          Complete reference for all operators in Carrion: arithmetic, logical, comparison, bitwise, and more.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Arithmetic Operators</SectionTitle>
        
        <OperatorTable>
          <thead>
            <tr>
              <TableHeader>Operator</TableHeader>
              <TableHeader>Description</TableHeader>
              <TableHeader>Example</TableHeader>
              <TableHeader>Result</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell code>+</TableCell>
              <TableCell>Addition</TableCell>
              <TableCell code>5 + 3</TableCell>
              <TableCell code>8</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>-</TableCell>
              <TableCell>Subtraction</TableCell>
              <TableCell code>5 - 3</TableCell>
              <TableCell code>2</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>*</TableCell>
              <TableCell>Multiplication</TableCell>
              <TableCell code>5 * 3</TableCell>
              <TableCell code>15</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>/</TableCell>
              <TableCell>Division</TableCell>
              <TableCell code>15 / 3</TableCell>
              <TableCell code>5.0</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>{'//'}</TableCell>
              <TableCell>Integer Division</TableCell>
              <TableCell code>{'17 // 3'}</TableCell>
              <TableCell code>5</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>%</TableCell>
              <TableCell>Modulo</TableCell>
              <TableCell code>17 % 3</TableCell>
              <TableCell code>2</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>**</TableCell>
              <TableCell>Exponentiation</TableCell>
              <TableCell code>2 ** 3</TableCell>
              <TableCell code>8</TableCell>
            </TableRow>
          </tbody>
        </OperatorTable>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Basic arithmetic
result = 10 + 5 * 2    // 20 (follows order of operations)
power = 2 ** 3         // 8
remainder = 17 % 5     // 2

// Unary operators
positive = +42         // 42
negative = -42         // -42

// Works with variables
x, y = (10, 20)
sum = x + y           // 30
negated = -x          // -10`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Assignment Operators</SectionTitle>
        
        <OperatorTable>
          <thead>
            <tr>
              <TableHeader>Operator</TableHeader>
              <TableHeader>Description</TableHeader>
              <TableHeader>Example</TableHeader>
              <TableHeader>Equivalent</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell code>=</TableCell>
              <TableCell>Basic assignment</TableCell>
              <TableCell code>x = 5</TableCell>
              <TableCell code>-</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>+=</TableCell>
              <TableCell>Add and assign</TableCell>
              <TableCell code>x += 3</TableCell>
              <TableCell code>x = x + 3</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>-=</TableCell>
              <TableCell>Subtract and assign</TableCell>
              <TableCell code>x -= 3</TableCell>
              <TableCell code>x = x - 3</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>*=</TableCell>
              <TableCell>Multiply and assign</TableCell>
              <TableCell code>x *= 3</TableCell>
              <TableCell code>x = x * 3</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>/=</TableCell>
              <TableCell>Divide and assign</TableCell>
              <TableCell code>x /= 3</TableCell>
              <TableCell code>x = x / 3</TableCell>
            </TableRow>
          </tbody>
        </OperatorTable>

        <SubSectionTitle>Increment and Decrement</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Compound assignment
x = 10
x += 5        // x is now 15
x *= 2        // x is now 30
x /= 3        // x is now 10

// Increment/decrement operators
a = 5
result1 = ++a  // a becomes 6, result1 is 6 (prefix)
result2 = a++  // result2 is 6, a becomes 7 (postfix)

b = 10
result3 = --b  // b becomes 9, result3 is 9 (prefix)
result4 = b--  // result4 is 9, b becomes 8 (postfix)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Comparison Operators</SectionTitle>
        
        <OperatorTable>
          <thead>
            <tr>
              <TableHeader>Operator</TableHeader>
              <TableHeader>Description</TableHeader>
              <TableHeader>Example</TableHeader>
              <TableHeader>Result</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell code>==</TableCell>
              <TableCell>Equal to</TableCell>
              <TableCell code>5 == 5</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>!=</TableCell>
              <TableCell>Not equal to</TableCell>
              <TableCell code>5 != 3</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>&lt;</TableCell>
              <TableCell>Less than</TableCell>
              <TableCell code>3 &lt; 5</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>&gt;</TableCell>
              <TableCell>Greater than</TableCell>
              <TableCell code>5 &gt; 3</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>&lt;=</TableCell>
              <TableCell>Less than or equal</TableCell>
              <TableCell code>3 &lt;= 5</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>&gt;=</TableCell>
              <TableCell>Greater than or equal</TableCell>
              <TableCell code>5 &gt;= 5</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
          </tbody>
        </OperatorTable>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Numeric comparisons
print(10 > 5)     // True
print(3 <= 3)     // True
print(7 != 8)     // True

// String comparisons (lexicographic)
print("apple" < "banana")  // True
print("hello" == "hello")  // True

// Chained comparisons
age = 25
valid = 18 <= age < 65     // True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Logical Operators</SectionTitle>
        
        <OperatorTable>
          <thead>
            <tr>
              <TableHeader>Operator</TableHeader>
              <TableHeader>Description</TableHeader>
              <TableHeader>Example</TableHeader>
              <TableHeader>Result</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell code>and</TableCell>
              <TableCell>Logical AND</TableCell>
              <TableCell code>True and False</TableCell>
              <TableCell code>False</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>or</TableCell>
              <TableCell>Logical OR</TableCell>
              <TableCell code>True or False</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>not</TableCell>
              <TableCell>Logical NOT</TableCell>
              <TableCell code>not True</TableCell>
              <TableCell code>False</TableCell>
            </TableRow>
          </tbody>
        </OperatorTable>

        <InfoBox>
          <InfoTitle>Short-Circuit Evaluation</InfoTitle>
          <InfoText>
            <InlineCode>and</InlineCode> stops if left is false. <InlineCode>or</InlineCode> stops if left is true. 
            The right operand is not evaluated in these cases.
          </InfoText>
        </InfoBox>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Basic logical operations
has_permission = True
is_admin = False

can_edit = has_permission and is_admin     // False
can_view = has_permission or is_admin      // True
cannot_edit = not can_edit                 // True

// Practical usage
if age >= 18 and has_id:
    print("Can enter")

// Complex conditions
if (score > 90 and attendance > 85) or is_honors_student:
    print("Qualifies for award")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Membership Operators</SectionTitle>
        
        <OperatorTable>
          <thead>
            <tr>
              <TableHeader>Operator</TableHeader>
              <TableHeader>Description</TableHeader>
              <TableHeader>Example</TableHeader>
              <TableHeader>Result</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell code>in</TableCell>
              <TableCell>Membership test</TableCell>
              <TableCell code>"a" in "apple"</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>not in</TableCell>
              <TableCell>Negative membership</TableCell>
              <TableCell code>"z" not in "apple"</TableCell>
              <TableCell code>True</TableCell>
            </TableRow>
          </tbody>
        </OperatorTable>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// String membership
letter = "a"
word = "banana"
print(letter in word)      // True
print("z" not in word)     // True

// Array membership
numbers = [1, 2, 3, 4, 5]
print(3 in numbers)        // True
print(6 not in numbers)    // True

// Hash membership (checks keys)
data = {"name": "John", "age": 30}
print("name" in data)      // True
print("email" not in data) // True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Bitwise Operators</SectionTitle>
        
        <OperatorTable>
          <thead>
            <tr>
              <TableHeader>Operator</TableHeader>
              <TableHeader>Description</TableHeader>
              <TableHeader>Example</TableHeader>
              <TableHeader>Result</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell code>&</TableCell>
              <TableCell>Bitwise AND</TableCell>
              <TableCell code>5 & 3</TableCell>
              <TableCell code>1</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>|</TableCell>
              <TableCell>Bitwise OR</TableCell>
              <TableCell code>5 | 3</TableCell>
              <TableCell code>7</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>^</TableCell>
              <TableCell>Bitwise XOR</TableCell>
              <TableCell code>5 ^ 3</TableCell>
              <TableCell code>6</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>~</TableCell>
              <TableCell>Bitwise NOT</TableCell>
              <TableCell code>~5</TableCell>
              <TableCell code>-6</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>&lt;&lt;</TableCell>
              <TableCell>Left shift</TableCell>
              <TableCell code>5 &lt;&lt; 1</TableCell>
              <TableCell code>10</TableCell>
            </TableRow>
            <TableRow>
              <TableCell code>&gt;&gt;</TableCell>
              <TableCell>Right shift</TableCell>
              <TableCell code>10 &gt;&gt; 1</TableCell>
              <TableCell code>5</TableCell>
            </TableRow>
          </tbody>
        </OperatorTable>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Bitwise operations
a = 5    // Binary: 101
b = 3    // Binary: 011

and_result = a & b    // 1 (Binary: 001)
or_result = a | b     // 7 (Binary: 111)
xor_result = a ^ b    // 6 (Binary: 110)
not_result = ~a       // -6 (Two's complement)

// Bit shifting
left_shift = a << 2   // 20 (Binary: 10100)
right_shift = a >> 1  // 2 (Binary: 10)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Operator Precedence</SectionTitle>
        <Text>
          Operators are evaluated in the following order (highest to lowest):
        </Text>

        <OperatorTable>
          <thead>
            <tr>
              <TableHeader>Priority</TableHeader>
              <TableHeader>Operators</TableHeader>
              <TableHeader>Description</TableHeader>
            </tr>
          </thead>
          <tbody>
            <TableRow>
              <TableCell>1</TableCell>
              <TableCell code>()</TableCell>
              <TableCell>Parentheses</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>2</TableCell>
              <TableCell code>**</TableCell>
              <TableCell>Exponentiation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>3</TableCell>
              <TableCell code>+, -, not, ~</TableCell>
              <TableCell>Unary operators</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>4</TableCell>
              <TableCell code>*, /, //, %</TableCell>
              <TableCell>Multiplicative</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>5</TableCell>
              <TableCell code>+, -</TableCell>
              <TableCell>Additive</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>6</TableCell>
              <TableCell code>&lt;&lt;, &gt;&gt;</TableCell>
              <TableCell>Shift</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>7</TableCell>
              <TableCell code>&</TableCell>
              <TableCell>Bitwise AND</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>8</TableCell>
              <TableCell code>^</TableCell>
              <TableCell>Bitwise XOR</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>9</TableCell>
              <TableCell code>|</TableCell>
              <TableCell>Bitwise OR</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>10</TableCell>
              <TableCell code>==, !=, &lt;, &gt;, &lt;=, &gt;=, in, not in</TableCell>
              <TableCell>Comparison</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>11</TableCell>
              <TableCell code>not</TableCell>
              <TableCell>Logical NOT</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>12</TableCell>
              <TableCell code>and</TableCell>
              <TableCell>Logical AND</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>13</TableCell>
              <TableCell code>or</TableCell>
              <TableCell>Logical OR</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>14</TableCell>
              <TableCell code>=, +=, -=, *=, /=</TableCell>
              <TableCell>Assignment</TableCell>
            </TableRow>
          </tbody>
        </OperatorTable>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Precedence examples
result = 2 + 3 * 4        // 14 (not 20)
result = (2 + 3) * 4      // 20
result = 2 ** 3 * 4       // 32 (exponentiation first)
result = not False and True  // True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Special Operators</SectionTitle>
        
        <SubSectionTitle>Dot Operator (Member Access)</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Method calls
text = "hello"
uppercase = text.upper()     // "HELLO"

numbers = [1, 2, 3]
length = numbers.length()    // 3

// Chained method calls
result = "  hello  ".strip().upper()  // "HELLO"`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Decorator Symbol (@)</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Decorator usage in abstract methods
arcane grim AbstractClass:
    @arcanespell
    spell abstract_method():
        ignore`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Best Practices</SectionTitle>
        
        <InfoBox>
          <InfoTitle>Use Parentheses for Clarity</InfoTitle>
          <InfoText>
            When operator precedence is unclear, use parentheses to make your intent explicit: 
            <InlineCode>(a + b) * c</InlineCode> is clearer than <InlineCode>a + b * c</InlineCode>
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Avoid Complex Expressions</InfoTitle>
          <InfoText>
            Break complex expressions into multiple lines with intermediate variables. This improves 
            readability and makes debugging easier.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Use Appropriate Operators</InfoTitle>
          <InfoText>
            Choose the right operator for the task: use <InlineCode>{'//'}</InlineCode> for integer division, 
            <InlineCode>%</InlineCode> for remainders, and <InlineCode>in</InlineCode> for membership tests.
          </InfoText>
        </InfoBox>
      </Section>
    </Container>
  );
};

export default Operators;
