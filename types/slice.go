package types

import "go/types"

type Slice struct {
	Elem Type
}

var _ Type = (*Slice)(nil)

func (s *Slice) Check(typ types.Type) bool {

	t, ok := typ.(*types.Slice)
	if !ok {
		return false
	}

	if s.Elem != nil && !s.Elem.Check(t.Elem()) {
		return false
	}

	return true
}
