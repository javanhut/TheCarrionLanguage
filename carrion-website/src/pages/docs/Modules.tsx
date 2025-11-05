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
    from { opacity: 0; transform: translateY(-20px); }
    to { opacity: 1; transform: translateY(0); }
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
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
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

const Text = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  font-size: 1.1rem;
  line-height: 1.8;
  margin-bottom: 1.5rem;
`;

const CodeBlock = styled.div`
  margin: 2rem 0;
  border-radius: ${({ theme }) => theme.borderRadius.large};
  overflow: hidden;
  box-shadow: ${({ theme }) => theme.shadows.large};
  transition: transform ${({ theme }) => theme.transitions.standard};
  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.3);
  }
`;

const InfoBox = styled.div`
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.1), rgba(139, 92, 246, 0.1));
  padding: 1.5rem;
  border-radius: ${({ theme }) => theme.borderRadius.medium};
  border-left: 4px solid ${({ theme }) => theme.colors.primary};
  margin: 2rem 0;
`;

const InfoTitle = styled.h4`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 0.5rem;
  font-weight: 600;
`;

const InfoText = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.6;
`;

const InlineCode = styled.code`
  background: rgba(6, 182, 212, 0.1);
  color: ${({ theme }) => theme.colors.primary};
  padding: 0.2rem 0.5rem;
  border-radius: 4px;
  font-family: 'Monaco', 'Courier New', monospace;
  font-size: 0.95em;
`;

const TreeView = styled.pre`
  background: ${({ theme }) => theme.colors.surface};
  padding: 1.5rem;
  border-radius: ${({ theme }) => theme.borderRadius.medium};
  color: ${({ theme }) => theme.colors.text.secondary};
  font-family: 'Monaco', 'Courier New', monospace;
  font-size: 0.95rem;
  line-height: 1.6;
  overflow-x: auto;
  border: 1px solid rgba(6, 182, 212, 0.2);
`;

const Modules: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Module System</Title>
        <Subtitle>
          Organize and reuse code across files with Carrion's module system. Build modular, maintainable applications.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Basic Import Syntax</SectionTitle>
        <Text>
          Import functionality from other files to organize your code and promote reusability.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Simple import
import "filename"

// Import with extension (optional for .crl files)
import "utilities.crl"
import "math_functions"

// Import from subdirectories
import "utils/helpers"
import "lib/data_structures"
import "../shared/common"  // Relative paths

// Import with aliases
import "very_long_module_name" as short_name
import "data_structures.Stack" as MyStack`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Creating Modules</SectionTitle>
        
        <SubSectionTitle>Utility Module Example</SubSectionTitle>
        <Text>
          Create a module with reusable functions and grimoires.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// File: math_utils.crl
spell add(a, b):
    return a + b

spell multiply(a, b):
    return a * b

spell factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)

grim Calculator:
    init():
        self.history = []
    
    spell calculate(operation, a, b):
        match operation:
            case "add":
                result = add(a, b)
            case "multiply":
                result = multiply(a, b)
            _:
                result = "Unknown operation"
        
        self.history.append(f"{operation}({a}, {b}) = {result}")
        return result`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Using the Module</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// File: main.crl
import "math_utils"

// Use imported functions
result1 = add(5, 3)
result2 = multiply(4, 7)
fact = factorial(5)

print(f"5 + 3 = {result1}")      // 8
print(f"4 * 7 = {result2}")      // 28
print(f"5! = {fact}")            // 120

// Use imported grimoire
calc = Calculator()
sum_result = calc.calculate("add", 10, 15)
print(f"Calculator result: {sum_result}")  // 25`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Selective Imports</SectionTitle>
        <Text>
          Import specific grimoires or functions when you don't need everything from a module.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// File: data_structures.crl
grim Stack:
    init():
        self.items = []
    
    spell push(item):
        self.items.append(item)
    
    spell pop():
        if len(self.items) > 0:
            return self.items.pop()
        return None

grim Queue:
    init():
        self.items = []
    
    spell enqueue(item):
        self.items.append(item)
    
    spell dequeue():
        if len(self.items) > 0:
            return self.items.pop(0)
        return None

// File: main.crl
import "data_structures.Stack"
// Only Stack is imported, Queue is not available

stack = Stack()
stack.push(1)
stack.push(2)
print(stack.pop())  // 2

// queue = Queue()  // Error: Queue not imported`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Module Organization Patterns</SectionTitle>
        
        <SubSectionTitle>Constants Module</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// File: constants.crl
PI = 3.14159265359
E = 2.71828182846
GOLDEN_RATIO = 1.61803398875

// Configuration constants
MAX_RETRY_ATTEMPTS = 3
DEFAULT_TIMEOUT = 30
API_VERSION = "v1.2.0"

// Color constants
COLOR_RED = "#FF0000"
COLOR_GREEN = "#00FF00"
COLOR_BLUE = "#0000FF"

grim Colors:
    RED = "#FF0000"
    GREEN = "#00FF00"
    BLUE = "#0000FF"
    
    spell hex_to_rgb(hex_color):
        // Convert hex to RGB
        return (255, 0, 0)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Configuration Module</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// File: config.crl
