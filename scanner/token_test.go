package scanner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TokenTestData struct {
	String     string
	TokenTypes []TokenType
	Lexemes    []string
	Lengths    []int
}

func GetTokenTypes() []TokenTestData {
	return []TokenTestData{
		{"(", []TokenType{LEFT_PAREN}, []string{"("}, []int{1}},
		{")", []TokenType{RIGHT_PAREN}, []string{")"}, []int{1}},
		{"{", []TokenType{LEFT_BRACE}, []string{"{"}, []int{1}},
		{"}", []TokenType{RIGHT_BRACE}, []string{"}"}, []int{1}},

		{",", []TokenType{COMMA}, []string{","}, []int{1}},
		{".", []TokenType{DOT}, []string{"."}, []int{1}},
		{"-", []TokenType{MINUS}, []string{"-"}, []int{1}},

		{"+", []TokenType{PLUS}, []string{"+"}, []int{1}},
		{";", []TokenType{SEMICOLON}, []string{";"}, []int{1}},

		{"/", []TokenType{SLASH}, []string{"/"}, []int{1}},
		{"*", []TokenType{STAR}, []string{"*"}, []int{1}},

		{"!", []TokenType{BANG}, []string{"!"}, []int{1}},
		{"!=", []TokenType{BANG_EQUAL}, []string{"!="}, []int{2}},

		{"=", []TokenType{EQUAL}, []string{"="}, []int{1}},
		{"==", []TokenType{EQUAL_EQUAL}, []string{"=="}, []int{2}},

		{">", []TokenType{GREATER}, []string{">"}, []int{1}},
		{">=", []TokenType{GREATER_EQUAL}, []string{">="}, []int{2}},

		{"<", []TokenType{LESS}, []string{"<"}, []int{1}},
		{"<=", []TokenType{LESS_EQUAL}, []string{"<="}, []int{2}},
		//
		{"fooBar", []TokenType{IDENTIFIER}, []string{"fooBar"}, []int{6}},

		{"\"This is a string\"", []TokenType{STRING}, []string{"This is a string"}, []int{16}},
		{"123", []TokenType{NUMBER}, []string{"123"}, []int{3}},

		{"and", []TokenType{AND}, []string{"and"}, []int{3}},
		{"class", []TokenType{CLASS}, []string{"class"}, []int{5}},
		{"else", []TokenType{ELSE}, []string{"else"}, []int{4}},
		{"false", []TokenType{FALSE}, []string{"false"}, []int{5}},
		{"for", []TokenType{FOR}, []string{"for"}, []int{3}},
		{"fun", []TokenType{FUN}, []string{"fun"}, []int{3}},
		{"if", []TokenType{IF}, []string{"if"}, []int{2}},
		{"nil", []TokenType{NIL}, []string{"nil"}, []int{3}},
		{"or", []TokenType{OR}, []string{"or"}, []int{2}},
		{"print", []TokenType{PRINT}, []string{"print"}, []int{5}},
		{"return", []TokenType{RETURN}, []string{"return"}, []int{6}},
		{"super", []TokenType{SUPER}, []string{"super"}, []int{5}},
		{"this", []TokenType{THIS}, []string{"this"}, []int{4}},
		{"true", []TokenType{TRUE}, []string{"true"}, []int{4}},
		{"var", []TokenType{VAR}, []string{"var"}, []int{3}},
		{"while", []TokenType{WHILE}, []string{"while"}, []int{5}},
	}
}

func TestToken_String(t *testing.T) {
	tts := GetTokenTypes()

	for i, tt := range tts {
		name := "Testing: " + tt.String + " Expecting correct parsing of element."
		t.Run(name, func(t *testing.T) {
			tkn := Token{
				Line:    i,
				Type:    tt.TokenTypes[0],
				Lexeme:  tt.Lexemes[0],
				Literal: "Test",
			}

			expected := fmt.Sprintf("Token: { Line: %d, Type: %s, Lexeme: \"%s\", Literal: %+v}", tkn.Line, tkn.Type, tkn.Lexeme, tkn.Literal)
			assert.Equal(t, expected, tkn.String(), "Expecting correct string output for token: "+expected)
		})
	}
}
