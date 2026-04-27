package p0590

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func postorderTraversal(slice *[]int, node *Node) {
	if node == nil {
		return
	}

	for i := 0; i < len(node.Children); i++ {
		postorderTraversal(slice, node.Children[i])
	}

	*slice = append(*slice, node.Val)
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	var slice []int

	postorderTraversal(&slice, root)

	return slice
}
