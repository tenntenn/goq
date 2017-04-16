package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ TypeMatcher = (*Int)(nil)
	_ Query       = (*Int)(nil)
)

// Int is a query for integer objects.
type Int struct {
	// Size is size of the integer value.
	Size *optional.Int
	// Unsigned means whether the value is unsigned or not.
	Unsigned *optional.Bool
}

// Match implements Type.Match.
func (q *Int) Match(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsInteger == 0 {
		return false
	}

	if !q.Size.Match(size(t)) {
		return false
	}

	if !q.Unsigned.Match(t.Info()&types.IsUnsigned != 0) {
		return false
	}

	return true
}

// Exec implements Query.Exec.
func (q *Int) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
