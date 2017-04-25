package distconv

type Meter float64
type Foot float64

const (
	oneMeterFoot Meter = 0.3048
)

func MToF(v Meter) Foot {
	return Foot(v * oneMeterFoot)
}
func FToM(v Foot) Meter {
	return Meter(v) / oneMeterFoot
}
