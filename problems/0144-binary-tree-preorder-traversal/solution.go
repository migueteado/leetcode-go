package p0144

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(slice *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	*slice = append(*slice, node.Val)
	preorder(slice, node.Left)
	preorder(slice, node.Right)
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var slice []int

	preorder(&slice, root)

	return slice
}
