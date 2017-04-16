package goq_test

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	. "github.com/tenntenn/goq"
)

func TestExec(t *testing.T) {
	const src = `package main
	func main() {
		n := 10
		println(n)
	}`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}

	if _, err := config.Check("main", fset, []*ast.File{f}, info); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	q := &Int{}
	results := Exec(info, q)
	if len(results) != 2 {
		t.Fatalf("the number of result must be 2 but %d", len(results))
	}

	if n := results[0].Object.Name(); n != "n" {
		t.Errorf(`exepect object name is "n" but %q`, n)
	}
}
