package main

import (
	"fmt"
	"math"
)

//给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
//
//
//示例：
//
//输入：
//
//1
//\
//3
///
//2
//
//输出：
//1
//
//解释：
//最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node := &TreeNode{Val:5,Left:&TreeNode{Val:4},Right:&TreeNode{Val:7}}
	fmt.Println(getMinimumDifference(node))
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := []int{}
	print(root, &s)
	result := math.MaxInt64
	length := len(s)
	for key, value := range s[:length - 1] {
		temp := s[key + 1] - value
		if  temp < result {
			result = temp
		}
	}
	return result
}

func print(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	print(node.Left, s)
	*s = append(*s, node.Val)
	print(node.Right, s)
}
