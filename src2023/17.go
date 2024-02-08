package src2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

func (s Solver2023) Day_17(part int, reader utils.Reader) int {
	grid, maxRows, maxCols := utils.ConstructIntGrid(reader)

	utils.Logger.Infoln(maxRows, maxCols)

	switch part {
	case 1:

		graph := NewGraph(utils.Logger)

		for i := 0; i < maxRows; i++ {
			for j := 0; j < maxCols; j++ {
				for k := 1; k <= 3; k++ {
					if j+k >= maxCols {
						continue
					}

					start := fmt.Sprintf("%d-%d", i, j)
					end := fmt.Sprintf("%d-%d", i, j+k)

					weight := grid[i][j+k]
					if k == 2 {
						weight += grid[i][j+k-1]
					} else if k == 3 {
						weight += grid[i][j+k-1] + grid[i][j+k-2]
					}
					graph.AddEdgeAndNodes(start, end, weight, false)
				}
			}
		}
		for j := 0; j < maxCols; j++ {
			for i := 0; i < maxRows; i++ {
				for k := 1; k <= 3; k++ {
					if i+k >= maxRows {
						continue
					}
					start := fmt.Sprintf("%d-%d", i, j)
					end := fmt.Sprintf("%d-%d", i+k, j)

					weight := grid[i+k][j]
					if k == 2 {
						weight += grid[i+k-1][j]
					} else if k == 3 {
						weight += grid[i+k-1][j] + grid[i+k-2][j]
					}

					graph.AddEdgeAndNodes(start, end, weight, false)
				}
			}
		}

		graph.PrintEdges()

		graph.Dijkstra17(graph.GetNode("0-0"))

		graph.PrintPath(graph.GetNode(fmt.Sprintf("%d-%d", maxRows-1, maxCols-1)))
		// graph.PrintPath(graph.GetNode("1-2"))

		path := graph.GetEntirePath(graph.GetNode(fmt.Sprintf("%d-%d", maxRows-1, maxCols-1)))
		heatLoss := 0
		for _, n := range path {

			r, c := strings.Split(n, "-")[0], strings.Split(n, "-")[1]
			row, _ := strconv.Atoi(r)
			col, _ := strconv.Atoi(c)

			heatLoss += grid[row][col]

			utils.Logger.Infof("%s -> %d", n, grid[row][col])
		}

		utils.Logger.Infoln(heatLoss)

	case 2:

	default:
		//shouldn't reach here
		return -1
	}

	return -1

}
