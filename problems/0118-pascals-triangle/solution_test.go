package p0118

import (
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:     "example 1: 5",
			numRows:  5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name:     "example 2: 1",
			numRows:  1,
			expected: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pascalTriangle(tt.numRows)
			if len(result) != len(tt.expected) {
				t.Fatalf("got %v, want %v", result, tt.expected)
			}
			for i := range result {
				if !reflect.DeepEqual(result[i], tt.expected[i]) {
					t.Fatalf("got %v, want %v", result, tt.expected)
				}
			}
		})
	}
}
