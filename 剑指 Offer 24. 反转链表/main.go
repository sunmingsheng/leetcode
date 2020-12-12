package main

import "fmt"

//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
// 
//
//限制：
//
//0 <= 节点个数 <= 5000
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	pHead := &ListNode{Val:1, Next: &ListNode{Val:2, Next:&ListNode{Val:3, Next:&ListNode{Val:4}}}}
	head := ReverseList(pHead)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// 4 3 2 1
func ReverseList( pHead *ListNode ) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	var prev *ListNode
	current := pHead
	next := pHead
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
