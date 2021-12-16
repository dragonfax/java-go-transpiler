.PHONY: run parser

GRAMMAR_FILES = JavaLexer.g4 JavaParser.g4
GO_SOURCE_FILES = $(shell find ./ -type f -name '*.go')
BINARY=java_visitor

run: $(BINARY)
	./$(BINARY) $(target)

debug:
	dlv --api-version 2 --headless --listen :40000 debug main.go -- $(target)

$(BINARY): go.* $(GO_SOURCE_FILES)
	go build -o java_visitor main.go

parser: $(GRAMMAR_FILES)
	antlr -o parser -Dlanguage=Go JavaLexer.g4 JavaParser.g4

test:
	go test ./...

