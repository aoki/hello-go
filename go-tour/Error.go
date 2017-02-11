package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %s", float64(e))
}

const Epsilon float64 = 0.000000000000001

func compute(x float64, z float64) float64 {
	zz := z - (z*z-x)/(2*z)
	if math.Abs(zz-z) < Epsilon {
		return zz
	}
	return compute(x, zz)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	return compute(f, 1.0), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
