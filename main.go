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
	testCmds := []string{
		"MEASure:VOLTage:RISE:TIME 0.2,4.5,.005\n",
		"MEASure:VOLTage:RISE:TIME 0.2V,4.5V,.005S\n",
		"*RST:*ESE?",
	}

	vm := scpi.VM{}

	for i, cmd := range testCmds {
		if i >= 1 {
			return 0, nil
		}
		fmt.Println("Source: ", cmd)

		res := vm.Interpret(cmd)
		if res != scpi.INTERPRET_OK {
			switch res {
			case scpi.INTERPRET_COMPILE_ERROR:
				return 65, fmt.Errorf("Compile Error")
			case scpi.INTERPRET_RUNTIME_ERROR:
				return 70, fmt.Errorf("Runtime Error")
			}
		}
	}
	return 0, nil
}
