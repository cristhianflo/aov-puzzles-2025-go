
func PartTwo(grid *internal.Grid) int {
	dial := NewDial(100)
	dial.Set(50)
	password := 0
	for r := 0; r < grid.Rows; r++ {
		row := grid.Row(r)

		direction, distance, err := ParseStep(string(row))

		if err != nil {
			log.Println(err)
		}

		rotations := dial.CalculateRotations(distance, string(direction))

		if string(direction) == "L" {
			distance = -(distance)
		}

		dial.Rotate(distance)

		password = password + int(rotations)
	}

	return password
}
