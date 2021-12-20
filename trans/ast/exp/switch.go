package exp

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/dragonfax/java_converter/input/parser"
)

type Switch struct {
	Condition ExpressionNode
	Cases     []*SwitchCase
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
	Labels     []ExpressionNode
	Statements []ExpressionNode
}

func NewSwitch(ctx *parser.StatementContext) *Switch {
	s := &Switch{Cases: make([]*SwitchCase, 0)}

	s.Condition = ExpressionProcessor(ctx.ParExpression().Expression())

	for _, group := range ctx.AllSwitchBlockStatementGroup() {
		c := &SwitchCase{}
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

func switchLabelsFromContext(ctx *parser.SwitchLabelContext) []ExpressionNode {
	if ctx.DEFAULT() != nil {
		return []ExpressionNode{NewIdentifierNode("default")}
	}

	if ctx.GetEnumConstantName() != nil {
		return []ExpressionNode{NewEnumRef(ctx.GetEnumConstantName().GetText())}
	}

	if ctx.GetConstantExpression() != nil {
		return []ExpressionNode{ExpressionProcessor(ctx.GetConstantExpression())}
	}
	return nil
}
