package main

import "fmt"

//给定一个二叉树，原地将它展开为链表。
//
//例如，给定二叉树
//
//1
/// \
//2   5
/// \   \
//3   4   6
//将其展开为：
//
//1
//\
//2
//\
//3
//\
//4
//\
//5
//\
//6
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{Val:1}
	flatten(head)
	fmt.Println(head)
}

func flatten(root *TreeNode)  {
	s := &[]*TreeNode{}
	print(root, s)
	for i := 0; i < len(*s); i++ {
		(*s)[i].Left = nil
		if i + 1 ==  len(*s) {
			(*s)[i].Right = nil
		} else {
			(*s)[i].Right = (*s)[i+1]
		}
	}
}

func print(node *TreeNode, s *[]*TreeNode) {
	if node == nil {
		return
	}
	*s = append(*s, node)
	print(node.Left, s)
	print(node.Right, s)
}
