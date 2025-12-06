package src2025

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

const (
	addition       = "+"
	multiplication = "*"
)

type array [][]int
type operations []string

func (s *Solver) Day_06(part int, reader utils.Reader) int {

	arr := make(array, 0)
	ops := make(operations, 0)
	row := 0

	for line := range reader.Stream() {
		parts := strings.Split(utils.StandardizeSpaces(line), " ")

		// last line reached!
		if parts[0] == addition || parts[0] == multiplication {
			ops = parts
			continue
		}

		arr = append(arr, make([]int, len(parts)))
		for partIdx, part := range parts {
			arr[row][partIdx] = utils.ToInt(part)
		}
		row++
	}

	logger.Debugf("Array=%+v", arr)
	logger.Debugf("Ops=%+v", ops)

	if len(arr[0]) != len(ops) {
		logger.Fatalf("mismatch in rows/cols")
	}

	var totalResults int
	for col := range len(arr[0]) {
		var problemResult int
		switch ops[col] {
		case addition:
			problemResult = 0
		case multiplication:
			problemResult = 1
		}

		for row := range len(arr) {
			// logger.Debugf("Element= %v", arr[row][col])
			switch ops[col] {
			case addition:
				problemResult += arr[row][col]
			case multiplication:
				problemResult *= arr[row][col]
			}
		}

		logger.Debugf("Col %d : Problem result %d", col, problemResult)
		totalResults += problemResult
	}

	return totalResults
}
