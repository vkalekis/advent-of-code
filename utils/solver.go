package utils

import (
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type Solution func(int, Reader) int

type Solver struct {
	year           int
	day            int
	part           int
	filepath       string
	reader         Reader
	logger         *zap.SugaredLogger
	available_days map[string]Solution
}

func NewSolver(year, day, part int, filepath string, example bool, logger *zap.SugaredLogger, available_days map[string]Solution) *Solver {
	if filepath == "" {
		filepath = fmt.Sprintf("data/%d/input%s", year, zfill(strconv.Itoa(day), 2))
	}
	if example {
		filepath = fmt.Sprintf("%s_example", filepath)
	}

	fr := NewFileReader(filepath)
	go fr.Read()

	return &Solver{
		year:           year,
		day:            day,
		part:           part,
		filepath:       filepath,
		reader:         fr,
		logger:         logger,
		available_days: available_days,
	}
}

func zfill(input string, width int) string {
	return fmt.Sprintf("%0"+fmt.Sprint(width)+"s", input)
}

func (s *Solver) Solve() error {

	day_specifier := fmt.Sprintf("%d_%s", s.year, zfill(strconv.Itoa(s.day), 2))

	if solution, ok := s.available_days[day_specifier]; ok {
		s.logger.Infof("Solving day %d - part %d using file %s", s.day, s.part, s.filepath)

		startTime := time.Now()
		res := solution(s.part, s.reader)
		endTime := time.Since(startTime)

		s.logger.Infof("Resulf of day %d - part %d : %d", s.day, s.part, res)
		s.logger.Infof("Took %v", endTime)
		return nil
	} else {
		return fmt.Errorf("provided day not in available days map")
	}
}
