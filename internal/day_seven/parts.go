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

func PartTwo(bg *BeamGrid) int {
	return ShootBeam(bg, bg.StartingPoint, 0)
}

func ShootBeam(bg *BeamGrid, beam [2]int, timelines int) int {
	total := timelines
	currentRow := beam[0] + 1

	if cache, ok := bg.BeamCache[beam]; ok {
		return cache
	}
	for currentRow < bg.Rows {
		currentBeam := [2]int{currentRow, beam[1]}
		if bg.Get(currentBeam[0], currentBeam[1]) == bg.SplitterSymbol {
			leftBeam := bg.BeamToLeft(currentBeam)
			rightBeam := bg.BeamToRight(currentBeam)
			leftResult := ShootBeam(bg, leftBeam, total)
			rightResult := ShootBeam(bg, rightBeam, total)
			bg.BeamCache[leftBeam] = leftResult
			bg.BeamCache[rightBeam] = rightResult
			total = leftResult + rightResult
			break
		} else {
			currentRow++
		}
	}

	if currentRow == bg.Rows {
		total++
	}
	return total
}
