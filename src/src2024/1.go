package src2024

import (
	"math"
	"sort"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func (s *Solver2024) Day_01(part int, reader utils.Reader) int {

	llist := make([]int, 0)
	rlist := make([]int, 0)

	for line := range reader.Stream() {

		logger.Debugln(line)

		parts := strings.Split(utils.StandardizeSpaces(line), " ")
		if len(parts) != 2 {
			continue
		}

		id := utils.ToInt(parts[0])
		llist = append(llist, id)
		id = utils.ToInt(parts[1])
		rlist = append(rlist, id)
	}

	if len(llist) != len(rlist) {
		logger.Errorf("Error on matching lengths of llist (%d) and rlist (%d)", len(llist), len(rlist))
		return -1
	}

	logger.Debugf("LList: %v", llist)
	logger.Debugf("RList: %v", rlist)

	switch part {
	case 1:
		sort.Slice(llist, func(i, j int) bool {
			return llist[i] < llist[j]
		})
		sort.Slice(rlist, func(i, j int) bool {
			return rlist[i] < rlist[j]
		})

		logger.Debugf("Sorted LList: %v", llist)
		logger.Debugf("Sorted RList: %v", rlist)

		dist := 0.0
		for i := range len(llist) {
			dist += math.Abs(float64(llist[i] - rlist[i]))
		}
		return int(dist)

	case 2:

		rlistFrequencies := make(map[int]int)
		for _, id := range rlist {
			if _, found := rlistFrequencies[id]; !found {
				rlistFrequencies[id] = 0
			}
			rlistFrequencies[id]++
		}

		logger.Debugf("RList frequencies: %+v", rlistFrequencies)

		similarity := 0
		for _, id := range llist {
			if freq, found := rlistFrequencies[id]; found {
				similarity += id * freq
			}
		}
		return similarity

	}

	return int(-1)
}
