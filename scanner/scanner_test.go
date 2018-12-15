package scanner

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestCode = `/*
Foooo bar

fffoooo bar bar / test

*/


// this is a comment
(( )){} // grouping stuff
!*+-/=<> <= == // operators
123
1225
12.356
identifier = "Fooo"
class StrangeName {
    var first = 123
}

iden_ti_fier = 123
0123.1223
"This is a very long string
that spans multiple lines

KK"
and  and_and
class  class_class
else  else_else
false  false_false
for  for_for

/*
Unterminated multi line comment
`

var expected = []Token{
	{Type: LEFT_PAREN, Line: 10, Lexeme: "("},
	{Line: 10, Type: LEFT_PAREN, Lexeme: "("},
	{Line: 10, Type: RIGHT_PAREN, Lexeme: ")"},
	{Line: 10, Type: RIGHT_PAREN, Lexeme: ")"},
	{Line: 10, Type: LEFT_BRACE, Lexeme: "{"},
	{Line: 10, Type: RIGHT_BRACE, Lexeme: "}"},
	{Line: 11, Type: BANG, Lexeme: "!"},
	{Line: 11, Type: STAR, Lexeme: "*"},
	{Line: 11, Type: PLUS, Lexeme: "+"},
	{Line: 11, Type: MINUS, Lexeme: "-"},
	{Line: 11, Type: SLASH, Lexeme: "/"},
	{Line: 11, Type: EQUAL, Lexeme: "="},
	{Line: 11, Type: LESS, Lexeme: "<"},
	{Line: 11, Type: GREATER, Lexeme: ">"},
	{Line: 11, Type: LESS_EQUAL, Lexeme: "<="},
	{Line: 11, Type: EQUAL_EQUAL, Lexeme: "=="},
	{Line: 12, Type: NUMBER, Lexeme: "123", Literal: 123},
	{Line: 12, Type: NUMBER, Lexeme: "1225", Literal: 1225},
	{Line: 12, Type: NUMBER, Lexeme: "12.356", Literal: 12.356},
	{Line: 12, Type: IDENTIFIER, Lexeme: "identifier"},
	{Line: 12, Type: EQUAL, Lexeme: "="},
	{Line: 12, Type: STRING, Lexeme: "Fooo", Literal: "Fooo"},
	{Line: 13, Type: CLASS, Lexeme: "class"},
	{Line: 13, Type: IDENTIFIER, Lexeme: "StrangeName"},
	{Line: 13, Type: LEFT_BRACE, Lexeme: "{"},
	{Line: 14, Type: VAR, Lexeme: "var"},
	{Line: 14, Type: IDENTIFIER, Lexeme: "first"},
	{Line: 14, Type: EQUAL, Lexeme: "="},
	{Line: 14, Type: NUMBER, Lexeme: "123", Literal: 123},
	{Line: 14, Type: RIGHT_BRACE, Lexeme: "}"},
	{Line: 16, Type: IDENTIFIER, Lexeme: "iden_ti_fier"},
	{Line: 16, Type: EQUAL, Lexeme: "="},
	{Line: 16, Type: NUMBER, Lexeme: "123", Literal: 123},
	{Line: 16, Type: NUMBER, Lexeme: "0123.1223", Literal: 123.1223},
	{Line: 16, Type: STRING, Lexeme: `This is a very long string
that spans multiple lines

KK`},
	{Line: 20, Type: AND, Lexeme: "and"},
	{Line: 20, Type: IDENTIFIER, Lexeme: "and_and"},
	{Line: 20, Type: CLASS, Lexeme: "class"},
	{Line: 20, Type: IDENTIFIER, Lexeme: "class_class"},
	{Line: 20, Type: ELSE, Lexeme: "else"},
	{Line: 20, Type: IDENTIFIER, Lexeme: "else_else"},
	{Line: 20, Type: FALSE, Lexeme: "false"},
	{Line: 20, Type: IDENTIFIER, Lexeme: "false_false"},
	{Line: 20, Type: FOR, Lexeme: "for"},
	{Line: 20, Type: IDENTIFIER, Lexeme: "for_for"},
	{Line: 22, Type: EOF, Lexeme: "EOF"},
}

