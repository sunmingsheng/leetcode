package main

import (
	"sort"
)

//给你一棵二叉树，请你返回层数最深的叶子节点的和。
//
//
//
//
//输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//输出：15
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/deepest-leaves-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	
}

//层序遍历
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := make(map[int]int)
    m[0] = root.Val
	printTree(root, 1, m)
	s := []int{}
	for key := range m {
		s = append(s, key)
	}
	sort.Ints(s)
	return m[s[len(s) - 1]]
}

func printTree(node *TreeNode, level int, m map[int]int) {
	if node.Right == nil && node.Left == nil {
		return
	}
	if _, ok := m[level]; !ok {
		m[level] = 0
	}
	if node.Right != nil {
		m[level] += node.Right.Val
	}
	if node.Left != nil {
		m[level] += node.Left.Val
	}
	if node.Right != nil {
		printTree(node.Right, level+1, m)
	}
	if node.Left != nil {
		printTree(node.Left, level+1, m)
	}
}
