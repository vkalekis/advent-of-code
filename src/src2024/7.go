package src2024

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

const (
	Multiplication = iota
	Addition
	Concatenation
)

type operation string

func (op operation) is() int {
	switch op {
	case "*":
		return Multiplication
	case "+":
		return Addition
	case "|":
		return Concatenation
	}
	return -1
}

type equation struct {
	wantResult int
	operands   []int
}

func constructEquations(reader utils.Reader) []equation {
	equations := make([]equation, 0)
	re := regexp.MustCompile(`^(\d+):((?: \d+)+)$`)

	for line := range reader.Stream() {

		logger.Debugf("%v", line)
		match := re.FindStringSubmatch(line)

		if match != nil {

			result := utils.ToInt(match[1])

			afterColon := strings.TrimSpace(match[2])

			operands := make([]int, 0)
			for _, numStr := range strings.Split(afterColon, " ") {
				operands = append(operands, utils.ToInt(numStr))
			}

			equations = append(equations, equation{
				wantResult: result,
				operands:   operands,
			})
		}
	}

	return equations
}

func (eq *equation) doMath(operations []operation) int {
	if len(operations)+1 != len(eq.operands) {
		return -1
	}

	result := eq.operands[0]
	for i := 1; i < len(eq.operands); i++ {
		switch operations[i-1].is() {
		case Multiplication:
			result = result * eq.operands[i]
		case Addition:
			result = result + eq.operands[i]
		case Concatenation:
			result = utils.ToInt(fmt.Sprintf("%d%d", result, eq.operands[i]))
		}
	}

	logger.Debugf("Operands: %v Result: %d", eq.operands, result)

	return result

}

// goBackwardsRecursive performs a recursive backtracking operation through an equation, starting from the last operand
// and working backwards to the first operand.
//   - It reverses the applied operations to find a valid sequence  of operations that leads to the expected result.
//   - At each step [if the operation is Multiplication/Concatenation], the function checks if the current operation
//     (applied to the current operand) can yield the current result. If not, the recursion stops, as there is no need to continue
//     searching for the rest of the operation sequence.
//
// The recursion stops when the first operand is reached, and is verified if the final result matches the first operand.
// Returns:
//   - true: if a valid sequence of operations is found that leads to the expected result.
//   - false: if no valid sequence can be found after exploring all possibilities at this step.
func (eq *equation) goBackWardsRecursive(part, i int, result int, op operation, opSequence *[]operation) bool {
	// if we're at the first operand, it must match the result
	if i == 0 {
		return eq.operands[0] == result
	}

	logger.Debugf("equation: %+v i=%d result=%d op=%v", eq, i, result, op)

	// we backpropagate and
	switch op.is() {
	case Multiplication:
		if eq.operands[i] == 0 || result%eq.operands[i] != 0 {
			return false
		}

		switch part {
		case 1:
			if eq.goBackWardsRecursive(part, i-1, result/eq.operands[i], "+", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, result/eq.operands[i], "*", opSequence) {
				*opSequence = append(*opSequence, op)
				return true
			}
		case 2:
			if eq.goBackWardsRecursive(part, i-1, result/eq.operands[i], "+", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, result/eq.operands[i], "*", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, result/eq.operands[i], "|", opSequence) {
				*opSequence = append(*opSequence, op)
				return true
			}
		}

		return false
	case Addition:

		switch part {
		case 1:
			if eq.goBackWardsRecursive(part, i-1, result-eq.operands[i], "+", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, result-eq.operands[i], "*", opSequence) {
				*opSequence = append(*opSequence, op)
				return true
			}
		case 2:
			if eq.goBackWardsRecursive(part, i-1, result-eq.operands[i], "+", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, result-eq.operands[i], "*", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, result-eq.operands[i], "|", opSequence) {
				*opSequence = append(*opSequence, op)
				return true
			}
		}

		return false

	case Concatenation:

		resultStr := fmt.Sprintf("%d", result)
		operandStr := fmt.Sprintf("%d", eq.operands[i])

		if !strings.HasSuffix(resultStr, operandStr) {
			return false
		}

		trimmed := strings.TrimSuffix(resultStr, operandStr)
		newResult := 0
		if trimmed != "" {
			var err error
			newResult, err = strconv.Atoi(trimmed)
			if err != nil {
				return false
			}
		}

		switch part {
		case 1:
			if eq.goBackWardsRecursive(part, i-1, newResult, "+", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, newResult, "*", opSequence) {
				*opSequence = append(*opSequence, op)
				return true
			}
		case 2:
			if eq.goBackWardsRecursive(part, i-1, newResult, "+", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, newResult, "*", opSequence) ||
				eq.goBackWardsRecursive(part, i-1, newResult, "|", opSequence) {
				*opSequence = append(*opSequence, op)
				return true
			}
		}

		return false
	default:
		return false
	}
}

func (s *Solver) Day_07(part int, reader utils.Reader) int {

	equations := constructEquations(reader)

	calibrationResult := 0

	for _, equation := range equations {
		opSequence := make([]operation, 0)

		var foundResult bool
		switch part {
		case 1:
			foundResult = equation.goBackWardsRecursive(part, len(equation.operands)-1, equation.wantResult, "+", &opSequence) ||
				equation.goBackWardsRecursive(part, len(equation.operands)-1, equation.wantResult, "*", &opSequence)
		case 2:
			foundResult = equation.goBackWardsRecursive(part, len(equation.operands)-1, equation.wantResult, "+", &opSequence) ||
				equation.goBackWardsRecursive(part, len(equation.operands)-1, equation.wantResult, "*", &opSequence) ||
				equation.goBackWardsRecursive(part, len(equation.operands)-1, equation.wantResult, "|", &opSequence)
		}

		logger.Infof("Equation: %v Found: %t OpSequence: %v", equation, foundResult, opSequence)

		if foundResult {
			calibrationResult += equation.wantResult
		}

		// logger.Infoln(equation.doMath(opSequence))
	}

	return calibrationResult
}
