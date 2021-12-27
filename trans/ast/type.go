package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

var primitiveTranslation = map[string]string{
	"float":   "float64",
	"double":  "float64",
	"boolean": "bool",
	"long":    "int64",
	"String":  "string",
	"void":    "",
}

type TypePath struct {
	*node.Base

	/* may just be the path to the type */
	Elements []*TypeElement

	Class *Class // shortcut, for resolution later
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

func (te *TypeElement) IsPrimitive() bool {
	switch te.ClassName {
	case "bypte", "float", "short", "int", "long", "double", "boolean", "char", "String", "void":
		return true
	default:
		return false
	}

}

func (tn *TypeElement) String() string {

	star := "*"
	if tn.IsPrimitive() {
		star = ""
	}

	className := tn.ClassName
	if tClassName, ok := primitiveTranslation[className]; ok {
		className = tClassName
	}

	if len(tn.TypeArguments) > 0 {

		list := make([]string, 0)
		for _, s := range tn.TypeArguments {
			list = append(list, s.String())
		}

		return fmt.Sprintf("%s%s[%s]", star, className, strings.Join(list, ","))
	}

	return star + className
}
