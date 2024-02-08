package utils

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// misc helper structs/functions

func ConstructGrid(reader Reader) (map[int][]string, int, int) {
	grid := make(map[int][]string)

	var maxRows, maxCols, row int

	for line := range reader.Stream() {

		splittedLine := strings.Split(line, "")

		grid[row] = splittedLine
		row++
	}

	maxRows = row
	maxCols = len(grid[0])
	return grid, maxRows, maxCols
}

func ConstructIntGrid(reader Reader) (map[int][]int, int, int) {
	grid := make(map[int][]int)

	var maxRows, maxCols, row int

	for line := range reader.Stream() {

		splittedLine := strings.Split(line, "")

		intLine := make([]int, 0)
		for _, ch := range splittedLine {
			intch, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			intLine = append(intLine, intch)
		}

		grid[row] = intLine
		row++
	}

	maxRows = row
	maxCols = len(grid[0])
	return grid, maxRows, maxCols
}

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

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func WriteMapToFile(data map[string]interface{}, filename string) error {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetCol(lines map[int][]string, col, maxRows int) []string {
	line := make([]string, 0)
	for r := 0; r < maxRows; r++ {
		line = append(line, lines[r][col])
	}
	return line
}
