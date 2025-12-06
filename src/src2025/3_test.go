package src2025

import (
	"reflect"
	"testing"
)

func Test_locateBestMax(t *testing.T) {
	tests := []struct {
		name        string
		arr         []int
		totalDigits int
		results     []int
	}{
		{
			name:        "2 digits",
			arr:         []int{5, 1, 4, 5, 6},
			totalDigits: 2,
			results:     []int{5, 6},
		},
		{
			name:        "2 digits - first iteration leaves only 1 digit",
			arr:         []int{5, 1, 4, 7, 6},
			totalDigits: 2,
			results:     []int{7, 6},
		},
		{
			name:        "3 digits",
			arr:         []int{5, 1, 4, 5, 6},
			totalDigits: 3,
			results:     []int{5, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := make([]int, 0)
			locateBestMax(tt.arr, tt.totalDigits, &results)

			if !reflect.DeepEqual(results, tt.results) {
				t.Errorf("locateBestMax(%v/%d) = %v, want = %v", tt.arr, tt.totalDigits, results, tt.results)
			}
		})
	}
}
