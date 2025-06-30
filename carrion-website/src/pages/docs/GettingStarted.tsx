import React from 'react';
import styled from 'styled-components';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import { Link } from 'react-router-dom';

const Container = styled.div`
  max-width: 1000px;
  margin: 0 auto;
  padding: 6rem 2rem 4rem;
  min-height: 100vh;
`;

const Header = styled.div`
  text-align: center;
  margin-bottom: 3rem;
`;

const Title = styled.h1`
  font-size: 3rem;
  margin-bottom: 1rem;
  background: ${({ theme }) => theme.gradients.primary};
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
`;

const Subtitle = styled.p`
  font-size: 1.3rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  max-width: 600px;
  margin: 0 auto;
`;

const Section = styled.section`
  margin-bottom: 3rem;
`;

const SectionTitle = styled.h2`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1.5rem;
  font-size: 2rem;
`;

const SubSectionTitle = styled.h3`
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 1rem;
  font-size: 1.5rem;
`;

const StepCard = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 15px;
  padding: 2rem;
  margin-bottom: 2rem;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-2px);
  }
`;

const StepNumber = styled.div`
  background: ${({ theme }) => theme.colors.primary};
  color: white;
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  margin-bottom: 1rem;
`;

const CodeBlock = styled.div`
  margin: 1rem 0;
`;

const QuickNavCard = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  padding: 1.5rem;
  margin: 2rem 0;
`;

const NavList = styled.ul`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  list-style: none;
  margin: 1rem 0;
`;

const NavItem = styled.li`
  a {
    color: ${({ theme }) => theme.colors.primary};
    text-decoration: none;
    font-weight: 500;
    
    &:hover {
      text-decoration: underline;
    }
  }
`;

const InfoBox = styled.div`
  background: rgba(0, 204, 153, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.primary};
  border-radius: 8px;
  padding: 1rem;
  margin: 1rem 0;
`;

const WarningBox = styled.div`
  background: rgba(255, 204, 0, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.warning};
  border-radius: 8px;
  padding: 1rem;
  margin: 1rem 0;
`;

