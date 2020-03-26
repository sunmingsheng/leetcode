package main

import (
	"fmt"
)

//输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
//
//例如：
//
//给定二叉树 [3,9,20,null,null,15,7]，
//
//3
/// \
//9  20
///  \
//15   7
//返回它的最大深度 3
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var maxDeep = 0

func main() {
	var root *TreeNode
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	travelTree(root, 0)
	return maxDeep
}

func travelTree(node *TreeNode, deep int) {
	if node == nil {
		if deep > maxDeep {
			maxDeep = deep
		}
		return
	}
	travelTree(node.Left, deep + 1)
	travelTree(node.Right, deep + 1)
}
