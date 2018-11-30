package main

/**
 * 左子树 右子树操作，将问题转换为左右子树的比较
 * DPS 深度递归是，函数里分left和fight，但递归的参数是传左子或右子，或双子
 * 这样 这个树就能递归到，因为传入树，分开叉
 * 递归函数的参数：传入树（传几个树），保留的值
 * 返回的值：就是计算的值
 * 状态初始值：递归是栈操作，是从最底向上操作,就是初始值
 * 状态转移：
		   包括递归参数转移
		   返回值转移
*/
// 传参，树
// 返回值是深度
func maxDepth(root *TreeNode) int {
	// 空树深度为0
	if root == nil {
		return 0
	}
	// 叶子深度是1，初始状态值
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// 状态转移 max(leftDepth, rightDepth)
	leftDepth := maxDepth(root.Left) + 1
	rightDepth := maxDepth(root.Right) + 1
	if leftDepth < rightDepth {
		return rightDepth
	} else {
		return leftDepth
	}
}
