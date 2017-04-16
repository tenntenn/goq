package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
	"github.com/tenntenn/optional/pattern"
)

/*
var (
	_ Query = (*Package)(nil)
)
*/

type Package struct {
	Path    *pattern.Pattern
	Name    *pattern.Pattern
	Imports *optional.Set
}

func (q *Package) Exec(o types.Object) bool {
	if o == nil || o.Pkg() == nil {
		return false
	}

	p := o.Pkg()

	if !q.Name.Match(p.Name()) {
		return false
	}

	if !q.Path.Match(p.Path()) {
		return false
	}

	if q.Imports != nil && !q.Imports.Match(&pkgImportsSeq{v: p}) {
		return false
	}

	return true
}
