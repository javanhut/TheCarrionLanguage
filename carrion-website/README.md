# Carrion Language Website

A modern, interactive website for the Carrion programming language built with React, TypeScript, and styled-components.

## Features

- **Modern React App**: Built with TypeScript for type safety
- **Interactive Playground**: Try Carrion code directly in the browser
- **Comprehensive Documentation**: Complete guides and tutorials
- **Responsive Design**: Works perfectly on all devices
- **Dark Theme**: Beautiful dark theme with Carrion's signature colors
- **Smooth Animations**: Powered by Framer Motion
- **Syntax Highlighting**: Code examples with proper highlighting

## Tech Stack

- **Frontend**: React 19 + TypeScript
- **Styling**: styled-components for CSS-in-JS
- **Routing**: React Router DOM for navigation
- **Animations**: Framer Motion for smooth transitions
- **Syntax Highlighting**: react-syntax-highlighter
- **Icons**: React Icons (Font Awesome)
- **Deployment**: GitHub Pages

## Development

### Prerequisites

- Node.js 16+
- npm or yarn

### Installation

```bash
# Clone the repository
git clone https://github.com/javanhut/TheCarrionLanguage.git
cd TheCarrionLanguage/carrion-website

# Install dependencies
npm install

# Start development server
npm start
```

The app will be available at `http://localhost:3000`

### Building

```bash
# Build for production
npm run build

# Build for GitHub Pages
npm run build:gh-pages
```

### Deployment

The website is automatically deployed to GitHub Pages on push to the main branch.

Manual deployment:
```bash
npm run deploy
```

## Project Structure

```
src/
├── components/          # Reusable React components
│   ├── layout/         # Layout components (Navbar, Footer)
│   └── common/         # Common UI components
├── pages/              # Page components
│   ├── docs/          # Documentation pages
│   ├── Home.tsx       # Landing page
│   ├── Features.tsx   # Features showcase
│   ├── Playground.tsx # Interactive code editor
│   ├── Download.tsx   # Download page
│   └── Community.tsx  # Community links
├── styles/            # Global styles and theme
│   ├── GlobalStyles.ts
│   └── theme.ts
└── App.tsx           # Main app component
```

## Documentation Pages

The website includes comprehensive documentation converted from Markdown to React components:

- **Getting Started**
  - Installation Guide
  - Quick Start Tutorial
  - Hello World Examples

- **Language Reference**
  - Syntax and Grammar
  - Data Types and Variables
  - Operators and Expressions
  - Control Flow

- **Advanced Features**
  - Grimoires (Classes)
  - Error Handling
  - Modules and Imports

- **Standard Library**
  - Munin Overview
  - Built-in Functions
  - String/Array Grimoires

## Playground

The interactive playground includes:

- **Code Editor**: Syntax-highlighted editor with Carrion code
- **Real-time Execution**: Simulated code execution
- **Example Programs**: Pre-built examples for learning
- **Output Display**: Formatted output with error handling

## Deployment

The website is deployed using GitHub Pages with the following workflow:

1. **Build**: Creates optimized production build
2. **Copy**: Creates 404.html for client-side routing
3. **Deploy**: Pushes to gh-pages branch
4. **Serve**: Available at https://javanhut.github.io/TheCarrionLanguage

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Available Scripts

- `npm start` - Start development server
- `npm build` - Build for production
- `npm test` - Run tests
- `npm run build:gh-pages` - Build with GitHub Pages config
- `npm run deploy` - Deploy to GitHub Pages

## Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## License

MIT License - Same as the Carrion language itself.

## Links

- [Live Website](https://javanhut.github.io/TheCarrionLanguage)
- [GitHub Repository](https://github.com/javanhut/TheCarrionLanguage)
- [Issue Tracker](https://github.com/javanhut/TheCarrionLanguage/issues)
- [Community Discussions](https://github.com/javanhut/TheCarrionLanguage/discussions)