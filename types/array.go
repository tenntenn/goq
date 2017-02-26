package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type Array struct {
	Elem Type
	Len  *int64
}

var _ Type = (*Array)(nil)

func (a *Array) Check(typ types.Type) bool {

	t, ok := typ.(*types.Array)
	if !ok {
		return false
	}

	if !option.Has(a.Len, t.Len()) {
		return false
	}

	if a.Elem != nil && !a.Elem.Check(t.Elem()) {
		return false
	}

	return true
}
