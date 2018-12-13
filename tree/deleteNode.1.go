package tree

import "fmt"

/**
返回节点就可以了，会根据该节点的指针关系关联出整个树
*/
func DeleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		min := root.Right
		// 右子树的最小值在左子树
		//fmt.Println("min.Left.Left:", min.Left.Left)
		for min.Left != nil {
			fmt.Println("min1:", min.Val, min)
			min = min.Left
		}
		root.Val = min.Val
		root.Right = DeleteNode(root.Right, min.Val)
	} else if root.Val > key {
		root.Left = DeleteNode(root.Left, key)
	} else {
		root.Right = DeleteNode(root.Right, key)
	}
	return root
}

func Delete3(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		minVal := func(tmp *TreeNode) int {
			for tmp.Left != nil {
				tmp = tmp.Left
			}
			return tmp.Val
		}(root.Right)
		root.Val = minVal
		root.Right = Delete3(root.Right, minVal)
	} else if root.Val > key {
		root.Left = Delete3(root.Left, key)
	} else {
		root.Right = Delete3(root.Right, key)
	}
	return root
}
