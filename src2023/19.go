package src2023

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

type comparison struct {
	operandA, op string
	operandB     int
	workflow     string
}

type workflow struct {
	name        string
	comparisons []comparison
}

type part map[string]int

type interval struct {
	lower, upper int
}
type intervals map[string][]interval

func parseLines(reader utils.Reader) (map[string]workflow, []part) {
	workflowRegex := regexp.MustCompile(`(\w+)\{(.+)\}`)
	operationRegex := regexp.MustCompile(`(\w+)([<>])(\d+):(\w+)`)
	partRegex := regexp.MustCompile(`\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}`)

	workflows := make(map[string]workflow, 0)
	parts := make([]part, 0)

	for line := range reader.Stream() {

		match := workflowRegex.FindStringSubmatch(line)

		comparisons := make([]comparison, 0)

		if len(match) > 2 {
			name := match[1]
			operations := strings.Split(match[2], ",")

			utils.Logger.Infoln("Name:", name)
			utils.Logger.Infoln("Operations:", operations)

			for ind := 0; ind < len(operations)-1; ind++ {

				match = operationRegex.FindStringSubmatch(operations[ind])

				if len(match) > 4 {
					operandB, _ := strconv.Atoi(match[3])

					comparisons = append(comparisons, comparison{
						operandA: match[1],
						op:       match[2],
						operandB: operandB,
						workflow: match[4],
					})
				}
			}

			comparisons = append(comparisons, comparison{
				workflow: operations[len(operations)-1],
			})

			workflows[name] = workflow{
				name:        name,
				comparisons: comparisons,
			}
		}

		if len(line) == 0 {
			break
		}
	}

	for line := range reader.Stream() {

		match := partRegex.FindStringSubmatch(line)

		utils.Logger.Infoln(match)

		if len(match) > 4 {
			x, _ := strconv.Atoi(match[1])
			m, _ := strconv.Atoi(match[2])
			a, _ := strconv.Atoi(match[3])
			s, _ := strconv.Atoi(match[4])

			parts = append(parts, part(map[string]int{
				"x": x,
				"m": m,
				"a": a,
				"s": s,
			}))

		}
	}
	return workflows, parts
}

func traversePath(parents map[string]string, start string) []string {
	path := make([]string, 0)
	for start != "in" {
		parent := parents[start]
		utils.Logger.Debugf(">>> %v %v", start, parent)

		path = append(path, start)
		start = parent
	}
	path = append(path, "in")
	return path
}

func do_dfs(graph map[string]map[string][]comparison, start string, visited map[string]bool, parents map[string]string, paths [][]string) [][]string {
	if visited[start] {
		return paths
	}

	visited[start] = true

	neighbors, found := graph[start]
	if !found {
		return paths
	}

	for neighbor, comparisons := range neighbors {
		parents[neighbor] = start
		for _, comparison := range comparisons {
			utils.Logger.Debugf("%s -> %s : comparison %v", start, neighbor, comparison)
			paths = do_dfs(graph, neighbor, visited, parents, paths)

			if strings.HasPrefix(neighbor, "A") {
				paths = append(paths, traversePath(parents, neighbor))
			}
		}

	}

	return paths
}

