package scpi

import "fmt"

type Token struct {
	tType  TokenType
	lexeme string
	start  int
	length int
	line   int
}

func (t Token) debug() {
	fmt.Printf("%s %d-%d, [%d]\n", t.lexeme, t.start, t.length, t.line)
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
	ERROR
)

//func (tt TokenType) TokenString() string {
//	return [...]string{":", ",", " ", ";", "\"", "+"}[tt-1]
//}

func (tt TokenType) enumIndex() int {
	return int(tt)
}
