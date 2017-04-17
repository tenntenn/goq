package goq

import "github.com/tenntenn/optional"

func tuppleEq(v1, v2 interface{}) bool {
	q, ok := v1.(Query)
	if !ok {
		return false
	}

	return q.Match(v2)
}

func NewTupple(len *optional.Int) *optional.Tuple {
	return optional.NewTupple(len, tuppleEq)
}
