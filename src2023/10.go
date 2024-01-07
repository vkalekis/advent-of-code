package src2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

type neighbor struct {
	dir    string
	labels []string
}

type node10 struct {
	label      string
	coords     utils.Coordinates
	n, s, e, w *node10
	visited    bool
	dist       int
}

func move(row, col int, dir string, maxRows, maxCols int) (string, int, int) {

	switch dir {
	case "n":
		if row > 0 {
			return fmt.Sprintf("%d-%d", row-1, col), row - 1, col
		}
	case "s":
		if row < maxRows-1 {
			return fmt.Sprintf("%d-%d", row+1, col), row + 1, col
		}
	case "w":
		if col > 0 {
			return fmt.Sprintf("%d-%d", row, col-1), row, col - 1
		}
	case "e":
		if col < maxCols-1 {
			return fmt.Sprintf("%d-%d", row, col+1), row, col + 1
		}
	}

	return fmt.Sprintf("%d-%d", row, col), row, col
}

func (cur *node10) getNeightbors() []neighbor {

	neighbors := make([]neighbor, 0)

	north := neighbor{
		dir:    "n",
		labels: []string{"|", "F", "7", "S"},
	}
	south := neighbor{
		dir:    "s",
		labels: []string{"|", "L", "J", "S"},
	}
	east := neighbor{
		dir:    "e",
		labels: []string{"-", "J", "7", "S"},
	}
	west := neighbor{
		dir:    "w",
		labels: []string{"-", "L", "F", "S"},
	}

	switch cur.label {
	case "|":
		neighbors = append(neighbors, north)
		neighbors = append(neighbors, south)
	case "-":
		neighbors = append(neighbors, east)
		neighbors = append(neighbors, west)
	case "L":
		neighbors = append(neighbors, north)
		neighbors = append(neighbors, east)
	case "J":
		neighbors = append(neighbors, north)
		neighbors = append(neighbors, west)
	case "7":
		neighbors = append(neighbors, south)
		neighbors = append(neighbors, west)
	case "F":
		neighbors = append(neighbors, south)
		neighbors = append(neighbors, east)
	}

	return neighbors
}

func (cur *node10) lookupneighbors(lines map[int][]string, visitedCoords map[string]*node10, row, col, maxRows, maxCols int) {

	utils.Logger.Infof(">>> Current: (%v %v) %v %+v", row, col, cur.label, cur)

	// for _, dir := range directions {
	for _, n := range cur.getNeightbors() {
		utils.Logger.Infof(">>%v", n)
		dir := n.dir

		ncoords, nrow, ncol := move(row, col, dir, maxRows, maxCols)
		utils.Logger.Infof(">>>> Debug coordinates of neighbords of (%v,%v): (%v,%v) in %v ", row, col, nrow, ncol, dir)
		if neigh, ok := visitedCoords[ncoords]; !ok {
			// if nrow != row && ncol != col {

			if lines[nrow][ncol] == "." {
				continue
			}
			found := false
			utils.Logger.Infof("???? %v %v", n.labels, lines[nrow][ncol])
			for _, v := range n.labels {
				utils.Logger.Infof("???? %v %v", v, lines[nrow][ncol])
				if lines[nrow][ncol] == v {
					found = true
					break
				}
			}
			if !found {
				continue
			}

			neigh := &node10{
				label: lines[nrow][ncol],
				coords: utils.Coordinates{
					Row: row,
					Col: col,
				},
			}
			utils.Logger.Infof(">>>>>> Here: %v (%v,%v) (%v,%v) %v -> %v", dir, row, col, nrow, ncol, cur.label, neigh.label)
			utils.Logger.Infof(">>>>>> Here: %v (%v,%v) (%v,%v) %p %p", dir, row, col, nrow, ncol, cur, neigh)

			switch dir {
			case "n":
				cur.n = neigh
				neigh.s = cur
			case "s":
				cur.s = neigh
				neigh.n = cur
			case "w":
				cur.w = neigh
				neigh.e = cur
			case "e":
				cur.e = neigh
				neigh.w = cur
			}

			visitedCoords[ncoords] = neigh
			visitedCoords[fmt.Sprintf("%d-%d", row, col)] = cur

			neigh.lookupneighbors(lines, visitedCoords, nrow, ncol, maxRows, maxCols)
		} else {
			found := false
			for _, v := range n.labels {
				if neigh.label == v {
					found = true
					break
				}
			}
			if !found {
				continue
			}

			utils.Logger.Infof(">>>>>> Here1: %v (%v %v) (%v %v) %v -> %v", dir, row, col, nrow, ncol, cur.label, neigh.label)

			switch dir {
			case "n":
				cur.n = neigh
				neigh.s = cur
			case "s":
				cur.s = neigh
				neigh.n = cur
			case "w":
				cur.w = neigh
				neigh.e = cur
			case "e":
				cur.e = neigh
				neigh.w = cur
			}

			visitedCoords[fmt.Sprintf("%d-%d", row, col)] = cur
			visitedCoords[fmt.Sprintf("%d-%d", nrow, ncol)] = neigh
		}
	}
}

