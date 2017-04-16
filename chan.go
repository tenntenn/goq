package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Chan)(nil)
)

// Chan is an query for chan objects.
type Chan struct {
	// Elem is type of send or receive values.
	Elem Query
	// Dir is direction of the chan.
	Dir *optional.Int
}

// Match implements Query.Match.
func (q *Chan) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Chan)
	if !ok {
		return false
	}

	if !q.Dir.Match(int(t.Dir())) {
		return false
	}

	if q.Elem != nil && !q.Match(t.Elem()) {
		return false
	}

	return true
}
