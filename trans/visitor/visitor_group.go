package visitor

import "github.com/dragonfax/java-go-transpiler/trans/ast"

func RunGroup(h *ast.Hierarchy) {
	ParentPass(h)
	NewScopePass(h).VisitNode(h)
	NewPrimitiveClassPass(h).VisitNode(h)
	NewPopulatePass(h).VisitNode(h)
	NewClassResolver(h).VisitNode(h)
	NewFetchTypesPass(h).VisitNode(h)
	NewVarResolver(h).VisitNode(h)
	BaseClassPass(h)

}
