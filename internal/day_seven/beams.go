package day_seven

import (
	"cristhianflo/aov-puzzles-2025/internal"
)

type BeamGrid struct {
	internal.Grid
	StartingPoint  [2]int
	SplitterSymbol rune
	SpaceSymbol    rune
}

func NewBeamGrid(path string) (*BeamGrid, error) {
	grid, err := internal.LoadGrid(path)
	if err != nil {
		return nil, err
	}

	var startingPoint [2]int

outer:
	for i := range grid.Rows {
		for j := range grid.Cols {
			cell := grid.Get(i, j)
			if cell == 'S' {
				startingPoint = [2]int{i, j}
				break outer
			}
		}
	}

	return &BeamGrid{
		Grid:           *grid,
		StartingPoint:  startingPoint,
		SplitterSymbol: '^',
		SpaceSymbol:    '.',
	}, nil
}

func (bg *BeamGrid) splitBeam(beam [2]int) [][2]int {
	beams := [][2]int{
		{beam[0], beam[1] - 1},
		{beam[0], beam[1] + 1},
	}

	var result [][2]int
	for _, newBeam := range beams {
		if bg.Get(newBeam[0], newBeam[1]) != bg.SpaceSymbol {
			continue
		}
		bg.Set(newBeam[0], newBeam[1], '|')
		result = append(result, newBeam)
	}

	return result
}
