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
  { id: 'arithmetic', title: 'Arithmetic' },
  { id: 'assignment', title: 'Assignment' },
  { id: 'comparison', title: 'Comparison' },
  { id: 'logical', title: 'Logical' },
  { id: 'membership', title: 'Membership' },
  { id: 'bitwise', title: 'Bitwise' },
  { id: 'precedence', title: 'Precedence' },
];

const Operators: React.FC = () => {
  return (
    <DocLayout
      title="Operators"
      description="Complete reference for arithmetic, logical, comparison, bitwise, and other operators in Carrion."
      sections={sections}
    >
      <Section id="arithmetic">
        <SectionTitle>Arithmetic Operators</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Operator</TableHead>
              <TableHead>Description</TableHead>
              <TableHead>Example</TableHead>
              <TableHead>Result</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>+</InlineCode></TableCell>
              <TableCell>Addition</TableCell>
              <TableCell><InlineCode>5 + 3</InlineCode></TableCell>
              <TableCell>8</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>-</InlineCode></TableCell>
              <TableCell>Subtraction</TableCell>
              <TableCell><InlineCode>5 - 3</InlineCode></TableCell>
              <TableCell>2</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>*</InlineCode></TableCell>
              <TableCell>Multiplication</TableCell>
              <TableCell><InlineCode>5 * 3</InlineCode></TableCell>
              <TableCell>15</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>/</InlineCode></TableCell>
              <TableCell>Division</TableCell>
              <TableCell><InlineCode>15 / 3</InlineCode></TableCell>
              <TableCell>5.0</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>{'//'}</InlineCode></TableCell>
              <TableCell>Integer Division</TableCell>
              <TableCell><InlineCode>17 // 3</InlineCode></TableCell>
              <TableCell>5</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>%</InlineCode></TableCell>
              <TableCell>Modulo</TableCell>
              <TableCell><InlineCode>17 % 3</InlineCode></TableCell>
              <TableCell>2</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>**</InlineCode></TableCell>
              <TableCell>Exponentiation</TableCell>
              <TableCell><InlineCode>2 ** 3</InlineCode></TableCell>
              <TableCell>8</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`result = 10 + 5 * 2    // 20 (follows order of operations)
power = 2 ** 3         // 8
remainder = 17 % 5     // 2

// Unary operators
positive = +42
negative = -42`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="assignment">
        <SectionTitle>Assignment Operators</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Operator</TableHead>
              <TableHead>Description</TableHead>
              <TableHead>Example</TableHead>
              <TableHead>Equivalent</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>=</InlineCode></TableCell>
              <TableCell>Basic assignment</TableCell>
              <TableCell><InlineCode>x = 5</InlineCode></TableCell>
              <TableCell>-</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>+=</InlineCode></TableCell>
              <TableCell>Add and assign</TableCell>
              <TableCell><InlineCode>x += 3</InlineCode></TableCell>
              <TableCell><InlineCode>x = x + 3</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>-=</InlineCode></TableCell>
              <TableCell>Subtract and assign</TableCell>
              <TableCell><InlineCode>x -= 3</InlineCode></TableCell>
              <TableCell><InlineCode>x = x - 3</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>*=</InlineCode></TableCell>
              <TableCell>Multiply and assign</TableCell>
              <TableCell><InlineCode>x *= 3</InlineCode></TableCell>
              <TableCell><InlineCode>x = x * 3</InlineCode></TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>/=</InlineCode></TableCell>
              <TableCell>Divide and assign</TableCell>
              <TableCell><InlineCode>x /= 3</InlineCode></TableCell>
              <TableCell><InlineCode>x = x / 3</InlineCode></TableCell>
            </TableRow>
          </tbody>
        </Table>

        <SubSection>
          <SubSectionTitle>Increment and Decrement</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`x = 5
result1 = ++x  // x becomes 6, result1 is 6 (prefix)
result2 = x++  // result2 is 6, x becomes 7 (postfix)

b = 10
result3 = --b  // b becomes 9, result3 is 9
result4 = b--  // result4 is 9, b becomes 8`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="comparison">
        <SectionTitle>Comparison Operators</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Operator</TableHead>
              <TableHead>Description</TableHead>
              <TableHead>Example</TableHead>
              <TableHead>Result</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>==</InlineCode></TableCell>
              <TableCell>Equal to</TableCell>
              <TableCell><InlineCode>5 == 5</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>!=</InlineCode></TableCell>
              <TableCell>Not equal to</TableCell>
              <TableCell><InlineCode>5 != 3</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>&lt;</InlineCode></TableCell>
              <TableCell>Less than</TableCell>
              <TableCell><InlineCode>3 &lt; 5</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>&gt;</InlineCode></TableCell>
              <TableCell>Greater than</TableCell>
              <TableCell><InlineCode>5 &gt; 3</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>&lt;=</InlineCode></TableCell>
              <TableCell>Less than or equal</TableCell>
              <TableCell><InlineCode>3 &lt;= 3</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>&gt;=</InlineCode></TableCell>
              <TableCell>Greater than or equal</TableCell>
              <TableCell><InlineCode>5 &gt;= 5</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// String comparisons (lexicographic)
print("apple" < "banana")  // True

// Chained comparisons
age = 25
valid = 18 <= age < 65     // True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="logical">
        <SectionTitle>Logical Operators</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Operator</TableHead>
              <TableHead>Description</TableHead>
              <TableHead>Example</TableHead>
              <TableHead>Result</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>and</InlineCode></TableCell>
              <TableCell>Logical AND</TableCell>
              <TableCell><InlineCode>True and False</InlineCode></TableCell>
              <TableCell>False</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>or</InlineCode></TableCell>
              <TableCell>Logical OR</TableCell>
              <TableCell><InlineCode>True or False</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>not</InlineCode></TableCell>
              <TableCell>Logical NOT</TableCell>
              <TableCell><InlineCode>not True</InlineCode></TableCell>
              <TableCell>False</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <InfoBox>
          <InfoTitle>Short-Circuit Evaluation</InfoTitle>
          <InfoText>
            <InlineCode>and</InlineCode> stops if left is false. <InlineCode>or</InlineCode> stops if left is true.
            The right operand is not evaluated in these cases.
          </InfoText>
        </InfoBox>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`has_permission = True
is_admin = False

can_edit = has_permission and is_admin  // False
can_view = has_permission or is_admin   // True
cannot_edit = not can_edit              // True

// Complex conditions
if (score > 90 and attendance > 85) or is_honors:
    print("Qualifies for award")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="membership">
        <SectionTitle>Membership Operators</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Operator</TableHead>
              <TableHead>Description</TableHead>
              <TableHead>Example</TableHead>
              <TableHead>Result</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>in</InlineCode></TableCell>
              <TableCell>Membership test</TableCell>
              <TableCell><InlineCode>"a" in "apple"</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>not in</InlineCode></TableCell>
              <TableCell>Negative membership</TableCell>
              <TableCell><InlineCode>"z" not in "apple"</InlineCode></TableCell>
              <TableCell>True</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// String membership
print("a" in "banana")      // True

// Array membership
numbers = [1, 2, 3, 4, 5]
print(3 in numbers)         // True
print(6 not in numbers)     // True

// Hash membership (checks keys)
data = {"name": "John", "age": 30}
print("name" in data)       // True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="bitwise">
        <SectionTitle>Bitwise Operators</SectionTitle>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Operator</TableHead>
              <TableHead>Description</TableHead>
              <TableHead>Example</TableHead>
              <TableHead>Result</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>&</InlineCode></TableCell>
              <TableCell>Bitwise AND</TableCell>
              <TableCell><InlineCode>5 & 3</InlineCode></TableCell>
              <TableCell>1</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>|</InlineCode></TableCell>
              <TableCell>Bitwise OR</TableCell>
              <TableCell><InlineCode>5 | 3</InlineCode></TableCell>
              <TableCell>7</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>^</InlineCode></TableCell>
              <TableCell>Bitwise XOR</TableCell>
              <TableCell><InlineCode>5 ^ 3</InlineCode></TableCell>
              <TableCell>6</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>~</InlineCode></TableCell>
              <TableCell>Bitwise NOT</TableCell>
              <TableCell><InlineCode>~5</InlineCode></TableCell>
              <TableCell>-6</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>&lt;&lt;</InlineCode></TableCell>
              <TableCell>Left shift</TableCell>
              <TableCell><InlineCode>5 &lt;&lt; 1</InlineCode></TableCell>
              <TableCell>10</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>&gt;&gt;</InlineCode></TableCell>
              <TableCell>Right shift</TableCell>
              <TableCell><InlineCode>10 &gt;&gt; 1</InlineCode></TableCell>
              <TableCell>5</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`a = 5    // Binary: 101
b = 3    // Binary: 011

and_result = a & b    // 1 (Binary: 001)
or_result = a | b     // 7 (Binary: 111)
xor_result = a ^ b    // 6 (Binary: 110)
left_shift = a << 2   // 20 (Binary: 10100)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="precedence">
        <SectionTitle>Operator Precedence</SectionTitle>
        <Paragraph>
          Operators are evaluated in this order (highest to lowest):
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Priority</TableHead>
              <TableHead>Operators</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell>1</TableCell>
              <TableCell><InlineCode>()</InlineCode></TableCell>
              <TableCell>Parentheses</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>2</TableCell>
              <TableCell><InlineCode>**</InlineCode></TableCell>
              <TableCell>Exponentiation</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>3</TableCell>
              <TableCell><InlineCode>+</InlineCode> <InlineCode>-</InlineCode> <InlineCode>not</InlineCode> <InlineCode>~</InlineCode></TableCell>
              <TableCell>Unary operators</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>4</TableCell>
              <TableCell><InlineCode>*</InlineCode> <InlineCode>/</InlineCode> <InlineCode>{'//'}</InlineCode> <InlineCode>%</InlineCode></TableCell>
              <TableCell>Multiplicative</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>5</TableCell>
              <TableCell><InlineCode>+</InlineCode> <InlineCode>-</InlineCode></TableCell>
              <TableCell>Additive</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>6</TableCell>
              <TableCell><InlineCode>&lt;&lt;</InlineCode> <InlineCode>&gt;&gt;</InlineCode></TableCell>
              <TableCell>Shift</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>7-9</TableCell>
              <TableCell><InlineCode>&</InlineCode> <InlineCode>^</InlineCode> <InlineCode>|</InlineCode></TableCell>
              <TableCell>Bitwise AND, XOR, OR</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>10</TableCell>
              <TableCell><InlineCode>==</InlineCode> <InlineCode>!=</InlineCode> <InlineCode>&lt;</InlineCode> <InlineCode>&gt;</InlineCode> <InlineCode>in</InlineCode></TableCell>
              <TableCell>Comparison</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>11-13</TableCell>
              <TableCell><InlineCode>not</InlineCode> <InlineCode>and</InlineCode> <InlineCode>or</InlineCode></TableCell>
              <TableCell>Logical</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>14</TableCell>
              <TableCell><InlineCode>=</InlineCode> <InlineCode>+=</InlineCode> <InlineCode>-=</InlineCode></TableCell>
              <TableCell>Assignment</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`result = 2 + 3 * 4       // 14 (not 20)
result = (2 + 3) * 4     // 20
result = 2 ** 3 * 4      // 32 (exponentiation first)
result = not False and True  // True`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Use Parentheses for Clarity</InfoTitle>
          <InfoText>
            When precedence is unclear, use parentheses to make your intent explicit.
            <InlineCode>(a + b) * c</InlineCode> is clearer than relying on precedence rules.
          </InfoText>
        </InfoBox>
      </Section>
    </DocLayout>
  );
};

export default Operators;
