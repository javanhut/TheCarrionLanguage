import React, { useState, useEffect, useCallback } from 'react';
import styled from 'styled-components';
import { motion } from 'framer-motion';
import { carrionWasm, CarrionResult, StdlibStatus } from '../utils/carrion-wasm';

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
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const VersionBadge = styled.span`
  font-size: 0.8rem;
  padding: 0.2rem 0.6rem;
  background: ${({ theme }) => theme.colors.primary}20;
  color: ${({ theme }) => theme.colors.primary};
  border-radius: 4px;
`;

const OutputContent = styled.pre<{ hasError?: boolean }>`
  padding: 1.5rem;
  min-height: 500px;
  margin: 0;
  font-family: ${({ theme }) => theme.fonts.code};
  font-size: 0.9rem;
  line-height: 1.6;
  color: ${({ hasError, theme }) => hasError ? '#ff6b6b' : theme.colors.text.primary};
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

const StatusIndicator = styled.span<{ status: 'loading' | 'ready' | 'error' }>`
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.3rem 0.8rem;
  border-radius: 20px;
  font-size: 0.85rem;

  ${({ status }) => {
    switch (status) {
      case 'loading':
        return `
          background: #f39c1220;
          color: #f39c12;
        `;
      case 'ready':
        return `
          background: #27ae6020;
          color: #27ae60;
        `;
      case 'error':
        return `
          background: #e74c3c20;
          color: #e74c3c;
        `;
    }
  }}
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
print("\\nCounting spell:")
for i in range(5):
    print("Abracadabra!")

print("\\nPower levels:")
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

# Another example
attempt:
    spell_name = "Test Spell"
    print(f"Casting {spell_name}...")
    result = 10 / 2
    print(f"Result: {result}")
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

type WasmStatus = 'loading' | 'ready' | 'error';

const Playground: React.FC = () => {
  const [code, setCode] = useState(codeExamples.helloWorld);
  const [output, setOutput] = useState('Loading Carrion interpreter...');
  const [isRunning, setIsRunning] = useState(false);
  const [wasmStatus, setWasmStatus] = useState<WasmStatus>('loading');
  const [version, setVersion] = useState('');
  const [hasError, setHasError] = useState(false);

  // Initialize WASM on component mount
  useEffect(() => {
    initWasm();
  }, []);

  const initWasm = async () => {
    try {
      setWasmStatus('loading');
      setOutput('Loading Carrion interpreter...');

      await carrionWasm.init();
      const ver = await carrionWasm.getVersion();
      setVersion(ver);

      // Check stdlib status
      const stdlibStatus: StdlibStatus = await carrionWasm.getStdlibStatus();

      setWasmStatus('ready');
      if (stdlibStatus.loaded) {
        setOutput('Carrion interpreter ready with Munin standard library. Click "Run" to execute your code.');
      } else {
        setOutput(`Carrion interpreter ready (stdlib warning: ${stdlibStatus.error}). Click "Run" to execute your code.`);
      }
      setHasError(false);
    } catch (error) {
      setWasmStatus('error');
      setOutput(`Failed to load Carrion interpreter: ${error instanceof Error ? error.message : String(error)}`);
      setHasError(true);
    }
  };

  const runCode = useCallback(async () => {
    if (wasmStatus !== 'ready') {
      setOutput('Interpreter not ready. Please wait for it to load.');
      return;
    }

    setIsRunning(true);
    setOutput('Executing...');
    setHasError(false);

    try {
      const result: CarrionResult = await carrionWasm.evaluate(code);

      if (result.success) {
        let outputText = '';

        // Add print output
        if (result.output) {
          outputText += result.output;
        }

        // Add final result if it's not empty
        if (result.result && result.result.trim()) {
          if (outputText) outputText += '\n';
          outputText += `=> ${result.result}`;
        }

        if (!outputText.trim()) {
          outputText = 'Program executed successfully (no output)';
        }

        setOutput(outputText);
        setHasError(false);
      } else {
        let errorOutput = '';
        if (result.output) {
          errorOutput += result.output + '\n\n';
        }
        errorOutput += `Error: ${result.error}`;
        setOutput(errorOutput);
        setHasError(true);
      }
    } catch (error) {
      setOutput(`Execution error: ${error instanceof Error ? error.message : String(error)}`);
      setHasError(true);
    } finally {
      setIsRunning(false);
    }
  }, [code, wasmStatus]);

  const loadExample = (exampleKey: keyof typeof codeExamples) => {
    setCode(codeExamples[exampleKey]);
    setOutput('Example loaded. Click "Run" to execute.');
    setHasError(false);
  };

  const clearCode = () => {
    setCode('');
    setOutput('');
    setHasError(false);
  };

  const resetEnvironment = async () => {
    await carrionWasm.reset();
    setOutput('Environment reset. Variables and functions cleared.');
    setHasError(false);
  };

  const getStatusText = () => {
    switch (wasmStatus) {
      case 'loading':
        return 'Loading...';
      case 'ready':
        return 'Ready';
      case 'error':
        return 'Error';
    }
  };

  return (
    <Container>
      <Header>
        <Title>Carrion Playground</Title>
        <Subtitle>
          Run real Carrion code directly in your browser
          <StatusIndicator status={wasmStatus} style={{ marginLeft: '1rem' }}>
            {getStatusText()}
          </StatusIndicator>
        </Subtitle>
      </Header>

      <InfoSection>
        <h2>Real Carrion Code Execution</h2>
        <p style={{ textAlign: 'center', marginBottom: '1.5rem', color: '#8892b0' }}>
          This playground runs the actual Carrion interpreter compiled to WebAssembly.
          Your code executes entirely in your browser - no server required.
        </p>
        <InfoGrid>
          <InfoCard>
            <h3>Full Language Support</h3>
            <p>Run real Carrion code with grimoires, spells, error handling, and all language features.</p>
          </InfoCard>
          <InfoCard>
            <h3>Instant Execution</h3>
            <p>Code runs locally in your browser using WebAssembly for near-native performance.</p>
          </InfoCard>
          <InfoCard>
            <h3>Learn by Doing</h3>
            <p>Experiment freely - the interpreter runs in a sandboxed environment with no network access.</p>
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
              <Button onClick={resetEnvironment} disabled={wasmStatus !== 'ready'}>
                Reset
              </Button>
              <Button onClick={clearCode}>
                Clear
              </Button>
              <Button primary onClick={runCode} disabled={isRunning || wasmStatus !== 'ready'}>
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
          <OutputHeader>
            <span>Output</span>
            {version && <VersionBadge>Carrion v{version}</VersionBadge>}
          </OutputHeader>
          <OutputContent hasError={hasError}>{output}</OutputContent>
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