func TestScanner_Scan(t *testing.T) {
	type testStruct struct {
		name         string
		input        string
		expectedType TokenType
	}

	tts := GetTokenTypes()

	for _, tt := range tts {
		name := "Testing: " + tt.String + " Expecting correct parsing of element."
		t.Run(name, func(t *testing.T) {
			scnr := Scanner{}
			scnr.Scan(tt.String)
			assert.Equal(t, scnr.Tokens[0].Type, tt.TokenTypes[0], "TokenType: "+name)
			assert.Equal(t, scnr.Tokens[0].Length, tt.Lengths[0], "Length: "+name)
			assert.Equal(t, scnr.Tokens[0].Lexeme, tt.Lexemes[0], "Literal: "+name)

			assert.Equal(t, scnr.Tokens[1].Type, EOF, "Expecting the last Token to be EOF")
		})
	}
}

func TestScanner_Scan2(t *testing.T) {
	scnr := Scanner{}
	scnr.Scan(TestCode)

	for i, itm := range expected {
		assert.Equal(t, itm.Line, scnr.Tokens[i].Line)
		assert.Equal(t, itm.Type, scnr.Tokens[i].Type)
		if scnr.Tokens[i].Lexeme != "" {
			assert.Equal(t, itm.Lexeme, scnr.Tokens[i].Lexeme)
		}

	}

}

func TestScanner_appendError(t *testing.T) {
	scnr := Scanner{}

	err1 := errors.New("Error 1")
	err2 := errors.New("Error 2")

	scnr.appendError(err1)
	scnr.appendError(err2)

	assert.Equal(t, err1, scnr.Errors[0], "Expecting correct first error: "+err1.Error())
	assert.Equal(t, err2, scnr.Errors[1], "Expecting correct second error: "+err2.Error())

	assert.True(t, scnr.HadError, "Expecting HadErrors to be set")
}

func TestScanner_appendToken(t *testing.T) {
	tkn1 := Token{Line: 1, Lexeme: "(", Type: LEFT_PAREN}
	tkn2 := Token{Line: 1, Lexeme: ")", Type: RIGHT_PAREN}
	tkn3 := Token{Line: 2, Lexeme: "{", Type: LEFT_BRACE}
	tkn4 := Token{Line: 2, Lexeme: "}", Type: RIGHT_BRACE}

	type args struct {
		tkn Token
	}
	tests := []struct {
		name     string
		args     args
		expected []Token
	}{
		{name: "Adding first Token: (. Expecting Tokens: (", args: args{tkn: tkn1}, expected: []Token{tkn1}},
		{name: "Adding second Token: ). Expecting Tokens: ()", args: args{tkn: tkn2}, expected: []Token{tkn1, tkn2}},
		{name: "Adding third Token: {. Expecting Tokens: (){", args: args{tkn: tkn3}, expected: []Token{tkn1, tkn2, tkn3}},
		{name: "Adding fourth Token: }. Expecting Tokens: (){}", args: args{tkn: tkn4}, expected: []Token{tkn1, tkn2, tkn3, tkn4}},
	}

	scnr := Scanner{}
	for _, tt := range tests {
		scnr.appendToken(tt.args.tkn)
		assert.Equal(t, tt.expected, scnr.Tokens, tt.name)
	}
}

func TestScanner_appendEOFToken(t *testing.T) {
	scnr := Scanner{}
	scnr.Scan("")

	expected := Token{
		Position: scnr.current,
		Lexeme:   "EOF",
		Length:   0,
		Line:     scnr.Line,
		Type:     EOF,
	}
	assert.Contains(t, scnr.Tokens, expected, "Expects, that an EOF Token is appended.")
}

func TestScanner_consume(t *testing.T) {
	line := "This it \" a test /* line*/"
	scnr := Scanner{
		lines:   line,
		Line:    1,
		current: 0,
		length:  len(line),
	}

	scnr.consume('"')
	assert.Equal(t, 8, scnr.current, "Expecting Scanner Consume to consume all characters up to the \"")

	scnr.consume('*')
	assert.Equal(t, 18, scnr.current, "Expecting Scanner Consume to consume all characters up to the *")

	scnr.consume('/')
	assert.Equal(t, 25, scnr.current, "Expecting Scanner Consume to consume all characters up to the /")
}

