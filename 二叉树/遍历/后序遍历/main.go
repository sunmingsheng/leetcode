package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	n := &TreeNode{Val:0}
	fmt.Println(postorderTraversal(n))
}

func postorderTraversal(root *TreeNode) []int {
	s := &[]int{}
	print(root, s)
	return *s
}

func print(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		print(node.Left, s)
	}
	if node.Right != nil {
		print(node.Right, s)
	}
	*s = append(*s, node.Val)
}