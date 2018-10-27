package scanner

import "fmt"

type Token struct {
	Type     TokenType
	Lexeme   string
	Literal  interface{}
	Line     int
	Position int
	Length   int
}

func (t Token) String() string {
	return fmt.Sprintf("Token: { Line: %d, Type: %s, Lexeme: \"%s\", Literal: %+v}", t.Line, t.Type, t.Lexeme, t.Literal)
}
