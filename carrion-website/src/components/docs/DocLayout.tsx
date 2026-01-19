import React, { useState, useEffect } from 'react';
import styled from 'styled-components';
import { Link, useLocation } from 'react-router-dom';

interface INavSection {
  title: string;
  items: NavItem[];
}

interface NavItem {
  label: string;
  path: string;
  children?: { label: string; id: string }[];
}

interface DocLayoutProps {
  children: React.ReactNode;
  title: string;
  description?: string;
  sections?: { id: string; title: string }[];
}

const docNavigation: INavSection[] = [
  {
    title: 'Getting Started',
    items: [
      { label: 'Introduction', path: '/docs/getting-started' },
      { label: 'Installation', path: '/docs/installation' },
      { label: 'Quick Start', path: '/docs/quick-start' },
      { label: 'REPL Guide', path: '/docs/repl-guide' },
    ],
  },
  {
    title: 'Language Basics',
    items: [
      { label: 'Language Reference', path: '/docs/language-reference' },
      { label: 'Operators', path: '/docs/operators' },
      { label: 'Control Flow', path: '/docs/control-flow' },
      { label: 'Error Handling', path: '/docs/error-handling' },
    ],
  },
  {
    title: 'Object-Oriented',
    items: [
      { label: 'Grimoires (Classes)', path: '/docs/grimoires' },
      { label: 'Modules', path: '/docs/modules' },
    ],
  },
  {
    title: 'Standard Library',
    items: [
      { label: 'Munin Overview', path: '/docs/standard-library' },
      { label: 'Built-in Functions', path: '/docs/builtin-functions' },
    ],
  },
];

// Flatten for prev/next navigation
const allPaths = docNavigation.flatMap(section => section.items.map(item => item.path));

const LayoutWrapper = styled.div`
  display: flex;
  min-height: 100vh;
  padding-top: 4.5rem;
`;

const Sidebar = styled.aside`
  width: 280px;
  min-width: 280px;
  background: ${({ theme }) => theme.colors.background.secondary};
  border-right: 1px solid ${({ theme }) => theme.colors.border};
  padding: 2rem 0;
  position: sticky;
  top: 4.5rem;
  height: calc(100vh - 4.5rem);
  overflow-y: auto;

  &::-webkit-scrollbar {
    width: 4px;
  }

  &::-webkit-scrollbar-track {
    background: transparent;
  }

  &::-webkit-scrollbar-thumb {
    background: ${({ theme }) => theme.colors.border};
    border-radius: 2px;
  }

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    display: none;
  }
`;

const MobileSidebarToggle = styled.button`
  display: none;
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  width: 56px;
  height: 56px;
  border-radius: 50%;
  background: ${({ theme }) => theme.colors.primary};
  color: white;
  border: none;
  cursor: pointer;
  z-index: 100;
  box-shadow: ${({ theme }) => theme.shadows.large};
  font-size: 1.5rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    display: flex;
    align-items: center;
    justify-content: center;
  }
`;

const MobileSidebar = styled.div<{ $isOpen: boolean }>`
  display: none;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    display: ${({ $isOpen }) => ($isOpen ? 'block' : 'none')};
    position: fixed;
    top: 4.5rem;
    left: 0;
    right: 0;
    bottom: 0;
    background: ${({ theme }) => theme.colors.background.secondary};
    z-index: 99;
    overflow-y: auto;
    padding: 2rem;
  }
`;

const NavSection = styled.div`
  margin-bottom: 1.5rem;
`;

const NavSectionTitle = styled.h3`
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: ${({ theme }) => theme.colors.text.muted};
  padding: 0 1.5rem;
  margin-bottom: 0.5rem;
`;

const NavList = styled.ul`
  list-style: none;
  padding: 0;
  margin: 0;
`;

const NavListItem = styled.li``;

const NavLink = styled(Link)<{ $active?: boolean }>`
  display: block;
  padding: 0.5rem 1.5rem;
  color: ${({ theme, $active }) => ($active ? theme.colors.primary : theme.colors.text.secondary)};
  text-decoration: none;
  font-size: 0.9rem;
  transition: all 0.2s ease;
  background: ${({ $active }) => ($active ? 'rgba(6, 182, 212, 0.1)' : 'transparent')};
  border-left: 3px solid ${({ theme, $active }) => ($active ? theme.colors.primary : 'transparent')};

  &:hover {
    color: ${({ theme }) => theme.colors.primary};
    background: rgba(6, 182, 212, 0.05);
  }
`;

const MainContent = styled.main`
  flex: 1;
  max-width: 900px;
  padding: 2rem 3rem 4rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.tablet}) {
    padding: 1.5rem;
  }
`;

const ContentWrapper = styled.div`
  display: flex;
  flex: 1;
`;

const TableOfContents = styled.aside`
  width: 220px;
  min-width: 220px;
  padding: 2rem 1rem;
  position: sticky;
  top: 4.5rem;
  height: calc(100vh - 4.5rem);
  overflow-y: auto;

  @media (max-width: ${({ theme }) => theme.breakpoints.desktop}) {
    display: none;
  }
`;

const TOCTitle = styled.h4`
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: ${({ theme }) => theme.colors.text.muted};
  margin-bottom: 1rem;
`;

const TOCList = styled.ul`
  list-style: none;
  padding: 0;
  margin: 0;
`;

const TOCItem = styled.li``;

