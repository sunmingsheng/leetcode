package main

import "fmt"

//给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.
//
//示例 1:
//
//输入:
//3
/// \
//9  20
///  \
//15   7
//输出: [3, 14.5, 11]
//解释:
//第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{Val:1,Left:&TreeNode{Val:2},Right:&TreeNode{Val:3}}
	fmt.Println(averageOfLevels(head))
}

func averageOfLevels(root *TreeNode) []float64 {
	m := make(map[int]*[]int)
	levelPrint(root, m, 1)
	s := []float64{}
	maxLevel := 1
	for key, _ := range m {
		if key > maxLevel {
			maxLevel = key
		}
	}
	for i := 1; i <= maxLevel; i++ {
		result := 0
		for _, value := range *m[i] {
			result += value
		}
		s = append(s, float64(result) / float64(len(*m[i])))
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
