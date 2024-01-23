package src2023

import (
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

func transpose(grid map[int]map[utils.Coordinates]string, maxRows, maxCols int) (map[int]map[utils.Coordinates]string, int, int) {
	transposedGrid := make(map[int]map[utils.Coordinates]string)

	for old := 0; old < maxRows; old++ {
		rowElements := grid[old]

		for coord, char := range rowElements {
			new := coord.Col
			newCoord := utils.Coordinates{
				Row: coord.Col,
				Col: coord.Row,
			}

			utils.Logger.Debugln(coord, newCoord, new, char)

			if _, found := transposedGrid[new]; !found {
				transposedGrid[new] = make(map[utils.Coordinates]string)
			}

			transposedGrid[new][newCoord] = char
		}
	}

	return transposedGrid, maxCols, maxRows
}

func go_up(grid map[int]map[utils.Coordinates]string) {
	for row := 0; row < len(grid); row++ {
		rowElements := grid[row]
		for coord, char := range rowElements {
			if char == "#" {
				continue
			}

			for r := row - 1; r >= 0; r-- {
				shiftedCoord := utils.Coordinates{
					Row: r,
					Col: coord.Col,
				}

				if grid[r][shiftedCoord] == "#" || grid[r][shiftedCoord] == "O" {
					utils.Logger.Debugf("Start: %+v R: %d Found: %s at %+v", coord, r, grid[r][shiftedCoord], shiftedCoord)
					delete(grid[row], coord)
					grid[r+1][utils.Coordinates{
						Row: r + 1,
						Col: coord.Col,
					}] = "O"

					// utils.Logger.Debugf("%+v", grid[r+1])
					break
				} else if r == 0 {
					utils.Logger.Debugf("Start: %+v Hit top at %+v", coord, shiftedCoord)
					delete(grid[row], coord)
					grid[r][utils.Coordinates{
						Row: r,
						Col: coord.Col,
					}] = "O"
				}
			}
		}
	}
}

func go_down(grid map[int]map[utils.Coordinates]string, maxRows int) {
	for row := maxRows - 1; row >= 0; row-- {
		rowElements := grid[row]
		for coord, char := range rowElements {
			if char == "#" {
				continue
			}

			for r := row + 1; r < maxRows; r++ {
				shiftedCoord := utils.Coordinates{
					Row: r,
					Col: coord.Col,
				}

				if grid[r][shiftedCoord] == "#" || grid[r][shiftedCoord] == "O" {
					utils.Logger.Debugf("Start: %+v R: %d/%d Found: %s at %+v", coord, r, maxRows, grid[r][shiftedCoord], shiftedCoord)
					delete(grid[row], coord)

					newCoord := utils.Coordinates{
						Row: r - 1,
						Col: coord.Col,
					}

					if _, found := grid[r-1]; !found {
						grid[r-1] = make(map[utils.Coordinates]string)
					}
					grid[r-1][newCoord] = "O"

					// utils.Logger.Debugf("%+v", grid[r+1])
					break
				} else if r == maxRows-1 {
					utils.Logger.Debugf("Start: %+v Hit top at %+v", coord, shiftedCoord)
					delete(grid[row], coord)
					grid[r][utils.Coordinates{
						Row: r,
						Col: coord.Col,
					}] = "O"
				}
			}
		}
	}
}

func printGrid(grid map[int]map[utils.Coordinates]string, maxRows, maxCols int) string {

	rowS := "\n"
	for row := 0; row < maxRows; row++ {

		rowElements := grid[row]

		for col := 0; col < maxCols; col++ {
			coord := utils.Coordinates{
				Row: row,
				Col: col,
			}
			if char, found := rowElements[coord]; found {
				rowS += char
			} else {
				rowS += "."
			}

		}
		rowS += "\n"
	}
	return rowS
}

func (s Solver2023) Day_14(part int, reader utils.Reader) int {
	row := 0

	grid := make(map[int]map[utils.Coordinates]string)
	maxCols := 0

	for line := range reader.Stream() {
		maxCols = len(strings.Split(line, ""))
		for col, char := range strings.Split(line, "") {
			if char != "." {
				coord := utils.Coordinates{
					Row: row,
					Col: col,
				}

				if _, found := grid[row]; !found {
					grid[row] = make(map[utils.Coordinates]string)
				}

				grid[row][coord] = char
			}
		}
		row++
	}

	maxRows := row

	// utils.Logger.Infof("Grid: %+v", grid)
	// for row := 0; row < len(grid); row++ {
	// 	utils.Logger.Infof("%d -> %v", row, grid[row])
	// }
	strGrid := printGrid(grid, maxRows, maxCols)
	utils.Logger.Infof(strGrid)

	calculateLoad := func() int {
		load := 0
		for row := 0; row < len(grid); row++ {
			rowElements := grid[row]
			for _, char := range rowElements {
				if char != "#" {
					load += maxRows - row
				}
			}
		}
		return load
	}

	switch part {
	case 1:
		go_up(grid)
		return calculateLoad()
	case 2:

		initGrid := make(map[int]map[utils.Coordinates]string)
		for row, v := range grid {
			if _, found := initGrid[row]; !found {
				initGrid[row] = make(map[utils.Coordinates]string)
			}
			initGrid[row] = v
		}

		states := make(map[string][]string)
		dirs := []string{"u", "w", "d", "e"}

		for i := 0; i < 1000000000; i++ {

			for _, dir := range dirs {
				switch dir {
				case "u":
					go_up(grid)
				case "w":
					grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
					go_up(grid)
					grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
				case "d":
					go_down(grid, maxRows)
				case "e":
					grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
					go_down(grid, maxRows)
					grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
				}

				strGrid := printGrid(grid, maxRows, maxCols)
				if details, found := states[strGrid]; found && dir == details[0] && dir == "e" {
					utils.Logger.Infof("%+v %d %v %v", strGrid, i, found, details)

					// on example 9 == 2 on east -> loop around with a step of 9-2=7
					// 0123456789

					// 10 -- 3
					// 11 -- 4
					// 12 -- 5
					// 13 -- 6
					// 14 -- 7
					// 15 -- 8
					// 16 -- 9
					// 17 -- 3
					// 18 -- 4
					// 19 -- 5
					foundIndex, _ := strconv.Atoi(details[1])
					circle := foundIndex - i
					shiftedIndex := foundIndex + (1000000000-foundIndex)%circle

					utils.Logger.Infof("%d is similar to %s - 1000000000 corresponds to %d", i, details[1], shiftedIndex)
					grid = initGrid
					for i := 0; i < shiftedIndex; i++ {

						for _, dir := range dirs {
							switch dir {
							case "u":
								go_up(grid)
							case "w":
								grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
								go_up(grid)
								grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
							case "d":
								go_down(grid, maxRows)
							case "e":
								grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
								go_down(grid, maxRows)
								grid, maxRows, maxCols = transpose(grid, maxRows, maxCols)
							}
						}
					}
					return calculateLoad()

				}
				states[printGrid(grid, maxRows, maxCols)] = []string{dir, strconv.Itoa(i)}
			}
		}
	default:
		//shouldn't reach here
		return -1
	}

	return -1

}
