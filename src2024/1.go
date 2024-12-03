package src2024

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

func (s *Solver2024) Day_01(part int, reader utils.Reader) int {

	llist := make([]int, 0)
	rlist := make([]int, 0)

	for line := range reader.Stream() {

		utils.Logger.Debugln(line)

		parts := strings.Split(utils.StandardizeSpaces(line), " ")
		if len(parts) != 2 {
			continue
		}

		id, _ := strconv.Atoi(parts[0])
		llist = append(llist, id)
		id, _ = strconv.Atoi(parts[1])
		rlist = append(rlist, id)
	}

	if len(llist) != len(rlist) {
		utils.Logger.Errorf("Error on matching lengths of llist (%d) and rlist (%d)", len(llist), len(rlist))
		return -1
	}

	utils.Logger.Debugf("LList: %v", llist)
	utils.Logger.Debugf("RList: %v", rlist)

	switch part {
	case 1:
		sort.Slice(llist, func(i, j int) bool {
			return llist[i] < llist[j]
		})
		sort.Slice(rlist, func(i, j int) bool {
			return rlist[i] < rlist[j]
		})

		utils.Logger.Debugf("Sorted LList: %v", llist)
		utils.Logger.Debugf("Sorted RList: %v", rlist)

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

		utils.Logger.Debugf("RList frequencies: %+v", rlistFrequencies)

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
