package main

import "fmt"

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
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
	head := &ListNode19{Val:1, Next:nil}
	fmt.Println(removeNthFromEnd(head, 1))
}

func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	oriHead := head
	s := []*ListNode19{}
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	if len(s) <= n {
		return oriHead.Next
	} else {
		s[len(s) - n - 1].Next = s[len(s) - n].Next
		s[len(s) - n].Next = nil
	}
	return oriHead
}
