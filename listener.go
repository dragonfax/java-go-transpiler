package main

import "github.com/dragonfax/delver_converter/parser"

type listener struct {
	*parser.BaseJavaParserListener

	inImport bool

	File     *File
	Filename string
}

func (s *listener) EnterCompilationUnit(ctx *parser.CompilationUnitContext) {
	s.File = NewFile()
}

func (s *listener) EnterQualifiedName(ctx *parser.QualifiedNameContext) {
	if s.File.QualifiedPackageName == "" {
		// first one, must be the package name
		s.File.QualifiedPackageName = ctx.GetText()
	}

	if s.inImport {
		s.File.Imports = append(s.File.Imports, ctx.GetText())
	}
}

func (s *listener) ExitCompilationUnit(ctx *parser.CompilationUnitContext) {
	// file is done.
}

func (s *listener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	s.inImport = true
}

func (s *listener) ExitImportDeclaration(ctx *parser.ImportDeclarationContext) {
	s.inImport = false
}

func (s *listener) EnterTypeType(ctx *parser.TypeTypeContext) {
	if s.File.BaseClass == "" {
		s.File.BaseClass = ctx.ClassOrInterfaceType().GetText()
	}
}
