package main

import (
	"cristhianflo/aov-puzzles-2025/internal/day_six"
	"fmt"
	"log"
)

func main() {
	mathGrid, err := day_six.NewMathGrid("input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	sum := day_six.PartOne(mathGrid)

	fmt.Printf("The sum of the math grid is: %d\n", sum)
}
