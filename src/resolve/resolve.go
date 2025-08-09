package resolve

import "github.com/javanhut/TheCarrionLanguage/src/ast"

// Binding describes a fast access path: walk Depth frames, then use Slot.
type Binding struct {
	Depth int
	Slot  int
}

type scope struct {
	parent *scope
	slots  map[uint32]int // sym -> slot
	next   int
}

func newScope(parent *scope) *scope {
	return &scope{parent: parent, slots: make(map[uint32]int)}
}

func (s *scope) declare(sym uint32) int {
	if slot, ok := s.slots[sym]; ok {
		return slot
	}
	slot := s.next
	s.slots[sym] = slot
	s.next++
	return slot
}

func (s *scope) resolve(sym uint32, depth int) (Binding, bool) {
	if slot, ok := s.slots[sym]; ok {
		return Binding{Depth: depth, Slot: slot}, true
	}
	if s.parent == nil {
		return Binding{}, false
	}
	return s.parent.resolve(sym, depth+1)
}

type Results struct {
	Idents map[*ast.Identifier]Binding
}

func Resolve(node ast.Node) (*Results, error) {
	a := &annotator{
		cur: newScope(nil),
		out: &Results{Idents: make(map[*ast.Identifier]Binding)},
	}
	a.walk(node)
	return a.out, nil
}

type annotator struct {
	cur *scope
	out *Results
}

func (a *annotator) withScope(fn func()) {
	a.cur = newScope(a.cur)
	fn()
	a.cur = a.cur.parent
}

func (a *annotator) walk(n ast.Node) {
	switch x := n.(type) {
	case *ast.Program:
		for _, s := range x.Statements {
			a.walk(s)
		}
	case *ast.BlockStatement:
		a.withScope(func() {
			for _, s := range x.Statements {
				a.walk(s)
			}
		})
	case *ast.AssignStatement:
		a.bindLHS(x.Name)
		if x.TypeHint != nil {
			a.walk(x.TypeHint)
		}
		if x.Value != nil {
			a.walk(x.Value)
		}
	case *ast.FunctionLiteral:
		a.withScope(func() {
			for _, p := range x.Parameters {
				a.cur.declare(p.Sym)
			}
			if x.Body != nil {
				a.walk(x.Body)
			}
		})
	case *ast.Identifier:
		// r-value read
		if b, ok := a.cur.resolve(x.Sym, 0); ok {
			a.out.Idents[x] = b
		}
	default:
		for _, ch := range childrenOf(n) {
			if ch != nil {
				a.walk(ch)
			}
		}
	}
}

func (a *annotator) bindLHS(e ast.Expression) {
	switch v := e.(type) {
	case *ast.Identifier:
		a.cur.declare(v.Sym)
	case *ast.TupleLiteral:
		for _, el := range v.Elements {
			a.bindLHS(el)
		}
	case *ast.IndexExpression:
		a.walk(v.Left)
		a.walk(v.Index)
	case *ast.DotExpression:
		a.walk(v.Left)
	default:
	}
}

func childrenOf(n ast.Node) []ast.Node {
	switch x := n.(type) {
	case *ast.ExpressionStatement:
		return []ast.Node{x.Expression}
	case *ast.InfixExpression:
		return []ast.Node{x.Left, x.Right}
	case *ast.PrefixExpression:
		return []ast.Node{x.Right}
	case *ast.CallExpression:
		out := []ast.Node{x.Function}
		for _, a := range x.Arguments {
			out = append(out, a)
		}
		return out
	case *ast.IfStatement:
		out := []ast.Node{x.Condition, x.Consequence}
		for _, b := range x.OtherwiseBranches {
			out = append(out, b.Condition, b.Consequence)
		}
		if x.Alternative != nil {
			out = append(out, x.Alternative)
		}
		return out
	case *ast.WhileStatement:
		return []ast.Node{x.Condition, x.Body}
	case *ast.ForStatement:
		out := []ast.Node{x.Variable, x.Iterable, x.Body}
		if x.Alternative != nil {
			out = append(out, x.Alternative)
		}
		return out
	case *ast.ReturnStatement:
		return []ast.Node{x.ReturnValue}
	case *ast.TupleLiteral:
		out := make([]ast.Node, 0, len(x.Elements))
		for _, el := range x.Elements {
			out = append(out, el)
		}
		return out
	}
	return nil
}
