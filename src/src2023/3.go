package src2023

import (
	"regexp"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type Number struct {
	x     int
	xend  int
	y     int
	value int
}

type Symbol struct {
	x     int
	xend  int
	y     int
	value string
}

func locatePositions(reader utils.Reader, numReg, symbReg *regexp.Regexp) ([]Number, []Symbol) {
	var i, lineInd int
	numbers := make([]Number, 0)
	symbols := make([]Symbol, 0)

	for line := range reader.Stream() {
		// find every number's position (line,column) and value
		nums := numReg.FindAllString(line, -1)
		numsInd := numReg.FindAllStringIndex(line, -1)

		for i = 0; i < len(nums); i++ {
			val := utils.ToInt(nums[i])
			numbers = append(numbers, Number{
				x:     numsInd[i][0],
				xend:  numsInd[i][1],
				y:     lineInd,
				value: val,
			})
		}

		// find every symbol's position (line,column) and value
		symbs := symbReg.FindAllString(line, -1)
		symbsInd := symbReg.FindAllStringIndex(line, -1)

		for i = 0; i < len(symbs); i++ {
			symbols = append(symbols, Symbol{
				x:     symbsInd[i][0],
				xend:  symbsInd[i][1],
				y:     lineInd,
				value: symbs[i],
			})
		}

		lineInd++
	}

	return numbers, symbols
}

func (s Solver) Day_03(part int, reader utils.Reader) int {

	numReg := regexp.MustCompile(`[0-9]+`)
	symbReg := regexp.MustCompile(`[^\w\s.]`)

	numbers, symbols := locatePositions(reader, numReg, symbReg)

	for _, num := range numbers {
		logger.Debugf("Number: %+v\n", num)
	}
	for _, symb := range symbols {
		logger.Debugf("Symbol: %+v\n", symb)
	}

	switch part {
	case 1:
		var sum int

		for _, num := range numbers {
			// locate numbers adjacent to symbols in x,y directions
			for _, sym := range symbols {
				if (sym.x >= num.x-1 && sym.xend <= num.xend+1) && (sym.y >= num.y-1 && sym.y <= num.y+1) {
					sum += num.value
				}
			}
		}
		return sum

	case 2:
		var gearRatios int

		for _, sym := range symbols {
			found := make([]int, 0)

			// locate numbers adjacent to symbols in x,y directions
			// and append the numbers value to a list per symbol
			for _, num := range numbers {
				if (sym.x >= num.x-1 && sym.xend <= num.xend+1) && (sym.y >= num.y-1 && sym.y <= num.y+1) {
					found = append(found, num.value)
				}
			}

			// if only two adjacent numbers exist for a symbol, multiply the numbers to find the gearRatio
			if len(found) == 2 {
				gearRatios += found[0] * found[1]
			}

		}

		return gearRatios
	default:
		// shouldn't reach here
		return -1
	}

}
