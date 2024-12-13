package src2023

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func detectFullZeroes(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func construct_levels(nums []int) (map[int][]int, int) {
	fullzeroes := false
	levels := make(map[int][]int)
	levels[0] = nums
	maxlevel := 1

	for !fullzeroes {
		diff := make([]int, len(nums)-1)

		for i := 1; i < len(nums); i++ {
			diff[i-1] = nums[i] - nums[i-1]
		}

		// logger.Infoln(diff)

		fullzeroes = detectFullZeroes(diff)
		nums = diff

		levels[maxlevel] = diff
		maxlevel++
	}
	return levels, maxlevel
}

func find_nextvalue(nums []int) int {

	levels, maxlevel := construct_levels(nums)

	logger.Infof("%+v", levels)

	// add 0 to the end of the last level
	levels[maxlevel-1] = append(levels[maxlevel-1], 0)
	for l := maxlevel - 2; l >= 0; l-- {
		// add the last values in the two lower levels going up the pyramid
		sum := levels[l][len(levels[l])-1] + levels[l+1][len(levels[l+1])-1]
		levels[l] = append(levels[l], sum)
	}

	logger.Infof("%+v -> %d", levels, levels[0][len(levels[0])-1])

	// next value
	return levels[0][len(levels[0])-1]
}

func find_prevvalue(nums []int) int {
	levels, maxlevel := construct_levels(nums)

	logger.Infof("%+v", levels)

	// add 0 to the begininng of the last level
	levels[maxlevel-1] = append([]int{0}, levels[maxlevel-1]...)

	for l := maxlevel - 2; l >= 0; l-- {
		// add the first values in the two lower levels going up the pyramid
		sum := levels[l][0] - levels[l+1][0]
		logger.Debugf("-- %d - %d = %d", levels[l][0], levels[l+1][0], levels[l][0]-levels[l+1][0])
		levels[l] = append([]int{sum}, levels[l]...)
	}

	logger.Infof("%+v --> %d", levels, levels[0][len(levels[0])-1])

	// previous value
	return levels[0][0]
}

func (s Solver2023) Day_09(part int, reader utils.Reader) int {

	nextValues := 0

	for line := range reader.Stream() {

		inpnums := strings.Split(line, " ")

		nums := make([]int, len(inpnums))
		for i, inum := range inpnums {
			num := utils.ToInt(inum)
			nums[i] = num
		}

		logger.Infoln(nums)

		switch part {
		case 1:
			nextValues += find_nextvalue(nums)
		case 2:
			nextValues += find_prevvalue(nums)
		default:
			// shouldn't reach here
			return -1

		}

	}
	return nextValues
}
