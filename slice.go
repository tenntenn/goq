package goq

import "go/types"

var (
	_ TypeMatcher = (*SliceQuery)(nil)
	_ Query       = (*SliceQuery)(nil)
)

type SliceQuery struct {
	Elem TypeMatcher
}

func (q *SliceQuery) Match(typ types.Type) bool {

	t, ok := typ.(*types.Slice)
	if !ok {
		return false
	}

	if q.Elem != nil && !q.Elem.Match(t.Elem()) {
		return false
	}

	return true
}

func (q *SliceQuery) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
