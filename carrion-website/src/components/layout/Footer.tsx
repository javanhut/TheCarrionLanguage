import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
// import { FaGithub, FaEnvelope } from 'react-icons/fa6';

const FooterContainer = styled.footer`
  background: ${({ theme }) => theme.colors.background.secondary};
  border-top: 1px solid ${({ theme }) => theme.colors.border};
  margin-top: auto;
  padding: 3rem 0 2rem;
`;

const FooterContent = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
`;

const FooterGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 3rem;
  margin-bottom: 3rem;
`;

const FooterSection = styled.div``;

const FooterTitle = styled.h4`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1.5rem;
  font-size: 1.2rem;
`;

const FooterList = styled.ul`
  list-style: none;
`;

const FooterListItem = styled.li`
  margin-bottom: 0.75rem;
`;

const FooterLink = styled(Link)`
  color: ${({ theme }) => theme.colors.text.secondary};
  transition: color ${({ theme }) => theme.transitions.fast};

  &:hover {
    color: ${({ theme }) => theme.colors.primary};
  }
`;

const ExternalLink = styled.a`
  color: ${({ theme }) => theme.colors.text.secondary};
  transition: color ${({ theme }) => theme.transitions.fast};
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;

  &:hover {
    color: ${({ theme }) => theme.colors.primary};
  }
`;

const FooterBottom = styled.div`
  text-align: center;
  padding-top: 2rem;
  border-top: 1px solid ${({ theme }) => theme.colors.border};
  color: ${({ theme }) => theme.colors.text.secondary};
`;

const LogoSection = styled.div`
  p {
    color: ${({ theme }) => theme.colors.text.secondary};
    margin-bottom: 1rem;
  }
`;

const LogoIcon = styled.div`
  font-size: 3rem;
  margin-top: 1rem;
`;

const Footer: React.FC = () => {
  return (
    <FooterContainer>
      <FooterContent>
        <FooterGrid>
          <FooterSection>
            <FooterTitle>Carrion Language</FooterTitle>
            <LogoSection>
              <p>Where code meets magic</p>
              <LogoIcon>🐦‍⬛</LogoIcon>
            </LogoSection>
          </FooterSection>

          <FooterSection>
            <FooterTitle>Quick Links</FooterTitle>
            <FooterList>
              <FooterListItem>
                <FooterLink to="/documentation">Documentation</FooterLink>
              </FooterListItem>
              <FooterListItem>
                <FooterLink to="/playground">Playground</FooterLink>
              </FooterListItem>
              <FooterListItem>
                <FooterLink to="/download">Downloads</FooterLink>
              </FooterListItem>
              <FooterListItem>
                <ExternalLink 
                  href="https://github.com/javanhut/TheCarrionLanguage" 
                  target="_blank" 
                  rel="noopener noreferrer"
                >
                  GitHub
                </ExternalLink>
              </FooterListItem>
            </FooterList>
          </FooterSection>

          <FooterSection>
            <FooterTitle>Resources</FooterTitle>
            <FooterList>
              <FooterListItem>
                <FooterLink to="/docs/getting-started">Getting Started</FooterLink>
              </FooterListItem>
              <FooterListItem>
                <FooterLink to="/docs/language-reference">Language Reference</FooterLink>
              </FooterListItem>
              <FooterListItem>
                <FooterLink to="/docs/standard-library">Standard Library</FooterLink>
              </FooterListItem>
              <FooterListItem>
                <ExternalLink 
                  href="https://github.com/javanhut/TheCarrionLanguage/blob/main/CONTRIBUTING.md" 
                  target="_blank" 
                  rel="noopener noreferrer"
                >
                  Contributing
                </ExternalLink>
              </FooterListItem>
            </FooterList>
          </FooterSection>

          <FooterSection>
            <FooterTitle>Community</FooterTitle>
            <FooterList>
              <FooterListItem>
                <ExternalLink 
                  href="https://github.com/javanhut/TheCarrionLanguage/discussions" 
                  target="_blank" 
                  rel="noopener noreferrer"
                >
                  Forum
                </ExternalLink>
              </FooterListItem>
              <FooterListItem>
                <ExternalLink 
                  href="https://github.com/javanhut/TheCarrionLanguage/issues" 
                  target="_blank" 
                  rel="noopener noreferrer"
                >
                  Issue Tracker
                </ExternalLink>
              </FooterListItem>
              <FooterListItem>
                <ExternalLink 
                  href="mailto:javanhut@carrionlang.com"
                >
                  Email
                </ExternalLink>
              </FooterListItem>
            </FooterList>
          </FooterSection>
        </FooterGrid>

        <FooterBottom>
          <p>&copy; 2024 Carrion Programming Language. Licensed under MIT License.</p>
          <p>Built with magic by the Carrion community</p>
        </FooterBottom>
      </FooterContent>
    </FooterContainer>
  );
};

export default Footer;