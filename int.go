package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Int)(nil)
)

// Int is a query for integer objects.
type Int struct {
	// Size is size of the integer value.
	Size *optional.Int
	// Unsigned means whether the value is unsigned or not.
	Unsigned *optional.Bool
}

// Match implements Query.Match.
func (q *Int) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Basic)
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
