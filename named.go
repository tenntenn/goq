package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Named)(nil)
)

type Named struct {
	Underlying TypeMatcher
	Methods    *optional.Tuple
	TypeName   *TypeName
}

func (q *Named) Exec(o types.Object) bool {
	if o == nil {
		return false
	}

	t := o.Type()
	if t == nil {
		return false
	}

	n, ok := t.(*types.Named)
	if !ok {
		return false
	}

	if q.Underlying != nil && !q.Underlying.Match(n.Underlying()) {
		return false
	}

	if !q.Methods.Match(&namedMethodsSeq{v: n}) {
		return false
	}

	if q.TypeName != nil && !q.TypeName.Exec(n.Obj()) {
		return false
	}

	return true
}
