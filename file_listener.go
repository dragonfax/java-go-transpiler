package main

import "github.com/dragonfax/delver_converter/parser"

type FileListener struct {
	*parser.BaseJavaParserListener

	inImport bool

	File     *File
	Filename string
}

func (s *FileListener) EnterCompilationUnit(ctx *parser.CompilationUnitContext) {
	s.File = NewFile()
}

func (s *FileListener) EnterQualifiedName(ctx *parser.QualifiedNameContext) {
	if s.File.QualifiedPackageName == "" {
		// first one, must be the package name
		s.File.QualifiedPackageName = ctx.GetText()
	}

	/* if s.inImport {
		s.File.Imports = append(s.File.Imports, ctx.GetText())
	}*/
}

func (s *FileListener) ExitCompilationUnit(ctx *parser.CompilationUnitContext) {
	// file is done.
}

func (s *FileListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	s.inImport = true
}

func (s *FileListener) ExitImportDeclaration(ctx *parser.ImportDeclarationContext) {
	s.inImport = false
}

func (s *FileListener) EnterTypeDeclaration(ctx *parser.TypeDeclarationContext) {
	stackListener.Push(NewClassListener(s.File))
}
