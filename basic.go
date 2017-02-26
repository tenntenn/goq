package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/types"
)

func Int(size *int) *IntQuery {
	return &IntQuery{
		Int: &types.Int{
			Size: size,
		},
	}
}

type IntQuery struct {
	*types.Int
}

var (
	_ types.Type = (*IntQuery)(nil)
	_ Query      = (*IntQuery)(nil)
)

func (q *IntQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Int.Check(o.Type())
}

func Float(size *int) *FloatQuery {
	return &FloatQuery{
		Float: &types.Float{
			Size: size,
		},
	}
}

type FloatQuery struct {
	*types.Float
}

var (
	_ types.Type = (*FloatQuery)(nil)
	_ Query      = (*FloatQuery)(nil)
)

func (q *FloatQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Float.Check(o.Type())
}

func Complex(size *int) *ComplexQuery {
	return &ComplexQuery{
		Complex: &types.Complex{
			Size: size,
		},
	}
}

type ComplexQuery struct {
	*types.Complex
}

var (
	_ types.Type = (*ComplexQuery)(nil)
	_ Query      = (*ComplexQuery)(nil)
)

func (q *ComplexQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Complex.Check(o.Type())
}

func Bool() *BoolQuery {
	return &BoolQuery{
		Bool: &types.Bool{},
	}
}

type BoolQuery struct {
	*types.Bool
}

var (
	_ types.Type = (*BoolQuery)(nil)
	_ Query      = (*BoolQuery)(nil)
)

func (q *BoolQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Bool.Check(o.Type())
}

func String() *StringQuery {
	return &StringQuery{
		String: &types.String{},
	}
}

type StringQuery struct {
	*types.String
}

var (
	_ types.Type = (*StringQuery)(nil)
	_ Query      = (*StringQuery)(nil)
)

func (q *StringQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.String.Check(o.Type())
}
