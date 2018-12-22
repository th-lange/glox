# Representing Code Challenge

## Rewrite - and no sugar!
> Earlier, I said that the |, *, and + forms we added to our grammar metasyntax were just syntactic sugar. Given this grammar:
```
expr → expr ( "(" ( expr ( "," expr )* )? ")" | "." IDENTIFIER )*
       | IDENTIFIER
       | NUMBER
```

### Result

```
expression  → literal
expression  → grouping
expression  → unary
expression  → binary
expression  → arguments


literal     → NUMBER
literal     → IDENTIFIER
literal     → true
literal     → false
literal     → nil

grouping    → "(" expression ")"

unary       → "-" expression 
unary       → "!" expression

binary      → expression operator expression

operator    → "==" 
operator    → "!=" 
operator    → "<" 
operator    → "<=" 
operator    → ">" 
operator    → ">=" 
operator    → "+" 
operator    → "-" 
operator    → "*" 
operator    → "/" 

arguments   → expression "," expression

    
```

## Equivalent to Visitor Pattern in Functional Language

> The Visitor pattern lets you emulate the functional style in an object-oriented language. Devise a corresponding pattern in a functional language. **It should let you bundle all of the operations on one type together and let you define new types easily**.
>
> (SML or Haskell would be ideal for this exercise, but Scheme or another Lisp works as well.)

I do not have enough knowledge to answer this without feeling bad. So I'll go with feeling bad for not answering it.

## RPN Visitor

This is almost done by the current visitor. Only change required would be to move the "name" to the end of the string builder

```go
// The Visitor Class
package visitor

import (
	"strings"

	"github.com/th-lange/glox/expression"
)

type RPNPrinter struct{}

func (visitor RPNPrinter) VisitBinary(expression expression.Binary) interface{} {
	return visitor.parenthesize(expression.Operator.ValueString(), expression.Left, expression.Right)
}

func (visitor RPNPrinter) VisitGrouping(expression expression.Grouping) interface{} {
	return visitor.parenthesize("group", expression.Expr)
}

func (visitor RPNPrinter) VisitLiteral(expression expression.Literal) interface{} {
	return expression.Value.ValueString()
}

func (visitor RPNPrinter) VisitUnary(expression expression.Unary) interface{} {
	return visitor.parenthesize(expression.Operator.Lexeme, expression.Right)
}

func (visitor RPNPrinter) parenthesize(name string, expression ...expression.Expression) string {
	sb := strings.Builder{}

	for _, itm := range expression {
		sb.WriteString(" " + itm.Accept(visitor).(string) + " ")
	}
	sb.WriteString(name)
	return sb.String()
}

```

The corresponding test:

```go
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
```