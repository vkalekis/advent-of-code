package src2024

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func (s *Solver2024) Day_05(part int, reader utils.Reader) int {
	orderRules := make(map[int][]int)
	updates := make([][]int, 0)

	for line := range reader.Stream() {
		logger.Debugln(line)

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				continue
			}
			page1 := utils.ToInt(parts[0])
			page2 := utils.ToInt(parts[1])

			if _, found := orderRules[page1]; !found {
				orderRules[page1] = make([]int, 0)
			}
			orderRules[page1] = append(orderRules[page1], page2)

		} else if len(line) == 0 {
			// separator
			continue
		} else {
			update := make([]int, 0)
			parts := strings.Split(line, ",")
			for _, part := range parts {
				page := utils.ToInt(part)
				update = append(update, page)
			}
			updates = append(updates, update)
		}
	}

	logger.Infof("%+v", orderRules)
	logger.Infof("%+v", updates)

	for _, update := range updates {
		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				if _, found := orderRules[update[i]]; found {

				}
			}
		}
	}
	return -1
}
