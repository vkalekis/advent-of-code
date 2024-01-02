package src2023

import (
	"math"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
	"go.uber.org/zap"
)

type race struct {
	duration int
	record   int
}

func bruteforceRace(r race, logger *zap.SugaredLogger) int {
	waysToWin := 0
	foundRecord := false

	for t := 0; t <= r.duration; t++ {
		dist := t * (r.duration - t)
		if dist > r.record {
			logger.Debugf("OVER: Time: %d Distance: %d Record: %d", t, dist, r.record)
			waysToWin++
			foundRecord = true
		} else {
			logger.Debugf("Time: %d Distance: %d Record: %d", t, dist, r.record)
			if foundRecord {
				break
			}
		}
	}
	return waysToWin
}

func smartSolveRace(r race, logger *zap.SugaredLogger) int {
	determinant := math.Sqrt(float64(r.duration)*float64(r.duration) - 4*float64(r.record))
	sol1 := (r.duration + int(determinant)) / 2
	sol2 := (r.duration - int(determinant)) / 2
	logger.Infof("Sol1=%v Sol2=%v Sol1-Sol2=%v", sol1, sol2, sol1-sol2+1)
	return sol1 - sol2 + 1
}

func Day_06(part int, reader utils.Reader, logger *zap.SugaredLogger) int {

	switch part {
	case 1:
		lineInd := 0
		races := make([]race, 0)

		for line := range reader.Stream() {
			splitLine := strings.Split(utils.StandardizeSpaces(line), " ")

			// first line corresponds to race durations
			if lineInd == 0 {
				for _, strDur := range splitLine[1:] {
					dur, err := strconv.Atoi(strDur)
					if err != nil {
						logger.Errorf("Error while parsing int: %v", err)
						return -1
					}
					races = append(races, race{
						duration: dur,
					})
				}
			} else {
				// second line corresponds to race record distance
				for ind, strDist := range splitLine[1:] {
					dist, err := strconv.Atoi(strDist)
					if err != nil {
						logger.Errorf("Error while parsing int: %v", err)
						return -1
					}
					races[ind].record = dist
				}
			}

			lineInd++
		}

		logger.Infof("Races: %+v", races)

		totalWaysToWin := 1

		for _, race := range races {
			waysToWin := bruteforceRace(race, logger)

			logger.Infof("Race %+v: Ways to win: %d", race, waysToWin)
			totalWaysToWin *= waysToWin
		}

		return totalWaysToWin

	case 2:

		lineInd := 0
		var r race

		for line := range reader.Stream() {
			splitLine := strings.Split(utils.StandardizeSpaces(line), " ")

			// first line corresponds to race durations
			if lineInd == 0 {
				totalDurStr := ""
				for _, d := range splitLine[1:] {
					totalDurStr += d
				}
				totalDur, err := strconv.Atoi(totalDurStr)
				if err != nil {
					logger.Errorf("Error while parsing int: %v", err)
					return -1
				}
				r = race{
					duration: totalDur,
				}

			} else {
				totalDistStr := ""
				for _, d := range splitLine[1:] {
					totalDistStr += d
				}
				totalDist, err := strconv.Atoi(totalDistStr)
				if err != nil {
					logger.Errorf("Error while parsing int: %v", err)
					return -1
				}
				r.record = totalDist
			}

			lineInd++
		}

		logger.Infof("Race: %+v", r)

		// waysToWin := bruteforceRace(r, logger)
		waysToWin := smartSolveRace(r, logger)

		return waysToWin
	default:
		// shouldn't reach here
		return -1
	}

}
