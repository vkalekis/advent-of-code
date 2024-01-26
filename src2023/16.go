package src2023

import (
	"strings"
	"sync"

	"github.com/vkalekis/advent-of-code/utils"
)

type node16 struct {
	coords utils.Coordinates
	dir    string
}

func getNeighbor(grid map[int][]string, start utils.Coordinates, maxRows, maxCols int, dir string) (string, utils.Coordinates) {
	row, col := start.Row, start.Col
	switch dir {
	case "n":
		if row > 0 {
			return grid[row-1][col], utils.NewCoordinates(row-1, col)
		}
	case "e":
		if col+1 < maxRows {
			return grid[row][col+1], utils.NewCoordinates(row, col+1)
		}
	case "s":
		if row+1 < maxRows {
			return grid[row+1][col], utils.NewCoordinates(row+1, col)
		}
	case "w":
		if col > 0 {
			return grid[row][col-1], utils.NewCoordinates(row, col-1)
		}
	}
	return "", utils.NewCoordinates(row, col)
}

func enqeueNextPosition(n node16, grid map[int][]string, q *utils.Queue[node16], maxRows, maxCols int) {
	row, col := n.coords.Row, n.coords.Col
	switch grid[row][col] {
	case "|":
		switch n.dir {
		case "e", "w":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row-1, col),
				dir:    "n",
			})
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row+1, col),
				dir:    "s",
			})
		case "n":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row-1, col),
				dir:    n.dir,
			})
		case "s":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row+1, col),
				dir:    n.dir,
			})
		}
	case "-":
		switch n.dir {
		case "n", "s":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col-1),
				dir:    "w",
			})
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col+1),
				dir:    "e",
			})
		case "e":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col+1),
				dir:    n.dir,
			})
		case "w":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col-1),
				dir:    n.dir,
			})
		}
	case "/":
		switch n.dir {
		case "e":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row-1, col),
				dir:    "n",
			})
		case "n":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col+1),
				dir:    "e",
			})
		case "w":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row+1, col),
				dir:    "s",
			})
		case "s":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col-1),
				dir:    "w",
			})
		}
	case "\\":
		switch n.dir {
		case "e":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row+1, col),
				dir:    "s",
			})
		case "n":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col-1),
				dir:    "w",
			})
		case "w":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row-1, col),
				dir:    "n",
			})
		case "s":
			q.Enqueue(node16{
				coords: utils.NewCoordinates(row, col+1),
				dir:    "e",
			})
		}

	default:
		neighbor, neighborCoords := getNeighbor(grid, n.coords, maxRows, maxCols, n.dir)
		if neighbor != "" {
			q.Enqueue(node16{
				coords: neighborCoords,
				dir:    n.dir,
			})
		}

	}

}

func deepCopy(grid map[int][]string) map[int][]string {
	copyGrid := make(map[int][]string)

	for k, v := range grid {
		copiedSlice := make([]string, len(v))
		copy(copiedSlice, v)
		copyGrid[k] = copiedSlice
	}

	return copyGrid
}