func mergeIntervals(intervals []interval) []interval {
	// [100,125] [40,150] -> [100,125]
	// [1,1800] [839,4000] -> [839,1800]

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].lower < intervals[j].upper
	})

	merged := []interval{intervals[0]}

	// utils.Logger.Debugf("%+v", intervals)

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		lastMerged := &merged[len(merged)-1]

		if curr.lower <= lastMerged.upper {
			// merge overlapping
			if curr.upper < lastMerged.upper {
				lastMerged.upper = curr.upper
			}
			if curr.lower > lastMerged.lower {
				lastMerged.lower = curr.lower
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}

func (s Solver2023) Day_19(p int, reader utils.Reader) int {

	workflows, parts := parseLines(reader)

	utils.Logger.Infof("%+v", workflows)
	utils.Logger.Infof("%+v", parts)

	switch p {
	case 1:

		acceptedPartsFields := 0

		for _, part := range parts {
			workflowsPath := "in -> "
			workflowName := "in"

			for workflowName != "A" && workflowName != "R" {

				workflow, found := workflows[workflowName]
				if !found {
					utils.Logger.Errorf("no in workflow found")
					return -2
				}

				nextWorkflowName := ""
				found = false
				for ind := 0; ind < len(workflow.comparisons)-1; ind++ {
					if found {
						break
					}
					operation := workflow.comparisons[ind]
					switch operation.op {
					case "<":
						if part[operation.operandA] < operation.operandB {
							nextWorkflowName = operation.workflow
							found = true
						}
					case ">":
						if part[operation.operandA] > operation.operandB {
							nextWorkflowName = operation.workflow
							found = true
						}
					}
				}
				if nextWorkflowName == "" {
					nextWorkflowName = workflow.comparisons[len(workflow.comparisons)-1].workflow
				}

				utils.Logger.Debugf("Part: %+v Workflow: %+v NextWorkflow: %+v", part, workflow, nextWorkflowName)
				workflowName = nextWorkflowName

				workflowsPath += fmt.Sprintf("%s -> ", workflowName)
			}

			utils.Logger.Infof("Part: %+v Workflows Path: %+v", part, workflowsPath)

			if workflowName == "A" {
				acceptedPartsFields += part["x"] + part["m"] + part["a"] + part["s"]
				utils.Logger.Infof("Part: %+v Value: %d", part, part["x"]+part["m"]+part["a"]+part["s"])
			}
		}
		return acceptedPartsFields
	case 2:

		graph := make(map[string]map[string][]comparison)

		aidx := 0

		for _, workflow := range workflows {

			if _, found := graph[workflow.name]; !found {
				graph[workflow.name] = make(map[string][]comparison)
			}

			for ind := 0; ind < len(workflow.comparisons); ind++ {
				if workflow.comparisons[ind].workflow == "A" {
					workflow.comparisons[ind].workflow = fmt.Sprintf("A_%d", aidx)
					aidx++
				}
				graph[workflow.name][workflow.comparisons[ind].workflow] = append(graph[workflow.name][workflow.comparisons[ind].workflow], workflow.comparisons[ind])
			}
		}

		paths := make([][]string, 0)
		paths = do_dfs(graph, "in", make(map[string]bool), make(map[string]string), paths)

		combinations := 0

		for _, path := range paths {

			i := intervals(map[string][]interval{
				"x": make([]interval, 0),
				"m": make([]interval, 0),
				"a": make([]interval, 0),
				"s": make([]interval, 0),
			})

			utils.Logger.Infof("Path: %+v", path)

			for ind := 1; ind < len(path); ind++ {
				prevWorkflow := workflows[path[ind]].comparisons
				utils.Logger.Debugf("Prev: %+v %s %s", prevWorkflow, path[ind-1], path[ind])
				for compInd := 0; compInd < len(prevWorkflow); compInd++ {
					comparison := prevWorkflow[compInd]

					if prevWorkflow[compInd].workflow != path[ind-1] {
						switch comparison.op {
						case "<":
							i[comparison.operandA] = append(i[comparison.operandA], interval{
								lower: comparison.operandB,
								upper: 4000,
							})
						case ">":
							i[comparison.operandA] = append(i[comparison.operandA], interval{
								lower: 1,
								upper: comparison.operandB,
							})
						}
					} else {
						switch comparison.op {
						case "<":
							i[comparison.operandA] = append(i[comparison.operandA], interval{
								lower: 1,
								upper: comparison.operandB - 1,
							})
						case ">":
							i[comparison.operandA] = append(i[comparison.operandA], interval{
								lower: comparison.operandB + 1,
								upper: 4000,
							})
						case "":
						}
						break
					}
				}
			}

			utils.Logger.Debugf("Intervals: %+v", i)

			localCombinations := 1
			for _, label := range []string{"x", "m", "a", "s"} {
				intervals := mergeIntervals(i[label])

				utils.Logger.Debugf("At %s : %+v", label, intervals)

				if len(intervals) == 0 {
					localCombinations *= 4000
				} else {
					localCombinations *= (intervals[0].upper - intervals[0].lower + 1)
				}

			}
			combinations += localCombinations
			utils.Logger.Debugf("Combinations %d : %d", localCombinations, combinations)
		}

		return combinations
	default:
		//shouldn't reach here
		return -1
	}
	return -1
}
