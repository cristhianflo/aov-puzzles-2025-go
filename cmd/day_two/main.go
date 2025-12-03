package main

import (
	"cristhianflo/aov-puzzles-2025/internal/day_two"
	"fmt"
	"log"
)

func main() {
	rangeList, err := day_two.NewRangeList("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := day_two.PartTwo(rangeList)

	fmt.Printf("The sum of the invalid IDs is: %d\n", sum)
}
