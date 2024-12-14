package src2024

import (
	"fmt"
	"regexp"
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

func (op operation) Is() int {
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
	result     int
	operands   []int
	operations []operation
}

type operationRange struct {
	start int
	end   int
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

func generateOperationCombinations(ops []string, x int, current string, result *[]string) {
	if x == 0 {
		*result = append(*result, current)
		return
	}

	for _, op := range ops {
		generateOperationCombinations(ops, x-1, current+op, result)
	}
}

func (eq *equation) doMath() bool {

	// mults := make([]operationRange, 0)

	result := eq.operands[0]
	for i := 1; i < len(eq.operands); i++ {
		switch eq.operations[i-1].Is() {
		case Multiplication:
			result = result * eq.operands[i]
		case Addition:
			result = result + eq.operands[i]
		case Concatenation:
			result = utils.ToInt(fmt.Sprintf("%d%d", result, eq.operands[i]))
		}

		if result > eq.wantResult {
			break
		}
	}

	eq.result = result

	logger.Debugf("Operands: %v Result: %d", eq.operands, result)

	return result == eq.wantResult

}

func (eq *equation) doBackWardMath() bool {

	// mults := make([]operationRange, 0)

	result := eq.wantResult
	foundWrongResult := false

LOOP:
	for i := len(eq.operands) - 1; i >= 1; i-- {

		switch eq.operations[i-1].Is() {
		case Multiplication:
			logger.Infoln(eq.wantResult, eq.operands[i], eq.operands[i-1], int(result/eq.operands[i]), result)
			if int(result/eq.operands[i])%eq.operands[i-1] != 0 {
				foundWrongResult = true
				break LOOP
			}
			result = int(result / eq.operands[i])

		case Addition:
			// logger.Infoln(eq.wantResult, eq.operands[i], eq.operands[i-1], eq.operands[i]+eq.operands[i-1], result)

			// if eq.operands[i]+eq.operands[i-1] != result {
			// 	foundWrongResult = true
			// 	break LOOP
			// }
			result -= eq.operands[i]
			if i == 1 && result != eq.operands[0] {
				foundWrongResult = true
			}

		case Concatenation:
			if utils.ToInt(fmt.Sprintf("%d%d", eq.operands[i], eq.operands[i-1])) == result {
				foundWrongResult = true
				break LOOP
			}
			result = eq.operands[i-1]
		}
	}

	logger.Debugf("Operands: %v FoundWrongResult: %t", eq.operands, foundWrongResult)

	return foundWrongResult
}

func (s *Solver2024) Day_07(part int, reader utils.Reader) int {

	equations := constructEquations(reader)

	calibrationResult := 0

	for _, equation := range equations {

		availableOperations := make([]string, 0)

		switch part {
		case 1:
			generateOperationCombinations([]string{"*", "+"}, len(equation.operands)-1, "", &availableOperations)
		case 2:
			generateOperationCombinations([]string{"*", "+", "|"}, len(equation.operands)-1, "", &availableOperations)
		}

		for _, operationsList := range availableOperations {
			operations := make([]operation, len(operationsList))
			for i := range operationsList {
				operations[i] = operation(operationsList[i])
			}

			// equation.operations = []operation(operations)
			// found := equation.doMath()
			// logger.Debugf("Equation: %+v Found: %t", equation, found)
			// if found {
			// 	calibrationResult += equation.wantResult
			// 	break
			// }

			equation.operations = []operation(operations)
			foundWrongResult := equation.doBackWardMath()
			logger.Infof("Equation: %+v FoundWrongResult: %t", equation, foundWrongResult)

			if !foundWrongResult {
				calibrationResult += equation.wantResult
				break
			}
		}
	}

	return calibrationResult
}
