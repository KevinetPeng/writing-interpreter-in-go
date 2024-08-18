package repl

import (
	"bufio"
	"fmt"
	"io"

	"interpreter/lexer"
	"interpreter/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(">> ")

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()

		l := lexer.NewLexer(line)

		var tok = l.NextToken()
		for tok.Type != token.EOF {
			fmt.Println(tok)
			tok = l.NextToken()
		}
	}
}
