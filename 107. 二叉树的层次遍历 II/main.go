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
	traversal(root, 0, m)
	for i := len(m)-1; i >= 0; i-- {
		s = append(s, m[i])
	}
	return s
}

func traversal(node *TreeNode, level int, m map[int][]int) {
	if node == nil {
		return
	}
	if _, ok := m[level]; !ok {
		m[level] = []int{}
	}
	m[level] = append(m[level], node.Val)
	traversal(node.Left, level + 1, m)
	traversal(node.Right, level + 1, m)
}