package lexer

import (
	"testing"

	"interpreter/token"
)

func TestNextToken(t *testing.T) {
	testInput := "=+,;(){} !*/<>-"

	expectedOutputs := [...]token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.NOT, Literal: "!"},
		{Type: token.MULTIPLY, Literal: "*"},
		{Type: token.DIVIDE, Literal: "/"},
		{Type: token.LESSTHAN, Literal: "<"},
		{Type: token.GREATERTHAN, Literal: ">"},
		{Type: token.MINUS, Literal: "-"},
		{Type: token.EOF, Literal: ""},
	}

	l := NewLexer(testInput)

	for i, testToken := range expectedOutputs {
		tok := l.NextToken()

		if tok.Type != testToken.Type {
			t.Fatalf("Token type '%v' does not match expected type of '%v' at index: %d", tok.Type, testToken.Type, i)
		}

		if tok.Literal != testToken.Literal {
			t.Fatalf("Token literal '%v' does not match expected literal of '%v' at index: %d", tok.Literal, testToken.Literal, i)
		}
	}
}

func TestNextTokenWithKeywords(t *testing.T) {
	testInput := `
	let five = 5;
	let ten = 10;
   	let add = fun(x, y) {
    	x + y;
	};
   	let result = add(five, ten);

	let variable_one = 10 / 5 * 2;
	let variable_two = variable_one - 1;

	if(variable_two > 1) {
		return true;
	} else {
		return false;
	}

	10 == 10; 
	10 != 9;
   	`

	expectedOutputs := [...]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "five"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "ten"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FUNCTION, Literal: "fun"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "result"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "five"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "ten"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "variable_one"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.DIVIDE, Literal: "/"},
		{Type: token.INT, Literal: "5"},
		{Type: token.MULTIPLY, Literal: "*"},
		{Type: token.INT, Literal: "2"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "variable_two"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENTIFIER, Literal: "variable_one"},
		{Type: token.MINUS, Literal: "-"},
		{Type: token.INT, Literal: "1"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.IF, Literal: "if"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "variable_two"},
		{Type: token.GREATERTHAN, Literal: ">"},
		{Type: token.INT, Literal: "1"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},

		{Type: token.INT, Literal: "10"},
		{Type: token.EQUALS, Literal: "=="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.INT, Literal: "10"},
		{Type: token.NOTEQUALS, Literal: "!="},
		{Type: token.INT, Literal: "9"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.EOF, Literal: ""},
	}

	l := NewLexer(testInput)

	for i, testToken := range expectedOutputs {
		tok := l.NextToken()

		if tok.Type != testToken.Type {
			t.Fatalf("Token type '%v' does not match expected type of '%v' at index: %d", tok.Type, testToken.Type, i)
		}

		if tok.Literal != testToken.Literal {
			t.Fatalf("Token literal '%v' does not match expected literal of '%v' at index: %d", tok.Literal, testToken.Literal, i)
		}
	}
}