func buildGraph(lines map[int][]string, maxRows, maxCols int) (map[string]*node10, map[string]interface{}, utils.Coordinates) {
	visitedCoords := make(map[string]*node10, 0)
	grounds := make(map[string]interface{})
	var scoordinates utils.Coordinates

	for row := 0; row < maxRows-1; row++ {
		line := lines[row]

		for col, c := range line {
			coords := fmt.Sprintf("%d-%d", row, col)
			if c == "." {
				grounds[coords] = struct{}{}
				continue
			}
			if c == "S" {
				scoordinates = utils.Coordinates{
					Row: row,
					Col: col,
				}
			}

			cur := &node10{
				label: c,
				coords: utils.Coordinates{
					Row: row,
					Col: col,
				},
			}

			if _, ok := visitedCoords[coords]; !ok {
				visitedCoords[coords] = cur
			}

			utils.Logger.Infof("(%v,%v) %v", row, col, cur.label)

			cur.lookupneighbors(lines, visitedCoords, row, col, maxRows, maxCols)

		}
	}

	return visitedCoords, grounds, scoordinates
}

func (n *node10) getSSymbol() string {
	if n.n != nil && n.s != nil {
		return "|"
	} else if n.e != nil && n.w != nil {
		return "-"
	} else if n.n != nil && n.e != nil {
		return "L"
	} else if n.n != nil && n.w != nil {
		return "J"
	} else if n.s != nil && n.w != nil {
		return "7"
	} else if n.s != nil && n.e != nil {
		return "F"
	} else {
		return "."
	}
}

func (root *node10) do_bfs() int {
	q := utils.NewQueue[*node10]()

	q.Enqueue(root)
	root.visited = true
	root.dist = 0

	maxDist := 0

	utils.Logger.Infof("Queue: %+v %v", q.Items(), q.IsEmpty())

	for !q.IsEmpty() {
		n, ok := q.Dequeue()
		utils.Logger.Debugf("Currently on: (%v,%v) %p %+v", n.coords.Row, n.coords.Col, n, n)

		if n.dist > maxDist {
			maxDist = n.dist
		}

		if !ok {
			utils.Logger.Errorf("Error on dequeue on empty queue")
			return -1
		}

		if n.n != nil && !n.n.visited {
			n.n.visited = true
			n.n.dist = n.dist + 1
			q.Enqueue(n.n)
		}
		if n.e != nil && !n.e.visited {
			n.e.visited = true
			n.e.dist = n.dist + 1
			q.Enqueue(n.e)
		}
		if n.w != nil && !n.w.visited {
			n.w.visited = true
			n.w.dist = n.dist + 1
			q.Enqueue(n.w)
		}
		if n.s != nil && !n.s.visited {
			n.s.visited = true
			n.s.dist = n.dist + 1
			q.Enqueue(n.s)
		}
	}

	return maxDist
}

