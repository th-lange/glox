package parser

import (
	"fmt"

	"github.com/th-lange/glox/expression"
	"github.com/th-lange/glox/scanner"
)

type parser struct {
	tokens *[]scanner.Token
	length int
	head   int
	errors []error
}

func NewParser(tokens *[]scanner.Token) *parser {
	prs := new(parser)
	prs.tokens = tokens
	prs.length = len(*tokens)
	prs.head = 0
	return prs
}

func (prs parser) String() string {
	return fmt.Sprintf("Parser: Head: %d, Length: %d, Current Token: %s", prs.head, prs.length, (*prs.tokens)[prs.head].String())
}

// expression     → equality ;
func (prs *parser) expression() expression.Expression {
	return prs.equality()
}

// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
func (prs *parser) equality() expression.Expression {
	expr := prs.comparison()
	for prs.advanceOnTokenTypeMatch(scanner.BANG_EQUAL, scanner.EQUAL_EQUAL) {
		operator := prs.previous()
		right := prs.comparison()
		expr = expression.Binary{expr, operator, right}
	}
	return expr
}

// comparison     → addition ( ( ">" | ">=" | "<" | "<=" ) addition )* ;
func (prs *parser) comparison() expression.Expression {
	expr := prs.addition()
	for prs.advanceOnTokenTypeMatch(scanner.GREATER, scanner.GREATER_EQUAL, scanner.LESS, scanner.LESS_EQUAL) {
		operator := prs.previous()
		right := prs.addition()
		expr = expression.Binary{expr, operator, right}
	}
	return expr
}

// addition       → multiplication ( ( "-" | "+" ) multiplication )* ;
func (prs *parser) addition() expression.Expression {
	expr := prs.multiplication()
	for prs.advanceOnTokenTypeMatch(scanner.MINUS, scanner.PLUS) {
		operator := prs.previous()
		right := prs.multiplication()
		expr = expression.Binary{expr, operator, right}
	}
	return expr
}

// multiplication → unary ( ( "/" | "*" ) unary )* ;
func (prs *parser) multiplication() expression.Expression {
	expr := prs.unary()
	for prs.advanceOnTokenTypeMatch(scanner.SLASH, scanner.STAR) {
		operator := prs.previous()
		right := prs.unary()
		expr = expression.Binary{expr, operator, right}
	}
	return expr
}

// unary          → ( "!" | "-" ) unary    |    primary ;
func (prs *parser) unary() expression.Expression {
	if prs.advanceOnTokenTypeMatch(scanner.BANG, scanner.MINUS) {
		operator := prs.previous()
		right := prs.unary()
		return expression.Unary{operator, right}
	}
	return prs.primary()
}

//primary        → NUMBER | STRING | "false" | "true" | "nil"   |    "(" expression ")" ;
func (prs *parser) primary() expression.Expression {
	if prs.advanceOnTokenTypeMatch(scanner.FALSE, scanner.TRUE, scanner.NIL, scanner.STRING, scanner.NUMBER) {
		return expression.Literal{prs.previous()}
	}
	if prs.advanceOnTokenTypeMatch(scanner.LEFT_PAREN) {
		expr := prs.expression()
		err := prs.consumeUntil(scanner.RIGHT_PAREN)
		if err != nil {
			// We should do more here!
			prs.errors = append(prs.errors, err)
		}
		return expression.Grouping{expr}
	}
	panic("Fond end of Grammar in parser.primary. \nStats: " + prs.String())

}

func (prs *parser) advanceOnTokenTypeMatch(tokenTypes ...scanner.TokenType) bool {
	for _, itm := range tokenTypes {
		if prs.check(itm) {
			prs.advance()
			return true
		}
	}
	return false
}

func (prs *parser) check(tokenType scanner.TokenType) bool {
	if prs.isAtEnd() {
		return false
	}
	return prs.current().Type == tokenType

}

func (prs *parser) isAtEnd() bool {
	return prs.head > prs.length

}

func (prs *parser) advance() {
	prs.head += 1
}

func (prs *parser) current() scanner.Token {
	if prs.head < 0 && prs.head > prs.length {
		panic("Cound not return prev element as HEAD is at 0")
	}
	return (*prs.tokens)[prs.head]
}

func (prs *parser) consumeUntil(tokenType scanner.TokenType) error {
	for prs.current().Type != tokenType {
		if prs.isAtEnd() {
			return ParserError{
				(*prs.tokens)[prs.head],
				prs.head,
				"Found EOF before expected Token",
			}
		}
		prs.advance()
	}
	return nil
}

func (prs *parser) previous() scanner.Token {
	if prs.head == 0 {
		panic("Cound not return prev element as HEAD is at 0")
	}
	return (*prs.tokens)[prs.head-1]
}
