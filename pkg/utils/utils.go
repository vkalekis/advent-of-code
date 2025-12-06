package utils

import (
	"encoding/json"
	"os"
	"slices"
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

// inefficient!
func CeilingDivision(a, b int) int {
	div := a / b
	if a%b != 0 {
		return div + 1
	}
	return div
}

// inefficient!
func FindFactors(n int) []int {
	factors := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}
	slices.Sort(factors)
	return factors
}

func MaxWithIndex(arr []int) (max int, idx int) {
	if len(arr) == 0 {
		return -1, -1
	}
	max = arr[0]
	idx = 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
			idx = i
		}
	}
	return
}
