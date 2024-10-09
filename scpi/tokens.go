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
	fmt.Printf("[%d] %v\t   %q\n", t.line, t.tType.tokenString(), t.lexeme)
}

type TokenType int

const (
	// Single Character
	BANG TokenType = iota + 1
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
	LESS
	GREATER
	COMMA
	DOT
	QUERY
	SLASH

	BANG_EQUAL
	WHITE_SPACE
	NEWLINE

	NUMBER
	NODE
	UNITS
	BOOL
	SPECIAL // MIN, MAX, INF, ...
	COMMON_CMD
	STRING

	ERROR
	EOF
)

func (tt TokenType) tokenString() string {
	return [...]string{
		"_",
		"BANG",
		"AT",
		"HASH",
		"DOLLAR",
		"PERCENT",
		"CARROT",
		"AMPERSAND",
		"STAR",
		"LEFT_PAREN",
		"RIGHT_PAREN",
		"MINUS",
		"PLUS",
		"UNDERSCORE",
		"EQUAL",
		"LEFT_SQUARE",
		"RIGHT_SQUARE",
		"LEFT_CURLY",
		"RIGHT_CURLY",
		"PIPE",
		"BACKASLASH",
		"COLON",
		"SEMICOLON",
		"LESS",
		"GREATER",
		"COMMA",
		"DOT",
		"QUERY",
		"SLASH",
		"BANG_EQUAL",
		"WHITE_SPACE",
		"NEWLINE",
		"NUMBER",
		"NODE",
		"UNITS",
		"BOOL",
		"SPECIAL",
		"COMMON_CMD",
		"STRING",
		"ERROR",
		"EOF"}[tt]
}

func (tt TokenType) enumIndex() int {
	return int(tt)
}
