package main

import (
	"fmt"

	"github.com/shell909090/cryptpass"
)

var pass1 = "QAcP07j3AA=="
var pass2 = "xd3BFUoV"
var pass3 = "abcdef"

func main() {
	cryptpass.PassPath = "cryptpass.key"
	fmt.Println(cryptpass.AutoPass(pass1))
	fmt.Println(cryptpass.AutoPass(pass2))
	fmt.Println(cryptpass.AutoPass(pass3))
}
