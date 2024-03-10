package vector

type Vector struct {
	X float64
	Y float64
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 Vector) Multiply(v2 Vector) Vector {
	return Vector{v1.X * v2.X, v1.Y * v2.Y}
}
