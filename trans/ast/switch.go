package ast

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type Switch struct {
	*node.BaseNode

	Condition node.Node
	Cases     []*SwitchCase
}

func (s *Switch) Children() []node.Node {
	return node.AppendNodeLists(s.Cases, s.Condition)
}

var switchTemplate = `
switch ( {{.Condition}} ) {
{{range .Cases}}
{{if ne (index .Labels 0).String "default"}}case{{end}} {{range $index, $element := .Labels}}{{if $index}},{{end}}{{ $element }}{{end}}:
{{range .Statements}}{{ . }}
{{end}}
{{end}}
}
`

var switchTemplateCompiled = template.Must(template.New("switch").Parse(switchTemplate))

func ExecuteTemplateToString(tmpl *template.Template, data interface{}) string {
	b := strings.Builder{}
	err := tmpl.Execute(&b, data)
	if err != nil {
		// return error as the results, so it ends up in the .go file.
		fmt.Printf("WARNING: error while executing template: %s\n", err.Error())
		return err.Error()
	}
	return b.String()
}

func (s *Switch) String() string {
	return ExecuteTemplateToString(switchTemplateCompiled, s)
}

type SwitchCase struct {
	*node.BaseNode

	Labels     []node.Node
	Statements []node.Node
}

func (sc *SwitchCase) String() string {
	return "case: "
}

func (sc *SwitchCase) Children() []node.Node {
	return node.AppendNodeLists(sc.Labels, sc.Statements...)
}

func NewSwitch(ctx *parser.StatementContext) *Switch {
	s := &Switch{BaseNode: node.NewNode(), Cases: make([]*SwitchCase, 0)}

	s.Condition = ExpressionProcessor(ctx.ParExpression().Expression())

	for _, group := range ctx.AllSwitchBlockStatementGroup() {
		c := &SwitchCase{BaseNode: node.NewNode()}
		for _, labelCtx := range group.AllSwitchLabel() {
			labels := switchLabelsFromContext(labelCtx)
			c.Labels = append(c.Labels, labels...)
		}
		for _, blockCtx := range group.AllBlockStatement() {
			stmts := BlockStatementProcessor(blockCtx)
			c.Statements = append(c.Statements, stmts...)
		}
		s.Cases = append(s.Cases, c)
	}

	// empty cases at the end (with no blocks)
	for _, labelCtx := range ctx.AllSwitchLabel() {
		labels := switchLabelsFromContext(labelCtx)
		if len(labels) > 0 {
			s.Cases = append(s.Cases, &SwitchCase{
				Labels: labels,
			})
		}
	}

	return s
}

func switchLabelsFromContext(ctx *parser.SwitchLabelContext) []node.Node {
	if ctx.DEFAULT() != nil {
		return []node.Node{NewIdentifierNode("default")}
	}

	if ctx.GetEnumConstantName() != nil {
		return []node.Node{NewEnumRef(ctx.GetEnumConstantName().GetText())}
	}

	if ctx.GetConstantExpression() != nil {
		return []node.Node{ExpressionProcessor(ctx.GetConstantExpression())}
	}
	return nil
}
