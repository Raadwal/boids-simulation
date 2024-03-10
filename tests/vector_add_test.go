package tests

import (
	"testing"

	"github.com/Raadwal/boids-simulation/internal/vector"
)

func TestVectorAdd_PositiveCoordinates(t *testing.T) {
	v1 := vector.Vector{X: 3, Y: 4}
	v2 := vector.Vector{X: 5, Y: 6}

	result := v1.Add(v2)
	expected := vector.Vector{X: 8, Y: 10}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestVectorAdd_NegativeCoordinates(t *testing.T) {
	v1 := vector.Vector{X: -2, Y: 5}
	v2 := vector.Vector{X: 3, Y: -1}

	result := v1.Add(v2)
	expected := vector.Vector{X: 1, Y: 4}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestVectorAdd_ZeroVectors(t *testing.T) {
	zeroVector := vector.Vector{X: 0, Y: 0}
	v := vector.Vector{X: 2, Y: 3}

	result := v.Add(zeroVector)
	expected := v

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestVectorAdd_IdenticalVectors(t *testing.T) {
	v := vector.Vector{X: 7, Y: -2}

	result := v.Add(v)
	expected := vector.Vector{X: 14, Y: -4}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
