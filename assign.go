package goq

import (
	"go/ast"
	"go/token"

	"github.com/tenntenn/optional"
)

var (
	_ Query = (*AssignStmtNode)(nil)
)

type AssignStmtNode struct {
	Lhs *optional.Set
	Tok *token.Token
	Rhs *optional.Set
}

func (q *AssignStmtNode) Match(v interface{}) bool {
	n, ok := v.(*ast.AssignStmt)
	if !ok {
		return false
	}

	if q.Tok != nil && *q.Tok != n.Tok {
		return false
	}

	if !q.Lhs.Match(&assignStmtLhs{v: n}) {
		return false
	}

	if !q.Rhs.Match(&assignStmtRhs{v: n}) {
		return false
	}

	return true
}

type assignStmtLhs struct {
	v *ast.AssignStmt
}

func (a assignStmtLhs) At(i int) interface{} {
	return a.v.Lhs[i]
}

func (s assignStmtLhs) Len() int {
	return len(s.v.Lhs)
}

type assignStmtRhs struct {
	v *ast.AssignStmt
}

func (a assignStmtRhs) At(i int) interface{} {
	return a.v.Rhs[i]
}

func (s assignStmtRhs) Len() int {
	return len(s.v.Rhs)
}
