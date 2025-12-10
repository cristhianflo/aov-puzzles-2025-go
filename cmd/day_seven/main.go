package main

import (
	"cristhianflo/aov-puzzles-2025/internal/day_seven"
	"fmt"
	"log"
)

func main() {
	beamGrid, err := day_seven.NewBeamGrid("input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	splits := day_seven.PartOne(beamGrid)

	fmt.Printf("The starting point is: %v\n", beamGrid.StartingPoint)

	fmt.Printf("The splits of the beam grid is: %d\n", splits)

	// for i := range beamGrid.Rows {
	// 	row := beamGrid.Row(i)
	// 	fmt.Println(string(row))
	// }
}
