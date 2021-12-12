package listen

import "github.com/dragonfax/delver_converter/parser"

type FileListener struct {
	*StackableListener
	*parser.BaseJavaParserListener

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

func (s *FileListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	s.File.PackageName = ctx.QualifiedName().GetText()
}

/* File shouldn't pop itself, its the root element of the stack. the parser cannot be nil
func (s *FileListener) ExitCompilationUnit(ctx *parser.CompilationUnitContext) {
	s.Pop(s, s)
}
*/

func (s *FileListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	s.File.Imports = append(s.File.Imports, ctx.QualifiedName().GetText())
}

func (s *FileListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.Push(s, NewClassListener(s.StackListener, s.File, ctx))
}
