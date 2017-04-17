package goq

import (
	"go/types"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*Signature)(nil)
)

type Signature struct {
	// Recv is type of the receiver.
	Recv *Var
	// Params are parameters of the function.
	Params *optional.Tuple
	// Results are results of the function.
	Results *optional.Tuple
	// Variadic is whether the function uses variadic parameters or not.
	Variadic *optional.Bool
}

// Match implements Query.Match.
func (tm *Signature) Match(v interface{}) bool {
	s, ok := toType(v).(*types.Signature)
	if !ok {
		return false
	}

	if !tm.Variadic.Match(s.Variadic()) {
		return false
	}

	if tm.Recv != nil && !tm.Recv.Match(s.Recv()) {
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
	rs := s.v.Results()
	if rs == nil {
		return nil
	}
	return rs.At(i)
}

func (s signatureResultsSeq) Len() int {
	rs := s.v.Results()
	if rs == nil {
		return 0
	}
	return rs.Len()
}
