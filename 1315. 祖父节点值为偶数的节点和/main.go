package main

import "fmt"
//
//给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：
//
//该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
//如果不存在祖父节点值为偶数的节点，那么返回 0 。
//
//输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//输出：18
//解释：图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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
