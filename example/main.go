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

	for ch := (*c).GetChar(); !(*c).IsEOF() && NUMBS.IsInRange(ch); {
		b = append(b, ch)
		(*c).Advance()
	}

	return lexer.NewToken(lexer.Number, b)
}

func main() {
	l, lerr := lexer.NewLexer(lexer.LexerSettings{
		Numbers: NUMBS,
		Ignore: IGNORE,
	})

	if lerr != nil {
		fmt.Printf("lerr: %v\n", lerr)
		return
	}

	l.LexNumber(lexNumbs)

	if t, terr := l.Tokenize([]byte("0123456789 23453222")); terr != nil {
		fmt.Printf("terr: %v\n", terr)
	} else {
		fmt.Printf("t: %v\n", t)
	}
}
