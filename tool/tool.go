package tool

import (
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java-go-transpiler/input/parser"
)

func MustByteListErr(buf []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return buf
}

func IsNilInterface(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func PanicDebug(msg string, ctx antlr.ParseTree) {
	panic(msg + ": " + ctx.GetText() + "\n\n" + ctx.ToStringTree(parser.RuleNames, nil))
}
