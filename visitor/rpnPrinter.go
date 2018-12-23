package visitor

import (
	"strings"

	"github.com/th-lange/glox/expression"
)

type RPNPrinter struct{}

func (visitor RPNPrinter) VisitBinary(expression expression.Binary) interface{} {
	return visitor.renderAsReversePolishNotation(expression.Operator.ValueString(), expression.Left, expression.Right)
}

func (visitor RPNPrinter) VisitGrouping(expression expression.Grouping) interface{} {
	return visitor.renderAsReversePolishNotation("group", expression.Expr)
}

func (visitor RPNPrinter) VisitLiteral(expression expression.Literal) interface{} {
	return expression.Value.ValueString()
}

func (visitor RPNPrinter) VisitUnary(expression expression.Unary) interface{} {
	return visitor.renderAsReversePolishNotation(expression.Operator.Lexeme, expression.Right)
}

func (visitor RPNPrinter) renderAsReversePolishNotation(name string, expression ...expression.Expression) string {
	sb := strings.Builder{}
	for _, itm := range expression {
		sb.WriteString(" " + itm.Accept(visitor).(string) + " ")
	}
	sb.WriteString(name)
	return sb.String()
}
