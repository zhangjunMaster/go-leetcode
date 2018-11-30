package main

/*
 * 怎样导出两个值？
 * 以前的都是返回一个值，现在是两个值，因此采用了数组，但是返回的两个值会叠加，产生更多的值
 * 怎样解决？
 * 传入数组，更能提高性能，不用每一个都要make一个数组
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
