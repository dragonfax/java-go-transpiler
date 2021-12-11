.PHONY = run

run: delver_converter
	./delver_converter

delver_converter: go.* *.go parser/*
	go build

parser/*: JavaLexer.g4 JavaParser.g4
	antlr -o parser -Dlanguage=Go JavaLexer.g4 JavaParser.g4


