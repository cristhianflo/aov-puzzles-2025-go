package day_four

import (
	"cristhianflo/aov-puzzles-2025/internal"
)

func PartOne(grid *internal.Grid) int {
	availableRolls := 0

	for row := range grid.Rows {
		for col := range grid.Row(row) {
			data := grid.Get(row, col)
			if string(data) != "@" {
				continue
			}
			rollCounter := 0
			neighbours := grid.Neighbors8(row, col)

			for _, neighbour := range neighbours {
				rune := grid.Get(neighbour[0], neighbour[1])
				if string(rune) == "@" {
					rollCounter++
				}
			}
			if rollCounter < 4 {

				availableRolls++
			}
		}

	}
	return availableRolls
}

func PartTwo(grid *internal.Grid) int {
	availableRolls := 0
	removedRolls := -1
	var removableRolls [][2]int

	for removedRolls != 0 {
		removedRolls = 0
		for row := range grid.Rows {
			for col := range grid.Row(row) {
				data := grid.Get(row, col)
				if string(data) != "@" {
					continue
				}
				rollCounter := 0
				neighbours := grid.Neighbors8(row, col)

				for _, neighbour := range neighbours {
					rune := grid.Get(neighbour[0], neighbour[1])
					if string(rune) == "@" {
						rollCounter++
					}
				}
				if rollCounter < 4 {
					removableRolls = append(removableRolls, [2]int{row, col})
				}
			}
		}

		for _, removeCoords := range removableRolls {
			grid.Set(removeCoords[0], removeCoords[1], 'x')
			removedRolls++
		}
		availableRolls += len(removableRolls)
		removableRolls = removableRolls[:0]
	}
	return availableRolls
}
