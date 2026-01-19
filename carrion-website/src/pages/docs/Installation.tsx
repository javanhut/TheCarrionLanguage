import React, { useState } from 'react';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import { Link } from 'react-router-dom';
import {
  DocLayout,
  Section,
  SectionTitle,
  SubSection,
  SubSectionTitle,
  Paragraph,
  CodeBlock,
  InfoBox,
  InfoTitle,
  InfoText,
  WarningBox,
  WarningTitle,
  TipBox,
  TipTitle,
  CardGrid,
  Card,
  CardTitle,
  CardDescription,
  TabContainer,
  TabList,
  Tab,
  TabContent,
  StepContainer,
  Step,
  StepNumber,
  StepContent,
  StepTitle,
  StepDescription,
  List,
  ListItem,
  InlineCode,
} from '../../components/docs';

const sections = [
  { id: 'quick-install', title: 'Quick Install' },
  { id: 'docker', title: 'Docker' },
  { id: 'one-liner', title: 'One-Line Installer' },
  { id: 'from-source', title: 'Build from Source' },
  { id: 'verification', title: 'Verification' },
  { id: 'troubleshooting', title: 'Troubleshooting' },
];

type Platform = 'linux' | 'macos' | 'windows';

const Installation: React.FC = () => {
  const [platform, setPlatform] = useState<Platform>('linux');

  const platformCommands = {
    linux: {
      download: `curl -L "https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.9/carrion_linux_amd64.tar.gz" -o carrion.tar.gz`,
      extract: 'tar -xzf carrion.tar.gz',
      install: 'sudo cp carrion /usr/local/bin/',
      chmod: 'chmod +x /usr/local/bin/carrion',
    },
    macos: {
      download: `curl -L "https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.9/carrion-src.tar.gz" -o carrion.tar.gz`,
      extract: 'tar -xzf carrion.tar.gz',
      install: 'sudo cp carrion /usr/local/bin/',
      chmod: 'chmod +x /usr/local/bin/carrion',
    },
    windows: {
      download: `curl -L "https://github.com/javanhut/TheCarrionLanguage/releases/download/v0.1.9/carrion_windows_amd64.zip" -o carrion.zip`,
      extract: '# Extract the zip file to a folder (e.g., C:\\carrion)',
      install: '# Add the folder to your PATH environment variable',
      chmod: '',
    },
  };

  return (
    <DocLayout
      title="Installation"
      description="Get Carrion up and running on your system in just a few steps."
      sections={sections}
    >
      <Section id="quick-install">
        <SectionTitle>Quick Install</SectionTitle>
        <Paragraph>
          Choose your operating system below and follow the installation steps.
        </Paragraph>

        <TabContainer>
          <TabList>
            <Tab $active={platform === 'linux'} onClick={() => setPlatform('linux')}>
              Linux
            </Tab>
            <Tab $active={platform === 'macos'} onClick={() => setPlatform('macos')}>
              macOS
            </Tab>
            <Tab $active={platform === 'windows'} onClick={() => setPlatform('windows')}>
              Windows
            </Tab>
          </TabList>

          <TabContent>
            <StepContainer>
              <Step>
                <StepNumber>1</StepNumber>
                <StepContent>
                  <StepTitle>Download</StepTitle>
                  <StepDescription>
                    <CodeBlock>
                      <SyntaxHighlighter language="bash" style={atomOneDark}>
                        {platformCommands[platform].download}
                      </SyntaxHighlighter>
                    </CodeBlock>
                  </StepDescription>
                </StepContent>
              </Step>

              <Step>
                <StepNumber>2</StepNumber>
                <StepContent>
                  <StepTitle>Extract</StepTitle>
                  <StepDescription>
                    <CodeBlock>
                      <SyntaxHighlighter language="bash" style={atomOneDark}>
                        {platformCommands[platform].extract}
                      </SyntaxHighlighter>
                    </CodeBlock>
                  </StepDescription>
                </StepContent>
              </Step>

              <Step>
                <StepNumber>3</StepNumber>
                <StepContent>
                  <StepTitle>Install</StepTitle>
                  <StepDescription>
                    <CodeBlock>
                      <SyntaxHighlighter language="bash" style={atomOneDark}>
                        {platformCommands[platform].install}
                        {platformCommands[platform].chmod && `\n${platformCommands[platform].chmod}`}
                      </SyntaxHighlighter>
                    </CodeBlock>
                  </StepDescription>
                </StepContent>
              </Step>

              <Step>
                <StepNumber>4</StepNumber>
                <StepContent>
                  <StepTitle>Verify</StepTitle>
                  <StepDescription>
                    <CodeBlock>
                      <SyntaxHighlighter language="bash" style={atomOneDark}>
{`carrion
# In the REPL, type: version()`}
                      </SyntaxHighlighter>
                    </CodeBlock>
                  </StepDescription>
                </StepContent>
              </Step>
            </StepContainer>
          </TabContent>
        </TabContainer>
      </Section>

      <Section id="docker">
        <SectionTitle>Docker</SectionTitle>
        <Paragraph>
          Run Carrion in a container without installing anything on your system.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Pull and run the latest version
docker pull javanhut/carrionlanguage:latest
docker run -it javanhut/carrionlanguage:latest

# Pull a specific version
docker pull javanhut/carrionlanguage:0.1.9
docker run -it javanhut/carrionlanguage:0.1.9

# Available tags: latest, 0.1.9, 0.1.8, 0.1.7, 0.1.6`}
          </SyntaxHighlighter>
        </CodeBlock>

        <TipBox>
          <TipTitle>Why Docker?</TipTitle>
          <InfoText>
            Docker provides an isolated environment, easy version switching, and works on any platform.
            No installation cleanup needed.
          </InfoText>
        </TipBox>
      </Section>

      <Section id="one-liner">
        <SectionTitle>One-Line Installer</SectionTitle>
        <Paragraph>
          For Linux and macOS users, use our installation script:
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Using curl
curl -fsSL https://raw.githubusercontent.com/javanhut/TheCarrionLanguage/main/install/install.sh | sh

# Using wget
wget -qO- https://raw.githubusercontent.com/javanhut/TheCarrionLanguage/main/install/install.sh | sh`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>What the script does</InfoTitle>
          <InfoText>
            Downloads the appropriate binary for your system, makes it executable, and places it
            in <InlineCode>/usr/local/bin</InlineCode> for global access.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="from-source">
        <SectionTitle>Build from Source</SectionTitle>
        <Paragraph>
          For developers who want the latest features or to contribute to Carrion.
        </Paragraph>

        <SubSection>
          <SubSectionTitle>Prerequisites</SubSectionTitle>
          <List>
            <ListItem>Go 1.19 or later</ListItem>
            <ListItem>Git</ListItem>
          </List>

          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
              {'go version  # Should show Go 1.19+'}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Build Steps</SubSectionTitle>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Clone the repository
git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage

# Build the interpreter
cd src
go build -o carrion

# Test the build
./carrion

# Optional: Install globally
sudo cp carrion /usr/local/bin/`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>
      </Section>

      <Section id="verification">
        <SectionTitle>Verification</SectionTitle>
        <Paragraph>
          After installation, verify that Carrion is working correctly.
        </Paragraph>

        <CodeBlock>
          <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Start the REPL
carrion

# In the REPL, check the version
>>> version()
Carrion v0.1.9

# Exit the REPL
>>> quit

# Run a test file
echo 'print("Hello, Carrion!")' > hello.crl
carrion hello.crl`}
          </SyntaxHighlighter>
        </CodeBlock>

        <InfoBox>
          <InfoTitle>Expected Output</InfoTitle>
          <InfoText>
            You should see the Carrion crow ASCII art when starting the REPL, and "Hello, Carrion!"
            when running the test file.
          </InfoText>
        </InfoBox>
      </Section>

      <Section id="troubleshooting">
        <SectionTitle>Troubleshooting</SectionTitle>

        <SubSection>
          <SubSectionTitle>Permission Denied</SubSectionTitle>
          <Paragraph>
            Make sure the binary is executable:
          </Paragraph>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
              {'chmod +x carrion'}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>Command Not Found</SubSectionTitle>
          <Paragraph>
            Ensure <InlineCode>/usr/local/bin</InlineCode> is in your PATH:
          </Paragraph>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
{`# Check PATH
echo $PATH

# Add to PATH temporarily
export PATH=$PATH:/usr/local/bin

# Add to PATH permanently (bash)
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
source ~/.bashrc`}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <SubSection>
          <SubSectionTitle>macOS Security Warning</SubSectionTitle>
          <Paragraph>
            On macOS, you might see a security warning for unsigned binaries:
          </Paragraph>
          <CodeBlock>
            <SyntaxHighlighter language="bash" style={atomOneDark}>
              {'sudo xattr -rd com.apple.quarantine carrion'}
            </SyntaxHighlighter>
          </CodeBlock>
        </SubSection>

        <WarningBox>
          <WarningTitle>Alpha Software</WarningTitle>
          <InfoText>
            Carrion is currently in alpha. Some features may be unstable. Please report any issues on our{' '}
            <a href="https://github.com/javanhut/TheCarrionLanguage/issues" target="_blank" rel="noopener noreferrer">
              GitHub repository
            </a>.
          </InfoText>
        </WarningBox>
      </Section>

      <Section>
        <SectionTitle>Next Steps</SectionTitle>
        <CardGrid>
          <Card as={Link} to="/docs/quick-start" style={{ textDecoration: 'none' }}>
            <CardTitle>Quick Start Tutorial</CardTitle>
            <CardDescription>Learn the basics of Carrion in a hands-on tutorial.</CardDescription>
          </Card>
          <Card as={Link} to="/playground" style={{ textDecoration: 'none' }}>
            <CardTitle>Online Playground</CardTitle>
            <CardDescription>Try Carrion in your browser without installing anything.</CardDescription>
          </Card>
          <Card as={Link} to="/docs/language-reference" style={{ textDecoration: 'none' }}>
            <CardTitle>Language Reference</CardTitle>
            <CardDescription>Explore the complete syntax and features of Carrion.</CardDescription>
          </Card>
        </CardGrid>
      </Section>
    </DocLayout>
  );
};

export default Installation;
