package main

import (
	"fmt"

	"github.com/vkalekis/advent-of-code/utils"
)

func main() {
	logger, err := utils.NewLogger("info")
	fmt.Println(err)
	graph := utils.NewGraph(logger)

	// a := utils.NewNode("A")
	// b := utils.NewNode("B")
	// c := utils.NewNode("C")
	// d := utils.NewNode("D")
	// e := utils.NewNode("E")
	// f := utils.NewNode("F")
	// g := utils.NewNode("G")

	// graph.AddNode(a)
	// graph.AddNode(b)
	// graph.AddNode(c)
	// graph.AddNode(d)
	// graph.AddNode(e)
	// graph.AddNode(f)
	// graph.AddNode(g)

	// graph.AddEdge(utils.NewEdge(a, b, 6))
	// graph.AddEdge(utils.NewEdge(a, d, 1))
	// graph.AddEdge(utils.NewEdge(a, c, 1))
	// graph.AddEdge(utils.NewEdge(a, f, 3))
	// graph.AddEdge(utils.NewEdge(b, e, 1))
	// graph.AddEdge(utils.NewEdge(c, d, 1))
	// graph.AddEdge(utils.NewEdge(e, g, 8))
	// graph.AddEdge(utils.NewEdge(d, g, 10))
	// graph.AddEdge(utils.NewEdge(f, g, 1))

	// graph.AddEdgeAndNodes("a", "b", 6)
	// graph.AddEdgeAndNodes("a", "d", 1)
	// graph.AddEdgeAndNodes("a", "c", 1)
	// graph.AddEdgeAndNodes("a", "f", 3)
	// graph.AddEdgeAndNodes("b", "e", 1)
	// graph.AddEdgeAndNodes("c", "d", 1)
	// graph.AddEdgeAndNodes("e", "g", 8)
	// graph.AddEdgeAndNodes("d", "g", 10)
	// graph.AddEdgeAndNodes("f", "g", 1)

	graph.AddEdgeAndNodes("s", "a", 1, true)
	graph.AddEdgeAndNodes("s", "b", 5, true)
	graph.AddEdgeAndNodes("a", "b", 2, true)
	graph.AddEdgeAndNodes("a", "d", 1, true)
	graph.AddEdgeAndNodes("a", "c", 2, true)
	graph.AddEdgeAndNodes("b", "d", 2, true)
	graph.AddEdgeAndNodes("c", "d", 3, true)
	graph.AddEdgeAndNodes("c", "e", 1, true)
	graph.AddEdgeAndNodes("d", "e", 2, true)

	// graph.Print()

	graph.Dijkstra(graph.GetNode("s"))

	graph.PrintExtensive()

	graph.PrintPath(graph.GetNode("e"))

	// q := utils.NewPriorityQueue()
	// q.Enqueue(utils.NewNode("aa", 21))
	// q.Enqueue(utils.NewNode("aaaa", 10))
	// x := utils.NewNode(">>aaaa", 3)
	// q.Enqueue(x)
	// q.Enqueue(utils.NewNode("aaaa", 4))
	// q.Enqueue(utils.NewNode("aaaa", 1))
	// q.Enqueue(utils.NewNode("aaaa", 20))
	// q.Enqueue(utils.NewNode("aaaa", -1))
	// q.Enqueue(utils.NewNode("aaaa", 5))
	// q.Enqueue(utils.NewNode("aaaa", 5))

	// for _, n := range q.Items() {
	// 	fmt.Printf("%+v\n", *n)
	// }

	// q.DecreaseDist(x, 2)
	// fmt.Println("\n\n")

	// for _, n := range q.Items() {
	// 	fmt.Printf("%+v\n", *n)
	// }

}
