package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type StructField struct {
	Name      *option.Parttern
	Exported  *bool
	Anonymous *bool
	Type      Type
}

func (f *StructField) check(v *types.Var) bool {
	if option.Has(f.Name, v.Name()) {
		return false
	}

	if option.Has(f.Exported, v.Exported()) {
		return false
	}

	if option.Has(f.Anonymous, v.Anonymous()) {
		return false
	}

	if f.Type != nil && !f.Type.Check(v.Type()) {
		return false
	}

	return true
}
