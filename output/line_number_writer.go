package output

import (
	"fmt"
	"io"
	"strings"
)

type lineNumberWriter struct {
	lineNumber int
	output     io.StringWriter
}

func newLineNumberWriter(output io.StringWriter) io.StringWriter {
	return &lineNumberWriter{lineNumber: 1, output: output}
}

func (lnw *lineNumberWriter) WriteString(s string) (int, error) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		lnw.output.WriteString(fmt.Sprintf("%d %s\n", lnw.lineNumber, line))
		lnw.lineNumber++
	}

	return 0, nil
}
