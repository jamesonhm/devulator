package scpi

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
	return s.errorToken("Unexpected character")
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) makeToken(tt TokenType) Token {
	return Token{
		tokenType: tt,
		start:     s.start,
		length:    s.current - s.start,
		line:      s.line,
	}
}
