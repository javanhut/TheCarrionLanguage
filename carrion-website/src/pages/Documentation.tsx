import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import { motion } from 'framer-motion';
// import { 
//   FaRocket, FaBook, FaCode, FaCogs, FaFlask, 
//   FaGraduationCap, FaTools, FaLayerGroup, FaBug 
// } from 'react-icons/fa6';

const Container = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  padding: 6rem 2rem 4rem;
  min-height: 100vh;
`;

const Header = styled.div`
  text-align: center;
  margin-bottom: 4rem;
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

const DocGrid = styled.div`
  display: grid;
  gap: 2rem;
`;

const DocSection = styled(motion.section)`
  background: ${({ theme }) => theme.colors.background.secondary};
  border-radius: 15px;
  padding: 2rem;
  border: 1px solid ${({ theme }) => theme.colors.border};
`;

const SectionHeader = styled.div`
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
`;

const SectionIcon = styled.div`
  font-size: 2rem;
  color: ${({ theme }) => theme.colors.primary};
`;

const SectionTitle = styled.h2`
  font-size: 1.8rem;
  color: ${({ theme }) => theme.colors.primary};
`;

const SectionDescription = styled.p`
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 1.5rem;
  line-height: 1.8;
`;

const LinkGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
`;

const DocLink = styled(Link)`
  background: ${({ theme }) => theme.colors.background.tertiary};
  padding: 1.2rem;
  border-radius: 10px;
  border: 1px solid ${({ theme }) => theme.colors.border};
  transition: all ${({ theme }) => theme.transitions.normal};
  display: flex;
  flex-direction: column;
  gap: 0.5rem;

  &:hover {
    transform: translateY(-2px);
    border-color: ${({ theme }) => theme.colors.primary};
    box-shadow: ${({ theme }) => theme.shadows.medium};
  }

  h3 {
    color: ${({ theme }) => theme.colors.text.primary};
    font-size: 1.1rem;
    font-weight: 600;
  }

  p {
    color: ${({ theme }) => theme.colors.text.secondary};
    font-size: 0.9rem;
    line-height: 1.6;
  }
`;

const QuickNav = styled.div`
  background: ${({ theme }) => theme.colors.background.tertiary};
  border-radius: 15px;
  padding: 2rem;
  margin-bottom: 2rem;
  border: 1px solid ${({ theme }) => theme.colors.border};
`;

const QuickNavTitle = styled.h3`
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 1rem;
`;

const QuickNavLinks = styled.div`
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
`;

const QuickLink = styled(Link)`
  background: ${({ theme }) => theme.colors.background.secondary};
  color: ${({ theme }) => theme.colors.text.primary};
  padding: 0.5rem 1rem;
  border-radius: 20px;
  transition: all ${({ theme }) => theme.transitions.fast};

  &:hover {
    background: ${({ theme }) => theme.colors.primary};
    color: white;
  }
`;

const Documentation: React.FC = () => {
  const sections = [
    {
      icon: '🚀',
      title: "Getting Started",
      description: "New to Carrion? Start your magical journey here with installation guides and tutorials.",
      links: [
        {
          to: "/docs/installation",
          title: "Installation Guide",
          description: "Set up Carrion on your system"
        },
        {
          to: "/docs/quick-start",
          title: "Quick Start Tutorial",
          description: "Learn the basics in 10 minutes"
        },
        {
          to: "/docs/getting-started",
          title: "First Steps",
          description: "Write your first Carrion program"
        }
      ]
    },
    {
      icon: '📚',
      title: "Language Fundamentals",
      description: "Master the core concepts of Carrion programming language.",
      links: [
        {
          to: "/docs/language-reference",
          title: "Language Reference",
          description: "Complete syntax and language features"
        },
        {
          to: "/docs/operators",
          title: "Operators",
          description: "Arithmetic, logical, and comparison operators"
        },
        {
          to: "/docs/control-flow",
          title: "Control Flow",
          description: "Conditionals, loops, and flow control"
        },
        {
          to: "/docs/builtin-functions",
          title: "Built-in Functions",
          description: "Core functions available in Carrion"
        }
      ]
    },
    {
      icon: '🔮',
      title: "Advanced Features",
      description: "Explore the powerful features that make Carrion magical.",
      links: [
        {
          to: "/docs/grimoires",
          title: "Grimoires (Classes)",
          description: "Object-oriented programming in Carrion"
        },
        {
          to: "/docs/error-handling",
          title: "Error Handling",
          description: "Attempt/ensnare/resolve error handling"
        },
        {
          to: "/docs/modules",
          title: "Modules",
          description: "Organize and import code"
        }
      ]
    },
    {
      icon: '📖',
      title: "Standard Library",
      description: "Harness the power of Munin, Carrion's standard library.",
      links: [
        {
          to: "/docs/standard-library",
          title: "Munin Overview",
          description: "Complete standard library reference"
        }
      ]
    }
  ];

  const popularTopics = [
    "Hello World",
    "Variables",
    "Functions",
    "Grimoires",
    "Error Handling",
    "Arrays",
    "Strings",
    "Loops"
  ];

  return (
    <Container>
      <Header>
        <Title>Documentation</Title>
        <Subtitle>Everything you need to master the Carrion programming language</Subtitle>
      </Header>

      <QuickNav>
        <QuickNavTitle>Popular Topics</QuickNavTitle>
        <QuickNavLinks>
          {popularTopics.map((topic) => (
            <QuickLink key={topic} to="/docs/quick-start">
              {topic}
            </QuickLink>
          ))}
        </QuickNavLinks>
      </QuickNav>

      <DocGrid>
        {sections.map((section, index) => (
          <DocSection
            key={index}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, delay: index * 0.1 }}
          >
            <SectionHeader>
              <SectionIcon>{section.icon}</SectionIcon>
              <SectionTitle>{section.title}</SectionTitle>
            </SectionHeader>
            <SectionDescription>{section.description}</SectionDescription>
            <LinkGrid>
              {section.links.map((link) => (
                <DocLink key={link.to} to={link.to}>
                  <h3>{link.title}</h3>
                  <p>{link.description}</p>
                </DocLink>
              ))}
            </LinkGrid>
          </DocSection>
        ))}
      </DocGrid>
    </Container>
  );
};

export default Documentation;