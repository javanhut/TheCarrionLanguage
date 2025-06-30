import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import { motion } from 'framer-motion';
// import { FaRocket, FaBook, FaDownload, FaPlay } from 'react-icons/fa6';
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';

const HeroSection = styled.section`
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: ${({ theme }) => theme.gradients.dark};
  padding: 6rem 2rem 4rem;
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: url('data:image/svg+xml,<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg"><defs><pattern id="grid" width="100" height="100" patternUnits="userSpaceOnUse"><path d="M 100 0 L 0 0 0 100" fill="none" stroke="rgba(255,255,255,0.03)" stroke-width="1"/></pattern></defs><rect width="100%" height="100%" fill="url(%23grid)"/></svg>');
    pointer-events: none;
  }
`;

const HeroContainer = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
  position: relative;
  z-index: 1;
`;

const HeroTitle = styled(motion.h1)`
  font-size: 4rem;
  font-weight: 800;
  margin-bottom: 1.5rem;
  background: ${({ theme }) => theme.gradients.primary};
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    font-size: 2.5rem;
  }
`;

const HeroSubtitle = styled(motion.p)`
  font-size: 1.5rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 2rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    font-size: 1.2rem;
  }
`;

const HeroAscii = styled(motion.pre)`
  font-family: monospace;
  font-size: 0.8rem;
  line-height: 1.2;
  color: ${({ theme }) => theme.colors.primary};
  margin: 2rem auto;
  opacity: 0.8;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    font-size: 0.6rem;
  }
`;

const HeroButtons = styled(motion.div)`
  display: flex;
  gap: 1.5rem;
  justify-content: center;
  margin-bottom: 3rem;
  flex-wrap: wrap;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    flex-direction: column;
    align-items: center;
  }
`;

const Button = styled(Link)<{ primary?: boolean }>`
  padding: 1rem 2rem;
  border-radius: 50px;
  font-weight: 600;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  transition: all ${({ theme }) => theme.transitions.normal};
  
  ${({ primary, theme }) => primary ? `
    background: ${theme.gradients.primary};
    color: white;
    box-shadow: ${theme.shadows.medium};

    &:hover {
      transform: translateY(-2px);
      box-shadow: ${theme.shadows.large};
    }
  ` : `
    background: transparent;
    color: ${theme.colors.primary};
    border: 2px solid ${theme.colors.primary};

    &:hover {
      background: ${theme.colors.primary};
      color: ${theme.colors.background.primary};
    }
  `}
`;

const CodeExample = styled(motion.div)`
  max-width: 700px;
  margin: 0 auto;
  text-align: left;
  background: ${({ theme }) => theme.colors.background.tertiary};
  border-radius: 10px;
  border: 1px solid ${({ theme }) => theme.colors.border};
  overflow: hidden;
`;

const FeatureSection = styled.section`
  padding: 5rem 2rem;
  background: ${({ theme }) => theme.colors.background.secondary};
`;

const SectionTitle = styled.h2`
  text-align: center;
  font-size: 3rem;
  margin-bottom: 3rem;
  background: ${({ theme }) => theme.gradients.secondary};
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
`;

const FeatureDescription = styled.div`
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
  color: ${({ theme }) => theme.colors.text.secondary};
`;

const CTASection = styled.section`
  padding: 5rem 2rem;
  background: ${({ theme }) => theme.colors.background.primary};
  text-align: center;
`;

const CTAContainer = styled.div`
  max-width: 800px;
  margin: 0 auto;
`;

const CTATitle = styled.h2`
  font-size: 2.5rem;
  margin-bottom: 1.5rem;
  color: ${({ theme }) => theme.colors.primary};
`;

const CTAText = styled.p`
  font-size: 1.2rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 2rem;
`;

const Home: React.FC = () => {
  const codeExample = `// Cast your first spell in Carrion
spell greet(name):
    return f"Welcome to the realm of magic, {name}!"

// Create a magical grimoire
grim MagicalCrow:
    init(name):
        self.name = name
        self.spells = []
    
    spell learn(spell_name):
        self.spells.append(spell_name)
        print(f"{self.name} learned {spell_name}!")

// Use your magic
munin = MagicalCrow("Munin")
munin.learn("Lightning Bolt")
print(greet("Fellow Mage"))`;


  return (
    <>
      <HeroSection>
        <HeroContainer>
          <HeroTitle
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
          >
            Welcome to Carrion
          </HeroTitle>
          <HeroSubtitle
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.2 }}
          >
            The Mystical Programming Language with a Crow's Wisdom
          </HeroSubtitle>
          <HeroAscii
            initial={{ opacity: 0 }}
            animate={{ opacity: 0.8 }}
            transition={{ duration: 1, delay: 0.4 }}
          >
{`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⡟⠋⢻⣷⣄⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣾⣿⣷⣿⣿⣿⣿⣿⣶⣾⣿⣿⠿⠿⠿⠶⠄⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠉⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⣿⣿⣿⣿⠟⠻⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣆⣤⠿⢶⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠑⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠸⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠙⠛⠋⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`}
          </HeroAscii>
          <HeroButtons
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.6 }}
          >
            <Button to="/download" primary>
               Download Latest
            </Button>
            <Button to="/docs/getting-started">
               Get Started
            </Button>
          </HeroButtons>
          <CodeExample
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.8 }}
          >
            <SyntaxHighlighter 
              language="python" 
              style={atomOneDark}
              customStyle={{
                margin: 0,
                padding: '1.5rem',
                fontSize: '0.9rem',
                background: 'transparent',
              }}
            >
              {codeExample}
            </SyntaxHighlighter>
          </CodeExample>
        </HeroContainer>
      </HeroSection>

      <FeatureSection>
        <SectionTitle>A Magical Programming Experience</SectionTitle>
        <FeatureDescription>
          <p style={{ fontSize: '1.2rem', lineHeight: '1.8', marginBottom: '2rem' }}>
            Carrion transforms mundane programming into an enchanting experience with its unique Norse mythology theme. 
            Classes become "grimoires" (spellbooks), methods become "spells", and the standard library "Munin" is named 
            after Odin's wise raven. Built with Go for performance and featuring Python-inspired syntax, Carrion makes 
            coding feel like crafting magic while maintaining familiar, readable patterns.
          </p>
          <p style={{ fontSize: '1.1rem', lineHeight: '1.8' }}>
            Whether you're a beginner learning your first programming language or an experienced developer looking for 
            something fresh and engaging, Carrion offers a complete object-oriented programming environment with 
            comprehensive error handling, module systems, and extensive documentation.
          </p>
        </FeatureDescription>
      </FeatureSection>

      <CTASection>
        <CTAContainer>
          <CTATitle>Ready to Start Your Magical Journey?</CTATitle>
          <CTAText>
            Join thousands of developers who are already crafting spells with Carrion.
            Experience programming like never before.
          </CTAText>
          <HeroButtons>
            <Button to="/playground" primary>
               Try Online
            </Button>
            <Button to="/documentation">
               Browse Docs
            </Button>
          </HeroButtons>
        </CTAContainer>
      </CTASection>
    </>
  );
};

export default Home;