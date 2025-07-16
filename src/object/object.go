package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
	"sync"

	"github.com/javanhut/TheCarrionLanguage/src/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NONE_OBJ         = "NONE"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	ARRAY_OBJ        = "ARRAY"
	BUILTIN_OBJ      = "BUILTIN"
	MAP_OBJ          = "MAP"
	TUPLE_OBJ        = "TUPLE"
	GRIMOIRE_OBJ         = "GRIMOIRE"
	INSTANCE_OBJ         = "INSTANCE"
	NAMESPACE_OBJ        = "NAMESPACE"
	GOROUTINE_OBJ        = "GOROUTINE"
	GOROUTINE_MANAGER_OBJ = "GOROUTINE_MANAGER"
	CAUGHT_ERROR_OBJ = "CAUGHT_ERROR"
	STOP_OBJ         = "STOP"
	SKIP_OBJ         = "SKIP"
	SUPER_OBJ        = "SUPER"
)

var NONE = &None{}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Float struct {
	Value float64
}

func (f *Float) Type() ObjectType { return FLOAT_OBJ }
func (f *Float) Inspect() string  { return fmt.Sprintf("%f", f.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type None struct {
	Value string
}

func (n *None) Type() ObjectType { return NONE_OBJ }
func (n *None) Inspect() string  { return fmt.Sprintf("%s", n.Value) }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Function struct {
   // Parameters holds function parameters, either simple identifiers or full Parameter nodes
   Parameters  []ast.Expression
   ReturnType  ast.Expression
	Body        *ast.BlockStatement
	Env         *Environment
	IsAbstract  bool
	IsPrivate   bool
	IsProtected bool
}

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("spell(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elems := make([]string, 0, len(ao.Elements))
	for _, e := range ao.Elements {
		elems = append(elems, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")
	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

func (f *Float) HashKey() HashKey {
	// Convert float to its bit representation for hashing
	bits := uint64(f.Value * 1000000) // Multiply by 1M to preserve 6 decimal places
	return HashKey{Type: f.Type(), Value: bits}
}

type HashPair struct {
	Key   Object
	Value Object
}
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return MAP_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type Hashable interface {
	HashKey() HashKey
}

type Tuple struct {
	Elements []Object
}

func (t *Tuple) Type() ObjectType { return TUPLE_OBJ }
func (t *Tuple) Inspect() string {
	var out bytes.Buffer
	elems := []string{}
	for _, e := range t.Elements {
		elems = append(elems, e.Inspect())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString(")")
	return out.String()
}

// Update Object (object.go)
type Grimoire struct {
	Name       string
	Methods    map[string]*Function
	InitMethod *Function
	Inherits   *Grimoire
	Env        *Environment // Add environment to store the grimoire's scope
	IsArcane   bool
}

func (s *Grimoire) Type() ObjectType { return GRIMOIRE_OBJ }
func (s *Grimoire) Inspect() string {
	return fmt.Sprintf("<grimoire %s>", s.Name)
}

// Ensure Instance type implements Object
type Instance struct {
	Grimoire *Grimoire
	Env      *Environment
}

func (i *Instance) Type() ObjectType { return INSTANCE_OBJ }

// Super represents a reference to the parent class for method calls
// Used for super.method() calls in inheritance hierarchies
type Super struct {
	Instance *Instance // The instance calling super
	Parent   *Grimoire // The parent grimoire to call methods on
}

func (s *Super) Type() ObjectType { return SUPER_OBJ }
func (s *Super) Inspect() string { return "super" }

// CaughtError wraps an error that has been caught by an ensnare clause
// This prevents it from being treated as a propagatable error
type CaughtError struct {
	OriginalError Object
}

func (ce *CaughtError) Type() ObjectType { return CAUGHT_ERROR_OBJ }
func (ce *CaughtError) Inspect() string { return ce.OriginalError.Inspect() }

// GetMessage returns the error message
func (ce *CaughtError) GetMessage() string {
	if errWithTrace, ok := ce.OriginalError.(*ErrorWithTrace); ok {
		return errWithTrace.Message
	}
	if customErr, ok := ce.OriginalError.(*CustomError); ok {
		return customErr.Message
	}
	if err, ok := ce.OriginalError.(*Error); ok {
		return err.Message
	}
	return ce.OriginalError.Inspect()
}
func (i *Instance) Inspect() string {
	// Special handling for primitive wrapper instances
	switch i.Grimoire.Name {
	case "Integer", "Float", "String", "Boolean":
		if value, ok := i.Env.Get("value"); ok {
			return value.Inspect()
		}
	case "Array":
		if elements, ok := i.Env.Get("elements"); ok {
			if arr, isArray := elements.(*Array); isArray {
				return arr.Inspect()
			}
		}
	}
	
	// Check if the instance has a to_string method
	if _, ok := i.Grimoire.Methods["to_string"]; ok {
		// We would need the evaluator to call this method properly
		// For now, fall through to default behavior
	}
	
	return fmt.Sprintf("<instance of %s>", i.Grimoire.Name)
}

// object/object.go

type Namespace struct {
	Env *Environment // Holds all exported members of the imported module
}

func (n *Namespace) Type() ObjectType { return NAMESPACE_OBJ }
func (n *Namespace) Inspect() string  { return "<namespace>" }

type Stop struct{}

func (s *Stop) Type() ObjectType { return STOP_OBJ }
func (s *Stop) Inspect() string  { return "stop" }

type Skip struct{}

func (s *Skip) Type() ObjectType { return SKIP_OBJ }
func (s *Skip) Inspect() string  { return "skip" }

var (
	STOP = &Stop{}
	SKIP = &Skip{}
)

// Goroutine represents a running goroutine in Carrion
type Goroutine struct {
	Name      string
	Done      chan bool
	Result    Object
	Error     Object
	IsRunning bool
	cleaned   bool // Track if cleanup has been performed
}

func (g *Goroutine) Type() ObjectType { return GOROUTINE_OBJ }
func (g *Goroutine) Inspect() string {
	if g.Name != "" {
		return fmt.Sprintf("goroutine(%s)", g.Name)
	}
	return "goroutine(anonymous)"
}

// Cleanup closes channels and releases resources when a goroutine finishes
func (g *Goroutine) Cleanup() {
	if g.cleaned {
		return // Already cleaned up
	}
	g.cleaned = true
	
	// Close the Done channel if it's not already closed
	if g.Done != nil {
		select {
		case <-g.Done:
			// Channel already closed or has a value
		default:
			close(g.Done)
		}
	}
	
	// Mark as not running
	g.IsRunning = false
	
	// Clear references to help GC
	g.Result = nil
	g.Error = nil
}

// IsCompleted checks if the goroutine has finished execution
func (g *Goroutine) IsCompleted() bool {
	if g.Done == nil {
		return true
	}
	
	select {
	case <-g.Done:
		return true
	default:
		return false
	}
}

// GoroutineManager manages all active goroutines
type GoroutineManager struct {
	mu                sync.RWMutex
	Goroutines        map[string]*Goroutine
	Anonymous         []*Goroutine
	MaxNamedSize      int  // Maximum number of named goroutines (0 = unlimited)
	MaxAnonymousSize  int  // Maximum number of anonymous goroutines (0 = unlimited)
	AutoCleanup       bool // Whether to automatically clean up completed goroutines
}

func NewGoroutineManager() *GoroutineManager {
	return &GoroutineManager{
		Goroutines:       make(map[string]*Goroutine),
		Anonymous:        make([]*Goroutine, 0),
		MaxNamedSize:     0,    // Unlimited by default
		MaxAnonymousSize: 0,    // Unlimited by default
		AutoCleanup:      true, // Enable auto cleanup by default
	}
}

// NewGoroutineManagerWithLimits creates a new GoroutineManager with specified limits
func NewGoroutineManagerWithLimits(maxNamed, maxAnonymous int, autoCleanup bool) *GoroutineManager {
	return &GoroutineManager{
		Goroutines:       make(map[string]*Goroutine),
		Anonymous:        make([]*Goroutine, 0),
		MaxNamedSize:     maxNamed,
		MaxAnonymousSize: maxAnonymous,
		AutoCleanup:      autoCleanup,
	}
}

func (gm *GoroutineManager) Type() ObjectType { return GOROUTINE_MANAGER_OBJ }
func (gm *GoroutineManager) Inspect() string {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	return fmt.Sprintf("GoroutineManager(named: %d, anonymous: %d)", 
		len(gm.Goroutines), len(gm.Anonymous))
}

// AddNamedGoroutine adds a named goroutine to the manager
func (gm *GoroutineManager) AddNamedGoroutine(name string, goroutine *Goroutine) error {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	
	// Auto cleanup if enabled
	if gm.AutoCleanup {
		gm.cleanupCompletedLocked()
	}
	
	// Check size limits
	if gm.MaxNamedSize > 0 && len(gm.Goroutines) >= gm.MaxNamedSize {
		return fmt.Errorf("named goroutine limit reached: %d", gm.MaxNamedSize)
	}
	
	gm.Goroutines[name] = goroutine
	return nil
}

// AddAnonymousGoroutine adds an anonymous goroutine to the manager
func (gm *GoroutineManager) AddAnonymousGoroutine(goroutine *Goroutine) error {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	
	// Auto cleanup if enabled
	if gm.AutoCleanup {
		gm.cleanupCompletedLocked()
	}
	
	// Check size limits
	if gm.MaxAnonymousSize > 0 && len(gm.Anonymous) >= gm.MaxAnonymousSize {
		return fmt.Errorf("anonymous goroutine limit reached: %d", gm.MaxAnonymousSize)
	}
	
	gm.Anonymous = append(gm.Anonymous, goroutine)
	return nil
}

// GetNamedGoroutine retrieves a named goroutine from the manager
func (gm *GoroutineManager) GetNamedGoroutine(name string) (*Goroutine, bool) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	goroutine, exists := gm.Goroutines[name]
	return goroutine, exists
}

// RemoveNamedGoroutine removes a named goroutine from the manager
func (gm *GoroutineManager) RemoveNamedGoroutine(name string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	delete(gm.Goroutines, name)
}

// GetAllNamedGoroutines returns a copy of all named goroutines
func (gm *GoroutineManager) GetAllNamedGoroutines() map[string]*Goroutine {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	result := make(map[string]*Goroutine, len(gm.Goroutines))
	for name, goroutine := range gm.Goroutines {
		result[name] = goroutine
	}
	return result
}

// GetAllAnonymousGoroutines returns a copy of all anonymous goroutines
func (gm *GoroutineManager) GetAllAnonymousGoroutines() []*Goroutine {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	result := make([]*Goroutine, len(gm.Anonymous))
	copy(result, gm.Anonymous)
	return result
}

// ClearAll removes all goroutines from the manager
func (gm *GoroutineManager) ClearAll() {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	
	// Cleanup all named goroutines
	for _, goroutine := range gm.Goroutines {
		goroutine.Cleanup()
	}
	
	// Cleanup all anonymous goroutines
	for _, goroutine := range gm.Anonymous {
		goroutine.Cleanup()
	}
	
	gm.Goroutines = make(map[string]*Goroutine)
	gm.Anonymous = make([]*Goroutine, 0)
}

// CleanupCompletedGoroutines removes and cleans up all completed goroutines
func (gm *GoroutineManager) CleanupCompletedGoroutines() {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.cleanupCompletedLocked()
}

// cleanupCompletedLocked performs cleanup without acquiring the lock (internal use)
func (gm *GoroutineManager) cleanupCompletedLocked() {
	// Clean up completed named goroutines
	for name, goroutine := range gm.Goroutines {
		if goroutine.IsCompleted() {
			goroutine.Cleanup()
			delete(gm.Goroutines, name)
		}
	}
	
	// Clean up completed anonymous goroutines
	newAnonymous := make([]*Goroutine, 0)
	for _, goroutine := range gm.Anonymous {
		if goroutine.IsCompleted() {
			goroutine.Cleanup()
		} else {
			newAnonymous = append(newAnonymous, goroutine)
		}
	}
	gm.Anonymous = newAnonymous
}

// RemoveAndCleanupNamed removes a named goroutine and performs cleanup
func (gm *GoroutineManager) RemoveAndCleanupNamed(name string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	
	if goroutine, exists := gm.Goroutines[name]; exists {
		goroutine.Cleanup()
		delete(gm.Goroutines, name)
	}
}

// GetCompletedCount returns the number of completed goroutines
func (gm *GoroutineManager) GetCompletedCount() (namedCompleted, anonymousCompleted int) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	
	for _, goroutine := range gm.Goroutines {
		if goroutine.IsCompleted() {
			namedCompleted++
		}
	}
	
	for _, goroutine := range gm.Anonymous {
		if goroutine.IsCompleted() {
			anonymousCompleted++
		}
	}
	
	return namedCompleted, anonymousCompleted
}

// SetMaxLimits updates the maximum size limits for both collections
func (gm *GoroutineManager) SetMaxLimits(maxNamed, maxAnonymous int) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.MaxNamedSize = maxNamed
	gm.MaxAnonymousSize = maxAnonymous
}

// SetAutoCleanup enables or disables automatic cleanup of completed goroutines
func (gm *GoroutineManager) SetAutoCleanup(enabled bool) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.AutoCleanup = enabled
}

// GetLimits returns the current size limits and auto-cleanup setting
func (gm *GoroutineManager) GetLimits() (maxNamed, maxAnonymous int, autoCleanup bool) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	return gm.MaxNamedSize, gm.MaxAnonymousSize, gm.AutoCleanup
}

// GetCapacityInfo returns current usage and capacity information
func (gm *GoroutineManager) GetCapacityInfo() (namedCount, namedMax, anonymousCount, anonymousMax int) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	return len(gm.Goroutines), gm.MaxNamedSize, len(gm.Anonymous), gm.MaxAnonymousSize
}

// IsAtCapacity checks if either collection is at its maximum capacity
func (gm *GoroutineManager) IsAtCapacity() (namedAtCapacity, anonymousAtCapacity bool) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	
	namedAtCapacity = gm.MaxNamedSize > 0 && len(gm.Goroutines) >= gm.MaxNamedSize
	anonymousAtCapacity = gm.MaxAnonymousSize > 0 && len(gm.Anonymous) >= gm.MaxAnonymousSize
	
	return namedAtCapacity, anonymousAtCapacity
}

// Reset completely resets the manager to a fresh state
func (gm *GoroutineManager) Reset() {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.Goroutines = make(map[string]*Goroutine)
	gm.Anonymous = make([]*Goroutine, 0)
}
