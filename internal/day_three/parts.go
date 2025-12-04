package day_three

import (
	"cristhianflo/aov-puzzles-2025/internal"
	"math"
)

func PartOne(grid *internal.Grid) int {
	joltage := 0
	for r := 0; r < grid.Rows; r++ {
		row := grid.Row(r)

		largestJoltage := 0
		firstIndex := -1
		maxJoltage := 0
		for i := 0; i < len(row); i++ {
			col := row[i]

			if firstIndex < i {
				if maxJoltage < int(col-'0') {
					maxJoltage = int(col - '0')
					firstIndex = i
				}
			}

			if i == len(row)-2 && largestJoltage < 10 {
				largestJoltage += maxJoltage * 10
				maxJoltage = 0
				i = -1
				continue
			}

			if i == len(row)-1 {
				largestJoltage += maxJoltage
				break
			}
		}
		joltage += largestJoltage

	}

	return joltage
}

func PartTwo(grid *internal.Grid) int {
	joltage := 0
	for r := 0; r < grid.Rows; r++ {
		row := grid.Row(r)

		largestJoltage := 0
		lastIndex := -1
		batteriesLeft := 12
		maxJoltage := 0
		for i := 0; i < len(row); i++ {
			col := row[i]

			if batteriesLeft == 0 {
				break
			}

			if lastIndex < i {
				if maxJoltage < int(col-'0') {
					maxJoltage = int(col - '0')
					lastIndex = i
				}
			}

			if i == len(row)-batteriesLeft {
				largestJoltage += maxJoltage * int(math.Pow(10, float64(batteriesLeft-1)))
				batteriesLeft--
				maxJoltage = 0
				i = -1
				continue
			}
		}
		joltage += largestJoltage

	}

	return joltage
}
