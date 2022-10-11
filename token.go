package golexer

type Token[T any] interface {
	GetValue() []byte
	GetType() T
}

type TokenType int

const (
	Number TokenType = iota
	String
	Identifier
	Keyword
	Delimiter
	Comment
)

type token struct {
	value []byte
	_type TokenType
}

func (t token) GetValue() []byte {
	return t.value
}

func (t token) GetType() TokenType {
	return t._type
}

func NewToken(t TokenType, v []byte) Token[TokenType] {
	return token{
		_type: t,
		value: v,
	}
}
