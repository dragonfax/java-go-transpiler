package test

import (
	"io/ioutil"
	"testing"

	"github.com/dragonfax/java_converter/output/trans"
	"github.com/stretchr/testify/assert"
)

const casesDir = "./cases"
const exampleFileSuffix = ".example"

func TestParser(t *testing.T) {

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
