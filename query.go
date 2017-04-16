package goq

var (
	_ Query = and(nil)
	_ Query = or(nil)
	_ Query = (*not)(nil)
)

// Query is an query to search objects.
type Query interface {
	Match(v interface{}) bool
}

// And concats queries into a query
// which Match returns true when all queries Match return true.
func And(qs ...Query) Query {
	return and(qs)
}

type and []Query

func (qs and) Match(v interface{}) bool {
	for i := range qs {
		if !qs[i].Match(v) {
			return false
		}
	}
	return true
}

// Or concats queries into a query
// which Match returns false when all queries Match return false.
func Or(qs ...Query) Query {
	return or(qs)
}

type or []Query

func (qs or) Match(v interface{}) bool {
	for i := range qs {
		if qs[i].Match(v) {
			return true
		}
	}
	return false
}

// Not return a query which Match returns true
// if given query's Match return false.
func Not(q Query) Query {
	return &not{
		Query: q,
	}
}

type not struct {
	Query
}

func (q *not) Match(v interface{}) bool {
	return !q.Query.Match(v)
}
