package scpi

import (
	"fmt"
)

func compile(source string) {
	runesrc := []rune(source)
	scanner := initScanner(runesrc)
	for {
		t := scanner.scanToken()
		fmt.Println("TYPE: ", t.tType)
		t.debug()
		if t.tType == EOF {
			break
		}
	}
}
