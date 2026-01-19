import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import { motion } from 'framer-motion';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';

const HeroSection = styled.section`
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: ${({ theme }) => theme.colors.background.primary};
  padding: 8rem 2rem 6rem;
  position: relative;
`;

const HeroContainer = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: center;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: 1fr;
    text-align: center;
  }
`;

const HeroContent = styled.div``;

const Badge = styled(motion.div)`
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: rgba(6, 182, 212, 0.1);
  border: 1px solid rgba(6, 182, 212, 0.2);
  border-radius: 50px;
  font-size: 0.875rem;
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1.5rem;
`;

const HeroTitle = styled(motion.h1)`
  font-size: 3.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  color: ${({ theme }) => theme.colors.text.primary};
  line-height: 1.15;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    font-size: 2.5rem;
  }
`;

const Highlight = styled.span`
  color: ${({ theme }) => theme.colors.primary};
`;

const HeroSubtitle = styled(motion.p)`
  font-size: 1.25rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 2rem;
  line-height: 1.7;
  max-width: 540px;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    max-width: 100%;
  }
`;

const HeroButtons = styled(motion.div)`
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    justify-content: center;
  }
`;

const PrimaryButton = styled(Link)`
  padding: 0.875rem 2rem;
  background: ${({ theme }) => theme.colors.primary};
  color: white;
  border-radius: 8px;
  font-weight: 500;
  transition: all 0.2s ease;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;

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
  font-weight: 500;
  transition: all 0.2s ease;

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    color: ${({ theme }) => theme.colors.primary};
  }
`;

const CodeExample = styled(motion.div)`
  background: #1a1b26;
  border-radius: 12px;
  border: 1px solid ${({ theme }) => theme.colors.border};
  overflow: hidden;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    max-width: 600px;
    margin: 0 auto;
  }
`;

const CodeHeader = styled.div`
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.875rem 1.25rem;
  background: rgba(255, 255, 255, 0.03);
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
`;

const CodeDot = styled.div<{ $color: string }>`
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: ${({ $color }) => $color};
`;

const CodeFileName = styled.span`
  margin-left: 0.75rem;
  font-size: 0.875rem;
  color: ${({ theme }) => theme.colors.text.muted};
`;

const FeaturesSection = styled.section`
  padding: 6rem 2rem;
  background: ${({ theme }) => theme.colors.background.secondary};
`;

const SectionContainer = styled.div`
  max-width: 1200px;
  margin: 0 auto;
`;

const SectionHeader = styled.div`
  text-align: center;
  margin-bottom: 4rem;
`;

const SectionTitle = styled.h2`
  font-size: 2.25rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 1rem;
`;

const SectionSubtitle = styled.p`
  font-size: 1.125rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  max-width: 600px;
  margin: 0 auto;
`;

const FeatureGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: repeat(2, 1fr);
  }

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    grid-template-columns: 1fr;
  }
`;

const FeatureCard = styled.div`
  padding: 2rem;
  background: ${({ theme }) => theme.colors.background.primary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 12px;
  transition: all 0.2s ease;

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-4px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
  }
`;

const FeatureIcon = styled.div`
  width: 48px;
  height: 48px;
  background: rgba(6, 182, 212, 0.1);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 1.25rem;
  font-size: 1.25rem;
  color: ${({ theme }) => theme.colors.primary};
  font-weight: 600;
`;

const FeatureTitle = styled.h3`
  font-size: 1.125rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.5rem;
`;

const FeatureDescription = styled.p`
  font-size: 0.95rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.6;
`;

const ToolingSection = styled.section`
  padding: 6rem 2rem;
  background: ${({ theme }) => theme.colors.background.primary};
`;

const ToolingGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1.5rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    grid-template-columns: 1fr;
  }
`;

const ToolingCard = styled.div`
  padding: 1.5rem;
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  display: flex;
  gap: 1rem;
  align-items: flex-start;
`;

const ToolingIcon = styled.div`
  width: 40px;
  height: 40px;
  min-width: 40px;
  background: rgba(6, 182, 212, 0.1);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: ${({ theme }) => theme.colors.primary};
  font-weight: 600;
  font-size: 0.875rem;
`;

const ToolingContent = styled.div``;

const ToolingTitle = styled.h4`
  font-size: 1rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.25rem;
`;

const ToolingDescription = styled.p`
  font-size: 0.875rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.5;
`;

const CTASection = styled.section`
  padding: 6rem 2rem;
  background: ${({ theme }) => theme.colors.background.secondary};
`;

const CTAContainer = styled.div`
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
  padding: 4rem;
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.08) 0%, rgba(139, 92, 246, 0.08) 100%);
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 16px;
`;

const CTATitle = styled.h2`
  font-size: 2rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 1rem;
`;

const CTAText = styled.p`
  font-size: 1.125rem;
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

