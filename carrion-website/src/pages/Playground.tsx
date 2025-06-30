import React, { useState } from 'react';
import styled from 'styled-components';
import { motion } from 'framer-motion';

const Container = styled.div`
  max-width: 1400px;
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
`;

const PlaygroundContainer = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin-bottom: 2rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: 1fr;
  }
`;

const EditorContainer = styled(motion.div)`
  background: ${({ theme }) => theme.colors.background.secondary};
  border-radius: 15px;
  overflow: hidden;
  border: 1px solid ${({ theme }) => theme.colors.border};
`;

const EditorHeader = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  padding: 1rem 1.5rem;
  display: flex;
  justify-content: between;
  align-items: center;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
`;

const EditorTitle = styled.div`
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: ${({ theme }) => theme.colors.text.primary};
  font-weight: 600;
  flex: 1;
`;

const EditorActions = styled.div`
  display: flex;
  gap: 1rem;
`;

const Button = styled.button<{ primary?: boolean }>`
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.6rem 1.2rem;
  border-radius: 8px;
  border: none;
  font-weight: 600;
  cursor: pointer;
  transition: all ${({ theme }) => theme.transitions.normal};

  ${({ primary, theme }) => primary ? `
    background: ${theme.colors.primary};
    color: white;

    &:hover:not(:disabled) {
      background: ${theme.colors.text.accent};
      transform: translateY(-1px);
    }
  ` : `
    background: ${theme.colors.background.secondary};
    color: ${theme.colors.text.primary};
    border: 1px solid ${theme.colors.border};

    &:hover:not(:disabled) {
      background: ${theme.colors.background.primary};
    }
  `}

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
`;

const CodeEditor = styled.textarea`
  width: 100%;
  min-height: 500px;
  padding: 1.5rem;
  background: ${({ theme }) => theme.colors.code};
  color: ${({ theme }) => theme.colors.text.primary};
  border: none;
  font-family: ${({ theme }) => theme.fonts.code};
  font-size: 0.9rem;
  line-height: 1.6;
  resize: vertical;
  outline: none;

  &::placeholder {
    color: ${({ theme }) => theme.colors.text.secondary};
  }
`;

const OutputContainer = styled(motion.div)`
  background: ${({ theme }) => theme.colors.background.secondary};
  border-radius: 15px;
  overflow: hidden;
  border: 1px solid ${({ theme }) => theme.colors.border};
`;

const OutputHeader = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  padding: 1rem 1.5rem;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
`;

const OutputContent = styled.pre`
  padding: 1.5rem;
  min-height: 500px;
  margin: 0;
  font-family: ${({ theme }) => theme.fonts.code};
  font-size: 0.9rem;
  line-height: 1.6;
  color: ${({ theme }) => theme.colors.text.primary};
  overflow-x: auto;
  white-space: pre-wrap;
  word-wrap: break-word;
`;

const InfoSection = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border-radius: 15px;
  padding: 2rem;
  margin-bottom: 2rem;
  border: 1px solid ${({ theme }) => theme.colors.border};
`;

const InfoGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-top: 1.5rem;
`;

const InfoCard = styled.div`
  h3 {
    color: ${({ theme }) => theme.colors.primary};
    margin-bottom: 0.5rem;
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  p {
    color: ${({ theme }) => theme.colors.text.secondary};
    line-height: 1.8;
  }
`;

const ExamplesSection = styled.div`
  margin-top: 3rem;
`;

const ExampleGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-top: 1rem;
`;

const ExampleButton = styled.button`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  padding: 1rem;
  text-align: left;
  cursor: pointer;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-2px);
  }

  h4 {
    color: ${({ theme }) => theme.colors.primary};
    margin-bottom: 0.3rem;
  }

  p {
    color: ${({ theme }) => theme.colors.text.secondary};
    font-size: 0.9rem;
  }
`;

// Carrion code examples
const codeExamples = {
  helloWorld: `# Welcome to Carrion!
print("Hello, Magical World!")

# Define a simple spell (function)
spell greet(name):
    return f"Greetings, {name}! Welcome to Carrion!"

print(greet("Fellow Mage"))`,

  variables: `# Variables in Carrion
name = "Munin"
age = 1000
is_magical = True
power_level = 9.5

print(f"Name: {name}")
print(f"Age: {age} years")
print(f"Magical: {is_magical}")
print(f"Power Level: {power_level}")

# Arrays and operations
spells = ["Fireball", "Lightning", "Heal"]
spells.append("Teleport")
print(f"Known spells: {spells}")`,

  grimoire: `# Define a grimoire (class)
grim MagicalCrow:
    init(name, power):
        self.name = name
        self.power = power
        self.spells = []
    
    spell learn_spell(spell_name):
        self.spells.append(spell_name)
        return f"{self.name} learned {spell_name}!"
    
    spell cast(target):
        if len(self.spells) > 0:
            spell = self.spells[0]
            return f"{self.name} casts {spell} at {target}!"
        else:
            return f"{self.name} has no spells to cast!"

# Create and use a magical crow
munin = MagicalCrow("Munin", 100)
print(munin.learn_spell("Lightning Bolt"))
print(munin.learn_spell("Healing Touch"))
print(munin.cast("the darkness"))`,

  controlFlow: `# Control flow in Carrion
power = 75

if power > 90:
    print("You are a Master Wizard!")
otherwise power > 60:
    print("You are a skilled mage!")
otherwise power > 30:
    print("You are an apprentice.")
else:
    print("Keep practicing your magic!")

# Loops
print("\nCounting spell:")
for i in range(5):
    print("Abracadabra!")

print("\nPower levels:")
levels = [10, 25, 50, 75, 100]
for level in levels:
    print(f"Power level: {level}")`,

  errorHandling: `# Error handling in Carrion
spell divide_magic(power, divisor):
    return power / divisor

attempt:
    result = divide_magic(100, 0)
    print(f"Result: {result}")
ensnare:
    print("Cannot divide by zero in the magical realm!")
resolve:
    print("Magic calculation complete!")

# Custom errors
attempt:
    spell_name = "Forbidden Spell"
    print(f"Casting {spell_name}...")
    # This would cause an error in real code
    raise "This spell is too dangerous!"
ensnare:
    print("Spell failed!")
resolve:
    print("Spell casting complete!")`,

  fibonacci: `# Fibonacci sequence in Carrion
spell fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

# Calculate first 10 Fibonacci numbers
print("Fibonacci Sequence:")
for i in range(10):
    result = fibonacci(i)
    print(f"F({i}) = {result}")

# Alternative using array
numbers = []
for i in range(10):
    numbers.append(fibonacci(i))

print("")
print(f"First 10 Fibonacci numbers: {numbers}")`
};

const Playground: React.FC = () => {
  const [code, setCode] = useState(codeExamples.helloWorld);
  const [output, setOutput] = useState('Click "Run" to execute your Carrion code...');
  const [isRunning, setIsRunning] = useState(false);
  const [apiAvailable, setApiAvailable] = useState(true);

  // Check API availability on component mount
  React.useEffect(() => {
    checkApiHealth();
  }, []);

  const checkApiHealth = async () => {
    try {
      // In production, we'll use simulation mode since the API requires a backend server
      const apiUrl = process.env.NODE_ENV === 'development' 
        ? 'http://localhost:3001/health' 
        : null; // No API in production for now
      
      if (!apiUrl) {
        setApiAvailable(false);
        return;
      }
      
      const response = await fetch(apiUrl);
      setApiAvailable(response.ok);
    } catch {
      setApiAvailable(false);
    }
  };

  const runCode = async () => {
    if (!apiAvailable) {
      setOutput('⚠️ Playground API is not available.\nUsing simulation mode...\n\n' + simulateCarrionExecution(code));
      return;
    }

    setIsRunning(true);
    setOutput('🐦‍⬛ Executing Carrion code...\n');

    try {
      const apiUrl = process.env.NODE_ENV === 'development' 
        ? 'http://localhost:3001/execute' 
        : null; // No API in production for now
      
      if (!apiUrl) {
        throw new Error('API not available in production');
      }
      
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ code }),
      });

      const result = await response.json();

      if (result.success) {
        setOutput(result.output || 'Program executed successfully (no output)');
      } else {
        setOutput(`❌ Execution Error:\n${result.stderr || result.error || 'Unknown error'}`);
      }
    } catch (error: any) {
      console.error('API Error:', error);
      setOutput(`🔌 Connection Error: ${error.message}\n\nFalling back to simulation mode...\n\n` + simulateCarrionExecution(code));
      setApiAvailable(false);
    } finally {
      setIsRunning(false);
    }
  };

  const simulateCarrionExecution = (code: string): string => {
    const output: string[] = [];
    const lines = code.split('\n');
    
    // Simple simulation - just look for print statements
    lines.forEach(line => {
      const trimmed = line.trim();
      
      // Skip comments and empty lines
      if (!trimmed || trimmed.startsWith('//')) return;
      
      // Simulate print statements
      if (trimmed.startsWith('print(')) {
        const match = trimmed.match(/print\((.*)\)/);
        if (match) {
          let value = match[1].trim();
          
          // Remove quotes for string literals
          if ((value.startsWith('"') && value.endsWith('"')) || 
              (value.startsWith("'") && value.endsWith("'"))) {
            value = value.slice(1, -1);
          }
          
          output.push(value);
        }
      }
    });

    // Add a note about simulation
    if (output.length === 0) {
      output.push('Code executed successfully (no output)');
    }
    
    output.push('\n---');
    output.push('Note: This is a simulated output. For full Carrion features,');
    output.push('please download and install the Carrion interpreter.');
    
    return output.join('\n');
  };

  const loadExample = (exampleKey: keyof typeof codeExamples) => {
    setCode(codeExamples[exampleKey]);
    setOutput('Example loaded. Click "Run" to execute.');
  };

  const clearCode = () => {
    setCode('');
    setOutput('');
  };

  return (
    <Container>
      <Header>
        <Title>Carrion Playground</Title>
        <Subtitle>Try Carrion directly in your browser</Subtitle>
      </Header>

      <InfoSection>
        <h2>🚀 Real Carrion Code Execution</h2>
        <p style={{ textAlign: 'center', marginBottom: '1.5rem', color: '#8892b0' }}>
          This playground executes actual Carrion code using the official interpreter in a secure Docker environment.
          {!apiAvailable && (
            <span style={{ color: '#f39c12', display: 'block', marginTop: '0.5rem' }}>
              ⚠️ API currently unavailable - using simulation mode
            </span>
          )}
        </p>
        <InfoGrid>
          <InfoCard>
            <h3>🔒 Secure Execution</h3>
            <p>Code runs in isolated Docker containers with resource limits and network isolation for security.</p>
          </InfoCard>
          <InfoCard>
            <h3>⚡ Real-time Results</h3>
            <p>See actual Carrion interpreter output, errors, and execution behavior instantly.</p>
          </InfoCard>
          <InfoCard>
            <h3>📚 Learn by Doing</h3>
            <p>Experiment with real Carrion features - grimoires, spells, error handling, and more.</p>
          </InfoCard>
        </InfoGrid>
      </InfoSection>

      <PlaygroundContainer>
        <EditorContainer
          initial={{ opacity: 0, x: -20 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ duration: 0.5 }}
        >
          <EditorHeader>
            <EditorTitle>
               main.crl
            </EditorTitle>
            <EditorActions>
              <Button onClick={clearCode}>
                 Clear
              </Button>
              <Button primary onClick={runCode} disabled={isRunning}>
                 {isRunning ? 'Running...' : 'Run'}
              </Button>
            </EditorActions>
          </EditorHeader>
          <CodeEditor
            value={code}
            onChange={(e) => setCode(e.target.value)}
            placeholder={`# Write your Carrion code here...
# Try: print("Hello, Carrion!")
# Or load an example from below`}
            spellCheck={false}
          />
        </EditorContainer>

        <OutputContainer
          initial={{ opacity: 0, x: 20 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ duration: 0.5 }}
        >
          <OutputHeader>Output</OutputHeader>
          <OutputContent>{output}</OutputContent>
        </OutputContainer>
      </PlaygroundContainer>

      <ExamplesSection>
        <h2>Example Programs</h2>
        <ExampleGrid>
          <ExampleButton onClick={() => loadExample('helloWorld')}>
            <h4>Hello World</h4>
            <p>Basic program structure and output</p>
          </ExampleButton>
          <ExampleButton onClick={() => loadExample('variables')}>
            <h4>Variables & Types</h4>
            <p>Working with different data types</p>
          </ExampleButton>
          <ExampleButton onClick={() => loadExample('grimoire')}>
            <h4>Grimoires (Classes)</h4>
            <p>Object-oriented programming</p>
          </ExampleButton>
          <ExampleButton onClick={() => loadExample('controlFlow')}>
            <h4>Control Flow</h4>
            <p>Conditionals and loops</p>
          </ExampleButton>
          <ExampleButton onClick={() => loadExample('errorHandling')}>
            <h4>Error Handling</h4>
            <p>Attempt/ensnare/resolve blocks</p>
          </ExampleButton>
          <ExampleButton onClick={() => loadExample('fibonacci')}>
            <h4>Fibonacci Sequence</h4>
            <p>Recursive functions example</p>
          </ExampleButton>
        </ExampleGrid>
      </ExamplesSection>
    </Container>
  );
};

export default Playground;