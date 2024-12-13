package src2023

import (
	"math"
	"strings"
	"sync"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type node11 struct {
	coords utils.Coordinates
	dist   int
}

func dilate_universe(lines map[int][]string, maxRows, maxCols *int, expansionRate int) ([]int, []int) {
	// locate rows without galaxies
	noGalaxiesRows := make([]int, 0)

	for row := 0; row < *maxRows; row++ {
		empty := true
		for col := 0; col < *maxCols; col++ {
			if lines[row][col] == "#" {
				empty = false
				break
			}
		}
		if empty {
			noGalaxiesRows = append(noGalaxiesRows, row)
		}
	}

	// locate cols without galaxies
	noGalaxiesCols := make([]int, 0)

	for col := 0; col < *maxCols; col++ {
		empty := true
		for row := 0; row < *maxRows; row++ {
			if lines[row][col] == "#" {
				empty = false
				break
			}
		}
		if empty {
			noGalaxiesCols = append(noGalaxiesCols, col)
		}
	}

	logger.Infof("Rows without galaxies: %v", noGalaxiesRows)
	logger.Infof("Cols without galaxies: %v", noGalaxiesCols)

	logger.Infof("First row: %v - len %v", lines[0], len(lines[0]))

	if expansionRate > 0 {
		for rate := 1; rate < expansionRate; rate++ {
			if rate%100 == 0 {
				logger.Debugf(">> %d", rate)
			}
			for row := 0; row < *maxRows; row++ {
				for retry, noGalaxiesCol := range noGalaxiesCols {
					noGalaxiesCol += retry * rate
					lines[row] = append(lines[row][:noGalaxiesCol], append([]string{"."}, lines[row][noGalaxiesCol:]...)...)
				}
			}
			*maxCols += len(noGalaxiesCols)
		}

		logger.Infof("First row: %v - len %v", lines[0], len(lines[0]))

		logger.Infof("Old rows %d - old cols: %d", len(lines), *maxCols)
		for rate := 1; rate < expansionRate; rate++ {

			for retry, noGalaxiesRow := range noGalaxiesRows {
				noGalaxiesRow += retry * rate

				for rrow := *maxRows - 1; rrow > noGalaxiesRow-1; rrow-- {
					logger.Debugf("NoGalaxiesRow: %d rrow: %d", noGalaxiesRow, rrow)
					lines[rrow+1] = lines[rrow]
				}

				*maxRows++
			}
		}
	}

	return noGalaxiesRows, noGalaxiesCols
}

func floodfill(lines map[int][]string, start utils.Coordinates, maxRows, maxCols int) map[utils.Coordinates]int {
	q := utils.NewQueue[*node11]()

	visitedNodes := make(map[utils.Coordinates]int, 0)

	q.Enqueue(&node11{
		coords: start,
		dist:   0,
	})
	visitedNodes[start] = 0

	for !q.IsEmpty() {
		n, ok := q.Dequeue()
		logger.Debugf("Currently on: (%v,%v) %+v Len queue: %d", n.coords.Row, n.coords.Col, n, len(q.Items()))

		if !ok {
			logger.Errorf("Error on dequeue on empty queue")
			break
		}

		neighbors := n.coords.GetNeighbors(maxRows, maxCols)

		for _, neighbor := range neighbors {
			if _, ok := visitedNodes[neighbor]; ok {
				continue
			}

			q.Enqueue(&node11{
				coords: neighbor,
				dist:   n.dist + 1,
			})
			visitedNodes[neighbor] = n.dist + 1
		}

		// logger.Debugf("Neighbors of %+v : %v", n, neighbors)
		// if len(q.Items()) == 10 {
		// 	os.Exit(123)
		// }
	}
	logger.Debugf("VisitedNodes: %+v", visitedNodes)

	return visitedNodes
}

func manhattan_distance(start utils.Coordinates, galaxiesCoords []utils.Coordinates) int {
	totalDist := 0
	for _, galaxyCoord := range galaxiesCoords {
		dist := math.Abs(float64(start.Row-galaxyCoord.Row)) + math.Abs(float64(start.Col-galaxyCoord.Col))
		totalDist += int(dist)
		logger.Debugf("Dist %+v -> %+v = %d", start, galaxyCoord, int(dist))

	}
	return totalDist
}

func naive_way(lines map[int][]string, maxRows, maxCols int, expansionRate int) int {
	_, _ = dilate_universe(lines, &maxRows, &maxCols, expansionRate)

	logger.Infof("New rows %d - new cols: %d", maxRows, maxCols)
	for i := 0; i < maxRows; i++ {
		logger.Debugf("Final rows after parsing: row[%d] = %v", i, lines[i])
	}

	galaxiesCoords := make([]utils.Coordinates, 0)
	for row := 0; row < maxRows; row++ {
		for col := 0; col < maxCols; col++ {
			if lines[row][col] == "#" {

				galaxiesCoords = append(galaxiesCoords, utils.NewCoordinates(row, col))
			}
		}
	}
	logger.Infoln(galaxiesCoords)

	// we double-count so we need to divide by 2 at the end
	totalDist := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, galaxyCoord := range galaxiesCoords {
		wg.Add(1)
		go func(galaxyCoord utils.Coordinates) {
			defer wg.Done()

			localTotalDist := manhattan_distance(galaxyCoord, galaxiesCoords)

			mu.Lock()
			totalDist += localTotalDist
			mu.Unlock()
		}(galaxyCoord)

	}

	wg.Wait()

	return totalDist / 2
}

func (s Solver2023) Day_11(part int, reader utils.Reader) int {

	lines := make(map[int][]string)
	lineInd := 0

	for line := range reader.Stream() {
		lines[lineInd] = strings.Split(line, "")
		lineInd++
	}

	maxRows, maxCols := lineInd, len(lines[0])

	logger.Infof("New rows %d - new cols: %d", maxRows, maxCols)
	for i := 0; i < maxRows; i++ {
		logger.Debugf("Final rows after parsing: row[%d] = %v", i, lines[i])
	}

	var expansionRate int
	var galaxiesCoords []utils.Coordinates

	noGalaxiesRows, noGalaxiesCols := dilate_universe(lines, &maxRows, &maxCols, -1)

	switch part {
	case 1:
		expansionRate = 1
	case 2:
		expansionRate = 999999
	default:
		//shouldn't reach here
		return -1
	}

	// return naive_way(lines, maxRows, maxCols, expansionRate+1)

	galaxiesCoords = make([]utils.Coordinates, 0)
	for row := 0; row < maxRows; row++ {
		for col := 0; col < maxCols; col++ {
			if lines[row][col] == "#" {
				grow, gcol := row, col

				for i := 0; i < expansionRate; i++ {
					for _, noGalaxiesRow := range noGalaxiesRows {
						if row > noGalaxiesRow {
							grow++
						}
					}
					for _, noGalaxiesCol := range noGalaxiesCols {
						if col > noGalaxiesCol {
							gcol++
						}
					}
				}

				galaxiesCoords = append(galaxiesCoords, utils.NewCoordinates(grow, gcol))

			}
		}
	}

	logger.Infof("Galaxies coords: %v", galaxiesCoords)

	// we double-count so we need to divide by 2 at the end
	totalDist := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, galaxyCoord := range galaxiesCoords {
		wg.Add(1)
		go func(galaxyCoord utils.Coordinates) {
			defer wg.Done()

			localTotalDist := manhattan_distance(galaxyCoord, galaxiesCoords)

			mu.Lock()
			totalDist += localTotalDist
			mu.Unlock()
		}(galaxyCoord)

	}

	wg.Wait()

	return totalDist / 2
}
