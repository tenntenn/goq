package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ TypeMatcher = (*Array)(nil)
	_ Query       = (*Array)(nil)
)

// Array is a query for array objects.
type Array struct {
	// Elem is type of the elements.
	Elem TypeMatcher
	// Len is length of the array.
	Len *optional.Int64
}

// Match implements Type.Match.
func (q *Array) Match(typ types.Type) bool {

	t, ok := typ.(*types.Array)
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

// Exec implements Query.Exec.
func (q *Array) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
