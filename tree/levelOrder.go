package tree

/**
* 横向的逐层查找.将left和right都塞到一个数组中，对数组进行递归
* 边界条件：【】内的*TreeNode为空，或者【】为nil
* 将每一层的节点都放到数组中，遍历这个数组，找到节点的left和right，再将这个放入到新数组中
* 本次解决使用递归
* 通过 for (len(nodes) >0) {
     nodes = ...
} 实现 while式的循环
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