const GettingStarted: React.FC = () => {
  const helloWorldExample = `# hello.crl
print("Hello, magical world of Carrion!")`;

  const basicVariablesExample = `# Variables in Carrion
name = "Wizard"
age = 25
height = 5.9
is_magical = True

# String interpolation
greeting = f"Hello, {name}! You are {age} years old."
print(greeting)`;

  const firstGrimExample = `# first_grim.crl
grim MagicalCreature:
    init(name, element):
        self.name = name
        self.element = element
        self.spells = []
    
    spell introduce():
        return f"I am {self.name}, master of {self.element} magic!"
    
    spell learn_spell(spell_name):
        self.spells.append(spell_name)
        print(f"{self.name} learned {spell_name}!")
    
    spell cast_spell(spell_name):
        if spell_name in self.spells:
            print(f"{self.name} casts {spell_name}!")
        else:
            print(f"{self.name} doesn't know {spell_name}")

# Create a magical creature
wizard = MagicalCreature("Gandalf", "Light")
print(wizard.introduce())

wizard.learn_spell("Lightning Bolt")
wizard.learn_spell("Healing Light")
wizard.cast_spell("Lightning Bolt")`;

  const errorHandlingExample = `# error_handling.crl
spell divide_numbers(a, b):
    attempt:
        result = a / b
        return result
    ensnare (ZeroDivisionError):
        print("Cannot divide by zero!")
        return None
    resolve:
        print("Division operation completed")

# Test error handling
print(divide_numbers(10, 2))  # Works fine
print(divide_numbers(10, 0))  # Handles error`;

  const replExample = `$ carrion
🐦‍⬛ Welcome to Carrion REPL v0.1.6
Type 'mimir' for help, 'quit' to exit

>>> print("Hello from REPL!")
Hello from REPL!

>>> name = "Coder"
>>> f"Welcome {name}!"
'Welcome Coder!'

>>> # Check version
>>> version()
Carrion v0.1.6

>>> # Get help
>>> mimir
Welcome to Mimir - The Carrion Help System
...

>>> quit`;

  return (
    <Container>
      <Header>
        <Title>Getting Started with Carrion</Title>
        <Subtitle>
          Welcome to the mystical world of Carrion! This guide will help you cast your first spells 
          and embark on your magical programming journey.
        </Subtitle>
      </Header>

      <QuickNavCard>
        <h3>🗺️ Quick Navigation</h3>
        <NavList>
          <NavItem><Link to="/docs/installation">📦 Installation Guide</Link></NavItem>
          <NavItem><Link to="/docs/language-reference">📚 Language Reference</Link></NavItem>
          <NavItem><Link to="/playground">🎮 Try Online Playground</Link></NavItem>
          <NavItem><Link to="/docs/quick-start">⚡ Quick Start Tutorial</Link></NavItem>
          <NavItem><Link to="/docs/standard-library">🔮 Standard Library</Link></NavItem>
          <NavItem><Link to="/community">💬 Join Community</Link></NavItem>
        </NavList>
      </QuickNavCard>

      <Section>
        <SectionTitle>Step 1: Install Carrion</SectionTitle>
        
        <StepCard>
          <StepNumber>1</StepNumber>
          <SubSectionTitle>Choose Your Installation Method</SubSectionTitle>
          
          <h4>🚀 Quick Install (Recommended)</h4>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Download for your OS from releases page
curl -L "https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.6/carrion_linux_amd64.tar.gz" -o carrion.tar.gz
tar -xzf carrion.tar.gz
sudo cp carrion /usr/local/bin/

# Test installation
carrion
# In REPL, type: version()`}
            </SyntaxHighlighter>
          </CodeBlock>

          <h4>🐳 Using Docker</h4>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`docker pull javanhut/carrionlanguage:latest
docker run -it javanhut/carrionlanguage:latest`}
            </SyntaxHighlighter>
          </CodeBlock>

          <InfoBox>
            <p><strong>Need help installing?</strong> Check out our detailed <Link to="/docs/installation">Installation Guide</Link> for step-by-step instructions for all operating systems.</p>
          </InfoBox>
        </StepCard>
      </Section>

      <Section>
        <SectionTitle>Step 2: Your First Carrion Program</SectionTitle>
        
        <StepCard>
          <StepNumber>2</StepNumber>
          <SubSectionTitle>Hello, Magical World!</SubSectionTitle>
          
          <p>Let's start with the traditional "Hello World" program in Carrion:</p>
          
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
              {helloWorldExample}
            </SyntaxHighlighter>
          </CodeBlock>

          <p>Save this as <code>hello.crl</code> and run it:</p>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`carrion hello.crl`}
            </SyntaxHighlighter>
          </CodeBlock>

          <p>You should see: <code>Hello, magical world of Carrion!</code></p>
        </StepCard>
      </Section>

      <Section>
        <SectionTitle>Step 3: Learning the Basics</SectionTitle>
        
        <StepCard>
          <StepNumber>3</StepNumber>
          <SubSectionTitle>Variables and Data Types</SubSectionTitle>
          
          <p>Carrion supports dynamic typing with familiar syntax:</p>
          
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
              {basicVariablesExample}
            </SyntaxHighlighter>
          </CodeBlock>

          <h4>📊 Supported Data Types:</h4>
          <ul>
            <li><strong>Integer:</strong> <code>42</code>, <code>-10</code></li>
            <li><strong>Float:</strong> <code>3.14</code>, <code>-2.5</code></li>
            <li><strong>String:</strong> <code>"Hello"</code>, <code>'World'</code>, <code>f"Formatted &#123;var&#125;"</code></li>
            <li><strong>Boolean:</strong> <code>True</code>, <code>False</code></li>
            <li><strong>Array:</strong> <code>[1, 2, 3]</code></li>
            <li><strong>Hash:</strong> <code>&#123;"key": "value"&#125;</code></li>
            <li><strong>None:</strong> <code>None</code></li>
          </ul>
        </StepCard>
      </Section>

      <Section>
        <SectionTitle>Step 4: Your First Grimoire (Class)</SectionTitle>
        
        <StepCard>
          <StepNumber>4</StepNumber>
          <SubSectionTitle>Creating Magical Objects</SubSectionTitle>
          
          <p>In Carrion, classes are called "grimoires" and methods are "spells". Let's create your first magical creature:</p>
          
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
              {firstGrimExample}
            </SyntaxHighlighter>
          </CodeBlock>

          <InfoBox>
            <p><strong>Magical Keywords:</strong></p>
            <ul>
              <li><code>grim</code> = class</li>
              <li><code>spell</code> = method/function</li>
              <li><code>init</code> = constructor</li>
            </ul>
          </InfoBox>
        </StepCard>
      </Section>

      <Section>
        <SectionTitle>Step 5: Error Handling Magic</SectionTitle>
        
        <StepCard>
          <StepNumber>5</StepNumber>
          <SubSectionTitle>Attempt, Ensnare, and Resolve</SubSectionTitle>
          
          <p>Carrion uses magical keywords for error handling:</p>
          
          <CodeBlock>
            <SyntaxHighlighter language="python" style={atomOneDark}>
              {errorHandlingExample}
            </SyntaxHighlighter>
          </CodeBlock>

          <InfoBox>
            <p><strong>Error Handling Keywords:</strong></p>
            <ul>
              <li><code>attempt</code> = try</li>
              <li><code>ensnare</code> = catch/except</li>
              <li><code>resolve</code> = finally</li>
            </ul>
          </InfoBox>
        </StepCard>
      </Section>

      <Section>
        <SectionTitle>Step 6: Interactive Programming with REPL</SectionTitle>
        
        <StepCard>
          <StepNumber>6</StepNumber>
          <SubSectionTitle>The Carrion REPL Experience</SubSectionTitle>
          
          <p>Carrion includes a powerful REPL (Read-Eval-Print Loop) for interactive programming:</p>
          
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
              {replExample}
            </SyntaxHighlighter>
          </CodeBlock>

          <h4>🔧 REPL Features:</h4>
          <ul>
            <li><strong>Tab Completion:</strong> Press Tab to auto-complete keywords and functions</li>
            <li><strong>Command History:</strong> Use arrow keys to navigate previous commands</li>
            <li><strong>Help System:</strong> Type <code>mimir</code> for interactive help</li>
            <li><strong>Built-in Commands:</strong> <code>version()</code>, <code>help()</code>, <code>modules()</code></li>
            <li><strong>Multi-line Input:</strong> Supports complex expressions and functions</li>
          </ul>
        </StepCard>
      </Section>

      <Section>
        <SectionTitle>🎯 What's Next?</SectionTitle>
        
        <p>Congratulations! You've taken your first steps into the magical world of Carrion. Here's where to go next:</p>

        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(250px, 1fr))', gap: '1.5rem', margin: '2rem 0' }}>
          <StepCard>
            <h3>📖 Dive Deeper</h3>
            <p>Explore the complete language features:</p>
            <ul>
              <li><Link to="/docs/language-reference">Language Reference</Link></li>
              <li><Link to="/docs/grimoires">Classes & Inheritance</Link></li>
              <li><Link to="/docs/control-flow">Control Flow</Link></li>
              <li><Link to="/docs/operators">Operators</Link></li>
            </ul>
          </StepCard>

          <StepCard>
            <h3>🔮 Standard Library</h3>
            <p>Discover the power of Munin:</p>
            <ul>
              <li><Link to="/docs/standard-library">Munin Overview</Link></li>
              <li><Link to="/docs/builtin-functions">Built-in Functions</Link></li>
              <li><Link to="/docs/modules">Modules System</Link></li>
            </ul>
          </StepCard>

          <StepCard>
            <h3>🎮 Practice & Play</h3>
            <p>Hone your magical skills:</p>
            <ul>
              <li><Link to="/playground">Online Playground</Link></li>
              <li><Link to="/docs/quick-start">Quick Start Tutorial</Link></li>
              <li><Link to="/community">Join the Community</Link></li>
            </ul>
          </StepCard>
        </div>

        <WarningBox>
          <p><strong>Note:</strong> Carrion is currently in active development (v0.1.6). Some features may be unstable. Please report any issues on our <a href="https://github.com/javanhut/TheCarrionLanguage/issues" target="_blank" rel="noopener noreferrer">GitHub repository</a>.</p>
        </WarningBox>
      </Section>

      <Section>
        <SectionTitle>💡 Tips for Success</SectionTitle>
        
        <StepCard>
          <SubSectionTitle>Best Practices</SubSectionTitle>
          <ul>
            <li><strong>Start Small:</strong> Begin with simple programs and gradually add complexity</li>
            <li><strong>Use the REPL:</strong> Perfect for testing ideas and learning syntax</li>
            <li><strong>Embrace the Magic:</strong> Don't be afraid of the magical terminology - it makes coding fun!</li>
            <li><strong>Read Error Messages:</strong> Carrion provides detailed error information to help you debug</li>
            <li><strong>Explore Examples:</strong> Check out the documentation for more code examples</li>
            <li><strong>Join the Community:</strong> Connect with other Carrion developers for help and inspiration</li>
          </ul>
        </StepCard>
      </Section>
    </Container>
  );
};

export default GettingStarted;