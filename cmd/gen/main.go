package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

/*

1. crawl the node dir for nodes that implmentat *Base
   just need a node name.
2. run the tempaltes over that data.

*/

var templateFiles = []string{"ast_visitor.g.go.tmpl", "base_ast_visitor.g.go.tmpl"}

const templateDir = "cmd/gen"
const targetDir = "trans/visitor"
const nodesDir = "trans/ast"

func main() {

	nodeList := make([]string, 0)

	err := filepath.WalkDir(nodesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet() // positions are relative to fset

		f, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
		if err != nil {
			return err
		}

		for _, decl := range f.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			if genDecl.Tok != token.TYPE {
				continue
			}

			spec, ok := genDecl.Specs[0].(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := spec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			field := structType.Fields.List[0]

			star, ok := field.Type.(*ast.StarExpr)
			if !ok {
				continue
			}

			sel, ok := star.X.(*ast.SelectorExpr)
			if !ok {
				continue
			}

			if sel.Sel.Name != "Base" || sel.X.(*ast.Ident).Name != "node" {
				continue
			}

			name := spec.Name.Name
			if name == "BaseOperator" {
				continue
			}

			nodeList = append(nodeList, name)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, templateFile := range templateFiles {
		targetFile := targetDir + "/" + strings.TrimSuffix(templateFile, ".tmpl")

		// read file
		body, err := ioutil.ReadFile(templateDir + "/" + templateFile)
		if err != nil {
			panic(err)
		}

		// apply template
		t, err := template.New("templateFile").Parse(string(body))
		if err != nil {
			panic(err)
		}
		writer := bytes.NewBuffer(make([]byte, 0, 1024*20))
		err = t.Execute(writer, nodeList)
		if err != nil {
			panic(err)
		}

		// gofmt source
		source, err := format.Source(writer.Bytes())
		if err != nil {
			panic(err)
		}

		// write file
		file, err := os.Create(targetFile)
		if err != nil {
			panic(err)
		}
		_, err = file.Write(source)
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	fmt.Println("files generated")
}
