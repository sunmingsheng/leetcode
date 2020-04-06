package main

import "fmt"

//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回锯齿形层次遍历如下：
//
//[
//[3],
//[20,9],
//[15,7]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	m := make(map[int]*[]int)
	levelPrint(root, m, 0, true)
	s := make([][]int, len(m))
	for key, items := range m {
		s[key] = *items
	}
	return s
}

func levelPrint(node *TreeNode, m map[int]*[]int, level int, flag bool) {
	if node == nil {
		return
	}
	if s, ok := m[level]; ok {
		*s = append(*s, node.Val)
	} else {
		s = &[]int{node.Val}
		m[level] = s
	}
	if flag {
		levelPrint(node.Left, m, level+1, false)
		levelPrint(node.Right, m, level+1, false)
	} else {
		levelPrint(node.Right, m, level+1, true)
		levelPrint(node.Left, m, level+1, true)
	}
}

//广度优先遍历和深度优先遍历