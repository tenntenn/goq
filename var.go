package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
	"github.com/tenntenn/optional/pattern"
)

var (
	_ Query = (*VarQuery)(nil)
)

type VarQuery struct {
	IsField   *optional.Bool
	Name      *pattern.Pattern
	Exported  *optional.Bool
	Anonymous *optional.Bool
	Type      TypeMatcher
}

func (q *VarQuery) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}

	v, ok := o.(*types.Var)
	if !ok {
		return false
	}

	if !q.IsField.Match(v.IsField()) {
		return false
	}

	if !q.Exported.Match(v.Exported()) {
		return false
	}

	if !q.Anonymous.Match(v.Anonymous()) {
		return false
	}

	if !q.Name.Match(v.Name()) {
		return false
	}

	if q.Type != nil && !q.Type.Match(v.Type()) {
		return false
	}

	return true
}
