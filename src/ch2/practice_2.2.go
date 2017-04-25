package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ringohub/hello-go/src/ch2/distconv"
	"github.com/ringohub/hello-go/src/ch2/tempconv"
	"github.com/ringohub/hello-go/src/ch2/weightconv"
)

var temperature = flag.Bool("t", false, "Temperature")
var weight = flag.Bool("w", false, "Weight")
var distance = flag.Bool("d", false, "Distance")

func main() {
	flag.Parse()
	fmt.Printf("%t, %t, %t\n", *temperature, *weight, *distance)
	var inputs []string
	if len(os.Args[2:]) == 0 {
		inputs = readStdIn()
	} else {
		inputs = os.Args[1:]
	}
	var values []float64
	for _, input := range inputs {
		v, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			fmt.Printf("Ignore this value \"%v\"\n", v)
		} else {
			values = append(values, v)
		}
	}

	fmt.Printf("%v, %d\n", inputs, len(inputs))
	fmt.Printf("%v, %d\n", values, len(values))

	if *weight {
		for _, v := range values {
			p := weightconv.Pound(v)
			fmt.Printf("%v lb, %v kg\n", p, weightconv.PToK(p))
		}
	} else if *distance {
		for _, v := range values {
			m := distconv.Meter(v)
			fmt.Printf("%v m, %v ft\n", m, distconv.MToF(m))
		}
	} else {
		for _, v := range values {
			c := tempconv.Celsius(v)
			fmt.Printf("%v, %v\n", c, tempconv.CToF(c))
		}
	}
}

// Read from Standard Input
func readStdIn() []string {
	input := bufio.NewScanner(os.Stdin)
	var inputs []string
	fmt.Println("Please input some numbers... (Press Ctrl+D to finish)", input.Text())
	for input.Scan() {
		inputs = append(inputs, input.Text())
	}
	return inputs
}
