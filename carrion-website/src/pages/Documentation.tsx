import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import { motion } from 'framer-motion';

const Container = styled.div`
  max-width: 1200px;
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

// Start Here Section - prominent for beginners
const StartHereSection = styled(motion.section)`
  background: linear-gradient(135deg, ${({ theme }) => theme.colors.primary}15 0%, ${({ theme }) => theme.colors.secondary}15 100%);
  border: 2px solid ${({ theme }) => theme.colors.primary};
  border-radius: 16px;
  padding: 2.5rem;
  margin-bottom: 3rem;
`;

const StartHereTitle = styled.h2`
  font-size: 1.5rem;
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1rem;
`;

const StartHereDescription = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 2rem;
  font-size: 1.1rem;
  line-height: 1.7;
`;

const StartHereLinks = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
`;

const StartHereLink = styled(Link)`
  background: ${({ theme }) => theme.colors.background.primary};
  padding: 1.5rem;
  border-radius: 12px;
  border: 1px solid ${({ theme }) => theme.colors.border};
  transition: all ${({ theme }) => theme.transitions.normal};
  display: flex;
  align-items: flex-start;
  gap: 1rem;

  &:hover {
    transform: translateY(-3px);
    border-color: ${({ theme }) => theme.colors.primary};
    box-shadow: ${({ theme }) => theme.shadows.medium};
  }
`;

const StepNumber = styled.div`
  background: ${({ theme }) => theme.colors.primary};
  color: ${({ theme }) => theme.colors.text.inverse};
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.9rem;
  flex-shrink: 0;
`;

const StepContent = styled.div`
  h3 {
    color: ${({ theme }) => theme.colors.text.primary};
    font-size: 1.1rem;
    font-weight: 600;
    margin-bottom: 0.3rem;
  }

  p {
    color: ${({ theme }) => theme.colors.text.secondary};
    font-size: 0.9rem;
    line-height: 1.5;
  }
`;

// Main documentation sections
const SectionGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
  margin-bottom: 3rem;
`;

const DocSection = styled(motion.section)`
  background: ${({ theme }) => theme.colors.background.secondary};
  border-radius: 12px;
  padding: 1.5rem;
  border: 1px solid ${({ theme }) => theme.colors.border};
`;

const SectionTitle = styled.h2`
  font-size: 1.3rem;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.5rem;
  padding-bottom: 0.75rem;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
`;

const SectionDescription = styled.p`
  color: ${({ theme }) => theme.colors.text.muted};
  font-size: 0.9rem;
  margin-bottom: 1rem;
  line-height: 1.6;
`;

const DocList = styled.ul`
  list-style: none;
  padding: 0;
  margin: 0;
`;

const DocListItem = styled.li`
  margin-bottom: 0.5rem;

  &:last-child {
    margin-bottom: 0;
  }
`;

const DocLink = styled(Link)`
  display: block;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  color: ${({ theme }) => theme.colors.text.primary};
  transition: all ${({ theme }) => theme.transitions.fast};

  &:hover {
    background: ${({ theme }) => theme.colors.background.tertiary};
    color: ${({ theme }) => theme.colors.primary};
  }

  span {
    display: block;
    color: ${({ theme }) => theme.colors.text.muted};
    font-size: 0.8rem;
    margin-top: 0.2rem;
  }
`;

// Try it section
const TryItSection = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border-radius: 12px;
  padding: 2rem;
  text-align: center;
  border: 1px solid ${({ theme }) => theme.colors.border};
`;

const TryItTitle = styled.h2`
  font-size: 1.5rem;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.75rem;
`;

const TryItDescription = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 1.5rem;
`;

const TryItButton = styled(Link)`
  display: inline-block;
  background: ${({ theme }) => theme.gradients.primary};
  color: white;
  padding: 0.75rem 2rem;
  border-radius: 8px;
  font-weight: 600;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:hover {
    transform: translateY(-2px);
    box-shadow: ${({ theme }) => theme.shadows.glow};
    color: white;
  }
`;

const Documentation: React.FC = () => {
  const beginnerSteps = [
    {
      step: 1,
      to: "/docs/installation",
      title: "Install Carrion",
      description: "Download and set up Carrion on your computer"
    },
    {
      step: 2,
      to: "/docs/getting-started",
      title: "Write Your First Program",
      description: "Learn the basics with a hands-on tutorial"
    },
    {
      step: 3,
      to: "/docs/language-reference",
      title: "Explore the Language",
      description: "Discover variables, functions, and more"
    }
  ];

  const sections = [
    {
      title: "Language Basics",
      description: "Core concepts every Carrion developer needs to know",
      links: [
        { to: "/docs/language-reference", title: "Language Reference", desc: "Complete syntax guide" },
        { to: "/docs/operators", title: "Operators", desc: "Math, logic, and comparisons" },
        { to: "/docs/control-flow", title: "Control Flow", desc: "If statements, loops, and more" },
        { to: "/docs/builtin-functions", title: "Built-in Functions", desc: "Functions available everywhere" }
      ]
    },
    {
      title: "Object-Oriented Programming",
      description: "Build complex programs with classes and objects",
      links: [
        { to: "/docs/grimoires", title: "Grimoires (Classes)", desc: "Create reusable code structures" },
        { to: "/docs/modules", title: "Modules", desc: "Organize code into files" },
        { to: "/docs/error-handling", title: "Error Handling", desc: "Handle errors gracefully" }
      ]
    },
    {
      title: "Standard Library",
      description: "Pre-built tools to supercharge your programs",
      links: [
        { to: "/docs/standard-library", title: "Munin Overview", desc: "Explore the standard library" }
      ]
    },
    {
      title: "Tools & Guides",
      description: "Additional resources for productive development",
      links: [
        { to: "/docs/quick-start", title: "Quick Start", desc: "Get coding in 10 minutes" },
        { to: "/docs/repl-guide", title: "REPL Guide", desc: "Interactive programming tips" }
      ]
    }
  ];

  return (
    <Container>
      <Header>
        <Title>Documentation</Title>
        <Subtitle>
          Learn Carrion from the ground up with our guides and reference materials
        </Subtitle>
      </Header>

      <StartHereSection
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <StartHereTitle>New to Carrion? Start Here</StartHereTitle>
        <StartHereDescription>
          Follow these three steps to get up and running with Carrion.
          No prior experience required.
        </StartHereDescription>
        <StartHereLinks>
          {beginnerSteps.map((item) => (
            <StartHereLink key={item.step} to={item.to}>
              <StepNumber>{item.step}</StepNumber>
              <StepContent>
                <h3>{item.title}</h3>
                <p>{item.description}</p>
              </StepContent>
            </StartHereLink>
          ))}
        </StartHereLinks>
      </StartHereSection>

      <SectionGrid>
        {sections.map((section, index) => (
          <DocSection
            key={index}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, delay: 0.1 + index * 0.1 }}
          >
            <SectionTitle>{section.title}</SectionTitle>
            <SectionDescription>{section.description}</SectionDescription>
            <DocList>
              {section.links.map((link) => (
                <DocListItem key={link.to}>
                  <DocLink to={link.to}>
                    {link.title}
                    <span>{link.desc}</span>
                  </DocLink>
                </DocListItem>
              ))}
            </DocList>
          </DocSection>
        ))}
      </SectionGrid>

      <TryItSection>
        <TryItTitle>Learn by Doing</TryItTitle>
        <TryItDescription>
          Try Carrion code directly in your browser - no installation needed
        </TryItDescription>
        <TryItButton to="/playground">
          Open Playground
        </TryItButton>
      </TryItSection>
    </Container>
  );
};

export default Documentation;
