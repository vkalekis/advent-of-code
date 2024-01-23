package src2023

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
)

type lens struct {
	label string
	f     int
}

type box struct {
	lenses []lens
}

func extractValues(r *regexp.Regexp, s string) (string, string, int) {
	matches := r.FindStringSubmatch(s)

	var symbols, op, num string

	if len(matches) > 0 {
		symbols = matches[1]
		op = matches[2]
		num = matches[3]

		if num != "" {
			n, err := strconv.Atoi(num)
			if err != nil {
				utils.Logger.Errorf("Error while parsing int: %v", err)
				return "", "", -1
			}
			return symbols, op, n
		}
	}
	return symbols, op, -1
}

func (s Solver2023) Day_15(part int, reader utils.Reader) int {

	var currentValue int32
	var sumCurrentValues int32

	runHash := func(s string, currentValue int32) int32 {
		for _, runeValue := range s {
			// utils.Logger.Infof("ASCII %c - %d", runeValue, runeValue)
			// utils.Logger.Infof("ASCII %d - %d", x, x+1)
			currentValue = (17 * (currentValue + runeValue)) % 256
			utils.Logger.Debugf("Received: %c/%d - CurrentValue: %d", runeValue, runeValue, currentValue)
		}

		return currentValue
	}

	switch part {
	case 1:
		for line := range reader.Stream() {
			splittedLine := strings.Split(line, ",")

			for _, s := range splittedLine {
				currentValue = runHash(s, 0)
				sumCurrentValues += currentValue
				utils.Logger.Infof("%s - CurrentValue: %d - Sum CurrentValues: %d", s, currentValue, sumCurrentValues)
			}

		}

		return int(sumCurrentValues)
	case 2:

		// symbols - -/= - optionally a number
		regexp := regexp.MustCompile(`([a-zA-Z]+)(-|=)?(\d+)?`)
		boxes := make([]box, 256)

		for line := range reader.Stream() {
			splittedLine := strings.Split(line, ",")

			for _, s := range splittedLine {

				lensLabel, symbol, f := extractValues(regexp, s)

				boxIndex := runHash(lensLabel, 0)
				utils.Logger.Infof("%s - BoxIndex %d", s, boxIndex)

				existingLensInd := -1
				for ind, l := range boxes[boxIndex].lenses {
					if l.label == lensLabel {
						existingLensInd = ind
						break
					}
				}

				switch symbol {
				case "-":
					if existingLensInd != -1 {
						boxes[boxIndex].lenses = append(boxes[boxIndex].lenses[:existingLensInd], boxes[boxIndex].lenses[existingLensInd+1:]...)
					}
				case "=":
					if f == -1 {
						continue
					}

					if existingLensInd != -1 {
						boxes[boxIndex].lenses[existingLensInd].label = lensLabel
						boxes[boxIndex].lenses[existingLensInd].f = f
					} else {
						l := lens{
							label: lensLabel,
							f:     f,
						}
						// boxes[boxIndex].lenses = append([]lens{l}, boxes[boxIndex].lenses...)
						boxes[boxIndex].lenses = append(boxes[boxIndex].lenses, l)
					}

				}

				utils.Logger.Infof("%s - BoxIndex %d Lenses: %+v", s, boxIndex, boxes[boxIndex].lenses)

			}

		}

		totalFocusingPower := 0
		focusingPower := 0
		for bind, b := range boxes {
			if len(b.lenses) > 0 {
				utils.Logger.Infof("BoxIndex %d Lenses: %+v", bind, b.lenses)

				for lind, l := range b.lenses {
					focusingPower = (bind + 1) * (lind + 1) * l.f
					utils.Logger.Infof("BoxIndex %d Lens: %+v FocusingPower: %d", bind, l, focusingPower)
					totalFocusingPower += focusingPower
				}
			}
		}
		return totalFocusingPower

	default:
		//shouldn't reach here
		return -1
	}

}
