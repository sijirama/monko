package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = "-> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if tok.Literal == "exit" {
				return
			}
			fmt.Printf("%+v\n", tok)
		}
	}
}
