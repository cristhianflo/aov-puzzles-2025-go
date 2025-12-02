package main

import (
	"cristhianflo/aov-puzzles-2025/internal"
	"cristhianflo/aov-puzzles-2025/internal/day_one"
	"fmt"
	"log"
)

func main() {
	grid, err := internal.LoadGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	password := day_one.PartTwo(grid)

	fmt.Printf("The password is: %d\n", password)
}
