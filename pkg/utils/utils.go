package utils

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
)

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func WriteMapToFile(data map[string]interface{}, filename string) error {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetCol(lines map[int][]string, col, maxRows int) []string {
	line := make([]string, 0)
	for r := 0; r < maxRows; r++ {
		line = append(line, lines[r][col])
	}
	return line
}

func ToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		logger.Fatalf("error during parsing %s: %v", s, err)
	}
	return v
}
