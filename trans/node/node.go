package node

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/dragonfax/java_converter/tool"
	"github.com/tkrajina/go-reflector/reflector"
)

type Node interface {
	String() string
	Children() []Node
	SetParent(Node)
	GetParent() Node
}

type Base struct {
	Parent Node
}

func New() *Base {
	return &Base{}
}

func (bn *Base) SetParent(p Node) {
	if bn == nil {
		panic("nil node in SetParent")
	}
	bn.Parent = p
}

func (bn *Base) GetParent() Node {
	return bn.Parent
}

func MarshalNode(node Node) interface{} {
	if tool.IsNilInterface(node) {
		return nil
	}

	childNodes := node.Children()
	children := make([]interface{}, 0)
	for _, childNode := range childNodes {
		if childNode != nil {
			children = append(children, MarshalNode(childNode))
		}
	}

	m := map[string]interface{}{
		"Name": Name(node),
		"go":   fmt.Sprintf("%T", node),
	}

	if len(children) > 0 {
		m["children"] = children
	}

	obj := reflector.New(node)
	for _, field := range obj.FieldsFlattened() {

		if !field.IsValid() {
			continue
		}

		if !field.IsExported() {
			continue
		}
		switch field.Kind() {
		case reflect.Array, reflect.Uintptr, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Struct, reflect.UnsafePointer:
			continue
		}

		value, err := field.Get()
		if err != nil {
			fmt.Printf(err.Error())
		}
		m[field.Name()] = value

	}

	return m
}

func JSONMarshalNode(node Node) string {
	data := MarshalNode(node)
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(js)
}

func Name(node Node) string {
	if nameNode, ok := node.(interface{ NodeName() string }); ok {
		return nameNode.NodeName()
	}

	obj := reflector.New(node)
	field := obj.Field("Name")
	if field.IsValid() {
		nameValue, _ := field.Get()
		if s, ok := nameValue.(string); ok {
			return s
		}
	}

	t := fmt.Sprintf("%T", node)

	return strings.TrimPrefix(t, "*ast.")
}
