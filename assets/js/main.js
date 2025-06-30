// Carrion Language Website JavaScript

// Mobile Navigation Toggle
const mobileMenu = document.getElementById('mobile-menu');
const navMenu = document.querySelector('.nav-menu');

mobileMenu.addEventListener('click', () => {
    navMenu.classList.toggle('active');
    
    // Animate hamburger menu
    const bars = mobileMenu.querySelectorAll('.bar');
    bars[0].style.transform = navMenu.classList.contains('active') ? 'rotate(-45deg) translate(-5px, 6px)' : '';
    bars[1].style.opacity = navMenu.classList.contains('active') ? '0' : '1';
    bars[2].style.transform = navMenu.classList.contains('active') ? 'rotate(45deg) translate(-5px, -6px)' : '';
});

// Close mobile menu when clicking a link
document.querySelectorAll('.nav-link').forEach(link => {
    link.addEventListener('click', () => {
        navMenu.classList.remove('active');
    });
});

// Smooth scrolling for anchor links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            });
        }
    });
});

// Navbar scroll effect
let lastScroll = 0;
const navbar = document.querySelector('.navbar');

window.addEventListener('scroll', () => {
    const currentScroll = window.pageYOffset;
    
    if (currentScroll > 100) {
        navbar.style.background = 'rgba(15, 15, 35, 0.98)';
        navbar.style.boxShadow = '0 5px 20px rgba(0, 0, 0, 0.3)';
    } else {
        navbar.style.background = 'rgba(15, 15, 35, 0.95)';
        navbar.style.boxShadow = 'none';
    }
    
    lastScroll = currentScroll;
});

// Playground functionality
const codeEditor = document.getElementById('code-editor');
const runButton = document.getElementById('run-code');
const outputConsole = document.getElementById('output-console');

// Simple Carrion interpreter simulator
class CarrionSimulator {
    constructor() {
        this.output = [];
        this.variables = {};
        this.functions = {};
        this.classes = {};
    }

    run(code) {
        this.output = [];
        const lines = code.split('\n');
        
        try {
            for (let i = 0; i < lines.length; i++) {
                const line = lines[i].trim();
                
                // Skip empty lines and comments
                if (!line || line.startsWith('//')) continue;
                
                // Handle print statements
                if (line.startsWith('print(')) {
                    const match = line.match(/print\((.*)\)/);
                    if (match) {
                        let value = match[1].trim();
                        
                        // Handle string literals
                        if (value.startsWith('"') || value.startsWith("'")) {
                            value = value.slice(1, -1);
                        }
                        // Handle f-strings
                        else if (value.startsWith('f"') || value.startsWith("f'")) {
                            value = this.evaluateFString(value);
                        }
                        // Handle variables
                        else if (this.variables[value] !== undefined) {
                            value = this.variables[value];
                        }
                        
                        this.output.push(value);
                    }
                }
                
                // Handle variable assignments
                else if (line.includes('=') && !line.includes('==')) {
                    const parts = line.split('=');
                    const varName = parts[0].trim();
                    let value = parts[1].trim();
                    
                    // Handle numeric values
                    if (!isNaN(value)) {
                        value = Number(value);
                    }
                    // Handle string values
                    else if (value.startsWith('"') || value.startsWith("'")) {
                        value = value.slice(1, -1);
                    }
                    // Handle function calls
                    else if (value.includes('(') && value.includes(')')) {
                        value = this.evaluateFunctionCall(value);
                    }
                    
                    this.variables[varName] = value;
                }
                
                // Handle spell (function) definitions
                else if (line.startsWith('spell ')) {
                    const match = line.match(/spell\s+(\w+)\((.*)\):/);
                    if (match) {
                        const funcName = match[1];
                        const params = match[2].split(',').map(p => p.trim());
                        
                        // Find the function body
                        let bodyLines = [];
                        i++;
                        while (i < lines.length && lines[i].startsWith('    ')) {
                            bodyLines.push(lines[i].substring(4));
                            i++;
                        }
                        i--; // Back up one line
                        
                        this.functions[funcName] = { params, body: bodyLines };
                    }
                }
                
                // Handle grimoire (class) definitions
                else if (line.startsWith('grim ')) {
                    const match = line.match(/grim\s+(\w+):/);
                    if (match) {
                        const className = match[1];
                        this.output.push(`Grimoire ${className} defined`);
                        
                        // Skip class body for now
                        i++;
                        while (i < lines.length && (lines[i].startsWith('    ') || lines[i].trim() === '')) {
                            i++;
                        }
                        i--;
                    }
                }
                
                // Handle repeat loops
                else if (line.startsWith('repeat ')) {
                    const match = line.match(/repeat\s+(\d+):/);
                    if (match) {
                        const count = parseInt(match[1]);
                        
                        // Find the loop body
                        let bodyLines = [];
                        i++;
                        while (i < lines.length && lines[i].startsWith('    ')) {
                            bodyLines.push(lines[i]);
                            i++;
                        }
                        i--; // Back up one line
                        
                        // Execute the loop
                        for (let j = 0; j < count; j++) {
                            bodyLines.forEach(bodyLine => {
                                // Recursively process the loop body
                                const tempCode = bodyLine.trim();
                                if (tempCode.startsWith('print(')) {
                                    const tempSimulator = new CarrionSimulator();
                                    tempSimulator.variables = {...this.variables};
                                    tempSimulator.run(tempCode);
                                    this.output.push(...tempSimulator.output);
                                }
                            });
                        }
                    }
                }
            }
            
            if (this.output.length === 0) {
                this.output.push('Code executed successfully with no output.');
            }
        } catch (error) {
            this.output.push(`Error: ${error.message}`);
        }
        
        return this.output.join('\n');
    }

