package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, e := range strings.Fields(s) {
		m[e] = m[e] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
