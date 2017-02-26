package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/types"
)

func Chan(elem types.Type) *ChanQuery {
	return &ChanQuery{
		&types.Chan{
			Elem: elem,
		},
	}
}

type ChanQuery struct {
	*types.Chan
}

func (q *ChanQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Chan.Check(o.Type())
}
