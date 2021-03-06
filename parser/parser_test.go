package parser

import (
	"testing"

	"github.com/th-lange/glox/expression"

	"github.com/stretchr/testify/assert"
	"github.com/th-lange/glox/scanner"
)

type BinaryTest struct {
	Operator []scanner.Token
}

func getParserResult() ([]scanner.Token, expression.Binary) {
	input := []scanner.Token{
		// Grouping
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},

		// Comparison
		{Line: 10, Type: scanner.EQUAL_EQUAL, Lexeme: "=="},

		// Grouping
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		// Comparison
		{Line: 10, Type: scanner.NUMBER, Lexeme: "1"},
		{Line: 10, Type: scanner.BANG_EQUAL, Lexeme: "!="},
		{Line: 10, Type: scanner.NUMBER, Lexeme: "2"},
		// Grouping End
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
	}

	expected := expression.Binary{
		Left: expression.Grouping{
			Expr: expression.Literal{
				Value: input[1],
			},
		},
		Operator: input[3],
		Right: expression.Grouping{
			Expr: expression.Binary{
				Left: expression.Literal{
					Value: input[5],
				},
				Operator: input[6],
				Right: expression.Literal{
					Value: input[7],
				},
			},
		},
	}
	return input, expected
}

func TestNewParser(t *testing.T) {
	tkns := make([]scanner.Token, 5, 5)

	res := NewParser(&tkns)
	assert.NotNil(t, res, "Expecting a new parser to be returned.")
	assert.Equal(t, 0, res.head, "Expecting head to be initialized to 0")
	assert.Equal(t, 4, res.last, "Expecting last to be set correctly to x - 1 (as we index with zero).")
}

func TestParser_IsAtEnd(t *testing.T) {
	// Case: Length = 0
	tkns := make([]scanner.Token, 0, 0)
	prs := NewParser(&tkns)
	result := prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true in a zero last tokenSlice")

	// Case: Length == head
	tkns = make([]scanner.Token, 5, 5)
	prs = NewParser(&tkns)
	result = prs.isAtEnd()
	assert.False(t, result, "Expecting IsAtEnd to return false in a parser of size 5 with head at 0 ")
	prs.head = 5
	result = prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true in a parser of size 5 with head at 5 ")

	// Case: head > last
	tkns = make([]scanner.Token, 5, 5)
	prs = NewParser(&tkns)
	prs.head = 12
	result = prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true if parser.head is higher then it's last ")

	// Case tokens is nil
	prs = NewParser(nil)
	result = prs.isAtEnd()
	assert.True(t, result, "Expecting IsAtEnd to return true if parser.head is higher then it's last ")

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
	input := []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
	}
	prs := NewParser(&input)

	expected := "Parser: Head: 0, Last: 1, Current Token: Token: { Line: 10, Type: LEFT_PAREN, Lexeme: \"(\", Literal: <nil>}"

	assert.Equal(t, expected, prs.String(), "Expecting the correct string output for the parsers String() method.")

	prs.advance()
	expected = "Parser: Head: 1, Last: 1, Current Token: Token: { Line: 10, Type: RIGHT_PAREN, Lexeme: \")\", Literal: <nil>}"
	assert.Equal(t, expected, prs.String(), "Expecting the correct string output for the parsers String() method.")

}

func TestParser_Primary_Simple(t *testing.T) {

	input := []scanner.Token{
		{Line: 10, Type: scanner.FALSE, Lexeme: "false"},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.NIL, Lexeme: "nil"},
		{Line: 10, Type: scanner.STRING, Lexeme: "some string"},
		{Line: 10, Type: scanner.NUMBER, Lexeme: "123.123"},
	}
	prs := NewParser(&input)

	expected := expression.Literal{}
	for _, inp := range input {
		result := prs.primary()
		expected.Value = inp
		assert.Equal(t, expected, result, "Expecting the correct type and Value for parsers primary() Method")
	}

}

func TestParser_Primary_Grouping_OK(t *testing.T) {

	input := []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
	}
	prs := NewParser(&input)

	expected := expression.Grouping{
		Expr: expression.Literal{
			Value: input[1],
		},
	}

	result := prs.primary()
	assert.Equal(t, expected, result, "Expecting a correctly grouped output.")
}

