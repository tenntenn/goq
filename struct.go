package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/types"
)

type StructQuery struct {
	*types.Struct
}

var (
	_ types.Type = (*StructQuery)(nil)
	_ Query      = (*StructQuery)(nil)
)

func (q *StructQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Check(o.Type())
}
