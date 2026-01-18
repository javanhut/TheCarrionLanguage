// src/object/suggestion_context.go
package object

import (
	"sort"
	"strings"
)

// SuggestionContext captures runtime state at error time to enable intelligent suggestions.
// Used by the error system to suggest "did you mean X?" when unknown identifiers are encountered.
type SuggestionContext struct {
	AvailableMethods   []string // Methods available on the target object
	AvailableVariables []string // Variables in current scope
	AvailableBuiltins  []string // Builtin function names
	TargetType         string   // The type of object being operated on
	AttemptedName      string   // What the user tried to call/access
	SimilarNames       []string // Pre-computed similar names (closest matches)
}

// NewSuggestionContext creates a new empty SuggestionContext with initialized slices.
func NewSuggestionContext() *SuggestionContext {
	return &SuggestionContext{
		AvailableMethods:   []string{},
		AvailableVariables: []string{},
		AvailableBuiltins:  []string{},
		SimilarNames:       []string{},
	}
}

// LevenshteinDistance computes the edit distance between two strings.
// Used to find similar names for "did you mean?" suggestions.
func LevenshteinDistance(a, b string) int {
	// Convert to lowercase for case-insensitive comparison
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	// Create a 2D slice for dynamic programming
	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
	}

	// Initialize first row and column
	for i := 0; i <= len(a); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		matrix[0][j] = j
	}

	// Fill in the rest of the matrix
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[len(a)][len(b)]
}

// min returns the minimum of three integers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// candidateWithDistance pairs a candidate name with its edit distance
type candidateWithDistance struct {
	name     string
	distance int
}

// FindSimilarNames returns candidates within maxDistance of target, sorted by distance.
// Returns at most 3 suggestions, closest first.
func FindSimilarNames(target string, candidates []string, maxDistance int) []string {
	if len(target) == 0 || len(candidates) == 0 {
		return []string{}
	}

	var matches []candidateWithDistance

	for _, candidate := range candidates {
		dist := LevenshteinDistance(target, candidate)
		if dist <= maxDistance && dist > 0 { // dist > 0 excludes exact matches
			matches = append(matches, candidateWithDistance{
				name:     candidate,
				distance: dist,
			})
		}
	}

	// Sort by distance (closest first)
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].distance < matches[j].distance
	})

	// Limit to top 3
	limit := 3
	if len(matches) < limit {
		limit = len(matches)
	}

	result := make([]string, limit)
	for i := 0; i < limit; i++ {
		result[i] = matches[i].name
	}

	return result
}

// GetObjectMethods returns the method names available on an object.
// For Grimoire and Instance types, returns method names from the Methods map.
// For other types, returns an empty slice (they use builtins instead).
func GetObjectMethods(obj Object) []string {
	if obj == nil {
		return []string{}
	}

	switch o := obj.(type) {
	case *Instance:
		if o.Grimoire == nil {
			return []string{}
		}
		return getGrimoireMethods(o.Grimoire)
	case *Grimoire:
		return getGrimoireMethods(o)
	default:
		return []string{}
	}
}

// getGrimoireMethods extracts method names from a Grimoire, including inherited methods.
func getGrimoireMethods(grim *Grimoire) []string {
	if grim == nil {
		return []string{}
	}

	methods := make(map[string]bool)

	// Add methods from this grimoire
	for name := range grim.Methods {
		methods[name] = true
	}

	// Add inherited methods (parent methods that aren't overridden)
	if grim.Inherits != nil {
		for _, name := range getGrimoireMethods(grim.Inherits) {
			if !methods[name] {
				methods[name] = true
			}
		}
	}

	// Convert to sorted slice
	result := make([]string, 0, len(methods))
	for name := range methods {
		result = append(result, name)
	}
	sort.Strings(result)

	return result
}

// GetBuiltinNames returns a list of all builtin function names in Carrion.
// Used to suggest builtins when a user calls an unknown function.
func GetBuiltinNames() []string {
	return []string{
		"len", "print", "println", "type", "input", "range",
		"str", "int", "float",
		"append", "keys", "values", "has_key", "delete",
		"slice", "join", "split", "replace",
		"upper", "lower", "strip", "trim",
		"starts_with", "ends_with", "contains",
		"index_of", "last_index_of",
		"sort", "reverse",
		"sum", "min", "max", "abs", "round", "floor", "ceil",
		"open", "read", "write", "close",
		"sleep", "time", "now",
	}
}

// PopulateSuggestionContext fills in a SuggestionContext with available methods and builtins.
// Call this when an error occurs to capture what was available at the time.
func PopulateSuggestionContext(ctx *SuggestionContext, obj Object, attemptedName string) {
	if ctx == nil {
		return
	}

	ctx.AttemptedName = attemptedName
	ctx.AvailableMethods = GetObjectMethods(obj)
	ctx.AvailableBuiltins = GetBuiltinNames()

	if obj != nil {
		ctx.TargetType = string(obj.Type())
	}

	// Compute similar names from methods and builtins combined
	allCandidates := append(ctx.AvailableMethods, ctx.AvailableBuiltins...)
	ctx.SimilarNames = FindSimilarNames(attemptedName, allCandidates, 3)
}

// GetEnvironmentNames returns all variable names visible from the given environment.
// Walks up the outer chain to include enclosing scopes.
// Inner scope names shadow outer scope names (no duplicates).
func GetEnvironmentNames(env *Environment) []string {
	if env == nil {
		return []string{}
	}

	names := make(map[string]bool)

	// Walk the scope chain, collecting names (inner shadows outer)
	for e := env; e != nil; e = e.GetOuter() {
		for _, name := range e.GetNames() {
			names[name] = true
		}
	}

	// Convert to sorted slice for consistent ordering
	result := make([]string, 0, len(names))
	for name := range names {
		result = append(result, name)
	}
	sort.Strings(result)

	return result
}

// BuildSuggestionContext is the main entry point for the evaluator.
// Creates a fully populated SuggestionContext with methods, builtins, and variables.
func BuildSuggestionContext(obj Object, attemptedName string, env *Environment) *SuggestionContext {
	ctx := NewSuggestionContext()

	// Populate with object methods and builtins
	PopulateSuggestionContext(ctx, obj, attemptedName)

	// Add environment variables if available
	if env != nil {
		ctx.AvailableVariables = GetEnvironmentNames(env)

		// Re-compute similar names to include variables
		allCandidates := make([]string, 0, len(ctx.AvailableMethods)+len(ctx.AvailableBuiltins)+len(ctx.AvailableVariables))
		allCandidates = append(allCandidates, ctx.AvailableMethods...)
		allCandidates = append(allCandidates, ctx.AvailableBuiltins...)
		allCandidates = append(allCandidates, ctx.AvailableVariables...)
		ctx.SimilarNames = FindSimilarNames(attemptedName, allCandidates, 3)
	}

	return ctx
}

// FormatSuggestion returns a human-readable suggestion string.
// Returns empty string if no suggestions available.
func FormatSuggestion(ctx *SuggestionContext) string {
	if ctx == nil || len(ctx.SimilarNames) == 0 {
		return ""
	}

	if len(ctx.SimilarNames) == 1 {
		return "Did you mean '" + ctx.SimilarNames[0] + "'?"
	}

	// Format multiple suggestions
	quoted := make([]string, len(ctx.SimilarNames))
	for i, name := range ctx.SimilarNames {
		quoted[i] = "'" + name + "'"
	}
	return "Did you mean one of: " + strings.Join(quoted, ", ") + "?"
}
