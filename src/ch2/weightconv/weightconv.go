// weightconv package interconvert between Pound and Kilogram.
package weightconv

type Pound float64
type Kilogram float64

const (
	oneKgPound Pound = 0.45359237
)

func PToK(v Pound) Kilogram {
	return Kilogram(v * oneKgPound)
}
func KToP(v Kilogram) Pound {
	return Pound(v) / oneKgPound
}
