package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/option"
)

func Package(path string) *PackageQuery {
	return &PackageQuery{
		Path: &option.Parttern{Value: &path},
	}
}

type PackageQuery struct {
	Path    *option.Parttern
	Name    *option.Parttern
	Imports []*PackageQuery
	// TODO(tenntenn): Scope
}

func (q *PackageQuery) exec(pkg *gtypes.Package) bool {
	if option.Has(q.Name, pkg.Name()) {
		return false
	}

	if option.Has(q.Path, pkg.Path()) {
		return false
	}

	return true
}

func (q *PackageQuery) imported(o gtypes.Object, pkg *PackageQuery) bool {
	for _, p := range o.Pkg().Imports() {
		if pkg.exec(p) {
			return true
		}
	}
	return false
}

func (q *PackageQuery) Exec(o gtypes.Object) bool {
	if o == nil || o.Pkg == nil {
		return false
	}

	if !q.exec(o.Pkg()) {
		return false
	}

	for _, pq := range q.Imports {
		if !q.imported(o, pq) {
			return false
		}
	}

	return true
}
