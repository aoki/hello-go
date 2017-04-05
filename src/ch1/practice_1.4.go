package main

import (
	"bufio"
	"fmt"
	"os"
)

type Info struct {
	Count     int
	Filenames []string
}

func main() {
	counts := make(map[string]*Info)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
				f.Close()
			}
			countLines(f, counts)
			f.Close()
		}
	}
	fmt.Println("")
	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.Count, line, n.Filenames)
		}
	}
}

func countLines(f *os.File, counts map[string]*Info) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, exists := counts[key]
		if exists {
			counts[key].Count++
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		} else {
			counts[key] = new(Info)
			counts[key].Count = 1
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		}
	}
}
