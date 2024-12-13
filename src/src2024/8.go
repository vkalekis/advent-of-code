package src2024

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func findAllAntennaPairs(antennas []utils.Coordinates) [][2]utils.Coordinates {
	var pairs [][2]utils.Coordinates

	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas); j++ {
			if i != j {
				pairs = append(pairs, [2]utils.Coordinates{antennas[i], antennas[j]})
			}
		}
	}

	return pairs
}

func (s *Solver2024) Day_08(part int, reader utils.Reader) int {

	antennas := make(map[string][]utils.Coordinates)
	var grid [][]string

	lineIdx := 0
	for line := range reader.Stream() {
		logger.Debugln(line)

		for cIdx, c := range line {
			if c != '.' && c != '#' {
				if _, found := antennas[string(c)]; !found {
					antennas[string(c)] = make([]utils.Coordinates, 0)
				}
				antennas[string(c)] = append(antennas[string(c)], utils.NewCoordinates(lineIdx, cIdx))
			}
		}

		grid = append(grid, strings.Split(line, ""))
		lineIdx++
	}

	rows, cols := len(grid), len(grid[0])

	uniqueAntinodes := make(map[utils.Coordinates]interface{})

	for c, as := range antennas {
		// we find all the pairs between the antennas of each label and also the mirrored pairs
		// eg. [ (1,1), (1,3) ] ->   [ (1,1), (1,3)], [ (1,3), (1,1)]
		pairs := findAllAntennaPairs(as)
		logger.Infof("Label:%s Antennas:%v Pairs:%v", c, as, pairs)

		switch part {
		case 1:
			for _, pair := range pairs {
				// diff = pair1-pair0
				// without the need to find the orientation:
				// +: pair0 + diff = pair1   -> rejected
				// -: pair0 - diff = antinode
				// as we consider both pairs of two antennas, this will traverse in both directions
				// (on the side of pair[0] and on the side of pair[1])
				diff := utils.NewCoordinates(
					pair[1].Row-pair[0].Row,
					pair[1].Col-pair[0].Col,
				)
				antinode := utils.NewCoordinates(
					pair[0].Row-diff.Row,
					pair[0].Col-diff.Col,
				)
				if antinode.Row < 0 || antinode.Row >= rows || antinode.Col < 0 || antinode.Col >= cols {
					continue
				}

				logger.Infof("Pair: %v Antinode: %v", pair, antinode)
				grid[antinode.Row][antinode.Col] = "#"
				uniqueAntinodes[antinode] = struct{}{}
			}
		case 2:
			for _, pair := range pairs {
				// diff = pair1-pair0
				// without the need to find the orientation:
				// +: pair0 + diff = pair1   -> it is considered an antinode now
				// -: pair0 - diff = antinode
				// as we consider both pairs of two antennas, this will traverse in both directions
				// (on the side of pair[0] and on the side of pair[1])

				// check if pair1 can be an antinode (this is valid in part 2)
				if pair[1].Row >= 0 && pair[1].Row < rows && pair[1].Col >= 0 && pair[1].Col < cols {
					logger.Infof("Pair: %v Antinode (pair1): %v", pair, pair[1])
					grid[pair[1].Row][pair[1].Col] = "#"
					uniqueAntinodes[pair[1]] = struct{}{}
				}

				for true {
					diff := utils.NewCoordinates(
						pair[1].Row-pair[0].Row,
						pair[1].Col-pair[0].Col,
					)
					antinode := utils.NewCoordinates(
						pair[0].Row-diff.Row,
						pair[0].Col-diff.Col,
					)
					if antinode.Row < 0 || antinode.Row >= rows || antinode.Col < 0 || antinode.Col >= cols {
						break
					}

					logger.Infof("Pair: %v Antinode: %v", pair, antinode)
					grid[antinode.Row][antinode.Col] = "#"
					uniqueAntinodes[antinode] = struct{}{}

					//                         pair[1]
					//              pair[0]
					//   antinode
					//
					//                |
					//               \ /
					//
					//              pair[1]
					//   pair[0]
					// in order to propagate the changes until we reach the bounds
					pair[1], pair[0] = pair[0], antinode
				}

			}

		}

	}

	for _, row := range grid {
		logger.Infof("%v", row)
	}

	return len(uniqueAntinodes)
}