    evaluateFString(fstring) {
        // Simple f-string evaluation
        let result = fstring.slice(2, -1); // Remove f" and "
        
        // Replace variables in {}
        result = result.replace(/\{(\w+)\}/g, (match, varName) => {
            return this.variables[varName] !== undefined ? this.variables[varName] : match;
        });
        
        return result;
    }

    evaluateFunctionCall(call) {
        // Simple function call evaluation
        const match = call.match(/(\w+)\((.*)\)/);
        if (match) {
            const funcName = match[1];
            const args = match[2];
            
            // Handle fibonacci example
            if (funcName === 'fibonacci') {
                const n = parseInt(args);
                return this.fibonacci(n);
            }
        }
        
        return call;
    }

    fibonacci(n) {
        if (n <= 1) return n;
        return this.fibonacci(n - 1) + this.fibonacci(n - 2);
    }
}

// Run code button handler
runButton.addEventListener('click', () => {
    const code = codeEditor.value;
    const simulator = new CarrionSimulator();
    
    outputConsole.textContent = 'Running...\n';
    
    // Simulate execution delay
    setTimeout(() => {
        const result = simulator.run(code);
        outputConsole.textContent = result;
    }, 500);
});

// Add tab support in code editor
codeEditor.addEventListener('keydown', (e) => {
    if (e.key === 'Tab') {
        e.preventDefault();
        const start = codeEditor.selectionStart;
        const end = codeEditor.selectionEnd;
        
        codeEditor.value = codeEditor.value.substring(0, start) + '    ' + codeEditor.value.substring(end);
        codeEditor.selectionStart = codeEditor.selectionEnd = start + 4;
    }
});

// Intersection Observer for animations
const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -100px 0px'
};

const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.style.opacity = '1';
            entry.target.style.transform = 'translateY(0)';
        }
    });
}, observerOptions);

// Observe feature cards and other elements
document.querySelectorAll('.feature-card, .doc-card, .community-card').forEach(card => {
    card.style.opacity = '0';
    card.style.transform = 'translateY(20px)';
    card.style.transition = 'opacity 0.6s ease, transform 0.6s ease';
    observer.observe(card);
});

// Add typing effect to hero subtitle
const subtitleElement = document.querySelector('.hero-subtitle');
const originalText = subtitleElement.textContent;
subtitleElement.textContent = '';

let charIndex = 0;
function typeSubtitle() {
    if (charIndex < originalText.length) {
        subtitleElement.textContent += originalText.charAt(charIndex);
        charIndex++;
        setTimeout(typeSubtitle, 50);
    }
}

// Start typing effect after page load
window.addEventListener('load', () => {
    setTimeout(typeSubtitle, 1000);
});

// Copy code functionality
document.querySelectorAll('pre code').forEach(block => {
    // Create copy button
    const button = document.createElement('button');
    button.className = 'copy-button';
    button.textContent = 'Copy';
    
    button.addEventListener('click', () => {
        navigator.clipboard.writeText(block.textContent).then(() => {
            button.textContent = 'Copied!';
            setTimeout(() => {
                button.textContent = 'Copy';
            }, 2000);
        });
    });
    
    const pre = block.parentElement;
    pre.style.position = 'relative';
    pre.appendChild(button);
});

// Add copy button styles
const style = document.createElement('style');
style.textContent = `
    .copy-button {
        position: absolute;
        top: 10px;
        right: 10px;
        background: var(--accent-primary);
        color: white;
        border: none;
        padding: 5px 10px;
        border-radius: 5px;
        font-size: 12px;
        cursor: pointer;
        opacity: 0;
        transition: opacity 0.3s ease;
    }
    
    pre:hover .copy-button {
        opacity: 1;
    }
    
    .copy-button:hover {
        background: var(--accent-hover);
    }
`;
document.head.appendChild(style);