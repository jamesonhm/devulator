package scpi

import "unicode"

// import "fmt"

type Scanner struct {
	source  string
	start   int
	current int
	line    int
}

func (s *Scanner) scanToken() Token {
	s.skipWitespace()
	s.start = s.current
	if s.isAtEnd() {
		return s.makeToken(EOF)
	}
	c := s.advance()
	//	fmt.Println("From Scanner, next c: ", string(c))
	if unicode.IsDigit(rune(c)) {
		return s.number()
	}

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
	case '*':
		return s.makeToken(STAR)
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
	case '\'':
		return s.makeToken(SQUOTE)
	case '"':
		return s.makeToken(DQUOTE)
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
	}

	return s.errorToken("Unexpected character")
}

func ternary[T any](cond bool, tval T, fval T) T {
	if cond {
		return tval
	}
	return fval
}

func initScanner(src string) Scanner {
	return Scanner{
		source:  src,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) peek(n int) byte {
	if s.current+n >= len(s.source) {
		return '\000'
	}
	return s.source[s.current+n]
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) makeToken(tt TokenType) Token {
	return Token{
		tType:  tt,
		lexeme: s.source[s.start:s.current],
		start:  s.start,
		length: s.current - s.start,
		line:   s.line,
	}
}

func (s *Scanner) errorToken(msg string) Token {
	return Token{
		tType:  ERROR,
		lexeme: msg,
		start:  s.start,
		length: len(msg),
		line:   s.line,
	}
}

func (s *Scanner) skipWitespace() {
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

func (s *Scanner) number() Token {
	for unicode.IsDigit(rune(s.peek(0))) {
		s.advance()
	}
	if s.peek(0) == '.' && unicode.IsDigit(rune(s.peek(1))) {
		s.advance()

		for unicode.IsDigit(rune(s.peek(0))) {
			s.advance()
		}
	}
	return s.makeToken(NUMBER)
}
