package tests

import (
	"testing"

	"github.com/Raadwal/boids-simulation/internal/vector"
)

func TestVectorMultiply_PositiveCoordinates(t *testing.T) {
	v1 := vector.Vector{X: 3, Y: 4}
	v2 := vector.Vector{X: 5, Y: 6}

	result := v1.Multiply(v2)
	expected := vector.Vector{X: 15, Y: 24}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestVectorMultiply_NegativeCoordinates(t *testing.T) {
	v1 := vector.Vector{X: -2, Y: 5}
	v2 := vector.Vector{X: 3, Y: -1}

	result := v1.Multiply(v2)
	expected := vector.Vector{X: -6, Y: -5}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestVectorMultiply_ZeroVectors(t *testing.T) {
	zeroVector := vector.Vector{X: 0, Y: 0}
	v := vector.Vector{X: 2, Y: 3}

	result := v.Multiply(zeroVector)
	expected := zeroVector

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestVectorMultiply_IdenticalVectors(t *testing.T) {
	v := vector.Vector{X: 7, Y: -2}

	result := v.Multiply(v)
	expected := vector.Vector{X: 49, Y: 4}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
