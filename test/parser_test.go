package test

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java_converter/input/ast/exp"
	"github.com/dragonfax/java_converter/input/parser"
)

/*
const casesDir = "./cases"
const exampleFileSuffix = ".example"

func TestFileParser(t *testing.T) {

	t.Skip()

	testPrefixes := []string{"implements", "interface", "static_variable", "math"}

	for _, testPrefix := range testPrefixes {
		javaFilename := casesDir + "/" + testPrefix + ".java" + exampleFileSuffix
		goFilename := casesDir + "/" + testPrefix + ".go" + exampleFileSuffix

		translatedGoCode := trans.TranslateFile(javaFilename)

		expectedGoCode, err := ioutil.ReadFile(goFilename)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, string(expectedGoCode), translatedGoCode)
	}

}
*/

func TestTypeNode(t *testing.T) {

	var lexer = parser.NewJavaLexer(nil)
	var p = parser.NewJavaParser(nil)

	content := `Level.Source`

	input := antlr.NewInputStream(content)
	lexer.SetInputStream(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(stream)
	p.BuildParseTrees = true
	ptree := p.TypeType()

	source := exp.NewTypeNode(ptree)

	if "Level.Source" != source.String() {
		t.Fatal("differnces")
	}
}
