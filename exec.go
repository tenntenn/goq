package goq

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Result struct {
	Node         ast.Node
	Object       types.Object
	TypeAndValue types.TypeAndValue
}

func Exec(info *types.Info, q Query) []*Result {

	resultSets := map[token.Pos]*Result{}

	for i, o := range info.Defs {
		if q.Match(o) || q.Match(i) {
			resultSets[i.Pos()] = &Result{
				Node:   i,
				Object: o,
			}
		}
	}

	for i, o := range info.Uses {
		if q.Match(o) || q.Match(i) {
			resultSets[i.Pos()] = &Result{
				Node:   i,
				Object: o,
			}
		}
	}

	for e, tv := range info.Types {
		if q.Match(tv) || q.Match(e) {
			resultSets[e.Pos()] = &Result{
				Node:         e,
				TypeAndValue: tv,
			}
		}
	}

	results := make([]*Result, 0, len(resultSets))
	for _, r := range resultSets {
		results = append(results, r)
	}

	return results
}
