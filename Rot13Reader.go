package main

import (
	"io"
	"os"
	"strings"
//	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(p []byte) (n int, err error) {
	_, err = r13.r.Read(p)

	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') {
			p[i] = 'A' + p[i] % 13
		} else if (p[i] >= 'a' && p[i] <= 'z') {
			p[i] = 'a' + p[i] % 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
