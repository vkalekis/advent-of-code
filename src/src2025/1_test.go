package src2025

import (
	"testing"

	"github.com/vkalekis/advent-of-code/pkg/logger"
)

func Test_Safe_Turn(t *testing.T) {
	tests := []struct {
		name         string
		instruction  string
		current, max int

		wantCurrent, wantZeroStops, wantZeroPasses int
	}{
		{
			name:           "standard right",
			instruction:    "R10",
			current:        5,
			max:            100,
			wantCurrent:    15,
			wantZeroStops:  0,
			wantZeroPasses: 0,
		},
		{
			name:           "standard left",
			instruction:    "L20",
			current:        65,
			max:            100,
			wantCurrent:    45,
			wantZeroStops:  0,
			wantZeroPasses: 0,
		},
		{
			name:           "right that cross 0",
			instruction:    "R40",
			current:        80,
			max:            100,
			wantCurrent:    20,
			wantZeroStops:  0,
			wantZeroPasses: 1,
		},
		{
			name:           "right that hits 0",
			instruction:    "R20",
			current:        80,
			max:            100,
			wantCurrent:    0,
			wantZeroStops:  1,
			wantZeroPasses: 1,
		},
		{
			name:           "left that cross 0",
			instruction:    "L80",
			current:        40,
			max:            100,
			wantCurrent:    60,
			wantZeroStops:  0,
			wantZeroPasses: 1,
		},
		{
			name:           "left that hits 0",
			instruction:    "L40",
			current:        40,
			max:            100,
			wantCurrent:    0,
			wantZeroStops:  1,
			wantZeroPasses: 1,
		},
		{
			name:           "1 loop right",
			instruction:    "R120",
			current:        20,
			max:            100,
			wantCurrent:    40,
			wantZeroStops:  0,
			wantZeroPasses: 1,
		},
		{
			name:           "2 loops right",
			instruction:    "R240",
			current:        20,
			max:            100,
			wantCurrent:    60,
			wantZeroStops:  0,
			wantZeroPasses: 2,
		},
		{
			name:           "3 loops right stop at 0",
			instruction:    "R380",
			current:        20,
			max:            100,
			wantCurrent:    0,
			wantZeroStops:  1,
			wantZeroPasses: 4,
		},
		{
			name:           "1 loop left",
			instruction:    "L150",
			current:        40,
			max:            100,
			wantCurrent:    90,
			wantZeroStops:  0,
			wantZeroPasses: 2,
		},
		{
			name:           "2 loops left",
			instruction:    "L200",
			current:        20,
			max:            100,
			wantCurrent:    20,
			wantZeroStops:  0,
			wantZeroPasses: 2,
		},
		{
			name:           "3 loops left stop at 0",
			instruction:    "L320",
			current:        20,
			max:            100,
			wantCurrent:    0,
			wantZeroStops:  1,
			wantZeroPasses: 4,
		},
		{
			name:           "start at 0 go right",
			instruction:    "R10",
			current:        0,
			max:            100,
			wantCurrent:    10,
			wantZeroStops:  0,
			wantZeroPasses: 0,
		},
		{
			name:           "start at 0 go left",
			instruction:    "L10",
			current:        0,
			max:            100,
			wantCurrent:    90,
			wantZeroStops:  0,
			wantZeroPasses: 0,
		},
		{
			name:           "arrived at exact multiple",
			instruction:    "L220",
			current:        20,
			max:            100,
			wantCurrent:    0,
			wantZeroStops:  1,
			wantZeroPasses: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = logger.NewLogger("debug")
			safe := newSafe(tt.current, tt.max)
			safe.turn(tt.instruction)

			if tt.wantCurrent != safe.current {
				t.Errorf("expected wantCurrent=%d, got %d", tt.wantCurrent, safe.current)
			}
			if tt.wantZeroStops != safe.zeroStops {
				t.Errorf("expected wantZeroStops=%d, got %d", tt.wantZeroStops, safe.zeroStops)
			}
			if tt.wantZeroPasses != safe.zeroPasses {
				t.Errorf("expected wantZeroPasses=%d, got %d", tt.wantZeroPasses, safe.zeroPasses)
			}
		})
	}
}
