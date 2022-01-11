.PHONY: run test generate

GRAMMAR_FILES = input/grammar/JavaLexer.g4 input/grammar/JavaParser.g4
GO_SOURCE_FILES = $(shell find ./ -type f -name '*.go')
GO=go1.18beta1
BINARY=convert
GENERATED_VISITOR_FILES = trans/visitor/ast_visitor.g.go trans/visitor/base_ast_visitor.g.go

run: $(BINARY)
	./$(BINARY) -source "$(source)" -target "$(target)" -package "$(package)"

debug:
	dlv --api-version 2 --headless --listen :40000 debug cmd/convert/main.go -- $(target)

$(BINARY): go.* $(GO_SOURCE_FILES)
	$(GO) build -o $(BINARY) cmd/convert/main.go

test:
	$(GO) test ./...

ANTLR_GO_STG=../antlr4/tool/resources/org/antlr/v4/tool/templates/codegen/Go/Go.stg
GENERATED_PARSER_FILES= input/parser/JavaLexer.interp input/parser/JavaLexer.tokens input/parser/JavaParser.interp input/parser/JavaParser.tokens input/parser/java_lexer.go input/parser/java_parser.go input/parser/javaparser_base_listener.go input/parser/javaparser_listener.go

stg.jar: $(ANTLR_GO_STG)
	( cd ../antlr4/tool/resources/ && jar c org/antlr/v4/tool/templates/codegen/Go/Go.stg ) > stg.jar

generate: $(GENERATED_PARSER_FILES) $(GENERATED_VISITOR_FILES)

$(GENERATED_PARSER_FILES): stg.jar $(GRAMMAR_FILES)
	CLASSPATH="stg.jar:/usr/local/Cellar/antlr/4.9.3/antlr-4.9.3-complete.jar:." \
		/usr/local/opt/openjdk/bin/java \
		org.antlr.v4.Tool \
		-o input/parser -visitor -Dlanguage=Go -Xexact-output-dir $(GRAMMAR_FILES)

$(GENERATED_VISITOR_FILES): cmd/gen/main.go cmd/gen/*.tmpl trans/node/*.go trans/ast/*.go
	$(GO) run cmd/gen/main.go
