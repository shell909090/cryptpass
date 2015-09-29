package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/shell909090/cryptpass"
)

func main() {
	var PassPath string
	flag.StringVar(&PassPath, "pass", "", "path of passkey")
	flag.Parse()

	if PassPath != "" {
		cryptpass.PassPath = PassPath
	}

	stdin := bufio.NewReader(os.Stdin)
	for {
		line, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error: %v", err)
			return
		}

		line = strings.TrimRight(line, "\n")
		s, err := cryptpass.EncryptPass(line)
		if err != nil {
			fmt.Println("Error: %v", err)
			return
		}

		fmt.Printf("%s\n", s)
	}
}
