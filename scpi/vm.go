package scpi

import "fmt"

type VM struct {
}

func (vm *VM) Interpret(src string) {

	scanner := intiScanner(src)
	for {
		t := scanner.scanToken()
		fmt.Println("TYPE: ", t.tType)
		if t.tType == EOF {
			break
		}
		t.debug()
	}
}
