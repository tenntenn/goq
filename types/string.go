package types

import "go/types"

type String struct{}

var _ Type = (*String)(nil)

func (*String) Check(typ types.Type) bool {

	t, ok := typ.(*types.Basic)
	if !ok {
		return false
	}

	if t.Info() != types.IsString {
		return false
	}

	return true
}
