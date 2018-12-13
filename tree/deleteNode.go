package tree

import "fmt"

/**
* 返回节点就可以了，会根据该节点的指针关系关联出整个树
* 递归会有初始值
* 递归的会又有个状态转移 root.Right = DeleteNode(root.Right, key)
 */
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		fmt.Println("v0")
		return root
	}

	if root.Left == nil && root.Right == nil {
		fmt.Println("v1", root)
		if root.Val == key {
			root = nil
		}
		return root
	}

	if key > root.Val {
		fmt.Println("v11", root.Right)
		// 关键在这个位置，递归的生成新节点root.Right = 节点
		root.Right = DeleteNode(root.Right, key)
	} else if key < root.Val {
		fmt.Println("v12")
		root.Left = DeleteNode(root.Left, key)
	} else {
		if root.Left != nil && root.Right != nil {
			fmt.Println("v2")

			// 第一步是从右子树找最小值
			min := root.Right
			// 右子树的最小值在左子树
			//fmt.Println("min.Left.Left:", min.Left.Left)
			for min.Left != nil {
				fmt.Println("min1:", min.Val, min)
				min = min.Left
			}
			fmt.Println("min:", min.Val, min)
			root.Val = min.Val
			root.Right = DeleteNode(root.Right, min.Val)
		} else {
			fmt.Println("v3")
			// 新节点怎么关联到树,是在root.Right = DeleteNode(root.Right, key)
			// 每一个判断会return
			var newRoot *TreeNode
			if root.Left != nil {
				newRoot = root.Left
				root.Left = nil
			} else {
				newRoot = root.Right
				root.Right = nil
			}
			//root = nil
			return newRoot
		}
	}
	return root
}
