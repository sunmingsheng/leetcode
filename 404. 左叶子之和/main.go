package main

import "fmt"

//计算给定二叉树的所有左叶子之和。
//
//示例：
//
//3
/// \
//9  20
///  \
//15   7
//
//在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-left-leaves
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(sumOfLeftLeaves(root))
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	print(root, &result)
	return result
}

func print(node *TreeNode, result *int) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		*result = *result + node.Left.Val
	}
	print(node.Left, result)
	print(node.Right, result)
}
