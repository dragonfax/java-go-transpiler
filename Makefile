
build:
	go build && ./delver_converter

grammer:
	antlr -o parser -Dlanguage=Go JavaLexer.g4 JavaParser.g4


