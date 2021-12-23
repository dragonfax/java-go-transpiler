module github.com/dragonfax/java_converter

go 1.18

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211208212222-82c441726976
	github.com/schollz/progressbar/v3 v3.8.5
	github.com/tkrajina/go-reflector v0.5.5
)

require (
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
)

replace github.com/antlr/antlr4/runtime/Go/antlr => ../antlr4/runtime/Go/antlr
