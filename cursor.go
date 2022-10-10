package golexer

import "errors"

type Cursor interface {
	Advance()
	GetChar() byte
	IsEOF() bool
	GetPosition() int
	GetLength() int
}

type cursor struct {
	char     byte
	content  []byte
	position int
	length   int
}

func (c *cursor) Advance() {
	if c.position < c.length {
		c.char = c.content[c.position]
		c.position += 1
	} else {
		c.char = EOF
	}
}

func (c *cursor) GetChar() byte {
	return c.char
}

func (c *cursor) IsEOF() bool {
	return c.char == EOF
}

func (c *cursor) GetPosition() int {
	return c.position
}

func (c *cursor) GetLength() int {
	return c.length
}

func newCursor(b []byte) (Cursor, error) {
	if b == nil {
		return nil, errors.New("the cursor's content cannot be nil")
	}

	return &cursor{
		char:     EOF,
		position: -1,
		content:  b,
		length:   len(b),
	}, nil

}
