package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/option"
	"github.com/tenntenn/typequery/types"
)

func Func(name string) *FuncQuery {
	return &FuncQuery{
		Name: &option.Parttern{Value: &name},
	}
}

type FuncQuery struct {
	*types.Func
	Name     *option.Parttern
	FullName *option.Parttern
	Exported *bool
}

var (
	_ types.Type = (*FuncQuery)(nil)
	_ Query      = (*FuncQuery)(nil)
)

func (q *FuncQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}

	f, ok := o.(*gtypes.Func)
	if !ok {
		return false
	}

	if !option.Has(q.Exported, f.Exported()) {
		return false
	}

	if !option.Has(q.Name, f.Name()) {
		return false
	}

	if !option.Has(q.FullName, f.FullName()) {
		return false
	}

	return q.Func.Check(f.Type())
}
