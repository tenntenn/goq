package option

import "reflect"

const (
	Tail = -1
)

type Slice struct {
	elems map[int]interface{}
	check func(v1, v2 interface{}) bool
	Len   *int
}

func NewSlice(elems map[int]interface{}, check func(v1, v2 interface{}) bool) *Slice {
	return &Slice{
		elems: elems,
		check: check,
	}
}

func (s *Slice) Has(slice interface{}) bool {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return true
	}

	n := v.Len()
	if !Has(s.Len, n) {
		return false
	}

	for i, e := range s.elems {
		if i > n {
			return false
		}
		var index int
		switch {
		case index < 0:
			index = n - i
		default:
			index = i
		}

		if index < 0 {
			return false
		}

		if !s.check(e, v.Index(index).Interface()) {
			return false
		}
	}

	return true
}
