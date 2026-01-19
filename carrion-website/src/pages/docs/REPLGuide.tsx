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
  TipBox,
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableCell,
  CardGrid,
  Card,
  CardTitle,
  CardDescription,
  InlineCode,
} from '../../components/docs';

const sections = [
  { id: 'intro', title: 'Introduction' },
  { id: 'features', title: 'Key Features' },
  { id: 'basic-usage', title: 'Basic Usage' },
  { id: 'multiline', title: 'Multi-line Mode' },
  { id: 'mimir', title: 'Mimir Help System' },
  { id: 'tips', title: 'Tips & Tricks' },
];

const REPLGuide: React.FC = () => {
  return (
    <DocLayout
      title="Interactive REPL Guide"
      description="Master the Carrion interactive shell for rapid development and experimentation."
      sections={sections}
    >
      <Section id="intro">
        <SectionTitle>What is the REPL?</SectionTitle>
        <Lead>
          The Carrion REPL (Read-Eval-Print Loop) provides an interactive environment where you can
          execute code, test features, and experiment with language concepts in real-time.
        </Lead>

        <InfoBox>
          <InfoTitle>Quick Start</InfoTitle>
          <InfoText>
            Simply run <InlineCode>carrion</InlineCode> in your terminal to launch the REPL.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="features">
        <SectionTitle>Key Features</SectionTitle>

        <CardGrid>
          <Card>
            <CardTitle>Clean Output</CardTitle>
            <CardDescription>
              Assignment statements and definitions don't clutter your screen - only expression
              results are displayed.
            </CardDescription>
          </Card>
          <Card>
            <CardTitle>Smart Completion</CardTitle>
            <CardDescription>
              Tab completion for keywords, functions, variables, and grimoire methods speeds
              up your workflow.
            </CardDescription>
          </Card>
          <Card>
            <CardTitle>Command History</CardTitle>
            <CardDescription>
              Navigate through previous commands with arrow keys. History persists between sessions.
            </CardDescription>
          </Card>
          <Card>
            <CardTitle>Multi-line Input</CardTitle>
            <CardDescription>
              Automatic detection of incomplete statements lets you write functions and classes naturally.
            </CardDescription>
          </Card>
        </CardGrid>
      </Section>

      <Section id="basic-usage">
        <SectionTitle>Basic Usage</SectionTitle>

        <SubSection>
          <SubSectionTitle>Starting the REPL</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`carrion`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Simple Expressions</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`>>> 2 + 2
4
>>> "Hello" + " " + "World"
Hello World
>>> [1, 2, 3] + [4, 5]
[1, 2, 3, 4, 5]`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Variable Assignment</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`>>> name = "Alice"         // No output
>>> age = 30                // No output
>>> print(f"{name} is {age}")
Alice is 30`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="multiline">
        <SectionTitle>Multi-line Mode</SectionTitle>
        <Paragraph>
          The REPL automatically enters multi-line mode when it detects an incomplete statement,
          such as a function or class definition.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`>>> spell factorial(n):
...     if n <= 1:
...         return 1
...     return n * factorial(n - 1)
...
>>> factorial(5)
120`}
          </SyntaxHighlighter>
        </CodeBlock>

        <TipBox>
          <InfoTitle>Tip</InfoTitle>
          <InfoText>
            Press Enter on an empty line to complete multi-line input.
          </InfoText>
        </TipBox>

        <SubSection>
          <SubSectionTitle>Defining Classes</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
{`>>> grim Counter:
...     init():
...         self.value = 0
...     spell increment():
...         self.value += 1
...         return self.value
...
>>> c = Counter()
>>> c.increment()
1
>>> c.increment()
2`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="mimir">
        <SectionTitle>Mimir Help System</SectionTitle>
        <Paragraph>
          Access comprehensive help directly from the REPL using the Mimir interactive documentation system.
        </Paragraph>

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Command</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <tbody>
            <TableRow>
              <TableCell><InlineCode>mimir</InlineCode></TableCell>
              <TableCell>Launch interactive help browser</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>mimir scry print</InlineCode></TableCell>
              <TableCell>Get documentation for a function</TableCell>
            </TableRow>
            <TableRow>
              <TableCell><InlineCode>mimir categories</InlineCode></TableCell>
              <TableCell>List all help categories</TableCell>
            </TableRow>
          </tbody>
        </Table>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark}>
{`>>> mimir                  // Interactive help browser
>>> mimir scry print      // Function documentation
>>> mimir categories      // List all categories`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section id="tips">
        <SectionTitle>Tips & Tricks</SectionTitle>

        <CardGrid>
          <Card>
            <CardTitle>Quick Testing</CardTitle>
            <CardDescription>
              Test algorithms and code snippets before adding them to your project files.
            </CardDescription>
          </Card>
          <Card>
            <CardTitle>Debug Values</CardTitle>
            <CardDescription>
              Inspect variables and expressions interactively during development.
            </CardDescription>
          </Card>
          <Card>
            <CardTitle>Prototype Classes</CardTitle>
            <CardDescription>
              Develop and refine grimoires in real-time before committing to files.
            </CardDescription>
          </Card>
          <Card>
            <CardTitle>History Search</CardTitle>
            <CardDescription>
              Press Ctrl+R for reverse history search to find previous commands.
            </CardDescription>
          </Card>
        </CardGrid>

        <SubSection>
          <SubSectionTitle>Keyboard Shortcuts</SubSectionTitle>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Shortcut</TableHead>
                <TableHead>Action</TableHead>
              </TableRow>
            </TableHeader>
            <tbody>
              <TableRow>
                <TableCell><InlineCode>Up/Down</InlineCode></TableCell>
                <TableCell>Navigate command history</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>Tab</InlineCode></TableCell>
                <TableCell>Auto-complete</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>Ctrl+R</InlineCode></TableCell>
                <TableCell>Reverse history search</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>Ctrl+C</InlineCode></TableCell>
                <TableCell>Cancel current input</TableCell>
              </TableRow>
              <TableRow>
                <TableCell><InlineCode>Ctrl+D</InlineCode></TableCell>
                <TableCell>Exit REPL</TableCell>
              </TableRow>
            </tbody>
          </Table>
        </SubSection>

        <InfoBox>
          <InfoTitle>Load Modules</InfoTitle>
          <InfoText>
            Import and test your code incrementally using <InlineCode>import</InlineCode> statements
            to load modules from your project.
          </InfoText>
        </InfoBox>
      </Section>
    </DocLayout>
  );
};

export default REPLGuide;
