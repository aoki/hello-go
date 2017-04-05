package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	cmdPath := strings.Split(os.Args[0], "/")
	fmt.Println(cmdPath[len(cmdPath)-1])
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
