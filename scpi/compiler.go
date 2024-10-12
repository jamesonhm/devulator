package scpi

import (
	"fmt"
	"os"
)

type parser struct {
	scanner   *scanner
	current   Token
	previous  Token
	hadError  bool
	panicMode bool
}

var compilingChunk *Chunk

func currentChunk() *Chunk {
	return compilingChunk
}

func initParser(source string) *parser {
	p := &parser{
		scanner:   initScanner(source),
		hadError:  false,
		panicMode: false,
	}
	return p
}

func (p *parser) advance() {
	p.previous = p.current

	for {
		p.current = p.scanner.scanToken()
		if p.current.tType != ERROR {
			break
		}
		p.errorAtCurrent(p.current.lexeme)
	}
}

func (p *parser) consume(tType TokenType, message string) {
	if p.current.tType == tType {
		p.advance()
		return
	}

	p.errorAtCurrent(message)
}

func (p *parser) check(tType TokenType) bool {
	return p.current.tType == tType
}

func (p *parser) match(tType TokenType) bool {
	if !p.check(tType) {
		return false
	}
	p.advance()
	return true
}

func (p *parser) errorAtCurrent(message string) {
	p.errorAt(p.current, message)
}

func (p *parser) errorPrev(message string) {
	p.errorAt(p.previous, message)
}

func (p *parser) errorAt(token Token, message string) {
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

func compile(source string, chunk *Chunk) bool {
	p := initParser(source)
	compilingChunk = chunk

	p.hadError = false
	p.panicMode = false
	p.advance()

	for !p.match(EOF) {
		programHeader(p)
	}

	p.consume(EOF, "Expect end of expression")
	// p.endCompiler()
	return !p.hadError
}

// programHeader produces a cmd struct
// includes the tree sequence, query, args
// parse -> vm uses a queue structure to eval programHeader units in fifo sequence order
func programHeader(p *parser) {
	if p.match(COLON) {
		// Leading colon starts the command at the root
		instrumentHeader(p)
	}
	if p.match(COMMON_CMD) {
		commonHeader(p)
	}

}

func instrumentHeader(p *parser) {
	return
}

func commonHeader(p *parser) {
	return
}
