package p0138

import (
	"testing"
)

// buildList constructs a linked list from a slice of [val, randomIndex] pairs.
// randomIndex of -1 means the random pointer is nil.
func buildList(pairs [][2]int) *Node {
	if len(pairs) == 0 {
		return nil
	}
	nodes := make([]*Node, len(pairs))
	for i, p := range pairs {
		nodes[i] = &Node{Val: p[0]}
	}
	for i := range nodes {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
		ri := pairs[i][1]
		if ri != -1 {
			nodes[i].Random = nodes[ri]
		}
	}
	return nodes[0]
}

// toSlice converts a linked list to [][2]int{val, randomIndex} for comparison.
func toSlice(head *Node) [][2]int {
	var nodes []*Node
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	index := make(map[*Node]int, len(nodes))
	for i, n := range nodes {
		index[n] = i
	}
	result := make([][2]int, len(nodes))
	for i, n := range nodes {
		ri := -1
		if n.Random != nil {
			ri = index[n.Random]
		}
		result[i] = [2]int{n.Val, ri}
	}
	return result
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name  string
		input [][2]int
	}{
		{
			name:  "example 1",
			input: [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			name:  "example 2",
			input: [][2]int{{1, 1}, {2, 1}},
		},
		{
			name:  "single node no random",
			input: [][2]int{{1, -1}},
		},
		{
			name:  "empty list",
			input: [][2]int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := buildList(tc.input)
			cloned := copyRandomList(original)

			got := toSlice(cloned)
			want := tc.input

			if len(got) != len(want) {
				t.Fatalf("length mismatch: got %d, want %d", len(got), len(want))
			}
			for i := range want {
				if got[i] != want[i] {
					t.Errorf("node %d: got %v, want %v", i, got[i], want[i])
				}
			}

			// Verify deep copy — no shared pointers
			origNodes := make(map[*Node]bool)
			for cur := original; cur != nil; cur = cur.Next {
				origNodes[cur] = true
			}
			for cur := cloned; cur != nil; cur = cur.Next {
				if origNodes[cur] {
					t.Errorf("cloned list shares a node pointer with the original")
				}
			}
		})
	}
}
