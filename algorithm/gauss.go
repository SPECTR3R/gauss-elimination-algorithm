package gauss

// Solve solves the system of linear equations using the gauss Elimination algorithm with scaled partial pivoting
func SolveSystem(a0 [][]float64, b0 []float64) ([]float64, error) {
	a := augmentedMatrix(a0, b0)

	for k := range a {
		if i, err := computeScaleFactor(a, k, len(a)); err != nil {
			return nil, err
		} else if i != k {
			a = swapRows(a, k, i)
		}
		a = gaussianElimination(a, k, len(a))
	}

	return backSubstitution(a, len(a)), nil
}
