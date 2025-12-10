package day_seven

import "fmt"

func PartOne(bg *BeamGrid) int {
	splitAmount := 0
	currentRow := bg.StartingPoint[0]
	beams := [][2]int{bg.StartingPoint}
	for currentRow < bg.Rows {
		var newBeams [][2]int
		for _, beam := range beams {
			beamShot := [2]int{currentRow, beam[1]}
			if bg.Get(beamShot[0], beamShot[1]) == bg.SplitterSymbol {
				fmt.Printf("Splitted at: %v\n", beamShot)
				splittedBeams := bg.splitBeam(beamShot)
				newBeams = append(newBeams, splittedBeams...)
				splitAmount++
			} else if bg.Get(beamShot[0], beamShot[1]) != '|' {
				newBeams = append(newBeams, beamShot)
				bg.Set(beamShot[0], beamShot[1], '|')
			}
		}
		if len(newBeams) > 0 {
			beams = newBeams
		}
		currentRow++
	}
	return splitAmount
}
