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
	if len(instruction) < 2 {
		return
	}

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

	var delta int
	switch dir {
	case left:
		delta = -steps
	case right:
		delta = steps
	}

	new := s.current + delta

	// -400    -300    -200    -100    0    100    200    300
	//					                  |
	//                       Start is always here
	//                                                251
	//                                                251/100 -> 2 we have crossed 0 two times (100,200)
	//      -330
	//      -330/100 -> 3 WRONG! we have crossed 0 4 times (0,-100,-200,-300)
	//        we need to do floor
	//       (-330-100)/100 -> (330+100)/100 -> 4
	//
	// edge case when we start at 0 and loop left:
	//  0 -> -20  -> zeropasses = (20+100)/100 = 1 whereas we didn't do any passes

	switch dir {
	case left:
		s.zeroPasses += (-new + s.max) / s.max
		if s.current == 0 {
			s.zeroPasses--
		}
	case right:
		s.zeroPasses += new / s.max
	}

	//  251 % 100 -> (251+100) % 100 = 51 anyway
	// -330 % 100 = -30 X  -> -30+100 = 70 % 100 = 70
	s.current = (new%s.max + s.max) % s.max

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
