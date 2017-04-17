package goq

import (
	"go/ast"
	"reflect"
)

var (
	_ Query = (*Node)(nil)
)

type Node struct {
	Type  reflect.Type
	Child reflect.Type
}

func (q *Node) Match(v interface{}) bool {
	n, ok := v.(ast.Node)
	if !ok {
		return false
	}

	if q.Type != nil && !reflect.TypeOf(n).AssignableTo(q.Type) {
		return false
	}

	if q.Child != nil {
		var hasChild bool
		ast.Inspect(n, func(cn ast.Node) bool {
			if hasChild {
				return false
			}
			hasChild = q.Match(cn)
			return true
		})

		if !hasChild {
			return false
		}
	}

	return true
}
