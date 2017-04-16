package goq

import "go/types"

var (
	_ TypeMatcher = (*Bool)(nil)
	_ Query       = (*Bool)(nil)
)

// Bool is a query for integer objects.
type Bool struct{}

// Match implements Type.Match.
func (q *Bool) Match(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsBoolean == 0 {
		return false
	}

	return true
}

// Exec implements Query.Exec.
func (q *Bool) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
