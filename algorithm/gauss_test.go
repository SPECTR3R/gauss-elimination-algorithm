package gauss

import (
	"reflect"
	"testing"
)

func TestAugmentedMatrix(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	b := []float64{6, 7, 8}

	want := [][]float64{
		{1, 2, 3, 6},
		{2, 3, 4, 7},
		{3, 4, 5, 8},
	}

	got := augmentedMatrix(a, b)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted %v, \ngot %v", want, got)
	}
}
