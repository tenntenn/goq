package typequery

import (
	gtypes "go/types"

	"github.com/tenntenn/typequery/option"
	"github.com/tenntenn/typequery/types"
)

func NamedType(underlying types.Type) *NamedTypeQuery {
	return &NamedTypeQuery{
		Underlying: underlying,
	}
}

type NamedTypeQuery struct {
	Package    *PackageQuery
	Name       *option.Parttern
	Underlying types.Type
}

var (
	_ types.Type = (*NamedTypeQuery)(nil)
	_ Query      = (*NamedTypeQuery)(nil)
)

func (q *NamedTypeQuery) Check(typ gtypes.Type) bool {
	if q.Underlying != nil && !q.Underlying.Check(typ.Underlying()) {
		return false
	}
	return true
}

func (q *NamedTypeQuery) Exec(o gtypes.Object) bool {
	if o == nil {
		return false
	}

	if !option.Has(q.Name, o.Name()) {
		return false
	}

	if q.Package != nil && !q.Exec(o) {
		return false
	}

	if o.Type() == nil {
		return false
	}

	if q.Underlying != nil && !q.Underlying.Check(o.Type().Underlying()) {
		return false
	}

	return true
}
