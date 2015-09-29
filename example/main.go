package main

import (
	"fmt"

	"github.com/shell909090/cryptpass"
)

var pass1 = "QAcP07j3AA=="
var pass2 = "xd3BFUoV"

func main() {
	t, _ := cryptpass.DecryptPass(pass1)
	fmt.Println(t)

	t, _ = cryptpass.DecryptPass(pass2)
	fmt.Println(t)
}
