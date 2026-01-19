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
  InlineCode,
} from '../../components/docs';

const sections = [
  { id: 'importing', title: 'Importing Modules' },
  { id: 'creating', title: 'Creating Modules' },
  { id: 'organization', title: 'Organization' },
  { id: 'best-practices', title: 'Best Practices' },
];

const Modules: React.FC = () => {
  return (
    <DocLayout
      title="Modules"
      description="Organize and share code with Carrion's module system."
      sections={sections}
    >
      <Section id="importing">
        <SectionTitle>Importing Modules</SectionTitle>

        <SubSection>
          <SubSectionTitle>Basic Import</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Import a module
import math

// Use functions from the module
result = math.sqrt(16)
print(result)  // 4.0`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Import with Alias</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Import with alias for convenience
import math as m
import my_long_module_name as short

result = m.sqrt(25)
short.some_function()`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Relative Imports</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// Import from same directory
import utils

// Import from subdirectory
import helpers.math_helpers

// Use imported module
helpers.math_helpers.calculate()`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="creating">
        <SectionTitle>Creating Modules</SectionTitle>
        <Paragraph>
          Any <InlineCode>.crl</InlineCode> file can be used as a module. Simply define functions,
          classes, and variables that you want to export.
        </Paragraph>

        <SubSection>
          <SubSectionTitle>Example Module: utils.crl</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// utils.crl

spell add(a, b):
    return a + b

spell multiply(a, b):
    return a * b

grim Calculator:
    init():
        self.result = 0

    spell add(value):
        self.result += value
        return self

    spell get_result():
        return self.result

// Module-level constant
PI = 3.14159`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Using the Module</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// main.crl
import utils

// Use functions
sum = utils.add(5, 3)
print(sum)  // 8

// Use classes
calc = utils.Calculator()
calc.add(10).add(5)
print(calc.get_result())  // 15

// Use constants
print(utils.PI)  // 3.14159`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="organization">
        <SectionTitle>Code Organization</SectionTitle>

        <SubSection>
          <SubSectionTitle>Project Structure</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`my_project/
  main.crl           # Entry point
  config.crl         # Configuration
  utils/
    helpers.crl      # Utility functions
    validators.crl   # Validation logic
  models/
    user.crl         # User grimoire
    product.crl      # Product grimoire
  services/
    auth.crl         # Authentication
    api.crl          # API calls`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Importing from Subdirectories</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`// main.crl
import utils.helpers
import utils.validators as validate
import models.user
import services.auth

// Use imported modules
helpers.format_string("hello")
validate.check_email("test@example.com")
user = models.user.User("Alice")
services.auth.login(user)`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="best-practices">
        <SectionTitle>Best Practices</SectionTitle>

        <InfoBox>
          <InfoTitle>One Responsibility per Module</InfoTitle>
          <InfoText>
            Keep modules focused on a single purpose. A module for "user authentication" should only
            contain auth-related code, not user profile management.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Use Meaningful Names</InfoTitle>
          <InfoText>
            Choose descriptive module names that indicate their purpose: <InlineCode>validators.crl</InlineCode>,
            <InlineCode>api_client.crl</InlineCode>, <InlineCode>string_helpers.crl</InlineCode>.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Avoid Circular Imports</InfoTitle>
          <InfoText>
            If module A imports module B, module B should not import module A. Restructure code to
            avoid circular dependencies.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Group Related Functions</InfoTitle>
          <InfoText>
            Keep related functions and classes together in the same module. Split into separate
            modules when they grow too large or serve different purposes.
          </InfoText>
        </InfoBox>
      </Section>
    </DocLayout>
  );
};

export default Modules;
