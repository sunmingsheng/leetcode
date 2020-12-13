package main

import "fmt"

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}


func main() {
	head := &ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3,Next:&ListNode{Val:4}}}}
	result := swapPairs(head)
	fmt.Println(result)
	fmt.Println(result.Next)
	fmt.Println(result.Next.Next)
	fmt.Println(result.Next.Next.Next)
}

//0 1 2 3 4
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	l := head
	m := head.Next
	if m == nil {
		return l
	}
	r := m.Next
	//交换l和m
	m.Next = l
	//递归处理
	l.Next = swapPairs(r)
	return m
}
