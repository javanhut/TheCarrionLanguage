import React, { useState, useEffect } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import styled from 'styled-components';
// import { FaBars, FaXmark } from 'react-icons/fa6';

const Nav = styled.nav<{ scrolled: boolean }>`
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  background: ${({ scrolled, theme }) => 
    scrolled ? 'rgba(15, 15, 35, 0.98)' : 'rgba(15, 15, 35, 0.95)'};
  backdrop-filter: blur(10px);
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  box-shadow: ${({ scrolled, theme }) => 
    scrolled ? theme.shadows.medium : 'none'};
  transition: all ${({ theme }) => theme.transitions.normal};
`;

const NavContainer = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
  gap: 2rem;
`;

const Logo = styled(Link)`
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.5rem;
  font-weight: bold;
  color: ${({ theme }) => theme.colors.primary};
  text-decoration: none;

  &:hover {
    color: ${({ theme }) => theme.colors.text.accent};
  }
`;

const LogoIcon = styled.span`
  font-size: 2rem;
`;

const CenterSection = styled.div`
  display: flex;
  align-items: center;
  gap: 2rem;
  flex: 1;
  justify-content: center;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    justify-content: flex-end;
  }
`;

const SearchContainer = styled.div`
  position: relative;
  
  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    display: none;
  }
`;

const SearchInput = styled.input`
  padding: 0.5rem 1rem;
  padding-left: 2.5rem;
  border-radius: 20px;
  border: 1px solid ${({ theme }) => theme.colors.border};
  background: ${({ theme }) => theme.colors.background.tertiary};
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: 0.9rem;
  width: 250px;
  transition: all ${({ theme }) => theme.transitions.normal};

  &:focus {
    outline: none;
    border-color: ${({ theme }) => theme.colors.primary};
    box-shadow: 0 0 0 2px rgba(0, 204, 153, 0.1);
    width: 300px;
  }

  &::placeholder {
    color: ${({ theme }) => theme.colors.text.secondary};
  }
`;

const SearchIcon = styled.div`
  position: absolute;
  left: 0.8rem;
  top: 50%;
  transform: translateY(-50%);
  color: ${({ theme }) => theme.colors.text.secondary};
  font-size: 0.9rem;
`;

const SearchResults = styled.div<{ show: boolean }>`
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-top: none;
  border-radius: 0 0 10px 10px;
  max-height: 300px;
  overflow-y: auto;
  display: ${({ show }) => (show ? 'block' : 'none')};
  z-index: 1001;
  box-shadow: ${({ theme }) => theme.shadows.medium};
`;

const SearchResultItem = styled.div`
  padding: 0.8rem 1rem;
  cursor: pointer;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  transition: background ${({ theme }) => theme.transitions.fast};

  &:hover {
    background: ${({ theme }) => theme.colors.background.tertiary};
  }

  &:last-child {
    border-bottom: none;
  }
`;

const ResultTitle = styled.div`
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.2rem;
`;

const ResultPath = styled.div`
  font-size: 0.8rem;
  color: ${({ theme }) => theme.colors.text.secondary};
`;

const NavMenu = styled.ul<{ isOpen: boolean }>`
  display: flex;
  list-style: none;
  gap: 2rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    position: fixed;
    left: ${({ isOpen }) => (isOpen ? '0' : '-100%')};
    top: 70px;
    flex-direction: column;
    background: ${({ theme }) => theme.colors.background.secondary};
    width: 100%;
    text-align: center;
    padding: 2rem 0;
    box-shadow: ${({ theme }) => theme.shadows.large};
    transition: left ${({ theme }) => theme.transitions.normal};
  }
`;

const NavItem = styled.li``;

const NavLink = styled(Link)<{ active: boolean }>`
  color: ${({ active, theme }) => 
    active ? theme.colors.primary : theme.colors.text.primary};
  font-weight: 500;
  position: relative;
  padding: 0.5rem 0;
  transition: color ${({ theme }) => theme.transitions.fast};

  &::after {
    content: '';
    position: absolute;
    bottom: -2px;
    left: 0;
    width: ${({ active }) => (active ? '100%' : '0')};
    height: 2px;
    background: ${({ theme }) => theme.colors.primary};
    transition: width ${({ theme }) => theme.transitions.normal};
  }

  &:hover {
    color: ${({ theme }) => theme.colors.primary};

    &::after {
      width: 100%;
    }
  }
`;

const MobileToggle = styled.div`
  display: none;
  cursor: pointer;
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: 1.5rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    display: block;
  }
