package goq

import "go/types"

func toType(v interface{}) types.Type {
	switch v := v.(type) {
	case types.Type:
		return v
	case types.Object:
		return v.Type()
	}
	return nil
}

type Type struct {
	Identical  types.Type
	Implements types.Type
	Default    Query
	Underlying Query
}

var (
	_ Query = (*Type)(nil)
)

// Match implemets Query.Match.
func (q *Type) Match(v interface{}) bool {
	t := toType(v)
	if t == nil {
		return false
	}

	if q.Identical != nil && !types.Identical(t, q.Identical) {
		return false
	}

	if q.Implements != nil {
		if i, ok := q.Implements.(*types.Interface); ok && !types.Implements(t, i) {
			return false
		}

		if i, ok := q.Implements.Underlying().(*types.Interface); ok && !types.Implements(t, i) {
			return false
		}
	}

	if q.Default != nil && !q.Default.Match(t) {
		return false
	}

	if q.Underlying != nil && !q.Underlying.Match(t) {
		return false
	}

	return true
}
