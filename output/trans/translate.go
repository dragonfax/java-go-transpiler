package trans

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java_converter/input/ast"
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/input/visitor"
)

var lexer = parser.NewJavaLexer(nil)
var p = parser.NewJavaParser(nil)

func init() {
	parser.RuleNames = p.RuleNames
}

func parseAST(path string) (*ast.File, error) {

	goVisitor := visitor.NewGoVisitor()

	input, _ := antlr.NewFileStream(path)
	lexer.SetInputStream(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(stream)
	p.BuildParseTrees = true
	file := goVisitor.Visit(p.CompilationUnit())

	return file.(*ast.File), nil
}

func dumpAST(file *ast.File) (string, error) {
	js, err := json.MarshalIndent(file, "", "  ")
	return string(js), err
}

// input filename to go-code string
func TranslateFile(filename string) (js string, goCode string, err error) {

	var goAST *ast.File
	goAST, err = parseAST(filename)
	if err != nil {
		return
	}

	js, err = dumpAST(goAST)
	if err != nil {
		return
	}

	goCode = goAST.String()

	goCode, err = goFmt(goCode)
	if err != nil {
		return
	}

	return js, goCode, err // goCode
}

// input filename to output filename
func TranslateFiles(filename, targetFilename string) {
	js, goCode, err := TranslateFile(filename)
	outputFile(targetFilename+".json", js)
	outputFile(targetFilename, goCode)
	if err != nil {
		if ferr, ok := err.(FormatingError); ok {
			outputFile(targetFilename+".err.txt", fmt.Sprintf("%s\n\n%s\n", ferr.Error(), ferr.Output))
		} else {
			outputFile(targetFilename+".err.txt", err.Error())
		}
	}
}

func outputFile(filename, code string) {
	err := ioutil.WriteFile(filename, []byte(code), 0640)
	if err != nil {
		panic(err)
	}
}

type FormatingError struct {
	error
	Output string
}

func goFmt(t string) (string, error) {
	source, err := format.Source([]byte(t))
	if err != nil {

		outputWriter := &strings.Builder{}

		// add line numbers to source output
		lineNumberWriter := newLineNumberWriter(outputWriter)
		lineNumberWriter.WriteString(t)

		return t, FormatingError{
			error:  err,
			Output: outputWriter.String(),
		}
	}

	return string(source), nil
}
