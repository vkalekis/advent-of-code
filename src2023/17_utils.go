package src2023

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
	"go.uber.org/zap"
)

type Node struct {
	label     string
	row, col  int
	distance  int
	prev      *Node
	neighbors []*Node
	dirs      map[string]int
}

func NewNode(label string) *Node {
	return &Node{
		label:     label,
		neighbors: make([]*Node, 0),
		dirs:      make(map[string]int),
	}
}

type Edge struct {
	src, dst *Node
	weight   int
	directed bool
}

func NewEdge(src, dst *Node, weight int, directed bool) *Edge {
	src.neighbors = append(src.neighbors, dst)
	if !directed {
		dst.neighbors = append(dst.neighbors, src)
	}

	return &Edge{
		src:      src,
		dst:      dst,
		weight:   weight,
		directed: directed,
	}
}

type Graph struct {
	nodes  map[string]*Node
	edges  map[*Node]map[*Node]int
	logger *zap.SugaredLogger
}

func NewGraph(logger *zap.SugaredLogger) *Graph {
	return &Graph{
		nodes:  make(map[string]*Node),
		edges:  make(map[*Node]map[*Node]int),
		logger: logger,
	}
}

func (g *Graph) AddNode(n *Node) {
	r, c := strings.Split(n.label, "-")[0], strings.Split(n.label, "-")[1]
	row, _ := strconv.Atoi(r)
	col, _ := strconv.Atoi(c)
	n.row = row
	n.col = col
	g.nodes[n.label] = n
}

func (g *Graph) AddEdge(e *Edge) {
	if _, found := g.edges[e.src]; !found {
		g.edges[e.src] = make(map[*Node]int, 0)
	}
	g.edges[e.src][e.dst] = e.weight

	if !e.directed {
		if _, found := g.edges[e.dst]; !found {
			g.edges[e.dst] = make(map[*Node]int, 0)
		}
		g.edges[e.dst][e.src] = e.weight

	}
}

func (g *Graph) AddEdgeAndNodes(srcLabel, dstLabel string, weight int, directed bool) {
	var srcNode, dstNode *Node
	for l, n := range g.nodes {
		if l == srcLabel {
			srcNode = n
		}
		if l == dstLabel {
			dstNode = n
		}
	}

	if srcNode == nil {
		g.AddNode(NewNode(srcLabel))
		srcNode = g.GetNode(srcLabel)
	}
	if dstNode == nil {
		g.AddNode(NewNode(dstLabel))
		dstNode = g.GetNode(dstLabel)
	}

	g.AddEdge(NewEdge(srcNode, dstNode, weight, directed))
}

func (g *Graph) GetNode(label string) *Node {
	return g.nodes[label]
}

func (g *Graph) PrintEdges() {

	for src, edges := range g.edges {
		for dst, weight := range edges {
			g.logger.Infof("%s -> %s with %d", src.label, dst.label, weight)
			if weight == 0 {
				os.Exit(1)
			}
		}
	}
}

func (g *Graph) PrintExtensive() {
	for _, n := range g.nodes {
		nodeInfo := ""
		nodeInfo += fmt.Sprintf("Label: %s -  Dist: %d", n.label, n.distance)
		if len(n.neighbors) > 0 {
			nodeInfo += fmt.Sprintf("    Neigh: ")
			for _, neigh := range n.neighbors {
				nodeInfo += fmt.Sprintf("%s", neigh.label)
			}
		}
		if n.prev != nil {
			nodeInfo += fmt.Sprintf(" Prev: %s", n.prev.label)
		}
		g.logger.Infof("%s", nodeInfo)
	}

}

func (g *Graph) Dijkstra(src *Node) {

	q := NewPriorityQueue()

	for _, n := range g.nodes {
		n.distance = math.MaxInt
		n.prev = nil
		q.Enqueue(n)
	}

	visited := make(map[*Node]interface{})

	q.DecreaseDist(src, 0)

	for !q.IsEmpty() {
		n, ok := q.Dequeue()
		if !ok {
			g.logger.Errorf("Error on dequeue on empty queue")
			break
		}
		if _, found := visited[n]; found {
			continue
		}

		g.logger.Infof("Dequeued: %s Dist: %d - Q length %d", n.label, n.distance, q.Length())

		visited[n] = struct{}{}

		for _, neigh := range n.neighbors {
			if _, found := visited[neigh]; found {
				continue
			}

			newDistance := n.distance + g.edges[n][neigh]

			g.logger.Infof("%s OldDist: %d Edge: %d NewDist: %d", neigh.label, n.distance, g.edges[n][neigh], newDistance)

			if newDistance < neigh.distance {
				g.logger.Infof("Updating %s dist from %d to %d", neigh.label, neigh.distance, newDistance)
				neigh.distance = newDistance
				neigh.prev = n
				q.DecreaseDist(neigh, newDistance)

				g.logger.Infof("Setting %s previous to %s", n.label, neigh.label)

				// for _, i := range q.items {
				// 	g.logger.Infof(">>> %s %d", i.label, i.distance)
				// }
			}
		}

	}
}

