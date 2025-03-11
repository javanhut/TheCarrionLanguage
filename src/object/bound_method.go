package object

import "fmt"

type BoundMethod struct {
	Instance *Instance
	Method   *Function
	Name     string
}

func (bm *BoundMethod) Type() ObjectType {
	return "BOUND_METHOD"
}

func (bm *BoundMethod) Inspect() string {
	return fmt.Sprintf("<bound method %s>", bm.Name)
}
