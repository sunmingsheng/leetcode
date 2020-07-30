package main

import (
	"fmt"
	"math"
	"sort"
)

//给定一个非空二叉树，返回其最大路径和。
//
//本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//示例 1:
//
//输入: [1,2,3]
//
//1
/// \
//2   3
//
//输出: 6
//示例 2:
//
//输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1, Left:&TreeNode{Val:3}, Right:&TreeNode{Val:1}}
	fmt.Println(maxPathSum(root))
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := math.MinInt64
	travel(root, &result)
	return result
}

func travel(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	a := node.Val
	b := travel(node.Right, maxSum)
	c := travel(node.Left, maxSum)
	d := a + b
	e := a + c
	f := a + b + c
	maxValue := max([]int{a,d,e,f})
	if maxValue > *maxSum {
		*maxSum = maxValue
	}
	node.Val = max([]int{a,d,e})
	return node.Val
}

func max(nums []int) int {
	length := len(nums)
	if length > 0 {
		sort.Ints(nums)
		return nums[length-1]
	}
	return 0
}
