package utils

import "fmt"

// misc helper structs/functions

const (
	Up = iota
	Down
	Left
	Right
)

type Coordinates struct {
	Row, Col int
}

func NewCoordinates(row, col int) Coordinates {
	return Coordinates{
		Row: row,
		Col: col,
	}
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d,%d)", c.Row, c.Col)
}

func (c Coordinates) IsValid(maxRows, maxCols int) bool {
	if c.Row < 0 || c.Row >= maxRows {
		return false
	}
	if c.Col < 0 || c.Col >= maxCols {
		return false
	}
	return true
}

func (c Coordinates) GetNeighbors(maxRows, maxCols int, onlyCardinal bool) []Coordinates {
	neighbors := make([]Coordinates, 0)
	if !c.IsValid(maxRows, maxCols) {
		return neighbors
	}

	if c.Row > 0 {
		neighbors = append(neighbors, Coordinates{
			Row: c.Row - 1,
			Col: c.Col,
		})
	}
	if c.Col > 0 {
		neighbors = append(neighbors, Coordinates{
			Row: c.Row,
			Col: c.Col - 1,
		})
	}
	if c.Row < maxRows-1 {
		neighbors = append(neighbors, Coordinates{
			Row: c.Row + 1,
			Col: c.Col,
		})
	}
	if c.Col < maxCols-1 {
		neighbors = append(neighbors, Coordinates{
			Row: c.Row,
			Col: c.Col + 1,
		})
	}

	if !onlyCardinal {
		if c.Row > 0 && c.Col > 0 {
			neighbors = append(neighbors, Coordinates{
				Row: c.Row - 1,
				Col: c.Col - 1,
			})
		}
		if c.Row > 0 && c.Col < maxCols-1 {
			neighbors = append(neighbors, Coordinates{
				Row: c.Row - 1,
				Col: c.Col + 1,
			})
		}
		if c.Row < maxRows-1 && c.Col > 0 {
			neighbors = append(neighbors, Coordinates{
				Row: c.Row + 1,
				Col: c.Col - 1,
			})
		}
		if c.Row < maxRows-1 && c.Col < maxCols-1 {
			neighbors = append(neighbors, Coordinates{
				Row: c.Row + 1,
				Col: c.Col + 1,
			})
		}
	}

	return neighbors
}

func (c Coordinates) Shift(dir string, steps int) Coordinates {
	switch dir {
	case "n":
		return Coordinates{c.Row - steps, c.Col}
	case "s":
		return Coordinates{c.Row + steps, c.Col}
	case "w":
		return Coordinates{c.Row, c.Col - steps}
	case "e":
		return Coordinates{c.Row, c.Col + steps}
	default:
		return c
	}
}

func areEqual(c1, c2 Coordinates) bool {
	return c1.Row == c2.Row && c1.Col == c2.Col
}

func GetIntermediateCoords(start, end Coordinates) []Coordinates {
	intermediate := make([]Coordinates, 0)
	if start.Row == end.Row {
		if start.Col > end.Col {
			// move left
			for col := start.Col; col >= end.Col; col-- {
				intermediate = append(intermediate, Coordinates{Row: start.Row, Col: col})
			}
		} else {
			// move right
			for col := start.Col; col <= end.Col; col++ {
				intermediate = append(intermediate, Coordinates{Row: start.Row, Col: col})
			}
		}
	} else if start.Col == end.Col {
		if start.Row > end.Row {
			// move up
			for row := start.Row; row >= end.Row; row-- {
				intermediate = append(intermediate, Coordinates{Row: row, Col: start.Col})
			}
		} else {
			// move down
			for row := start.Row; row <= end.Row; row++ {
				intermediate = append(intermediate, Coordinates{Row: row, Col: start.Col})
			}
		}
	}
	return intermediate
}

func RemoveDuplicates(coords []Coordinates) []Coordinates {
	seen := make(map[Coordinates]bool)
	var unique []Coordinates

	for _, coord := range coords {
		if !seen[coord] {
			seen[coord] = true
			unique = append(unique, coord)
		}
	}

	return unique
}
