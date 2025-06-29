package debug

// Config holds debug configuration settings
type Config struct {
	Enabled   bool
	Lexer     bool
	Parser    bool
	Evaluator bool
}

// NewConfig creates a new debug configuration with default values
func NewConfig() *Config {
	return &Config{
		Enabled:   false,
		Lexer:     false,
		Parser:    false,
		Evaluator: false,
	}
}

// EnableAll enables all debug options
func (c *Config) EnableAll() {
	c.Lexer = true
	c.Parser = true
	c.Evaluator = true
}

// ShouldDebugLexer returns true if lexer debugging is enabled
func (c *Config) ShouldDebugLexer() bool {
	return c.Enabled && c.Lexer
}

// ShouldDebugParser returns true if parser debugging is enabled
func (c *Config) ShouldDebugParser() bool {
	return c.Enabled && c.Parser
}

// ShouldDebugEvaluator returns true if evaluator debugging is enabled
func (c *Config) ShouldDebugEvaluator() bool {
	return c.Enabled && c.Evaluator
}