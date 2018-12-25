package parser

import (
	"strconv"

	"github.com/th-lange/glox/scanner"
)

type ParserError struct {
	ErrorStart    scanner.Token
	TokenPosition int
	Message       string
}

func (se ParserError) Error() string {
	return "[Token Position " + strconv.Itoa(se.TokenPosition) + "] Last Valid Element" + (se.ErrorStart).String() + "! " + se.Message
}
