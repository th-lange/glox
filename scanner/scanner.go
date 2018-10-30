package scanner

import (
	"fmt"
	"io"
)

type Scanner struct {
	Errors   []error
	Tokens   []Token
	HadError bool
	Line     int
	current  int
	length   int
	lines    string
}

func (scnr *Scanner) Scan(lines string) {

	println(lines)
	scnr.Errors = make([]error, 0, 16)
	scnr.Tokens = make([]Token, 0, 32)
	scnr.current = 0
	scnr.Line = 1
	scnr.length = len(lines)
	scnr.lines = lines

	scnr.HadError = false
	scnr.scanTokens(lines)

	for _, itm := range scnr.Tokens {
		fmt.Println(itm)
	}

}

func (scnr *Scanner) appendError(err error) {
	scnr.Errors = append(scnr.Errors, err)
}

func (scnr *Scanner) appendToken(tkn Token) {
	scnr.Tokens = append(scnr.Tokens, tkn)
}

func (scnr *Scanner) scanTokens(line string) {
	fmt.Println("In ScanTokens!")
	for {
		cur, peek := scnr.nextChars()
		if cur == 0 {
			break
		}
		err := scnr.getNextToken(cur, peek)
		if err == io.EOF {
			scnr.appendEOFToken()
			return
		} else if err != nil {
			scnr.appendError(err)
		}
	}
}

func (scnr *Scanner) nextChars() (rune, rune) {
	if scnr.current+1 < scnr.length-1 {
		return rune(scnr.lines[scnr.current]), rune(scnr.lines[scnr.current+1])
	} else if scnr.current < scnr.length-1 {
		return rune(scnr.lines[scnr.current]), 0
	}
	return 0, 0

}

func (scnr *Scanner) getNextToken(cur, peek rune) error {
	tkn := Token{
		Position: scnr.current,
		Line:     scnr.Line,
		Lexeme:   string(cur),
		Length:   1,
	}

	switch cur {
	case ' ', '\r', '\t':
		scnr.current += 1
		return nil
	case '\n':
		scnr.current += 1
		scnr.Line += 1
		return nil
	case '(':
		tkn.Type = LEFT_PAREN
	case ')':
		tkn.Type = RIGHT_PAREN
	case '{':
		tkn.Type = LEFT_BRACE
	case '}':
		tkn.Type = RIGHT_BRACE
	case ',':
		tkn.Type = COMMA
	case '.':
		tkn.Type = DOT
	case '-':
		tkn.Type = MINUS
	case '+':
		tkn.Type = PLUS
	case ';':
		tkn.Type = SEMICOLON
	case '*':
		tkn.Type = STAR
	case '!':
		if peek == '=' {
			tkn.Type = BANG_EQUAL
			tkn.Length = 2
			tkn.Lexeme = "!="
			scnr.current += 1
		} else {
			tkn.Type = BANG
		}
	case '=':
		if peek == '=' {
			tkn.Type = EQUAL_EQUAL
			tkn.Length = 2
			tkn.Lexeme = "=="
			scnr.current += 1
		} else {
			tkn.Type = EQUAL
		}

	case '<':
		if peek == '=' {
			tkn.Type = LESS_EQUAL
			tkn.Lexeme = "<="
			tkn.Length = 2
			scnr.current += 1
		} else {
			tkn.Type = LESS
		}

	case '>':
		if peek == '=' {
			tkn.Type = GREATER_EQUAL
			tkn.Lexeme = ">="
			tkn.Length = 2
			scnr.current += 1
		} else {
			tkn.Type = GREATER
		}
	case '/':
		if peek == '/' {
			// consume til eol / eof
			// do not create lexeme
			return scnr.consume('\n')
		} else {
			tkn.Type = SLASH
		}
	case '"':
		err := scnr.string(&tkn)
		if err != nil {
			return err
		}
	default:
		return ScannerError{
			Position: scnr.current,
			Message:  "Unexpected character: " + string(cur),
		}

	}
	scnr.current += 1

	scnr.appendToken(tkn)
	return nil

}

func (scnr *Scanner) appendEOFToken() {
	scnr.appendToken(Token{
		Position: scnr.current,
		Lexeme:   "EOF",
		Length:   0,
		Type:     EOF,
	})
}

func (scnr *Scanner) consume(limiter rune) error {
	for scnr.current < scnr.length-1 && scnr.lines[scnr.current] != uint8(limiter) {
		if scnr.lines[scnr.current] == '\n' {
			scnr.Line += 1
		}
		scnr.current += 1
	}

	if scnr.current == scnr.length-1 {
		return io.EOF
	}
	return nil
}

func (scnr *Scanner) string(tkn *Token) error {
	// we will allow multi line strings
	scnr.current += 1
	err := scnr.consume('"')
	if err == io.EOF && scnr.lines[scnr.current] != '"' {
		return ScannerError{
			Line:     scnr.Line,
			Position: scnr.current,
			Message:  "Unterminated string",
		}
	}
	tkn.Type = STRING
	tkn.Position += 1 // Remove leading "
	tkn.Length = scnr.current - tkn.Position
	tkn.Literal = scnr.lines[tkn.Position : tkn.Position+tkn.Length]
	return nil
}
