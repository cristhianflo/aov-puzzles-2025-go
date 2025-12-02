package main

import (
	"cristhianflo/aov-puzzles-2025/internal"
	"cristhianflo/aov-puzzles-2025/internal/day_one"
	"fmt"
	"log"
	"strconv"
)

func main() {
	grid, err := internal.LoadGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	partOne(grid)
}

func partOne(grid *internal.Grid) {
	dial := day_one.NewDial(100)
	dial.Set(50)
	password := 0
	for r := 0; r < grid.Rows; r++ {
		row := grid.Row(r)

		direction, distance, err := ParseStep(string(row))
		if err != nil {
			log.Println(err)
		}

		if string(direction) == "L" {
			distance = distance * -1
		}

		dial.Rotate(distance)

		if dial.Value() == 0 {
			password++
		}
	}

	fmt.Printf("The password is: %d\n", password)
}

func ParseStep(s string) (rune, int, error) {
	if len(s) < 2 {
		return 0, 0, fmt.Errorf("invalid step: %s", s)
	}

	dir := rune(s[0])
	num, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, 0, err
	}

	return dir, num, nil
}
