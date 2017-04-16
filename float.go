package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Float)(nil)
)

// Float is a query for float objects.
type Float struct {
	// Size is size of the float value.
	Size *optional.Int
}

// Match implements Query.Match.
func (q *Float) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsFloat == 0 {
		return false
	}

	if !q.Size.Match(size(t)) {
		return false
	}

	return true
}

// Exec implements Query.Exec.
func (q *Float) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
