package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type Complex struct {
	Size *int
}

var _ Type = (*Complex)(nil)

func (c *Complex) Check(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info() != types.IsComplex {
		return false
	}

	if !option.Has(c.Size, size(t)) {
		return false
	}

	return true
}