const Home: React.FC = () => {
  const codeExample = `# Define a class with methods
grim HttpClient:
    init(base_url):
        self.base_url = base_url
        self.timeout = 30

    spell get(endpoint):
        url = f"{self.base_url}/{endpoint}"
        return self.request("GET", url)

    spell post(endpoint, data):
        url = f"{self.base_url}/{endpoint}"
        return self.request("POST", url, data)

# Error handling
attempt:
    client = HttpClient("https://api.example.com")
    response = client.get("users")
ensnare (NetworkError):
    print("Connection failed")
resolve:
    print("Request completed")`;

  return (
    <>
      <HeroSection>
        <HeroContainer>
          <HeroContent>
            <Badge
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5 }}
            >
              v0.1.9 — Now with Sindri, Mimir & Bifrost
            </Badge>
            <HeroTitle
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: 0.1 }}
            >
              A modern language for <Highlight>clarity</Highlight> and <Highlight>productivity</Highlight>
            </HeroTitle>
            <HeroSubtitle
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: 0.2 }}
            >
              Carrion combines Python's readability with unique syntax and integrated
              tooling. Object-oriented, dynamically typed, and built in Go.
            </HeroSubtitle>
            <HeroButtons
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: 0.3 }}
            >
              <PrimaryButton to="/docs/installation">
                Get Started
              </PrimaryButton>
              <SecondaryButton to="/playground">
                Try in Browser
              </SecondaryButton>
            </HeroButtons>
          </HeroContent>
          <CodeExample
            initial={{ opacity: 0, x: 20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ duration: 0.6, delay: 0.4 }}
          >
            <CodeHeader>
              <CodeDot $color="#ff5f56" />
              <CodeDot $color="#ffbd2e" />
              <CodeDot $color="#27ca40" />
              <CodeFileName>example.crl</CodeFileName>
            </CodeHeader>
            <SyntaxHighlighter
              language="python"
              style={atomOneDark}
              customStyle={{
                margin: 0,
                padding: '1.5rem',
                fontSize: '0.875rem',
                background: 'transparent',
                lineHeight: 1.6,
              }}
            >
              {codeExample}
            </SyntaxHighlighter>
          </CodeExample>
        </HeroContainer>
      </HeroSection>

      <FeaturesSection>
        <SectionContainer>
          <SectionHeader>
            <SectionTitle>Language Features</SectionTitle>
            <SectionSubtitle>
              Everything you need to build applications, from quick scripts to complex systems.
            </SectionSubtitle>
          </SectionHeader>

          <FeatureGrid>
            <FeatureCard>
              <FeatureIcon>OOP</FeatureIcon>
              <FeatureTitle>Object-Oriented</FeatureTitle>
              <FeatureDescription>
                Full support for classes, inheritance, polymorphism, and encapsulation
                with clean, readable syntax.
              </FeatureDescription>
            </FeatureCard>

            <FeatureCard>
              <FeatureIcon>Err</FeatureIcon>
              <FeatureTitle>Error Handling</FeatureTitle>
              <FeatureDescription>
                Structured exception handling with attempt/ensnare/resolve blocks
                and detailed error reporting.
              </FeatureDescription>
            </FeatureCard>

            <FeatureCard>
              <FeatureIcon>Mod</FeatureIcon>
              <FeatureTitle>Module System</FeatureTitle>
              <FeatureDescription>
                Import system with support for local files, packages, and selective
                imports for clean namespaces.
              </FeatureDescription>
            </FeatureCard>

            <FeatureCard>
              <FeatureIcon>Lib</FeatureIcon>
              <FeatureTitle>Standard Library</FeatureTitle>
              <FeatureDescription>
                Munin standard library with collections, file I/O, math functions,
                and OS interface built-in.
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

            <FeatureCard>
              <FeatureIcon>Pat</FeatureIcon>
              <FeatureTitle>Pattern Matching</FeatureTitle>
              <FeatureDescription>
                Modern match/case statements for clean conditional logic and
                structured data handling.
              </FeatureDescription>
            </FeatureCard>
          </FeatureGrid>
        </SectionContainer>
      </FeaturesSection>

      <ToolingSection>
        <SectionContainer>
          <SectionHeader>
            <SectionTitle>Integrated Tooling</SectionTitle>
            <SectionSubtitle>
              Development tools built into the language ecosystem.
            </SectionSubtitle>
          </SectionHeader>

          <ToolingGrid>
            <ToolingCard>
              <ToolingIcon>S</ToolingIcon>
              <ToolingContent>
                <ToolingTitle>Sindri Testing</ToolingTitle>
                <ToolingDescription>
                  Built-in testing framework with automatic discovery and assertions.
                </ToolingDescription>
              </ToolingContent>
            </ToolingCard>

            <ToolingCard>
              <ToolingIcon>M</ToolingIcon>
              <ToolingContent>
                <ToolingTitle>Mimir Documentation</ToolingTitle>
                <ToolingDescription>
                  Interactive help system accessible directly from the REPL.
                </ToolingDescription>
              </ToolingContent>
            </ToolingCard>

            <ToolingCard>
              <ToolingIcon>B</ToolingIcon>
              <ToolingContent>
                <ToolingTitle>Bifrost Packages</ToolingTitle>
                <ToolingDescription>
                  Package manager with Git integration and dependency resolution.
                </ToolingDescription>
              </ToolingContent>
            </ToolingCard>
          </ToolingGrid>
        </SectionContainer>
      </ToolingSection>

      <CTASection>
        <CTAContainer>
          <CTATitle>Ready to get started?</CTATitle>
          <CTAText>
            Install Carrion in seconds and start building. Documentation and
            tutorials available to help you along the way.
          </CTAText>
          <CTAButtons>
            <PrimaryButton to="/docs/quick-start">Quick Start Guide</PrimaryButton>
            <SecondaryButton to="/documentation">Browse Documentation</SecondaryButton>
          </CTAButtons>
        </CTAContainer>
      </CTASection>
    </>
  );
};

export default Home;
