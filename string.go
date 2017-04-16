package goq

import "go/types"

var (
	_ TypeMatcher = (*String)(nil)
	_ Query       = (*String)(nil)
)

// String is a query for integer objects.
type String struct{}

// Match implements Type.Match.
func (q *String) Match(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsString == 0 {
		return false
	}

	return true
}

// Exec implements Query.Exec.
func (q *String) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
