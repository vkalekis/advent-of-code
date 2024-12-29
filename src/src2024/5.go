package src2024

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type update []int
type rules map[int][]int

func constructUpdatesAndRules(reader utils.Reader) ([]update, rules) {
	// pageX before pageY
	rules := make(rules)

	updates := make([]update, 0)

	for line := range reader.Stream() {
		logger.Debugln(line)

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				continue
			}
			page1 := utils.ToInt(parts[0])
			page2 := utils.ToInt(parts[1])

			if _, found := rules[page1]; !found {
				rules[page1] = make([]int, 0)
			}
			rules[page1] = append(rules[page1], page2)
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

	return updates, rules
}

func (u update) isInOrder(rules rules) bool {

	for i := 0; i < len(u); i++ {
		if rule, found := rules[u[i]]; found {
			// logger.Infof("Found: %d Rule:%v", update[i], rule)

			for _, ruleNum := range rule {
				// update[i] is supposed to be before ruleNum
				// if we find the ruleNum before update[i] -> rule is not valid
				// the update sequence is not in order
				for j := 0; j < i; j++ {
					if u[j] == ruleNum {
						// logger.Infof("Found: %d at %d", ruleNum, j)
						return false
					}
				}
			}

		}
	}

	return true
}

func (s *Solver2024) Day_05(part int, reader utils.Reader) int {

	updates, rules := constructUpdatesAndRules(reader)

	logger.Infof("Rules: %+v", rules)

	medianSum := 0

	for _, update := range updates {
		isInOrder := update.isInOrder(rules)

		logger.Infof("Update: %v InOrder: %t", update, isInOrder)

		switch part {
		case 1:
			if isInOrder {
				medianSum += update[len(update)/2]
			}
		case 2:
			if !isInOrder {
				// logger.Infof("Update before swap: %v", update)

				swapOnePair := func(update []int) {
					for i := 0; i < len(update); i++ {
						if rule, found := rules[update[i]]; found {
							// logger.Infof("Found: %d Rule:%v", update[i], rule)

							for _, ruleNum := range rule {

								// update[i] is supposed to be before ruleNum
								// if we find the ruleNum before update[i] -> rule is not valid
								// the update sequence is not in order, we swap the elements at indexes i,j
								for j := 0; j < i; j++ {
									if update[j] == ruleNum {
										// logger.Infof("Swapping: %d %d", update[i], update[j])
										update[j] = update[i]
										update[i] = ruleNum
										return
									}
								}
							}

						}
					}
				}

				for !update.isInOrder(rules) {
					swapOnePair(update)
				}

				// logger.Infof("Update after swap: %v", update)
				medianSum += update[len(update)/2]
			}
		}

	}
	return medianSum
}
