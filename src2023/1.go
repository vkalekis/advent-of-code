package src2023

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/vkalekis/advent-of-code/utils"
)

func (s *Solver2023) Day_01(part int, reader utils.Reader) int {

	var sum int64
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

	var intRune, first, last int64
	var strRune string
	var err error

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
				strRune = string(runes[ind])
				intRune, err = strconv.ParseInt(strRune, 10, 64)
				if err != nil {
					utils.Logger.Errorf("Error while parsing int: %v", err)
					return -1
				}
				first = intRune
				break
			}
		}
		// find last digit
		for ind = len(runes) - 1; ind >= 0; ind-- {
			if unicode.IsDigit(runes[ind]) {
				strRune = string(runes[ind])
				intRune, err = strconv.ParseInt(strRune, 10, 64)
				if err != nil {
					utils.Logger.Errorf("Error while parsing int: %v", err)
					return -1
				}
				last = intRune
				break
			}
		}

		utils.Logger.Debugln(origLine, line, first, last, first*10+last)

		sum += first*10 + last
	}

	return int(sum)
}
