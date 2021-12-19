module github.com/dragonfax/java_converter

go 1.18

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211208212222-82c441726976
)

require (
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/antlr/antlr4/runtime/Go/antlr => ../antlr4/runtime/Go/antlr
