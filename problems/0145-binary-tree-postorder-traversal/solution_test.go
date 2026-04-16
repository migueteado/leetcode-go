package p0145

import (
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "example 1: [1,null,2,3]",
			root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{3, 2, 1},
		},
		{
			name:     "example 2: empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "example 3: single node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := postorderTraversal(tt.root)
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
