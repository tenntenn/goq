package types

import "go/types"

type Bool struct{}

var _ Type = (*Bool)(nil)

func (*Bool) Check(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info() != types.IsBoolean {
		return false
	}

	return true
}
