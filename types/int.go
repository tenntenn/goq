package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type Int struct {
	Size     *int
	Unsigned *bool
}

var _ Type = (*Int)(nil)

func (n *Int) Check(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info()&types.IsInteger == 0 {
		return false
	}

	if !option.Has(n.Size, size(t)) {
		return false
	}

	if !option.Has(n.Unsigned, t.Info()&types.IsUnsigned != 0) {
		return false
	}

	return true
}
