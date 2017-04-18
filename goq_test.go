package goq_test

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	"testing"

	. "github.com/tenntenn/goq"
	"github.com/tenntenn/optional"
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
	results := New(fset, []*ast.File{f}, info).Query(q)
	if len(results) != 2 {
		t.Fatalf("the number of result must be 2 but %d", len(results))
	}

	if n := results[0].Object.Name(); n != "n" {
		t.Errorf(`exepect object name is "n" but %q`, n)
	}
}

func TestError(t *testing.T) {
	const src = `package main
	type Err string
	func (err Err) Error() string {return string(err)}
	func f() (int, Err) { // 4
		return 0, Err("hoge")
	}
	func main() {
		println(f()) // 8
		println(func() Err { // 9
			return Err("fuga")
		}())
		println(func() error { // 12
			return nil
		}())
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
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
		Types: map[ast.Expr]types.TypeAndValue{},
	}

	files := []*ast.File{f}
	if _, err := config.Check("main", fset, files, info); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	errorType := types.Universe.Lookup("error").Type()
	results := New(fset, []*ast.File{f}, info).Query(&Signature{
		Results: NewTupple(nil).Put(optional.Last, &Var{
			Type: And(&Type{
				Implements: errorType,
			}, Not(&Type{
				Identical: errorType,
			})),
		}),
	}).Filter(Or(&Node{
		Files: files,
		Path: []Query{
			&Node{
				Type: reflect.TypeOf((*ast.FuncLit)(nil)),
			},
		},
	}, &Node{
		Files: files,
		Path: []Query{
			&Node{
				Type: reflect.TypeOf((*ast.FuncDecl)(nil)),
			},
			&Node{
				Type: reflect.TypeOf((*ast.Ident)(nil)),
			},
		},
	}))

	if len(results) != 2 {
		t.Fatalf("the number of result must be 2 but %d", len(results))
	}

	if pos := results[0].Node.Pos(); fset.Position(pos).Line != 4 {
		t.Errorf(`exepect first line is 4 but %d`, fset.Position(pos).Line)
	}

	if pos := results[1].Node.Pos(); fset.Position(pos).Line != 9 {
		t.Errorf(`exepect first line is 9 but %d`, fset.Position(pos).Line)
	}
}
