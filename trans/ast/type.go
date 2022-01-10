package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java-go-transpiler/input/parser"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

type TypePath struct {
	*node.Base

	/* may just be the path to the type */
	Elements []*TypeElement

	Class *Class // shortcut, for resolution later
}

// TypePath isn't an expression per say, but it has a GetType for expressions to use.
func (tp *TypePath) GetType() *Class {
	// always returns the last element of the path
	return tp.Elements[len(tp.Elements)-1].Class
}

func NewType(elements []*TypeElement) *TypePath {
	return &TypePath{Base: node.New(), Elements: elements}
}

func NewTypeOrVoid(typ *parser.TypeTypeOrVoidContext) *TypePath {
	typCtx := typ
	if typCtx.VOID() != nil {
		return NewType([]*TypeElement{NewTypeElement("void", nil)})
	} else {
		return NewTypeNodeFromContext(typCtx.TypeType())
	}
}

func NewTypeNodeFromContext(typ *parser.TypeTypeContext) *TypePath {
	if typ == nil {
		return nil
	}
	ctx := typ
	if ctx.PrimitiveType() != nil {
		// simple primitive type, easy to parse
		return NewType([]*TypeElement{
			NewTypeElement(ctx.PrimitiveType().GetText(), nil)})
	}

	classCtx := ctx.ClassOrInterfaceType()

	// multile components to one type. Say Car.WheelEnum
	elements := make([]*TypeElement, 0)
	for i, typID := range classCtx.AllIDENTIFIER() {
		className := typID.GetText()

		typeComp := classCtx.TypeArguments(i)
		if typeComp == nil {
			// no typ arguments for this element of the type
			elements = append(elements, NewTypeElement(className, nil))
			continue
		}

		typeCompCtx := typeComp

		// (typeArguments) multiple type args between <> seperated by commas
		thisTypeArguments := make([]*TypePath, 0)
		for _, typCompArg := range typeCompCtx.AllTypeArgument() {

			childType := NewTypeNodeFromContext(typCompArg.TypeType())
			thisTypeArguments = append(thisTypeArguments, childType)
		}

		node := NewTypeElement(className, thisTypeArguments)
		elements = append(elements, node)
	}

	return NewType(elements)
}

func (tn TypePath) String() string {
	list := make([]string, 0)
	for _, ten := range tn.Elements {
		list = append(list, ten.String())
	}
	return strings.Join(list, ".")
}

func (tn *TypePath) Children() []node.Node {
	if tn == nil {
		return nil
	}
	return node.ListOfNodesToNodeList(tn.Elements)
}

// used when defining the type of variable/return value/parameter/etc
type TypeElement struct {
	*node.Base
	*BaseClassScope

	ClassName     string
	TypeArguments []*TypePath

	Class *Class
}

func NewTypeElement(className string, typeArguments []*TypePath) *TypeElement {
	node := &TypeElement{
		Base:           node.New(),
		BaseClassScope: NewClassScope(),
		ClassName:      className,
		TypeArguments:  typeArguments,
	}
	return node
}

func (te *TypeElement) NodeName() string {
	if te.Class == nil {
		return te.ClassName + " /* unresolved */"
	}
	return te.Class.Name
}

func (te *TypeElement) Children() []node.Node {
	return node.ListOfNodesToNodeList(te.TypeArguments)
}

func (tn *TypeElement) String() string {

	star := "*"
	resolved := ""
	if tn.Class == nil {
		resolved = " /* Unresolved */"
	}

	if len(tn.TypeArguments) > 0 {

		list := make([]string, 0)
		for _, s := range tn.TypeArguments {
			list = append(list, s.String())
		}

		return fmt.Sprintf("%s%s[%s]%s", star, tn.ClassName, strings.Join(list, ","), resolved)
	}

	return star + tn.ClassName
}
