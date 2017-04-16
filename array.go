package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Array)(nil)
)

// Array is a query for array objects.
type Array struct {
	// Elem is type of the elements.
	Elem Query
	// Len is length of the array.
	Len *optional.Int64
}

// Match implements Query.Match.
func (q *Array) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Array)
	if !ok {
		return false
	}

	if !q.Len.Match(t.Len()) {
		return false
	}

	if q.Elem != nil && !q.Elem.Match(t.Elem()) {
		return false
	}

	return true
}
