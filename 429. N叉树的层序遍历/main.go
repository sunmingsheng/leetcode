package main

import "fmt"

//给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
//
//例如，给定一个 3叉树 :
//
//返回其层序遍历:
//
//[
//[1],
//[3,2,4],
//[5,6]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type Node struct {
	Val int
	Children []*Node
}

func main() {
	root := &Node{Val:1}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *Node) [][]int {
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
	for i := 1; i <= maxLevel; i++ {
		s = append(s, m[i])
	}
	return s
}

func print(node *Node, level int, m *map[int][]int) {
	if node == nil {
		return
	}
	if s, ok := (*m)[level]; ok {
		(*m)[level] = append(s, node.Val)
	} else {
		(*m)[level] = []int{node.Val}
	}
	for _, child := range node.Children {
		print(child, level + 1, m)
	}
}