func (s Solver2023) Day_16(part int, reader utils.Reader) int {

	grid := make(map[int][]string)

	var maxRows, maxCols, row int

	for line := range reader.Stream() {

		splittedLine := strings.Split(line, "")
		utils.Logger.Infoln(splittedLine)

		grid[row] = splittedLine
		row++
	}

	maxRows = row
	maxCols = len(grid[0])

	rows := make([]int, maxRows)
	cols := make([]int, maxCols)
	for r := 0; r < maxRows; r++ {
		rows[r] = r
	}
	for c := 0; c < maxCols; c++ {
		cols[c] = c
	}

	switch part {
	case 1:

		visited := make(map[utils.Coordinates]string)

		start := utils.NewCoordinates(0, 0)

		q := utils.NewQueue[node16]()

		q.Enqueue(node16{
			coords: start,
			dir:    "e",
		})

		for !q.IsEmpty() {
			n, ok := q.Dequeue()
			if !ok {
				utils.Logger.Errorf("Error on dequeue on empty queue")
				break
			}

			if !n.coords.IsValid(maxRows, maxCols) {
				utils.Logger.Errorf("%+v", n.coords)
				continue
			}
			if dir, ok := visited[n.coords]; ok && dir == n.dir {
				continue
			}

			utils.Logger.Infof("Poped: %+v %s - Q: %v", n, grid[n.coords.Row][n.coords.Col], q.Items())

			visited[n.coords] = n.dir

			enqeueNextPosition(n, grid, q, maxRows, maxCols)

			if grid[n.coords.Row][n.coords.Col] == "." {
				grid[n.coords.Row][n.coords.Col] = "#"
			}
			// grid[n.coords.Row][n.coords.Col] = "#"

			utils.Logger.Infof("Defl: %v", q.Items())
			utils.Logger.Infof("GridR %d GridC: %d", n.coords.Row, n.coords.Col)

			// utils.Logger.Infoln(" ", cols)
			// for r := 0; r < maxRows; r++ {
			// 	utils.Logger.Infoln(rows[r], grid[r])
			// }

			// buf := bufio.NewReader(os.Stdin)
			// fmt.Print("> ")
			// _, _ = buf.ReadBytes('\n')

		}

		energizedTiles := len(visited)
		return energizedTiles
	case 2:
		startNodes := make([]node16, 0)
		for c := 0; c < maxCols; c++ {
			startNodes = append(startNodes, node16{
				coords: utils.NewCoordinates(0, c),
				dir:    "s",
			})
			startNodes = append(startNodes, node16{
				coords: utils.NewCoordinates(maxRows-1, c),
				dir:    "n",
			})
		}
		for r := 0; r < maxRows; r++ {
			startNodes = append(startNodes, node16{
				coords: utils.NewCoordinates(r, 0),
				dir:    "e",
			})
			startNodes = append(startNodes, node16{
				coords: utils.NewCoordinates(r, maxCols-1),
				dir:    "w",
			})
		}

		wg := sync.WaitGroup{}
		mtx := sync.Mutex{}
		maxEnergizedTiles := 0

		for _, start := range startNodes {

			wg.Add(1)

			localGrid := grid

			utils.Logger.Infof("Starting with %+v", start)

			go func(localGrid map[int][]string, start node16) {

				defer wg.Done()

				visited := make(map[utils.Coordinates]string)

				q := utils.NewQueue[node16]()

				q.Enqueue(node16{
					coords: start.coords,
					dir:    start.dir,
				})

				for !q.IsEmpty() {
					n, ok := q.Dequeue()
					if !ok {
						utils.Logger.Errorf("Error on dequeue on empty queue")
						break
					}

					if !n.coords.IsValid(maxRows, maxCols) {
						continue
					}
					if dir, ok := visited[n.coords]; ok && dir == n.dir {
						continue
					}

					utils.Logger.Debugf("Poped: %+v %s - Q: %v", n, localGrid[n.coords.Row][n.coords.Col], q.Items())

					visited[n.coords] = n.dir

					enqeueNextPosition(n, localGrid, q, maxRows, maxCols)

					if localGrid[n.coords.Row][n.coords.Col] == "." {
						localGrid[n.coords.Row][n.coords.Col] = "#"
					}

					utils.Logger.Debugf("Defl: %v", q.Items())
					utils.Logger.Debugf("GridR %d GridC: %d", n.coords.Row, n.coords.Col)
				}

				energizedTiles := len(visited)

				mtx.Lock()
				defer mtx.Unlock()
				if energizedTiles > maxEnergizedTiles {
					maxEnergizedTiles = energizedTiles
				}
			}(localGrid, start)

		}

		wg.Wait()

		return maxEnergizedTiles
	default:
		//shouldn't reach here
		return -1
	}

	return -1

}
