package tests

import (
	"testing"

	"github.com/Raadwal/boids-simulation/internal/vector"
)

func TestVectorMultiply_PositiveCoordinates(t *testing.T) {
	v1 := vector.Vector{X: 3, Y: 4}
	v2 := vector.Vector{X: 5, Y: 6}

	v1.Multiply(v2)
	expected := vector.Vector{X: 15, Y: 24}

	if v1 != expected {
		t.Errorf("Expected %v, got %v", expected, v1)
	}
}

func TestVectorMultiply_NegativeCoordinates(t *testing.T) {
	v1 := vector.Vector{X: -2, Y: 5}
	v2 := vector.Vector{X: 3, Y: -1}

	v1.Multiply(v2)
	expected := vector.Vector{X: -6, Y: -5}

	if v1 != expected {
		t.Errorf("Expected %v, got %v", expected, v1)
	}
}

func TestVectorMultiply_ZeroVector(t *testing.T) {
	v1 := vector.Vector{X: 2, Y: 3}
	v2 := vector.Vector{X: 0, Y: 0}

	v1.Multiply(v2)
	expected := vector.Vector{X: 0, Y: 0}

	if v1 != expected {
		t.Errorf("Expected %v, got %v", expected, v1)
	}
}

func TestVectorMultiply_Self(t *testing.T) {
	v := vector.Vector{X: 7, Y: -2}

	v.Multiply(v)
	expected := vector.Vector{X: 49, Y: 4}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}
