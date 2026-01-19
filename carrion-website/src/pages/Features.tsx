import React from 'react';
import styled from 'styled-components';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import { Link } from 'react-router-dom';

const Container = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  padding: 6rem 2rem 4rem;
  min-height: 100vh;
`;

const Header = styled.div`
  text-align: center;
  margin-bottom: 5rem;
`;

const Title = styled.h1`
  font-size: 3rem;
  margin-bottom: 1.5rem;
  color: ${({ theme }) => theme.colors.text.primary};
  font-weight: 700;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    font-size: 2.25rem;
  }
`;

const Subtitle = styled.p`
  font-size: 1.25rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  max-width: 650px;
  margin: 0 auto;
  line-height: 1.7;
`;

const Section = styled.section`
  margin-bottom: 5rem;
`;

const SectionHeader = styled.div`
  margin-bottom: 2.5rem;
`;

const SectionTitle = styled.h2`
  font-size: 1.75rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.75rem;
`;

const SectionDescription = styled.p`
  font-size: 1.05rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  max-width: 600px;
`;

const FeatureGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1.5rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: repeat(2, 1fr);
  }

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    grid-template-columns: 1fr;
  }
`;

const FeatureCard = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 12px;
  padding: 1.75rem;
  transition: all 0.2s ease;

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-3px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
  }
`;

const FeatureIcon = styled.div`
  width: 40px;
  height: 40px;
  background: rgba(6, 182, 212, 0.1);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 1rem;
  color: ${({ theme }) => theme.colors.primary};
  font-size: 1.25rem;
`;

const FeatureTitle = styled.h3`
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.5rem;
  font-size: 1.1rem;
  font-weight: 600;
`;

const FeatureDescription = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.6;
  font-size: 0.95rem;
`;

const HighlightSection = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 16px;
  padding: 3rem;
  margin-bottom: 5rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    padding: 2rem;
  }
`;

const HighlightGrid = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
  align-items: center;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: 1fr;
    gap: 2rem;
  }
`;

const HighlightContent = styled.div``;

const HighlightTitle = styled.h2`
  font-size: 1.75rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 1rem;
`;

const HighlightText = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.7;
  margin-bottom: 1.5rem;
`;

const CodeBlock = styled.div`
  border-radius: 10px;
  overflow: hidden;
  background: #1a1b26;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
`;

const ComparisonSection = styled.div`
  margin-bottom: 5rem;
`;

const KeywordTable = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 12px;
  overflow: hidden;
`;

const KeywordRow = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr 2fr;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};

  &:last-child {
    border-bottom: none;
  }

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    grid-template-columns: 1fr 1fr;
  }
`;

const KeywordCell = styled.div<{ $header?: boolean }>`
  padding: 0.875rem 1.25rem;
  font-size: 0.95rem;
  color: ${({ theme, $header }) => $header ? theme.colors.text.primary : theme.colors.text.secondary};
  font-weight: ${({ $header }) => $header ? '600' : 'normal'};
  background: ${({ theme, $header }) => $header ? theme.colors.background.tertiary : 'transparent'};

  code {
    background: rgba(6, 182, 212, 0.1);
    color: ${({ theme }) => theme.colors.primary};
    padding: 0.15rem 0.4rem;
    border-radius: 4px;
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.9em;
  }

  &:nth-child(3) {
    @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
      display: none;
    }
  }
`;

const StatGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1.5rem;
  margin-bottom: 4rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: repeat(2, 1fr);
  }
`;

const StatCard = styled.div`
  text-align: center;
  padding: 1.5rem;
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 12px;
`;

const StatValue = styled.div`
  font-size: 2rem;
  font-weight: 700;
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 0.5rem;
`;

const StatLabel = styled.div`
  font-size: 0.9rem;
  color: ${({ theme }) => theme.colors.text.secondary};
`;

const CTASection = styled.div`
  text-align: center;
  padding: 3rem;
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.1) 0%, rgba(139, 92, 246, 0.1) 100%);
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 16px;
`;

const CTATitle = styled.h2`
  font-size: 1.75rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 1rem;
`;

const CTAText = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 2rem;
  max-width: 500px;
  margin-left: auto;
  margin-right: auto;
`;

const CTAButtons = styled.div`
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
`;

const PrimaryButton = styled(Link)`
  padding: 0.875rem 2rem;
  background: ${({ theme }) => theme.colors.primary};
  color: white;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.2s ease;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 15px rgba(6, 182, 212, 0.4);
  }
`;

const SecondaryButton = styled(Link)`
  padding: 0.875rem 2rem;
  background: transparent;
  color: ${({ theme }) => theme.colors.text.primary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 8px;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.2s ease;

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    color: ${({ theme }) => theme.colors.primary};
  }
