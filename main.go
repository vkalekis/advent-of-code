package main

import (
	"flag"
	"fmt"

	"github.com/vkalekis/advent-of-code/src2023"
	"github.com/vkalekis/advent-of-code/utils"
)

var available_days map[string]utils.Solution
var solver2023 src2023.Solver2023

func main() {

	day := flag.Int("day", 1, "day number in AoC calendar")
	part := flag.Int("part", 1, "1/2: Run first or second part of the solution")
	level := flag.String("level", "info", "logger debug level")
	filepath := flag.String("file", "", "read specific filepath instead of the predefined format")
	example := flag.Bool("example", false, "use the example file")
	flag.Parse()

	if err := checkArgs(*day, *part, *level); err != nil {
		fmt.Errorf("Error during parsing of args: %v", err)
		return
	}

	logger, _ := utils.NewLogger(*level)

	solver2023 = src2023.NewSolver2023()

	initDaysMap()

	solver := utils.NewSolver(*day, *part, *filepath, *example, logger, available_days)

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
		"day_01": solver2023.Day_01,
		"day_02": solver2023.Day_02,
		"day_03": solver2023.Day_03,
		"day_04": solver2023.Day_04,
		// // "day_05": src2023.Day_05,
		"day_06": solver2023.Day_06,
		"day_07": solver2023.Day_07,
		"day_08": solver2023.Day_08,
		"day_09": solver2023.Day_09,
		"day_10": solver2023.Day_10,
		"day_11": solver2023.Day_11,
	}
}
