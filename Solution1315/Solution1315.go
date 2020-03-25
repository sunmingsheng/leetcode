package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	left := &TreeNode{Val:1}
	right := &TreeNode{Val:3}
	head := &TreeNode{Val:2}
	head.Left = left
	head.Right = right
	fmt.Println(sumEvenGrandparent(head))
}

var results = 0

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	findSunNodes(root)
	return results
}

func findSunNodes(node *TreeNode) {
	if node.Val % 2 == 0 {
		if node.Left != nil {
			if node.Left.Left != nil {
				results += node.Left.Left.Val
			}
			if node.Left.Right != nil {
				results += node.Left.Right.Val
			}
		}
		if node.Right != nil {
			if node.Right.Left != nil {
				results += node.Right.Left.Val
			}
			if node.Right.Right != nil {
				results += node.Right.Right.Val
			}
		}
	}
	if node.Left != nil {
		findSunNodes(node.Left)
	}
	if node.Right != nil {
		findSunNodes(node.Right)
	}
}
