import 'styled-components';

declare module 'styled-components' {
  export interface DefaultTheme {
    colors: {
      primary: string;
      secondary: string;
      accent: string;
      surface: string;
      background: {
        primary: string;
        secondary: string;
        tertiary: string;
        card: string;
      };
      text: {
        primary: string;
        secondary: string;
        muted: string;
        accent: string;
      };
      border: string;
      borderAccent: string;
      code: string;
      link: string;
      linkHover: string;
      success: string;
      error: string;
      warning: string;
    };
    fonts: {
      primary: string;
      code: string;
    };
    breakpoints: {
      mobile: string;
      tablet: string;
      desktop: string;
    };
    transitions: {
      fast: string;
      normal: string;
      slow: string;
      standard?: string;
    };
    shadows: {
      small: string;
      medium: string;
      large: string;
      xl: string;
      glow: string;
      glowStrong: string;
    };
    gradients: {
      primary: string;
      secondary: string;
      success: string;
      dark: string;
      card: string;
    };
    spacing: {
      xs: string;
      sm: string;
      md: string;
      lg: string;
      xl: string;
      xxl: string;
    };
    borderRadius: {
      small: string;
      sm: string;
      medium: string;
      md: string;
      large: string;
      lg: string;
      xl: string;
      xxl: string;
      full: string;
    };
  }
}