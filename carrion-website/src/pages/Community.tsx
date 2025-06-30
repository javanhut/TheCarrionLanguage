import React from 'react';
import styled from 'styled-components';

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

const Card = styled.div`
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

const InfoBox = styled.div`
  background: rgba(0, 204, 153, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.primary};
  border-radius: 8px;
  padding: 1.5rem;
  margin: 2rem 0;
  text-align: center;
`;

const WarningBox = styled.div`
  background: rgba(255, 204, 0, 0.1);
  border: 1px solid ${({ theme }) => theme.colors.warning};
  border-radius: 8px;
  padding: 1.5rem;
  margin: 2rem 0;
`;

const GitHubLink = styled.a`
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  background: ${({ theme }) => theme.colors.primary};
  color: white;
  padding: 1rem 2rem;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    background: ${({ theme }) => theme.colors.text.accent};
    transform: translateY(-1px);
  }
`;

const LinkGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin: 2rem 0;
`;

const LinkCard = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  padding: 1.5rem;
  text-align: center;
`;

const Community: React.FC = () => {
  return (
    <Container>
      <Header>
        <Title>Community</Title>
        <Subtitle>
          Connect with the Carrion programming language community and contribute to its growth
        </Subtitle>
      </Header>

      <WarningBox>
        <h3>🚧 Development Status</h3>
        <p><strong>Important:</strong> Carrion is currently in active development. We are <strong>not accepting code contributions</strong> at this time as the language architecture and core features are still being established.</p>
      </WarningBox>

      <Section>
        <SectionTitle>🐛 Report Issues & Bugs</SectionTitle>
        <Card>
          <h3>Found a Problem?</h3>
          <p>While we're not accepting code contributions yet, we <strong>welcome bug reports and issue submissions</strong>! Your feedback helps improve Carrion.</p>
          
          <h4>What to Report:</h4>
          <ul>
            <li><strong>Bugs:</strong> Unexpected behavior or crashes</li>
            <li><strong>Documentation Issues:</strong> Unclear or incorrect documentation</li>
            <li><strong>Feature Requests:</strong> Ideas for future language features</li>
            <li><strong>Installation Problems:</strong> Issues with setup or deployment</li>
            <li><strong>Playground Bugs:</strong> Problems with the online code editor</li>
          </ul>

          <div style={{ textAlign: 'center', marginTop: '2rem' }}>
            <GitHubLink 
              href="https://github.com/javanhut/TheCarrionLanguage/issues/new"
              target="_blank"
              rel="noopener noreferrer"
            >
              🐛 Report an Issue
            </GitHubLink>
          </div>
        </Card>
      </Section>

      <Section>
        <SectionTitle>📚 Getting Help</SectionTitle>
        <LinkGrid>
          <LinkCard>
            <h3>📖 Documentation</h3>
            <p>Start with our comprehensive guides and language reference.</p>
            <a href="/docs/getting-started">Getting Started Guide</a>
          </LinkCard>
          
          <LinkCard>
            <h3>🎮 Playground</h3>
            <p>Try Carrion code directly in your browser with real execution.</p>
            <a href="/playground">Online Playground</a>
          </LinkCard>
          
          <LinkCard>
            <h3>💾 Installation</h3>
            <p>Download and install Carrion on your system.</p>
            <a href="/docs/installation">Installation Guide</a>
          </LinkCard>
        </LinkGrid>
      </Section>

      <Section>
        <SectionTitle>🔗 Project Links</SectionTitle>
        <Card>
          <h3>GitHub Repository</h3>
          <p>Visit the official Carrion repository to view source code, track development progress, and submit issues.</p>
          
          <div style={{ textAlign: 'center', marginTop: '1.5rem' }}>
            <GitHubLink 
              href="https://github.com/javanhut/TheCarrionLanguage"
              target="_blank"
              rel="noopener noreferrer"
            >
              🐙 View on GitHub
            </GitHubLink>
          </div>
        </Card>
      </Section>

      <Section>
        <SectionTitle>🚀 Future Contributions</SectionTitle>
        <InfoBox>
          <h3>Want to Contribute Code?</h3>
          <p>We appreciate your interest in contributing to Carrion! Once the core language design is stabilized, we plan to open up for community contributions.</p>
          <p><strong>For now, please:</strong></p>
          <ul style={{ textAlign: 'left', maxWidth: '400px', margin: '1rem auto' }}>
            <li>⭐ Star the repository to show support</li>
            <li>📝 Report bugs and issues</li>
            <li>📖 Help improve documentation</li>
            <li>🎯 Suggest features and improvements</li>
            <li>🗣️ Spread the word about Carrion</li>
          </ul>
        </InfoBox>
      </Section>

      <Section>
        <SectionTitle>📊 Project Statistics</SectionTitle>
        <Card>
          <h3>Carrion Language Development</h3>
          <ul>
            <li><strong>Current Version:</strong> v0.1.6</li>
            <li><strong>Language:</strong> Go (Backend), TypeScript/React (Website)</li>
            <li><strong>License:</strong> Check repository for details</li>
            <li><strong>Platform Support:</strong> Linux, macOS, Windows</li>
            <li><strong>Container Support:</strong> Docker, Podman</li>
          </ul>
          
          <p style={{ marginTop: '1.5rem' }}>
            Follow the repository to stay updated on development progress and feature releases.
          </p>
        </Card>
      </Section>
    </Container>
  );
};

export default Community;