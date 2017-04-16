package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Named)(nil)
)

type Named struct {
	Underlying Query
	Methods    *optional.Tuple
	TypeName   *TypeName
}

// Match implemets Query.Match.
func (q *Named) Match(v interface{}) bool {

	n, ok := toType(v).(*types.Named)
	if !ok {
		return false
	}

	if q.Underlying != nil && !q.Underlying.Match(n.Underlying()) {
		return false
	}

	if !q.Methods.Match(&namedMethodsSeq{v: n}) {
		return false
	}

	if q.TypeName != nil && !q.TypeName.Match(n.Obj()) {
		return false
	}

	return true
}
