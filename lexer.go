package golexer

import (
	"fmt"
)

type LexerFunction[T any] func(*Cursor) Token[T]

type Lexer[T any] interface {
	LexNumber(LexerFunction[T])
	LexString(LexerFunction[T])
	LexIdentifier(LexerFunction[T])
	LexDelimiter(LexerFunction[T])
	LexComment(LexerFunction[T])
	Tokenize([]byte) ([]Token[T], error)
}

type LexerSettings struct {
	Numbers    ByteRange
	Chars      ByteRange
	String     BytePoints
	Delimiter  BytePoints
	Ignore     BytePoints
	Comment    BytePoints
	PlainChars BytePoints
}

type lexerFunctionList[T any] []LexerFunction[T]

func (l *lexerFunctionList[T]) lex(c *Cursor, t *[]Token[T]) {
	for _, f := range *l {
		*t = append(*t, f(c))
	}
}

type lexer struct {
	settings      LexerSettings
	lexNumber     lexerFunctionList[TokenType]
	lexString     lexerFunctionList[TokenType]
	lexIdentifier lexerFunctionList[TokenType]
	lexComment    lexerFunctionList[TokenType]
	lexDelimiter  lexerFunctionList[TokenType]
}

func (l *lexer) LexNumber(lex LexerFunction[TokenType]) {
	l.lexNumber = append(l.lexNumber, lex)
}

func (l *lexer) LexString(lex LexerFunction[TokenType]) {
	l.lexString = append(l.lexNumber, lex)
}

func (l *lexer) LexIdentifier(lex LexerFunction[TokenType]) {
	l.lexIdentifier = append(l.lexNumber, lex)
}

func (l *lexer) LexDelimiter(lex LexerFunction[TokenType]) {
	l.lexDelimiter = append(l.lexNumber, lex)
}

func (l *lexer) LexComment(lex LexerFunction[TokenType]) {
	l.lexComment = append(l.lexNumber, lex)
}

func (l *lexer) Tokenize(b []byte) ([]Token[TokenType], error) {
	var tokens []Token[TokenType]
	c, cerr := newCursor(b)

	if cerr != nil {
		return nil, cerr
	}

	c.Advance()

	for !c.IsEOF() {
		_c := c.GetChar()

		switch {
		case l.settings.Ignore.HasPoint(_c):
			c.Advance()

		case l.settings.Numbers.IsInRange(_c):
			l.lexNumber.lex(&c, &tokens)

		case l.settings.String.HasPoint(_c):
			c.Advance()
			l.lexString.lex(&c, &tokens)

		case l.settings.Chars.IsInRange(_c) || l.settings.PlainChars.HasPoint(_c):
			l.lexIdentifier.lex(&c, &tokens)

		case l.settings.Comment.HasPoint(_c):
			c.Advance()
			l.lexComment.lex(&c, &tokens)

		case l.settings.Delimiter.HasPoint(_c):
			l.lexDelimiter.lex(&c, &tokens)

		default:
			return nil, fmt.Errorf("unexpected token %s", string(_c))
		}

	}

	return tokens, nil
}

func NewLexer(s LexerSettings) (Lexer[TokenType], error) {

	if s.Numbers == nil {
		s.Numbers = defaultByteRange{}
	}

	if s.Chars == nil {
		s.Chars = defaultByteRange{}
	}

	if s.Ignore == nil {
		s.Ignore = defaultBytePoint{}
	}

	if s.Comment == nil {
		s.Comment = defaultBytePoint{}
	}

	if s.Delimiter == nil {
		s.Delimiter = defaultBytePoint{}
	}

	if s.PlainChars == nil {
		s.PlainChars = defaultBytePoint{}
	}

	if s.String == nil {
		s.String = defaultBytePoint{}
	}

	return &lexer{
		settings:      s,
		lexNumber:     make(lexerFunctionList[TokenType], 0),
		lexString:     make(lexerFunctionList[TokenType], 0),
		lexIdentifier: make(lexerFunctionList[TokenType], 0),
		lexDelimiter:  make(lexerFunctionList[TokenType], 0),
		lexComment:    make(lexerFunctionList[TokenType], 0),
	}, nil

}
