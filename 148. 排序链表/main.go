package main

import "fmt"

//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//示例 1:
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2:
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sort-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{Val:5}
	head.Next = &ListNode{Val:2}
	head.Next.Next = &ListNode{Val:3}
	fmt.Println(sortList(&head))
}

//归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	middleNode := findMiddleNode(head)
	p1 := sortList(head)
	p2 := sortList(middleNode)
	//合并p1 p2
	dummyHead := &ListNode{Val:-1}
	p := dummyHead
	//返回结果
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummyHead.Next
}

//快慢指针遍历链表
func findMiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	//快慢指针遍历链表
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	middleNode := slow.Next
	//链表断开操作
	slow.Next = nil
	return middleNode
}
