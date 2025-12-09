package day_six

func PartOne(mathGrid *MathGrid) int {
	sum := 0

	for i := range mathGrid.Cols {
		result := mathGrid.SolveColumn(i)
		sum += result
	}

	return sum
}
