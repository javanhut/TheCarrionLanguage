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

const Card = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 16px;
  padding: 2.5rem;
  margin-bottom: 2rem;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-4px);
    box-shadow: ${({ theme }) => theme.shadows.large},
                0 0 30px rgba(0, 204, 153, 0.15);
  }
`;

const Grid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
  margin: 2rem 0;
`;

const FeatureCard = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 12px;
  padding: 2rem;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateX(5px);
    box-shadow: -5px 0 20px rgba(0, 204, 153, 0.1);
  }
`;

const FeatureTitle = styled.h4`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1rem;
  font-size: 1.4rem;
  font-weight: 600;
`;

const Text = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.8;
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
`;

const CodeBlock = styled.div`
  margin: 2rem 0;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: ${({ theme }) => theme.shadows.medium};
`;

const InfoBox = styled.div`
  background: rgba(0, 204, 153, 0.1);
  border-left: 4px solid ${({ theme }) => theme.colors.primary};
  border-radius: 8px;
  padding: 1.5rem 2rem;
  margin: 2rem 0;
  
  p {
    margin: 0;
    color: ${({ theme }) => theme.colors.text.primary};
  }

  strong {
    color: ${({ theme }) => theme.colors.primary};
  }

  code {
    background: ${({ theme }) => theme.colors.code};
    padding: 0.3rem 0.6rem;
    border-radius: 4px;
  }
`;

const REPLGuide: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Interactive REPL Guide</Title>
        <Subtitle>
          Master the Carrion interactive shell for rapid development and experimentation
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>What is the REPL?</SectionTitle>
        <Card>
          <Text>
            The Carrion REPL (Read-Eval-Print Loop) provides an interactive environment where you can 
            execute code, test features, and experiment with language concepts in real-time.
          </Text>
          <InfoBox>
            <p><strong>Quick Start:</strong> Simply run <code>carrion</code> in your terminal to launch the REPL.</p>
          </InfoBox>
        </Card>
      </Section>

      <Section>
        <SectionTitle>Key Features</SectionTitle>
        <Grid>
          <FeatureCard>
            <FeatureTitle>Clean Output</FeatureTitle>
            <Text>
              Assignment statements and definitions don't clutter your screen - only expression 
              results are displayed.
            </Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>Smart Completion</FeatureTitle>
            <Text>
              Tab completion for keywords, functions, variables, and grimoire methods speeds 
              up your workflow.
            </Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>Command History</FeatureTitle>
            <Text>
              Navigate through previous commands with arrow keys. History persists between sessions.
            </Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>Multi-line Input</FeatureTitle>
            <Text>
              Automatic detection of incomplete statements lets you write functions and classes naturally.
            </Text>
          </FeatureCard>
        </Grid>
      </Section>

      <Section>
        <SectionTitle>Basic Usage</SectionTitle>
        
        <SubSectionTitle>Starting the REPL</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter 
            language="bash" 
            style={atomOneDark} 
            customStyle={{ margin: 0, borderRadius: '12px', fontSize: '1rem', padding: '1.5rem' }}
          >
            {`carrion`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Simple Expressions</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter 
            language="python" 
            style={atomOneDark} 
            customStyle={{ margin: 0, borderRadius: '12px', fontSize: '1rem', padding: '1.5rem' }}
          >
            {`>>> 2 + 2
4
>>> "Hello" + " " + "World"
Hello World
>>> [1, 2, 3] + [4, 5]
[1, 2, 3, 4, 5]`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Variable Assignment</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter 
            language="python" 
            style={atomOneDark} 
            customStyle={{ margin: 0, borderRadius: '12px', fontSize: '1rem', padding: '1.5rem' }}
          >
            {`>>> name = "Alice"         // No output
>>> age = 30                // No output  
>>> print(f"{name} is {age}")
Alice is 30`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Multi-line Mode</SectionTitle>
        <Card>
          <Text>
            The REPL automatically enters multi-line mode for function and class definitions:
          </Text>
          <CodeBlock>
            <SyntaxHighlighter 
              language="python" 
              style={atomOneDark} 
              customStyle={{ margin: 0, borderRadius: '12px', fontSize: '1rem', padding: '1.5rem' }}
            >
              {`>>> spell factorial(n):
...     if n <= 1:
...         return 1
...     return n * factorial(n - 1)
... 
>>> factorial(5)
120`}
            </SyntaxHighlighter>
          </CodeBlock>
          <InfoBox>
            <p><strong>Tip:</strong> Press Enter on an empty line to complete multi-line input.</p>
          </InfoBox>
        </Card>
      </Section>

      <Section>
        <SectionTitle>Mimir Integration</SectionTitle>
        <Card>
          <Text>
            Access comprehensive help directly from the REPL:
          </Text>
          <CodeBlock>
            <SyntaxHighlighter 
              language="python" 
              style={atomOneDark} 
              customStyle={{ margin: 0, borderRadius: '12px', fontSize: '1rem', padding: '1.5rem' }}
            >
              {`>>> mimir                  // Interactive help browser
>>> mimir scry print      // Function documentation
>>> mimir categories      // List all categories`}
            </SyntaxHighlighter>
          </CodeBlock>
        </Card>
      </Section>

      <Section>
        <SectionTitle>Tips & Tricks</SectionTitle>
        <Grid>
          <FeatureCard>
            <FeatureTitle>Quick Testing</FeatureTitle>
            <Text>Test algorithms before adding them to your project</Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>Debug Values</FeatureTitle>
            <Text>Inspect variables and expressions interactively</Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>Prototype Classes</FeatureTitle>
            <Text>Develop and refine grimoires in real-time</Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>Load Modules</FeatureTitle>
            <Text>Import and test your code incrementally</Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>History Search</FeatureTitle>
            <Text>Press Ctrl+R for reverse history search</Text>
          </FeatureCard>
          <FeatureCard>
            <FeatureTitle>Explore Standard Library</FeatureTitle>
            <Text>Test Munin grimoires and builtin functions on the fly</Text>
          </FeatureCard>
        </Grid>
      </Section>
    </Container>
  );
};

export default REPLGuide;
