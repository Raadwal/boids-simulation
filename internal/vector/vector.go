package vector

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v1 *Vector) Add(v2 Vector) {
	v1.X += v2.X
	v1.Y += v2.Y
}

func (v1 *Vector) Subtract(v2 Vector) {
	v1.X -= v2.X
	v1.Y -= v2.Y
}

func (v1 *Vector) Multiply(v2 Vector) {
	v1.X *= v2.X
	v1.Y *= v2.Y
}

func (v *Vector) MultiplyByScalar(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

func (v *Vector) DivideByScalar(scalar float64) {
	v.X /= scalar
	v.Y /= scalar
}

func (v *Vector) Clip(min, max float64) {
	v.X = math.Max(min, math.Min(max, v.X))
	v.Y = math.Max(min, math.Min(max, v.Y))
}

func (v *Vector) Normalize() {
	factor := math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
	v.X /= factor
	v.Y /= factor
}

func (v1 *Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v2.X-v1.X, 2) + math.Pow(v2.Y-v1.Y, 2))
}
