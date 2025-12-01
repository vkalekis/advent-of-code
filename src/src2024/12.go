package src2024

import (
	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func findConnectedRegions(g utils.Grid) map[utils.Coordinates][]utils.Coordinates {
	visited := make(map[utils.Coordinates]bool)
	regions := make(map[utils.Coordinates][]utils.Coordinates)

	var dfs func(g utils.Grid, cur utils.Coordinates, region *[]utils.Coordinates)
	dfs = func(g utils.Grid, cur utils.Coordinates, region *[]utils.Coordinates) {
		visited[cur] = true
		*region = append(*region, cur)
		for _, n := range cur.GetNeighbors(len(g), len(g[0])) {
			if !visited[n] && g[cur.Row][cur.Col] == g[n.Row][n.Col] {
				dfs(g, n, region)
			}
		}
	}

	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			c := utils.NewCoordinates(row, col)
			if !visited[c] {
				var region []utils.Coordinates
				dfs(g, c, &region)
				logger.Debugf("Start: %+v / %s Regions: %+v", c, g[c.Row][c.Col], region)

				regions[c] = region
			}
		}
	}

	return regions
}

func (s *Solver) Day_12(part int, reader utils.Reader) int {

	grid := utils.GenerateGrid(reader)
	regions := findConnectedRegions(grid)

	logger.Infof("Regions: %+v", regions)

	totalPrice := 0
	for regionStart, regionCoords := range regions {
		plant := grid[regionStart.Row][regionStart.Col]
		// for the perimeter we only consider the neighbor tiles that are not tiles of the same plant
		//   | |
		// - A A -
		//   | A -
		//     |
		perimeter := 0
		for _, regionCoord := range regionCoords {
			neighbors := regionCoord.GetNeighbors(len(grid), len(grid[0]))
			// logger.Infof("Plant: %s Neighbors: %d", plant, neighbors)
			// if the tile is an edge tile, the neighbors will be <4 but the edges are considered perimeter
			perimeter += 4 - len(neighbors)
			for _, neighbor := range neighbors {
				if grid[neighbor.Row][neighbor.Col] != plant {
					perimeter++
				}
			}

		}

		logger.Infof("Plant: %s Perimeter: %d Area: %d Price: %d", plant, perimeter, len(regionCoords), perimeter*len(regionCoords))
		totalPrice += perimeter * len(regionCoords)
	}

	return totalPrice
}
