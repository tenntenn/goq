package typequery

import "go/types"

type Query interface {
	Exec(o types.Object) bool
}

func And(qs ...Query) Query {
	return andQuery(qs)
}

type andQuery []Query

func (qs andQuery) Exec(o types.Object) bool {
	for i := range qs {
		if !qs[i].Exec(o) {
			return false
		}
	}
	return true
}

func Or(qs ...Query) Query {
	return orQuery(qs)
}

type orQuery []Query

func (qs orQuery) Exec(o types.Object) bool {
	for i := range qs {
		if qs[i].Exec(o) {
			return true
		}
	}
	return false
}
