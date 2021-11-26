package gauss

func augmentedMatrix(a [][]float64, b []float64) [][]float64 {
	// make augmented matrix
	m := len(b)
	ab := make([][]float64, m)

	for i, ai := range a {
		row := make([]float64, m+1)
		copy(row, ai)
		row[m] = b[i]
		ab[i] = row
	}
	return ab
}
