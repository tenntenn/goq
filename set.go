package goq

import "github.com/tenntenn/optional"

func setEq(v1, v2 interface{}) bool {
	q, ok := v1.(Query)
	if !ok {
		return false
	}

	return q.Match(v2)
}

func NewSet(len *optional.Int) *optional.Set {
	return optional.NewSet(len, setEq)
}
