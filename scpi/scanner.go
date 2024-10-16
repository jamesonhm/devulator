package scpi

import (
	"fmt"
	"strings"
	"unicode"
)

type scanner struct {
	source  []rune
	start   int
	current int
	line    int
}

func newScanner(src string) *scanner {
	src = trimLeftSpace(strings.ToUpper(src))
	runesrc := []rune(src)
	return &scanner{
		source:  runesrc,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *scanner) scanToken() Token {
	// s.skipWitespace()
	s.start = s.current
	if s.isAtEnd() {
		return s.makeToken(EOF)
	}
	c := s.advance()
	//	fmt.Println("From Scanner, next c: ", string(c))
	// if strings.ContainsRune(" \t", c) {
	// 	return s.whitespace()
	// }
	if unicode.IsLetter(c) || c == '*' {
		return s.identifier()
	}
	if unicode.IsDigit(c) {
		return s.number()
	}
	// TODO: add numeric cases for binary, hexadecimal, octal

	switch c {
	// 	case '!':
	// 		return s.makeToken(BANG)
	case '@':
		return s.makeToken(AT)
	case '#':
		return s.makeToken(HASH)
	case '$':
		return s.makeToken(DOLLAR)
	case '%':
		return s.makeToken(PERCENT)
	case '^':
		return s.makeToken(CARROT)
	case '&':
		return s.makeToken(AMPERSAND)
	case '(':
		return s.makeToken(LEFT_PAREN)
	case ')':
		return s.makeToken(RIGHT_PAREN)
	case '-':
		return s.makeToken(MINUS)
	case '_':
		return s.makeToken(UNDERSCORE)
	case '=':
		return s.makeToken(EQUAL)
	case '[':
		return s.makeToken(LEFT_SQUARE)
	case ']':
		return s.makeToken(RIGHT_SQUARE)
	case '{':
		return s.makeToken(LEFT_CURLY)
	case '}':
		return s.makeToken(RIGHT_CURLY)
	case '\\':
		return s.makeToken(BACKASLASH)
	case '|':
		return s.makeToken(PIPE)
	case ',':
		return s.makeToken(COMMA)
	case '?':
		return s.makeToken(QUERY)
	case ':':
		return s.makeToken(COLON)
	case ';':
		return s.makeToken(SEMICOLON)
	case '"':
		return s.string('"')
	case '\'':
		return s.string('\'')
	case '<':
		return s.makeToken(LESS)
	case '>':
		return s.makeToken(GREATER)
	case '.':
		return s.makeToken(DOT)
	case '/':
		return s.makeToken(SLASH)
	case '!':
		return s.makeToken(ternary(s.match('='), BANG_EQUAL, BANG))
	case ' ':
		return s.makeToken(WHITE_SPACE)
	case '\n':
		s.line++
		return s.makeToken(NEWLINE)
	// case '*':
	// 	return s.starWord()
	default:
		return s.errorToken(fmt.Sprintf("Unrecognized Keyword: %q", c))
	}
}

func ternary[T any](cond bool, tval T, fval T) T {
	if cond {
		return tval
	}
	return fval
}

func trimLeftSpace(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

func (s *scanner) advance() rune {
	s.current++
	return s.source[s.current-1]
}

func (s *scanner) peek(n int) rune {
	if s.current+n >= len(s.source) {
		return '\000'
	}
	return s.source[s.current+n]
}

func (s *scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) makeToken(tt TokenType) Token {
	return Token{
		tType:  tt,
		lexeme: string(s.source[s.start:s.current]),
		start:  s.start,
		length: s.current - s.start,
		line:   s.line,
	}
}

func (s *scanner) errorToken(msg string) Token {
	return Token{
		tType:  ERROR,
		lexeme: msg,
		start:  s.start,
		length: len(msg),
		line:   s.line,
	}
}

func (s *scanner) skipWitespace() {
	if s.isAtEnd() {
		return
	}
	for {
		c := s.peek(0)
		switch c {
		case ' ', '\t':
			s.advance()
		case '\n':
			s.line++
			s.advance()
		default:
			return
		}
	}
}

func (s *scanner) string(qtype rune) Token {
	for s.peek(0) != qtype && !s.isAtEnd() {
		if s.peek(0) == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		return s.errorToken("Unterminated string.")
	}
	s.advance()
	return s.makeToken(STRING)
}

func (s *scanner) whitespace() Token {
	for {
		c := s.peek(0)
		switch c {
		case ' ', '\t':
			s.advance()
		default:
			return s.makeToken(WHITE_SPACE)
		}
	}
}

func (s *scanner) word() string {
	for unicode.IsLetter(s.peek(0)) {
		s.advance()
	}
	return string(s.source[s.start:s.current])
}

func (s *scanner) identifier() Token {
	word := s.word()
	if wordInSet(word, SPECIALS) {
		return s.makeToken(SPECIAL)
	} else if wordInSet(word, BOOLS) {
		return s.makeToken(BOOL)
		// 	}
		// 	else if wordInSet(word, COMMONS) {
		// 		return s.makeToken(COMMON_CMD)
	} else {
		return s.makeToken(NODE)
	}
}

func (s *scanner) number() Token {
	for unicode.IsDigit(s.peek(0)) {
		s.advance()
	}
	if s.peek(0) == '.' && unicode.IsDigit(s.peek(1)) {
		s.advance()

		for unicode.IsDigit(s.peek(0)) {
			s.advance()
		}
	}
	return s.makeToken(NUMBER)
}
