package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/types"
)

func Slice(elem types.Type) *SliceQuery {
	return &SliceQuery{
		&types.Slice{
			Elem: elem,
		},
	}
}

type SliceQuery struct {
	*types.Slice
}

func (q *SliceQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Slice.Check(o.Type())
}
