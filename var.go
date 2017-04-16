package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
	"github.com/tenntenn/optional/pattern"
)

var (
	_ Query = (*Var)(nil)
)

type Var struct {
	IsField   *optional.Bool
	Name      *pattern.Pattern
	Exported  *optional.Bool
	Anonymous *optional.Bool
	Type      Query
}

// Match implements Query.Match.
func (q *Var) Match(v interface{}) bool {

	o, ok := v.(*types.Var)
	if !ok {
		return false
	}

	if !q.IsField.Match(o.IsField()) {
		return false
	}

	if !q.Exported.Match(o.Exported()) {
		return false
	}

	if !q.Anonymous.Match(o.Anonymous()) {
		return false
	}

	if !q.Name.Match(o.Name()) {
		return false
	}

	if q.Type != nil && !q.Type.Match(o.Type()) {
		return false
	}

	return true
}
