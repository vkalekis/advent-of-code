package utils

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

type Solution func(int, Reader, *zap.SugaredLogger) int

type Solver struct {
	day            int
	part           int
	reader         Reader
	logger         *zap.SugaredLogger
	available_days map[string]Solution
}

func NewSolver(day int, part int, reader Reader, logger *zap.SugaredLogger, available_days map[string]Solution) *Solver {
	return &Solver{
		day:            day,
		part:           part,
		reader:         reader,
		logger:         logger,
		available_days: available_days,
	}
}

func zfill(input string, width int) string {
	return fmt.Sprintf("%0"+fmt.Sprint(width)+"s", input)
}

func (s *Solver) Solve() error {

	day_specifier := fmt.Sprintf("day_%s", zfill(strconv.Itoa(s.day), 2))

	if solution, ok := s.available_days[day_specifier]; ok {
		s.logger.Infof("Solving day %d - part %d", s.day, s.part)
		res := solution(s.part, s.reader, s.logger)
		s.logger.Infof("Resulf of day %d - part %d : %d", s.day, s.part, res)
		return nil
	} else {
		return fmt.Errorf("provided day not in available days map")
	}

}