func TestScanner_consumeUntilTwo(t *testing.T) {
	line := "This it // a test /* line*/"
	scnr := Scanner{
		lines:   line,
		Line:    1,
		current: 0,
		length:  len(line),
	}

	scnr.consumeUntilTwo('/', '/')
	assert.Equal(t, 10, scnr.current, "Expecting Scanner Consume to consume all characters up to the //")

	scnr.consumeUntilTwo('/', '*')
	assert.Equal(t, 20, scnr.current, "Expecting Scanner Consume to consume all characters up to the /*")
}

func TestScanner_string_ok(t *testing.T) {
	line := "variable = \"This is a super string\""
	scnr := Scanner{
		lines:   line,
		Line:    1,
		current: 11,
		length:  len(line),
	}
	tkn := Token{
		Position: scnr.current,
		Line:     scnr.Line,
		Lexeme:   "\"",
		Length:   1,
	}

	err := scnr.string(&tkn)
	assert.NoError(t, err, "Expecting no error after parsing strings.")
	assert.Equal(t, STRING, tkn.Type, "Expected TokenType to be string after string parsing.")
	assert.Equal(t, "This is a super string", tkn.Lexeme, "Expecting the correct string to be parsed.")
}

func TestScanner_string_error(t *testing.T) {
	line := "variable = \"This is a super string"
	scnr := Scanner{
		lines:   line,
		Line:    1,
		current: 11,
		length:  len(line),
	}
	tkn := Token{
		Position: scnr.current,
		Line:     scnr.Line,
		Lexeme:   "\"",
		Length:   1,
	}

	err := scnr.string(&tkn)
	assert.Error(t, err, "Expecting no error after parsing strings.")
}

func TestScanner_number_number(t *testing.T) {
	line := "variable = 123 "
	scnr := Scanner{
		lines:   line,
		Line:    1,
		current: 11,
		length:  len(line),
	}
	tkn := Token{
		Position: scnr.current,
		Line:     scnr.Line,
		Lexeme:   "1",
		Length:   1,
	}

	err := scnr.number(&tkn)
	assert.NoError(t, err, "Expecting no error after parsing strings.")
	assert.Equal(t, NUMBER, tkn.Type, "Expected TokenType to be number after number parsing.")
	assert.Equal(t, "123", tkn.Lexeme, "Expecting the correct number to be parsed.")
}

func TestScanner_number_double(t *testing.T) {
	line := "variable = 123.123 "
	scnr := Scanner{
		lines:   line,
		Line:    1,
		current: 11,
		length:  len(line),
	}
	tkn := Token{
		Position: scnr.current,
		Line:     scnr.Line,
		Lexeme:   "1",
		Length:   1,
	}

	err := scnr.number(&tkn)
	assert.NoError(t, err, "Expecting no error after parsing strings.")
	assert.Equal(t, NUMBER, tkn.Type, "Expected TokenType to be double after number parsing.")
	assert.Equal(t, "123.123", tkn.Lexeme, "Expecting the correct double to be parsed.")
}

func TestScanner_identifier(t *testing.T) {
	line := " variable = 123.123 "
	scnr := Scanner{
		lines:   line,
		Line:    1,
		current: 1,
		length:  len(line),
	}
	tkn := Token{
		Position: scnr.current,
		Line:     scnr.Line,
		Lexeme:   "v",
		Length:   1,
	}

	scnr.identifier(&tkn)
	assert.Equal(t, IDENTIFIER, tkn.Type, "Expecting the type to be set to IDENTIFIER.")
	assert.Equal(t, "variable", tkn.Lexeme, "Expecting the type to be set to variable.")
	assert.Equal(t, 9, scnr.current, "Expecting current to be set to the correct value.")
}

func TestScanner_isAlphaNumeric(t *testing.T) {
	assert.True(t, isAlphaNumeric('1'), "Expecting 1 to be considered alphaNumeric")
	assert.True(t, isAlphaNumeric('x'), "Expecting x to be considered alphaNumeric")
	assert.True(t, isAlphaNumeric('_'), "Expecting _ to be considered alphaNumeric")
	assert.True(t, isAlphaNumeric('-'), "Expecting - to be considered alphaNumeric")

	assert.False(t, isAlphaNumeric('*'), "Expecting * not to be considered alphaNumeric")
	assert.False(t, isAlphaNumeric('/'), "Expecting / not to be considered alphaNumeric")
	assert.False(t, isAlphaNumeric('>'), "Expecting < not to be considered alphaNumeric")
	assert.False(t, isAlphaNumeric('<'), "Expecting > not to be considered alphaNumeric")
}
