package src2025

import (
	"sort"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func mergeRanges(ranges []utils.IdRange) []utils.IdRange {
	if len(ranges) == 0 {
		return nil
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	merged := []utils.IdRange{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		tail := merged[len(merged)-1]
		// overlapping
		// |-----|
		//     |------|
		//  ->
		// |----------|
		// also the ranges are inclusive: 1-5 6-7 -> 1-7
		if ranges[i].Min <= tail.Max+1 {
			tail.Max = utils.Max(ranges[i].Max, tail.Max)
			merged[len(merged)-1] = tail
		} else {
			merged = append(merged, ranges[i])
		}

		logger.Debugf("Merged: %v", merged)

	}
	return merged
}

func (s *Solver) Day_05(part int, reader utils.Reader) int {

	idRanges := make([]utils.IdRange, 0)
	ids := make([]int, 0)

	var freshIngredients int
	var switchedToIds bool

	for line := range reader.Stream() {
		if len(line) == 0 {
			switchedToIds = true
			continue
		}

		if !switchedToIds {
			parts := strings.Split(utils.Clean(line), "-")
			if len(parts) != 2 {
				continue
			}
			idRanges = append(idRanges, utils.IdRange{
				Min: utils.ToInt(parts[0]),
				Max: utils.ToInt(parts[1]),
			})
			continue
		}

		ids = append(ids, utils.ToInt(utils.Clean(line)))
	}

	logger.Debugf("IdRanges: %v \nIds: %v", idRanges, ids)

	switch part {
	case 1:
		for _, id := range ids {
			var fresh bool
			for _, idRange := range idRanges {
				if id >= idRange.Min && id <= idRange.Max {
					fresh = true
					break
				}
			}

			if fresh {
				logger.Debugf("Fresh ingredient %d", id)
				freshIngredients++
			}
		}
	case 2:
		merged := mergeRanges(idRanges)
		logger.Debugf("Merged: %v", merged)

		for _, mergedRange := range merged {
			freshIngredients += mergedRange.TotalIds()
		}
	}

	return freshIngredients
}
