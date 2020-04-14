package main

import "fmt"

//给定一个二叉树，在树的最后一行找到最左边的值。
//
//示例 1:
//
//输入:
//
//2
/// \
//1   3
//
//输出:
//1
// 
//
//示例 2:
//
//输入:
//
//1
/// \
//2   3
///   / \
//4   5   6
///
//7
//
//输出:
//7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-bottom-left-tree-value
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(findBottomLeftValue(root))
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	level := 1
	maxLevel := 1
	m := make(map[int][]int)
	//层序遍历
	levelPrint(root, level, &maxLevel, &m)
	fmt.Println(m)
	return m[maxLevel][0]
}

func levelPrint(node *TreeNode, level int, maxLevel *int, m *map[int][]int) {
	if level > *maxLevel {
		*maxLevel = *maxLevel + 1
	}
	if _, ok := (*m)[level]; ok {
		(*m)[level] = append((*m)[level], node.Val)
	} else {
		(*m)[level] = []int{node.Val}
	}
	if node.Left != nil {
		levelPrint(node.Left, level + 1, maxLevel, m)
	}
	if node.Right != nil {
		levelPrint(node.Right, level + 1, maxLevel, m)
	}
}
