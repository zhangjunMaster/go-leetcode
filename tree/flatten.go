package tree

/**
 * 1.递归有个出口
 * 2.边界条件 罗列好
 * 3.树有两个树，对这个两个树 root.Left = ... 和root.Right = ....做处理
	 递归是有返回值的
	 如果没有返回值，也是使用 这种root.Right = root.Left操作
 * 4.从出口开始算,对root.Left 和 root.Right做示意图
	采用后续是因为 2->3->1
	中续 2->1->3
	前序 1->2->3
	  1
	2	3
	tmp := 1.Right //3
	1.Right = 1.Left //2
	1.Left=nil
	tmp 放哪儿？
	tmp放在1.Right的最后面，则是for 遍历到 right为空
	如果有保留的则设置 tmp
*/
// 给定一个二叉树，原地将它展开为链表。

func Flatten(root *TreeNode) {
	// 出口
	if root == nil {
		return
	}
	// 后序遍历
	Flatten(root.Left)
	Flatten(root.Right)
	// 保留root.Right
	tmp := root.Right
	// 将左子树过载右子树上
	root.Right = root.Left
	// 删除左子树
	root.Left = nil
	// 找到右子树的最底值，将tmp挂在上面
	// 这个位置是与最简单的相差异的
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}
