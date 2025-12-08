package main

import (
	"cristhianflo/aov-puzzles-2025/internal"
	"cristhianflo/aov-puzzles-2025/internal/day_four"
	"fmt"
	"log"
)

func main() {
	grid, err := internal.LoadGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rolls := day_four.PartTwo(grid)

	fmt.Printf("The total rolls of paper are: %d\n", rolls)
}
