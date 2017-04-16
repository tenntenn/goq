package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

func eq(v1, v2 interface{}) bool {
	if t1, ok := v1.(TypeMatcher); ok {
		if t2, ok := v2.(types.Type); ok {
			return t1.Match(t2)
		}
	}

	if q, ok := v1.(Query); ok {
		if o, ok := v2.(types.Object); ok {
			return q.Exec(o)
		}
	}

	return false
}

type Set struct {
}

func NewSet() *optional.Set {
	return optional.NewSet(nil, eq)
}
