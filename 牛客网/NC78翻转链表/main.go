package main

//输入一个链表，反转链表后，输出新链表的表头。

type ListNode struct{
	Val int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
//1->2->3
//next = 2
//1->nil
//prev=1
//pHead=2
func ReverseList( pHead *ListNode ) *ListNode {
	var prev *ListNode
	for pHead != nil {
		next := pHead.Next
		pHead.Next = prev
		prev = pHead
		pHead = next
	}
	return prev
}