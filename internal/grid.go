package internal

import (
	"bufio"
	"os"
)

type Grid struct {
	Rows int
	Cols int
	Data [][]rune
}

func LoadGrid(path string) (*Grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	rows := len(data)
	cols := 0
	if rows > 0 {
		cols = len(data[0])
	}

	return &Grid{
		Rows: rows,
		Cols: cols,
		Data: data,
	}, nil
}

func (g *Grid) Row(r int) []rune {
	return g.Data[r]
}

func (g *Grid) InBounds(row, col int) bool {
	return row >= 0 && row < g.Rows && col >= 0 && col < g.Cols
}

func (g *Grid) Get(row, col int) rune {
	return g.Data[row][col]
}

func (g *Grid) Neighbors8(row, col int) [][2]int {
	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	out := make([][2]int, 0, 8)
	for _, dir := range dirs {
		nr, nc := row+dir[0], col+dir[1]
		if g.InBounds(nr, nc) {
			out = append(out, [2]int{nr, nc})
		}
	}
	return out
}
