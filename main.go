package main

import (
	"flag"
	"fmt"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
	"github.com/vkalekis/advent-of-code/src/src2023"
	"github.com/vkalekis/advent-of-code/src/src2024"
	"github.com/vkalekis/advent-of-code/src/src2025"
)

var availableDays map[string]utils.Solution
var solver2023 src2023.Solver
var solver2024 src2024.Solver
var solver2025 src2025.Solver

var availableYears = []int{2023, 2024, 2025}

func main() {

	year := flag.Int("year", 2024, "year")
	day := flag.Int("day", 1, "day number in AoC calendar")
	part := flag.Int("part", 1, "1/2: Run first or second part of the solution")
	level := flag.String("level", "info", "logger debug level")
	filepath := flag.String("file", "", "read specific filepath instead of the predefined format")
	example := flag.Bool("example", false, "use the example file")
	flag.Parse()

	if err := checkArgs(*year, *day, *part, *level); err != nil {
		fmt.Printf("Error during parsing of args: %v", err)
		return
	}

	_ = logger.NewLogger(*level)

	solver2023 = src2023.NewSolver()
	solver2024 = src2024.NewSolver()
	solver2025 = src2025.NewSolver()

	initDaysMap()

	solver := utils.NewSolver(*year, *day, *part, *filepath, *example, availableDays)

	if err := solver.Solve(); err != nil {
		logger.Errorf("Error during solving: %v", err)
	}

}

func checkArgs(year, day, part int, level string) error {
	checkYear := func() (found bool) {
		for i := range availableYears {
			if availableYears[i] == year {
				found = true
				return
			}
		}
		return
	}()

	if !checkYear {
		return fmt.Errorf("provided year %d is not valid", year)
	}
	if day < 0 {
		return fmt.Errorf("provided day %d is not valid", day)
	}
	if part != 1 && part != 2 {
		return fmt.Errorf("provided solution part %d is not valid", part)
	}
	if level != "debug" && level != "info" && level != "warn" && level != "error" {
		return fmt.Errorf("provided logger level %s is not valid", level)
	}
	return nil
}

func initDaysMap() {
	availableDays = map[string]utils.Solution{
		"2023_01": solver2023.Day_01,
		"2023_02": solver2023.Day_02,
		"2023_03": solver2023.Day_03,
		"2023_04": solver2023.Day_04,
		// // "2023_day_05": src2023.Day_05,
		"2023_06": solver2023.Day_06,
		"2023_07": solver2023.Day_07,
		"2023_08": solver2023.Day_08,
		"2023_09": solver2023.Day_09,
		"2023_10": solver2023.Day_10,
		"2023_11": solver2023.Day_11,
		// "2023_13": solver2023.Day_13,
		"2023_14": solver2023.Day_14,
		"2023_15": solver2023.Day_15,
		// "2023_16": solver2023.Day_16,
		"2023_18": solver2023.Day_18,
		"2023_19": solver2023.Day_19,
		"2023_20": solver2023.Day_20,
		// ---
		"2024_01": solver2024.Day_01,
		"2024_02": solver2024.Day_02,
		"2024_03": solver2024.Day_03,
		"2024_04": solver2024.Day_04,
		"2024_05": solver2024.Day_05,
		"2024_06": solver2024.Day_06,
		"2024_07": solver2024.Day_07,
		"2024_08": solver2024.Day_08,
		"2024_09": solver2024.Day_09,
		"2024_10": solver2024.Day_10,
		"2024_11": solver2024.Day_11,
		"2024_12": solver2024.Day_12,
		"2024_13": solver2024.Day_13,
		"2024_14": solver2024.Day_14,
		// ---
		"2025_01": solver2025.Day_01,
		"2025_02": solver2025.Day_02,
		"2025_03": solver2025.Day_03,
		"2025_05": solver2025.Day_05,
	}
}
