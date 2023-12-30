package main

import (
	"flag"
	"fmt"

	"github.com/vkalekis/advent-of-code/src2023"
	"github.com/vkalekis/advent-of-code/utils"
)

var available_days map[string]utils.Solution

func main() {

	day := flag.Int("day", 1, "day number in AoC calendar")
	part := flag.Int("part", 1, "1/2: Run first or second part of the solution")
	level := flag.String("level", "info", "logger debug level")
	filepath := flag.String("file", "", "read specific filepath instead of the predefined format")
	flag.Parse()

	if err := checkArgs(*day, *part, *level); err != nil {
		fmt.Errorf("Error during parsing of args: %v", err)
		return
	}

	logger, _ := utils.NewLogger(*level)

	initDaysMap()

	solver := utils.NewSolver(*day, *part, *filepath, logger, available_days)

	if err := solver.Solve(); err != nil {
		logger.Errorf("Error during solving: %v", err)
	}

}

func checkArgs(day, part int, level string) error {
	if day < 0 {
		return fmt.Errorf("provided day is not valid")
	}
	if part != 1 && part != 2 {
		return fmt.Errorf("provided solution part is not valid")
	}
	if level != "debug" && level != "info" && level != "warn" && level != "error" {
		return fmt.Errorf("provided logger level is not valid")
	}
	return nil
}

func initDaysMap() {
	available_days = map[string]utils.Solution{
		"day_01": src2023.Day_01,
		"day_02": src2023.Day_02,
	}
}
