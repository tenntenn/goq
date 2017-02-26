package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type Chan struct {
	Elem Type
	Dir  *types.ChanDir
}

var _ Type = (*Chan)(nil)

func (c *Chan) Check(typ types.Type) bool {

	t, ok := typ.(*types.Chan)
	if !ok {
		return false
	}

	if !option.Has(c.Dir, t.Dir()) {
		return false
	}

	if c.Elem != nil && !c.Elem.Check(t.Elem()) {
		return false
	}

	return true
}
