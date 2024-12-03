package src2024

import (
	"regexp"
	"strconv"

	"github.com/vkalekis/advent-of-code/utils"
)

func calculateTotal(line string) int {
	mulRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := mulRegex.FindAllStringSubmatch(line, -1)

	utils.Logger.Debugln(matches)

	total := 0
	for _, match := range matches {
		if len(match) == 3 {
			operandA, _ := strconv.Atoi(match[1])
			operandB, _ := strconv.Atoi(match[2])
			utils.Logger.Debugf("Found: mul(%d,%d)", operandA, operandB)

			total += operandA * operandB
		}
	}

	return total
}

func (s *Solver2024) Day_03(part int, reader utils.Reader) int {

	var longLine string

	for line := range reader.Stream() {
		utils.Logger.Debugln(line)
		longLine += line
	}

	total := 0

	switch part {
	case 1:
		return calculateTotal(longLine)
	case 2:

		matchesRegex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
		matches := matchesRegex.FindAllString(longLine, -1)
		utils.Logger.Debugln(matches)

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
		utils.Logger.Debugln(keptMuls)
		total = calculateTotal(keptMuls)
	}

	return total
}
