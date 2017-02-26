package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/types"
)

func Array(elem types.Type) *ArrayQuery {
	return &ArrayQuery{
		&types.Array{
			Elem: elem,
		},
	}
}

type ArrayQuery struct {
	*types.Array
}

func (q *ArrayQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Array.Check(o.Type())
}
