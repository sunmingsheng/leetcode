package main

import "fmt"

//给定一个 N 叉树，找到其最大深度。
//
//最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
//例如，给定一个 3叉树 :

type Node struct {
	Val int
	Children []*Node
}

func main() {
	root := &Node{Val:1}
	fmt.Println(maxDepth(root))
}


func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxLen := 0
	path := []Node{*root}
	findPath(root,path,&maxLen)
	return maxLen
}

func findPath(node *Node, path []Node, maxLen *int) {
	if len(path) > *maxLen {
		*maxLen = len(path)
	}
	if node == nil {
		return
	} else {
		for _, value := range node.Children {
			if value != nil {
				findPath(value, append(path, *value), maxLen)
			}
		}
	}
}
