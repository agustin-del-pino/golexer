package main

import (
	"fmt"

	lexer "github.com/agustin-del-pino/golexer"
)

var NUMBS = lexer.NewSingleByteRange(0x30, 0x39)
var IGNORE = lexer.NewBPoints(0x20)

func lexNumbs(c *lexer.Cursor) lexer.Token[lexer.TokenType] {
	var b []byte

	b = append(b, (*c).GetChar())

	(*c).Advance()

	for NUMBS.IsInRange((*c).GetChar()) {
		b = append(b, (*c).GetChar())
		(*c).Advance()
	}

	return lexer.NewToken(lexer.Number, b)
}

func main() {
	l, lerr := lexer.NewLexer(lexer.LexerSettings{
		Numbers: NUMBS,
		Ignore:  IGNORE,
	})

	if lerr != nil {
		fmt.Printf("lerr: %v\n", lerr)
		return
	}

	l.LexNumber(lexNumbs)

	if t, terr := l.Tokenize([]byte("001100 987654321 123 456 789")); terr != nil {
		fmt.Printf("terr: %v\n", terr)
	} else {
		fmt.Printf("t: %v\n", t)
	}
}
