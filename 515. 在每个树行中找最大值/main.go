package main

import "fmt"

//您需要在二叉树的每一行中找到最大的值。
//
//示例：
//
//输入:
//
//1
/// \
//3   2
/// \   \
//5   3   9
//
//输出: [1, 3, 9]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(largestValues(root))
}

func largestValues(root *TreeNode) []int {
	m := make(map[int]int)
	print(root, 0, &m)
	s := make([]int, len(m))
	for key, value := range m {
		s[key] = value
	}
	return s
}

func print(node *TreeNode, level int, m *map[int]int) {
	if node == nil {
		return
	}
	if max, ok := (*m)[level]; ok {
		if node.Val > max {
			(*m)[level] = node.Val
		}
	} else {
		(*m)[level] = node.Val
	}
	print(node.Left, level + 1, m)
	print(node.Right, level + 1, m)
}
