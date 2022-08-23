package algorithm

import (
	"errors"
	"math"
)

// augmentedMatrix creates an augmented matrix from the given matrix and vector.
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

// computeScaleFactor computes the scale factor S_i = |a_ki|/max(row) in each row of the given matrix, and returns the index of the row with the greatest scale factor.
func computeScaleFactor(a [][]float64, k, m int) (int, error) {
	max := -1.0
	iMax := 0

	for i := k; i < m; i++ {
		row := a[i]
		s := -1.

		for j := k; j < m; j++ {
			x := math.Abs(row[j])
			if x > s {
				s = x
			}
		}

		if abs := math.Abs(row[k]) / s; abs > max {
			iMax = i
			max = abs
		}
	}

	// return error if pivot value is == 0
	if a[iMax][k] == 0 {
		return 0, errors.New("matrix is unsolvable")
	}

	return iMax, nil
}

// swapRows swaps rows of given matrix using.
func swapRows(a [][]float64, idx1, idx2 int) [][]float64 {
	a[idx1], a[idx2] = a[idx2], a[idx1]
	return a
}

// gaussianElimination operates the elements of pivot k row with all the rows below it to create the zeroes of the row-eschelon form.
func gaussianElimination(a [][]float64, k, m int) [][]float64 {
	for i := k + 1; i < m; i++ {
		for j := k + 1; j <= m; j++ {
			a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
		}
		a[i][k] = 0
	}
	return a
}

// backSubstitution returns the solutions of the system with the row-eschelon form.
func backSubstitution(a [][]float64, m int) []float64 {
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x
}