`;

const Navbar: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [scrolled, setScrolled] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');
  const [searchResults, setSearchResults] = useState<any[]>([]);
  const [showResults, setShowResults] = useState(false);
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 100);
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  useEffect(() => {
    setIsOpen(false);
  }, [location]);

  const navItems = [
    { path: '/', label: 'Home' },
    { path: '/features', label: 'Features' },
    { path: '/documentation', label: 'Documentation' },
    { path: '/playground', label: 'Playground' },
    { path: '/download', label: 'Download' },
    { path: '/community', label: 'Community' },
  ];

  const searchableContent = [
    { title: 'Getting Started', path: '/docs/getting-started', category: 'Documentation' },
    { title: 'Language Reference', path: '/docs/language-reference', category: 'Documentation' },
    { title: 'Installation Guide', path: '/docs/installation', category: 'Documentation' },
    { title: 'Quick Start', path: '/docs/quick-start', category: 'Documentation' },
    { title: 'Examples', path: '/docs/examples', category: 'Documentation' },
    { title: 'API Documentation', path: '/docs/api', category: 'Documentation' },
    { title: 'Download Latest', path: '/download', category: 'Download' },
    { title: 'Features', path: '/features', category: 'Features' },
    { title: 'Playground', path: '/playground', category: 'Playground' },
    { title: 'Community', path: '/community', category: 'Community' },
    { title: 'Variables and Types', path: '/docs/variables-types', category: 'Documentation' },
    { title: 'Functions (Spells)', path: '/docs/functions', category: 'Documentation' },
    { title: 'Classes (Grimoires)', path: '/docs/classes', category: 'Documentation' },
    { title: 'Error Handling', path: '/docs/error-handling', category: 'Documentation' },
    { title: 'Standard Library (Munin)', path: '/docs/stdlib', category: 'Documentation' },
  ];

  const handleSearch = (query: string) => {
    setSearchQuery(query);
    if (query.trim().length > 1) {
      const filtered = searchableContent.filter(item =>
        item.title.toLowerCase().includes(query.toLowerCase()) ||
        item.category.toLowerCase().includes(query.toLowerCase())
      );
      setSearchResults(filtered.slice(0, 8));
      setShowResults(true);
    } else {
      setSearchResults([]);
      setShowResults(false);
    }
  };

  const handleResultClick = (path: string) => {
    navigate(path);
    setSearchQuery('');
    setShowResults(false);
  };

  const handleSearchKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && searchResults.length > 0) {
      handleResultClick(searchResults[0].path);
    } else if (e.key === 'Escape') {
      setSearchQuery('');
      setShowResults(false);
    }
  };

  return (
    <Nav scrolled={scrolled}>
      <NavContainer>
        <Logo to="/">
          <LogoIcon>🐦‍⬛</LogoIcon>
          Carrion
        </Logo>
        <CenterSection>
          <SearchContainer>
            <SearchIcon>🔍</SearchIcon>
            <SearchInput
              type="text"
              placeholder="Search documentation..."
              value={searchQuery}
              onChange={(e) => handleSearch(e.target.value)}
              onKeyDown={handleSearchKeyDown}
              onBlur={() => setTimeout(() => setShowResults(false), 200)}
              onFocus={() => searchQuery.length > 1 && setShowResults(true)}
            />
            <SearchResults show={showResults}>
              {searchResults.map((result, index) => (
                <SearchResultItem
                  key={index}
                  onClick={() => handleResultClick(result.path)}
                >
                  <ResultTitle>{result.title}</ResultTitle>
                  <ResultPath>{result.category}</ResultPath>
                </SearchResultItem>
              ))}
              {searchResults.length === 0 && searchQuery.length > 1 && (
                <SearchResultItem>
                  <ResultTitle>No results found</ResultTitle>
                  <ResultPath>Try a different search term</ResultPath>
                </SearchResultItem>
              )}
            </SearchResults>
          </SearchContainer>
        </CenterSection>
        <NavMenu isOpen={isOpen}>
          {navItems.map((item) => (
            <NavItem key={item.path}>
              <NavLink
                to={item.path}
                active={location.pathname === item.path}
              >
                {item.label}
              </NavLink>
            </NavItem>
          ))}
        </NavMenu>
        <MobileToggle onClick={() => setIsOpen(!isOpen)}>
          {isOpen ? '✕' : '☰'}
        </MobileToggle>
      </NavContainer>
    </Nav>
  );
};

export default Navbar;