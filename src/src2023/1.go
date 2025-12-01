package src2023

import (
	"strings"
	"unicode"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func (s *Solver) Day_01(part int, reader utils.Reader) int {

	var sum int
	numsMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	var r, first, last int

	for line := range reader.Stream() {

		origLine := line
		// replace every string number occurence with the equivalent form of:
		// `firstChar Number lastChar` : eg. `eight` becomes `e8t`
		// in order to reuse the same logic from part 1 searching
		// for digits in the start and end of the string
		if part == 2 {
			for num, numRepl := range numsMap {
				line = strings.Replace(line, num, numRepl, -1)
			}
		}

		runes := []rune(line)

		var ind int
		// find first digit
		for ind = 0; ind < len(runes); ind++ {
			if unicode.IsDigit(runes[ind]) {
				r = utils.ToInt(string(runes[ind]))
				first = r
				break
			}
		}
		// find last digit
		for ind = len(runes) - 1; ind >= 0; ind-- {
			if unicode.IsDigit(runes[ind]) {
				r = utils.ToInt(string(runes[ind]))
				last = r
				break
			}
		}

		logger.Debugln(origLine, line, first, last, first*10+last)

		sum += first*10 + last
	}

	return int(sum)
}
