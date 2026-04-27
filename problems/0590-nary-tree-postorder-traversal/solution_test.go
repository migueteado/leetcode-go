package p0590

import (
	"testing"
)

func TestPostorder(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		expected []int
	}{
		{
			name: "example 1: [1,null,3,2,4,null,5,6]",
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 3,
						Children: []*Node{
							{Val: 5},
							{Val: 6},
						},
					},
					{Val: 2},
					{Val: 4},
				},
			},
			expected: []int{5, 6, 3, 2, 4, 1},
		},
		{
			name:     "example 2: empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "example 3: single node",
			root:     &Node{Val: 1},
			expected: []int{1},
		},
		{
			name: "example 4: linear chain",
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 2,
						Children: []*Node{
							{Val: 3},
						},
					},
				},
			},
			expected: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := postorder(tt.root)
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
