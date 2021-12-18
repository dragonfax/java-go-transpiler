.PHONY: run debug test generate

GENERATED_PARSER_FILES= parser/JavaLexer.interp parser/JavaLexer.tokens parser/JavaParser.interp parser/JavaParser.tokens parser/java_lexer.go parser/java_parser.go parser/javaparser_base_listener.go parser/javaparser_listener.go
GRAMMAR_FILES = JavaLexer.g4 JavaParser.g4
GO_SOURCE_FILES = $(shell find ./ -type f -name '*.go')
ANTLR_GO_RUNTIME=../antlr4/runtime/Go
ANTLR_GO_FILES = $(shell find $(ANTLR_GO_RUNTIME) type f -name '*.go')
BINARY=java_visitor

run: $(BINARY)
	./$(BINARY)

debug:
	dlv --api-version 2 --headless --listen :40000 debug main.go

$(BINARY): go.* $(GO_SOURCE_FILES) $(GENERATED_PARSER_FILES)
	go build -o $(BINARY) main.go

generate: $(GENERATED_PARSER_FILES)

$(GENERATED_PARSER_FILES): stg.jar $(GRAMMAR_FILES)
	CLASSPATH="stg.jar:/usr/local/Cellar/antlr/4.9.3/antlr-4.9.3-complete.jar:." \
		/usr/local/opt/openjdk/bin/java \
		-jar /usr/local/Cellar/antlr/4.9.3/antlr-4.9.3-complete.jar \
		-o parser -Dlanguage=Go -visitor -cp stg.jar JavaLexer.g4 JavaParser.g4

test:
	go test ./...


ANTLR_GO_STG=../antlr4/tool/resources/org/antlr/v4/tool/templates/codegen/Go/Go.stg

stg.jar: $(ANTLR_GO_STG)
	( cd ../antlr4/tool/resources/ && jar c org/antlr/v4/tool/templates/codegen/Go/Go.stg ) > stg.jar