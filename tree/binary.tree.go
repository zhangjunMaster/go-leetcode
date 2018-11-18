package main

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
	return &TreeNode{Val: Val, Left: &TreeNode{}, Right: &TreeNode{}}
}

// recursive call self
// TreeNode have the property that contains recursion

func (t *TreeNode) Add(Val int) bool {
	// stop recursive
	// if t == nil how to t = &TreeNode
	if t.Left == nil && t.Right == nil {
		t.Val = Val
		t.Left = &TreeNode{}
		t.Right = &TreeNode{}
		return true
	} else {
		// call self
		if Val < t.Val {
			t.Left.Add(Val)
		} else {
			t.Right.Add(Val)
		}
	}
	return false
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
	if t.Left == nil && t.Right == nil {
		return
	}
	t.Left.InOrderSearch()
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

func main() {
	fmt.Println("s:", s, &s)
	fmt.Println(s[0])
	t := NewTreeNode(10)
	t.Add(18)
	t.Add(7)
	t.Add(0)
	t.Add(30)
	t.InOrderSearch()
	fmt.Println(s)
	fmt.Println("min:", sMin(s))
}
