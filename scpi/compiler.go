package scpi

import (
	"fmt"
	"os"
)

type precedence int

const (
	PREC_NONE   precedence = iota
	PREC_OR                // or, exor
	PREC_AND               // and
	PREC_NOT               // not
	PREC_TERM              // + , -
	PREC_FACTOR            // *, /, MOD, DIV
	PREC_EXP               // ^
	PREC_UNARY             // +, -
	PREC_GROUP             // ()
	PREC_PRIMARY
)

type parseFn func(p *parser)

type parseRule struct {
	prefix parseFn
	infix  parseFn
	prec   precedence
}

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

func newParser(source string) *parser {
	p := &parser{
		scanner:   newScanner(source),
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

///////////////////////////////////////////////////////////////////////////////////

func compile(source string, chunk *Chunk) bool {
	p := newParser(source)
	compilingChunk = chunk

	p.hadError = false
	p.panicMode = false
	p.advance()
	fmt.Printf("token after first advance: %v\n", p.current)
	var count int = 0
	for !p.match(EOF) {
		if count > 5 {
			break
		}
		programHeader(p)
		count++
	}

	p.consume(EOF, "Expect end of expression")
	// p.endCompiler()
	return !p.hadError
}

// programHeader produces a header tree
// includes the tree sequence, query, args
// parse -> vm uses a queue structure to eval programHeader units in fifo sequence order
func programHeader(p *parser) {
	if p.check(COMMON_CMD) {
		commonHeader(p)
	}
	instrumentHeader(p)

}

func instrumentHeader(p *parser) {
	fmt.Println("Inst Header: ")
	cmdTree := make([]Token, 0)
	var count int = 0
	for {
		if count > 5 {
			break
		}
		fmt.Printf("instr loop, current token: %q, %s\n", p.current.lexeme, p.current.tType.tokenString())
		if p.match(COLON) {
			continue
		} else if p.match(NODE) {
			cmdTree = append(cmdTree, p.previous)
		} else if p.match(WHITE_SPACE) {
			argList(p)
			// Output call with args
		} else if p.match(QUERY) {
			fmt.Println("Query")
		} else if p.match(SEMICOLON) {
			// Output call without args
			break
		} else if p.match(NEWLINE) {
			// Output call without args
			break
		}
		count++
	}
	fmt.Println("cmd tree: ", cmdTree)
}

func commonHeader(p *parser) {
	fmt.Println("common Header")
	p.advance()
}

func argList(p *parser) { // []Token {
	fmt.Printf("arglist: ")
	args := make([]value, 0)
	args = append(args, p.current)
	var count int = 0
	for p.match(COMMA) {
		if count > 5 {
			break
		}
		args = append(args, p.current)
		p.advance()
		count++
	}
	fmt.Println(args)
	p.advance()
}

var rules = map[TokenType]parseRule{
	LEFT_PAREN:  parseRule{grouping, nil, PREC_GROUP},
	RIGHT_PAREN: parseRule{nil, nil, PREC_NONE},
	MINUS:       parseRule{unary, binary, PREC_TERM},
	NUMBER:      parseRule{number, nil, PREC_NONE},
}

func getRule(tType TokenType) parseRule {
	rule, ok := rules[tType]
	if !ok {
		return parseRule{}
	}
	return rule
}

func parsePrecedence(p *parser, prec precedence) {
	p.advance()
	var prefixRule parseFn = getRule(p.previous.tType).prefix
	if prefixRule == nil {
		p.errorPrev("Expect Expression")
		return
	}
	prefixRule(p)

	for prec < getRule(p.current.tType).prec {
		p.advance()
		var infixRule parseFn = getRule(p.previous.tType).infix
		infixRule(p)
	}
}

func expression(p *parser) {
	parsePrecedence(p, PREC_OR)
}

func grouping(p *parser) {
	expression(p)
	p.consume(RIGHT_PAREN, "Expect right paren")
}
