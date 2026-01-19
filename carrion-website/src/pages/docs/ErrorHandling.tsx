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
  { id: 'basics', title: 'Basics' },
  { id: 'specific-errors', title: 'Catching Specific Errors' },
  { id: 'raising', title: 'Raising Errors' },
  { id: 'custom-errors', title: 'Custom Errors' },
  { id: 'best-practices', title: 'Best Practices' },
];

const ErrorHandling: React.FC = () => {
  return (
    <DocLayout
      title="Error Handling"
      description="Handle errors gracefully with attempt, ensnare, and resolve."
      sections={sections}
    >
      <Section id="basics">
        <SectionTitle>Error Handling Basics</SectionTitle>
        <Paragraph>
          Carrion uses magical keywords for error handling that map to familiar concepts:
        </Paragraph>

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
              <TableCell><InlineCode>attempt</InlineCode></TableCell>
              <TableCell><InlineCode>try</InlineCode></TableCell>
              <TableCell>Try to execute code that might fail</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>ensnare</InlineCode></TableCell>
              <TableCell><InlineCode>except</InlineCode> / <InlineCode>catch</InlineCode></TableCell>
              <TableCell>Catch and handle errors</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>resolve</InlineCode></TableCell>
              <TableCell><InlineCode>finally</InlineCode></TableCell>
              <TableCell>Always execute (cleanup)</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>raise</InlineCode></TableCell>
              <TableCell><InlineCode>raise</InlineCode> / <InlineCode>throw</InlineCode></TableCell>
              <TableCell>Throw an error</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell divide(a, b):
    attempt:
        result = a / b
        return result
    ensnare (ZeroDivisionError):
        print("Cannot divide by zero!")
        return None
    resolve:
        print("Division operation completed")

// Usage
print(divide(10, 2))  // 5.0, "Division operation completed"
print(divide(10, 0))  // "Cannot divide by zero!", None`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="specific-errors">
        <SectionTitle>Catching Specific Errors</SectionTitle>

        <SubSection>
          <SubSectionTitle>Single Error Type</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`attempt:
    value = int(user_input)
ensnare (ValueError):
    print("Please enter a valid number")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Multiple Error Types</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`attempt:
    file = File()
    data = file.read("config.txt")
    config = parse_json(data)
ensnare (FileNotFoundError):
    print("Config file not found")
    config = default_config()
ensnare (JsonParseError):
    print("Invalid JSON in config")
    config = default_config()
ensnare:
    print("Unexpected error occurred")
    config = default_config()`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Catching All Errors</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`attempt:
    risky_operation()
ensnare:
    print("Something went wrong")
    // Handle any error`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <InfoBox>
          <InfoTitle>Specific vs General</InfoTitle>
          <InfoText>
            Catch specific errors first, then use a general <InlineCode>ensnare</InlineCode> as a fallback.
            Avoid catching all errors silently - always log or handle them appropriately.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="raising">
        <SectionTitle>Raising Errors</SectionTitle>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell validate_age(age):
    if age < 0:
        raise ValueError("Age cannot be negative")
    if age > 150:
        raise ValueError("Age seems unrealistic")
    return True

spell withdraw(amount, balance):
    if amount <= 0:
        raise ValueError("Amount must be positive")
    if amount > balance:
        raise Error("InsufficientFunds", "Not enough balance")
    return balance - amount

// Using the functions
attempt:
    validate_age(-5)
ensnare (ValueError) as e:
    print(f"Validation error: {e}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="custom-errors">
        <SectionTitle>Custom Errors</SectionTitle>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Define custom error grimoire
grim ValidationError:
    init(field, message):
        self.field = field
        self.message = message

    spell to_string():
        return f"ValidationError: {self.field} - {self.message}"

grim AuthenticationError:
    init(reason):
        self.reason = reason

    spell to_string():
        return f"AuthenticationError: {self.reason}"

// Using custom errors
spell validate_user(username, password):
    if len(username) < 3:
        raise ValidationError("username", "Too short")
    if len(password) < 8:
        raise ValidationError("password", "Must be 8+ characters")
    return True

// Handle custom errors
attempt:
    validate_user("ab", "short")
ensnare (ValidationError) as e:
    print(f"Validation failed: {e.field} - {e.message}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="best-practices">
        <SectionTitle>Best Practices</SectionTitle>

        <SubSection>
          <SubSectionTitle>Resource Cleanup with Resolve</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell process_file(filename):
    file = None
    attempt:
        file = File()
        data = file.read(filename)
        return process(data)
    ensnare (FileNotFoundError):
        print(f"File not found: {filename}")
        return None
    resolve:
        // This always runs - perfect for cleanup
        if file:
            file.close()
        print("Cleanup complete")`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Input Validation</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`spell get_valid_input(prompt, validator):
    while True:
        attempt:
            value = input(prompt)
            if validator(value):
                return value
            print("Invalid input, try again")
        ensnare:
            print("Error processing input, try again")

// Usage
age = get_valid_input(
    "Enter age: ",
    (x) -> int(x) >= 0 and int(x) <= 150
)`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <InfoBox>
          <InfoTitle>Key Guidelines</InfoTitle>
          <InfoText>
            1. Use specific error types when possible.
            2. Always clean up resources in <InlineCode>resolve</InlineCode>.
            3. Don't catch errors you can't handle properly.
            4. Provide meaningful error messages.
          </InfoText>
        </InfoBox>
      </Section>
    </DocLayout>
  );
};

export default ErrorHandling;
