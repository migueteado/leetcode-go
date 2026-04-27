package p0119

import (
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected []int
	}{
		{
			name:     "example 1: 3",
			numRows:  3,
			expected: []int{1, 3, 3, 1},
		},
		{
			name:     "example 2: 0",
			numRows:  0,
			expected: []int{1},
		},
		{
			name:     "example 3: 1",
			numRows:  1,
			expected: []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pascalsTriangle2(tt.numRows)
			if len(result) != len(tt.expected) {
				t.Fatalf("got %v, want %v", result, tt.expected)
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Fatalf("got %v, want %v", result, tt.expected)
				}
			}
		})
	}
}
