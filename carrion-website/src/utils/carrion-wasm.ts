// TypeScript wrapper for Carrion WASM module

declare global {
  interface Window {
    Go: new () => GoInstance;
    carrionEval: (code: string) => CarrionResult;
    carrionReset: () => void;
    carrionVersion: () => string;
    carrionStdlibStatus: () => StdlibStatus;
  }
}

export interface StdlibStatus {
  loaded: boolean;
  error: string;
}

interface GoInstance {
  argv: string[];
  env: Record<string, string>;
  exit: (code: number) => void;
  importObject: WebAssembly.Imports;
  exited: boolean;
  mem: DataView;
  run(instance: WebAssembly.Instance): Promise<void>;
}

export interface CarrionResult {
  success: boolean;
  error: string;
  output: string;
  result?: string;
}

class CarrionWASM {
  private go: GoInstance | null = null;
  private wasmInstance: WebAssembly.Instance | null = null;
  private isLoading: boolean = false;
  private isReady: boolean = false;
  private readyPromise: Promise<void> | null = null;
  private readyResolve: (() => void) | null = null;

  constructor() {
    // Set up the ready promise
    this.readyPromise = new Promise((resolve) => {
      this.readyResolve = resolve;
    });
  }

  async init(): Promise<void> {
    if (this.isReady) return;
    if (this.isLoading) {
      return this.readyPromise!;
    }

    this.isLoading = true;

    try {
      // Load wasm_exec.js if not already loaded
      if (typeof window.Go === 'undefined') {
        await this.loadScript('/wasm_exec.js');
      }

      // Create Go instance
      this.go = new window.Go();

      // Listen for the carrionReady event
      const readyListener = () => {
        this.isReady = true;
        this.isLoading = false;
        if (this.readyResolve) {
          this.readyResolve();
        }
        window.removeEventListener('carrionReady', readyListener);
      };
      window.addEventListener('carrionReady', readyListener);

      // Load and instantiate WASM
      const wasmResponse = await fetch('/carrion.wasm');
      const wasmBuffer = await wasmResponse.arrayBuffer();
      const result = await WebAssembly.instantiate(wasmBuffer, this.go.importObject);
      this.wasmInstance = result.instance;

      // Run the Go program (this starts the WASM and keeps it alive)
      this.go.run(this.wasmInstance);

      // Wait for ready with timeout
      const timeout = new Promise<void>((_, reject) => {
        setTimeout(() => reject(new Error('WASM initialization timeout')), 10000);
      });

      await Promise.race([this.readyPromise!, timeout]);

    } catch (error) {
      this.isLoading = false;
      throw error;
    }
  }

  private loadScript(src: string): Promise<void> {
    return new Promise((resolve, reject) => {
      const script = document.createElement('script');
      script.src = src;
      script.onload = () => resolve();
      script.onerror = () => reject(new Error(`Failed to load ${src}`));
      document.head.appendChild(script);
    });
  }

  async evaluate(code: string): Promise<CarrionResult> {
    await this.init();

    if (!window.carrionEval) {
      return {
        success: false,
        error: 'Carrion WASM not initialized',
        output: '',
      };
    }

    try {
      return window.carrionEval(code);
    } catch (error) {
      return {
        success: false,
        error: error instanceof Error ? error.message : String(error),
        output: '',
      };
    }
  }

  async reset(): Promise<void> {
    await this.init();
    if (window.carrionReset) {
      window.carrionReset();
    }
  }

  async getVersion(): Promise<string> {
    await this.init();
    if (window.carrionVersion) {
      return window.carrionVersion();
    }
    return 'unknown';
  }

  async getStdlibStatus(): Promise<StdlibStatus> {
    await this.init();
    if (window.carrionStdlibStatus) {
      return window.carrionStdlibStatus();
    }
    return { loaded: false, error: 'WASM not initialized' };
  }

  getIsReady(): boolean {
    return this.isReady;
  }
}

// Singleton instance
export const carrionWasm = new CarrionWASM();
