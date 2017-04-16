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

// FIXME: remove it
type pkg struct {
	types.Object
	v *types.Package
}

func (p *pkg) Pkg() *types.Package {
	return p.v
}

func (q *TypeName) Exec(o types.Object) bool {
	tn, ok := o.(*types.TypeName)
	if !ok {
		return false
	}

	if !q.Exported.Match(tn.Exported()) {
		return false
	}

	if !q.Name.Match(tn.Name()) {
		return false
	}

	if !q.Pkg.Exec(&pkg{v: tn.Pkg()}) {
		return false
	}

	return true
}
