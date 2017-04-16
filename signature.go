package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ TypeMatcher = (*Signature)(nil)
)

type Signature struct {
	// Recv is type of the receiver.
	Recv *VarQuery
	// Params are parameters of the function.
	Params *optional.Tuple
	// Results are results of the function.
	Results *optional.Tuple
	// Variadic is whether the function uses variadic parameters or not.
	Variadic *optional.Bool
}

func (tm *Signature) Match(t types.Type) bool {
	s, ok := t.(*types.Signature)
	if !ok {
		return false
	}

	if !tm.Variadic.Match(s.Variadic()) {
		return false
	}

	if tm.Recv != nil && !tm.Recv.Exec(s.Recv()) {
		return false
	}

	if !tm.Params.Match(&signatureParamsSeq{v: s}) {
		return false
	}

	if !tm.Results.Match(&signatureResultsSeq{v: s}) {
		return false
	}

	return true
}

type signatureParamsSeq struct {
	v *types.Signature
}

func (s signatureParamsSeq) At(i int) interface{} {
	return s.v.Params().At(i)
}

func (s signatureParamsSeq) Len() int {
	return s.v.Params().Len()
}

type signatureResultsSeq struct {
	v *types.Signature
}

func (s signatureResultsSeq) At(i int) interface{} {
	return s.v.Results().At(i)
}

func (s signatureResultsSeq) Len() int {
	return s.v.Results().Len()
}
