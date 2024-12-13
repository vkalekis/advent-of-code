package src2024

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func findStartingPoint(grid [][]string) (utils.Coordinates, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case "^":
				return utils.NewCoordinates(i, j), utils.Up
			case "<":
				return utils.NewCoordinates(i, j), utils.Left
			case ">":
				return utils.NewCoordinates(i, j), utils.Right
			case "v":
				return utils.NewCoordinates(i, j), utils.Down
			}
		}
	}

	return utils.NewCoordinates(-1, -1), -1
}

func moveGuard(grid [][]string, dir *int, guardCoords *utils.Coordinates) (bool, []utils.Coordinates) {
	hitBound := false

	visited := make([]utils.Coordinates, 0)

	switch *dir {
	case utils.Up:
		for i := guardCoords.Row; i >= 0; i-- {
			if grid[i][guardCoords.Col] == "#" {
				*dir = utils.Right
				guardCoords.Row = i + 1
				break
			}

			visited = append(visited, utils.NewCoordinates(i, guardCoords.Col))

			if i == 0 {
				logger.Infof("-> %d", *dir)
				guardCoords.Row = 0
				hitBound = true
			}
		}
	case utils.Down:
		for i := guardCoords.Row; i < len(grid); i++ {
			if grid[i][guardCoords.Col] == "#" {
				*dir = utils.Left
				guardCoords.Row = i - 1
				break
			}

			visited = append(visited, utils.NewCoordinates(i, guardCoords.Col))

			if i == len(grid)-1 {
				logger.Infof("-> %d", *dir)
				guardCoords.Row = len(grid) - 1
				hitBound = true
			}
		}
	case utils.Left:
		for j := guardCoords.Col; j >= 0; j-- {
			if grid[guardCoords.Row][j] == "#" {
				*dir = utils.Up
				guardCoords.Col = j + 1
				break
			}

			visited = append(visited, utils.NewCoordinates(guardCoords.Row, j))

			if j == 0 {
				logger.Infof("-> %d", *dir)
				guardCoords.Col = 0
				hitBound = true
			}
		}
	case utils.Right:
		for j := guardCoords.Col; j < len(grid[0]); j++ {
			if grid[guardCoords.Row][j] == "#" {
				*dir = utils.Down
				guardCoords.Col = j - 1
				break
			}

			visited = append(visited, utils.NewCoordinates(guardCoords.Row, j))

			if j == len(grid[0])-1 {
				logger.Infof("-> %d", *dir)
				guardCoords.Col = len(grid) - 1
				hitBound = true
			}
		}
	}

	return hitBound, visited
}

func getGuardPath(grid [][]string, guardCoords utils.Coordinates, dir int) []utils.Coordinates {
	visited := make([]utils.Coordinates, 0)

	var hitBound bool
	var v []utils.Coordinates

	for !hitBound {
		hitBound, v = moveGuard(grid, &dir, &guardCoords)

		logger.Infof("Dir: %s Guard: %s", dir, guardCoords)

		visited = append(visited, v...)
	}

	return visited
}

func (s *Solver2024) Day_06(part int, reader utils.Reader) int {

	var grid [][]string

	for line := range reader.Stream() {
		logger.Debugln(line)
		grid = append(grid, strings.Split(line, ""))
	}

	for _, row := range grid {
		logger.Infof("%v", row)
	}

	guardCoords, dir := findStartingPoint(grid)
	logger.Infof("Dir: %s Guard: %s", dir, guardCoords)

	switch part {
	case 1:

		visited := getGuardPath(grid, guardCoords, dir)

		for _, row := range grid {
			logger.Infof("%v", row)
		}

		return len(utils.RemoveDuplicates(visited))
	case 2:
		originalGuardCoords := utils.NewCoordinates(guardCoords.Row, guardCoords.Col)

		visited := getGuardPath(grid, guardCoords, dir)
		visited = utils.RemoveDuplicates(visited)

		totalLoops := 0

		// skip the original guard position
		for i := 1; i < len(visited); i++ {

			guardCoords = utils.NewCoordinates(originalGuardCoords.Row, originalGuardCoords.Col)
			dir = utils.Up

			logger.Infof("Obstacle at %s Guard at %s", visited[i], guardCoords)

			// add a new obstacle
			grid[visited[i].Row][visited[i].Col] = "#"

			// keep the visited obstacles, along with the direction the guard had when he collided with the obstacle
			visitedObstacles := make(map[utils.Coordinates][]int)
			visitedObstacles[guardCoords] = []int{dir}

			var hitBound bool
			for !hitBound {
				hitBound, _ = moveGuard(grid, &dir, &guardCoords)

				if obstacleDirs, found := visitedObstacles[guardCoords]; !found {
					visitedObstacles[guardCoords] = []int{dir}
				} else {
					// if the guard has reached a position with the same direction 2 times -> loop
					foundDir := false
					for _, obstacleDir := range obstacleDirs {
						if obstacleDir == dir {
							foundDir = true
							break
						}
					}

					if foundDir {
						logger.Infof("Loop! Dir:%d Guard: %s", dir, guardCoords)
						totalLoops++

						// for _, row := range grid {
						// 	logger.Infof("%v", row)
						// }

						break
					}
				}

			}

			grid[visited[i].Row][visited[i].Col] = "."
		}

		return totalLoops
	}

	return -1
}
