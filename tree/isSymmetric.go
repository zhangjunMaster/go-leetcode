package tree

/**
 * 树的操作最好的是用递归
 * 函数调用函数本身，在调用本身的时候，传入t.left t.right
 * 函数中传入左树和右树，判断左树的left值和右树的right值，左右的right值和右树的left值
 * 不要局限于前中后序遍历
 * 前中后 DFS 主要的特性是深度优先，总是不停的往下找，走到没路才罢休。
 * 层序是 BFS 从root开始扩展，每一层都是精密的搜索完整了才下一个
 */

/*
 * isMirror中是比较t1的left和t2的right，而不是t1.left和t1.right,递归下来会一直像劈叉似的延伸
 * 深度是对单个树进行递归，从而使得左边一直到左边 left -> left -> left
   状态函数：isMirror(left, right)
*/
func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
