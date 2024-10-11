package scpi

import "os"

import (
	"fmt"
)

type Parser struct {
	scanner   Scanner
	current   Token
	previous  Token
	hadError  bool
	panicMode bool
}

func (p *Parser) compile(source string, chunk *Chunk) bool {
	p.scanner = initScanner(source)

	p.hadError = false
	p.panicMode = false
	p.advance()

	p.consume(EOF, "Expect end of expression")
	return !p.hadError
}

func (p *Parser) advance() {
	p.previous = p.current

	for {
		p.current = p.scanner.scanToken()
		if p.current.tType != ERROR {
			break
		}
		p.errorAtCurrent(p.current.lexeme)
	}
}

func (p *Parser) consume(tType TokenType, message string) {
	if p.current.tType == tType {
		p.advance()
		return
	}

	p.errorAtCurrent(message)
}

func (p *Parser) errorAtCurrent(message string) {
	p.errorAt(p.current, message)
}

func (p *Parser) errorPrev(message string) {
	p.errorAt(p.previous, message)
}

func (p *Parser) errorAt(token Token, message string) {
	if p.panicMode {
		return
	}
	p.panicMode = true
	fmt.Fprintf(os.Stderr, "[line %d] Error", token.line)

	if token.tType == EOF {
		fmt.Fprintf(os.Stderr, " at end")
	} else if token.tType == ERROR {

	} else {
		fmt.Fprintf(os.Stderr, " at '%.*s'", token.length, token.lexeme)
	}

	fmt.Fprintf(os.Stderr, ": %s\n", message)
	p.hadError = true
}
