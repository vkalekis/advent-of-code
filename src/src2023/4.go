package src2023

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func parseLine(line string) ([]string, []string) {
	splittedLine := strings.Split(line, "|")

	winning := strings.Split(utils.StandardizeSpaces(splittedLine[0]), " ")[2:]
	played := strings.Split(utils.StandardizeSpaces(splittedLine[1]), " ")

	return winning, played
}

func (s Solver) Day_04(part int, reader utils.Reader) int {

	switch part {

	case 1:
		totalPoints := 0
		cardInd := 1

		for line := range reader.Stream() {

			winning, played := parseLine(line)
			logger.Debugf("Card %d: Winning: %v - Played: %v", cardInd, winning, played)

			// 0  points for 0 won cards
			// 1  point  for 1 won card
			// *2 points for every extra won card
			points := 0
			for _, num := range played {
				for _, win := range winning {
					if win == num {
						if points == 0 {
							points = 1
						} else {
							points *= 2
						}

					}
				}
			}

			totalPoints += points
			cardInd++
		}

		return totalPoints

	case 2:
		cardCount := make(map[int]int)
		cardInd := 1

		for line := range reader.Stream() {

			winning, played := parseLine(line)
			logger.Debugf("Card %d: Winning: %v - Played: %v", cardInd, winning, played)

			winCount := 0
			for _, num := range played {
				for _, win := range winning {
					if win == num {
						winCount++
					}
				}
			}

			// increment the card count for the played card
			if _, ok := cardCount[cardInd]; !ok {
				cardCount[cardInd] = 1
			} else {
				cardCount[cardInd]++
			}

			// for every copy of the played card (in the case there are multiple played cards)
			// increment the card count for the won cards
			//   the won cards will be played when their turn arrives
			for copy := 1; copy <= cardCount[cardInd]; copy++ {
				for i := 0; i < winCount; i++ {
					if _, ok := cardCount[cardInd+i+1]; !ok {
						cardCount[cardInd+i+1] = 1
					} else {
						cardCount[cardInd+i+1]++
					}
				}
			}

			logger.Debugf("Card %d: WinCount %d - CardCount: %d", cardInd, winCount, cardCount)

			cardInd++
		}

		totalCards := 0
		for _, v := range cardCount {
			totalCards += v
		}
		return totalCards

	default:
		// shouldn't reach here
		return -1

	}
}
