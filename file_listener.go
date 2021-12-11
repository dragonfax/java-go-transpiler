package main

import "github.com/dragonfax/delver_converter/parser"

type FileListener struct {
	*StackableListener
	*parser.BaseJavaParserListener

	inImport bool

	File *File
}

func NewFileListener(sl *StackListener, filename string) *FileListener {
	f := &FileListener{
		StackableListener: NewStackableListener(sl),
		File:              NewFile(),
	}
	f.File.Filename = filename
	return f
}

func (s *FileListener) EnterCompilationUnit(ctx *parser.CompilationUnitContext) {
	s.File = NewFile()
}

func (s *FileListener) EnterQualifiedName(ctx *parser.QualifiedNameContext) {
	if s.File.QualifiedPackageName == "" {
		// first one, must be the package name
		s.File.QualifiedPackageName = ctx.GetText()
	}

	// don't need imports.
	/* if s.inImport {
		s.File.Imports = append(s.File.Imports, ctx.GetText())
	}*/
}

/* File shouldn't pop itself, its the root element of the stack. the parser cannot be nil
func (s *FileListener) ExitCompilationUnit(ctx *parser.CompilationUnitContext) {
	s.Pop(s, s)
}
*/

func (s *FileListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	s.inImport = true
}

func (s *FileListener) ExitImportDeclaration(ctx *parser.ImportDeclarationContext) {
	s.inImport = false
}

func (s *FileListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.Push(s, NewClassListener(s.StackListener, s.File))
}
