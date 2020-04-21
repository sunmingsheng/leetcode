package main

import (
	"fmt"
	"math"
)

//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回它的最小深度  2.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1, Right:&TreeNode{Val:2}}
	fmt.Println(minDepth(root))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDep := math.MaxInt32
	print(root, 1, &minDep)
	return minDep
}

func print(node *TreeNode, depth int, minDep *int) {
	if node == nil {
		return
	}
	if node.Right == nil && node.Left == nil {
		if depth < *minDep {
			*minDep = depth
		}
		return
	}
	print(node.Left, depth+1, minDep)
	print(node.Right, depth+1, minDep)
}
