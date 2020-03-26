package main

import "fmt"

//给定一个 N 叉树，返回其节点值的后序遍历。
//
//例如，给定一个 3叉树 :
//
//返回其后序遍历: [5,6,3,2,4,1].

type Node struct {
	Val int
	Children []*Node
}

func main() {
	node := &Node{Val:0}
	fmt.Println(postorder(node))
}

func postorder(root *Node) []int {
 	s := &[]int{}
 	print(root, s)
 	return *s
}

func print(node *Node, s *[]int) {
  	if node == nil {
		return
	}
	for _, item := range node.Children {
		print(item, s)
	}
	*s = append(*s, node.Val)
}
