package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/th-lange/glox/expression"
	"github.com/th-lange/glox/scanner"
)

func TestRPNPrinter_VisitBinary(t *testing.T) {

	rpnPrinter := RPNPrinter{}

	tokenPlus := scanner.Token{Lexeme: "+", Type: scanner.PLUS}
	tokenMinus := scanner.Token{Lexeme: "-", Type: scanner.MINUS}
	tokenMul := scanner.Token{Lexeme: "*", Type: scanner.STAR}
	token1 := scanner.Token{Lexeme: "1", Type: scanner.NUMBER}
	token2 := scanner.Token{Lexeme: "2", Type: scanner.NUMBER}
	token3 := scanner.Token{Lexeme: "3", Type: scanner.NUMBER}
	token4 := scanner.Token{Lexeme: "4", Type: scanner.NUMBER}

	firstBin := expression.Binary{
		expression.Literal{token1},
		tokenPlus,
		expression.Literal{token2},
	}
	secondBin := expression.Binary{
		expression.Literal{token4},
		tokenMinus,
		expression.Literal{token3},
	}

	binary := expression.Binary{
		Left:     firstBin,
		Operator: tokenMul,
		Right:    secondBin,
	}

	result := rpnPrinter.VisitBinary(binary)
	assert.Equal(t, "  1  2 +   4  3 - *", result, "Expecting correct RPNotation from printer")

}
