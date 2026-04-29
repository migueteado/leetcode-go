package p0287

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: [1,3,4,2,2]",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "example 2: [3,1,3,4,2]",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name:     "duplicate at start: [1,1,2,3]",
			nums:     []int{1, 1, 2, 3},
			expected: 1,
		},
		{
			name:     "duplicate at end: [2,3,4,1,4]",
			nums:     []int{2, 3, 4, 1, 4},
			expected: 4,
		},
		{
			name:     "invalid input out of range: [1,5,2]",
			nums:     []int{1, 5, 2},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDuplicate(tt.nums)
			if result != tt.expected {
				t.Fatalf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
