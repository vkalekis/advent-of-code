package src2024

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type node struct {
	c      utils.Coordinates
	height int
}

type graph map[node][]node

func generateGrid(reader utils.Reader) [][]string {
	var grid [][]string

	for line := range reader.Stream() {
		logger.Debugln(line)
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func generateGraph(grid [][]string) graph {

	g := make(graph)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			c := utils.NewCoordinates(row, col)
			n := node{
				c:      c,
				height: utils.ToInt(grid[row][col]),
			}
			g[n] = make([]node, 0)
			for _, nn := range c.GetNeighbors(len(grid), len(grid[0])) {
				g[n] = append(g[n], node{
					c:      nn,
					height: utils.ToInt(grid[nn.Row][nn.Col]),
				})
			}
		}
	}

	return g
}

func findPathStarts(grid [][]string) []utils.Coordinates {
	pathStarts := make([]utils.Coordinates, 0)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == "0" {
				c := utils.NewCoordinates(row, col)
				pathStarts = append(pathStarts, c)
			}
		}
	}

	return pathStarts
}

func dfs(g graph, start node, visited map[node]bool, score *int, considerVisitedNodes bool) {

	visited[start] = true
	// logger.Infof("Start: %v", start)

	if start.height == 9 {
		// found the end of the path with gradual steps (+1)
		*score++
	}

	for _, n := range g[start] {
		if considerVisitedNodes {
			if !visited[n] && n.height == start.height+1 {
				dfs(g, n, visited, score, considerVisitedNodes)
			}
		} else {
			if n.height == start.height+1 {
				dfs(g, n, visited, score, considerVisitedNodes)
			}
		}
	}
}

func (s *Solver2024) Day_10(part int, reader utils.Reader) int {

	grid := generateGrid(reader)
	g := generateGraph(grid)

	for k, v := range g {
		logger.Debugf("%+v -> %+v", k, v)
	}

	considerVisitedNodes := false
	switch part {
	case 1:
		considerVisitedNodes = true
	case 2:
	}

	totalScore := 0
	for _, pathStart := range findPathStarts(grid) {
		visited := make(map[node]bool)
		score := 0
		start := node{
			c: pathStart,
		}
		dfs(g, start, visited, &score, considerVisitedNodes)

		logger.Infof("pathStart:%v -> %d", start, score)
		totalScore += score
	}

	return totalScore
}