func TestParser_Unary_BANG(t *testing.T) {
	input := []scanner.Token{
		{Line: 10, Type: scanner.BANG, Lexeme: "!"},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
	}

	expected := expression.Unary{
		Operator: input[0],
		Right: expression.Literal{
			Value: input[1],
		},
	}
	prs := NewParser(&input)

	result := prs.unary()
	assert.Equal(t, expected, result, "Expecting a correct unary BANG output.")
}

func TestParser_Unary_MINUS(t *testing.T) {
	input := []scanner.Token{
		{Line: 10, Type: scanner.MINUS, Lexeme: "!"},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
	}

	expected := expression.Unary{
		Operator: input[0],
		Right: expression.Literal{
			Value: input[1],
		},
	}
	prs := NewParser(&input)

	result := prs.unary()
	assert.Equal(t, expected, result, "Expecting a correct unary MINUS output.")
}

func TestParser_Multiplication(t *testing.T) {
	testCases := BinaryTest{
		[]scanner.Token{
			{Line: 10, Type: scanner.SLASH, Lexeme: "/"},
			{Line: 10, Type: scanner.STAR, Lexeme: "*"},
		},
	}

	for _, test := range testCases.Operator {
		input := []scanner.Token{
			{Line: 10, Type: scanner.NUMBER, Lexeme: "4"},
			test,
			{Line: 10, Type: scanner.NUMBER, Lexeme: "2"},
		}

		expected := expression.Binary{
			Left: expression.Literal{
				Value: input[0],
			},
			Operator: input[1],
			Right: expression.Literal{
				Value: input[2],
			},
		}
		prs := NewParser(&input)

		result := prs.multiplication()
		assert.Equal(t, expected, result, "Expecting a correct output for: "+test.Lexeme)
	}
}

func TestParser_Addition(t *testing.T) {
	testCases := BinaryTest{
		[]scanner.Token{
			{Line: 10, Type: scanner.MINUS, Lexeme: "-"},
			{Line: 10, Type: scanner.PLUS, Lexeme: "+"},
		},
	}

	for _, test := range testCases.Operator {
		input := []scanner.Token{
			{Line: 10, Type: scanner.NUMBER, Lexeme: "1"},
			test,
			{Line: 10, Type: scanner.NUMBER, Lexeme: "2"},
		}

		expected := expression.Binary{
			Left: expression.Literal{
				Value: input[0],
			},
			Operator: input[1],
			Right: expression.Literal{
				Value: input[2],
			},
		}
		prs := NewParser(&input)

		result := prs.addition()
		assert.Equal(t, expected, result, "Expecting a correct output for: "+test.Lexeme)
	}
}

func TestParser_Comparison(t *testing.T) {

	testCases := BinaryTest{
		[]scanner.Token{
			{Line: 10, Type: scanner.GREATER, Lexeme: ">"},
			{Line: 10, Type: scanner.GREATER_EQUAL, Lexeme: ">="},
			{Line: 10, Type: scanner.LESS, Lexeme: "<"},
			{Line: 10, Type: scanner.LESS_EQUAL, Lexeme: "<="},
		},
	}

	for _, test := range testCases.Operator {
		input := []scanner.Token{
			{Line: 10, Type: scanner.NUMBER, Lexeme: "1"},
			test,
			{Line: 10, Type: scanner.NUMBER, Lexeme: "2"},
		}

		expected := expression.Binary{
			Left: expression.Literal{
				Value: input[0],
			},
			Operator: input[1],
			Right: expression.Literal{
				Value: input[2],
			},
		}
		prs := NewParser(&input)

		result := prs.comparison()
		assert.Equal(t, expected, result, "Expecting a correct output for"+test.Lexeme)
	}
}

func TestParser_Equality(t *testing.T) {
	testCases := BinaryTest{
		[]scanner.Token{
			{Line: 10, Type: scanner.BANG_EQUAL, Lexeme: "!="},
			{Line: 10, Type: scanner.EQUAL_EQUAL, Lexeme: "=="},
		},
	}

	for _, test := range testCases.Operator {
		input := []scanner.Token{
			{Line: 10, Type: scanner.NUMBER, Lexeme: "1"},
			test,
			{Line: 10, Type: scanner.NUMBER, Lexeme: "1"},
		}

		expected := expression.Binary{
			Left: expression.Literal{
				Value: input[0],
			},
			Operator: input[1],
			Right: expression.Literal{
				Value: input[2],
			},
		}
		prs := NewParser(&input)

		result := prs.equality()
		assert.Equal(t, expected, result, "Expecting a correct output for"+test.Lexeme)
	}
}