func (g *Graph) Dijkstra17(src *Node) {

	q := NewPriorityQueue()

	for _, n := range g.nodes {
		n.distance = math.MaxInt
		n.prev = nil
		q.Enqueue(n)
	}

	visited := make(map[*Node]interface{})

	q.DecreaseDist(src, 0)
	src.dirs["e"] = 0

	for !q.IsEmpty() {
		n, ok := q.Dequeue()
		if !ok {
			g.logger.Errorf("Error on dequeue on empty queue")
			break
		}
		if _, found := visited[n]; found {
			continue
		}

		g.logger.Infof("Dequeued: %s Dist: %d - Q length %d", n.label, n.distance, q.Length())

		visited[n] = struct{}{}

		for _, neigh := range n.neighbors {
			if _, found := visited[neigh]; found {
				utils.Logger.Infoln("skipping ", neigh.label)
				continue
			}

			var dir string
			var steps int
			if neigh.row == n.row && neigh.col > n.col {
				dir = "e"
				steps = neigh.col - n.col
			} else if neigh.row == n.row && neigh.col < n.col {
				dir = "w"
				steps = n.col - neigh.col
			} else if neigh.col == n.col && neigh.row > n.row {
				dir = "s"
				steps = neigh.row - n.row
			} else if neigh.col == n.col && neigh.row < n.row {
				dir = "n"
				steps = n.row - neigh.row
			}

			// utils.Logger.Infof(">Steps %s %+v - %s %+v", n.label, n.dirs, neigh.label, neigh.dirs)
			if nsteps, nfound := n.dirs[dir]; nfound {
				if nsteps == 3 || nsteps+steps >= 3 {
					utils.Logger.Infoln("skipping ", neigh.label)
					continue
				} else {
					neigh.dirs[dir] = nsteps + steps
				}

			} else {
				neigh.dirs[dir] = steps
			}
			// utils.Logger.Infof(">>Steps %s %+v - %s %+v", n.label, n.dirs, neigh.label, neigh.dirs)

			newDistance := n.distance + g.edges[n][neigh]

			g.logger.Infof("%s OldDist: %d Edge: %d NewDist: %d", neigh.label, n.distance, g.edges[n][neigh], newDistance)

			if newDistance < neigh.distance {
				g.logger.Infof("Updating %s dist from %d to %d", neigh.label, neigh.distance, newDistance)
				neigh.distance = newDistance
				neigh.prev = n
				q.DecreaseDist(neigh, newDistance)

				g.logger.Infof("Setting %s previous to %s", n.label, neigh.label)
				// g.logger.Infof("Steps: %s-%s-%d %s-%s-%d", n.label, n.dir, n.steps, neigh.label, neigh.dir, neigh.steps)

			}
		}

		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		_, _ = buf.ReadBytes('\n')
	}
}

func (g *Graph) PrintPath(start *Node) {
	g.logger.Infoln(start.label, start.distance)
	for start.prev != nil {
		g.logger.Infoln(start.prev.label, start.prev.distance)
		start = start.prev
	}
}

func (g *Graph) GetEntirePath(start *Node) []string {
	path := make([]string, 0)

	// g.logger.Infoln(start.label, start.distance)

	for start.prev != nil {
		r, c := strings.Split(start.label, "-")[0], strings.Split(start.label, "-")[1]
		startrow, _ := strconv.Atoi(r)
		startcol, _ := strconv.Atoi(c)

		r, c = strings.Split(start.prev.label, "-")[0], strings.Split(start.prev.label, "-")[1]
		prevrow, _ := strconv.Atoi(r)
		prevcol, _ := strconv.Atoi(c)

		if startrow < prevrow {
			startrow, prevrow = prevrow, startrow
		}
		if startcol < prevcol {
			startcol, prevcol = prevcol, startcol
		}

		// g.logger.Infoln(startrow, startcol, prevrow, prevcol)

		for row := startrow; row >= prevrow; row-- {
			for col := startcol; col >= prevcol; col-- {

				if row == prevrow && col == prevcol {
					continue
				}
				// g.logger.Infoln("...", row, col)
				path = append(path, fmt.Sprintf("%d-%d", row, col))
			}
		}
		// g.logger.Infoln(start.prev.label, start.prev.distance)
		start = start.prev
	}

	path = append(path, "0-0")

	return path
}

type PriorityQueue struct {
	items  []*Node
	length int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		items:  make([]*Node, 0),
		length: 0,
	}
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *PriorityQueue) Enqueue(item *Node) {
	if len(q.items) == 0 {
		q.items = append(q.items, item)
		q.length++
		return
	}

	insertionInd := 0
	for _, qitem := range q.items {
		if qitem.distance < item.distance {
			insertionInd++
			continue
		}
	}

	q.items = append(q.items, &Node{})

	copy(q.items[insertionInd+1:], q.items[insertionInd:])
	q.items[insertionInd] = item
	q.length++
}

func (q *PriorityQueue) DecreaseDist(item *Node, dist int) {
	item.distance = dist
	sort.Slice(q.items, func(i, j int) bool {
		return q.items[i].distance < q.items[j].distance
	})
}

func (q *PriorityQueue) Dequeue() (*Node, bool) {
	if q.length > 0 {
		item := q.items[0]
		q.items = q.items[1:]
		q.length--
		return item, true
	}
	var zero *Node
	return zero, false
}

func (q *PriorityQueue) Items() []*Node {
	return q.items
}

func (q *PriorityQueue) Length() int {
	return q.length
}
