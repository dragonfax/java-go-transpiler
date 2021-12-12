.PHONY = run

GRAMMAR_FILES = grammar/JavaLexer.g4 grammar/JavaParser.g4

run: delver_converter
	rm ../delver_converted/*.go || true
	./delver_converter ../delver_converted $(file)

delver_converter: go.* */*.go
	go build -o delver_converter cmd/main.go

parser/*: $(GRAMMAR_FILES)
	cd grammar && antlr -o ../parser -Dlanguage=Go JavaLexer.g4 JavaParser.g4

test:
	go test ./...

