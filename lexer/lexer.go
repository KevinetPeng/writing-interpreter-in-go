package lexer

import (
	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	char         byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// Lexer Methods
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPosition]
		l.position = l.nextPosition
		l.nextPosition++
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}

	return l.input[l.nextPosition]
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) parseWord() token.Token {
	word := string(l.char)

	for isLetter(l.peekChar()) {
		l.readChar()
		word += string(l.char)
	}

	if tokenType, ok := token.Keywords[word]; ok {
		return token.Token{
			Type:    tokenType,
			Literal: word,
		}
	}

	return token.Token{
		Type:    token.IDENTIFIER,
		Literal: word,
	}
}

func (l *Lexer) parseNumber() token.Token {
	numberString := string(l.char)

	for isDigit(l.peekChar()) {
		l.readChar()
		numberString += string(l.char)
	}

	return token.Token{
		Type:    token.INT,
		Literal: numberString,
	}
}

func (l *Lexer) consumeWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	l.readChar()
	l.consumeWhitespace()

	// Handle multi-character operators
	if l.char == '=' && l.peekChar() == '=' {
		l.readChar()
		return token.Token{
			Type:    token.EQUALS,
			Literal: "==",
		}
	}

	if l.char == '!' && l.peekChar() == '=' {
		l.readChar()
		return token.Token{
			Type:    token.NOTEQUALS,
			Literal: "!=",
		}
	}

	if _, exists := token.OperatorsAndDelimiters[token.TokenType(l.char)]; exists {
		return token.Token{
			Type:    token.TokenType(l.char),
			Literal: string(l.char),
		}
	}

	if isLetter(l.char) {
		return l.parseWord()
	}

	if isDigit(l.char) {
		return l.parseNumber()
	}

	if l.char == 0 {
		return token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	}

	return token.Token{Type: token.ILLEGAL, Literal: ""}
}
