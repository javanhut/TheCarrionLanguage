import React from 'react';
import styled from 'styled-components';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
// import { FaLinux, FaApple, FaWindows, FaDownload, FaTerminal, FaGithub } from 'react-icons/fa6';

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
`;

const Section = styled.section`
  margin-bottom: 3rem;
`;

const SectionTitle = styled.h2`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1.5rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
`;

const InstallGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-bottom: 3rem;
`;

const InstallCard = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 15px;
  padding: 2rem;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-2px);
  }
`;

const CardHeader = styled.div`
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
`;

const OSIcon = styled.div`
  font-size: 2rem;
  color: ${({ theme }) => theme.colors.primary};
`;

const CardTitle = styled.h3`
  color: ${({ theme }) => theme.colors.text.primary};
`;

const DownloadButton = styled.a`
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  background: ${({ theme }) => theme.colors.primary};
  color: white;
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
  transition: all ${({ theme }) => theme.transitions.normal};
  margin-bottom: 1rem;

  &:hover {
    background: ${({ theme }) => theme.colors.text.accent};
    transform: translateY(-1px);
  }
`;

const CodeBlock = styled.div`
  margin: 1rem 0;
`;

const StepList = styled.ol`
  margin-left: 1.5rem;
  
  li {
    margin-bottom: 1rem;
    line-height: 1.8;
    color: ${({ theme }) => theme.colors.text.primary};
  }
`;

const WarningBox = styled.div`
  background: rgba(255, 204, 0, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.warning};
  border-radius: 8px;
  padding: 1rem;
  margin: 1rem 0;
  
  p {
    color: ${({ theme }) => theme.colors.text.primary};
    margin: 0;
  }
`;

const InfoBox = styled.div`
  background: rgba(0, 204, 153, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.primary};
  border-radius: 8px;
  padding: 1rem;
  margin: 1rem 0;
  
  p {
    color: ${({ theme }) => theme.colors.text.primary};
    margin: 0;
  }
`;

const Installation: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Installation Guide</Title>
        <Subtitle>Get Carrion up and running on your system</Subtitle>
      </Header>

      <Section>
        <SectionTitle>
           Quick Install Options
        </SectionTitle>
        
        <InstallGrid>
          <InstallCard>
            <CardHeader>
              <OSIcon></OSIcon>
              <CardTitle>Linux</CardTitle>
            </CardHeader>
            <DownloadButton 
              href="https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.6/carrion_linux_amd64.tar.gz"
              download
            >
               Download for Linux
            </DownloadButton>
            <CodeBlock>
              <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Download and install
curl -L "https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.6/carrion_linux_amd64.tar.gz" -o carrion_linux_amd64.tar.gz
tar -xzf carrion_linux_amd64.tar.gz
sudo cp carrion /usr/local/bin/
chmod +x /usr/local/bin/carrion

# Test installation
carrion
# In the REPL, type: version()`}
              </SyntaxHighlighter>
            </CodeBlock>
          </InstallCard>

          <InstallCard>
            <CardHeader>
              <OSIcon></OSIcon>
              <CardTitle>macOS</CardTitle>
            </CardHeader>
            <DownloadButton 
              href="https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.6/carrion-src.tar.gz"
              download
            >
               Download for macOS (Universal)
            </DownloadButton>
            <CodeBlock>
              <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Download and install
curl -L "https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.6/carrion-src.tar.gz" -o carrion-src.tar.gz
tar -xzf carrion-src.tar.gz
sudo cp carrion /usr/local/bin/
chmod +x /usr/local/bin/carrion

# Test installation
carrion
# In the REPL, type: version()`}
              </SyntaxHighlighter>
            </CodeBlock>
          </InstallCard>

          <InstallCard>
            <CardHeader>
              <OSIcon></OSIcon>
              <CardTitle>Windows</CardTitle>
            </CardHeader>
            <DownloadButton 
              href="https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.6/carrion_windows_amd64.zip"
              download
            >
               Download for Windows
            </DownloadButton>
            <p>Download the .zip file, extract it, and add the folder to your PATH environment variable.</p>
            <CodeBlock>
              <SyntaxHighlighter language="powershell" style={atomOneDark}>
{`# Download and extract
curl -L "https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.6/carrion_windows_amd64.zip" -o carrion_windows_amd64.zip
# Extract the zip file to a folder (e.g., C:\\carrion)
# Add the folder to your PATH environment variable
# Open Command Prompt or PowerShell and test:
carrion
# In the REPL, type: version()`}
              </SyntaxHighlighter>
            </CodeBlock>
          </InstallCard>
        </InstallGrid>
      </Section>

      <Section>
        <SectionTitle>
           Docker Installation
        </SectionTitle>
        <p>Use Docker to run Carrion without installing it directly on your system:</p>
        
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Pull latest version
docker pull javanhut/carrionlanguage:latest
docker run -it javanhut/carrionlanguage:latest
# This automatically starts the Carrion REPL
# Type version() to check the version

