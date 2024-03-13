package tests

import (
	"math"
	"testing"

	"github.com/Raadwal/boids-simulation/internal/vector"
)

func TestVectorDivideByScalar_Positive(t *testing.T) {
	v := vector.Vector{X: 3, Y: 6}
	scalar := 3.0

	v.DivideByScalar(scalar)
	expected := vector.Vector{X: 1, Y: 2}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestVectorDivideByScalar_Negative(t *testing.T) {
	v := vector.Vector{X: 3, Y: 6}
	scalar := -3.0

	v.DivideByScalar(scalar)
	expected := vector.Vector{X: -1, Y: -2}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestVectorDivideByScalar_Zero(t *testing.T) {
	v := vector.Vector{X: 3, Y: 6}
	scalar := 0.0

	v.DivideByScalar(scalar)

	if !math.IsInf(v.X, 0) || !math.IsInf(v.Y, 0) {
		t.Errorf("Expected {+Inf +Inf}, got %v", v)
	}
}
