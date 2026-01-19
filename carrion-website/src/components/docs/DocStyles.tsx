import styled from 'styled-components';

// Section components
export const Section = styled.section`
  margin-bottom: 3rem;
  scroll-margin-top: 6rem;
`;

export const SectionTitle = styled.h2`
  font-size: 1.75rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 1rem;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  scroll-margin-top: 5rem;
`;

export const SubSection = styled.div`
  margin-bottom: 2rem;
`;

export const SubSectionTitle = styled.h3`
  font-size: 1.25rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.75rem;
  scroll-margin-top: 5rem;
`;

// Text components
export const Paragraph = styled.p`
  font-size: 1rem;
  line-height: 1.75;
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 1rem;
`;

export const Lead = styled.p`
  font-size: 1.15rem;
  line-height: 1.7;
  color: ${({ theme }) => theme.colors.text.secondary};
  margin-bottom: 1.5rem;
`;

// Code components
export const CodeBlock = styled.div`
  margin: 1.25rem 0;
  border-radius: 8px;
  overflow: hidden;
  background: #1a1b26;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);

  pre {
    margin: 0 !important;
    padding: 1.25rem !important;
    font-size: 0.9rem !important;
    line-height: 1.6 !important;
  }
`;

export const InlineCode = styled.code`
  background: rgba(6, 182, 212, 0.1);
  color: ${({ theme }) => theme.colors.primary};
  padding: 0.15rem 0.4rem;
  border-radius: 4px;
  font-family: 'JetBrains Mono', 'Monaco', 'Courier New', monospace;
  font-size: 0.9em;
`;

// Info/Alert boxes
export const InfoBox = styled.div`
  background: rgba(6, 182, 212, 0.08);
  border: 1px solid rgba(6, 182, 212, 0.2);
  border-left: 4px solid ${({ theme }) => theme.colors.primary};
  border-radius: 6px;
  padding: 1rem 1.25rem;
  margin: 1.25rem 0;
`;

export const InfoTitle = styled.h4`
  font-size: 0.9rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.primary};
  margin-bottom: 0.4rem;
`;

export const InfoText = styled.p`
  font-size: 0.95rem;
  line-height: 1.6;
  color: ${({ theme }) => theme.colors.text.secondary};
  margin: 0;
`;

export const WarningBox = styled(InfoBox)`
  background: rgba(251, 191, 36, 0.08);
  border-color: rgba(251, 191, 36, 0.2);
  border-left-color: #fbbf24;
`;

export const WarningTitle = styled(InfoTitle)`
  color: #fbbf24;
`;

export const TipBox = styled(InfoBox)`
  background: rgba(34, 197, 94, 0.08);
  border-color: rgba(34, 197, 94, 0.2);
  border-left-color: #22c55e;
`;

export const TipTitle = styled(InfoTitle)`
  color: #22c55e;
`;

// List components
export const List = styled.ul`
  margin: 1rem 0;
  padding-left: 1.5rem;
  color: ${({ theme }) => theme.colors.text.secondary};
`;

export const ListItem = styled.li`
  margin-bottom: 0.5rem;
  line-height: 1.7;
`;

export const OrderedList = styled.ol`
  margin: 1rem 0;
  padding-left: 1.5rem;
  color: ${({ theme }) => theme.colors.text.secondary};
`;

// Table components
export const Table = styled.table`
  width: 100%;
  border-collapse: collapse;
  margin: 1.25rem 0;
  font-size: 0.95rem;
`;

export const TableHeader = styled.thead`
  background: ${({ theme }) => theme.colors.background.secondary};
`;

export const TableRow = styled.tr`
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
`;

export const TableHead = styled.th`
  padding: 0.75rem 1rem;
  text-align: left;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
`;

export const TableCell = styled.td`
  padding: 0.75rem 1rem;
  color: ${({ theme }) => theme.colors.text.secondary};
`;

// Card components
export const CardGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.25rem;
  margin: 1.5rem 0;
`;

export const Card = styled.div`
  background: ${({ theme }) => theme.colors.background.secondary};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: 10px;
  padding: 1.5rem;
  transition: all 0.2s ease;

  &:hover {
    border-color: ${({ theme }) => theme.colors.primary};
    transform: translateY(-2px);
  }
`;

export const CardTitle = styled.h4`
  font-size: 1.1rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.5rem;
`;

export const CardDescription = styled.p`
  font-size: 0.95rem;
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.6;
`;

// Keyword/Term highlighting
export const Keyword = styled.span`
  background: ${({ theme }) => theme.colors.primary};
  color: ${({ theme }) => theme.colors.text.inverse};
  padding: 0.1rem 0.5rem;
  border-radius: 4px;
  font-weight: 500;
  font-size: 0.9em;
`;

export const Terminology = styled.span`
  color: ${({ theme }) => theme.colors.primary};
  font-weight: 500;
`;

// Feature comparison grid
export const ComparisonTable = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.75rem;
  margin: 1.25rem 0;
`;

export const ComparisonItem = styled.div`
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  background: ${({ theme }) => theme.colors.background.secondary};
  border-radius: 6px;
  font-size: 0.95rem;
`;

export const ComparisonLabel = styled.span`
  color: ${({ theme }) => theme.colors.text.muted};
  min-width: 100px;
`;

export const ComparisonValue = styled.code`
  color: ${({ theme }) => theme.colors.primary};
  font-family: 'JetBrains Mono', 'Monaco', 'Courier New', monospace;
`;

// Tabs component
export const TabContainer = styled.div`
  margin: 1.5rem 0;
`;

export const TabList = styled.div`
  display: flex;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};
  margin-bottom: 1rem;
`;

export const Tab = styled.button<{ $active?: boolean }>`
  padding: 0.75rem 1.25rem;
  background: transparent;
  border: none;
  border-bottom: 2px solid ${({ $active, theme }) => ($active ? theme.colors.primary : 'transparent')};
  color: ${({ $active, theme }) => ($active ? theme.colors.primary : theme.colors.text.secondary)};
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    color: ${({ theme }) => theme.colors.primary};
  }
`;

export const TabContent = styled.div``;

// Divider
export const Divider = styled.hr`
  border: none;
  border-top: 1px solid ${({ theme }) => theme.colors.border};
  margin: 2rem 0;
`;

// Anchor for internal links
export const Anchor = styled.a`
  color: ${({ theme }) => theme.colors.primary};
  text-decoration: none;

  &:hover {
    text-decoration: underline;
  }
`;

// Steps for tutorials
export const StepContainer = styled.div`
  margin: 1.5rem 0;
`;

export const Step = styled.div`
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  padding-bottom: 1.5rem;
  border-bottom: 1px solid ${({ theme }) => theme.colors.border};

  &:last-child {
    border-bottom: none;
    margin-bottom: 0;
    padding-bottom: 0;
  }
`;

export const StepNumber = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  min-width: 32px;
  background: ${({ theme }) => theme.colors.primary};
  color: ${({ theme }) => theme.colors.text.inverse};
  border-radius: 50%;
  font-weight: 600;
  font-size: 0.9rem;
`;

export const StepContent = styled.div`
  flex: 1;
`;

export const StepTitle = styled.h4`
  font-size: 1.1rem;
  font-weight: 600;
  color: ${({ theme }) => theme.colors.text.primary};
  margin-bottom: 0.5rem;
`;

export const StepDescription = styled.div`
  color: ${({ theme }) => theme.colors.text.secondary};
  line-height: 1.7;
`;
