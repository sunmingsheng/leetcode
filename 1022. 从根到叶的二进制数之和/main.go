package main

import (
	"fmt"
	"math"
)

//给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
//
//对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
//
//以 10^9 + 7 为模，返回这些数字之和。
//
// 
//
//示例
//
//输入：[1,0,1,0,1,0,1]
//输出：22
//解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1,Left:&TreeNode{Val:0},Right:&TreeNode{Val:1}}
	fmt.Println(sumRootToLeaf(root))
}

var ss = [][]int{}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	s := []int{}
	print(root, s)
	//二进制转十进制
	for _, s := range ss {
		length := len(s)
		for key, value := range s {
			result += value * int(math.Pow(2, float64(length - key - 1)))
		}
	}
	return result
}

func print(node *TreeNode, s []int) {
	s = append(s, node.Val)
	if node.Right == nil && node.Left == nil {
		ss = append(ss, s)
		return
	}
	print(node.Left, s)
	print(node.Right, s)
}

