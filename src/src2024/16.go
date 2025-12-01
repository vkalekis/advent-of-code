package src2024

import (
	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func dfs_16(g utils.Graph, cur utils.Node, visited map[utils.Node]bool, path []utils.Node, allPaths *[][]utils.Node) {
	// append the current to the path
	path = append(path, cur)

	if cur.Data.(string) == "E" {
		*allPaths = append(*allPaths, path)
		return
	}

	visited[cur] = true

	for _, n := range g[cur] {
		if !visited[n] && n.Data.(string) != "#" {
			dfs_16(g, n, visited, path, allPaths)
		}
	}

	// unmark the current node to allow other paths using this node
	visited[cur] = false
}

func getPathScore(path []utils.Node) int {
	score := 0
	for i := 0; i < len(path)-2; i++ {
		if path[i].C.Row == path[i+2].C.Row || path[i].C.Col == path[i+2].C.Col {
			// didn't change direction
			score += 1
		} else {
			// changed direction, turned to some other direction
			score += 1001
		}
	}

	return score
}

func (s *Solver) Day_16(part int, reader utils.Reader) int {

	grid := utils.GenerateGrid(reader)
	g := utils.NewGraph(grid, func(e string) interface{} {
		return e
	})

	for k, v := range g {
		logger.Debugf("%+v -> %+v", k, v)
	}

	var start, end utils.Node
	for node := range g {
		if node.Data.(string) == "S" {
			start = node
		} else if node.Data.(string) == "E" {
			end = node
		}
	}

	logger.Infof("Start: %+v End: %+v", start, end)

	visited := make(map[utils.Node]bool)
	allPaths := [][]utils.Node{}
	dfs_16(g, start, visited, []utils.Node{}, &allPaths)

	for _, elem := range allPaths[0] {
		logger.Infof("%+v", elem)

		grid[elem.C.Row][elem.C.Col] = "~"
	}
	logger.Infof("%+v", getPathScore(allPaths[0]))
	for _, v := range grid {
		logger.Infof("%+v", v)
	}
	return -1
}
