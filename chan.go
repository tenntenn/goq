package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ TypeMatcher = (*Chan)(nil)
	_ Query       = (*Chan)(nil)
)

// Chan is an query for chan objects.
type Chan struct {
	// Elem is type of send or receive values.
	Elem TypeMatcher
	// Dir is direction of the chan.
	Dir *optional.Int
}

// Match implements Type.Match.
func (q *Chan) Match(typ types.Type) bool {

	t, ok := typ.(*types.Chan)
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

// Exec implements Query.Exec.
func (q *Chan) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
