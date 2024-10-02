package scpi

type Token struct {
	tokenType TokenType
	start     int
	length    int
	line      int
}

type TokenType int

const (
	COLON TokenType = iota
	COMMA
	WHITE_SPACE
	SEMICOLON
	DQUOTE
	PLUS
	MINUS
	DOT
	NUMBER
	NODE
	QUERY
	BANG
	AT
	HASH
	DOLLAR
	PERCENT
	CARROT
	LEFT_PAREN
	RIGHT_PAREN
	EOF
)

func (tt TokenType) TokenString() string {
	return [...]string{":", ",", " ", ";", "\"", "+"}[tt-1]
}

func (tt TokenType) EnumIndex() int {
	return int(tt)
}
