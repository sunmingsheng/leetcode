package main

import "fmt"

//给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(increasingBST(root))
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	s := []int{}
	inOrder(root, &s)
	head := &TreeNode{
		Val:s[0],
	}
	currentNode := head
	for _, value := range s[1:] {
		newNode := &TreeNode{
			Val:value,
		}
		currentNode.Right = newNode
		currentNode = newNode
	}
	return head
}

func inOrder(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, s)
	*s = append(*s, node.Val)
	inOrder(node.Right, s)
}
