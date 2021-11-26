package gauss

import (
	"errors"
	"math"
)

func augmentedMatrix(a0 [][]float64, b0 []float64) [][]float64 {
	// make augmented matrix
	m := len(b0)
	a := make([][]float64, m)

	for i, ai := range a0 {
		row := make([]float64, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	return a
}

func computeScaleFactor(a [][]float64, k, m int) (int, error) {
	max := -1.0
	iMax := 0

	for i := k; i < m; i++ {
		row := a[i]
		// compute scale factor s = max abs in row
		s := -1.
		for j := k; j < m; j++ {
			x := math.Abs(row[j])
			if x > s {
				s = x
			}
		}
		// scale the abs used to pick the pivot.
		if abs := math.Abs(row[k]) / s; abs > max {
			iMax = i
			max = abs
		}
	}

	if a[iMax][k] == 0 {
		return 0, errors.New("matrix is unsolvable")
	}

	return iMax, nil
}
