package src2025

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

const (
	toiletPaperSymbol = "@"
	emptySymbol       = "."
)

func findAccessableToiletPapers(grid [][]string) []utils.Coordinates {
	accessableToiletPapers := make([]utils.Coordinates, 0)

	rows, cols := len(grid), len(grid[0])
	for row := range grid {
		for col := range grid {
			if grid[row][col] != toiletPaperSymbol {
				continue
			}

			coord := utils.Coordinates{
				Row: row, Col: col,
			}
			//logger.Debugf("Coord: %+v Neighbors: %+v", coord, coord.GetNeighbors(rows, cols, false))

			var toiletPaperNeighbors int
			for _, neighbor := range coord.GetNeighbors(rows, cols, false) {
				if grid[neighbor.Row][neighbor.Col] == toiletPaperSymbol {
					toiletPaperNeighbors++
				}
			}
			if toiletPaperNeighbors < 4 {
				accessableToiletPapers = append(accessableToiletPapers, coord)
			}
		}
	}

	return accessableToiletPapers
}

func printToiletPaperGrid(grid [][]string) {
	for row := range grid {
		logger.Debugf("%v", grid[row])
	}
}

func (s *Solver) Day_04(part int, reader utils.Reader) int {

	grid := make([][]string, 0)
	row := 0

	for line := range reader.Stream() {
		grid = append(grid, make([]string, len(line)))

		grid[row] = strings.Split(line, "")
		row++
	}

	logger.Debugf("Grid: %v", grid)

	switch part {
	case 1:
		return len(findAccessableToiletPapers(grid))
	case 2:

		var totalRemovedToiletPapers int
		var accessableToilerPapers []utils.Coordinates

		for true {

			//
			// printToiletPaperGrid(grid)
			//

			accessableToilerPapers = findAccessableToiletPapers(grid)
			if len(accessableToilerPapers) == 0 {
				break
			}

			logger.Debugf("Removing %d toilet paper(s)", len(accessableToilerPapers))
			totalRemovedToiletPapers += len(accessableToilerPapers)

			for _, tp := range accessableToilerPapers {
				grid[tp.Row][tp.Col] = emptySymbol
			}
		}
		return totalRemovedToiletPapers
	}

	return -1
}
