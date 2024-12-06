package src2024

import (
	"fmt"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

func convXMASWorker(inCh <-chan utils.Row, outCh chan int) {
	count := 0
	for row := range inCh {
		utils.Logger.Debugf("%v", row)
		if len(row) < 4 {
			continue
		}
		for i := 0; i <= len(row)-4; i++ {
			word := fmt.Sprintf("%s%s%s%s", row[i], row[i+1], row[i+2], row[i+3])
			utils.Logger.Debugf("Row: %v Word:%s", row, word)
			if word == "XMAS" || word == "SAMX" {
				count++
			}
		}
	}

	utils.Logger.Infof("Total count: %d", count)
	outCh <- count
}

// overkill but, eh, it was similar to the first part!
func conv2DXMASWorker(inCh <-chan utils.Array, outCh chan int) {
	count := 0
	for a := range inCh {
		for i := 0; i < len(a)-2; i++ {
			for j := 0; j < len(a[i])-2; j++ {
				// utils.Logger.Infof("%d/%d %d/%d", i, len(a), j, len(a[0]))

				// M.M  M.S  S.S  S.M
				// .A.  .A.  .A.  .A.
				// S.S  M.S  M.M  S.M
				if (a[i][j] == "M" && a[i][j+2] == "M" && a[i+1][j+1] == "A" && a[i+2][j] == "S" && a[i+2][j+2] == "S") ||
					(a[i][j] == "M" && a[i][j+2] == "S" && a[i+1][j+1] == "A" && a[i+2][j] == "M" && a[i+2][j+2] == "S") ||
					(a[i][j] == "S" && a[i][j+2] == "S" && a[i+1][j+1] == "A" && a[i+2][j] == "M" && a[i+2][j+2] == "M") ||
					(a[i][j] == "S" && a[i][j+2] == "M" && a[i+1][j+1] == "A" && a[i+2][j] == "S" && a[i+2][j+2] == "M") {
					count++
				}
			}
		}
	}

	utils.Logger.Infof("Total count: %d", count)
	outCh <- count
}

func (s *Solver2024) Day_04(part int, reader utils.Reader) int {
	var a utils.Array

	for line := range reader.Stream() {
		utils.Logger.Debugln(line)
		a = append(a, strings.Split(line, ""))
	}

	outCh := make(chan int)

	switch part {
	case 1:
		inCh := make(chan utils.Row)

		go convXMASWorker(inCh, outCh)

		for _, row := range a.ExtractRows() {
			inCh <- utils.Row(row)
		}
		for _, row := range a.ExtractColumns() {
			inCh <- utils.Row(row)
		}
		for _, row := range a.ExtractMainDiagonals() {
			inCh <- utils.Row(row)
		}
		for _, row := range a.ExtractAntiDiagonals() {
			inCh <- utils.Row(row)
		}

		close(inCh)
	case 2:
		inCh := make(chan utils.Array)

		go conv2DXMASWorker(inCh, outCh)

		inCh <- a

		close(inCh)
	}

	totalCount := <-outCh

	return totalCount
}
