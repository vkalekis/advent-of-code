package utils

// misc helper structs/functions

type Coordinates struct {
	Row, Col int
}

func NewCoordinates(row, col int) Coordinates {
	return Coordinates{
		Row: row,
		Col: col,
	}
}

func (coords Coordinates) IsValid(maxRows, maxCols int) bool {
	if coords.Row < 0 || coords.Row >= maxRows {
		return false
	}
	if coords.Col < 0 || coords.Col >= maxCols {
		return false
	}
	return true
}

func (coords Coordinates) GetNeighbors(maxRows, maxCols int) []Coordinates {
	neighbors := make([]Coordinates, 0)
	if !coords.IsValid(maxRows, maxCols) {
		return neighbors
	}

	if coords.Row > 0 {
		neighbors = append(neighbors, Coordinates{
			Row: coords.Row - 1,
			Col: coords.Col,
		})
	}
	if coords.Col > 0 {
		neighbors = append(neighbors, Coordinates{
			Row: coords.Row,
			Col: coords.Col - 1,
		})
	}
	if coords.Row < maxRows-1 {
		neighbors = append(neighbors, Coordinates{
			Row: coords.Row + 1,
			Col: coords.Col,
		})
	}
	if coords.Col < maxCols-1 {
		neighbors = append(neighbors, Coordinates{
			Row: coords.Row,
			Col: coords.Col + 1,
		})
	}

	return neighbors
}

func (coords Coordinates) Shift(dir string, steps int) Coordinates {
	switch dir {
	case "n":
		return NewCoordinates(coords.Row-steps, coords.Col)
	case "s":
		return NewCoordinates(coords.Row+steps, coords.Col)
	case "w":
		return NewCoordinates(coords.Row, coords.Col-steps)
	case "e":
		return NewCoordinates(coords.Row, coords.Col+steps)
	default:
		return coords
	}
}
