package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
	"github.com/tenntenn/optional/pattern"
)

var (
	_ Query = (*Func)(nil)
)

// Func is a query for function objects.
type Func struct {
	// Name is name of the function.
	Name *pattern.Pattern
	// FullName is full name of the function.
	FullName *pattern.Pattern
	// Exported is whether the function is exported or not.
	Exported  *optional.Bool
	Signature *Signature
}

// Exec implements Query.Exec.
func (q *Func) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}

	f, ok := o.(*types.Func)
	if !ok {
		return false
	}

	if !q.Exported.Match(f.Exported()) {
		return false
	}

	if !q.Name.Match(f.Name()) {
		return false
	}

	if !q.FullName.Match(f.FullName()) {
		return false
	}

	return q.Signature.Match(f.Type())
}
