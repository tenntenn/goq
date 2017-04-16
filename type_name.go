package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
	"github.com/tenntenn/optional/pattern"
)

var (
	_ Query = (*TypeName)(nil)
)

type TypeName struct {
	Exported *optional.Bool
	Name     *pattern.Pattern
	Pkg      *Package
}

// Match implements Query.Match
func (q *TypeName) Match(v interface{}) bool {
	tn, ok := v.(*types.TypeName)
	if !ok {
		return false
	}

	if !q.Exported.Match(tn.Exported()) {
		return false
	}

	if !q.Name.Match(tn.Name()) {
		return false
	}

	if !q.Pkg.Match(tn.Pkg()) {
		return false
	}

	return true
}
