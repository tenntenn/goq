package goq

import "go/types"

var (
	_ Query = (*Slice)(nil)
)

type Slice struct {
	Elem Query
}

// Match implements Query.Match.
func (q *Slice) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Slice)
	if !ok {
		return false
	}

	if q.Elem != nil && !q.Elem.Match(t.Elem()) {
		return false
	}

	return true
}
