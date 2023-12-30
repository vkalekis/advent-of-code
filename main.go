package main

import (
	"flag"
	"fmt"

	"github.com/vkalekis/advent-of-code/src2023"
	"github.com/vkalekis/advent-of-code/utils"
)

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

func main() {

	day := flag.Int("day", 1, "day number in AoC calendar")
	part := flag.Int("part", 1, "1/2: Run first or second part of the solution")
	level := flag.String("level", "info", "logger debug level")
	flag.Parse()

	if err := checkArgs(*day, *part, *level); err != nil {
		fmt.Errorf("Error during parsing of args: %v", err)
		return
	}

	logger, _ := utils.NewLogger(*level)

	available_days := map[string]utils.Solution{
		"day_01": src2023.Day_01,
	}

	fr := utils.NewFileReader("data/2023/input01")

	solver := utils.NewSolver(*day, *part, fr, logger, available_days)

	go fr.Read()

	if err := solver.Solve(); err != nil {
		logger.Errorf("Error during solving: %v", err)
	}

}
