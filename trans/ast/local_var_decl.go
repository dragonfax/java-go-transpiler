package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

/* A local variable declaration,
 * but also re-used for a formal parameter in a method or constructor declaration.
 * Not in a method call though, thats just an expression.
 */
type LocalVarDecl struct {
	*node.Base
	*BaseMethodScope

	Name       string
	Expression Expression
	TypePath   *TypePath
	Ellipses   bool
	IsArgument bool
}

func NewLocalVarDecl(typ *TypePath, name string, expression Expression) *LocalVarDecl {
	return &LocalVarDecl{
		Base:            node.New(),
		BaseMethodScope: NewMethodScope(),
		TypePath:        typ,
		Name:            name,
		Expression:      expression,
	}
}

func NewArgument(typ *TypePath, name string, ellipses bool) *LocalVarDecl {
	if typ == nil {
		panic(" no variable type")
	}
	if name == "" {
		panic("no variable name")
	}
	return &LocalVarDecl{
		Base:            node.New(),
		BaseMethodScope: NewMethodScope(),
		TypePath:        typ,
		Name:            name,
		Ellipses:        ellipses,
		IsArgument:      true,
	}
}

func (f *LocalVarDecl) Children() []node.Node {
	list := []node.Node{}
	if f.TypePath != nil {
		list = append(list, f.TypePath)
	}
	if f.Expression != nil {
		list = append(list, f.Expression)
	}
	return list
}

func (f *LocalVarDecl) String() string {
	if f.Ellipses {
		// only formal parameters have these
		return fmt.Sprintf("%s %s...", f.Name, f.TypePath)
	}

	if f.Expression != nil {
		// definitaly a local var declaration
		return fmt.Sprintf("%s := %s", f.Name, f.Expression)
	}

	varTerm := "var "
	if f.IsArgument {
		varTerm = ""
	}

	// might be a formal parameter, or a local var declaration
	return fmt.Sprintf("%s%s %s", varTerm, f.Name, f.TypePath)
}

func NewLocalVarDeclNodeList(decl *parser.LocalVariableDeclarationContext) []*LocalVarDecl {

	typ := NewTypeNodeFromContext(decl.TypeType())

	list := make([]*LocalVarDecl, 0)
	for _, varDecl := range decl.VariableDeclarators().AllVariableDeclarator() {

		varDeclCtx := varDecl

		var exp node.Node
		if varDeclCtx.VariableInitializer() != nil {
			varInitCtx := varDeclCtx.VariableInitializer()
			exp = variableInitializerProcessor(varInitCtx)
		}

		name := varDeclCtx.VariableDeclaratorId().GetText()

		node := NewLocalVarDecl(typ, name, exp)

		list = append(list, node)
	}

	return list
}

func variableInitializerProcessor(ctx *parser.VariableInitializerContext) node.Node {
	var exp node.Node
	if ctx.Expression() != nil {
		exp = ExpressionProcessor(ctx.Expression())
	}
	if ctx.ArrayInitializer() != nil {
		exp = NewArrayLiteral(ctx.ArrayInitializer())
	}

	return exp
}
