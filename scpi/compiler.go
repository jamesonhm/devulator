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

var compilingChunk *Chunk

func currentChunk() *Chunk {
	return compilingChunk
}

func (p *Parser) compile(source string, chunk *Chunk) bool {
	p.scanner = initScanner(source)
	compilingChunk = chunk

	fmt.Println("compiler, scanner init done")

	p.hadError = false
	p.panicMode = false
	p.advance()

	fmt.Println("first advance, curr token: ", p.current)

	p.command()

	p.consume(EOF, "Expect end of expression")
	// p.endCompiler()
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

// command produces a cmd struct
// includes the tree sequence, query, args
// parse -> vm uses a queue structure to eval command units in fifo sequence order
func (p *Parser) command() {
	var cmdseq = make([]Token, 0)
	for {
		if p.match(WHITE_SPACE) {
			fmt.Println("reached end of cmd tree, handle args")
			break
		}
		if p.match(SEMICOLON) {
			fmt.Println("semicolon - command separator")
		}
		if p.match(COLON) {
			// Leading colon starts the command at the root
			continue
		}
		if p.match(QUERY) {
			fmt.Println("Query")
		}
		if p.match(NEWLINE) {
			fmt.Println("End of line")
			break
		}
		cmdseq = append(cmdseq, p.current)
		p.advance()
	}
	fmt.Printf("%v", cmdseq)
	fmt.Printf("\n")
}

func (p *Parser) check(tType TokenType) bool {
	return p.current.tType == tType
}

func (p *Parser) match(tType TokenType) bool {
	if !p.check(tType) {
		return false
	}
	p.advance()
	return true
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
