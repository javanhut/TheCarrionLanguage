export const theme = {
  colors: {
    primary: '#06b6d4',      // Modern cyan
    secondary: '#8b5cf6',     // Modern purple
    accent: '#10b981',        // Modern green
    surface: '#1f2937',       // Surface/card background
    background: {
      primary: '#0a0e27',     // Deeper dark blue
      secondary: '#111827',   // Dark gray-blue
      tertiary: '#1e293b',    // Slate dark
      card: '#1f2937',        // Card background
    },
    text: {
      primary: '#f1f5f9',     // Bright white-gray
      secondary: '#94a3b8',   // Medium gray
      muted: '#64748b',       // Muted gray
      accent: '#06b6d4',      // Cyan accent
      inverse: '#0a0e27',     // Dark text for light backgrounds
    },
    border: '#334155',        // Slate border
    borderAccent: '#475569',  // Lighter border
    code: '#0f172a',          // Deep code background
    link: '#06b6d4',
    linkHover: '#22d3ee',
    success: '#10b981',
    error: '#ef4444',
    warning: '#f59e0b',
  },
  fonts: {
    primary: '"Inter", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif',
    code: '"Fira Code", "JetBrains Mono", "Courier New", monospace',
  },
  breakpoints: {
    mobile: '768px',
    tablet: '1024px',
    desktop: '1400px',
  },
  transitions: {
    fast: '0.15s cubic-bezier(0.4, 0, 0.2, 1)',
    normal: '0.3s cubic-bezier(0.4, 0, 0.2, 1)',
    slow: '0.5s cubic-bezier(0.4, 0, 0.2, 1)',
    standard: '0.3s cubic-bezier(0.4, 0, 0.2, 1)',
  },
  shadows: {
    small: '0 1px 2px 0 rgb(0 0 0 / 0.05)',
    medium: '0 4px 6px -1px rgb(0 0 0 / 0.1)',
    large: '0 10px 15px -3px rgb(0 0 0 / 0.1)',
    xl: '0 20px 25px -5px rgb(0 0 0 / 0.1)',
    glow: '0 0 30px rgba(6, 182, 212, 0.3)',
    glowStrong: '0 0 40px rgba(6, 182, 212, 0.5)',
  },
  gradients: {
    primary: 'linear-gradient(135deg, #06b6d4 0%, #8b5cf6 100%)',
    secondary: 'linear-gradient(135deg, #8b5cf6 0%, #ec4899 100%)',
    success: 'linear-gradient(135deg, #10b981 0%, #06b6d4 100%)',
    dark: 'linear-gradient(135deg, #0a0e27 0%, #111827 50%, #1e293b 100%)',
    card: 'linear-gradient(145deg, #1e293b 0%, #1f2937 100%)',
  },
  spacing: {
    xs: '0.5rem',
    sm: '1rem',
    md: '1.5rem',
    lg: '2rem',
    xl: '3rem',
    xxl: '4rem',
  },
  borderRadius: {
    small: '0.375rem',
    sm: '0.375rem',
    medium: '0.5rem',
    md: '0.5rem',
    large: '0.75rem',
    lg: '0.75rem',
    xl: '1rem',
    xxl: '1.5rem',
    full: '9999px',
  },
};