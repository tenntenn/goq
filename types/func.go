package types

import (
	"go/types"

	"github.com/tenntenn/typequery/option"
)

type Func struct {
	Recv     Type
	Params   *option.Slice
	Results  *option.Slice
	Variadic *bool
}

var _ Type = (*Func)(nil)

func checkType(v1, v2 interface{}) bool {
	t1, ok := v1.(Type)
	if !ok {
		return false
	}

	t2, ok := v2.(types.Type)
	if !ok {
		return false
	}

	return t1.Check(t2)
}

func toOptionSlice(ts map[int]Type) *option.Slice {
	elems := make(map[int]interface{}, len(ts))
	for k, v := range ts {
		elems[k] = v
	}
	return option.NewSlice(elems, checkType)
}

func tupleToTypes(tuple *types.Tuple) []types.Type {
	ts := make([]types.Type, tuple.Len())
	for i := 0; i < len(ts); i++ {
		ts[i] = tuple.At(i).Type()
	}
	return ts
}

func NewFunc(recv Type, params, results map[int]Type) *Func {
	return &Func{
		Recv:    recv,
		Params:  toOptionSlice(params),
		Results: toOptionSlice(results),
	}
}

func (f *Func) Check(typ types.Type) bool {

	t, ok := typ.(*types.Signature)
	if !ok {
		return false
	}

	if recv := t.Recv(); f.Recv != nil &&
		recv != nil && f.Recv.Check(recv.Type()) {
		return false
	}

	if !f.Params.Has(tupleToTypes(t.Params())) {
		return false
	}

	if !f.Results.Has(tupleToTypes(t.Results())) {
		return false
	}

	return true
}
