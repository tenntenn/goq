# Type Query

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

	if _, err := config.Check("main", fset, []*ast.File{f}, info); err != nil {
		return err
	}

	q := Int(nil)
	objects := Exec(info, q)
    for _, o := range objects {
        fmt.Println(o, "at", fset.Pos(o.Pos()))
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
