package trans

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type lineNumberWriter struct {
	lineNumber int
	output     *os.File
}

func newLineNumberWriter(output *os.File) io.StringWriter {
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