const TOCLink = styled.button<{ $active?: boolean }>`
  display: block;
  padding: 0.35rem 0;
  padding-left: 0.75rem;
  color: ${({ theme, $active }) => ($active ? theme.colors.primary : theme.colors.text.muted)};
  text-decoration: none;
  font-size: 0.85rem;
  border: none;
  border-left: 2px solid ${({ theme, $active }) => ($active ? theme.colors.primary : theme.colors.border)};
  background: transparent;
  cursor: pointer;
  text-align: left;
  width: 100%;
  transition: all 0.2s ease;

  &:hover {
    color: ${({ theme }) => theme.colors.primary};
    border-left-color: ${({ theme }) => theme.colors.primary};
  }
`;

const PageHeader = styled.header`
  margin-bottom: 3rem;
  padding-bottom: 2rem;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
`;

const PageTitle = styled.h1`
  font-size: 2.5rem;
  font-weight: 700;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.75rem;

  @media (max-width: ${({ theme }) => theme.breakpoints.mobile}) {
    font-size: 2rem;
  }
`;

const PageDescription = styled.p`
  font-size: 1.15rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.6;
`;

const PrevNextNav = styled.div`
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  margin-top: 4rem;
  padding-top: 2rem;
  border-top: 1px solid ${({ theme }) => theme.colors.border};
`;

const PrevNextLink = styled(Link)<{ $direction: 'prev' | 'next' }>`
  display: flex;
  flex-direction: column;
  padding: 1rem 1.5rem;
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 8px;
  text-decoration: none;
  transition: all 0.2s ease;
  text-align: ${({ $direction }) => ($direction === 'next' ? 'right' : 'left')};
  flex: 1;
  max-width: 50%;

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-2px);
  }
`;

const PrevNextLabel = styled.span`
  font-size: 0.75rem;
  color: ${({ theme }) => theme.colors.text.muted};
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 0.25rem;
`;

const PrevNextTitle = styled.span`
  font-size: 1rem;
  font-weight: 500;
  color: ${({ theme }) => theme.colors.primary};
`;

const DocLayout: React.FC<DocLayoutProps> = ({ children, title, description, sections }) => {
  const location = useLocation();
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [activeSection, setActiveSection] = useState<string>('');

  const currentIndex = allPaths.indexOf(location.pathname);
  const prevPath = currentIndex > 0 ? allPaths[currentIndex - 1] : null;
  const nextPath = currentIndex < allPaths.length - 1 ? allPaths[currentIndex + 1] : null;

  const findLabel = (path: string): string => {
    for (const section of docNavigation) {
      const item = section.items.find(i => i.path === path);
      if (item) return item.label;
    }
    return '';
  };

  useEffect(() => {
    setMobileMenuOpen(false);
  }, [location.pathname]);

  const scrollToSection = (id: string) => {
    const element = document.getElementById(id);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
    }
  };

  useEffect(() => {
    if (!sections || sections.length === 0) return;

    const handleScroll = () => {
      const sectionElements = sections.map(s => document.getElementById(s.id));
      const scrollPosition = window.scrollY + 100;

      for (let i = sectionElements.length - 1; i >= 0; i--) {
        const element = sectionElements[i];
        if (element && element.offsetTop <= scrollPosition) {
          setActiveSection(sections[i].id);
          return;
        }
      }
      setActiveSection(sections[0]?.id || '');
    };

    window.addEventListener('scroll', handleScroll);
    handleScroll();

    return () => window.removeEventListener('scroll', handleScroll);
  }, [sections]);

  const renderNav = () => (
    <>
      {docNavigation.map((section, idx) => (
        <NavSection key={idx}>
          <NavSectionTitle>{section.title}</NavSectionTitle>
          <NavList>
            {section.items.map((item, itemIdx) => (
              <NavListItem key={itemIdx}>
                <NavLink to={item.path} $active={location.pathname === item.path}>
                  {item.label}
                </NavLink>
              </NavListItem>
            ))}
          </NavList>
        </NavSection>
      ))}
    </>
  );

  return (
    <LayoutWrapper>
      <Sidebar>{renderNav()}</Sidebar>

      <MobileSidebar $isOpen={mobileMenuOpen}>
        {renderNav()}
      </MobileSidebar>

      <ContentWrapper>
        <MainContent>
          <PageHeader>
            <PageTitle>{title}</PageTitle>
            {description && <PageDescription>{description}</PageDescription>}
          </PageHeader>

          {children}

          <PrevNextNav>
            {prevPath ? (
              <PrevNextLink to={prevPath} $direction="prev">
                <PrevNextLabel>Previous</PrevNextLabel>
                <PrevNextTitle>{findLabel(prevPath)}</PrevNextTitle>
              </PrevNextLink>
            ) : (
              <div />
            )}
            {nextPath ? (
              <PrevNextLink to={nextPath} $direction="next">
                <PrevNextLabel>Next</PrevNextLabel>
                <PrevNextTitle>{findLabel(nextPath)}</PrevNextTitle>
              </PrevNextLink>
            ) : (
              <div />
            )}
          </PrevNextNav>
        </MainContent>

        {sections && sections.length > 0 && (
          <TableOfContents>
            <TOCTitle>On This Page</TOCTitle>
            <TOCList>
              {sections.map((section) => (
                <TOCItem key={section.id}>
                  <TOCLink
                    onClick={() => scrollToSection(section.id)}
                    $active={activeSection === section.id}
                  >
                    {section.title}
                  </TOCLink>
                </TOCItem>
              ))}
            </TOCList>
          </TableOfContents>
        )}
      </ContentWrapper>

      <MobileSidebarToggle onClick={() => setMobileMenuOpen(!mobileMenuOpen)}>
        {mobileMenuOpen ? 'X' : '='}
      </MobileSidebarToggle>
    </LayoutWrapper>
  );
};

export default DocLayout;
