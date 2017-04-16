package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ TypeMatcher = (*Struct)(nil)
	_ Query       = (*Struct)(nil)
)

type Struct struct {
	NumFields *optional.Int
	Fields    []*VarQuery
}

func (*Struct) hasField(typ *types.Struct, f *VarQuery) bool {
	for i := 0; i < typ.NumFields(); i++ {
		if t := f.Type; t != nil &&
			t.Match(typ.Field(i).Type()) {
			return true
		}
	}
	return false
}

func (q *Struct) Match(typ types.Type) bool {
	t, ok := typ.(*types.Struct)
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

func (q *Struct) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
