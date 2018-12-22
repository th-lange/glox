package visitor

import (
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/th-lange/glox/scanner"

	"github.com/th-lange/glox/expression"
)

var prettyPrinter PrettyPrinter

var firstSetBinary expression.Binary
var firstExpectedBinary string
var secondSetBinary expression.Binary
var secondExpectedBinary string

func setUpBinary() {
	prettyPrinter = PrettyPrinter{}

	tokenPlus := scanner.Token{Lexeme: "+", Type: scanner.PLUS}
	tokenDiv := scanner.Token{Lexeme: "/", Type: scanner.SLASH}
	token123 := scanner.Token{Lexeme: "123", Type: scanner.NUMBER}
	token321 := scanner.Token{Lexeme: "321", Type: scanner.NUMBER}
	token1 := scanner.Token{Lexeme: "1", Type: scanner.NUMBER}
	token2 := scanner.Token{Lexeme: "2", Type: scanner.NUMBER}

	// First Set

	firstFirst := expression.Literal{token123}
	firstSecond := expression.Literal{token321}
	firstSetBinary = expression.Binary{firstFirst, tokenPlus, firstSecond}
	firstExpectedBinary = " ( + 123  321  ) "

	// Second Set

	grouping := expression.Grouping{
		expression.Binary{
			Left:     expression.Literal{Value: token1},
			Operator: tokenPlus,
			Right:    expression.Literal{Value: token1},
		},
	}

	secondSecond := expression.Literal{
		Value: token2,
	}
	secondSetBinary = expression.Binary{
		Left:     grouping,
		Operator: tokenDiv,
		Right:    secondSecond,
	}
	secondExpectedBinary = " ( /  ( group  ( + 1  1  )   )   2  ) "
}

func setUpGrouping() {

}

func TestPrettyPrinter_VisitBinary(t *testing.T) {
	setUpBinary()
	result := prettyPrinter.VisitBinary(firstSetBinary)
	assert.Equal(t, firstExpectedBinary, result, "Expecting results to match")

	result = prettyPrinter.VisitBinary(secondSetBinary)
	assert.Equal(t, secondExpectedBinary, result, "Expecting results to match")

}

func TestPrettyPrinter_VisitGrouping(t *testing.T) {
	prettyPrinter = PrettyPrinter{}

	tokenPlus := scanner.Token{Lexeme: "+", Type: scanner.PLUS}
	token1 := scanner.Token{Lexeme: "1", Type: scanner.NUMBER}
	token2 := scanner.Token{Lexeme: "2", Type: scanner.NUMBER}

	firstSetGrouping := expression.Grouping{
		expression.Binary{
			Left:     expression.Literal{Value: token1},
			Operator: tokenPlus,
			Right:    expression.Literal{Value: token2},
		},
	}

	firstExpectedGrouping := " ( group  ( + 1  2  )   ) "
	result := prettyPrinter.VisitGrouping(firstSetGrouping)
	assert.Equal(t, firstExpectedGrouping, result, "Expecting results to match")
}

func TestPrettyPrinter_VisitLiteral(t *testing.T) {
	tokenPlus := scanner.Token{Lexeme: "+", Type: scanner.PLUS}
	token1 := scanner.Token{Lexeme: "1", Type: scanner.NUMBER}

	result := prettyPrinter.VisitLiteral(expression.Literal{tokenPlus})
	assert.Equal(t, "+", result, "Expecting correct value for Literal.")

	result = prettyPrinter.VisitLiteral(expression.Literal{token1})
	assert.Equal(t, "1", result, "Expecting correct value for Literal.")
}

func TestPrettyPrinter_VisitUnary(t *testing.T) {
	tokenPlus := scanner.Token{Lexeme: "+", Type: scanner.PLUS}
	token1 := scanner.Token{Lexeme: "1", Type: scanner.NUMBER}

	unary := expression.Unary{
		tokenPlus,
		expression.Literal{token1},
	}

	result := prettyPrinter.VisitUnary(unary)
	assert.Equal(t, " ( + 1  ) ", result, "Expecting correct value for Literal.")
}
