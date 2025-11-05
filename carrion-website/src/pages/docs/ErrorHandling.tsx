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
  &:nth-child(5) { animation-delay: 0.4s; }
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

const ErrorHandling: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Error Handling</Title>
        <Subtitle>
          Master Carrion's magical error handling system with attempt-ensnare-resolve patterns for robust, reliable code.
        </Subtitle>
      </Header>

      <Section>
        <SectionTitle>Basic Syntax</SectionTitle>
        <Text>
          Carrion uses magical terminology for error handling: <InlineCode>attempt</InlineCode> for try, 
          <InlineCode>ensnare</InlineCode> for catch, and <InlineCode>resolve</InlineCode> for finally.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Basic error handling structure
attempt:
    // Risky code that might fail
    result = 10 / 0
ensnare:
    // Handle any error
    print("An error occurred!")
    result = 0

// With finally block (resolve)
attempt:
    file = File()
    content = file.read("data.txt")
ensnare:
    print("Error reading file!")
resolve:
    // Always runs, even if no error
    print("Cleanup completed")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Handling Specific Errors</SectionTitle>
        <Text>
          Catch specific error types to handle different failures appropriately.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`// Multiple error types
attempt:
    number = int(input("Enter a number: "))
    result = 100 / number
    print(f"Result: {result}")
ensnare (ValueError):
    print("Invalid number format!")
ensnare (ZeroDivisionError):
    print("Cannot divide by zero!")
ensnare:
    print("An unexpected error occurred!")

// Access error details
attempt:
    risky_operation()
ensnare (error):
    print(f"Error type: {type(error)}")
    print(f"Error message: {error.message}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Raising Errors</SectionTitle>
        <Text>
          Create and raise custom errors to signal problems in your code.
        </Text>

        <SubSectionTitle>Basic Error Raising</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell validate_age(age):
    if age < 0:
        raise Error("Validation", "Age cannot be negative")
    if age > 150:
        raise Error("Validation", "Age seems unrealistic")
    return True

// Usage
attempt:
    validate_age(-5)
ensnare (Error):
    print("Validation failed!")

// Custom error types
ValidationError = Error("ValidationError", "Input validation failed")
NetworkError = Error("NetworkError", "Network operation failed")

spell connect_to_server(url):
    if not url.startswith("http"):
        raise ValidationError
    
    if url == "http://broken-server.com":
        raise NetworkError
    
    return "Connected successfully"

attempt:
    result = connect_to_server("invalid-url")
ensnare (ValidationError):
    print("Invalid URL format")
ensnare (NetworkError):
    print("Network connection failed")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Error Recovery Patterns</SectionTitle>
        
        <SubSectionTitle>Retry Logic</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell connect_with_retry(url, max_attempts = 3):
    attempts = 0
    
    while attempts < max_attempts:
        attempt:
            return connect_to_server(url)
        ensnare (NetworkError):
            attempts += 1
            print(f"Connection attempt {attempts} failed")
            
            if attempts < max_attempts:
                print("Retrying in 2 seconds...")
                os = OS()
                os.sleep(2)
            else:
                print("Max retry attempts reached")
                raise Error("Connection", "Failed after retries")

// Usage
attempt:
    connection = connect_with_retry("http://unreliable-server.com")
    print("Successfully connected!")
ensnare:
    print("Connection failed permanently")`}
          </SyntaxHighlighter>
        </CodeBlock>

        <SubSectionTitle>Nested Error Handling</SubSectionTitle>
        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell process_user_data(user_data):
    attempt:
        // Validate data structure
        attempt:
            validate_structure(user_data)
        ensnare (ValidationError):
            print("Data structure validation failed")
            raise Error("Processing", "Invalid data structure")
        
        // Process individual fields
        for field in user_data:
            attempt:
                process_field(field)
            ensnare:
                print(f"Warning: Failed to process field {field}")
                skip  // Continue with next field
    
    ensnare (Error):
        print("Critical error in data processing")
        return False
    
    return True`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Error Handling in Grimoires</SectionTitle>
        <Text>
          Implement robust error handling in object-oriented code.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`grim BankAccount:
    init(account_number, initial_balance = 0):
        attempt:
            if initial_balance < 0:
                raise Error("InvalidBalance", "Balance cannot be negative")
            
            self.account_number = account_number
            self.balance = initial_balance
            self.transaction_history = []
        
        ensnare:
            print(f"Account creation failed: {error.message}")
            raise  // Re-raise to prevent invalid object
    
    spell withdraw(amount):
        attempt:
            if amount <= 0:
                raise Error("InvalidAmount", "Amount must be positive")
            
            if amount > self.balance:
                raise Error("InsufficientFunds", "Not enough balance")
            
            self.balance -= amount
            self.transaction_history.append(f"Withdrawal: -{amount}")
            return True
        
        ensnare (Error):
            print(f"Withdrawal failed: {error.message}")
            return False
    
    spell deposit(amount):
        attempt:
            if amount <= 0:
                raise Error("InvalidAmount", "Amount must be positive")
            
            self.balance += amount
            self.transaction_history.append(f"Deposit: +{amount}")
            return True
        
        ensnare (Error):
            print(f"Deposit failed: {error.message}")
            return False

