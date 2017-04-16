package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ TypeMatcher = (*Complex)(nil)
	_ Query       = (*Complex)(nil)
)

// Complex is a query for float objects.
type Complex struct {
	// Size is size of the float value.
	Size *optional.Int
}

// Match implements Type.Match.
func (q *Complex) Match(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsComplex == 0 {
		return false
	}

	if !q.Size.Match(size(t)) {
		return false
	}

	return true
}

// Exec implements Query.Exec.
func (q *Complex) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
