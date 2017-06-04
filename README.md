# goq [![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godoc] [![Travis](https://img.shields.io/travis/tenntenn/goq.svg?style=flat-square)][travis]

[godoc]: http://godoc.org/github.com/tenntenn/goq
[travis]: https://travis-ci.org/tenntenn/goq

`goq` is a Go package to access nodes of Go's AST by queries.

See usage and example in [GoDoc](https://godoc.org/github.com/tenntenn/goq).

*NOTE*: This package is experimental and may make backward-incompatible changes.

## Install

Use go get:

```
$ go get github.com/tenntenn/goq
```

## Usage

All usage are described in [GoDoc](https://godoc.org/github.com/tenntenn/goq).

```
const src = `package main
func main() {
	n := 10
	println(n)
}`

func run() error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		return err
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}

	files := []*ast.File{f}
	if _, err := config.Check("main", fset, files, info); err != nil {
		return err
	}

	results := goq.New(fset, info, files).Query(&Int{})
	for _, r := range results {
		fmt.Println(r.Object, "at", fset.Pos(r.Node.Pos()))
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
```
