// +build ignore
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ringohub/hello-go/src/ch2/popcount"
)

func main() {
	var num uint64
	if len(os.Args[1:]) == 0 {
		num = 0
	} else {
		v, _ := strconv.ParseUint(os.Args[1], 10, 64)
		num = v
	}
	fmt.Println(num)
	fmt.Println(popcount.PopCount(num))
}
