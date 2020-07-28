package main

import (
	"fmt"
)

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 
//
//示例：
//
//给你这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3}}}
	head = reverseKGroup(head, 3)
	fmt.Println("开始打印")
	printList(head)
}

func printList(node *ListNode) {
	for {
		if node != nil {
			println(node.Val)
			node = node.Next
		} else {
			break
		}
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	lists := []*ListNode{}
	for i := 0; i < k; i++ {
		if head == nil {
			return lists[0]
		}
		lists = append(lists, head)
		head = head.Next
	}
	nextA := lists[k - 1].Next
	for i := 0; i < k / 2; i++ {
		j := k - i - 1
		temp := lists[i]
		lists[i] = lists[j]
		lists[j] = temp
	}
	for i := 0; i < k - 1; i++ {
		lists[i].Next = lists[i+1]
	}
	lists[k - 1].Next = nextA
	if lists[k-1].Next != nil {
		lists[k-1].Next = reverseKGroup(lists[k-1].Next, k)
	}
	return lists[0]
}