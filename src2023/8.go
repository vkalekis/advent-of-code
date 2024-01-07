package src2023

import (
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

type node struct {
	left  *node
	right *node
	label string
}

func (s Solver2023) Day_08(part int, reader utils.Reader) int {

	directions := ""
	nodesMap := make(map[string]node)

	lineInd := 0

	var startingNodes []node

	for line := range reader.Stream() {
		switch lineInd {
		case 0:
			directions = line
		case 1:
		default:
			splitLine := strings.Split(utils.StandardizeSpaces(line), " ")

			nodesMap[splitLine[0]] = node{
				left: &node{
					label: splitLine[2][1:4],
				},
				right: &node{
					label: splitLine[3][0:3],
				},
				label: splitLine[0],
			}

		}

		lineInd++

	}

	startingNodes = []node{nodesMap["AAA"]}

	utils.Logger.Debugf("Directions: %s", directions)
	utils.Logger.Debugf("Nodes: %+v", nodesMap)
	utils.Logger.Infof("StartingNodes: %v", startingNodes)

	found := false
	steps := 0

	for _, current := range startingNodes {

		utils.Logger.Infof("Starting from node %s", current.label)

		for !found {
			for _, runedir := range directions {
				dir := string(runedir)

				utils.Logger.Debugf("Current: %v", current)

				switch dir {
				case "L":
					new_node := (*current.left).label
					current = nodesMap[new_node]
				case "R":
					new_node := (*current.right).label
					current = nodesMap[new_node]
				}

				utils.Logger.Infof("Took %s direction and ended up on node %s", dir, current.label)

				if current.label == "ZZZ" {
					found = true
				}

				steps++
				if steps > 10 {
					found = true
				}
			}
		}

	}

	return steps
}
