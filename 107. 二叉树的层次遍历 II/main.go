package main

import "fmt"

//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其自底向上的层次遍历为：
//
//[
//[15,7],
//[9,20],
//[3]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(levelOrderBottom(root))
}


func levelOrderBottom(root *TreeNode) [][]int {
	s := [][]int{}
	if root == nil {
		return s
	}
	m := make(map[int][]int)
	print(root, 1, &m)
	maxLevel := 1
	for key, _ := range m {
		if key > maxLevel {
			maxLevel = key
		}
	}
	for i := maxLevel; i >= 1; i-- {
		s = append(s, m[i])
	}
	return s
}

func print(node *TreeNode, level int, m *map[int][]int) {
	if node == nil {
		return
	}
	if s, ok := (*m)[level]; ok {
		(*m)[level] = append(s, node.Val)
	} else {
		(*m)[level] = []int{node.Val}
	}
	print(node.Left, level + 1, m)
	print(node.Right, level + 1, m)
}
