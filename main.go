package main

import (
	"fmt"
	"os"

	"github.com/jamesonhm/devulator/scpi"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(70)
	}
}

func run() error {
	src := ":,(!@#$%)"
	fmt.Println("Source: ", src)
	vm := scpi.VM{}
	return vm.Interpret(src)
}
