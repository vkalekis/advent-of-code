package utils

type Row []string
type Col []string
type Array [][]string

func (a Array) ExtractRow(i int) Row {
	if i >= 0 && i < len(a) {
		return a[i]
	}
	return nil
}

func (a Array) ExtractRows() Array {
	var rows Array

	for i := 0; i < len(a); i++ {
		rows = append(rows, a.ExtractRow(i))
	}

	return rows
}

func (a Array) ExtractColumn(j int) Col {
	var column Col
	for _, row := range a {
		if j >= 0 && j < len(row) {
			column = append(column, row[j])
		}
	}
	return column
}

func (a Array) ExtractColumns() Array {
	var columns Array

	for j := 0; j < len(a[0]); j++ {
		columns = append(columns, a.ExtractColumn(j))
	}

	return columns
}

// func (a Array) ReverseRows() Array {
// 	for i := range a {
// 		reverse(a[i])
// 	}
// 	return a
// }

func Reverse(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func (a Array) ExtractMainDiagonals() Array {
	rows := len(a)
	cols := len(a[0])
	var diagonals Array

	// starting from the first column
	for startRow := 0; startRow < rows; startRow++ {
		var diag Row
		for i, j := startRow, 0; i < rows && j < cols; i, j = i+1, j+1 {
			diag = append(diag, a[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	// starting from the first row
	for startCol := 1; startCol < cols; startCol++ {
		var diag Row
		for i, j := 0, startCol; i < rows && j < cols; i, j = i+1, j+1 {
			diag = append(diag, a[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	return diagonals
}

func (a Array) ExtractAntiDiagonals() Array {
	rows := len(a)
	cols := len(a[0])
	var diagonals Array

	// starting from the last column
	for startRow := 0; startRow < rows; startRow++ {
		var diag Row
		for i, j := startRow, cols-1; i < rows && j >= 0; i, j = i+1, j-1 {
			diag = append(diag, a[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	// starting from the first row (top side)
	for startCol := cols - 2; startCol >= 0; startCol-- {
		var diag Row
		for i, j := 0, startCol; i < rows && j >= 0; i, j = i+1, j-1 {
			diag = append(diag, a[i][j])
		}
		diagonals = append(diagonals, diag)
	}

	return diagonals
}
