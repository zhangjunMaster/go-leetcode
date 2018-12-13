package tree

func PreorderTraversal(root *TreeNode) []int {
	results := []int{}
	if root == nil {
		return results
	}
	results = append(results, root.Val)
	results = append(results, PreorderTraversal(root.Left)...)
	results = append(results, PreorderTraversal(root.Right)...)
	return results
}
