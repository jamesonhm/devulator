package main

import (
	"fmt"

	"github.com/jamesonhm/devulator/scpi"
)

func main() {
	fmt.Println("Manual main")
	fmt.Println(scpi.Scan("empty"))
	fmt.Println(scpi.Lex("empty"))
}
