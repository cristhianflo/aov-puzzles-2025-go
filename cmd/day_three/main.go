package main

import (
	"cristhianflo/aov-puzzles-2025/internal"
	"cristhianflo/aov-puzzles-2025/internal/day_three"
	"fmt"
	"log"
)

func main() {
	grid, err := internal.LoadGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	joltage := day_three.PartOne(grid)

	fmt.Printf("The total output joltage is: %d\n", joltage)
}
