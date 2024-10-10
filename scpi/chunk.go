package scpi

import "fmt"

type opCode uint8

type Chunk struct {
	code []opCode
}

func InitChunk() Chunk {
	return Chunk{
		code: make([]opCode, 0),
	}
}

func (c *Chunk) WriteChunk(b opCode) {
	c.code = append(c.code, b)
}

func (c Chunk) DisassembleChunk(name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(c.code); {
		offset = c.disassembleInstruction(offset)
	}
}

func (c Chunk) disassembleInstruction(offset int) int {
	fmt.Printf("%04d ", offset)

	var instruction opCode = c.code[offset]

	switch instruction {
	case OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

const (
	OP_RETURN opCode = iota + 1
)
