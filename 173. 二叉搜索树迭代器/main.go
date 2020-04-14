package main

import "fmt"

//实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
//
//调用 next() 将返回二叉搜索树中的下一个最小的数。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{Val:1}
	s := Constructor(head)
	fmt.Println(s)
}

type BSTIterator struct {
	s []int
	size int
}


func Constructor(root *TreeNode) BSTIterator {
	s := []int{}
	print(root, &s)
	return BSTIterator{
		s:    s,
		size: len(s),
	}
}

func print(node *TreeNode,s *[]int) {
	if node == nil {
		return
	}
	print(node.Left, s)
	*s = append(*s, node.Val)
	print(node.Right, s)
}


func (this *BSTIterator) Next() int {
	result := this.s[0]
	this.s = this.s[1:]
	this.size -= 1
	return result
}

func (this *BSTIterator) HasNext() bool {
	if this.size > 0 {
		return true
	}
	return false
}