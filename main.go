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

func run() (int, error) {
	src := `  :,(!@#$%)  *!=
	@@@ _()
	12
	23.4
	abort:AbO:CalculaTE`

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
