package main

import (
	"fmt"
	"os"

	"github.com/jamesonhm/devulator/scpi"
)

func main() {
	code, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(code)
	}
}

// func run() (int, error) {
// 	chunk := scpi.InitChunk()
// 	constloc := chunk.AddConstant(1.2)
// 	chunk.WriteChunk(scpi.OP_CONSTANT, 123)
// 	chunk.WriteChunk(constloc, 123)
//
// 	chunk.WriteChunk(scpi.OP_RETURN, 123)
//
// 	chunk.DisassembleChunk("test chunk")
// 	return 0, nil
// }

func run() (int, error) {
	// 	src := `  COMMand:stuFF?;*ABC  *!=
	// 	:DATA 4,"this is a string";
	// 	12
	// 	23.4
	// 	abort:AbO:CalculaTE`
	src := `TEST:A 1
	:SOMETHING:OTHER;"
	*RST:*ESE?`
	fmt.Println("Source: ", src)

	vm := scpi.VM{}

	res := vm.Interpret(src)
	if res != scpi.INTERPRET_OK {
		switch res {
		case scpi.INTERPRET_COMPILE_ERROR:
			return 65, fmt.Errorf("Compile Error")
		case scpi.INTERPRET_RUNTIME_ERROR:
			return 70, fmt.Errorf("Runtime Error")
		}
	}
	return 0, nil
}
