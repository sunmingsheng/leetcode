package main

import "fmt"

//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
//输入: [1,null,2,3]
//1
//\
//2
///
//3
//
//输出: [1,2,3]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{Val:1}
	fmt.Println(preorderTraversal(head))
}

func preorderTraversal(root *TreeNode) []int {
	s := []int{}
	preOrder(root, &s)
	return s
}

func preOrder(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	*s = append(*s, node.Val)
	preOrder(node.Left, s)
	preOrder(node.Right, s)
}
