package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// root的左右树为空时，只计算一边树
	if root.Left == nil {
		rightDepth := minDepth(root.Right) + 1
		return rightDepth
	}

	if root.Right == nil {
		leftDepth := minDepth(root.Left) + 1
		return leftDepth
	}
	// root的左右树不为空时
	leftDepth := minDepth(root.Left) + 1
	rightDepth := minDepth(root.Right) + 1
	if leftDepth > rightDepth {
		return rightDepth
	} else {
		return leftDepth
	}
}
