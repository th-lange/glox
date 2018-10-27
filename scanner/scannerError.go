package scanner

import (
	"strconv"
)

type ScannerError struct {
	Line     int
	Position int
	Message  string
}

func (se ScannerError) Error() string {
	return "[Line " + strconv.Itoa(se.Line) + "] HadError" + strconv.Itoa(se.Position) + ": " + se.Message
}
