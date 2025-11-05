const fs = require('fs');
const path = require('path');

const SITE_URL = 'https://wiki.carrionlang.com';
const BUILD_DIR = path.join(__dirname, 'build');

function formatDate(date) {
  return date.toISOString().split('T')[0];
}

const routes = [
  { path: '/', priority: '1.0', changefreq: 'daily' },
  { path: '/features', priority: '0.8', changefreq: 'weekly' },
  { path: '/documentation', priority: '0.9', changefreq: 'weekly' },
  { path: '/playground', priority: '0.7', changefreq: 'monthly' },
  { path: '/download', priority: '0.8', changefreq: 'weekly' },
  { path: '/community', priority: '0.6', changefreq: 'monthly' },
  { path: '/docs/getting-started', priority: '0.9', changefreq: 'weekly' },
  { path: '/docs/installation', priority: '0.9', changefreq: 'weekly' },
  { path: '/docs/quick-start', priority: '0.9', changefreq: 'weekly' },
  { path: '/docs/repl-guide', priority: '0.8', changefreq: 'weekly' },
  { path: '/docs/language-reference', priority: '0.9', changefreq: 'weekly' },
  { path: '/docs/standard-library', priority: '0.8', changefreq: 'weekly' },
  { path: '/docs/grimoires', priority: '0.8', changefreq: 'weekly' },
  { path: '/docs/error-handling', priority: '0.8', changefreq: 'weekly' },
  { path: '/docs/modules', priority: '0.8', changefreq: 'weekly' },
  { path: '/docs/builtin-functions', priority: '0.8', changefreq: 'weekly' },
  { path: '/docs/operators', priority: '0.8', changefreq: 'weekly' },
  { path: '/docs/control-flow', priority: '0.8', changefreq: 'weekly' }
];

function generateSitemap() {
  const currentDate = formatDate(new Date());
  
  let sitemap = '<?xml version="1.0" encoding="UTF-8"?>\n';
  sitemap += '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n';
  
  for (const route of routes) {
    const url = SITE_URL + (route.path === '/' ? '' : route.path);
    
    sitemap += '  <url>\n';
    sitemap += `    <loc>${url}</loc>\n`;
    sitemap += `    <lastmod>${currentDate}</lastmod>\n`;
    sitemap += `    <changefreq>${route.changefreq}</changefreq>\n`;
    sitemap += `    <priority>${route.priority}</priority>\n`;
    sitemap += '  </url>\n';
  }
  
  sitemap += '</urlset>\n';
  
  if (!fs.existsSync(BUILD_DIR)) {
    fs.mkdirSync(BUILD_DIR, { recursive: true });
  }
  
  const sitemapPath = path.join(BUILD_DIR, 'sitemap.xml');
  fs.writeFileSync(sitemapPath, sitemap);
  
  console.log(`Sitemap generated successfully at ${sitemapPath}`);
  console.log(`Total URLs: ${routes.length}`);
}

try {
  generateSitemap();
} catch (error) {
  console.error('Error generating sitemap:', error);
  process.exit(1);
}
