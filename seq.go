package goq

import "go/types"

type pkgImportsSeq struct {
	v *types.Package
}

func (s *pkgImportsSeq) At(i int) interface{} {
	return s.v.Imports()[i]
}

func (s *pkgImportsSeq) Len() int {
	return len(s.v.Imports())
}

type namedMethodsSeq struct {
	v *types.Named
}

func (s *namedMethodsSeq) At(i int) interface{} {
	return s.v.Method(i)
}

func (s *namedMethodsSeq) Len() int {
	return s.v.NumMethods()
}

func methodEq(v1, v2 interface{}) bool {
	t1, ok := v1.(*Func)
	if !ok {
		return false
	}

	t2, ok := v2.(*types.Func)
	if !ok {
		return false
	}

	return t1.Exec(t2)
}
