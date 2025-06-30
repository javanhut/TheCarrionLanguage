import 'styled-components';

declare module 'styled-components' {
  export interface DefaultTheme {
    colors: {
      primary: string;
      secondary: string;
      background: {
        primary: string;
        secondary: string;
        tertiary: string;
      };
      text: {
        primary: string;
        secondary: string;
        accent: string;
      };
      border: string;
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
    };
    shadows: {
      small: string;
      medium: string;
      large: string;
      glow: string;
    };
    gradients: {
      primary: string;
      secondary: string;
      dark: string;
    };
  }
}