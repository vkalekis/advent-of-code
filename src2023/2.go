package src2023

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
	"go.uber.org/zap"
)

func getCount(round, label string, regex *regexp.Regexp) int {
	foundStr := regex.FindString(round)
	if len(foundStr) == 0 {
		return 0
	}
	foundVal, _ := strconv.Atoi(strings.Replace(foundStr, label, "", -1))
	return foundVal
}

func Day_02(part int, reader utils.Reader, logger *zap.SugaredLogger) int {
	gameIdReg := regexp.MustCompile("Game [0-9]+")
	redRegex := regexp.MustCompile("[0-9]+ red")
	greenRegex := regexp.MustCompile("[0-9]+ green")
	blueRegex := regexp.MustCompile("[0-9]+ blue")

	var validGame bool
	var reds, greens, blues int
	var minReds, minGreens, minBlues int
	var validGameIds, powers int

	for line := range reader.Stream() {

		minReds, minGreens, minBlues = 0, 0, 0
		validGame = true

		split1 := strings.Split(line, ":")
		rounds := strings.Split(split1[1], ";")

		gameId := strings.Replace(gameIdReg.FindString(split1[0]), "Game ", "", -1)
		gameIdInt, _ := strconv.Atoi(gameId)

		for _, round := range rounds {

			reds = getCount(round, " red", redRegex)
			greens = getCount(round, " green", greenRegex)
			blues = getCount(round, " blue", blueRegex)

			logger.Debugf("Reds=%d Greens=%d Blues=%d", reds, greens, blues)
			if reds > 12 || greens > 13 || blues > 14 {
				validGame = false
				// the min value of each ball color doesn't need to get calculated in part 1 -> stop the loop
				if part == 1 {
					break
				}
			}

			if reds > minReds {
				minReds = reds
			}
			if greens > minGreens {
				minGreens = greens
			}
			if blues > minBlues {
				minBlues = blues
			}

		}

		logger.Infof("Game: %d - Valid: %v", gameIdInt, validGame)
		logger.Debugf("Rounds: %v", rounds)
		if validGame {
			validGameIds += gameIdInt
		}

		if part == 2 {
			logger.Debugf("MinReds=%d MinGreens=%d MinBlues=%d", minReds, minGreens, minBlues)
			powers += minReds * minGreens * minBlues
		}

	}

	switch part {
	case 1:
		return validGameIds
	case 2:
		return powers
	default:
		// shouldn't reach here
		return -1
	}

}
