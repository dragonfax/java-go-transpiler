package output

import (
	"errors"
	"go/format"
	"strings"
)

var FormatingError = errors.New("formating error")

func goFmt(t string) (string, error) {
	source, err := format.Source([]byte(t))
	if err != nil {

		outputWriter := &strings.Builder{}

		// add line numbers to source output
		lineNumberWriter := newLineNumberWriter(outputWriter)
		lineNumberWriter.WriteString(t)

		return outputWriter.String(), FormatingError
	}

	return string(source), nil
}