grim Config:
    init():
        self.settings = {
            "debug": False,
            "log_level": "INFO",
            "database_url": "localhost:5432",
            "cache_enabled": True
        }
    
    spell get(key, default = None):
        return self.settings.get(key, default)
    
    spell set(key, value):
        self.settings[key] = value
    
    spell load_from_file(filename):
        file = File()
        if file.exists(filename):
            content = file.read(filename)
            print(f"Loaded config from {filename}")
    
    spell save_to_file(filename):
        file = File()
        content = str(self.settings)
        file.write(filename, content)
        print(f"Saved config to {filename}")

// Global configuration instance
app_config = Config()`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Project Structure</SectionTitle>
        <Text>
          Organize larger projects with a clear directory structure.
        </Text>

        <TreeView>
{`project/
├── main.crl                 Main entry point
├── config.crl              Configuration settings
├── constants.crl           Application constants
├── utils/
│   ├── string_utils.crl    String utilities
│   ├── math_utils.crl      Math functions
│   └── file_utils.crl      File operations
├── models/
│   ├── user.crl           User data model
│   ├── product.crl        Product model
│   └── order.crl          Order model
├── services/
│   ├── user_service.crl   User business logic
│   ├── auth_service.crl   Authentication
│   └── data_service.crl   Data access
└── tests/
    ├── test_utils.crl     Test utilities
    ├── test_models.crl    Model tests
    └── test_services.crl  Service tests`}
        </TreeView>

        <SubSectionTitle>Main Application Structure</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// File: main.crl
import "config"
import "services.user_service"
import "services.auth_service"
import "utils.string_utils"

// Initialize application
app_config.load_from_file("app.config")
auth = AuthService()
user_service = UserService()

spell main():
    print("Welcome to Carrion Application")
    
    username = input("Username: ")
    password = input("Password: ")
    
    if auth.authenticate(username, password):
        user = user_service.get_user(username)
        formatted_name = capitalize_words(user.full_name)
        print(f"Welcome, {formatted_name}!")
    else:
        print("Authentication failed")

main()`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Advanced Import Patterns</SectionTitle>
        
        <SubSectionTitle>Conditional Imports</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Import based on conditions
debug_mode = True

if debug_mode:
    import "debug_utilities"
    enable_debug_logging()
else:
    import "production_utilities"
    enable_performance_monitoring()`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Dynamic Module Loading</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell load_database_driver(database_type):
    match database_type:
        case "mysql":
            import "drivers.mysql_driver"
            return MySQLDriver()
        case "postgresql":
            import "drivers.postgresql_driver" 
            return PostgreSQLDriver()
        case "sqlite":
            import "drivers.sqlite_driver"
            return SQLiteDriver()
        _:
            raise Error("Database", f"Unsupported: {database_type}")

// Usage
db_type = input("Enter database type: ")
driver = load_database_driver(db_type)`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Module Initialization</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// File: logger.crl
print("Logger module loaded")  // Runs on import

grim Logger:
    init(name):
        self.name = name
        self.messages = []
    
    spell log(level, message):
        timestamp = get_current_time()
        formatted = f"[{timestamp}] {level}: {message}"
        self.messages.append(formatted)
        print(formatted)
    
    spell debug(message):
        self.log("DEBUG", message)
    
    spell info(message):
        self.log("INFO", message)
    
    spell error(message):
        self.log("ERROR", message)

// Create default logger
default_logger = Logger("default")
default_logger.info("Logger module initialized")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Error Handling with Imports</SectionTitle>
        <Text>
          Handle missing or problematic modules gracefully.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell safe_import(module_name):
    attempt:
        import module_name
        return True
    ensnare:
        print(f"Failed to import {module_name}")
        return False

// Graceful fallback for optional features
if safe_import("advanced_graphics"):
    use_advanced_graphics = True
    print("Advanced graphics enabled")
else:
    use_advanced_graphics = False
    print("Using basic graphics")

// Check dependencies
spell check_dependencies():
    required = ["math_utils", "string_utils", "data_structures"]
    
    for module in required:
        if not safe_import(module):
            print(f"Error: Required module '{module}' not found")
            return False
    
    print("All dependencies satisfied")
    return True

if not check_dependencies():
    print("Cannot start application")
    exit(1)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Best Practices</SectionTitle>
        
        <InfoBox>
          <InfoTitle>One Responsibility Per Module</InfoTitle>
          <InfoText>
            Each module should have a single, well-defined purpose. Keep related functionality together 
            and separate unrelated concerns.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Avoid Circular Dependencies</InfoTitle>
          <InfoText>
            Don't create situations where Module A imports Module B and Module B imports Module A. 
            Refactor to extract shared functionality into a third module.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Use Clear Module Names</InfoTitle>
          <InfoText>
            Choose descriptive names that clearly indicate the module's purpose: 
            <InlineCode>user_service.crl</InlineCode>, <InlineCode>math_utils.crl</InlineCode>, 
            <InlineCode>database_config.crl</InlineCode>
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Document Module Interfaces</InfoTitle>
          <InfoText>
            Add docstrings to modules describing their purpose, key functions, and usage examples. 
            This helps other developers understand how to use your code.
          </InfoText>
        </InfoBox>
      </Section>
    </Container>
  );
};

export default Modules;
