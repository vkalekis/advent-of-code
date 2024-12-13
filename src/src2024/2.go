package src2024

import (
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type report []int

func (r report) isSafe() bool {
	mode := ""

	for i := 0; i < len(r)-1; i++ {
		if mode == "" {
			// get starting mode from first two elements
			if r[i+1] > r[i] {
				mode = "asc"
			} else if r[i+1] < r[i] {
				mode = "desc"
			} else {
				// not a valid mode
				// logger.Infoln(report)
				return false
			}
		}

		var diff int
		if mode == "asc" {
			diff = r[i+1] - r[i]
		} else if mode == "desc" {
			diff = r[i] - r[i+1]
		}

		if !(diff >= 1 && diff <= 3) {
			// diff greater than the one requested
			return false
		}
	}

	return true
}

func (s *Solver2024) Day_02(part int, reader utils.Reader) int {
	reports := make([]report, 0)
	for line := range reader.Stream() {

		logger.Debugln(line)

		levels := make([]int, 0)

		for _, l := range strings.Split(utils.StandardizeSpaces(line), " ") {
			iL := utils.ToInt(l)
			levels = append(levels, iL)
		}
		logger.Debugln(levels)

		reports = append(reports, report(levels))
	}

	logger.Infoln(len(reports))

	safeCount := 0
	for _, r := range reports {
		switch part {
		case 1:
			if r.isSafe() {
				safeCount++
			}
		case 2:
			if !r.isSafe() {
				// brute force, start removing each element and check if the subreport is safe
				found := false
				for i := 0; i < len(r); i++ {
					subReport := make(report, len(r))
					copy(subReport, r)

					subReport = append(subReport[:i], subReport[i+1:]...)

					// logger.Infoln("Original:", report, "Modified:", subReport)
					if subReport.isSafe() {
						found = true
						break
					}
				}

				if found {
					safeCount++
				}
			} else {
				safeCount++
			}
		}

	}

	return safeCount
}