# Pull specific version (e.g., 0.1.6)
docker pull javanhut/carrionlanguage:0.1.6
docker run -it javanhut/carrionlanguage:0.1.6

# Available tags: latest, 0.1.6, 0.1.5, 0.1.4, 0.1.3, 0.1.2, 0.1.1, 0.1.0`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <p><strong>Docker benefits:</strong> No installation required, isolated environment, easy version switching, works on any platform with Docker installed.</p>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>
           One-Line Installer
        </SectionTitle>
        <p>For Linux and macOS users, use our convenient installation script:</p>
        
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Using curl
curl -fsSL https://raw.githubusercontent.com/javanhut/TheCarrionLanguage/main/install/install.sh | sh

# Using wget
wget -qO- https://raw.githubusercontent.com/javanhut/TheCarrionLanguage/main/install/install.sh | sh`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <p><strong>What this script does:</strong> Downloads the appropriate binary for your system, makes it executable, and places it in /usr/local/bin for global access.</p>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>
           Build from Source
        </SectionTitle>
        <p>For developers who want the latest features or to contribute to Carrion:</p>

        <StepList>
          <li>
            <strong>Prerequisites:</strong> Ensure you have Go 1.19 or later installed
            <CodeBlock>
              <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Check Go version
go version`}
              </SyntaxHighlighter>
            </CodeBlock>
          </li>
          
          <li>
            <strong>Clone the repository:</strong>
            <CodeBlock>
              <SyntaxHighlighter language="bash" style={atomOneDark}>
{`git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage`}
              </SyntaxHighlighter>
            </CodeBlock>
          </li>
          
          <li>
            <strong>Build the interpreter:</strong>
            <CodeBlock>
              <SyntaxHighlighter language="bash" style={atomOneDark}>
{`cd src
go build -o carrion`}
              </SyntaxHighlighter>
            </CodeBlock>
          </li>
          
          <li>
            <strong>Test the installation:</strong>
            <CodeBlock>
              <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Run the REPL
./carrion

# Or run a file
echo 'print("Hello, Carrion!")' > test.crl
./carrion test.crl`}
              </SyntaxHighlighter>
            </CodeBlock>
          </li>
          
          <li>
            <strong>Optional: Install globally:</strong>
            <CodeBlock>
              <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Copy to system PATH
sudo cp carrion /usr/local/bin/`}
              </SyntaxHighlighter>
            </CodeBlock>
          </li>
        </StepList>
      </Section>

      <Section>
        <SectionTitle>Verification</SectionTitle>
        <p>After installation, verify that Carrion is working correctly:</p>
        
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Start the REPL to check version
carrion
# In the REPL, type: version()
# Press Ctrl+C or Ctrl+D to exit

# Create and run a simple program
echo 'print("Magic works!")' > hello.crl
carrion hello.crl`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <p><strong>Note:</strong> The <code>carrion --version</code> flag is not yet implemented. Use the REPL command <code>version()</code> to check the installed version.</p>
        </InfoBox>

        <InfoBox>
          <p><strong>Expected output:</strong> You should see the Carrion crow ASCII art when starting the REPL, and "Magic works!" when running the test file.</p>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>Troubleshooting</SectionTitle>
        
        <h3>Common Issues</h3>
        
        <h4>Permission Denied Error</h4>
        <p>If you get a permission denied error, make sure the binary is executable:</p>
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`chmod +x carrion`}
          </SyntaxHighlighter>
        </CodeBlock>

        <h4>Command Not Found</h4>
        <p>If the system can't find the <code>carrion</code> command, ensure it's in your PATH:</p>
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Check if /usr/local/bin is in PATH
echo $PATH

# Add to PATH temporarily
export PATH=$PATH:/usr/local/bin

# Add to PATH permanently (bash)
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
source ~/.bashrc`}
          </SyntaxHighlighter>
        </CodeBlock>

        <h4>macOS Security Warning</h4>
        <p>On macOS, you might see a security warning for unsigned binaries:</p>
        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Allow the binary to run
sudo xattr -rd com.apple.quarantine carrion`}
          </SyntaxHighlighter>
        </CodeBlock>

        <WarningBox>
          <p><strong>Note:</strong> Carrion is currently in alpha. Some features may be unstable. Please report any issues on our GitHub repository.</p>
        </WarningBox>
      </Section>

      <Section>
        <SectionTitle>Next Steps</SectionTitle>
        <p>Now that you have Carrion installed, you're ready to start your magical programming journey!</p>
        
        <ul>
          <li>📚 <a href="/docs/quick-start">Quick Start Tutorial</a> - Learn the basics in 10 minutes</li>
          <li>🎮 <a href="/playground">Try the Online Playground</a> - Experiment with code in your browser</li>
          <li>📖 <a href="/docs/language-reference">Language Reference</a> - Complete syntax guide</li>
          <li>💬 <a href="/community">Join the Community</a> - Get help and share your creations</li>
        </ul>
      </Section>
    </Container>
  );
};

export default Installation;