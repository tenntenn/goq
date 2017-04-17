package goq

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Goq struct {
	fset  *token.FileSet
	info  *types.Info
	files []*ast.File
}

func New(fset *token.FileSet, files []*ast.File, info *types.Info) *Goq {
	return &Goq{
		fset:  fset,
		files: files,
		info:  info,
	}
}

func (gq *Goq) Query(q Query) Results {
	resultSets := map[token.Pos]*Result{}

	for i, o := range gq.info.Defs {
		if q.Match(o) || q.Match(i) {
			resultSets[i.Pos()] = &Result{
				Goq:    gq,
				Node:   i,
				Object: o,
			}
		}
	}

	for i, o := range gq.info.Uses {
		if q.Match(o) || q.Match(i) {
			resultSets[i.Pos()] = &Result{
				Goq:    gq,
				Node:   i,
				Object: o,
			}
		}
	}

	for e, tv := range gq.info.Types {
		if q.Match(tv.Type) || q.Match(e) {
			resultSets[e.Pos()] = &Result{
				Goq:          gq,
				Node:         e,
				TypeAndValue: tv,
			}
		}
	}

	for n, o := range gq.info.Implicits {
		if q.Match(o) || q.Match(n) {
			resultSets[n.Pos()] = &Result{
				Goq:    gq,
				Node:   n,
				Object: o,
			}
		}
	}

	results := make([]*Result, 0, len(resultSets))
	for _, r := range resultSets {
		results = append(results, r)
	}

	return results
}

type Result struct {
	Goq          *Goq
	Node         ast.Node
	Object       types.Object
	TypeAndValue types.TypeAndValue
}

func (r *Result) filter(q Query) bool {
	return q.Match(r.Object) || q.Match(r.TypeAndValue) || q.Match(r.Node)
}

type Results []*Result

func (rs Results) Filter(q Query) Results {
	filtered := make(Results, 0, len(rs))
	for _, r := range rs {
		if r.filter(q) {
			filtered = append(filtered, r)
		}
	}
	return filtered[:len(filtered):len(filtered)]
}