`;

const Features: React.FC = () => {
  const carrionExample = `grim Animal:
    init(name):
        self.name = name

    spell speak():
        return f"{self.name} makes a sound"

grim Dog(Animal):
    init(name, breed):
        super.init(name)
        self.breed = breed

    spell speak():
        return f"{self.name} barks"

# Error handling
attempt:
    risky_operation()
ensnare (SpecificError):
    handle_error()
resolve:
    cleanup()`;

  return (
    <Container>
      <Header>
        <Title>Language Features</Title>
        <Subtitle>
          A modern, dynamically-typed language built for clarity and productivity.
          Familiar syntax with powerful capabilities.
        </Subtitle>
      </Header>

      <StatGrid>
        <StatCard>
          <StatValue>v0.1.9</StatValue>
          <StatLabel>Latest Release</StatLabel>
        </StatCard>
        <StatCard>
          <StatValue>Go</StatValue>
          <StatLabel>Built With</StatLabel>
        </StatCard>
        <StatCard>
          <StatValue>3</StatValue>
          <StatLabel>Platforms</StatLabel>
        </StatCard>
        <StatCard>
          <StatValue>MIT</StatValue>
          <StatLabel>License</StatLabel>
        </StatCard>
      </StatGrid>

      <HighlightSection>
        <HighlightGrid>
          <HighlightContent>
            <HighlightTitle>Clean, Expressive Syntax</HighlightTitle>
            <HighlightText>
              Carrion combines Python's readability with unique keywords that make code
              intent clear. Object-oriented programming feels natural with grimoires (classes)
              and spells (methods).
            </HighlightText>
            <HighlightText>
              Full support for inheritance, abstract classes, error handling, and modern
              language features you expect from a production-ready language.
            </HighlightText>
          </HighlightContent>
          <CodeBlock>
            <SyntaxHighlighter
              language="python"
              style={atomOneDark}
              customStyle={{ margin: 0, padding: '1.5rem', fontSize: '0.9rem' }}
            >
              {carrionExample}
            </SyntaxHighlighter>
          </CodeBlock>
        </HighlightGrid>
      </HighlightSection>

      <Section>
        <SectionHeader>
          <SectionTitle>Core Features</SectionTitle>
          <SectionDescription>
            Everything you need to build applications, from simple scripts to complex systems.
          </SectionDescription>
        </SectionHeader>

        <FeatureGrid>
          <FeatureCard>
            <FeatureIcon>OOP</FeatureIcon>
            <FeatureTitle>Object-Oriented</FeatureTitle>
            <FeatureDescription>
              Classes, inheritance, polymorphism, and encapsulation. Build modular,
              maintainable code with familiar OOP patterns.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Dyn</FeatureIcon>
            <FeatureTitle>Dynamic Typing</FeatureTitle>
            <FeatureDescription>
              Rapid development with dynamic typing and comprehensive runtime
              error reporting with file and line information.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Err</FeatureIcon>
            <FeatureTitle>Error Handling</FeatureTitle>
            <FeatureDescription>
              Structured error handling with attempt/ensnare/resolve blocks.
              Catch specific errors or handle them broadly.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Lib</FeatureIcon>
            <FeatureTitle>Standard Library</FeatureTitle>
            <FeatureDescription>
              Munin standard library with collections, file I/O, math functions,
              OS interface, and more out of the box.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Mod</FeatureIcon>
            <FeatureTitle>Module System</FeatureTitle>
            <FeatureDescription>
              Import system with support for local files, packages, and selective
              imports for clean namespace management.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Str</FeatureIcon>
            <FeatureTitle>String Interpolation</FeatureTitle>
            <FeatureDescription>
              F-strings and multiple string formats for flexible text formatting
              and expression embedding.
            </FeatureDescription>
          </FeatureCard>
        </FeatureGrid>
      </Section>

      <Section>
        <SectionHeader>
          <SectionTitle>Developer Tooling</SectionTitle>
          <SectionDescription>
            Integrated tools that make development faster and more enjoyable.
          </SectionDescription>
        </SectionHeader>

        <FeatureGrid>
          <FeatureCard>
            <FeatureIcon>Test</FeatureIcon>
            <FeatureTitle>Sindri Testing</FeatureTitle>
            <FeatureDescription>
              Built-in testing framework with automatic test discovery, colored output,
              and the check() assertion function.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Doc</FeatureIcon>
            <FeatureTitle>Mimir Documentation</FeatureTitle>
            <FeatureDescription>
              Interactive help system accessible from the REPL. Look up any function
              or module with instant documentation.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Pkg</FeatureIcon>
            <FeatureTitle>Bifrost Packages</FeatureTitle>
            <FeatureDescription>
              Package manager with Git integration. Install, manage, and distribute
              packages with dependency resolution.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>CLI</FeatureIcon>
            <FeatureTitle>Interactive REPL</FeatureTitle>
            <FeatureDescription>
              Feature-rich REPL with tab completion, command history, and smart
              output display for rapid prototyping.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Abs</FeatureIcon>
            <FeatureTitle>Abstract Classes</FeatureTitle>
            <FeatureDescription>
              Define interfaces with arcane grimoires and abstract methods for
              robust inheritance patterns.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>Pat</FeatureIcon>
            <FeatureTitle>Pattern Matching</FeatureTitle>
            <FeatureDescription>
              Modern match/case statements for clean conditional logic and
              structured data handling.
            </FeatureDescription>
          </FeatureCard>
        </FeatureGrid>
      </Section>

      <ComparisonSection>
        <SectionHeader>
          <SectionTitle>Syntax Reference</SectionTitle>
          <SectionDescription>
            Carrion's keywords compared to their equivalents in other languages.
          </SectionDescription>
        </SectionHeader>

        <KeywordTable>
          <KeywordRow>
            <KeywordCell $header>Carrion</KeywordCell>
            <KeywordCell $header>Equivalent</KeywordCell>
            <KeywordCell $header>Description</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>grim</code></KeywordCell>
            <KeywordCell><code>class</code></KeywordCell>
            <KeywordCell>Define a class</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>spell</code></KeywordCell>
            <KeywordCell><code>def / function</code></KeywordCell>
            <KeywordCell>Define a method or function</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>attempt</code></KeywordCell>
            <KeywordCell><code>try</code></KeywordCell>
            <KeywordCell>Begin error handling block</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>ensnare</code></KeywordCell>
            <KeywordCell><code>catch / except</code></KeywordCell>
            <KeywordCell>Handle caught errors</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>resolve</code></KeywordCell>
            <KeywordCell><code>finally</code></KeywordCell>
            <KeywordCell>Always execute cleanup</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>otherwise</code></KeywordCell>
            <KeywordCell><code>elif / else if</code></KeywordCell>
            <KeywordCell>Conditional else-if</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>skip</code></KeywordCell>
            <KeywordCell><code>continue</code></KeywordCell>
            <KeywordCell>Skip to next loop iteration</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>stop</code></KeywordCell>
            <KeywordCell><code>break</code></KeywordCell>
            <KeywordCell>Exit loop</KeywordCell>
          </KeywordRow>
          <KeywordRow>
            <KeywordCell><code>arcane grim</code></KeywordCell>
            <KeywordCell><code>abstract class</code></KeywordCell>
            <KeywordCell>Define abstract base class</KeywordCell>
          </KeywordRow>
        </KeywordTable>
      </ComparisonSection>

      <Section>
        <SectionHeader>
          <SectionTitle>Why Carrion?</SectionTitle>
          <SectionDescription>
            Built for developers who value clarity and productivity.
          </SectionDescription>
        </SectionHeader>

        <FeatureGrid>
          <FeatureCard>
            <FeatureIcon>01</FeatureIcon>
            <FeatureTitle>Easy to Learn</FeatureTitle>
            <FeatureDescription>
              Familiar Python-like syntax means you can be productive in hours,
              not weeks. Great for beginners and experienced developers alike.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>02</FeatureIcon>
            <FeatureTitle>Modern Design</FeatureTitle>
            <FeatureDescription>
              Built from scratch with modern language features. No legacy baggage,
              just clean, consistent design decisions.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>03</FeatureIcon>
            <FeatureTitle>Cross Platform</FeatureTitle>
            <FeatureDescription>
              Runs on Linux, macOS, and Windows. Available as binaries, from source,
              or via Docker containers.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>04</FeatureIcon>
            <FeatureTitle>Go Performance</FeatureTitle>
            <FeatureDescription>
              Built in Go for reliable performance and excellent concurrency
              capabilities under the hood.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>05</FeatureIcon>
            <FeatureTitle>Active Development</FeatureTitle>
            <FeatureDescription>
              Regular releases with new features. Roadmap includes JIT compilation,
              VM implementation, and type annotations.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>06</FeatureIcon>
            <FeatureTitle>Open Source</FeatureTitle>
            <FeatureDescription>
              MIT licensed and fully open source. Contribute, fork, or use it
              however you need for your projects.
            </FeatureDescription>
          </FeatureCard>
        </FeatureGrid>
      </Section>

      <CTASection>
        <CTATitle>Ready to Get Started?</CTATitle>
        <CTAText>
          Install Carrion and start building in minutes. Check out the documentation
          for tutorials and examples.
        </CTAText>
        <CTAButtons>
          <PrimaryButton to="/docs/installation">Installation Guide</PrimaryButton>
          <SecondaryButton to="/docs/quick-start">Quick Start</SecondaryButton>
        </CTAButtons>
      </CTASection>
    </Container>
  );
};

export default Features;
