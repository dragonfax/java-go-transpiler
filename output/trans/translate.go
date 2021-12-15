package trans

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/aymerick/raymond"
	"github.com/dragonfax/java_converter/input/ast"
	"github.com/dragonfax/java_converter/input/listen"
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

const golangTemplateFilename = "output/templates/golang.tmpl"

var lexer = parser.NewJavaLexer(nil)
var p = parser.NewJavaParser(nil)

func init() {
	parser.RuleNames = p.RuleNames
}

var golangTemplate = tool.MustByteListErr(ioutil.ReadFile(golangTemplateFilename))

func parseAST(path string) (*ast.File, error) {

	stackListener := listen.NewStackListener()

	fileListener := listen.NewFileListener(stackListener, path)
	stackListener.Push(fileListener)

	input, _ := antlr.NewFileStream(path)
	lexer.SetInputStream(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(stackListener, p.CompilationUnit())

	if stackListener.Len() != 1 {
		if stackListener.Len() == 2 {
			panic(fmt.Sprintf("left over listener %T", stackListener.Peek()))
		}
		return nil, errors.New("wrong number of listeners left over, " + fmt.Sprintf("%d", stackListener.Len()))
	}

	/*
		js, err := json.MarshalIndent(p.CompilationUnit(), "", "  ")
		if err != nil {
			panic(err)
		}
	*/
	// fmt.Println(p.CompilationUnit().ToStringTree())

	return fileListener.File, nil
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

	goCode, err = generateGo(goAST)
	if err != nil {
		return
	}

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

func generateGo(file *ast.File) (string, error) {
	return raymond.Render(string(golangTemplate), file)
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
