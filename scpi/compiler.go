package scpi

import (
	"fmt"
)

func compile(source string) {
	scanner := initScanner(source)
	for {
		t := scanner.scanToken()
		fmt.Println("TYPE: ", t.tType)
		t.debug()
		if t.tType == EOF {
			break
		}
	}
}
