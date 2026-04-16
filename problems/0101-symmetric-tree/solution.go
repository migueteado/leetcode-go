package p0101

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	// Both are nil, is the end of the leaf in both cases
	if left == nil && right == nil {
		return true
	}
	// Only one is nil, is the end of only one leaf, subtrees differ
	if left == nil || right == nil {
		return false
	}
	// Both values have to be the same
	return left.Val == right.Val &&
		// And then we continue comparing in mirror
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return isMirror(root.Left, root.Right)
}
