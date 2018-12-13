package tree

func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
			return root
		}
		root.Right = InsertIntoBST(root.Right, val)
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
			return root
		}
		root.Left = InsertIntoBST(root.Left, val)
	}
	return root
}
