package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ TypeMatcher = (*Interface)(nil)
	_ Query       = (*Interface)(nil)
)

type embededsSeq struct {
	v *types.Interface
}

func (s *embededsSeq) At(i int) interface{} {
	return s.v.Embedded(i)
}

func (s *embededsSeq) Len() int {
	return s.v.NumEmbeddeds()
}

type methodsSeq struct {
	v *types.Interface
}

func (s *methodsSeq) At(i int) interface{} {
	return s.v.Method(i)
}

func (s *methodsSeq) Len() int {
	return s.v.NumMethods()
}

type Interface struct {
	Empty     *optional.Bool
	Methods   *optional.Set
	Embeddeds *optional.Set
}

func (q *Interface) Match(typ types.Type) bool {

	t, ok := typ.(*types.Interface)
	if !ok {
		return false
	}

	if !q.Embeddeds.Match(&embededsSeq{v: t}) {
		return false
	}

	if !q.Methods.Match(&methodsSeq{v: t}) {
		return false
	}

	return true
}

func (q *Interface) Exec(o types.Object) bool {
	if o == nil || o.Type() == nil {
		return false
	}
	return q.Match(o.Type())
}
