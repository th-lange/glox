package parser

import (
	"strconv"

	"github.com/th-lange/glox/scanner"
)

// Indicates that the PARSED code is erroneous
type ParsingError struct {
	ErrorStart    scanner.Token
	TokenPosition int
	Message       string
	SyncEnd       *scanner.Token
}

func (pe ParsingError) Error() string {
	endMsg := ""
	if pe.SyncEnd != nil {
		endMsg = "\nSynced Up to Element: " + pe.SyncEnd.String()
	}
	return "[Token Position " + strconv.Itoa(pe.TokenPosition) + "] Last Valid Element" + (pe.ErrorStart).String() + "!\n" + pe.Message + endMsg
}

func NewError(message string, sync bool, prs *parser) ParsingError {

	err := ParsingError{prs.current(), prs.head, message, nil}
	if sync {
		prs.synchronize()
		err.SyncEnd = &(*prs.tokens)[prs.head]
	}
	return err
}

// Indicates issues with the PARSER code is erroneous
type InvalidArgumentError struct {
	message string
}

func (iae InvalidArgumentError) Error() string {
	return iae.message
}
