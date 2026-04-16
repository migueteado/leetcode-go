package p0101

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name: "example 1: symmetric [1,2,2,3,4,4,3]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3},
				},
			},
			expected: true,
		},
		{
			name: "example 2: not symmetric [1,2,2,null,3,null,3]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: false,
		},
		{
			name:     "nil root",
			root:     nil,
			expected: false,
		},
		{
			name:     "single node",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "different values",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSymmetric(tt.root)
			if result != tt.expected {
				t.Fatalf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
