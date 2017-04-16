package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
	"github.com/tenntenn/optional/pattern"
)

var (
	_ Query = (*Package)(nil)
)

func toPackage(v interface{}) *types.Package {
	switch v := v.(type) {
	case *types.Package:
		return v
	case types.Object:
		return v.Pkg()
	}
	return nil
}

type Package struct {
	Path    *pattern.Pattern
	Name    *pattern.Pattern
	Imports *optional.Set
}

// Match implements Query.Match.
func (q *Package) Match(v interface{}) bool {

	p := toPackage(v)
	if p == nil {
		return false
	}

	if !q.Name.Match(p.Name()) {
		return false
	}

	if !q.Path.Match(p.Path()) {
		return false
	}

	if q.Imports != nil && !q.Imports.Match(p) {
		return false
	}

	return true
}
