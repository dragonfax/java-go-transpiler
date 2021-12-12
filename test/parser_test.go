package test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const casesDir = "./cases"

func TestParser(t *testing.T) {

	testPrefixes := []string{"implements", "interface", "static_variable"}

	for _, testPrefix := range testPrefixes {
		java, err := ioutil.ReadFile(casesDir + "/" + testPrefix + ".java")
		if err != nil {
			panic(err)
		}
		goCode, err := ioutil.ReadFile(casesDir + "/" + testPrefix + ".go")
		if err != nil {
			panic(err)
		}
		result := trans.Translate(java)
		assert.Equal(t, goCode, result)
	}

}
