package src2024

import (
	"fmt"
	"os"
	"regexp"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type robotInfo struct {
	pos        utils.Coordinates
	vrow, vcol int
}

type robotGrid struct {
	robotsInfo []robotInfo
	robotsPos  map[utils.Coordinates]int
	rows, cols int
}

func constructRobotsGrid(reader utils.Reader, rows, cols int) robotGrid {
	robotGrid := robotGrid{
		rows: rows,
		cols: cols,
	}

	for line := range reader.Stream() {
		re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

		match := re.FindStringSubmatch(line)
		if len(match) > 1 {

			// x and y are inverted in the problem input wrt the standard definitions in utils pkg
			robotGrid.robotsInfo = append(robotGrid.robotsInfo, robotInfo{
				pos:  utils.NewCoordinates(utils.ToInt(match[2]), utils.ToInt(match[1])),
				vrow: utils.ToInt(match[4]),
				vcol: utils.ToInt(match[3]),
			})
		}
	}

	return robotGrid
}

func (g *robotGrid) view(s int) {

	grid := make([][]string, g.rows)
	for r := range g.rows {
		grid[r] = make([]string, g.cols)
		for c := range g.cols {
			if robotCount, found := g.robotsPos[utils.NewCoordinates(r, c)]; found {
				grid[r][c] = fmt.Sprintf("%d", robotCount)
			} else {
				grid[r][c] = "."
			}
		}
		// logger.Infof("%v", grid[r])
	}

	file, err := os.Create(fmt.Sprintf("./grid_%d.txt", s))
	if err != nil {
		return
	}
	defer file.Close()
	for _, row := range grid {
		_, err := fmt.Fprintln(file, row)
		if err != nil {
			return
		}
	}
}

func (g *robotGrid) moveRobots(seconds int) {
	endPos := make(map[utils.Coordinates]int)
	for _, robotInfo := range g.robotsInfo {

		var new utils.Coordinates
		new.Row = ((robotInfo.pos.Row+seconds*robotInfo.vrow)%g.rows + g.rows) % g.rows
		new.Col = ((robotInfo.pos.Col+seconds*robotInfo.vcol)%g.cols + g.cols) % g.cols

		endPos[new]++
	}

	g.robotsPos = endPos
}

func (g *robotGrid) countRobotsInQuadrants() int {

	// 1 | 2
	// - - -
	// 3 | 4
	findQuadrant := func(c utils.Coordinates) int {
		if c.Row < (g.rows-1)/2 {
			if c.Col < (g.cols-1)/2 {
				return 1
			} else if c.Col >= (g.cols+1)/2 {
				return 2
			}
		} else if c.Row >= (g.rows+1)/2 {
			if c.Col < (g.cols-1)/2 {
				return 3
			} else if c.Col >= (g.cols+1)/2 {
				return 4
			}
		}
		return 0
	}

	qCount := make(map[int]int)
	for pos, robotCount := range g.robotsPos {
		q := findQuadrant(pos)
		// logger.Infof("Pos: %v RobotCount: %d Q: %d", pos, robotCount, q)
		if q != 0 {
			qCount[q] += robotCount
		}
	}

	// logger.Infof("QCount: %+v", qCount)

	safetyFactor := 1
	for _, q := range qCount {
		safetyFactor *= q
	}

	return safetyFactor
}

func (g *robotGrid) findConnectedElements() int {
	visited := make(map[utils.Coordinates]bool)
	dirs := []utils.Coordinates{{Row: 0, Col: 1}, {Row: 0, Col: -1}}

	var dfs func(utils.Coordinates, *int)
	dfs = func(c utils.Coordinates, connectedComponents *int) {
		visited[c] = true
		*connectedComponents++
		for _, dir := range dirs {
			neighbor := utils.NewCoordinates(c.Row+dir.Row, c.Col+dir.Col)
			if _, exists := g.robotsPos[neighbor]; exists && !visited[neighbor] {
				dfs(neighbor, connectedComponents)
			}
		}
	}

	maxConnectedComponents := 0

	for c := range g.robotsPos {
		connectedComponents := 0
		if !visited[c] {
			dfs(c, &connectedComponents)
		}
		if connectedComponents > maxConnectedComponents {
			maxConnectedComponents = connectedComponents
		}
	}

	return maxConnectedComponents
}

func (s *Solver2024) Day_14(part int, reader utils.Reader) int {

	// g := constructRobotsGrid(reader, 7, 11)
	g := constructRobotsGrid(reader, 103, 101)

	switch part {
	case 1:
		g.moveRobots(100)
		logger.Infof("EndRobotPos: %+v", g.robotsPos)

		safetyFactor := g.countRobotsInQuadrants()

		return safetyFactor
	case 2:
		for s := range 1000000 {
			g.moveRobots(s)
			ce := g.findConnectedElements()
			// manually inspect the frames that have >10 connected elements in the horizontal direction
			if ce > 10 {
				logger.Infof("s=%d connectedElements: %d", s, ce)
				g.view(s)
			}
		}
	}

	return -1
}