func loc(visitedCoords map[string]*node10, row, col, maxRows, maxCols int) (int, int) {
	var count1, count2 int
	prev := ""
	c := col - 1

	for c >= 0 {
		if n, ok := visitedCoords[fmt.Sprintf("%v-%v", row, c)]; ok {
			// logger.Infof("--- %v %v %v %v %v", prev, n.label, n.row, n.col, prev == "" && n.label != "-" && n.label != "|")
			if n.label == "S" {
				n.label = n.getSSymbol()
			}
			if !n.visited {
				c--
				continue
			}

			if n.label == "|" {
				// fmt.Println("HIT |", count1)
				count1++
				c--
				continue
			} else if n.label == "-" {
				c--
				continue
			}

			if prev == "" {
				if n.label != "-" && n.label != "|" {
					prev = n.label
				}
			} else {

				if prev == "7" && n.label == "F" {
					//pass case
					prev = ""
				} else if prev == "J" && n.label == "L" {
					// pass
					prev = ""
				} else if prev == "J" && n.label == "F" {
					// fmt.Println("HIT JF", count1)
					count1++
					prev = ""
				} else if prev == "7" && n.label == "L" {
					// fmt.Println("HIT 7L", count1)
					count1++
					prev = ""
				} else {
					// panic(fmt.Sprintf("%v %v %v %v", prev, n.label, n.row, n.col))
					// prev = n.label
				}

			}
		}
		c--
	}

	c = col + 1
	prev = ""
	for c < maxCols {

		if n, ok := visitedCoords[fmt.Sprintf("%v-%v", row, c)]; ok {
			if n.label == "S" {
				n.label = n.getSSymbol()
			}
			if !n.visited {
				c++
				continue
			}

			if n.label == "|" {
				// fmt.Println("HIT |")
				count2++
				c++
				continue
			} else if n.label == "-" {
				c++
				continue
			}

			if prev == "" {

				prev = n.label

			} else {

				if prev == "F" && n.label == "7" {
					//pass case
					prev = ""
				} else if prev == "L" && n.label == "J" {
					// pass
					prev = ""
				} else if prev == "F" && n.label == "J" {
					// fmt.Println("HIT FJ")
					count2++
					prev = ""
				} else if prev == "L" && n.label == "7" {
					// fmt.Println("HIT L7")
					count2++
					prev = ""
				} else {
					// panic(fmt.Sprintf("a%v %v %v %v", prev, n.label, n.row, n.col))
					// prev = n.label
				}

			}
		}
		c++
	}

	return count1, count2
}

func (s Solver2023) Day_10(part int, reader utils.Reader) int {

	lines := make(map[int][]string)
	lineInd := 0

	for line := range reader.Stream() {
		lines[lineInd] = strings.Split(line, "")
		lineInd++
	}

	maxRows, maxCols := lineInd, len(lines[0])
	utils.Logger.Debugf("Lines: %+v MR: %v MC: %v", lines, maxRows, maxCols)

	var srow, scol int

	visitedCoords, grounds, scoordinates := buildGraph(lines, maxRows, maxCols)

	srow = scoordinates.Row
	scol = scoordinates.Col

	utils.Logger.Infof("S is here: (%v,%v) and is: %v",
		srow, scol, visitedCoords[fmt.Sprintf("%v-%v", srow, scol)].getSSymbol())

	maxDist := visitedCoords[fmt.Sprintf("%v-%v", srow, scol)].do_bfs()

	// output map to file
	output := make(map[string]interface{})
	for _, v := range visitedCoords {
		output[fmt.Sprintf("%d-%d", v.coords.Row, v.coords.Col)] = v.dist
	}
	if err := utils.WriteMapToFile(output, "outputs/output10.json"); err != nil {
		utils.Logger.Errorf("Could not write output to file %s: %v", "output10.json", err)
	}

	switch part {
	case 1:
		return maxDist

	case 2:
		total := 0
		for gr, _ := range grounds {

			row, _ := strconv.Atoi(strings.Split(gr, "-")[0])
			col, _ := strconv.Atoi(strings.Split(gr, "-")[1])

			countL, countR := loc(visitedCoords, row, col, maxRows, maxCols)

			if countL%2 == 1 && countR%2 == 1 {
				utils.Logger.Infof("%v %v    %v %v", row, col, countR, countR)
				total++
			}
		}

		for _, v := range visitedCoords {
			if !v.visited {
				countL, countR := loc(visitedCoords, v.coords.Row, v.coords.Col, maxRows, maxCols)
				if countL%2 == 1 && countR%2 == 1 {
					utils.Logger.Infof("%v %v    %v %v", v.coords.Row, v.coords, countL, countR)
					total++
				}
			}
		}

		return total
	default:
		// shouldn't reach here
		return -1
	}
}
