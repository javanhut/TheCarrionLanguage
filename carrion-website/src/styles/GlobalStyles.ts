import { createGlobalStyle } from 'styled-components';

const GlobalStyles = createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  html {
    scroll-behavior: smooth;
  }

  body {
    font-family: ${({ theme }) => theme.fonts.primary};
    background-color: ${({ theme }) => theme.colors.background.primary};
    color: ${({ theme }) => theme.colors.text.primary};
    line-height: 1.6;
    overflow-x: hidden;
    min-height: 100vh;
  }

  #root {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }

  h1, h2, h3, h4, h5, h6 {
    line-height: 1.2;
    font-weight: 600;
  }

  h1 {
    font-size: 3rem;
    @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
      font-size: 2rem;
    }
  }

  h2 {
    font-size: 2.5rem;
    @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
      font-size: 1.75rem;
    }
  }

  h3 {
    font-size: 2rem;
    @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
      font-size: 1.5rem;
    }
  }

  a {
    color: ${({ theme }) => theme.colors.link};
    text-decoration: none;
    transition: color ${({ theme }) => theme.transitions.fast};

    &:hover {
      color: ${({ theme }) => theme.colors.linkHover};
    }
  }

  code {
    font-family: ${({ theme }) => theme.fonts.code};
    background: ${({ theme }) => theme.colors.code};
    padding: 0.2em 0.4em;
    border-radius: 4px;
    font-size: 0.9em;
  }

  pre {
    background: ${({ theme }) => theme.colors.code};
    border-radius: 8px;
    padding: 1rem;
    overflow-x: auto;
    
    code {
      background: transparent;
      padding: 0;
    }
  }

  button {
    cursor: pointer;
    font-family: inherit;
    transition: all ${({ theme }) => theme.transitions.normal};
  }

  ul {
    list-style: none;
  }

  /* Custom scrollbar */
  ::-webkit-scrollbar {
    width: 10px;
  }

  ::-webkit-scrollbar-track {
    background: ${({ theme }) => theme.colors.background.secondary};
  }

  ::-webkit-scrollbar-thumb {
    background: ${({ theme }) => theme.colors.primary};
    border-radius: 5px;
  }

  ::-webkit-scrollbar-thumb:hover {
    background: ${({ theme }) => theme.colors.text.accent};
  }

  /* Selection colors */
  ::selection {
    background: ${({ theme }) => theme.colors.primary};
    color: ${({ theme }) => theme.colors.background.primary};
  }

  /* Animations */
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes slideIn {
    from {
      transform: translateX(-100%);
    }
    to {
      transform: translateX(0);
    }
  }

  @keyframes pulse {
    0%, 100% {
      opacity: 1;
    }
    50% {
      opacity: 0.7;
    }
  }
`;

export default GlobalStyles;