package visitor

import (
	"strings"

	"github.com/th-lange/glox/expression"
)

type PrettyPrinter struct{}

func (visitor PrettyPrinter) VisitBinary(expression expression.Binary) interface{} {
	return visitor.parenthesize(expression.Operator.ValueString(), expression.Left, expression.Right)
}

func (visitor PrettyPrinter) VisitGrouping(expression expression.Grouping) interface{} {
	return visitor.parenthesize("group", expression.Expr)
}

func (visitor PrettyPrinter) VisitLiteral(expression expression.Literal) interface{} {
	return expression.Value.ValueString()
}

func (visitor PrettyPrinter) VisitUnary(expression expression.Unary) interface{} {
	return visitor.parenthesize(expression.Operator.Lexeme, expression.Right)
}

func (visitor PrettyPrinter) parenthesize(name string, expression ...expression.Expression) string {
	sb := strings.Builder{}

	sb.WriteString(" ( " + name)
	for _, itm := range expression {
		sb.WriteString(" " + itm.Accept(visitor).(string) + " ")
	}
	sb.WriteString(" ) ")
	return sb.String()
}
