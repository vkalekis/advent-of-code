package main

// 55093 somehow

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/vkalekis/advent-of-code/utils"
)

func main() {

	logger, _ := utils.NewLogger()

	fr := utils.NewFileReader("inputs/input1")
	// fr := utils.NewDummyReader()
	go fr.Read()

	doB := true
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

	for line := range fr.StreamCh {
		origLine := line
		if doB {
			for num, numRepl := range numsMap {
				line = strings.Replace(line, num, numRepl, -1)
			}
		}

		runes := []rune(line)

		var ind int
		for ind = 0; ind < len(runes); ind++ {
			if unicode.IsDigit(runes[ind]) {
				strRune = string(runes[ind])
				intRune, err = strconv.ParseInt(strRune, 10, 64)
				if err != nil {
					logger.Errorf("Error while parsing int: %v", err)
					return
				}
				first = intRune
				break
			}
		}
		for ind = len(runes) - 1; ind >= 0; ind-- {
			if unicode.IsDigit(runes[ind]) {
				strRune = string(runes[ind])
				intRune, err = strconv.ParseInt(strRune, 10, 64)
				if err != nil {
					logger.Errorf("Error while parsing int: %v", err)
					return
				}
				last = intRune
				break
			}
		}

		logger.Debugln(origLine, line, first, last, first*10+last)

		sum += first*10 + last
	}

	logger.Infoln(sum)
}
