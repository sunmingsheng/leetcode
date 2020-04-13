package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	node := Node{Val:1}
	fmt.Println(preorder(&node))
}

func preorder(root *Node) []int {
	s := []int{}
	print(root, &s)
	return s
}

func print(node *Node, s *[]int) {
	if node == nil {
		return
	}
	*s = append(*s, node.Val)
	for _, node := range node.Children {
		print(node, s)
	}
}
