package main

import (
	"fmt"

	"github.com/shell909090/cryptpass"
)

func main() {
	s, _ := cryptpass.EncryptPass("abc")
	fmt.Println(s)
	t, _ := cryptpass.DecryptPass(s)
	fmt.Println(t)

	s, _ = cryptpass.EncryptPass("中文")
	fmt.Println(s)
	t, _ = cryptpass.DecryptPass(s)
	fmt.Println(t)
}
