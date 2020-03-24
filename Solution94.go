package main

import "fmt"


//给定一个二叉树，返回它的中序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//  1
//   \
//     2
///   /
//  3
//
//输出: [1,3,2]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	 head := &TreeNode{Val:1}
	 fmt.Println(inorderTraversal(head))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	results := []int{}
	printNode(root, &results)
	return results
}

func printNode(node *TreeNode, results *[]int) {
	if node == nil {
		return
	}
	printNode(node.Left, results)
	*results = append(*results, node.Val)
	printNode(node.Right, results)
}