.PHONY = run

GRAMMAR_FILES = input/grammar/JavaLexer.g4 input/grammar/JavaParser.g4

run: java_converter
	# rm ../java_converted/*.go || true
	./java_converter ../java_converted $(file)

java_converter: go.* */*.go
	go build -o java_converter cmd/main.go

parser/*: $(GRAMMAR_FILES)
	cd grammar && antlr -o ../parser -Dlanguage=Go JavaLexer.g4 JavaParser.g4

test:
	go test ./...

