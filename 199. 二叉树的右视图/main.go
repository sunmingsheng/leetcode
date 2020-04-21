package main

import "fmt"

//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//示例:
//
//输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//1            <---
///   \
//2     3         <---
//\     \
//5     4       <---
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(rightSideView(root))
}

func rightSideView(root *TreeNode) []int {
	s := []int{}
	if root == nil {
		return s
	}
	maxLevel := 0
	m := make(map[int]int)
	print(root, 0, &maxLevel, &m)
	for i := 0; i <= maxLevel; i++ {
		s = append(s, m[i])
	}
	return s
}

func print(node *TreeNode, level int, maxLevel *int, m *map[int]int) {
	if node == nil {
		return
	}
	if level > *maxLevel {
		*maxLevel = level
	}
	(*m)[level] = node.Val
	print(node.Left, level + 1, maxLevel, m)
	print(node.Right, level + 1, maxLevel, m)
}
