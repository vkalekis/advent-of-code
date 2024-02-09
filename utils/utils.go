package utils

import (
	"encoding/json"
	"os"
	"strings"
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
