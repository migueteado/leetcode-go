package p0083

import (
	"reflect"
	"testing"
)

func buildList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(slice))
	for i, s := range slice {
		nodes[i] = &ListNode{Val: s}
	}

	for i := range nodes {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
	}

	return nodes[0]
}

func toSlice(head *ListNode) []int {
	var result []int
	var curr *ListNode = head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "Test 1", input: []int{1, 1, 2}, expected: []int{1, 2}},
		{name: "Test 2", input: []int{1, 1, 2, 3, 3}, expected: []int{1, 2, 3}},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			head := buildList(testCase.input)
			output := deleteDuplicates(head)
			if !reflect.DeepEqual(toSlice(output), testCase.expected) {
				t.Errorf("Expected %v, got %v", testCase.expected, toSlice(output))
			}
		})
	}
}
