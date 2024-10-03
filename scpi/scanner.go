package scpi

import "fmt"

type Scanner struct {
	source  string
	start   int
	current int
	line    int
}

func (s *Scanner) scanToken() Token {
	s.start = s.current
	if s.isAtEnd() {
		return s.makeToken(EOF)
	}
	c := s.advance()
	fmt.Println("From Scanner, next c: ", string(c))

	switch c {
	case ':':
		fmt.Println("Matched colon")
		return s.makeToken(COLON)
	case ',':
		fmt.Println("Matched comma")
		return s.makeToken(COMMA)
	case '?':
		return s.makeToken(QUERY)
	case '!':
		return s.makeToken(BANG)
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
	case '(':
		return s.makeToken(LEFT_PAREN)
	case ')':
		return s.makeToken(RIGHT_PAREN)
	}

	return s.errorToken("Unexpected character")
}

func intiScanner(src string) Scanner {
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
