package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/option"
	"github.com/tenntenn/typequery/types"
)

func Var(name string) *VarQuery {
	return &VarQuery{
		Name: &option.Parttern{Value: &name},
	}
}

type VarQuery struct {
	types.Type
	Name     *option.Parttern
	Package  *PackageQuery
	Exported *bool
}

var (
	_ types.Type = (*VarQuery)(nil)
	_ Query      = (*VarQuery)(nil)
)

func (q *VarQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}

	v, ok := o.(*gtypes.Var)
	if !ok {
		return false
	}

	if v.IsField() {
		return false
	}

	if !option.Has(q.Exported, v.Exported()) {
		return false
	}

	if !option.Has(q.Name, v.Name()) {
		return false
	}

	return q.Type.Check(v.Type())
}
