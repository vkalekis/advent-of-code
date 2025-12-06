package src2025

import (
	"reflect"
	"testing"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

func Test_mergeRanges(t *testing.T) {
	tests := []struct {
		name       string
		ranges     []utils.IdRange
		wantRanges []utils.IdRange
	}{
		{
			name: "no relation",
			ranges: []utils.IdRange{
				{Min: 0, Max: 5},
				{Min: 7, Max: 15},
			},
			wantRanges: []utils.IdRange{
				{Min: 0, Max: 5},
				{Min: 7, Max: 15},
			},
		},
		{
			name: "overlapping",
			ranges: []utils.IdRange{
				{Min: 1, Max: 5},
				{Min: 0, Max: 5},
			},
			wantRanges: []utils.IdRange{
				{Min: 0, Max: 5},
			},
		},
		{
			name: "extension left",
			ranges: []utils.IdRange{
				{Min: 10, Max: 50},
				{Min: 0, Max: 9},
			},
			wantRanges: []utils.IdRange{
				{Min: 0, Max: 50},
			},
		},
		{
			name: "extension right",
			ranges: []utils.IdRange{
				{Min: 0, Max: 9},
				{Min: 10, Max: 50},
			},
			wantRanges: []utils.IdRange{
				{Min: 0, Max: 50},
			},
		},
		{
			name: "some overlap",
			ranges: []utils.IdRange{
				{Min: 0, Max: 9},
				{Min: 5, Max: 50},
			},
			wantRanges: []utils.IdRange{
				{Min: 0, Max: 50},
			},
		},
		{
			name: "some overlap",
			ranges: []utils.IdRange{
				{Min: 5, Max: 20},
				{Min: 7, Max: 10},
			},
			wantRanges: []utils.IdRange{
				{Min: 5, Max: 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = logger.NewLogger("debug")
			got := mergeRanges(tt.ranges)
			if !reflect.DeepEqual(got, tt.wantRanges) {
				t.Errorf("mergeRanges() = %v, want %v", got, tt.wantRanges)
			}
		})
	}
}
