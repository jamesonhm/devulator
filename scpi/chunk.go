package scpi

//import "fmt"

type Chunk struct {
	code      []uint8
	lines     []int
	constants any
}

func InitChunk() Chunk {
	return Chunk{
		code:      make([]uint8, 1),
		lines:     make([]int, 0),
		constants: make([]int, 0), //initValueArray(),
	}
}

//func (c *Chunk) WriteChunk(b uint8, line int) {
//	c.code = append(c.code, b)
//	c.lines = append(c.lines, line)
//}
//
//func (c *Chunk) AddConstant(value Value) uint8 {
//	c.constants.writeValue(value)
//	return uint8(len(c.constants.values) - 1)
//}
//
//func (c Chunk) DisassembleChunk(name string) {
//	fmt.Printf("== %s ==\n", name)
//
//	for offset := 0; offset < len(c.code); {
//		offset = c.disassembleInstruction(offset)
//	}
//}
//
//func (c Chunk) disassembleInstruction(offset int) int {
//	fmt.Printf("%04d ", offset)
//	if offset > 0 && c.lines[offset] == c.lines[offset-1] {
//		fmt.Printf("   | ")
//	} else {
//		fmt.Printf("%4d ", c.lines[offset])
//	}
//
//	var instruction uint8 = c.code[offset]
//
//	switch instruction {
//	case OP_CONSTANT:
//		return constantInstruction("OP_CONSTANT", c, offset)
//	case OP_RETURN:
//		return simpleInstruction("OP_RETURN", offset)
//	default:
//		fmt.Printf("Unknown opcode %d\n", instruction)
//		return offset + 1
//	}
//}
//
//func simpleInstruction(name string, offset int) int {
//	fmt.Printf("%s\n", name)
//	return offset + 1
//}
//
//func constantInstruction(name string, c Chunk, offset int) int {
//	constloc := uint8(c.code[offset+1])
//	fmt.Printf("%-16s %4d '", name, constloc)
//	printValue(c.constants.values[constloc])
//	fmt.Printf("'\n")
//	return offset + 2
//}
//
//const (
//	OP_RETURN uint8 = iota + 1
//	OP_CONSTANT
//)
