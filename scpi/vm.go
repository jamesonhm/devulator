package scpi

//import "fmt"

type VM struct {
}

type InterpretResult int

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

func (vm *VM) Interpret(src string) InterpretResult {
	compile(src)
	return INTERPRET_OK
}
