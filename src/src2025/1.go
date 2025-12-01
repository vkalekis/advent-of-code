package src2025

import (
	"strconv"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

const (
	left  = "L"
	right = "R"
)

type safe struct {
	current, max          int
	zeroStops, zeroPasses int
}

func newSafe(currnent, max int) *safe {
	logger.Debugf("Safe initial state: cur=%d/max=%d", currnent, max)
	return &safe{
		current: currnent,
		max:     max,
	}
}

func (s *safe) turn(instruction string) {
	var dir string
	var steps int

	dir = string(instruction[0])
	if dir != left && dir != right {
		return
	}
	steps, err := strconv.Atoi(string(instruction[1:]))
	if err != nil {
		return
	}

	switch dir {
	case left:
		s.current -= steps

		// if the current is negative or zero, then we have crossed zero (most probably!)
		if s.current <= 0 {
			s.zeroPasses += -s.current/s.max + 1

			// and we reset
			//  5 + L10 ->  -5 ->  -5 + 100 = 95
			// and we need to do it -s.current/s.max + 1 times to wrap around to the positives
			// -257 -> we must loop (2+1) = 3 times   ( 2 = 257 div 100)

			// (!) but if we started at 0, we double count one transition
			// 0 + L10 -> -10 -> 10 div 100 + 1 = 1 where as we didn't do any transitions
			if s.current+steps == 0 {
				s.zeroPasses--
			}

			s.current += s.max * (-s.current/s.max + 1)
			s.current = s.current % s.max

		}

	case right:
		s.current += steps
		// if the current is over max, then we have crossed zero...
		if s.current >= s.max {
			// ...as many times as the div...
			// 50 + 5001 = 5501 -> we have crossed it 5501 div 100 = 55 times (!)
			s.zeroPasses += s.current / s.max
		}
		// ...and we reset
		// 98 + R10 -> 108 -> 108 % 100 = 8
		s.current = s.current % s.max
	}

	if s.current == 0 {
		s.zeroStops++
	}

	logger.Debugf("Safe after %s: %d (stops: %d, passes: %d)", instruction, s.current, s.zeroStops, s.zeroPasses)
}

func (s *Solver) Day_01(part int, reader utils.Reader) int {

	safe := newSafe(50, 100)

	for line := range reader.Stream() {
		safe.turn(line)
	}

	switch part {
	case 1:
		return safe.zeroStops
	case 2:
		return safe.zeroPasses
	}
	return -1
}
