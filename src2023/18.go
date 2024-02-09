package src2023

import (
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

type line16 struct {
	dir   string
	steps int
}

func traverseLagoon(lines []line16) int {
	start := utils.NewCoordinates(0, 0)
	coords := make([]utils.Coordinates, 0)
	coords = append(coords, start)

	ind := 0

	for _, line := range lines {

		utils.Logger.Infof("Shifting: %+v %v", line, coords[ind])
		shiftedCoord := coords[ind].Shift(line.dir, line.steps)
		coords = append(coords, shiftedCoord)
		ind++
	}

	// using pick's theorem + shoelase
	// Pick: Area = Internal + boundary/2 - 1  -> A=I+B/2-1 -> I+B/2=A+1 -> I+B=A+B/2+1
	// Area is calculated using shoelace formula
	area := 0
	boundaries := 0

	for ind := 1; ind < len(coords); ind++ {
		factor := -coords[ind-1].Row*coords[ind].Col + coords[ind-1].Col*coords[ind].Row
		utils.Logger.Infof("%+v->%+v %d", coords[ind-1], coords[ind], factor)
		area += factor

		minRow, maxRow := coords[ind-1].Row, coords[ind].Row
		if coords[ind-1].Row > coords[ind].Row {
			minRow, maxRow = coords[ind].Row, coords[ind-1].Row
		}
		minCol, maxCol := coords[ind-1].Col, coords[ind].Col
		if coords[ind-1].Col > coords[ind].Col {
			minCol, maxCol = coords[ind].Col, coords[ind-1].Col
		}

		for r := minRow; r <= maxRow; r++ {
			for c := minCol; c <= maxCol; c++ {
				boundaries++
			}
		}
		// don't double count
		// eg. ##### <- it will get doublecounted as a last and a first
		//         #
		//         #
		boundaries--

	}
	utils.Logger.Infof("%+v", 0.5*float64(area))
	utils.Logger.Infof("%+v", boundaries)

	return int(0.5*float64(area) + 0.5*float64(boundaries) + 1)
}

func (s Solver2023) Day_18(part int, reader utils.Reader) int {

	lines := make([]line16, 0)

	switch part {
	case 1:

		dirsMap := map[string]string{
			"U": "n",
			"D": "s",
			"R": "e",
			"L": "w",
		}

		for line := range reader.Stream() {
			splittedLine := strings.Split(utils.StandardizeSpaces(line), " ")

			steps, _ := strconv.Atoi(splittedLine[1])

			lines = append(lines, line16{
				dir:   dirsMap[splittedLine[0]],
				steps: steps,
			})
		}

		return traverseLagoon(lines)
	case 2:

		for line := range reader.Stream() {
			splittedLine := strings.Split(utils.StandardizeSpaces(line), " ")

			hex := strings.TrimSuffix(strings.TrimPrefix(splittedLine[2], "("), ")")

			steps, _ := strconv.ParseInt(hex[1:6], 16, 32)

			var dir string
			switch hex[6] {
			case '0':
				dir = "e"
			case '1':
				dir = "s"
			case '2':
				dir = "w"
			case '3':
				dir = "n"
			}

			utils.Logger.Debugf("%s %d", hex[1:6], steps)

			lines = append(lines, line16{
				dir:   dir,
				steps: int(steps),
			})
		}

		utils.Logger.Infof("%+v", lines)

		return traverseLagoon(lines)

	default:
		//shouldn't reach here
		return -1
	}
}