// Usage
attempt:
    account = BankAccount("12345", 1000)
    
    if account.withdraw(500):
        print("Withdrawal successful")
    
    if not account.withdraw(2000):
        print("Large withdrawal blocked")
    
ensnare:
    print("Account operations failed")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Error Propagation</SectionTitle>
        <Text>
          Control how errors flow through your program with re-raising and wrapping.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell read_config_file(filename):
    attempt:
        file = File()
        if not file.exists(filename):
            raise Error("FileNotFound", f"Config file '{filename}' not found")
        
        content = file.read(filename)
        if len(content) == 0:
            raise Error("EmptyFile", "Configuration file is empty")
        
        return parse_config(content)
    
    ensnare:
        print(f"Failed to read config: {error.message}")
        raise  // Re-raise the error

spell initialize_application():
    attempt:
        config = read_config_file("app.config")
        setup_database(config)
        start_services(config)
        return True
    
    ensnare:
        print("Application initialization failed")
        return False

// Usage
if not initialize_application():
    print("Cannot start application")
    exit(1)`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Assertions and Debugging</SectionTitle>
        <Text>
          Use assertions to validate assumptions during development.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell factorial(n):
    check(n >= 0, "Factorial undefined for negative numbers")
    check(type(n) == "INTEGER", "Factorial requires integer")
    
    if n <= 1:
        return 1
    return n * factorial(n - 1)

// Debug assertions
DEBUG_MODE = True

spell debug_assert(condition, message):
    if DEBUG_MODE and not condition:
        raise Error("AssertionError", f"Assertion failed: {message}")

spell process_array(arr):
    debug_assert(arr is not None, "Array cannot be None")
    debug_assert(len(arr) > 0, "Array cannot be empty")
    debug_assert(type(arr) == "ARRAY", "Input must be array")
    
    total = 0
    for item in arr:
        debug_assert(type(item) in ["INTEGER", "FLOAT"], 
                     "Array items must be numbers")
        total += item
    
    return total`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>

      <Section>
        <SectionTitle>Best Practices</SectionTitle>
        
        <InfoBox>
          <InfoTitle>Specific Error Messages</InfoTitle>
          <InfoText>
            Provide clear, actionable error messages that explain what went wrong and how to fix it. 
            Avoid generic messages like "Error occurred".
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Fail Fast, Recover Gracefully</InfoTitle>
          <InfoText>
            Validate inputs early and raise errors immediately when detecting problems. Use 
            <InlineCode>ensnare</InlineCode> blocks for recovery strategies and fallback values.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Always Clean Up Resources</InfoTitle>
          <InfoText>
            Use the <InlineCode>resolve</InlineCode> block to ensure cleanup code runs regardless of errors. 
            This is essential for file handles, network connections, and other resources.
          </InfoText>
        </InfoBox>

        <InfoBox>
          <InfoTitle>Log Errors Appropriately</InfoTitle>
          <InfoText>
            Log error details for debugging, but show user-friendly messages to users. Never expose 
            sensitive information in error messages.
          </InfoText>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>Complete Example</SectionTitle>
        <Text>
          A comprehensive example showing multiple error handling techniques together.
        </Text>

        <CodeBlock>
          <SyntaxHighlighter language="python" style={atomOneDark} customStyle={{ margin: 0, padding: '1.5rem' }}>
{`spell safe_file_processor(input_file, output_file):
    """Process file with comprehensive error handling."""
    file = File()
    processed_lines = 0
    
    attempt:
        // Validate inputs
        if not input_file or not output_file:
            raise Error("InvalidInput", "File names cannot be empty")
        
        if input_file == output_file:
            raise Error("InvalidInput", "Input and output must differ")
        
        // Check file existence
        if not file.exists(input_file):
            raise Error("FileNotFound", f"Input file '{input_file}' not found")
        
        // Read and process
        content = file.read(input_file)
        lines = content.split("\\n")
        
        processed = []
        for i, line in enumerate(lines):
            attempt:
                // Process each line
                processed_line = line.strip().upper()
                if len(processed_line) > 0:
                    processed.append(processed_line)
                    processed_lines += 1
            ensnare:
                print(f"Warning: Failed to process line {i}")
                skip  // Continue with next line
        
        // Write output
        result = "\\n".join(processed)
        file.write(output_file, result)
        
        return {
            "success": True,
            "lines_processed": processed_lines,
            "output_file": output_file
        }
    
    ensnare (Error):
        return {
            "success": False,
            "error": error.message,
            "lines_processed": processed_lines
        }
    
    resolve:
        // Always log completion
        print(f"Processing completed. Lines processed: {processed_lines}")

// Usage
result = safe_file_processor("input.txt", "output.txt")
if result["success"]:
    print(f"Success! Processed {result['lines_processed']} lines")
else:
    print(f"Failed: {result['error']}")`}
          </SyntaxHighlighter>
        </CodeBlock>
      </Section>
    </Container>
  );
};

export default ErrorHandling;
