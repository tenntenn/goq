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
	objects := Exec(info, q)
	if len(objects) != 1 {
		t.Fatalf("the number of result must be 1 but %d", len(objects))
	}

	if n := objects[0].Name(); n != "n" {
		t.Errorf(`exepect object name is "n" but %q`, n)
	}
}
