package tree

/**
 * 给定一个二叉树和一个目标和，
 * 判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 * 树的递归，是传入什么，返回什么
 */

type Path struct {
	path []int
	sum  int
}

// 这种是从叶子向顶点，不符合要求
// 递归的会又有个状态转移 p.path = append(p.path, root.Val)
func sumVal2(root *TreeNode) []*Path {
	var result []*Path
	if root == nil {
		return []*Path{&Path{[]int{0}, 0}}
	}
	if root.Left == nil && root.Right == nil {
		return []*Path{&Path{[]int{root.Val}, root.Val}}
	}
	for _, p := range sumVal2(root.Left) {
		p.sum += root.Val
		p.path = append(p.path, root.Val)
		result = append(result, p)
	}
	for _, p := range sumVal2(root.Right) {
		p.sum += root.Val
		p.path = append(p.path, root.Val)
		result = append(result, p)
	}
	return result
}

func pathSum2(root *TreeNode, sum int) [][]int {
	var paths [][]int
	result := sumVal2(root)
	for _, p := range result {
		if p.sum == sum {
			paths = append(paths, p.path)
		}
	}
	return paths
}

func sumVal3() {

}
