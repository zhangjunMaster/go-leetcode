package tree

/*
 * 怎样导出两个值？----暂时不能导出
 * 怎样对左右树保存 -----如果是数值使用数组
 * 树分左右树，-----因此 root.left和root.right都要递归到
 * ------ 一个递归只有一个结果
 * 以前的都是返回一个值，现在是两个值，因此采用了数组，但是返回的两个值会叠加，产生更多的值
 * 怎样解决？
 * 递归产生两个结果以后会裂变，所以一般是返回一个结果，但是可以传入两个值
 * func hasPathSumHelper(root *TreeNode,currentSum int, sum int)bool
 * 如果使用引用类型的数组，最后不用返回值，而是使用指针的传参
 * 这种累加的可以用现有的值
 */

func sumVal(root *TreeNode) *[]int {
	var result []int
	if root == nil {
		return &result
	}
	if root.Left == nil && root.Right == nil {
		return &[]int{root.Val}
	}
	for _, v := range *sumVal(root.Left) {
		result = append(result, v+root.Val)
	}
	for _, v := range *sumVal(root.Right) {
		result = append(result, v+root.Val)
	}
	return &result
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sums := sumVal(root)
	for _, v := range *sums {
		if v == sum {
			return true
		}
	}
	return false
}
