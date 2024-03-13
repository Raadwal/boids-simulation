package tests

import (
	"testing"

	"github.com/Raadwal/boids-simulation/internal/vector"
)

func TestVectorSubtract_PositiveCoordinates(t *testing.T) {
	v1 := vector.Vector{X: 3, Y: 4}
	v2 := vector.Vector{X: 5, Y: 6}

	v1.Subtract(v2)
	expected := vector.Vector{X: -2, Y: -2}

	if v1 != expected {
		t.Errorf("Expected %v, got %v", expected, v1)
	}
}

func TestVectorSubtract_NegativeCoordinates(t *testing.T) {
	v1 := vector.Vector{X: -2, Y: 5}
	v2 := vector.Vector{X: 3, Y: -1}

	v1.Subtract(v2)
	expected := vector.Vector{X: -5, Y: 6}

	if v1 != expected {
		t.Errorf("Expected %v, got %v", expected, v1)
	}
}

func TestVectorSubtract_ZeroVector(t *testing.T) {
	v1 := vector.Vector{X: 2, Y: 3}
	v2 := vector.Vector{X: 0, Y: 0}

	v1.Subtract(v2)
	expected := vector.Vector{X: 2, Y: 3}

	if v1 != expected {
		t.Errorf("Expected %v, got %v", expected, v1)
	}
}

func TestVectorSubtract_Self(t *testing.T) {
	v := vector.Vector{X: 7, Y: -2}

	v.Subtract(v)
	expected := vector.Vector{X: 0, Y: 0}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}
