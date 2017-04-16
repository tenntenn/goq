package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Struct)(nil)
)

type Struct struct {
	NumFields *optional.Int
	Fields    []*Var
}

func (*Struct) hasField(typ *types.Struct, f *Var) bool {
	for i := 0; i < typ.NumFields(); i++ {
		if t := f.Type; t != nil &&
			t.Match(typ.Field(i).Type()) {
			return true
		}
	}
	return false
}

func (q *Struct) Match(v interface{}) bool {
	t, ok := toType(v).(*types.Struct)
	if !ok {
		return false
	}

	if !q.NumFields.Match(t.NumFields()) {
		return false
	}

	for _, f := range q.Fields {
		if !q.hasField(t, f) {
			return false
		}
	}

	return true
}
