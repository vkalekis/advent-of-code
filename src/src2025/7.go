package src2025

import (
	"maps"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

const (
	start    = "S"
	splitter = "^"
)

func findAllIndices(s, substr string) []int {
	var indices []int
	var start int
	for {
		index := strings.Index(s[start:], substr)
		if index == -1 {
			break
		}
		indices = append(indices, start+index)
		start += index + 1
	}
	return indices
}

func (s *Solver) Day_07(part int, reader utils.Reader) int {
	tachyons := make(map[int]int)

	var lineIdx, columns int
	var totalSplits int

	for line := range reader.Stream() {
		iter := make(map[int]int)
		maps.Copy(iter, tachyons)

		if strings.Contains(line, start) {
			startIdxs := findAllIndices(line, start)
			iter[startIdxs[0]] = 1
			columns = len(line)
		} else {
			for idx, ch := range line {
				if string(ch) == splitter {
					// logger.Debugf("LineIdx: %d Iter: %+v SplitterIdx: %d", lineIdx, iter, idx)
					if _, ok := iter[idx]; ok {
						//    |
						//    ^
						//   | |
						//
						//    1
						//  1   1

						//     |   |
						//     ^   ^
						//   |   |   |
						//
						//     1   1
						//   1   2   1

						if idx-1 >= 0 {
							iter[idx-1] += iter[idx]
						}
						if idx+1 < columns {
							iter[idx+1] += iter[idx]
						}
						delete(iter, idx)
						totalSplits++
					}
					// logger.Debugf("Tachyons: %v", iter)
				}
			}
		}
		tachyons = iter
		logger.Debugln(line, tachyons)
		lineIdx++
	}

	switch part {
	case 1:
		return totalSplits
	case 2:
		var timelines int
		for _, c := range tachyons {
			timelines += c
		}
		return timelines
	}
	return -1
}
