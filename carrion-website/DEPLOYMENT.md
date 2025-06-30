# Carrion Website Deployment

This website is configured to deploy automatically to GitHub Pages with the custom domain `wiki.carrionlang.com`.

## Deployment Configuration

- **Custom Domain**: `wiki.carrionlang.com` (configured via `public/CNAME`)
- **Routing**: Uses HashRouter for GitHub Pages compatibility
- **Build**: Optimized React production build with 404.html fallback

## Automatic Deployment

The site deploys automatically via GitHub Actions when changes are pushed to the `gh-pages` branch.

### GitHub Actions Workflow
- Triggers on push to `gh-pages` branch
- Builds the React app
- Deploys to GitHub Pages
- Uses the custom domain from CNAME file

## Manual Deployment

To deploy manually:

```bash
# Build the production version
npm run build:gh-pages

# Deploy using gh-pages (if configured)
npm run deploy
```

## Local Development

```bash
# Install dependencies
npm install

# Start development server
npm start

# Build for production
npm run build
```

## Features

- ✅ React 19 with TypeScript
- ✅ Styled Components with theming
- ✅ Framer Motion animations
- ✅ React Router with HashRouter for GitHub Pages
- ✅ Interactive playground (simulation mode in production)
- ✅ Responsive design
- ✅ Custom domain support
- ✅ Automatic deployment

## Playground Notes

- In development: Connects to local API server for real Carrion execution
- In production: Uses simulation mode (no backend required)
- Real execution requires the playground API server to be running

## Domain Configuration

The custom domain `wiki.carrionlang.com` should be configured in your DNS provider to point to GitHub Pages:

```
CNAME wiki.carrionlang.com javanhut.github.io
```

Or for apex domains:
```
A wiki.carrionlang.com 185.199.108.153
A wiki.carrionlang.com 185.199.109.153
A wiki.carrionlang.com 185.199.110.153
A wiki.carrionlang.com 185.199.111.153
```