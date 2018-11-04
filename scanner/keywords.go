package scanner

// var scannerKeywords = make(map[string]TokenType){}
//
// scannerKeywords = map[string]TokenType {
// 	"f00"
// }

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func adjustForKeywords(tkn *Token) {
	if value, ok := keywords[tkn.Lexeme]; ok {
		tkn.Type = value
	}
}
