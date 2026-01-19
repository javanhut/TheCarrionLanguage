package object

import "github.com/javanhut/TheCarrionLanguage/src/debug"

// environment.go
type Environment struct {
	store       map[string]Object
	outer       *Environment
	debugConfig *debug.Config
	globalVars  map[string]bool // tracks which variables are declared as global
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	g := make(map[string]bool)
	return &Environment{store: s, outer: nil, globalVars: g}
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
	
	// Copy global variable markers
	for name, isGlobal := range e.globalVars {
		clone.globalVars[name] = isGlobal
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

// MarkGlobal marks a variable as global in the current environment
func (e *Environment) MarkGlobal(name string) {
	if e.globalVars == nil {
		e.globalVars = make(map[string]bool)
	}
	e.globalVars[name] = true
}

// IsGlobal checks if a variable is marked as global in this environment
func (e *Environment) IsGlobal(name string) bool {
	return e.globalVars[name]
}

// SetGlobal sets a variable in the global scope (outermost environment)
func (e *Environment) SetGlobal(name string, val Object) Object {
	// Find the outermost environment (global scope)
	globalEnv := e
	for globalEnv.outer != nil {
		globalEnv = globalEnv.outer
	}
	globalEnv.store[name] = val
	return val
}

// SetWithGlobalCheck sets a variable, checking if it should be set in global scope
func (e *Environment) SetWithGlobalCheck(name string, val Object) Object {
	if e.IsGlobal(name) {
		return e.SetGlobal(name, val)
	}
	return e.Set(name, val)
}

// GetStore returns a copy of the environment's store for external access
func (e *Environment) GetStore() map[string]Object {
	result := make(map[string]Object)
	for name, obj := range e.store {
		result[name] = obj
	}
	return result
}
