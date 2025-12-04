package src2025

import (
	"context"
	"reflect"
	"testing"
)

func TestGenerateIds(t *testing.T) {
	tests := []struct {
		name     string
		req      repeatedBlockReq
		expected []int
	}{
		{
			name: "1-digit repeated IDs",
			req: repeatedBlockReq{
				blockSize: 1,
				blocks:    2,
			},
			expected: []int{
				11, 22, 33, 44, 55, 66, 77, 88, 99,
			},
		},
		// the other cases are too extreme to list
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ch := generateRepeated(ctx, tt.req)

			var got []int
			for id := range ch {
				got = append(got, id)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("generateIds(%v) = %v, want = %v", tt.req, got, tt.expected)
			}
		})
	}
}
