package main

import (
	"flag"
	"fmt"

	"github.com/shell909090/cryptpass"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		s, err := cryptpass.EncryptPass(arg)
		if err != nil {
			fmt.Println("Error: %v", err)
			return
		}
		fmt.Printf("encrypt: %s =>\n%s\n", arg, s)
	}
}
