package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 递归中怎样将状态导出来,这里用了两个递归，一个递归获取值，一个递归根据上一个递归获取判断（状态）
/**
 * 所有的树的DPS的问题，都要化成左右树的问题
 * 关于状态，可能一个递归实现不了，就要分成两个，一个递归获取树的深度，一个比较深度
 * 递归只干一件事，返回值和参数，对于变量保存状态不好操作
 * 而动态规划不是递归，是用变量保存状态的
 */

func deepLength(root *TreeNode) int {

	// 空树
	if root == nil {
		return 0
	}
	// 初始值，叶子是
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// 左右树的深度
	leftDepth := deepLength(root.Left) + 1
	rightDepth := deepLength(root.Right) + 1
	// 比较左右树的深度
	if leftDepth == -1 || rightDepth == -1 {
		return -2
	}
	if abs(leftDepth, rightDepth) > 1 {
		return -2
	}
	// 通过对比左右树，得出最大的深度
	return max(leftDepth, rightDepth)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if deepLength(root) == -2 {
		return false
	}

	return true
}

// 只查一个root的深度
func deepLength2(root *TreeNode) int {

	// 空树
	if root == nil {
		return 0
	}
	// 初始值，叶子是
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// 左右树的深度
	leftDepth := deepLength(root.Left) + 1
	rightDepth := deepLength(root.Right) + 1
	// 通过对比左右树，得出最大的深度
	return max(leftDepth, rightDepth)
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 比较这个root的深度
	if abs(deepLength2(root.Left), deepLength2(root.Right)) > 1 {
		return false
	}
	// 递归比较下一个left和right的深度
	return isBalanced2(root.Left) && isBalanced2(root.Right)
}
