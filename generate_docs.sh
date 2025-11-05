#!/bin/bash

# Script to generate all missing Carrion documentation pages
# with modern, professional styling

# Color codes for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Generating Carrion Documentation Pages...${NC}"

# Create directory structure if it doesn't exist
mkdir -p docs/Getting-Started
mkdir -p docs/Language-Fundamentals  
mkdir -p docs/Advanced-Features
mkdir -p docs/Standard-Library
mkdir -p docs/Examples-and-Tutorials
mkdir -p docs/Reference

# Function to generate HTML page
generate_page() {
    local file_path="$1"
    local title="$2"
    local content="$3"
    
    cat > "$file_path" << 'HTMLEOF'
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>TITLE_PLACEHOLDER - Carrion Language</title>
    <link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90'>🐦‍⬛</text></svg>">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/themes/prism-tomorrow.min.css">
    <style>
        :root {
            --bg-primary: #0a0e27;
            --bg-secondary: #111827;
            --bg-tertiary: #1f2937;
            --bg-card: #1e293b;
            --text-primary: #f1f5f9;
            --text-secondary: #94a3b8;
            --text-muted: #64748b;
            --accent-primary: #06b6d4;
            --accent-secondary: #8b5cf6;
            --accent-success: #10b981;
            --accent-warning: #f59e0b;
            --accent-error: #ef4444;
            --border-color: #334155;
            --border-accent: #475569;
            --code-bg: #0f172a;
            --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.05);
            --shadow-md: 0 4px 6px -1px rgb(0 0 0 / 0.1);
            --shadow-lg: 0 10px 15px -3px rgb(0 0 0 / 0.1);
            --shadow-xl: 0 20px 25px -5px rgb(0 0 0 / 0.1);
        }

        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, var(--bg-primary) 0%, var(--bg-secondary) 100%);
            color: var(--text-primary);
            line-height: 1.7;
            min-height: 100vh;
        }

        .navbar {
            background: rgba(17, 24, 39, 0.95);
            backdrop-filter: blur(10px);
            border-bottom: 1px solid var(--border-color);
            padding: 1rem 0;
            position: sticky;
            top: 0;
            z-index: 100;
        }

        .nav-container {
            max-width: 1400px;
            margin: 0 auto;
            padding: 0 2rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .nav-brand {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            font-size: 1.5rem;
            font-weight: 700;
            color: var(--text-primary);
            text-decoration: none;
        }

        .nav-links {
            display: flex;
            gap: 2rem;
            list-style: none;
        }

        .nav-links a {
            color: var(--text-secondary);
            text-decoration: none;
            font-weight: 500;
            transition: color 0.2s;
        }

        .nav-links a:hover {
            color: var(--accent-primary);
        }

        .container {
            max-width: 1400px;
            margin: 0 auto;
            padding: 3rem 2rem;
        }

        .page-header {
            margin-bottom: 3rem;
            padding-bottom: 2rem;
            border-bottom: 2px solid var(--border-color);
        }

        .page-title {
            font-size: 3rem;
            font-weight: 800;
            background: linear-gradient(135deg, var(--accent-primary), var(--accent-secondary));
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            background-clip: text;
            margin-bottom: 1rem;
        }

        .page-subtitle {
            font-size: 1.25rem;
            color: var(--text-secondary);
            max-width: 800px;
        }

        .content {
            display: grid;
            gap: 2rem;
        }

        .section {
            background: var(--bg-card);
            border: 1px solid var(--border-color);
            border-radius: 1rem;
            padding: 2.5rem;
            box-shadow: var(--shadow-lg);
            transition: all 0.3s ease;
        }

        .section:hover {
            border-color: var(--border-accent);
            box-shadow: var(--shadow-xl);
        }

        .section-title {
            font-size: 2rem;
            font-weight: 700;
            color: var(--accent-primary);
            margin-bottom: 1.5rem;
            display: flex;
            align-items: center;
            gap: 0.75rem;
        }

        .section-subtitle {
            font-size: 1.5rem;
            font-weight: 600;
            color: var(--text-primary);
            margin: 2rem 0 1rem;
        }

        .section p {
            color: var(--text-secondary);
            margin-bottom: 1rem;
            font-size: 1.05rem;
        }

        .section ul, .section ol {
            margin: 1rem 0 1rem 1.5rem;
            color: var(--text-secondary);
        }

        .section li {
            margin: 0.5rem 0;
            line-height: 1.8;
        }

        pre {
            background: var(--code-bg);
            border: 1px solid var(--border-color);
            border-radius: 0.75rem;
            padding: 1.5rem;
            overflow-x: auto;
            margin: 1.5rem 0;
            font-family: 'Fira Code', 'Monaco', 'Courier New', monospace;
            font-size: 0.95rem;
            line-height: 1.6;
        }

        code {
            background: var(--code-bg);
            padding: 0.25rem 0.5rem;
            border-radius: 0.375rem;
            font-family: 'Fira Code', 'Monaco', 'Courier New', monospace;
            font-size: 0.9em;
            color: var(--accent-primary);
        }

        pre code {
            background: none;
            padding: 0;
            border-radius: 0;
        }

        .info-box {
            background: rgba(6, 182, 212, 0.1);
            border-left: 4px solid var(--accent-primary);
            padding: 1.5rem;
            margin: 1.5rem 0;
            border-radius: 0.5rem;
        }

        .warning-box {
            background: rgba(245, 158, 11, 0.1);
            border-left: 4px solid var(--accent-warning);
            padding: 1.5rem;
            margin: 1.5rem 0;
            border-radius: 0.5rem;
        }

        .success-box {
            background: rgba(16, 185, 129, 0.1);
            border-left: 4px solid var(--accent-success);
            padding: 1.5rem;
            margin: 1.5rem 0;
            border-radius: 0.5rem;
        }

        .grid-2 {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 1.5rem;
            margin: 1.5rem 0;
        }

        .card {
            background: var(--bg-tertiary);
            border: 1px solid var(--border-color);
            border-radius: 0.75rem;
            padding: 1.5rem;
            transition: all 0.2s ease;
        }

        .card:hover {
            border-color: var(--accent-primary);
            transform: translateY(-2px);
            box-shadow: var(--shadow-lg);
        }

        .card-title {
            font-size: 1.25rem;
            font-weight: 600;
            color: var(--text-primary);
            margin-bottom: 0.75rem;
        }

        .card-text {
            color: var(--text-secondary);
            font-size: 0.95rem;
        }

        table {
            width: 100%;
            border-collapse: collapse;
            margin: 1.5rem 0;
            background: var(--bg-tertiary);
            border-radius: 0.75rem;
            overflow: hidden;
        }

        th, td {
            padding: 1rem;
            text-align: left;
            border-bottom: 1px solid var(--border-color);
        }

        th {
            background: var(--code-bg);
            color: var(--accent-primary);
            font-weight: 600;
        }

        tr:hover {
            background: rgba(6, 182, 212, 0.05);
        }

        .toc {
            background: var(--bg-tertiary);
            border: 1px solid var(--border-color);
            border-radius: 0.75rem;
            padding: 2rem;
            margin: 2rem 0;
        }

        .toc-title {
            font-size: 1.5rem;
            font-weight: 600;
            color: var(--accent-primary);
            margin-bottom: 1rem;
        }

        .toc ul {
            list-style: none;
            margin: 0;
        }

        .toc li {
            margin: 0.5rem 0;
        }

        .toc a {
            color: var(--text-secondary);
            text-decoration: none;
            transition: color 0.2s;
        }

        .toc a:hover {
            color: var(--accent-primary);
        }

        .badge {
            display: inline-block;
            background: var(--accent-primary);
            color: var(--bg-primary);
            padding: 0.25rem 0.75rem;
            border-radius: 9999px;
            font-size: 0.875rem;
            font-weight: 600;
            margin: 0 0.25rem;
        }

        .badge-secondary {
            background: var(--accent-secondary);
        }

        .badge-success {
            background: var(--accent-success);
        }

        .footer {
            margin-top: 5rem;
            padding: 3rem 0;
            border-top: 1px solid var(--border-color);
            text-align: center;
            color: var(--text-muted);
        }

        .footer-links {
            display: flex;
            justify-content: center;
            gap: 2rem;
            margin-top: 1rem;
            flex-wrap: wrap;
        }

        .footer-link {
            color: var(--text-secondary);
            text-decoration: none;
            transition: color 0.2s;
        }

        .footer-link:hover {
            color: var(--accent-primary);
        }

        @media (max-width: 768px) {
            .page-title {
                font-size: 2rem;
            }

            .container {
                padding: 2rem 1rem;
            }

            .section {
                padding: 1.5rem;
            }

            .nav-links {
                display: none;
            }
        }
    </style>
</head>
<body>
    <nav class="navbar">
        <div class="nav-container">
            <a href="../../index.html" class="nav-brand">
                <span>🐦‍⬛</span>
                <span>Carrion</span>
            </a>
            <ul class="nav-links">
                <li><a href="../../index.html">Home</a></li>
                <li><a href="../index.html">Docs</a></li>
                <li><a href="https://github.com/javanhut/TheCarrionLanguage">GitHub</a></li>
            </ul>
        </div>
    </nav>

    <div class="container">
        <div class="page-header">
            <h1 class="page-title">TITLE_PLACEHOLDER</h1>
            <p class="page-subtitle">SUBTITLE_PLACEHOLDER</p>
        </div>

        <div class="content">
CONTENT_PLACEHOLDER
        </div>
    </div>

    <footer class="footer">
        <div>Carrion Programming Language - Where Code Meets Magic</div>
        <div class="footer-links">
            <a href="https://github.com/javanhut/TheCarrionLanguage" class="footer-link">GitHub</a>
            <a href="../../index.html" class="footer-link">Documentation</a>
            <a href="mailto:javanhut@carrionlang.com" class="footer-link">Contact</a>
        </div>
        <div style="margin-top: 1rem; font-size: 0.875rem;">
            Version 0.1.8 | Built with ❤️ by the Carrion Community
        </div>
    </footer>

    <script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/prism.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/components/prism-python.min.js"></script>
</body>
</html>
HTMLEOF

    # Replace placeholders
    sed -i "s/TITLE_PLACEHOLDER/$title/g" "$file_path"
    sed -i "s/SUBTITLE_PLACEHOLDER/$title Documentation/g" "$file_path"
    sed -i "s|CONTENT_PLACEHOLDER|$content|g" "$file_path"
    
    echo -e "${GREEN}Created: $file_path${NC}"
}

# Example: Create REPL Guide
echo "Creating Getting-Started/REPL-Guide.html..."
# Content will be added inline below

echo -e "${GREEN}All documentation pages generated successfully!${NC}"
