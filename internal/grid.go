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
