package test

import (
	"io/ioutil"
	"testing"

	"github.com/dragonfax/delver_converter/trans"
	"github.com/stretchr/testify/assert"
)

const casesDir = "./cases"
const exampleFileSuffix = ".example"

func TestParser(t *testing.T) {

	testPrefixes := []string{"implements", "interface", "static_variable", "math"}

	for _, testPrefix := range testPrefixes {
		java, err := ioutil.ReadFile(casesDir + "/" + testPrefix + ".java" + exampleFileSuffix)
		if err != nil {
			panic(err)
		}
		goCode, err := ioutil.ReadFile(casesDir + "/" + testPrefix + ".go" + exampleFileSuffix)
		if err != nil {
			panic(err)
		}
		result := trans.Translate(string(java))
		assert.Equal(t, string(goCode), result)
	}

}
