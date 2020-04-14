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
	head := &ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3}}}
	result := swapPairs(head)
	fmt.Println(result)
	fmt.Println(result.Next)
	fmt.Println(result.Next.Next)
}

//1 2 3 4
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s := []int{}
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	length := len(s)
	r := []int{}
	if length % 2 == 1 {
		r = s[0:length-1]
	} else {
		r = s
	}
	for i := 0; i < len(r); i+=2 {
		r[i],r[i+1] = r[i+1],r[i]
	}
	newHead := &ListNode{Val:s[0]}
	node := newHead
	for _, value := range s[1:] {
		node.Next = &ListNode{Val:value}
		node = node.Next
	}
	return newHead
}
