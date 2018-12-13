package tree

import "fmt"

/**
binary TreeNode
1.最多有两个子树
2.左子树，右子树是有顺序的，次序不能颠倒
3.叶子
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var s []int

func NewTreeNode(Val int) *TreeNode {
	return &TreeNode{Val: Val, Left: nil, Right: nil}
}

// recursive call self
// TreeNode have the property that contains recursion

func (t *TreeNode) Add(Val int) *TreeNode {
	// stop recursive
	// if t == nil how to t = &TreeNode

	if t.Left == nil && t.Right == nil {
		//fmt.Println(Val, t)
		if Val < t.Val {
			fmt.Println("val2", Val)
			t.Left = NewTreeNode(Val)
		} else {
			t.Right = NewTreeNode(Val)
		}
	} else {
		// call self
		if Val < t.Val {
			if t.Left == nil {
				t.Left = NewTreeNode(Val)
			} else {
				t.Left.Add(Val)
			}
		} else {
			if t.Right == nil {
				t.Right = NewTreeNode(Val)
			} else {
				t.Right.Add(Val)
			}
		}

	}
	return t
}

func (t *TreeNode) PreSearch() {
	fmt.Println(t.Val)
	if t.Left == nil && t.Right == nil {
		return
	}
	t.Left.PreSearch()
	t.Right.PreSearch()
}

func (t *TreeNode) InOrderSearch() {
	if t == nil {
		return
	}

	t.Left.InOrderSearch()
	fmt.Println(t.Val)
	s = append(s, t.Val)
	t.Right.InOrderSearch()
}

func (t *TreeNode) PostSearch() {
	if t.Left == nil && t.Right == nil {
		return
	}
	t.Left.InOrderSearch()
	t.Right.InOrderSearch()
	fmt.Println(t.Val)
}

func sMin(s []int) (min int) {
	fmt.Println(len(s))
	min = s[1] - s[0]
	for i := 1; i < len(s); i++ {
		fmt.Println(i)
		if s[i]-s[i-1] < min {
			min = s[i] - s[i-1]
		}
	}
	return min
}
