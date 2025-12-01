package src2024

import (
	"fmt"
	"math"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type stone int

func blink(stones []stone) []stone {
	newIteration := make([]stone, 0)
	for _, s := range stones {

		if s == 0 {
			newIteration = append(newIteration, 1)
		} else if (int(math.Floor(math.Log10(float64(s))))+1)%2 == 0 {
			numberLength := int(math.Floor(math.Log10(float64(s)))) + 1
			leftPart := utils.ToInt(fmt.Sprintf("%d", s)[:numberLength/2])
			rightPart := utils.ToInt(fmt.Sprintf("%d", s)[numberLength/2:])

			newIteration = append(newIteration, stone(leftPart))
			newIteration = append(newIteration, stone(rightPart))
		} else {
			newIteration = append(newIteration, s*2024)
		}
	}

	return newIteration
}

func blinkMap(stones map[stone]int) map[stone]int {
	newIteration := make(map[stone]int)

	for s, count := range stones {

		if s == 0 {
			newIteration[1] += count
		} else if (int(math.Floor(math.Log10(float64(s))))+1)%2 == 0 {
			numberLength := int(math.Floor(math.Log10(float64(s)))) + 1
			leftPart := utils.ToInt(fmt.Sprintf("%d", s)[:numberLength/2])
			rightPart := utils.ToInt(fmt.Sprintf("%d", s)[numberLength/2:])

			newIteration[stone(leftPart)] += count
			newIteration[stone(rightPart)] += count
		} else {
			newIteration[s*2024] += count
		}
	}

	return newIteration
}
func (s *Solver) Day_11(part int, reader utils.Reader) int {

	stones := make([]stone, 0)
	for line := range reader.Stream() {
		strStones := strings.Split(utils.StandardizeSpaces(line), " ")
		for _, strStone := range strStones {
			stones = append(stones, stone(utils.ToInt(strStone)))
		}
	}

	logger.Infof("Stones: %v", stones)

	switch part {
	case 1:
		for r := 1; r <= 25; r++ {
			stones = blink(stones)
			logger.Infof("Blinked %d: Total: %d", r, len(stones))
			// logger.Infof("Blinked %d: Total: %d Stones: %v", r, len(stones), stones)
		}

		return len(stones)
	case 2:
		stonesMap := make(map[stone]int)
		for _, s := range stones {
			stonesMap[s] = 1
		}

		for r := 1; r <= 75; r++ {
			stonesMap = blinkMap(stonesMap)

			total := 0
			for _, c := range stonesMap {
				total += c
			}
			logger.Infof("Blinked %d: Total: %d", r, total)
			// logger.Infof("Blinked %d: Total: %d Stones: %v", r, total, stonesMap)
		}

		total := 0
		for _, c := range stonesMap {
			total += c
		}
		return total
	}

	return -1
}
