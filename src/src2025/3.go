package src2025

import (
	"math"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

// locateBestMax finds the digits of the maximum valued number (with totalDigits digits) in the original array
// the digits in this number are in the order found in the original array
//
// Logic: while trying to find the max number with X digits:
// check in the subset [0 - len(arr)-(X-1)) to leave (at least) X-1 digits in the right section of the table for the next iteration
//
//	 e.g. trying to find the max 3 digit number
//
//		[ 7 8 6 1 2 3 5 ]
//
// 1st iteration: search for the max in the array [ 8 7 6 1 2 ], leaving room for at least 2 digits for the next iterations
// found 8, the next iteration will take as input the rest of the array: [ 6 1 2 3 5 ]
// 2nd iteration: search for the max in the array [ 6 1 2 3 ], leaving room for 1 digit for the next iteration
// found 6, the next iteration takes as input the array [ 1 2 3 5 ]
// 3rd iteration: search for the max in the array [ 1 2 3 5 ] -> 5
// joltage = 865
func locateBestMax(arr []int, totalDigits int, results *[]int) {
	var max, idx int

	if totalDigits == 1 {
		max, _ = utils.MaxWithIndex(arr)
		*results = append(*results, max)
		return
	}

	logger.Debugf("%d: Going from 0-%d", totalDigits, len(arr)-(totalDigits-1))

	max, idx = utils.MaxWithIndex(arr[:len(arr)-(totalDigits-1)])
	logger.Debugf("Max: %d Idx: %d", max, idx)
	*results = append(*results, max)

	locateBestMax(arr[idx+1:], totalDigits-1, results)
}

// getNumberFromDigitsSlice constructs the base 10 value from a digits slice.
// [ 1 2 3 ] -> 123
func getNumberFromDigitsSlice(a []int) int {
	num := 0
	for i := range a {
		num += int(math.Pow10(len(a)-i-1)) * a[i]
	}
	return num
}

func (s *Solver) Day_03(part int, reader utils.Reader) int {

	var totalJoltage int

	for line := range reader.Stream() {

		var bank []int
		for _, battery := range strings.Split(line, "") {
			bank = append(bank, utils.ToInt(battery))
		}
		logger.Debugf("Bank: %v", bank)

		results := make([]int, 0)

		switch part {
		case 1:
			locateBestMax(bank, 2, &results)
		case 2:
			locateBestMax(bank, 12, &results)
		}

		logger.Debugf("Results: %v", results)

		totalJoltage += getNumberFromDigitsSlice(results)
	}

	return totalJoltage
}
