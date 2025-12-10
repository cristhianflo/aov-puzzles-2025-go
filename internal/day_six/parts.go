package day_six

import "fmt"

func PartOne(mathGrid *MathGrid) int {
	sum := 0

	for i := range mathGrid.Cols {
		result := mathGrid.SolveColumn(i)
		sum += result
	}

	return sum
}

func PartTwo(mathGrid *MathGrid) int {
	sum := 0

	var numbers []int = make([]int, mathGrid.Rows-1)
	var operator Cell
	numAmount := 0
	fmt.Println(mathGrid.Col(0))
	for i := range mathGrid.Cols {
		col := mathGrid.Col(i)
		calculate := true
		for _, cell := range col {
			if col[0].Type == CellBlank && cell.Type == CellBlank {
				continue
			} else if cell.Type == CellNumber {
				numbers[numAmount] = numbers[numAmount]*10 + cell.Number
			} else if cell.Type == CellOperation {
				operator = cell
			}
			calculate = false
		}

		if numAmount < mathGrid.Rows {
			numAmount++
		}

		if calculate {
			result := numbers[0]
			switch operator.Operator {
			case '+':
				for _, number := range numbers[1:] {
					result += number
				}
			case '*':
				for _, number := range numbers[1:] {
					if number == 0 {
						continue
					}
					result *= number
				}
			}
			sum += result
			numAmount = 0
			clear(numbers)
		}
	}

	return sum
}
