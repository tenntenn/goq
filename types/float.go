package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type Float struct {
	Size *int
}

var _ Type = (*Float)(nil)

func (f *Float) Check(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info() != types.IsFloat {
		return false
	}

	if !option.Has(f.Size, size(t)) {
		return false
	}

	return true
}
