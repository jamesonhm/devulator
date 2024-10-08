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
	// Single Character
	BANG TokenType = iota
	AT
	HASH
	DOLLAR
	PERCENT
	CARROT
	AMPERSAND
	STAR
	LEFT_PAREN
	RIGHT_PAREN
	MINUS
	PLUS
	UNDERSCORE
	EQUAL
	LEFT_SQUARE
	RIGHT_SQUARE
	LEFT_CURLY
	RIGHT_CURLY
	PIPE
	BACKASLASH
	COLON
	SEMICOLON
	SQUOTE
	DQUOTE
	LESS
	GREATER
	COMMA
	DOT
	QUERY
	SLASH

	BANG_EQUAL
	WHITE_SPACE

	NUMBER
	IDENTIFIER

	ERROR
	EOF
)

//func (tt TokenType) TokenString() string {
//	return [...]string{":", ",", " ", ";", "\"", "+"}[tt-1]
//}

func (tt TokenType) enumIndex() int {
	return int(tt)
}
