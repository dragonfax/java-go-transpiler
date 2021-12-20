package main

import (
	"fmt"
	"os"

	"github.com/dragonfax/java_converter/output"
)

func main() {

	source := os.Args[1]
	err := output.Translate(source)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
