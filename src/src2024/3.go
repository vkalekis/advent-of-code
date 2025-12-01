package src2024

import (
	"regexp"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func calculateTotal(line string) int {
	mulRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := mulRegex.FindAllStringSubmatch(line, -1)

	logger.Debugln(matches)

	total := 0
	for _, match := range matches {
		if len(match) == 3 {
			operandA := utils.ToInt(match[1])
			operandB := utils.ToInt(match[2])
			logger.Debugf("Found: mul(%d,%d)", operandA, operandB)

			total += operandA * operandB
		}
	}

	return total
}

func (s *Solver) Day_03(part int, reader utils.Reader) int {

	var longLine string

	for line := range reader.Stream() {
		logger.Debugln(line)
		longLine += line
	}

	total := 0

	switch part {
	case 1:
		return calculateTotal(longLine)
	case 2:

		matchesRegex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
		matches := matchesRegex.FindAllString(longLine, -1)
		logger.Debugln(matches)

		enabled := true
		keptMuls := ""
		for _, match := range matches {
			switch match {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					keptMuls += match
				}
			}
		}
		logger.Debugln(keptMuls)
		total = calculateTotal(keptMuls)
	}

	return total
}
