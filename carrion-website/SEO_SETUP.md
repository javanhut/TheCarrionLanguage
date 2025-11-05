# SEO Setup for wiki.carrionlang.com

This guide explains how to get your Carrion Language wiki indexed by Google and other search engines.

## What Has Been Implemented

### 1. Sitemap Generation
- **File**: `generate-sitemap.js`
- **Purpose**: Automatically generates a sitemap.xml file listing all pages on the site
- **Included URLs**: All main pages and documentation pages from the React Router configuration
- **Build Integration**: Automatically runs during the GitHub Pages deployment workflow

### 2. Robots.txt Configuration
- **File**: `public/robots.txt`
- **Changes**: Added sitemap reference to help search engines discover all pages
- **Allows**: All search engine bots to crawl the entire site

### 3. SEO Meta Tags
- **File**: `public/index.html`
- **Added**:
  - Descriptive title and meta description
  - Keywords for better discoverability
  - Open Graph tags for social media sharing
  - Twitter Card tags
  - Canonical URL
  - Author information

### 4. GitHub Actions Workflow
- **File**: `.github/workflows/deploy-wiki.yml`
- **Change**: Added sitemap generation step to automatically create sitemap.xml during deployment

## Next Steps: Google Search Console Setup

To complete the SEO setup and get your site indexed by Google, follow these steps:

### 1. Verify Domain Ownership

Visit [Google Search Console](https://search.google.com/search-console) and add your property:

1. Click "Add Property"
2. Enter: `wiki.carrionlang.com`
3. Choose verification method (DNS verification recommended):
   - Add TXT record to your DNS settings
   - Or add HTML meta tag to your site
   - Or upload HTML verification file

### 2. Submit Sitemap

Once verified:

1. Go to "Sitemaps" in the left sidebar
2. Enter: `sitemap.xml`
3. Click "Submit"
4. Google will begin crawling your pages

### 3. Request Indexing

For immediate results:

1. Go to "URL Inspection" in the left sidebar
2. Enter: `https://wiki.carrionlang.com/`
3. Click "Request Indexing"
4. Repeat for important pages like:
   - `/docs/getting-started`
   - `/docs/installation`
   - `/docs/quick-start`

### 4. Monitor Performance

Check these sections regularly:

- **Coverage**: See which pages are indexed
- **Performance**: Track clicks, impressions, and rankings
- **Enhancements**: Check for mobile usability and other issues

## Additional SEO Best Practices

### Internal Linking
- Add links between related documentation pages
- Include breadcrumb navigation
- Link from homepage to key documentation pages

### Content Quality
- Keep documentation up-to-date
- Add code examples with proper syntax highlighting
- Include clear headings and subheadings
- Write descriptive alt text for images

### External Promotion
- Link to wiki.carrionlang.com from:
  - GitHub repository README
  - Main carrionlang.com site (if exists)
  - Social media profiles
  - Developer community posts

### Regular Updates
- Update content regularly to signal freshness to Google
- Add new documentation pages as features are released
- Keep the changelog updated

## Expected Timeline

- **Initial Crawl**: 1-3 days after submitting sitemap
- **Full Indexing**: 1-4 weeks for all pages
- **Ranking Improvements**: 2-6 months as authority builds

## Troubleshooting

### Site Not Being Indexed

Check:
1. DNS settings are correct for wiki.carrionlang.com
2. Site is publicly accessible (not password protected)
3. No server errors (check in Search Console)
4. Robots.txt is not blocking Google
5. Sitemap is accessible at https://wiki.carrionlang.com/sitemap.xml

### Low Rankings

Improve by:
1. Adding more unique, quality content
2. Getting backlinks from other sites
3. Improving page load speed
4. Ensuring mobile-friendly design
5. Increasing user engagement (time on site, low bounce rate)

## Testing

Before deploying:

```bash
# Test sitemap generation locally
cd carrion-website
node generate-sitemap.js

# Verify sitemap was created
ls -la build/sitemap.xml
cat build/sitemap.xml

# Test that site builds successfully
npm run build
```

## Monitoring Tools

- **Google Search Console**: Primary tool for monitoring indexing and search performance
- **Google Analytics**: Track visitor behavior and traffic sources
- **PageSpeed Insights**: Monitor and improve page load performance
- **Mobile-Friendly Test**: Ensure mobile compatibility

## Support

For issues or questions about SEO setup:
1. Check Google Search Console Help documentation
2. Review GitHub Pages documentation
3. Check DNS configuration with your domain provider
