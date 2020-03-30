package main

import "fmt"

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其层次遍历结果：
//
//[
//[3],
//[9,20],
//[15,7]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{Val:1}
	fmt.Println(levelOrder(head))
}

func levelOrder(root *TreeNode) [][]int {
	m := make(map[int]*[]int)
	levelPrint(root, m, 0)
	s := make([][]int, len(m))
	for key, items := range m {
		s[key] = *items
	}
	return s
}

func levelPrint(node *TreeNode, m map[int]*[]int, level int) {
	if node == nil {
		return
	}
	if s, ok := m[level]; ok {
		*s = append(*s, node.Val)
	} else {
		s = &[]int{node.Val}
		m[level] = s
	}
	levelPrint(node.Left, m, level + 1)
	levelPrint(node.Right, m, level + 1)
}
