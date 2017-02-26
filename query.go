package typequery

import "go/types"

var (
	_ Query = andQuery(nil)
	_ Query = orQuery(nil)
	_ Query = (*notQuery)(nil)
)

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

func Not(q Query) Query {
	return &notQuery{
		Query: q,
	}
}

type notQuery struct {
	Query
}

func (q *notQuery) Exec(o types.Object) bool {
	return !q.Query.Exec(o)
}
