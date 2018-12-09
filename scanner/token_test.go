package scanner

type TokenTestData struct {
	String     string
	TokenTypes []TokenType
}

func GetTokenTypes() []TokenTestData {
	return [] TokenTestData{
		TokenTestData{"(", []TokenType{LEFT_PAREN}},
		TokenTestData{")", []TokenType{RIGHT_PAREN}},
		TokenTestData{"{", []TokenType{LEFT_BRACE}},
		TokenTestData{"}", []TokenType{RIGHT_BRACE}},

		TokenTestData{",", []TokenType{COMMA}},
		TokenTestData{".", []TokenType{DOT}},
		TokenTestData{"-", []TokenType{MINUS}},

		TokenTestData{"+", []TokenType{PLUS}},
		TokenTestData{";", []TokenType{SEMICOLON}},

		TokenTestData{"/", []TokenType{SLASH}},
		TokenTestData{"*", []TokenType{STAR}},

		TokenTestData{"!", []TokenType{BANG}},
		TokenTestData{"!=", []TokenType{BANG_EQUAL}},

		TokenTestData{"=", []TokenType{EQUAL}},
		TokenTestData{"==", []TokenType{EQUAL_EQUAL}},

		TokenTestData{">", []TokenType{GREATER}},
		TokenTestData{">=", []TokenType{GREATER_EQUAL}},

		TokenTestData{"<", []TokenType{LESS}},
		TokenTestData{"<=", []TokenType{LESS_EQUAL}},

		TokenTestData{"fooBar", []TokenType{IDENTIFIER}},

		TokenTestData{"\"This is a string\"", []TokenType{STRING}},
		TokenTestData{"123", []TokenType{NUMBER}},

		TokenTestData{"and", []TokenType{AND}},
		TokenTestData{"class", []TokenType{CLASS}},
		TokenTestData{"else", []TokenType{ELSE}},
		TokenTestData{"false", []TokenType{FALSE}},
		TokenTestData{"for", []TokenType{FOR}},
		TokenTestData{"fun", []TokenType{FUN}},
		TokenTestData{"if", []TokenType{IF}},
		TokenTestData{"nil", []TokenType{NIL}},
		TokenTestData{"or", []TokenType{OR}},
		TokenTestData{"print", []TokenType{PRINT}},
		TokenTestData{"return", []TokenType{RETURN}},
		TokenTestData{"super", []TokenType{SUPER}},
		TokenTestData{"this", []TokenType{THIS}},
		TokenTestData{"true", []TokenType{TRUE}},
		TokenTestData{"var", []TokenType{VAR}},
		TokenTestData{"while", []TokenType{WHILE}},
	}
}
