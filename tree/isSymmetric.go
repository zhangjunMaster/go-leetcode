package main

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

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

/**
 * 横向的逐层查找
 * 边界条件：【】内的*TreeNode为空，或者【】为nil
 * 将每一层的节点都放到数组中，遍历这个数组，找到节点的left和right，再将这个放入到新数组中
 */
func getVal(nodes []*TreeNode, result [][]int) [][]int {
	var nodeVal []int
	// 新数组
	var newNodes []*TreeNode
	if len(nodes) == 0 {
		return result
	}
	// 遍历旧数组
	for _, v := range nodes {
		if v == nil {
			return result
		}
		if v.Left != nil {
			newNodes = append(newNodes, v.Left)
		}
		if v.Right != nil {
			newNodes = append(newNodes, v.Right)
		}
		nodeVal = append(nodeVal, v.Val)
	}
	result = append(result, nodeVal)
	return getVal(newNodes, result)
}

func levelOrder(root *TreeNode) [][]int {
	var nodes []*TreeNode
	var result [][]int
	nodes = append(nodes, root)
	result = getVal(nodes, result)
	return result
}

func mian() {

}
