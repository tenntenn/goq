package goq

import (
	"go/ast"
	"reflect"
)

var (
	_ Query = (*Node)(nil)
)

type Node struct {
	Files []*ast.File
	Type  reflect.Type
	Path  []Query
}

func (q *Node) Match(v interface{}) bool {
	n, ok := v.(ast.Node)
	if !ok {
		return false
	}

	if q.Type != nil && !reflect.TypeOf(n).AssignableTo(q.Type) {
		return false
	}

	if q.Path != nil && !q.has(n) {
		return false
	}

	return true
}

func (q *Node) has(n ast.Node) bool {
	for _, f := range q.Files {
		path := NodePath(f, n)
		for i := range path {
			matched := true
			for j := range q.Path {
				if len(path) <= i+j || !q.Path[j].Match(path[i+j]) {
					matched = false
					break
				}
			}
			if matched {
				return true
			}
		}
	}
	return false
}

type nodeTracer struct {
	nodePath []ast.Node
	target   ast.Node
	stop     bool
}

func (nt *nodeTracer) Visit(n ast.Node) ast.Visitor {
	if nt.stop {
		return nil
	}

	if n == nil {
		nt.nodePath = nt.nodePath[:len(nt.nodePath)-1]
		return nil
	}

	nt.nodePath = append(nt.nodePath, n)

	if n == nt.target {
		nt.stop = true
		return nil
	}

	return nt
}

func NodePath(root, node ast.Node) []ast.Node {
	nt := &nodeTracer{target: node}
	ast.Walk(nt, root)
	return nt.nodePath
}
