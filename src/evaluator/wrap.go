package evaluator

import (
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func wrapBuiltinType(obj object.Object, typeName string, env *object.Environment) object.Object {
	// Look up the spellbook by name (e.g. "Array") in the environment.
	sbObj, ok := env.Get(typeName)
	if !ok {
		return obj
	}
	spellbook, ok := sbObj.(*object.Spellbook)
	if !ok {
		return obj
	}
	// Create a new instance of the spellbook.
	instance := &object.Instance{
		Spellbook: spellbook,
		Env:       object.NewEnclosedEnvironment(spellbook.Env),
	}
	// Store the original object under a known key (e.g. "inner").
	instance.Env.Set("inner", obj)
	return instance
}
