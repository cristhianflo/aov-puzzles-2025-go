package day_six

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

type CellType int

const (
	CellNumber    CellType = 0
	CellOperation CellType = 1
)

type Cell struct {
	Type     CellType
	Number   int
	Operator rune
}

type MathGrid struct {
	Rows int
	Cols int
	Data [][]Cell
}

func NewMathGrid(path string) (*MathGrid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]Cell

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text() + "\n")
		var row []Cell
		for i := 0; i < len(line); i++ {
			char := line[i]
			var cell Cell
			if unicode.IsSpace(char) {
				continue
			} else if unicode.IsDigit(char) {
				numString := ""

				for unicode.IsDigit(line[i]) && i < len(line)-1 {
					numString += string(line[i])
					i++
				}
				num, err := strconv.Atoi(numString)
				if err != nil {
					return nil, err
				}
				cell.Type = CellNumber
				cell.Number = num
			} else if char == '+' || char == '*' {
				cell.Type = CellOperation
				cell.Operator = char
			}
			row = append(row, cell)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	rows := len(data)
	cols := 0
	if rows > 0 {
		cols = len(data[0])
	}

	return &MathGrid{
		Rows: rows,
		Cols: cols,
		Data: data,
	}, nil
}

func (g *MathGrid) Col(c int) []Cell {
	if c < 0 && c > g.Cols {
		return nil
	}
	var dataCols []Cell

	for i := range g.Rows {
		col := g.Get(i, c)
		dataCols = append(dataCols, col)
	}

	return dataCols
}

func (g *MathGrid) Get(row, col int) Cell {
	return g.Data[row][col]
}

func (g *MathGrid) SolveColumn(c int) int {
	column := g.Col(c)
	cell := column[len(column)-1]
	result := column[0].Number
	for _, number := range column[1 : len(column)-1] {
		switch cell.Operator {
		case '+':
			result += number.Number
		case '*':
			result *= number.Number
		}
	}
	return result
}