func TestParser_Expression(t *testing.T) {
	input, expected := getParserResult()
	prs := NewParser(&input)
	result := prs.expression()
	assert.Equal(t, expected, result, "Expecting a correct output for parsing complex expressions.")
}

func TestParser_AdvanceOnTokenTypeMatch(t *testing.T) {
	input := []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
	}
	prs := NewParser(&input)

	assert.Equal(t, 0, prs.head, "Expecting initial setting of prs.head to 0")
	assert.Equal(t, 2, prs.last, "Expecting correct setting (and non-modification) of the prs.tail value")

	result := prs.advanceOnTokenTypeMatch(scanner.EQUAL_EQUAL, scanner.EQUAL, scanner.BANG_EQUAL)

	assert.Equal(t, 0, prs.head, "Expecting correct adjustment if token is found")
	assert.Equal(t, 2, prs.last, "Expecting correct setting (and non-modification) of the prs.tail value")
	assert.False(t, result, "Expecting false if token is not present")

	result = prs.advanceOnTokenTypeMatch(scanner.EQUAL_EQUAL, scanner.LEFT_PAREN, scanner.TRUE)

	assert.Equal(t, 1, prs.head, "Expecting correct adjustment if token is found")
	assert.Equal(t, 2, prs.last, "Expecting correct setting (and non-modification) of the prs.tail value")
	assert.True(t, result, "Expecting true if token is present")

	prs.advanceOnTokenTypeMatch(scanner.EQUAL_EQUAL, scanner.TRUE)

	assert.Equal(t, 2, prs.head, "Expecting correct adjustment if token is found")
	assert.Equal(t, 2, prs.last, "Expecting correct setting (and non-modification) of the prs.tail value")
	assert.True(t, result, "Expecting true if token is present")

	prs.advanceOnTokenTypeMatch(scanner.RIGHT_PAREN, scanner.TRUE)

	assert.Equal(t, 3, prs.head, "Expecting correct adjustment if token is found")
	assert.Equal(t, 2, prs.last, "Expecting correct setting (and non-modification) of the prs.tail value")
	assert.True(t, result, "Expecting true if token is present")
}

func TestParser_Check(t *testing.T) {
	input := []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
	}
	prs := NewParser(&input)

	result := prs.check(scanner.RIGHT_PAREN)
	assert.False(t, result, "Expecting false if elements do not match")

	result = prs.check(scanner.LEFT_PAREN)
	assert.True(t, result, "Expecting true if elements do match")

}

func TestParser_Synchronize(t *testing.T) {
	// Test Case 1: Empty parser
	input := []scanner.Token{}
	prs := NewParser(&input)

	prs.synchronize()

	// Test Case 2: Parser with no break conditions
	input = []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
	}
	prs = NewParser(&input)

	prs.synchronize()
	assert.Equal(t, 3, prs.head, "Expecting synchronize() to walk to the end if no break conditions are found: CLASS, FUN, VAR, FOR, IF, WHILE, PRINT, RETURN")

	// Test Case 3: Last element is a break condition
	input = []scanner.Token{
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
		{Line: 10, Type: scanner.FUN, Lexeme: "fun"},
	}
	prs = NewParser(&input)
	prs.synchronize()
	assert.Equal(t, 4, prs.head, "Expecting synchronize() to walk to the end if no break conditions are found: CLASS, FUN, VAR, FOR, IF, WHILE, PRINT, RETURN")

	// Test Case 4: Next element is a break condition
	input = []scanner.Token{
		{Line: 10, Type: scanner.FUN, Lexeme: "fun"},
		{Line: 10, Type: scanner.LEFT_PAREN, Lexeme: "("},
		{Line: 10, Type: scanner.TRUE, Lexeme: "true"},
		{Line: 10, Type: scanner.RIGHT_PAREN, Lexeme: ")"},
	}
	prs = NewParser(&input)
	prs.synchronize()
	assert.Equal(t, 1, prs.head, "Expecting synchronize() to walk to the end if no break conditions are found: CLASS, FUN, VAR, FOR, IF, WHILE, PRINT, RETURN")
}

func TestParser_Parse(t *testing.T) {
	input, expected := getParserResult()
	prs := NewParser(&input)
	result := prs.Parse()
	assert.Equal(t, expected, result, "Expecting a correct output for parsing complex expressions.")
}
