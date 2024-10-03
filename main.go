package main

import (
	//"fmt"

	"github.com/jamesonhm/devulator/scpi"
)

func main() {
	vm := scpi.VM{}
	vm.Interpret(":,(!@#$%)")
}
