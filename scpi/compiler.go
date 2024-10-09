package scpi

// import (
// 	"fmt"
// )

func compile(source string) {
	scanner := initScanner(source)
	for {
		t := scanner.scanToken()
		t.debug()
		if t.tType == EOF {
			break
		}
	}
}
