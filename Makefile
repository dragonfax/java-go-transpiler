.PHONY: run test

GRAMMAR_FILES = input/grammar/JavaLexer.g4 input/grammar/JavaParser.g4
GO_SOURCE_FILES = $(shell find ./ -type f -name '*.go')

run: java_converter
	./java_converter $(target)

debug:
	dlv --api-version 2 --headless --listen :40000 debug cmd/main.go -- $(target)

java_converter: go.* $(GO_SOURCE_FILES)
	go build -o java_converter cmd/main.go

parser/*: $(GRAMMAR_FILES)
	cd input/grammar && antlr -o ../../parser -Dlanguage=Go JavaLexer.g4 JavaParser.g4

test:
	go test ./...

