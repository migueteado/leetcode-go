package p0145

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(slice *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	postOrder(slice, node.Left)
	postOrder(slice, node.Right)
	*slice = append(*slice, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int

	postOrder(&result, root)

	return result
}
