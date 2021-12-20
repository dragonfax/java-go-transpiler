package output

import (
	"errors"
	"io/ioutil"
	"os"
)

func outputFile(filename, code string) {
	err := ioutil.WriteFile(filename, []byte(code), 0640)
	if err != nil {
		panic(err)
	}
}

func removeFileIfExists(filename string) {
	// remove target file if its already there.
	_, err := os.Stat(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	} else {
		// file exists. thats a problem.
		err = os.Remove(filename)
		if err != nil {
			panic(err)
		}
	}

}
