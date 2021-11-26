package algorithm

import (
	"reflect"
	"testing"
)

func TestAugmentedMatrix(t *testing.T) {
	a0 := [][]float64{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	b0 := []float64{6, 7, 8}

	want := [][]float64{
		{1, 2, 3, 6},
		{2, 3, 4, 7},
		{3, 4, 5, 8},
	}

	got := augmentedMatrix(a0, b0)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted %v, \ngot %v", want, got)
	}
}

func TestComputeScaleFactor(t *testing.T) {
	a := [][]float64{
		{2.11, -4.21, 0.921, 2.01},
		{4.01, 10.2, -1.12, -3.09},
		{1.09, .987, .832, 4.21},
	}

	want := 2

	got, err := computeScaleFactor(a, 0, 3)
	if err != nil {
		t.Fatalf("did not expect error, but got %v", err)
	}

	if got != want {
		t.Fatalf("\nWanted %v, \ngot %v", want, got)
	}
}

func TestSwapRows(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}

	want := [][]float64{
		{3, 4, 5},
		{2, 3, 4},
		{1, 2, 3},
	}

	a = swapRows(a, 0, 2)

	if !reflect.DeepEqual(a, want) {
		t.Fatalf("\nWanted %v, \ngot %v", want, a)
	}
}

func TestGaussianElimination(t *testing.T) {
	a := [][]float64{
		{1, 1, 1, 2},
		{1, 2, 3, 5},
		{2, 3, 4, 11},
	}

	want := [][]float64{
		{1, 1, 1, 2},
		{0, 1, 2, 3},
		{0, 1, 2, 7},
	}
	got := gaussianElimination(a, 0, 3)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted %v, \n   got %v", got, want)
	}
}

func TestBackSubstitution(t *testing.T) {
	a := [][]float64{
		{1, 0, 1, 0},
		{0, 4, -4, -4},
		{0, 0, -4, -4},
	}

	want := []float64{-1, 0, 1}
	got := backSubstitution(a, 3)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted %v, \n   got %v", got, want)
	}
}
