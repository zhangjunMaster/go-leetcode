package main

import (
	"fmt"
	"go-leetcode/tree"
)

func main() {
	t := tree.NewTreeNode(1)
	s := []int{2, 2, 5, 5, 6, 6}
	for _, v := range s {
		t = t.Add(v)
	}
	arr := tree.PreorderTraversal(t)
	for _, v := range arr {
		fmt.Println("arr:", v)
	}
}
