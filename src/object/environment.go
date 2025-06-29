package object

import "github.com/javanhut/TheCarrionLanguage/src/debug"

// environment.go
type Environment struct {
	store       map[string]Object
	outer       *Environment
	debugConfig *debug.Config
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Environment) GetNames() []string {
	names := make([]string, 0)
	for name := range e.store {
		names = append(names, name)
	}
	return names
}

func (e *Environment) GetOuter() *Environment {
	return e.outer
}

// Clone creates a deep copy of the environment to prevent shared references
func (e *Environment) Clone() *Environment {
	clone := NewEnvironment()
	
	// Copy all variables from this environment
	for name, obj := range e.store {
		clone.store[name] = obj
	}
	
	// Recursively clone the outer environment if it exists
	if e.outer != nil {
		clone.outer = e.outer.Clone()
	}
	
	// Copy debug config if present
	if e.debugConfig != nil {
		clone.debugConfig = e.debugConfig
	}
	
	return clone
}

// SetDebugConfig sets the debug configuration for this environment
func (e *Environment) SetDebugConfig(config *debug.Config) {
	e.debugConfig = config
}

// GetDebugConfig returns the debug configuration
func (e *Environment) GetDebugConfig() *debug.Config {
	if e.debugConfig != nil {
		return e.debugConfig
	}
	if e.outer != nil {
		return e.outer.GetDebugConfig()
	}
	return nil
}
