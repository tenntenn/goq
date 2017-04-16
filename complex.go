package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Complex)(nil)
)

// Complex is a query for float objects.
type Complex struct {
	// Size is size of the float value.
	Size *optional.Int
}

// Match implements Query.Match.
func (q *Complex) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Basic)
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
