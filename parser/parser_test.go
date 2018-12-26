package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/th-lange/glox/scanner"
)

//func getNewParser()

func TestNewParser(t *testing.T) {
	tkns := make([]scanner.Token, 5, 5)

	res := NewParser(&tkns)
	assert.NotNil(t, res, "Expecting a new parser to be returned.")
	assert.Equal(t, 0, res.head, "Expecting head to be initialized to 0")
	assert.Equal(t, 4, res.length, "Expecting length to be set correctly to x - 1 (as we index with zero).")
}

func TestParser_IsAtEnd(t *testing.T) {
	// Case: Length = 0
	tkns := make([]scanner.Token, 0, 0)
	prs := NewParser(&tkns)
	result := prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true in a zero length tokenSlice")

	// Case: Length == head
	tkns = make([]scanner.Token, 5, 5)
	prs = NewParser(&tkns)
	result = prs.isAtEnd()
	assert.False(t, result, "Expecting IsAtEnd to return false in a parser of size 5 with head at 0 ")
	prs.head = 5
	result = prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true in a parser of size 5 with head at 5 ")

	// Case: head > length
	tkns = make([]scanner.Token, 5, 5)
	prs = NewParser(&tkns)
	prs.head = 12
	result = prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true if parser.head is higher then it's length ")

	// Case tokens is nil
	prs = NewParser(nil)
	result = prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true if parser.head is higher then it's length ")

}

func TestParser_Advance(t *testing.T) {
	tkns := make([]scanner.Token, 5, 5)
	prs := NewParser(&tkns)

	assert.Equal(t, 0, prs.head, "Expecting parser to be initialized with 0")
	prs.advance()
	prs.advance()
	prs.advance()
	assert.Equal(t, 3, prs.head, "Expecting a triple advance to leave the parser head at 3")
}

func TestParser_Current(t *testing.T) {
	expected := []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
		{Line: 10, Type: scanner.LEFT_BRACE, Lexeme: "{"},
		{Line: 10, Type: scanner.RIGHT_BRACE, Lexeme: "}"},
	}
	prs := NewParser(&expected)

	for _, exp := range expected {
		assert.Equal(t, exp, prs.current())
		prs.advance()
	}
}

func TestParser_Consume(t *testing.T) {
	expected := []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
		{Line: 10, Type: scanner.LEFT_BRACE, Lexeme: "{"},
		{Line: 10, Type: scanner.RIGHT_BRACE, Lexeme: "{"},
		{Line: 10, Type: scanner.RIGHT_BRACE, Lexeme: "{"},
		{Line: 10, Type: scanner.RIGHT_BRACE, Lexeme: "{"},
	}
	prs := NewParser(&expected)

	err := prs.consume(scanner.LEFT_PAREN)
	assert.NoError(t, err, "Expecting correct consumption of LEFT_PAREN")
	assert.Equal(t, 1, prs.head, "Expecting HEAD of parser to be set to 1 after consumption")
	err = prs.consume(scanner.RIGHT_PAREN)
	assert.NoError(t, err, "Expecting correct consumption of RIGHT_PAREN")
	assert.Equal(t, 2, prs.head, "Expecting HEAD of parser to be set to 1 after consumption")
	err = prs.consume(scanner.LEFT_BRACE)
	assert.Equal(t, 3, prs.head, "Expecting HEAD of parser to be set to 1 after consumption")
	assert.NoError(t, err, "Expecting correct consumption of LEFT_BRACE")

	err = prs.consume(scanner.EQUAL)
	assert.Error(t, err, "Expecting error when consume is called with a TokenType that does not follow.")

}

func TestParser_Previous(t *testing.T) {
	expected := []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
		{Line: 10, Type: scanner.LEFT_BRACE, Lexeme: "{"},
		{Line: 10, Type: scanner.RIGHT_BRACE, Lexeme: "}"},
	}
	prs := NewParser(&expected)

	for i := 1; i < 6; i += 1 {
		prs.advance()
		assert.Equal(t, expected[i-1], prs.previous())
	}

}

func TestParser_String(t *testing.T) {
}
func TestParser_Parse(t *testing.T) {
}
func TestParser_Expression(t *testing.T) {
}
func TestParser_Equality(t *testing.T) {
}
func TestParser_Comparison(t *testing.T) {
}
func TestParser_Addition(t *testing.T) {
}
func TestParser_Multiplication(t *testing.T) {
}
func TestParser_Unary(t *testing.T) {
}
func TestParser_Primary(t *testing.T) {
}
func TestParser_AdvanceOnTokenTypeMatch(t *testing.T) {
}
func TestParser_Check(t *testing.T) {
}
func TestParser_Synchronize(t *testing.T) {
}
