package goq

import "go/types"

var (
	_ Query = (*Bool)(nil)
)

// Bool is a query for integer objects.
type Bool struct{}

// Match implements Query.Match.
func (q *Bool) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsBoolean == 0 {
		return false
	}

	return true
}
