package object

type BoundMethod struct {
	Instance *Instance
	Method   *Function
}

func (bm *BoundMethod) Type() ObjectType {
	return "BOUND_METHOD"
}

func (bm *BoundMethod) Inspect() string {
	return "<bound method>"
}
