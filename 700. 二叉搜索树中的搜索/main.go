package main

import "fmt"

//给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{Val:1}
	fmt.Println(searchBST(head, 1))
}


func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node != nil {
		if node.Val == val {
			return node
		}
		if node.Val > val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
