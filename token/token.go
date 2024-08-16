package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF" //End of file

	// Identifiers + literals
	IDENTIFIER TokenType = "IDENTIFIER"
	INT        TokenType = "INT"

	// Operators
	ASSIGN      TokenType = "="
	PLUS        TokenType = "+"
	MINUS       TokenType = "-"
	NOT         TokenType = "!"
	MULTIPLY    TokenType = "*"
	DIVIDE      TokenType = "/"
	LESSTHAN    TokenType = "<"
	GREATERTHAN TokenType = ">"
	EQUALS      TokenType = "=="
	NOTEQUALS   TokenType = "!="

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
)

// Using map of empty structs as a set
var OperatorsAndDelimiters = map[TokenType]struct{}{
	ASSIGN:      {},
	PLUS:        {},
	MINUS:       {},
	NOT:         {},
	MULTIPLY:    {},
	DIVIDE:      {},
	LESSTHAN:    {},
	GREATERTHAN: {},

	COMMA:     {},
	SEMICOLON: {},
	LPAREN:    {},
	RPAREN:    {},
	LBRACE:    {},
	RBRACE:    {},
}

var Keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}
