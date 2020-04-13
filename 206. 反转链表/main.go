package main

import "fmt"

//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val:1}
	head.Next = &ListNode{Val:2}
	fmt.Println(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	dummyNode := &ListNode{Val:-1}
	for head != nil {
		temp := head.Next
		head.Next = dummyNode.Next
		dummyNode.Next = head
		head = temp
	}
	return dummyNode.Next
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	node := head
//	s := []*ListNode{}
//	for node != nil {
//		s = append(s, node)
//		node = node.Next
//	}
//	length := len(s)
//	head = s[length - 1]
//	for i := length - 2; i >=0; i-- {
//		head.Next = s[i]
//	}
//	return head
//}
