/* definitions for the java.lang packgae */
package java_lang

import "github.com/dragonfax/java-go-transpiler/trans/ast"

var Package = ast.NewPackage("java.lang")

func init() {

	stdOut := ast.NewClass()
	stdOut.Name = "StdOut"
	stdOut.Methods = []*ast.Method{ast.NewMethod("println", nil, nil, nil)}
	stdOut.PackageScope = Package

	Package.AddClass(stdOut)

	system := ast.NewClass()
	system.Name = "System"
	system.Fields = []*ast.Field{ast.NewField(stdOut, "out", nil)}
	system.PackageScope = Package

	Package.AddClass(system)
}
