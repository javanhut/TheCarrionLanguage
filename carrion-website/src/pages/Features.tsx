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
`;

const Title = styled.h1`
  font-size: 3.5rem;
  margin-bottom: 1rem;
  background: ${({ theme }) => theme.gradients.primary};
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
`;

const Subtitle = styled.p`
  font-size: 1.4rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  max-width: 600px;
  margin: 0 auto;
  line-height: 1.6;
`;

const Section = styled.section`
  margin-bottom: 4rem;
`;

const SectionTitle = styled.h2`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 2rem;
  font-size: 2.5rem;
  text-align: center;
`;

const SubSectionTitle = styled.h3`
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 1.5rem;
  font-size: 1.8rem;
`;

const FeatureGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
  margin-bottom: 3rem;
`;

const FeatureCard = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 15px;
  padding: 2rem;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-5px);
    box-shadow: ${({ theme }) => theme.shadows.large};
  }
`;

const FeatureIcon = styled.div`
  font-size: 3rem;
  margin-bottom: 1rem;
`;

const FeatureTitle = styled.h3`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1rem;
  font-size: 1.4rem;
`;

const FeatureDescription = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.6;
`;

const ComparisonTable = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 15px;
  overflow: hidden;
  margin: 2rem 0;
`;

const TableRow = styled.div<{ header?: boolean }>`
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  background: ${({ header, theme }) => header ? theme.colors.background.tertiary : 'transparent'};
  
  &:not(:last-child) {
    border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  }
`;

const TableCell = styled.div<{ header?: boolean }>`
  padding: 1rem 1.5rem;
  color: ${({ header, theme }) => header ? theme.colors.primary : theme.colors.text.primary};
  font-weight: ${({ header }) => header ? '600' : 'normal'};
  border-right: 1px solid ${({ theme }) => theme.colors.border};

  &:last-child {
    border-right: none;
  }

  code {
    background: ${({ theme }) => theme.colors.background.primary};
    padding: 0.2rem 0.4rem;
    border-radius: 4px;
    font-size: 0.9rem;
  }
`;

const CodeComparison = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin: 2rem 0;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: 1fr;
  }
`;

const CodeBlock = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  overflow: hidden;
`;

const CodeHeader = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  padding: 1rem 1.5rem;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  font-weight: 600;
  color: ${({ theme }) => theme.colors.primary};
`;

const KeywordList = styled.ul`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  list-style: none;
  margin: 2rem 0;
`;

const KeywordItem = styled.li`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 8px;
  padding: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const Keyword = styled.code`
  color: ${({ theme }) => theme.colors.primary};
  font-weight: 600;
`;

const Arrow = styled.span`
  color: ${({ theme }) => theme.colors.text.secondary};
  margin: 0 0.5rem;
`;

const Translation = styled.span`
  color: ${({ theme }) => theme.colors.text.primary};
`;

const Features: React.FC = () => {
  const pythonExample = `# Python
class Animal:
    def __init__(self, name):
        self.name = name
    
    def speak(self):
        return f"{self.name} makes a sound"

class Dog(Animal):
    def __init__(self, name, breed):
        super().__init__(name)
        self.breed = breed
    
    def speak(self):
        return f"{self.name} barks"

# Error handling
try:
    risky_operation()
except SpecificError:
    handle_error()
finally:
    cleanup()`;

  const carrionExample = `# Carrion
