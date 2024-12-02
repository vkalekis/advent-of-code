package src2024

import (
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

func isReportSafe(report []int) bool {
	mode := ""

	for i := 0; i < len(report)-1; i++ {
		if mode == "" {
			if report[i+1] > report[i] {
				mode = "asc"
			} else if report[i+1] < report[i] {
				mode = "desc"
			} else {
				// not a valid mode
				// utils.Logger.Infoln(report)
				return false
			}
		}

		var diff int
		if mode == "asc" {
			diff = report[i+1] - report[i]
		} else if mode == "desc" {
			diff = report[i] - report[i+1]
		}

		if !(diff >= 1 && diff <= 3) {
			// diff greater than the one requested
			return false
		}
	}

	return true
}

func (s *Solver2024) Day_02(part int, reader utils.Reader) int {
	reports := make([][]int, 0)
	for line := range reader.Stream() {

		utils.Logger.Debugln(line)

		levels := make([]int, 0)

		for _, l := range strings.Split(utils.StandardizeSpaces(line), " ") {
			iL, _ := strconv.Atoi(l)
			levels = append(levels, iL)
		}
		utils.Logger.Debugln(levels)

		reports = append(reports, levels)
	}

	utils.Logger.Infoln(len(reports))

	safeCount := 0
	for _, report := range reports {
		switch part {
		case 1:
			if isReportSafe(report) {
				safeCount++
			}
		case 2:

			if !isReportSafe(report) {
				// brute force, start removing each element and check if the subreport is safe
				found := false
				for i := 0; i < len(report); i++ {
					subReport := make([]int, len(report))
					copy(subReport, report)

					subReport = append(subReport[:i], subReport[i+1:]...)

					// utils.Logger.Infoln("Original:", report, "Modified:", subReport)
					if isReportSafe(subReport) {
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
