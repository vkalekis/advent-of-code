package src2024

import (
	"regexp"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

// linearEquation represents a linear equation with variables x,y in the form of:
// a*x+b*y=c
type linearEquation struct {
	a, b, c int
}

type systemOfEquations struct {
	eq1, eq2 linearEquation
}

type linearEquationSolution struct {
	x, y int
}

func (s systemOfEquations) solve() (linearEquationSolution, bool) {

	sol := linearEquationSolution{}

	// using cramer's rule
	// a1*x + b1*y = c1
	// a2*x + b2*y = c2
	//
	//     | a1 b1 |
	// D = | a2 b2 | = a1*b2 - a2*b1
	d := s.eq1.a*s.eq2.b - s.eq1.b*s.eq2.a

	if d == 0 {
		// no unique solution
		return sol, false
	} else {
		//      | c1 b1 |
		// Dx = | c2 b2 | = c1*b2 - c2*b1
		//
		//      | a1 c1 |
		// Dy = | a2 c2 | = a1*c2 - a2*c1
		//
		// solution: x = Dx/D, y = Dy/D

		dx := s.eq1.c*s.eq2.b - s.eq2.c*s.eq1.b
		dy := s.eq1.a*s.eq2.c - s.eq2.a*s.eq1.c

		if dx%d != 0 || dy%d != 0 {
			// the result is not an integer, the solutios represents integer button
			// presses -> return invalid solution
			return sol, false
		}
		sol.x = dx / d
		sol.y = dy / d
	}

	return sol, true
}

func constructSystemsOfEquations(reader utils.Reader, conversionError bool) []systemOfEquations {
	systemOfEqs := make([]systemOfEquations, 0)

	system := systemOfEquations{
		eq1: linearEquation{},
		eq2: linearEquation{},
	}

	for line := range reader.Stream() {
		if strings.HasPrefix(line, "Button A:") {
			re := regexp.MustCompile(`X\+(\d+)`)

			match := re.FindStringSubmatch(line)
			if len(match) > 1 {
				system.eq1.a = utils.ToInt(match[1])
			}

			re = regexp.MustCompile(`Y\+(\d+)`)

			match = re.FindStringSubmatch(line)
			if len(match) > 1 {
				system.eq2.a = utils.ToInt(match[1])
			}
		} else if strings.HasPrefix(line, "Button B:") {
			re := regexp.MustCompile(`X\+(\d+)`)

			match := re.FindStringSubmatch(line)
			if len(match) > 1 {
				system.eq1.b = utils.ToInt(match[1])
			}

			re = regexp.MustCompile(`Y\+(\d+)`)

			match = re.FindStringSubmatch(line)
			if len(match) > 1 {
				system.eq2.b = utils.ToInt(match[1])
			}

		} else if strings.HasPrefix(line, "Prize:") {
			re := regexp.MustCompile(`X\=(\d+)`)

			match := re.FindStringSubmatch(line)
			if len(match) > 1 {

				system.eq1.c = utils.ToInt(match[1])
				if conversionError {
					system.eq1.c += 10000000000000
				}
			}

			re = regexp.MustCompile(`Y\=(\d+)`)

			match = re.FindStringSubmatch(line)
			if len(match) > 1 {
				system.eq2.c = utils.ToInt(match[1])
				if conversionError {
					system.eq2.c += 10000000000000
				}
			}
		} else {
			systemOfEqs = append(systemOfEqs, system)
		}
	}

	// append the final system (no new line in the end)
	systemOfEqs = append(systemOfEqs, system)

	return systemOfEqs
}

func (s *Solver2024) Day_13(part int, reader utils.Reader) int {

	conversionError := false
	switch part {
	case 1:
	case 2:
		conversionError = true
	}

	systemsOfEqs := constructSystemsOfEquations(reader, conversionError)

	totalTokens := 0
	for _, systemOfEqs := range systemsOfEqs {
		sol, found := systemOfEqs.solve()
		logger.Infof("System: %+v, Sol: %+v, found: %t, tokens: %d",
			systemOfEqs, sol, found, 3*sol.x+sol.y)
		if found {
			totalTokens += 3*sol.x + sol.y
		}
	}

	return totalTokens
}
