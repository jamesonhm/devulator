package scpi

import (
	"fmt"
)

func compile(source string) {
	scanner := intiScanner(source)
	for {
		t := scanner.scanToken()
		fmt.Println("TYPE: ", t.tType)
		if t.tType == EOF {
			break
		}
		t.debug()
	}
}
