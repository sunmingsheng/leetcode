package main

import "fmt"

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表:  1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//
//给定的 n 保证是有效的。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode19 struct {
	Val int
	Next *ListNode19
}

func main() {
	head := &ListNode19{Val:1, Next:&ListNode19{Val:2,Next:&ListNode19{3, &ListNode19{Val:4, Next:&ListNode19{Val:5}}}}}
	head = removeNthFromEnd(head, 2)
	fmt.Println(head)
	fmt.Println(head.Next)
	fmt.Println(head.Next.Next)
	fmt.Println(head.Next.Next.Next)
}

//0-->1->2->3->4->5
func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	dummyHead := &ListNode19{}
	dummyHead.Next = head
	l1, l2 := dummyHead, dummyHead
	for i := 0; i <= n; i++ {
		l1 = l1.Next
	}
	for l1 != nil {
		l1 = l1.Next
		l2 = l2.Next
	}
	l2.Next = l2.Next.Next
	return dummyHead.Next
}
