package object

import "fmt"

type StaticMethod struct {
	Grimoire *Grimoire
	Method   *Function
	Name     string
}

func (sm *StaticMethod) Type() ObjectType {
	return "STATIC_METHOD"
}

func (sm *StaticMethod) Inspect() string {
	return fmt.Sprintf("<static method %s>", sm.Name)
}