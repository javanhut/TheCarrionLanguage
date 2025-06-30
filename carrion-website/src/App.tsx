import React from 'react';
import { HashRouter as Router, Routes, Route } from 'react-router-dom';
import { ThemeProvider } from 'styled-components';
import GlobalStyles from './styles/GlobalStyles';
import { theme } from './styles/theme';

// Layout Components
import Navbar from './components/layout/Navbar';
import Footer from './components/layout/Footer';

// Page Components
import Home from './pages/Home';
import Features from './pages/Features';
import Documentation from './pages/Documentation';
import Playground from './pages/Playground';
import Download from './pages/Download';
import Community from './pages/Community';

// Documentation Pages
import GettingStarted from './pages/docs/GettingStarted';
import Installation from './pages/docs/Installation';
import QuickStart from './pages/docs/QuickStart';
import LanguageReference from './pages/docs/LanguageReference';
import StandardLibrary from './pages/docs/StandardLibrary';
import Grimoires from './pages/docs/Grimoires';
import ErrorHandling from './pages/docs/ErrorHandling';
import Modules from './pages/docs/Modules';
import BuiltinFunctions from './pages/docs/BuiltinFunctions';
import Operators from './pages/docs/Operators';
import ControlFlow from './pages/docs/ControlFlow';

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Router>
        <GlobalStyles />
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/features" element={<Features />} />
          <Route path="/documentation" element={<Documentation />} />
          <Route path="/playground" element={<Playground />} />
          <Route path="/download" element={<Download />} />
          <Route path="/community" element={<Community />} />
          
          {/* Documentation Routes */}
          <Route path="/docs/getting-started" element={<GettingStarted />} />
          <Route path="/docs/installation" element={<Installation />} />
          <Route path="/docs/quick-start" element={<QuickStart />} />
          <Route path="/docs/language-reference" element={<LanguageReference />} />
          <Route path="/docs/standard-library" element={<StandardLibrary />} />
          <Route path="/docs/grimoires" element={<Grimoires />} />
          <Route path="/docs/error-handling" element={<ErrorHandling />} />
          <Route path="/docs/modules" element={<Modules />} />
          <Route path="/docs/builtin-functions" element={<BuiltinFunctions />} />
          <Route path="/docs/operators" element={<Operators />} />
          <Route path="/docs/control-flow" element={<ControlFlow />} />
        </Routes>
        <Footer />
      </Router>
    </ThemeProvider>
  );
}

export default App;
