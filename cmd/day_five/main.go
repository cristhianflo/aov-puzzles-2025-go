package main

import (
	"cristhianflo/aov-puzzles-2025/internal/day_five"
	"fmt"
	"log"
)

func main() {
	cafeteria, err := day_five.NewCafeteria("input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	freshIngredients := day_five.PartTwo(cafeteria)

	fmt.Printf("The amount of fresh ingredients in the cafeteria is: %d\n", freshIngredients)
}