grim Animal:
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
          Discover what makes Carrion unique - a powerful programming language 
          that combines familiar syntax with magical terminology and advanced features.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Core Language Features</SectionTitle>
        <FeatureGrid>
          <FeatureCard>
            <FeatureIcon>🧙‍♂️</FeatureIcon>
            <FeatureTitle>Magical Syntax</FeatureTitle>
            <FeatureDescription>
              Transform mundane programming concepts into magical terminology. 
              Classes become "grimoires", methods become "spells", and error handling 
              uses mystical keywords like "attempt" and "ensnare".
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🏗️</FeatureIcon>
            <FeatureTitle>Object-Oriented Programming</FeatureTitle>
            <FeatureDescription>
              Full OOP support with classes, inheritance, encapsulation, polymorphism, 
              and abstraction. Create complex hierarchies with clean, readable syntax.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🔥</FeatureIcon>
            <FeatureTitle>Dynamic Typing</FeatureTitle>
            <FeatureDescription>
              Write code faster with dynamic typing while maintaining type safety 
              through runtime checks and comprehensive error reporting.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>⚡</FeatureIcon>
            <FeatureTitle>Built in Go</FeatureTitle>
            <FeatureDescription>
              Leverages Go's performance and reliability for the interpreter, 
              providing fast execution and excellent concurrency capabilities.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>📚</FeatureIcon>
            <FeatureTitle>Rich Standard Library</FeatureTitle>
            <FeatureDescription>
              "Munin" standard library provides enhanced collections, file I/O, 
              math functions, OS interface, and debugging utilities out of the box.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🛡️</FeatureIcon>
            <FeatureTitle>Robust Error Handling</FeatureTitle>
            <FeatureDescription>
              Comprehensive error handling with magical keywords and detailed 
              error reporting that pinpoints issues with file and line information.
            </FeatureDescription>
          </FeatureCard>
        </FeatureGrid>
      </Section>

      <Section>
        <SectionTitle>Magical Keywords Translation</SectionTitle>
        <p style={{ textAlign: 'center', marginBottom: '2rem', color: '#8892b0' }}>
          Carrion transforms traditional programming concepts into magical terminology:
        </p>
        <KeywordList>
          <KeywordItem>
            <Keyword>grim</Keyword>
            <Arrow>→</Arrow>
            <Translation>class</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>spell</Keyword>
            <Arrow>→</Arrow>
            <Translation>method/function</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>attempt</Keyword>
            <Arrow>→</Arrow>
            <Translation>try</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>ensnare</Keyword>
            <Arrow>→</Arrow>
            <Translation>catch/except</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>resolve</Keyword>
            <Arrow>→</Arrow>
            <Translation>finally</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>skip</Keyword>
            <Arrow>→</Arrow>
            <Translation>continue</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>stop</Keyword>
            <Arrow>→</Arrow>
            <Translation>break</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>otherwise</Keyword>
            <Arrow>→</Arrow>
            <Translation>elif</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>arcane grim</Keyword>
            <Arrow>→</Arrow>
            <Translation>abstract class</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>Munin</Keyword>
            <Arrow>→</Arrow>
            <Translation>Standard Library</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>Mimir</Keyword>
            <Arrow>→</Arrow>
            <Translation>Interactive Help</Translation>
          </KeywordItem>
          <KeywordItem>
            <Keyword>.crl</Keyword>
            <Arrow>→</Arrow>
            <Translation>File Extension</Translation>
          </KeywordItem>
        </KeywordList>
      </Section>

      <Section>
        <SectionTitle>Carrion vs Python</SectionTitle>
        <p style={{ textAlign: 'center', marginBottom: '2rem', color: '#8892b0' }}>
          While Carrion shares Python's accessibility, it brings unique features and syntax:
        </p>
        
        <ComparisonTable>
          <TableRow header>
            <TableCell header>Feature</TableCell>
            <TableCell header>Python</TableCell>
            <TableCell header>Carrion</TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>Comments</strong></TableCell>
            <TableCell><code># Single line</code></TableCell>
            <TableCell><code># Single line</code><br/><code>```Multi-line```</code></TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>Classes</strong></TableCell>
            <TableCell><code>class MyClass:</code></TableCell>
            <TableCell><code>grim MyGrimoire:</code></TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>Methods</strong></TableCell>
            <TableCell><code>def method(self):</code></TableCell>
            <TableCell><code>spell method():</code></TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>Error Handling</strong></TableCell>
            <TableCell><code>try/except/finally</code></TableCell>
            <TableCell><code>attempt/ensnare/resolve</code></TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>Loop Control</strong></TableCell>
            <TableCell><code>continue/break</code></TableCell>
            <TableCell><code>skip/stop</code></TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>Elif Statement</strong></TableCell>
            <TableCell><code>elif condition:</code></TableCell>
            <TableCell><code>otherwise condition:</code></TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>Increment</strong></TableCell>
            <TableCell><code>i += 1</code></TableCell>
            <TableCell><code>i++, ++i, i += 1</code></TableCell>
          </TableRow>
          <TableRow>
            <TableCell><strong>File Extension</strong></TableCell>
            <TableCell><code>.py</code></TableCell>
            <TableCell><code>.crl</code></TableCell>
          </TableRow>
        </ComparisonTable>
      </Section>

      <Section>
        <SectionTitle>Code Comparison</SectionTitle>
        <CodeComparison>
          <CodeBlock>
            <CodeHeader>🐍 Python</CodeHeader>
            <SyntaxHighlighter
              language="python"
              style={atomOneDark}
              customStyle={{ margin: 0, background: 'transparent' }}
            >
              {pythonExample}
            </SyntaxHighlighter>
          </CodeBlock>
          <CodeBlock>
            <CodeHeader>🐦‍⬛ Carrion</CodeHeader>
            <SyntaxHighlighter
              language="python"
              style={atomOneDark}
              customStyle={{ margin: 0, background: 'transparent' }}
            >
              {carrionExample}
            </SyntaxHighlighter>
          </CodeBlock>
        </CodeComparison>
      </Section>

      <Section>
        <SubSectionTitle>Advanced Features</SubSectionTitle>
        <FeatureGrid>
          <FeatureCard>
            <FeatureIcon>🎯</FeatureIcon>
            <FeatureTitle>Pattern Matching</FeatureTitle>
            <FeatureDescription>
              Modern pattern matching with <code>match/case</code> statements 
              for clean conditional logic and data structure handling.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🔮</FeatureIcon>
            <FeatureTitle>Abstract Classes</FeatureTitle>
            <FeatureDescription>
              Create abstract base classes with <code>arcane grim</code> and 
              <code>@arcanespell</code> decorators for robust inheritance patterns.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🌟</FeatureIcon>
            <FeatureTitle>String Interpolation</FeatureTitle>
            <FeatureDescription>
              Multiple string formats including f-strings and i-strings 
              for flexible text formatting and expression embedding.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>📊</FeatureIcon>
            <FeatureTitle>Rich Data Types</FeatureTitle>
            <FeatureDescription>
              Support for integers, floats, strings, booleans, arrays, hashes, 
              and tuples with comprehensive built-in methods.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🔧</FeatureIcon>
            <FeatureTitle>Interactive REPL</FeatureTitle>
            <FeatureDescription>
              Feature-rich REPL with tab completion, command history, 
              built-in help system ("Mimir"), and debugging utilities.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>📦</FeatureIcon>
            <FeatureTitle>Enhanced Collections</FeatureTitle>
            <FeatureDescription>
              Extended collection classes with additional methods like 
              <code>.contains()</code>, <code>.slice()</code>, and mathematical operations.
            </FeatureDescription>
          </FeatureCard>
        </FeatureGrid>
      </Section>

      <Section>
        <SubSectionTitle>Why Choose Carrion?</SubSectionTitle>
        <FeatureGrid>
          <FeatureCard>
            <FeatureIcon>🎓</FeatureIcon>
            <FeatureTitle>Beginner Friendly</FeatureTitle>
            <FeatureDescription>
              Familiar Python-like syntax makes it easy to learn, while the magical 
              terminology creates an engaging and memorable learning experience.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🚀</FeatureIcon>
            <FeatureTitle>Modern Design</FeatureTitle>
            <FeatureDescription>
              Built from the ground up with modern language features, avoiding 
              historical baggage while incorporating lessons learned from other languages.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>🌍</FeatureIcon>
            <FeatureTitle>Cross Platform</FeatureTitle>
            <FeatureDescription>
              Runs on Linux, macOS, and Windows. Available as binaries, source code, 
              and Docker containers for maximum flexibility.
            </FeatureDescription>
          </FeatureCard>

          <FeatureCard>
            <FeatureIcon>📈</FeatureIcon>
            <FeatureTitle>Active Development</FeatureTitle>
            <FeatureDescription>
              Regular updates and improvements with a roadmap including JIT compilation, 
              VM implementation, and enhanced type systems.
            </FeatureDescription>
          </FeatureCard>
        </FeatureGrid>
      </Section>
    </Container>
  );
};

export default Features;