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
	*BaseMemberScope

	Name       string
	Expression node.Node
	Type       *Type
	Ellipses   bool
}

func NewLocalVarDecl(typ *Type, name string, expression node.Node) *LocalVarDecl {
	return &LocalVarDecl{
		Base:            node.New(),
		BaseMemberScope: NewMemberScope(),
		Type:            typ,
		Name:            name,
		Expression:      expression,
	}
}

func NewArgument(typ *Type, name string, ellipses bool) *LocalVarDecl {
	if typ == nil {
		panic(" no variable type")
	}
	if name == "" {
		panic("no variable name")
	}
	return &LocalVarDecl{
		Base:            node.New(),
		BaseMemberScope: NewMemberScope(),
		Type:            typ,
		Name:            name,
		Ellipses:        ellipses,
	}
}

func (f *LocalVarDecl) Children() []node.Node {
	list := []node.Node{}
	if f.Type != nil {
		list = append(list, f.Type)
	}
	if f.Expression != nil {
		list = append(list, f.Expression)
	}
	return list
}

func (f *LocalVarDecl) String() string {
	if f.Ellipses {
		// only formal parameters have these
		return fmt.Sprintf("%s %s...", f.Name, f.Type)
	}

	if f.Expression != nil {
		// definitaly a local var declaration
		return fmt.Sprintf("%s := %s", f.Name, f.Expression)
	}

	// might be a formal parameter, or a local var declaration
	return fmt.Sprintf("var %s %s", f.Name, f.Type)
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
