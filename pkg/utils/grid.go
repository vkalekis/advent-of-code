package utils

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
)

type Row []string
type Col []string
type Grid [][]string

func GenerateGrid(reader Reader) Grid {
	var g Grid

	for line := range reader.Stream() {
		logger.Debugln(line)
		g = append(g, strings.Split(line, ""))
	}
	return g
}

func (g Grid) ExtractRow(i int) Row {
	if i >= 0 && i < len(g) {
		return g[i]
	}
	return nil
}

func (g Grid) ExtractRows() Grid {
	var rows Grid

	for i := 0; i < len(g); i++ {
		rows = append(rows, g.ExtractRow(i))
	}

	return rows
}

func (g Grid) ExtractColumn(j int) Col {
	var column Col
	for _, row := range g {
		if j >= 0 && j < len(row) {
			column = append(column, row[j])
		}
	}
	return column
}

func (g Grid) ExtractColumns() Grid {
	var columns Grid

	for j := 0; j < len(g[0]); j++ {
		columns = append(columns, g.ExtractColumn(j))
	}

	return columns
}

// func (g Grid) ReverseRows() Array {
// 	for i := range g {
// 		reverse(g[i])
// 	}
// 	return g
// }

func Reverse(sl []string) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func (g Grid) ExtractMainDiagonals() Grid {
	rows := len(g)
	cols := len(g[0])
	var diagonals Grid

	// starting from the first column
	for startRow := 0; startRow < rows; startRow++ {
		var diag Row
		for i, j := startRow, 0; i < rows && j < cols; i, j = i+1, j+1 {
			diag = append(diag, g[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	// starting from the first row
	for startCol := 1; startCol < cols; startCol++ {
		var diag Row
		for i, j := 0, startCol; i < rows && j < cols; i, j = i+1, j+1 {
			diag = append(diag, g[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	return diagonals
}

func (g Grid) ExtractAntiDiagonals() Grid {
	rows := len(g)
	cols := len(g[0])
	var diagonals Grid

	// starting from the last column
	for startRow := 0; startRow < rows; startRow++ {
		var diag Row
		for i, j := startRow, cols-1; i < rows && j >= 0; i, j = i+1, j-1 {
			diag = append(diag, g[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	// starting from the first row (top side)
	for startCol := cols - 2; startCol >= 0; startCol-- {
		var diag Row
		for i, j := 0, startCol; i < rows && j >= 0; i, j = i+1, j-1 {
			diag = append(diag, g[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	return diagonals
}
