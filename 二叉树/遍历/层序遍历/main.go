package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) Put(value *TreeNode) {
	*q = append(*q, value)
}

func (q *Queue) Pop() *TreeNode {
	if q.isEmpty() {
		return nil
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func main() {
	left := &TreeNode{Val:1}
	right := &TreeNode{Val:4,Left:&TreeNode{Val:3}}
	head := &TreeNode{Val:2}
	head.Left = left
	head.Right = right
	fmt.Println(levelOrder(head))
}

//func levelOrder(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	q := &Queue{}
//	q.Put(root)
//	//fmt.Println(root.Val)
//	for !q.isEmpty() {
//		node := q.Pop()
//		//fmt.Println(node.Val)
//		if node.Left != nil {
//			q.Put(node.Left)
//		}
//		if node.Right != nil {
//			q.Put(node.Right)
//		}
//	}
//}

func levelOrder( root *TreeNode ) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		temp := []int{}
		newQ := []*TreeNode{}
		for i := 0; i< len(q);i++ {
			temp = append(temp, q[i].Val)
			if q[i].Left != nil {
				newQ = append(newQ, q[i].Left)
			}
			if q[i].Right != nil {
				newQ = append(newQ, q[i].Right)
			}
		}
		q = newQ
		res = append(res, temp)
	}
	return res
}
