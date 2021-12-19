module github.com/dragonfax/java_converter

go 1.18

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211208212222-82c441726976
	github.com/aymerick/raymond v2.0.2+incompatible
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/antlr/antlr4/runtime/Go/antlr => ../antlr4/runtime/Go/antlr
