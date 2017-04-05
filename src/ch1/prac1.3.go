// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	normal()
	end := time.Now()
	fmt.Printf("normal:\t%f\n", end.Sub(start).Seconds())

	start = time.Now()
	join()
	end = time.Now()
	fmt.Printf("join:\t%f\n", end.Sub(start).Seconds())
}

func normal() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func join() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
