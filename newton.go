package main

import (
	"fmt"
	"math"
)

const Epsilon float64 = 0.000000000000001

func compute(x float64, z float64) float64 {
	zz := z - (z*z-x)/(2*z)
	if math.Abs(zz-z) < Epsilon {
		return zz
	}
	return compute(x, zz)
}

func Sqrt(x float64) float64 {
	return compute(x, 1.0)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
