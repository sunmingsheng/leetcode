package main

import "fmt"

//如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
//
//只有给定的树是单值二叉树时，才返回 true；否则返回 false。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(isUnivalTree(root))
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result := root.Val
	return print(root, result)
}

func print(node *TreeNode, result int) bool {
	if node == nil {
		return true
	}
	if node.Val != result {
		return false
	}
	return print(node.Left, result) && print(node.Right, result)
}
