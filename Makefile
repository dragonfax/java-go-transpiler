.PHONY: generate

CLASSPATH="/usr/local/Cellar/antlr/4.9.3/antlr-4.9.3-complete.jar:."

generate:
	antlr -o parser JavaLexer.g4 JavaParser.g4
