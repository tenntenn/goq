package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type Struct struct {
	NumFields *int
	Fields    []*StructField
}

func (s *Struct) hasField(typ *types.Struct, f *StructField) bool {
	for i := 0; i < typ.NumFields(); i++ {
		if f.check(typ.Field(i)) {
			return true
		}
	}
	return false
}

func (s *Struct) Check(typ types.Type) bool {
	t, ok := typ.(*types.Struct)
	if !ok {
		return false
	}

	if !option.Has(s.NumFields, t.NumFields()) {
		return false
	}

	for _, f := range s.Fields {
		if !s.hasField(t, f) {
			return false
		}
	}

	return true
}
