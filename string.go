package goq

import "go/types"

var (
	_ Query = (*String)(nil)
)

// String is a query for integer objects.
type String struct{}

// Match implements Query.Match.
func (q *String) Match(v interface{}) bool {

	t, ok := toType(v).(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsString == 0 {
		return false
	}

	return true
}
