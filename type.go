package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/types"
)

func Type(typ types.Type) *TypeQuery {
	return &TypeQuery{
		Type: typ,
	}
}

type TypeQuery struct {
	Type types.Type
}

var (
	_ types.Type = (*TypeQuery)(nil)
	_ Query      = (*TypeQuery)(nil)
)

func (q *TypeQuery) Check(typ gtypes.Type) bool {
	if typ == nil {
		return false
	}
	return q.Type.Check(typ) || q.Type.Check(typ.Underlying())
}

func (q *TypeQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Type.Check(o.Type()) || q.Type.Check(o.Type().Underlying())
}
